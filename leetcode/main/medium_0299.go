package main

import (
	"fmt"
	"strconv"
)

func getHint(secret string, guess string) string {
	result := ""
	countA := 0
	countB := 0
	if len(secret) > 0 && len(guess) > 0 {
		used := make(map[int]bool)
		hadsec := "" // handled secret
		hadgue := "" // handled guess

		// 找出A的个数,同时构造出B的计算数据
		for i := 0; i < len(secret); i++ {
			if secret[i] == guess[i] {
				countA++
				used[i] = true
			} else {
				hadsec = hadsec + string(secret[i])
				hadgue = hadgue + string(guess[i])
			}
		}

		cousec := map[uint8]int{'0': 0, '1': 0, '2': 0, '3': 0, '4': 0, '5': 0, '6': 0, '7': 0, '8': 0, '9': 0} // count secret
		cougue := map[uint8]int{'0': 0, '1': 0, '2': 0, '3': 0, '4': 0, '5': 0, '6': 0, '7': 0, '8': 0, '9': 0} // count guess

		for i := 0; i < len(hadsec); i++ {
			cousec[hadsec[i]]++
			cougue[hadgue[i]]++
		}

		for k, _ := range cousec {
			if cousec[k] < cougue[k] {
				countB = countB + cousec[k]
			} else {
				countB = countB + cougue[k]
			}
		}
	}

	result = result + strconv.Itoa(countA) + "A" + strconv.Itoa(countB) + "B"
	return result
}

func main() {
	fmt.Println(getHint("1807", "7810"))

}
