package main

import (
	"errors"
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

func toggleCustomerId(context *gin.Context) {
	id := context.Param("id")
	customer, err := getCustomerById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Customer not Found"})
	}

	customer.Completed = !customer.Completed

	context.IndentedJSON(http.StatusOK, customer)

}

func getCustomerById(id string) (*customer, error) {
	for i, t := range customers {
		if t.ID == id {
			return &customers[i], nil
		}
	}
	return nil, errors.New("Customer Not Found")
}

func getCustomer(context *gin.Context) {
	id := context.Param("id")
	customer, err := getCustomerById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Customer not Found"})
	}

	context.IndentedJSON(http.StatusOK, customer)
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
	router.GET("/customers/:id", getCustomer)
	router.PATCH("/customers/:id", toggleCustomerId)
	router.POST("/addcustomer", addCustomer)
	router.Run("localhost:9090")
}
