package handler

import (
	"chineseHealthy/pkg/response"
	"net/http"

	//"api_v2/common/response"
	"chineseHealthy/apps/medicine/api/internal/logic"
	"chineseHealthy/apps/medicine/api/internal/svc"
)

func ImageUpdataHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewImageUpdataLogic(r.Context(), svcCtx)

		resp, err := l.ImageUpdata(r)
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//
		//}
		response.HttpResult(r, w, resp, err)

	}
}
