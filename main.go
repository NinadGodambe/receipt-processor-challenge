package main

import (
	"receipt-processor-challenge/controller"
	"log"
	"github.com/gin-gonic/gin"
)

func main() {
	// Creating a router instance
	router := gin.Default()

	// Initializing the controller
	receiptController := controller.NewReceiptController()

	// Defining the routes and starting the server
	router.POST("/receipts/process", receiptController.ProcessReceipt)
	router.GET("/receipts/:id/points", receiptController.GetPoints)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
