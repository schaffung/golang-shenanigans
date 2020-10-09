package main

import "fmt"

func reverse(s []int) {
	// Just takes int values

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	// just a slice...
	i := []int{1, 2, 3, 4}

	// reversing the slice...
	reverse(i)

	fmt.Println("So the result of this reverse operations is : ", i)
}
