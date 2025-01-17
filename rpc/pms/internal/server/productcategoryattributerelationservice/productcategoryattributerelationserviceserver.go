// Code generated by goctl. DO NOT EDIT.
// Source: pms.proto

package server

import (
	"context"

	"zero-admin/rpc/pms/internal/logic/productcategoryattributerelationservice"
	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pmsclient"
)

type ProductCategoryAttributeRelationServiceServer struct {
	svcCtx *svc.ServiceContext
	pmsclient.UnimplementedProductCategoryAttributeRelationServiceServer
}

func NewProductCategoryAttributeRelationServiceServer(svcCtx *svc.ServiceContext) *ProductCategoryAttributeRelationServiceServer {
	return &ProductCategoryAttributeRelationServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *ProductCategoryAttributeRelationServiceServer) ProductCategoryAttributeRelationAdd(ctx context.Context, in *pmsclient.ProductCategoryAttributeRelationAddReq) (*pmsclient.ProductCategoryAttributeRelationAddResp, error) {
	l := productcategoryattributerelationservicelogic.NewProductCategoryAttributeRelationAddLogic(ctx, s.svcCtx)
	return l.ProductCategoryAttributeRelationAdd(in)
}

func (s *ProductCategoryAttributeRelationServiceServer) ProductCategoryAttributeRelationList(ctx context.Context, in *pmsclient.ProductCategoryAttributeRelationListReq) (*pmsclient.ProductCategoryAttributeRelationListResp, error) {
	l := productcategoryattributerelationservicelogic.NewProductCategoryAttributeRelationListLogic(ctx, s.svcCtx)
	return l.ProductCategoryAttributeRelationList(in)
}

func (s *ProductCategoryAttributeRelationServiceServer) ProductCategoryAttributeRelationUpdate(ctx context.Context, in *pmsclient.ProductCategoryAttributeRelationUpdateReq) (*pmsclient.ProductCategoryAttributeRelationUpdateResp, error) {
	l := productcategoryattributerelationservicelogic.NewProductCategoryAttributeRelationUpdateLogic(ctx, s.svcCtx)
	return l.ProductCategoryAttributeRelationUpdate(in)
}

func (s *ProductCategoryAttributeRelationServiceServer) ProductCategoryAttributeRelationDelete(ctx context.Context, in *pmsclient.ProductCategoryAttributeRelationDeleteReq) (*pmsclient.ProductCategoryAttributeRelationDeleteResp, error) {
	l := productcategoryattributerelationservicelogic.NewProductCategoryAttributeRelationDeleteLogic(ctx, s.svcCtx)
	return l.ProductCategoryAttributeRelationDelete(in)
}
