package controllers

import (
	"florist-gin/business/users"
	"florist-gin/controller/users/request"
	"florist-gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	usecase users.UserUseCaseInterface
}

func NewUserController(uc users.UserUseCaseInterface) *UserController {
	return &UserController{
		usecase: uc,
	}
}

func (controller *UserController) SignUp(c *gin.Context) {
	// initialize a new context bcs in the usecase require context not gin context
	// check the method
	if c.Request.Method != "POST" {
		utils.ErrorResponseWithoutMessages(c, http.StatusMethodNotAllowed, "Invalid HTTP method")
	}

	var userSignUp request.UserEdit

	err := c.Bind(&userSignUp)

	if err != nil {
		utils.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Error binding the user data")
	}

	// why pointer to user sign up? kenapa juga to use case di bind ke pointer
	user, errRepo := controller.usecase.SignUp(*userSignUp.ToUsecase())

	if errRepo != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error in repo", errRepo)
	}

	utils.SuccessResponse(c, user)
}

func (controller *UserController) Login(c *gin.Context) {
	// check the method
	if c.Request.Method != "GET" {
		utils.ErrorResponseWithoutMessages(c, http.StatusMethodNotAllowed, "Invalid HTTP method")
	}

	var userLogin request.UserLogin

	err := c.Bind(&userLogin)

	if err != nil {
		utils.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Error binding the user data")
	}

	// why pointer to user sign up? kenapa juga to use case di bind ke pointer
	user, errRepo := controller.usecase.SignUp(*userLogin.ToUsecase())

	if errRepo != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error in repo", errRepo)
	}

	utils.SuccessResponse(c, user)
}
