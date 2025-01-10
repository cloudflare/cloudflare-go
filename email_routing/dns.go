// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
	"github.com/tidwall/gjson"
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
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/email/routing/dns", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Disable your Email Routing zone. Also removes additional MX records previously
// required for Email Routing to work.
func (r *DNSService) Delete(ctx context.Context, body DNSDeleteParams, opts ...option.RequestOption) (res *DNSDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/email/routing/dns", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Unlock MX Records previously locked by Email Routing.
func (r *DNSService) Edit(ctx context.Context, params DNSEditParams, opts ...option.RequestOption) (res *Settings, err error) {
	var env DNSEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/email/routing/dns", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Show the DNS records needed to configure your Email Routing zone.
func (r *DNSService) Get(ctx context.Context, params DNSGetParams, opts ...option.RequestOption) (res *DNSGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/email/routing/dns", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
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

type DNSDeleteResponse struct {
	// This field can have the runtime type of [[]shared.ResponseInfo].
	Errors interface{} `json:"errors,required"`
	// This field can have the runtime type of [[]shared.ResponseInfo].
	Messages interface{} `json:"messages,required"`
	// Whether the API call was successful
	Success DNSDeleteResponseSuccess `json:"success,required"`
	// This field can have the runtime type of [[]DNSRecord].
	Result interface{} `json:"result"`
	// This field can have the runtime type of
	// [DNSDeleteResponseEmailDNSSettingsResponseCollectionResultInfo].
	ResultInfo interface{}           `json:"result_info"`
	JSON       dnsDeleteResponseJSON `json:"-"`
	union      DNSDeleteResponseUnion
}

// dnsDeleteResponseJSON contains the JSON metadata for the struct
// [DNSDeleteResponse]
type dnsDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r dnsDeleteResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DNSDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DNSDeleteResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DNSDeleteResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are
// [email_routing.DNSDeleteResponseEmailAPIResponseCommon],
// [email_routing.DNSDeleteResponseEmailDNSSettingsResponseCollection].
func (r DNSDeleteResponse) AsUnion() DNSDeleteResponseUnion {
	return r.union
}

// Union satisfied by [email_routing.DNSDeleteResponseEmailAPIResponseCommon] or
// [email_routing.DNSDeleteResponseEmailDNSSettingsResponseCollection].
type DNSDeleteResponseUnion interface {
	implementsEmailRoutingDNSDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSDeleteResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DNSDeleteResponseEmailAPIResponseCommon{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DNSDeleteResponseEmailDNSSettingsResponseCollection{}),
		},
	)
}

type DNSDeleteResponseEmailAPIResponseCommon struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DNSDeleteResponseEmailAPIResponseCommonSuccess `json:"success,required"`
	JSON    dnsDeleteResponseEmailAPIResponseCommonJSON    `json:"-"`
}

// dnsDeleteResponseEmailAPIResponseCommonJSON contains the JSON metadata for the
// struct [DNSDeleteResponseEmailAPIResponseCommon]
type dnsDeleteResponseEmailAPIResponseCommonJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSDeleteResponseEmailAPIResponseCommon) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsDeleteResponseEmailAPIResponseCommonJSON) RawJSON() string {
	return r.raw
}

func (r DNSDeleteResponseEmailAPIResponseCommon) implementsEmailRoutingDNSDeleteResponse() {}

// Whether the API call was successful
type DNSDeleteResponseEmailAPIResponseCommonSuccess bool

const (
	DNSDeleteResponseEmailAPIResponseCommonSuccessTrue DNSDeleteResponseEmailAPIResponseCommonSuccess = true
)

func (r DNSDeleteResponseEmailAPIResponseCommonSuccess) IsKnown() bool {
	switch r {
	case DNSDeleteResponseEmailAPIResponseCommonSuccessTrue:
		return true
	}
	return false
}

type DNSDeleteResponseEmailDNSSettingsResponseCollection struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success    DNSDeleteResponseEmailDNSSettingsResponseCollectionSuccess    `json:"success,required"`
	Result     []DNSRecord                                                   `json:"result"`
	ResultInfo DNSDeleteResponseEmailDNSSettingsResponseCollectionResultInfo `json:"result_info"`
	JSON       dnsDeleteResponseEmailDNSSettingsResponseCollectionJSON       `json:"-"`
}

// dnsDeleteResponseEmailDNSSettingsResponseCollectionJSON contains the JSON
// metadata for the struct [DNSDeleteResponseEmailDNSSettingsResponseCollection]
type dnsDeleteResponseEmailDNSSettingsResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSDeleteResponseEmailDNSSettingsResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsDeleteResponseEmailDNSSettingsResponseCollectionJSON) RawJSON() string {
	return r.raw
}

func (r DNSDeleteResponseEmailDNSSettingsResponseCollection) implementsEmailRoutingDNSDeleteResponse() {
}

// Whether the API call was successful
type DNSDeleteResponseEmailDNSSettingsResponseCollectionSuccess bool

const (
	DNSDeleteResponseEmailDNSSettingsResponseCollectionSuccessTrue DNSDeleteResponseEmailDNSSettingsResponseCollectionSuccess = true
)

func (r DNSDeleteResponseEmailDNSSettingsResponseCollectionSuccess) IsKnown() bool {
	switch r {
	case DNSDeleteResponseEmailDNSSettingsResponseCollectionSuccessTrue:
		return true
	}
	return false
}

type DNSDeleteResponseEmailDNSSettingsResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                           `json:"total_count"`
	JSON       dnsDeleteResponseEmailDNSSettingsResponseCollectionResultInfoJSON `json:"-"`
}

// dnsDeleteResponseEmailDNSSettingsResponseCollectionResultInfoJSON contains the
// JSON metadata for the struct
// [DNSDeleteResponseEmailDNSSettingsResponseCollectionResultInfo]
type dnsDeleteResponseEmailDNSSettingsResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSDeleteResponseEmailDNSSettingsResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsDeleteResponseEmailDNSSettingsResponseCollectionResultInfoJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DNSDeleteResponseSuccess bool

const (
	DNSDeleteResponseSuccessTrue DNSDeleteResponseSuccess = true
)

func (r DNSDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case DNSDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type DNSGetResponse struct {
	// This field can have the runtime type of [[]shared.ResponseInfo].
	Errors interface{} `json:"errors,required"`
	// This field can have the runtime type of [[]shared.ResponseInfo].
	Messages interface{} `json:"messages,required"`
	// Whether the API call was successful
	Success DNSGetResponseSuccess `json:"success,required"`
	// This field can have the runtime type of
	// [DNSGetResponseEmailEmailRoutingDNSQueryResponseResult], [[]DNSRecord].
	Result interface{} `json:"result"`
	// This field can have the runtime type of
	// [DNSGetResponseEmailEmailRoutingDNSQueryResponseResultInfo],
	// [DNSGetResponseEmailDNSSettingsResponseCollectionResultInfo].
	ResultInfo interface{}        `json:"result_info"`
	JSON       dnsGetResponseJSON `json:"-"`
	union      DNSGetResponseUnion
}

// dnsGetResponseJSON contains the JSON metadata for the struct [DNSGetResponse]
type dnsGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r dnsGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DNSGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DNSGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DNSGetResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are
// [email_routing.DNSGetResponseEmailEmailRoutingDNSQueryResponse],
// [email_routing.DNSGetResponseEmailDNSSettingsResponseCollection].
func (r DNSGetResponse) AsUnion() DNSGetResponseUnion {
	return r.union
}

// Union satisfied by
// [email_routing.DNSGetResponseEmailEmailRoutingDNSQueryResponse] or
// [email_routing.DNSGetResponseEmailDNSSettingsResponseCollection].
type DNSGetResponseUnion interface {
	implementsEmailRoutingDNSGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DNSGetResponseEmailEmailRoutingDNSQueryResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DNSGetResponseEmailDNSSettingsResponseCollection{}),
		},
	)
}

type DNSGetResponseEmailEmailRoutingDNSQueryResponse struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success    DNSGetResponseEmailEmailRoutingDNSQueryResponseSuccess    `json:"success,required"`
	Result     DNSGetResponseEmailEmailRoutingDNSQueryResponseResult     `json:"result"`
	ResultInfo DNSGetResponseEmailEmailRoutingDNSQueryResponseResultInfo `json:"result_info"`
	JSON       dnsGetResponseEmailEmailRoutingDNSQueryResponseJSON       `json:"-"`
}

// dnsGetResponseEmailEmailRoutingDNSQueryResponseJSON contains the JSON metadata
// for the struct [DNSGetResponseEmailEmailRoutingDNSQueryResponse]
type dnsGetResponseEmailEmailRoutingDNSQueryResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSGetResponseEmailEmailRoutingDNSQueryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsGetResponseEmailEmailRoutingDNSQueryResponseJSON) RawJSON() string {
	return r.raw
}

func (r DNSGetResponseEmailEmailRoutingDNSQueryResponse) implementsEmailRoutingDNSGetResponse() {}

// Whether the API call was successful
type DNSGetResponseEmailEmailRoutingDNSQueryResponseSuccess bool

const (
	DNSGetResponseEmailEmailRoutingDNSQueryResponseSuccessTrue DNSGetResponseEmailEmailRoutingDNSQueryResponseSuccess = true
)

func (r DNSGetResponseEmailEmailRoutingDNSQueryResponseSuccess) IsKnown() bool {
	switch r {
	case DNSGetResponseEmailEmailRoutingDNSQueryResponseSuccessTrue:
		return true
	}
	return false
}

type DNSGetResponseEmailEmailRoutingDNSQueryResponseResult struct {
	Errors []DNSGetResponseEmailEmailRoutingDNSQueryResponseResultError `json:"errors"`
	Record []DNSRecord                                                  `json:"record"`
	JSON   dnsGetResponseEmailEmailRoutingDNSQueryResponseResultJSON    `json:"-"`
}

// dnsGetResponseEmailEmailRoutingDNSQueryResponseResultJSON contains the JSON
// metadata for the struct [DNSGetResponseEmailEmailRoutingDNSQueryResponseResult]
type dnsGetResponseEmailEmailRoutingDNSQueryResponseResultJSON struct {
	Errors      apijson.Field
	Record      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSGetResponseEmailEmailRoutingDNSQueryResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsGetResponseEmailEmailRoutingDNSQueryResponseResultJSON) RawJSON() string {
	return r.raw
}

type DNSGetResponseEmailEmailRoutingDNSQueryResponseResultError struct {
	Code string `json:"code"`
	// List of records needed to enable an Email Routing zone.
	Missing DNSRecord                                                      `json:"missing"`
	JSON    dnsGetResponseEmailEmailRoutingDNSQueryResponseResultErrorJSON `json:"-"`
}

// dnsGetResponseEmailEmailRoutingDNSQueryResponseResultErrorJSON contains the JSON
// metadata for the struct
// [DNSGetResponseEmailEmailRoutingDNSQueryResponseResultError]
type dnsGetResponseEmailEmailRoutingDNSQueryResponseResultErrorJSON struct {
	Code        apijson.Field
	Missing     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSGetResponseEmailEmailRoutingDNSQueryResponseResultError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsGetResponseEmailEmailRoutingDNSQueryResponseResultErrorJSON) RawJSON() string {
	return r.raw
}

type DNSGetResponseEmailEmailRoutingDNSQueryResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                       `json:"total_count"`
	JSON       dnsGetResponseEmailEmailRoutingDNSQueryResponseResultInfoJSON `json:"-"`
}

// dnsGetResponseEmailEmailRoutingDNSQueryResponseResultInfoJSON contains the JSON
// metadata for the struct
// [DNSGetResponseEmailEmailRoutingDNSQueryResponseResultInfo]
type dnsGetResponseEmailEmailRoutingDNSQueryResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSGetResponseEmailEmailRoutingDNSQueryResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsGetResponseEmailEmailRoutingDNSQueryResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type DNSGetResponseEmailDNSSettingsResponseCollection struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success    DNSGetResponseEmailDNSSettingsResponseCollectionSuccess    `json:"success,required"`
	Result     []DNSRecord                                                `json:"result"`
	ResultInfo DNSGetResponseEmailDNSSettingsResponseCollectionResultInfo `json:"result_info"`
	JSON       dnsGetResponseEmailDNSSettingsResponseCollectionJSON       `json:"-"`
}

// dnsGetResponseEmailDNSSettingsResponseCollectionJSON contains the JSON metadata
// for the struct [DNSGetResponseEmailDNSSettingsResponseCollection]
type dnsGetResponseEmailDNSSettingsResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSGetResponseEmailDNSSettingsResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsGetResponseEmailDNSSettingsResponseCollectionJSON) RawJSON() string {
	return r.raw
}

func (r DNSGetResponseEmailDNSSettingsResponseCollection) implementsEmailRoutingDNSGetResponse() {}

// Whether the API call was successful
type DNSGetResponseEmailDNSSettingsResponseCollectionSuccess bool

const (
	DNSGetResponseEmailDNSSettingsResponseCollectionSuccessTrue DNSGetResponseEmailDNSSettingsResponseCollectionSuccess = true
)

func (r DNSGetResponseEmailDNSSettingsResponseCollectionSuccess) IsKnown() bool {
	switch r {
	case DNSGetResponseEmailDNSSettingsResponseCollectionSuccessTrue:
		return true
	}
	return false
}

type DNSGetResponseEmailDNSSettingsResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                        `json:"total_count"`
	JSON       dnsGetResponseEmailDNSSettingsResponseCollectionResultInfoJSON `json:"-"`
}

// dnsGetResponseEmailDNSSettingsResponseCollectionResultInfoJSON contains the JSON
// metadata for the struct
// [DNSGetResponseEmailDNSSettingsResponseCollectionResultInfo]
type dnsGetResponseEmailDNSSettingsResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSGetResponseEmailDNSSettingsResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsGetResponseEmailDNSSettingsResponseCollectionResultInfoJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DNSGetResponseSuccess bool

const (
	DNSGetResponseSuccessTrue DNSGetResponseSuccess = true
)

func (r DNSGetResponseSuccess) IsKnown() bool {
	switch r {
	case DNSGetResponseSuccessTrue:
		return true
	}
	return false
}

type DNSNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Domain of your zone.
	Name param.Field[string] `json:"name,required"`
}

func (r DNSNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DNSNewResponseEnvelopeSuccess `json:"success,required"`
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

// Whether the API call was successful
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
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type DNSEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Domain of your zone.
	Name param.Field[string] `json:"name,required"`
}

func (r DNSEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DNSEditResponseEnvelopeSuccess `json:"success,required"`
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

// Whether the API call was successful
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
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Domain of your zone.
	Subdomain param.Field[string] `query:"subdomain"`
}

// URLQuery serializes [DNSGetParams]'s query parameters as `url.Values`.
func (r DNSGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
