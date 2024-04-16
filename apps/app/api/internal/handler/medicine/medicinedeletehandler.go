package medicine

import (
	"chineseHealthy/pkg/response"
	"net/http"

	"chineseHealthy/apps/app/api/internal/logic/medicine"
	"chineseHealthy/apps/app/api/internal/svc"
	"chineseHealthy/apps/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func MedicineDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MedicineDelRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := medicine.NewMedicineDeleteLogic(r.Context(), svcCtx)
		resp, err := l.MedicineDelete(&req)
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
		response.HttpResult(r, w, resp, err)

	}
}
