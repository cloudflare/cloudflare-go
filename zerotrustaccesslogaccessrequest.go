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

// ZeroTrustAccessLogAccessRequestService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZeroTrustAccessLogAccessRequestService] method instead.
type ZeroTrustAccessLogAccessRequestService struct {
	Options []option.RequestOption
}

// NewZeroTrustAccessLogAccessRequestService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZeroTrustAccessLogAccessRequestService(opts ...option.RequestOption) (r *ZeroTrustAccessLogAccessRequestService) {
	r = &ZeroTrustAccessLogAccessRequestService{}
	r.Options = opts
	return
}

// Gets a list of Access authentication audit logs for an account.
func (r *ZeroTrustAccessLogAccessRequestService) List(ctx context.Context, identifier string, opts ...option.RequestOption) (res *[]ZeroTrustAccessLogAccessRequestListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessLogAccessRequestListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/logs/access_requests", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustAccessLogAccessRequestListResponse struct {
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
	UserEmail string                                          `json:"user_email" format:"email"`
	JSON      zeroTrustAccessLogAccessRequestListResponseJSON `json:"-"`
}

// zeroTrustAccessLogAccessRequestListResponseJSON contains the JSON metadata for
// the struct [ZeroTrustAccessLogAccessRequestListResponse]
type zeroTrustAccessLogAccessRequestListResponseJSON struct {
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

func (r *ZeroTrustAccessLogAccessRequestListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessLogAccessRequestListResponseEnvelope struct {
	Errors   []ZeroTrustAccessLogAccessRequestListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessLogAccessRequestListResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustAccessLogAccessRequestListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    ZeroTrustAccessLogAccessRequestListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustAccessLogAccessRequestListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustAccessLogAccessRequestListResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustAccessLogAccessRequestListResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustAccessLogAccessRequestListResponseEnvelope]
type zeroTrustAccessLogAccessRequestListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessLogAccessRequestListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessLogAccessRequestListResponseEnvelopeErrors struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    zeroTrustAccessLogAccessRequestListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessLogAccessRequestListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessLogAccessRequestListResponseEnvelopeErrors]
type zeroTrustAccessLogAccessRequestListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessLogAccessRequestListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessLogAccessRequestListResponseEnvelopeMessages struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    zeroTrustAccessLogAccessRequestListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessLogAccessRequestListResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessLogAccessRequestListResponseEnvelopeMessages]
type zeroTrustAccessLogAccessRequestListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessLogAccessRequestListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessLogAccessRequestListResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessLogAccessRequestListResponseEnvelopeSuccessTrue ZeroTrustAccessLogAccessRequestListResponseEnvelopeSuccess = true
)

type ZeroTrustAccessLogAccessRequestListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                           `json:"total_count"`
	JSON       zeroTrustAccessLogAccessRequestListResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustAccessLogAccessRequestListResponseEnvelopeResultInfoJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessLogAccessRequestListResponseEnvelopeResultInfo]
type zeroTrustAccessLogAccessRequestListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessLogAccessRequestListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
