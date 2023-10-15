package options

type ImageStoreOptions struct {
	ManifestsPath string          `json:"manifests-path"       mapstructure:"manifests-path"`
	Bucket        string          `json:"bucket"               mapstructure:"bucket"`
	Environment   string          `json:"environment"           mapstructure:"environment"`
	Platform      string          `json:"platform"           mapstructure:"platform"`
	Registry      RegistryOptions `json:"registry"          mapstructure:"registry"`
}

type RegistryOptions struct {
	Default DefaultRegistry `json:"default"               mapstructure:"default"`
}

type DefaultRegistry struct {
	Username string `json:"username"               mapstructure:"username"`
	Password string `json:"password"               mapstructure:"password"`
}

func NewImageStoreOptions() *ImageStoreOptions {
	return &ImageStoreOptions{
		ManifestsPath: "",
		Bucket:        "devops",
		Environment:   "develop",
		Platform:      "platform",
		Registry: RegistryOptions{
			Default: DefaultRegistry{
				Username: "",
				Password: "",
			},
		},
	}
}
