package main

import (
	"fmt"

	"github.com/llarhub/openssl"
)

func main() {
	b := make([]byte, 10)

	openssl.RAND_Bytes(b)
	fmt.Printf("%x\n", b)

	openssl.RAND_PrivBytes(b)
	fmt.Printf("%x\n", b)
}
