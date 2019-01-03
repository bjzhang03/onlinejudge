package main

import "fmt"

var table = []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536, 131072, 262144, 524288, 1048576, 2097152, 4194304,
	8388608, 16777216, 33554432, 67108864, 134217728, 268435456, 536870912, 1073741824, 2147483648, 4294967296, 8589934592, 17179869184, 34359738368,
	68719476736, 137438953472, 274877906944, 549755813888, 1099511627776, 2199023255552, 4398046511104, 8796093022208, 17592186044416, 35184372088832,
	70368744177664, 140737488355328, 281474976710656, 562949953421312, 1125899906842624, 2251799813685248, 4503599627370496, 9007199254740992, 18014398509481984,
	36028797018963968, 72057594037927936, 144115188075855872, 288230376151711744, 576460752303423488, 1152921504606846976, 2305843009213693952, 4611686018427387904}

func bytesToint(b []byte) int {
	result := 0
	for i := len(b) - 1; i >= 0; i-- {
		if b[i] != 0 {
			result = result + table[len(b)-1-i]
		}
	}
	//fmt.Println("result = ", result)
	return result
}

func grayCode(n int) []int {
	if n > 0 {
		result := []int{}
		save := [][]byte{{0}, {1}}
		for i := 1; i < n; i++ {
			nextsaveleft := [][]byte{}
			for j := 0; j < len(save); j++ {
				tmp := []byte{}
				tmp = append(tmp, save[j]...)
				tmp = append(tmp, 0)
				nextsaveleft = append(nextsaveleft, tmp)
			}
			nextsaveright := [][]byte{}
			for j := 0; j < len(save); j++ {
				tmp := []byte{}
				tmp = append(tmp, save[j]...)
				tmp = append(tmp, 1)
				nextsaveright = append(nextsaveright, tmp)
			}
			for j := len(nextsaveright) - 1; j >= 0; j-- {
				nextsaveleft = append(nextsaveleft, nextsaveright[j])
			}
			save = [][]byte{}
			save = append(save, nextsaveleft...)
			//fmt.Println("save = ", save)
		}

		for i := 0; i < len(save); i++ {
			result = append(result, bytesToint(save[i]))
		}
		return result
	}
	return []int{0}

}

func main() {
	fmt.Println(grayCode(5))
}
