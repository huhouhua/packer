package process

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	genericoptions "ruijie.com.cn/devops/packer/internal/pkg/options"
	"ruijie.com.cn/devops/packer/pkg/clients/minio"
	"sync"
)

// Config 配置genericpipeline的结构
type Config struct {
	minioOptions  *genericoptions.MinioOptions
	webhook       *genericoptions.WebHookOptions
	ManifestsPath string
	Bucket        string
	Environment   string
	Platform      string
}

func NewConfig(minioOptions *genericoptions.MinioOptions, webhookOptions *genericoptions.WebHookOptions) *Config {
	return &Config{
		minioOptions: minioOptions,
		webhook:      webhookOptions,
	}
}

type CompletedConfig struct {
	*Config
}

func (c *Config) Complete() CompletedConfig {
	return CompletedConfig{c}
}

func (c CompletedConfig) New() (*GenericProcess, error) {
	s := &GenericProcess{
		Once:            sync.Once{},
		MinioClientFunc: c.createMinioClient,
		Bucket:          c.Bucket,
		Environment:     c.Environment,
		Platform:        c.Platform,
		ManifestsPath:   c.ManifestsPath,
		webhookOptions:  c.webhook,
	}
	return s, nil
}

func (c CompletedConfig) createMinioClient() (minio.IClient, error) {
	client, err := minio.NewClient(
		c.minioOptions.Endpoint,
		c.minioOptions.AccessKey,
		c.minioOptions.SecretKey,
		c.minioOptions.IsHttps,
		c.Bucket,
		c.minioOptions.MaxRetry)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s create minio client failed!", err.Error()))
	}
	return client, nil
}

func Load() {
	viper.SetConfigName("packer-server")
	viper.AddConfigPath(".")
	viper.AddConfigPath("configs")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("packer")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("加载配置文件失败！ %s", err.Error()))
	}
}
