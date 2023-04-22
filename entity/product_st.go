package entity

type ProductSt struct {
	ID          string `json:"id"`
	ProductName string `json:"product_name"`
	Price       int    `json:"price"`
	ProductCtg  string `json:"product_ctg"`
	Stock       int    `json:"stock"`
}
