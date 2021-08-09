package main

import (
	"crypto/sha256"
	"fmt"
)

func powDemo() {
	data := "hello world"

	for i := 0; i < 1000000; i++ {
		hash := sha256.Sum256([]byte(data + fmt.Sprint(i)))
		fmt.Printf("hash: %x, nonce: %d\n", hash, i)
	}
}
