package main

import (
	"fmt"
	"userapp/models"
	"userapp/utils"
)

func main() {
	fmt.Println("hello user app")

	user := models.NewUser("张三", "zhangsan@example.com", 25)

	//验证数据
	if !utils.ValidateEmail(user.Email) {
		fmt.Println("邮箱格式不正确")
		return
	}

	if !utils.ValidateAge(user.Age) {
		fmt.Println("年纪不符合")
		return
	}

	fmt.Printf("用户创建成功：%+v\n", user)
}
