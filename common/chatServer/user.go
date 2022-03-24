package chatServer

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"strings"
)

type WsUser struct {
	Name   string
	Addr   string
	Ch     chan string
	Conn   *websocket.Conn
	server *ServerStruct
}

// NewUser 创建一个用户的API
func NewUser(conn *websocket.Conn, server *ServerStruct) *WsUser {
	user := &WsUser{
		Name:   conn.RemoteAddr().String(),
		Addr:   conn.RemoteAddr().String(),
		Conn:   conn,
		Ch:     make(chan string),
		server: server,
	}
	//启动监听当前user channel消息的go程
	go user.ListenMessage()
	return user
}

// ListenMessage 监听当前用户channel的方法，一旦有消息，就直接发送给对端客户端
func (u *WsUser) ListenMessage() {
	for {
		msg := <-u.Ch
		err := u.Conn.WriteMessage(len(msg), []byte(u.Name+msg+"\n"))
		if err != nil {
			log.Printf("err:%v", err)
		}
	}
}

// Online 用户上线功能
func (u *WsUser) Online() {
	//用户上线后加入到Map当中
	u.server.mapLock.Lock()
	u.server.OnlineMap[u.Conn.RemoteAddr().String()] = u
	u.server.mapLock.Unlock()
	log.Printf("%v", Server.OnlineMap)
	//广播当前用户已上线！
	u.server.BroadCast(u, "已上线！")
}

// Offline 用户的下线业务
func (u *WsUser) Offline() {
	//用户下线后删除Map中数据
	u.server.mapLock.Lock()
	delete(u.server.OnlineMap, u.Name)
	u.server.mapLock.Unlock()
	//广播当前用户已上线！
	u.server.BroadCast(u, "已下线！")
}

// SendMsg 发送信息
func (u *WsUser) SendMsg(msg string) {
	msg = msg + "\n"
	fmt.Printf("在发送消息！")
	err := u.Conn.WriteMessage(1, []byte(msg))
	if err != nil {
		return
	}
}
func (u *WsUser) name() {

}

// DoMessage 用户处理消息的业务
func (u *WsUser) DoMessage(msg string) {
	switch {
	case msg == "who":
		u.server.mapLock.Lock()
		for _, v := range u.server.OnlineMap {
			u.SendMsg(v.Conn.RemoteAddr().String())
		}
		u.server.mapLock.Unlock()
		//更新name
	case len(msg) > 7 && msg[:7] == "rename|":
		//消息格式为rename|*****
		newName := strings.Split(msg, "|")[1]
		u.Name = newName
		u.SendMsg("更新姓名为：" + newName)
	case len(msg) > 4 && msg[:3] == "to|":
		//消息格式to|******|消息内容
		//获取对方用户名
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			u.SendMsg("此用户名无效！，请使用\"to|张三|你好！\"\n")
			return
		}
		//根据用户名，得到对方User对象
		remoteUser, ok := u.server.OnlineMap[remoteName]
		if !ok {
			u.SendMsg("该用户不存在！")
			return
		}
		//获取消息内容，通过对方User对象将消息内容发送过去
		content := strings.Split(msg, "|")[2]
		if content == "" {
			u.SendMsg("不能发送空的消息！\n")
			return
		}
		remoteUser.SendMsg(u.Name + "给您发送：" + content)
	default:
		u.server.BroadCast(u, msg)
	}

}
