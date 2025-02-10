package dal

import (
	"github.com/lgxyc/gomall/app/order/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
