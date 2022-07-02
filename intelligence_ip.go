package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type IPIntelligence struct {
	IP           string       `json:"ip"`
	BelongsToRef BelongsToRef `json:"belongs_to_ref"`
	RiskTypes    []RiskTypes  `json:"risk_types"`
}
type BelongsToRef struct {
	ID          string `json:"id"`
	Value       int    `json:"value"`
	Type        string `json:"type"`
	Country     string `json:"country"`
	Description string `json:"description"`
}
type RiskTypes struct {
	ID              int    `json:"id"`
	SuperCategoryID int    `json:"super_category_id"`
	Name            string `json:"name"`
}

type IPPassiveDNS struct {
	ReverseRecords []ReverseRecords `json:"reverse_records,omitempty"`
	Count          int              `json:"count,omitempty"`
	Page           int              `json:"page,omitempty"`
	PerPage        int              `json:"per_page,omitempty"`
}

type ReverseRecords struct {
	FirstSeen string `json:"first_seen"`
	LastSeen  string `json:"last_seen"`
	Hostname  string `json:"hostname"`
}

type IPIntelligenceParameters struct {
	AccountID string `url:"-"`
	IPv4      string `url:"ipv4,omitempty"`
	IPv6      string `url:"ipv6,omitempty"`
}

type IPIntelligenceResponse struct {
	Response
	Result []IPIntelligence `json:"result,omitempty"`
}

type IPIntelligenceListParameters struct {
	AccountID string
}

type IPIntelligenceItem struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type IPIntelligenceListResponse struct {
	Response
	Result []IPIntelligenceItem `json:"result,omitempty"`
}

type IPIntelligencePassiveDNSParameters struct {
	AccountID string     `url:"-"`
	IPv4      string     `url:"ipv4,omitempty"`
	Start     *time.Time `url:"start,omitempty"`
	End       *time.Time `url:"end,omitempty"`
	Page      int        `url:"page,omitempty"`
	PerPage   int        `url:"per_page,omitempty"`
}

type IPIntelligencePassiveDNSResponse struct {
	Response
	Result IPPassiveDNS `json:"result,omitempty"`
}

// IntelligenceGetIPOverview gets information about ipv4 or ipv6 address
//
// API Reference: https://api.cloudflare.com/#ip-intelligence-get-ip-overview
func (api *API) IntelligenceGetIPOverview(ctx context.Context, params IPIntelligenceParameters) ([]IPIntelligence, error) {
	// Make sure Account ID is set
	if params.AccountID == "" {
		return []IPIntelligence{}, ErrMissingAccountID
	}

	uri := buildURI(fmt.Sprintf("/accounts/%s/intel/ip", params.AccountID), params)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []IPIntelligence{}, err
	}
	var ipDetails IPIntelligenceResponse
	if err := json.Unmarshal(res, &ipDetails); err != nil {
		return []IPIntelligence{}, err
	}
	return ipDetails.Result, nil
}

// IntelligenceGetIPList gets intelligence ip-lists
//
// API Reference: https://api.cloudflare.com/#ip-list-get-ip-lists
func (api *API) IntelligenceGetIPList(ctx context.Context, params IPIntelligenceListParameters) ([]IPIntelligenceItem, error) {
	// Make sure Account ID is set
	if params.AccountID == "" {
		return []IPIntelligenceItem{}, ErrMissingAccountID
	}
	uri := fmt.Sprintf("/accounts/%s/intel/ip-list", params.AccountID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []IPIntelligenceItem{}, err
	}

	var ipListItem IPIntelligenceListResponse
	if err := json.Unmarshal(res, &ipListItem); err != nil {
		return []IPIntelligenceItem{}, err
	}
	return ipListItem.Result, nil
}

// IntelligencePassiveDNS gets a history of DNS for an ip
//
// API Reference: https://api.cloudflare.com/#passive-dns-by-ip-get-passive-dns-by-ip
func (api *API) IntelligencePassiveDNS(ctx context.Context, params IPIntelligencePassiveDNSParameters) (IPPassiveDNS, error) {
	// Make sure Account ID is set
	if params.AccountID == "" {
		return IPPassiveDNS{}, ErrMissingAccountID
	}

	uri := buildURI(fmt.Sprintf("/accounts/%s/intel/dns", params.AccountID), params)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return IPPassiveDNS{}, err
	}
	var passiveDNS IPIntelligencePassiveDNSResponse
	if err := json.Unmarshal(res, &passiveDNS); err != nil {
		return IPPassiveDNS{}, err
	}
	return passiveDNS.Result, nil
}
