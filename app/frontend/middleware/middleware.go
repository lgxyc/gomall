package middleware

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

func Register(h *server.Hertz) {
	// 鉴权
	h.Use(GlobalAuth())
}
