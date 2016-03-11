package cloudflare

import (
	"encoding/json"
	"errors"
	"net/url"
)

func (api *API) ListWAFPackages(zoneID string) ([]WAFPackage, error) {
	var p WAFPackagesResponse
	var packages []WAFPackage
	var res []byte
	var err error
	res, err = api.makeRequest("GET", "/zones/"+zoneID+"/firewall/waf/packages", nil)
	if err != nil {
		return []WAFPackage{}, err
	}
	err = json.Unmarshal(res, &p)
	if err != nil {
		return []WAFPackage{}, err
	}
	if !p.Success {
		return []WAFPackage{}, err
	}
	for pi, _ := range p.Result {
		packages = append(packages, p.Result[pi])
	}
	return packages, err
}

func (api *API) ListWAFRules(zoneID string, packageID string, v ...url.Values) ([]WAFRule, error) {
	var query url.Values
	var r WAFRulesResponse
	var rules []WAFRule
	var res []byte
	var err error
	if len(v) > 1 {
		return []WAFRule{}, errors.New("Too many arguments")
	} else if len(v) == 0 {
		query = url.Values{}
	} else {
		query = v[0]
	}
	res, err = api.makeRequest("GET", "/zones/"+zoneID+"/firewall/waf/packages/"+packageID+"/rules?"+query.Encode(), nil)
	if err != nil {
		return []WAFRule{}, err
	}
	err = json.Unmarshal(res, &r)
	if err != nil {
		return []WAFRule{}, err
	}
	if !r.Success {
		return []WAFRule{}, err
	}
	for ri, _ := range r.Result {
		rules = append(rules, r.Result[ri])
	}
	return rules, err
}
