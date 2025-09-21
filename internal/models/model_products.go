package models

type Products struct{

	ProductId string `form:"primaryKey" json:"productId"`
	Name string `json: "name"`
	Price string `json: "price"`
	VendorId string `json: "vendorId"`
}