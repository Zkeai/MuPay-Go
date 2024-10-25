package service

import (
	"context"
	"errors"
	"github.com/Zkeai/MuPay-Go/internal/repo/db"
	"github.com/google/uuid"
	"strings"
	"time"
)

func (s *Service) CreateCommodity(ctx context.Context, commodity *db.YuCommodity) (string, error) {
	// 辅助函数：如果值为零，设置默认值
	setDefault := func(value, defaultValue interface{}) interface{} {
		switch v := value.(type) {
		case int:
			if v == 0 {
				return defaultValue
			}
		case float64:
			if v == 0 {
				return defaultValue
			}
		case string:
			if v == "" {
				return defaultValue
			}
		case time.Time:
			if v.IsZero() {
				return defaultValue
			}
		}
		return value
	}

	// 使用辅助函数为非必需字段设置默认值
	code := strings.ReplaceAll(uuid.New().String(), "-", "")
	commodity.CategoryId = setDefault(commodity.CategoryId, 1).(int64)
	commodity.Name = setDefault(commodity.Name, commodity.Name).(string)
	commodity.Description = setDefault(commodity.Description, "默认商品描述").(string)
	commodity.Cover = setDefault(commodity.Cover, "/image/default_cover.png").(*string)
	commodity.FactoryPrice = setDefault(commodity.FactoryPrice, 0.00).(float64)
	commodity.Price = setDefault(commodity.Price, 0.00).(float64)
	commodity.UserPrice = setDefault(commodity.UserPrice, 0.00).(float64)
	commodity.Status = setDefault(commodity.Status, 1).(int8)
	commodity.Owner = setDefault(commodity.Owner, 0).(int64)
	commodity.ApiStatus = setDefault(commodity.ApiStatus, 0).(int8)
	commodity.Code = setDefault(commodity.Code, code).(string)
	commodity.DeliveryWay = setDefault(commodity.DeliveryWay, 0).(int8)
	commodity.DeliveryAutoMode = setDefault(commodity.DeliveryAutoMode, 0).(int8)
	commodity.DeliveryMessage = setDefault(commodity.DeliveryMessage, "").(string)
	commodity.ContactType = setDefault(commodity.ContactType, 0).(int8)
	commodity.PasswordStatus = setDefault(commodity.PasswordStatus, 0).(int8)
	commodity.Sort = setDefault(commodity.Sort, 100).(int16)
	commodity.Coupon = setDefault(commodity.Coupon, 0).(int8)
	commodity.SharedId = setDefault(commodity.SharedId, 0).(int64)
	commodity.SharedPremium = setDefault(commodity.SharedPremium, 0.00).(float32)
	commodity.SharedPremiumType = setDefault(commodity.SharedPremiumType, 0).(int8)
	commodity.SeckillStatus = setDefault(commodity.SeckillStatus, 0).(int8)
	commodity.SeckillStartTime = setDefault(commodity.SeckillStartTime, time.Now()).(*time.Time)
	commodity.SeckillEndTime = setDefault(commodity.SeckillEndTime, time.Now().Add(24*time.Hour)).(*time.Time)
	commodity.DraftStatus = setDefault(commodity.DraftStatus, 0).(int8)
	commodity.DraftPremium = setDefault(commodity.DraftPremium, 0.00).(*float64)
	commodity.InventoryHidden = setDefault(commodity.InventoryHidden, 0).(int8)
	commodity.LeaveMessage = setDefault(commodity.LeaveMessage, "发货留言").(*string)
	commodity.Recommend = setDefault(commodity.Recommend, 0).(int8)
	commodity.SendEmail = setDefault(commodity.SendEmail, 0).(int8)
	commodity.OnlyUser = setDefault(commodity.OnlyUser, 0).(int8)
	commodity.PurchaseCount = setDefault(commodity.PurchaseCount, 0).(int64)
	commodity.Widget = setDefault(commodity.Widget, nil).(*string)
	commodity.LevelPrice = setDefault(commodity.LevelPrice, "").(*string)
	commodity.LevelDisable = setDefault(commodity.LevelDisable, 0).(int8)
	commodity.Minimum = setDefault(commodity.Minimum, 0).(int64)
	commodity.Maximum = setDefault(commodity.Maximum, 0).(int64)
	commodity.SharedSync = setDefault(commodity.SharedSync, 0).(int8)
	commodity.Config = setDefault(commodity.Config, nil).(*string)
	commodity.Hide = setDefault(commodity.Hide, 0).(int8)
	commodity.InventorySync = setDefault(commodity.InventorySync, 0).(int8)

	category, err := s.repo.CreateCommodity(ctx, commodity)
	if err != nil {
		return "", err
	}

	return category, nil
}

func (s *Service) GetCommodity(ctx context.Context, id int64) ([]db.YuCommodity, error) {

	commoditys, err := s.repo.GetCommodityByID(ctx, id)
	if err != nil {
		println(err.Error())
		if err.Error() == `sql: no rows in result set` {
			return nil, errors.New("当前分类没有任何商品")
		}
		return nil, err
	}

	return commoditys, nil
}
