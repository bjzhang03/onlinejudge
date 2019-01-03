package main

import (
	"fmt"
	"strings"
)

func isValidSerialization(preorder string) bool {
	result := false
	if len(preorder) > 0 {
		strs := strings.Split(preorder, ",")
		fmt.Println(strs)

	}
	return result
}

func main() {

	fmt.Println(isValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#"))

}
