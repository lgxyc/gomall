package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	common "github.com/lgxyc/gomall/app/frontend/hertz_gen/frontend/common"
	product "github.com/lgxyc/gomall/app/frontend/hertz_gen/frontend/product"
)

type SearchProductListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchProductListService(Context context.Context, RequestContext *app.RequestContext) *SearchProductListService {
	return &SearchProductListService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProductListService) Run(req *product.SearchProductListReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
