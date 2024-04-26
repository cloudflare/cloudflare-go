// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// MonitorReferenceService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewMonitorReferenceService] method
// instead.
type MonitorReferenceService struct {
	Options []option.RequestOption
}

// NewMonitorReferenceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMonitorReferenceService(opts ...option.RequestOption) (r *MonitorReferenceService) {
	r = &MonitorReferenceService{}
	r.Options = opts
	return
}

// Get the list of resources that reference the provided monitor.
func (r *MonitorReferenceService) Get(ctx context.Context, monitorID string, query MonitorReferenceGetParams, opts ...option.RequestOption) (res *[]MonitorReferenceGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MonitorReferenceGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors/%s/references", query.AccountID, monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MonitorReferenceGetResponse struct {
	ReferenceType MonitorReferenceGetResponseReferenceType `json:"reference_type"`
	ResourceID    string                                   `json:"resource_id"`
	ResourceName  string                                   `json:"resource_name"`
	ResourceType  string                                   `json:"resource_type"`
	JSON          monitorReferenceGetResponseJSON          `json:"-"`
}

// monitorReferenceGetResponseJSON contains the JSON metadata for the struct
// [MonitorReferenceGetResponse]
type monitorReferenceGetResponseJSON struct {
	ReferenceType apijson.Field
	ResourceID    apijson.Field
	ResourceName  apijson.Field
	ResourceType  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *MonitorReferenceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorReferenceGetResponseJSON) RawJSON() string {
	return r.raw
}

type MonitorReferenceGetResponseReferenceType string

const (
	MonitorReferenceGetResponseReferenceTypeStar     MonitorReferenceGetResponseReferenceType = "*"
	MonitorReferenceGetResponseReferenceTypeReferral MonitorReferenceGetResponseReferenceType = "referral"
	MonitorReferenceGetResponseReferenceTypeReferrer MonitorReferenceGetResponseReferenceType = "referrer"
)

func (r MonitorReferenceGetResponseReferenceType) IsKnown() bool {
	switch r {
	case MonitorReferenceGetResponseReferenceTypeStar, MonitorReferenceGetResponseReferenceTypeReferral, MonitorReferenceGetResponseReferenceTypeReferrer:
		return true
	}
	return false
}

type MonitorReferenceGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type MonitorReferenceGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// List of resources that reference a given monitor.
	Result []MonitorReferenceGetResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    MonitorReferenceGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo MonitorReferenceGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       monitorReferenceGetResponseEnvelopeJSON       `json:"-"`
}

// monitorReferenceGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [MonitorReferenceGetResponseEnvelope]
type monitorReferenceGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorReferenceGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorReferenceGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MonitorReferenceGetResponseEnvelopeSuccess bool

const (
	MonitorReferenceGetResponseEnvelopeSuccessTrue MonitorReferenceGetResponseEnvelopeSuccess = true
)

func (r MonitorReferenceGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MonitorReferenceGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type MonitorReferenceGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                           `json:"total_count"`
	JSON       monitorReferenceGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// monitorReferenceGetResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [MonitorReferenceGetResponseEnvelopeResultInfo]
type monitorReferenceGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorReferenceGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorReferenceGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
