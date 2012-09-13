package controllers

import (
	"code.google.com/p/go.net/websocket"
	"helper"
	"html/template"
	"net/http"
	"strings"
	"time"
)

func init() {
	App.RegisterController("chatroom/", PageChatroom{})
}

const (
	TEXT_MTYPE   = "text_mtype"
	STATUS_MTYPE = "status_mtype"
	TIME_FORMAT  = "01-02 15:04:05"
)

var runningActiveRoom *ActiveRoom = &ActiveRoom{}

type PageChatroom struct {
	Application
}

type RenderData struct {
	Email         string
	WebSocketHost string
}

func (p *PageChatroom) Index(w http.ResponseWriter, r *http.Request) {
}

func (p *PageChatroom) Join(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if mail, ok := p.POST["email"]; ok {
			render := RenderData{Email: mail, WebSocketHost: r.Host}
			p.Body = render
		} else {
			p.Template = "chatroom/error.html"
		}
	} else {
		p.Body = "错误的请求"
		p.Template = "chatroom/error.html"
	}
	return
}

func BuildConnection(ws *websocket.Conn) {
	email := ws.Request().URL.Query().Get("email")

	if email == "" {
		return
	}
	email = template.HTMLEscapeString(email)
	onlineUser := &OnlineUser{
		InRoom:     runningActiveRoom,
		Connection: ws,
		Send:       make(chan Message, 256),
		UserInfo: &User{
			Email:    email,
			Name:     strings.Split(email, "@")[0],
			Gravatar: helper.UrlSize(email, 20),
		},
	}
	runningActiveRoom.OnlineUsers[email] = onlineUser

	m := Message{
		MType: STATUS_MTYPE,
		UserStatus: UserStatus{
			Users: runningActiveRoom.GetOnlineUsers(),
		},
	}
	runningActiveRoom.Broadcast <- m

	go onlineUser.PushToClient()
	onlineUser.PullFromClient()

	onlineUser.killUserResource()
}

type ActiveRoom struct {
	OnlineUsers map[string]*OnlineUser
	Broadcast   chan Message
	CloseSign   chan bool
}

type OnlineUser struct {
	InRoom     *ActiveRoom
	Connection *websocket.Conn
	UserInfo   *User
	Send       chan Message
}

type User struct {
	Name     string
	Email    string
	Gravatar string
}

type Message struct {
	MType       string
	TextMessage TextMessage
	UserStatus  UserStatus
}

type TextMessage struct {
	Content  string
	UserInfo *User
	Time     string
}

type UserStatus struct {
	Users []*User
}

func InitChatRoom() {
	runningActiveRoom = &ActiveRoom{
		OnlineUsers: make(map[string]*OnlineUser),
		Broadcast:   make(chan Message),
		CloseSign:   make(chan bool),
	}
	go runningActiveRoom.run()
}

// Core function of room
func (this *ActiveRoom) run() {
	for {
		select {
		case b := <-this.Broadcast:
			for _, online := range this.OnlineUsers {
				online.Send <- b
			}
		case c := <-this.CloseSign:
			if c == true {
				close(this.Broadcast)
				close(this.CloseSign)
				return
			}
		}
	}
}

func (this *OnlineUser) PullFromClient() {
	for {
		var content string
		err := websocket.Message.Receive(this.Connection, &content)
		// If user closes or refreshes the browser, a err will occur
		if err != nil {
			return
		}

		m := Message{
			MType: TEXT_MTYPE,
			TextMessage: TextMessage{
				UserInfo: this.UserInfo,
				Time:     humanCreatedAt(),
				Content:  template.HTMLEscapeString(content),
			},
		}
		this.InRoom.Broadcast <- m
	}
}

func (this *OnlineUser) PushToClient() {
	for b := range this.Send {
		err := websocket.JSON.Send(this.Connection, b)
		if err != nil {
			break
		}
	}
}

func (this *OnlineUser) killUserResource() {
	this.Connection.Close()
	delete(this.InRoom.OnlineUsers, this.UserInfo.Email)
	close(this.Send)

	m := Message{
		MType: STATUS_MTYPE,
		UserStatus: UserStatus{
			Users: runningActiveRoom.GetOnlineUsers(),
		},
	}
	runningActiveRoom.Broadcast <- m
}

func (this *ActiveRoom) GetOnlineUsers() (users []*User) {
	for _, online := range this.OnlineUsers {
		users = append(users, online.UserInfo)
	}
	return
}

func humanCreatedAt() string {
	return time.Now().Format(TIME_FORMAT)
}
