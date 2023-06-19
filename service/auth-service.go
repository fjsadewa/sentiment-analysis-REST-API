package service

import (
	"log"
	"sentiment/dto"
	"sentiment/models"
	"sentiment/repository"

	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

//AuthService adalah interface yang mendefinisikan operasi-operasi yang terkait dengan otentikasi dan pengguna.
type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.RegisterDTO) models.User
	FindByEmail(email string) models.User
	IsDuplicateEmail(email string) bool
}

//authService adalah struct yang mengimplementasikan interface AuthService.
type authService struct {
	userRepository repository.UserRepository
}

//NewAuthService digunakan untuk membuat instance baru dari AuthService.
func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}

/*
VerifyCredential digunakan untuk memverifikasi kredensial pengguna (email dan password).
Jika kredensial valid, method akan mengembalikan objek user yang sesuai.
Jika kredensial tidak valid, method akan mengembalikan false.
*/
func (service *authService) VerifyCredential(email string, password string) interface{} {
	res := service.userRepository.VerifyCredential(email, password)
	if v, ok := res.(models.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return false
}

//CreateUser digunakan untuk membuat user baru berdasarkan data dari DTO.
func (service *authService) CreateUser(user dto.RegisterDTO) models.User {
	userToCreate := models.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.userRepository.InsertUser(userToCreate)
	return res
}

//FindByEmail digunakan untuk mencari data user berdasarkan email dari repository.
func (service *authService) FindByEmail(email string) models.User {
	return service.userRepository.FindByEmail(email)
}

/*
IsDuplicateEmail digunakan untuk memeriksa apakah email sudah ada dalam database.
Jika email sudah ada, method akan mengembalikan true.
Jika email belum ada, method akan mengembalikan false.
*/
func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

/*
comparePassword digunakan untuk membandingkan password yang di-hash dengan password mentah.
Jika password sama, method akan mengembalikan true.
Jika password berbeda, method akan mengembalikan false.
*/
func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
