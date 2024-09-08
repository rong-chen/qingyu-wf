package global

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var MySql *gorm.DB
var Redis *redis.Client
