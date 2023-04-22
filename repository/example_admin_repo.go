package repository

import (
	"database/sql"
	"fmt"
	"log"

	"go_inven_ctrl/entity"
)

type ExampleAdminRepo interface {
	GetAllExampleAdmin() any
	GetByIdExampleAdmin(id string) any
	CreateExampleAdmin(newExampleAdmin *entity.ExampleAdmin) string
	UpdateExampleAdmin(exampleAdmin *entity.ExampleAdmin) string
	DeleteExampleAdmin(id string) string
}

type exampleAdminRepo struct {
	db *sql.DB
}

func (r *exampleAdminRepo) GetAllExampleAdmin() any {
	fmt.Println("test")
	var exampleAdmins []entity.ExampleAdmin

	query := "SELECT id,name, email,phone,photo from st_team order by id asc"
	rows, err := r.db.Query(query)

	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var exampleAdmin entity.ExampleAdmin

		if err := rows.Scan(&exampleAdmin.ID, &exampleAdmin.Name, &exampleAdmin.Email, &exampleAdmin.Phone, &exampleAdmin.Photo); err != nil {
			log.Println(err)
		}

		exampleAdmins = append(exampleAdmins, exampleAdmin)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}
	fmt.Println(exampleAdmins)
	if len(exampleAdmins) == 0 {
		return "no data"
	}

	return exampleAdmins

}

func (r *exampleAdminRepo) GetByIdExampleAdmin(id string) any {
	var exampleadmin entity.ExampleAdmin

	query := "SELECT id, name, email, phone, photo from st_team WHERE id = $1"

	row := r.db.QueryRow(query, id)

	if err := row.Scan(&exampleadmin.ID, &exampleadmin.Name, &exampleadmin.Email, &exampleadmin.Phone, &exampleadmin.Photo); err != nil {
		log.Println(err)
	}

	if exampleadmin.ID == "" {
		return "example admin not found"
	}

	return exampleadmin

}

func (r *exampleAdminRepo) CreateExampleAdmin(newExampleAdmin *entity.ExampleAdmin) string {
	query := "INSERT INTO st_team ( id, name, email,password,phone,photo) VALUES($1,$2,$3,$4,$5,$6)"
	_, err := r.db.Exec(query, newExampleAdmin.ID, newExampleAdmin.Name, newExampleAdmin.Email, newExampleAdmin.Password, newExampleAdmin.Phone, newExampleAdmin.Photo)

	if err != nil {
		log.Println(err)
		return "failed to create example admin"
	}

	return "Example created successfully"
}

func (r *exampleAdminRepo) UpdateExampleAdmin(exampleAdmin *entity.ExampleAdmin) string {
	res := r.GetByIdExampleAdmin(exampleAdmin.ID)
	if res == "example admin not found" {
		return res.(string)
	}

	query := "UPDATE st_team SET  name = $1, email = $2, password =$3, phone =$4, photo =$5 WHERE id = $6 ;"
	_, err := r.db.Exec(query, exampleAdmin.Name, exampleAdmin.Email, exampleAdmin.Password, exampleAdmin.Phone, exampleAdmin.Photo, exampleAdmin.ID)

	if err != nil {
		log.Println(err)
		return "failed to update ExampleAdmin"
	}

	return fmt.Sprintf("ExampleAdmin with id %s updated successfully", exampleAdmin.ID)

}

func (r *exampleAdminRepo) DeleteExampleAdmin(id string) string {
	res := r.GetByIdExampleAdmin(id)
	if res == "example admin not found" {
		return res.(string)
	}

	query := "DELETE FROM st_team WHERE id =$1"
	_, err := r.db.Exec(query, id)

	if err != nil {
		log.Println(err)
		return "failed to delete productSt"
	}

	return fmt.Sprintf("ExampleAdmin with id %s deleted successfully", id)
}

func NewExampleAdminRepo(db *sql.DB) ExampleAdminRepo {
	repo := new(exampleAdminRepo)

	repo.db = db

	return repo
}
