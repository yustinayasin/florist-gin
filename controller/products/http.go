package controllers

import (
	"florist-gin/business/products"
	"florist-gin/controller/products/request"
	"florist-gin/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	usecase products.ProductUseCaseInterface
}

func NewProductController(uc products.ProductUseCaseInterface) *ProductController {
	return &ProductController{
		usecase: uc,
	}
}

func (controller *ProductController) AddProduct(c *gin.Context) {
	if c.Request.Method != "POST" {
		utils.ErrorResponseWithoutMessages(c, http.StatusMethodNotAllowed, "Invalid HTTP method")
		return
	}

	fmt.Println("test")

	var productAddProduct request.Product

	err := c.BindJSON(&productAddProduct)

	if err != nil {
		utils.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Error binding the product data")
		return
	}

	fmt.Println(productAddProduct)

	product, errRepo := controller.usecase.AddProduct(*productAddProduct.ToUsecase())

	if errRepo != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error in reposss", errRepo)
		return
	}

	utils.SuccessResponse(c, product, []string{"Product successfully created"})
}

func (controller *ProductController) EditProduct(c *gin.Context) {
	if c.Request.Method != "PUT" {
		utils.ErrorResponseWithoutMessages(c, http.StatusMethodNotAllowed, "Invalid HTTP method")
		return
	}

	var productEdit request.Product

	productId, err := strconv.Atoi(c.Param("productId"))

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Product ID must be an integer", err)
		c.Abort()
		return
	}

	err = c.BindJSON(&productEdit)

	if err != nil {
		utils.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Error binding the product data")
		return
	}

	product, errRepo := controller.usecase.EditProduct(*productEdit.ToUsecase(), productId)

	if errRepo != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error in repo", errRepo)
		return
	}

	utils.SuccessResponse(c, product, []string{"Product successfully edited"})
}

func (controller *ProductController) DeleteProduct(c *gin.Context) {
	if c.Request.Method != "DELETE" {
		utils.ErrorResponseWithoutMessages(c, http.StatusMethodNotAllowed, "Invalid HTTP method")
		return
	}

	productId, err := strconv.Atoi(c.Param("productId"))

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Product ID must be an integer", err)
		c.Abort()
		return
	}

	product, errRepo := controller.usecase.DeleteProduct(productId)

	if errRepo != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error in repo", errRepo)
		return
	}

	utils.SuccessResponse(c, product, []string{"Product successfully deleted"})
}

func (controller *ProductController) GetProductDetail(c *gin.Context) {
	if c.Request.Method != "GET" {
		utils.ErrorResponseWithoutMessages(c, http.StatusMethodNotAllowed, "Invalid HTTP method")
		return
	}

	productId, err := strconv.Atoi(c.Param("productId"))

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Product ID must be an integer", err)
		c.Abort()
		return
	}

	product, errRepo := controller.usecase.GetProductDetail(productId)

	if errRepo != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error in repo", errRepo)
		return
	}

	utils.SuccessResponse(c, product, []string{"Successfully get product"})
}
