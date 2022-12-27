package internal

import (
	"chat_socket/server/define"
	"chat_socket/server/helper"
	"chat_socket/server/internal/handler"
	"strings"
	"time"

	"encoding/json"
	"fmt"
	"net"
)

type Client struct {
	conn net.Conn
	hub  *Hub
	send chan string
	addr string
	uc   *define.Userclaim
}

func ServerWs(conn net.Conn, hub *Hub) {

	user := Client{
		conn: conn,
		hub:  hub,
		send: make(chan string),
		addr: conn.RemoteAddr().String(),
		uc:   new(define.Userclaim),
	}

	go user.readPump()

	go user.writePump()

	//user.hub.broadcast <- makeMsg(&user, string(jsons))

}

func makeMsg(user *Client, msg string) (buf string) {
	//t := time.Now()

	tU := time.Now().Unix()
	tFStr := time.Unix(tU, 0).Format("15:04:05")
	buf = fmt.Sprintf("[%s]{%s}：%s", tFStr, user.uc.Name, msg)
	fmt.Println(buf)
	return
}

func (user *Client) readPump() {
	defer func() {
		user.hub.unregister <- user
		user.conn.Close()
	}()
	buf := make([]byte, 1024)
	for {
		n, err := user.conn.Read(buf)
		if n == 0 {
			fmt.Printf("用户 {%s} 退出\n", user.uc.Name)
			return
		}
		if err != nil {
			fmt.Println("conn.Read err !=nil:", err, n)
			return
		}
		msg := string(buf[:n])
		switch {
		//登录验证
		case len(msg) > 5 && msg[:6] == "/login":
			//fmt.Println(msg)
			token, err := handler.UserLoginHandler(msg)
			if err != nil {
				fmt.Println("登录失败", err)
				return
			}
			jsons, err := json.Marshal(token)
			if err != nil {
				fmt.Println("json.Marshal(token)失败", err)
				return
			}
			user.uc, err = helper.AnalyzeToken(token.Token)
			if err != nil {
				fmt.Println(" helper.AnalyzeToken(token.Token)失败", err)
				return
			}
			//用户加入房间
			user.hub.register <- user
			user.hub.broadcast <- makeMsg(user, "上线")
			//fmt.Printf("[%s]{%s}:%s\n", user.addr, user.name, "上线")
			user.send <- makeMsg(user, string(jsons))

			//房间在线人数
		case msg == "/online":
			user.conn.Write([]byte("online user list\n"))
			for v, _ := range user.hub.clients {
				userInfo := fmt.Sprintf("[%s]{%s}\n", v.addr, v.uc.Name)
				user.conn.Write([]byte(userInfo))
			}

		case len(msg) > 15 && msg[:3] == "to#":
			//私信模式
			//格式 to#ip:port#用户消息
			content := strings.Split(msg, "#")
			msg = "来自 " + user.uc.Name + "的私信:  " + content[2]

			for c, _ := range user.hub.clients {
				if c.uc.Name == content[1] {
					c.send <- msg
				}
			}

		default:
			fmt.Println(msg)
			user.hub.broadcast <- makeMsg(user, msg)
		}

	}
}

func (user *Client) writePump() {

	for {
		select {
		//	用户重置计时器
		case <-time.After(time.Second * 150):
			delete(user.hub.clients, user) //将用户从在线列表移除
			//messageList <- makeMsg(user, user.name+" time out")
			fmt.Println("用户重置计时器")

			user.hub.broadcast <- makeMsg(user, user.uc.Name+" 退出房间")
			user.conn.Close()
			return

		case msg := <-user.send:
			user.conn.Write([]byte(msg + "\n"))
		}
	}

}
