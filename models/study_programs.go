package models

import "time"

//StudyPrograms merupakan model untuk tabel program studi
type StudyPrograms struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Code      string    `gorm:"type:text" json:"code"`
	Name      string    `gorm:"type:text" json:"name"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"update_at,omitempty"`
	DeletedAt time.Time `gorm:"index"  json:"deleted_at,omitempty"`
}

//TableName digunakan untuk mengembalikan nama tabel yang sesuai
func (StudyPrograms) TableName() string {
	return "study_programs"
}

/*
 * Keterangan:
 *
 * Package models berisi definisi model data yang sesuai dengan struktur tabel pada database.
 * Struct StudyPrograms merepresentasikan tabel program studi pada database.
 *
 * Properti-properti pada struct StudyPrograms adalah sebagai berikut:
 * - ID: ID unik program studi (kolom primaryKey).
 * - Code: Kode program studi.
 * - Name: Nama program studi.
 * - CreatedAt: Waktu pembuatan program studi.
 * - UpdatedAt: Waktu terakhir program studi diperbarui.
 * - DeletedAt: Waktu penghapusan program studi (kolom index, akan bernilai NULL jika belum dihapus).
 *
 * Method TableName digunakan untuk mengembalikan nama tabel yang sesuai dengan struktur.
 *
 * Pastikan untuk menggunakan model ini sesuai dengan struktur tabel program studi pada database.
 * Anda dapat menyertakan komentar ini dalam bahasa Indonesia pada kode yang relevan.
 */