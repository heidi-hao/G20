package main

import (
	"fmt"
)

func main() {
	//统计学生的成绩
	scores := []int{85, 92, 78, 95, 88, 73, 90, 85, 79, 91}

	var sum, max, min int
	max = scores[0]
	min = scores[0]

	for _, score := range scores {
		sum += score
		if score > max {
			max = score
		}
		if score < max {
			min = score
		}
	}

	average := float64(sum) / float64(len(scores))
	fmt.Printf("平均分:%.2f\n", average)
	fmt.Printf("最高分:%d\n", max)
	fmt.Printf("最高分:%d\n", min)

	for i := 0; i < len(scores)-1; i++ {
		for j := 0; j < len(scores)-1-i; j++ {
			if scores[j] > scores[j+1] {
				scores[j+1], scores[j] = scores[j], scores[j+1]
			}
		}
	}
	fmt.Println(scores)

}
