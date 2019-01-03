package main

import (
	"fmt"
	"sort"
	"strconv"
)

func compareStr(itema string, itemb string) bool {
	result := false
	if len(itema) < len(itemb) {
		//fmt.Println("<")
		tmpa := itema
		for len(itema) < len(itemb) {
			itema = itema + string(itema[0])
		}
		for i := 0; i < len(itema); i++ {
			if itema[i] > itemb[i] {
				result = true
				return result
			} else if itema[i] < itemb[i] {
				result = false
				return result
			}
		}
		// {121,12}
		// {8308,830}
		//fmt.Println("itema = ", itema)
		if tmpa[len(tmpa)-1] < tmpa[0] {
			return false
		}
		return true
	} else if len(itema) == len(itemb) {
		//fmt.Println("=")
		for i := 0; i < len(itema); i++ {
			if itema[i] > itemb[i] {
				result = true
				return result
			} else if itema[i] < itemb[i] {
				result = false
				return result
			}
		}
	} else if len(itema) > len(itemb) {
		//fmt.Println(">")
		tmpb := itemb
		for len(itema) > len(itemb) {
			itemb = itemb + string(itemb[0])
		}
		for i := 0; i < len(itema); i++ {
			if itema[i] > itemb[i] {
				result = true
				return result
			} else if itema[i] < itemb[i] {
				result = false
				return result
			}
		}
		// {121,12}
		// {8308,830}
		// fmt.Println("tmpb = ", tmpb)
		if tmpb[len(tmpb)-1] <= tmpb[0] {
			return true
		}
		return false
	}
	// fmt.Println("itema = ", itema, " ,itemb = ", itemb, " ,result = ", result)
	return result
}

func largestNumber(nums []int) string {
	result := ""
	if len(nums) > 0 {
		numsstr := []string{}
		sum := 0
		for i := 0; i < len(nums); i++ {
			numsstr = append(numsstr, strconv.Itoa(nums[i]))
			sum = sum + nums[i]
		}

		if sum > 0 {
			sort.Slice(numsstr[:], func(i, j int) bool {
				return compareStr(numsstr[i], numsstr[j])
			})

			//fmt.Println(numsstr)
			for i := 0; i < len(numsstr); i++ {
				result = result + numsstr[i]
			}
		} else if sum == 0 {
			result = "0"
		}
	}
	return result
}

func main() {

	//nums := []int{121, 12}
	nums := []int{830, 8308}

	fmt.Println(largestNumber(nums))

}
