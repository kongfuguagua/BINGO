package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	DLApp      zrpc.RpcClientConf
	DLModel    zrpc.RpcClientConf
	DataSource string
	Cache      cache.CacheConf
}
