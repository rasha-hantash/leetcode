package main

func isValid(s string) bool {
	var stack []rune

	parentMap := map[rune]rune{
		']': '[',
		')': '(',
		'}': '{',
	}

	for _, char := range s {
		if char == '{' || char == '[' || char == '(' {
			stack = append(stack, char)
		}
		if _, exist := parentMap[char]; exist {
			if len(stack) > 0 && stack[len(stack)-1] == parentMap[char] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
