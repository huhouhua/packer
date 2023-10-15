package process

import (
	"ruijie.com.cn/devops/packer/internal/packerserver/process/check"
	"ruijie.com.cn/devops/packer/internal/packerserver/process/download"
	"ruijie.com.cn/devops/packer/internal/packerserver/process/manifest"
	"ruijie.com.cn/devops/packer/internal/packerserver/process/upload"
	genericprocess "ruijie.com.cn/devops/packer/internal/pkg/process"
	"sync"
)

var once = sync.Once{}

func Install(g *genericprocess.GenericProcess) {
	once.Do(func() {
		installStoreProcess()
		installPackageProcess()
		err := g.Register(process)
		if err != nil {
			return
		}
	})
}

// installStoreProcess 注册获取镜像清单过程
func installStoreProcess() {
	handlerStoreFunc = genericprocess.NewBuilder[genericprocess.ManifestContext]().
		UseWhen(manifest.NewMiddleware()).
		Build()
}

// installPackageProcess 注册打包过程
func installPackageProcess() {
	handlerPackageFunc = genericprocess.NewBuilder[genericprocess.PackageContext]().
		UseWhen(check.NewMiddleware()).
		UseWhen(download.NewMiddleware()).
		UseWhen(upload.NewMiddleware()).
		Build()
}
