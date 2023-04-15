package website

import (
	"context"
	"errors"

	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/database/model"
	"github.com/star-horizon/anonymous-box-saas/database/repo"
	"github.com/star-horizon/anonymous-box-saas/internal/hashids"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/base"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/dash"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/dash/authservice"
)

var tracer = otel.Tracer(ServiceName)

const (
	ServiceName = "website-service"
)

type WebsiteServiceImpl struct {
	fx.In
	WebsiteRepo   repo.WebsiteRepo
	AuthSvcClient authservice.Client
	HashidsSvc    hashids.Service
}

func NewWebsiteService(impl WebsiteServiceImpl) dash.WebsiteService {
	return &impl
}

// CreateWebsite implements dash.WebsiteService.CreateWebsite
func (s *WebsiteServiceImpl) CreateWebsite(ctx context.Context, req *dash.CreateWebsiteRequest) (*dash.CreateWebsiteResponse, error) {
	ctx, span := tracer.Start(ctx, "create-website")
	defer span.End()

	userHash, err := s.HashidsSvc.Encode(ctx, req.GetUserId(), hashids.HashTypeUser)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("encode user id failed")
		return nil, err
	}

	website := &model.Website{
		UserRefer:             req.GetUserId(),
		Key:                   userHash,
		IsPublic:              req.GetIsPublic(),
		Name:                  req.GetName(),
		Description:           req.GetDescription(),
		AvatarIcon:            req.GetAvatarIcon(),
		Background:            req.GetBackground(),
		Language:              req.GetLanguage(),
		AllowAnonymousComment: req.GetAllowAnonymous(),
	}

	if err := s.WebsiteRepo.CreateByUserID(ctx, req.GetUserId(), website); err != nil {
		logrus.WithContext(ctx).WithError(err).Error("create website failed")
		return nil, err
	}

	return &dash.CreateWebsiteResponse{
		Key: userHash,
	}, nil
}

// GetWebsite implements dash.WebsiteService.GetWebsite
func (s *WebsiteServiceImpl) GetWebsite(ctx context.Context, req *dash.GetWebsiteRequest) (*dash.GetWebsiteResponse, error) {
	ctx, span := tracer.Start(ctx, "get-website")
	defer span.End()

	website, err := s.WebsiteRepo.GetByID(ctx, req.GetId())
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("get website failed")
		return nil, err
	}

	return &dash.GetWebsiteResponse{
		Id:             website.ID,
		Key:            website.Key,
		Name:           website.Name,
		Description:    website.Description,
		AvatarIcon:     website.AvatarIcon,
		Background:     website.Background,
		Language:       website.Language,
		IsPublic:       website.IsPublic,
		AllowAnonymous: website.AllowAnonymousComment,
	}, nil
}

// UpdateWebsite implements dash.WebsiteService.UpdateWebsite
func (s *WebsiteServiceImpl) UpdateWebsite(ctx context.Context, req *dash.UpdateWebsiteRequest) (*base.Empty, error) {
	ctx, span := tracer.Start(ctx, "update-website")
	defer span.End()

	website, err := s.WebsiteRepo.GetByID(ctx, req.GetId())
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("get website failed")
		return nil, err
	}

	if website.UserRefer != req.GetUserId() {
		logrus.WithContext(ctx).Error("user not match")
		return nil, errors.New("user not match")
	}

	website.Key = req.GetKey()
	website.Name = req.GetName()
	website.Description = req.GetDescription()
	website.AvatarIcon = req.GetAvatarIcon()
	website.Background = req.GetBackground()
	website.Language = req.GetLanguage()
	website.IsPublic = req.GetIsPublic()
	website.AllowAnonymousComment = req.GetAllowAnonymous()

	if err := s.WebsiteRepo.UpdateByID(ctx, website.ID, website); err != nil {
		logrus.WithContext(ctx).WithError(err).Error("update website failed")
		return nil, err
	}

	return &base.Empty{}, nil
}

func (s *WebsiteServiceImpl) ListWebsites(ctx context.Context, req *dash.ListWebsitesRequest) (*dash.ListWebsitesResponse, error) {
	ctx, span := tracer.Start(ctx, "list-websites")
	defer span.End()

	websites, count, err := s.WebsiteRepo.ListByUserID(
		ctx,
		req.GetUserId(),
		int((req.GetPagination().GetPage()-1)*req.GetPagination().GetPerPage()),
		int(req.GetPagination().GetPerPage()),
	)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("list websites failed")
		return nil, err
	}

	return &dash.ListWebsitesResponse{
		Total: count,
		Websites: lo.Map(websites, func(website *model.Website, index int) *dash.GetWebsiteResponse {
			return &dash.GetWebsiteResponse{
				Id:             website.ID,
				Key:            website.Key,
				Name:           website.Name,
				Description:    website.Description,
				AvatarIcon:     website.AvatarIcon,
				Background:     website.Background,
				Language:       website.Language,
				IsPublic:       website.IsPublic,
				AllowAnonymous: website.AllowAnonymousComment,
			}
		}),
	}, nil
}
