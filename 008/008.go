package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%08b\n", swapBits(0b10101010)) // expected 01010101
	fmt.Printf("%08b\n", swapBits(0b11100010)) // expected 11010001
}

func swapBits(a uint8) uint8 {
	return ((a&0b10101010)>>1 | (a&0b1010101)<<1)
}
