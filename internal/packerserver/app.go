package packerserver

import (
	"ruijie.com.cn/devops/packer/internal/packerserver/config"
	"ruijie.com.cn/devops/packer/internal/packerserver/options"
	"ruijie.com.cn/devops/packer/pkg/app"
)

func NewApp() *app.App {
	opts := options.NewOptions()
	application := app.NewApp("packer",
		app.WithVersion("0.1.0"),
		app.WithDescription("Make a deployment package"),
		app.WithRunFunc(run(opts)))
	return application
}
func run(opts *options.Options) app.RunFunc {
	return func() error {
		cfg, err := config.CreateConfigFromOptions(opts)
		if err != nil {
			return err
		}
		return Run(cfg)
	}
}
