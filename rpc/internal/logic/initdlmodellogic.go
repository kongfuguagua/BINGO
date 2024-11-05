package logic

import (
	"context"

	"dl/rpc/dl"
	"dl/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitDLModelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitDLModelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDLModelLogic {
	return &InitDLModelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InitDLModelLogic) InitDLModel(in *dl.SetDLModelRequest) (*dl.DLModel, error) {
	// todo: add your logic here and delete this line

	return &dl.DLModel{}, nil
}
