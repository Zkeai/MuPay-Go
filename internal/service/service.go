package service

import "github.com/Zkeai/MuPay-Go/internal/conf"
import "github.com/Zkeai/MuPay-Go/internal/repo"

type Service struct {
	conf *conf.Conf
	repo *repo.Repo
}

func NewService(conf *conf.Conf) *Service {
	return &Service{
		conf: conf,
		repo: repo.NewRepo(conf),
	}
}
