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
	Message string `json:"message"`
	Object struct {
		Message string `json:"message"`
		Value   int    `json:"value"`
	} `json:"object"`
}

func TestMain(m *testing.M) {
	testTargetGeode = os.Getenv("GEODE_URL")
	m.Run()
}

func TestGoodNewClient(t *testing.T) {
	client, err := NewClient(testTargetGeode, true)
	assert.Nil(t, err)
	assert.Equal(t, client.GeodeUrl, testTargetGeode, "client connection is good.")
	// if we're good at this point, save it for reuse.
	geodeTestClient = client
}

func TestBadNewClient(t *testing.T) {
	_, err := NewClient("http://localhost:123", true)
	assert.NotNil(t, err, "bad client is bad.")
}

func TestClient_GetServers(t *testing.T) {
	servers, err := geodeTestClient.GetServers()
	assert.Nil(t, err, "no error when getting test servers.")
	assert.EqualValues(t, testTargetGeode, servers[0])
}

func TestClient_GetRegions(t *testing.T) {
	regions, err := geodeTestClient.GetRegions()
	assert.Nil(t, err)
	assert.EqualValues(t, regions.RegionInfo[0].Name, "testRegion")
}

func TestClient_GetRegion(t *testing.T) {
	data, err := geodeTestClient.GetRegion("testRegion", 50, false)
	assert.Nil(t, err, "failed to get data from testRegion.")
	var tmpData testRegion
	if err := json.Unmarshal(data, &tmpData); err != nil {
		assert.Nil(t, err, "cannot unmarshal json data.")
	}
}
