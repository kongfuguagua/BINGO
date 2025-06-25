package handler

import (
	"net/http"

	"dl/api/internal/logic"
	"dl/api/internal/svc"
	"dl/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateDLHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DLCreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCreateDLLogic(r.Context(), svcCtx)
		resp, err := l.CreateDL(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
