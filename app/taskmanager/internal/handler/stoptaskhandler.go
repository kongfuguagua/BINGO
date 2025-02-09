package handler

import (
	"net/http"

	"dl/app/taskmanager/internal/logic"
	"dl/app/taskmanager/internal/svc"
	"dl/app/taskmanager/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func StopTaskHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TaskStopRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewStopTaskLogic(r.Context(), svcCtx)
		resp, err := l.StopTask(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
