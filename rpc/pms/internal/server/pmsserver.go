// Code generated by goctl. DO NOT EDIT.
// Source: pms.proto

package server

import (
	"context"

	"zero-admin/rpc/pms/internal/logic"
	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pmsclient"
)

type PmsServer struct {
	svcCtx *svc.ServiceContext
	pmsclient.UnimplementedPmsServer
}

func NewPmsServer(svcCtx *svc.ServiceContext) *PmsServer {
	return &PmsServer{
		svcCtx: svcCtx,
	}
}

func (s *PmsServer) ProductAdd(ctx context.Context, in *pmsclient.ProductAddReq) (*pmsclient.ProductAddResp, error) {
	l := logic.NewProductAddLogic(ctx, s.svcCtx)
	return l.ProductAdd(in)
}

func (s *PmsServer) ProductList(ctx context.Context, in *pmsclient.ProductListReq) (*pmsclient.ProductListResp, error) {
	l := logic.NewProductListLogic(ctx, s.svcCtx)
	return l.ProductList(in)
}

func (s *PmsServer) ProductListByIds(ctx context.Context, in *pmsclient.ProductByIdsReq) (*pmsclient.ProductListResp, error) {
	l := logic.NewProductListByIdsLogic(ctx, s.svcCtx)
	return l.ProductListByIds(in)
}

func (s *PmsServer) ProductUpdate(ctx context.Context, in *pmsclient.ProductUpdateReq) (*pmsclient.ProductUpdateResp, error) {
	l := logic.NewProductUpdateLogic(ctx, s.svcCtx)
	return l.ProductUpdate(in)
}

func (s *PmsServer) ProductDelete(ctx context.Context, in *pmsclient.ProductDeleteReq) (*pmsclient.ProductDeleteResp, error) {
	l := logic.NewProductDeleteLogic(ctx, s.svcCtx)
	return l.ProductDelete(in)
}

func (s *PmsServer) ProductDetailById(ctx context.Context, in *pmsclient.ProductDetailByIdReq) (*pmsclient.ProductDetailByIdResp, error) {
	l := logic.NewProductDetailByIdLogic(ctx, s.svcCtx)
	return l.ProductDetailById(in)
}

func (s *PmsServer) AlbumAdd(ctx context.Context, in *pmsclient.AlbumAddReq) (*pmsclient.AlbumAddResp, error) {
	l := logic.NewAlbumAddLogic(ctx, s.svcCtx)
	return l.AlbumAdd(in)
}

func (s *PmsServer) AlbumList(ctx context.Context, in *pmsclient.AlbumListReq) (*pmsclient.AlbumListResp, error) {
	l := logic.NewAlbumListLogic(ctx, s.svcCtx)
	return l.AlbumList(in)
}

func (s *PmsServer) AlbumUpdate(ctx context.Context, in *pmsclient.AlbumUpdateReq) (*pmsclient.AlbumUpdateResp, error) {
	l := logic.NewAlbumUpdateLogic(ctx, s.svcCtx)
	return l.AlbumUpdate(in)
}

func (s *PmsServer) AlbumDelete(ctx context.Context, in *pmsclient.AlbumDeleteReq) (*pmsclient.AlbumDeleteResp, error) {
	l := logic.NewAlbumDeleteLogic(ctx, s.svcCtx)
	return l.AlbumDelete(in)
}

func (s *PmsServer) AlbumPicAdd(ctx context.Context, in *pmsclient.AlbumPicAddReq) (*pmsclient.AlbumPicAddResp, error) {
	l := logic.NewAlbumPicAddLogic(ctx, s.svcCtx)
	return l.AlbumPicAdd(in)
}

func (s *PmsServer) AlbumPicList(ctx context.Context, in *pmsclient.AlbumPicListReq) (*pmsclient.AlbumPicListResp, error) {
	l := logic.NewAlbumPicListLogic(ctx, s.svcCtx)
	return l.AlbumPicList(in)
}

func (s *PmsServer) AlbumPicUpdate(ctx context.Context, in *pmsclient.AlbumPicUpdateReq) (*pmsclient.AlbumPicUpdateResp, error) {
	l := logic.NewAlbumPicUpdateLogic(ctx, s.svcCtx)
	return l.AlbumPicUpdate(in)
}

func (s *PmsServer) AlbumPicDelete(ctx context.Context, in *pmsclient.AlbumPicDeleteReq) (*pmsclient.AlbumPicDeleteResp, error) {
	l := logic.NewAlbumPicDeleteLogic(ctx, s.svcCtx)
	return l.AlbumPicDelete(in)
}

func (s *PmsServer) BrandAdd(ctx context.Context, in *pmsclient.BrandAddReq) (*pmsclient.BrandAddResp, error) {
	l := logic.NewBrandAddLogic(ctx, s.svcCtx)
	return l.BrandAdd(in)
}

func (s *PmsServer) BrandList(ctx context.Context, in *pmsclient.BrandListReq) (*pmsclient.BrandListResp, error) {
	l := logic.NewBrandListLogic(ctx, s.svcCtx)
	return l.BrandList(in)
}

func (s *PmsServer) BrandListByIds(ctx context.Context, in *pmsclient.BrandListByIdsReq) (*pmsclient.BrandListResp, error) {
	l := logic.NewBrandListByIdsLogic(ctx, s.svcCtx)
	return l.BrandListByIds(in)
}

func (s *PmsServer) BrandUpdate(ctx context.Context, in *pmsclient.BrandUpdateReq) (*pmsclient.BrandUpdateResp, error) {
	l := logic.NewBrandUpdateLogic(ctx, s.svcCtx)
	return l.BrandUpdate(in)
}

func (s *PmsServer) BrandDelete(ctx context.Context, in *pmsclient.BrandDeleteReq) (*pmsclient.BrandDeleteResp, error) {
	l := logic.NewBrandDeleteLogic(ctx, s.svcCtx)
	return l.BrandDelete(in)
}

func (s *PmsServer) CommentAdd(ctx context.Context, in *pmsclient.CommentAddReq) (*pmsclient.CommentAddResp, error) {
	l := logic.NewCommentAddLogic(ctx, s.svcCtx)
	return l.CommentAdd(in)
}

func (s *PmsServer) CommentList(ctx context.Context, in *pmsclient.CommentListReq) (*pmsclient.CommentListResp, error) {
	l := logic.NewCommentListLogic(ctx, s.svcCtx)
	return l.CommentList(in)
}

func (s *PmsServer) CommentUpdate(ctx context.Context, in *pmsclient.CommentUpdateReq) (*pmsclient.CommentUpdateResp, error) {
	l := logic.NewCommentUpdateLogic(ctx, s.svcCtx)
	return l.CommentUpdate(in)
}

func (s *PmsServer) CommentDelete(ctx context.Context, in *pmsclient.CommentDeleteReq) (*pmsclient.CommentDeleteResp, error) {
	l := logic.NewCommentDeleteLogic(ctx, s.svcCtx)
	return l.CommentDelete(in)
}

func (s *PmsServer) CommentReplayAdd(ctx context.Context, in *pmsclient.CommentReplayAddReq) (*pmsclient.CommentReplayAddResp, error) {
	l := logic.NewCommentReplayAddLogic(ctx, s.svcCtx)
	return l.CommentReplayAdd(in)
}

func (s *PmsServer) CommentReplayList(ctx context.Context, in *pmsclient.CommentReplayListReq) (*pmsclient.CommentReplayListResp, error) {
	l := logic.NewCommentReplayListLogic(ctx, s.svcCtx)
	return l.CommentReplayList(in)
}

func (s *PmsServer) CommentReplayUpdate(ctx context.Context, in *pmsclient.CommentReplayUpdateReq) (*pmsclient.CommentReplayUpdateResp, error) {
	l := logic.NewCommentReplayUpdateLogic(ctx, s.svcCtx)
	return l.CommentReplayUpdate(in)
}

func (s *PmsServer) CommentReplayDelete(ctx context.Context, in *pmsclient.CommentReplayDeleteReq) (*pmsclient.CommentReplayDeleteResp, error) {
	l := logic.NewCommentReplayDeleteLogic(ctx, s.svcCtx)
	return l.CommentReplayDelete(in)
}

func (s *PmsServer) FeightTemplateAdd(ctx context.Context, in *pmsclient.FeightTemplateAddReq) (*pmsclient.FeightTemplateAddResp, error) {
	l := logic.NewFeightTemplateAddLogic(ctx, s.svcCtx)
	return l.FeightTemplateAdd(in)
}

func (s *PmsServer) FeightTemplateList(ctx context.Context, in *pmsclient.FeightTemplateListReq) (*pmsclient.FeightTemplateListResp, error) {
	l := logic.NewFeightTemplateListLogic(ctx, s.svcCtx)
	return l.FeightTemplateList(in)
}

func (s *PmsServer) FeightTemplateUpdate(ctx context.Context, in *pmsclient.FeightTemplateUpdateReq) (*pmsclient.FeightTemplateUpdateResp, error) {
	l := logic.NewFeightTemplateUpdateLogic(ctx, s.svcCtx)
	return l.FeightTemplateUpdate(in)
}

func (s *PmsServer) FeightTemplateDelete(ctx context.Context, in *pmsclient.FeightTemplateDeleteReq) (*pmsclient.FeightTemplateDeleteResp, error) {
	l := logic.NewFeightTemplateDeleteLogic(ctx, s.svcCtx)
	return l.FeightTemplateDelete(in)
}

func (s *PmsServer) MemberPriceAdd(ctx context.Context, in *pmsclient.MemberPriceAddReq) (*pmsclient.MemberPriceAddResp, error) {
	l := logic.NewMemberPriceAddLogic(ctx, s.svcCtx)
	return l.MemberPriceAdd(in)
}

func (s *PmsServer) MemberPriceList(ctx context.Context, in *pmsclient.MemberPriceListReq) (*pmsclient.MemberPriceListResp, error) {
	l := logic.NewMemberPriceListLogic(ctx, s.svcCtx)
	return l.MemberPriceList(in)
}

func (s *PmsServer) MemberPriceUpdate(ctx context.Context, in *pmsclient.MemberPriceUpdateReq) (*pmsclient.MemberPriceUpdateResp, error) {
	l := logic.NewMemberPriceUpdateLogic(ctx, s.svcCtx)
	return l.MemberPriceUpdate(in)
}

func (s *PmsServer) MemberPriceDelete(ctx context.Context, in *pmsclient.MemberPriceDeleteReq) (*pmsclient.MemberPriceDeleteResp, error) {
	l := logic.NewMemberPriceDeleteLogic(ctx, s.svcCtx)
	return l.MemberPriceDelete(in)
}

func (s *PmsServer) ProductAttributeCategoryAdd(ctx context.Context, in *pmsclient.ProductAttributeCategoryAddReq) (*pmsclient.ProductAttributeCategoryAddResp, error) {
	l := logic.NewProductAttributeCategoryAddLogic(ctx, s.svcCtx)
	return l.ProductAttributeCategoryAdd(in)
}

func (s *PmsServer) ProductAttributeCategoryList(ctx context.Context, in *pmsclient.ProductAttributeCategoryListReq) (*pmsclient.ProductAttributeCategoryListResp, error) {
	l := logic.NewProductAttributeCategoryListLogic(ctx, s.svcCtx)
	return l.ProductAttributeCategoryList(in)
}

func (s *PmsServer) ProductAttributeCategoryUpdate(ctx context.Context, in *pmsclient.ProductAttributeCategoryUpdateReq) (*pmsclient.ProductAttributeCategoryUpdateResp, error) {
	l := logic.NewProductAttributeCategoryUpdateLogic(ctx, s.svcCtx)
	return l.ProductAttributeCategoryUpdate(in)
}

func (s *PmsServer) ProductAttributeCategoryDelete(ctx context.Context, in *pmsclient.ProductAttributeCategoryDeleteReq) (*pmsclient.ProductAttributeCategoryDeleteResp, error) {
	l := logic.NewProductAttributeCategoryDeleteLogic(ctx, s.svcCtx)
	return l.ProductAttributeCategoryDelete(in)
}

func (s *PmsServer) ProductAttributeAdd(ctx context.Context, in *pmsclient.ProductAttributeAddReq) (*pmsclient.ProductAttributeAddResp, error) {
	l := logic.NewProductAttributeAddLogic(ctx, s.svcCtx)
	return l.ProductAttributeAdd(in)
}

func (s *PmsServer) ProductAttributeList(ctx context.Context, in *pmsclient.ProductAttributeListReq) (*pmsclient.ProductAttributeListResp, error) {
	l := logic.NewProductAttributeListLogic(ctx, s.svcCtx)
	return l.ProductAttributeList(in)
}

func (s *PmsServer) ProductAttributeUpdate(ctx context.Context, in *pmsclient.ProductAttributeUpdateReq) (*pmsclient.ProductAttributeUpdateResp, error) {
	l := logic.NewProductAttributeUpdateLogic(ctx, s.svcCtx)
	return l.ProductAttributeUpdate(in)
}

func (s *PmsServer) ProductAttributeDelete(ctx context.Context, in *pmsclient.ProductAttributeDeleteReq) (*pmsclient.ProductAttributeDeleteResp, error) {
	l := logic.NewProductAttributeDeleteLogic(ctx, s.svcCtx)
	return l.ProductAttributeDelete(in)
}

func (s *PmsServer) ProductAttributeValueAdd(ctx context.Context, in *pmsclient.ProductAttributeValueAddReq) (*pmsclient.ProductAttributeValueAddResp, error) {
	l := logic.NewProductAttributeValueAddLogic(ctx, s.svcCtx)
	return l.ProductAttributeValueAdd(in)
}

func (s *PmsServer) ProductAttributeValueList(ctx context.Context, in *pmsclient.ProductAttributeValueListReq) (*pmsclient.ProductAttributeValueListResp, error) {
	l := logic.NewProductAttributeValueListLogic(ctx, s.svcCtx)
	return l.ProductAttributeValueList(in)
}

func (s *PmsServer) ProductAttributeValueUpdate(ctx context.Context, in *pmsclient.ProductAttributeValueUpdateReq) (*pmsclient.ProductAttributeValueUpdateResp, error) {
	l := logic.NewProductAttributeValueUpdateLogic(ctx, s.svcCtx)
	return l.ProductAttributeValueUpdate(in)
}

func (s *PmsServer) ProductAttributeValueDelete(ctx context.Context, in *pmsclient.ProductAttributeValueDeleteReq) (*pmsclient.ProductAttributeValueDeleteResp, error) {
	l := logic.NewProductAttributeValueDeleteLogic(ctx, s.svcCtx)
	return l.ProductAttributeValueDelete(in)
}

func (s *PmsServer) ProductCategoryAttributeRelationAdd(ctx context.Context, in *pmsclient.ProductCategoryAttributeRelationAddReq) (*pmsclient.ProductCategoryAttributeRelationAddResp, error) {
	l := logic.NewProductCategoryAttributeRelationAddLogic(ctx, s.svcCtx)
	return l.ProductCategoryAttributeRelationAdd(in)
}

func (s *PmsServer) ProductCategoryAttributeRelationList(ctx context.Context, in *pmsclient.ProductCategoryAttributeRelationListReq) (*pmsclient.ProductCategoryAttributeRelationListResp, error) {
	l := logic.NewProductCategoryAttributeRelationListLogic(ctx, s.svcCtx)
	return l.ProductCategoryAttributeRelationList(in)
}

func (s *PmsServer) ProductCategoryAttributeRelationUpdate(ctx context.Context, in *pmsclient.ProductCategoryAttributeRelationUpdateReq) (*pmsclient.ProductCategoryAttributeRelationUpdateResp, error) {
	l := logic.NewProductCategoryAttributeRelationUpdateLogic(ctx, s.svcCtx)
	return l.ProductCategoryAttributeRelationUpdate(in)
}

func (s *PmsServer) ProductCategoryAttributeRelationDelete(ctx context.Context, in *pmsclient.ProductCategoryAttributeRelationDeleteReq) (*pmsclient.ProductCategoryAttributeRelationDeleteResp, error) {
	l := logic.NewProductCategoryAttributeRelationDeleteLogic(ctx, s.svcCtx)
	return l.ProductCategoryAttributeRelationDelete(in)
}

func (s *PmsServer) ProductCategoryAdd(ctx context.Context, in *pmsclient.ProductCategoryAddReq) (*pmsclient.ProductCategoryAddResp, error) {
	l := logic.NewProductCategoryAddLogic(ctx, s.svcCtx)
	return l.ProductCategoryAdd(in)
}

func (s *PmsServer) ProductCategoryList(ctx context.Context, in *pmsclient.ProductCategoryListReq) (*pmsclient.ProductCategoryListResp, error) {
	l := logic.NewProductCategoryListLogic(ctx, s.svcCtx)
	return l.ProductCategoryList(in)
}

func (s *PmsServer) ProductCategoryUpdate(ctx context.Context, in *pmsclient.ProductCategoryUpdateReq) (*pmsclient.ProductCategoryUpdateResp, error) {
	l := logic.NewProductCategoryUpdateLogic(ctx, s.svcCtx)
	return l.ProductCategoryUpdate(in)
}

func (s *PmsServer) ProductCategoryDelete(ctx context.Context, in *pmsclient.ProductCategoryDeleteReq) (*pmsclient.ProductCategoryDeleteResp, error) {
	l := logic.NewProductCategoryDeleteLogic(ctx, s.svcCtx)
	return l.ProductCategoryDelete(in)
}

func (s *PmsServer) ProductFullReductionAdd(ctx context.Context, in *pmsclient.ProductFullReductionAddReq) (*pmsclient.ProductFullReductionAddResp, error) {
	l := logic.NewProductFullReductionAddLogic(ctx, s.svcCtx)
	return l.ProductFullReductionAdd(in)
}

func (s *PmsServer) ProductFullReductionList(ctx context.Context, in *pmsclient.ProductFullReductionListReq) (*pmsclient.ProductFullReductionListResp, error) {
	l := logic.NewProductFullReductionListLogic(ctx, s.svcCtx)
	return l.ProductFullReductionList(in)
}

func (s *PmsServer) ProductFullReductionUpdate(ctx context.Context, in *pmsclient.ProductFullReductionUpdateReq) (*pmsclient.ProductFullReductionUpdateResp, error) {
	l := logic.NewProductFullReductionUpdateLogic(ctx, s.svcCtx)
	return l.ProductFullReductionUpdate(in)
}

func (s *PmsServer) ProductFullReductionDelete(ctx context.Context, in *pmsclient.ProductFullReductionDeleteReq) (*pmsclient.ProductFullReductionDeleteResp, error) {
	l := logic.NewProductFullReductionDeleteLogic(ctx, s.svcCtx)
	return l.ProductFullReductionDelete(in)
}

func (s *PmsServer) ProductLadderAdd(ctx context.Context, in *pmsclient.ProductLadderAddReq) (*pmsclient.ProductLadderAddResp, error) {
	l := logic.NewProductLadderAddLogic(ctx, s.svcCtx)
	return l.ProductLadderAdd(in)
}

func (s *PmsServer) ProductLadderList(ctx context.Context, in *pmsclient.ProductLadderListReq) (*pmsclient.ProductLadderListResp, error) {
	l := logic.NewProductLadderListLogic(ctx, s.svcCtx)
	return l.ProductLadderList(in)
}

func (s *PmsServer) ProductLadderUpdate(ctx context.Context, in *pmsclient.ProductLadderUpdateReq) (*pmsclient.ProductLadderUpdateResp, error) {
	l := logic.NewProductLadderUpdateLogic(ctx, s.svcCtx)
	return l.ProductLadderUpdate(in)
}

func (s *PmsServer) ProductLadderDelete(ctx context.Context, in *pmsclient.ProductLadderDeleteReq) (*pmsclient.ProductLadderDeleteResp, error) {
	l := logic.NewProductLadderDeleteLogic(ctx, s.svcCtx)
	return l.ProductLadderDelete(in)
}

func (s *PmsServer) ProductOperateLogAdd(ctx context.Context, in *pmsclient.ProductOperateLogAddReq) (*pmsclient.ProductOperateLogAddResp, error) {
	l := logic.NewProductOperateLogAddLogic(ctx, s.svcCtx)
	return l.ProductOperateLogAdd(in)
}

func (s *PmsServer) ProductOperateLogList(ctx context.Context, in *pmsclient.ProductOperateLogListReq) (*pmsclient.ProductOperateLogListResp, error) {
	l := logic.NewProductOperateLogListLogic(ctx, s.svcCtx)
	return l.ProductOperateLogList(in)
}

func (s *PmsServer) ProductOperateLogUpdate(ctx context.Context, in *pmsclient.ProductOperateLogUpdateReq) (*pmsclient.ProductOperateLogUpdateResp, error) {
	l := logic.NewProductOperateLogUpdateLogic(ctx, s.svcCtx)
	return l.ProductOperateLogUpdate(in)
}

func (s *PmsServer) ProductOperateLogDelete(ctx context.Context, in *pmsclient.ProductOperateLogDeleteReq) (*pmsclient.ProductOperateLogDeleteResp, error) {
	l := logic.NewProductOperateLogDeleteLogic(ctx, s.svcCtx)
	return l.ProductOperateLogDelete(in)
}

func (s *PmsServer) ProductVertifyRecordAdd(ctx context.Context, in *pmsclient.ProductVertifyRecordAddReq) (*pmsclient.ProductVertifyRecordAddResp, error) {
	l := logic.NewProductVertifyRecordAddLogic(ctx, s.svcCtx)
	return l.ProductVertifyRecordAdd(in)
}

func (s *PmsServer) ProductVertifyRecordList(ctx context.Context, in *pmsclient.ProductVertifyRecordListReq) (*pmsclient.ProductVertifyRecordListResp, error) {
	l := logic.NewProductVertifyRecordListLogic(ctx, s.svcCtx)
	return l.ProductVertifyRecordList(in)
}

func (s *PmsServer) ProductVertifyRecordUpdate(ctx context.Context, in *pmsclient.ProductVertifyRecordUpdateReq) (*pmsclient.ProductVertifyRecordUpdateResp, error) {
	l := logic.NewProductVertifyRecordUpdateLogic(ctx, s.svcCtx)
	return l.ProductVertifyRecordUpdate(in)
}

func (s *PmsServer) ProductVertifyRecordDelete(ctx context.Context, in *pmsclient.ProductVertifyRecordDeleteReq) (*pmsclient.ProductVertifyRecordDeleteResp, error) {
	l := logic.NewProductVertifyRecordDeleteLogic(ctx, s.svcCtx)
	return l.ProductVertifyRecordDelete(in)
}

func (s *PmsServer) SkuStockAdd(ctx context.Context, in *pmsclient.SkuStockAddReq) (*pmsclient.SkuStockAddResp, error) {
	l := logic.NewSkuStockAddLogic(ctx, s.svcCtx)
	return l.SkuStockAdd(in)
}

func (s *PmsServer) SkuStockList(ctx context.Context, in *pmsclient.SkuStockListReq) (*pmsclient.SkuStockListResp, error) {
	l := logic.NewSkuStockListLogic(ctx, s.svcCtx)
	return l.SkuStockList(in)
}

func (s *PmsServer) SkuStockUpdate(ctx context.Context, in *pmsclient.SkuStockUpdateReq) (*pmsclient.SkuStockUpdateResp, error) {
	l := logic.NewSkuStockUpdateLogic(ctx, s.svcCtx)
	return l.SkuStockUpdate(in)
}

func (s *PmsServer) SkuStockDelete(ctx context.Context, in *pmsclient.SkuStockDeleteReq) (*pmsclient.SkuStockDeleteResp, error) {
	l := logic.NewSkuStockDeleteLogic(ctx, s.svcCtx)
	return l.SkuStockDelete(in)
}

func (s *PmsServer) CollectList(ctx context.Context, in *pmsclient.CollectListReq) (*pmsclient.CollectListResp, error) {
	l := logic.NewCollectListLogic(ctx, s.svcCtx)
	return l.CollectList(in)
}

func (s *PmsServer) CollectAddOrDelete(ctx context.Context, in *pmsclient.CollectAddOrDeleteReq) (*pmsclient.CollectAddOrDeleteResp, error) {
	l := logic.NewCollectAddOrDeleteLogic(ctx, s.svcCtx)
	return l.CollectAddOrDelete(in)
}
