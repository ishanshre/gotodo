package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/ishanshre/gotodo/pkg/models"
	"github.com/joho/godotenv"
)

type Storage interface {
	CreateToDo(*models.ToDo) error
	DeleteToDo(int) error
	UpdateToDo(int, *models.ToDo) error
	GetTodos() ([]*models.ToDo, error)
	GetToDoById(int) (*models.ToDo, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("Error in loading the environment files %v", err)
	}
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_CONN_STRING"))
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.createTodoTable()
}

func (s *PostgresStore) createTodoTable() error {
	query := `CREATE TABLE IF NOT EXISTS todo (
		id SERIAL PRIMARY KEY,
		body VARCHAR(500),
		created_at TIMESTAMPTZ
	)`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateToDo(todo *models.ToDo) error {
	query := `
		INSERT INTO todo (body, created_at)
		VALUES ($1, $2)
	`
	_, err := s.db.Query(query, todo.Body, todo.CreatedAt)
	if err != nil {
		return err
	}
	log.Println("Todo Created Successfully")
	return nil
}
func (s *PostgresStore) DeleteToDo(id int) error {
	query := `
		DELETE FROM todo
		WHERE id = $1;
	`
	s.db.Exec("COMMIT")
	_, err := s.db.Query(query, id)
	if err != nil {
		return err
	}
	return nil
}
func (s *PostgresStore) UpdateToDo(id int, updateTodo *models.ToDo) error {
	query := `
		UPDATE todo
		SET body = $1, created_at = $2
		WHERE id = $3
	`
	s.db.Exec("COMMIT")
	_, err := s.db.Query(query, updateTodo.Body, updateTodo.CreatedAt, id)
	if err != nil {
		return err
	}
	log.Printf("Todo with id %v updated successfully\n", id)
	return nil
}
func (s *PostgresStore) GetTodos() ([]*models.ToDo, error) {
	rows, err := s.db.Query("SELECT * FROM todo")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	todos := []*models.ToDo{}
	for rows.Next() {
		todo, err := scanTodos(rows)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}
func (s *PostgresStore) GetToDoById(id int) (*models.ToDo, error) {
	query := `SELECT * FROM todo WHERE id = $1`
	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanTodos(rows)
	}
	return nil, fmt.Errorf("todo %d not found", id)
}

func scanTodos(rows *sql.Rows) (*models.ToDo, error) {
	todo := new(models.ToDo)
	err := rows.Scan(&todo.Id, &todo.Body, &todo.CreatedAt)
	return todo, err
}
