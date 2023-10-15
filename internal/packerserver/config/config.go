package config

import "ruijie.com.cn/devops/packer/internal/packerserver/options"

type Config struct {
	*options.Options
}

func CreateConfigFromOptions(opts *options.Options) (*Config, error) {
	return &Config{opts}, nil
}
