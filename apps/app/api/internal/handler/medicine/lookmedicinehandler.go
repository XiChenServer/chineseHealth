package medicine

import (
	"chineseHealthy/pkg/response"
	"net/http"

	"chineseHealthy/apps/app/api/internal/logic/medicine"
	"chineseHealthy/apps/app/api/internal/svc"
	"chineseHealthy/apps/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func LookMedicineHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LookMedicineVagueRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := medicine.NewLookMedicineLogic(r.Context(), svcCtx)
		resp, err := l.LookMedicine(&req)
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
		response.HttpResult(r, w, resp, err)

	}
}
