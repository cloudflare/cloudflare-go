package cloudflare

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/pkg/errors"
)

// AccountSettings outlines the available options for an account.
type AccountSettings struct {
	EnforceTwoFactor bool `json:"enforce_twofactor"`
}

// Account represents the root object that owns resources.
type Account struct {
	ID       string          `json:"id,omitempty"`
	Name     string          `json:"name,omitempty"`
	Settings AccountSettings `json:"settings"`
}

// AccountResponse represents the response from the accounts endpoint for a
// single account ID.
type AccountResponse struct {
	Result Account `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

// AccountListResponse represents the response from the list accounts endpoint.
type AccountListResponse struct {
	Result []Account `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

// ListAccounts returns all accounts the logged in user has access to.
//
// API reference: https://api.cloudflare.com/#accounts-list-accounts
func (api *API) ListAccounts(pageOpts PaginationOptions) ([]Account, ResultInfo, error) {
	v := url.Values{}
	if pageOpts.PerPage > 0 {
		v.Set("per_page", strconv.Itoa(pageOpts.PerPage))
	}
	if pageOpts.Page > 0 {
		v.Set("page", strconv.Itoa(pageOpts.Page))
	}

	uri := "/accounts"
	if len(v) > 0 {
		uri = uri + "?" + v.Encode()
	}

	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return []Account{}, ResultInfo{}, errors.Wrap(err, errMakeRequestError)
	}

	var accListResponse AccountListResponse
	err = json.Unmarshal(res, &accListResponse)
	if err != nil {
		return []Account{}, ResultInfo{}, errors.Wrap(err, errUnmarshalError)
	}
	return accListResponse.Result, accListResponse.ResultInfo, nil
}

// Account returns a single account based on the ID.Account
//
// API reference: https://api.cloudflare.com/#accounts-account-details
func (api *API) Account(accountID string) (Account, ResultInfo, error) {
	uri := "/accounts/" + accountID

	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return Account{}, ResultInfo{}, errors.Wrap(err, errMakeRequestError)
	}

	var accResponse AccountResponse
	err = json.Unmarshal(res, &accResponse)
	if err != nil {
		return Account{}, ResultInfo{}, errors.Wrap(err, errUnmarshalError)
	}

	return accResponse.Result, accResponse.ResultInfo, nil
}
