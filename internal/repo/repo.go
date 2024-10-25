package repo

import (
	"context"
	"github.com/Zkeai/MuPay-Go/internal/conf"
	"github.com/Zkeai/MuPay-Go/internal/repo/db"
)

type Repo struct {
	db *db.DB
}

func NewRepo(conf *conf.Conf) *Repo {
	return &Repo{
		db: db.NewDB(conf.DB),
	}
}

func (r *Repo) UserRegister(ctx context.Context, wallet string) (*db.YuUser, error) {

	return r.db.InsertUser(ctx, wallet)
}

func (r *Repo) UserQuery(ctx context.Context, wallet string) (*db.YuUser, error) {

	return r.db.QueryUser(ctx, wallet)
}

func (r *Repo) BusinessRegister(ctx context.Context, userID int64) (string, error) {

	return r.db.InsertBusiness(ctx, userID)
}

func (r *Repo) BusinessQuery(ctx context.Context, host string) (*db.YuBusiness, error) {

	return r.db.QueryBusiness(ctx, host)
}

func (r *Repo) CategoryAdd(ctx context.Context, name string, userID int64, sort *int8, icon string) (string, error) {

	return r.db.InsertCategory(ctx, name, userID, sort, icon)
}

func (r *Repo) CategoryQuery(ctx context.Context, userid int64) ([]db.YuCategory, error) {

	return r.db.QueryCategory(ctx, userid)
}

func (r *Repo) CreateCommodity(ctx context.Context, commodity *db.YuCommodity) (string, error) {

	return r.db.CreateCommodity(ctx, commodity)
}

func (r *Repo) GetCommodityByID(ctx context.Context, id int64) ([]db.YuCommodity, error) {

	return r.db.GetCommodityByID(ctx, id)
}

func (r *Repo) UpdateCommodity(ctx context.Context, commodity db.YuCommodity) (string, error) {

	return r.db.UpdateCommodity(ctx, commodity)
}

func (r *Repo) DeleteCommodity(ctx context.Context, id int64) (string, error) {

	return r.db.DeleteCommodity(ctx, id)
}

func (r *Repo) CreatePay(ctx context.Context, pay *db.YuPay) (string, error) {

	return r.db.InsertPay(ctx, pay)
}

func (r *Repo) GetPay(ctx context.Context, id int) ([]db.YuPay, error) {

	return r.db.GetYuPayByID(ctx, id)
}

func (r *Repo) UpdatePay(ctx context.Context, pay *db.YuPay) error {

	return r.db.UpdateYuPay(ctx, pay)
}

func (r *Repo) DeletePay(ctx context.Context, id int) error {

	return r.db.DeleteYuPay(ctx, id)
}
