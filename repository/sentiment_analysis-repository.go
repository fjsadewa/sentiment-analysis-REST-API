package repository

import (
	"sentiment/models"

	"gorm.io/gorm"
)

//SentimenAnalysisRepository adalah interface yang mendefinisikan operasi-operasi yang dapat dilakukan pada entitas SentimentAnalysis.
type SentimenAnalysisRepository interface {
	InsertSentimenAnalysis(sentimentAnalysis models.SentimentAnalysis) models.SentimentAnalysis
	UpdateSentimenAnalysis(sentimentAnalysis models.SentimentAnalysis) models.SentimentAnalysis
	GetAllSentimenAnalysis() []models.SentimentAnalysis
	GetSentimenAnalysisByID(id uint64) models.SentimentAnalysis
	DeleteSentimenAnalysis(id uint64) error
}

//sentimenAnalysisConnection adalah struct yang mengimplementasikan interface SentimenAnalysisRepository.
type sentimenAnalysisConnection struct {
	connection *gorm.DB
}

//NewSentimenAnalysisRepository digunakan untuk membuat instance baru dari SentimenAnalysisRepository.
func NewSentimenAnalysisRepository(db *gorm.DB) SentimenAnalysisRepository {
	return &sentimenAnalysisConnection{
		connection: db,
	}
}

//InsertSentimenAnalysis digunakan untuk menyimpan data analisis sentimen ke dalam database.
func (db *sentimenAnalysisConnection) InsertSentimenAnalysis(sentimentAnalysis models.SentimentAnalysis) models.SentimentAnalysis {
	db.connection.Save(&sentimentAnalysis)
	return sentimentAnalysis
}

//UpdateSentimenAnalysis digunakan untuk memperbarui data analisis sentimen yang ada dalam database.
func (db *sentimenAnalysisConnection) UpdateSentimenAnalysis(sentimentAnalysis models.SentimentAnalysis) models.SentimentAnalysis {
	db.connection.Save(&sentimentAnalysis)
	return sentimentAnalysis
}

//GetAllSentimenAnalysis digunakan untuk mengambil semua data analisis sentimen dari database.
func (db *sentimenAnalysisConnection) GetAllSentimenAnalysis() []models.SentimentAnalysis {
	var sentimentAnalysis []models.SentimentAnalysis
	db.connection.Preload("Comment").Preload("Comment.Course").Preload("Comment.Lecturer").Preload("Comment.Lecturer.StudyProgram").Find(&sentimentAnalysis)
	return sentimentAnalysis
}

//GetSentimenAnalysisByID digunakan untuk mengambil data analisis sentimen berdasarkan ID dari database.
func (db *sentimenAnalysisConnection) GetSentimenAnalysisByID(id uint64) models.SentimentAnalysis {
	var sentimenAnalysis models.SentimentAnalysis
	db.connection.Preload("Comment").Preload("Comment.Course").Preload("Comment.Lecturer").Preload("Comment.Lecturer.StudyProgram").Find(&sentimenAnalysis, id)
	return sentimenAnalysis
}

//DeleteSentimenAnalysis digunakan untuk menghapus data analisis sentimen dari database berdasarkan ID.
func (db *sentimenAnalysisConnection) DeleteSentimenAnalysis(id uint64) error {
	sentimenAnalysis := models.SentimentAnalysis{}
	result := db.connection.Delete(&sentimenAnalysis, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

/*
 * Keterangan:
 *
 * Package repository berisi definisi repository yang bertanggung jawab untuk mengakses data dari database.
 * Interface SentimenAnalysisRepository adalah kontrak yang mendefinisikan operasi-operasi yang dapat dilakukan pada entitas SentimentAnalysis.
 * Struct sentimenAnalysisConnection mengimplementasikan interface SentimenAnalysisRepository dan mengakses database menggunakan objek *gorm.DB.
 * Method-method pada sentimenAnalysisConnection digunakan untuk menjalankan operasi-operasi pada entitas SentimentAnalysis.
 * NewSentimenAnalysisRepository digunakan untuk membuat instance baru dari SentimenAnalysisRepository dengan menginisialisasi sentimenAnalysisConnection.
 *
 * Pastikan untuk menggunakan repository ini sesuai dengan kebutuhan dan struktur data yang sesuai.
 * Anda dapat menyertakan komentar ini dalam bahasa Indonesia pada kode yang relevan.
 */
