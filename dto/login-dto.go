package dto

//LoginDTO merupakan objek transfer data untuk login pengguna.
type LoginDTO struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required,min=6"`
}

/*
 * LoginDTO merupakan objek transfer data untuk login pengguna.
 *
 * Email: Alamat email pengguna yang digunakan untuk login.
 * Password: Kata sandi pengguna yang digunakan untuk login. Harus terdiri dari minimal 6 karakter.
 *
 * Tag JSON menentukan nama properti ketika data dikirim dalam format JSON.
 * Tag form menentukan nama properti ketika data dikirim melalui formulir HTML.
 * Tag binding menentukan aturan validasi untuk pengikatan properti.
 */