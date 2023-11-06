package main

import (
	"fmt"
	"net"
	"time"
	"go_code/go_code/chatroom/server/model"
)

func process(conn net.Conn){
	defer conn.Close()
	
	processor := &Processor{
		Conn : conn,
	}
	err := processor.process2()
	if err != nil{
		fmt.Println("客户端和服务器通讯协程错误=",err)
		return
	}
	
}
func initUserDao(){
	model.MyUserDao = model.NewUserDao(pool)
}

func main(){
	initPool("localhost:6379",16,0,300 * time.Second)
	initUserDao()

	fmt.Println("服务器[新的结构]在8889端口监听...")
	listen,err := net.Listen("tcp","127.0.0.1:8889")
	defer listen.Close()
	if err != nil{
		fmt.Println("net.listen err=",err)
		return
	}
	for{
		fmt.Println("等待客户端链接服务器...")
		conn,err := listen.Accept()
		if err != nil{
			fmt.Println("listen.accept err=",err)
		}

		go process(conn)
	}
}