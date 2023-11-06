package process2

import (
	"fmt"
)

var (
	userMgr *UserMgr
)

type UserMgr struct{
	onlineUsers map[int]*UserProcess
}

func init(){
	userMgr = &UserMgr{
		onlineUsers : make(map[int]*UserProcess,1024),
	}
}

func (this *UserMgr) AddOnlineUser(up *UserProcess){

	this.onlineUsers[up.UserId] = up 
}

func (this *UserMgr) DeOnlineUser(userId int){

	delete(this.onlineUsers,userId) 
}

func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess{
	return this.onlineUsers
}

func (this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess,err error){

	up,ok := this.onlineUsers[userId]
	if !ok{
		err = fmt.Errorf("用户%d不存在",userId)
		return
	}
	return
}