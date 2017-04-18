package common

import (
	"testing"
)

func TestLoadEndpointFromFile(t *testing.T) {

}

func TestDescribeOpenAPIEndpoint(t *testing.T) {
	client := NewLocationClient("", "")
	ep := client.DescribeOpenAPIEndpoint(Hangzhou, "ecs")
	t.Log(">>>>>>>>>:", ep)
}
