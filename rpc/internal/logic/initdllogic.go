package logic

import (
	"context"

	"dl/rpc/dl"
	"dl/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitDLLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitDLLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDLLogic {
	return &InitDLLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InitDLLogic) InitDL(in *dl.DLapp) (*dl.DLapp, error) {
	// todo: add your logic here and delete this line

	return &dl.DLapp{}, nil
}
