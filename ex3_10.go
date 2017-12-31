package main

import (
	"bytes"
	"crypto/des"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	key, _ := hex.DecodeString(os.Args[1])
	plaintext, _ := hex.DecodeString(os.Args[2])
	cipher, _ := des.NewCipher(key)
	ciphertext := make([]byte, 8, 8)
	cipher.Encrypt(ciphertext, plaintext)
	fmt.Printf("Ciphertext of original key and value: %v\n", hex.EncodeToString(ciphertext))

	key2 := make([]byte, 8, 8)
	for i, v := range key {
		key2[i] = v ^ 255
	}
	plaintext2 := make([]byte, 8, 8)
	for i, v := range plaintext {
		plaintext2[i] = v ^ 255
	}
	cipher2, _ := des.NewCipher(key2)
	ciphertext2 := make([]byte, 8, 8)
	cipher2.Encrypt(ciphertext2, plaintext2)
	fmt.Printf("Ciphertext of key and value complements: %v\n", hex.EncodeToString(ciphertext2))

	ciphertext3 := make([]byte, 8, 8)
	for i, v := range ciphertext2 {
		ciphertext3[i] = v ^ 255
	}
	fmt.Printf("Complement of second ciphertext: %v\n", hex.EncodeToString(ciphertext3))
	fmt.Printf("Successfully demonstrated complementation property? %v\n", bytes.Equal(ciphertext, ciphertext3))
}
