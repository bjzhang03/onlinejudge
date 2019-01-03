package main

import "fmt"

func add(a uint8, b uint8, carry uint8) (uint8, uint8) {
	result := uint8(0)
	// 初始化map数据
	Count := make(map[uint8]int)
	Count['1'] = 0
	Count['0'] = 0
	// 记录1，0出现的次数
	Count[a] = Count[a] + 1
	Count[b] = Count[b] + 1
	Count[carry] = Count[carry] + 1
	// 进行计算的算法
	if 3 == Count['1'] {
		result = '1'
		carry = '1'
	} else if 2 == Count['1'] {
		result = '0'
		carry = '1'
	} else if 1 == Count['1'] {
		result = '1'
		carry = '0'
	} else {
		result = '0'
		carry = '0'
	}
	// 返回结果
	return result, carry
}

func addBinary(a string, b string) string {
	result := ""
	starta := len(a) - 1
	startb := len(b) - 1
	carry := uint8('0')
	for ; starta >= 0 && startb >= 0;
	{
		tmp, tmpcarry := add(a[starta], b[startb], carry)
		result = string(tmp) + result
		carry = tmpcarry
		starta--
		startb--
	}

	for ; starta >= 0; {
		tmp, tmpcarry := add(a[starta], '0', carry)
		result = string(tmp) + result
		carry = tmpcarry
		starta--
	}

	for ; startb >= 0; {
		tmp, tmpcarry := add(b[startb], '0', carry)
		result = string(tmp) + result
		carry = tmpcarry
		startb--
	}

	if carry == '1' {
		result = string(carry) + result
	}
	return result

}
func main() {
	a := "11"
	b := "1"

	fmt.Println(addBinary(a, b))

}
