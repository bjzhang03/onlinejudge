package main

import (
	"fmt"
	"math"
)

// 标记一个点
type point struct {
	x int
	y int
}

// 标记一个矩形
type rectangle struct {
	leftx  int
	lefty  int
	rightx int
	righty int
}

// 计算矩形面积
func area(r rectangle) int {
	return int(math.Abs(float64(r.leftx-r.rightx)) * math.Abs(float64(r.lefty-r.righty)))
}

func haspoint(r rectangle, p point) bool {
	result := false
	if r.leftx <= p.x && r.rightx >= p.x && r.lefty <= p.y && r.righty >= p.y {
		result = true
	}
	return result
}

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	result := 0
	// 计算当前的面积
	result = int(math.Abs(float64((D-B)*(C-A))) + math.Abs(float64((G-E)*(H-F))))

	return result
}

func main() {
	fmt.Println(computeArea(-3, 0, 3, 4, 0, -1, 9, 2))

}
