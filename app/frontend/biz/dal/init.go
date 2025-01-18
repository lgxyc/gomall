package dal

import (
	"github.com/lgxyc/gomall/app/frontend/biz/dal/mysql"
	"github.com/lgxyc/gomall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}