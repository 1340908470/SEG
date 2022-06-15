package services

import (
	"fmt"
	"gorm.io/gorm"
	"seg/pkg/auth"
	"seg/pkg/auth/models"
	"seg/pkg/base/rpc"
)

type TestService struct {
	db *gorm.DB
}

func NewTestService(db *gorm.DB) *TestService {
	return &TestService{db: db}
}

type TestReq struct {
	Name string `json:"name"`
}

type TestRes struct {
	Message string `json:"message"`
}

func (s *TestService) Test(ctx *rpc.Context, req *TestReq, res *TestRes) (err error) {
	//1. 获取User信息
	var user *models.User

	//获取创建者信息
	var tokenUser auth.TokenUser
	tokenUser, err = auth.ExtractTokenUser(ctx)
	if err != nil {
		return err
	}

	tx := s.db.Begin()
	user = models.FindUserById(tx, tokenUser.ID)

	print(user)

	res.Message = fmt.Sprintf("hello, %v", req.Name)

	go func() {
		msgCI, _ := ctx.SocketMsgC.LoadOrStore(user.ID.ToString(), make(chan interface{}, 20))
		msgC := msgCI.(chan interface{})
		msgC <- res
	}()
	return nil
}
