package main

func matchParentheses(tmpa uint8, tmpb uint8) bool {
	result := false
	if tmpa == '{' && tmpb == '}' {
		result = true;
	}
	if tmpa == '[' && tmpb == ']' {
		result = true;
	}
	if tmpa == '(' && tmpb == ')' {
		result = true;
	}
	return result;
}

func isValid(s string) bool {
	result := true;
	var stack []uint8
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 {
			if matchParentheses(stack[len(stack)-1], s[i]) {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, s[i])
			}
		} else {
			stack = append(stack, s[i])
		}
	}
	if len(stack) > 0 {
		result = false
	}
	return result
}

func main() {
	isValid("111")

}
