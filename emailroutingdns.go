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
func (r *EmailRoutingDNSService) Get(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *[]EmailRoutingDNSGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingDNSGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/dns", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List of records needed to enable an Email Routing zone.
type EmailRoutingDNSGetResponse struct {
	// DNS record content.
	Content string `json:"content"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name"`
	// Required for MX, SRV and URI records. Unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// Time to live, in seconds, of the DNS record. Must be between 60 and 86400, or 1
	// for 'automatic'.
	TTL EmailRoutingDNSGetResponseTTL `json:"ttl"`
	// DNS record type.
	Type EmailRoutingDNSGetResponseType `json:"type"`
	JSON emailRoutingDNSGetResponseJSON `json:"-"`
}

// emailRoutingDNSGetResponseJSON contains the JSON metadata for the struct
// [EmailRoutingDNSGetResponse]
type emailRoutingDNSGetResponseJSON struct {
	Content     apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingDNSGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time to live, in seconds, of the DNS record. Must be between 60 and 86400, or 1
// for 'automatic'.
//
// Union satisfied by [shared.UnionFloat] or [EmailRoutingDNSGetResponseTTLNumber].
type EmailRoutingDNSGetResponseTTL interface {
	ImplementsEmailRoutingDNSGetResponseTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EmailRoutingDNSGetResponseTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type EmailRoutingDNSGetResponseTTLNumber float64

const (
	EmailRoutingDNSGetResponseTTLNumber1 EmailRoutingDNSGetResponseTTLNumber = 1
)

// DNS record type.
type EmailRoutingDNSGetResponseType string

const (
	EmailRoutingDNSGetResponseTypeA      EmailRoutingDNSGetResponseType = "A"
	EmailRoutingDNSGetResponseTypeAaaa   EmailRoutingDNSGetResponseType = "AAAA"
	EmailRoutingDNSGetResponseTypeCname  EmailRoutingDNSGetResponseType = "CNAME"
	EmailRoutingDNSGetResponseTypeHTTPS  EmailRoutingDNSGetResponseType = "HTTPS"
	EmailRoutingDNSGetResponseTypeTxt    EmailRoutingDNSGetResponseType = "TXT"
	EmailRoutingDNSGetResponseTypeSrv    EmailRoutingDNSGetResponseType = "SRV"
	EmailRoutingDNSGetResponseTypeLoc    EmailRoutingDNSGetResponseType = "LOC"
	EmailRoutingDNSGetResponseTypeMx     EmailRoutingDNSGetResponseType = "MX"
	EmailRoutingDNSGetResponseTypeNs     EmailRoutingDNSGetResponseType = "NS"
	EmailRoutingDNSGetResponseTypeCert   EmailRoutingDNSGetResponseType = "CERT"
	EmailRoutingDNSGetResponseTypeDnskey EmailRoutingDNSGetResponseType = "DNSKEY"
	EmailRoutingDNSGetResponseTypeDs     EmailRoutingDNSGetResponseType = "DS"
	EmailRoutingDNSGetResponseTypeNaptr  EmailRoutingDNSGetResponseType = "NAPTR"
	EmailRoutingDNSGetResponseTypeSmimea EmailRoutingDNSGetResponseType = "SMIMEA"
	EmailRoutingDNSGetResponseTypeSshfp  EmailRoutingDNSGetResponseType = "SSHFP"
	EmailRoutingDNSGetResponseTypeSvcb   EmailRoutingDNSGetResponseType = "SVCB"
	EmailRoutingDNSGetResponseTypeTlsa   EmailRoutingDNSGetResponseType = "TLSA"
	EmailRoutingDNSGetResponseTypeUri    EmailRoutingDNSGetResponseType = "URI"
)

type EmailRoutingDNSGetResponseEnvelope struct {
	Errors   []EmailRoutingDNSGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EmailRoutingDNSGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []EmailRoutingDNSGetResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    EmailRoutingDNSGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo EmailRoutingDNSGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       emailRoutingDNSGetResponseEnvelopeJSON       `json:"-"`
}

// emailRoutingDNSGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [EmailRoutingDNSGetResponseEnvelope]
type emailRoutingDNSGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingDNSGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingDNSGetResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    emailRoutingDNSGetResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingDNSGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [EmailRoutingDNSGetResponseEnvelopeErrors]
type emailRoutingDNSGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingDNSGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingDNSGetResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    emailRoutingDNSGetResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingDNSGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [EmailRoutingDNSGetResponseEnvelopeMessages]
type emailRoutingDNSGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingDNSGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmailRoutingDNSGetResponseEnvelopeSuccess bool

const (
	EmailRoutingDNSGetResponseEnvelopeSuccessTrue EmailRoutingDNSGetResponseEnvelopeSuccess = true
)

type EmailRoutingDNSGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                          `json:"total_count"`
	JSON       emailRoutingDNSGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// emailRoutingDNSGetResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [EmailRoutingDNSGetResponseEnvelopeResultInfo]
type emailRoutingDNSGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingDNSGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
