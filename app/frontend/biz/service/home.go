package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	common "github.com/lgxyc/gomall/app/frontend/hertz_gen/frontend/common"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (map[string]any, error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	resp := make(map[string]any)

	items := []map[string]any{
		{"Name": "T-shirt", "Price": 100, "Picture": "/static/image/t-shirt-1.jpeg"},
		{"Name": "T-shirt-2", "Price": 200, "Picture": "/static/image/t-shirt-2.jpeg"},
		{"Name": "T-shirt", "Price": 300, "Picture": "/static/image/t-shirt-1.jpeg"},
		{"Name": "T-shirt", "Price": 100, "Picture": "/static/image/t-shirt-1.jpeg"},
		{"Name": "T-shirt", "Price": 100, "Picture": "/static/image/t-shirt-1.jpeg"},
		{"Name": "T-shirt", "Price": 100, "Picture": "/static/image/t-shirt-1.jpeg"},
	}
	resp["Title"] = "Hot Sales"
	resp["Items"] = items
	return resp, nil
}
