package process

import (
	"fmt"
	"encoding/json"
	"go_code/go_code/chatroom/common/message"
	"go_code/go_code/chatroom/client/utils"
)

type SmsProcess struct{

}

func (this *SmsProcess) SendGroupMes(content string) (err error) {

	var mes message.Message
	mes.Type = message.SmsMesType

	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	data,err := json.Marshal(smsMes)
	if err != nil{
		fmt.Println("SendGroupMes json.Marshal fail=",err.Error())
		return
	}

	mes.Data = string(data)

	data,err = json.Marshal(mes)
	if err != nil{
		fmt.Println("SendGroupMes json.Marshal fail=",err.Error())
		return
	}

	tf := &utils.Transfer{
		Conn : CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil{
		fmt.Println("SendGroupMes err=",err.Error())
		return
	}
	return
}