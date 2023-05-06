package user

import (
	"net/http"

	"management-system/common/response"
	"management-system/service/app/api/internal/logic/user"
	"management-system/service/app/api/internal/svc"
	"management-system/service/app/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadAvatarHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadAvatarReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUploadAvatarLogic(r.Context(), svcCtx,r)
		resp, err := l.UploadAvatar(&req)
		response.SendResponse(w, r, resp, err)
	}
}
