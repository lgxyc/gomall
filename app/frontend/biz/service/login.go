package service

import (
	"context"
	"log"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	auth "github.com/lgxyc/gomall/app/frontend/hertz_gen/frontend/auth"
	"github.com/lgxyc/gomall/app/frontend/middleware"
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
	session.Set(middleware.SessionUserId, 1)
	if err = session.Save(); err != nil {
		log.Fatal(err)
	}
	// 返回上一级页面
	redirect = "/"
	if req.Next != "" {
		redirect = req.Next
	}
	return redirect, err
}
