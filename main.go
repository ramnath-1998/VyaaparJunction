package main

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/ramnath.1998/vyaaparjunction/db"
	"github.com/ramnath.1998/vyaaparjunction/ent"
	"github.com/ramnath.1998/vyaaparjunction/routes"
)

func CreateProduct(client *ent.Client) (*ent.ProductCategory, error) {

	cement, err :=
		client.ProductCategory.Create().
			SetCategoryName("Cement Category").
			SetIdentifier(uuid.New()).
			Save(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed creating cement: %w", err)
	}

	return cement, nil
}

func init() {
	db.DBConnect()
	routes.RunRoutes()
}

func main() {
	category, err := CreateProduct(&db.DbClient)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(category)
	defer db.DbClient.Close()
}
