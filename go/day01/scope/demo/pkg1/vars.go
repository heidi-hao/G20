package pkg1

// 导出变量 首字母大写

var (
	Var1 = "This is Var1"
	Var2 = "This is Var2"

	//未导出变量 首字母小写
	var3 = "This is var3"
)

func GetVar3() string {
	return var3
}
