package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Annoucement"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/myPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup      system.ApiGroup
	ExampleApiGroup     example.ApiGroup
	MyPkgApiGroup       myPkg.ApiGroup
	AnnoucementApiGroup Annoucement.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
