package options

type MinioOptions struct {
	AccessKey string `json:"access-key"       mapstructure:"access-key"`
	SecretKey string `json:"secret-key"       mapstructure:"secret-key"`
	IsHttps   bool   `json:"is-https"       mapstructure:"is-https"`
	Endpoint  string `json:"endpoint"       mapstructure:"endpoint"`
	MaxRetry  int    `json:"max-retry"       mapstructure:"max-retry"`
}

func NewMinioOptions() *MinioOptions {
	return &MinioOptions{
		AccessKey: "",
		SecretKey: "",
		Endpoint:  "",
		IsHttps:   false,
		MaxRetry:  0,
	}
}
