package repository

import (
	"database/sql"
	"fmt"
	"log"

	"go_inven_ctrl/entity"
)

type ProductStRepo interface {
	GetAllProductSt() any
	GetByIdProductSt(id string) any
	CreateProductSt(newProductSt *entity.ProductSt) string
	UpdateProductSt(productSt *entity.ProductSt) string
	DeleteProductSt(id string) string
}

type productStRepo struct {
	db *sql.DB
}

func (r *productStRepo) GetAllProductSt() any {
	fmt.Println("test")
	var productsSt []entity.ProductSt

	query := "SELECT  id,product_name,price,product_category,stock FROM product_st order by id asc"
	rows, err := r.db.Query(query)

	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var productSt entity.ProductSt

		if err := rows.Scan(&productSt.ID, &productSt.ProductName, &productSt.Price, &productSt.ProductCtg, &productSt.Stock); err != nil {
			log.Println(err)
		}

		productsSt = append(productsSt, productSt)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}
	fmt.Println(productsSt)
	if len(productsSt) == 0 {
		return "no data"
	}

	return productsSt

}

func (r *productStRepo) GetByIdProductSt(id string) any {
	var productSt entity.ProductSt

	query := "SELECT id, product_name, price,product_category ,stock FROM product_st WHERE id = $1"

	row := r.db.QueryRow(query, id)

	if err := row.Scan(&productSt.ID, &productSt.ProductName, &productSt.Price, &productSt.ProductCtg, &productSt.Stock); err != nil {
		log.Println(err)
	}

	if productSt.ID == "" {
		return "productSt not found"
	}

	return productSt

}

func (r *productStRepo) CreateProductSt(newProductSt *entity.ProductSt) string {
	query := "INSERT INTO product_st ( id, product_name, product_category, price) VALUES($1,$2,$3,$4)"
	_, err := r.db.Exec(query, newProductSt.ID, newProductSt.ProductName, newProductSt.ProductCtg, newProductSt.Price)

	if err != nil {
		log.Println(err)
		return "failed to create ProductSt"
	}

	return "productSt created successfully"
}

func (r *productStRepo) UpdateProductSt(productSt *entity.ProductSt) string {
	res := r.GetByIdProductSt(productSt.ID)
	fmt.Println(productSt)
	if res == "productSt not found" {
		return res.(string)
	}

	query := "UPDATE product_st SET product_name = $1, price = $2,product_category = $3 WHERE id = $4 ;"
	_, err := r.db.Exec(query, productSt.ProductName, productSt.Price, productSt.ProductCtg, productSt.ID)

	if err != nil {
		log.Println(err)
		return "failed to update ProductSt"
	}

	return fmt.Sprintf("ProductSt with id %s updated successfully", productSt.ID)

}

func (r *productStRepo) DeleteProductSt(id string) string {
	res := r.GetByIdProductSt(id)
	if res == "productSt not found" {
		return res.(string)
	}

	query := "DELETE FROM product_st WHERE id =$1"
	_, err := r.db.Exec(query, id)

	if err != nil {
		log.Println(err)
		return "failed to delete productSt"
	}

	return fmt.Sprintf("ProductSt with id %s deleted successfully", id)
}

func NewProductStRepo(db *sql.DB) ProductStRepo {
	repo := new(productStRepo)

	repo.db = db

	return repo
}
