package conf

import (
	"github.com/Zkeai/MuPay-Go/common/database"
	chttp "github.com/Zkeai/MuPay-Go/common/net/cttp"
	"github.com/Zkeai/MuPay-Go/common/redis"
)

type Conf struct {
	Server *chttp.Config    `yaml:"server"`
	DB     *database.Config `yaml:"db"`
	Rdb    *redis.Config    `yaml:"redis"`
}
