package service

import (
	"log"
	"sentiment/dto"
	"sentiment/models"
	"sentiment/repository"

	"github.com/mashingan/smapping"
)

/*
UserService adalah sebuah interface yang mendefinisikan method-method
untuk menangani operasi terkait pengguna.
*/
type UserService interface {
	Update(user dto.UserUpdateDTO) models.User
	Profile(userID string) models.User
}

//userService adalah sebuah struct yang mengimplementasikan interface UserService.
type userService struct {
	userRepository repository.UserRepository
}

//NewUserService membuat sebuah instance baru dari userService.
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

//Update mengubah data pengguna yang ada dalam database.
func (service *userService) Update(user dto.UserUpdateDTO) models.User {
	userToUpdate := models.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updatedUser := service.userRepository.UpdateUser((userToUpdate))
	return updatedUser
}

//Profile mengambil data pengguna berdasarkan ID pengguna dari database.
func (service *userService) Profile(userID string) models.User {
	return service.userRepository.ProfileUser(userID)
}
