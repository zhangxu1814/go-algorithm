package num_decoding

//TODO need fix bug

// func numDecodings(s string) int {
// 	if s == "" {
// 		return 0
// 	}
// 	arr := []byte(s)
// 	if arr[0] == '0' {
// 		return -1
// 	}
// 	count := 0
// 	if len(arr) == 1 {
// 		count = 1
// 	} else {
// 		if arr[1] == '0' || arr[0] > '2' || arr[0] == 2 && arr[1] > '6' {
// 			count = 1 + numDecodings(string(arr[1:]))
// 		} else {
// 			count = 2 + numDecodings(string(arr[1:])) + numDecodings(string(arr[2:]))
// 		}
// 	}

// 	return count
// }
