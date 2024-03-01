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

// ZeroTrustAccessUserFailedLoginService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZeroTrustAccessUserFailedLoginService] method instead.
type ZeroTrustAccessUserFailedLoginService struct {
	Options []option.RequestOption
}

// NewZeroTrustAccessUserFailedLoginService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZeroTrustAccessUserFailedLoginService(opts ...option.RequestOption) (r *ZeroTrustAccessUserFailedLoginService) {
	r = &ZeroTrustAccessUserFailedLoginService{}
	r.Options = opts
	return
}

// Get all failed login attempts for a single user.
func (r *ZeroTrustAccessUserFailedLoginService) List(ctx context.Context, identifier string, id string, opts ...option.RequestOption) (res *[]ZeroTrustAccessUserFailedLoginListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessUserFailedLoginListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/users/%s/failed_logins", identifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustAccessUserFailedLoginListResponse struct {
	Expiration int64                                          `json:"expiration"`
	Metadata   interface{}                                    `json:"metadata"`
	JSON       zeroTrustAccessUserFailedLoginListResponseJSON `json:"-"`
}

// zeroTrustAccessUserFailedLoginListResponseJSON contains the JSON metadata for
// the struct [ZeroTrustAccessUserFailedLoginListResponse]
type zeroTrustAccessUserFailedLoginListResponseJSON struct {
	Expiration  apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessUserFailedLoginListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessUserFailedLoginListResponseEnvelope struct {
	Errors   []ZeroTrustAccessUserFailedLoginListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessUserFailedLoginListResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustAccessUserFailedLoginListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    ZeroTrustAccessUserFailedLoginListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustAccessUserFailedLoginListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustAccessUserFailedLoginListResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustAccessUserFailedLoginListResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustAccessUserFailedLoginListResponseEnvelope]
type zeroTrustAccessUserFailedLoginListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessUserFailedLoginListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessUserFailedLoginListResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zeroTrustAccessUserFailedLoginListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessUserFailedLoginListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessUserFailedLoginListResponseEnvelopeErrors]
type zeroTrustAccessUserFailedLoginListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessUserFailedLoginListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessUserFailedLoginListResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    zeroTrustAccessUserFailedLoginListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessUserFailedLoginListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessUserFailedLoginListResponseEnvelopeMessages]
type zeroTrustAccessUserFailedLoginListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessUserFailedLoginListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessUserFailedLoginListResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessUserFailedLoginListResponseEnvelopeSuccessTrue ZeroTrustAccessUserFailedLoginListResponseEnvelopeSuccess = true
)

type ZeroTrustAccessUserFailedLoginListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                          `json:"total_count"`
	JSON       zeroTrustAccessUserFailedLoginListResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustAccessUserFailedLoginListResponseEnvelopeResultInfoJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessUserFailedLoginListResponseEnvelopeResultInfo]
type zeroTrustAccessUserFailedLoginListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessUserFailedLoginListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
