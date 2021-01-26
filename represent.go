package datautil

import "fmt"

// DataRepresentation 用于显示数据的前部分行。
// 要求传入的数据第一行是标题。
// 传入第一个参数是数据，第二个参数是显示前多少行。
func DataRepresentation(data [][]string, top int) {
	for ix, value := range data {
		for itemIndex, item := range value {
			if itemIndex == 0 {
				//if ix == 0 {
				//	fmt.Println("Head:")
				//} else if ix == 1 {
				//	fmt.Println("Body:")
				//}
				// 第一列开头没有空格
				fmt.Print("| " + item)
			} else {
				fmt.Print(" | " + item)
			}
		}
		fmt.Println(" |")
		if ix == top {
			break
		}
	}
}
