package process2
import (
	"fmt"
	"net"
	"encoding/json"
	"go_code/go_code/chatroom/common/message"
	"go_code/go_code/chatroom/server/utils"
)

type SmsProcess struct{

}

func (this *SmsProcess) SendGroupMes(mes *message.Message){

	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data),&smsMes)
	if err != nil{
		fmt.Println("json.Unmarshal err=",err)
		return
	}

	data,err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=",err)
		return
	}

	for id,up := range userMgr.onlineUsers{

		if id == smsMes.UserId{
			continue
		}

		this.SendMesToEachOnlineUser(data,up.Conn)

	}

}

func (this *SmsProcess) SendMesToEachOnlineUser(data []byte,conn net.Conn){

	tf := &utils.Transfer{
		Conn : conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败 err=",err)
	}
}