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

// UserLoadBalancerMonitorReferenceService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewUserLoadBalancerMonitorReferenceService] method instead.
type UserLoadBalancerMonitorReferenceService struct {
	Options []option.RequestOption
}

// NewUserLoadBalancerMonitorReferenceService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewUserLoadBalancerMonitorReferenceService(opts ...option.RequestOption) (r *UserLoadBalancerMonitorReferenceService) {
	r = &UserLoadBalancerMonitorReferenceService{}
	r.Options = opts
	return
}

// Get the list of resources that reference the provided monitor.
func (r *UserLoadBalancerMonitorReferenceService) LoadBalancerMonitorsListMonitorReferences(ctx context.Context, identifier string, opts ...option.RequestOption) (res *[]UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/monitors/%s/references", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponse struct {
	ReferenceType UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseReferenceType `json:"reference_type"`
	ResourceID    string                                                                                         `json:"resource_id"`
	ResourceName  string                                                                                         `json:"resource_name"`
	ResourceType  string                                                                                         `json:"resource_type"`
	JSON          userLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseJSON          `json:"-"`
}

// userLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponse]
type userLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseJSON struct {
	ReferenceType apijson.Field
	ResourceID    apijson.Field
	ResourceName  apijson.Field
	ResourceType  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseReferenceType string

const (
	UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseReferenceTypeStar     UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseReferenceType = "*"
	UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseReferenceTypeReferral UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseReferenceType = "referral"
	UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseReferenceTypeReferrer UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseReferenceType = "referrer"
)

type UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseEnvelope struct {
	Errors   []UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeMessages `json:"messages,required"`
	// List of resources that reference a given monitor.
	Result []UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeResultInfo `json:"result_info"`
	JSON       userLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeJSON       `json:"-"`
}

// userLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseEnvelope]
type userLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeErrors struct {
	Code    int64                                                                                               `json:"code,required"`
	Message string                                                                                              `json:"message,required"`
	JSON    userLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeErrorsJSON `json:"-"`
}

// userLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeErrors]
type userLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeMessages struct {
	Code    int64                                                                                                 `json:"code,required"`
	Message string                                                                                                `json:"message,required"`
	JSON    userLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeMessagesJSON `json:"-"`
}

// userLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeMessages]
type userLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeSuccess bool

const (
	UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeSuccessTrue UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeSuccess = true
)

type UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                                 `json:"total_count"`
	JSON       userLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeResultInfoJSON `json:"-"`
}

// userLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeResultInfo]
type userLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorReferenceLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
