package repo

import (
	"context"

	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/database/dal"
	"github.com/star-horizon/anonymous-box-saas/database/model"
	"github.com/star-horizon/anonymous-box-saas/pkg/cache"
)

type CommentRepo interface {
	GetByID(ctx context.Context, id uint64) (*model.Comment, error)                                            // GetByID returns the comment with the specified comment ID.
	ListByWebsiteID(ctx context.Context, websiteID uint64, offset, limit int) ([]*model.Comment, int64, error) // ListByWebsiteID returns the list of comments with the specified website ID.
	CreateByWebsiteID(ctx context.Context, websiteID uint64, comment *model.Comment) error                     // CreateByWebsiteID creates a new comment with the specified website ID.
	UpdateContentByID(ctx context.Context, id uint64, content string) error                                    // UpdateContentByID updates the content of the comment with the specified comment ID.
	DeleteByIDAndUser(ctx context.Context, id, userId uint64) error                                            // DeleteByID deletes the comment with the specified comment ID.
}

type commentRepo struct {
	fx.In
	Query *dal.Query
	Cache cache.Driver
}

func NewCommentRepo(repo commentRepo) CommentRepo {
	return &repo
}

// GetByID implements CommentRepo.GetByID.
func (r commentRepo) GetByID(ctx context.Context, id uint64) (*model.Comment, error) {
	return r.Query.Comment.
		WithContext(ctx).
		Where(r.Query.Comment.ID.Eq(id)).
		First()
}

// ListByWebsiteID implements CommentRepo.ListByWebsiteID.
func (r commentRepo) ListByWebsiteID(ctx context.Context, websiteID uint64, offset, limit int) ([]*model.Comment, int64, error) {
	ctx, span := tracer.Start(ctx, "list-comments-by-website-id")
	defer span.End()

	comments, count, err := r.Query.Comment.
		WithContext(ctx).
		Where(r.Query.Comment.WebsiteRefer.Eq(websiteID)).
		FindByPage(offset, limit)
	if err != nil {
		span.RecordError(err)
		return nil, 0, err
	}

	return comments, count, nil
}

// CreateByWebsiteID implements CommentRepo.CreateByWebsiteID.
func (r commentRepo) CreateByWebsiteID(ctx context.Context, websiteID uint64, comment *model.Comment) error {
	ctx, span := tracer.Start(ctx, "create-comment-by-website-id")
	defer span.End()

	comment.WebsiteRefer = websiteID

	if err := r.Query.Comment.
		WithContext(ctx).
		Create(comment); err != nil {
		span.RecordError(err)
		return err
	}

	return nil
}

// UpdateContentByID implements CommentRepo.UpdateContentByID.
func (r commentRepo) UpdateContentByID(ctx context.Context, id uint64, content string) error {
	ctx, span := tracer.Start(ctx, "update-comment-content-by-id")
	defer span.End()

	if _, err := r.Query.Comment.
		WithContext(ctx).
		Where(r.Query.Comment.ID.Eq(id)).
		Update(r.Query.Comment.Content, content); err != nil {
		span.RecordError(err)
		return err
	}

	return nil
}

// DeleteByIDAndUser implements CommentRepo.DeleteByIDAndUser.
func (r commentRepo) DeleteByIDAndUser(ctx context.Context, id, userId uint64) error {
	ctx, span := tracer.Start(ctx, "delete-comment-by-id")
	defer span.End()

	if _, err := r.Query.Comment.
		WithContext(ctx).
		Where(r.Query.Comment.ID.Eq(id)).
		Having(
			r.Query.WithContext(ctx).Comment.
				Columns(r.Query.Comment.WebsiteRefer).
				In(
					r.Query.WithContext(ctx).Website.Where(
						r.Query.Website.UserRefer.Eq(userId),
					).Select(r.Query.Website.ID),
				),
		).
		Delete(); err != nil {
		span.RecordError(err)
		return err
	}

	return nil
}
