package main

import (
	"todo-app/configs"
	"todo-app/handler"
	"todo-app/repository"
	"todo-app/service"
	"todo-app/utils/logger"

	_ "github.com/lib/pq"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	logFile, err := logger.InitLogger()
	if err != nil {
		logrus.WithError(err).Fatal("Error loading logrus")
	}
	logrus.Info("logFile initialized successfully")
	defer logFile.Close()

	cfg, err := configs.InitConfig()
	if err != nil {
		logrus.WithError(err).Fatal("error initializing configs")
	}
	logrus.Info("Configs initialized successfully")

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		Username: cfg.DB.Username,
		DBName:   cfg.DB.Username,
		SSLMode:  cfg.DB.SSLMode,
		Password: cfg.DB.Password,
	})
	if err != nil {
		logrus.WithError(err).Fatal("error initializing db")
	}
	logrus.Info("DB connection successfully")

	repository := repository.NewTodoRepository(db)
	service := service.NewTodoService(repository)
	handler := handler.NewTodoHandler(service)

	r := gin.Default()
	r.POST("/todo", handler.CreateTodo)
	r.GET("/tasks", handler.GetTodos)
	r.PUT("/task/:id", handler.UpdateTodo)
	r.DELETE("/task/:id", handler.DeleteTodo)

	if err := r.Run(":8080"); err != nil {
		logrus.WithError(err).Fatal("Failed to run server: ")
	}
}
