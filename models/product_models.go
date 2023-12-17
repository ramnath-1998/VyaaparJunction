package models

type Product struct {
	ProductName string `json:"productName"`
	Identifier  string `json:"identifier"`
	CreatedOn   string `json:"createdOn,omitempty"`
}
