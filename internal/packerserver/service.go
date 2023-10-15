package packerserver

import (
	"log"
	"ruijie.com.cn/devops/packer/internal/packerserver/config"
	"ruijie.com.cn/devops/packer/internal/packerserver/process"
	genericprocess "ruijie.com.cn/devops/packer/internal/pkg/process"
	"ruijie.com.cn/devops/packer/pkg/shutdown"
)

type builder struct {
	gs             *shutdown.GracefulShutdown
	genericProcess *genericprocess.GenericProcess
	cfg            *config.Config
}

type preparedBuilder struct {
	*builder
}

func createBuilder(cfg *config.Config) (*builder, error) {
	gs := shutdown.New()

	genericConfig, err := buildGenericConfig(cfg)
	if err != nil {
		return nil, err
	}
	genericProcess, err := genericConfig.Complete().New()
	if err != nil {
		return nil, err
	}
	server := &builder{
		gs:             gs,
		cfg:            cfg,
		genericProcess: genericProcess,
	}
	return server, nil
}

func (s *builder) PrepareRun() preparedBuilder {
	process.Install(s.genericProcess)

	s.gs.AddShutdownCallback(shutdown.ShutdownFunc(func(string) error {
		s.genericProcess.Close()
		return nil
	}))
	return preparedBuilder{s}
}

func (s preparedBuilder) Run() error {
	if err := s.gs.Start(); err != nil {
		log.Fatalf("启动管理器失败: %s", err.Error())
	}
	return s.genericProcess.Run()
}

func buildGenericConfig(cfg *config.Config) (genericConfig *genericprocess.Config, lastErr error) {
	genericprocess.Load()
	genericConfig = genericprocess.NewConfig(
		cfg.Minio,
		cfg.WebHook)

	if lastErr = cfg.ApplyTo(genericConfig); lastErr != nil {
		return
	}
	return
}
