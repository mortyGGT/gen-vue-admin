package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/Annoucement"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type AnnoucementSearch struct{
    Annoucement.Annoucement
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
