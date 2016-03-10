package cloudflare

import (
	"encoding/json"
)

var api *API

func (api *API) ListWafPackages(zoneId string) ([]WafPackage, error) {
	var p WafPackagesResponse
	var packages []WafPackage
	var res []byte
	var err error
	res, err = api.makeRequest("GET", "/zones/"+zoneId+"/firewall/waf/packages", nil)
	if err != nil {
		return []WafPackage{}, err
	}
	err = json.Unmarshal(res, &p)
	if err != nil {
		return []WafPackage{}, err
	}
	if !p.Success {
		return []WafPackage{}, err
	}
	for pi, _ := range p.Result {
		packages = append(packages, p.Result[pi])
	}
	return packages, err
}
