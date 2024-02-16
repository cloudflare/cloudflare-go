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
func (r *AccessUserFailedLoginService) ZeroTrustUsersGetFailedLogins(ctx context.Context, identifier string, id string, opts ...option.RequestOption) (res *[]AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/users/%s/failed_logins", identifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponse struct {
	Expiration int64                                                          `json:"expiration"`
	Metadata   interface{}                                                    `json:"metadata"`
	JSON       accessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseJSON `json:"-"`
}

// accessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseJSON contains the JSON
// metadata for the struct
// [AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponse]
type accessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseJSON struct {
	Expiration  apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseEnvelope struct {
	Errors   []AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseEnvelopeMessages `json:"messages,required"`
	Result   []AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseEnvelopeResultInfo `json:"result_info"`
	JSON       accessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseEnvelopeJSON       `json:"-"`
}

// accessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseEnvelope]
type accessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseEnvelopeErrors struct {
	Code    int64                                                                        `json:"code,required"`
	Message string                                                                       `json:"message,required"`
	JSON    accessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseEnvelopeErrorsJSON `json:"-"`
}

// accessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseEnvelopeErrors]
type accessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseEnvelopeMessages struct {
	Code    int64                                                                          `json:"code,required"`
	Message string                                                                         `json:"message,required"`
	JSON    accessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseEnvelopeMessagesJSON `json:"-"`
}

// accessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseEnvelopeMessages]
type accessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseEnvelopeSuccess bool

const (
	AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseEnvelopeSuccessTrue AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseEnvelopeSuccess = true
)

type AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                          `json:"total_count"`
	JSON       accessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseEnvelopeResultInfoJSON `json:"-"`
}

// accessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseEnvelopeResultInfo]
type accessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserFailedLoginZeroTrustUsersGetFailedLoginsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
