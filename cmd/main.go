package main

import (
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	clinic "sycret-test-task"
	"sycret-test-task/internal/handler"
	"sycret-test-task/internal/service"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("Error occured during config initializing: %s", err.Error())
	}

	services := service.NewServices()
	handlers := handler.NewHandlers(services)

	srv := new(clinic.Server)
	if err := srv.Run(viper.GetString("PORT"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server")
	}
}

func initConfig() error {
	viper.SetConfigFile("C:/test/сonfig/.env")
	return viper.ReadInConfig()
}
