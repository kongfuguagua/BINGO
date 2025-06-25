package logic

import (
	"context"

	"dl/rpc/model/internal/svc"
	"dl/rpc/model/modeler"

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

func (l *InitDLModelLogic) InitDLModel(in *modeler.SetDLModelRequest) (*modeler.DLModel, error) {
	// todo: add your logic here and delete this line

	return &modeler.DLModel{}, nil
}
