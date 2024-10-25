package db

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Zkeai/MuPay-Go/common/logger"
	"time"
)

type YuBusiness struct {
	ID            int64     `json:"id" gorm:"id"`                         // 主键id
	UserId        int64     `json:"user_id" gorm:"user_id"`               // 用户id
	ShopName      string    `json:"shop_name" gorm:"shop_name"`           // 店铺名称
	Title         string    `json:"title" gorm:"title"`                   // 浏览器标题
	Notice        string    `json:"notice" gorm:"notice"`                 // 店铺公告
	ServiceUrl    *string   `json:"service_url" gorm:"service_url"`       // 网页客服链接
	Subdomain     *string   `json:"subdomain" gorm:"subdomain"`           // 子域名
	Topdomain     *string   `json:"topdomain" gorm:"topdomain"`           // 顶级域名
	MasterDisplay int8      `json:"master_display" gorm:"master_display"` // 主站显示：0=否，1=是
	CreateTime    time.Time `json:"create_time" gorm:"create_time"`       // 创建时间
}

const (
	insertBusinessSQL = `INSERT INTO yu_business (user_id,shop_name,title,notice) VALUES (?,?,?,?)`
	queryBusinessSql  = "SELECT  user_id, shop_name, title, notice, service_url, master_display, create_time FROM yu_business WHERE topdomain = ? OR (topdomain IS NULL AND subdomain = ?) LIMIT 1"
)

func (db *DB) InsertBusiness(ctx context.Context, userID int64) (string, error) {

	_, err := db.db.Exec(ctx, insertBusinessSQL, userID, "木鱼店铺", "木鱼店铺", "本程序为开源程序，使用者造成的一切法律后果与作者无关。")
	if err != nil {
		return "", err
	}

	return "success", nil
}

func (db *DB) QueryBusiness(ctx context.Context, host string) (*YuBusiness, error) {
	row := db.db.QueryRow(ctx, queryBusinessSql, host, host)
	u := &YuBusiness{}
	err := row.Scan(
		&u.UserId,
		&u.ShopName,
		&u.Title,
		&u.Notice,
		&u.ServiceUrl,
		&u.MasterDisplay,
		&u.CreateTime,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		logger.Error(err)
		return nil, err
	}
	return u, err
}
