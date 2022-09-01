package parnoe_programirovanie

import (
	"bufio"
	"fmt"
	"os"
)

func Start() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscan(in, &testCount)

	for i := 0; i < testCount; i++ {
		var arr []int
		var n int
		fmt.Fscan(in, &n)
		for j := 0; j < n; j++ {
			var a int
			fmt.Fscan(in, &a)
			arr = append(arr, a)
		}
		vgihhh := madgic(arr)
		fmt.Fprintln(out, vgihhh)
	}
}
func madgic(arr []int) []int {
	return arr
}

//in
//
//3
//6
//2 1 3 1 1 4
//2
//5 5
//8
//1 4 2 5 4 2 6 3

//out
//
//1 2
//3 6
//4 5
//
//1 2
//
//1 3
//2 5
//4 7
//6 8
