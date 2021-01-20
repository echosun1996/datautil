package datautil

import "fmt"

// DataRepresentation 用于显示二维数组前‘top’条数据的方法
// 要求传入的数据第一行是标题。
// 传入第一个参数是数据，第二个参数是显示前多少行。
func DataRepresentation(data [][]string, top int) {
	if top < 0 {
		panic("'top' should >0!")
	}
	if top >= len(data) {
		panic("'top' should smaller than data's size!")
	}
	for ix, value := range data {
		for itemIndex, item := range value {
			if itemIndex == 0 {
				if ix == 0 {
					fmt.Print("Head:\n|   ")
				} else if ix == 1 {
					fmt.Println("Body:")
				}
				if itemIndex == 0 && ix != 0 {
					fmt.Print("| ", ix)
				}
			}
			fmt.Print(" | " + item)

		}
		fmt.Println(" |")
		if ix == top {
			break
		}
	}
}
