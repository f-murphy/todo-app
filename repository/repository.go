package repository

import (
	"time"
	"todo-app/models"

	"github.com/jmoiron/sqlx"
)

type TodoRepositoryInterface interface {
	findAll() ([]*models.Todo, error)
	Create(todo models.Todo) (models.Todo, error)
	Update(todo models.Todo) (models.Todo, error)
	Delete(id uint) error
}

type TodoRepository struct {
	db *sqlx.DB
}

func NewTodoRepository(db *sqlx.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) findAll() ([]*models.Todo, error) {
	todos := []*models.Todo{}
	query := `SELECT * from Todos`

	err := r.db.Select(&todos, query)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *TodoRepository) Create(todo models.Todo) (models.Todo, error) {
	var id uint
	now := time.Now()
	err := r.db.QueryRow("INSERT INTO Todos (title, content, completed, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		todo.Title, todo.Content, todo.Completed, now, now).Scan(&id)
	if err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}

func (r *TodoRepository) Update(todo models.Todo) (models.Todo, error) {
	now := time.Now()
	_, err := r.db.Exec("UPDATE Todos SET title = $1, content = $2, completed = $3, updated_at = $4 WHERE id = $5",
		todo.Title, todo.Content, todo.Completed, now, todo.ID)
	if err != nil {
		return models.Todo{}, err
	}
	todo.UpdatedAt = now
	return todo, nil
}

func (r *TodoRepository) Delete(id uint) error {
	_, err := r.db.Exec("DELETE FROM Todos WHERE id = $1", id)
	return err
}
