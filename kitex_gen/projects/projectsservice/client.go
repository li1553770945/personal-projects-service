// Code generated by Kitex v0.7.2. DO NOT EDIT.

package projectsservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	projects "github.com/li1553770945/personal-projects-service/kitex_gen/projects"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetProjects(ctx context.Context, req *projects.ProjectsReq, callOptions ...callopt.Option) (r *projects.ProjectsResp, err error)
	GetProjectNum(ctx context.Context, callOptions ...callopt.Option) (r *projects.ProjectNumResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kProjectsServiceClient{
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

type kProjectsServiceClient struct {
	*kClient
}

func (p *kProjectsServiceClient) GetProjects(ctx context.Context, req *projects.ProjectsReq, callOptions ...callopt.Option) (r *projects.ProjectsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetProjects(ctx, req)
}

func (p *kProjectsServiceClient) GetProjectNum(ctx context.Context, callOptions ...callopt.Option) (r *projects.ProjectNumResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetProjectNum(ctx)
}
