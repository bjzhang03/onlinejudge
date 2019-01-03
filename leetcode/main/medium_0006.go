package main

import "fmt"
// 更牛逼的找规律的方法
// https://blog.csdn.net/langmanqishizaijia/article/details/51154790
func convert(s string, numRows int) string {
	result := ""
	if len(s) == 0 || len(s) == 1 || numRows == 0 || numRows == 1 {
		// 字符长度为0
		result = s
	} else if len(s) > 0 && numRows == 2 {
		tmp1 := ""
		for i := 0; i < len(s); i = i + 2 {
			tmp1 = tmp1 + string(s[i])
		}
		tmp2 := ""
		for i := 1; i < len(s); i = i + 2 {
			tmp2 = tmp2 + string(s[i])
		}
		result = tmp1 + tmp2
	} else if len(s) > 0 && numRows > 2 {
		save := make([][]uint8, numRows)
		for i := 0; i < numRows; i++ {
			save[i] = make([]uint8, len(s)/(numRows-2)+1)
		}
		sindex := 0
		for j := 0; j < len(s); j++ {
			if sindex >= len(s) {
				break
			}
			for i := 0; i < numRows; i++ {
				if sindex < len(s) {
					save[i][j] = s[sindex]
					sindex++
				}
			}
			// 到下一列的数据
			j++
			for i := numRows - 2; i > 0; i-- {
				if sindex < len(s) {
					save[i][j] = s[sindex]
					sindex++
				}
			}
		}
		for i := 0; i < numRows; i++ {
			for j := 0; j < len(s); j++ {
				if save[i][j] > 0 {
					result = result + string(save[i][j])
				}
			}
		}
	}
	return result
}

func main() {

	fmt.Println(convert("PAYPALISHIRING", 3))

}
