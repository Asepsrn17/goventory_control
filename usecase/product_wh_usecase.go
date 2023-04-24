package usecase

import (
	"go_inven_ctrl/entity"
	"go_inven_ctrl/repository"
)

type ProductWhUsecase interface {
	FindProducts() any
	FindProductById(id int) any
	FindProductByName(name string) (*entity.ProductWh, error)
	Input(newProduct *entity.ProductWh) (*entity.ProductWh, error)
	Edit(product *entity.ProductWh) string
	Output(id int) string
}

type productWhUsecase struct {
	productWhRepo repository.ProductWhRepo
}

func NewProductWhUsecase(productWhRepo repository.ProductWhRepo) ProductWhUsecase {
	return &productWhUsecase{
		productWhRepo: productWhRepo,
	}
}

func (u *productWhUsecase) FindProducts() any {
	return u.productWhRepo.GetAll()
}

func (u *productWhUsecase) FindProductById(id int) any {
	return u.productWhRepo.GetById(id)
}

func (u *productWhUsecase) FindProductByName(name string) (*entity.ProductWh, error) {
	return u.productWhRepo.GetByName(name)
}

func (u *productWhUsecase) Input(newProduct *entity.ProductWh) (*entity.ProductWh, error) {
	return u.productWhRepo.Create(newProduct)
}

func (u *productWhUsecase) Edit(product *entity.ProductWh) string {
	return u.productWhRepo.Update(product)
}

func (u *productWhUsecase) Output(id int) string {
	return u.productWhRepo.Delete(id)
}
