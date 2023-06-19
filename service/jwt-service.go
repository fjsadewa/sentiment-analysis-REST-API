package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//JWTService adalah kontrak tentang apa yang dapat dilakukan oleh JWTService.
type JWTService interface {
	GenerateToken(userID string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

//NewJWTService adalah metode untuk membuat instance baru dari JWTService.
func NewJWTService() JWTService {
	return &jwtService{
		issuer:    "hadad123",
		secretKey: getSecretKey(),
	}
}

/*
getSecretKey adalah fungsi untuk mendapatkan kunci rahasia dari environment variable JWT_SECRET.
Jika environment variable tidak ada, maka kunci rahasia default akan digunakan.
*/
func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		secretKey = "slkfmgsdlfkmsdpfgpismdfg"
	}
	return secretKey
}

//GenerateToken digunakan untuk menghasilkan token JWT berdasarkan userID.
func (j *jwtService) GenerateToken(UserID string) string {
	expirationTime := time.Now().Add(3600 * time.Second)

	claims := &jwtCustomClaim{
		UserID,
		jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

/*
ValidateToken digunakan untuk memvalidasi token JWT.
Mengembalikan objek jwt.Token jika token valid, atau error jika terjadi kesalahan.
*/
func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}
