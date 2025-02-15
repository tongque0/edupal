package dal

import (
	"github.com/tongque0/edupal/biz/dal/cache"
	"github.com/tongque0/edupal/biz/dal/mysql"
)

func Init() {
	mysql.Init()
	cache.Init()
}
