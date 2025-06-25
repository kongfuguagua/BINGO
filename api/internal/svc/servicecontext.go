package svc

import (
	"dl/api/internal/config"
	"dl/db"
	"dl/rpc/app/dlfunction"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	DLclient dlfunction.DLfunction
	Model    db.WorldModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		DLclient: dlfunction.NewDLfunction(zrpc.MustNewClient(c.DLApp)),
		Model:    db.NewWorldModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
