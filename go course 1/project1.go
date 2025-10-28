package main

import "fmt"

func main() {
	sum := 0
	count := 1
	var avg float32
	for {
		var num int
		fmt.Printf("please input number(input 0 to stop)")
		fmt.Scan(&num)
		sum = sum + num
		if num == 0 {
			break
		}
		count++
	}
	avg = average(sum, count)
	if avg >= 60 {
		fmt.Printf("average score is%.2f ,qualified", avg)
	} else {
		fmt.Printf("average score is%.2f ,unqualified", avg)
	}
}
func average(sum int, count int) float32 {
	return float32(sum) / float32(count)
}
