package logic

import (
	"context"

	"dl/rpc/dl"
	"dl/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitDLDataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitDLDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDLDataLogic {
	return &InitDLDataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InitDLDataLogic) InitDLData(in *dl.SetDLDataRequest) (*dl.DLDataOBJ, error) {
	// todo: 这里大概要做校验数据对象是否可靠

	return &dl.DLDataOBJ{}, nil
}
