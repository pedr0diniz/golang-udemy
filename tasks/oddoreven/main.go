package main

import (
	"fmt"
)

func main() {
	oddOrEven(sliceOf10Ints())
}

func sliceOf10Ints() []int {
	sl := []int{}
	for i := 0; i < 10; i++ {
		sl = append(sl, i)
	}
	return sl
}

func oddOrEven(sl []int) {
	for _, i := range sl {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}

}
