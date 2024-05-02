package controllers

import (
	"florist-gin/business/ordersproducts"
	"florist-gin/controller/ordersproducts/request"
	"florist-gin/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrdersProductsController struct {
	usecase ordersproducts.OrdersProductsUseCaseInterface
}

func NewOrdersProductsController(uc ordersproducts.OrdersProductsUseCaseInterface) *OrdersProductsController {
	return &OrdersProductsController{
		usecase: uc,
	}
}

func (controller *OrdersProductsController) AddOrdersProducts(c *gin.Context) {
	if c.Request.Method != "POST" {
		utils.ErrorResponseWithoutMessages(c, http.StatusMethodNotAllowed, "Invalid HTTP method")
		return
	}

	var orderAdd request.OrdersProducts

	err := c.Bind(&orderAdd)

	if err != nil {
		utils.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Error binding the order data")
		return
	}

	order, errRepo := controller.usecase.AddOrdersProducts(*orderAdd.ToUsecase())

	if errRepo != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error in repo", errRepo)
		return
	}

	utils.SuccessResponse(c, order, []string{"OrdersProducts successfully created"})
}

func (controller *OrdersProductsController) EditOrdersProducts(c *gin.Context) {
	if c.Request.Method != "PUT" {
		utils.ErrorResponseWithoutMessages(c, http.StatusMethodNotAllowed, "Invalid HTTP method")
		return
	}

	var ordersProductsEdit request.OrdersProducts

	ordersProductsId, err := strconv.ParseUint(c.Param("ordersProductsId"), 10, 32)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "OrdersProducts ID must be an integer", err)
		c.Abort()
		return
	}

	parseUint32 := int(ordersProductsId)

	err = c.Bind(&ordersProductsEdit)

	if err != nil {
		utils.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Error binding the order data")
		return
	}

	order, errRepo := controller.usecase.EditOrdersProducts(*ordersProductsEdit.ToUsecase(), parseUint32)

	if errRepo != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error in repo", errRepo)
		return
	}

	utils.SuccessResponse(c, order, []string{"OrdersProducts successfully edited"})
}

func (controller *OrdersProductsController) DeleteOrdersProducts(c *gin.Context) {
	if c.Request.Method != "DELETE" {
		utils.ErrorResponseWithoutMessages(c, http.StatusMethodNotAllowed, "Invalid HTTP method")
		return
	}

	ordersProductsId, err := strconv.ParseUint(c.Param("ordersProductsId"), 10, 32)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "OrdersProducts ID must be an integer", err)
		c.Abort()
		return
	}

	parseUint32 := int(ordersProductsId)

	order, errRepo := controller.usecase.DeleteOrdersProducts(parseUint32)

	if errRepo != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error in repo", errRepo)
		return
	}

	utils.SuccessResponse(c, order, []string{"OrdersProducts successfully deleted"})
}

func (controller *OrdersProductsController) GetOrdersProductsDetail(c *gin.Context) {
	if c.Request.Method != "GET" {
		utils.ErrorResponseWithoutMessages(c, http.StatusMethodNotAllowed, "Invalid HTTP method")
		return
	}

	ordersProductsId, err := strconv.ParseUint(c.Param("ordersProductsId"), 10, 32)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "OrdersProducts ID must be an integer", err)
		c.Abort()
		return
	}

	parseUint32 := int(ordersProductsId)

	order, errRepo := controller.usecase.GetOrdersProductsDetail(parseUint32)

	if errRepo != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error in repo", errRepo)
		return
	}

	utils.SuccessResponse(c, order, []string{"Successfully get order"})
}
