package repository

import (
	"sentiment/models"

	"gorm.io/gorm"
)

//LecturerRepository adalah interface yang mendefinisikan operasi-operasi yang dapat dilakukan pada entitas Lecturer.
type LecturerRepository interface {
	InsertLecturer(lecturers models.Lecturers) models.Lecturers
	UpdateLecturer(lecturers models.Lecturers) models.Lecturers
	GetAllLecturer() []models.Lecturers
	GetLecturerByID(id uint64) models.Lecturers
	DeleteLecturer(id uint64) error
}

//lecturerConnection adalah struct yang mengimplementasikan interface LecturerRepository.
type lecturerConnection struct {
	connection *gorm.DB
}

//NewLecturerRepository digunakan untuk membuat instance baru dari LecturerRepository.
func NewLecturerRepository(db *gorm.DB) LecturerRepository {
	return &lecturerConnection{
		connection: db,
	}
}

//InsertLecturer digunakan untuk menyimpan data dosen ke dalam database.
func (db *lecturerConnection) InsertLecturer(lecturers models.Lecturers) models.Lecturers {
	db.connection.Save(&lecturers)
	return lecturers
}

//UpdateLecturer digunakan untuk memperbarui data dosen yang ada dalam database.
func (db *lecturerConnection) UpdateLecturer(lecturers models.Lecturers) models.Lecturers {
	db.connection.Save(&lecturers)
	return lecturers
}

//GetAllLecturer digunakan untuk mengambil semua data dosen dari database.
func (db *lecturerConnection) GetAllLecturer() []models.Lecturers {
	var lecturers []models.Lecturers
	db.connection.Preload("StudyProgram").Find(&lecturers)
	return lecturers
}

//GetLecturerByID digunakan untuk mengambil data dosen berdasarkan ID dari database.
func (db *lecturerConnection) GetLecturerByID(id uint64) models.Lecturers {
	var lecturer models.Lecturers
	db.connection.Preload("StudyProgram").Find(&lecturer, id)
	return lecturer
}

//DeleteLecturer digunakan untuk menghapus data dosen dari database berdasarkan ID.
func (db *lecturerConnection) DeleteLecturer(id uint64) error {
	lecturer := models.Lecturers{}
	result := db.connection.Delete(&lecturer, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

/*
 * Keterangan:
 *
 * Package repository berisi definisi repository yang bertanggung jawab untuk mengakses data dari database.
 * Interface LecturerRepository adalah kontrak yang mendefinisikan operasi-operasi yang dapat dilakukan pada entitas Lecturer.
 * Struct lecturerConnection mengimplementasikan interface LecturerRepository dan mengakses database menggunakan objek *gorm.DB.
 * Method-method pada lecturerConnection digunakan untuk menjalankan operasi-operasi pada entitas Lecturer.
 * NewLecturerRepository digunakan untuk membuat instance baru dari LecturerRepository dengan menginisialisasi lecturerConnection.
 *
 * Pastikan untuk menggunakan repository ini sesuai dengan kebutuhan dan struktur data yang sesuai.
 * Anda dapat menyertakan komentar ini dalam bahasa Indonesia pada kode yang relevan.
 */