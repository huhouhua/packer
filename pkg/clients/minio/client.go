package minio

import (
	"github.com/minio/minio-go"
	"github.com/minio/minio-go/pkg/credentials"
	"io"
	"net/http"
	"ruijie.com.cn/devops/packer/pkg/logx"
	"ruijie.com.cn/devops/packer/pkg/util"
	"time"
)

type ClientFunc func() (IClient, error)

var _ IClient = &Client{}

type IClient interface {
	PutObjectWithStream(objectPath string, fileSize int64, stream io.Reader) (bool, error)
	GetObjectStream(objectPath string) (*minio.Object, error)
	Exist(objectPath string) (bool, error)
}

type Client struct {
	*minio.Client
	bucket   string
	maxRetry int
}

func NewClient(endpoint string, accessKey string, secretKey string, isHttps bool, bucket string, maxRetry int) (IClient, error) {
	client, err := minio.NewWithOptions(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: isHttps,
	})
	return &Client{
		client,
		bucket,
		maxRetry,
	}, err
}
func (c *Client) PutObjectWithStream(objectPath string, fileSize int64, stream io.Reader) (bool, error) {
	opt := minio.PutObjectOptions{
		ContentType: "application/octet-stream",
		Progress:    NewProgressPar(fileSize)}
	retry := 0
	size := int64(0)
	var err error = nil
	for retry <= c.maxRetry {
		logx.Info("start %s upload file size  %s ...", objectPath, util.FileSizeToFormat(fileSize))
		size, err = c.PutObject(c.bucket, objectPath, stream, fileSize, opt)
		if err != nil {
			retry++
			logx.WarningWithMagenta("%s upload Failed! start retry for %d", err, retry)
			time.Sleep(time.Second * 3)
		} else {
			logx.SuccessWithGreen("%s of size %s Successfully", objectPath, util.FileSizeToFormat(size))
			break
		}
	}
	return retry <= c.maxRetry, err
}
func (c *Client) GetObjectStream(objectPath string) (*minio.Object, error) {
	opt := minio.GetObjectOptions{}
	//opt.Set("contentType", "application/json")
	return c.GetObject(c.bucket, objectPath, opt)
}
func (c *Client) Exist(objectPath string) (bool, error) {
	opt := minio.StatObjectOptions{}
	_, err := c.StatObject(c.bucket, objectPath, opt)
	if err != nil {
		resp := err.(minio.ErrorResponse)
		if resp.StatusCode == http.StatusNotFound {
			return false, err
		}
		return true, err
	}
	return true, nil
}
