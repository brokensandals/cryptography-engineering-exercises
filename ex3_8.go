package main

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
)

func main() {
	ciphertext, _ := hex.DecodeString("539B333B39706D149028CFE1D9D4A407")
	key, _ := hex.DecodeString("8000000000000000000000000000000000000000000000000000000000000001")
	cipher, _ := aes.NewCipher(key)
	plaintext := make([]byte, 16, 16)
	cipher.Decrypt(plaintext, ciphertext)
	fmt.Println(string(hex.EncodeToString(plaintext)))
}
