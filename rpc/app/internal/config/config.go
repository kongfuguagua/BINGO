package config

import (
	"dl/pkg/zetcd"

	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	zetcd.ZEtcdConf
}
