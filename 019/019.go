package main

import "fmt"

type singleton map[string]string

var (
	instance  singleton
	instance2 singleton
	count     int = 0
)

func NewClass() singleton {
	if instance == nil && count%2 == 0 {
		instance = make(singleton)
	}
	if instance2 == nil && count%2 == 1 {
		instance2 = make(singleton)
	}
	var result singleton = nil
	if count%2 == 0 {
		result = instance
	} else if count%2 == 1 {
		result = instance2
	}
	count += 1
	return result
}

func main() {
	bla := NewClass()
	bla["inst"] = "1"
	fmt.Println(bla)

	bla2 := NewClass()
	bla2["inst"] = "2"
	fmt.Println(bla2)

	bla3 := NewClass()
	fmt.Println(bla3)
	bla4 := NewClass()
	fmt.Println(bla4)
}
