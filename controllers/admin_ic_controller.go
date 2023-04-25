package controllers

import (
	"go_inven_ctrl/entity"
	"go_inven_ctrl/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminIcController struct {
	usecase usecase.AdminIcUsecase
}

func (c *AdminIcController) FindAdminIc(ctx *gin.Context) {
	res := c.usecase.FindAdminIc()

	ctx.JSON(http.StatusOK, res)
}

func (c *AdminIcController) FindAdminIcByid(ctx *gin.Context) {
	id := ctx.Param("id")
	res := c.usecase.FindAdminIcByid(id)
	if res == "admin not found" {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "admin not found"})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *AdminIcController) Register(ctx *gin.Context) {

	var newAdminIc entity.AdminIc

	if err := ctx.BindJSON(&newAdminIc); err != nil {
		ctx.JSON(http.StatusBadRequest, "invalid input")
		return
	}
	res := c.usecase.Register(&newAdminIc)
	ctx.JSON(http.StatusCreated, res)
}

func (c *AdminIcController) Edit(ctx *gin.Context) {

	var adminIc entity.AdminIc

	if err := ctx.BindJSON(&adminIc); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := c.usecase.Edit(&adminIc)
	ctx.JSON(http.StatusOK, res)
}


func (c *AdminIcController) Unreg(ctx *gin.Context) {
	id := ctx.Param("id")
	res := c.usecase.Unreg(id)
	if res == "admin not found" {
		ctx.JSON(http.StatusBadRequest, "invalid input ID")
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func NewAdminIcController(c usecase.AdminIcUsecase) * AdminIcController {
	controller := AdminIcController {
		usecase: c,
	}
	return &controller
}