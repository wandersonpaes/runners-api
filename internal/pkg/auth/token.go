package auth

import (
	"log"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var secretKey []byte

func SetUp() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	secretKey = []byte(os.Getenv("SECRET_KEY"))
}

func CreateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 1).Unix()
	permissions["userID"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString([]byte(secretKey))
}
