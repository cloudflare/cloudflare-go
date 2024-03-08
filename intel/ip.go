// File generated from our OpenAPI spec by Stainless.

package intel

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/tidwall/gjson"
)

// IPService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewIPService] method instead.
type IPService struct {
	Options []option.RequestOption
}

// NewIPService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewIPService(opts ...option.RequestOption) (r *IPService) {
	r = &IPService{}
	r.Options = opts
	return
}

// Get IP Overview
func (r *IPService) Get(ctx context.Context, params IPGetParams, opts ...option.RequestOption) (res *[]IntelSchemasIP, err error) {
	opts = append(r.Options[:], opts...)
	var env IPGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/ip", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IntelSchemasIP struct {
	// Specifies a reference to the autonomous systems (AS) that the IP address belongs
	// to.
	BelongsToRef IntelSchemasIPBelongsToRef `json:"belongs_to_ref"`
	IP           IntelSchemasIPIP           `json:"ip" format:"ipv4"`
	RiskTypes    interface{}                `json:"risk_types"`
	JSON         intelSchemasIPJSON         `json:"-"`
}

// intelSchemasIPJSON contains the JSON metadata for the struct [IntelSchemasIP]
type intelSchemasIPJSON struct {
	BelongsToRef apijson.Field
	IP           apijson.Field
	RiskTypes    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IntelSchemasIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelSchemasIPJSON) RawJSON() string {
	return r.raw
}

// Specifies a reference to the autonomous systems (AS) that the IP address belongs
// to.
type IntelSchemasIPBelongsToRef struct {
	ID          interface{} `json:"id"`
	Country     string      `json:"country"`
	Description string      `json:"description"`
	// Infrastructure type of this ASN.
	Type  IntelSchemasIPBelongsToRefType `json:"type"`
	Value string                         `json:"value"`
	JSON  intelSchemasIPBelongsToRefJSON `json:"-"`
}

// intelSchemasIPBelongsToRefJSON contains the JSON metadata for the struct
// [IntelSchemasIPBelongsToRef]
type intelSchemasIPBelongsToRefJSON struct {
	ID          apijson.Field
	Country     apijson.Field
	Description apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelSchemasIPBelongsToRef) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelSchemasIPBelongsToRefJSON) RawJSON() string {
	return r.raw
}

// Infrastructure type of this ASN.
type IntelSchemasIPBelongsToRefType string

const (
	IntelSchemasIPBelongsToRefTypeHostingProvider IntelSchemasIPBelongsToRefType = "hosting_provider"
	IntelSchemasIPBelongsToRefTypeIsp             IntelSchemasIPBelongsToRefType = "isp"
	IntelSchemasIPBelongsToRefTypeOrganization    IntelSchemasIPBelongsToRefType = "organization"
)

// Union satisfied by [shared.UnionString] or [shared.UnionString].
type IntelSchemasIPIP interface {
	ImplementsIntelIntelSchemasIpip()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IntelSchemasIPIP)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type IPGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	IPV4      param.Field[string] `query:"ipv4"`
	IPV6      param.Field[string] `query:"ipv6"`
}

// URLQuery serializes [IPGetParams]'s query parameters as `url.Values`.
func (r IPGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IPGetResponseEnvelope struct {
	Errors   []IPGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IPGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []IntelSchemasIP                `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    IPGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo IPGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       ipGetResponseEnvelopeJSON       `json:"-"`
}

// ipGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [IPGetResponseEnvelope]
type ipGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type IPGetResponseEnvelopeErrors struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    ipGetResponseEnvelopeErrorsJSON `json:"-"`
}

// ipGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [IPGetResponseEnvelopeErrors]
type ipGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IPGetResponseEnvelopeMessages struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    ipGetResponseEnvelopeMessagesJSON `json:"-"`
}

// ipGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [IPGetResponseEnvelopeMessages]
type ipGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IPGetResponseEnvelopeSuccess bool

const (
	IPGetResponseEnvelopeSuccessTrue IPGetResponseEnvelopeSuccess = true
)

type IPGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                             `json:"total_count"`
	JSON       ipGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// ipGetResponseEnvelopeResultInfoJSON contains the JSON metadata for the struct
// [IPGetResponseEnvelopeResultInfo]
type ipGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
