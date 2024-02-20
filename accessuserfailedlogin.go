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
func (r *AccessUserFailedLoginService) List(ctx context.Context, identifier string, id string, opts ...option.RequestOption) (res *[]AccessUserFailedLoginListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessUserFailedLoginListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/users/%s/failed_logins", identifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessUserFailedLoginListResponse struct {
	Expiration int64                                 `json:"expiration"`
	Metadata   interface{}                           `json:"metadata"`
	JSON       accessUserFailedLoginListResponseJSON `json:"-"`
}

// accessUserFailedLoginListResponseJSON contains the JSON metadata for the struct
// [AccessUserFailedLoginListResponse]
type accessUserFailedLoginListResponseJSON struct {
	Expiration  apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserFailedLoginListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserFailedLoginListResponseEnvelope struct {
	Errors   []AccessUserFailedLoginListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessUserFailedLoginListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AccessUserFailedLoginListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AccessUserFailedLoginListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AccessUserFailedLoginListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       accessUserFailedLoginListResponseEnvelopeJSON       `json:"-"`
}

// accessUserFailedLoginListResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessUserFailedLoginListResponseEnvelope]
type accessUserFailedLoginListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserFailedLoginListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserFailedLoginListResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accessUserFailedLoginListResponseEnvelopeErrorsJSON `json:"-"`
}

// accessUserFailedLoginListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AccessUserFailedLoginListResponseEnvelopeErrors]
type accessUserFailedLoginListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserFailedLoginListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserFailedLoginListResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    accessUserFailedLoginListResponseEnvelopeMessagesJSON `json:"-"`
}

// accessUserFailedLoginListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AccessUserFailedLoginListResponseEnvelopeMessages]
type accessUserFailedLoginListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserFailedLoginListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessUserFailedLoginListResponseEnvelopeSuccess bool

const (
	AccessUserFailedLoginListResponseEnvelopeSuccessTrue AccessUserFailedLoginListResponseEnvelopeSuccess = true
)

type AccessUserFailedLoginListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                 `json:"total_count"`
	JSON       accessUserFailedLoginListResponseEnvelopeResultInfoJSON `json:"-"`
}

// accessUserFailedLoginListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [AccessUserFailedLoginListResponseEnvelopeResultInfo]
type accessUserFailedLoginListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserFailedLoginListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
