package Annoucement

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Annoucement"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    AnnoucementReq "github.com/flipped-aurora/gin-vue-admin/server/model/Annoucement/request"
)

type AnnoucementService struct {
}

// CreateAnnoucement 创建Annoucement记录
// Author [piexlmax](https://github.com/piexlmax)
func (annouceService *AnnoucementService) CreateAnnoucement(annouce Annoucement.Annoucement) (err error) {
	err = global.GVA_DB.Create(&annouce).Error
	return err
}

// DeleteAnnoucement 删除Annoucement记录
// Author [piexlmax](https://github.com/piexlmax)
func (annouceService *AnnoucementService)DeleteAnnoucement(annouce Annoucement.Annoucement) (err error) {
	err = global.GVA_DB.Delete(&annouce).Error
	return err
}

// DeleteAnnoucementByIds 批量删除Annoucement记录
// Author [piexlmax](https://github.com/piexlmax)
func (annouceService *AnnoucementService)DeleteAnnoucementByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]Annoucement.Annoucement{},"id in ?",ids.Ids).Error
	return err
}

// UpdateAnnoucement 更新Annoucement记录
// Author [piexlmax](https://github.com/piexlmax)
func (annouceService *AnnoucementService)UpdateAnnoucement(annouce Annoucement.Annoucement) (err error) {
	err = global.GVA_DB.Save(&annouce).Error
	return err
}

// GetAnnoucement 根据id获取Annoucement记录
// Author [piexlmax](https://github.com/piexlmax)
func (annouceService *AnnoucementService)GetAnnoucement(id uint) (annouce Annoucement.Annoucement, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&annouce).Error
	return
}

// GetAnnoucementInfoList 分页获取Annoucement记录
// Author [piexlmax](https://github.com/piexlmax)
func (annouceService *AnnoucementService)GetAnnoucementInfoList(info AnnoucementReq.AnnoucementSearch) (list []Annoucement.Annoucement, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&Annoucement.Annoucement{})
    var annouces []Annoucement.Annoucement
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&annouces).Error
	return  annouces, total, err
}
