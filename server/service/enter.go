package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/Annoucement"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/myPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup      system.ServiceGroup
	ExampleServiceGroup     example.ServiceGroup
	MyPkgServiceGroup       myPkg.ServiceGroup
	AnnoucementServiceGroup Annoucement.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
