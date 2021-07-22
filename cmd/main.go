package main

import (
	"log"

	"github.com/forever-eight/todo.git"
	"github.com/forever-eight/todo.git/internal/app/handler"
	"github.com/forever-eight/todo.git/internal/app/repository"
	"github.com/forever-eight/todo.git/internal/app/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	err := srv.Run("8080", handlers.InitRoutes())
	if err != nil {
		log.Println("error occurred while running http server ", err)
	}
}
