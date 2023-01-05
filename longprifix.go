package main

import (
	"fmt"
	"sort"
)

func main() {

	strc := []string{"flower", "flow", "flight"}

	fmt.Println(longestCommonPrefix(strc))

	result := (longestCommonPrefix(strc))

	fmt.Println(result)
}
func longestCommonPrefix(strs []string) string {
	var longestPrefix string = ""
	var endPrefix = false

	if len(strs) > 0 {
		sort.Strings(strs)
		first := string(strs[0])
		last := string(strs[len(strs)-1])

		for i := 0; i < len(first); i++ {
			if !endPrefix && string(last[i]) == string(first[i]) {
				longestPrefix += string(last[i])
			} else {
				endPrefix = true
			}
		}
	}
	return longestPrefix
}
