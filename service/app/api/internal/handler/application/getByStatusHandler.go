package application

import (
	"net/http"

	"management-system/common/response"
	"management-system/service/app/api/internal/logic/application"
	"management-system/service/app/api/internal/svc"
	"management-system/service/app/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetByStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetByStatusReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := application.NewGetByStatusLogic(r.Context(), svcCtx)
		resp, err := l.GetByStatus(&req)
		response.SendResponse(w, r, resp, err)
	}
}
