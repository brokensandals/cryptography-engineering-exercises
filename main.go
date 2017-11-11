package main

import (
	"bytes"
	"crypto/aes"
	"crypto/des"
	"encoding/hex"
	"fmt"
	"os"
)

func ex3_8() {
	ciphertext, _ := hex.DecodeString("539B333B39706D149028CFE1D9D4A407")
	key, _ := hex.DecodeString("8000000000000000000000000000000000000000000000000000000000000001")
	cipher, _ := aes.NewCipher(key)
	plaintext := make([]byte, 16, 16)
	cipher.Decrypt(plaintext, ciphertext)
	fmt.Println(string(hex.EncodeToString(plaintext)))
}

func ex3_9() {
	key, _ := hex.DecodeString("8000000000000000000000000000000000000000000000000000000000000001")
	plaintext, _ := hex.DecodeString("296C93FDF499AAEB4194BABC2E63561D")
	cipher, _ := aes.NewCipher(key)
	ciphertext := make([]byte, 16, 16)
	cipher.Encrypt(ciphertext, plaintext)
	fmt.Println(string(hex.EncodeToString(ciphertext)))
}

func ex3_10() {
	key, _ := hex.DecodeString(os.Args[2])
	plaintext, _ := hex.DecodeString(os.Args[3])
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

func ex4_4() {
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

func main() {
	ex := os.Args[1]
	switch ex {
	case "3.8":
		ex3_8()
	case "3.9":
		ex3_9()
	case "3.10":
		ex3_10()
	case "4.4":
		ex4_4()
	}
}
