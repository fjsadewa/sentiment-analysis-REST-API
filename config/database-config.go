package config

import (
	"fmt"
	"os"
	"sentiment/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
Function SetupDatabaseConnection berfungsi untuk mengatur
koneksi database yang akan digunakan
*/
func SetupDatabaseConnection() *gorm.DB {
	//Memanggil file konfigurasi env
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	//Memmanggil value pengaturan koneksi dari env
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	//Membangun string DSN (Data Source Name) untuk menghubungkan ke basis data MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)

	//Membuka koneksi basis data menggunakan string DSN
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}
	//Melakukan migrasi otomatis tabel-tabel terkait model-model yang diberikan
	db.AutoMigrate(&models.User{}, &models.Comments{}, &models.Courses{}, &models.SentimentAnalysis{}, &models.Lecturers{}, &models.StudyPrograms{})
	//Mengembalikan objek db yang merupakan koneksi basis data yang telah dibuka
	return db
}

/*
Function CloseDatabaseConnection berfungsi untuk menutup koneksi database yang telah dibuka/tersambung
*/
func CloseDatabaseConnection(db *gorm.DB) {
	//Memperoleh objek dbSQL yang mewakili koneksi basis data
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	//Menutup konesi database
	dbSQL.Close()
}
