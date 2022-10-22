package main

import (
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
	router.GET("/todos", getAPIContent)
	router.Run("localhost:9090")
}

func getAPIContent(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}
