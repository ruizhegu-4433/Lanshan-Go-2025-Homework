package main

import "fmt"

func main() {
	fmt.Println("input numbers(input 0 to to stop)")
	var a []int
	result := make(map[int]int)
	for {
		var num int
		fmt.Scan(&num)
		a = append(a, num)
		if num == 0 {
			break
		}
	}
	result = arr_to_map(a)
	fmt.Println(result)
}
func arr_to_map(a []int) map[int]int {
	counts := make(map[int]int)
	for i := 0; i < len(a); i++ {
		var num int
		count := 0
		num = a[i]
		for _, numbers := range a {
			if num == numbers {
				count++
			}
			counts[num] = count
		}
	}
	return counts
}
