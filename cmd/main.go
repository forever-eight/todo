package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/forever-eight/todo.git"
	"github.com/forever-eight/todo.git/internal/app/handler"
	"github.com/forever-eight/todo.git/internal/app/repository"
	"github.com/forever-eight/todo.git/internal/app/service"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	err := initConfig()
	if err != nil {
		logrus.Fatalf("config initialization error: %s", err)
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("env file error: %s", err)
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
		logrus.Fatalf("db initialization error: %s", err)
	}

	repos := repository.NewRepository(db)
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
