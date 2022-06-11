package services

import (
	"fmt"
	"github.com/lantu-dev/puki/pkg/auth"
	"github.com/lantu-dev/puki/pkg/auth/models"
	"github.com/lantu-dev/puki/pkg/base/rpc"
	"gorm.io/gorm"
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
	return nil
}
