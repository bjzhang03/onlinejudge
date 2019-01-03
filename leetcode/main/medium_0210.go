package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	result := []int{}
	if numCourses > 0 {
		save := make([][]int, numCourses)
		for i := 0; i < numCourses; i++ {
			save[i] = make([]int, numCourses)
		}
		for _, item := range prerequisites {
			save[item[1]][item[0]] = 1
		}
		//fmt.Println("save = ", save)
		dgzero := []int{} // degree zero
		// 找到入度为0的节点
		for i := 0; i < numCourses; i++ {
			flag := true
			for j := 0; j < numCourses; j++ {
				if save[j][i] == 1 {
					flag = false
					break
				}
			}
			if flag {
				// 添加第i个节点进来
				dgzero = append(dgzero, i)
			}
		}
		for len(dgzero) > 0 {
			// 将某一行的数据设置为0，也就是抹去这一个节点的出度
			candidatenode := []int{}
			for i := 0; i < numCourses; i++ {
				if save[dgzero[0]][i] == 1 {
					candidatenode = append(candidatenode, i)
				}
				save[dgzero[0]][i] = 0
			}
			// 标记某个节点已经被使用过了
			result = append(result, dgzero[0])
			// 去掉当前的已经使用过的节点的数据
			dgzero = dgzero[1:]
			// 找到入度为0的节点
			for _, candidate := range candidatenode {
				flag := true
				for i := 0; i < numCourses; i++ {
					if save[i][candidate] == 1 {
						flag = false
					}
				}
				if flag {
					dgzero = append(dgzero, candidate)
				}
			}
		}
		// 出现无法选修的情况则把所有数据置空
		if len(result) != numCourses {
			result = []int{}
		}
	}
	return result
}

func main() {

}
