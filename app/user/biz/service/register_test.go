package service

import (
	"context"
	"testing"

	"github.com/joho/godotenv"
	"github.com/lgxyc/gomall/app/user/biz/dal/mysql"
	user "github.com/lgxyc/gomall/rpc_gen/kitex_gen/user"
)

func TestRegister_Run(t *testing.T) {
	// 加载环境变量
	godotenv.Load("../../.env")
	// 初始化mysql
	mysql.Init()
	ctx := context.Background()
	s := NewRegisterService(ctx)
	// init req and assert value

	req := &user.RegisterReq{
		Email:           "demo@dsdslk",
		Password:        "FKJSLKD",
		PasswordConfirm: "xx",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
}
