package handler

import (
	"net/http"

	"chineseHealthy/apps/medicine/api/internal/logic"
	"chineseHealthy/apps/medicine/api/internal/svc"
	"chineseHealthy/apps/medicine/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	//"api_v2/common/response"
)

func ImageDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImageDelRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewImageDeleteLogic(r.Context(), svcCtx)
		resp, err := l.ImageDelete(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
		//response.Response(r, w, resp, err)

	}
}