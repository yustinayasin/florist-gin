package controllers

import (
	"florist-gin/business/carts"
	"florist-gin/business/users"
	"florist-gin/controller/carts/response"
	"florist-gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CartController struct {
	usecase carts.CartUseCaseInterface
}

func NewCartController(uc carts.CartUseCaseInterface) *CartController {
	return &CartController{
		usecase: uc,
	}
}

func (controller *CartController) GetCart(c *gin.Context) {
	if c.Request.Method != "GET" {
		utils.ErrorResponseWithoutMessages(c, http.StatusMethodNotAllowed, "Invalid HTTP method")
		return
	}

	// cartUserId, err := strconv.ParseUint(c.Param("cartUserId"), 10, 32)

	// if err != nil {
	// 	utils.ErrorResponse(c, http.StatusBadRequest, "Cart user ID must be an integer", err)
	// 	c.Abort()
	// 	return
	// }

	// paramUint32 := int(cartUserId)

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

	cart, errRepo := controller.usecase.GetCart(user.Id)

	if errRepo != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error in repo", errRepo)
		return
	}

	cartResponse := response.FromUsecase(cart)

	utils.SuccessResponse(c, cartResponse, []string{"Successfully get the cart"})
}
