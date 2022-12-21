package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func TranslateApiToEn(keyword string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}
	res, err := http.Get(os.Getenv("TRANSLATE_API_ENDPOINT") + fmt.Sprintf("?text=%s&source=%s&target=%s", keyword, "ja", "en"))
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func TranslateApiToJa(keyword string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}
	res, err := http.Get(os.Getenv("TRANSLATE_API_ENDPOINT") + fmt.Sprintf("?text=%s&source=%s&target=%s", keyword, "en", "ja"))
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
