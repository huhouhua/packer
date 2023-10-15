package manifest

import (
	"errors"
	"fmt"
	"io"
	"path"
	"ruijie.com.cn/devops/packer/internal/pkg/process"
	"ruijie.com.cn/devops/packer/pkg/logx"
	"ruijie.com.cn/devops/packer/pkg/util"
)

type Middleware struct {
}

func NewMiddleware() *Middleware {
	return &Middleware{}
}

func (m *Middleware) Invoke(prev any, ctx *process.ManifestContext, next process.HandlerFunc[process.ManifestContext]) error {
	p := path.Join(ctx.Platform, ctx.Environment, ctx.ManifestsPath)
	client, err := ctx.MinioClientFunc()
	if err != nil {
		return err
	}
	logx.Info("download  manifests from %s", p)

	reader, err := client.GetObjectStream(p)
	if err != nil {
		return errors.New(fmt.Sprintf("%s download failed!", err.Error()))
	}
	defer reader.Close()

	obj, err := reader.Stat()
	if err != nil {
		return errors.New(fmt.Sprintf("%s reader failed!", err.Error()))
	}
	ma, err := io.ReadAll(reader)
	err = ctx.UnmarshalWithManifest(ma)
	if err != nil {
		return errors.New(fmt.Sprintf("%s unmarshal failed!", err.Error()))
	}
	logx.Info("manifests file size %s", util.FileSizeToFormat(obj.Size))
	logx.SuccessWithGreen("download  manifests Successful")
	return next("manifest", ctx)
}
