package svc

import (
	"dl/api/internal/config"
	"dl/rpc/dlfunction"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	DLcreate dlfunction.DLfunction
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		DLcreate: dlfunction.NewDLfunction(zrpc.MustNewClient(c.DLApp)),
	}
}
