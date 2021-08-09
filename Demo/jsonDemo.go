package main

import (
	"fmt"
	"strings"
)

func main() {
	strs := []string{"hello", "world", "itcast"}

	str := strings.Join(strs, "=")

	fmt.Printf("str: %s\n", str)
}
