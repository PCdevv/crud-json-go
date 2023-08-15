package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"errors"

	"strconv"
)

type todo struct {
	ID        int    `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: 1, Item: "Clean Room", Completed: false},
	{ID: 2, Item: "Read Book", Completed: false},
	{ID: 3, Item: "Record Video Room", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func getTodoById(id int) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("Todo not found")
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	numId, err := strconv.Atoi(id)
	todo, err := getTodoById(numId)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func toggleTodoSatus(context *gin.Context) {
	id := context.Param("id")
	numId, err := strconv.Atoi(id)
	todo, err := getTodoById(numId)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}

func addTodos(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

func deleteTodo(context *gin.Context) {
	id := context.Param("id")
	numId, err := strconv.Atoi(id)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid todo ID"})
		return
	}

	for i, t := range todos {
		if t.ID == numId {
			todos = append(todos[:i], todos[i+1:]...)
			context.IndentedJSON(http.StatusOK, gin.H{"message": "Todo deleted"})
			return
		}
	}

	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.POST("/todos", addTodos)
	router.PATCH("/todos/:id", toggleTodoSatus)
	router.DELETE("/todos/:id", deleteTodo)
	router.Run("localhost:9090")
}
