package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

// ErrMissingASN is for when ASN is required but not set
var ErrMissingASN = errors.New("required asn missing")

type ASNInfo struct {
	ASN         int      `json:"asn"`
	Description string   `json:"description"`
	Country     string   `json:"country"`
	Type        string   `json:"type"`
	DomainCount int      `json:"domain_count"`
	TopDomains  []string `json:"top_domains"`
}

type IntelligenceASNOverviewParameters struct {
	AccountID string
	ASN       int
}

type IntelligenceASNResponse struct {
	Response
	Result []ASNInfo `json:"result,omitempty"`
}

type IntelligenceASNSubnetsParameters struct {
	AccountID string
	ASN       int
}

type IntelligenceASNSubnetResponse struct {
	ASN          int      `json:"asn,omitempty"`
	IPCountTotal int      `json:"ip_count_total,omitempty"`
	Subnets      []string `json:"subnets,omitempty"`
	Count        int      `json:"count,omitempty"`
	Page         int      `json:"page,omitempty"`
	PerPage      int      `json:"per_page,omitempty"`
}

// IntelligenceASNOverview get overview for an ASN number
//
// API Reference: https://api.cloudflare.com/#asn-intelligence-get-asn-overview
func (api *API) IntelligenceASNOverview(ctx context.Context, params IntelligenceASNOverviewParameters) ([]ASNInfo, error) {
	// Make sure Account ID is set
	if params.AccountID == "" {
		return []ASNInfo{}, ErrMissingAccountID
	}

	// Make sure ASN is set
	if params.ASN == 0 {
		return []ASNInfo{}, ErrMissingASN
	}

	uri := fmt.Sprintf("/accounts/%s/intel/asn/%d", params.AccountID, params.ASN)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []ASNInfo{}, err
	}
	var asnInfoResponse IntelligenceASNResponse
	if err := json.Unmarshal(res, &asnInfoResponse); err != nil {
		return []ASNInfo{}, err
	}
	return asnInfoResponse.Result, nil
}

// IntelligenceASNSubnets gets all subnets of an ASN
//
// API Reference: https://api.cloudflare.com/#asn-intelligence-get-asn-subnets
func (api *API) IntelligenceASNSubnets(ctx context.Context, params IntelligenceASNSubnetsParameters) (IntelligenceASNSubnetResponse, error) {
	// Make sure Account ID is set
	if params.AccountID == "" {
		return IntelligenceASNSubnetResponse{}, ErrMissingAccountID
	}

	// Make sure ASN is set
	if params.ASN == 0 {
		return IntelligenceASNSubnetResponse{}, ErrMissingASN
	}

	uri := fmt.Sprintf("/accounts/%s/intel/asn/%d/subnets", params.AccountID, params.ASN)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return IntelligenceASNSubnetResponse{}, err
	}
	var intelligenceASNSubnetResponse IntelligenceASNSubnetResponse
	if err := json.Unmarshal(res, &intelligenceASNSubnetResponse); err != nil {
		return IntelligenceASNSubnetResponse{}, err
	}
	return intelligenceASNSubnetResponse, nil

}
