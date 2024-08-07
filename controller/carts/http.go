package controllers

import (
	"florist-gin/business/carts"
	"florist-gin/utils"
	"net/http"
	"strconv"

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

	cartId, err := strconv.ParseUint(c.Param("cartId"), 10, 32)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Cart ID must be an integer", err)
		c.Abort()
		return
	}

	paramUint32 := int(cartId)

	cart, errRepo := controller.usecase.GetCart(paramUint32)

	if errRepo != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error in repo", errRepo)
		return
	}

	utils.SuccessResponse(c, cart, []string{"Successfully get the cart"})
}
