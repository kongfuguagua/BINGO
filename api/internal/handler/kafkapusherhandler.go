package handler

import (
	"net/http"

	"dl/api/internal/logic"
	"dl/api/internal/svc"
	"dl/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func KafkaPusherHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.KafkaPusherRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewKafkaPusherLogic(r.Context(), svcCtx)
		resp, err := l.KafkaPusher(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
