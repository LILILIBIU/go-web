package chatServer

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"runtime"
	"sync"
	"time"
)

// ServerStruct 服务结构体
type ServerStruct struct {
	OnlineMap map[string]*WsUser
	mapLock   sync.RWMutex
	//TODO 在线用户列表预计在Redis实现,先在这用Map实现，可以抽离列表
	Message chan string
}

// NewServer 新建Server服务
func NewServer() *ServerStruct {
	server := &ServerStruct{
		OnlineMap: make(map[string]*WsUser),
		Message:   make(chan string),
	}
	return server
}

// ListenMessage 监听广播channel的go程，一旦有消息马上转发给所有User
func (s *ServerStruct) ListenMessage() {
	go func() {
		for {
			log.Printf("ListenMessage 正常启动！")
			msg := <-s.Message
			//将Msg发送给所有在线用户
			s.mapLock.RLock()
			for _, v := range s.OnlineMap {
				v.Ch <- msg
			}
			s.mapLock.RUnlock()
			log.Println("ListenMessage 发送完毕！")
		}
	}()
}

// BroadCast 广播方法
func (s *ServerStruct) BroadCast(user *WsUser, msg string) {
	sendMsg := "[" + user.Name + "]" + user.Name + ":" + msg
	//log.Printf("在BroadCast里面！\n")
	//log.Println(sendMsg)
	//把消息发送给广播channel
	s.Message <- sendMsg
	//log.Println(sendMsg)

}

// Handler 用户进入的主处理函数
func (s *ServerStruct) Handler(conn *websocket.Conn, c *gin.Context) {
	user := NewUser(conn, s)
	log.Printf("连接成功！")
	user.Online()
	//监听用户是否活跃channel
	isLive := make(chan struct{})
	isClose := make(chan struct{})
	go func() {
		for {
			select {
			case <-isClose:
				runtime.Goexit()
			default:
				msg := <-user.Ch
				err := user.Conn.WriteMessage(1, []byte(msg))
				if err != nil {
					log.Printf("user.Conn.WriteMessage err:%v\n", err)
				}
			}
		}
	}()
	//接收客户端发来的消息
	go func() {
		for {
			select {
			case <-isClose:
				runtime.Goexit()
			default:
				//读取ws中的数据
				_, buf, err := conn.ReadMessage()
				if err != nil {
					log.Printf("读取ws中的数据,err:%v", err)
					user.Offline()
					runtime.Goexit()
					break
				}
				isLive <- struct{}{}
				//log.Printf("%T", buf)
				msg := string(buf)
				log.Println("读数据之后！")
				//log.Println(msg)
				user.DoMessage(msg)
			}
		}
	}()
	//当前阻塞
LOOP:
	for {
		select {
		case <-isLive:
			//重置计时器
		case <-time.After(time.Second * 30):
			//已经超时
			user.SendMsg("窗口超时！")
			isClose <- struct{}{}
			isClose <- struct{}{}
			user.Offline()
			err := conn.Close()
			if err != nil {
				fmt.Printf("conn.Close err:%v\n", err)
				break LOOP
			}
			//退出当前Handler
			runtime.GC()
			break LOOP
		}
	}
}

// Start 启动服务
func (s *ServerStruct) Start() {
	//启动监听MSG的go程
	s.ListenMessage()
	//do handler
	//close 监听 socket
}
