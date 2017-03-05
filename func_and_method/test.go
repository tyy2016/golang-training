package main

import "fmt"

type books struct{
	id int
	name string
}

//函数
func show(a books){
	fmt.Print(a.id)
	fmt.Println(a.name)
}

//方法
func (b *books)show() {
	fmt.Print(b.id)
	fmt.Println(b.name)
}

func main(){
	books1 := books{11,"math"}
	books2 := books{12,"english"}
//	fmt.Println(books1)
//	fmt.Println(books2)
	show(books1)
	books2.show()
}
