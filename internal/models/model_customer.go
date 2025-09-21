package models

type Customer struct{
	CustomerId string `form:"primaryKey" json:"customerId"`
	FirstName string `json: "firstName"`
	LastName string `json: "lastName"`
	Email string `json: "email"`
	Phone string `json: "phone"`
	Address string `json: "address"`
}