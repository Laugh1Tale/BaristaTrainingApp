package main

import (
	"barTrApp/pkg/handler"
	baristaTrainingApp "barTrApp/pkg/server"
	"log"
)

func main() {
	handler := new(handler.Handler)
	server := new(baristaTrainingApp.Server)
	if err := server.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("error ocured while running http server")
	}
}

//swag for documentaion swagger auto gen
//swag init -g internal/app/app.go
