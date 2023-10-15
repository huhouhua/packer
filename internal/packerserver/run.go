package packerserver

import "ruijie.com.cn/devops/packer/internal/packerserver/config"

func Run(cfg *config.Config) error {
	server, err := createBuilder(cfg)
	if err != nil {
		return err
	}
	return server.PrepareRun().Run()
}
