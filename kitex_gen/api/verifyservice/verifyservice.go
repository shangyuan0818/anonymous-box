// Code generated by Kitex v0.5.1. DO NOT EDIT.

package verifyservice

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	api "github.com/star-horizon/anonymous-box-saas/kitex_gen/api"
	proto "google.golang.org/protobuf/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func serviceInfo() *kitex.ServiceInfo {
	return verifyServiceServiceInfo
}

var verifyServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "VerifyService"
	handlerType := (*api.VerifyService)(nil)
	methods := map[string]kitex.MethodInfo{
		"ApplyEmailVerify": kitex.NewMethodInfo(applyEmailVerifyHandler, newApplyEmailVerifyArgs, newApplyEmailVerifyResult, false),
		"VerifyEmail":      kitex.NewMethodInfo(verifyEmailHandler, newVerifyEmailArgs, newVerifyEmailResult, false),
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

func applyEmailVerifyHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.ApplyEmailVerifyRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.VerifyService).ApplyEmailVerify(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *ApplyEmailVerifyArgs:
		success, err := handler.(api.VerifyService).ApplyEmailVerify(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ApplyEmailVerifyResult)
		realResult.Success = success
	}
	return nil
}
func newApplyEmailVerifyArgs() interface{} {
	return &ApplyEmailVerifyArgs{}
}

func newApplyEmailVerifyResult() interface{} {
	return &ApplyEmailVerifyResult{}
}

type ApplyEmailVerifyArgs struct {
	Req *api.ApplyEmailVerifyRequest
}

func (p *ApplyEmailVerifyArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.ApplyEmailVerifyRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ApplyEmailVerifyArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ApplyEmailVerifyArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ApplyEmailVerifyArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ApplyEmailVerifyArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *ApplyEmailVerifyArgs) Unmarshal(in []byte) error {
	msg := new(api.ApplyEmailVerifyRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ApplyEmailVerifyArgs_Req_DEFAULT *api.ApplyEmailVerifyRequest

func (p *ApplyEmailVerifyArgs) GetReq() *api.ApplyEmailVerifyRequest {
	if !p.IsSetReq() {
		return ApplyEmailVerifyArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ApplyEmailVerifyArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ApplyEmailVerifyArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ApplyEmailVerifyResult struct {
	Success *emptypb.Empty
}

var ApplyEmailVerifyResult_Success_DEFAULT *emptypb.Empty

func (p *ApplyEmailVerifyResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(emptypb.Empty)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ApplyEmailVerifyResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ApplyEmailVerifyResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ApplyEmailVerifyResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ApplyEmailVerifyResult")
	}
	return proto.Marshal(p.Success)
}

func (p *ApplyEmailVerifyResult) Unmarshal(in []byte) error {
	msg := new(emptypb.Empty)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ApplyEmailVerifyResult) GetSuccess() *emptypb.Empty {
	if !p.IsSetSuccess() {
		return ApplyEmailVerifyResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ApplyEmailVerifyResult) SetSuccess(x interface{}) {
	p.Success = x.(*emptypb.Empty)
}

func (p *ApplyEmailVerifyResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ApplyEmailVerifyResult) GetResult() interface{} {
	return p.Success
}

func verifyEmailHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.VerifyEmailRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.VerifyService).VerifyEmail(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *VerifyEmailArgs:
		success, err := handler.(api.VerifyService).VerifyEmail(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*VerifyEmailResult)
		realResult.Success = success
	}
	return nil
}
func newVerifyEmailArgs() interface{} {
	return &VerifyEmailArgs{}
}

func newVerifyEmailResult() interface{} {
	return &VerifyEmailResult{}
}

type VerifyEmailArgs struct {
	Req *api.VerifyEmailRequest
}

func (p *VerifyEmailArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.VerifyEmailRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *VerifyEmailArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *VerifyEmailArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *VerifyEmailArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in VerifyEmailArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *VerifyEmailArgs) Unmarshal(in []byte) error {
	msg := new(api.VerifyEmailRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var VerifyEmailArgs_Req_DEFAULT *api.VerifyEmailRequest

func (p *VerifyEmailArgs) GetReq() *api.VerifyEmailRequest {
	if !p.IsSetReq() {
		return VerifyEmailArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *VerifyEmailArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *VerifyEmailArgs) GetFirstArgument() interface{} {
	return p.Req
}

type VerifyEmailResult struct {
	Success *emptypb.Empty
}

var VerifyEmailResult_Success_DEFAULT *emptypb.Empty

func (p *VerifyEmailResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(emptypb.Empty)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *VerifyEmailResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *VerifyEmailResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *VerifyEmailResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in VerifyEmailResult")
	}
	return proto.Marshal(p.Success)
}

func (p *VerifyEmailResult) Unmarshal(in []byte) error {
	msg := new(emptypb.Empty)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *VerifyEmailResult) GetSuccess() *emptypb.Empty {
	if !p.IsSetSuccess() {
		return VerifyEmailResult_Success_DEFAULT
	}
	return p.Success
}

func (p *VerifyEmailResult) SetSuccess(x interface{}) {
	p.Success = x.(*emptypb.Empty)
}

func (p *VerifyEmailResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *VerifyEmailResult) GetResult() interface{} {
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

func (p *kClient) ApplyEmailVerify(ctx context.Context, Req *api.ApplyEmailVerifyRequest) (r *emptypb.Empty, err error) {
	var _args ApplyEmailVerifyArgs
	_args.Req = Req
	var _result ApplyEmailVerifyResult
	if err = p.c.Call(ctx, "ApplyEmailVerify", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) VerifyEmail(ctx context.Context, Req *api.VerifyEmailRequest) (r *emptypb.Empty, err error) {
	var _args VerifyEmailArgs
	_args.Req = Req
	var _result VerifyEmailResult
	if err = p.c.Call(ctx, "VerifyEmail", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
