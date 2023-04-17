// Code generated by Kitex v0.5.1. DO NOT EDIT.
package boxcommentservice

import (
	server "github.com/cloudwego/kitex/server"
	box "github.com/star-horizon/anonymous-box-saas/kitex_gen/box"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler box.BoxCommentService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}