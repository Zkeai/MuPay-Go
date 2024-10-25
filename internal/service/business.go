package service

import (
	"context"
	"github.com/Zkeai/MuPay-Go/internal/repo/db"
)

func (s *Service) BusinessRegister(ctx context.Context, userID int64) (string, error) {

	business, err := s.repo.BusinessRegister(ctx, userID)
	if err != nil {
		return "", err
	}

	return business, nil
}

func (s *Service) BusinessQuery(ctx context.Context, host string) (*db.YuBusiness, error) {

	businessModel, err := s.repo.BusinessQuery(ctx, host)
	if err != nil {
		return nil, err
	}
	return businessModel, nil
}
