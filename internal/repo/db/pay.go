package db

import (
	"context"
	"database/sql"
	"github.com/Zkeai/MuPay-Go/common/logger"
	"time"
)

type YuPay struct {
	ID         int       `json:"id" gorm:"id"`                   //id
	Name       string    `json:"name" gorm:"name"`               //名称
	Code       string    `json:"code" gorm:"code"`               //支付代码
	Commodity  uint8     `json:"commodity" gorm:"commodity"`     //前台状态
	Recharge   uint8     `json:"recharge" gorm:"recharge"`       //充值状态
	Handle     string    `json:"handle" gorm:"handle"`           //支付平台
	Sort       uint16    `json:"sort" gorm:"sort"`               //排序
	Equipment  uint8     `json:"equipment" gorm:"equipment"`     //设备：0=通用 1=手机 2=电脑
	Cost       float64   `json:"cost" gorm:"cost"`               //手续费
	CostType   *uint8    `json:"cost_type" gorm:"cost_type"`     //手续费模式：0=单笔固定，1=百分比
	CreateTime time.Time `json:"create_time" gorm:"create_time"` //创建时间
}

const (
	insertPaySQL = `INSERT INTO yu_pay (name, code, commodity, recharge, handle, sort, equipment, cost, cost_type)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	queryPaySQL  = `SELECT id, name,create_time,icon,status FROM yu_category WHERE owner = ? AND hide != 1 ORDER BY sort ASC`
	updatePaySQL = `UPDATE yu_pay SET name = ?, code = ?, commodity = ?, recharge = ?, handle = ?, sort = ?, equipment = ?, cost = ?, cost_type = ? WHERE id = ?`
	deletePaySQL = `DELETE FROM yu_pay WHERE id = ?`
)

func (db *DB) InsertPay(ctx context.Context, pay *YuPay) (string, error) {

	_, err := db.db.Exec(ctx, insertPaySQL, pay.Name, pay.Code, pay.Commodity, pay.Recharge, pay.Handle, pay.Sort, pay.Equipment, pay.Cost, pay.CostType)

	return "success", err
}

func (db *DB) GetYuPayByID(ctx context.Context, id int) ([]YuPay, error) {
	rows, err := db.db.Query(ctx, queryPaySQL, id)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			logger.Error(err)
		}
	}(rows)
	var pays []YuPay
	for rows.Next() {
		var pay YuPay
		if err := rows.Scan(
			&pay.ID, &pay.Name, &pay.Code, &pay.Commodity, &pay.Recharge,
			&pay.Handle, &pay.Sort, &pay.Equipment, &pay.Cost, &pay.CostType, &pay.CreateTime,
		); err != nil {
			return nil, err
		}
		pays = append(pays, pay)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return pays, nil
}

func (db *DB) UpdateYuPay(ctx context.Context, pay *YuPay) error {

	_, err := db.db.Exec(ctx, updatePaySQL, pay.Name, pay.Code, pay.Commodity, pay.Recharge, pay.Handle, pay.Sort, pay.Equipment, pay.Cost, pay.CostType, pay.ID)
	return err
}

func (db *DB) DeleteYuPay(ctx context.Context, id int) error {

	_, err := db.db.Exec(ctx, deletePaySQL, id)
	return err
}
