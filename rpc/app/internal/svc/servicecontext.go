package svc

import (
	"dl/pkg/zetcd"
	"dl/rpc/app/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	ZEtcdClient *zetcd.ZEtcdClient
}

// NewServiceContext creates a new ServiceContext with the provided configuration.
func NewServiceContext(c config.Config) *ServiceContext {
	ZEtcdClient, err := zetcd.NewZEtcdClient(c.ZEtcdConf)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:      c,
		ZEtcdClient: ZEtcdClient,
	}
}
