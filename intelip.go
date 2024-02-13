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
func (r *IntelIPService) IPIntelligenceGetIPOverview(ctx context.Context, accountID string, query IntelIPIPIntelligenceGetIPOverviewParams, opts ...option.RequestOption) (res *[]IntelIpipIntelligenceGetIPOverviewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IntelIpipIntelligenceGetIPOverviewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/ip", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IntelIpipIntelligenceGetIPOverviewResponse struct {
	// Specifies a reference to the autonomous systems (AS) that the IP address belongs
	// to.
	BelongsToRef IntelIpipIntelligenceGetIPOverviewResponseBelongsToRef `json:"belongs_to_ref"`
	IP           IntelIpipIntelligenceGetIPOverviewResponseIP           `json:"ip" format:"ipv4"`
	RiskTypes    interface{}                                            `json:"risk_types"`
	JSON         intelIpipIntelligenceGetIPOverviewResponseJSON         `json:"-"`
}

// intelIpipIntelligenceGetIPOverviewResponseJSON contains the JSON metadata for
// the struct [IntelIpipIntelligenceGetIPOverviewResponse]
type intelIpipIntelligenceGetIPOverviewResponseJSON struct {
	BelongsToRef apijson.Field
	IP           apijson.Field
	RiskTypes    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IntelIpipIntelligenceGetIPOverviewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies a reference to the autonomous systems (AS) that the IP address belongs
// to.
type IntelIpipIntelligenceGetIPOverviewResponseBelongsToRef struct {
	ID          interface{} `json:"id"`
	Country     string      `json:"country"`
	Description string      `json:"description"`
	// Infrastructure type of this ASN.
	Type  IntelIpipIntelligenceGetIPOverviewResponseBelongsToRefType `json:"type"`
	Value string                                                     `json:"value"`
	JSON  intelIpipIntelligenceGetIPOverviewResponseBelongsToRefJSON `json:"-"`
}

// intelIpipIntelligenceGetIPOverviewResponseBelongsToRefJSON contains the JSON
// metadata for the struct [IntelIpipIntelligenceGetIPOverviewResponseBelongsToRef]
type intelIpipIntelligenceGetIPOverviewResponseBelongsToRefJSON struct {
	ID          apijson.Field
	Country     apijson.Field
	Description apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIpipIntelligenceGetIPOverviewResponseBelongsToRef) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Infrastructure type of this ASN.
type IntelIpipIntelligenceGetIPOverviewResponseBelongsToRefType string

const (
	IntelIpipIntelligenceGetIPOverviewResponseBelongsToRefTypeHostingProvider IntelIpipIntelligenceGetIPOverviewResponseBelongsToRefType = "hosting_provider"
	IntelIpipIntelligenceGetIPOverviewResponseBelongsToRefTypeIsp             IntelIpipIntelligenceGetIPOverviewResponseBelongsToRefType = "isp"
	IntelIpipIntelligenceGetIPOverviewResponseBelongsToRefTypeOrganization    IntelIpipIntelligenceGetIPOverviewResponseBelongsToRefType = "organization"
)

// Union satisfied by [shared.UnionString] or [shared.UnionString].
type IntelIpipIntelligenceGetIPOverviewResponseIP interface {
	ImplementsIntelIpipIntelligenceGetIPOverviewResponseIP()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IntelIpipIntelligenceGetIPOverviewResponseIP)(nil)).Elem(),
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

type IntelIPIPIntelligenceGetIPOverviewParams struct {
	IPV4 param.Field[string] `query:"ipv4"`
	IPV6 param.Field[string] `query:"ipv6"`
}

// URLQuery serializes [IntelIPIPIntelligenceGetIPOverviewParams]'s query
// parameters as `url.Values`.
func (r IntelIPIPIntelligenceGetIPOverviewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IntelIpipIntelligenceGetIPOverviewResponseEnvelope struct {
	Errors   []IntelIpipIntelligenceGetIPOverviewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IntelIpipIntelligenceGetIPOverviewResponseEnvelopeMessages `json:"messages,required"`
	Result   []IntelIpipIntelligenceGetIPOverviewResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    IntelIpipIntelligenceGetIPOverviewResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo IntelIpipIntelligenceGetIPOverviewResponseEnvelopeResultInfo `json:"result_info"`
	JSON       intelIpipIntelligenceGetIPOverviewResponseEnvelopeJSON       `json:"-"`
}

// intelIpipIntelligenceGetIPOverviewResponseEnvelopeJSON contains the JSON
// metadata for the struct [IntelIpipIntelligenceGetIPOverviewResponseEnvelope]
type intelIpipIntelligenceGetIPOverviewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIpipIntelligenceGetIPOverviewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIpipIntelligenceGetIPOverviewResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    intelIpipIntelligenceGetIPOverviewResponseEnvelopeErrorsJSON `json:"-"`
}

// intelIpipIntelligenceGetIPOverviewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [IntelIpipIntelligenceGetIPOverviewResponseEnvelopeErrors]
type intelIpipIntelligenceGetIPOverviewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIpipIntelligenceGetIPOverviewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelIpipIntelligenceGetIPOverviewResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    intelIpipIntelligenceGetIPOverviewResponseEnvelopeMessagesJSON `json:"-"`
}

// intelIpipIntelligenceGetIPOverviewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [IntelIpipIntelligenceGetIPOverviewResponseEnvelopeMessages]
type intelIpipIntelligenceGetIPOverviewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIpipIntelligenceGetIPOverviewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type IntelIpipIntelligenceGetIPOverviewResponseEnvelopeSuccess bool

const (
	IntelIpipIntelligenceGetIPOverviewResponseEnvelopeSuccessTrue IntelIpipIntelligenceGetIPOverviewResponseEnvelopeSuccess = true
)

type IntelIpipIntelligenceGetIPOverviewResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                          `json:"total_count"`
	JSON       intelIpipIntelligenceGetIPOverviewResponseEnvelopeResultInfoJSON `json:"-"`
}

// intelIpipIntelligenceGetIPOverviewResponseEnvelopeResultInfoJSON contains the
// JSON metadata for the struct
// [IntelIpipIntelligenceGetIPOverviewResponseEnvelopeResultInfo]
type intelIpipIntelligenceGetIPOverviewResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelIpipIntelligenceGetIPOverviewResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
