package Annoucement

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/Annoucement"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    AnnoucementReq "github.com/flipped-aurora/gin-vue-admin/server/model/Annoucement/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type AnnoucementApi struct {
}

var annouceService = service.ServiceGroupApp.AnnoucementServiceGroup.AnnoucementService


// CreateAnnoucement 创建Annoucement
// @Tags Annoucement
// @Summary 创建Annoucement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Annoucement.Annoucement true "创建Annoucement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /annouce/createAnnoucement [post]
func (annouceApi *AnnoucementApi) CreateAnnoucement(c *gin.Context) {
	var annouce Annoucement.Annoucement
	err := c.ShouldBindJSON(&annouce)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := annouceService.CreateAnnoucement(annouce); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAnnoucement 删除Annoucement
// @Tags Annoucement
// @Summary 删除Annoucement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Annoucement.Annoucement true "删除Annoucement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /annouce/deleteAnnoucement [delete]
func (annouceApi *AnnoucementApi) DeleteAnnoucement(c *gin.Context) {
	var annouce Annoucement.Annoucement
	err := c.ShouldBindJSON(&annouce)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := annouceService.DeleteAnnoucement(annouce); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAnnoucementByIds 批量删除Annoucement
// @Tags Annoucement
// @Summary 批量删除Annoucement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Annoucement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /annouce/deleteAnnoucementByIds [delete]
func (annouceApi *AnnoucementApi) DeleteAnnoucementByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := annouceService.DeleteAnnoucementByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAnnoucement 更新Annoucement
// @Tags Annoucement
// @Summary 更新Annoucement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Annoucement.Annoucement true "更新Annoucement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /annouce/updateAnnoucement [put]
func (annouceApi *AnnoucementApi) UpdateAnnoucement(c *gin.Context) {
	var annouce Annoucement.Annoucement
	err := c.ShouldBindJSON(&annouce)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := annouceService.UpdateAnnoucement(annouce); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAnnoucement 用id查询Annoucement
// @Tags Annoucement
// @Summary 用id查询Annoucement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Annoucement.Annoucement true "用id查询Annoucement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /annouce/findAnnoucement [get]
func (annouceApi *AnnoucementApi) FindAnnoucement(c *gin.Context) {
	var annouce Annoucement.Annoucement
	err := c.ShouldBindQuery(&annouce)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reannouce, err := annouceService.GetAnnoucement(annouce.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reannouce": reannouce}, c)
	}
}

// GetAnnoucementList 分页获取Annoucement列表
// @Tags Annoucement
// @Summary 分页获取Annoucement列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query AnnoucementReq.AnnoucementSearch true "分页获取Annoucement列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /annouce/getAnnoucementList [get]
func (annouceApi *AnnoucementApi) GetAnnoucementList(c *gin.Context) {
	var pageInfo AnnoucementReq.AnnoucementSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := annouceService.GetAnnoucementInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
