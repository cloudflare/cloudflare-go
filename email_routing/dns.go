// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// DNSService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDNSService] method instead.
type DNSService struct {
	Options []option.RequestOption
}

// NewDNSService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDNSService(opts ...option.RequestOption) (r *DNSService) {
	r = &DNSService{}
	r.Options = opts
	return
}

// Enable you Email Routing zone. Add and lock the necessary MX and SPF records.
func (r *DNSService) New(ctx context.Context, params DNSNewParams, opts ...option.RequestOption) (res *Settings, err error) {
	var env DNSNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/email/routing/dns", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Disable your Email Routing zone. Also removes additional MX records previously
// required for Email Routing to work.
func (r *DNSService) Delete(ctx context.Context, body DNSDeleteParams, opts ...option.RequestOption) (res *Settings, err error) {
	var env DNSDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/email/routing/dns", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Unlock MX Records previously locked by Email Routing.
func (r *DNSService) Edit(ctx context.Context, params DNSEditParams, opts ...option.RequestOption) (res *Settings, err error) {
	var env DNSEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/email/routing/dns", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Show the DNS records needed to configure your Email Routing zone.
func (r *DNSService) Get(ctx context.Context, params DNSGetParams, opts ...option.RequestOption) (res *[]DNSRecord, err error) {
	var env DNSGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/email/routing/dns", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// List of records needed to enable an Email Routing zone.
type DNSRecord struct {
	// DNS record content.
	Content string `json:"content"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name"`
	// Required for MX, SRV and URI records. Unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// Time to live, in seconds, of the DNS record. Must be between 60 and 86400, or 1
	// for 'automatic'.
	TTL DNSRecordTTL `json:"ttl"`
	// DNS record type.
	Type DNSRecordType `json:"type"`
	JSON dnsRecordJSON `json:"-"`
}

// dnsRecordJSON contains the JSON metadata for the struct [DNSRecord]
type dnsRecordJSON struct {
	Content     apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordJSON) RawJSON() string {
	return r.raw
}

// Time to live, in seconds, of the DNS record. Must be between 60 and 86400, or 1
// for 'automatic'.
type DNSRecordTTL float64

const (
	DNSRecordTTL1 DNSRecordTTL = 1
)

func (r DNSRecordTTL) IsKnown() bool {
	switch r {
	case DNSRecordTTL1:
		return true
	}
	return false
}

// DNS record type.
type DNSRecordType string

const (
	DNSRecordTypeA      DNSRecordType = "A"
	DNSRecordTypeAAAA   DNSRecordType = "AAAA"
	DNSRecordTypeCNAME  DNSRecordType = "CNAME"
	DNSRecordTypeHTTPS  DNSRecordType = "HTTPS"
	DNSRecordTypeTXT    DNSRecordType = "TXT"
	DNSRecordTypeSRV    DNSRecordType = "SRV"
	DNSRecordTypeLOC    DNSRecordType = "LOC"
	DNSRecordTypeMX     DNSRecordType = "MX"
	DNSRecordTypeNS     DNSRecordType = "NS"
	DNSRecordTypeCERT   DNSRecordType = "CERT"
	DNSRecordTypeDNSKEY DNSRecordType = "DNSKEY"
	DNSRecordTypeDS     DNSRecordType = "DS"
	DNSRecordTypeNAPTR  DNSRecordType = "NAPTR"
	DNSRecordTypeSMIMEA DNSRecordType = "SMIMEA"
	DNSRecordTypeSSHFP  DNSRecordType = "SSHFP"
	DNSRecordTypeSVCB   DNSRecordType = "SVCB"
	DNSRecordTypeTLSA   DNSRecordType = "TLSA"
	DNSRecordTypeURI    DNSRecordType = "URI"
)

func (r DNSRecordType) IsKnown() bool {
	switch r {
	case DNSRecordTypeA, DNSRecordTypeAAAA, DNSRecordTypeCNAME, DNSRecordTypeHTTPS, DNSRecordTypeTXT, DNSRecordTypeSRV, DNSRecordTypeLOC, DNSRecordTypeMX, DNSRecordTypeNS, DNSRecordTypeCERT, DNSRecordTypeDNSKEY, DNSRecordTypeDS, DNSRecordTypeNAPTR, DNSRecordTypeSMIMEA, DNSRecordTypeSSHFP, DNSRecordTypeSVCB, DNSRecordTypeTLSA, DNSRecordTypeURI:
		return true
	}
	return false
}

type DNSNewParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// Domain of your zone.
	Name param.Field[string] `json:"name"`
}

func (r DNSNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSNewResponseEnvelope struct {
	Errors   []DNSNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DNSNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DNSNewResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  Settings                      `json:"result"`
	JSON    dnsNewResponseEnvelopeJSON    `json:"-"`
}

// dnsNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSNewResponseEnvelope]
type dnsNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DNSNewResponseEnvelopeErrors struct {
	Code             int64                              `json:"code" api:"required"`
	Message          string                             `json:"message" api:"required"`
	DocumentationURL string                             `json:"documentation_url"`
	Source           DNSNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             dnsNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// dnsNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [DNSNewResponseEnvelopeErrors]
type dnsNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DNSNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DNSNewResponseEnvelopeErrorsSource struct {
	Pointer string                                 `json:"pointer"`
	JSON    dnsNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dnsNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the struct
// [DNSNewResponseEnvelopeErrorsSource]
type dnsNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DNSNewResponseEnvelopeMessages struct {
	Code             int64                                `json:"code" api:"required"`
	Message          string                               `json:"message" api:"required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           DNSNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             dnsNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// dnsNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [DNSNewResponseEnvelopeMessages]
type dnsNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DNSNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DNSNewResponseEnvelopeMessagesSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    dnsNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dnsNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [DNSNewResponseEnvelopeMessagesSource]
type dnsNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DNSNewResponseEnvelopeSuccess bool

const (
	DNSNewResponseEnvelopeSuccessTrue DNSNewResponseEnvelopeSuccess = true
)

func (r DNSNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DNSNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DNSDeleteParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type DNSDeleteResponseEnvelope struct {
	Errors   []DNSDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DNSDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DNSDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  Settings                         `json:"result"`
	JSON    dnsDeleteResponseEnvelopeJSON    `json:"-"`
}

// dnsDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSDeleteResponseEnvelope]
type dnsDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DNSDeleteResponseEnvelopeErrors struct {
	Code             int64                                 `json:"code" api:"required"`
	Message          string                                `json:"message" api:"required"`
	DocumentationURL string                                `json:"documentation_url"`
	Source           DNSDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             dnsDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// dnsDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [DNSDeleteResponseEnvelopeErrors]
type dnsDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DNSDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DNSDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                    `json:"pointer"`
	JSON    dnsDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dnsDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [DNSDeleteResponseEnvelopeErrorsSource]
type dnsDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DNSDeleteResponseEnvelopeMessages struct {
	Code             int64                                   `json:"code" api:"required"`
	Message          string                                  `json:"message" api:"required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           DNSDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             dnsDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// dnsDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [DNSDeleteResponseEnvelopeMessages]
type dnsDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DNSDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DNSDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    dnsDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dnsDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [DNSDeleteResponseEnvelopeMessagesSource]
type dnsDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DNSDeleteResponseEnvelopeSuccess bool

const (
	DNSDeleteResponseEnvelopeSuccessTrue DNSDeleteResponseEnvelopeSuccess = true
)

func (r DNSDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DNSDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DNSEditParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// Domain of your zone.
	Name param.Field[string] `json:"name"`
}

func (r DNSEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSEditResponseEnvelope struct {
	Errors   []DNSEditResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DNSEditResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DNSEditResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  Settings                       `json:"result"`
	JSON    dnsEditResponseEnvelopeJSON    `json:"-"`
}

// dnsEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSEditResponseEnvelope]
type dnsEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DNSEditResponseEnvelopeErrors struct {
	Code             int64                               `json:"code" api:"required"`
	Message          string                              `json:"message" api:"required"`
	DocumentationURL string                              `json:"documentation_url"`
	Source           DNSEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             dnsEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// dnsEditResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [DNSEditResponseEnvelopeErrors]
type dnsEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DNSEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DNSEditResponseEnvelopeErrorsSource struct {
	Pointer string                                  `json:"pointer"`
	JSON    dnsEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dnsEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [DNSEditResponseEnvelopeErrorsSource]
type dnsEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DNSEditResponseEnvelopeMessages struct {
	Code             int64                                 `json:"code" api:"required"`
	Message          string                                `json:"message" api:"required"`
	DocumentationURL string                                `json:"documentation_url"`
	Source           DNSEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             dnsEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// dnsEditResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [DNSEditResponseEnvelopeMessages]
type dnsEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DNSEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DNSEditResponseEnvelopeMessagesSource struct {
	Pointer string                                    `json:"pointer"`
	JSON    dnsEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dnsEditResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [DNSEditResponseEnvelopeMessagesSource]
type dnsEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DNSEditResponseEnvelopeSuccess bool

const (
	DNSEditResponseEnvelopeSuccessTrue DNSEditResponseEnvelopeSuccess = true
)

func (r DNSEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DNSEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DNSGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// Deprecated. When supplied, the response shape differs from the documented
	// default and is not modeled in generated SDKs. Do not rely on this parameter.
	Subdomain param.Field[string] `query:"subdomain"`
}

// URLQuery serializes [DNSGetParams]'s query parameters as `url.Values`.
func (r DNSGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DNSGetResponseEnvelope struct {
	Errors   []DNSGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DNSGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success    DNSGetResponseEnvelopeSuccess    `json:"success" api:"required"`
	Result     []DNSRecord                      `json:"result"`
	ResultInfo DNSGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       dnsGetResponseEnvelopeJSON       `json:"-"`
}

// dnsGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSGetResponseEnvelope]
type dnsGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DNSGetResponseEnvelopeErrors struct {
	Code             int64                              `json:"code" api:"required"`
	Message          string                             `json:"message" api:"required"`
	DocumentationURL string                             `json:"documentation_url"`
	Source           DNSGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dnsGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dnsGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [DNSGetResponseEnvelopeErrors]
type dnsGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DNSGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DNSGetResponseEnvelopeErrorsSource struct {
	Pointer string                                 `json:"pointer"`
	JSON    dnsGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dnsGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the struct
// [DNSGetResponseEnvelopeErrorsSource]
type dnsGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DNSGetResponseEnvelopeMessages struct {
	Code             int64                                `json:"code" api:"required"`
	Message          string                               `json:"message" api:"required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           DNSGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dnsGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dnsGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [DNSGetResponseEnvelopeMessages]
type dnsGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DNSGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DNSGetResponseEnvelopeMessagesSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    dnsGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dnsGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [DNSGetResponseEnvelopeMessagesSource]
type dnsGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DNSGetResponseEnvelopeSuccess bool

const (
	DNSGetResponseEnvelopeSuccessTrue DNSGetResponseEnvelopeSuccess = true
)

func (r DNSGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DNSGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DNSGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64 `json:"total_count"`
	// The number of total pages in the entire result set.
	TotalPages float64                              `json:"total_pages"`
	JSON       dnsGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// dnsGetResponseEnvelopeResultInfoJSON contains the JSON metadata for the struct
// [DNSGetResponseEnvelopeResultInfo]
type dnsGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
