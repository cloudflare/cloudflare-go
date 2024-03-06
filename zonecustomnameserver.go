// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// ZoneCustomNameserverService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneCustomNameserverService]
// method instead.
type ZoneCustomNameserverService struct {
	Options []option.RequestOption
}

// NewZoneCustomNameserverService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneCustomNameserverService(opts ...option.RequestOption) (r *ZoneCustomNameserverService) {
	r = &ZoneCustomNameserverService{}
	r.Options = opts
	return
}

// Set metadata for account-level custom nameservers on a zone.
//
// If you would like new zones in the account to use account custom nameservers by
// default, use PUT /accounts/:identifier to set the account setting
// use_account_custom_ns_by_default to true.
func (r *ZoneCustomNameserverService) Update(ctx context.Context, params ZoneCustomNameserverUpdateParams, opts ...option.RequestOption) (res *ZoneCustomNameserverUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneCustomNameserverUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/custom_ns", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get metadata for account-level custom nameservers on a zone.
func (r *ZoneCustomNameserverService) Get(ctx context.Context, query ZoneCustomNameserverGetParams, opts ...option.RequestOption) (res *ZoneCustomNameserverGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneCustomNameserverGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/custom_ns", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [ZoneCustomNameserverUpdateResponseUnknown],
// [ZoneCustomNameserverUpdateResponseArray] or [shared.UnionString].
type ZoneCustomNameserverUpdateResponse interface {
	ImplementsZoneCustomNameserverUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneCustomNameserverUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ZoneCustomNameserverUpdateResponseArray []interface{}

func (r ZoneCustomNameserverUpdateResponseArray) ImplementsZoneCustomNameserverUpdateResponse() {}

// Union satisfied by [ZoneCustomNameserverGetResponseUnknown],
// [ZoneCustomNameserverGetResponseArray] or [shared.UnionString].
type ZoneCustomNameserverGetResponse interface {
	ImplementsZoneCustomNameserverGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneCustomNameserverGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ZoneCustomNameserverGetResponseArray []interface{}

func (r ZoneCustomNameserverGetResponseArray) ImplementsZoneCustomNameserverGetResponse() {}

type ZoneCustomNameserverUpdateParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Whether zone uses account-level custom nameservers.
	Enabled param.Field[bool] `json:"enabled"`
	// The number of the name server set to assign to the zone.
	NSSet param.Field[float64] `json:"ns_set"`
}

func (r ZoneCustomNameserverUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneCustomNameserverUpdateResponseEnvelope struct {
	Errors   []ZoneCustomNameserverUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneCustomNameserverUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   ZoneCustomNameserverUpdateResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    ZoneCustomNameserverUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZoneCustomNameserverUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zoneCustomNameserverUpdateResponseEnvelopeJSON       `json:"-"`
}

// zoneCustomNameserverUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneCustomNameserverUpdateResponseEnvelope]
type zoneCustomNameserverUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomNameserverUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCustomNameserverUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneCustomNameserverUpdateResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zoneCustomNameserverUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneCustomNameserverUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZoneCustomNameserverUpdateResponseEnvelopeErrors]
type zoneCustomNameserverUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomNameserverUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCustomNameserverUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneCustomNameserverUpdateResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zoneCustomNameserverUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneCustomNameserverUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneCustomNameserverUpdateResponseEnvelopeMessages]
type zoneCustomNameserverUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomNameserverUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCustomNameserverUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneCustomNameserverUpdateResponseEnvelopeSuccess bool

const (
	ZoneCustomNameserverUpdateResponseEnvelopeSuccessTrue ZoneCustomNameserverUpdateResponseEnvelopeSuccess = true
)

type ZoneCustomNameserverUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                  `json:"total_count"`
	JSON       zoneCustomNameserverUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// zoneCustomNameserverUpdateResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [ZoneCustomNameserverUpdateResponseEnvelopeResultInfo]
type zoneCustomNameserverUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomNameserverUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCustomNameserverUpdateResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZoneCustomNameserverGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneCustomNameserverGetResponseEnvelope struct {
	Errors   []ZoneCustomNameserverGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneCustomNameserverGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ZoneCustomNameserverGetResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success ZoneCustomNameserverGetResponseEnvelopeSuccess `json:"success,required"`
	// Whether zone uses account-level custom nameservers.
	Enabled bool `json:"enabled"`
	// The number of the name server set to assign to the zone.
	NSSet      float64                                           `json:"ns_set"`
	ResultInfo ZoneCustomNameserverGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zoneCustomNameserverGetResponseEnvelopeJSON       `json:"-"`
}

// zoneCustomNameserverGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneCustomNameserverGetResponseEnvelope]
type zoneCustomNameserverGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	Enabled     apijson.Field
	NSSet       apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomNameserverGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCustomNameserverGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneCustomNameserverGetResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zoneCustomNameserverGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneCustomNameserverGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZoneCustomNameserverGetResponseEnvelopeErrors]
type zoneCustomNameserverGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomNameserverGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCustomNameserverGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneCustomNameserverGetResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zoneCustomNameserverGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneCustomNameserverGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZoneCustomNameserverGetResponseEnvelopeMessages]
type zoneCustomNameserverGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomNameserverGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCustomNameserverGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneCustomNameserverGetResponseEnvelopeSuccess bool

const (
	ZoneCustomNameserverGetResponseEnvelopeSuccessTrue ZoneCustomNameserverGetResponseEnvelopeSuccess = true
)

type ZoneCustomNameserverGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                               `json:"total_count"`
	JSON       zoneCustomNameserverGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// zoneCustomNameserverGetResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [ZoneCustomNameserverGetResponseEnvelopeResultInfo]
type zoneCustomNameserverGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomNameserverGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCustomNameserverGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
