package database

import (
	"context"

	"github.com/ashleysamuel7/go-microservice/internal/models"
)

func (c Client) GetAllProducts(ctx context.Context, vendor_id string) ([]models.Products, error){
	var products [] models.Products
	result:= c.DB.WithContext(ctx).
		Where(models.Products{VendorId: vendor_id}).
		Find(&products)
	return products, result.Error
}