package handler

import (
	"net/http"
	"strconv"
	"todo-app/models"
	"todo-app/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type TodoHandler struct {
	service service.TodosServiceInterface
}

func NewTodoHandler(service service.TodosServiceInterface) *TodoHandler {
	return &TodoHandler{service: service}
}

func (h *TodoHandler) GetTodos(c *gin.Context) {
	todos, err := h.service.FindAll()

	if err != nil {
		logrus.WithError(err).Error("error getting all todos")
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	logrus.Info("All todos successfully received")
	c.JSON(http.StatusOK, todos)
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var todo models.Todo
	err := c.BindJSON(&todo)
	if err != nil {
		logrus.WithError(err).Error("error binding JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.service.Create(todo)
	if err != nil {
		logrus.WithError(err).Error("error creating task")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logrus.Info("Create todo success")
	c.JSON(http.StatusCreated, gin.H{"todo id": id})
}

func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")
	todoID, err := strconv.Atoi(id)
	if err != nil {
		logrus.WithError(err).Error("Invalid ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo.ID = todoID
	updatedTodo, err := h.service.Update(todo)
	if err != nil {
		logrus.WithError(err).Error("Error updating todo")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logrus.Info("Update todo success")
	c.JSON(http.StatusOK, updatedTodo)
}

func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	todoID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.Delete(uint(todoID)); err != nil {
		logrus.WithError(err).Error("Error deleting todo")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logrus.Info("Delete todo success")
	c.JSON(http.StatusNoContent, nil)
}
