// / 排序
package main

import (
	"fmt"
	"sort"
)

func main() {
	// 数组切片
	strs := []string{"m", "aa", "c", "b"}
	// 排序
	sort.Strings(strs)
	fmt.Println(strs)

	ints := []int{5, 1, 3, 8}
	// 对 int 数组进行切片
	sort.Ints(ints)
	fmt.Println(ints)

	// 检查数据是否有序
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
