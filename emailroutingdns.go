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

// EmailRoutingDNSService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewEmailRoutingDNSService] method
// instead.
type EmailRoutingDNSService struct {
	Options []option.RequestOption
}

// NewEmailRoutingDNSService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmailRoutingDNSService(opts ...option.RequestOption) (r *EmailRoutingDNSService) {
	r = &EmailRoutingDNSService{}
	r.Options = opts
	return
}

// Show the DNS records needed to configure your Email Routing zone.
func (r *EmailRoutingDNSService) EmailRoutingSettingsEmailRoutingDNSSettings(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/email/routing/dns", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponse struct {
	Errors     []EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseError    `json:"errors"`
	Messages   []EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseMessage  `json:"messages"`
	Result     []EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResult   `json:"result"`
	ResultInfo EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseSuccess `json:"success"`
	JSON    emailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseJSON    `json:"-"`
}

// emailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseJSON contains
// the JSON metadata for the struct
// [EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponse]
type emailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseError struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    emailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseErrorJSON `json:"-"`
}

// emailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseErrorJSON
// contains the JSON metadata for the struct
// [EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseError]
type emailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseMessage struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    emailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseMessageJSON `json:"-"`
}

// emailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseMessageJSON
// contains the JSON metadata for the struct
// [EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseMessage]
type emailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// List of records needed to enable an Email Routing zone.
type EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResult struct {
	// DNS record content.
	Content string `json:"content"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name"`
	// Required for MX, SRV and URI records. Unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// Time to live, in seconds, of the DNS record. Must be between 60 and 86400, or 1
	// for 'automatic'.
	TTL EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTTL `json:"ttl"`
	// DNS record type.
	Type EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType `json:"type"`
	JSON emailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultJSON `json:"-"`
}

// emailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultJSON
// contains the JSON metadata for the struct
// [EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResult]
type emailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultJSON struct {
	Content     apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time to live, in seconds, of the DNS record. Must be between 60 and 86400, or 1
// for 'automatic'.
//
// Union satisfied by [shared.UnionFloat] or
// [EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTTLNumber].
type EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTTL interface {
	ImplementsEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTTLNumber float64

const (
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTTLNumber1 EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTTLNumber = 1
)

// DNS record type.
type EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType string

const (
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTypeA      EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType = "A"
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTypeAaaa   EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType = "AAAA"
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTypeCname  EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType = "CNAME"
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTypeHTTPs  EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType = "HTTPS"
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTypeTxt    EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType = "TXT"
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTypeSrv    EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType = "SRV"
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTypeLoc    EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType = "LOC"
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTypeMx     EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType = "MX"
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTypeNs     EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType = "NS"
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTypeCert   EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType = "CERT"
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTypeDnskey EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType = "DNSKEY"
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTypeDs     EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType = "DS"
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTypeNaptr  EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType = "NAPTR"
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTypeSmimea EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType = "SMIMEA"
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTypeSshfp  EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType = "SSHFP"
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTypeSvcb   EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType = "SVCB"
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTypeTlsa   EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType = "TLSA"
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultTypeUri    EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultType = "URI"
)

type EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                          `json:"total_count"`
	JSON       emailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultInfoJSON `json:"-"`
}

// emailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultInfoJSON
// contains the JSON metadata for the struct
// [EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultInfo]
type emailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseSuccess bool

const (
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseSuccessTrue EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseSuccess = true
)
