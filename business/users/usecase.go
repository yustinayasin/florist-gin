package users

import (
	"errors"
	"florist-gin/helpers"
	"fmt"
)

type GeneratorToken interface {
	GenerateToken(userId int) string
}

type UserUseCase struct {
	Repo UserRepoInterface
	Jwt  GeneratorToken
}

func NewUseCase(userRepo UserRepoInterface, tokenGenerator GeneratorToken) UserUseCaseInterface {
	return &UserUseCase{
		Repo: userRepo,
		Jwt:  tokenGenerator,
	}
}

func (userUseCase *UserUseCase) SignUp(user User) (User, error) {
	if user.Name == "" {
		return User{}, errors.New("name cannot be empty")
	}

	if user.Email == "" {
		return User{}, errors.New("email cannot be empty")
	}

	if user.Password == "" {
		return User{}, errors.New("password cannot be empty")
	}

	user.TypeId = 2

	hash, _ := helpers.HashPassword(user.Password)

	user.Password = hash

	userRepo, err := userUseCase.Repo.SignUp(user)

	if err != nil {
		return User{}, err
	}

	return userRepo, nil
}

func (userUseCase *UserUseCase) Login(user User) (User, error) {
	if user.Email == "" {
		return User{}, errors.New("email cannot be empty")
	}

	if user.Password == "" {
		return User{}, errors.New("password cannot be empty")
	}

	userRepo, err := userUseCase.Repo.Login(user)

	fmt.Println(err)

	if err != nil {
		return User{}, err
	}

	match := helpers.CheckPasswordHash(user.Password, userRepo.Password)

	if !match {
		return User{}, errors.New("password doesn't match")
	}

	userRepo.Token = userUseCase.Jwt.GenerateToken(userRepo.Id)

	return userRepo, nil
}

func (userUseCase *UserUseCase) EditUser(user User, id int) (User, error) {
	if id == 0 {
		return User{}, errors.New("user ID cannot be empty")
	}

	if user.Name == "" {
		return User{}, errors.New("name cannot be empty")
	}

	if user.Email == "" {
		return User{}, errors.New("email cannot be empty")
	}

	if user.Password == "" {
		return User{}, errors.New("password cannot be empty")
	}

	if user.Address == "" {
		return User{}, errors.New("address cannot be empty")
	}

	if user.PostalCode == "" {
		return User{}, errors.New("postal code cannot be empty")
	}

	if user.PhoneNumber == "" {
		return User{}, errors.New("phone number cannot be empty")
	}

	hash, _ := helpers.HashPassword(user.Password)
	user.Password = hash

	userRepo, err := userUseCase.Repo.EditUser(user, id)

	if err != nil {
		return User{}, err
	}

	return userRepo, nil
}

func (userUseCase *UserUseCase) DeleteUser(id int) (User, error) {
	if id == 0 {
		return User{}, errors.New("user ID cannot be empty")
	}

	userRepo, err := userUseCase.Repo.DeleteUser(id)

	if err != nil {
		return User{}, err
	}

	return userRepo, nil
}

func (userUseCase *UserUseCase) GetUser(id int) (User, error) {
	if id == 0 {
		return User{}, errors.New("user ID cannot be empty")
	}

	userRepo, err := userUseCase.Repo.GetUser(id)

	if err != nil {
		return User{}, err
	}

	return userRepo, nil
}
