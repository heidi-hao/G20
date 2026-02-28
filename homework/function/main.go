package main

import "fmt"

func main() {

	//通过内存地址访问Tag的值
	b1 := Book{
		Title:  "hello world",
		Author: "heidi",
		Page:   100,
		Tag: []string{
			"abc",
			"def",
			"hjk",
		},
	}
	fmt.Println("b1:", b1.Tag)

	//使用二维切片表示一组学生的各科成绩，计算这组学生的学科平均分
	//scores: 所有的学生
	//scores[i]:第i个学生的所有成绩
	//socres[i][j]:第i个学生，第j门课的成绩 行：学生 列：成绩
	scores := [][]int{
		{88, 88, 90},
		{66, 99, 94},
		{75, 84, 98},
		{93, 77, 66},
	}

	subjectCnt := len(scores[0]) //科目数 scores[0] 第一个学生的成绩 有三门科目
	fmt.Println(subjectCnt)
	studentCnt := len(scores) //学生数 有四行 说明有四个学生
	fmt.Println(studentCnt)

	//创建一个长度为科目数的切片，用来累加每科总分 sum = [0,0,0] subjectCnt=3
	sum := make([]int, subjectCnt)
	fmt.Println(sum)

	//双重循环：核心逻辑
	//外层循环(学生) i=0 第一个学生 i=1 第二个学生  i=2 第三个学生 i=3 第四个学生
	//内层循环(科目) j=0 数学 j=1 语文 j=2 英语
	for i := 0; i < studentCnt; i++ {
		for j := 0; j < subjectCnt; j++ {
			sum[j] += scores[i][j] // //当i=0的时候 要j=0，1，2，3相加
		}

	}
	//计算平均分并输出
	for j := 0; j < subjectCnt; j++ {
		fmt.Println("第%d科平均分:%.2f\n", j+1,
			float64(sum[j])/float64(studentCnt),
		)
	}

}

type Book struct {
	Title  string
	Author string
	Page   uint
	Tag    []string
}
