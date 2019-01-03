package main

func findTheDifference(s string, t string) byte {
	if len(t) > len(s) {
		countS := map[uint8]int{'a': 0, 'b': 0, 'c': 0, 'd': 0, 'e': 0, 'f': 0, 'g': 0, 'h': 0, 'i': 0, 'j': 0, 'k': 0, 'l': 0, 'm': 0, 'n': 0, 'o': 0, 'p': 0, 'q': 0, 'r': 0, 's': 0, 't': 0, 'u': 0, 'v': 0, 'w': 0, 'x': 0, 'y': 0, 'z': 0}
		countT := map[uint8]int{'a': 0, 'b': 0, 'c': 0, 'd': 0, 'e': 0, 'f': 0, 'g': 0, 'h': 0, 'i': 0, 'j': 0, 'k': 0, 'l': 0, 'm': 0, 'n': 0, 'o': 0, 'p': 0, 'q': 0, 'r': 0, 's': 0, 't': 0, 'u': 0, 'v': 0, 'w': 0, 'x': 0, 'y': 0, 'z': 0}

		for i := 0; i < len(s); i++ {
			countS[s[i]]++
		}
		for i := 0; i < len(t); i++ {
			countT[t[i]]++
		}
		for k, v := range countT {
			if v > countS[k] {
				return byte(v)
			}
		}
	}
	return 1
}

func main() {

}
