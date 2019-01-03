package main

// https://blog.csdn.net/u014673347/article/details/46944469
func rangeBitwiseAnd(m int, n int) int {
	count := 0
	for n != m {
		m >>= 1
		n >>= 1
		count++
	}
	return (m << uint8(count))
}

func main() {

}
