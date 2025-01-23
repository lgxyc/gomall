package dal

import (
	"github.com/lgxyc/gomall/app/checkout/biz/dal/mysql"
	"github.com/lgxyc/gomall/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
