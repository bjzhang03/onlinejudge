package main

func isAnagram(s string, t string) bool {
	result := true
	if len(s) != len(t) {
		result = false
	} else {
		smap := make(map[uint8]int)
		tmap := make(map[uint8]int)

		for i := 0; i < len(s); i++ {
			if _, ok := smap[s[i]]; !ok {
				smap[s[i]] = 1
			} else {
				smap[s[i]]++
			}
		}

		for i := 0; i < len(t); i++ {
			if _, ok := tmap[t[i]]; !ok {
				tmap[t[i]] = 1
			} else {
				tmap[t[i]]++
			}
		}

		if len(smap) != len(tmap) {
			result = false
		} else {
			for k, v := range smap {
				if _, ok := tmap[k]; !ok {
					result = false
					break
				}
				if v != tmap[k] {
					result = false
					break
				}
			}
		}
	}
	return result

}

func main() {

}
