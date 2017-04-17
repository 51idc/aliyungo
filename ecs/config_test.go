package ecs

import "github.com/51idc/aliyungo/common"

//Modify with your Access Key Id and Access Key Secret

const (
	TestAccessKeyId     = ""
	TestAccessKeySecret = ""
	TestInstanceId      = "MY_TEST_INSTANCEID"
	TestSecurityGroupId = "MY_TEST_SECURITY_GROUP_ID"
	TestImageId         = "MY_IMAGE_ID"
	TestAccountId       = "MY_TEST_ACCOUNT_ID" //Get from https://account.console.aliyun.com
	TestRegionID        = common.APNorthEast1
	TestRegionID2       = common.Zhangjiakou
	TestRegionID3       = common.Hangzhou
	TestRegionID4       = common.Qingdao
	TestInstanceType    = "ecs.n4.large"
	TestVSwitchID       = "MY_TEST_VSWITCHID"

	TestIAmRich = false
	TestQuick   = false
)

var testClient *Client

func NewTestClient() *Client {
	if testClient == nil {
		testClient = NewClient(TestAccessKeyId, TestAccessKeySecret)
	}
	return testClient
}

var testDebugClient *Client

func NewTestClientForDebug() *Client {
	if testDebugClient == nil {
		testDebugClient = NewClient(TestAccessKeyId, TestAccessKeySecret)
		testDebugClient.SetDebug(true)
	}
	return testDebugClient
}

var testLocationClient *Client

func NetTestLocationClientForDebug() *Client {
	if testLocationClient == nil {
		testLocationClient = NewECSClient(TestAccessKeyId, TestAccessKeySecret, TestRegionID2)
		testLocationClient.SetDebug(true)
	}

	return testLocationClient
}
