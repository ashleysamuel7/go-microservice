package database

import (
	"context"

	"github.com/ashleysamuel7/go-microservice/internal/models"
)

func (c Client) GetAllVendors(ctx context.Context) ([]models.Products, error){
	var products [] models.Products
	result:= c.DB.WithContext(ctx).
		Find(&products)
	return products, result.Error
}