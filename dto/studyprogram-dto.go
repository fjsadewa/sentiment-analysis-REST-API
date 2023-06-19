package dto

//StudyProgramDTO merupakan objek transfer data untuk program studi.
type StudyProgramDTO struct {
	ID   uint   `json:"id" form:"id" binding:"required"`
	Code string `json:"code" form:"code" binding:"required"`
	Name string `json:"name" form:"name" binding:"required"`
}

//StudyProgramCreateDTO merupakan objek transfer data untuk membuat program studi.
type StudyProgramCreateDTO struct {
	Code string `json:"code" form:"code" binding:"required"`
	Name string `json:"name" form:"name" binding:"required"`
}

//StudyProgramUpdateDTO merupakan objek transfer data untuk memperbarui program studi.
type StudyProgramUpdateDTO struct {
	ID   uint   `json:"id" form:"id" binding:"required"`
	Code string `json:"code" form:"code" binding:"required"`
	Name string `json:"name" form:"name" binding:"required"`
}

//StudyProgramIDDTO merupakan objek transfer data yang hanya berisi ID program studi.
type StudyProgramIDDTO struct {
	ID uint64 `json:"id" form:"id"`
}

//StudyProgramCodeDTO merupakan objek transfer data yang hanya berisi kode program studi.
type StudyProgramCodeDTO struct {
	Code string `json:"code" form:"code" binding:"required"`
}

//ErrorDTO merupakan objek transfer data untuk kesalahan yang terjadi pada program studi.
type ErrorDTO struct {
	Error ErrorStudyProgramDTO `json:"error"`
}

//ErrorStudyProgramDTO merupakan objek transfer data yang berisi pesan kesalahan program studi.
type ErrorStudyProgramDTO struct {
	Message string `json:"message"`
}


/*
 * StudyProgramDTO merupakan objek transfer data untuk program studi.
 *
 * ID: ID program studi. Wajib diisi.
 * Code: Kode program studi. Wajib diisi.
 * Name: Nama program studi. Wajib diisi.
 *
 * StudyProgramCreateDTO merupakan objek transfer data untuk membuat program studi.
 *
 * Code: Kode program studi. Wajib diisi.
 * Name: Nama program studi. Wajib diisi.
 *
 * StudyProgramUpdateDTO merupakan objek transfer data untuk memperbarui program studi.
 *
 * ID: ID program studi yang akan diperbarui. Wajib diisi.
 * Code: Kode program studi. Wajib diisi.
 * Name: Nama program studi. Wajib diisi.
 *
 * StudyProgramIDDTO merupakan objek transfer data yang hanya berisi ID program studi.
 *
 * ID: ID program studi. Opsional.
 *
 * StudyProgramCodeDTO merupakan objek transfer data yang hanya berisi kode program studi.
 *
 * Code: Kode program studi. Wajib diisi.
 *
 * ErrorDTO merupakan objek transfer data untuk kesalahan yang terjadi pada program studi.
 *
 * Error: Objek ErrorStudyProgramDTO yang berisi pesan kesalahan.
 *
 * ErrorStudyProgramDTO merupakan objek transfer data yang berisi pesan kesalahan program studi.
 *
 * Message: Pesan kesalahan. Wajib diisi.
 *
 * Tag JSON menentukan nama properti ketika data dikirim dalam format JSON.
 * Tag form menentukan nama properti ketika data dikirim melalui formulir HTML.
 * Tag binding menentukan aturan validasi untuk pengikatan properti.
 */