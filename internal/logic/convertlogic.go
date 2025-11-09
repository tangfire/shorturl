package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"shorturl/internal/svc"
	"shorturl/internal/types"
	"shorturl/pkg/connect"
	"shorturl/pkg/md5"
	"shorturl/pkg/merr"
	"shorturl/pkg/urltool"
)

type ConvertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConvertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConvertLogic {
	return &ConvertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Convert 转链业务逻辑: 输入一个长链接 -> 转为短链接
func (l *ConvertLogic) Convert(req *types.ConvertRequest) (resp *types.ConvertResponse, err error) {
	// 1. 校验输入的数据
	// 1.1 数据不能为空
	// 1.2 输入的长链接必须是一个能请求通的网址
	if ok := connect.Get(req.LongUrl); !ok {
		return nil, errors.New("无效的链接")
	}
	// 1.3 判断之前是否已经转链过(数据库是否已存在该长链接)
	// 1.3.1 给长链接上成md5
	md5Value := md5.Sum([]byte(req.LongUrl))
	// 1.3.2 拿md5去数据库中查是否存在
	u, err := l.svcCtx.ShortUrlModel.FindOneByMd5(l.ctx, sql.NullString{String: md5Value, Valid: true})
	if !errors.Is(err, sqlx.ErrNotFound) {
		if err == nil {
			return nil, fmt.Errorf("该链接已经被转为:%s", u.Surl.String)
		}
		logx.Errorw("ShortUrlModel.FindOneByMd5 failed", merr.NormalErr(err))
		return nil, err
	}

	// 1.4 输入的不能是一个短链接(避免循环转链)
	// 输入的是一个完整的url
	basePath, err := urltool.GetBasePath(req.LongUrl)
	if err != nil {
		logx.Errorw("urltool.GetBasePath failed", logx.LogField{
			Key:   "lurl",
			Value: req.LongUrl,
		}, merr.NormalErr(err))
		return nil, err
	}
	_, err = l.svcCtx.ShortUrlModel.FindOneBySurl(l.ctx, sql.NullString{String: basePath, Valid: true})
	if !errors.Is(err, sqlx.ErrNotFound) {
		if err == nil {
			return nil, fmt.Errorf("该链接已经是短链接了")
		}
		logx.Errorw("ShortUrlModel.FindOneBySurl failed", merr.NormalErr(err))
		return nil, err
	}
	// 2. 取号 基于MySQL实现的发号器
	// 每来一个转链请求,我们就使用REPLACE
	seq, err := l.svcCtx.Sequence.Next()
	if err != nil {
		logx.Errorw("Sequence.Next failed", merr.NormalErr(err))
		return nil, err
	}
	fmt.Println(seq)

	// 3. 号码转短链

	// 4. 存储长短链接映射关系

	// 5. 返回响应

	return
}
