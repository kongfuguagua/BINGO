package logic

import (
	"context"

	"dl/app/taskmanager/internal/svc"
	"dl/app/taskmanager/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StartTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStartTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartTaskLogic {
	return &StartTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StartTaskLogic) StartTask(req *types.TaskStartRequest) (resp *types.TaskStartResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
