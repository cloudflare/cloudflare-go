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
func (r *AccountAccessUserFailedLoginService) ZeroTrustUsersGetFailedLogins(ctx context.Context, identifier interface{}, id interface{}, opts ...option.RequestOption) (res *FailedLoginResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/access/users/%v/failed_logins", identifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type FailedLoginResponse struct {
	Errors     []FailedLoginResponseError    `json:"errors"`
	Messages   []FailedLoginResponseMessage  `json:"messages"`
	Result     []FailedLoginResponseResult   `json:"result"`
	ResultInfo FailedLoginResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success FailedLoginResponseSuccess `json:"success"`
	JSON    failedLoginResponseJSON    `json:"-"`
}

// failedLoginResponseJSON contains the JSON metadata for the struct
// [FailedLoginResponse]
type failedLoginResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FailedLoginResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FailedLoginResponseError struct {
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    failedLoginResponseErrorJSON `json:"-"`
}

// failedLoginResponseErrorJSON contains the JSON metadata for the struct
// [FailedLoginResponseError]
type failedLoginResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FailedLoginResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FailedLoginResponseMessage struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    failedLoginResponseMessageJSON `json:"-"`
}

// failedLoginResponseMessageJSON contains the JSON metadata for the struct
// [FailedLoginResponseMessage]
type failedLoginResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FailedLoginResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FailedLoginResponseResult struct {
	Expiration int64                         `json:"expiration"`
	Metadata   interface{}                   `json:"metadata"`
	JSON       failedLoginResponseResultJSON `json:"-"`
}

// failedLoginResponseResultJSON contains the JSON metadata for the struct
// [FailedLoginResponseResult]
type failedLoginResponseResultJSON struct {
	Expiration  apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FailedLoginResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FailedLoginResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                           `json:"total_count"`
	JSON       failedLoginResponseResultInfoJSON `json:"-"`
}

// failedLoginResponseResultInfoJSON contains the JSON metadata for the struct
// [FailedLoginResponseResultInfo]
type failedLoginResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FailedLoginResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FailedLoginResponseSuccess bool

const (
	FailedLoginResponseSuccessTrue FailedLoginResponseSuccess = true
)
