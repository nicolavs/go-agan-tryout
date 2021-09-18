package config

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ValidToken(t *jwt.Token, rolesNeeded []string) bool {
	claims := t.Claims.(jwt.MapClaims)
	fmt.Println(claims)
	roles := claims["roles"].([]string)
	if Contains(roles, "administrator") {
		return true
	}

	for _, r := range rolesNeeded {
		if Contains(roles, r) {
			return true
		}
	}

	return false
}
