package main

import (
	middleware "florist-gin/app/middleware"
	"florist-gin/app/routes"
	"florist-gin/helpers"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/gorm"

	cartsProductsUsecase "florist-gin/business/cartsproducts"
	cartsProductsController "florist-gin/controller/cartsproducts"
	cartsProductsRepo "florist-gin/drivers/databases/cartsproducts"
	categoryRepo "florist-gin/drivers/databases/categories"

	userUsecase "florist-gin/business/users"
	userController "florist-gin/controller/users"
	userRepo "florist-gin/drivers/databases/users"

	productUsecase "florist-gin/business/products"
	productController "florist-gin/controller/products"
	productRepo "florist-gin/drivers/databases/products"

	cartUsecase "florist-gin/business/carts"
	cartController "florist-gin/controller/carts"
	cartRepo "florist-gin/drivers/databases/carts"
)

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&userRepo.User{},
		&categoryRepo.Category{},
		&cartRepo.Cart{},
		&productRepo.Product{},
		&cartsProductsRepo.CartsProducts{},
	)
}

func main() {
	err := godotenv.Load("config.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	secretJWT := os.Getenv("JWT_SECRET")
	expiresDuration := os.Getenv("JWT_EXPIRED")
	expiresDurationInt, _ := strconv.Atoi(expiresDuration)

	db, err := helpers.NewDatabase()

	if err != nil {
		log.Fatal(err)
	}

	dbMigrate(db)

	jwtConf := middleware.ConfigJWT{
		SecretJWT:       secretJWT,
		ExpiresDuration: expiresDurationInt,
	}

	userRepoInterface := userRepo.NewUserRepository(db, cartRepo.CartRepository{Db: db})
	userUseCaseInterface := userUsecase.NewUseCase(userRepoInterface, jwtConf)
	userControllerInterface := userController.NewUserController(userUseCaseInterface)

	productRepoInterface := productRepo.NewProductRepository(db)
	productUseCaseInterface := productUsecase.NewUseCase(productRepoInterface)
	productControllerInterface := productController.NewProductController(productUseCaseInterface)

	cartRepoInterface := cartRepo.NewCartRepository(db)
	cartUseCaseInterface := cartUsecase.NewUseCase(cartRepoInterface)
	cartControllerInterface := cartController.NewCartController(cartUseCaseInterface)

	cartsProductsRepoInterface := cartsProductsRepo.NewCartsProductsRepository(db)
	cartsProductsUseCaseInterface := cartsProductsUsecase.NewUseCase(cartsProductsRepoInterface)
	cartsProductsControllerInterface := cartsProductsController.NewCartsProductsController(cartsProductsUseCaseInterface)

	routesInit := routes.RouteControllerList{
		UserController:          *userControllerInterface,
		ProductController:       *productControllerInterface,
		CartController:          *cartControllerInterface,
		CartsProductsController: *cartsProductsControllerInterface,
		JWTConfig:               &jwtConf,
	}

	routesInit.RouteRegister(userRepoInterface)
}
