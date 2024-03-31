package handler

import (
	"net/http"

	//"api_v2/common/response"
	"chineseHealthy/apps/medicine/api/internal/logic"
	"chineseHealthy/apps/medicine/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ImageUpdataHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewImageUpdataLogic(r.Context(), svcCtx)
		resp, err := l.ImageUpdata(r)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
		//response.Response(r, w, resp, err)

	}
}
