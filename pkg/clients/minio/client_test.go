package minio

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"testing"
)

const (
	accessKey     = "admin"
	secretKey     = "huhouhua"
	endpoint      = "localhost:9000"
	bucket        = "devops"
	manifestsPath = "platform/develop/manifests/开发系统.rar"
	uploadPath    = "platform/test/manifests/开发系统.rar"
)

func Test_PutObjectWithStream(t *testing.T) {
	client, err := NewClient(endpoint, accessKey, secretKey, false, bucket, 10)
	require.NoError(t, err)
	require.NotNil(t, client, "client not nil")

	reader, err := client.GetObjectStream(manifestsPath)

	require.NoError(t, err)
	require.NotNil(t, reader)
	defer reader.Close()

	objectInfo, err := reader.Stat()
	require.NoError(t, err)
	require.NotNil(t, objectInfo)

	suc, err := client.PutObjectWithStream(uploadPath, objectInfo.Size, reader)
	require.NoError(t, err)
	require.True(t, suc, "upload failed", err)
}
func Test_GetObjectStream(t *testing.T) {
	path := "platform/develop/manifests/manifests.json"
	client, err := NewClient(endpoint, accessKey, secretKey, false, bucket, 10)
	require.NoError(t, err)
	require.NotNil(t, client, "client not nil")
	reader, err := client.GetObjectStream(path)

	require.NoError(t, err)
	require.NotNil(t, reader)
	defer reader.Close()

	manifests, err := io.ReadAll(reader)
	require.NoError(t, err)

	t.Logf("file size %d \n ", len(manifests))
	t.Log(string(manifests))
}
func Test_Exist(t *testing.T) {
	path := "platform/develop/manifests/manifests.json"
	client, err := NewClient(endpoint, accessKey, secretKey, false, bucket, 10)
	require.NoError(t, err)
	require.NotNil(t, client, "client not nil")
	found, err := client.Exist(path)
	require.NoError(t, err, "not found!")
	assert.True(t, found)
}
