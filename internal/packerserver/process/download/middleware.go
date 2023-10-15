package download

import (
	"github.com/opencontainers/go-digest"
	"io"
	"ruijie.com.cn/devops/packer/internal/pkg/process"
	"ruijie.com.cn/devops/packer/pkg/logx"
)

type Middleware struct {
}

func NewMiddleware() *Middleware {
	return &Middleware{}

}

func (m *Middleware) Invoke(prev any, ctx *process.PackageContext, next process.HandlerFunc[process.PackageContext]) error {
	reader, err := m.download(ctx)
	newErr := next(reader, ctx)
	if newErr != nil {
		return newErr
	}
	return err
}
func (m *Middleware) download(ctx *process.PackageContext) (io.Reader, error) {
	defer func() {
		if err := recover(); err != nil {
			logx.ErrorWithRed(" digest download failed! %s ", err)
			ctx.Statistics.Fail = append(ctx.Statistics.Fail, ctx.Layer)
		}
	}()
	logx.Info("start download digest %s", ctx.Layer.Digest)
	d, _ := digest.Parse(ctx.Layer.Digest)
	reader, err := ctx.Registry.DownloadBlob(ctx.Repo.RepoName, d)

	if err != nil {
		logx.ErrorWithRed("download failed! %s", err.Error())
		ctx.Statistics.Fail = append(ctx.Statistics.Fail, ctx.Layer)
		return nil, err
	}
	logx.SuccessWithGreen("download digest %s Successful", d.String())
	return reader, nil
}
