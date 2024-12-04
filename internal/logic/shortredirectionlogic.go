package logic

import (
	"context"
	"math/rand"

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
	// 先查询一下缓存
	longUrl, err := l.svcCtx.Redis.Get(l.svcCtx.Config.Name + "_" + req.ShortKey)
	if err != nil {
		l.Error(err)
	}
	if longUrl == "" {
		shortUrls, err := l.svcCtx.Model.FindOneByShortUrl(l.ctx, req.ShortKey)
		if err != nil {
			return nil, err
		}
		longUrl = shortUrls.LongUrl
		// 异步添加一个缓存
		go func(key, longUrl string) {
			l.svcCtx.Redis.Setex(key, longUrl, 60*120+rand.Intn(60*120))
		}(l.svcCtx.Config.Name+"_"+shortUrls.ShortUrl, longUrl)
	}
	resp = &types.ShortRedirectionResp{
		LongUrl: longUrl,
	}
	return
}
