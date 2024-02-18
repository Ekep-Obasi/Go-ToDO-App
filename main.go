package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Ekep-Obasi/utils"
)

func todoController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	requestMethod := r.Method
	queryParam := r.URL.Query().Get("id")

	// for GET /todo & /todo?id={id} request
	if requestMethod == "GET" {

		// for request with query id
		if queryParam != "" {

			todo, err := utils.GetTodoByID(queryParam)

			// invalid query param
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}

			todoJSON, err := utils.ToJSON(todo)

			// error during marshaling
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Write(todoJSON)
			return
		}

		// for request without query id

		allTodosJSON, err := utils.ToJSON(utils.GetAllTodos())

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(allTodosJSON)
		return
	} else if requestMethod == "POST" {

		body, err := io.ReadAll(r.Body)

		if err != nil {
			defer r.Body.Close()

			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var todo utils.Todo

		error := utils.ToSTRUCT(body, &todo)

		if error != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		newTodoJSON, err := utils.ToJSON(utils.CreateTodo(todo.ID))

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write(newTodoJSON)
		return
	}
}

func main() {

	var cwd, _ = os.Getwd()
	var staticFilePath = cwd + "/static"

	// defining route handlers
	http.Handle("/", http.FileServer(http.Dir(staticFilePath))) // <-- static route
	http.HandleFunc("/todos", todoController)

	fmt.Println("<--Server is running on Port 5000-->")

	http.ListenAndServe(":5000", nil)
}
