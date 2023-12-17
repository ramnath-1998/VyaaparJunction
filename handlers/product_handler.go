package handlers

import (
	"context"
	"log"

	"github.com/ramnath.1998/vyaaparjunction/db"
	"github.com/ramnath.1998/vyaaparjunction/models"
)

func HandleGetProducts() []models.Product {

	products, err := db.DbClient.Product.Query().All(context.Background())
	if err != nil {
		log.Fatalf("failed querying products: %v", err)
	}

	result := []models.Product{}
	for _, eachProduct := range products {
		productModel := models.Product{
			ProductName: eachProduct.ProductName,
			Identifier:  eachProduct.Identifier.String(),
			CreatedOn:   eachProduct.CreatedOn.Format("2006-01-02T15:04:05"),
		}
		result = append(result, productModel)
	}

	return result
}
