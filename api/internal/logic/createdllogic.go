package logic

import (
	"context"

	"dl/api/internal/svc"
	"dl/api/internal/types"
	"dl/rpc/dlfunction"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDLLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateDLLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDLLogic {
	return &CreateDLLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateDLLogic) CreateDL(req *types.DLCreateRequest) (resp *types.DLCreateResponse, err error) {
	spec := &dlfunction.DLCreateSpec{
		Namespace: req.Namespace,
		DLName:    req.DLName,
		ModelName: req.Spec.Model.Name,
	}
	in := &dlfunction.DLCreateRequest{
		Spec: spec,
	}
	dlinfo, err := l.svcCtx.DLcreate.CreateDL(l.ctx, in)
	if err != nil {
		return nil, err
	}
	out := &types.DLCreateResponse{
		DLInfo: types.DLApp{
			Metadata: types.DLMetadata{
				Id:        dlinfo.DlApp.Metadata.Id,
				Namespace: dlinfo.DlApp.Metadata.Namespace,
				DLName:    dlinfo.DlApp.Metadata.Namespace,
			},
			Spec: types.DLSpec{
				Model: types.DLModel{
					Name: dlinfo.DlApp.Spec.Model.Name,
				},
			},
		},
	}
	return out, nil
}
