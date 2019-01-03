package main

import "fmt"

func judge(s string, t string) bool {
	result := true
	save := make(map[uint8]uint8)
	sindex := 0
	tindex := 0
	for sindex < len(s) {
		if _, ok := save[s[sindex]]; !ok {
			save[s[sindex]] = t[tindex]
		} else if save[s[sindex]] != t[tindex] {
			result = false
			break
		}
		sindex++
		tindex++
	}
	return result
}

func isIsomorphic(s string, t string) bool {
	result := true
	if len(s) != len(t) {
		result = false
	} else {

		result = judge(s, t) && judge(t, s)

	}
	return result
}

func main() {

	fmt.Println(isIsomorphic("ag", "aa"))

}
