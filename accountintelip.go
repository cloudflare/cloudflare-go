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

// AccountIntelIPService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountIntelIPService] method
// instead.
type AccountIntelIPService struct {
	Options []option.RequestOption
}

// NewAccountIntelIPService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountIntelIPService(opts ...option.RequestOption) (r *AccountIntelIPService) {
	r = &AccountIntelIPService{}
	r.Options = opts
	return
}

// Get IP Overview
func (r *AccountIntelIPService) IPIntelligenceGetIPOverview(ctx context.Context, accountIdentifier string, query AccountIntelIPIPIntelligenceGetIPOverviewParams, opts ...option.RequestOption) (res *AccountIntelIpipIntelligenceGetIPOverviewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/intel/ip", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountIntelIpipIntelligenceGetIPOverviewResponse struct {
	Errors     []AccountIntelIpipIntelligenceGetIPOverviewResponseError    `json:"errors"`
	Messages   []AccountIntelIpipIntelligenceGetIPOverviewResponseMessage  `json:"messages"`
	Result     []AccountIntelIpipIntelligenceGetIPOverviewResponseResult   `json:"result"`
	ResultInfo AccountIntelIpipIntelligenceGetIPOverviewResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountIntelIpipIntelligenceGetIPOverviewResponseSuccess `json:"success"`
	JSON    accountIntelIpipIntelligenceGetIPOverviewResponseJSON    `json:"-"`
}

// accountIntelIpipIntelligenceGetIPOverviewResponseJSON contains the JSON metadata
// for the struct [AccountIntelIpipIntelligenceGetIPOverviewResponse]
type accountIntelIpipIntelligenceGetIPOverviewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelIpipIntelligenceGetIPOverviewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelIpipIntelligenceGetIPOverviewResponseError struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    accountIntelIpipIntelligenceGetIPOverviewResponseErrorJSON `json:"-"`
}

// accountIntelIpipIntelligenceGetIPOverviewResponseErrorJSON contains the JSON
// metadata for the struct [AccountIntelIpipIntelligenceGetIPOverviewResponseError]
type accountIntelIpipIntelligenceGetIPOverviewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelIpipIntelligenceGetIPOverviewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelIpipIntelligenceGetIPOverviewResponseMessage struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    accountIntelIpipIntelligenceGetIPOverviewResponseMessageJSON `json:"-"`
}

// accountIntelIpipIntelligenceGetIPOverviewResponseMessageJSON contains the JSON
// metadata for the struct
// [AccountIntelIpipIntelligenceGetIPOverviewResponseMessage]
type accountIntelIpipIntelligenceGetIPOverviewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelIpipIntelligenceGetIPOverviewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelIpipIntelligenceGetIPOverviewResponseResult struct {
	// Specifies a reference to the autonomous systems (AS) that the IP address belongs
	// to.
	BelongsToRef AccountIntelIpipIntelligenceGetIPOverviewResponseResultBelongsToRef `json:"belongs_to_ref"`
	IP           AccountIntelIpipIntelligenceGetIPOverviewResponseResultIP           `json:"ip" format:"ipv4"`
	RiskTypes    interface{}                                                         `json:"risk_types"`
	JSON         accountIntelIpipIntelligenceGetIPOverviewResponseResultJSON         `json:"-"`
}

// accountIntelIpipIntelligenceGetIPOverviewResponseResultJSON contains the JSON
// metadata for the struct
// [AccountIntelIpipIntelligenceGetIPOverviewResponseResult]
type accountIntelIpipIntelligenceGetIPOverviewResponseResultJSON struct {
	BelongsToRef apijson.Field
	IP           apijson.Field
	RiskTypes    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountIntelIpipIntelligenceGetIPOverviewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies a reference to the autonomous systems (AS) that the IP address belongs
// to.
type AccountIntelIpipIntelligenceGetIPOverviewResponseResultBelongsToRef struct {
	ID          interface{} `json:"id"`
	Country     string      `json:"country"`
	Description string      `json:"description"`
	// Infrastructure type of this ASN.
	Type  AccountIntelIpipIntelligenceGetIPOverviewResponseResultBelongsToRefType `json:"type"`
	Value string                                                                  `json:"value"`
	JSON  accountIntelIpipIntelligenceGetIPOverviewResponseResultBelongsToRefJSON `json:"-"`
}

// accountIntelIpipIntelligenceGetIPOverviewResponseResultBelongsToRefJSON contains
// the JSON metadata for the struct
// [AccountIntelIpipIntelligenceGetIPOverviewResponseResultBelongsToRef]
type accountIntelIpipIntelligenceGetIPOverviewResponseResultBelongsToRefJSON struct {
	ID          apijson.Field
	Country     apijson.Field
	Description apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelIpipIntelligenceGetIPOverviewResponseResultBelongsToRef) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Infrastructure type of this ASN.
type AccountIntelIpipIntelligenceGetIPOverviewResponseResultBelongsToRefType string

const (
	AccountIntelIpipIntelligenceGetIPOverviewResponseResultBelongsToRefTypeHostingProvider AccountIntelIpipIntelligenceGetIPOverviewResponseResultBelongsToRefType = "hosting_provider"
	AccountIntelIpipIntelligenceGetIPOverviewResponseResultBelongsToRefTypeIsp             AccountIntelIpipIntelligenceGetIPOverviewResponseResultBelongsToRefType = "isp"
	AccountIntelIpipIntelligenceGetIPOverviewResponseResultBelongsToRefTypeOrganization    AccountIntelIpipIntelligenceGetIPOverviewResponseResultBelongsToRefType = "organization"
)

// Union satisfied by [shared.UnionString] or [shared.UnionString].
type AccountIntelIpipIntelligenceGetIPOverviewResponseResultIP interface {
	ImplementsAccountIntelIpipIntelligenceGetIPOverviewResponseResultIP()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountIntelIpipIntelligenceGetIPOverviewResponseResultIP)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountIntelIpipIntelligenceGetIPOverviewResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                         `json:"total_count"`
	JSON       accountIntelIpipIntelligenceGetIPOverviewResponseResultInfoJSON `json:"-"`
}

// accountIntelIpipIntelligenceGetIPOverviewResponseResultInfoJSON contains the
// JSON metadata for the struct
// [AccountIntelIpipIntelligenceGetIPOverviewResponseResultInfo]
type accountIntelIpipIntelligenceGetIPOverviewResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelIpipIntelligenceGetIPOverviewResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountIntelIpipIntelligenceGetIPOverviewResponseSuccess bool

const (
	AccountIntelIpipIntelligenceGetIPOverviewResponseSuccessTrue AccountIntelIpipIntelligenceGetIPOverviewResponseSuccess = true
)

type AccountIntelIPIPIntelligenceGetIPOverviewParams struct {
	Ipv4 param.Field[string] `query:"ipv4"`
	Ipv6 param.Field[string] `query:"ipv6"`
}

// URLQuery serializes [AccountIntelIPIPIntelligenceGetIPOverviewParams]'s query
// parameters as `url.Values`.
func (r AccountIntelIPIPIntelligenceGetIPOverviewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
