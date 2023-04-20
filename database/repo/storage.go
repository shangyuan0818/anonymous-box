package repo

import (
	"context"
	"fmt"

	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/database/dal"
	"github.com/star-horizon/anonymous-box-saas/database/model"
	"github.com/star-horizon/anonymous-box-saas/pkg/cache"
)

type StorageRepo interface {
	GetByID(ctx context.Context, id uint64) (*model.Storage, error)                                                    // GetByID returns the storage with the specified storage ID.
	ListByInUse(ctx context.Context, inUse bool, offset, limit int) ([]*model.Storage, int64, error)                   // ListByUsed returns the list of storages with the specified used.
	ListBySizeAvailable(ctx context.Context, size int64, offset, limit int) ([]*model.Storage, int64, error)           // ListBySizeAvailable returns the list of storages with the specified size available.
	ListByType(ctx context.Context, storageType model.StorageType, offset, limit int) ([]*model.Storage, int64, error) // ListByType returns the list of storages with the specified storage type.
	List(ctx context.Context, limit, offset int) ([]*model.Storage, int64, error)                                      // List returns the list of storages.
	Create(ctx context.Context, storage *model.Storage) error                                                          // Create creates a new storage.
	AddSizeUsedByID(ctx context.Context, id uint64, size int64) error                                                  // AddSizeUsedByID adds the specified size to the storage with the specified storage ID.
	UpdateByID(ctx context.Context, id uint64, storage *model.Storage) error                                           // UpdateByID updates the storage with the specified storage ID.
	DeleteByID(ctx context.Context, id uint64) error                                                                   // DeleteByID deletes the storage with the specified storage ID.
}

type storageRepo struct {
	fx.In
	Query *dal.Query
	Cache cache.Driver
}

func NewStorageRepo(repo storageRepo) StorageRepo {
	return &repo
}

// GetByID implements StorageRepo.GetByID.
func (r *storageRepo) GetByID(ctx context.Context, id uint64) (*model.Storage, error) {
	ctx, span := tracer.Start(ctx, "get-storage-by-id")
	defer span.End()

	cacheKey := fmt.Sprint("database:storage:id:", id)
	if v, exist := r.Cache.Get(ctx, cacheKey); exist {
		if storage, ok := v.(model.Storage); ok {
			return &storage, nil
		}
	}

	storage, err := r.Query.Storage.
		WithContext(ctx).
		Where(r.Query.Storage.ID.Eq(id)).
		First()
	if err != nil {
		return nil, err
	}

	_ = r.Cache.Set(ctx, cacheKey, *storage, 0)

	return storage, nil
}

// ListByInUse implements StorageRepo.ListByInUse.
func (r *storageRepo) ListByInUse(ctx context.Context, inUse bool, offset, limit int) ([]*model.Storage, int64, error) {
	ctx, span := tracer.Start(ctx, "list-storage-by-used")
	defer span.End()

	storages, count, err := r.Query.Storage.
		WithContext(ctx).
		Where(r.Query.Storage.IsInUse.Is(inUse)).
		FindByPage(offset, limit)
	if err != nil {
		return nil, 0, err
	}

	return storages, count, nil
}

// ListBySizeAvailable implements StorageRepo.ListBySizeAvailable.
func (r *storageRepo) ListBySizeAvailable(ctx context.Context, size int64, offset, limit int) ([]*model.Storage, int64, error) {
	ctx, span := tracer.Start(ctx, "list-storage-by-size-available")
	defer span.End()

	storages, count, err := r.Query.Storage.
		WithContext(ctx).
		Where(
			r.Query.Storage.IsInUse.Is(true),
			r.Query.Storage.MaxSize.GteCol(r.Query.Storage.UsedSize.Add(size)),
		).
		FindByPage(offset, limit)
	if err != nil {
		return nil, 0, err
	}

	return storages, count, nil
}

// ListByType implements StorageRepo.ListByType.
func (r *storageRepo) ListByType(ctx context.Context, storageType model.StorageType, offset, limit int) ([]*model.Storage, int64, error) {
	ctx, span := tracer.Start(ctx, "list-storage-by-type")
	defer span.End()

	storages, count, err := r.Query.Storage.
		WithContext(ctx).
		Where(r.Query.Storage.Type.Eq(string(storageType))).
		FindByPage(offset, limit)
	if err != nil {
		return nil, 0, err
	}

	return storages, count, nil
}

// List implements StorageRepo.List.
func (r *storageRepo) List(ctx context.Context, limit, offset int) ([]*model.Storage, int64, error) {
	ctx, span := tracer.Start(ctx, "list-storage")
	defer span.End()

	storages, count, err := r.Query.Storage.
		WithContext(ctx).
		FindByPage(offset, limit)
	if err != nil {
		return nil, 0, err
	}

	return storages, count, nil
}

// Create implements StorageRepo.Create.
func (r *storageRepo) Create(ctx context.Context, storage *model.Storage) error {
	ctx, span := tracer.Start(ctx, "create-storage")
	defer span.End()

	if err := r.Query.Storage.WithContext(ctx).Create(storage); err != nil {
		return err
	}

	return nil
}

// AddSizeUsedByID implements StorageRepo.AddSizeUsedByID.
func (r *storageRepo) AddSizeUsedByID(ctx context.Context, id uint64, size int64) error {
	ctx, span := tracer.Start(ctx, "add-size-used-by-id")
	defer span.End()

	if err := r.Cache.Delete(ctx, fmt.Sprint("database:storage:id:", id)); err != nil {
		return err
	}

	if _, err := r.Query.Storage.WithContext(ctx).
		Where(r.Query.Storage.ID.Eq(id)).
		Update(
			r.Query.Storage.UsedSize,
			r.Query.Storage.UsedSize.Add(size),
		); err != nil {
		return err
	}

	return nil
}

// UpdateByID implements StorageRepo.UpdateByID.
func (r *storageRepo) UpdateByID(ctx context.Context, id uint64, storage *model.Storage) error {
	ctx, span := tracer.Start(ctx, "update-storage-by-id")
	defer span.End()

	if err := r.Cache.Delete(ctx, fmt.Sprint("database:storage:id:", id)); err != nil {
		return err
	}

	if _, err := r.Query.Storage.WithContext(ctx).
		Where(r.Query.Storage.ID.Eq(id)).
		Updates(storage); err != nil {
		return err
	}

	return nil
}

// DeleteByID implements StorageRepo.DeleteByID.
func (r *storageRepo) DeleteByID(ctx context.Context, id uint64) error {
	ctx, span := tracer.Start(ctx, "delete-storage-by-id")
	defer span.End()

	if err := r.Cache.Delete(ctx, fmt.Sprint("database:storage:id:", id)); err != nil {
		return err
	}

	if _, err := r.Query.Storage.WithContext(ctx).
		Where(r.Query.Storage.ID.Eq(id)).
		Delete(); err != nil {
		return err
	}

	return nil
}
