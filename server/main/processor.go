package main
import (
	"fmt"
	"net"
	"go_code/go_code/chatroom/common/message"
	"go_code/go_code/chatroom/server/utils"
	"go_code/go_code/chatroom/server/process"
	"io"
)

type Processor struct{
	Conn  net.Conn
}

func (this *Processor) serverProcessMes(mes *message.Message)(err error){

	fmt.Println("mes=",mes)

	switch mes.Type{
		case message.LoginMesType :
			up := &process2.UserProcess{
				Conn : this.Conn,
			}
			err = up.ServerProcessLogin(mes)
		case message.RegisterMesType :
			up := &process2.UserProcess{
				Conn : this.Conn,
			}
			err = up.ServerProcessRegister(mes)
		case message.SmsMesType:
			smsProcess := &process2.SmsProcess{}
			smsProcess.SendGroupMes(mes)
		default :
			fmt.Println("消息类型不存在，无法处理。")
	}
	return 
}
func (this *Processor) process2()(err error){

	for{
		tf := &utils.Transfer{
			Conn : this.Conn,
		}
		mes,err := tf.ReadPkg()
		if err != nil{
			if err == io.EOF{
				fmt.Println("客户端退出,服务器端也正常关闭..")
				return err
			}else{
			fmt.Println("readpkg err=",err)
			return err 
			}
		}
		//fmt.Println("mes=",mes)
		err = this.serverProcessMes(&mes)
		if err != nil{
			return err
		}
	}
}