package main

import "fmt"

func main() {
	// 定义学生和成绩
	students := []string{"张三", "李四", "王五", "赵六", "钱七"}
	scores := []int{85, 92, 78, 95, 88}

	//显示所有学生的成绩
	for i := 0; i < len(students)-1; i++ {
		fmt.Printf("%d. %s: %d分\n", i+1, students[i], scores[i])
	}

	// 2. 计算并显示：平均分、最高分、最低分
	var sum, max, min int
	max = scores[0]
	min = scores[0]

	for _, socre := range scores {
		sum += socre
		if socre > max {
			max = socre

		}
		if socre < min {
			min = socre
		}
	}
	average := float64(sum) / float64(len(students))
	fmt.Println("最高分:", max)
	fmt.Println("最低分:", min)
	fmt.Println("平均分:", average)

	//3.统计各分数段的人
	excellent := 0 // 90-100
	good := 0      // 80-89
	medium := 0    // 70-79
	pass := 0      // 60-69
	fail := 0      // <60
	for _, socre := range scores {
		if socre >= 90 && socre <= 100 {
			excellent++
		} else if socre >= 80 {
			good++
		} else if socre >= 70 {
			medium++
		} else if socre >= 60 {
			pass++
		} else if socre < 60 {
			fail++
		}
		fmt.Println("90-100:优秀 ", excellent)
		fmt.Println("80-89:良好 ", good)
		fmt.Println("70-79:中等 ", medium)
		fmt.Println("60-69:及格 ", pass)
		fmt.Println("<60:不及格 ", fail)
	}

	//成绩排名
	for i := 0; i < len(scores)-1; i++ {
		for j := 0; j < len(scores)-1-i; j++ {
			if scores[j] < scores[j+1] {
				scores[j], scores[j+1] = scores[j+1], scores[j]
				//同时交换学生
				students[j], students[j+1] = students[j+1], students[j]
			}
		}
	}
	fmt.Println(scores)
	for i := 0; i < len(scores); i++ {
		fmt.Printf("%d. %s: %d分\n", i+1, students[i], scores[i])
	}

}
