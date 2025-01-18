package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hertz-contrib/sessions"
	common "github.com/lgxyc/gomall/app/frontend/hertz_gen/frontend/common"
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
	session.Delete("user_id")
	if err = session.Save(); err != nil {
		klog.Fatal(err)
	}
	return nil, err
}
