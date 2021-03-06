package dns

import (
	"github.com/51idc/aliyungo/common"
	"os"
)

type Client struct {
	common.Client
}

const (
	// DNSDefaultEndpoint is the default API endpoint of DNS services
	DNSDefaultEndpoint = "http://dns.aliyuncs.com"
	DNSAPIVersion      = "2015-01-09"
)

// NewClient creates a new instance of DNS client
func NewClient(accessKeyId, accessKeySecret string) *Client {
	endpoint := os.Getenv("DNS_ENDPOINT")
	if endpoint == "" {
		endpoint = DNSDefaultEndpoint
	}
	return NewClientWithEndpoint(endpoint, accessKeyId, accessKeySecret)
}

// NewCustomClient creates a new instance of ECS client with customized API endpoint
func NewCustomClient(accessKeyId, accessKeySecret string, endpoint string) *Client {
	client := &Client{}
	client.Init(endpoint, DNSAPIVersion, accessKeyId, accessKeySecret)
	return client
}

func NewClientWithEndpoint(endpoint string, accessKeyId, accessKeySecret string) *Client {
	client := &Client{}
	client.Init(endpoint, DNSAPIVersion, accessKeyId, accessKeySecret)
	return client
}
