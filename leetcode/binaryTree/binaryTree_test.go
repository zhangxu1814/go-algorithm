package binaryTree

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	//println([]byte(strconv.Itoa(1000)))
	//s := "12345"
	//println(s[0])

	//s := []string{"test"}
	//println(s[0])
	//println(s)
	//a := s[1:][0]
	//println(a)

	//a := "test" + "test"
	//println(a)

	//cur := -1<<31 - 1
	//println(cur)

	//a := []int{5, 1, 1, 2, 0, 0}
	//sortArray(a)
	//fmt.Println(a)

	//print(2 << 2)

	s := []string{}
	s1 := "aaaaaa"
	s = append(s, s1)
	s1 += "bbb"
	s = append(s, s1)
	fmt.Println(s)

}
