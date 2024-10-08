package users

import (
	"errors"
	"florist-gin/business/users"
	"florist-gin/drivers/databases/carts"
	usersDB "florist-gin/drivers/databases/users"
	cartsDB "florist-gin/drivers/repositories/carts"
	"florist-gin/helpers"

	"gorm.io/gorm"
)

type UserRepository struct {
	Db       *gorm.DB
	CartRepo cartsDB.CartRepository
}

func NewUserRepository(database *gorm.DB, cartRepo cartsDB.CartRepository) users.UserRepoInterface {
	return &UserRepository{
		Db:       database,
		CartRepo: cartRepo,
	}
}

func (repo *UserRepository) SignUp(user users.User) (users.User, error) {
	userDB := usersDB.FromUsecase(user)

	result := repo.Db.Create(&userDB)

	if result.Error != nil {
		return users.User{}, result.Error
	}

	cart := carts.Cart{
		UserId: userDB.Id,
	}

	result = repo.CartRepo.Db.Preload("carts").Create(&cart)

	if result.Error != nil {
		return users.User{}, result.Error
	}

	return userDB.ToUsecase(), nil
}

func (repo *UserRepository) Login(user users.User) (users.User, error) {
	userDB := usersDB.FromUsecase(user)

	result := repo.Db.Where("email = ?", userDB.Email).Preload("Type").First(&userDB)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return users.User{}, errors.New("user not found")
		}
		return users.User{}, errors.New("error in database")
	}

	match := helpers.CheckPasswordHash(user.Password, userDB.Password)

	if !match {
		return users.User{}, errors.New("password doesn't match")
	}

	return userDB.ToUsecase(), nil
}

func (repo *UserRepository) EditUser(user users.User, id int) (users.User, error) {
	userDb := usersDB.FromUsecase(user)

	var newUser usersDB.User

	result := repo.Db.Preload("Type").First(&newUser, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return users.User{}, errors.New("User not found")
		}
		return users.User{}, errors.New("error in database")
	}

	newUser.Email = userDb.Email
	newUser.Password = userDb.Password
	newUser.Name = userDb.Name
	newUser.Address = userDb.Address
	newUser.PhoneNumber = userDb.PhoneNumber
	newUser.PostalCode = userDb.PostalCode

	repo.Db.Save(&newUser)
	return newUser.ToUsecase(), nil
}

func (repo *UserRepository) DeleteUser(id int) (users.User, error) {
	var userDb usersDB.User

	resultFind := repo.Db.First(&userDb, id)

	if resultFind.Error != nil {
		return users.User{}, errors.New("user not found")
	}

	result := repo.Db.Delete(&userDb, id)

	if result.Error != nil {
		return users.User{}, errors.New("user not found")
	}

	return userDb.ToUsecase(), nil
}

func (repo *UserRepository) GetUser(id int) (users.User, error) {
	var userDb usersDB.User

	resultFind := repo.Db.First(&userDb, id)

	if resultFind.Error != nil {
		return users.User{}, errors.New("user not found")
	}

	return userDb.ToUsecase(), nil
}
