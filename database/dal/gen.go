// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:         db,
		Attachment: newAttachment(db, opts...),
		Comment:    newComment(db, opts...),
		Setting:    newSetting(db, opts...),
		Storage:    newStorage(db, opts...),
		User:       newUser(db, opts...),
		Website:    newWebsite(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Attachment attachment
	Comment    comment
	Setting    setting
	Storage    storage
	User       user
	Website    website
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:         db,
		Attachment: q.Attachment.clone(db),
		Comment:    q.Comment.clone(db),
		Setting:    q.Setting.clone(db),
		Storage:    q.Storage.clone(db),
		User:       q.User.clone(db),
		Website:    q.Website.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:         db,
		Attachment: q.Attachment.replaceDB(db),
		Comment:    q.Comment.replaceDB(db),
		Setting:    q.Setting.replaceDB(db),
		Storage:    q.Storage.replaceDB(db),
		User:       q.User.replaceDB(db),
		Website:    q.Website.replaceDB(db),
	}
}

type queryCtx struct {
	Attachment IAttachmentDo
	Comment    ICommentDo
	Setting    ISettingDo
	Storage    IStorageDo
	User       IUserDo
	Website    IWebsiteDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Attachment: q.Attachment.WithContext(ctx),
		Comment:    q.Comment.WithContext(ctx),
		Setting:    q.Setting.WithContext(ctx),
		Storage:    q.Storage.WithContext(ctx),
		User:       q.User.WithContext(ctx),
		Website:    q.Website.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
