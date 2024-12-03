package handler

import (
	"net/http"

	"ShortURL/internal/logic"
	"ShortURL/internal/svc"
	"ShortURL/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func shortRedirectionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShortRedirectionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewShortRedirectionLogic(r.Context(), svcCtx)
		resp, err := l.ShortRedirection(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			http.Redirect(w, r, resp.LongUrl, http.StatusFound)
			// httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
