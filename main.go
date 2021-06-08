package main

import (
	"fmt"
	"math"
	"reflect"
	"sort"
)

func main() {
	//1544
	fmt.Println(math.Abs(float64('P' - 'p')))

	//42
	m := make(map[int][]int)
	// m[0] = append(m[0], 0)
	// m[1] = append(m[1], 1)
	// m[0] = append(m[0], 2)
	// m[2] = append(m[2], 3)
	// m[1] = append(m[1], 4)
	// m[0] = append(m[0], 5)
	// m[1] = append(m[1], 6)
	// m[3] = append(m[3], 7)
	// m[2] = append(m[2], 8)
	// m[1] = append(m[1], 9)
	// m[2] = append(m[2], 10)
	// m[1] = append(m[1], 11)
	m[4] = append(m[4], 0)
	m[2] = append(m[2], 1)
	m[0] = append(m[0], 2)
	m[3] = append(m[3], 3)
	m[2] = append(m[2], 4)
	m[5] = append(m[5], 5)
	h := []int{}
	for k := range m {
		h = append(h, k)
	}
	sort.Sort(sort.IntSlice(h))
	for _, v := range h {
		println(v)
	}

	args := reflect.New(reflect.TypeOf(5))
	fmt.Println("args:", args.Type())
}
