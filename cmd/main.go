package main

import (
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	todo "github.com/ohpasha/rest-api"
	"github.com/ohpasha/rest-api/pkg/handler"
	"github.com/ohpasha/rest-api/pkg/repository"
	"github.com/ohpasha/rest-api/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

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

	if error := srv.Run(viper.GetString("port"), handlers.InitRouters()); error != nil {
		logrus.Fatalf("error: %s", error.Error())
	}
}

func initConfigs() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
