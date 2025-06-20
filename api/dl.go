package main

import (
	"context"
	"flag"
	"fmt"

	"dl/api/internal/config"
	"dl/api/internal/handler"
	"dl/api/internal/svc"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/dl-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)

	svcCtx := svc.NewServiceContext(c)
	ServiceGroup := service.NewServiceGroup()
	defer ServiceGroup.Stop()

	handler.RegisterHandlers(server, svcCtx)

	ServiceGroup.Add(server)

	kqueue := kq.MustNewQueue(c.KqConsumerConf, kq.WithHandle(func(ctx context.Context, key, val string) error {
		fmt.Println("key: ", key, "val: ", val)
		return nil
	}))

	ServiceGroup.Add(kqueue)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	ServiceGroup.Start()
	// server.Start()
}
