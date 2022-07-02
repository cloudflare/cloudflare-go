package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

var (
	// ErrMissingDomain is for when domain is needed but not given
	ErrMissingDomain = errors.New("required domain missing")
)

type DomainDetails struct {
	Domain                string                `json:"domain"`
	ResolvesToRefs        []ResolvesToRefs      `json:"resolves_to_refs"`
	PopularityRank        int                   `json:"popularity_rank"`
	Application           Application           `json:"application"`
	RiskTypes             []interface{}         `json:"risk_types"`
	ContentCategories     []ContentCategories   `json:"content_categories"`
	AdditionalInformation AdditionalInformation `json:"additional_information"`
}
type ResolvesToRefs struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}
type Application struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type ContentCategories struct {
	ID              int    `json:"id"`
	SuperCategoryID int    `json:"super_category_id"`
	Name            string `json:"name"`
}
type AdditionalInformation struct {
	SuspectedMalwareFamily string `json:"suspected_malware_family"`
}

type DomainHistory struct {
	Domain          string            `json:"domain"`
	Categorizations []Categorizations `json:"categorizations"`
}
type Categories struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type Categorizations struct {
	Categories []Categories `json:"categories"`
	Start      string       `json:"start"`
	End        string       `json:"end"`
}
type GetDomainDetailsParameters struct {
	AccountID string `url:"-"`
	Domain    string `url:"domain,omitempty"`
}

type DomainDetailsResponse struct {
	Response
	Result DomainDetails `json:"result,omitempty"`
}

type GetBulkDomainDetailsParameters struct {
	AccountID string   `url:"-"`
	Domains   []string `url:"domain"`
}

type GetBulkDomainDetailsResponse struct {
	Response
	Result []DomainDetails `json:"result,omitempty"`
}

type GetDomainHistoryParameters struct {
	AccountID string `url:"-"`
	Domain    string `url:"domain,omitempty"`
}

type GetDomainHistoryResponse struct {
	Response
	Result []DomainHistory `json:"result,omitempty"`
}

// IntelligenceDomainDetails gets domain information.
//
// API Reference: https://api.cloudflare.com/#domain-intelligence-get-domain-details
func (api *API) IntelligenceDomainDetails(ctx context.Context, params GetDomainDetailsParameters) (DomainDetails, error) {
	// Make sure Account ID is set
	if params.AccountID == "" {
		return DomainDetails{}, ErrMissingAccountID
	}
	// Make sure Domain is set
	if params.Domain == "" {
		return DomainDetails{}, ErrMissingDomain
	}
	uri := buildURI(fmt.Sprintf("/accounts/%s/intel/domain", params.AccountID), params)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return DomainDetails{}, err
	}
	var domainDetails DomainDetailsResponse
	if err := json.Unmarshal(res, &domainDetails); err != nil {
		return DomainDetails{}, err
	}
	return domainDetails.Result, nil
}

// IntelligenceBulkDomainDetails gets domain information for a list of domains.
//
// API Reference: https://api.cloudflare.com/#domain-intelligence-get-multiple-domain-details
func (api *API) IntelligenceBulkDomainDetails(ctx context.Context, params GetBulkDomainDetailsParameters) ([]DomainDetails, error) {
	// Make sure Account ID is set
	if params.AccountID == "" {
		return []DomainDetails{}, ErrMissingAccountID
	}
	// Make sure domains are set
	if len(params.Domains) == 0 {
		return []DomainDetails{}, ErrMissingDomain
	}
	uri := buildURI(fmt.Sprintf("/accounts/%s/intel/domain/bulk", params.AccountID), params)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []DomainDetails{}, err
	}
	var domainDetails GetBulkDomainDetailsResponse
	if err := json.Unmarshal(res, &domainDetails); err != nil {
		return []DomainDetails{}, err
	}
	return domainDetails.Result, nil
}

// IntelligenceDomainHistory get domain history for given domain
//
// API Reference: https://api.cloudflare.com/#domain-history-get-domain-history
func (api *API) IntelligenceDomainHistory(ctx context.Context, params GetDomainHistoryParameters) ([]DomainHistory, error) {
	// Make sure Account ID is set
	if params.AccountID == "" {
		return []DomainHistory{}, ErrMissingAccountID
	}
	// Make sure Domain is set
	if params.Domain == "" {
		return []DomainHistory{}, ErrMissingDomain
	}
	uri := buildURI(fmt.Sprintf("/accounts/%s/intel/domain-history", params.AccountID), params)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []DomainHistory{}, err
	}
	var domainDetails GetDomainHistoryResponse
	if err := json.Unmarshal(res, &domainDetails); err != nil {
		return []DomainHistory{}, err
	}
	return domainDetails.Result, nil
}
