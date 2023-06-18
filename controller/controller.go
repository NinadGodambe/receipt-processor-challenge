package controller

import (
	"receipt-processor-challenge/model"
	"receipt-processor-challenge/util"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

// Defines a struct with a map to hold the receipts
type ReceiptController struct {
	receipts map[string]model.Receipt
}

// NewReceiptController creates a new instance of the ReceiptController
func NewReceiptController() *ReceiptController {
	return &ReceiptController{
		receipts: make(map[string]model.Receipt),
	}
}

// ProcessReceipt handles the request to process a receipt
func (rc *ReceiptController) ProcessReceipt(c *gin.Context) {
	var receipt model.Receipt
	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generating a unique ID for the receipt
	receiptID := util.GenerateUUID()
	
	// Storing the receipt in the map with the generated ID
	rc.receipts[receiptID] = receipt
	c.JSON(http.StatusOK, gin.H{"id": receiptID})
}

// GetPoints handles retrieving the points earned for a given receipt
func (rc *ReceiptController) GetPoints(c *gin.Context) {
	receiptID := c.Param("id")

	// Retrieving the receipt from the map based on the given ID
	receipt, exists := rc.receipts[receiptID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Receipt not found"})
		return
	}

	// Calculating the points for the receipt using the CalculatePoints function
	points := util.CalculatePoints(&receipt)
	fmt.Println("Points:", points)
	c.JSON(http.StatusOK, gin.H{"points": points})
}

