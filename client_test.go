package gofire

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testTargetGeode string
var geodeTestClient *Client

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
	assert.Nil(t, err, )
	assert.EqualValues(t, regions.RegionInfo[0].Name, "testRegion")
}