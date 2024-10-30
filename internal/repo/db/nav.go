package db

import (
	"context"
	"database/sql"
	"github.com/Zkeai/MuPay-Go/common/logger"
)

type NavItem struct {
	ID         int    `json:"id"`
	ItemKey    string `json:"itemKey"`
	Text       string `json:"text"`
	Icon       string `json:"icon"`
	ParentID   int    `json:"parentId"`
	Roles      string `json:"roles"`
	OrderIndex int    `json:"orderIndex"`
}

const (
	insertNavItemSQL = `INSERT INTO yu_nav (item_key, text, icon, parent_id, roles, order_index)
VALUES (?, ?, ?, ?, ?, ?);`
	queryNavItemSQL = `SELECT * FROM yu_nav WHERE JSON_CONTAINS(CAST(roles AS JSON), ?, '$');`
)

func (db *DB) InsertNavItem(ctx context.Context, nav *NavItem) (string, error) {

	_, err := db.db.Exec(ctx, insertNavItemSQL, nav.ItemKey, nav.Text, nav.Icon, nav.ParentID, nav.Roles, nav.OrderIndex)
	if err != nil {
		return "error", err
	}

	return "success", nil
}

func (db *DB) QueryNavItem(ctx context.Context, role string) ([]NavItem, error) {
	rows, err := db.db.Query(ctx, queryNavItemSQL, role)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			logger.Error(err)
		}
	}(rows)
	var navItems []NavItem
	for rows.Next() {
		var navItem NavItem
		if err := rows.Scan(&navItem.ID, &navItem.ItemKey, &navItem.Text, &navItem.Icon, &navItem.ParentID, &navItem.OrderIndex, &navItem.Roles); err != nil {
			return nil, err
		}
		navItems = append(navItems, navItem)
	}

	return navItems, err
}
