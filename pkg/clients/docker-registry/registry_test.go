package docker_registry

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

const (
	_url     = "https://registry.cn-shanghai.aliyuncs.com"
	username = "xxxx"
	password = "xxxxx"
)

func Test_login(t *testing.T) {
	_, err := New(_url, "", "", func(format string, args ...interface{}) {
		t.Log(format, args)
	})
	assert.NoError(t, err)
	_, err = NewInsecure(_url, "", "", func(format string, args ...interface{}) {
		t.Log(format, args)
	})
	assert.NoError(t, err)
}
func Test_auth_login(t *testing.T) {
	_, err := New(_url, username, password, func(format string, args ...interface{}) {
		t.Log(format, args)
	})
	assert.NoError(t, err)
	_, err = NewInsecure(_url, username, password, func(format string, args ...interface{}) {
		t.Log(format, args)
	})
	assert.NoError(t, err)
}

func Test_transport_Login(t *testing.T) {
	round := WrapTransport(http.DefaultTransport, _url, username, password)
	r, err := http.NewRequest(http.MethodGet, _url, nil)
	require.NoError(t, err)
	_, err = round.RoundTrip(r)
	require.NoError(t, err)
}
