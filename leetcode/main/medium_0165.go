package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	result := 0
	if len(version1) > 0 && len(version2) > 0 {
		ver1str := strings.Split(version1, ".")
		ver2str := strings.Split(version2, ".")

		ver1num := []int{}
		ver2num := []int{}

		for i := 0; i < len(ver1str); i++ {
			tmp, _ := strconv.Atoi(ver1str[i])
			ver1num = append(ver1num, tmp)
		}

		for i := 0; i < len(ver2str); i++ {
			tmp, _ := strconv.Atoi(ver2str[i])
			ver2num = append(ver2num, tmp)
		}

		if len(ver1num) > len(ver2num) {
			for len(ver1num) > len(ver2num) {
				ver2num = append(ver2num, 0)
			}

		} else if len(ver2num) > len(ver1num) {
			for len(ver2num) > len(ver1num) {
				ver1num = append(ver1num, 0)
			}
		}

		//fmt.Println(ver1num, ver2num)

		for i := 0; i < len(ver1num); i++ {
			if ver1num[i] > ver2num[i] {
				result = 1
				return result
			} else if ver1num[i] < ver2num[i] {
				result = -1
				return result
			}
		}

	}
	return result
}

func main() {

	fmt.Println(compareVersion("07.5.2.4.0.0", "7.5.2.4"))

}
