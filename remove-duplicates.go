# Write a program to remove duplicate integers from a slice
package main

import "fmt"

func removeDuplicates(s []int) []int {
    // 1. Pre-allocate capacity to match input length
    // length is 0, but capacity is len(s)
    t := make([]int, 0, len(s)) 
    
    // 2. Use struct{} in the map to save memory 
    // map[int]struct{} takes 0 bytes for the value
    m := make(map[int]struct{}) 
    
    for _, i := range s {
        if _, exists := m[i]; !exists {
            m[i] = struct{}{}
            t = append(t, i) // No reallocations happen here!
        }
    }
    return t
}

func main() {
	sl := []int{1, 2, 3, 1, 2, 3}
	fmt.Println(removeDuplicates(sl))
}
