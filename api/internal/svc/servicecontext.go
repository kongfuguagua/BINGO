package svc

import (
	"dl/api/internal/config"
	"dl/rpc/app/dlfunction"
	"dl/rpc/model/dlmodeler"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	DLclient  dlfunction.DLfunction
	DLmodeler dlmodeler.DLModeler
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		DLclient:  dlfunction.NewDLfunction(zrpc.MustNewClient(c.DLApp)),
		DLmodeler: dlmodeler.NewDLModeler(zrpc.MustNewClient(c.DLModel)),
	}
}
