package server

import (
	chttp "github.com/Zkeai/MuPay-Go/common/net/cttp"
	"github.com/Zkeai/MuPay-Go/internal/conf"

	"github.com/Zkeai/MuPay-Go/internal/handler"
	"github.com/Zkeai/MuPay-Go/internal/service"
)

func NewHTTP(conf *conf.Conf) *chttp.Server {
	s := chttp.NewServer(conf.Server)

	svc := service.NewService(conf)
	handler.InitRouter(s, svc)

	err := s.Start()

	if err != nil {
		panic(err)
	}

	return s
}
