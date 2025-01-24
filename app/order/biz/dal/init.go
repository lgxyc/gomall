package dal

import (
	"github.com/lgxyc/gomall/app/order/biz/dal/mysql"
	"github.com/lgxyc/gomall/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
