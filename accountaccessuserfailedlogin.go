// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountAccessUserFailedLoginService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountAccessUserFailedLoginService] method instead.
type AccountAccessUserFailedLoginService struct {
	Options []option.RequestOption
}

// NewAccountAccessUserFailedLoginService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAccessUserFailedLoginService(opts ...option.RequestOption) (r *AccountAccessUserFailedLoginService) {
	r = &AccountAccessUserFailedLoginService{}
	r.Options = opts
	return
}

// Get all failed login attempts for a single user.
func (r *AccountAccessUserFailedLoginService) ZeroTrustUsersGetFailedLogins(ctx context.Context, identifier string, id string, opts ...option.RequestOption) (res *AccountAccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/users/%s/failed_logins", identifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountAccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponse struct {
	Errors     []AccountAccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseError    `json:"errors"`
	Messages   []AccountAccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseMessage  `json:"messages"`
	Result     []AccountAccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseResult   `json:"result"`
	ResultInfo AccountAccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountAccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseSuccess `json:"success"`
	JSON    accountAccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseJSON    `json:"-"`
}

// accountAccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseJSON contains
// the JSON metadata for the struct
// [AccountAccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponse]
type accountAccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseError struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    accountAccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseErrorJSON `json:"-"`
}

// accountAccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountAccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseError]
type accountAccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseMessage struct {
	Code    int64                                                                        `json:"code,required"`
	Message string                                                                       `json:"message,required"`
	JSON    accountAccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseMessageJSON `json:"-"`
}

// accountAccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountAccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseMessage]
type accountAccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseResult struct {
	Expiration int64                                                                       `json:"expiration"`
	Metadata   interface{}                                                                 `json:"metadata"`
	JSON       accountAccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseResultJSON `json:"-"`
}

// accountAccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseResultJSON
// contains the JSON metadata for the struct
// [AccountAccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseResult]
type accountAccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseResultJSON struct {
	Expiration  apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                         `json:"total_count"`
	JSON       accountAccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseResultInfoJSON `json:"-"`
}

// accountAccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountAccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseResultInfo]
type accountAccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseSuccess bool

const (
	AccountAccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseSuccessTrue AccountAccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseSuccess = true
)
