package main

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
)

func main() {
	key, _ := hex.DecodeString("8000000000000000000000000000000000000000000000000000000000000001")
	cipher, _ := aes.NewCipher(key)
	ciphertext, _ := hex.DecodeString("87F348FF79B811AF3857D6718E5F0F917C3D26F77377635A5E43E9B5CC5D05926E26FFC5220DC7D405F1708670E6E017")
	plaintext := make([]byte, 32, 32)
	for i := 16; i < 48; i += 16 {
		prevcipherblock := ciphertext[i-16 : i]
		curcipherblock := ciphertext[i : i+16]
		cipher.Decrypt(plaintext[i-16:i], curcipherblock)
		for j := 0; j < 16; j++ {
			plaintext[i-16+j] ^= prevcipherblock[j]
		}
	}
	fmt.Println(string(plaintext))
}
