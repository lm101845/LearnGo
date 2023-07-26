/**
 * @Author: liming
 * @Date: 2023/7/26 14:34
 */

// 知识点：常量
package main

const ee = 2.7182

//常量定义之后就不能修改了

// 多个常量也可以一起声明：
const (
	statusOK = 200
	notFound = 404
)

// const同时声明多个常量时，如果省略了值则表示和上面一行的值相同
const (
	n1 = 100
	n2
	n3
)

// iota是go语言的常量计数器，只能在常量的表达式中使用。
const (
	a1 = iota
	a2
	_
	//使用_可以跳过某些值
	a4
)

const (
	b1 = iota
	b2 = iota
	_
	b4 = iota
)

const (
	c1 = iota
	c2 = 100
	c3 = iota
	c4 = iota
)

// 多个iota定义在一行
// iota在const关键字出现时将被重置为0。
// const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。 使用iota能简化定义，在定义枚举时很有用。
const (
	a, b = iota + 1, iota + 2 //1,2
	c, d                      //2,3
	e, f                      //3,4
)

// 定义数量级 （这里的<<表示左移操作，1<<10表示将1的二进制表示向左移10位，
// 也就是由1变成了10000000000，也就是十进制的1024。
// 同理2<<2表示将2的二进制表示向左移2位，也就是由10变成了1000，也就是十进制的8。）
const (
	_  = iota
	KB = 1 << (10 * iota) //2的10次方
	MB = 1 << (10 * iota) //2的20次方
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	const pi = 3.1415
}
