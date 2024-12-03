package logic

import (
	"context"

	"ShortURL/internal/svc"
	"ShortURL/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShortRedirectionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShortRedirectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortRedirectionLogic {
	return &ShortRedirectionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShortRedirectionLogic) ShortRedirection(req *types.ShortRedirectionReq) (resp *types.ShortRedirectionResp, err error) {
	// todo: add your logic here and delete this line

	shortUrls, err := l.svcCtx.Model.FindOneByShortUrl(l.ctx, req.ShortKey)
	if err != nil {
		return nil, err
	}
	resp = &types.ShortRedirectionResp{
		LongUrl: shortUrls.LongUrl,
	}
	return
}
