package db

import (
	"context"
	"database/sql"
	"github.com/Zkeai/MuPay-Go/common/logger"
	"time"
)

type YuCommodity struct {
	ID                int64      `json:"id" gorm:"id"`                                   // 主键id
	CategoryId        int64      `json:"category_id" gorm:"category_id"`                 // 商品分类ID
	Name              string     `json:"name" gorm:"name"`                               // 商品名称
	Description       string     `json:"description" gorm:"description"`                 // 商品说明
	Cover             *string    `json:"cover" gorm:"cover"`                             // 商品封面图片
	FactoryPrice      float64    `json:"factory_price" gorm:"factory_price"`             // 成本单价
	Price             float64    `json:"price" gorm:"price"`                             // 商品单价(未登录)
	UserPrice         float64    `json:"user_price" gorm:"user_price"`                   // 商品单价(会员价)
	Status            int8       `json:"status" gorm:"status"`                           // 状态：0=下架，1=上架
	Owner             int64      `json:"owner" gorm:"owner"`                             // 所属会员：0=系统，其他等于会员UID
	CreateTime        time.Time  `json:"create_time" gorm:"create_time"`                 // 创建时间
	ApiStatus         int8       `json:"api_status" gorm:"api_status"`                   // API对接：0=关闭，1=启用
	Code              string     `json:"code" gorm:"code"`                               // 商品代码(API对接)
	DeliveryWay       int8       `json:"delivery_way" gorm:"delivery_way"`               // 发货方式：0=自动发货，1=手动发货/插件发货
	DeliveryAutoMode  int8       `json:"delivery_auto_mode" gorm:"delivery_auto_mode"`   // 自动发卡模式：0=旧卡先发，1=随机发卡，2=新卡先发
	DeliveryMessage   string     `json:"delivery_message" gorm:"delivery_message"`       // 手动发货显示信息
	ContactType       int8       `json:"contact_type" gorm:"contact_type"`               // 联系方式：0=任意，1=手机，2=邮箱，3=QQ
	PasswordStatus    int8       `json:"password_status" gorm:"password_status"`         // 订单密码：0=关闭，1=启用
	Sort              int16      `json:"sort" gorm:"sort"`                               // 排序
	Coupon            int8       `json:"coupon" gorm:"coupon"`                           // 优惠卷：0=关闭，1=启用
	SharedId          int64      `json:"shared_id" gorm:"shared_id"`                     // 共享平台ID
	SharedCode        *string    `json:"shared_code" gorm:"shared_code"`                 // 共享平台-商品代码
	SharedPremium     float32    `json:"shared_premium" gorm:"shared_premium"`           // 商品加价
	SharedPremiumType int8       `json:"shared_premium_type" gorm:"shared_premium_type"` // 加价模式
	SeckillStatus     int8       `json:"seckill_status" gorm:"seckill_status"`           // 商品秒杀：0=关闭，1=开启
	SeckillStartTime  *time.Time `json:"seckill_start_time" gorm:"seckill_start_time"`   // 秒杀开始时间
	SeckillEndTime    *time.Time `json:"seckill_end_time" gorm:"seckill_end_time"`       // 秒杀结束时间
	DraftStatus       int8       `json:"draft_status" gorm:"draft_status"`               // 指定卡密购买：0=关闭，1=启用
	DraftPremium      *float64   `json:"draft_premium" gorm:"draft_premium"`             // 指定卡密购买时溢价
	InventoryHidden   int8       `json:"inventory_hidden" gorm:"inventory_hidden"`       // 隐藏库存：0=否，1=是
	LeaveMessage      *string    `json:"leave_message" gorm:"leave_message"`             // 发货留言
	Recommend         int8       `json:"recommend" gorm:"recommend"`                     // 推荐商品：0=否，1=是
	SendEmail         int8       `json:"send_email" gorm:"send_email"`                   // 发送邮件：0=否，1=是
	OnlyUser          int8       `json:"only_user" gorm:"only_user"`                     // 限制登录购买：0=否，1=是
	PurchaseCount     int64      `json:"purchase_count" gorm:"purchase_count"`           // 限制购买数量：0=无限制
	Widget            *string    `json:"widget" gorm:"widget"`                           // 控件
	LevelPrice        *string    `json:"level_price" gorm:"level_price"`                 // 会员等级-定制价格
	LevelDisable      int8       `json:"level_disable" gorm:"level_disable"`             // 禁用会员等级折扣，0=关闭，1=启用
	Minimum           int64      `json:"minimum" gorm:"minimum"`                         // 最低购买数量，0=无限制
	Maximum           int64      `json:"maximum" gorm:"maximum"`                         // 最大购买数量，0=无限制
	SharedSync        int8       `json:"shared_sync" gorm:"shared_sync"`                 // 同步平台价格：0=关，1=开
	Config            *string    `json:"config" gorm:"config"`                           // 配置文件
	Hide              int8       `json:"hide" gorm:"hide"`                               // 隐藏：1=隐藏，0=不隐藏
	InventorySync     int8       `json:"inventory_sync" gorm:"inventory_sync"`           // 同步库存数量: 0=关，1=开
}

const (
	insertCommoditySQL = `INSERT INTO yu_commodity 
		(category_id, name, description, cover, factory_price, price, user_price, status, owner, api_status, code, delivery_way, delivery_auto_mode, delivery_message, contact_type, password_status, sort, coupon, shared_id, shared_code, shared_premium, shared_premium_type, seckill_status, seckill_start_time, seckill_end_time, draft_status, draft_premium, inventory_hidden, leave_message, recommend, send_email, only_user, purchase_count, widget, level_price, level_disable, minimum, maximum, shared_sync, config, hide, inventory_sync) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	queryCommoditySQL = `SELECT * FROM yu_commodity WHERE category_id = ?`
	updateCommodity   = `UPDATE yu_commodity SET category_id=?, name=?, description=?, cover=?, factory_price=?, price=?, user_price=?, status=?, owner=?, api_status=?, code=?, delivery_way=?, delivery_auto_mode=?, delivery_message=?, contact_type=?, password_status=?, sort=?, coupon=?, shared_id=?, shared_code=?, shared_premium=?, shared_premium_type=?, seckill_status=?, seckill_start_time=?, seckill_end_time=?, draft_status=?, draft_premium=?, inventory_hidden=?, leave_message=?, recommend=?, send_email=?, only_user=?, purchase_count=?, widget=?, level_price=?, level_disable=?, minimum=?, maximum=?, shared_sync=?, config=?, hide=?, inventory_sync=? WHERE id=?`
	deleteCommodity   = `DELETE FROM yu_commodity WHERE id = ?`
)

func (db *DB) CreateCommodity(ctx context.Context, commodity *YuCommodity) (string, error) {

	_, err := db.db.Exec(ctx, insertCommoditySQL, commodity.CategoryId, commodity.Name, commodity.Description, commodity.Cover, commodity.FactoryPrice, commodity.Price, commodity.UserPrice, commodity.Status, commodity.Owner, commodity.ApiStatus, commodity.Code, commodity.DeliveryWay, commodity.DeliveryAutoMode, commodity.DeliveryMessage, commodity.ContactType, commodity.PasswordStatus, commodity.Sort, commodity.Coupon, commodity.SharedId, commodity.SharedCode, commodity.SharedPremium, commodity.SharedPremiumType, commodity.SeckillStatus, commodity.SeckillStartTime, commodity.SeckillEndTime, commodity.DraftStatus, commodity.DraftPremium, commodity.InventoryHidden, commodity.LeaveMessage, commodity.Recommend, commodity.SendEmail, commodity.OnlyUser, commodity.PurchaseCount, commodity.Widget, commodity.LevelPrice, commodity.LevelDisable, commodity.Minimum, commodity.Maximum, commodity.SharedSync, commodity.Config, commodity.Hide, commodity.InventorySync)

	if err != nil {
		return "", err
	}

	return "success", nil
}

func (db *DB) GetCommodityByID(ctx context.Context, id int64) ([]YuCommodity, error) {
	rows, err := db.db.Query(ctx, queryCommoditySQL, id)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			logger.Error(err)
		}
	}(rows)

	var commoditys []YuCommodity

	for rows.Next() {
		var commodity YuCommodity
		if err := rows.Scan(&commodity.ID, &commodity.CategoryId, &commodity.Name, &commodity.Description, &commodity.Cover, &commodity.FactoryPrice, &commodity.Price, &commodity.UserPrice, &commodity.Status, &commodity.Owner, &commodity.CreateTime, &commodity.ApiStatus, &commodity.Code, &commodity.DeliveryWay, &commodity.DeliveryAutoMode, &commodity.DeliveryMessage, &commodity.ContactType, &commodity.PasswordStatus, &commodity.Sort, &commodity.Coupon, &commodity.SharedId, &commodity.SharedCode, &commodity.SharedPremium, &commodity.SharedPremiumType, &commodity.SeckillStatus, &commodity.SeckillStartTime, &commodity.SeckillEndTime, &commodity.DraftStatus, &commodity.DraftPremium, &commodity.InventoryHidden, &commodity.LeaveMessage, &commodity.Recommend, &commodity.SendEmail, &commodity.OnlyUser, &commodity.PurchaseCount, &commodity.Widget, &commodity.LevelPrice, &commodity.LevelDisable, &commodity.Minimum, &commodity.Maximum, &commodity.SharedSync, &commodity.Config, &commodity.Hide, &commodity.InventorySync); err != nil {
			return nil, err
		}

		commoditys = append(commoditys, commodity)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return commoditys, nil
}

func (db *DB) UpdateCommodity(ctx context.Context, commodity YuCommodity) (string, error) {

	_, err := db.db.Exec(ctx, updateCommodity, commodity.CategoryId, commodity.Name, commodity.Description, commodity.Cover, commodity.FactoryPrice, commodity.Price, commodity.UserPrice, commodity.Status, commodity.Owner, commodity.ApiStatus, commodity.Code, commodity.DeliveryWay, commodity.DeliveryAutoMode, commodity.DeliveryMessage, commodity.ContactType, commodity.PasswordStatus, commodity.Sort, commodity.Coupon, commodity.SharedId, commodity.SharedCode, commodity.SharedPremium, commodity.SharedPremiumType, commodity.SeckillStatus, commodity.SeckillStartTime, commodity.SeckillEndTime, commodity.DraftStatus, commodity.DraftPremium, commodity.InventoryHidden, commodity.LeaveMessage, commodity.Recommend, commodity.SendEmail, commodity.OnlyUser, commodity.PurchaseCount, commodity.Widget, commodity.LevelPrice, commodity.LevelDisable, commodity.Minimum, commodity.Maximum, commodity.SharedSync, commodity.Config, commodity.Hide, commodity.InventorySync, commodity.ID)
	if err != nil {
		return "fail", err
	}
	return "success", nil
}

func (db *DB) DeleteCommodity(ctx context.Context, id int64) (string, error) {
	_, err := db.db.Exec(ctx, deleteCommodity, id)

	if err != nil {
		return "fail", err
	}
	return "success", nil
}
