package logic

import (
	"context"

	"dl/rpc/app/dl"
	"dl/rpc/app/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDLByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDLByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDLByIdLogic {
	return &GetDLByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDLByIdLogic) GetDLById(in *dl.DLGetRequestById) (*dl.DLGetResponseById, error) {
	// todo: add your logic here and delete this line

	return &dl.DLGetResponseById{}, nil
}
