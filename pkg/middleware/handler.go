package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ishanshre/gotodo/pkg/models"
)

func (s *APIServer) handleToDos(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetToDos(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateToDo(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleToDosById(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetToDoById(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteToDo(w, r)
	}
	if r.Method == "PUT" {
		return s.handleUpdateToDo(w, r)
	}
	return fmt.Errorf("method not allowed: %s", r.Method)
}

func (s *APIServer) handleGetToDos(w http.ResponseWriter, r *http.Request) error {
	todos, err := s.store.GetTodos()
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, todos)
}

func (s *APIServer) handleGetToDoById(w http.ResponseWriter, r *http.Request) error {
	id, err := getId(r)
	if err != nil {
		return err
	}
	todo, err := s.store.GetToDoById(id)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, todo)
}

func (s *APIServer) handleCreateToDo(w http.ResponseWriter, r *http.Request) error {
	createNewTodo := new(models.ToDo) // create  a new varaible to store todo struct
	if err := json.NewDecoder(r.Body).Decode(&createNewTodo); err != nil {
		return nil
	} // decode user input into createTodo varialbe
	todos := models.NewToDo(createNewTodo.Body)
	if err := s.store.CreateToDo(todos); err != nil {
		return err
	}
	return writeJSON(w, http.StatusCreated, todos)
}

func (s *APIServer) handleUpdateToDo(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteToDo(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func getId(r *http.Request) (int, error) {
	idstr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return id, fmt.Errorf("id parsing error: %s", err)
	}
	return id, nil
}
