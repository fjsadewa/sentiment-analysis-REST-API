package repository

import (
	"errors"
	"sentiment/models"

	"gorm.io/gorm"
)

//StudyProgramRepository adalah interface yang mendefinisikan operasi-operasi yang dapat dilakukan pada entitas StudyPrograms.
type StudyProgramRepository interface {
	InsertStudyProgram(studyProgram models.StudyPrograms) models.StudyPrograms
	UpdateStudyProgram(studyProgram models.StudyPrograms) models.StudyPrograms
	DeleteStudyProgram(id uint64) error
	AllStudyProgram() []models.StudyPrograms
	FindStudyProgramByID(id uint64) (models.StudyPrograms, error)
	FindStudyProgramByCode(studyProgramName string) models.StudyPrograms
}

//studyProgramConnection adalah struct yang mengimplementasikan interface StudyProgramRepository.
type studyProgramConnection struct {
	connection *gorm.DB
}

//NewStudyProgramRepository digunakan untuk membuat instance baru dari StudyProgramRepository.
func NewStudyProgramRepository(db *gorm.DB) StudyProgramRepository {
	return &studyProgramConnection{
		connection: db,
	}
}

//InsertStudyProgram digunakan untuk menyimpan data program studi ke dalam database.
func (db *studyProgramConnection) InsertStudyProgram(studyProgram models.StudyPrograms) models.StudyPrograms {
	db.connection.Save(&studyProgram)
	return studyProgram
}

//UpdateStudyProgram digunakan untuk memperbarui data program studi yang ada dalam database.
func (db *studyProgramConnection) UpdateStudyProgram(studyProgram models.StudyPrograms) models.StudyPrograms {
	db.connection.Save(&studyProgram)
	return studyProgram
}

//DeleteStudyProgram digunakan untuk menghapus data program studi dari database berdasarkan ID.
func (db *studyProgramConnection) DeleteStudyProgram(id uint64) error {
	studyProgram := models.StudyPrograms{}
	result := db.connection.Delete(&studyProgram, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//AllStudyProgram digunakan untuk mengambil semua data program studi dari database.
func (db *studyProgramConnection) AllStudyProgram() []models.StudyPrograms {
	var studyPrograms []models.StudyPrograms
	db.connection.Find(&studyPrograms)
	return studyPrograms
}

//FindStudyProgramByID digunakan untuk mencari data program studi berdasarkan ID dari database.
func (db *studyProgramConnection) FindStudyProgramByID(id uint64) (models.StudyPrograms, error) {
	var studyProgram models.StudyPrograms
	db.connection.Find(&studyProgram, id)
	if studyProgram.ID == 0 {
		return studyProgram, errors.New("study program not found")
	}
	return studyProgram, nil
}

//FindStudyProgramByCode digunakan untuk mencari data program studi berdasarkan kode dari database.
func (db *studyProgramConnection) FindStudyProgramByCode(studyProgramCode string) models.StudyPrograms {
	var studyProgram models.StudyPrograms
	db.connection.Where("code = ?", studyProgramCode).Take(&studyProgram)
	return studyProgram
}

/*
 * Keterangan:
 *
 * Package repository berisi definisi repository yang bertanggung jawab untuk mengakses data dari database.
 * Interface StudyProgramRepository adalah kontrak yang mendefinisikan operasi-operasi yang dapat dilakukan pada entitas StudyPrograms.
 * Struct studyProgramConnection mengimplementasikan interface StudyProgramRepository dan mengakses database menggunakan objek *gorm.DB.
 * Method-method pada studyProgramConnection digunakan untuk menjalankan operasi-operasi pada entitas StudyPrograms.
 * NewStudyProgramRepository digunakan untuk membuat instance baru dari studyProgramConnection.
 * Setiap method pada studyProgramConnection mengakses database menggunakan objek *gorm.DB dan menjalankan operasi-operasi yang sesuai.
 * Jika terjadi kesalahan dalam menjalankan operasi-operasi tersebut, method akan mengembalikan error.
 */