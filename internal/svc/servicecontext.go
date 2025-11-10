package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"shorturl/internal/config"
	"shorturl/model"
	"shorturl/sequence"
)

type ServiceContext struct {
	Config        config.Config
	ShortUrlModel model.ShortUrlMapModel // short_url_map
	Sequence      sequence.Sequence
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.ShortUrlDB.DSN)
	return &ServiceContext{
		Config:        c,
		ShortUrlModel: model.NewShortUrlMapModel(conn),
		Sequence:      sequence.NewMySQL(c.Sequence.DSN), // sequence
	}
}
