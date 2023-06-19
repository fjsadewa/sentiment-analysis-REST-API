package models

import "time"

//Courses merupakan model untuk tabel kursus
type Courses struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Code      string    `gorm:"type:text" json:"code"`
	Name      string    `gorm:"type:text" json:"name"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"update_at,omitempty"`
	DeletedAt time.Time `gorm:"index"  json:"deleted_at,omitempty"`
}

//TableName digunakan untuk mengembalikan nama tabel yang sesuai
func (Courses) TableName() string {
	return "courses"
}

/*
 * Keterangan:
 *
 * Package models berisi definisi model data yang sesuai dengan struktur tabel pada database.
 * Struct Courses merepresentasikan tabel kursus pada database.
 *
 * Properti-properti pada struct Courses adalah sebagai berikut:
 * - ID: ID unik kursus (kolom primaryKey).
 * - Code: Kode kursus.
 * - Name: Nama kursus.
 * - CreatedAt: Waktu pembuatan kursus.
 * - UpdatedAt: Waktu terakhir kursus diperbarui.
 * - DeletedAt: Waktu penghapusan kursus (kolom index, akan bernilai NULL jika belum dihapus).
 *
 * Method TableName digunakan untuk mengembalikan nama tabel yang sesuai dengan struktur.
 *
 * Pastikan untuk menggunakan model ini sesuai dengan struktur tabel kursus pada database.
 * Anda dapat menyertakan komentar ini dalam bahasa Indonesia pada kode yang relevan.
 */