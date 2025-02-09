package logic

import (
	"context"

	"dl/app/taskmanager/internal/svc"
	"dl/app/taskmanager/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTaskLogic {
	return &CreateTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTaskLogic) CreateTask(req *types.TaskCreateRequest) (resp *types.TaskCreateResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
