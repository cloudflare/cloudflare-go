package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)
// SecondaryZoneDNSTSIG contains the structure for a secondary DNS zone TSIG.
type SecondaryZoneDNSTSIG struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name"`
	Secret string `json:"secret"`
	Algo   string `json:"algo"`
}

// SecondaryZoneDNSTSIGDetailResponse is the API response for a single secondary
// DNS zone TSIG.
type SecondaryZoneDNSTSIGDetailResponse struct {
	Response
	Result SecondaryZoneDNSTSIG `json:"result"`
}

// SecondaryZoneDNSTSIGListResponse is the API response for all secondary DNS
// zone TSIGs.
type SecondaryZoneDNSTSIGListResponse struct {
	Response
	Result []SecondaryZoneDNSTSIG `json:"result"`
}

// GetSecondaryDNSTSIG gets a single account level TSIG for a secondary DNS
// configuration.
//
// API reference: https://api.cloudflare.com/#secondary-dns-tsig--tsig-details
func (api *API) GetSecondaryDNSTSIG(ctx context.Context, accountID, tsigID string) (SecondaryZoneDNSTSIG, error) {
	uri := fmt.Sprintf("/accounts/%s/secondary_dns/tsigs/%s", accountID, tsigID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return SecondaryZoneDNSTSIG{}, err
	}

	var r SecondaryZoneDNSTSIGDetailResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return SecondaryZoneDNSTSIG{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// SecondaryDNSTSIGs gets all account level TSIG for a secondary DNS
// configuration.
//
// API reference: https://api.cloudflare.com/#secondary-dns-tsig--list-tsigs
func (api *API) SecondaryDNSTSIGs(ctx context.Context, accountID string) ([]SecondaryZoneDNSTSIG, error) {
	uri := fmt.Sprintf("/accounts/%s/secondary_dns/tsigs", accountID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []SecondaryZoneDNSTSIG{}, err
	}

	var r SecondaryZoneDNSTSIGListResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return []SecondaryZoneDNSTSIG{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// CreateSecondaryDNSZoneTSIG creates a secondary DNS zone TSIG at the account
// level.
//
// API reference: https://api.cloudflare.com/#secondary-dns-tsig--create-tsig
func (api *API) CreateSecondaryDNSZoneTSIG(ctx context.Context, accountID string, tsig SecondaryZoneDNSTSIG) (SecondaryZoneDNSTSIG, error) {
	uri := fmt.Sprintf("/accounts/%s/secondary_dns/tsigs", accountID)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri,
		SecondaryZoneDNSTSIG{
			Name:   tsig.Name,
			Secret: tsig.Secret,
			Algo:   tsig.Algo,
		},
	)

	if err != nil {
		return SecondaryZoneDNSTSIG{}, err
	}

	result := SecondaryZoneDNSTSIGDetailResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return SecondaryZoneDNSTSIG{}, errors.Wrap(err, errUnmarshalError)
	}

	return result.Result, nil
}

