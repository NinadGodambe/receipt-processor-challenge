package controller_test

import (
	"receipt-processor-challenge/controller"
	"receipt-processor-challenge/model"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
)

// TestProcessReceipt tests the ProcessReceipt function.
func TestProcessReceipt(t *testing.T) {
	// Creating a new controller instance
	ctrl := controller.NewReceiptController()

	// Creating a receipt for request
	receipt := model.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []model.Item{
			{
				ShortDescription: "Mountain Dew 12PK",
				Price:            "6.49",
			},
			{
				ShortDescription: "Emils Cheese Pizza",
				Price:            "12.25",
			},
		},
		Total: "18.74",
	}
	payload, err := json.Marshal(receipt)
	assert.NoError(t, err)

	// Creating a new POST request and capturing the response
	req, err := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(payload))
	assert.NoError(t, err)
	res := httptest.NewRecorder()

	// Creating a router, sending the request and cheching the status code
	router := gin.Default()
	router.POST("/receipts/process", ctrl.ProcessReceipt)
	router.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)

}