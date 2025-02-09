package handler

import (
	"net/http"

	"dl/app/taskmanager/internal/logic"
	"dl/app/taskmanager/internal/svc"
	"dl/app/taskmanager/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func InfoTaskHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TaskInfoRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewInfoTaskLogic(r.Context(), svcCtx)
		resp, err := l.InfoTask(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
