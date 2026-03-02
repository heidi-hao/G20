package main

import "fmt"

func main() {
	// 定义学生和成绩
	students := []string{"张三", "李四", "王五", "赵六", "钱七"}
	scores := []int{85, 92, 78, 95, 88}

	//显示所有学生的信息
	PlayAllInfo(students, scores)
	//统计信息 最高分 平均分 最低分
	AverageScore(students, scores)

}

func PlayAllInfo(students []string, scores []int) {
	for i := 0; i < len(students); i++ {
		fmt.Printf("%d. %s: %d分\n", i+1, students[i], scores[i])

	}

}

func AverageScore(students []string, scores []int) {
	var sum int
	for _, score := range scores {
		sum += score
	}
	fmt.Println(sum)
	fmt.Println("平均分:", float64(sum)/float64(len(students)))

}
