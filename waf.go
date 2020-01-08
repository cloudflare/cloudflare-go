package cloudflare

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/pkg/errors"
)

// WAFPackage represents a WAF package configuration.
type WAFPackage struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	ZoneID        string `json:"zone_id"`
	DetectionMode string `json:"detection_mode"`
	Sensitivity   string `json:"sensitivity"`
	ActionMode    string `json:"action_mode"`
}

// WAFPackagesResponse represents the response from the WAF packages endpoint.
type WAFPackagesResponse struct {
	Response
	Result     []WAFPackage `json:"result"`
	ResultInfo ResultInfo   `json:"result_info"`
}

// WAFPackageResponse represents the response from the WAF package endpoint.
type WAFPackageResponse struct {
	Response
	Result     WAFPackage `json:"result"`
	ResultInfo ResultInfo `json:"result_info"`
}

// WAFPackageOptions represents options to edit a WAF package.
type WAFPackageOptions struct {
	Sensitivity string `json:"sensitivity,omitempty"`
	ActionMode  string `json:"action_mode,omitempty"`
}

// WAFGroup represents a WAF rule group.
type WAFGroup struct {
	ID                 string   `json:"id"`
	Name               string   `json:"name"`
	Description        string   `json:"description"`
	RulesCount         int      `json:"rules_count"`
	ModifiedRulesCount int      `json:"modified_rules_count"`
	PackageID          string   `json:"package_id"`
	Mode               string   `json:"mode"`
	AllowedModes       []string `json:"allowed_modes"`
}

// WAFGroupsResponse represents the response from the WAF groups endpoint.
type WAFGroupsResponse struct {
	Response
	Result     []WAFGroup `json:"result"`
	ResultInfo ResultInfo `json:"result_info"`
}

// WAFGroupResponse represents the response from the WAF group endpoint.
type WAFGroupResponse struct {
	Response
	Result     WAFGroup   `json:"result"`
	ResultInfo ResultInfo `json:"result_info"`
}

// WAFRule represents a WAF rule.
type WAFRule struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
	PackageID   string `json:"package_id"`
	Group       struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"group"`
	Mode         string   `json:"mode"`
	DefaultMode  string   `json:"default_mode"`
	AllowedModes []string `json:"allowed_modes"`
}

// WAFRulesResponse represents the response from the WAF rules endpoint.
type WAFRulesResponse struct {
	Response
	Result     []WAFRule  `json:"result"`
	ResultInfo ResultInfo `json:"result_info"`
}

// WAFRuleResponse represents the response from the WAF rule endpoint.
type WAFRuleResponse struct {
	Response
	Result     WAFRule    `json:"result"`
	ResultInfo ResultInfo `json:"result_info"`
}

// WAFRuleOptions is a subset of WAFRule, for editable options.
type WAFRuleOptions struct {
	Mode string `json:"mode"`
}

// WAFOverridesResponse represents the response form the WAF overrides endpoint.
type WAFOverridesResponse struct {
	Response
	Result     []WAFOverride `json:"result"`
	ResultInfo ResultInfo    `json:"result_info"`
}

// WAFOverrideResponse represents the response form the WAF override endpoint.
type WAFOverrideResponse struct {
	Response
	Result     WAFOverride `json:"result"`
	ResultInfo ResultInfo  `json:"result_info"`
}

// WAFOverride represents a WAF override.
type WAFOverride struct {
	ID             string            `json:"id,omitempty"`
	Description    string            `json:"description"`
	URLs           []string          `json:"urls"`
	Priority       int               `json:"priority"`
	Groups         map[string]string `json:"groups"`
	RewriteActions map[string]string `json:"rewrite_actions"`
	Rules          map[string]string `json:"rules"`
	Paused         bool              `json:"paused"`
}

// ListWAFPackages returns a slice of the WAF packages for the given zone.
//
// API Reference: https://api.cloudflare.com/#waf-rule-packages-list-firewall-packages
func (api *API) ListWAFPackages(zoneID string) ([]WAFPackage, error) {
	// Construct a query string
	v := url.Values{}
	// Request as many WAF packages as possible per page - API max is 100
	v.Set("per_page", "100")

	var packages []WAFPackage
	var res []byte
	var err error
	page := 1

	// Loop over makeRequest until what we've fetched all records
	for {
		v.Set("page", strconv.Itoa(page))
		query := "?" + v.Encode()
		uri := "/zones/" + zoneID + "/firewall/waf/packages" + query
		res, err = api.makeRequest("GET", uri, nil)
		if err != nil {
			return []WAFPackage{}, errors.Wrap(err, errMakeRequestError)
		}

		var p WAFPackagesResponse
		err = json.Unmarshal(res, &p)
		if err != nil {
			return []WAFPackage{}, errors.Wrap(err, errUnmarshalError)
		}

		if !p.Success {
			// TODO: Provide an actual error message instead of always returning nil
			return []WAFPackage{}, err
		}

		packages = append(packages, p.Result...)
		if p.ResultInfo.Page >= p.ResultInfo.TotalPages {
			break
		}

		// Loop around and fetch the next page
		page++
	}

	return packages, nil
}

// WAFPackage returns a WAF package for the given zone.
//
// API Reference: https://api.cloudflare.com/#waf-rule-packages-firewall-package-details
func (api *API) WAFPackage(zoneID, packageID string) (WAFPackage, error) {
	uri := "/zones/" + zoneID + "/firewall/waf/packages/" + packageID
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return WAFPackage{}, errors.Wrap(err, errMakeRequestError)
	}

	var r WAFPackageResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WAFPackage{}, errors.Wrap(err, errUnmarshalError)
	}

	return r.Result, nil
}

// UpdateWAFPackage lets you update the a WAF Package.
//
// API Reference: https://api.cloudflare.com/#waf-rule-packages-edit-firewall-package
func (api *API) UpdateWAFPackage(zoneID, packageID string, opts WAFPackageOptions) (WAFPackage, error) {
	uri := "/zones/" + zoneID + "/firewall/waf/packages/" + packageID
	res, err := api.makeRequest("PATCH", uri, opts)
	if err != nil {
		return WAFPackage{}, errors.Wrap(err, errMakeRequestError)
	}

	var r WAFPackageResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WAFPackage{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// ListWAFGroups returns a slice of the WAF groups for the given WAF package.
//
// API Reference: https://api.cloudflare.com/#waf-rule-groups-list-rule-groups
func (api *API) ListWAFGroups(zoneID, packageID string) ([]WAFGroup, error) {
	// Construct a query string
	v := url.Values{}
	// Request as many WAF groups as possible per page - API max is 100
	v.Set("per_page", "100")

	var groups []WAFGroup
	var res []byte
	var err error
	page := 1

	// Loop over makeRequest until what we've fetched all records
	for {
		v.Set("page", strconv.Itoa(page))
		query := "?" + v.Encode()
		uri := "/zones/" + zoneID + "/firewall/waf/packages/" + packageID + "/groups" + query
		res, err = api.makeRequest("GET", uri, nil)
		if err != nil {
			return []WAFGroup{}, errors.Wrap(err, errMakeRequestError)
		}

		var r WAFGroupsResponse
		err = json.Unmarshal(res, &r)
		if err != nil {
			return []WAFGroup{}, errors.Wrap(err, errUnmarshalError)
		}

		if !r.Success {
			// TODO: Provide an actual error message instead of always returning nil
			return []WAFGroup{}, err
		}

		groups = append(groups, r.Result...)
		if r.ResultInfo.Page >= r.ResultInfo.TotalPages {
			break
		}

		// Loop around and fetch the next page
		page++
	}
	return groups, nil
}

// WAFGroup returns a WAF rule group from the given WAF package.
//
// API Reference: https://api.cloudflare.com/#waf-rule-groups-rule-group-details
func (api *API) WAFGroup(zoneID, packageID, groupID string) (WAFGroup, error) {
	uri := "/zones/" + zoneID + "/firewall/waf/packages/" + packageID + "/groups/" + groupID
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return WAFGroup{}, errors.Wrap(err, errMakeRequestError)
	}

	var r WAFGroupResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WAFGroup{}, errors.Wrap(err, errUnmarshalError)
	}

	return r.Result, nil
}

// UpdateWAFGroup lets you update the mode of a WAF Group.
//
// API Reference: https://api.cloudflare.com/#waf-rule-groups-edit-rule-group
func (api *API) UpdateWAFGroup(zoneID, packageID, groupID, mode string) (WAFGroup, error) {
	opts := WAFRuleOptions{Mode: mode}
	uri := "/zones/" + zoneID + "/firewall/waf/packages/" + packageID + "/groups/" + groupID
	res, err := api.makeRequest("PATCH", uri, opts)
	if err != nil {
		return WAFGroup{}, errors.Wrap(err, errMakeRequestError)
	}

	var r WAFGroupResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WAFGroup{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// ListWAFRules returns a slice of the WAF rules for the given WAF package.
//
// API Reference: https://api.cloudflare.com/#waf-rules-list-rules
func (api *API) ListWAFRules(zoneID, packageID string) ([]WAFRule, error) {
	// Construct a query string
	v := url.Values{}
	// Request as many WAF rules as possible per page - API max is 100
	v.Set("per_page", "100")

	var rules []WAFRule
	var res []byte
	var err error
	page := 1

	// Loop over makeRequest until what we've fetched all records
	for {
		v.Set("page", strconv.Itoa(page))
		query := "?" + v.Encode()
		uri := "/zones/" + zoneID + "/firewall/waf/packages/" + packageID + "/rules" + query
		res, err = api.makeRequest("GET", uri, nil)
		if err != nil {
			return []WAFRule{}, errors.Wrap(err, errMakeRequestError)
		}

		var r WAFRulesResponse
		err = json.Unmarshal(res, &r)
		if err != nil {
			return []WAFRule{}, errors.Wrap(err, errUnmarshalError)
		}

		if !r.Success {
			// TODO: Provide an actual error message instead of always returning nil
			return []WAFRule{}, err
		}

		rules = append(rules, r.Result...)
		if r.ResultInfo.Page >= r.ResultInfo.TotalPages {
			break
		}

		// Loop around and fetch the next page
		page++
	}

	return rules, nil
}

// WAFRule returns a WAF rule from the given WAF package.
//
// API Reference: https://api.cloudflare.com/#waf-rules-rule-details
func (api *API) WAFRule(zoneID, packageID, ruleID string) (WAFRule, error) {
	uri := "/zones/" + zoneID + "/firewall/waf/packages/" + packageID + "/rules/" + ruleID
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return WAFRule{}, errors.Wrap(err, errMakeRequestError)
	}

	var r WAFRuleResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WAFRule{}, errors.Wrap(err, errUnmarshalError)
	}

	return r.Result, nil
}

// UpdateWAFRule lets you update the mode of a WAF Rule.
//
// API Reference: https://api.cloudflare.com/#waf-rules-edit-rule
func (api *API) UpdateWAFRule(zoneID, packageID, ruleID, mode string) (WAFRule, error) {
	opts := WAFRuleOptions{Mode: mode}
	uri := "/zones/" + zoneID + "/firewall/waf/packages/" + packageID + "/rules/" + ruleID
	res, err := api.makeRequest("PATCH", uri, opts)
	if err != nil {
		return WAFRule{}, errors.Wrap(err, errMakeRequestError)
	}

	var r WAFRuleResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WAFRule{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// ListWAFOverrides returns a slice of the WAF overrides.
//
// API Reference: https://api.cloudflare.com/#waf-overrides-list-uri-controlled-waf-configurations
func (api *API) ListWAFOverrides(zoneID string) ([]WAFOverride, error) {
	var overrides []WAFOverride
	var res []byte
	var err error

	uri := "/zones/" + zoneID + "/firewall/waf/overrides"
	res, err = api.makeRequest("GET", uri, nil)
	if err != nil {
		return []WAFOverride{}, errors.Wrap(err, errMakeRequestError)
	}

	var r WAFOverridesResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return []WAFOverride{}, errors.Wrap(err, errUnmarshalError)
	}

	if !r.Success {
		// TODO: Provide an actual error message instead of always returning nil
		return []WAFOverride{}, err
	}

	for ri := range r.Result {
		overrides = append(overrides, r.Result[ri])
	}
	return overrides, nil
}

// WAFOverride returns a WAF override from the given override ID.
//
// API Reference: https://api.cloudflare.com/#waf-overrides-update-uri-controlled-waf-configuration
func (api *API) WAFOverride(zoneID, overrideID string) (WAFOverride, error) {
	uri := "/zones/" + zoneID + "/firewall/waf/overrides/" + overrideID
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return WAFOverride{}, errors.Wrap(err, errMakeRequestError)
	}

	var r WAFOverrideResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WAFOverride{}, errors.Wrap(err, errUnmarshalError)
	}

	return r.Result, nil
}

// CreateWAFOverride creates a new WAF override.
//
// API reference: https://api.cloudflare.com/#waf-overrides-create-a-uri-controlled-waf-configuration
func (api *API) CreateWAFOverride(zoneID string, override WAFOverride) (WAFOverride, error) {
	uri := "/zones/" + zoneID + "/firewall/waf/overrides"
	res, err := api.makeRequest("POST", uri, override)
	if err != nil {
		return WAFOverride{}, errors.Wrap(err, errMakeRequestError)
	}
	var r WAFOverrideResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return WAFOverride{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// UpdateWAFOverride updates an existing WAF override.
//
// API reference: https://api.cloudflare.com/#waf-overrides-update-uri-controlled-waf-configuration
func (api *API) UpdateWAFOverride(zoneID, overrideID string, override WAFOverride) (WAFOverride, error) {
	uri := "/zones/" + zoneID + "/firewall/waf/overrides/" + overrideID

	res, err := api.makeRequest("PUT", uri, override)
	if err != nil {
		return WAFOverride{}, errors.Wrap(err, errMakeRequestError)
	}

	var r WAFOverrideResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WAFOverride{}, errors.Wrap(err, errUnmarshalError)
	}

	return r.Result, nil
}

// DeleteWAFOverride deletes a WAF override for a zone.
//
// API reference: https://api.cloudflare.com/#waf-overrides-delete-lockdown-rule
func (api *API) DeleteWAFOverride(zoneID, overrideID string) error {
	uri := "/zones/" + zoneID + "/firewall/waf/overrides/" + overrideID
	res, err := api.makeRequest("DELETE", uri, nil)
	if err != nil {
		return errors.Wrap(err, errMakeRequestError)
	}
	var r WAFOverrideResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return errors.Wrap(err, errUnmarshalError)
	}
	return nil
}
