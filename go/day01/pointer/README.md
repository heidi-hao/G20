#指针概念

## 场景

+ func sum(a,b) total :这种没有指针参与
+ func sum(a,b,sum): sum 是一个变量，用于接收函数的处理结果
+ fmt.Scan: 接收用户输入的参数
+ flag.Parse(): 接收命令行参数
+ Restful API Handler(req,*resp):resp变量必须是指针，用于接收请求处理后的结果



 ## 概念

 Go里面的参数 都是值传递 不是引用传递