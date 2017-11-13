package gofire

import (
	"encoding/json"
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
func (cl Client) GetRegion(region string, limit int, all bool) ([]byte, error) {
	req, err := cl.getRequestBuilder("/" + region)
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

