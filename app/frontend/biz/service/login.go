package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hertz-contrib/sessions"
	auth "github.com/lgxyc/gomall/app/frontend/hertz_gen/frontend/auth"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginReq) (redirect string, err error) {
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", 1)
	if err = session.Save(); err != nil {
		klog.Fatal(err)
	}
	// 返回上一级页面
	redirect = "/"
	if req.Next != "" {
		redirect = req.Next
	}
	return redirect, err
}
