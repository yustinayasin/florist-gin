package controllers

import (
	"florist-gin/business/cartsproducts"
	"florist-gin/business/users"
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

	// Retrieve the user from the context
	userInterface, exists := c.Get("user")

	if !exists {
		utils.ErrorResponseWithoutMessages(c, http.StatusUnauthorized, "User not found in context")
		return
	}

	// Type assert the user to the correct type. userInterface variable it's not a pointer
	user, ok := userInterface.(users.User) // Adjust the type according to your actual user struct
	if !ok {
		utils.ErrorResponseWithoutMessages(c, http.StatusInternalServerError, "Failed to type assert user")
		return
	}

	var cartsProducts request.CartsProducts

	err := c.Bind(&cartsProducts)

	if err != nil {
		utils.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Error binding the carts products data")
		return
	}

	cartsproducts, errRepo := controller.usecase.AddProductToCart(*cartsProducts.ToUsecase(), user.Id)

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

	cartsProductsId, err := strconv.ParseUint(c.Param("cartsProductsId"), 10, 32)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Carts Products ID must be an integer", err)
		c.Abort()
		return
	}

	parseUint32 := int(cartsProductsId)

	err = c.Bind(&cartProducts)

	if err != nil {
		utils.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Error binding the carts products data")
		return
	}

	cartsproducts, errRepo := controller.usecase.EditProductFromCart(*cartProducts.ToUsecase(), parseUint32)

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

	cartsProductsId, err := strconv.ParseUint(c.Param("cartsProductsId"), 10, 32)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Carts Products ID must be an integer", err)
		c.Abort()
		return
	}

	parseUint32 := int(cartsProductsId)

	cartsproducts, errRepo := controller.usecase.DeleteProductFromCart(parseUint32)

	if errRepo != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error in repo", errRepo)
		return
	}

	utils.SuccessResponse(c, cartsproducts, []string{"Product successfully deleted from the cart"})
}
