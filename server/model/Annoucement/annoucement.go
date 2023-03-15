// 自动生成模板Annoucement
package Annoucement

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// Annoucement 结构体
type Annoucement struct {
      global.GVA_MODEL
      Title  string `json:"title" form:"title" gorm:"column:title;comment:;"`
      Content  string `json:"content" form:"content" gorm:"column:content;comment:;"`
      Create_date  *time.Time `json:"create_date" form:"create_date" gorm:"column:create_date;comment:;"`
}


// TableName Annoucement 表名
func (Annoucement) TableName() string {
  return "annoucement"
}

