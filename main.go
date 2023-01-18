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

func addCustomer(context *gin.Context) {
	var newCustomer customer
	if err := context.BindJSON(&newCustomer); err != nil {
		return
	}

	customers = append(customers, newCustomer)

	context.IndentedJSON(http.StatusCreated, newCustomer)

}

func main() {
	router := gin.Default()
	router.GET("/customers", getCustomers)
	router.POST("/addcustomer", addCustomer)
	router.Run("localhost:9090")
}
