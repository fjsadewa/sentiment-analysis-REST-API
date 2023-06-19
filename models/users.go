package models

import "time"

//User merupakan model untuk tabel pengguna (user)
type User struct {
	ID        uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Username  string    `gorm:"type:varchar(255)" json:"username"`
	Email     string    `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password  string    `gorm:"->;<-;not null" json:"password"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"update_at,omitempty"`
	DeletedAt time.Time `gorm:"index"  json:"deleted_at,omitempty"`
	Token     string    `gorm:"-" json:"token,omitempty"`
}

//TableName digunakan untuk mengembalikan nama tabel yang sesuai
func (User) TableName() string {
	return "users"
}

/*
 * Keterangan:
 *
 * Package models berisi definisi model data yang sesuai dengan struktur tabel pada database.
 * Struct User merepresentasikan tabel pengguna (user) pada database.
 *
 * Properti-properti pada struct User adalah sebagai berikut:
 * - ID: ID unik pengguna (kolom primary key dengan auto increment).
 * - Username: Nama pengguna.
 * - Email: Alamat email pengguna (dijadikan sebagai unique index).
 * - Password: Kata sandi pengguna.
 * - CreatedAt: Waktu pembuatan pengguna.
 * - UpdatedAt: Waktu terakhir pengguna diperbarui.
 * - DeletedAt: Waktu penghapusan pengguna (kolom index, akan bernilai NULL jika belum dihapus).
 * - Token: Token pengguna (tidak disimpan dalam database).
 *
 * Method TableName digunakan untuk mengembalikan nama tabel yang sesuai dengan struktur.
 *
 * Pastikan untuk menggunakan model ini sesuai dengan struktur tabel pengguna pada database.
 * Anda dapat menyertakan komentar ini dalam bahasa Indonesia pada kode yang relevan.
 */