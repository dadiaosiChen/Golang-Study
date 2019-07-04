package main

import(
	"flag"
	"fmt"
)

//第一个参数，string类型，程序内参数调用名为argstr1，命令源码参数输入为arg1
var argstr1 string
//初始化参数argstr1，第一个为参数程序内地址，第二个为参数名，第三个为参数默认值，第四个为参数解释，在--help时会使用到
func init() {
	flag.StringVar(&argstr1,"arg1","default1","test string argument1")
}

//同第一个参数
var argint1 int

func init() {
	flag.IntVar(&argint1,"arg2",10086,"test int argument2")
}

func main() {
	//解析参数
	flag.Parse()
	fmt.Printf("The first arg is %s, the second arg is %d.",argstr1,argint1)
}