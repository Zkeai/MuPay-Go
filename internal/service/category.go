package service

import (
	"context"
	"github.com/Zkeai/MuPay-Go/internal/repo/db"
)

func (s *Service) CategoryAdd(ctx context.Context, name string, userID int64, sort *int8, icon string) (string, error) {
	if sort == nil {
		defaultSort := int8(0)
		sort = &defaultSort
	}
	if icon == "" {
		icon = "icon-shangpinguanli"
	}
	category, err := s.repo.CategoryAdd(ctx, name, userID, sort, icon)
	if err != nil {
		return "", err
	}

	return category, nil
}

func (s *Service) CategoryQuery(ctx context.Context, userid int64) ([]db.YuCategory, error) {

	categoryModel, err := s.repo.CategoryQuery(ctx, userid)
	if err != nil {
		return nil, err
	}
	return categoryModel, nil
}
