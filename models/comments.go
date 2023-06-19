package models

import "time"

//Comments merupakan model untuk tabel komentar
type Comments struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Comment    string    `gorm:"type:text" json:"comment"`
	CourseID   uint64    `gorm:"not null" json:"-"`
	Course     Courses   `gorm:"foreignKey:CourseID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"course_id"`
	LecturerID uint64    `gorm:"not null" json:"-"`
	Lecturer   Lecturers `gorm:"foreignKey:LecturerID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"lecturer_id"`
	SchoolYear uint      `gorm:"not null" json:"school_year"`
	Semester   uint      `gorm:"not null" json:"semester"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"update_at,omitempty"`
	DeletedAt  time.Time `gorm:"index"  json:"deleted_at,omitempty"`
}

//TableName digunakan untuk mengembalikan nama tabel yang sesuai
func (Comments) TableName() string {
	return "comments"
}

/*
 * Keterangan:
 *
 * Package models berisi definisi model data yang sesuai dengan struktur tabel pada database.
 * Struct Comments merepresentasikan tabel komentar pada database.
 *
 * Properti-properti pada struct Comments adalah sebagai berikut:
 * - ID: ID unik komentar (kolom primaryKey).
 * - Comment: Isi komentar (tipe text).
 * - CourseID: ID kursus yang terkait dengan komentar (kolom foreignKey).
 * - Course: Objek kursus yang terkait dengan komentar.
 * - LecturerID: ID dosen yang terkait dengan komentar (kolom foreignKey).
 * - Lecturer: Objek dosen yang terkait dengan komentar.
 * - SchoolYear: Tahun ajaran komentar dibuat.
 * - Semester: Semester komentar dibuat.
 * - CreatedAt: Waktu pembuatan komentar.
 * - UpdatedAt: Waktu terakhir komentar diperbarui.
 * - DeletedAt: Waktu penghapusan komentar (kolom index, akan bernilai NULL jika belum dihapus).
 *
 * Method TableName digunakan untuk mengembalikan nama tabel yang sesuai dengan struktur.
 *
 * Pastikan untuk menggunakan model ini sesuai dengan struktur tabel komentar pada database.
 * Anda dapat menyertakan komentar ini dalam bahasa Indonesia pada kode yang relevan.
 */
