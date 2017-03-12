package main

import "fmt"

type User struct {
	Name string
	Email string
}

type Notifier interface {
	Notify() error
	Sending() error
}

func (u User) Notify() error {
	fmt.Printf("Notifying to %s @ %s \n",u.Name,u.Email)
	return nil
}

func (u User) Sending() error{
	fmt.Printf("Sending to %s @ %s \n",u.Name,u.Email)
	return nil
}

func Notification(notify Notifier) error{
	return notify.Notify()
}

func Send(notify Notifier) error{
	return notify.Sending()
}

func main(){
	user := User{"tyy","tyy@123.com"}
/*part one---------------------------------------
//将方法Notify()和Sending()都声明成指针类型,传值类型参数进去
	Notification(user)
	Send(user)
# command-line-arguments
./main.go:35: cannot use user (type User) as type Notifier in argument to Notification:
        User does not implement Notifier (Notify method has pointer receiver)
./main.go:36: cannot use user (type User) as type Notifier in argument to Send:
        User does not implement Notifier (Notify method has pointer receiver)
*/

//part two---------------------------------------
//将方法Notify()和Sending()都声明成指针类型,传指针类型参数进去
	Notification(&user)
	Send(&user)
//under this situation run successfully
  


//part three---------------------------------------
//将方法Notify()和Sending()都声明成值类型,传值类型参数进去
	Notification(user)
	Send(user)
//under this situation run successfully


//part four----------------------------------------
//将方法Notify()和Sending()都声明成值类型,传指针类型参数进去
	Notification(&user)
	Send(&user)
//under this situation run successfully
}
