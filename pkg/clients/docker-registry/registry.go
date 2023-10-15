package docker_registry

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type LogfCallback func(format string, args ...interface{})

/*
 * Discard log messages silently.
 */
func Quiet(format string, args ...interface{}) {
	/* discard logs */
}

/*
 * Pass log messages along to Go's "log" module.
 */
func Log(format string, args ...interface{}) {
	log.Printf(format, args...)
}

type Registry struct {
	URL    string
	Client *http.Client
	Logf   LogfCallback
}

/*
 * 使用给定的URL和凭据创建一个新的Registry，然后Ping()执行在返回之前验证注册表是否可用
 *
 *
 * 您也可以通过填充字段来手动构建Registry
 * 这传递http。DefaultTransport为WrapTransport
 * http.Client.
 */
func New(registryURL, username, password string, Logf LogfCallback) (*Registry, error) {
	transport := http.DefaultTransport
	return newFromTransport(registryURL, username, password, transport, Logf)
}

/*
 * 创建一个新的 Registry, 与New一样，使用http。禁用的传输SSL证书验证
 */
func NewInsecure(registryURL, username, password string, Logf LogfCallback) (*Registry, error) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			// TODO: Why?
			InsecureSkipVerify: true, //nolint:gosec
		},
	}

	return newFromTransport(registryURL, username, password, transport, Logf)
}

/*
 * 给定一个现有的http。RoundTripper，如http。DefaultTransport，构建
 * 传输栈是向Docker注册表API进行身份验证所必需的。这
 * 添加了对OAuth承载令牌和HTTP基本验证的支持，并设置
 * 此库所依赖的错误处理。
 */
func WrapTransport(transport http.RoundTripper, url, username, password string) http.RoundTripper {
	tokenTransport := &TokenTransport{
		Transport: transport,
		Username:  username,
		Password:  password,
	}
	basicAuthTransport := &BasicTransport{
		Transport: tokenTransport,
		URL:       url,
		Username:  username,
		Password:  password,
	}
	errorTransport := &ErrorTransport{
		Transport: basicAuthTransport,
	}
	return errorTransport
}

func newFromTransport(registryURL, username, password string, transport http.RoundTripper, logf LogfCallback) (*Registry, error) {
	url := strings.TrimSuffix(registryURL, "/")
	transport = WrapTransport(transport, url, username, password)
	registry := &Registry{
		URL: url,
		Client: &http.Client{
			Transport: transport,
		},
		Logf: logf,
	}

	if err := registry.Ping(); err != nil {
		return nil, err
	}

	return registry, nil
}

func (r *Registry) url(pathTemplate string, args ...interface{}) string {
	pathSuffix := fmt.Sprintf(pathTemplate, args...)
	url := fmt.Sprintf("%s%s", r.URL, pathSuffix)
	return url
}

func (r *Registry) Ping() error {
	url := r.url("/v2/")
	r.Logf("registry.ping url=%s", url)
	resp, err := r.Client.Get(url)
	if resp != nil {
		defer resp.Body.Close()
	}
	return err
}
