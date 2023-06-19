package dto

//RegisterDTO merupakan objek transfer data untuk registrasi pengguna.
type RegisterDTO struct {
	Username string `json:"username" form:"username" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required,min=6"`
}

/*
 * RegisterDTO merupakan objek transfer data untuk registrasi pengguna.
 *
 * Username: Nama pengguna yang akan digunakan untuk registrasi. Wajib diisi.
 * Email: Alamat email pengguna yang akan digunakan untuk registrasi. Wajib diisi dan harus berformat email yang valid.
 * Password: Kata sandi pengguna yang akan digunakan untuk registrasi. Wajib diisi dan harus terdiri dari minimal 6 karakter.
 *
 * Tag JSON menentukan nama properti ketika data dikirim dalam format JSON.
 * Tag form menentukan nama properti ketika data dikirim melalui formulir HTML.
 * Tag binding menentukan aturan validasi untuk pengikatan properti.
 */