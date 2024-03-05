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

// EmailRoutingRoutingDNSService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewEmailRoutingRoutingDNSService]
// method instead.
type EmailRoutingRoutingDNSService struct {
	Options []option.RequestOption
}

// NewEmailRoutingRoutingDNSService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEmailRoutingRoutingDNSService(opts ...option.RequestOption) (r *EmailRoutingRoutingDNSService) {
	r = &EmailRoutingRoutingDNSService{}
	r.Options = opts
	return
}

// Show the DNS records needed to configure your Email Routing zone.
func (r *EmailRoutingRoutingDNSService) Get(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *[]EmailDNSRecord, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingRoutingDNSGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/dns", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List of records needed to enable an Email Routing zone.
type EmailDNSRecord struct {
	// DNS record content.
	Content string `json:"content"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name"`
	// Required for MX, SRV and URI records. Unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// Time to live, in seconds, of the DNS record. Must be between 60 and 86400, or 1
	// for 'automatic'.
	TTL EmailDNSRecordTTL `json:"ttl"`
	// DNS record type.
	Type EmailDNSRecordType `json:"type"`
	JSON emailDNSRecordJSON `json:"-"`
}

// emailDNSRecordJSON contains the JSON metadata for the struct [EmailDNSRecord]
type emailDNSRecordJSON struct {
	Content     apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailDNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time to live, in seconds, of the DNS record. Must be between 60 and 86400, or 1
// for 'automatic'.
//
// Union satisfied by [shared.UnionFloat] or [EmailDNSRecordTTLNumber].
type EmailDNSRecordTTL interface {
	ImplementsEmailDNSRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EmailDNSRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type EmailDNSRecordTTLNumber float64

const (
	EmailDNSRecordTTLNumber1 EmailDNSRecordTTLNumber = 1
)

// DNS record type.
type EmailDNSRecordType string

const (
	EmailDNSRecordTypeA      EmailDNSRecordType = "A"
	EmailDNSRecordTypeAaaa   EmailDNSRecordType = "AAAA"
	EmailDNSRecordTypeCname  EmailDNSRecordType = "CNAME"
	EmailDNSRecordTypeHTTPS  EmailDNSRecordType = "HTTPS"
	EmailDNSRecordTypeTxt    EmailDNSRecordType = "TXT"
	EmailDNSRecordTypeSrv    EmailDNSRecordType = "SRV"
	EmailDNSRecordTypeLoc    EmailDNSRecordType = "LOC"
	EmailDNSRecordTypeMx     EmailDNSRecordType = "MX"
	EmailDNSRecordTypeNs     EmailDNSRecordType = "NS"
	EmailDNSRecordTypeCert   EmailDNSRecordType = "CERT"
	EmailDNSRecordTypeDnskey EmailDNSRecordType = "DNSKEY"
	EmailDNSRecordTypeDs     EmailDNSRecordType = "DS"
	EmailDNSRecordTypeNaptr  EmailDNSRecordType = "NAPTR"
	EmailDNSRecordTypeSmimea EmailDNSRecordType = "SMIMEA"
	EmailDNSRecordTypeSshfp  EmailDNSRecordType = "SSHFP"
	EmailDNSRecordTypeSvcb   EmailDNSRecordType = "SVCB"
	EmailDNSRecordTypeTlsa   EmailDNSRecordType = "TLSA"
	EmailDNSRecordTypeUri    EmailDNSRecordType = "URI"
)

type EmailRoutingRoutingDNSGetResponseEnvelope struct {
	Errors   []EmailRoutingRoutingDNSGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EmailRoutingRoutingDNSGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []EmailDNSRecord                                    `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    EmailRoutingRoutingDNSGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo EmailRoutingRoutingDNSGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       emailRoutingRoutingDNSGetResponseEnvelopeJSON       `json:"-"`
}

// emailRoutingRoutingDNSGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [EmailRoutingRoutingDNSGetResponseEnvelope]
type emailRoutingRoutingDNSGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingDNSGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRoutingDNSGetResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    emailRoutingRoutingDNSGetResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingRoutingDNSGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [EmailRoutingRoutingDNSGetResponseEnvelopeErrors]
type emailRoutingRoutingDNSGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingDNSGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRoutingDNSGetResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    emailRoutingRoutingDNSGetResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingRoutingDNSGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [EmailRoutingRoutingDNSGetResponseEnvelopeMessages]
type emailRoutingRoutingDNSGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingDNSGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmailRoutingRoutingDNSGetResponseEnvelopeSuccess bool

const (
	EmailRoutingRoutingDNSGetResponseEnvelopeSuccessTrue EmailRoutingRoutingDNSGetResponseEnvelopeSuccess = true
)

type EmailRoutingRoutingDNSGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                 `json:"total_count"`
	JSON       emailRoutingRoutingDNSGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// emailRoutingRoutingDNSGetResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [EmailRoutingRoutingDNSGetResponseEnvelopeResultInfo]
type emailRoutingRoutingDNSGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingDNSGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
