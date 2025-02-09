package logic

import (
	"context"

	"dl/app/taskmanager/internal/svc"
	"dl/app/taskmanager/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListTaskLogic {
	return &ListTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListTaskLogic) ListTask() (resp *types.TaskListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
