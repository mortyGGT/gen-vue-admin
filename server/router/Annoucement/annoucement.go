package Annoucement

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AnnoucementRouter struct {
}

// InitAnnoucementRouter 初始化 Annoucement 路由信息
func (s *AnnoucementRouter) InitAnnoucementRouter(Router *gin.RouterGroup) {
	annouceRouter := Router.Group("annouce").Use(middleware.OperationRecord())
	annouceRouterWithoutRecord := Router.Group("annouce")
	var annouceApi = v1.ApiGroupApp.AnnoucementApiGroup.AnnoucementApi
	{
		annouceRouter.POST("createAnnoucement", annouceApi.CreateAnnoucement)   // 新建Annoucement
		annouceRouter.DELETE("deleteAnnoucement", annouceApi.DeleteAnnoucement) // 删除Annoucement
		annouceRouter.DELETE("deleteAnnoucementByIds", annouceApi.DeleteAnnoucementByIds) // 批量删除Annoucement
		annouceRouter.PUT("updateAnnoucement", annouceApi.UpdateAnnoucement)    // 更新Annoucement
	}
	{
		annouceRouterWithoutRecord.GET("findAnnoucement", annouceApi.FindAnnoucement)        // 根据ID获取Annoucement
		annouceRouterWithoutRecord.GET("getAnnoucementList", annouceApi.GetAnnoucementList)  // 获取Annoucement列表
	}
}
