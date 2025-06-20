package logic

import (
	"context"
	"fmt"

	"dl/api/internal/svc"
	"dl/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type KafkaPusherLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKafkaPusherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KafkaPusherLogic {
	return &KafkaPusherLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KafkaPusherLogic) KafkaPusher(req *types.KafkaPusherRequest) (resp *types.KafkaPusherResponse, err error) {
	// todo: add your logic here and delete this line
	data := req.Data
	out := &types.KafkaPusherResponse{
		Result: "",
	}
	if err := l.svcCtx.KqPusherClient.Push(l.ctx, data); err != nil {
		logx.Errorf("KqPusherClient Push Error , err :%v", err)
		out.Result = "failed"
		return out, err
	}
	fmt.Printf("KqPusherClient Push Success, data :%v", data)
	out.Result = "success"
	return out, nil
}
