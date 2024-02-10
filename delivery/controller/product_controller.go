package controller

import (
	"main/config"
	"main/model"
	"main/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productUc usecase.ProductUsecase
	rg        *gin.RouterGroup
}

func (p *ProductController) createHandler(c *gin.Context) {
	var payload model.Product

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product, err := p.productUc.RegisterNewProduct(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
}

func (p *ProductController) listHandler(c *gin.Context) {
	product, err := p.productUc.FindAllProduct()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
}

func (p *ProductController) Route() {
	p.rg.POST(config.ProductPost, p.createHandler)
	p.rg.GET(config.ProductGet, p.listHandler)
}

func NewProductController(productUc usecase.ProductUsecase, rg *gin.RouterGroup) *ProductController {
	return &ProductController{productUc: productUc, rg: rg}
}
