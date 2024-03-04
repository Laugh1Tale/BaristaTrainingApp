package main

import (
	"barTrApp/pkg/handler"
	"barTrApp/pkg/repository"
	baristaTrainingApp "barTrApp/pkg/server"
	"barTrApp/pkg/service"
	"log"

	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repository := repository.NewRepository()
	services := service.NewService(repository)
	handlers := handler.NewHandler(services)

	server := new(baristaTrainingApp.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error ocured while running http server")
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

//swag for documentaion swagger auto gen
//swag init -g internal/app/app.go
