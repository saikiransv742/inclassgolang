package main

import (
	"fmt"
	"math"
)

func primenumberornot(num int) {
	if num < 2 {
		fmt.Println("Number must be greater than 2.")
		return
	}
	sq_root := int(math.Sqrt(float64(num)))
	for i := 2; i <= sq_root; i++ {
		if num%i == 0 {
			fmt.Println("Not a Prime Number")
			return
		}
	}
	fmt.Println("Prime Number")
	return
}
func main() {
	primenumberornot(0)
	primenumberornot(2)
	primenumberornot(12)
	primenumberornot(55)

}
