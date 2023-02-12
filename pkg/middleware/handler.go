package middleware

import (
	"fmt"
	"net/http"

	"github.com/ishanshre/gotodo/models"
)

func (s *APIServer) handleToDos(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetToDos(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetToDos(w http.ResponseWriter, r *http.Request) error {
	todo := models.NewToDo("Learn Golang")
	return writeJSON(w, http.StatusOK, todo)
}

func (s *APIServer) handleGetToDoById(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleCreateToDo(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleUpdateToDo(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteToDo(w http.ResponseWriter, r *http.Request) error {
	return nil
}
