package logic

import (
	"context"
	"errors"
	"math/rand"

	"ShortURL/internal/model"
	"ShortURL/internal/svc"
	"ShortURL/internal/types"
	"ShortURL/internal/utils/dictionary"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	var shortKey string
	var shortUrls *model.ShortUrls
	for i := 0; i < 5; i++ { // 如果冲突就重复5次，5次冲突就告警和退出
		shortKey = dictionary.Dictionary.Get()
		// 查询是否冲突
		shortUrls, err = l.svcCtx.Model.FindOneByShortUrl(l.ctx, shortKey)
		if err != nil {
			if err == sqlx.ErrNotFound {
				break
			} else {
				return nil, err
			}
		}
	}
	if shortUrls != nil && shortUrls.Id != 0 {
		return nil, errors.New("请重试")
	}
	data := model.ShortUrls{
		LongUrl:  req.LongUrl,
		ShortUrl: shortKey,
	}
	_, err = l.svcCtx.Model.Insert(l.ctx, &data)
	if err != nil {
		return
	}
	resp = &types.RegisterResp{
		Host:     l.svcCtx.Config.ServerName,
		ShortKey: shortKey,
	}
	resp.ShortUrl = resp.Host + "/s/" + resp.ShortKey
	// 异步添加一个缓存
	go func(key, longUrl string) {
		l.svcCtx.Redis.Setex(key, longUrl, 60*120+rand.Intn(60*120))
	}(l.svcCtx.Config.Name+"_"+resp.ShortKey, req.LongUrl)
	return resp, err
}
