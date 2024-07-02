package security

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"log"
)

// Decrypt function
func Decrypt(encryptedToken, ivBase64, keyBase64 string) (string, error) {
	// Decode base64
	encryptedData, err := base64.StdEncoding.DecodeString(encryptedToken)
	if err != nil {
		return "", fmt.Errorf("failed to decode encrypted token: %v", err)
	}
	iv, err := base64.StdEncoding.DecodeString(ivBase64)
	if err != nil {
		return "", fmt.Errorf("failed to decode IV: %v", err)
	}
	key, err := base64.StdEncoding.DecodeString(keyBase64)
	if err != nil {
		return "", fmt.Errorf("failed to decode key: %v", err)
	}
	iv := []byte(ivBase64)
	key := []byte(keyBase64)

	// Ensure key length is 16, 24, or 32 bytes for AES
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return "", fmt.Errorf("invalid key size")
	}

	// Create AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %v", err)
	}

	// Create a new CBC mode cipher
	mode := cipher.NewCBCDecrypter(block, iv)

	// Decrypt the data
	decrypted := make([]byte, len(encryptedData))
	mode.CryptBlocks(decrypted, encryptedData)

	// Remove padding (Assuming PKCS7 padding)
	paddingLen := int(decrypted[len(decrypted)-1])
	if paddingLen > len(decrypted) {
		return "", fmt.Errorf("padding length is too large")
	}
	decrypted = decrypted[:len(decrypted)-paddingLen]

	return string(decrypted), nil
}
func main() {
	// Example values (replace these with your actual values)
	encryptedToken := "wdPeMY9/xi6ieRfYOW11cF9IOqBDMNraML5WK9huZQ611zT6Q+bfOjIs2R3aiGk8NzqvvSshKaB42DfGXUlviN0sU8Dk/MJu7/+0CyDUbvjmJaxTaUL7MHw5g8bLZkemQcSfHVVXouNmakvhphP9Kg=="
	ivBase64 := "00c5a4f64285ee4a38173a14e9fb6019"
	keyBase64 := "30ee197ffc44308c"

	decryptedToken, err := Decrypt(encryptedToken, ivBase64, keyBase64)
	if err != nil {
		log.Fatalf("Decryption failed: %v", err)
	}

	fmt.Println("Decrypted token:", decryptedToken)
}
