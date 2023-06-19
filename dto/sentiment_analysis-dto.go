package dto

//SentimenAnalysisCreateDTO merupakan objek transfer data untuk membuat analisis sentimen.
type SentimenAnalysisCreateDTO struct {
	SentimentType   string  `json:"sentiment_type" form:"sentiment_type" binding:"required"`
	CommentID       uint64  `json:"comment_id" form:"comment_id" binding:"required"`
	ConfidenceLevel float64 `json:"confidence_level" form:"confidence_level" binding:"required"`
}

//SentimenAnalysisUpdateDTO merupakan objek transfer data untuk memperbarui analisis sentimen.
type SentimenAnalysisUpdateDTO struct {
	ID              uint64  `json:"id" form:"id" binding:"required"`
	SentimentType   string  `json:"sentiment_type" form:"sentiment_type" binding:"required"`
	CommentID       uint64  `json:"comment_id" form:"comment_id" binding:"required"`
	ConfidenceLevel float64 `json:"confidence_level" form:"confidence_level" binding:"required"`
}

//SentimenAnalysisDTO merupakan objek transfer data untuk analisis sentimen.
type SentimenAnalysisDTO struct {
	ID               uint64  `json:"id" form:"id" binding:"required"`
	SentimentType    string  `json:"sentiment_type" form:"sentiment_type" binding:"required"`
	ConfidenceLevel  float64 `json:"confidence_level" form:"confidence_level" binding:"required"`
	CommentID        uint64  `json:"comment_id" form:"comment_id" binding:"required"`
	Comment          string  `json:"comment" form:"comment" binding:"required"`
	CourseID         uint64  `json:"course_id" form:"course_id" binding:"required"`
	CourseCode       string  `json:"course_code" form:"course_code" binding:"required"`
	CourseName       string  `json:"course_name" form:"course_name" binding:"required"`
	LecturerID       uint64  `json:"lecturer_id" form:"lecturer_id" binding:"required"`
	LecturerCode     string  `json:"lecturer_code" form:"lecturer_code" binding:"required"`
	LecturerName     string  `json:"lecturer_name" form:"lecturer_name" binding:"required"`
	StudyProgramID   uint64  `json:"study_program_id" form:"study_program_id" binding:"required"`
	StudyProgramCode string  `json:"study_program_code" form:"study_program_code" binding:"required"`
	StudyProgramName string  `json:"study_program_name" form:"study_program_name" binding:"required"`
	SchoolYear       uint    `json:"school_year" form:"school_year" binding:"required"`
	Semester         uint    `json:"semester" form:"semester" binding:"required"`
}

/*
 * SentimenAnalysisCreateDTO merupakan objek transfer data untuk membuat analisis sentimen.
 *
 * SentimentType: Jenis sentimen dari analisis. Wajib diisi.
 * CommentID: ID komentar yang akan dianalisis. Wajib diisi.
 * ConfidenceLevel: Tingkat kepercayaan terhadap hasil analisis. Wajib diisi.
 *
 * SentimenAnalysisUpdateDTO merupakan objek transfer data untuk memperbarui analisis sentimen.
 *
 * ID: ID analisis sentimen yang akan diperbarui. Wajib diisi.
 * SentimentType: Jenis sentimen dari analisis. Wajib diisi.
 * CommentID: ID komentar yang akan dianalisis. Wajib diisi.
 * ConfidenceLevel: Tingkat kepercayaan terhadap hasil analisis. Wajib diisi.
 *
 * SentimenAnalysisDTO merupakan objek transfer data untuk analisis sentimen.
 *
 * ID: ID analisis sentimen. Wajib diisi.
 * SentimentType: Jenis sentimen dari analisis. Wajib diisi.
 * ConfidenceLevel: Tingkat kepercayaan terhadap hasil analisis. Wajib diisi.
 * CommentID: ID komentar yang dianalisis. Wajib diisi.
 * Comment: Isi komentar yang dianalisis. Wajib diisi.
 * CourseID: ID kursus terkait dengan komentar. Wajib diisi.
 * CourseCode: Kode kursus terkait dengan komentar. Wajib diisi.
 * CourseName: Nama kursus terkait dengan komentar. Wajib diisi.
 * LecturerID: ID dosen terkait dengan komentar. Wajib diisi.
 * LecturerCode: Kode dosen terkait dengan komentar. Wajib diisi.
 * LecturerName: Nama dosen terkait dengan komentar. Wajib diisi.
 * StudyProgramID: ID program studi terkait dengan komentar. Wajib diisi.
 * StudyProgramCode: Kode program studi terkait dengan komentar. Wajib diisi.
 * StudyProgramName: Nama program studi terkait dengan komentar. Wajib diisi.
 * SchoolYear: Tahun ajaran terkait dengan komentar. Wajib diisi.
 * Semester: Semester terkait dengan komentar. Wajib diisi.
 *
 * Tag JSON menentukan nama properti ketika data dikirim dalam format JSON.
 * Tag form menentukan nama properti ketika data dikirim melalui formulir HTML.
 * Tag binding menentukan aturan validasi untuk pengikatan properti.
 */