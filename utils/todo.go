package utils

import (
	"errors"
)

type Todo struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Complete bool   `json:"complete"`
}

var Todos []Todo

func CreateTodo(title string) Todo {
	uniqueId := generateUniqueID()

	var newTodo = Todo{ID: uniqueId, Title: title, Complete: false}

	Todos = append(Todos, newTodo)

	return newTodo
}

func DeleteTodo(id string) {
	for i, todo := range Todos {
		if todo.ID == id {
			Todos = append(Todos[:i], Todos[i+1:]...)
		}
	}
}

func GetAllTodos() []Todo {
	return Todos
}

func GetTodoByID(id string) (Todo, error) {
	var expectedTodo Todo

	for _, todo := range Todos {
		if todo.ID == id {
			expectedTodo = todo
			return expectedTodo, nil
		}
	}

	return expectedTodo, errors.New("unable to find todo with id: " + id)
}

func ToggleCompleteStatus(id string) error {
	for i, todo := range Todos {
		if todo.ID == id {
			Todos[i].Complete = !todo.Complete
			return nil
		}
	}
	return errors.New("Unable to find todo with id: " + id)
}

func UpdateTodos(id string, updatedTitle string) (Todo, error) {

	var updateTodo Todo

	for i, todo := range Todos {
		if todo.ID == id {
			updateTodo = Todos[i]
			updateTodo.Title = updatedTitle

			return todo, nil
		}
	}
	return updateTodo, errors.New("Unable to find todo with id: " + id)
}

// populate the Todos onload
func init() {
	for _, taskName := range TodoTaksList {
		CreateTodo(taskName)
	}
}
