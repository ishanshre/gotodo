package storage

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/ishanshre/gotodo/models"
	"github.com/joho/godotenv"
)

type Storage interface {
	CreateToDo(*models.ToDo) error
	DeleteToDo(int) error
	UpdateToDo(*models.ToDo) error
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

func (s *PostgresStore) CreateToDo(*models.ToDo) error {
	return nil
}
func (s *PostgresStore) DeleteToDo(int) error {
	return nil
}
func (s *PostgresStore) UpdateToDo(*models.ToDo) error {
	return nil
}
func (s *PostgresStore) GetTodos() ([]*models.ToDo, error) {
	return nil, nil
}
func (s *PostgresStore) GetToDoById(int) (*models.ToDo, error) {
	return nil, nil
}
