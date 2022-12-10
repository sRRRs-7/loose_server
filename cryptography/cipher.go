package cryptography

import (
	"crypto/sha256"
	"fmt"
)

func HashPassword(password string) (string, error) {
	sha := sha256.New()
	_, err := sha.Write([]byte(password))
	if err != nil {
		return "", fmt.Errorf("sha256 encrypt error: %v", err)
	}
	hash := sha.Sum(nil)

	s := fmt.Sprintf("%x \n", hash)
	return s, nil
}

func VerifyHash(password, hashPassword string) (bool, error) {
	if password != hashPassword {
		return false, fmt.Errorf("verify password and hashPassword error")
	}
	if []byte(password)[7] != []byte(hashPassword)[7] {
		return false, fmt.Errorf("verify password and hashPassword error")
	}
	return true, nil
}