package socket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"net/http"
	"sync"
	"time"
)

var up = &websocket.Upgrader{
	//定义读写缓冲区大小
	WriteBufferSize: 1024,
	ReadBufferSize:  1024,
	//校验请求
	CheckOrigin: func(r *http.Request) bool {
		//如果不是get请求，返回错误
		if r.Method != "GET" {
			fmt.Println("请求方式错误")
			return false
		}
		//还可以根据其他需求定制校验规则
		return true
	},
}

var socketDb *gorm.DB

func Setup(db *gorm.DB) {
	socketDb = db
}

type TUser struct {
	ID int64 `json:"i"`
}

func SocketHandleFunc(w http.ResponseWriter, r *http.Request, m *sync.Map) {
	//通过升级后的升级器得到链接
	conn, err := up.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("获取连接失败:", err)
		return
	}

	var userID string

	for {
		_, uid, _ := conn.ReadMessage()
		if uid != nil {
			userID = string(uid)
			m.Store(userID, make(chan interface{}, 20))
			break
		}
	}

	for {
		msgCI, _ := m.LoadOrStore(userID, make(chan interface{}, 20))
		msgC := msgCI.(chan interface{})
		select {
		case mI := <-msgC:
			conn.WriteJSON(mI)
			time.Sleep(1000)
		default:
			time.Sleep(1000)
		}
	}
}
