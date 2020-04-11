package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Response struct {
	FindItemsByKeywordsResponse FindItemsByKeywordsResponse
}

type FindItemsByKeywordsResponse struct {
	SearchResult SearchResult
}

type SearchResult struct {
	Item []Item
}

type Item struct {
	Title           []string `json:"title"`
	GlobalID        []string `json:"globalId"`
	Subtitle        []string `json:"subtitle"`
	PrimaryCategory Category
	Gallery         []string `json:"galleryURL"`
	ItemURL         []string `json:"viewItemURL"`
	PaymentMethod   []string `json:"paymentMethod"`
	Location        []string `json:"location"`
	Country         []string `json:"country"`
	ShipToLocation  []string `json:"shipToLocation"`
	ShippingType    []string `json:"shippingType"`
	Condition       Condition
	ShipCost        ItemPrice
	Price           ItemPrice
}

type Category struct {
	CategoryID   []string `json:"categoryId"`
	CategoryName []string `json:"categoryName"`
}

type Condition struct {
	ConditionID          []string `json:"conditionId"`
	ConditionDisplayName []string `json:"conditionDisplayName"`
}

type ItemResponse struct {
	Title           []string
	GlobalID        []string
	Subtitle        []string
	PrimaryCategory Category
	Gallery         []string
	ItemURL         []string
	PaymentMethod   []string
	Location        []string
	Country         []string
	ShipToLocation  []string
	ShippingType    []string
	Condition       Condition
	ShipCost        ItemPrice
	Price           ItemPrice
}

type ItemPrice struct {
	Currency string `json:"@currencyId"`
	Value    string `json:"__value__"`
}

type ToDoList struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Task   string             `json:"task,omitempty"`
	Status bool               `json:"status,omitempty"`
}
