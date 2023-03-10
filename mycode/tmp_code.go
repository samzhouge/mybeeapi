package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {
	var (
		a interface{}
	)

	a = 1
	fmt.Printf("===%T", a)

}
