package dal

import (
	"github.com/lgxyc/gomall/app/product/biz/dal/mysql"
	"github.com/lgxyc/gomall/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
