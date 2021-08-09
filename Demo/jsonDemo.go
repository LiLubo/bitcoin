package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	strs := []string{"hello", "world", "lubo"}

	str := strings.Join(strs, "=")

	fmt.Printf("str: %s\n", str)

	joinRes := bytes.Join([][]byte{[]byte("hello"), []byte("world"), []byte("lubo")}, []byte("="))
	fmt.Printf("joinRes: %s\n", joinRes)
}
