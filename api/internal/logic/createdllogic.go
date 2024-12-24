package logic

import (
	"context"

	"dl/api/internal/svc"
	"dl/api/internal/types"
	"dl/rpc/app/dlfunction"

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
	dlinfo, err := l.svcCtx.DLclient.CreateDL(l.ctx, in)
	if err != nil {
		return nil, err
	}
	// dlmodel, err := l.svcCtx.DLmodeler.InitDLModel(l.ctx, &dlmodeler.SetDLModelRequest{Path: dlinfo.DlApp.Spec.Model.Path})
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
					// Path:       dlmodel.Path,
					// Status:     dlmodel.Status,
					// InputType:  dlmodel.InputType,
					// OutputType: dlmodel.OutputType,
				},
			},
		},
	}
	return out, nil
}
