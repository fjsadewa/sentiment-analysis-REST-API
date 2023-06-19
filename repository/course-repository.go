package repository

import (
	"sentiment/models"

	"gorm.io/gorm"
)

//CourseRepository adalah interface yang mendefinisikan operasi-operasi yang dapat dilakukan pada entitas Course.
type CourseRepository interface {
	InsertCourse(courses models.Courses) models.Courses
	UpdateCourse(courses models.Courses) models.Courses
	GetAllCourse() []models.Courses
	GetCourseByID(id uint64) models.Courses
	DeleteCourse(id uint64) error
}

//courseConnection adalah struct yang mengimplementasikan interface CourseRepository.
type courseConnection struct {
	connection *gorm.DB
}

//NewCourseRepository digunakan untuk membuat instance baru dari CourseRepository.
func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseConnection{
		connection: db,
	}
}

//InsertCourse digunakan untuk menyimpan data kursus ke dalam database.
func (db *courseConnection) InsertCourse(courses models.Courses) models.Courses {
	db.connection.Save(&courses)
	return courses
}

//UpdateCourse digunakan untuk memperbarui data kursus yang ada dalam database.
func (db *courseConnection) UpdateCourse(courses models.Courses) models.Courses {
	db.connection.Save(&courses)
	return courses
}

//GetAllCourse digunakan untuk mengambil semua data kursus dari database.
func (db *courseConnection) GetAllCourse() []models.Courses {
	var courses []models.Courses
	db.connection.Preload("StudyProgram").Find(&courses)
	return courses
}

//GetCourseByID digunakan untuk mengambil data kursus berdasarkan ID dari database.
func (db *courseConnection) GetCourseByID(id uint64) models.Courses {
	var course models.Courses
	db.connection.Preload("StudyProgram").Find(&course, id)
	return course
}

//DeleteCourse digunakan untuk menghapus data kursus dari database berdasarkan ID.
func (db *courseConnection) DeleteCourse(id uint64) error {
	course := models.Courses{}
	result := db.connection.Delete(&course, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
