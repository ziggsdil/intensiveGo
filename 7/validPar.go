package main

import "fmt"

func main() {
	str := "(){}[]"
	fmt.Println(isValid(str))
}

func isValid(s string) bool {
	pairs := map[byte]byte{
		']': '[',
		'}': '{',
		')': '(',
	}

	stack := make([]byte, 0)
	for _, ch := range []byte(s) {
		pair, ok := pairs[ch]
		if !ok {
			stack = append(stack, ch)
			continue
		}
		if len(stack) == 0 {
			return false
		}
		if stack[len(stack)-1] != pair {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}
