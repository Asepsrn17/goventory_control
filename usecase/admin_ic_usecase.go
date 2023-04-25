package usecase

import (
	"go_inven_ctrl/entity"
	"go_inven_ctrl/repository"
)


type AdminIcUsecase interface {
	FindAdminIc() any
	FindAdminIcByid(id string) any
	Register(newAdminIc *entity.AdminIc) string
	Edit(adminIc *entity.AdminIc) string
	Unreg(id string) string
	Login(email, password string) (*entity.AdminIc, error)
}

type adminIcUsecase struct {
	adminIcRepo repository.AdminIcRepo
}

func (adm *adminIcUsecase) FindAdminIc() any {
	return adm.adminIcRepo.GetAll()
}

func (adm *adminIcUsecase) FindAdminIcByid(id string) any {
	return adm.adminIcRepo.GetByid(id)
}

func (adm *adminIcUsecase) Register(newAdminIc *entity.AdminIc) string {
	return adm.adminIcRepo.Create(newAdminIc)
}

func (adm *adminIcUsecase) Edit(adminIc *entity.AdminIc) string {
	return adm.adminIcRepo.Update(adminIc)
}

func (adm *adminIcUsecase) Unreg(id string) string {
	return adm.adminIcRepo.Delete(id)
}

func (adm *adminIcUsecase) Login(email, password string) (*entity.AdminIc, error) {
	return adm.adminIcRepo.Verify(email, password)
}

func NewAdminIcUsecase(adminIcRepo repository.AdminIcRepo) AdminIcUsecase {
	return &adminIcUsecase {
		adminIcRepo: adminIcRepo,
	}
}