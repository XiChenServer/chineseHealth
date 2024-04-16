package image

import (
	"chineseHealthy/pkg/response"
	"net/http"

	"chineseHealthy/apps/app/api/internal/logic/image"
	"chineseHealthy/apps/app/api/internal/svc"
)

func ImageUpdataHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := image.NewImageUpdataLogic(r.Context(), svcCtx)
		resp, err := l.ImageUpdata()
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
		response.HttpResult(r, w, resp, err)

	}
}
