package main

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
)

func main() {
	key, _ := hex.DecodeString("8000000000000000000000000000000000000000000000000000000000000001")
	cipher, _ := aes.NewCipher(key)
	plaintext, _ := hex.DecodeString("626C6F636B2063697068657273202020686173682066756E6374696F6E732078626C6F636B2063697068657273202020")
	ciphertext := make([]byte, 48, 48)
	for i := 0; i < len(plaintext); i += 16 {
		cipher.Encrypt(ciphertext[i:i+16], plaintext[i:i+16])
	}
	fmt.Println(string(hex.EncodeToString(ciphertext)))
}
