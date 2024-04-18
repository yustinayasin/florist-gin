package controllers

import (
	"encoding/base64"
	"florist-gin/business/products"
	"florist-gin/controller/products/request"
	"florist-gin/controller/products/response"
	"florist-gin/utils"
	"fmt"
	"io"
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

	// Open the file
	photo, err := file.Open()
	if err != nil {
		utils.ErrorResponseWithoutMessages(c, http.StatusInternalServerError, "Failed to open file")
		return
	}
	defer photo.Close()

	// Read the file content into a byte slice
	photoData, err := io.ReadAll(photo)
	if err != nil {
		utils.ErrorResponseWithoutMessages(c, http.StatusInternalServerError, "Failed to read file content")
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

	err = c.Bind(&productAddProduct)

	if err != nil {
		utils.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Error binding the product data")
		return
	}

	productAddProduct.FileName = uniqueFileName
	productAddProduct.File = photoData

	product, errRepo := controller.usecase.AddProduct(*productAddProduct.ToUsecase())

	if errRepo != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error in repo", errRepo)
		return
	}

	encodedImage := base64.StdEncoding.EncodeToString(product.File)

	product.File = []byte(encodedImage)

	utils.SuccessResponse(c, product, []string{"Product successfully created"})
}

func (controller *ProductController) EditProduct(c *gin.Context) {
	if c.Request.Method != "PUT" {
		utils.ErrorResponseWithoutMessages(c, http.StatusMethodNotAllowed, "Invalid HTTP method")
		return
	}

	// Retrieve the file from the form data
	file, errPhoto := c.FormFile("photo")
	if errPhoto != nil {
		utils.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "No file uploaded")
		return
	}

	// Open the file
	photo, err := file.Open()
	if err != nil {
		utils.ErrorResponseWithoutMessages(c, http.StatusInternalServerError, "Failed to open file")
		return
	}
	defer photo.Close()

	// Read the file content into a byte slice
	photoData, err := io.ReadAll(photo)
	if err != nil {
		utils.ErrorResponseWithoutMessages(c, http.StatusInternalServerError, "Failed to read file content")
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

	var productEdit request.Product

	productId, err := strconv.Atoi(c.Param("productId"))

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Product ID must be an integer", err)
		c.Abort()
		return
	}

	err = c.Bind(&productEdit)

	if err != nil {
		utils.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Error binding the product data")
		return
	}

	productEdit.FileName = uniqueFileName
	productEdit.File = photoData

	product, errRepo := controller.usecase.EditProduct(*productEdit.ToUsecase(), productId)

	if errRepo != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error in repo", errRepo)
		return
	}

	encodedImage := base64.StdEncoding.EncodeToString(product.File)

	product.File = []byte(encodedImage)

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

	parseUint32 := int(productId)

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

	productResponse := response.FromUsecase(product)

	utils.SuccessResponse(c, productResponse, []string{"Successfully get product"})
}

func (controller *ProductController) GetAllProduct(c *gin.Context) {
	if c.Request.Method != "GET" {
		utils.ErrorResponseWithoutMessages(c, http.StatusMethodNotAllowed, "Invalid HTTP method")
		return
	}

	categoryIdStr := c.Query("categoryId")
	var categoryId int

	if categoryIdStr != "" {
		var err error
		categoryId, err = strconv.Atoi(categoryIdStr)

		if err != nil {
			utils.ErrorResponse(c, http.StatusBadRequest, "Category ID must be an integer", err)
			c.Abort()
			return
		}
	}

	product, errRepo := controller.usecase.GetAllProduct(categoryId)

	if errRepo != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error in repo", errRepo)
		return
	}

	productResponse := response.FromUsecaseList(product)

	utils.SuccessResponse(c, productResponse, []string{"Successfully get product"})
}
