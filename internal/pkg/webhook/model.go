package webhook

type NotificationRequest struct {
	PackageId string
	IsSuccess bool
	Message   string
}
