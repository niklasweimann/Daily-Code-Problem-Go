package main

import "fmt"

func main() {
	playHanoi(3, "1", "2", "3")
}

func playHanoi(n int, from_rod string,
	to_rod string, aux_rod string) {
	if n == 0 {
		return
	}

	playHanoi(n-1, from_rod, aux_rod, to_rod)
	fmt.Printf("Move Disk #%v to %v\n", n, to_rod)
	playHanoi(n-1, aux_rod, to_rod, from_rod)
}
