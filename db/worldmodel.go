package db

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ WorldModel = (*customWorldModel)(nil)

type (
	// WorldModel is an interface to be customized, add more methods here,
	// and implement the added methods in customWorldModel.
	WorldModel interface {
		worldModel
	}

	customWorldModel struct {
		*defaultWorldModel
	}
)

// NewWorldModel returns a model for the database table.
func NewWorldModel(conn sqlx.SqlConn, c cache.CacheConf) WorldModel {
	return &customWorldModel{
		defaultWorldModel: newWorldModel(conn, c),
	}
}
