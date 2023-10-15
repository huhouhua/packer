package check

import (
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
	check, err := m.check(ctx)
	if check {
		return err
	}
	newErr := next(false, ctx)
	if newErr != nil {
		return newErr
	}
	return err
}
func (*Middleware) check(ctx *process.PackageContext) (bool, error) {
	defer func() {
		if err := recover(); err != nil {
			logx.ErrorWithRed("check failed! %s ", err)
		}
	}()
	client, err := ctx.MinioClientFunc()
	if err != nil {
		return false, err
	}
	p := path.Join(ctx.Platform, ctx.Environment, ctx.Layer.BlobPath)
	logx.Info("start check %s whether there ", ctx.Layer.BlobPath)
	check, err := client.Exist(p)
	if err != nil && check {
		logx.Warning("%s already exists  skipping ... but there is have  %s", ctx.Layer.BlobPath, err.Error())
		ctx.Statistics.Skip = append(ctx.Statistics.Skip, ctx.Layer)
		return true, err
	}
	if check {
		logx.Warning("%s already exists  skipping ... ", ctx.Layer.BlobPath)
		ctx.Statistics.Skip = append(ctx.Statistics.Skip, ctx.Layer)
		return true, nil
	}
	logx.Info("check %s not exist ", ctx.Layer.BlobPath)
	return false, nil
}
