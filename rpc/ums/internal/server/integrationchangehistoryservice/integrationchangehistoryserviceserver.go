// Code generated by goctl. DO NOT EDIT.
// Source: ums.proto

package server

import (
	"context"

	"zero-admin/rpc/ums/internal/logic/integrationchangehistoryservice"
	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"
)

type IntegrationChangeHistoryServiceServer struct {
	svcCtx *svc.ServiceContext
	umsclient.UnimplementedIntegrationChangeHistoryServiceServer
}

func NewIntegrationChangeHistoryServiceServer(svcCtx *svc.ServiceContext) *IntegrationChangeHistoryServiceServer {
	return &IntegrationChangeHistoryServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *IntegrationChangeHistoryServiceServer) IntegrationChangeHistoryAdd(ctx context.Context, in *umsclient.IntegrationChangeHistoryAddReq) (*umsclient.IntegrationChangeHistoryAddResp, error) {
	l := integrationchangehistoryservicelogic.NewIntegrationChangeHistoryAddLogic(ctx, s.svcCtx)
	return l.IntegrationChangeHistoryAdd(in)
}

func (s *IntegrationChangeHistoryServiceServer) IntegrationChangeHistoryList(ctx context.Context, in *umsclient.IntegrationChangeHistoryListReq) (*umsclient.IntegrationChangeHistoryListResp, error) {
	l := integrationchangehistoryservicelogic.NewIntegrationChangeHistoryListLogic(ctx, s.svcCtx)
	return l.IntegrationChangeHistoryList(in)
}

func (s *IntegrationChangeHistoryServiceServer) IntegrationChangeHistoryUpdate(ctx context.Context, in *umsclient.IntegrationChangeHistoryUpdateReq) (*umsclient.IntegrationChangeHistoryUpdateResp, error) {
	l := integrationchangehistoryservicelogic.NewIntegrationChangeHistoryUpdateLogic(ctx, s.svcCtx)
	return l.IntegrationChangeHistoryUpdate(in)
}

func (s *IntegrationChangeHistoryServiceServer) IntegrationChangeHistoryDelete(ctx context.Context, in *umsclient.IntegrationChangeHistoryDeleteReq) (*umsclient.IntegrationChangeHistoryDeleteResp, error) {
	l := integrationchangehistoryservicelogic.NewIntegrationChangeHistoryDeleteLogic(ctx, s.svcCtx)
	return l.IntegrationChangeHistoryDelete(in)
}