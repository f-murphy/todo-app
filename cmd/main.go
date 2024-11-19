package main

import (
	"fmt"
	"log"
	"todo-app/configs"
	"todo-app/handler"
	"todo-app/repository"
	"todo-app/service"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := configs.InitConfig()
	if err != nil {
		fmt.Println("error initializing configs")
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		Username: cfg.DB.Username,
		DBName:   cfg.DB.Username,
		SSLMode:  cfg.DB.SSLMode,
		Password: cfg.DB.Password,
	})
	if err != nil {
		log.Fatal("error initializing db")
	}
	defer db.Close()
	fmt.Print("Database connected successfully")

	repository := repository.NewTodoRepository(db)
	service := service.NewTodoService(repository)
	handler := handler.NewTodoHandler(service)

	r := gin.Default()
	r.POST("/todo", handler.CreateTodo)
	r.GET("/todo", handler.GetTodos)
	r.PUT("/todo/:id", handler.UpdateTodo)
	r.DELETE("/todo/:id", handler.DeleteTodo)

	if err := r.Run(":8080"); err != nil {
		fmt.Println("Failed to run server: ", err)
	}
}
