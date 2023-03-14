package main

import (
	"fmt"
	"strconv"
)

func convertToBin(v int) (result string) {
	for ; v > 0; v /= 2 {
		result = strconv.Itoa(v%2) + result
	}
	return
}

func main() {
	fmt.Println(convertToBin(5))
	fmt.Println(convertToBin(13))
	fmt.Println(fmt.Errorf("%s", "a"))
}
