package logic

import (
	"context"

	"dl/api/internal/svc"
	"dl/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDLLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDLLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDLLogic {
	return &GetDLLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDLLogic) GetDL(req *types.DLGetRequest) (resp *types.DLGetResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
