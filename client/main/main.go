package main
import (
	"fmt"
	"go_code/go_code/chatroom/client/process"
)

var userId int
var userPwd string
var userName string

func main(){
	 var key int
	// var loop = true

	for true{
		fmt.Println("-----------------欢迎登陆多人聊天系统-----------------")
		fmt.Println("\t\t\t 1 登录聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("请选择(1 - 3):")

		fmt.Scanf("%d\n",&key)

		switch key {
			case 1:
				fmt.Println("登录聊天室")
				fmt.Println("请输入用户id:")
				fmt.Scanf("%d\n",&userId)
				fmt.Println("请输入用户密码:")
				fmt.Scanf("%s\n",&userPwd)
				
				up := &process.UserProcess{}
				up.Login(userId,userPwd)
			case 2:
				fmt.Println("注册用户")
				fmt.Println("请输入用户Id:")
				fmt.Scanln(&userId)
				fmt.Println("请输入用户密码:")
				fmt.Scanln(&userPwd)
				fmt.Println("请输入用户名字:")
				fmt.Scanln(&userName)
				up := &process.UserProcess{}
				up.Register(userId,userPwd,userName)
				
			case 3:
				fmt.Println("退出系统")
				//loop = false
			default:
				fmt.Println("输入有误，请重新输入！")
			
		}
	}
	// if key == 1{
	// 	fmt.Println("请输入用户id:")
	// 	fmt.Scanf("%d\n",&UserId)
	// 	fmt.Println("请输入用户密码:")
	// 	fmt.Scanf("%s\n",&UserPwd)

	// 	//login(UserId,UserPwd)

	// 	// if err != nil{
	// 	// 	fmt.Println("登陆失败")
	// 	// }else{
	// 	// 	fmt.Println("登陆成功")
	// 	// }
	// }else if key == 2{
	// 	fmt.Println("进行用户注册")
	// }
}