package main

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
)

func main() {
	key, _ := hex.DecodeString("8000000000000000000000000000000000000000000000000000000000000001")
	plaintext, _ := hex.DecodeString("296C93FDF499AAEB4194BABC2E63561D")
	cipher, _ := aes.NewCipher(key)
	ciphertext := make([]byte, 16, 16)
	cipher.Encrypt(ciphertext, plaintext)
	fmt.Println(string(hex.EncodeToString(ciphertext)))
}
