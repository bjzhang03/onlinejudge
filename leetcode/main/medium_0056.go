package main

// Definition for an interval.
type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	result := []Interval{}
	if len(intervals) > 0 {
		flag := true
		usedinterval := make(map[int]bool)
		for i := 0; i < len(intervals); i++ {
			for j := i + 1; j < len(intervals); j++ {
				// 以前没有被使用过的数据
				if _, ok := usedinterval[j]; !ok {
					if intervals[j].Start > intervals[i].Start && intervals[j].End < intervals[i].End {
						usedinterval[j] = true
					}
				}

			}
			usedinterval[i] = true
		}

		if !flag {
			result = merge(result)
		}
	}
	return result
}

func main() {

}
