package main

import "fmt"

func climbStairs(n int) int {
	result := 0
	if 1 == n {
		result = 1
	} else if 2 == n {
		result = 2
	} else {
		tmpa := 1
		result = 2
		for i := 3; i <= n; i++ {
			tmp := tmpa + result
			tmpa = result
			result = tmp
		}
	}
	return result
}
func main() {

	fmt.Println(climbStairs(4))

}
