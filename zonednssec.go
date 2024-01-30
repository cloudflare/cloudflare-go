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

// ZoneDnssecService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewZoneDnssecService] method instead.
type ZoneDnssecService struct {
	Options []option.RequestOption
}

// NewZoneDnssecService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneDnssecService(opts ...option.RequestOption) (r *ZoneDnssecService) {
	r = &ZoneDnssecService{}
	r.Options = opts
	return
}

// Details about DNSSEC status and configuration.
func (r *ZoneDnssecService) Get(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneDnssecGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/dnssec", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Enable or disable DNSSEC.
func (r *ZoneDnssecService) Update(ctx context.Context, zoneIdentifier string, body ZoneDnssecUpdateParams, opts ...option.RequestOption) (res *ZoneDnssecUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/dnssec", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type ZoneDnssecGetResponse struct {
	Errors   []ZoneDnssecGetResponseError   `json:"errors,required"`
	Messages []ZoneDnssecGetResponseMessage `json:"messages,required"`
	Result   ZoneDnssecGetResponseResult    `json:"result,required"`
	// Whether the API call was successful
	Success ZoneDnssecGetResponseSuccess `json:"success,required"`
	JSON    zoneDnssecGetResponseJSON    `json:"-"`
}

// zoneDnssecGetResponseJSON contains the JSON metadata for the struct
// [ZoneDnssecGetResponse]
type zoneDnssecGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDnssecGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneDnssecGetResponseError struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    zoneDnssecGetResponseErrorJSON `json:"-"`
}

// zoneDnssecGetResponseErrorJSON contains the JSON metadata for the struct
// [ZoneDnssecGetResponseError]
type zoneDnssecGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDnssecGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneDnssecGetResponseMessage struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    zoneDnssecGetResponseMessageJSON `json:"-"`
}

// zoneDnssecGetResponseMessageJSON contains the JSON metadata for the struct
// [ZoneDnssecGetResponseMessage]
type zoneDnssecGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDnssecGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [ZoneDnssecGetResponseResultUnknown],
// [ZoneDnssecGetResponseResultArray] or [shared.UnionString].
type ZoneDnssecGetResponseResult interface {
	ImplementsZoneDnssecGetResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneDnssecGetResponseResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ZoneDnssecGetResponseResultArray []interface{}

func (r ZoneDnssecGetResponseResultArray) ImplementsZoneDnssecGetResponseResult() {}

// Whether the API call was successful
type ZoneDnssecGetResponseSuccess bool

const (
	ZoneDnssecGetResponseSuccessTrue ZoneDnssecGetResponseSuccess = true
)

type ZoneDnssecUpdateResponse struct {
	Errors   []ZoneDnssecUpdateResponseError   `json:"errors,required"`
	Messages []ZoneDnssecUpdateResponseMessage `json:"messages,required"`
	Result   ZoneDnssecUpdateResponseResult    `json:"result,required"`
	// Whether the API call was successful
	Success ZoneDnssecUpdateResponseSuccess `json:"success,required"`
	JSON    zoneDnssecUpdateResponseJSON    `json:"-"`
}

// zoneDnssecUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneDnssecUpdateResponse]
type zoneDnssecUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDnssecUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneDnssecUpdateResponseError struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    zoneDnssecUpdateResponseErrorJSON `json:"-"`
}

// zoneDnssecUpdateResponseErrorJSON contains the JSON metadata for the struct
// [ZoneDnssecUpdateResponseError]
type zoneDnssecUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDnssecUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneDnssecUpdateResponseMessage struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    zoneDnssecUpdateResponseMessageJSON `json:"-"`
}

// zoneDnssecUpdateResponseMessageJSON contains the JSON metadata for the struct
// [ZoneDnssecUpdateResponseMessage]
type zoneDnssecUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDnssecUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [ZoneDnssecUpdateResponseResultUnknown],
// [ZoneDnssecUpdateResponseResultArray] or [shared.UnionString].
type ZoneDnssecUpdateResponseResult interface {
	ImplementsZoneDnssecUpdateResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneDnssecUpdateResponseResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ZoneDnssecUpdateResponseResultArray []interface{}

func (r ZoneDnssecUpdateResponseResultArray) ImplementsZoneDnssecUpdateResponseResult() {}

// Whether the API call was successful
type ZoneDnssecUpdateResponseSuccess bool

const (
	ZoneDnssecUpdateResponseSuccessTrue ZoneDnssecUpdateResponseSuccess = true
)

type ZoneDnssecUpdateParams struct {
	// If true, multi-signer DNSSEC is enabled on the zone, allowing multiple providers
	// to serve a DNSSEC-signed zone at the same time. This is required for DNSKEY
	// records (except those automatically generated by Cloudflare) to be added to the
	// zone.
	//
	// See
	// [Multi-signer DNSSEC](https://developers.cloudflare.com/dns/dnssec/multi-signer-dnssec/)
	// for details.
	DnssecMultiSigner param.Field[bool] `json:"dnssec_multi_signer"`
	// If true, allows Cloudflare to transfer in a DNSSEC-signed zone including
	// signatures from an external provider, without requiring Cloudflare to sign any
	// records on the fly.
	//
	// Note that this feature has some limitations. See
	// [Cloudflare as Secondary](https://developers.cloudflare.com/dns/zone-setups/zone-transfers/cloudflare-as-secondary/setup/#dnssec)
	// for details.
	DnssecPresigned param.Field[bool] `json:"dnssec_presigned"`
	// Status of DNSSEC, based on user-desired state and presence of necessary records.
	Status param.Field[ZoneDnssecUpdateParamsStatus] `json:"status"`
}

func (r ZoneDnssecUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Status of DNSSEC, based on user-desired state and presence of necessary records.
type ZoneDnssecUpdateParamsStatus string

const (
	ZoneDnssecUpdateParamsStatusActive   ZoneDnssecUpdateParamsStatus = "active"
	ZoneDnssecUpdateParamsStatusDisabled ZoneDnssecUpdateParamsStatus = "disabled"
)
