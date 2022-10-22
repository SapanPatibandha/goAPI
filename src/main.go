package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{Item: "Learn C#", ID: "3", Completed: true},
	{Item: "Learn GO", ID: "1", Completed: false},
	{Item: "Learn AWS", ID: "2", Completed: true},
}

func main() {

	router := gin.Default()
	router.GET("/todos", getTODOS)
	router.GET("/todos/:id", getTODO)
	router.PATCH("/todos/:id", patchTODO)
	router.POST("/todos", addTODO)
	router.Run("localhost:9090")
}

func getTODOS(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTODO(context *gin.Context) {
	var newToDo todo

	err := context.BindJSON(&newToDo)

	if err != nil {
		return
	}

	todos = append(todos, newToDo)

	context.IndentedJSON(http.StatusCreated, newToDo)
}

func getTODObyID(id string) (*todo, error) {
	for i, val := range todos {
		if val.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("Requested ID not found")
}

func getTODO(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTODObyID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message:": "Record not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func patchTODO(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTODObyID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message:": "Record not found"})
		return
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}
