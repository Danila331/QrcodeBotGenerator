package pkg

import (
	"fmt"
	"image/color"
	"net/http"

	"github.com/skip2/go-qrcode"
)

func ValidUrl(url string) bool {

	response, err := http.Head(url)

	if err != nil {
		return false
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return false
	}
	return true
}

func GeneratorQr(url string) (string, error) {
	filename := fmt.Sprintf("./images/%d.png", Count)
	Count++
	err := qrcode.WriteColorFile(url, qrcode.Medium, 400, color.White, color.Black, filename)
	if err != nil {
		fmt.Println("Error creating QR code:", err)
		return "", err
	}
	return filename, nil
}
