package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type WHOIS struct {
	Domain            string   `json:"domain,omitempty"`
	CreatedDate       string   `json:"created_date,omitempty"`
	UpdatedDate       string   `json:"updated_date,omitempty"`
	Registrant        string   `json:"registrant,omitempty"`
	RegistrantOrg     string   `json:"registrant_org,omitempty"`
	RegistrantCountry string   `json:"registrant_country,omitempty"`
	RegistrantEmail   string   `json:"registrant_email,omitempty"`
	Registrar         string   `json:"registrar,omitempty"`
	Nameservers       []string `json:"nameservers,omitempty"`
}

type WHOISParameters struct {
	AccountID string `url:"-"`
	Domain    string `url:"domain"`
}

type WHOISResponse struct {
	Response
	Result WHOIS `json:"result,omitempty"`
}

// IntelligenceWHOIS gets whois information for a domain.
//
// API Reference: https://api.cloudflare.com/#whois-record-get-whois-record
func (api *API) IntelligenceWHOIS(ctx context.Context, params WHOISParameters) (WHOIS, error) {
	// Make sure Account ID is set
	if params.AccountID == "" {
		return WHOIS{}, ErrMissingAccountID
	}

	// Make sure domain  is set
	if params.Domain == "" {
		return WHOIS{}, ErrMissingDomain
	}

	uri := buildURI(fmt.Sprintf("/accounts/%s/intel/whois", params.AccountID), params)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return WHOIS{}, err
	}
	var whoisResponse WHOISResponse
	if err := json.Unmarshal(res, &whoisResponse); err != nil {
		return WHOIS{}, err
	}

	return whoisResponse.Result, nil
}
