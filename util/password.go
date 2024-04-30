package util

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const (
	salt = "test"
)

func HashPassword(password string) (string, error) {
	hashPssword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	return string(hashPssword), nil
}

func CheckHashPassword(hashPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}

func PSHA256(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	hashedPassword := hex.EncodeToString(h.Sum([]byte(salt)))
	return hashedPassword
}

func CheckPHashPassword(hashPassword string, password string) error {
	if hashPassword != PSHA256(password) {
		return errors.New("密码错误")
	}
	return nil
}
