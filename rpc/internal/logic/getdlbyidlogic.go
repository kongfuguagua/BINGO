package logic

import (
	"context"

	"dl/rpc/dl"
	"dl/rpc/internal/svc"

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
	// todo: 这里大概要做数据库操作，通过id查spec

	return &dl.DLGetResponseById{}, nil
}
