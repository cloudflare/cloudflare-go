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

// AccessUserFailedLoginService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccessUserFailedLoginService]
// method instead.
type AccessUserFailedLoginService struct {
	Options []option.RequestOption
}

// NewAccessUserFailedLoginService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccessUserFailedLoginService(opts ...option.RequestOption) (r *AccessUserFailedLoginService) {
	r = &AccessUserFailedLoginService{}
	r.Options = opts
	return
}

// Get all failed login attempts for a single user.
func (r *AccessUserFailedLoginService) ZeroTrustUsersGetFailedLogins(ctx context.Context, identifier string, id string, opts ...option.RequestOption) (res *AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/users/%s/failed_logins", identifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponse struct {
	Errors     []AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseError    `json:"errors"`
	Messages   []AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseMessage  `json:"messages"`
	Result     []AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseResult   `json:"result"`
	ResultInfo AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseSuccess `json:"success"`
	JSON    accessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseJSON    `json:"-"`
}

// accessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseJSON contains the JSON
// metadata for the struct
// [AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponse]
type accessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseError struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    accessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseErrorJSON `json:"-"`
}

// accessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseErrorJSON contains the
// JSON metadata for the struct
// [AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseError]
type accessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseMessage struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    accessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseMessageJSON `json:"-"`
}

// accessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseMessageJSON contains
// the JSON metadata for the struct
// [AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseMessage]
type accessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseResult struct {
	Expiration int64                                                                `json:"expiration"`
	Metadata   interface{}                                                          `json:"metadata"`
	JSON       accessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseResultJSON `json:"-"`
}

// accessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseResultJSON contains
// the JSON metadata for the struct
// [AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseResult]
type accessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseResultJSON struct {
	Expiration  apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                  `json:"total_count"`
	JSON       accessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseResultInfoJSON `json:"-"`
}

// accessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseResultInfo]
type accessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseSuccess bool

const (
	AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseSuccessTrue AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseSuccess = true
)
