package dal

import (
	"github.com/lgxyc/gomall/app/cart/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
