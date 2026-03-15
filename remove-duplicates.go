# Write a program to remove duplicate integers from a slice
package main

import "fmt"

func removeDuplicates(s []int) []int {
	m := make(map[int]bool)
	t := []int{}
	for _, i := range s {
		_, ok := m[i]
		if !ok {
			t = append(t, i)
			m[i] = true
			fmt.Printf("Adding %d in map\n", i)
		}
	}
	return t
}

func main() {
	sl := []int{1, 2, 3, 1, 2, 3}
	fmt.Println(removeDuplicates(sl))
}
