package utils

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"


func RandomString(n int) string {
	var sb strings.Builder
	sb.Grow(10)
	l := len(alphabet)

	for i := 0; i <= n; i++ {
		c := alphabet[rand.Intn(l)]
		sb.WriteByte(c)
	}
	return sb.String()

}


func RandomInteger(min , max int64) int64 {
	return min + rand.Int63n(max-min+1)
}


func RandomEmail() string {
	s := RandomString(10)
	domain := "@email.com"
	return s + domain
}