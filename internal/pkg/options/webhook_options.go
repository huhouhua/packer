package options

type WebHookOptions struct {
	Url       string `json:"url"       mapstructure:"url"`
	PackageId string `json:"package-id"       mapstructure:"package-id"`
}

func NewWebHookOptions() *WebHookOptions {
	return &WebHookOptions{Url: "", PackageId: ""}
}
