package user

import (
	"net/http"

	"management-system/common/response"
	"management-system/service/app/api/internal/logic/user"
	"management-system/service/app/api/internal/svc"
	"management-system/service/app/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func EditUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EditUserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewEditUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.EditUserInfo(&req)
		response.SendResponse(w, r, resp, err)
	}
}
