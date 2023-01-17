package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "net/http"
)

type customer struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var customers = []customer{
	{ID: "1", Item: "Customer 1", Completed: false},
	{ID: "2", Item: "Customer 2", Completed: false},
	{ID: "3", Item: "Customer 3", Completed: true},
}

func getCustomers(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, customers)
}

func main() {
	router := gin.Default()
	router.GET("/customers", getCustomers)
	router.Run("localhost:9090")
}
