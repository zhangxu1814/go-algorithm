package impl_strstr

import (
	"reflect"
)

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	l1 := len(haystack)
	l2 := len(needle)
	if l1 < l2 {
		return -1
	}
	n := needle[:]
	i := 0
	for ; i <= l1-l2; i++ {
		if reflect.DeepEqual(haystack[i:l2+i], n) {
			return i
		}
	}

	return -1
}

//TODO
// func strStrOptimization(haystack string, needle string) int {

// }

//TODO
// func strStrUsingKMP(haystack string, needle string) int {

// }
