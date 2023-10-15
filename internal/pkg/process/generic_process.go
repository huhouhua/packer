package process

import (
	"context"
	"errors"
	genericoptions "ruijie.com.cn/devops/packer/internal/pkg/options"
	"ruijie.com.cn/devops/packer/internal/pkg/webhook"
	docker_registry "ruijie.com.cn/devops/packer/pkg/clients/docker-registry"
	"ruijie.com.cn/devops/packer/pkg/clients/minio"
	"ruijie.com.cn/devops/packer/pkg/logx"
	"sync"
	"time"
)

type GenericProcess struct {
	sync.Once
	manifests       docker_registry.ManifestSlice
	processFunc     RegisterFunc[Context]
	MinioClientFunc minio.ClientFunc
	webhookOptions  *genericoptions.WebHookOptions
	Bucket          string
	Environment     string
	Platform        string
	ManifestsPath   string
}

func (s *GenericProcess) ApplyTo(manifests docker_registry.ManifestSlice) {
	s.manifests = manifests
}

func (s *GenericProcess) Register(process RegisterFunc[Context]) error {
	s.Do(func() {
		s.processFunc = process
	})
	return nil
}

func (s *GenericProcess) process(ctx *Context) error {
	return s.processFunc(ctx)
}
func (s *GenericProcess) Run() error {
	logx.Info("start package.....")
	logx.Info("time out set is 20 minute ...")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Minute)
	defer cancel()
	workDone := make(chan struct{}, 1)
	pCtx := &Context{
		s.ManifestsPath,
		s.Bucket,
		s.Environment,
		s.Platform,
		s.MinioClientFunc,
	}
	var err error = nil
	go func() {
		err = s.process(pCtx)
		workDone <- struct{}{}
	}()
	hook := webhook.NewWebHook(s.webhookOptions)
	for {
		select {
		case <-workDone:
			logx.Info("package end!")
			if err == nil {
				hook.Notification("", true)
			} else {
				hook.Notification(err.Error(), false)
			}
			return err
		case <-ctx.Done():
			logx.Error("package time out ")
			hook.Notification("package time out", false)
			return errors.New("package time out")
		default:
		}
	}
}
func (s *GenericProcess) Close() {
	logx.Warning("Shutdown")
}
