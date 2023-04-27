package main

import (
	"fmt"

	"github.com/Martande8055/encrypt-decrypt/encription"

	"github.com/Martande8055/encrypt-decrypt/decription"
)

func main() {
	enc := encription.Encrypt("saurabh")
	fmt.Println(enc)
	fmt.Println(decription.Decript(enc))
}
