package logic

import (
	"context"

	"shorturl/internal/svc"
	"shorturl/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

	// 1.3 判断之前是否已经转链过(数据库是否已存在该长链接)

	// 1.4 输入的不能是一个短链接(避免循环转链)

	// 2. 取号

	// 3. 号码转短链

	// 4. 存储长短链接映射关系

	// 5. 返回响应

	return
}
