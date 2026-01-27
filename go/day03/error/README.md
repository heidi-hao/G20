
# 异常处理
程序运行过程中，可能会有异常，比如：输入参数错误 文件不存在 数据库连接失败等等

## error
异常处理的流派：
+ 通过返回值: resp, error := func()(resp,error)
+ 通过try catch:  try { ... }catch(e) { ... },面向对象的语言居多：java python js

通过一个返回值来处理异常
resp, err := func() (resp,error)
if err != nil {
    fmt.Println("error:",err)
    return nil
}




## 程序怎么抛出error
1.直接返回 底层的error到上层(return error)
return err

2.通过error包来New一个error(errors.New)
files,err := os.ReadDir(path)
if err != nil{
    return errors.New(fmt.Springf("file read error:%s",err))
}

2.直接封装下，封装成自己的error(fmt.Errorf)
files,err := os.ReadDir(path)
if err != nil{
    return fmt.Errorf("File Read Error:%s",err)
}

3.组合异常，在底层的异常上，添加自己定义的异常(errors.Join)
files, err := os.ReadDir(path)
if err != nil{
    return errors.Join(errors.New("walkDir error"),err)
}

异常的比较(使用errors.Is)
var (
    ErrInvalid = errors.New("invald argument")
    ErrPermission = errors.New("permission denied")
    ErrExist = errors.New("file already exists")
    ErrNotExist = errors.New("file dose not exist")
    ErrClosed = errors.New("file already clodes")
)
//遍历目录，如果失败了，就区分“目录不存在”和“其他错误”，分别处理。
if err != nil{
    if errors.Is(err,os.ErrNotExist){
        fmt.Println(ErrFileNotFound)
    }else{
        fmt.Println("%s",err)
    }
    os.Exit(1)
}

## panic (程序奔溃)
panic 非常危险 空指针(NPE)
windows蓝屏这种，特效危险的操作，如果继续运行可能会造成安全隐患，内存的非法访问
A程序的内存(内存地址A，用完后把内存地址释放，OS这一个物理内存 分配给了程序B，A再次访问已经释放的内存地址A)
奔溃的其中一个场景，内存非法访问

go是可以绕开指针类型系统(nil),直接使用unsafe包，访问地址

//panic
//索引4 就是非法内存
arrs := []int{1,2,3}
fmt.Pringln(arrs[4])
panic: runtime error: index out of range [4] with length 3

goroutine 1 [running]:
main.main()
	/Users/yumaojun/Projects/go-course/go20/day03/error/main.go:28 +0x44
exit status 2


## recover
你写一个APIServer 不小心panic：runtime error: index out of range [4] with length 3、程序奔溃 无法继续处理用户请求 你网站有1000个接口 产生这个报错的是其中的一个接口的一个流程 需要阻止程序崩溃，打印日志就行 
捕获panic信号 捕获是一个特殊的逻辑 程序退出前 去检查有没有panic信号 放执行前面和后面都不行 需要放到函数调用结束后执行(Hook)
func recoverHandler() {
    if r := recover(); r != nil {
        fmt.Println("Recoverd from panic:",r)
        os.Exis(1)
    }
}
func main(){
    defer recoverHandler()
    ...
}


## defer(函数Hook)
函数执行完成后，调用你的defer函数进行资源清理或者崩溃恢复
一个流程(协程)里面，可以有个defer语句，声明函数执行完成后的Hook调用



