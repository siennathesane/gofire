package gofire

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Ping tests the connectivity of a Geode server.
func (cl Client) Ping() (int, error, bool) {
	req, err := cl.getRequestBuilder("/ping")
	if err != nil {
		return 400, errors.New(fmt.Sprintf("error building request. error: %s", err)), false
	}
	resp, err := cl.intClient.Do(req)
	if err != nil {
		return 400, err, false
	}
	switch resp.StatusCode {
	case 500:
		return 500, errors.New("encountered error at server. check the Geode exception trace"), false
	case 404:
		return 404, errors.New("the Developer REST API service is not available"), false
	case 200:
		return 200, nil, true
	default:
		return 400, errors.New("client error"), false
	}
}

// Gets the servers in the cluster.
func (cl Client) GetServers() ([]string, error) {
	servers := make([]string, 0)
	req, err := cl.getRequestBuilder("/servers")
	if err != nil {
		return servers, errors.New(fmt.Sprintf("error building request. error: %s", err))
	}
	resp, err := cl.intClient.Do(req)
	if err != nil {
		return servers, err
	}
	if err = json.NewDecoder(resp.Body).Decode(&servers); err != nil {
		return servers, err
	}
	return servers, nil
}
