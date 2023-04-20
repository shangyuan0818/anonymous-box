package dash_comment

import (
	"context"

	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/database/model"
	"github.com/star-horizon/anonymous-box-saas/database/repo"
	"github.com/star-horizon/anonymous-box-saas/internal/hashids"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/base"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/dash"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/dash/websiteservice"
)

var tracer = otel.Tracer(ServiceName)

const (
	ServiceName = "dash-comment-service"
)

type CommentServiceImpl struct {
	fx.In
	CommentRepo      repo.CommentRepo
	WebsiteSvcClient websiteservice.Client
	HashidsSvc       hashids.Service
}

func NewCommentService(impl CommentServiceImpl) dash.CommentService {
	return &impl
}

// GetComment implements dash.CommentService
func (s *CommentServiceImpl) GetComment(ctx context.Context, req *dash.GetCommentRequest) (*dash.Comment, error) {
	ctx, span := tracer.Start(ctx, "get-comment")
	defer span.End()

	comment, err := s.CommentRepo.GetByID(ctx, req.GetId())
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("get comment failed")
		return nil, err
	}

	return &dash.Comment{
		Id:           req.GetId(),
		WebsiteRefer: comment.WebsiteRefer,
		Name:         lo.If(comment.Name.Valid, comment.Name.String).Else(""),
		Email:        lo.If(comment.Email.Valid, comment.Email.String).Else(""),
		Url:          lo.If(comment.Url.Valid, comment.Url.String).Else(""),
		Content:      comment.Content,
		CreatedAt: &base.Timestamp{
			Seconds: comment.CreatedAt.Unix(),
			Nanos:   int32(comment.CreatedAt.Nanosecond()),
		},
		UpdatedAt: &base.Timestamp{
			Seconds: comment.UpdatedAt.Unix(),
			Nanos:   int32(comment.UpdatedAt.Nanosecond()),
		},
	}, nil
}

// ListComments implements dash.CommentService
func (s *CommentServiceImpl) ListComments(ctx context.Context, req *dash.ListCommentsRequest) (*dash.ListCommentsResponse, error) {
	ctx, span := tracer.Start(ctx, "list-comments")
	defer span.End()

	comments, count, err := s.CommentRepo.ListByWebsiteID(
		ctx,
		req.GetWebsiteRefer(),
		int(req.GetPagination().GetPerPage()),
		int((req.GetPagination().GetPage()-1)*req.GetPagination().GetPerPage()),
	)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("list comments failed")
		return nil, err
	}

	return &dash.ListCommentsResponse{
		Total: count,
		Comments: lo.Map(comments, func(comment *model.Comment, _ int) *dash.Comment {
			return &dash.Comment{
				Id:           comment.ID,
				WebsiteRefer: comment.WebsiteRefer,
				Name:         lo.If(comment.Name.Valid, comment.Name.String).Else(""),
				Email:        lo.If(comment.Email.Valid, comment.Email.String).Else(""),
				Url:          lo.If(comment.Url.Valid, comment.Url.String).Else(""),
				Content:      comment.Content,
				CreatedAt: &base.Timestamp{
					Seconds: comment.CreatedAt.Unix(),
					Nanos:   int32(comment.CreatedAt.Nanosecond()),
				},
				UpdatedAt: &base.Timestamp{
					Seconds: comment.UpdatedAt.Unix(),
					Nanos:   int32(comment.UpdatedAt.Nanosecond()),
				},
			}
		}),
	}, nil
}

// DeleteComment implements dash.CommentService
func (s *CommentServiceImpl) DeleteComment(ctx context.Context, req *dash.DeleteCommentRequest) (*base.Empty, error) {
	ctx, span := tracer.Start(ctx, "delete-comment")
	defer span.End()

	if err := s.CommentRepo.DeleteByIDAndUser(ctx, req.GetId(), req.GetUserId()); err != nil {
		logrus.WithContext(ctx).WithError(err).Error("delete comment failed")
		return nil, err
	}

	return &base.Empty{}, nil
}
