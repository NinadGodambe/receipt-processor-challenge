package model


// Create a new Receipt object
type Receipt struct {
	Retailer      string  `json:"retailer"`
	PurchaseDate  string  `json:"purchaseDate"`
	PurchaseTime  string  `json:"purchaseTime"`
	Items         []Item  `json:"items"`
	Total         string  `json:"total"`
}

// Create a new Item object
type Item struct {
	ShortDescription string  `json:"shortDescription"`
	Price            string  `json:"price"`
}
