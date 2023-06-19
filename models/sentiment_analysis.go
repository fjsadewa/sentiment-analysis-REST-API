package models

import "time"

//Lecturers merupakan model untuk tabel dosen
type SentimentAnalysis struct {
	ID              uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	SentimentType   string    `gorm:"type:text" json:"sentiment_type"`
	ConfidenceLevel float64   `gorm:"not null" json:"confidence_level"`
	CommentID       uint      `gorm:"not null" json:"-"`
	Comment         Comments  `gorm:"foreignKey:CommentID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"comment_id"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"update_at,omitempty"`
	DeletedAt       time.Time `gorm:"index"  json:"deleted_at,omitempty"`
}

//TableName digunakan untuk mengembalikan nama tabel yang sesuai
func (SentimentAnalysis) TableName() string {
	return "sentiment_analysis"
}

/*
 * Keterangan:
 *
 * Package models berisi definisi model data yang sesuai dengan struktur tabel pada database.
 * Struct SentimentAnalysis merepresentasikan tabel analisis sentimen pada database.
 *
 * Properti-properti pada struct SentimentAnalysis adalah sebagai berikut:
 * - ID: ID unik analisis sentimen (kolom primaryKey).
 * - SentimentType: Jenis sentimen pada analisis.
 * - ConfidenceLevel: Tingkat kepercayaan pada analisis sentimen (kolom not null).
 * - CommentID: ID komentar yang terkait dengan analisis (kolom not null).
 * - Comment: Objek Comments yang merupakan relasi ke komentar yang terkait (foreign key).
 * - CreatedAt: Waktu pembuatan analisis sentimen.
 * - UpdatedAt: Waktu terakhir analisis sentimen diperbarui.
 * - DeletedAt: Waktu penghapusan analisis sentimen (kolom index, akan bernilai NULL jika belum dihapus).
 *
 * Method TableName digunakan untuk mengembalikan nama tabel yang sesuai dengan struktur.
 *
 * Pastikan untuk menggunakan model ini sesuai dengan struktur tabel analisis sentimen pada database.
 * Anda dapat menyertakan komentar ini dalam bahasa Indonesia pada kode yang relevan.
 */