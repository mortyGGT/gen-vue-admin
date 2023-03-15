package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/Annoucement"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/myPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System      system.RouterGroup
	Example     example.RouterGroup
	MyPkg       myPkg.RouterGroup
	Annoucement Annoucement.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
