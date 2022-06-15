package rpc

import (
	"net/http"
	"sync"
)

type Context struct {
	ID         string
	Endpoint   string
	Request    *http.Request
	Writer     http.ResponseWriter
	SocketMsgC *sync.Map
}
