package logic

import (
	"context"

	"dl/rpc/app/dl"
	"dl/rpc/app/internal/svc"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDLLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateDLLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDLLogic {
	return &CreateDLLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateDLLogic) CreateDL(in *dl.DLCreateRequest) (*dl.DLCreateResponse, error) {
	DLModel := dl.DLModel{
		Name:       in.Spec.ModelName,
		Path:       "",
		Status:     false,
		InputType:  "",
		OutputType: "",
	}
	DLDataOBJ := dl.DLDataOBJ{
		Status: false,
	}
	DLSpec := dl.DLSpec{Model: &DLModel, DataObj: &DLDataOBJ}
	DLId := uuid.New().String()
	DLMetadata := dl.DLMetadata{
		Namespace: in.Spec.Namespace,
		Id:        DLId,
		DLName:    in.Spec.DLName,
	}
	DlApp := &dl.DLapp{Metadata: &DLMetadata, Spec: &DLSpec}
	return &dl.DLCreateResponse{DlApp: DlApp}, nil
}
