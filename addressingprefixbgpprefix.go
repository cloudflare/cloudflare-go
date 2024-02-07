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

// AddressingPrefixBgpPrefixService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAddressingPrefixBgpPrefixService] method instead.
type AddressingPrefixBgpPrefixService struct {
	Options []option.RequestOption
}

// NewAddressingPrefixBgpPrefixService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAddressingPrefixBgpPrefixService(opts ...option.RequestOption) (r *AddressingPrefixBgpPrefixService) {
	r = &AddressingPrefixBgpPrefixService{}
	r.Options = opts
	return
}

// Retrieve a single BGP Prefix according to its identifier
func (r *AddressingPrefixBgpPrefixService) Get(ctx context.Context, accountIdentifier string, prefixIdentifier string, bgpPrefixIdentifier string, opts ...option.RequestOption) (res *AddressingPrefixBgpPrefixGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressingPrefixBgpPrefixGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/prefixes/%s", accountIdentifier, prefixIdentifier, bgpPrefixIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update the properties of a BGP Prefix, such as the on demand advertisement
// status (advertised or withdrawn).
func (r *AddressingPrefixBgpPrefixService) Update(ctx context.Context, accountIdentifier string, prefixIdentifier string, bgpPrefixIdentifier string, body AddressingPrefixBgpPrefixUpdateParams, opts ...option.RequestOption) (res *AddressingPrefixBgpPrefixUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressingPrefixBgpPrefixUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/prefixes/%s", accountIdentifier, prefixIdentifier, bgpPrefixIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all BGP Prefixes within the specified IP Prefix. BGP Prefixes are used to
// control which specific subnets are advertised to the Internet. It is possible to
// advertise subnets more specific than an IP Prefix by creating more specific BGP
// Prefixes.
func (r *AddressingPrefixBgpPrefixService) List(ctx context.Context, accountIdentifier string, prefixIdentifier string, opts ...option.RequestOption) (res *[]AddressingPrefixBgpPrefixListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressingPrefixBgpPrefixListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/prefixes", accountIdentifier, prefixIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AddressingPrefixBgpPrefixGetResponse struct {
	// Identifier
	ID string `json:"id"`
	// Autonomous System Number (ASN) the prefix will be advertised under.
	Asn           int64                                             `json:"asn,nullable"`
	BgpSignalOpts AddressingPrefixBgpPrefixGetResponseBgpSignalOpts `json:"bgp_signal_opts"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr       string                                       `json:"cidr"`
	CreatedAt  time.Time                                    `json:"created_at" format:"date-time"`
	ModifiedAt time.Time                                    `json:"modified_at" format:"date-time"`
	OnDemand   AddressingPrefixBgpPrefixGetResponseOnDemand `json:"on_demand"`
	JSON       addressingPrefixBgpPrefixGetResponseJSON     `json:"-"`
}

// addressingPrefixBgpPrefixGetResponseJSON contains the JSON metadata for the
// struct [AddressingPrefixBgpPrefixGetResponse]
type addressingPrefixBgpPrefixGetResponseJSON struct {
	ID            apijson.Field
	Asn           apijson.Field
	BgpSignalOpts apijson.Field
	Cidr          apijson.Field
	CreatedAt     apijson.Field
	ModifiedAt    apijson.Field
	OnDemand      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AddressingPrefixBgpPrefixGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBgpPrefixGetResponseBgpSignalOpts struct {
	// Whether control of advertisement of the prefix to the Internet is enabled to be
	// performed via BGP signal
	Enabled bool `json:"enabled"`
	// Last time BGP signaling control was toggled. This field is null if BGP signaling
	// has never been enabled.
	ModifiedAt time.Time                                             `json:"modified_at,nullable" format:"date-time"`
	JSON       addressingPrefixBgpPrefixGetResponseBgpSignalOptsJSON `json:"-"`
}

// addressingPrefixBgpPrefixGetResponseBgpSignalOptsJSON contains the JSON metadata
// for the struct [AddressingPrefixBgpPrefixGetResponseBgpSignalOpts]
type addressingPrefixBgpPrefixGetResponseBgpSignalOptsJSON struct {
	Enabled     apijson.Field
	ModifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBgpPrefixGetResponseBgpSignalOpts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBgpPrefixGetResponseOnDemand struct {
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
	JSON           addressingPrefixBgpPrefixGetResponseOnDemandJSON `json:"-"`
}

// addressingPrefixBgpPrefixGetResponseOnDemandJSON contains the JSON metadata for
// the struct [AddressingPrefixBgpPrefixGetResponseOnDemand]
type addressingPrefixBgpPrefixGetResponseOnDemandJSON struct {
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	OnDemandEnabled      apijson.Field
	OnDemandLocked       apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AddressingPrefixBgpPrefixGetResponseOnDemand) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBgpPrefixUpdateResponse struct {
	// Identifier
	ID string `json:"id"`
	// Autonomous System Number (ASN) the prefix will be advertised under.
	Asn           int64                                                `json:"asn,nullable"`
	BgpSignalOpts AddressingPrefixBgpPrefixUpdateResponseBgpSignalOpts `json:"bgp_signal_opts"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr       string                                          `json:"cidr"`
	CreatedAt  time.Time                                       `json:"created_at" format:"date-time"`
	ModifiedAt time.Time                                       `json:"modified_at" format:"date-time"`
	OnDemand   AddressingPrefixBgpPrefixUpdateResponseOnDemand `json:"on_demand"`
	JSON       addressingPrefixBgpPrefixUpdateResponseJSON     `json:"-"`
}

// addressingPrefixBgpPrefixUpdateResponseJSON contains the JSON metadata for the
// struct [AddressingPrefixBgpPrefixUpdateResponse]
type addressingPrefixBgpPrefixUpdateResponseJSON struct {
	ID            apijson.Field
	Asn           apijson.Field
	BgpSignalOpts apijson.Field
	Cidr          apijson.Field
	CreatedAt     apijson.Field
	ModifiedAt    apijson.Field
	OnDemand      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AddressingPrefixBgpPrefixUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBgpPrefixUpdateResponseBgpSignalOpts struct {
	// Whether control of advertisement of the prefix to the Internet is enabled to be
	// performed via BGP signal
	Enabled bool `json:"enabled"`
	// Last time BGP signaling control was toggled. This field is null if BGP signaling
	// has never been enabled.
	ModifiedAt time.Time                                                `json:"modified_at,nullable" format:"date-time"`
	JSON       addressingPrefixBgpPrefixUpdateResponseBgpSignalOptsJSON `json:"-"`
}

// addressingPrefixBgpPrefixUpdateResponseBgpSignalOptsJSON contains the JSON
// metadata for the struct [AddressingPrefixBgpPrefixUpdateResponseBgpSignalOpts]
type addressingPrefixBgpPrefixUpdateResponseBgpSignalOptsJSON struct {
	Enabled     apijson.Field
	ModifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBgpPrefixUpdateResponseBgpSignalOpts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBgpPrefixUpdateResponseOnDemand struct {
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
	OnDemandLocked bool                                                `json:"on_demand_locked"`
	JSON           addressingPrefixBgpPrefixUpdateResponseOnDemandJSON `json:"-"`
}

// addressingPrefixBgpPrefixUpdateResponseOnDemandJSON contains the JSON metadata
// for the struct [AddressingPrefixBgpPrefixUpdateResponseOnDemand]
type addressingPrefixBgpPrefixUpdateResponseOnDemandJSON struct {
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	OnDemandEnabled      apijson.Field
	OnDemandLocked       apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AddressingPrefixBgpPrefixUpdateResponseOnDemand) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBgpPrefixListResponse struct {
	// Identifier
	ID string `json:"id"`
	// Autonomous System Number (ASN) the prefix will be advertised under.
	Asn           int64                                              `json:"asn,nullable"`
	BgpSignalOpts AddressingPrefixBgpPrefixListResponseBgpSignalOpts `json:"bgp_signal_opts"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr       string                                        `json:"cidr"`
	CreatedAt  time.Time                                     `json:"created_at" format:"date-time"`
	ModifiedAt time.Time                                     `json:"modified_at" format:"date-time"`
	OnDemand   AddressingPrefixBgpPrefixListResponseOnDemand `json:"on_demand"`
	JSON       addressingPrefixBgpPrefixListResponseJSON     `json:"-"`
}

// addressingPrefixBgpPrefixListResponseJSON contains the JSON metadata for the
// struct [AddressingPrefixBgpPrefixListResponse]
type addressingPrefixBgpPrefixListResponseJSON struct {
	ID            apijson.Field
	Asn           apijson.Field
	BgpSignalOpts apijson.Field
	Cidr          apijson.Field
	CreatedAt     apijson.Field
	ModifiedAt    apijson.Field
	OnDemand      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AddressingPrefixBgpPrefixListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBgpPrefixListResponseBgpSignalOpts struct {
	// Whether control of advertisement of the prefix to the Internet is enabled to be
	// performed via BGP signal
	Enabled bool `json:"enabled"`
	// Last time BGP signaling control was toggled. This field is null if BGP signaling
	// has never been enabled.
	ModifiedAt time.Time                                              `json:"modified_at,nullable" format:"date-time"`
	JSON       addressingPrefixBgpPrefixListResponseBgpSignalOptsJSON `json:"-"`
}

// addressingPrefixBgpPrefixListResponseBgpSignalOptsJSON contains the JSON
// metadata for the struct [AddressingPrefixBgpPrefixListResponseBgpSignalOpts]
type addressingPrefixBgpPrefixListResponseBgpSignalOptsJSON struct {
	Enabled     apijson.Field
	ModifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBgpPrefixListResponseBgpSignalOpts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBgpPrefixListResponseOnDemand struct {
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
	JSON           addressingPrefixBgpPrefixListResponseOnDemandJSON `json:"-"`
}

// addressingPrefixBgpPrefixListResponseOnDemandJSON contains the JSON metadata for
// the struct [AddressingPrefixBgpPrefixListResponseOnDemand]
type addressingPrefixBgpPrefixListResponseOnDemandJSON struct {
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	OnDemandEnabled      apijson.Field
	OnDemandLocked       apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AddressingPrefixBgpPrefixListResponseOnDemand) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBgpPrefixGetResponseEnvelope struct {
	Errors   []AddressingPrefixBgpPrefixGetResponseEnvelopeErrors   `json:"errors"`
	Messages []AddressingPrefixBgpPrefixGetResponseEnvelopeMessages `json:"messages"`
	Result   AddressingPrefixBgpPrefixGetResponse                   `json:"result"`
	// Whether the API call was successful
	Success AddressingPrefixBgpPrefixGetResponseEnvelopeSuccess `json:"success"`
	JSON    addressingPrefixBgpPrefixGetResponseEnvelopeJSON    `json:"-"`
}

// addressingPrefixBgpPrefixGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [AddressingPrefixBgpPrefixGetResponseEnvelope]
type addressingPrefixBgpPrefixGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBgpPrefixGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBgpPrefixGetResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    addressingPrefixBgpPrefixGetResponseEnvelopeErrorsJSON `json:"-"`
}

// addressingPrefixBgpPrefixGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AddressingPrefixBgpPrefixGetResponseEnvelopeErrors]
type addressingPrefixBgpPrefixGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBgpPrefixGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBgpPrefixGetResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    addressingPrefixBgpPrefixGetResponseEnvelopeMessagesJSON `json:"-"`
}

// addressingPrefixBgpPrefixGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AddressingPrefixBgpPrefixGetResponseEnvelopeMessages]
type addressingPrefixBgpPrefixGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBgpPrefixGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressingPrefixBgpPrefixGetResponseEnvelopeSuccess bool

const (
	AddressingPrefixBgpPrefixGetResponseEnvelopeSuccessTrue AddressingPrefixBgpPrefixGetResponseEnvelopeSuccess = true
)

type AddressingPrefixBgpPrefixUpdateParams struct {
	OnDemand param.Field[AddressingPrefixBgpPrefixUpdateParamsOnDemand] `json:"on_demand"`
}

func (r AddressingPrefixBgpPrefixUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressingPrefixBgpPrefixUpdateParamsOnDemand struct {
	Advertised param.Field[bool] `json:"advertised"`
}

func (r AddressingPrefixBgpPrefixUpdateParamsOnDemand) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressingPrefixBgpPrefixUpdateResponseEnvelope struct {
	Errors   []AddressingPrefixBgpPrefixUpdateResponseEnvelopeErrors   `json:"errors"`
	Messages []AddressingPrefixBgpPrefixUpdateResponseEnvelopeMessages `json:"messages"`
	Result   AddressingPrefixBgpPrefixUpdateResponse                   `json:"result"`
	// Whether the API call was successful
	Success AddressingPrefixBgpPrefixUpdateResponseEnvelopeSuccess `json:"success"`
	JSON    addressingPrefixBgpPrefixUpdateResponseEnvelopeJSON    `json:"-"`
}

// addressingPrefixBgpPrefixUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [AddressingPrefixBgpPrefixUpdateResponseEnvelope]
type addressingPrefixBgpPrefixUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBgpPrefixUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBgpPrefixUpdateResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    addressingPrefixBgpPrefixUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// addressingPrefixBgpPrefixUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AddressingPrefixBgpPrefixUpdateResponseEnvelopeErrors]
type addressingPrefixBgpPrefixUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBgpPrefixUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBgpPrefixUpdateResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    addressingPrefixBgpPrefixUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// addressingPrefixBgpPrefixUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [AddressingPrefixBgpPrefixUpdateResponseEnvelopeMessages]
type addressingPrefixBgpPrefixUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBgpPrefixUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressingPrefixBgpPrefixUpdateResponseEnvelopeSuccess bool

const (
	AddressingPrefixBgpPrefixUpdateResponseEnvelopeSuccessTrue AddressingPrefixBgpPrefixUpdateResponseEnvelopeSuccess = true
)

type AddressingPrefixBgpPrefixListResponseEnvelope struct {
	Errors     []AddressingPrefixBgpPrefixListResponseEnvelopeErrors   `json:"errors"`
	Messages   []AddressingPrefixBgpPrefixListResponseEnvelopeMessages `json:"messages"`
	Result     []AddressingPrefixBgpPrefixListResponse                 `json:"result"`
	ResultInfo AddressingPrefixBgpPrefixListResponseEnvelopeResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AddressingPrefixBgpPrefixListResponseEnvelopeSuccess `json:"success"`
	JSON    addressingPrefixBgpPrefixListResponseEnvelopeJSON    `json:"-"`
}

// addressingPrefixBgpPrefixListResponseEnvelopeJSON contains the JSON metadata for
// the struct [AddressingPrefixBgpPrefixListResponseEnvelope]
type addressingPrefixBgpPrefixListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBgpPrefixListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBgpPrefixListResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    addressingPrefixBgpPrefixListResponseEnvelopeErrorsJSON `json:"-"`
}

// addressingPrefixBgpPrefixListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AddressingPrefixBgpPrefixListResponseEnvelopeErrors]
type addressingPrefixBgpPrefixListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBgpPrefixListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBgpPrefixListResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    addressingPrefixBgpPrefixListResponseEnvelopeMessagesJSON `json:"-"`
}

// addressingPrefixBgpPrefixListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AddressingPrefixBgpPrefixListResponseEnvelopeMessages]
type addressingPrefixBgpPrefixListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBgpPrefixListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBgpPrefixListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                     `json:"total_count"`
	JSON       addressingPrefixBgpPrefixListResponseEnvelopeResultInfoJSON `json:"-"`
}

// addressingPrefixBgpPrefixListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [AddressingPrefixBgpPrefixListResponseEnvelopeResultInfo]
type addressingPrefixBgpPrefixListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBgpPrefixListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressingPrefixBgpPrefixListResponseEnvelopeSuccess bool

const (
	AddressingPrefixBgpPrefixListResponseEnvelopeSuccessTrue AddressingPrefixBgpPrefixListResponseEnvelopeSuccess = true
)
