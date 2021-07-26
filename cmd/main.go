package main

import (
	"log"

	"github.com/spf13/viper"

	"github.com/forever-eight/todo.git"
	"github.com/forever-eight/todo.git/internal/app/handler"
	"github.com/forever-eight/todo.git/internal/app/repository"
	"github.com/forever-eight/todo.git/internal/app/service"
)

func main() {
	err := initConfig()
	if err != nil {
		log.Fatalf("config initialization error: %s", err)
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	err = srv.Run(viper.GetString("port"), handlers.InitRoutes())
	if err != nil {
		log.Println("error occurred while running http server ", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
