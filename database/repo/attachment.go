package repo

import (
	"context"
	"fmt"

	"github.com/star-horizon/anonymous-box-saas/database/dal"
	"github.com/star-horizon/anonymous-box-saas/database/model"
	"github.com/star-horizon/anonymous-box-saas/pkg/cache"
)

type AttachmentRepo interface {
	GetByID(ctx context.Context, id uint64) (*model.Attachment, error)                // GetByID returns the attachment with the specified attachment ID.
	Create(ctx context.Context, storageId uint64, attachment *model.Attachment) error // Create creates a new attachment.
	DeleteByID(ctx context.Context, id uint64) error                                  // DeleteByID deletes the attachment with the specified attachment ID.
}

type attachmentRepo struct {
	Query *dal.Query
	Cache cache.Driver
}

func NewAttachmentRepo(repo attachmentRepo) AttachmentRepo {
	return &repo
}

// GetByID implements AttachmentRepo.GetByID.
func (r *attachmentRepo) GetByID(ctx context.Context, id uint64) (*model.Attachment, error) {
	ctx, span := tracer.Start(ctx, "get-attachment-by-id")
	defer span.End()

	if v, exist := r.Cache.Get(ctx, fmt.Sprint("database:attachment:id:", id)); exist {
		if attachment, ok := v.(model.Attachment); ok {
			return &attachment, nil
		}
	}

	attachment, err := r.Query.Attachment.WithContext(ctx).Where(r.Query.Attachment.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}

	_ = r.Cache.Set(ctx, fmt.Sprint("database:attachment:id:", id), *attachment, 0)

	return attachment, nil
}

// Create implements AttachmentRepo.Create.
func (r *attachmentRepo) Create(ctx context.Context, storageId uint64, attachment *model.Attachment) error {
	ctx, span := tracer.Start(ctx, "create-attachment")
	defer span.End()

	attachment.StorageID = storageId

	if err := r.Query.Attachment.WithContext(ctx).Create(attachment); err != nil {
		return err
	}

	_ = r.Cache.Set(ctx, fmt.Sprint("database:attachment:id:", attachment.ID), *attachment, 0)

	return nil
}

// DeleteByID implements AttachmentRepo.DeleteByID.
func (r *attachmentRepo) DeleteByID(ctx context.Context, id uint64) error {
	ctx, span := tracer.Start(ctx, "delete-attachment-by-id")
	defer span.End()

	if _, err := r.Query.Attachment.WithContext(ctx).Where(r.Query.Attachment.ID.Eq(id)).Delete(); err != nil {
		return err
	}

	_ = r.Cache.Delete(ctx, fmt.Sprint("database:attachment:id:", id))

	return nil
}
