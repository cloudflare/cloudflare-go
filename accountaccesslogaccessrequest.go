// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountAccessLogAccessRequestService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountAccessLogAccessRequestService] method instead.
type AccountAccessLogAccessRequestService struct {
	Options []option.RequestOption
}

// NewAccountAccessLogAccessRequestService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAccessLogAccessRequestService(opts ...option.RequestOption) (r *AccountAccessLogAccessRequestService) {
	r = &AccountAccessLogAccessRequestService{}
	r.Options = opts
	return
}

// Gets a list of Access authentication audit logs for an account.
func (r *AccountAccessLogAccessRequestService) AccessAuthenticationLogsGetAccessAuthenticationLogs(ctx context.Context, identifier string, opts ...option.RequestOption) (res *AccountAccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/logs/access_requests", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountAccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponse struct {
	Errors     []AccountAccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseError    `json:"errors"`
	Messages   []AccountAccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseMessage  `json:"messages"`
	Result     []AccountAccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseResult   `json:"result"`
	ResultInfo AccountAccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountAccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseSuccess `json:"success"`
	JSON    accountAccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseJSON    `json:"-"`
}

// accountAccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseJSON
// contains the JSON metadata for the struct
// [AccountAccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponse]
type accountAccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseError struct {
	Code    int64                                                                                             `json:"code,required"`
	Message string                                                                                            `json:"message,required"`
	JSON    accountAccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseErrorJSON `json:"-"`
}

// accountAccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountAccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseError]
type accountAccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseMessage struct {
	Code    int64                                                                                               `json:"code,required"`
	Message string                                                                                              `json:"message,required"`
	JSON    accountAccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseMessageJSON `json:"-"`
}

// accountAccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountAccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseMessage]
type accountAccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseResult struct {
	// The event that occurred, such as a login attempt.
	Action string `json:"action"`
	// The result of the authentication event.
	Allowed bool `json:"allowed"`
	// The URL of the Access application.
	AppDomain string `json:"app_domain"`
	// The unique identifier for the Access application.
	AppUid interface{} `json:"app_uid"`
	// The IdP used to authenticate.
	Connection string    `json:"connection"`
	CreatedAt  time.Time `json:"created_at" format:"date-time"`
	// The IP address of the authenticating user.
	IPAddress string `json:"ip_address"`
	// The unique identifier for the request to Cloudflare.
	RayID string `json:"ray_id"`
	// The email address of the authenticating user.
	UserEmail string                                                                                             `json:"user_email" format:"email"`
	JSON      accountAccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseResultJSON `json:"-"`
}

// accountAccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseResultJSON
// contains the JSON metadata for the struct
// [AccountAccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseResult]
type accountAccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseResultJSON struct {
	Action      apijson.Field
	Allowed     apijson.Field
	AppDomain   apijson.Field
	AppUid      apijson.Field
	Connection  apijson.Field
	CreatedAt   apijson.Field
	IPAddress   apijson.Field
	RayID       apijson.Field
	UserEmail   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                                `json:"total_count"`
	JSON       accountAccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseResultInfoJSON `json:"-"`
}

// accountAccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountAccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseResultInfo]
type accountAccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseSuccess bool

const (
	AccountAccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseSuccessTrue AccountAccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseSuccess = true
)
