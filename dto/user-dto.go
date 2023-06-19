package dto

//UserUpdateDTO merupakan objek transfer data untuk memperbarui pengguna.
type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id"`
	Username string `json:"username" form:"username" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password,omitempty" form:"password,omitempty"`
}

/*
 * UserUpdateDTO merupakan objek transfer data untuk memperbarui pengguna.
 *
 * ID: ID pengguna.
 * Username: Nama pengguna. Wajib diisi.
 * Email: Alamat email pengguna. Wajib diisi dan harus berformat email yang valid.
 * Password: Kata sandi pengguna. Opsional. Properti ini ditandai dengan omitempty,
 *           yang berarti jika properti ini kosong atau tidak ada, maka akan diabaikan.
 *
 * Tag JSON menentukan nama properti ketika data dikirim dalam format JSON.
 * Tag form menentukan nama properti ketika data dikirim melalui formulir HTML.
 * Tag binding menentukan aturan validasi untuk pengikatan properti.
 *
 * Jika Anda ingin memperbarui pengguna, pastikan untuk menyertakan ID pengguna yang valid.
 * Anda dapat memperbarui properti Username dan Email. Jika ingin memperbarui kata sandi,
 * Anda dapat menyertakan properti Password yang tidak kosong.
 */