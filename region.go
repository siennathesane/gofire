package gofire

import "encoding/json"

type RegionInfo struct {
	Name string `json:"name"`
	RegionType string `json:"type"`
	KeyConstraint string `json:"key-constraint"`
	ValueConstraint string `json:"value-constraint"`
}

type Regions struct {
	RegionInfo []RegionInfo `json:"regions"`
}

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
	if err := json.NewDecoder(resp.Body).Decode(&regions); err != nil {
		return regions, err
	}
	return regions, nil
}
