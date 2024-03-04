package main

import (
	"barTrApp/pkg/handler"
	"barTrApp/pkg/repository"
	baristaTrainingApp "barTrApp/pkg/server"
	"barTrApp/pkg/service"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load("../../.env"); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
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
		logrus.Fatalf("failed to initialize db %s", err.Error())
	}

	repository := repository.NewRepository(db)
	services := service.NewService(repository)
	handlers := handler.NewHandler(services)

	server := new(baristaTrainingApp.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error ocured while running http server")
	}
}

func initConfig() error {
	viper.AddConfigPath("../../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

//swag for documentaion swagger auto gen
//swag init -g internal/app/app.go
