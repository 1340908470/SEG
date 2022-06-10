package services

import (
	"fmt"
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
	res.Message = fmt.Sprintf("hello, %v", req.Name)
	return nil
}
