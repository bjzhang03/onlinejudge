package main

import (
	"fmt"
	"math"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	result := ""
	if denominator == 0 {
		result = result + "NaN"
	} else {
		// 处理两个数字符号不一样的情况
		if numerator*denominator < 0 {
			result = result + "-"
		}
		// 使用它们的绝对值来进行计算
		numerator = int(math.Abs(float64(numerator)))
		denominator = int(math.Abs(float64(denominator)))
		// 先处理整数部分的数据
		if numerator > denominator {
			result = result + strconv.Itoa(numerator/denominator)
			numerator = numerator % denominator
		} else {
			result = result + "0"
		}

		if numerator > 0 {
			// 处理小数部分
			save := make(map[int]int)
			result = result + "."
			//fmt.Println("numerator = ", numerator, " , denominator = ", denominator)
			rres := ""
			for numerator > 0 {
				//fmt.Println("save = ", save, " , rres = ", rres, "numerator = ", numerator, " , denominator = ", denominator)
				numerator = numerator * 10
				for numerator < denominator {
					rres = rres + "0"
					numerator = numerator * 10
				}

				if _, ok := save[numerator]; ok {
					// 单独处理多个0的情况
					index := save[numerator]
					for index-1 >= 0 && rres[len(rres)-1] == rres[index-1] {
						rres = rres[:len(rres)-1]
						index--
					}
					left := rres[:index]

					left = left + "("
					left = left + rres[index:]
					left = left + ")"
					result = result + left
					return result
				}

				save[numerator] = len(rres)
				rres = rres + strconv.Itoa(numerator/denominator)
				numerator = numerator % denominator

			}
			result = result + rres
			return result
		}

	}
	return result
}

func main() {

	fmt.Print(fractionToDecimal(-1, 13))

}
