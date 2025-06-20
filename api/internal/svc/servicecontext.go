package svc

import (
	"dl/api/internal/config"
	"dl/rpc/app/dlfunction"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	DLclient       dlfunction.DLfunction
	KqPusherClient *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		DLclient:       dlfunction.NewDLfunction(zrpc.MustNewClient(c.DLApp)),
		KqPusherClient: kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic),
	}
}
