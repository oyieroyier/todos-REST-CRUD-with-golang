package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []Todo{
	{ID: "1", Title: "Learn some Go", Completed: true},
	{ID: "2", Title: "Build a REST API in Go", Completed: true},
	{ID: "3", Title: "Deploy the REST API", Completed: false},
}

// Get all todos
func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

// Filter todo by ID
func filterTodoById(id string) (*Todo, error) {
	for index, value := range todos {
		if value.ID == id {
			return &todos[index], nil
		}
	}
	// If it does not find a Todo with the ID, it will return nil and an error.
	return nil, errors.New("selected todo not found")
}

// Get the single Todo by id.
func getTodoById(context *gin.Context) {
	id := context.Param("id")

	todo, err := filterTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)

}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodoById)

	err := router.Run(":8080")

	if err != nil {
		return
	}
}
