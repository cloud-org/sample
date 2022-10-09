package logic

import (
	"context"
	"gin-reflect-handler/internal/svc"
	"gin-reflect-handler/internal/types"

	"github.com/gin-gonic/gin"
	"github.com/tal-tech/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson"
)

type Target struct {
	svcCtx *svc.ServiceContext
	ctx    context.Context
}

func NewTargetLogic(ctx *gin.Context, svcCtx *svc.ServiceContext) *Target {
	return &Target{
		svcCtx: svcCtx,
		ctx:    ctx,
	}
}

func (t *Target) GetAttrList(req *types.AttrListReq) (bson.M, error) {
	// todo: your logic

	logx.Infof("objectId: %s", req.ObjectId)

	return bson.M{"objectId": req.ObjectId}, nil
}
