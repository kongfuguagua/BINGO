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
	// todo: 这里大概要做校验model是否可用，并获取信息

	return &dl.DLModel{}, nil
}
