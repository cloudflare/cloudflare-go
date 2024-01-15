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
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/prefixes/%s", accountIdentifier, prefixIdentifier, bgpPrefixIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the properties of a BGP Prefix, such as the on demand advertisement
// status (advertised or withdrawn).
func (r *AddressingPrefixBgpPrefixService) Update(ctx context.Context, accountIdentifier string, prefixIdentifier string, bgpPrefixIdentifier string, body AddressingPrefixBgpPrefixUpdateParams, opts ...option.RequestOption) (res *AddressingPrefixBgpPrefixUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/prefixes/%s", accountIdentifier, prefixIdentifier, bgpPrefixIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all BGP Prefixes within the specified IP Prefix. BGP Prefixes are used to
// control which specific subnets are advertised to the Internet. It is possible to
// advertise subnets more specific than an IP Prefix by creating more specific BGP
// Prefixes.
func (r *AddressingPrefixBgpPrefixService) List(ctx context.Context, accountIdentifier string, prefixIdentifier string, opts ...option.RequestOption) (res *AddressingPrefixBgpPrefixListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/prefixes", accountIdentifier, prefixIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AddressingPrefixBgpPrefixGetResponse struct {
	Errors   []AddressingPrefixBgpPrefixGetResponseError   `json:"errors"`
	Messages []AddressingPrefixBgpPrefixGetResponseMessage `json:"messages"`
	Result   AddressingPrefixBgpPrefixGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AddressingPrefixBgpPrefixGetResponseSuccess `json:"success"`
	JSON    addressingPrefixBgpPrefixGetResponseJSON    `json:"-"`
}

// addressingPrefixBgpPrefixGetResponseJSON contains the JSON metadata for the
// struct [AddressingPrefixBgpPrefixGetResponse]
type addressingPrefixBgpPrefixGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBgpPrefixGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBgpPrefixGetResponseError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    addressingPrefixBgpPrefixGetResponseErrorJSON `json:"-"`
}

// addressingPrefixBgpPrefixGetResponseErrorJSON contains the JSON metadata for the
// struct [AddressingPrefixBgpPrefixGetResponseError]
type addressingPrefixBgpPrefixGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBgpPrefixGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBgpPrefixGetResponseMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    addressingPrefixBgpPrefixGetResponseMessageJSON `json:"-"`
}

// addressingPrefixBgpPrefixGetResponseMessageJSON contains the JSON metadata for
// the struct [AddressingPrefixBgpPrefixGetResponseMessage]
type addressingPrefixBgpPrefixGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBgpPrefixGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBgpPrefixGetResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// Autonomous System Number (ASN) the prefix will be advertised under.
	ASN           int64                                                   `json:"asn,nullable"`
	BgpSignalOpts AddressingPrefixBgpPrefixGetResponseResultBgpSignalOpts `json:"bgp_signal_opts"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr       string                                             `json:"cidr"`
	CreatedAt  time.Time                                          `json:"created_at" format:"date-time"`
	ModifiedAt time.Time                                          `json:"modified_at" format:"date-time"`
	OnDemand   AddressingPrefixBgpPrefixGetResponseResultOnDemand `json:"on_demand"`
	JSON       addressingPrefixBgpPrefixGetResponseResultJSON     `json:"-"`
}

// addressingPrefixBgpPrefixGetResponseResultJSON contains the JSON metadata for
// the struct [AddressingPrefixBgpPrefixGetResponseResult]
type addressingPrefixBgpPrefixGetResponseResultJSON struct {
	ID            apijson.Field
	ASN           apijson.Field
	BgpSignalOpts apijson.Field
	Cidr          apijson.Field
	CreatedAt     apijson.Field
	ModifiedAt    apijson.Field
	OnDemand      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AddressingPrefixBgpPrefixGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBgpPrefixGetResponseResultBgpSignalOpts struct {
	// Whether control of advertisement of the prefix to the Internet is enabled to be
	// performed via BGP signal
	Enabled bool `json:"enabled"`
	// Last time BGP signaling control was toggled. This field is null if BGP signaling
	// has never been enabled.
	ModifiedAt time.Time                                                   `json:"modified_at,nullable" format:"date-time"`
	JSON       addressingPrefixBgpPrefixGetResponseResultBgpSignalOptsJSON `json:"-"`
}

// addressingPrefixBgpPrefixGetResponseResultBgpSignalOptsJSON contains the JSON
// metadata for the struct
// [AddressingPrefixBgpPrefixGetResponseResultBgpSignalOpts]
type addressingPrefixBgpPrefixGetResponseResultBgpSignalOptsJSON struct {
	Enabled     apijson.Field
	ModifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBgpPrefixGetResponseResultBgpSignalOpts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBgpPrefixGetResponseResultOnDemand struct {
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
	OnDemandLocked bool                                                   `json:"on_demand_locked"`
	JSON           addressingPrefixBgpPrefixGetResponseResultOnDemandJSON `json:"-"`
}

// addressingPrefixBgpPrefixGetResponseResultOnDemandJSON contains the JSON
// metadata for the struct [AddressingPrefixBgpPrefixGetResponseResultOnDemand]
type addressingPrefixBgpPrefixGetResponseResultOnDemandJSON struct {
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	OnDemandEnabled      apijson.Field
	OnDemandLocked       apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AddressingPrefixBgpPrefixGetResponseResultOnDemand) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressingPrefixBgpPrefixGetResponseSuccess bool

const (
	AddressingPrefixBgpPrefixGetResponseSuccessTrue AddressingPrefixBgpPrefixGetResponseSuccess = true
)

type AddressingPrefixBgpPrefixUpdateResponse struct {
	Errors   []AddressingPrefixBgpPrefixUpdateResponseError   `json:"errors"`
	Messages []AddressingPrefixBgpPrefixUpdateResponseMessage `json:"messages"`
	Result   AddressingPrefixBgpPrefixUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AddressingPrefixBgpPrefixUpdateResponseSuccess `json:"success"`
	JSON    addressingPrefixBgpPrefixUpdateResponseJSON    `json:"-"`
}

// addressingPrefixBgpPrefixUpdateResponseJSON contains the JSON metadata for the
// struct [AddressingPrefixBgpPrefixUpdateResponse]
type addressingPrefixBgpPrefixUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBgpPrefixUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBgpPrefixUpdateResponseError struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    addressingPrefixBgpPrefixUpdateResponseErrorJSON `json:"-"`
}

// addressingPrefixBgpPrefixUpdateResponseErrorJSON contains the JSON metadata for
// the struct [AddressingPrefixBgpPrefixUpdateResponseError]
type addressingPrefixBgpPrefixUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBgpPrefixUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBgpPrefixUpdateResponseMessage struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    addressingPrefixBgpPrefixUpdateResponseMessageJSON `json:"-"`
}

// addressingPrefixBgpPrefixUpdateResponseMessageJSON contains the JSON metadata
// for the struct [AddressingPrefixBgpPrefixUpdateResponseMessage]
type addressingPrefixBgpPrefixUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBgpPrefixUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBgpPrefixUpdateResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// Autonomous System Number (ASN) the prefix will be advertised under.
	ASN           int64                                                      `json:"asn,nullable"`
	BgpSignalOpts AddressingPrefixBgpPrefixUpdateResponseResultBgpSignalOpts `json:"bgp_signal_opts"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr       string                                                `json:"cidr"`
	CreatedAt  time.Time                                             `json:"created_at" format:"date-time"`
	ModifiedAt time.Time                                             `json:"modified_at" format:"date-time"`
	OnDemand   AddressingPrefixBgpPrefixUpdateResponseResultOnDemand `json:"on_demand"`
	JSON       addressingPrefixBgpPrefixUpdateResponseResultJSON     `json:"-"`
}

// addressingPrefixBgpPrefixUpdateResponseResultJSON contains the JSON metadata for
// the struct [AddressingPrefixBgpPrefixUpdateResponseResult]
type addressingPrefixBgpPrefixUpdateResponseResultJSON struct {
	ID            apijson.Field
	ASN           apijson.Field
	BgpSignalOpts apijson.Field
	Cidr          apijson.Field
	CreatedAt     apijson.Field
	ModifiedAt    apijson.Field
	OnDemand      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AddressingPrefixBgpPrefixUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBgpPrefixUpdateResponseResultBgpSignalOpts struct {
	// Whether control of advertisement of the prefix to the Internet is enabled to be
	// performed via BGP signal
	Enabled bool `json:"enabled"`
	// Last time BGP signaling control was toggled. This field is null if BGP signaling
	// has never been enabled.
	ModifiedAt time.Time                                                      `json:"modified_at,nullable" format:"date-time"`
	JSON       addressingPrefixBgpPrefixUpdateResponseResultBgpSignalOptsJSON `json:"-"`
}

// addressingPrefixBgpPrefixUpdateResponseResultBgpSignalOptsJSON contains the JSON
// metadata for the struct
// [AddressingPrefixBgpPrefixUpdateResponseResultBgpSignalOpts]
type addressingPrefixBgpPrefixUpdateResponseResultBgpSignalOptsJSON struct {
	Enabled     apijson.Field
	ModifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBgpPrefixUpdateResponseResultBgpSignalOpts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBgpPrefixUpdateResponseResultOnDemand struct {
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
	OnDemandLocked bool                                                      `json:"on_demand_locked"`
	JSON           addressingPrefixBgpPrefixUpdateResponseResultOnDemandJSON `json:"-"`
}

// addressingPrefixBgpPrefixUpdateResponseResultOnDemandJSON contains the JSON
// metadata for the struct [AddressingPrefixBgpPrefixUpdateResponseResultOnDemand]
type addressingPrefixBgpPrefixUpdateResponseResultOnDemandJSON struct {
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	OnDemandEnabled      apijson.Field
	OnDemandLocked       apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AddressingPrefixBgpPrefixUpdateResponseResultOnDemand) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressingPrefixBgpPrefixUpdateResponseSuccess bool

const (
	AddressingPrefixBgpPrefixUpdateResponseSuccessTrue AddressingPrefixBgpPrefixUpdateResponseSuccess = true
)

type AddressingPrefixBgpPrefixListResponse struct {
	Errors     []AddressingPrefixBgpPrefixListResponseError    `json:"errors"`
	Messages   []AddressingPrefixBgpPrefixListResponseMessage  `json:"messages"`
	Result     []AddressingPrefixBgpPrefixListResponseResult   `json:"result"`
	ResultInfo AddressingPrefixBgpPrefixListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AddressingPrefixBgpPrefixListResponseSuccess `json:"success"`
	JSON    addressingPrefixBgpPrefixListResponseJSON    `json:"-"`
}

// addressingPrefixBgpPrefixListResponseJSON contains the JSON metadata for the
// struct [AddressingPrefixBgpPrefixListResponse]
type addressingPrefixBgpPrefixListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBgpPrefixListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBgpPrefixListResponseError struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    addressingPrefixBgpPrefixListResponseErrorJSON `json:"-"`
}

// addressingPrefixBgpPrefixListResponseErrorJSON contains the JSON metadata for
// the struct [AddressingPrefixBgpPrefixListResponseError]
type addressingPrefixBgpPrefixListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBgpPrefixListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBgpPrefixListResponseMessage struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    addressingPrefixBgpPrefixListResponseMessageJSON `json:"-"`
}

// addressingPrefixBgpPrefixListResponseMessageJSON contains the JSON metadata for
// the struct [AddressingPrefixBgpPrefixListResponseMessage]
type addressingPrefixBgpPrefixListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBgpPrefixListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBgpPrefixListResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// Autonomous System Number (ASN) the prefix will be advertised under.
	ASN           int64                                                    `json:"asn,nullable"`
	BgpSignalOpts AddressingPrefixBgpPrefixListResponseResultBgpSignalOpts `json:"bgp_signal_opts"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr       string                                              `json:"cidr"`
	CreatedAt  time.Time                                           `json:"created_at" format:"date-time"`
	ModifiedAt time.Time                                           `json:"modified_at" format:"date-time"`
	OnDemand   AddressingPrefixBgpPrefixListResponseResultOnDemand `json:"on_demand"`
	JSON       addressingPrefixBgpPrefixListResponseResultJSON     `json:"-"`
}

// addressingPrefixBgpPrefixListResponseResultJSON contains the JSON metadata for
// the struct [AddressingPrefixBgpPrefixListResponseResult]
type addressingPrefixBgpPrefixListResponseResultJSON struct {
	ID            apijson.Field
	ASN           apijson.Field
	BgpSignalOpts apijson.Field
	Cidr          apijson.Field
	CreatedAt     apijson.Field
	ModifiedAt    apijson.Field
	OnDemand      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AddressingPrefixBgpPrefixListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBgpPrefixListResponseResultBgpSignalOpts struct {
	// Whether control of advertisement of the prefix to the Internet is enabled to be
	// performed via BGP signal
	Enabled bool `json:"enabled"`
	// Last time BGP signaling control was toggled. This field is null if BGP signaling
	// has never been enabled.
	ModifiedAt time.Time                                                    `json:"modified_at,nullable" format:"date-time"`
	JSON       addressingPrefixBgpPrefixListResponseResultBgpSignalOptsJSON `json:"-"`
}

// addressingPrefixBgpPrefixListResponseResultBgpSignalOptsJSON contains the JSON
// metadata for the struct
// [AddressingPrefixBgpPrefixListResponseResultBgpSignalOpts]
type addressingPrefixBgpPrefixListResponseResultBgpSignalOptsJSON struct {
	Enabled     apijson.Field
	ModifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBgpPrefixListResponseResultBgpSignalOpts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBgpPrefixListResponseResultOnDemand struct {
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
	OnDemandLocked bool                                                    `json:"on_demand_locked"`
	JSON           addressingPrefixBgpPrefixListResponseResultOnDemandJSON `json:"-"`
}

// addressingPrefixBgpPrefixListResponseResultOnDemandJSON contains the JSON
// metadata for the struct [AddressingPrefixBgpPrefixListResponseResultOnDemand]
type addressingPrefixBgpPrefixListResponseResultOnDemandJSON struct {
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	OnDemandEnabled      apijson.Field
	OnDemandLocked       apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AddressingPrefixBgpPrefixListResponseResultOnDemand) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixBgpPrefixListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                             `json:"total_count"`
	JSON       addressingPrefixBgpPrefixListResponseResultInfoJSON `json:"-"`
}

// addressingPrefixBgpPrefixListResponseResultInfoJSON contains the JSON metadata
// for the struct [AddressingPrefixBgpPrefixListResponseResultInfo]
type addressingPrefixBgpPrefixListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixBgpPrefixListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressingPrefixBgpPrefixListResponseSuccess bool

const (
	AddressingPrefixBgpPrefixListResponseSuccessTrue AddressingPrefixBgpPrefixListResponseSuccess = true
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
