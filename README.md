# go-learning
* 开发环境 GOPATH
    1. go 1.8 以后Unix默认$HOME/go
    2. Windows默认%USERPROFILE%/go
    3. Mac上修改～/.bash_profile设置
* go command
    1. go version               (查看go版本)
    2. go run hello_world.go    (运行go)
    3. go build hello_world.go  (编译go) 
* 应用程序入口
    1. 必须是main包: package main
    2. 必须是main方法： func main()
    3. 文件名不一定是main.go
* 退出返回值(与其他语言区别)
    1. Go中main函数不支持任何返回值
    2. 通过os.Exit返回状态
* 获取命令行参数(与其他语言区别)
    1. main函数不支持传入参数func main(arg []string)
    2. 在程序中直接通过os.Args获取命令行参数
* 编写测试程序
    1. 源码文件以_test结尾：xxx_test.go
    2. 测试方法名以Test开头:func TestXXX(t *testing.T){}
* 变量赋值(与其他主要语言差异)
    1. 赋值可以进行自动类型推断
    2. 在一个赋值语句中可以对多个变量进行同时赋值
* 常量定义(与其他主要语言差异)
    1. 快速设置连续值(iota)
* 基本数据类型
    1. bool
    2. string
    3. int int8 int16 int32 int64
    4. uint uint8 uint16 uint32 uint64 uintptr
    5. byte    // alias for uint8
    6. rune    // alias for int32, represents a Unicode code point
    7. float32 float64
    8. complex64 complex128
* 类型转换(与其他主要语言差异)
    1. Go语言不允许隐式类型转换
    2. 别名和原有类型也不能进行隐式类型转换
* 类型预定义值
    1. math.MaxInt64
    2. math.MaxFloat64
    3. math.MaxUint32
* 指针类型(与其他主要语言差异)
    1. 不支持指针运算
    2. string是值类型,其默认初始值为空字符串,而不是nil
* 算数运算符
    1. + - * / % ++ --
    2. Go语言没有前置++, --
* 比较运算符
    1. == != > < >= <=
* 用==比较数组
    1. 相同维数且含有相同个数元素的数组才可以比较
    2. 每个元素都相同的才相等
* 逻辑运算符
    1. && || !
* 位运算符(与其他主要语言差异)
    1. & | ^ << >>
    2. &^ 按位置0
* 循环(与其他主要语言差异)
    1. Go语言仅支持循环关键字
    2. for j := 7; j <= 9; j++
* if条件(与其他主要语言差异)
    1. condition表达式结果必须为布尔值
    2. 支持变量赋值:
    3. if var declaration; condition {
        
    }
* switch(与其他主要语言差异)
    1. 条件表达式不限制为常量或者整数
    2. 单个case中,可以出现多个结果选项,使用逗号分隔;
    3. 与C语言规则相反,Go不需要break来明确退出一个case;
    4. 可以不设定switch之后的条件表达式,在此种情况下,整个switch结构与多个if...else...的逻辑等同























