// Create a dynamic 2-dimensioanl slice
// @author Prabhash Rathore

package main

import "fmt"

func create_2d_slice(m, n int) [][]uint8 {
	array := make([][]uint8, m)

	for i := 0; i < m; i++ {
		array[i] = make([]uint8, n)
		for j := 0; j < n; j++ {
			array[i][j] = uint8(i + j)
		}
	}

	return array
}

func main() {
	output := create_2d_slice(4, 3)
	fmt.Println("2d slice is:")
	fmt.Println(output)
}
