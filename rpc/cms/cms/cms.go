// Code generated by goctl. DO NOT EDIT.
// Source: cms.proto

package cms

import (
	"context"

	"zero-admin/rpc/cms/cmsclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	SubjectAddReq       = cmsclient.SubjectAddReq
	SubjectAddResp      = cmsclient.SubjectAddResp
	SubjectDeleteReq    = cmsclient.SubjectDeleteReq
	SubjectDeleteResp   = cmsclient.SubjectDeleteResp
	SubjectListByIdsReq = cmsclient.SubjectListByIdsReq
	SubjectListData     = cmsclient.SubjectListData
	SubjectListReq      = cmsclient.SubjectListReq
	SubjectListResp     = cmsclient.SubjectListResp
	SubjectUpdateReq    = cmsclient.SubjectUpdateReq
	SubjectUpdateResp   = cmsclient.SubjectUpdateResp

	Cms interface {
		SubjectAdd(ctx context.Context, in *SubjectAddReq, opts ...grpc.CallOption) (*SubjectAddResp, error)
		SubjectDelete(ctx context.Context, in *SubjectDeleteReq, opts ...grpc.CallOption) (*SubjectDeleteResp, error)
		SubjectUpdate(ctx context.Context, in *SubjectUpdateReq, opts ...grpc.CallOption) (*SubjectUpdateResp, error)
		SubjectList(ctx context.Context, in *SubjectListReq, opts ...grpc.CallOption) (*SubjectListResp, error)
		SubjectListByIds(ctx context.Context, in *SubjectListByIdsReq, opts ...grpc.CallOption) (*SubjectListResp, error)
	}

	defaultCms struct {
		cli zrpc.Client
	}
)

func NewCms(cli zrpc.Client) Cms {
	return &defaultCms{
		cli: cli,
	}
}

func (m *defaultCms) SubjectAdd(ctx context.Context, in *SubjectAddReq, opts ...grpc.CallOption) (*SubjectAddResp, error) {
	client := cmsclient.NewCmsClient(m.cli.Conn())
	return client.SubjectAdd(ctx, in, opts...)
}

func (m *defaultCms) SubjectDelete(ctx context.Context, in *SubjectDeleteReq, opts ...grpc.CallOption) (*SubjectDeleteResp, error) {
	client := cmsclient.NewCmsClient(m.cli.Conn())
	return client.SubjectDelete(ctx, in, opts...)
}

func (m *defaultCms) SubjectUpdate(ctx context.Context, in *SubjectUpdateReq, opts ...grpc.CallOption) (*SubjectUpdateResp, error) {
	client := cmsclient.NewCmsClient(m.cli.Conn())
	return client.SubjectUpdate(ctx, in, opts...)
}

func (m *defaultCms) SubjectList(ctx context.Context, in *SubjectListReq, opts ...grpc.CallOption) (*SubjectListResp, error) {
	client := cmsclient.NewCmsClient(m.cli.Conn())
	return client.SubjectList(ctx, in, opts...)
}

func (m *defaultCms) SubjectListByIds(ctx context.Context, in *SubjectListByIdsReq, opts ...grpc.CallOption) (*SubjectListResp, error) {
	client := cmsclient.NewCmsClient(m.cli.Conn())
	return client.SubjectListByIds(ctx, in, opts...)
}
