// Remove Duplicates from Sorted Arra
package main

import "fmt"

func main() {

	numbers := []int{1, 2, 2, 4, 4, 5, 7, 9}

	fmt.Println(darray(numbers))

	result := darray(numbers)

	fmt.Println(result)

}

func darray(num []int) int {
	if len(num) == 0 {
		return 0
	}
	i := 1
	last := num[0]
	for _, curr := range num[1:] {
		if curr > last {
			num[i] = curr
			last = curr
			i++
		}
	}
	return i

}
