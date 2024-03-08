// File generated from our OpenAPI spec by Stainless.

package email_routing

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/tidwall/gjson"
)

// RoutingDNSService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewRoutingDNSService] method instead.
type RoutingDNSService struct {
	Options []option.RequestOption
}

// NewRoutingDNSService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRoutingDNSService(opts ...option.RequestOption) (r *RoutingDNSService) {
	r = &RoutingDNSService{}
	r.Options = opts
	return
}

// Show the DNS records needed to configure your Email Routing zone.
func (r *RoutingDNSService) Get(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *[]RoutingDNSGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RoutingDNSGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/dns", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List of records needed to enable an Email Routing zone.
type RoutingDNSGetResponse struct {
	// DNS record content.
	Content string `json:"content"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name"`
	// Required for MX, SRV and URI records. Unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// Time to live, in seconds, of the DNS record. Must be between 60 and 86400, or 1
	// for 'automatic'.
	TTL RoutingDNSGetResponseTTL `json:"ttl"`
	// DNS record type.
	Type RoutingDNSGetResponseType `json:"type"`
	JSON routingDNSGetResponseJSON `json:"-"`
}

// routingDNSGetResponseJSON contains the JSON metadata for the struct
// [RoutingDNSGetResponse]
type routingDNSGetResponseJSON struct {
	Content     apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingDNSGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingDNSGetResponseJSON) RawJSON() string {
	return r.raw
}

// Time to live, in seconds, of the DNS record. Must be between 60 and 86400, or 1
// for 'automatic'.
//
// Union satisfied by [shared.UnionFloat] or
// [email_routing.RoutingDNSGetResponseTTLNumber].
type RoutingDNSGetResponseTTL interface {
	ImplementsEmailRoutingRoutingDNSGetResponseTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RoutingDNSGetResponseTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RoutingDNSGetResponseTTLNumber(0)),
		},
	)
}

type RoutingDNSGetResponseTTLNumber float64

const (
	RoutingDNSGetResponseTTLNumber1 RoutingDNSGetResponseTTLNumber = 1
)

// DNS record type.
type RoutingDNSGetResponseType string

const (
	RoutingDNSGetResponseTypeA      RoutingDNSGetResponseType = "A"
	RoutingDNSGetResponseTypeAAAA   RoutingDNSGetResponseType = "AAAA"
	RoutingDNSGetResponseTypeCNAME  RoutingDNSGetResponseType = "CNAME"
	RoutingDNSGetResponseTypeHTTPS  RoutingDNSGetResponseType = "HTTPS"
	RoutingDNSGetResponseTypeTXT    RoutingDNSGetResponseType = "TXT"
	RoutingDNSGetResponseTypeSRV    RoutingDNSGetResponseType = "SRV"
	RoutingDNSGetResponseTypeLOC    RoutingDNSGetResponseType = "LOC"
	RoutingDNSGetResponseTypeMX     RoutingDNSGetResponseType = "MX"
	RoutingDNSGetResponseTypeNS     RoutingDNSGetResponseType = "NS"
	RoutingDNSGetResponseTypeCert   RoutingDNSGetResponseType = "CERT"
	RoutingDNSGetResponseTypeDNSKEY RoutingDNSGetResponseType = "DNSKEY"
	RoutingDNSGetResponseTypeDS     RoutingDNSGetResponseType = "DS"
	RoutingDNSGetResponseTypeNAPTR  RoutingDNSGetResponseType = "NAPTR"
	RoutingDNSGetResponseTypeSmimea RoutingDNSGetResponseType = "SMIMEA"
	RoutingDNSGetResponseTypeSSHFP  RoutingDNSGetResponseType = "SSHFP"
	RoutingDNSGetResponseTypeSVCB   RoutingDNSGetResponseType = "SVCB"
	RoutingDNSGetResponseTypeTLSA   RoutingDNSGetResponseType = "TLSA"
	RoutingDNSGetResponseTypeURI    RoutingDNSGetResponseType = "URI"
)

type RoutingDNSGetResponseEnvelope struct {
	Errors   []RoutingDNSGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RoutingDNSGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []RoutingDNSGetResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    RoutingDNSGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo RoutingDNSGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       routingDNSGetResponseEnvelopeJSON       `json:"-"`
}

// routingDNSGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RoutingDNSGetResponseEnvelope]
type routingDNSGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingDNSGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingDNSGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RoutingDNSGetResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    routingDNSGetResponseEnvelopeErrorsJSON `json:"-"`
}

// routingDNSGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RoutingDNSGetResponseEnvelopeErrors]
type routingDNSGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingDNSGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingDNSGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RoutingDNSGetResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    routingDNSGetResponseEnvelopeMessagesJSON `json:"-"`
}

// routingDNSGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RoutingDNSGetResponseEnvelopeMessages]
type routingDNSGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingDNSGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingDNSGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RoutingDNSGetResponseEnvelopeSuccess bool

const (
	RoutingDNSGetResponseEnvelopeSuccessTrue RoutingDNSGetResponseEnvelopeSuccess = true
)

type RoutingDNSGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                     `json:"total_count"`
	JSON       routingDNSGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// routingDNSGetResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [RoutingDNSGetResponseEnvelopeResultInfo]
type routingDNSGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingDNSGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingDNSGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
