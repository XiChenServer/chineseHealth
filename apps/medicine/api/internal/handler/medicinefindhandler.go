package handler

import (
	"chineseHealthy/pkg/response"
	"net/http"

	"chineseHealthy/apps/medicine/api/internal/logic"
	"chineseHealthy/apps/medicine/api/internal/svc"
	"chineseHealthy/apps/medicine/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	//"api_v2/common/response"
)

func MedicineFindHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MedicineFindRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewMedicineFindLogic(r.Context(), svcCtx)
		resp, err := l.MedicineFind(&req)
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//
		//}
		response.HttpResult(r, w, resp, err)

	}
}
