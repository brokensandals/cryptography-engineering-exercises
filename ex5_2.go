package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

func main() {
	input, err := hex.DecodeString("48656C6C6F2C20776F726C642E202020")
	if err != nil {
		panic(err)
	}

	hash := sha512.New()
	_, err = hash.Write(input)
	if err != nil {
		panic(err)
	}

	result := hash.Sum(nil)

	fmt.Println(hex.EncodeToString(result))
}
