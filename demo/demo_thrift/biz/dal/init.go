package dal

import (
	"github.com/lgxyc/gomall/demo/demo_thrift/biz/dal/mysql"
	"github.com/lgxyc/gomall/demo/demo_thrift/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
