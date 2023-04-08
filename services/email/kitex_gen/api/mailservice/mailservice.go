// Code generated by Kitex v0.5.1. DO NOT EDIT.

package mailservice

import (
	"context"
	"fmt"
	api "github.com/ahdark-services/anonymous-box-saas/services/email/kitex_gen/api"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return mailServiceServiceInfo
}

var mailServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "MailService"
	handlerType := (*api.MailService)(nil)
	methods := map[string]kitex.MethodInfo{
		"SendMail": kitex.NewMethodInfo(sendMailHandler, newSendMailArgs, newSendMailResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "api",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.5.1",
		Extra:           extra,
	}
	return svcInfo
}

func sendMailHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.SendMailRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.MailService).SendMail(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *SendMailArgs:
		success, err := handler.(api.MailService).SendMail(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*SendMailResult)
		realResult.Success = success
	}
	return nil
}
func newSendMailArgs() interface{} {
	return &SendMailArgs{}
}

func newSendMailResult() interface{} {
	return &SendMailResult{}
}

type SendMailArgs struct {
	Req *api.SendMailRequest
}

func (p *SendMailArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.SendMailRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *SendMailArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *SendMailArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *SendMailArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SendMailArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *SendMailArgs) Unmarshal(in []byte) error {
	msg := new(api.SendMailRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SendMailArgs_Req_DEFAULT *api.SendMailRequest

func (p *SendMailArgs) GetReq() *api.SendMailRequest {
	if !p.IsSetReq() {
		return SendMailArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SendMailArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *SendMailArgs) GetFirstArgument() interface{} {
	return p.Req
}

type SendMailResult struct {
	Success *api.SendMailResponse
}

var SendMailResult_Success_DEFAULT *api.SendMailResponse

func (p *SendMailResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.SendMailResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *SendMailResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *SendMailResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *SendMailResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SendMailResult")
	}
	return proto.Marshal(p.Success)
}

func (p *SendMailResult) Unmarshal(in []byte) error {
	msg := new(api.SendMailResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SendMailResult) GetSuccess() *api.SendMailResponse {
	if !p.IsSetSuccess() {
		return SendMailResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SendMailResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.SendMailResponse)
}

func (p *SendMailResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SendMailResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) SendMail(ctx context.Context, Req *api.SendMailRequest) (r *api.SendMailResponse, err error) {
	var _args SendMailArgs
	_args.Req = Req
	var _result SendMailResult
	if err = p.c.Call(ctx, "SendMail", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
