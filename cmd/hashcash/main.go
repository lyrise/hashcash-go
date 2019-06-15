package main

import (
	hashcash "github.com/lyrise/hashcash-go"
)

func main() {
	var buffer [32]byte

	for i := 0; i < len(buffer); i++ {
		buffer[i] = byte(i)
	}

	hashcash.SimpleSha2_256(buffer[:])
}
