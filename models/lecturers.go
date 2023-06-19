package models

import "time"

//Lecturers merupakan model untuk tabel dosen
type Lecturers struct {
	ID             uint64        `gorm:"primaryKey;autoIncrement" json:"id"`
	Code           string        `gorm:"type:text" json:"code"`
	Name           string        `gorm:"type:text" json:"name"`
	StudyProgramID uint64        `gorm:"not null" json:"-"`
	StudyProgram   StudyPrograms `gorm:"foreignKey:StudyProgramID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"study_program_id"`
	CreatedAt      time.Time     `json:"created_at,omitempty"`
	UpdatedAt      time.Time     `json:"update_at,omitempty"`
	DeletedAt      time.Time     `gorm:"index"  json:"deleted_at,omitempty"`
}

//TableName digunakan untuk mengembalikan nama tabel yang sesuai
func (Lecturers) TableName() string {
	return "lecturers"
}

/*
 * Keterangan:
 *
 * Package models berisi definisi model data yang sesuai dengan struktur tabel pada database.
 * Struct Lecturers merepresentasikan tabel dosen pada database.
 *
 * Properti-properti pada struct Lecturers adalah sebagai berikut:
 * - ID: ID unik dosen (kolom primaryKey).
 * - Code: Kode dosen.
 * - Name: Nama dosen.
 * - StudyProgramID: ID program studi yang terkait dengan dosen (kolom not null).
 * - StudyProgram: Objek StudyPrograms yang merupakan relasi ke program studi yang terkait (foreign key).
 * - CreatedAt: Waktu pembuatan dosen.
 * - UpdatedAt: Waktu terakhir dosen diperbarui.
 * - DeletedAt: Waktu penghapusan dosen (kolom index, akan bernilai NULL jika belum dihapus).
 *
 * Method TableName digunakan untuk mengembalikan nama tabel yang sesuai dengan struktur.
 *
 * Pastikan untuk menggunakan model ini sesuai dengan struktur tabel dosen pada database.
 * Anda dapat menyertakan komentar ini dalam bahasa Indonesia pada kode yang relevan.
 */