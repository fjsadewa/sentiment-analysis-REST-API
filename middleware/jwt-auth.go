package middleware

import (
	"log"
	"net/http"
	"sentiment/helper"
	"sentiment/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

/*
AuthorizeJWT adalah middleware untuk memverifikasi
token JWT dalam header Authorization pada permintaan HTTP.
*/
func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		//Mengambil nilai header "Authorization" dari permintaan
		authHeader := c.GetHeader("Authorization")
		//Jika nilai authHeader kosong, menampilkan pesan kesalahan dan memberikan respons JSON dengan status HTTP 400 Bad Request
		if authHeader == "" {
			response := helper.BuildErrorResponse("Failed to process request", "No token found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		//Memvalidasi token menggunakan objek jwtService.ValidateToken()
		token, err := jwtService.ValidateToken(authHeader)
		//Jika token valid, mengakses klaim-klaim dalam token dan mencetak informasi relevan ke log
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]: ", claims["user_id"])
			log.Println("Claim[issuer] : ", claims["issuer"])
		}else {
			//Jika token tidak valid, mencetak pesan kesalahan dan memberikan respons JSON dengan status HTTP 401 Unauthorized 
			log.Println(err)
			response := helper.BuildErrorResponse("Token is not valid", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
