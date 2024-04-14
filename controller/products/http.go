package controllers

import (
	"florist-gin/business/products"
	"florist-gin/controller/products/request"
	"florist-gin/utils"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

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

	// Retrieve the file from the form data
	file, errPhoto := c.FormFile("photo")
	if errPhoto != nil {
		utils.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "No file uploaded")
		return
	}

	// Specify the maximum file size (in bytes)
	maxFileSize := int64(10 * 1024 * 1024) // 10 MB

	// Check if the file size exceeds the maximum size
	if file.Size > maxFileSize {
		utils.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "File size exceeds the maximum allowed size")
		return
	}

	// Retrieve the file name from the file header
	fileName := file.Filename

	// Create a unique file name to avoid overwriting existing files
	uniqueFileName := fmt.Sprintf("%s-%d%s", fileName[:len(fileName)-len(filepath.Ext(fileName))], time.Now().UnixNano(), filepath.Ext(fileName))

	fileReader, errOpenFile := file.Open()

	if errOpenFile != nil {
		utils.ErrorResponseWithoutMessages(c, http.StatusInternalServerError, "Failed to open file")
		return
	}
	defer fileReader.Close()

	var productAddProduct request.Product

	err := c.BindJSON(&productAddProduct)

	if err != nil {
		utils.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Error binding the product data")
		return
	}

	productAddProduct.FileName = uniqueFileName
	productAddProduct.File = fileReader

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

	productId, err := strconv.ParseUint(c.Param("productId"), 10, 32)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Product ID must be an integer", err)
		c.Abort()
		return
	}

	parseUint32 := uint32(productId)

	err = c.BindJSON(&productEdit)

	if err != nil {
		utils.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Error binding the product data")
		return
	}

	product, errRepo := controller.usecase.EditProduct(*productEdit.ToUsecase(), parseUint32)

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

	productId, err := strconv.ParseUint(c.Param("productId"), 10, 32)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Product ID must be an integer", err)
		c.Abort()
		return
	}

	parseUint32 := uint32(productId)

	product, errRepo := controller.usecase.DeleteProduct(parseUint32)

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

	productId, err := strconv.ParseUint(c.Param("productId"), 10, 32)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Product ID must be an integer", err)
		c.Abort()
		return
	}

	parseUint32 := uint32(productId)

	product, errRepo := controller.usecase.GetProductDetail(parseUint32)

	if errRepo != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error in repo", errRepo)
		return
	}

	utils.SuccessResponse(c, product, []string{"Successfully get product"})
}
