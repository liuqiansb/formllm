package dao

import (
	"context"
	"database/sql"
	"github.com/jinzhu/gorm"
)

type Dao struct {
	DB *gorm.DB
}

type DBInterface interface {
	BeginTransaction(ctx context.Context) (tx *gorm.DB)
	UserDao
}

func (d *Dao) BeginTransaction(ctx context.Context) (tx *gorm.DB) {
	tx = d.DB.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
		ReadOnly:  false,
	})
	return
}
