// File generated from our OpenAPI spec by Stainless.

package addressing

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// PrefixBGPPrefixService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewPrefixBGPPrefixService] method
// instead.
type PrefixBGPPrefixService struct {
	Options []option.RequestOption
}

// NewPrefixBGPPrefixService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPrefixBGPPrefixService(opts ...option.RequestOption) (r *PrefixBGPPrefixService) {
	r = &PrefixBGPPrefixService{}
	r.Options = opts
	return
}

// List all BGP Prefixes within the specified IP Prefix. BGP Prefixes are used to
// control which specific subnets are advertised to the Internet. It is possible to
// advertise subnets more specific than an IP Prefix by creating more specific BGP
// Prefixes.
func (r *PrefixBGPPrefixService) List(ctx context.Context, prefixID string, query PrefixBGPPrefixListParams, opts ...option.RequestOption) (res *[]PrefixBGPPrefixListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PrefixBGPPrefixListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/prefixes", query.AccountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update the properties of a BGP Prefix, such as the on demand advertisement
// status (advertised or withdrawn).
func (r *PrefixBGPPrefixService) Edit(ctx context.Context, prefixID string, bgpPrefixID string, params PrefixBGPPrefixEditParams, opts ...option.RequestOption) (res *PrefixBGPPrefixEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PrefixBGPPrefixEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/prefixes/%s", params.AccountID, prefixID, bgpPrefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieve a single BGP Prefix according to its identifier
func (r *PrefixBGPPrefixService) Get(ctx context.Context, prefixID string, bgpPrefixID string, query PrefixBGPPrefixGetParams, opts ...option.RequestOption) (res *PrefixBGPPrefixGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PrefixBGPPrefixGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/prefixes/%s", query.AccountID, prefixID, bgpPrefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PrefixBGPPrefixListResponse struct {
	// Identifier
	ID string `json:"id"`
	// Autonomous System Number (ASN) the prefix will be advertised under.
	ASN           int64                                    `json:"asn,nullable"`
	BGPSignalOpts PrefixBGPPrefixListResponseBGPSignalOpts `json:"bgp_signal_opts"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr       string                              `json:"cidr"`
	CreatedAt  time.Time                           `json:"created_at" format:"date-time"`
	ModifiedAt time.Time                           `json:"modified_at" format:"date-time"`
	OnDemand   PrefixBGPPrefixListResponseOnDemand `json:"on_demand"`
	JSON       prefixBGPPrefixListResponseJSON     `json:"-"`
}

// prefixBGPPrefixListResponseJSON contains the JSON metadata for the struct
// [PrefixBGPPrefixListResponse]
type prefixBGPPrefixListResponseJSON struct {
	ID            apijson.Field
	ASN           apijson.Field
	BGPSignalOpts apijson.Field
	Cidr          apijson.Field
	CreatedAt     apijson.Field
	ModifiedAt    apijson.Field
	OnDemand      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PrefixBGPPrefixListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPPrefixListResponseJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPPrefixListResponseBGPSignalOpts struct {
	// Whether control of advertisement of the prefix to the Internet is enabled to be
	// performed via BGP signal
	Enabled bool `json:"enabled"`
	// Last time BGP signaling control was toggled. This field is null if BGP signaling
	// has never been enabled.
	ModifiedAt time.Time                                    `json:"modified_at,nullable" format:"date-time"`
	JSON       prefixBGPPrefixListResponseBGPSignalOptsJSON `json:"-"`
}

// prefixBGPPrefixListResponseBGPSignalOptsJSON contains the JSON metadata for the
// struct [PrefixBGPPrefixListResponseBGPSignalOpts]
type prefixBGPPrefixListResponseBGPSignalOptsJSON struct {
	Enabled     apijson.Field
	ModifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPPrefixListResponseBGPSignalOpts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPPrefixListResponseBGPSignalOptsJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPPrefixListResponseOnDemand struct {
	// Prefix advertisement status to the Internet. This field is only not 'null' if on
	// demand is enabled.
	Advertised bool `json:"advertised,nullable"`
	// Last time the advertisement status was changed. This field is only not 'null' if
	// on demand is enabled.
	AdvertisedModifiedAt time.Time `json:"advertised_modified_at,nullable" format:"date-time"`
	// Whether advertisement of the prefix to the Internet may be dynamically enabled
	// or disabled.
	OnDemandEnabled bool `json:"on_demand_enabled"`
	// Whether advertisement status of the prefix is locked, meaning it cannot be
	// changed.
	OnDemandLocked bool                                    `json:"on_demand_locked"`
	JSON           prefixBGPPrefixListResponseOnDemandJSON `json:"-"`
}

// prefixBGPPrefixListResponseOnDemandJSON contains the JSON metadata for the
// struct [PrefixBGPPrefixListResponseOnDemand]
type prefixBGPPrefixListResponseOnDemandJSON struct {
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	OnDemandEnabled      apijson.Field
	OnDemandLocked       apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *PrefixBGPPrefixListResponseOnDemand) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPPrefixListResponseOnDemandJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPPrefixEditResponse struct {
	// Identifier
	ID string `json:"id"`
	// Autonomous System Number (ASN) the prefix will be advertised under.
	ASN           int64                                    `json:"asn,nullable"`
	BGPSignalOpts PrefixBGPPrefixEditResponseBGPSignalOpts `json:"bgp_signal_opts"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr       string                              `json:"cidr"`
	CreatedAt  time.Time                           `json:"created_at" format:"date-time"`
	ModifiedAt time.Time                           `json:"modified_at" format:"date-time"`
	OnDemand   PrefixBGPPrefixEditResponseOnDemand `json:"on_demand"`
	JSON       prefixBGPPrefixEditResponseJSON     `json:"-"`
}

// prefixBGPPrefixEditResponseJSON contains the JSON metadata for the struct
// [PrefixBGPPrefixEditResponse]
type prefixBGPPrefixEditResponseJSON struct {
	ID            apijson.Field
	ASN           apijson.Field
	BGPSignalOpts apijson.Field
	Cidr          apijson.Field
	CreatedAt     apijson.Field
	ModifiedAt    apijson.Field
	OnDemand      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PrefixBGPPrefixEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPPrefixEditResponseJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPPrefixEditResponseBGPSignalOpts struct {
	// Whether control of advertisement of the prefix to the Internet is enabled to be
	// performed via BGP signal
	Enabled bool `json:"enabled"`
	// Last time BGP signaling control was toggled. This field is null if BGP signaling
	// has never been enabled.
	ModifiedAt time.Time                                    `json:"modified_at,nullable" format:"date-time"`
	JSON       prefixBGPPrefixEditResponseBGPSignalOptsJSON `json:"-"`
}

// prefixBGPPrefixEditResponseBGPSignalOptsJSON contains the JSON metadata for the
// struct [PrefixBGPPrefixEditResponseBGPSignalOpts]
type prefixBGPPrefixEditResponseBGPSignalOptsJSON struct {
	Enabled     apijson.Field
	ModifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPPrefixEditResponseBGPSignalOpts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPPrefixEditResponseBGPSignalOptsJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPPrefixEditResponseOnDemand struct {
	// Prefix advertisement status to the Internet. This field is only not 'null' if on
	// demand is enabled.
	Advertised bool `json:"advertised,nullable"`
	// Last time the advertisement status was changed. This field is only not 'null' if
	// on demand is enabled.
	AdvertisedModifiedAt time.Time `json:"advertised_modified_at,nullable" format:"date-time"`
	// Whether advertisement of the prefix to the Internet may be dynamically enabled
	// or disabled.
	OnDemandEnabled bool `json:"on_demand_enabled"`
	// Whether advertisement status of the prefix is locked, meaning it cannot be
	// changed.
	OnDemandLocked bool                                    `json:"on_demand_locked"`
	JSON           prefixBGPPrefixEditResponseOnDemandJSON `json:"-"`
}

// prefixBGPPrefixEditResponseOnDemandJSON contains the JSON metadata for the
// struct [PrefixBGPPrefixEditResponseOnDemand]
type prefixBGPPrefixEditResponseOnDemandJSON struct {
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	OnDemandEnabled      apijson.Field
	OnDemandLocked       apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *PrefixBGPPrefixEditResponseOnDemand) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPPrefixEditResponseOnDemandJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPPrefixGetResponse struct {
	// Identifier
	ID string `json:"id"`
	// Autonomous System Number (ASN) the prefix will be advertised under.
	ASN           int64                                   `json:"asn,nullable"`
	BGPSignalOpts PrefixBGPPrefixGetResponseBGPSignalOpts `json:"bgp_signal_opts"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr       string                             `json:"cidr"`
	CreatedAt  time.Time                          `json:"created_at" format:"date-time"`
	ModifiedAt time.Time                          `json:"modified_at" format:"date-time"`
	OnDemand   PrefixBGPPrefixGetResponseOnDemand `json:"on_demand"`
	JSON       prefixBGPPrefixGetResponseJSON     `json:"-"`
}

// prefixBGPPrefixGetResponseJSON contains the JSON metadata for the struct
// [PrefixBGPPrefixGetResponse]
type prefixBGPPrefixGetResponseJSON struct {
	ID            apijson.Field
	ASN           apijson.Field
	BGPSignalOpts apijson.Field
	Cidr          apijson.Field
	CreatedAt     apijson.Field
	ModifiedAt    apijson.Field
	OnDemand      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PrefixBGPPrefixGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPPrefixGetResponseJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPPrefixGetResponseBGPSignalOpts struct {
	// Whether control of advertisement of the prefix to the Internet is enabled to be
	// performed via BGP signal
	Enabled bool `json:"enabled"`
	// Last time BGP signaling control was toggled. This field is null if BGP signaling
	// has never been enabled.
	ModifiedAt time.Time                                   `json:"modified_at,nullable" format:"date-time"`
	JSON       prefixBGPPrefixGetResponseBGPSignalOptsJSON `json:"-"`
}

// prefixBGPPrefixGetResponseBGPSignalOptsJSON contains the JSON metadata for the
// struct [PrefixBGPPrefixGetResponseBGPSignalOpts]
type prefixBGPPrefixGetResponseBGPSignalOptsJSON struct {
	Enabled     apijson.Field
	ModifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPPrefixGetResponseBGPSignalOpts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPPrefixGetResponseBGPSignalOptsJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPPrefixGetResponseOnDemand struct {
	// Prefix advertisement status to the Internet. This field is only not 'null' if on
	// demand is enabled.
	Advertised bool `json:"advertised,nullable"`
	// Last time the advertisement status was changed. This field is only not 'null' if
	// on demand is enabled.
	AdvertisedModifiedAt time.Time `json:"advertised_modified_at,nullable" format:"date-time"`
	// Whether advertisement of the prefix to the Internet may be dynamically enabled
	// or disabled.
	OnDemandEnabled bool `json:"on_demand_enabled"`
	// Whether advertisement status of the prefix is locked, meaning it cannot be
	// changed.
	OnDemandLocked bool                                   `json:"on_demand_locked"`
	JSON           prefixBGPPrefixGetResponseOnDemandJSON `json:"-"`
}

// prefixBGPPrefixGetResponseOnDemandJSON contains the JSON metadata for the struct
// [PrefixBGPPrefixGetResponseOnDemand]
type prefixBGPPrefixGetResponseOnDemandJSON struct {
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	OnDemandEnabled      apijson.Field
	OnDemandLocked       apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *PrefixBGPPrefixGetResponseOnDemand) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPPrefixGetResponseOnDemandJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPPrefixListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PrefixBGPPrefixListResponseEnvelope struct {
	Errors   []PrefixBGPPrefixListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PrefixBGPPrefixListResponseEnvelopeMessages `json:"messages,required"`
	Result   []PrefixBGPPrefixListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    PrefixBGPPrefixListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo PrefixBGPPrefixListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       prefixBGPPrefixListResponseEnvelopeJSON       `json:"-"`
}

// prefixBGPPrefixListResponseEnvelopeJSON contains the JSON metadata for the
// struct [PrefixBGPPrefixListResponseEnvelope]
type prefixBGPPrefixListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPPrefixListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPPrefixListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPPrefixListResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    prefixBGPPrefixListResponseEnvelopeErrorsJSON `json:"-"`
}

// prefixBGPPrefixListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [PrefixBGPPrefixListResponseEnvelopeErrors]
type prefixBGPPrefixListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPPrefixListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPPrefixListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPPrefixListResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    prefixBGPPrefixListResponseEnvelopeMessagesJSON `json:"-"`
}

// prefixBGPPrefixListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [PrefixBGPPrefixListResponseEnvelopeMessages]
type prefixBGPPrefixListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPPrefixListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPPrefixListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PrefixBGPPrefixListResponseEnvelopeSuccess bool

const (
	PrefixBGPPrefixListResponseEnvelopeSuccessTrue PrefixBGPPrefixListResponseEnvelopeSuccess = true
)

type PrefixBGPPrefixListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                           `json:"total_count"`
	JSON       prefixBGPPrefixListResponseEnvelopeResultInfoJSON `json:"-"`
}

// prefixBGPPrefixListResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [PrefixBGPPrefixListResponseEnvelopeResultInfo]
type prefixBGPPrefixListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPPrefixListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPPrefixListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPPrefixEditParams struct {
	// Identifier
	AccountID param.Field[string]                            `path:"account_id,required"`
	OnDemand  param.Field[PrefixBGPPrefixEditParamsOnDemand] `json:"on_demand"`
}

func (r PrefixBGPPrefixEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrefixBGPPrefixEditParamsOnDemand struct {
	Advertised param.Field[bool] `json:"advertised"`
}

func (r PrefixBGPPrefixEditParamsOnDemand) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrefixBGPPrefixEditResponseEnvelope struct {
	Errors   []PrefixBGPPrefixEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PrefixBGPPrefixEditResponseEnvelopeMessages `json:"messages,required"`
	Result   PrefixBGPPrefixEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success PrefixBGPPrefixEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    prefixBGPPrefixEditResponseEnvelopeJSON    `json:"-"`
}

// prefixBGPPrefixEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [PrefixBGPPrefixEditResponseEnvelope]
type prefixBGPPrefixEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPPrefixEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPPrefixEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPPrefixEditResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    prefixBGPPrefixEditResponseEnvelopeErrorsJSON `json:"-"`
}

// prefixBGPPrefixEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [PrefixBGPPrefixEditResponseEnvelopeErrors]
type prefixBGPPrefixEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPPrefixEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPPrefixEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPPrefixEditResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    prefixBGPPrefixEditResponseEnvelopeMessagesJSON `json:"-"`
}

// prefixBGPPrefixEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [PrefixBGPPrefixEditResponseEnvelopeMessages]
type prefixBGPPrefixEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPPrefixEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPPrefixEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PrefixBGPPrefixEditResponseEnvelopeSuccess bool

const (
	PrefixBGPPrefixEditResponseEnvelopeSuccessTrue PrefixBGPPrefixEditResponseEnvelopeSuccess = true
)

type PrefixBGPPrefixGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PrefixBGPPrefixGetResponseEnvelope struct {
	Errors   []PrefixBGPPrefixGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PrefixBGPPrefixGetResponseEnvelopeMessages `json:"messages,required"`
	Result   PrefixBGPPrefixGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success PrefixBGPPrefixGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    prefixBGPPrefixGetResponseEnvelopeJSON    `json:"-"`
}

// prefixBGPPrefixGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PrefixBGPPrefixGetResponseEnvelope]
type prefixBGPPrefixGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPPrefixGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPPrefixGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPPrefixGetResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    prefixBGPPrefixGetResponseEnvelopeErrorsJSON `json:"-"`
}

// prefixBGPPrefixGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [PrefixBGPPrefixGetResponseEnvelopeErrors]
type prefixBGPPrefixGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPPrefixGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPPrefixGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPPrefixGetResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    prefixBGPPrefixGetResponseEnvelopeMessagesJSON `json:"-"`
}

// prefixBGPPrefixGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [PrefixBGPPrefixGetResponseEnvelopeMessages]
type prefixBGPPrefixGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPPrefixGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPPrefixGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PrefixBGPPrefixGetResponseEnvelopeSuccess bool

const (
	PrefixBGPPrefixGetResponseEnvelopeSuccessTrue PrefixBGPPrefixGetResponseEnvelopeSuccess = true
)
