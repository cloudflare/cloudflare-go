package cloudflare

import (
	"encoding/json"
)

var api *API

func (api *API) ListWafPackages(zoneID string) ([]WafPackage, error) {
	var p WafPackagesResponse
	var packages []WafPackage
	var res []byte
	var err error
	res, err = api.makeRequest("GET", "/zones/"+zoneID+"/firewall/waf/packages", nil)
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

func (api *API) ListWafRules(zoneID string, packageID string) ([]WafRule, error) {
	var r WafRulesResponse
	var rules []WafRule
	var res []byte
	var err error
	res, err = api.makeRequest("GET", "/zones/"+zoneID+"/firewall/waf/packages/"+packageID+"/rules", nil)
	if err != nil {
		return []WafRule{}, err
	}
	err = json.Unmarshal(res, &r)
	if err != nil {
		return []WafRule{}, err
	}
	if !r.Success {
		return []WafRule{}, err
	}
	for ri, _ := range r.Result {
		rules = append(rules, r.Result[ri])
	}
	return rules, err
}
