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
	// todo: add your logic here and delete this line

	return &dl.DLDataOBJ{}, nil
}
