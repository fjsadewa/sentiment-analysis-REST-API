package dto

//CommentCreateDTO digunakan untuk membuat objek komentar baru
type CommentCreateDTO struct {
	SentimentType   string  `json:"sentiment_type" form:"sentiment_type"`
	ConfidenceLevel float64 `json:"confidence_level" form:"confidence_level"`
	Comment         string  `json:"comment" form:"comment" binding:"required"`
	CourseID        uint64  `json:"course_id" form:"course_id" binding:"required"`
	LecturerID      uint64  `json:"lecturer_id" form:"lecturer_id" binding:"required"`
	SchoolYear      uint    `json:"school_year" form:"school_year" binding:"required"`
	Semester        uint    `json:"semester" form:"semester" binding:"required"`
}

//CommentUpdateDTO digunakan untuk memperbarui objek komentar yang ada. 
type CommentUpdateDTO struct {
	ID              uint64  `json:"id" form:"id" binding:"required"`
	SentimentType   string  `json:"sentiment_type" form:"sentiment_type"`
	ConfidenceLevel float64 `json:"confidence_level" form:"confidence_level"`
	Comment         string  `json:"comment" form:"comment" binding:"required"`
	CourseID        uint64  `json:"course_id" form:"course_id" binding:"required"`
	LecturerID      uint64  `json:"lecturer_id" form:"lecturer_id" binding:"required"`
	SchoolYear      uint    `json:"school_year" form:"school_year" binding:"required"`
	Semester        uint    `json:"semester" form:"semester" binding:"required"`
}

//CommentDTO digunakan untuk mengirim informasi komentar ke komponen lain.
type CommentDTO struct {
	CommentID        uint64 `json:"comment_id" form:"id" binding:"required"`
	Comment          string `json:"comment" form:"comment" binding:"required"`
	LecturerID       uint64 `json:"lecturer_id" form:"lecturer_id" binding:"required"`
	LecturerCode     string `json:"lecturer_code" form:"lecturer_code" binding:"required"`
	LecturerName     string `json:"lecturer_name" form:"lecturer_name" binding:"required"`
	CourseID         uint64 `json:"course_id" form:"course_id" binding:"required"`
	CourseCode       string `json:"course_code" form:"course_code" binding:"required"`
	CourseName       string `json:"course_name" form:"course_name" binding:"required"`
	StudyProgramID   uint64 `json:"study_program_id" form:"study_program_id" binding:"required"`
	StudyProgramCode string `json:"study_program_code" form:"study_program_code" binding:"required"`
	StudyProgramName string `json:"study_program_name" form:"study_program_name" binding:"required"`
	SchoolYear       uint   `json:"school_year" form:"school_year" binding:"required"`
	Semester         uint   `json:"semester" form:"semester" binding:"required"`
}

/*
 * Keterangan:
 *
 * CommentCreateDTO digunakan saat membuat komentar baru.
 * CommentUpdateDTO digunakan saat memperbarui komentar yang sudah ada.
 * CommentDTO digunakan untuk mengirim informasi komentar ke komponen lain.
 *
 * Setiap DTO memiliki properti yang sesuai dengan kebutuhan dan aturan validasi tertentu.
 * Properti yang diberi tag "json" menentukan nama properti saat dikirim dalam format JSON.
 * Properti yang diberi tag "form" menentukan nama properti saat dikirim melalui formulir HTML.
 * Properti yang diberi tag "binding" menentukan aturan validasi untuk pengikatan properti.
 * Properti dengan tag "required" menandakan bahwa properti tersebut wajib diisi.
 *
 * Pastikan untuk menggunakan DTO yang sesuai dengan keperluan Anda saat membuat atau memperbarui komentar,
 * serta ketika mengirim informasi komentar ke komponen lain.
 */