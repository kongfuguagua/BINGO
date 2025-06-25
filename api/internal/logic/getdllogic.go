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
	l.Logger.Infof("Received id: %s", req.Id)
	record, err := l.svcCtx.Model.FindOne(l.ctx, req.Id)

	if err != nil {
		return nil, err
	}

	resp = &types.DLGetResponse{
		DL: types.DLApp{
			Metadata: types.DLMetadata{
				Id:        record.ID,
				Namespace: record.Namespace,
				DLName:    record.DLName,
			},
			Spec: types.DLSpec{},
		},
	}

	return resp, nil
}
