package dal

import (
	"github.com/lgxyc/gomall/app/user/biz/dal/mysql"
	"github.com/lgxyc/gomall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
