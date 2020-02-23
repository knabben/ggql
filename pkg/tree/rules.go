package tree

import (
	"fmt"
	"sort"
)

func Run(a int64) int64 {
	if a == 0 {
		return 0
	}
	fmt.Println(a)
	return Run(a - 1)
}


func Run1(s []int, i int) bool {
	sort.Ints(s)
	fmt.Println(s)

	middle := len(s) / 2
	if middle == 0 {
		return false
	}

	if s[middle] == i {
		return true
	}

	if i > s[middle] {
		return Run1(s[middle:len(s)], i)
	} else if i < s[middle] {
		return Run1(s[:middle], i)
	}
	return false
}
