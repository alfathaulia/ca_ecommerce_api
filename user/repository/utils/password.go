package util

import (
	"fmt"
	"math/rand"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

const alphaber = "abcdefghijklmnopqrstuvwxyz"

// RandomString menegenare sebuah random string dari length of n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphaber)

	for i := 0; i < n; i++ {
		c := alphaber[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// HashPassword returns the bcrypt hash of the password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

// CheckPassword checks if the provided password is correct or not
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
