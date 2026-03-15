package main

import "fmt"

func slice() {
	s := []int{1, 2, 3}
	t := s
	t[0] = 99
	fmt.Println(s[0], t[0])
}

func zeroValues() {
	var x int
	var s string
	var b bool
	fmt.Println(x, s, b)
}

func rangeExample() {
	nums := []int{10, 20, 30}
	for i, v := range nums {
		fmt.Printf("index=%d value=%d\n", i, v)
	}
}

func variadicSum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	zeroValues()
	slice()
	rangeExample()
	fmt.Println(variadicSum(1, 2, 3, 4, 5))
	r := Rectangle{3.0, 4.0}
	fmt.Println(r.Area())
}
