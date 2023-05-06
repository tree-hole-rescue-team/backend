package user

import (
	"net/http"

	"management-system/common/response"
	"management-system/service/app/api/internal/logic/user"
	"management-system/service/app/api/internal/svc"
	"management-system/service/app/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ImportExcelHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImportExcelReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewImportExcelLogic(r.Context(), svcCtx,r)
		resp, err := l.ImportExcel(&req)
		response.SendResponse(w, r, resp, err)
	}
}
