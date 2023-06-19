package repository

import (
	"log"
	"sentiment/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//UserRepository adalah interface yang mendefinisikan operasi-operasi yang dapat dilakukan pada entitas User.
type UserRepository interface {
	InsertUser(user models.User) models.User
	UpdateUser(user models.User) models.User
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) models.User
	ProfileUser(userID string) models.User
}

//userConnection adalah struct yang mengimplementasikan interface UserRepository.
type userConnection struct {
	connection *gorm.DB
}

//NewUserRepository digunakan untuk membuat instance baru dari UserRepository.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

/*
InsertUser digunakan untuk menyimpan data user ke dalam database.
Password user akan di-hash sebelum disimpan.
*/
func (db *userConnection) InsertUser(user models.User) models.User {
	user.Password = hashAndSalt([]byte(user.Password))
	db.connection.Save(&user)
	return user
}

/*
UpdateUser digunakan untuk memperbarui data user yang ada dalam database.
Jika password user diisi, password akan di-hash sebelum disimpan.
Jika password tidak diisi, password akan tetap menggunakan password sebelumnya.
*/
func (db *userConnection) UpdateUser(user models.User) models.User {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser models.User
		db.connection.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}
	db.connection.Save(&user)
	return user
}

/*
VerifyCredential digunakan untuk memverifikasi kredensial pengguna (email dan password).
Jika kredensial valid, method akan mengembalikan objek user yang sesuai.
Jika kredensial tidak valid, method akan mengembalikan nil.
*/
func (db *userConnection) VerifyCredential(email string, password string) interface{} {
	var user models.User
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

/*
IsDuplicateEmail digunakan untuk memeriksa apakah email sudah ada dalam database.
Jika email sudah ada, method akan mengembalikan objek *gorm.DB yang berisi data user dengan email yang sama.
Jika email belum ada, method akan mengembalikan nil.
*/
func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user models.User
	return db.connection.Where("email = ?", email).Take(&user)
}

//FindByEmail digunakan untuk mencari data user berdasarkan email dari database.
func (db *userConnection) FindByEmail(email string) models.User {
	var user models.User
	db.connection.Where("email = ?", email).Take(&user)
	return user
}

//ProfileUser digunakan untuk mengambil data profil user berdasarkan ID dari database.
func (db *userConnection) ProfileUser(userID string) models.User {
	var user models.User
	db.connection.Find(&user, userID)
	return user
}

//hashAndSalt digunakan untuk melakukan hash dan salt pada password.
func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to has a password")
	}
	return string(hash)
}
