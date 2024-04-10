package controllers

import (
	"florist-gin/business/users"
	"florist-gin/controller/users/request"
	"florist-gin/utils"
	"net/http"
	"strconv"

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
	if c.Request.Method != "POST" {
		utils.ErrorResponseWithoutMessages(c, http.StatusMethodNotAllowed, "Invalid HTTP method")
		return
	}

	var userSignUp request.UserEdit

	err := c.BindJSON(&userSignUp)

	if err != nil {
		utils.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Error binding the user data")
		return
	}

	// why pointer to user sign up? kenapa juga to use case di bind ke pointer
	user, errRepo := controller.usecase.SignUp(*userSignUp.ToUsecase())

	if errRepo != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error in repo", errRepo)
		return
	}

	utils.SuccessResponse(c, user, []string{"User successfully created"})
}

func (controller *UserController) Login(c *gin.Context) {
	if c.Request.Method != "POST" {
		utils.ErrorResponseWithoutMessages(c, http.StatusMethodNotAllowed, "Invalid HTTP method")
		return
	}

	var userLogin request.UserLogin

	err := c.BindJSON(&userLogin)

	if err != nil {
		utils.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Error binding the user data")
		return
	}

	user, errRepo := controller.usecase.Login(*userLogin.ToUsecase())

	if errRepo != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error in repo", errRepo)
		return
	}

	utils.SuccessResponse(c, user, []string{"User successfully login"})
}

func (controller *UserController) EditUser(c *gin.Context) {
	if c.Request.Method != "PUT" {
		utils.ErrorResponseWithoutMessages(c, http.StatusMethodNotAllowed, "Invalid HTTP method")
		return
	}

	var userEdit request.UserEdit

	userId, err := strconv.Atoi(c.Param("userId"))

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "User ID must be an integer", err)
		c.Abort()
		return
	}

	err = c.BindJSON(&userEdit)

	if err != nil {
		utils.ErrorResponseWithoutMessages(c, http.StatusBadRequest, "Error binding the user data")
		return
	}

	// why pointer to user sign up? kenapa juga to use case di bind ke pointer
	user, errRepo := controller.usecase.EditUser(*userEdit.ToUsecase(), userId)

	if errRepo != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error in repo", errRepo)
		return
	}

	utils.SuccessResponse(c, user, []string{"User successfully edited"})
}

func (controller *UserController) DeleteUser(c *gin.Context) {
	if c.Request.Method != "DELETE" {
		utils.ErrorResponseWithoutMessages(c, http.StatusMethodNotAllowed, "Invalid HTTP method")
		return
	}

	userId, err := strconv.Atoi(c.Param("userId"))

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "User ID must be an integer", err)
		c.Abort()
		return
	}

	user, errRepo := controller.usecase.DeleteUser(userId)

	if errRepo != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error in repo", errRepo)
		return
	}

	utils.SuccessResponse(c, user, []string{"User successfully deleted"})
}

func (controller *UserController) GetUser(c *gin.Context) {
	if c.Request.Method != "GET" {
		utils.ErrorResponseWithoutMessages(c, http.StatusMethodNotAllowed, "Invalid HTTP method")
		return
	}

	userId, err := strconv.Atoi(c.Param("userId"))

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "User ID must be an integer", err)
		c.Abort()
		return
	}

	user, errRepo := controller.usecase.GetUser(userId)

	if errRepo != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Error in repo", errRepo)
		return
	}

	utils.SuccessResponse(c, user, []string{"Successfully get user"})
}
