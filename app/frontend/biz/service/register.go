package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hertz-contrib/sessions"
	auth "github.com/lgxyc/gomall/app/frontend/hertz_gen/frontend/auth"
	common "github.com/lgxyc/gomall/app/frontend/hertz_gen/frontend/common"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *auth.RegisterReq) (resp *common.Empty, err error) {
	sessions := sessions.Default(h.RequestContext)
	sessions.Set("user_id", 1)
	if err = sessions.Save(); err != nil {
		klog.Fatal(err)
	}
	return nil, err
}
