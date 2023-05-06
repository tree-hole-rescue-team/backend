package role

import (
	"net/http"

	"management-system/common/response"
	"management-system/service/app/api/internal/logic/role"
	"management-system/service/app/api/internal/svc"
	"management-system/service/app/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ChangeRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChangeRoleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewChangeRoleLogic(r.Context(), svcCtx)
		resp, err := l.ChangeRole(&req)
		response.SendResponse(w, r, resp, err)
	}
}
