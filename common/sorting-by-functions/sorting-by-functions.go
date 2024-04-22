// / 自定义规则排序
package main

import (
	"fmt"
	"sort"
)

// / sort 排序 也是调用被排序类型的 Len/Less/Swap 等函数
// / 我们重写这部分函数即可
type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
