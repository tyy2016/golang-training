package main

import "fmt"

func testfunc(x int)int{
	return x+1
}

func pointfunc(y *int)int{
	return *y+1
}

type info struct{
	Name string
	Age  int
}

func (c info)testtestfunc(){
	fmt.Println(c.Name)
	fmt.Println(c.Age)
}

func (d *info)pointpointfunc(){
	fmt.Println(d.Name)
	fmt.Println(d.Age)
}

func main(){
//	a := 2
//	fmt.Println("testfunc:",testfunc(&a))
//	函数参数为值类型,则不能将指针作为参数传递
//	./main.go:15: cannot use &a (type *int) as type int in argument to testfunc

//	fmt.Println("testfunc:",pointfunc(a))
//	函数参数为指针类型,则不能将值类型参数传递
//	./main.go:18: cannot use a (type int) as type *int in argument to pointfunc

//	information是info类的一个变量,将information变量的值给testtestfunc方法
	information := info{"mush",24}
	information.testtestfunc()
//	接收者为指针类型的方法,information依然可以调用指针类型的方法
	information.pointpointfunc()

//	secinformation是info类的一个指针变量,可以支持pointpointfunc()方法,
	secinformation := &info{"room",25}
	secinformation.pointpointfunc()
//	接收者是值类型的方法,指针变量依然可以调用值类型的方法
	secinformation.testtestfunc()
}
