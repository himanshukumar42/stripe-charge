package models

import "github.com/jinzhu/gorm"

type Address struct {
	City       string `json:"city"`
	Country    string `json:"country"`
	Line1      string `json:"line1"`
	Line2      string `json:"line2"`
	PostalCode string `json:"postal_code"`
	State      string `json:"state"`
}

type ShippingDetails struct {
	Address        Address `json:"address"`
	Name           string  `json:"name"`
	Phone          string  `json:"phone"`
	TrackingNumber string  `json:"tracking_number"`
}
type Charge struct {
	gorm.Model
	Amount       int64           `json:"amount"`
	ReceiptEmail string          `json:"receiptMail"`
	ProductName  string          `json:"productName"`
	Shipping     ShippingDetails `json:"shipping"`
}

func (c *Charge) TableName() string {
	return "charge"

}
