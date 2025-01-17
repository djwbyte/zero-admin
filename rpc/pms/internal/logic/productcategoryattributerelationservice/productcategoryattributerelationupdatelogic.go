package productcategoryattributerelationservicelogic

import (
	"context"
	"zero-admin/rpc/model/pmsmodel"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductCategoryAttributeRelationUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductCategoryAttributeRelationUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCategoryAttributeRelationUpdateLogic {
	return &ProductCategoryAttributeRelationUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductCategoryAttributeRelationUpdateLogic) ProductCategoryAttributeRelationUpdate(in *pmsclient.ProductCategoryAttributeRelationUpdateReq) (*pmsclient.ProductCategoryAttributeRelationUpdateResp, error) {
	err := l.svcCtx.PmsProductCategoryAttributeRelationModel.Update(l.ctx, &pmsmodel.PmsProductCategoryAttributeRelation{
		Id:                 in.Id,
		ProductCategoryId:  in.ProductCategoryId,
		ProductAttributeId: in.ProductAttributeId,
	})
	if err != nil {
		return nil, err
	}

	return &pmsclient.ProductCategoryAttributeRelationUpdateResp{}, nil
}
