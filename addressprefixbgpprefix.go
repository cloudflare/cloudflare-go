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

// AddressPrefixBGPPrefixService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAddressPrefixBGPPrefixService]
// method instead.
type AddressPrefixBGPPrefixService struct {
	Options []option.RequestOption
}

// NewAddressPrefixBGPPrefixService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAddressPrefixBGPPrefixService(opts ...option.RequestOption) (r *AddressPrefixBGPPrefixService) {
	r = &AddressPrefixBGPPrefixService{}
	r.Options = opts
	return
}

// List all BGP Prefixes within the specified IP Prefix. BGP Prefixes are used to
// control which specific subnets are advertised to the Internet. It is possible to
// advertise subnets more specific than an IP Prefix by creating more specific BGP
// Prefixes.
func (r *AddressPrefixBGPPrefixService) List(ctx context.Context, prefixID string, query AddressPrefixBGPPrefixListParams, opts ...option.RequestOption) (res *[]AddressPrefixBGPPrefixListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressPrefixBGPPrefixListResponseEnvelope
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
func (r *AddressPrefixBGPPrefixService) Edit(ctx context.Context, prefixID string, bgpPrefixID string, params AddressPrefixBGPPrefixEditParams, opts ...option.RequestOption) (res *AddressPrefixBGPPrefixEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressPrefixBGPPrefixEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/prefixes/%s", params.AccountID, prefixID, bgpPrefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieve a single BGP Prefix according to its identifier
func (r *AddressPrefixBGPPrefixService) Get(ctx context.Context, prefixID string, bgpPrefixID string, query AddressPrefixBGPPrefixGetParams, opts ...option.RequestOption) (res *AddressPrefixBGPPrefixGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressPrefixBGPPrefixGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/prefixes/%s", query.AccountID, prefixID, bgpPrefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AddressPrefixBGPPrefixListResponse struct {
	// Identifier
	ID string `json:"id"`
	// Autonomous System Number (ASN) the prefix will be advertised under.
	ASN           int64                                           `json:"asn,nullable"`
	BGPSignalOpts AddressPrefixBGPPrefixListResponseBGPSignalOpts `json:"bgp_signal_opts"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr       string                                     `json:"cidr"`
	CreatedAt  time.Time                                  `json:"created_at" format:"date-time"`
	ModifiedAt time.Time                                  `json:"modified_at" format:"date-time"`
	OnDemand   AddressPrefixBGPPrefixListResponseOnDemand `json:"on_demand"`
	JSON       addressPrefixBGPPrefixListResponseJSON     `json:"-"`
}

// addressPrefixBGPPrefixListResponseJSON contains the JSON metadata for the struct
// [AddressPrefixBGPPrefixListResponse]
type addressPrefixBGPPrefixListResponseJSON struct {
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

func (r *AddressPrefixBGPPrefixListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixBGPPrefixListResponseBGPSignalOpts struct {
	// Whether control of advertisement of the prefix to the Internet is enabled to be
	// performed via BGP signal
	Enabled bool `json:"enabled"`
	// Last time BGP signaling control was toggled. This field is null if BGP signaling
	// has never been enabled.
	ModifiedAt time.Time                                           `json:"modified_at,nullable" format:"date-time"`
	JSON       addressPrefixBGPPrefixListResponseBGPSignalOptsJSON `json:"-"`
}

// addressPrefixBGPPrefixListResponseBGPSignalOptsJSON contains the JSON metadata
// for the struct [AddressPrefixBGPPrefixListResponseBGPSignalOpts]
type addressPrefixBGPPrefixListResponseBGPSignalOptsJSON struct {
	Enabled     apijson.Field
	ModifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPPrefixListResponseBGPSignalOpts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixBGPPrefixListResponseOnDemand struct {
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
	OnDemandLocked bool                                           `json:"on_demand_locked"`
	JSON           addressPrefixBGPPrefixListResponseOnDemandJSON `json:"-"`
}

// addressPrefixBGPPrefixListResponseOnDemandJSON contains the JSON metadata for
// the struct [AddressPrefixBGPPrefixListResponseOnDemand]
type addressPrefixBGPPrefixListResponseOnDemandJSON struct {
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	OnDemandEnabled      apijson.Field
	OnDemandLocked       apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AddressPrefixBGPPrefixListResponseOnDemand) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixBGPPrefixEditResponse struct {
	// Identifier
	ID string `json:"id"`
	// Autonomous System Number (ASN) the prefix will be advertised under.
	ASN           int64                                           `json:"asn,nullable"`
	BGPSignalOpts AddressPrefixBGPPrefixEditResponseBGPSignalOpts `json:"bgp_signal_opts"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr       string                                     `json:"cidr"`
	CreatedAt  time.Time                                  `json:"created_at" format:"date-time"`
	ModifiedAt time.Time                                  `json:"modified_at" format:"date-time"`
	OnDemand   AddressPrefixBGPPrefixEditResponseOnDemand `json:"on_demand"`
	JSON       addressPrefixBGPPrefixEditResponseJSON     `json:"-"`
}

// addressPrefixBGPPrefixEditResponseJSON contains the JSON metadata for the struct
// [AddressPrefixBGPPrefixEditResponse]
type addressPrefixBGPPrefixEditResponseJSON struct {
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

func (r *AddressPrefixBGPPrefixEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixBGPPrefixEditResponseBGPSignalOpts struct {
	// Whether control of advertisement of the prefix to the Internet is enabled to be
	// performed via BGP signal
	Enabled bool `json:"enabled"`
	// Last time BGP signaling control was toggled. This field is null if BGP signaling
	// has never been enabled.
	ModifiedAt time.Time                                           `json:"modified_at,nullable" format:"date-time"`
	JSON       addressPrefixBGPPrefixEditResponseBGPSignalOptsJSON `json:"-"`
}

// addressPrefixBGPPrefixEditResponseBGPSignalOptsJSON contains the JSON metadata
// for the struct [AddressPrefixBGPPrefixEditResponseBGPSignalOpts]
type addressPrefixBGPPrefixEditResponseBGPSignalOptsJSON struct {
	Enabled     apijson.Field
	ModifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPPrefixEditResponseBGPSignalOpts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixBGPPrefixEditResponseOnDemand struct {
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
	OnDemandLocked bool                                           `json:"on_demand_locked"`
	JSON           addressPrefixBGPPrefixEditResponseOnDemandJSON `json:"-"`
}

// addressPrefixBGPPrefixEditResponseOnDemandJSON contains the JSON metadata for
// the struct [AddressPrefixBGPPrefixEditResponseOnDemand]
type addressPrefixBGPPrefixEditResponseOnDemandJSON struct {
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	OnDemandEnabled      apijson.Field
	OnDemandLocked       apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AddressPrefixBGPPrefixEditResponseOnDemand) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixBGPPrefixGetResponse struct {
	// Identifier
	ID string `json:"id"`
	// Autonomous System Number (ASN) the prefix will be advertised under.
	ASN           int64                                          `json:"asn,nullable"`
	BGPSignalOpts AddressPrefixBGPPrefixGetResponseBGPSignalOpts `json:"bgp_signal_opts"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr       string                                    `json:"cidr"`
	CreatedAt  time.Time                                 `json:"created_at" format:"date-time"`
	ModifiedAt time.Time                                 `json:"modified_at" format:"date-time"`
	OnDemand   AddressPrefixBGPPrefixGetResponseOnDemand `json:"on_demand"`
	JSON       addressPrefixBGPPrefixGetResponseJSON     `json:"-"`
}

// addressPrefixBGPPrefixGetResponseJSON contains the JSON metadata for the struct
// [AddressPrefixBGPPrefixGetResponse]
type addressPrefixBGPPrefixGetResponseJSON struct {
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

func (r *AddressPrefixBGPPrefixGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixBGPPrefixGetResponseBGPSignalOpts struct {
	// Whether control of advertisement of the prefix to the Internet is enabled to be
	// performed via BGP signal
	Enabled bool `json:"enabled"`
	// Last time BGP signaling control was toggled. This field is null if BGP signaling
	// has never been enabled.
	ModifiedAt time.Time                                          `json:"modified_at,nullable" format:"date-time"`
	JSON       addressPrefixBGPPrefixGetResponseBGPSignalOptsJSON `json:"-"`
}

// addressPrefixBGPPrefixGetResponseBGPSignalOptsJSON contains the JSON metadata
// for the struct [AddressPrefixBGPPrefixGetResponseBGPSignalOpts]
type addressPrefixBGPPrefixGetResponseBGPSignalOptsJSON struct {
	Enabled     apijson.Field
	ModifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPPrefixGetResponseBGPSignalOpts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixBGPPrefixGetResponseOnDemand struct {
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
	OnDemandLocked bool                                          `json:"on_demand_locked"`
	JSON           addressPrefixBGPPrefixGetResponseOnDemandJSON `json:"-"`
}

// addressPrefixBGPPrefixGetResponseOnDemandJSON contains the JSON metadata for the
// struct [AddressPrefixBGPPrefixGetResponseOnDemand]
type addressPrefixBGPPrefixGetResponseOnDemandJSON struct {
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	OnDemandEnabled      apijson.Field
	OnDemandLocked       apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AddressPrefixBGPPrefixGetResponseOnDemand) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixBGPPrefixListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressPrefixBGPPrefixListResponseEnvelope struct {
	Errors   []AddressPrefixBGPPrefixListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressPrefixBGPPrefixListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AddressPrefixBGPPrefixListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AddressPrefixBGPPrefixListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AddressPrefixBGPPrefixListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       addressPrefixBGPPrefixListResponseEnvelopeJSON       `json:"-"`
}

// addressPrefixBGPPrefixListResponseEnvelopeJSON contains the JSON metadata for
// the struct [AddressPrefixBGPPrefixListResponseEnvelope]
type addressPrefixBGPPrefixListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPPrefixListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixBGPPrefixListResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    addressPrefixBGPPrefixListResponseEnvelopeErrorsJSON `json:"-"`
}

// addressPrefixBGPPrefixListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AddressPrefixBGPPrefixListResponseEnvelopeErrors]
type addressPrefixBGPPrefixListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPPrefixListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixBGPPrefixListResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    addressPrefixBGPPrefixListResponseEnvelopeMessagesJSON `json:"-"`
}

// addressPrefixBGPPrefixListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AddressPrefixBGPPrefixListResponseEnvelopeMessages]
type addressPrefixBGPPrefixListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPPrefixListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressPrefixBGPPrefixListResponseEnvelopeSuccess bool

const (
	AddressPrefixBGPPrefixListResponseEnvelopeSuccessTrue AddressPrefixBGPPrefixListResponseEnvelopeSuccess = true
)

type AddressPrefixBGPPrefixListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                  `json:"total_count"`
	JSON       addressPrefixBGPPrefixListResponseEnvelopeResultInfoJSON `json:"-"`
}

// addressPrefixBGPPrefixListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [AddressPrefixBGPPrefixListResponseEnvelopeResultInfo]
type addressPrefixBGPPrefixListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPPrefixListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixBGPPrefixEditParams struct {
	// Identifier
	AccountID param.Field[string]                                   `path:"account_id,required"`
	OnDemand  param.Field[AddressPrefixBGPPrefixEditParamsOnDemand] `json:"on_demand"`
}

func (r AddressPrefixBGPPrefixEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressPrefixBGPPrefixEditParamsOnDemand struct {
	Advertised param.Field[bool] `json:"advertised"`
}

func (r AddressPrefixBGPPrefixEditParamsOnDemand) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressPrefixBGPPrefixEditResponseEnvelope struct {
	Errors   []AddressPrefixBGPPrefixEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressPrefixBGPPrefixEditResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressPrefixBGPPrefixEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressPrefixBGPPrefixEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressPrefixBGPPrefixEditResponseEnvelopeJSON    `json:"-"`
}

// addressPrefixBGPPrefixEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [AddressPrefixBGPPrefixEditResponseEnvelope]
type addressPrefixBGPPrefixEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPPrefixEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixBGPPrefixEditResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    addressPrefixBGPPrefixEditResponseEnvelopeErrorsJSON `json:"-"`
}

// addressPrefixBGPPrefixEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AddressPrefixBGPPrefixEditResponseEnvelopeErrors]
type addressPrefixBGPPrefixEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPPrefixEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixBGPPrefixEditResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    addressPrefixBGPPrefixEditResponseEnvelopeMessagesJSON `json:"-"`
}

// addressPrefixBGPPrefixEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AddressPrefixBGPPrefixEditResponseEnvelopeMessages]
type addressPrefixBGPPrefixEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPPrefixEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressPrefixBGPPrefixEditResponseEnvelopeSuccess bool

const (
	AddressPrefixBGPPrefixEditResponseEnvelopeSuccessTrue AddressPrefixBGPPrefixEditResponseEnvelopeSuccess = true
)

type AddressPrefixBGPPrefixGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressPrefixBGPPrefixGetResponseEnvelope struct {
	Errors   []AddressPrefixBGPPrefixGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressPrefixBGPPrefixGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressPrefixBGPPrefixGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressPrefixBGPPrefixGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressPrefixBGPPrefixGetResponseEnvelopeJSON    `json:"-"`
}

// addressPrefixBGPPrefixGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [AddressPrefixBGPPrefixGetResponseEnvelope]
type addressPrefixBGPPrefixGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPPrefixGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixBGPPrefixGetResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    addressPrefixBGPPrefixGetResponseEnvelopeErrorsJSON `json:"-"`
}

// addressPrefixBGPPrefixGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AddressPrefixBGPPrefixGetResponseEnvelopeErrors]
type addressPrefixBGPPrefixGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPPrefixGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixBGPPrefixGetResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    addressPrefixBGPPrefixGetResponseEnvelopeMessagesJSON `json:"-"`
}

// addressPrefixBGPPrefixGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AddressPrefixBGPPrefixGetResponseEnvelopeMessages]
type addressPrefixBGPPrefixGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixBGPPrefixGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressPrefixBGPPrefixGetResponseEnvelopeSuccess bool

const (
	AddressPrefixBGPPrefixGetResponseEnvelopeSuccessTrue AddressPrefixBGPPrefixGetResponseEnvelopeSuccess = true
)
