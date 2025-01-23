package utils

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/lgxyc/gomall/app/frontend/infra/rpc"
	frontendutils "github.com/lgxyc/gomall/app/frontend/utils"
	rpccart "github.com/lgxyc/gomall/rpc_gen/kitex_gen/cart"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	c.String(code, err.Error())
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, data)
}

func WarpResponse(ctx context.Context, c *app.RequestContext, content map[string]any) map[string]any {
	userId := frontendutils.GetUserIdFromCtx(ctx)
	content["user_id"] = userId

	if userId > 0 {
		getCartResp, err := rpc.CartClient.GetCart(ctx, &rpccart.GetCartReq{UserId: frontendutils.GetUserIdFromCtx(ctx)})
		if err == nil && getCartResp != nil {
			content["cart_num"] = len(getCartResp.Cart.Items)
		}
	}
	return content
}
