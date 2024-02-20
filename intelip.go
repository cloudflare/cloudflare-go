// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// IntelIPService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewIntelIPService] method instead.
type IntelIPService struct {
	Options []option.RequestOption
}

// NewIntelIPService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewIntelIPService(opts ...option.RequestOption) (r *IntelIPService) {
	r = &IntelIPService{}
	r.Options = opts
	return
}

// Get IP Overview
func (r *IntelIPService) Get(ctx context.Context, accountID string, query IntelIPGetParams, opts ...option.RequestOption) (res *[]IntelIPGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IntelIPGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/ip", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IntelIPGetResponse struct {
	// Specifies a reference to the autonomous systems (AS) that the IP address belongs
	// to.
	BelongsToRef IntelIPGetResponseBelongsToRef `json:"belongs_to_ref"`
	IP           IntelIPGetResponseIP           `json:"ip" format:"ipv4"`
	RiskTypes    interface{}                    `json:"risk_types"`
	JSON         intelIPGetResponseJSON         `json:"-"`
}

// intelIPGetResponseJSON contains the JSON metadata for the struct
// [IntelIPGetResponse]
type intelIPGetResponseJSON struct {
	BelongsToRef apijson.Field
	IP           apijson.Field
	RiskTypes    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IntelIPGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies a reference to the autonomous systems (AS) that the IP address belongs
// to.
type IntelIPGetResponseBelongsToRef struct {
	ID          interface{} `json:"id"`
	Country     string      `json:"country"`
	Description string      `json:"description"`
	// Infrastructure type of this ASN.
	Type  IntelIPGetResponseBelongsToRefType `json:"type"`
	Value string                             `json:"value"`
	JSON  intelIPGetResponseBelongsToRefJSON `json:"-"`
}

// intelIPGetResponseBelongsToRefJSON contains the JSON metadata for the struct
// [IntelIPGetResponseBelongsToRef]
type intelIPGetResponseBelongsToRefJSON struct {
	ID          apijson.Field
	Country     apijson.Field
	Description apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIPGetResponseBelongsToRef) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Infrastructure type of this ASN.
type IntelIPGetResponseBelongsToRefType string

const (
	IntelIPGetResponseBelongsToRefTypeHostingProvider IntelIPGetResponseBelongsToRefType = "hosting_provider"
	IntelIPGetResponseBelongsToRefTypeIsp             IntelIPGetResponseBelongsToRefType = "isp"
	IntelIPGetResponseBelongsToRefTypeOrganization    IntelIPGetResponseBelongsToRefType = "organization"
)

// Union satisfied by [shared.UnionString] or [shared.UnionString].
type IntelIPGetResponseIP interface {
	ImplementsIntelIPGetResponseIP()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IntelIPGetResponseIP)(nil)).Elem(),
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

type IntelIPGetParams struct {
	IPV4 param.Field[string] `query:"ipv4"`
	IPV6 param.Field[string] `query:"ipv6"`
}

// URLQuery serializes [IntelIPGetParams]'s query parameters as `url.Values`.
func (r IntelIPGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IntelIPGetResponseEnvelope struct {
	Errors   []IntelIPGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IntelIPGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []IntelIPGetResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    IntelIPGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo IntelIPGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       intelIPGetResponseEnvelopeJSON       `json:"-"`
}

// intelIPGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [IntelIPGetResponseEnvelope]
type intelIPGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIPGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIPGetResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    intelIPGetResponseEnvelopeErrorsJSON `json:"-"`
}

// intelIPGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [IntelIPGetResponseEnvelopeErrors]
type intelIPGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIPGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIPGetResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    intelIPGetResponseEnvelopeMessagesJSON `json:"-"`
}

// intelIPGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [IntelIPGetResponseEnvelopeMessages]
type intelIPGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIPGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type IntelIPGetResponseEnvelopeSuccess bool

const (
	IntelIPGetResponseEnvelopeSuccessTrue IntelIPGetResponseEnvelopeSuccess = true
)

type IntelIPGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                  `json:"total_count"`
	JSON       intelIPGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// intelIPGetResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [IntelIPGetResponseEnvelopeResultInfo]
type intelIPGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIPGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
