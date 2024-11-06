package svc

import (
	"dl/api/internal/config"
	"dl/rpc/dlfunction"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	DLApp  dlfunction.DLfunction
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DLApp:  dlfunction.NewDLfunction(zrpc.MustNewClient(c.DLApp)),
	}
}
