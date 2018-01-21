package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %v NUM_BITS\n", os.Args[0])
		os.Exit(1)
	}

	bits, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil || bits < 8 || bits > 48 || bits % 8 != 0 {
		fmt.Println("Number of bits must be a multiple of 8 between 8 and 48 inclusive")
		os.Exit(1)
	}

	keep := bits / 8

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	seen := make(map[string]string)

	bytes := make([]byte, keep, keep)
	var message string
	var hash string
	for {
		rng.Read(bytes)
		message = hex.EncodeToString(bytes)
		hashbytes := sha512.Sum512(bytes)
		hash = hex.EncodeToString(hashbytes[0:keep])
		if _, ok := seen[hash]; ok && seen[hash] != message {
			break
		}
		seen[hash] = message
	}

	fmt.Printf("Examined %v values\n", len(seen))
	fmt.Printf("Two values produced hash %v:\n%v\n%v\n", hash, seen[hash], message)
}
