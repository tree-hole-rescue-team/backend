package rescue

import (
	"net/http"

	"management-system/common/response"
	"management-system/service/app/api/internal/logic/rescue"
	"management-system/service/app/api/internal/svc"
	"management-system/service/app/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func PushRescueInfoWithStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PushWithStatusReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := rescue.NewPushRescueInfoWithStatusLogic(r.Context(), svcCtx)
		resp, err := l.PushRescueInfoWithStatus(&req)
		response.SendResponse(w, r, resp, err)
	}
}
