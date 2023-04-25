package repository

import (
	"database/sql"
	"fmt"
	"go_inven_ctrl/entity"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type AdminIcRepo interface {
	GetAll() any
	GetByid(id string) any
	Create(newAdminIc *entity.AdminIc) string
	Update(adminIc *entity.AdminIc) string
	Delete(id string) string
	Verify(username, password string) (*entity.AdminIc, error)
}

type adminIcRepo struct {
	db *sql.DB
}

func (a *adminIcRepo) GetAll() any {
	var adminIcs []entity.AdminIc

	query := "SELECT id, name, email, phone, photo, password FROM ic_team"
	rows, err := a.db.Query(query)

	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var adminIc entity.AdminIc

		if err := rows.Scan(&adminIc.ID, &adminIc.Name, &adminIc.Email, &adminIc.Phone, &adminIc.Photo, &adminIc.Password); err != nil {
			log.Println(err)
		}
		adminIcs = append(adminIcs, adminIc)
	}

	if len(adminIcs) == 0 {
		return "no data"
	}
	return adminIcs
}

func (a *adminIcRepo) GetByid(id string) any {
	var adminIcInDb entity.AdminIc

	query := "SELECT id, name, email, phone, photo, password FROM ic_team WHERE id = $1"
	row := a.db.QueryRow(query, id)

	err := row.Scan(&adminIcInDb.ID, &adminIcInDb.Name, &adminIcInDb.Email, &adminIcInDb.Phone, &adminIcInDb.Photo, &adminIcInDb.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return "admin not found"
		}
		log.Println(err)
	}

	return adminIcInDb
}


func (a *adminIcRepo) Create(newAdminIc *entity.AdminIc) string {
	var emailExist bool

	query := "SELECT EXISTS(SELECT 1 FROM ic_team WHERE email=$1)"
	err := a.db.QueryRow(query, newAdminIc.Email).Scan(&emailExist)
	if err != nil {
		log.Println(err)
		return "failed to create admin"
	}
	if emailExist {
		return "email already exist"
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newAdminIc.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return "failed to create admin"
	}

	query1 := "INSERT INTO ic_team(id, name, email, phone, photo, password) VALUES ($1, $2, $3, $4, $5, $6)"
	_, exeErr := a.db.Exec(query1, newAdminIc.ID, newAdminIc.Name, newAdminIc.Email, newAdminIc.Phone, newAdminIc.Photo, string(hashedPassword))

	if exeErr != nil {
		log.Println(exeErr)
		return "failed to create user"
	}
	
	return "user craeted successfully"
}

func (a *adminIcRepo) Update(adminIc *entity.AdminIc) string {
	res := a.GetByid(adminIc.ID)

	if res == "admin not found" {
		return res.(string)
	}

	query := "UPDATE ic_team SET name = $1, email = $2, phone = $3, photo = $4, password = $5 WHERE id = $6"
	_, err := a.db.Exec(query, adminIc.Name, adminIc.Email, adminIc.Phone, adminIc.Photo, adminIc.Password, adminIc.ID)

	if err != nil {
		log.Println(err)
	}
	return fmt.Sprintf("admin ic with id %s updated successfully", adminIc.ID)

}

func (a *adminIcRepo) Delete(id string) string {
	res := a.GetByid(id)

	if res == "admin not found" {
		return res.(string)
	}

	query := "DELETE FROM ic_team WHERE id = $1"
	_, err := a.db.Exec(query, id)

	if err != nil {
		log.Println(err)
		return "failed to delete admin ic"
	}

	return fmt.Sprintf("admin ic with id %s deleted successfully", id)
}

func (a *adminIcRepo) Verify(email, password string) (*entity.AdminIc, error) {
	var adminIc entity.AdminIc

	query := "SELECT id, name, phone, photo FROM ic_team WHERE email = $1 AND password = $2"

	row := a.db.QueryRow(query, email, password)

	err := row.Scan(&adminIc.ID, &adminIc.Name)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("admin not found")
		} else {
			return nil, err
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(adminIc.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("invalid password")
	}
	return &adminIc, nil
}

func NewAdminIcRepo(db *sql.DB) AdminIcRepo {
	repo := new(adminIcRepo)
	repo.db = db
	return repo
}