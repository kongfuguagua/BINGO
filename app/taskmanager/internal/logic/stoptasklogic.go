package logic

import (
	"context"

	"dl/app/taskmanager/internal/svc"
	"dl/app/taskmanager/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StopTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStopTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StopTaskLogic {
	return &StopTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StopTaskLogic) StopTask(req *types.TaskStopRequest) (resp *types.TaskStopResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
