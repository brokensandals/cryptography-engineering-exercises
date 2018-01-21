package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %v HEX_MESSAGE\n", os.Args[0])
		os.Exit(1)
	}

	target := strings.ToLower(os.Args[1])
	targetbytes, err := hex.DecodeString(target)
	if err != nil {
		panic(err)
	}

	keep := uint(len(target) / 2)

	candidatebytes := make([]byte, 8, 8)
	steps := 0
	for match := false; !match; {
		for i := uint(0); i < 8; i++ {
			candidatebytes[i] = byte(steps >> (i * 8))
		}
		steps++
		sha := sha512.Sum512(candidatebytes)

		match = true
		for i := uint(0); i < keep && match; i++ {
			if targetbytes[i] != sha[i] {
				match = false
			}
		}
	}

	fmt.Printf("Took %v steps\nPreimage: %v\nASCII: %v\n", steps, hex.EncodeToString(candidatebytes), string(candidatebytes))
}
