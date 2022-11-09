package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	todo "github.com/ohpasha/rest-api"
	"github.com/ohpasha/rest-api/pkg/handler"
	"github.com/ohpasha/rest-api/pkg/repository"
	"github.com/ohpasha/rest-api/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title           Todo REST Api
// @version         1.0
// @description     server api for Todo list

// @contact.name   @uhpasha
// @contact.email  pashok.yoba@gmail.com

// @host      localhost:8000
// @BasePath  /

// @securityDefinitions.apikey  ApiKeyAuth
// @in header
// @name Authorization
func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfigs(); err != nil {
		logrus.Fatalf("can't read config file: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("can't load .env: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("cann't inizialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)
	srv := new(todo.Server)

	go func() {
		if error := srv.Run(viper.GetString("port"), handlers.InitRouters()); error != nil {
			logrus.Fatalf("error: %s", error.Error())
		}
	}()

	logrus.Print("server started")

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	logrus.Print("server shutting down")

	if err := srv.Stop(context.Background()); err != nil {
		logrus.Error("error with stopping server; %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Error("error with stopping db; %s", err.Error())
	}

}

func initConfigs() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
