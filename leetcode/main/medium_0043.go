package main

import (
	"fmt"
	"strconv"
)

func addstr(a string, b string) string {
	result := "0"
	if len(a) > 0 && len(b) > 0 {
		result = ""
		sa := len(a) - 1
		sb := len(b) - 1
		carry := 0
		for sa >= 0 && sb >= 0 {
			da, _ := strconv.Atoi(string(a[sa]))
			db, _ := strconv.Atoi(string(b[sb]))

			currentsum := (carry + da + db) % 10
			carry = (carry + da + db) / 10
			result = strconv.Itoa(currentsum) + result
			sa--
			sb--
		}

		for sa >= 0 {
			da, _ := strconv.Atoi(string(a[sa]))

			currentsum := (carry + da) % 10
			carry = (carry + da) / 10
			result = strconv.Itoa(currentsum) + result
			sa--
		}

		for sb >= 0 {
			db, _ := strconv.Atoi(string(b[sb]))

			currentsum := (carry + db) % 10
			carry = (carry + db) / 10
			result = strconv.Itoa(currentsum) + result
			sb--
		}

		if carry > 0 {
			result = strconv.Itoa(carry) + result
		}
	}
	return result
}

func multiplystr(nums string, p uint8) string {
	result := "0"
	if len(nums) > 0 && p > '0' {
		sn := len(nums) - 1
		carry := 0
		result = ""
		dp, _ := strconv.Atoi(string(p))
		for sn >= 0 {
			dn, _ := strconv.Atoi(string(nums[sn]))
			currentsum := (carry + dn*dp) % 10
			carry = (carry + dn*dp) / 10
			result = strconv.Itoa(currentsum) + result
			sn--
		}
		// 处理进位问题
		if carry > 0 {
			result = strconv.Itoa(carry) + result
		}
	}
	return result

}

func multiply(num1 string, num2 string) string {
	result := "0"
	if len(num1) > 0 && len(num2) > 0 {
		// 第一个数字一定是不为0的
		if len(num2) > len(num1) {
			tmp := num2
			num2 = num1
			num1 = tmp
		}
		result = addstr(result, result)
		for i := len(num2) - 1; i >= 0; i-- {
			tmp := multiplystr(num1, num2[i])
			for j := len(num2) - 1; j > i; j-- {
				tmp = tmp + "0"
			}
			// fmt.Println("tmp = ", tmp)
			result = addstr(result, tmp)
		}
	}
	return result
}

func main() {

	fmt.Println(addstr("9", "00"))

	fmt.Println(multiplystr("123", '0'))

	fmt.Println(multiply("0", "1"))

}
