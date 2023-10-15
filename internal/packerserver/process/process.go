package process

import (
	"errors"
	genericprocess "ruijie.com.cn/devops/packer/internal/pkg/process"
	statistics "ruijie.com.cn/devops/packer/internal/pkg/statistics"
	docker "ruijie.com.cn/devops/packer/pkg/clients/docker-registry"
	"ruijie.com.cn/devops/packer/pkg/clients/minio"
	"ruijie.com.cn/devops/packer/pkg/logx"
	"time"
)

var handlerStoreFunc genericprocess.HandlerFunc[genericprocess.ManifestContext]
var handlerPackageFunc genericprocess.HandlerFunc[genericprocess.PackageContext]

func process(ctx *genericprocess.Context) error {
	t := time.Now()
	manifests, err := generate(ctx)
	if err != nil {
		return err
	}
	newStatistic := statistics.NewStatistics()
	var newErr error = nil
	for i, m := range manifests {
		statistics.Print(m, "image information", true)
		sta, err := handler(m, ctx.MinioClientFunc, ctx.Environment, ctx.Platform)
		if newErr == nil {
			newErr = err
		}
		newStatistic.Add(i, statistics.NewValuePair(m, sta))
	}
	elapsed := time.Since(t)
	newStatistic.Print()
	logx.Info("total package time is %s ", elapsed.String())
	if !newStatistic.IsSuccess {
		return errors.New("fail！")
	}
	if newErr != nil {
		return newErr
	}
	return nil
}

func handler(manifest docker.Manifest, minioClientFunc minio.ClientFunc, environment string, platform string) (*statistics.Manifest, error) {
	statisticManifest := statistics.NewManifest()
	ctx := &genericprocess.PackageContext{
		Statistics:      statisticManifest,
		Layer:           manifest.Config.Layer,
		Repo:            manifest.RepoTag,
		Environment:     environment,
		Platform:        platform,
		MinioClientFunc: minioClientFunc,
	}
	reg, err := docker.New(manifest.RepoTag.ExternalRegistry, manifest.RepoTag.UserName, manifest.RepoTag.PassWord, func(format string, args ...interface{}) {
		logx.Info(format, args)
	})
	if err != nil {
		return statisticManifest, err
	}
	ctx.Registry = reg
	err = handlerConfig(ctx)
	if err != nil {
		return statisticManifest, err
	}
	var newErr error = nil
	for _, layer := range manifest.Layers {
		ctx.Layer = layer
		err := handlerPackageFunc("", ctx)
		if newErr == nil {
			newErr = err
		}
	}
	if newErr != nil {
		return statisticManifest, newErr
	}
	return statisticManifest, nil
}

func handlerConfig(ctx *genericprocess.PackageContext) error {
	return handlerPackageFunc("", ctx)
}

func generate(ctx *genericprocess.Context) (docker.ManifestSlice, error) {
	logx.Info("start generate image manifest ...")
	defer func() {
		logx.Info("generate end")
	}()
	manCtx := genericprocess.NewManifestContext(ctx.Platform, ctx.MinioClientFunc, ctx.ManifestsPath, ctx.Environment)
	err := handlerStoreFunc("", manCtx)
	if err != nil {
		return nil, err
	}
	return manCtx.Manifests, nil
}
