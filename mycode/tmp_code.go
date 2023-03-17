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

func sumArgs(numbers ...int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

func main() {
	fmt.Println(convertToBin(5))
	fmt.Println(convertToBin(13))
	fmt.Println(fmt.Errorf("%s", "a"))

	fmt.Println(sumArgs(1, 2, 3))

	s1 := []int{2, 4, 6, 8}
	s2 := make([]int, 10, 32)

	copy(s2, s1)
	printSlice(s2)

	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	m := map[string]string{
		"a": "aa",
		"b": "bb",
	}

	if name, ok := m["a"]; ok {
		fmt.Println(name)
	} else {
		fmt.Println("no ok")
	}

}
