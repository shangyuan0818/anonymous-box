// Code generated by Kitex v0.5.1. DO NOT EDIT.

package attachmentservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	transport "github.com/cloudwego/kitex/transport"
	api "github.com/star-horizon/anonymous-box-saas/kitex_gen/api"
	base "github.com/star-horizon/anonymous-box-saas/kitex_gen/base"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	RequestAttachmentUpload(ctx context.Context, Req *api.RequestAttachmentUploadRequest, callOptions ...callopt.Option) (r *api.RequestAttachmentUploadResponse, err error)
	UploadAttachment(ctx context.Context, callOptions ...callopt.Option) (stream AttachmentService_UploadAttachmentClient, err error)
	GetAttachment(ctx context.Context, Req *api.GetAttachmentRequest, callOptions ...callopt.Option) (r *api.GetAttachmentResponse, err error)
	DeleteAttachment(ctx context.Context, Req *api.DeleteAttachmentRequest, callOptions ...callopt.Option) (r *base.Empty, err error)
}

type AttachmentService_UploadAttachmentClient interface {
	streaming.Stream
	Send(*api.UploadAttachmentRequest) error
	CloseAndRecv() (*api.UploadAttachmentResponse, error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, client.WithTransportProtocol(transport.GRPC))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kAttachmentServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kAttachmentServiceClient struct {
	*kClient
}

func (p *kAttachmentServiceClient) RequestAttachmentUpload(ctx context.Context, Req *api.RequestAttachmentUploadRequest, callOptions ...callopt.Option) (r *api.RequestAttachmentUploadResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RequestAttachmentUpload(ctx, Req)
}

func (p *kAttachmentServiceClient) UploadAttachment(ctx context.Context, callOptions ...callopt.Option) (stream AttachmentService_UploadAttachmentClient, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UploadAttachment(ctx)
}

func (p *kAttachmentServiceClient) GetAttachment(ctx context.Context, Req *api.GetAttachmentRequest, callOptions ...callopt.Option) (r *api.GetAttachmentResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetAttachment(ctx, Req)
}

func (p *kAttachmentServiceClient) DeleteAttachment(ctx context.Context, Req *api.DeleteAttachmentRequest, callOptions ...callopt.Option) (r *base.Empty, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteAttachment(ctx, Req)
}
