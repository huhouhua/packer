package webhook

import (
	genericoptions "ruijie.com.cn/devops/packer/internal/pkg/options"
	"time"
)

type WebHook struct {
	opt *genericoptions.WebHookOptions
}

func NewWebHook(opt *genericoptions.WebHookOptions) *WebHook {
	return &WebHook{opt: opt}
}

func (w *WebHook) Notification(msg string, success bool) {
	go func(hookOptions *genericoptions.WebHookOptions) {
		client := NewClient()
		_ = client.Notification(hookOptions.Url, NotificationRequest{
			IsSuccess: success,
			Message:   msg,
			PackageId: hookOptions.PackageId,
		})
	}(w.opt)
	time.Sleep(time.Second * 3)
}
