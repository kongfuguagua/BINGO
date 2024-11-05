package logic

import (
	"context"

	"dl/rpc/dl"
	"dl/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDLLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateDLLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDLLogic {
	return &CreateDLLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateDLLogic) CreateDL(in *dl.DLCreateRequest) (*dl.DLCreateResponse, error) {
	// todo: add your logic here and delete this line

	return &dl.DLCreateResponse{}, nil
}
