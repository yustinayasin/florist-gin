package controllers

import (
	"florist-gin/business/orders"
	"florist-gin/controller/orders/request"
	"florist-gin/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	usecase orders.OrderUseCaseInterface
}

func NewOrderController(uc orders.OrderUseCaseInterface) *OrderController {
	return &OrderController{
		usecase: uc,
	}
}

func (controller *OrderController) AddOrder(c *gin.Context) {
	if c.Request.Method != "POST" {
		utils.ErrorResponseWithoutMessages(c, http.StatusMethodNotAllowed, "Invalid HTTP method")
		return
	}

	var orderAdd request.Order

	err := c.Bind(&orderAdd)

	if err != nil {
		utils.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Error binding the order data")
		return
	}

	order, errRepo := controller.usecase.AddOrder(*orderAdd.ToUsecase())

	if errRepo != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error in repo", errRepo)
		return
	}

	utils.SuccessResponse(c, order, []string{"Order successfully created"})
}

func (controller *OrderController) EditOrder(c *gin.Context) {
	if c.Request.Method != "PUT" {
		utils.ErrorResponseWithoutMessages(c, http.StatusMethodNotAllowed, "Invalid HTTP method")
		return
	}

	var orderEdit request.Order

	orderId, err := strconv.ParseUint(c.Param("orderId"), 10, 32)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Order ID must be an integer", err)
		c.Abort()
		return
	}

	parseUint32 := int(orderId)

	err = c.Bind(&orderEdit)

	if err != nil {
		utils.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Error binding the order data")
		return
	}

	if orderEdit.Status != "true" && orderEdit.Status != "false" {
		utils.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Status should be boolean value")
		return
	}

	order, errRepo := controller.usecase.EditOrder(*orderEdit.ToUsecase(), parseUint32)

	if errRepo != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error in repo", errRepo)
		return
	}

	utils.SuccessResponse(c, order, []string{"Order successfully edited"})
}

func (controller *OrderController) DeleteOrder(c *gin.Context) {
	if c.Request.Method != "DELETE" {
		utils.ErrorResponseWithoutMessages(c, http.StatusMethodNotAllowed, "Invalid HTTP method")
		return
	}

	orderId, err := strconv.ParseUint(c.Param("orderId"), 10, 32)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Order ID must be an integer", err)
		c.Abort()
		return
	}

	parseUint32 := int(orderId)

	order, errRepo := controller.usecase.DeleteOrder(parseUint32)

	if errRepo != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error in repo", errRepo)
		return
	}

	utils.SuccessResponse(c, order, []string{"Order successfully deleted"})
}

func (controller *OrderController) GetOrderDetail(c *gin.Context) {
	if c.Request.Method != "GET" {
		utils.ErrorResponseWithoutMessages(c, http.StatusMethodNotAllowed, "Invalid HTTP method")
		return
	}

	orderId, err := strconv.ParseUint(c.Param("orderId"), 10, 32)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Order ID must be an integer", err)
		c.Abort()
		return
	}

	parseUint32 := int(orderId)

	order, errRepo := controller.usecase.GetOrderDetail(parseUint32)

	if errRepo != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error in repo", errRepo)
		return
	}

	utils.SuccessResponse(c, order, []string{"Successfully get order"})
}
