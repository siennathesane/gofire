package gofire

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testTargetGeode string
var geodeTestClient *Client

type testRegion struct {
	// the struct tag has to match the region name or this will fail.
	Region interface{} `json:"testRegion"`
}

type testData struct {
	Message string `json:"message"`
	Value   int    `json:"value"`
}

type testNestedData struct {
	Message string   `json:"message"`
	Object  testData `json:"object"`
}

func TestMain(m *testing.M) {
	testTargetGeode = os.Getenv("GEODE_URL")
	m.Run()
}

func TestGoodNewClient(t *testing.T) {
	// this also tests gofire.Client.Ping()
	client, err := NewClient(testTargetGeode, true)
	if !assert.Nil(t, err) {
		assert.FailNow(t, "cannot talk to geode! is GEODE_URL set?")
	}
	assert.Equal(t, client.GeodeUrl, testTargetGeode, "client connection is good.")
	// if we're good at this point, save it for reuse.
	geodeTestClient = client
	geodeTestClient.Region = "testRegion"
}

func TestBadNewClient(t *testing.T) {
	_, err := NewClient("http://localhost:123", true)
	assert.NotNil(t, err, "bad client is bad.")
}

func TestClientGetGroup(t *testing.T) {
	t.Run("GetServers", testClient_GetServers)
	t.Run("GetRegions", testClient_GetRegions)
	t.Run("GetRegion", testClient_GetRegion)
	t.Run("GetKeys", testClient_GetKeys)
	t.Run("GetKey", testClient_Get)
	t.Run("NumEntries", testClient_NumEntries)
}

func TestClientPutGroup(t *testing.T) {
	t.Run("PutGroup", func(t *testing.T) {
		t.Run("PutKey", testClient_Put)
	})
}

func TestClientDeleteGroup(t *testing.T) {
	t.Run("DeleteKey", testClient_Delete)
}

func testClient_GetServers(t *testing.T) {
	servers, err := geodeTestClient.GetServers()
	assert.Nil(t, err, "no error when getting test servers.")
	assert.EqualValues(t, testTargetGeode, servers[0])
}

func testClient_GetRegions(t *testing.T) {
	regions, err := geodeTestClient.GetRegions()
	assert.Nil(t, err)
	assert.EqualValues(t, regions.RegionInfo[0].Name, "testRegion")
}

func testClient_GetRegion(t *testing.T) {
	data, err := geodeTestClient.GetRegion(50, false)
	assert.Nil(t, err, "failed to get data from testRegion.")
	var tmpData testRegion
	if err := json.Unmarshal(data, &tmpData); err != nil {
		assert.Nil(t, err, "cannot unmarshal json data.")
	}
}

func testClient_GetKeys(t *testing.T) {
	testKeyList := &KeyList{
		Keys: []string{"testTypeDataKey", "testNestedKey", "testKey"},
	}
	keys, err := geodeTestClient.GetKeys("")
	assert.Nil(t, err, "error getting keys.")
	assert.EqualValues(t, testKeyList.Keys, keys.Keys)

	_, err = geodeTestClient.GetKeys("doesNotExist")
	assert.NotNil(t, err, "should return a 404 on the region.")
}

func testClient_Get(t *testing.T) {
	var tmpTestData testData
	if err := geodeTestClient.Get("", &tmpTestData); err != nil {
		assert.Nil(t, err, "cannot properly parse testData.")
	}

	var tmpTestNestedData testNestedData
	if err := geodeTestClient.Get("", &tmpTestNestedData); err != nil {
		assert.Nil(t, err, "cannot properly parse testNestedData.")
	}

	if err := geodeTestClient.Get("doesNotExist", &tmpTestData); err != nil {
		assert.NotNil(t, err, "this region should not exist.")
	}
}

func testClient_NumEntries(t *testing.T) {
	// TODO (mxplusb): add another case once setting keys is implemented.
	testCount, err := geodeTestClient.NumEntries("")
	if err != nil {
		assert.Nil(t, err, "can't get a count of entries from geode.")
	}
	assert.Equal(t, 3, testCount)
}

func testClient_Put(t *testing.T) {
	testMap := testNestedData{
		Message: "Hello, world.",
		Object: testData{
			Message: "I'm a potato.",
			Value:   2,
		},
	}

	err := geodeTestClient.Put("testMap", testMap)
	assert.Nil(t, err)

	var tmpTestData testNestedData
	err = geodeTestClient.Get("testMap", &tmpTestData)
	assert.Nil(t, err, "error getting testMap.")
	assert.EqualValues(t, testMap.Object.Value, tmpTestData.Object.Value)
}

func testClient_Delete(t *testing.T) {
	testMap := testNestedData{
		Message: "Hello, world.",
		Object: testData{
			Message: "I'm a potato.",
			Value:   2,
		},
	}

	err := geodeTestClient.Put("testMap", testMap)
	assert.Nil(t, err)

	err = geodeTestClient.Delete("testMap")
	assert.Nil(t, err)
}