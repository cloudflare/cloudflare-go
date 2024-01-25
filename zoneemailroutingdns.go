// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// ZoneEmailRoutingDNSService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneEmailRoutingDNSService]
// method instead.
type ZoneEmailRoutingDNSService struct {
	Options []option.RequestOption
}

// NewZoneEmailRoutingDNSService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneEmailRoutingDNSService(opts ...option.RequestOption) (r *ZoneEmailRoutingDNSService) {
	r = &ZoneEmailRoutingDNSService{}
	r.Options = opts
	return
}

// Show the DNS records needed to configure your Email Routing zone.
func (r *ZoneEmailRoutingDNSService) EmailRoutingSettingsEmailRoutingDNSSettings(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/email/routing/dns", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponse struct {
	Errors     []ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseError    `json:"errors"`
	Messages   []ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseMessage  `json:"messages"`
	Result     []ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResult   `json:"result"`
	ResultInfo ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseSuccess `json:"success"`
	JSON    zoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseJSON    `json:"-"`
}

// zoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseJSON
// contains the JSON metadata for the struct
// [ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponse]
type zoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseError struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    zoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseErrorJSON `json:"-"`
}

// zoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseError]
type zoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseMessage struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    zoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseMessageJSON `json:"-"`
}

// zoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseMessage]
type zoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// List of records needed to enable an Email Routing zone.
type ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResult struct {
	// DNS record content.
	Content string `json:"content"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name"`
	// Required for MX, SRV and URI records. Unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// Time to live, in seconds, of the DNS record. Must be between 60 and 86400, or 1
	// for 'automatic'.
	Ttl ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTtl `json:"ttl"`
	// DNS record type.
	Type ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType `json:"type"`
	JSON zoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultJSON `json:"-"`
}

// zoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResult]
type zoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultJSON struct {
	Content     apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	Ttl         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time to live, in seconds, of the DNS record. Must be between 60 and 86400, or 1
// for 'automatic'.
//
// Union satisfied by [shared.UnionFloat] or
// [ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTtlNumber].
type ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTtl interface {
	ImplementsZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTtlNumber float64

const (
	ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTtlNumber1 ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTtlNumber = 1
)

// DNS record type.
type ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType string

const (
	ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTypeA      ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType = "A"
	ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTypeAaaa   ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType = "AAAA"
	ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTypeCname  ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType = "CNAME"
	ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTypeHTTPs  ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType = "HTTPS"
	ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTypeTxt    ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType = "TXT"
	ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTypeSrv    ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType = "SRV"
	ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTypeLoc    ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType = "LOC"
	ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTypeMx     ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType = "MX"
	ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTypeNs     ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType = "NS"
	ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTypeCert   ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType = "CERT"
	ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTypeDnskey ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType = "DNSKEY"
	ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTypeDs     ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType = "DS"
	ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTypeNaptr  ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType = "NAPTR"
	ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTypeSmimea ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType = "SMIMEA"
	ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTypeSshfp  ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType = "SSHFP"
	ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTypeSvcb   ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType = "SVCB"
	ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTypeTlsa   ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType = "TLSA"
	ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTypeUri    ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType = "URI"
)

type ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                              `json:"total_count"`
	JSON       zoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultInfoJSON `json:"-"`
}

// zoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultInfoJSON
// contains the JSON metadata for the struct
// [ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultInfo]
type zoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseSuccess bool

const (
	ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseSuccessTrue ZoneEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseSuccess = true
)
