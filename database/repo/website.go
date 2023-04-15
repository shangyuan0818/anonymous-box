package repo

import (
	"context"
	"errors"
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/database/dal"
	"github.com/star-horizon/anonymous-box-saas/database/model"
	"github.com/star-horizon/anonymous-box-saas/pkg/cache"
)

type WebsiteRepo interface {
	GetByID(ctx context.Context, id uint64) (*model.Website, error)
	GetByKey(ctx context.Context, key string) (*model.Website, error)
	GetByUserID(ctx context.Context, userID uint64) ([]*model.Website, error)
	ListByUserID(ctx context.Context, userID uint64, offset, limit int) ([]*model.Website, int64, error)
	CreateByUserID(ctx context.Context, userID uint64, website *model.Website) error
	UpdateByID(ctx context.Context, id uint64, website *model.Website) error
	DeleteByID(ctx context.Context, id uint64) error
}

type websiteRepo struct {
	fx.In
	Query *dal.Query
	Cache cache.Driver
}

func NewWebsiteRepo(repo websiteRepo) WebsiteRepo {
	return &repo
}

// GetByID implements WebsiteRepo.GetByID.
func (r *websiteRepo) GetByID(ctx context.Context, id uint64) (*model.Website, error) {
	ctx, span := tracer.Start(ctx, "get-website-by-id")
	defer span.End()

	if v, exist := r.Cache.Get(ctx, fmt.Sprint("database:website:id:", id)); exist {
		if website, ok := v.(model.Website); ok {
			span.AddEvent("get-from-cache", trace.WithAttributes(
				attribute.String("status", "hit"),
			))

			return &website, nil
		}
	} else {
		span.AddEvent("get-from-cache", trace.WithAttributes(
			attribute.String("status", "miss"),
		))
	}

	website, err := r.Query.Website.
		WithContext(ctx).
		Where(r.Query.Website.ID.Eq(id)).
		First()
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	return website, nil
}

// GetByKey implements WebsiteRepo.GetByKey.
func (r *websiteRepo) GetByKey(ctx context.Context, key string) (*model.Website, error) {
	ctx, span := tracer.Start(ctx, "get-website-by-key")
	defer span.End()

	if v, exist := r.Cache.Get(ctx, fmt.Sprint("database:website:key:", key)); exist {
		if website, ok := v.(model.Website); ok {
			span.AddEvent("get-from-cache", trace.WithAttributes(
				attribute.String("status", "hit"),
			))

			return &website, nil
		}
	} else {
		span.AddEvent("get-from-cache", trace.WithAttributes(
			attribute.String("status", "miss"),
		))
	}

	website, err := r.Query.Website.
		WithContext(ctx).
		Where(r.Query.Website.Key.Eq(key)).
		First()
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	return website, nil
}

// GetByUserID implements WebsiteRepo.GetByUserID.
func (r *websiteRepo) GetByUserID(ctx context.Context, userID uint64) ([]*model.Website, error) {
	ctx, span := tracer.Start(ctx, "get-website-by-user-id")
	defer span.End()

	websites, err := r.Query.Website.
		WithContext(ctx).
		Where(r.Query.Website.UserRefer.Eq(userID)).
		Find()
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	return websites, nil
}

// ListByUserID implements WebsiteRepo.ListByUserID.
func (r *websiteRepo) ListByUserID(ctx context.Context, userID uint64, offset, limit int) ([]*model.Website, int64, error) {
	ctx, span := tracer.Start(ctx, "list-website-by-user-id")
	defer span.End()

	websites, count, err := r.Query.Website.
		WithContext(ctx).
		Where(r.Query.Website.UserRefer.Eq(userID)).
		FindByPage(offset, limit)
	if err != nil {
		span.RecordError(err)
		return nil, 0, err
	}

	return websites, count, nil
}

// CreateByUserID implements WebsiteRepo.CreateByUserID.
func (r *websiteRepo) CreateByUserID(ctx context.Context, userID uint64, website *model.Website) error {
	ctx, span := tracer.Start(ctx, "create-website-by-user-id")
	defer span.End()

	if website == nil {
		return errors.New("website is nil")
	}

	website.UserRefer = userID

	if err := r.Query.Website.
		WithContext(ctx).
		Create(website); err != nil {
		span.RecordError(err)
		return err
	}

	return nil
}

// UpdateByID implements WebsiteRepo.UpdateByID.
func (r *websiteRepo) UpdateByID(ctx context.Context, id uint64, website *model.Website) error {
	ctx, span := tracer.Start(ctx, "update-website-by-id")
	defer span.End()

	if website == nil {
		return errors.New("website is nil")
	}

	website.ID = id

	if _, err := r.Query.Website.
		WithContext(ctx).
		Updates(website); err != nil {
		span.RecordError(err)
		return err
	}

	return nil
}

// DeleteByID implements WebsiteRepo.DeleteByID.
func (r *websiteRepo) DeleteByID(ctx context.Context, id uint64) error {
	ctx, span := tracer.Start(ctx, "delete-website-by-id")
	defer span.End()

	if _, err := r.Query.Website.
		WithContext(ctx).
		Where(r.Query.Website.ID.Eq(id)).
		Delete(); err != nil {
		span.RecordError(err)
		return err
	}

	return nil
}
