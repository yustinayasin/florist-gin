package routes

import (
	"florist-gin/app/middleware"
	"florist-gin/business/users"
	cartController "florist-gin/controller/carts"
	cartsProductsController "florist-gin/controller/cartsproducts"
	orderController "florist-gin/controller/orders"
	ordersProductsController "florist-gin/controller/ordersproducts"
	productController "florist-gin/controller/products"
	userController "florist-gin/controller/users"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RouteControllerList struct {
	UserController           userController.UserController
	ProductController        productController.ProductController
	CartController           cartController.CartController
	CartsProductsController  cartsProductsController.CartsProductsController
	OrderController          orderController.OrderController
	OrdersProductsController ordersProductsController.OrdersProductsController
	JWTConfig                *middleware.ConfigJWT
}

func (controller RouteControllerList) RouteRegister(userRepoInterface users.UserRepoInterface) {
	router := gin.Default()

	// CORS configuration
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://poppy-florist.yustinayasin.com"}, // Replace with your frontend origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

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
		product.GET("/:productId", controller.ProductController.GetProductDetail)
		product.GET("/", controller.ProductController.GetAllProduct)
	}

	cart := router.Group("/cart")
	{
		cart.GET("/", middleware.RequireAuth(controller.CartController.GetCart, *controller.JWTConfig, userRepoInterface))
	}

	cartsproducts := router.Group("/cartsproducts")
	{
		cartsproducts.POST("/add", middleware.RequireAuth(controller.CartsProductsController.AddProductToCart, *controller.JWTConfig, userRepoInterface))
		cartsproducts.PUT("/:cartsProductsId", middleware.RequireAuth(controller.CartsProductsController.EditProductFromCart, *controller.JWTConfig, userRepoInterface))
		cartsproducts.DELETE("/:cartsProductsId", middleware.RequireAuth(controller.CartsProductsController.DeleteProductFromCart, *controller.JWTConfig, userRepoInterface))
	}

	order := router.Group("/order")
	{
		order.POST("/add", middleware.RequireAuth(controller.OrderController.AddOrder, *controller.JWTConfig, userRepoInterface))
		order.PUT("/:orderId", middleware.RequireAuth(controller.OrderController.EditOrder, *controller.JWTConfig, userRepoInterface))
		order.DELETE("/:orderId", middleware.RequireAuth(controller.OrderController.DeleteOrder, *controller.JWTConfig, userRepoInterface))
		order.GET("/:orderId", middleware.RequireAuth(controller.OrderController.GetOrderDetail, *controller.JWTConfig, userRepoInterface))
	}

	ordersproducts := router.Group("/ordersproducts")
	{
		ordersproducts.POST("/add", middleware.RequireAuth(controller.OrdersProductsController.AddOrdersProducts, *controller.JWTConfig, userRepoInterface))
		ordersproducts.PUT("/:ordersProductsId", middleware.RequireAuth(controller.OrdersProductsController.EditOrdersProducts, *controller.JWTConfig, userRepoInterface))
		ordersproducts.DELETE("/:ordersProductsId", middleware.RequireAuth(controller.OrdersProductsController.DeleteOrdersProducts, *controller.JWTConfig, userRepoInterface))
		ordersproducts.GET("/:ordersProductsId", middleware.RequireAuth(controller.OrdersProductsController.GetOrdersProductsDetail, *controller.JWTConfig, userRepoInterface))
	}

	port := ":8080"
	err := router.Run(port)
	if err != nil {
		log.Println("Failed to start server:", err)
		os.Exit(1)
	}
}
