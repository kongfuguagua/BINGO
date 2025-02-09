package logic

import (
	"context"

	"dl/app/taskmanager/internal/svc"
	"dl/app/taskmanager/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInfoTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoTaskLogic {
	return &InfoTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InfoTaskLogic) InfoTask(req *types.TaskInfoRequest) (resp *types.TaskInfoResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
