package main

import (
	middleware "florist-gin/app/middleware"
	"florist-gin/app/routes"
	"florist-gin/helpers"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
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

	orderUsecase "florist-gin/business/orders"
	orderController "florist-gin/controller/orders"
	orderRepo "florist-gin/drivers/databases/orders"
)

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&userRepo.User{},
		&categoryRepo.Category{},
		&cartRepo.Cart{},
		&productRepo.Product{},
		&cartsProductsRepo.CartsProducts{},
		&orderRepo.Order{},
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

	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY")
	secretAccessKey := os.Getenv("MINIO_SECRET_KEY")
	useSSL := true

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	userRepoInterface := userRepo.NewUserRepository(db, cartRepo.CartRepository{Db: db})
	userUseCaseInterface := userUsecase.NewUseCase(userRepoInterface, jwtConf)
	userControllerInterface := userController.NewUserController(userUseCaseInterface)

	productRepoInterface := productRepo.NewProductRepository(db, minioClient)
	productUseCaseInterface := productUsecase.NewUseCase(productRepoInterface)
	productControllerInterface := productController.NewProductController(productUseCaseInterface)

	cartRepoInterface := cartRepo.NewCartRepository(db)
	cartUseCaseInterface := cartUsecase.NewUseCase(cartRepoInterface)
	cartControllerInterface := cartController.NewCartController(cartUseCaseInterface)

	cartsProductsRepoInterface := cartsProductsRepo.NewCartsProductsRepository(db)
	cartsProductsUseCaseInterface := cartsProductsUsecase.NewUseCase(cartsProductsRepoInterface)
	cartsProductsControllerInterface := cartsProductsController.NewCartsProductsController(cartsProductsUseCaseInterface)

	orderRepoInterface := orderRepo.NewOrderRepository(db)
	orderUseCaseInterface := orderUsecase.NewUseCase(orderRepoInterface)
	orderControllerInterface := orderController.NewOrderController(orderUseCaseInterface)

	routesInit := routes.RouteControllerList{
		UserController:          *userControllerInterface,
		ProductController:       *productControllerInterface,
		CartController:          *cartControllerInterface,
		CartsProductsController: *cartsProductsControllerInterface,
		OrderController:         *orderControllerInterface,
		JWTConfig:               &jwtConf,
	}

	routesInit.RouteRegister(userRepoInterface)
}
