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
func (r *ZoneEmailRoutingDNSService) EmailRoutingSettingsEmailRoutingDNSSettings(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *DNSSettingsResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/email/routing/dns", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type DNSSettingsResponseCollection struct {
	Errors     []DNSSettingsResponseCollectionError    `json:"errors"`
	Messages   []DNSSettingsResponseCollectionMessage  `json:"messages"`
	Result     []DNSSettingsResponseCollectionResult   `json:"result"`
	ResultInfo DNSSettingsResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success DNSSettingsResponseCollectionSuccess `json:"success"`
	JSON    dnsSettingsResponseCollectionJSON    `json:"-"`
}

// dnsSettingsResponseCollectionJSON contains the JSON metadata for the struct
// [DNSSettingsResponseCollection]
type dnsSettingsResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSettingsResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSSettingsResponseCollectionError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    dnsSettingsResponseCollectionErrorJSON `json:"-"`
}

// dnsSettingsResponseCollectionErrorJSON contains the JSON metadata for the struct
// [DNSSettingsResponseCollectionError]
type dnsSettingsResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSettingsResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSSettingsResponseCollectionMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    dnsSettingsResponseCollectionMessageJSON `json:"-"`
}

// dnsSettingsResponseCollectionMessageJSON contains the JSON metadata for the
// struct [DNSSettingsResponseCollectionMessage]
type dnsSettingsResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSettingsResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// List of records needed to enable an Email Routing zone.
type DNSSettingsResponseCollectionResult struct {
	// DNS record content.
	Content string `json:"content"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name"`
	// Required for MX, SRV and URI records. Unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// Time to live, in seconds, of the DNS record. Must be between 60 and 86400, or 1
	// for 'automatic'.
	Ttl DNSSettingsResponseCollectionResultTtl `json:"ttl"`
	// DNS record type.
	Type DNSSettingsResponseCollectionResultType `json:"type"`
	JSON dnsSettingsResponseCollectionResultJSON `json:"-"`
}

// dnsSettingsResponseCollectionResultJSON contains the JSON metadata for the
// struct [DNSSettingsResponseCollectionResult]
type dnsSettingsResponseCollectionResultJSON struct {
	Content     apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	Ttl         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSettingsResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time to live, in seconds, of the DNS record. Must be between 60 and 86400, or 1
// for 'automatic'.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSSettingsResponseCollectionResultTtlNumber].
type DNSSettingsResponseCollectionResultTtl interface {
	ImplementsDNSSettingsResponseCollectionResultTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSSettingsResponseCollectionResultTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSSettingsResponseCollectionResultTtlNumber float64

const (
	DNSSettingsResponseCollectionResultTtlNumber1 DNSSettingsResponseCollectionResultTtlNumber = 1
)

// DNS record type.
type DNSSettingsResponseCollectionResultType string

const (
	DNSSettingsResponseCollectionResultTypeA      DNSSettingsResponseCollectionResultType = "A"
	DNSSettingsResponseCollectionResultTypeAaaa   DNSSettingsResponseCollectionResultType = "AAAA"
	DNSSettingsResponseCollectionResultTypeCname  DNSSettingsResponseCollectionResultType = "CNAME"
	DNSSettingsResponseCollectionResultTypeHTTPs  DNSSettingsResponseCollectionResultType = "HTTPS"
	DNSSettingsResponseCollectionResultTypeTxt    DNSSettingsResponseCollectionResultType = "TXT"
	DNSSettingsResponseCollectionResultTypeSrv    DNSSettingsResponseCollectionResultType = "SRV"
	DNSSettingsResponseCollectionResultTypeLoc    DNSSettingsResponseCollectionResultType = "LOC"
	DNSSettingsResponseCollectionResultTypeMx     DNSSettingsResponseCollectionResultType = "MX"
	DNSSettingsResponseCollectionResultTypeNs     DNSSettingsResponseCollectionResultType = "NS"
	DNSSettingsResponseCollectionResultTypeCert   DNSSettingsResponseCollectionResultType = "CERT"
	DNSSettingsResponseCollectionResultTypeDnskey DNSSettingsResponseCollectionResultType = "DNSKEY"
	DNSSettingsResponseCollectionResultTypeDs     DNSSettingsResponseCollectionResultType = "DS"
	DNSSettingsResponseCollectionResultTypeNaptr  DNSSettingsResponseCollectionResultType = "NAPTR"
	DNSSettingsResponseCollectionResultTypeSmimea DNSSettingsResponseCollectionResultType = "SMIMEA"
	DNSSettingsResponseCollectionResultTypeSshfp  DNSSettingsResponseCollectionResultType = "SSHFP"
	DNSSettingsResponseCollectionResultTypeSvcb   DNSSettingsResponseCollectionResultType = "SVCB"
	DNSSettingsResponseCollectionResultTypeTlsa   DNSSettingsResponseCollectionResultType = "TLSA"
	DNSSettingsResponseCollectionResultTypeUri    DNSSettingsResponseCollectionResultType = "URI"
)

type DNSSettingsResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                     `json:"total_count"`
	JSON       dnsSettingsResponseCollectionResultInfoJSON `json:"-"`
}

// dnsSettingsResponseCollectionResultInfoJSON contains the JSON metadata for the
// struct [DNSSettingsResponseCollectionResultInfo]
type dnsSettingsResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSettingsResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DNSSettingsResponseCollectionSuccess bool

const (
	DNSSettingsResponseCollectionSuccessTrue DNSSettingsResponseCollectionSuccess = true
)
