package logic

import (
	"context"

	"ShortURL/internal/model"
	"ShortURL/internal/svc"
	"ShortURL/internal/types"
	"ShortURL/internal/utils/dictionary"

	"github.com/zeromicro/go-zero/core/logx"
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
	// todo: add your logic here and delete this line
	randStr := dictionary.Dictionary.Get()
	data := model.ShortUrls{
		LongUrl:  req.LongUrl,
		ShortUrl: randStr,
	}
	_, err = l.svcCtx.Model.Insert(l.ctx, &data)
	if err != nil {
		return
	}
	resp = &types.RegisterResp{
		Host:     l.svcCtx.Config.ServerName,
		ShortKey: randStr,
	}
	resp.ShortUrl = resp.Host + "/s/" + resp.ShortKey
	return resp, err
}
