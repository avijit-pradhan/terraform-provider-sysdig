package v2

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type SysdigRequest struct {
	config     *config
	httpClient *http.Client
}

type SysdigCommon interface {
	Common
}

type SysdigMonitor interface {
	SysdigCommon
}

type SysdigSecure interface {
	SysdigCommon
}

func (sr *SysdigRequest) Request(ctx context.Context, method string, url string, payload io.Reader) (*http.Response, error) {
	r, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}

	r = r.WithContext(ctx)
	r.Header.Set(AuthorizationHeader, fmt.Sprintf("Bearer %s", sr.config.token))
	r.Header.Set(ContentTypeHeader, ContentTypeJSON)

	return request(sr.httpClient, sr.config, r)
}

func NewSysdigMonitor(opts ...ClientOption) SysdigMonitor {
	return newSysdigClient(opts...)
}

func NewSysdigSecure(opts ...ClientOption) SysdigSecure {
	return newSysdigClient(opts...)
}

func newSysdigClient(opts ...ClientOption) *Client {
	cfg := configure(opts...)
	return &Client{
		config: cfg,
		requester: &SysdigRequest{
			config:     cfg,
			httpClient: newHTTPClient(cfg),
		},
	}
}
