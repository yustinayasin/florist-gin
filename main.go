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
	cartsProductsDB "florist-gin/drivers/databases/cartsproducts"
	categoryRepo "florist-gin/drivers/databases/categories"
	cartsProductsRepo "florist-gin/drivers/repositories/cartsproducts"

	userUsecase "florist-gin/business/users"
	userController "florist-gin/controller/users"
	userDB "florist-gin/drivers/databases/users"
	userRepo "florist-gin/drivers/repositories/users"

	productUsecase "florist-gin/business/products"
	productController "florist-gin/controller/products"
	productDB "florist-gin/drivers/databases/products"
	productRepo "florist-gin/drivers/repositories/products"

	cartUsecase "florist-gin/business/carts"
	cartController "florist-gin/controller/carts"
	cartDB "florist-gin/drivers/databases/carts"
	cartRepo "florist-gin/drivers/repositories/carts"

	orderUsecase "florist-gin/business/orders"
	orderController "florist-gin/controller/orders"
	orderDB "florist-gin/drivers/databases/orders"
	orderRepo "florist-gin/drivers/repositories/orders"

	ordersProductsUsecase "florist-gin/business/ordersproducts"
	ordersProductsController "florist-gin/controller/ordersproducts"
	ordersProductsDB "florist-gin/drivers/databases/ordersproducts"
	ordersProductsRepo "florist-gin/drivers/repositories/ordersproducts"
)

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&userDB.User{},
		&categoryRepo.Category{},
		&cartDB.Cart{},
		&productDB.Product{},
		&cartsProductsDB.CartsProducts{},
		&orderDB.Order{},
		&ordersProductsDB.OrdersProducts{},
	)
}

func main() {
	err := godotenv.Load("/src/.env")

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
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatalln("Error initialize minio client: ", err)
	}

	userRepoInterface := userRepo.NewUserRepository(db, cartRepo.CartRepository{Db: db})
	userUseCaseInterface := userUsecase.NewUseCase(userRepoInterface, jwtConf)
	userControllerInterface := userController.NewUserController(userUseCaseInterface)

	productRepoInterface := productRepo.NewProductRepository(db, minioClient)
	productUseCaseInterface := productUsecase.NewUseCase(productRepoInterface)
	productControllerInterface := productController.NewProductController(productUseCaseInterface)

	cartRepoInterface := cartRepo.NewCartRepository(db, minioClient)
	cartUseCaseInterface := cartUsecase.NewUseCase(cartRepoInterface)
	cartControllerInterface := cartController.NewCartController(cartUseCaseInterface)

	cartsProductsDBInterface := cartsProductsRepo.NewCartsProductsRepository(db, cartRepo.CartRepository{Db: db}, productRepo.ProductRepository{Db: db})
	cartsProductsUseCaseInterface := cartsProductsUsecase.NewUseCase(cartsProductsDBInterface)
	cartsProductsControllerInterface := cartsProductsController.NewCartsProductsController(cartsProductsUseCaseInterface)

	orderRepoInterface := orderRepo.NewOrderRepository(db)
	orderUseCaseInterface := orderUsecase.NewUseCase(orderRepoInterface)
	orderControllerInterface := orderController.NewOrderController(orderUseCaseInterface)

	ordersProductsRepoInterface := ordersProductsRepo.NewOrdersProductsRepository(db)
	ordersProductsUseCaseInterface := ordersProductsUsecase.NewUseCase(ordersProductsRepoInterface)
	ordersProductsControllerInterface := ordersProductsController.NewOrdersProductsController(ordersProductsUseCaseInterface)

	routesInit := routes.RouteControllerList{
		UserController:           *userControllerInterface,
		ProductController:        *productControllerInterface,
		CartController:           *cartControllerInterface,
		CartsProductsController:  *cartsProductsControllerInterface,
		OrderController:          *orderControllerInterface,
		OrdersProductsController: *ordersProductsControllerInterface,
		JWTConfig:                &jwtConf,
	}

	routesInit.RouteRegister(userRepoInterface)
}
