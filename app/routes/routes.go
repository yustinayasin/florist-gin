package routes

import (
	"florist-gin/app/middleware"
	"florist-gin/business/users"
	cartController "florist-gin/controller/carts"
	cartsProductsController "florist-gin/controller/cartsproducts"
	orderController "florist-gin/controller/orders"
	productController "florist-gin/controller/products"
	userController "florist-gin/controller/users"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type RouteControllerList struct {
	UserController          userController.UserController
	ProductController       productController.ProductController
	CartController          cartController.CartController
	CartsProductsController cartsProductsController.CartsProductsController
	OrderController         orderController.OrderController
	JWTConfig               *middleware.ConfigJWT
}

func (controller RouteControllerList) RouteRegister(userRepoInterface users.UserRepoInterface) {
	router := gin.Default()

	user := router.Group("/user")
	{
		user.POST("/login", controller.UserController.Login)
		user.POST("/signup", controller.UserController.SignUp)
		user.PUT("/:userId", middleware.RequireAuth(controller.UserController.EditUser, *controller.JWTConfig, userRepoInterface))
		user.DELETE("/:userId", middleware.RequireAuth(controller.UserController.DeleteUser, *controller.JWTConfig, userRepoInterface))
		user.GET("/:userId", middleware.RequireAuth(controller.UserController.GetUser, *controller.JWTConfig, userRepoInterface))
	}

	product := router.Group("/product")
	{
		product.POST("/add", middleware.RequireAuthAdmin(controller.ProductController.AddProduct, *controller.JWTConfig, userRepoInterface))
		product.PUT("/:productId", middleware.RequireAuthAdmin(controller.ProductController.EditProduct, *controller.JWTConfig, userRepoInterface))
		product.DELETE("/:productId", middleware.RequireAuthAdmin(controller.ProductController.DeleteProduct, *controller.JWTConfig, userRepoInterface))
		product.GET("/:productId", middleware.RequireAuth(controller.ProductController.GetProductDetail, *controller.JWTConfig, userRepoInterface))
		product.GET("/", middleware.RequireAuth(controller.ProductController.GetAllProduct, *controller.JWTConfig, userRepoInterface))
	}

	cart := router.Group("/cart")
	{
		cart.GET("/:cartId", middleware.RequireAuth(controller.CartController.GetCart, *controller.JWTConfig, userRepoInterface))
	}

	cartsproducts := router.Group("/cartsproducts")
	{
		cartsproducts.POST("/add", middleware.RequireAuth(controller.CartsProductsController.AddProductToCart, *controller.JWTConfig, userRepoInterface))
		cartsproducts.PUT("/:cartsProductsId", middleware.RequireAuth(controller.CartsProductsController.EditProductFromCart, *controller.JWTConfig, userRepoInterface))
		cartsproducts.DELETE("/:cartsProductsId", middleware.RequireAuth(controller.CartsProductsController.DeleteProductFromCart, *controller.JWTConfig, userRepoInterface))
	}

	port := ":8081"
	err := router.Run(port)
	if err != nil {
		log.Println("Failed to start server:", err)
		os.Exit(1)
	}
}
