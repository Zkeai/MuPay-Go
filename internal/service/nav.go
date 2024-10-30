package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Zkeai/MuPay-Go/common/logger"
	"github.com/Zkeai/MuPay-Go/common/middleware"
	"github.com/Zkeai/MuPay-Go/common/redis"
	"github.com/Zkeai/MuPay-Go/internal/repo/db"
	"strconv"
)

type NavItem struct {
	ID         int       `json:"id"`
	ItemKey    string    `json:"itemKey"`
	Text       string    `json:"text"`
	Icon       string    `json:"icon"`
	ParentID   int       `json:"parentId,omitempty"`
	Roles      string    `json:"roles,omitempty"`
	OrderIndex int       `json:"orderIndex,omitempty"`
	Items      []NavItem `json:"items,omitempty"`
}

func (s *Service) NavAdd(ctx context.Context, nav *db.NavItem) (string, error) {

	nav_, err := s.repo.CreateNav(ctx, nav)

	if err != nil {
		return "err", err
	}
	return nav_, nil
}

func (s *Service) QueryNav(ctx context.Context, walletAddress string) (json.RawMessage, error) {

	//通过address得到权限
	result, err := redis.GetClient().Get(redis.Ctx, walletAddress).Result()
	var sessionData middleware.SessionData
	err = json.Unmarshal([]byte(result), &sessionData)
	if err != nil {
		logger.Error("Failed to unmarshal JSON data: %v", err)
		return json.RawMessage(""), err
	}
	role := sessionData.Role
	roleStr := strconv.Itoa(role)
	nav_, err := s.repo.GetNav(ctx, roleStr)

	// Step 1: Build a map of each nav item
	navMap := make(map[int]*NavItem)
	for _, item := range nav_ {
		navMap[item.ID] = &NavItem{
			ID:      item.ID,
			ItemKey: item.ItemKey,
			Text:    item.Text,
			Icon:    item.Icon,
		}
	}

	// Step 2: Establish parent-child relationships
	for _, item := range nav_ {
		if item.ParentID > 0 {
			if parent, exists := navMap[item.ParentID]; exists {
				parent.Items = append(parent.Items, *navMap[item.ID])
			}
		}
	}

	// Step 3: Collect only top-level items for the final output
	var navSlice []NavItem
	for _, item := range nav_ {
		if item.ParentID == 0 {
			navSlice = append(navSlice, *navMap[item.ID])
		} else {
			if parent, exists := navMap[item.ParentID]; exists {
				parent.Items = append(parent.Items, *navMap[item.ID])
			}
		}
	}

	// Step 4: Marshal the result into JSON
	jsonData, err := json.MarshalIndent(navSlice, "", "  ")
	if err != nil {
		return json.RawMessage(""), fmt.Errorf("error marshalling JSON: %v", err)
	}

	// 将字节数组转换为字符串并返回

	jsonString := json.RawMessage(jsonData)

	return jsonString, nil
}
