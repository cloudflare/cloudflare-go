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
func (r *EmailRoutingDNSService) EmailRoutingSettingsEmailRoutingDNSSettings(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *[]EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/dns", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List of records needed to enable an Email Routing zone.
type EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponse struct {
	// DNS record content.
	Content string `json:"content"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name"`
	// Required for MX, SRV and URI records. Unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// Time to live, in seconds, of the DNS record. Must be between 60 and 86400, or 1
	// for 'automatic'.
	TTL EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseTTL `json:"ttl"`
	// DNS record type.
	Type EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseType `json:"type"`
	JSON emailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseJSON `json:"-"`
}

// emailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseJSON contains
// the JSON metadata for the struct
// [EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponse]
type emailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseJSON struct {
	Content     apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time to live, in seconds, of the DNS record. Must be between 60 and 86400, or 1
// for 'automatic'.
//
// Union satisfied by [shared.UnionFloat] or
// [EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseTTLNumber].
type EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseTTL interface {
	ImplementsEmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseTTLNumber float64

const (
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseTTLNumber1 EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseTTLNumber = 1
)

// DNS record type.
type EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseType string

const (
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseTypeA      EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseType = "A"
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseTypeAaaa   EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseType = "AAAA"
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseTypeCname  EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseType = "CNAME"
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseTypeHTTPs  EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseType = "HTTPS"
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseTypeTxt    EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseType = "TXT"
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseTypeSrv    EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseType = "SRV"
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseTypeLoc    EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseType = "LOC"
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseTypeMx     EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseType = "MX"
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseTypeNs     EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseType = "NS"
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseTypeCert   EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseType = "CERT"
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseTypeDnskey EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseType = "DNSKEY"
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseTypeDs     EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseType = "DS"
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseTypeNaptr  EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseType = "NAPTR"
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseTypeSmimea EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseType = "SMIMEA"
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseTypeSshfp  EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseType = "SSHFP"
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseTypeSvcb   EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseType = "SVCB"
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseTypeTlsa   EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseType = "TLSA"
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseTypeUri    EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseType = "URI"
)

type EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseEnvelope struct {
	Errors   []EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseEnvelopeMessages `json:"messages,required"`
	Result   []EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseEnvelopeResultInfo `json:"result_info"`
	JSON       emailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseEnvelopeJSON       `json:"-"`
}

// emailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseEnvelope]
type emailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseEnvelopeErrors struct {
	Code    int64                                                                                `json:"code,required"`
	Message string                                                                               `json:"message,required"`
	JSON    emailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseEnvelopeErrors]
type emailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseEnvelopeMessages struct {
	Code    int64                                                                                  `json:"code,required"`
	Message string                                                                                 `json:"message,required"`
	JSON    emailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseEnvelopeMessages]
type emailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseEnvelopeSuccess bool

const (
	EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseEnvelopeSuccessTrue EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseEnvelopeSuccess = true
)

type EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                  `json:"total_count"`
	JSON       emailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseEnvelopeResultInfoJSON `json:"-"`
}

// emailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseEnvelopeResultInfo]
type emailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingDNSEmailRoutingSettingsEmailRoutingDNSSettingsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
