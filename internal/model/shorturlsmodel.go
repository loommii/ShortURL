package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ShortUrlsModel = (*customShortUrlsModel)(nil)

type (
	// ShortUrlsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customShortUrlsModel.
	ShortUrlsModel interface {
		shortUrlsModel
		withSession(session sqlx.Session) ShortUrlsModel
	}

	customShortUrlsModel struct {
		*defaultShortUrlsModel
	}
)

// NewShortUrlsModel returns a model for the database table.
func NewShortUrlsModel(conn sqlx.SqlConn) ShortUrlsModel {
	return &customShortUrlsModel{
		defaultShortUrlsModel: newShortUrlsModel(conn),
	}
}

func (m *customShortUrlsModel) withSession(session sqlx.Session) ShortUrlsModel {
	return NewShortUrlsModel(sqlx.NewSqlConnFromSession(session))
}
