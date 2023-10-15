package upload

import (
	"errors"
	"fmt"
	"io"
	"path"
	"ruijie.com.cn/devops/packer/internal/pkg/process"
	"ruijie.com.cn/devops/packer/pkg/logx"
)

type Middleware struct {
}

func NewMiddleware() *Middleware {
	return &Middleware{}

}
func (m *Middleware) Invoke(prev any, ctx *process.PackageContext, next process.HandlerFunc[process.PackageContext]) error {
	if prev == nil {
		return next(false, ctx)
	}
	reader := prev.(io.Reader)
	err := m.upload(reader, ctx)
	if err != nil {
		return err
	}
	return next(false, ctx)
}
func (m *Middleware) upload(reader io.Reader, ctx *process.PackageContext) error {
	defer func() {
		if err := recover(); err != nil {
			logx.ErrorWithRed("upload failed! %s ", err)
			ctx.Statistics.Fail = append(ctx.Statistics.Fail, ctx.Layer)
		}
	}()
	p := path.Join(ctx.Platform, ctx.Environment, ctx.Layer.BlobPath)
	client, err := ctx.MinioClientFunc()
	if err != nil {
		return err
	}
	suc, err := client.PutObjectWithStream(p, ctx.Layer.Size, reader)
	if err != nil {
		logx.ErrorWithRed("upload failed! %s", err.Error())
		ctx.Statistics.Fail = append(ctx.Statistics.Fail, ctx.Layer)
		return err
	}
	if !suc {
		ctx.Statistics.Fail = append(ctx.Statistics.Fail, ctx.Layer)
		return errors.New(fmt.Sprintf("%s upload  failed!", p))
	}
	ctx.Statistics.Success = append(ctx.Statistics.Success, ctx.Layer)
	return nil
}
