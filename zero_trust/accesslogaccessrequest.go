// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// AccessLogAccessRequestService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccessLogAccessRequestService]
// method instead.
type AccessLogAccessRequestService struct {
	Options []option.RequestOption
}

// NewAccessLogAccessRequestService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccessLogAccessRequestService(opts ...option.RequestOption) (r *AccessLogAccessRequestService) {
	r = &AccessLogAccessRequestService{}
	r.Options = opts
	return
}

// Gets a list of Access authentication audit logs for an account.
func (r *AccessLogAccessRequestService) List(ctx context.Context, identifier string, opts ...option.RequestOption) (res *[]AccessRequests, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessLogAccessRequestListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/logs/access_requests", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessRequests struct {
	// The event that occurred, such as a login attempt.
	Action string `json:"action"`
	// The result of the authentication event.
	Allowed bool `json:"allowed"`
	// The URL of the Access application.
	AppDomain string `json:"app_domain"`
	// The unique identifier for the Access application.
	AppUid string `json:"app_uid"`
	// The IdP used to authenticate.
	Connection string    `json:"connection"`
	CreatedAt  time.Time `json:"created_at" format:"date-time"`
	// The IP address of the authenticating user.
	IPAddress string `json:"ip_address"`
	// The unique identifier for the request to Cloudflare.
	RayID string `json:"ray_id"`
	// The email address of the authenticating user.
	UserEmail string             `json:"user_email" format:"email"`
	JSON      accessRequestsJSON `json:"-"`
}

// accessRequestsJSON contains the JSON metadata for the struct [AccessRequests]
type accessRequestsJSON struct {
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

func (r *AccessRequests) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRequestsJSON) RawJSON() string {
	return r.raw
}

type AccessLogAccessRequestListResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   []AccessRequests                                          `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AccessLogAccessRequestListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AccessLogAccessRequestListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       accessLogAccessRequestListResponseEnvelopeJSON       `json:"-"`
}

// accessLogAccessRequestListResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessLogAccessRequestListResponseEnvelope]
type accessLogAccessRequestListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessLogAccessRequestListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessLogAccessRequestListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessLogAccessRequestListResponseEnvelopeSuccess bool

const (
	AccessLogAccessRequestListResponseEnvelopeSuccessTrue AccessLogAccessRequestListResponseEnvelopeSuccess = true
)

func (r AccessLogAccessRequestListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessLogAccessRequestListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessLogAccessRequestListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                  `json:"total_count"`
	JSON       accessLogAccessRequestListResponseEnvelopeResultInfoJSON `json:"-"`
}

// accessLogAccessRequestListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [AccessLogAccessRequestListResponseEnvelopeResultInfo]
type accessLogAccessRequestListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessLogAccessRequestListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessLogAccessRequestListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
