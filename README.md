# Go web Demo

使用golang mysql搭建的多人聊天系统

本项目提供二进制代码和源码可供参考

## 如何拉取代码

```shell
git clone https://github.com/LILILIBIU/go-web.git
```



## 如何启动

编译文件

```shell
cd go-web
go run .
```

启动项目默认使用端口号:8888

如有需要可以在application.toml配置文件中更改server.port以更改项目端口号，其他配置信息都可在配置文件中更改



## 如何进入

在启动项目后可以访问

```
http://127.0.0.1:8888/ws/WebSocket
```

以查看基于websocket的线上聊天室Demo



## 更多

在这个项目中可以多人进入聊天室

### 聊天室指令

#### who

who 命令可以查看在线人员

#### rename

rename|****** 可以在线更改自己的昵称

#### to|name

to|zhangsan|******可以指定人员进行聊天，聊天内容仅双方可见，原理是命令可以直接联通两人的私人信息channel进行通信。
