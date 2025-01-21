package dal

import (
	"github.com/lgxyc/gomall/app/payment/biz/dal/mysql"
	"github.com/lgxyc/gomall/app/payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
