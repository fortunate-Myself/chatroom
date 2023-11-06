package process

import (
	"fmt"
	"net"
	"go_code/go_code/chatroom/common/message"
	"go_code/go_code/chatroom/client/utils"
	"encoding/binary"
	"encoding/json"
	"os"

)

type UserProcess struct{

}

func (this *UserProcess) Register(userId int,userPwd,userName string)(err error){
	conn,err := net.Dial("tcp","127.0.0.1:8889")
	if err != nil{
		fmt.Println("net.dial err=",err)
		return 
	}
	defer conn.Close()

	var mes message.Message
	mes.Type = message.RegisterMesType

	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	data,err := json.Marshal(registerMes)
	if err != nil{
		fmt.Println("json.Marshal  err=",err)
		return 
	}
	mes.Data = string(data)
	data,err = json.Marshal(mes)
	if err != nil{
		fmt.Println("json.Marshal err=",err)
		return 
	}

	tf := &utils.Transfer{
		Conn : conn,
	}
	err = tf.WritePkg(data)
	if err != nil{
		fmt.Println("注册发送包出错! err=",err)
	}
	mes,err = tf.ReadPkg()
	if err != nil{
		fmt.Println("readPkg(conn) err=",err)
		return
	}

	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data),&registerResMes)
	
	if registerResMes.Code == 200{	
		fmt.Println("注册成功,你重新登陆")
		os.Exit(0)
	}else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return 
}



func (this *UserProcess) Login(userId int,userPwd string)(err error){
	
	conn,err := net.Dial("tcp","127.0.0.1:8889")
	if err != nil{
		fmt.Println("net.dial err=",err)
		return 
	}
	defer conn.Close()

	
	var mes message.Message
	mes.Type = message.LoginMesType

	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	data,err := json.Marshal(loginMes)
	if err != nil{
		fmt.Println("json.Marshal  err=",err)
		return 
	}

	mes.Data = string(data)

	data,err = json.Marshal(mes)
	if err != nil{
		fmt.Println("json.Marshal err=",err)
		return 
	}

	var pkgLen uint32
	pkgLen = uint32(len(data))
	
	var buf [4]byte 
	binary.BigEndian.PutUint32(buf[0:4],pkgLen)

	n,err := conn.Write(buf[:4])
	if n != 4 || err != nil{
		fmt.Println("conn.Write(buf) fail",err)
		return 
	}
	fmt.Printf("客户端的消息长度=%d,内容=%s\n",len(data),string(data))
	_,err = conn.Write(data)
	if  err != nil{
		fmt.Println("conn.Write(data) fail",err)
		return 
	}

	tf := &utils.Transfer{
		Conn : conn,
	}
	mes,err = tf.ReadPkg()
	if err != nil{
		fmt.Println("readPkg(conn) err=",err)
		return
	}

	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data),&loginResMes)
	
	if loginResMes.Code == 200{	

		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline 

		//fmt.Println("登陆成功")
		fmt.Println("当前在线用户列表如下:")
		for _,v := range loginResMes.UsersId{
			
			if v == userId {
				continue
			}

			fmt.Println("用户id:\t",v)

			user := &message.User{
				UserId : v,
				UserStatus : message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Print("\n\n")
		go serverProcessMes(conn)
		for{
			ShowMenu()
		}
	}else {
		fmt.Println(loginResMes.Error)
		
	}

	return 
}