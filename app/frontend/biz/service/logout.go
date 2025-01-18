package service

import (
	"context"
	"log"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	common "github.com/lgxyc/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/lgxyc/gomall/app/frontend/middleware"
)

type LogoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLogoutService(Context context.Context, RequestContext *app.RequestContext) *LogoutService {
	return &LogoutService{RequestContext: RequestContext, Context: Context}
}

func (h *LogoutService) Run(req *common.Empty) (resp *common.Empty, err error) {
	session := sessions.Default(h.RequestContext)
	session.Delete(middleware.SessionUserId)
	if err = session.Save(); err != nil {
		log.Fatalln(err)
	}
	return nil, err
}
