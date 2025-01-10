// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// AccessLogAccessRequestService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessLogAccessRequestService] method instead.
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
func (r *AccessLogAccessRequestService) List(ctx context.Context, params AccessLogAccessRequestListParams, opts ...option.RequestOption) (res *[]AccessRequests, err error) {
	var env AccessLogAccessRequestListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/logs/access_requests", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
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
	AppUID string `json:"app_uid"`
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
	AppUID      apijson.Field
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

type AccessLogAccessRequestListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The chronological sorting order for the logs.
	Direction param.Field[AccessLogAccessRequestListParamsDirection] `query:"direction"`
	// The maximum number of log entries to retrieve.
	Limit param.Field[int64] `query:"limit"`
	// The earliest event timestamp to query.
	Since param.Field[time.Time] `query:"since" format:"date-time"`
	// The latest event timestamp to query.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [AccessLogAccessRequestListParams]'s query parameters as
// `url.Values`.
func (r AccessLogAccessRequestListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The chronological sorting order for the logs.
type AccessLogAccessRequestListParamsDirection string

const (
	AccessLogAccessRequestListParamsDirectionDesc AccessLogAccessRequestListParamsDirection = "desc"
	AccessLogAccessRequestListParamsDirectionAsc  AccessLogAccessRequestListParamsDirection = "asc"
)

func (r AccessLogAccessRequestListParamsDirection) IsKnown() bool {
	switch r {
	case AccessLogAccessRequestListParamsDirectionDesc, AccessLogAccessRequestListParamsDirectionAsc:
		return true
	}
	return false
}

type AccessLogAccessRequestListResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccessLogAccessRequestListResponseEnvelopeSuccess `json:"success,required"`
	Result  []AccessRequests                                  `json:"result"`
	JSON    accessLogAccessRequestListResponseEnvelopeJSON    `json:"-"`
}

// accessLogAccessRequestListResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessLogAccessRequestListResponseEnvelope]
type accessLogAccessRequestListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
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
