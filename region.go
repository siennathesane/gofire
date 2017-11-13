package gofire

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type RegionInfo struct {
	Name            string `json:"name"`
	RegionType      string `json:"type"`
	KeyConstraint   string `json:"key-constraint"`
	ValueConstraint string `json:"value-constraint"`
}

type Regions struct {
	RegionInfo []RegionInfo `json:"regions"`
}

// Lists out the regions a given server hosts.
func (cl Client) GetRegions() (Regions, error) {
	regions := Regions{
		RegionInfo: make([]RegionInfo, 0),
	}
	req, err := cl.getRequestBuilder("/")
	if err != nil {
		return regions, err
	}
	resp, err := cl.intClient.Do(req)
	if err != nil {
		return regions, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&regions); err != nil {
		return regions, err
	}
	return regions, nil
}

// Read data for the region. This is not a good API for golang since it's purely dynamic. If you want to use it, you have
// to decode it, heh. Author's recommendation: don't use this API.
func (cl Client) GetRegion(limit int, all bool) ([]byte, error) {
	// TODO (mxplusb): finish this API interface...
	req, err := cl.getRequestBuilder("/" + cl.Region)
	if err != nil {
		return []byte(""), nil
	}
	resp, err := cl.intClient.Do(req)
	if err != nil {
		return []byte(""), err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}
	return body, nil
}

// Used to hold the list of keys.
type KeyList struct {
	Keys []string `json:"keys"`
}

func (cl *Client) resetRegion(region string) {
	cl.Region = region
}

// Gets a list of keys for a given region. Uses the configured region unless specified.
func (cl Client) GetKeys(region string) (KeyList, error) {
	switch {
	case region == cl.Region:
	case region == "":
	case region != cl.Region:
		defer cl.resetRegion(cl.Region)
		cl.Region = region
	}

	var keys KeyList
	req, err := cl.getRequestBuilder(fmt.Sprintf("/%s/keys", cl.Region))
	if err != nil {
		return KeyList{}, err
	}
	resp, err := cl.intClient.Do(req)
	if err != nil {
		return keys, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
	case 404:
		return keys, errors.New("specified region does not exist")
	case 500:
		return keys, errors.New(fmt.Sprintf("error encountered at GemFire server. check the HTTP response body for a stack trace of the exception. error: %s", err))
	}

	if err := json.NewDecoder(resp.Body).Decode(&keys); err != nil {
		return KeyList{}, err
	}
	return keys, nil
}
