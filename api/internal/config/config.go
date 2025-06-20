package config

import (
	"github.com/zeromicro/go-queue/kq"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	DLApp        zrpc.RpcClientConf
	DLModel      zrpc.RpcClientConf
	KqPusherConf struct {
		Brokers []string
		Topic   string
	}
	KqConsumerConf kq.KqConf
}
