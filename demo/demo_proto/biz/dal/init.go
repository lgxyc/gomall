package dal

import (
	"github.com/lgxyc/gomall/demo/demo_proto/biz/dal/mysql"
	"github.com/lgxyc/gomall/demo/demo_proto/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
