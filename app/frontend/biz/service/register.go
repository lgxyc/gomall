package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	auth "github.com/lgxyc/gomall/app/frontend/hertz_gen/frontend/auth"
	common "github.com/lgxyc/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/lgxyc/gomall/app/frontend/infra/rpc"
	"github.com/lgxyc/gomall/rpc_gen/kitex_gen/user"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *auth.RegisterReq) (resp *common.Empty, err error) {
	registerResp, err := rpc.UserClient.Register(h.Context,
		&user.RegisterReq{
			Email:           req.Email,
			Password:        req.Password,
			PasswordConfirm: req.PasswordConfirm,
		})
	if err != nil {
		return nil, err
	}
	sessions := sessions.Default(h.RequestContext)
	sessions.Set("user_id", registerResp.UserId)
	if err = sessions.Save(); err != nil {
		return nil, err
	}
	return nil, err
}
