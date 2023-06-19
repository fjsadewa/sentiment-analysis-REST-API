package repository

import (
	"sentiment/models"

	"gorm.io/gorm"
)

//CommentRepository adalah interface yang mendefinisikan operasi-operasi yang dapat dilakukan pada entitas Comment.
type CommentRepository interface {
	InsertComment(comments models.Comments) models.Comments
	InsertToSentimentAnalysis(sentimentAnalysis models.SentimentAnalysis) models.SentimentAnalysis
	UpdateComment(comments models.Comments) models.Comments
	GetAllComment() []models.Comments
	GetCommentByID(id uint64) models.Comments
	DeleteComment(id uint64) error
}

//commentConnection adalah struct yang mengimplementasikan interface CommentRepository.
type commentConnection struct {
	connection *gorm.DB
}

//NewCommentRepository digunakan untuk membuat instance baru dari CommentRepository.
func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentConnection{
		connection: db,
	}
}

//InsertComment digunakan untuk menyimpan data komentar ke dalam database.
func (db *commentConnection) InsertComment(comments models.Comments) models.Comments {
	db.connection.Save(&comments)
	return comments
}

//UpdateComment digunakan untuk memperbarui data komentar yang ada dalam database.
func (db *commentConnection) UpdateComment(comments models.Comments) models.Comments {
	db.connection.Save(&comments)
	return comments
}

//InsertToSentimentAnalysis digunakan untuk menyimpan data analisis sentimen ke dalam database.
func (db *commentConnection) InsertToSentimentAnalysis(sentimentAnalysis models.SentimentAnalysis) models.SentimentAnalysis {
	db.connection.Save(&sentimentAnalysis)
	return sentimentAnalysis
}

//GetAllComment digunakan untuk mengambil semua data komentar dari database.
func (db *commentConnection) GetAllComment() []models.Comments {
	var comments []models.Comments
	db.connection.Preload("Course").Preload("Lecturer").Find(&comments)
	return comments
}

//GetCommentByID digunakan untuk mengambil data komentar berdasarkan ID dari database.
func (db *commentConnection) GetCommentByID(id uint64) models.Comments {
	var comment models.Comments
	db.connection.Preload("Course").Preload("Lecturer").Find(&comment, id)
	return comment
}

//DeleteComment digunakan untuk menghapus data komentar dari database berdasarkan ID.
func (db *commentConnection) DeleteComment(id uint64) error {
	comment := models.Comments{}
	result := db.connection.Delete(&comment, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

/*
 * Keterangan:
 *
 * Package repository berisi definisi repository yang bertanggung jawab untuk mengakses data dari database.
 * Interface CommentRepository adalah kontrak yang mendefinisikan operasi-operasi yang dapat dilakukan pada entitas Comment.
 * Struct commentConnection mengimplementasikan interface CommentRepository dan mengakses database menggunakan objek *gorm.DB.
 * Method-method pada commentConnection digunakan untuk menjalankan operasi-operasi pada entitas Comment.
 * NewCommentRepository digunakan untuk membuat instance baru dari CommentRepository dengan menginisialisasi commentConnection.
 *
 * Pastikan untuk menggunakan repository ini sesuai dengan kebutuhan dan struktur data yang sesuai.
 * Anda dapat menyertakan komentar ini dalam bahasa Indonesia pada kode yang relevan.
 */