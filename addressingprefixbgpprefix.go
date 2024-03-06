// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AddressingPrefixBGPPrefixService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAddressingPrefixBGPPrefixService] method instead.
type AddressingPrefixBGPPrefixService struct {
	Options []option.RequestOption
}

// NewAddressingPrefixBGPPrefixService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAddressingPrefixBGPPrefixService(opts ...option.RequestOption) (r *AddressingPrefixBGPPrefixService) {
	r = &AddressingPrefixBGPPrefixService{}
	r.Options = opts
	return
}

// List all BGP Prefixes within the specified IP Prefix. BGP Prefixes are used to
// control which specific subnets are advertised to the Internet. It is possible to
// advertise subnets more specific than an IP Prefix by creating more specific BGP
// Prefixes.
func (r *AddressingPrefixBGPPrefixService) List(ctx context.Context, prefixID string, query AddressingPrefixBGPPrefixListParams, opts ...option.RequestOption) (res *[]AddressingPrefixBGPPrefixListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressingPrefixBGPPrefixListResponseEnvelope
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
func (r *AddressingPrefixBGPPrefixService) Edit(ctx context.Context, prefixID string, bgpPrefixID string, params AddressingPrefixBGPPrefixEditParams, opts ...option.RequestOption) (res *AddressingPrefixBGPPrefixEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressingPrefixBGPPrefixEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/prefixes/%s", params.AccountID, prefixID, bgpPrefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieve a single BGP Prefix according to its identifier
func (r *AddressingPrefixBGPPrefixService) Get(ctx context.Context, prefixID string, bgpPrefixID string, query AddressingPrefixBGPPrefixGetParams, opts ...option.RequestOption) (res *AddressingPrefixBGPPrefixGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressingPrefixBGPPrefixGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/prefixes/%s", query.AccountID, prefixID, bgpPrefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AddressingPrefixBGPPrefixListResponse struct {
	// Identifier
	ID string `json:"id"`
	// Autonomous System Number (ASN) the prefix will be advertised under.
	ASN           int64                                              `json:"asn,nullable"`
	BGPSignalOpts AddressingPrefixBGPPrefixListResponseBGPSignalOpts `json:"bgp_signal_opts"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr       string                                        `json:"cidr"`
	CreatedAt  time.Time                                     `json:"created_at" format:"date-time"`
	ModifiedAt time.Time                                     `json:"modified_at" format:"date-time"`
	OnDemand   AddressingPrefixBGPPrefixListResponseOnDemand `json:"on_demand"`
	JSON       addressingPrefixBGPPrefixListResponseJSON     `json:"-"`
}

// addressingPrefixBGPPrefixListResponseJSON contains the JSON metadata for the
// struct [AddressingPrefixBGPPrefixListResponse]
type addressingPrefixBGPPrefixListResponseJSON struct {
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

func (r *AddressingPrefixBGPPrefixListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixBGPPrefixListResponseJSON) RawJSON() string {
	return r.raw
}

type AddressingPrefixBGPPrefixListResponseBGPSignalOpts struct {
	// Whether control of advertisement of the prefix to the Internet is enabled to be
	// performed via BGP signal
	Enabled bool `json:"enabled"`
	// Last time BGP signaling control was toggled. This field is null if BGP signaling
	// has never been enabled.
	ModifiedAt time.Time                                              `json:"modified_at,nullable" format:"date-time"`
	JSON       addressingPrefixBGPPrefixListResponseBGPSignalOptsJSON `json:"-"`
}

// addressingPrefixBGPPrefixListResponseBGPSignalOptsJSON contains the JSON
// metadata for the struct [AddressingPrefixBGPPrefixListResponseBGPSignalOpts]
type addressingPrefixBGPPrefixListResponseBGPSignalOptsJSON struct {
	Enabled     apijson.Field
	ModifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBGPPrefixListResponseBGPSignalOpts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixBGPPrefixListResponseBGPSignalOptsJSON) RawJSON() string {
	return r.raw
}

type AddressingPrefixBGPPrefixListResponseOnDemand struct {
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
	OnDemandLocked bool                                              `json:"on_demand_locked"`
	JSON           addressingPrefixBGPPrefixListResponseOnDemandJSON `json:"-"`
}

// addressingPrefixBGPPrefixListResponseOnDemandJSON contains the JSON metadata for
// the struct [AddressingPrefixBGPPrefixListResponseOnDemand]
type addressingPrefixBGPPrefixListResponseOnDemandJSON struct {
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	OnDemandEnabled      apijson.Field
	OnDemandLocked       apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AddressingPrefixBGPPrefixListResponseOnDemand) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixBGPPrefixListResponseOnDemandJSON) RawJSON() string {
	return r.raw
}

type AddressingPrefixBGPPrefixEditResponse struct {
	// Identifier
	ID string `json:"id"`
	// Autonomous System Number (ASN) the prefix will be advertised under.
	ASN           int64                                              `json:"asn,nullable"`
	BGPSignalOpts AddressingPrefixBGPPrefixEditResponseBGPSignalOpts `json:"bgp_signal_opts"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr       string                                        `json:"cidr"`
	CreatedAt  time.Time                                     `json:"created_at" format:"date-time"`
	ModifiedAt time.Time                                     `json:"modified_at" format:"date-time"`
	OnDemand   AddressingPrefixBGPPrefixEditResponseOnDemand `json:"on_demand"`
	JSON       addressingPrefixBGPPrefixEditResponseJSON     `json:"-"`
}

// addressingPrefixBGPPrefixEditResponseJSON contains the JSON metadata for the
// struct [AddressingPrefixBGPPrefixEditResponse]
type addressingPrefixBGPPrefixEditResponseJSON struct {
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

func (r *AddressingPrefixBGPPrefixEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixBGPPrefixEditResponseJSON) RawJSON() string {
	return r.raw
}

type AddressingPrefixBGPPrefixEditResponseBGPSignalOpts struct {
	// Whether control of advertisement of the prefix to the Internet is enabled to be
	// performed via BGP signal
	Enabled bool `json:"enabled"`
	// Last time BGP signaling control was toggled. This field is null if BGP signaling
	// has never been enabled.
	ModifiedAt time.Time                                              `json:"modified_at,nullable" format:"date-time"`
	JSON       addressingPrefixBGPPrefixEditResponseBGPSignalOptsJSON `json:"-"`
}

// addressingPrefixBGPPrefixEditResponseBGPSignalOptsJSON contains the JSON
// metadata for the struct [AddressingPrefixBGPPrefixEditResponseBGPSignalOpts]
type addressingPrefixBGPPrefixEditResponseBGPSignalOptsJSON struct {
	Enabled     apijson.Field
	ModifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBGPPrefixEditResponseBGPSignalOpts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixBGPPrefixEditResponseBGPSignalOptsJSON) RawJSON() string {
	return r.raw
}

type AddressingPrefixBGPPrefixEditResponseOnDemand struct {
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
	OnDemandLocked bool                                              `json:"on_demand_locked"`
	JSON           addressingPrefixBGPPrefixEditResponseOnDemandJSON `json:"-"`
}

// addressingPrefixBGPPrefixEditResponseOnDemandJSON contains the JSON metadata for
// the struct [AddressingPrefixBGPPrefixEditResponseOnDemand]
type addressingPrefixBGPPrefixEditResponseOnDemandJSON struct {
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	OnDemandEnabled      apijson.Field
	OnDemandLocked       apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AddressingPrefixBGPPrefixEditResponseOnDemand) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixBGPPrefixEditResponseOnDemandJSON) RawJSON() string {
	return r.raw
}

type AddressingPrefixBGPPrefixGetResponse struct {
	// Identifier
	ID string `json:"id"`
	// Autonomous System Number (ASN) the prefix will be advertised under.
	ASN           int64                                             `json:"asn,nullable"`
	BGPSignalOpts AddressingPrefixBGPPrefixGetResponseBGPSignalOpts `json:"bgp_signal_opts"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr       string                                       `json:"cidr"`
	CreatedAt  time.Time                                    `json:"created_at" format:"date-time"`
	ModifiedAt time.Time                                    `json:"modified_at" format:"date-time"`
	OnDemand   AddressingPrefixBGPPrefixGetResponseOnDemand `json:"on_demand"`
	JSON       addressingPrefixBGPPrefixGetResponseJSON     `json:"-"`
}

// addressingPrefixBGPPrefixGetResponseJSON contains the JSON metadata for the
// struct [AddressingPrefixBGPPrefixGetResponse]
type addressingPrefixBGPPrefixGetResponseJSON struct {
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

func (r *AddressingPrefixBGPPrefixGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixBGPPrefixGetResponseJSON) RawJSON() string {
	return r.raw
}

type AddressingPrefixBGPPrefixGetResponseBGPSignalOpts struct {
	// Whether control of advertisement of the prefix to the Internet is enabled to be
	// performed via BGP signal
	Enabled bool `json:"enabled"`
	// Last time BGP signaling control was toggled. This field is null if BGP signaling
	// has never been enabled.
	ModifiedAt time.Time                                             `json:"modified_at,nullable" format:"date-time"`
	JSON       addressingPrefixBGPPrefixGetResponseBGPSignalOptsJSON `json:"-"`
}

// addressingPrefixBGPPrefixGetResponseBGPSignalOptsJSON contains the JSON metadata
// for the struct [AddressingPrefixBGPPrefixGetResponseBGPSignalOpts]
type addressingPrefixBGPPrefixGetResponseBGPSignalOptsJSON struct {
	Enabled     apijson.Field
	ModifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBGPPrefixGetResponseBGPSignalOpts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixBGPPrefixGetResponseBGPSignalOptsJSON) RawJSON() string {
	return r.raw
}

type AddressingPrefixBGPPrefixGetResponseOnDemand struct {
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
	OnDemandLocked bool                                             `json:"on_demand_locked"`
	JSON           addressingPrefixBGPPrefixGetResponseOnDemandJSON `json:"-"`
}

// addressingPrefixBGPPrefixGetResponseOnDemandJSON contains the JSON metadata for
// the struct [AddressingPrefixBGPPrefixGetResponseOnDemand]
type addressingPrefixBGPPrefixGetResponseOnDemandJSON struct {
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	OnDemandEnabled      apijson.Field
	OnDemandLocked       apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AddressingPrefixBGPPrefixGetResponseOnDemand) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixBGPPrefixGetResponseOnDemandJSON) RawJSON() string {
	return r.raw
}

type AddressingPrefixBGPPrefixListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressingPrefixBGPPrefixListResponseEnvelope struct {
	Errors   []AddressingPrefixBGPPrefixListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressingPrefixBGPPrefixListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AddressingPrefixBGPPrefixListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AddressingPrefixBGPPrefixListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AddressingPrefixBGPPrefixListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       addressingPrefixBGPPrefixListResponseEnvelopeJSON       `json:"-"`
}

// addressingPrefixBGPPrefixListResponseEnvelopeJSON contains the JSON metadata for
// the struct [AddressingPrefixBGPPrefixListResponseEnvelope]
type addressingPrefixBGPPrefixListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBGPPrefixListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixBGPPrefixListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AddressingPrefixBGPPrefixListResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    addressingPrefixBGPPrefixListResponseEnvelopeErrorsJSON `json:"-"`
}

// addressingPrefixBGPPrefixListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AddressingPrefixBGPPrefixListResponseEnvelopeErrors]
type addressingPrefixBGPPrefixListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBGPPrefixListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixBGPPrefixListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AddressingPrefixBGPPrefixListResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    addressingPrefixBGPPrefixListResponseEnvelopeMessagesJSON `json:"-"`
}

// addressingPrefixBGPPrefixListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AddressingPrefixBGPPrefixListResponseEnvelopeMessages]
type addressingPrefixBGPPrefixListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBGPPrefixListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixBGPPrefixListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressingPrefixBGPPrefixListResponseEnvelopeSuccess bool

const (
	AddressingPrefixBGPPrefixListResponseEnvelopeSuccessTrue AddressingPrefixBGPPrefixListResponseEnvelopeSuccess = true
)

type AddressingPrefixBGPPrefixListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                     `json:"total_count"`
	JSON       addressingPrefixBGPPrefixListResponseEnvelopeResultInfoJSON `json:"-"`
}

// addressingPrefixBGPPrefixListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [AddressingPrefixBGPPrefixListResponseEnvelopeResultInfo]
type addressingPrefixBGPPrefixListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBGPPrefixListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixBGPPrefixListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type AddressingPrefixBGPPrefixEditParams struct {
	// Identifier
	AccountID param.Field[string]                                      `path:"account_id,required"`
	OnDemand  param.Field[AddressingPrefixBGPPrefixEditParamsOnDemand] `json:"on_demand"`
}

func (r AddressingPrefixBGPPrefixEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressingPrefixBGPPrefixEditParamsOnDemand struct {
	Advertised param.Field[bool] `json:"advertised"`
}

func (r AddressingPrefixBGPPrefixEditParamsOnDemand) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressingPrefixBGPPrefixEditResponseEnvelope struct {
	Errors   []AddressingPrefixBGPPrefixEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressingPrefixBGPPrefixEditResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressingPrefixBGPPrefixEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressingPrefixBGPPrefixEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressingPrefixBGPPrefixEditResponseEnvelopeJSON    `json:"-"`
}

// addressingPrefixBGPPrefixEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [AddressingPrefixBGPPrefixEditResponseEnvelope]
type addressingPrefixBGPPrefixEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBGPPrefixEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixBGPPrefixEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AddressingPrefixBGPPrefixEditResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    addressingPrefixBGPPrefixEditResponseEnvelopeErrorsJSON `json:"-"`
}

// addressingPrefixBGPPrefixEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AddressingPrefixBGPPrefixEditResponseEnvelopeErrors]
type addressingPrefixBGPPrefixEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBGPPrefixEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixBGPPrefixEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AddressingPrefixBGPPrefixEditResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    addressingPrefixBGPPrefixEditResponseEnvelopeMessagesJSON `json:"-"`
}

// addressingPrefixBGPPrefixEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AddressingPrefixBGPPrefixEditResponseEnvelopeMessages]
type addressingPrefixBGPPrefixEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBGPPrefixEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixBGPPrefixEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressingPrefixBGPPrefixEditResponseEnvelopeSuccess bool

const (
	AddressingPrefixBGPPrefixEditResponseEnvelopeSuccessTrue AddressingPrefixBGPPrefixEditResponseEnvelopeSuccess = true
)

type AddressingPrefixBGPPrefixGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressingPrefixBGPPrefixGetResponseEnvelope struct {
	Errors   []AddressingPrefixBGPPrefixGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressingPrefixBGPPrefixGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressingPrefixBGPPrefixGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressingPrefixBGPPrefixGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressingPrefixBGPPrefixGetResponseEnvelopeJSON    `json:"-"`
}

// addressingPrefixBGPPrefixGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [AddressingPrefixBGPPrefixGetResponseEnvelope]
type addressingPrefixBGPPrefixGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBGPPrefixGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixBGPPrefixGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AddressingPrefixBGPPrefixGetResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    addressingPrefixBGPPrefixGetResponseEnvelopeErrorsJSON `json:"-"`
}

// addressingPrefixBGPPrefixGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AddressingPrefixBGPPrefixGetResponseEnvelopeErrors]
type addressingPrefixBGPPrefixGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBGPPrefixGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixBGPPrefixGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AddressingPrefixBGPPrefixGetResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    addressingPrefixBGPPrefixGetResponseEnvelopeMessagesJSON `json:"-"`
}

// addressingPrefixBGPPrefixGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AddressingPrefixBGPPrefixGetResponseEnvelopeMessages]
type addressingPrefixBGPPrefixGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBGPPrefixGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixBGPPrefixGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressingPrefixBGPPrefixGetResponseEnvelopeSuccess bool

const (
	AddressingPrefixBGPPrefixGetResponseEnvelopeSuccessTrue AddressingPrefixBGPPrefixGetResponseEnvelopeSuccess = true
)
