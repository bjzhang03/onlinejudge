package main

import "fmt"

func numTrees(n int) int {
	result := 0
	if n > 0 {
		save := make([]int, n+1)
		save[0] = 1
		for i := 1; i <= n; i++ {
			currentsum := 0
			for left := 0; left < i; left++ {
				currentsum = currentsum + save[left]*save[i-1-left]
			}
			save[i] = currentsum
		}
		//fmt.Println(save)
		result = save[n]
	}
	return result
}

func main() {
	fmt.Println(numTrees(30))

}
