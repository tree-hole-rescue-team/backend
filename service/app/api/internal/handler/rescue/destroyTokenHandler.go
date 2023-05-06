package rescue

import (
	"net/http"

	"management-system/common/response"
	"management-system/service/app/api/internal/logic/rescue"
	"management-system/service/app/api/internal/svc"
	"management-system/service/app/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DestroyTokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DestroyTokenReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := rescue.NewDestroyTokenLogic(r.Context(), svcCtx)
		resp, err := l.DestroyToken(&req)
		response.SendResponse(w, r, resp, err)
	}
}
