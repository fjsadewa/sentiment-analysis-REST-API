package dto

//LectureCreateDTO digunakan untuk membuat objek dosen baru
type LecturerCreateDTO struct {
	Code           string `json:"code" form:"code" binding:"required"`
	Name           string `json:"name" form:"name" binding:"required"`
	StudyProgramID uint64 `json:"study_program_id" form:"study_program_id" binding:"required"`
}

//LectureUpdateDTO digunakan untuk memperbarui objek dosen yang ada.
type LecturerUpdateDTO struct {
	ID             uint64 `json:"id" form:"id" binding:"required"`
	Code           string `json:"code" form:"code" binding:"required"`
	Name           string `json:"name" form:"name" binding:"required"`
	StudyProgramID uint64 `json:"study_program_id" form:"study_program_id" binding:"required"`
}

/*
 * Keterangan:
 *
 * LecturerCreateDTO digunakan saat membuat dosen baru.
 * LecturerUpdateDTO digunakan saat memperbarui dosen yang sudah ada.
 *
 * Setiap DTO memiliki properti yang sesuai dengan kebutuhan dan aturan validasi tertentu.
 * Properti yang diberi tag "json" menentukan nama properti saat dikirim dalam format JSON.
 * Properti yang diberi tag "form" menentukan nama properti saat dikirim melalui formulir HTML.
 * Properti yang diberi tag "binding" menentukan aturan validasi untuk pengikatan properti.
 * Properti dengan tag "required" menandakan bahwa properti tersebut wajib diisi.
 *
 * Pastikan untuk menggunakan DTO yang sesuai dengan keperluan Anda saat membuat atau memperbarui dosen.
 * Properti yang harus diisi dengan benar adalah "code", "name", dan "study_program_id".
 */