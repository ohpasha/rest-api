package main

import (
	"log"

	_ "github.com/lib/pq"
	todo "github.com/ohpasha/rest-api"
	"github.com/ohpasha/rest-api/pkg/handler"
	"github.com/ohpasha/rest-api/pkg/repository"
	"github.com/ohpasha/rest-api/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfigs(); err != nil {
		log.Fatalf("can't read config file: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("cann't inizialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)
	srv := new(todo.Server)

	if error := srv.Run(viper.GetString("port"), handlers.InitRouters()); error != nil {
		log.Fatalf("error: %s", error.Error())
	}
}

func initConfigs() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
