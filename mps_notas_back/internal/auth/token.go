package auth

import (
	"errors"
	"fmt"
	"mps_notas_back/internal/config"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(usuerId int) (string, error) {
	permitions := jwt.MapClaims{}

	permitions["authorized"] = true
	permitions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permitions["user_id"] = usuerId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permitions)
	return token.SignedString(config.SecretKey)

}

// ValidateToken validate if the token are valid
func ValidateToken(r *http.Request) error {
	strToken := extractToken(r)
	token, err := jwt.Parse(strToken, getVerificationKey)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return errors.New("invalid token")
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	splitedString := strings.Split(token, " ")
	if len(splitedString) == 2 && splitedString[0] == "Bearer" {
		return splitedString[1]
	}
	return ""
}

func getVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("signin method not expected %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}

// ExtractUserId get the request and extract the id of the user
func ExtractUserId(r *http.Request) (int, error) {
	strToken := extractToken(r)
	token, err := jwt.Parse(strToken, getVerificationKey)
	if err != nil {
		return 0, err
	}

	if permitions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		userId, err := strconv.ParseInt(string(fmt.Sprintf("%.0f", permitions["user_id"])), 10, 64)
		if err != nil {
			return 0, err
		}
		return int(userId), nil
	}
	return 0, fmt.Errorf("unexpected error")
}
