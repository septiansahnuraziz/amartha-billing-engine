package httpclient

import (
	"crypto/tls"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
)

// HTTPConnectionOptions options for the http connection
type HTTPConnectionOptions struct {
	TLSHandshakeTimeout   time.Duration
	TLSInsecureSkipVerify bool
	Timeout               time.Duration
	UseOpenTelemetry      bool
}

var defaultHTTPConnectionOptions = &HTTPConnectionOptions{
	TLSHandshakeTimeout:   100 * time.Second,
	TLSInsecureSkipVerify: false,
	Timeout:               200 * time.Second,
	UseOpenTelemetry:      false,
}

// NewHTTPConnection new http client
func NewHTTPConnection(opt *HTTPConnectionOptions) *http.Client {
	options := applyHTTPConnectionOptions(opt)

	httpClient := &http.Client{
		Timeout: options.Timeout,
		Transport: &http.Transport{
			TLSHandshakeTimeout: options.TLSHandshakeTimeout,
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: options.TLSInsecureSkipVerify},
		},
	}

	if !options.UseOpenTelemetry {
		return httpClient
	}

	return httpClient
}

func applyHTTPConnectionOptions(opt *HTTPConnectionOptions) *HTTPConnectionOptions {
	if opt != nil {
		return opt
	}

	return defaultHTTPConnectionOptions
}

func BuildHTTPRequest(method string, urlStr string, body io.Reader, headers *[]HttpHeaderDTO) (*http.Request, error) {
	var nrTrans newrelic.Transaction
	req, err := http.NewRequest(method, urlStr, body)

	segment := buildNRExternalSegment(&nrTrans, req)
	defer segment.End()

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	if headers != nil {
		if len(*headers) > 0 {
			for _, header := range *headers {
				req.Header.Add(header.Key, header.Value)
			}
		}
	}

	req = newrelic.RequestWithTransactionContext(req, &nrTrans)

	return req, err
}

func BuildHTTPRequestWithToken(tokenType string, token string, method string, urlStr string, body io.Reader) (*http.Request, error) {
	var nrTrans newrelic.Transaction
	req, err := http.NewRequest(method, urlStr, body)
	segment := buildNRExternalSegment(&nrTrans, req)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", tokenType+" "+token)
	req = newrelic.RequestWithTransactionContext(req, &nrTrans)

	defer segment.End()
	return req, err
}

func buildNRExternalSegment(nrTrans *newrelic.Transaction, req *http.Request) *newrelic.ExternalSegment {
	segment := newrelic.StartExternalSegment(nrTrans, req)
	if u := req.URL; nil != u {
		var firstSegment string
		if s := strings.Split(strings.TrimPrefix(u.Path, "/"), "/"); len(s) > 0 {
			firstSegment = s[0]
		}
		nu := url.URL{
			Scheme: u.Scheme,
			Host:   u.Host + "_" + firstSegment,
		}
		segment.URL = nu.String()
	}
	return segment
}
