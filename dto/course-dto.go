package dto

//CourseCreate DTO digunakan untuk membuat objek kursus baru.
type CourseCreateDTO struct {
	Code string `json:"code" form:"code" binding:"required"`
	Name string `json:"name" form:"name" binding:"required"`
}

//CourseUpdateDTO  digunakan untuk memperbarui objek kursus yang ada.
type CourseUpdateDTO struct {
	ID   uint64 `json:"id" form:"id" binding:"required"`
	Code string `json:"code" form:"code" binding:"required"`
	Name string `json:"name" form:"name" binding:"required"`
}

/*
 * Keterangan:
 *
 * CourseCreateDTO digunakan saat membuat kursus baru.
 * CourseUpdateDTO digunakan saat memperbarui kursus yang sudah ada.
 *
 * Setiap DTO memiliki properti yang sesuai dengan kebutuhan dan aturan validasi tertentu.
 * Properti yang diberi tag "json" menentukan nama properti saat dikirim dalam format JSON.
 * Properti yang diberi tag "form" menentukan nama properti saat dikirim melalui formulir HTML.
 * Properti yang diberi tag "binding" menentukan aturan validasi untuk pengikatan properti.
 * Properti dengan tag "required" menandakan bahwa properti tersebut wajib diisi.
 *
 * Pastikan untuk menggunakan DTO yang sesuai dengan keperluan Anda saat membuat atau memperbarui kursus.
 * Properti yang harus diisi dengan benar adalah "code" dan "name".
 */