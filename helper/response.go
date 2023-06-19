package helper

import (
	"strings"
)

//Response digunakan untuk bentuk JSON yang memiliki struktur tetap
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"` // interface for dynamic data
	Data    interface{} `json:"data"`
}

//EmptyObj digunakan ketika data tidak ingin bernilai null pada JSON
type EmptyObj struct{}

//BuildResponse digunakan untuk menyisipkan nilai data ke dalam respons sukses yang dinamis
func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
	return res
}

//BuildErrorResponse digunakan untuk menyisipkan nilai data ke dalam respons gagal yang dinamis
func BuildErrorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Errors:  splittedError,
		Data:    data,
	}
	return res
}

/*
 * Keterangan:
 *
 * Paket helper berisi fungsi-fungsi yang membantu dalam pembuatan respons JSON yang konsisten.
 * Struct Response digunakan untuk bentuk JSON dengan struktur tetap.
 * Struct EmptyObj digunakan ketika data tidak ingin memiliki nilai null pada JSON.
 *
 * Fungsi BuildResponse digunakan untuk menyisipkan nilai data ke dalam respons sukses yang dinamis.
 * Fungsi BuildErrorResponse digunakan untuk menyisipkan nilai data ke dalam respons gagal yang dinamis.
 * Pada BuildErrorResponse, pesan kesalahan dipecah menjadi array menggunakan pemisah "\n".
 *
 * Pastikan untuk menggunakan fungsi-fungsi ini sesuai kebutuhan saat membangun respons JSON.
 * Anda dapat menggunakan Response dengan memasukkan status, pesan, kesalahan (jika ada), dan data yang relevan.
 * Pastikan untuk menyertakan komentar ini dalam bahasa Indonesia pada kode yang relevan.
 */