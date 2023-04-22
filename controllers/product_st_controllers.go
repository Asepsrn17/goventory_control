package controllers

import (
	"fmt"
	"net/http"

	"go_inven_ctrl/entity"
	"go_inven_ctrl/usecase"

	// "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type ProductStController struct {
	usecase usecase.ProductStUsecase
}

func (c *ProductStController) FindAllProductsSt(ctx *gin.Context) {
	// claims := ctx.MustGet("claims").(jwt.MapClaims)
	// username := claims["username"].(string)

	res := c.usecase.FindAllProductsSt()
	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

func (c *ProductStController) FindProductsStById(ctx *gin.Context) {
	// claims := ctx.MustGet("claims").(jwt.MapClaims)
	// username := claims["username"].(string)

	id := ctx.Param("id")
	fmt.Println(id)
	res := c.usecase.FindProductStById(id)
	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

func (c *ProductStController) RegisterProductSt(ctx *gin.Context) {
	var newProductSt entity.ProductSt

	if err := ctx.BindJSON(&newProductSt); err != nil {
		ctx.JSON(http.StatusBadRequest, "invalid input")
		fmt.Println(err)
		return
	}

	res := c.usecase.RegisterProductSt(&newProductSt)

	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

func (c *ProductStController) EditProductSt(ctx *gin.Context) {
	// claims := ctx.MustGet("claims").(jwt.MapClaims)
	// username := claims["username"].(string)

	var productSt entity.ProductSt

	if err := ctx.BindJSON(&productSt); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		return
	}

	res := c.usecase.EditProductSt(&productSt)
	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

func (c *ProductStController) UnregProductSt(ctx *gin.Context) {
	// claims := ctx.MustGet("claims").(jwt.MapClaims)
	// username := claims["username"].(string)

	id := ctx.Param("id")

	res := c.usecase.UnregProductSt(id)
	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

func NewProductStController(u usecase.ProductStUsecase) *ProductStController {
	controller := ProductStController{
		usecase: u,
	}

	return &controller
}
