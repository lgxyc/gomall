package service

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hertz-contrib/sessions"
	auth "github.com/lgxyc/gomall/app/frontend/hertz_gen/frontend/auth"
	"github.com/lgxyc/gomall/app/frontend/infra/rpc"
	"github.com/lgxyc/gomall/rpc_gen/kitex_gen/user"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginReq) (redirect string, err error) {
	loginResp, err := rpc.UserClient.Login(h.Context,
		&user.LoginReq{
			Email:    req.Email,
			Password: req.Password,
		})
	if err != nil {
		return "/", err
	}
	fmt.Println("func loginService userId:", loginResp.UserId)
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", loginResp.UserId)
	if err = session.Save(); err != nil {
		klog.Error(err)
	}
	// 返回上一级页面
	redirect = "/"
	if req.Next != "" {
		redirect = req.Next
	}
	return redirect, err
}
