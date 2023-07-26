/**
 * @Author: liming
 * @Date: 2023/7/26 14:17
 */

// 知识点：变量
package main

import "fmt"

// 声明变量
//var name string
//var age int
//var isOk bool

// 批量声明
var (
	name string
	age  int
	isOk bool
)

func main() {
	name = "李想"
	age = 18
	isOk = true
	//go语言中变量声明了必须要使用，不使用则编译不过去
	fmt.Println(name) //打印完指定内容后，会在后面加一个换行符
	fmt.Print(name)   //没有换行符
	fmt.Printf("name:%s", name)

	//go语言推荐使用驼峰命名

	// 声明变量同时赋值
	var s1 string = "李明"
	fmt.Println(s1)

	// 类型推导(根据值判断该变量是什么类型)
	var s2 = 20
	fmt.Println(s2)

	//简短变量声明(只能在函数内部用，不能在外部使用)
	s3 := "嘿嘿嘿"
	fmt.Println(s3)

	//s3 := "哈哈哈"
	//同一个作用域中不能重复声明同名的变量

	//匿名变量是一个特殊的变量:_(后面学了函数再说)
}
