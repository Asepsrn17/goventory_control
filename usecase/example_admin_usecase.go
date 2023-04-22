package usecase

import (
	"go_inven_ctrl/entity"
	"go_inven_ctrl/repository"
)

type ExampleAdminUsecase interface {
	FindAllExampleAdmin() any
	FindByIdExampleAdmin(id string) any
	RegisterExampleAdmin(newExampleAdmin *entity.ExampleAdmin) string
	EditExampleAdmin(exampleAdmin *entity.ExampleAdmin) string
	UnregExampleAdmin(id string) string
}

type exampleAdminUsecase struct {
	exampleAdminRepo repository.ExampleAdminRepo
}

func (u *exampleAdminUsecase) FindAllExampleAdmin() any {
	return u.exampleAdminRepo.GetAllExampleAdmin()
}

func (u *exampleAdminUsecase) FindByIdExampleAdmin(id string) any {
	return u.exampleAdminRepo.GetByIdExampleAdmin(id)
}

func (u *exampleAdminUsecase) RegisterExampleAdmin(newExampleAdmin *entity.ExampleAdmin) string {
	return u.exampleAdminRepo.CreateExampleAdmin(newExampleAdmin)
}

func (u *exampleAdminUsecase) EditExampleAdmin(exampleAdmin *entity.ExampleAdmin) string {
	return u.exampleAdminRepo.UpdateExampleAdmin(exampleAdmin)
}

func (u *exampleAdminUsecase) UnregExampleAdmin(id string) string {
	return u.exampleAdminRepo.DeleteExampleAdmin(id)
}

func NewExampleAdminUsecase(exampleAdmin repository.ExampleAdminRepo) ExampleAdminUsecase {
	return &exampleAdminUsecase{
		exampleAdminRepo: exampleAdmin,
	}
}
