package controllers

import (
	"florist-gin/business/cartsproducts"
	"florist-gin/controller/cartsproducts/request"
	"florist-gin/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartsProductsController struct {
	usecase cartsproducts.CartsProductsUseCaseInterface
}

func NewCartsProductsController(uc cartsproducts.CartsProductsUseCaseInterface) *CartsProductsController {
	return &CartsProductsController{
		usecase: uc,
	}
}

func (controller *CartsProductsController) AddProductToCart(c *gin.Context) {
	if c.Request.Method != "POST" {
		utils.ErrorResponseWithoutMessages(c, http.StatusMethodNotAllowed, "Invalid HTTP method")
		return
	}

	var cartsProducts request.CartsProducts

	err := c.BindJSON(&cartsProducts)

	if err != nil {
		utils.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Error binding the carts products data")
		return
	}

	cartsproducts, errRepo := controller.usecase.AddProductToCart(*cartsProducts.ToUsecase())

	if errRepo != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error in repo", errRepo)
		return
	}

	utils.SuccessResponse(c, cartsproducts, []string{"Product successfully added to the cart"})
}

func (controller *CartsProductsController) EditProductFromCart(c *gin.Context) {
	if c.Request.Method != "PUT" {
		utils.ErrorResponseWithoutMessages(c, http.StatusMethodNotAllowed, "Invalid HTTP method")
		return
	}

	var cartProducts request.CartsProducts

	cartsProductsId, err := strconv.Atoi(c.Param("cartsProductsId"))

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Carts Products ID must be an integer", err)
		c.Abort()
		return
	}

	err = c.BindJSON(&cartProducts)

	if err != nil {
		utils.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Error binding the carts products data")
		return
	}

	cartsproducts, errRepo := controller.usecase.EditProductFromCart(*cartProducts.ToUsecase(), cartsProductsId)

	if errRepo != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error in repo", errRepo)
		return
	}

	utils.SuccessResponse(c, cartsproducts, []string{"Product successfully edited from the cart"})
}

func (controller *CartsProductsController) DeleteProductFromCart(c *gin.Context) {
	if c.Request.Method != "DELETE" {
		utils.ErrorResponseWithoutMessages(c, http.StatusMethodNotAllowed, "Invalid HTTP method")
		return
	}

	cartsProductsId, err := strconv.Atoi(c.Param("cartsProductsId"))

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Carts Products ID must be an integer", err)
		c.Abort()
		return
	}

	cartsproducts, errRepo := controller.usecase.DeleteProductFromCart(cartsProductsId)

	if errRepo != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error in repo", errRepo)
		return
	}

	utils.SuccessResponse(c, cartsproducts, []string{"Product successfully deleted from the cart"})
}
