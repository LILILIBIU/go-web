package chatServer

var Server = &ServerStruct{
	OnlineMap: make(map[string]*WsUser),
	Message:   make(chan string),
}

func InitChatServer() {
	//Server =NewServer()
	Server.Start()
}
