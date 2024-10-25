package db

import (
	"context"
	"database/sql"
	"github.com/Zkeai/MuPay-Go/common/logger"
	"time"
)

type YuCategory struct {
	ID              int64     `json:"id" gorm:"id"`                               // 主键id
	Name            string    `json:"name" gorm:"name"`                           // 商品分类名称
	Sort            int16     `json:"sort" gorm:"sort"`                           // 排序
	CreateTime      time.Time `json:"create_time" gorm:"create_time"`             // 创建时间
	Owner           int64     `json:"owner" gorm:"owner"`                         // 所属会员：0=系统，其他等于会员UID
	Icon            string    `json:"icon" gorm:"icon"`                           // 分类图标
	Status          int8      `json:"status" gorm:"status"`                       // 状态：0=停用，1=启用
	Hide            int8      `json:"hide" gorm:"hide"`                           // 隐藏：1=隐藏，0=不隐藏
	UserLevelConfig string    `json:"user_level_config" gorm:"user_level_config"` // 会员配置
}

const (
	insertCategorySQL = `INSERT INTO yu_category (name,owner,sort,icon) VALUES (?,?,?,?)`
	queryCategorySQL  = `SELECT id, name,create_time,icon,status FROM yu_category WHERE owner = ? AND hide != 1 ORDER BY sort ASC`
)

func (db *DB) InsertCategory(ctx context.Context, name string, userID int64, sort *int8, icon string) (string, error) {

	_, err := db.db.Exec(ctx, insertCategorySQL, name, userID, sort, icon)
	if err != nil {
		return "", err
	}

	return "success", nil
}

func (db *DB) QueryCategory(ctx context.Context, userid int64) ([]YuCategory, error) {
	rows, err := db.db.Query(ctx, queryCategorySQL, userid)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			logger.Error(err)
		}
	}(rows)

	var categories []YuCategory
	for rows.Next() {
		var category YuCategory
		if err := rows.Scan(&category.ID, &category.Name, &category.CreateTime, &category.Icon, &category.Status); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}
