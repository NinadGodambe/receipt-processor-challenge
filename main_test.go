package main_test

import (
	"receipt-processor-challenge/controller"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
)

func TestMain(t *testing.T) {
	// Creating a new default router
	router := gin.Default()

	// Initializing the controller
	receiptController := controller.NewReceiptController()

	// Defining the routes and creating a POST request
	router.POST("/receipts/process", receiptController.ProcessReceipt)
	router.GET("/receipts/:id/points", receiptController.GetPoints)
	req, err := http.NewRequest("POST", "/receipts/process", nil)
	assert.NoError(t, err)

	// Recording the response
	res := httptest.NewRecorder()

	// Sending the request and checking the Http Status Code
	router.ServeHTTP(res, req)
	assert.Equal(t, http.StatusBadRequest, res.Code)
}

