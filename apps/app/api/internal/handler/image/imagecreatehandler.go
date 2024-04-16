package image

import (
	"chineseHealthy/pkg/response"
	"net/http"

	"chineseHealthy/apps/app/api/internal/logic/image"
	"chineseHealthy/apps/app/api/internal/svc"
	"chineseHealthy/apps/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ImageCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImageCreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := image.NewImageCreateLogic(r.Context(), svcCtx)
		resp, err := l.ImageCreate(&req)
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
		response.HttpResult(r, w, resp, err)

	}
}
