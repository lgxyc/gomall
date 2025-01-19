package dal

import (
	"github.com/lgxyc/gomall/app/cart/biz/dal/mysql"
	"github.com/lgxyc/gomall/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
