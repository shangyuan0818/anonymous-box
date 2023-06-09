package repo

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/database/dal"
	"github.com/star-horizon/anonymous-box-saas/database/model"
	"github.com/star-horizon/anonymous-box-saas/pkg/cache"
)

type UserRepo interface {
	GetByID(ctx context.Context, id uint64) (*model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)

	Create(ctx context.Context, user *model.User) error
	UpdateByID(ctx context.Context, id uint64, user *model.User) error
	UpdatePassword(ctx context.Context, id uint64, password string) error
	Delete(ctx context.Context, id uint64) error
}

type userRepo struct {
	fx.In
	Query *dal.Query
	Cache cache.Driver
}

func NewUserRepo(repo userRepo) UserRepo {
	return &repo
}

// GetByID implements UserRepo.GetByID.
func (r *userRepo) GetByID(ctx context.Context, id uint64) (*model.User, error) {
	ctx, span := tracer.Start(ctx, "get-user-by-id")
	defer span.End()

	if v, exist := r.Cache.Get(ctx, fmt.Sprint("database:user:id:", id)); exist {
		if user, ok := v.(model.User); ok {
			span.AddEvent("get-from-cache", trace.WithAttributes(
				attribute.String("status", "hit"),
			))

			return &user, nil
		}
	} else {
		span.AddEvent("get-from-cache", trace.WithAttributes(
			attribute.String("status", "miss"),
		))
	}

	user, err := r.Query.User.WithContext(ctx).Where(r.Query.User.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}

	if err := r.Cache.Set(ctx, fmt.Sprint("database:user:id:", id), *user, 0); err != nil {
		span.RecordError(err)
	}

	return user, nil
}

// GetByUsername implements UserRepo.GetByUsername.
func (r *userRepo) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	ctx, span := tracer.Start(ctx, "get-user-by-username")
	defer span.End()

	user, err := r.Query.User.WithContext(ctx).Where(r.Query.User.Username.Eq(username)).First()
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetByEmail implements UserRepo.GetByEmail.
func (r *userRepo) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	ctx, span := tracer.Start(ctx, "get-user-by-email")
	defer span.End()

	user, err := r.Query.User.WithContext(ctx).Where(r.Query.User.Email.Eq(email)).First()
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Create implements UserRepo.Create.
func (r *userRepo) Create(ctx context.Context, user *model.User) error {
	ctx, span := tracer.Start(ctx, "create-user")
	defer span.End()

	if err := r.Query.User.WithContext(ctx).Create(user); err != nil {
		span.RecordError(err)
	}

	return nil
}

// UpdateByID implements UserRepo.UpdateByID.
func (r *userRepo) UpdateByID(ctx context.Context, id uint64, user *model.User) error {
	ctx, span := tracer.Start(ctx, "update-user")
	defer span.End()

	if err := r.Cache.Delete(ctx, fmt.Sprint("database:user:id:", id)); err != nil {
		return err
	}

	if _, err := r.Query.User.WithContext(ctx).
		Where(r.Query.User.ID.Eq(id)).
		Updates(user); err != nil {
		return err
	}

	return nil
}

// UpdatePassword implements UserRepo.UpdatePassword.
func (r *userRepo) UpdatePassword(ctx context.Context, id uint64, password string) error {
	ctx, span := tracer.Start(ctx, "update-user-password")
	defer span.End()

	if err := r.Cache.Delete(ctx, fmt.Sprint("database:user:id:", id)); err != nil {
		return err
	}

	if _, err := r.Query.User.WithContext(ctx).Where(r.Query.User.ID.Eq(id)).Update(r.Query.User.Password, password); err != nil {
		return err
	}

	return nil
}

// Delete implements UserRepo.Delete.
func (r *userRepo) Delete(ctx context.Context, id uint64) error {
	ctx, span := tracer.Start(ctx, "delete-user")
	defer span.End()

	if err := r.Cache.Delete(ctx, fmt.Sprint("database:user:id:", id)); err != nil {
		return err
	}

	if _, err := r.Query.User.WithContext(ctx).Where(r.Query.User.ID.Eq(id)).Delete(); err != nil {
		return err
	}

	return nil
}
