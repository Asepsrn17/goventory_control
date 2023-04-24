package usecase

import (
	"go_inven_ctrl/entity"
	"go_inven_ctrl/repository"
	"go_inven_ctrl/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type WarehouseTeamUsecase interface {
	FindEmployees() any
	FindEmployeeById(id int) any
	Login(req *entity.EmployeeReq) (*entity.LoginEmployeeRes, error)
	Register(req *entity.WarehouseTeam) (*entity.CreateEmployeeRes, error)
	Edit(employee *entity.WarehouseTeam) string
	Unreg(id int) string
}

type warehouseTeamUsecase struct {
	warehouseTeamRepo repository.WarehouseTeamRepo
}

func NewWarehouseTeamUsecase(warehouseTeamRepo repository.WarehouseTeamRepo) WarehouseTeamUsecase {
	return &warehouseTeamUsecase{
		warehouseTeamRepo: warehouseTeamRepo,
	}
}

func (u *warehouseTeamUsecase) FindEmployees() any {
	return u.warehouseTeamRepo.GetAll()
}

func (u *warehouseTeamUsecase) FindEmployeeById(id int) any {
	return u.warehouseTeamRepo.GetById(id)
}

type MyJWTClaim struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.Claims
}

var jwtKey = []byte("enigmacamp")

func (u *warehouseTeamUsecase) Login(req *entity.EmployeeReq) (*entity.LoginEmployeeRes, error) {
	s, err := u.warehouseTeamRepo.GetByName(req.Name)
	if err != nil {
		return &entity.LoginEmployeeRes{}, err
	}

	err = utils.CheckPassword(req.Password, s.Password)
	if err != nil {
		return &entity.LoginEmployeeRes{}, err
	}

	// Create new JWT token with user ID and username as claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaim{
		ID:    s.ID,
		Name:  s.Name,
		Email: s.Email,
		Claims: jwt.MapClaims{
			"name":  s.Name,
			"email": s.Email,
			"exp":   time.Now().Add(time.Minute * 1).Unix(),
		},
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return &entity.LoginEmployeeRes{}, err
	}

	return &entity.LoginEmployeeRes{AccessToken: tokenString, ID: s.ID, Name: s.Name, Email: s.Email}, nil
}

func (u *warehouseTeamUsecase) Register(req *entity.WarehouseTeam) (*entity.CreateEmployeeRes, error) {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	s := &entity.WarehouseTeam{
		ID:       req.ID,
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Phone:    req.Phone,
		Photo:    req.Photo,
	}

	r, err := u.warehouseTeamRepo.Create((s))
	if err != nil {
		return nil, err
	}

	res := &entity.CreateEmployeeRes{
		ID:    r.ID,
		Name:  r.Name,
		Email: r.Email,
	}

	return res, nil
}

func (u *warehouseTeamUsecase) Edit(employee *entity.WarehouseTeam) string {
	return u.warehouseTeamRepo.Update(employee)
}

func (u *warehouseTeamUsecase) Unreg(id int) string {
	return u.warehouseTeamRepo.Delete(id)
}
