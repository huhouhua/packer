package options

import (
	"github.com/spf13/viper"
	genericoptions "ruijie.com.cn/devops/packer/internal/pkg/options"
	genericprocess "ruijie.com.cn/devops/packer/internal/pkg/process"
)

type Options struct {
	ImageStore *genericoptions.ImageStoreOptions `json:"image-store"   mapstructure:"image-store"`
	Minio      *genericoptions.MinioOptions      `json:"minio"   mapstructure:"minio"`
	WebHook    *genericoptions.WebHookOptions    `json:"webhook"   mapstructure:"webhook"`
}

func NewOptions() *Options {
	return &Options{
		ImageStore: genericoptions.NewImageStoreOptions(),
		Minio:      genericoptions.NewMinioOptions(),
		WebHook:    genericoptions.NewWebHookOptions(),
	}
}

func (s *Options) ApplyTo(config *genericprocess.Config) error {
	err := viper.Unmarshal(s)
	if err != nil {
		return err
	}
	config.Bucket = s.ImageStore.Bucket
	config.Environment = s.ImageStore.Environment
	config.Platform = s.ImageStore.Platform
	config.ManifestsPath = s.ImageStore.ManifestsPath
	return nil
}
