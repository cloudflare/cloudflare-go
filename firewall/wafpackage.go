// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// WAFPackageService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewWAFPackageService] method instead.
type WAFPackageService struct {
	Options []option.RequestOption
	Groups  *WAFPackageGroupService
	Rules   *WAFPackageRuleService
}

// NewWAFPackageService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWAFPackageService(opts ...option.RequestOption) (r *WAFPackageService) {
	r = &WAFPackageService{}
	r.Options = opts
	r.Groups = NewWAFPackageGroupService(opts...)
	r.Rules = NewWAFPackageRuleService(opts...)
	return
}

// Fetches WAF packages for a zone.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *WAFPackageService) List(ctx context.Context, zoneIdentifier string, query WAFPackageListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[WAFPackageListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/firewall/waf/packages", zoneIdentifier)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Fetches WAF packages for a zone.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *WAFPackageService) ListAutoPaging(ctx context.Context, zoneIdentifier string, query WAFPackageListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[WAFPackageListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, zoneIdentifier, query, opts...))
}

// Fetches the details of a WAF package.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *WAFPackageService) Get(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *WAFPackageGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Union satisfied by
// [firewall.WAFPackageListResponseLegacyJhsAPIResponseCollection] or
// [firewall.WAFPackageListResponseObject].
type WAFPackageListResponse interface {
	implementsFirewallWAFPackageListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WAFPackageListResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WAFPackageListResponseLegacyJhsAPIResponseCollection{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WAFPackageListResponseObject{}),
		},
	)
}

type WAFPackageListResponseLegacyJhsAPIResponseCollection struct {
	Errors   []shared.ResponseInfo                                      `json:"errors,required"`
	Messages []shared.ResponseInfo                                      `json:"messages,required"`
	Result   WAFPackageListResponseLegacyJhsAPIResponseCollectionResult `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    WAFPackageListResponseLegacyJhsAPIResponseCollectionSuccess    `json:"success,required"`
	ResultInfo WAFPackageListResponseLegacyJhsAPIResponseCollectionResultInfo `json:"result_info"`
	JSON       wafPackageListResponseLegacyJhsAPIResponseCollectionJSON       `json:"-"`
}

// wafPackageListResponseLegacyJhsAPIResponseCollectionJSON contains the JSON
// metadata for the struct [WAFPackageListResponseLegacyJhsAPIResponseCollection]
type wafPackageListResponseLegacyJhsAPIResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFPackageListResponseLegacyJhsAPIResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageListResponseLegacyJhsAPIResponseCollectionJSON) RawJSON() string {
	return r.raw
}

func (r WAFPackageListResponseLegacyJhsAPIResponseCollection) implementsFirewallWAFPackageListResponse() {
}

// Union satisfied by
// [firewall.WAFPackageListResponseLegacyJhsAPIResponseCollectionResultUnknown],
// [firewall.WAFPackageListResponseLegacyJhsAPIResponseCollectionResultArray] or
// [shared.UnionString].
type WAFPackageListResponseLegacyJhsAPIResponseCollectionResult interface {
	ImplementsFirewallWAFPackageListResponseLegacyJhsAPIResponseCollectionResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WAFPackageListResponseLegacyJhsAPIResponseCollectionResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WAFPackageListResponseLegacyJhsAPIResponseCollectionResultArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type WAFPackageListResponseLegacyJhsAPIResponseCollectionResultArray []interface{}

func (r WAFPackageListResponseLegacyJhsAPIResponseCollectionResultArray) ImplementsFirewallWAFPackageListResponseLegacyJhsAPIResponseCollectionResult() {
}

// Whether the API call was successful
type WAFPackageListResponseLegacyJhsAPIResponseCollectionSuccess bool

const (
	WAFPackageListResponseLegacyJhsAPIResponseCollectionSuccessTrue WAFPackageListResponseLegacyJhsAPIResponseCollectionSuccess = true
)

func (r WAFPackageListResponseLegacyJhsAPIResponseCollectionSuccess) IsKnown() bool {
	switch r {
	case WAFPackageListResponseLegacyJhsAPIResponseCollectionSuccessTrue:
		return true
	}
	return false
}

type WAFPackageListResponseLegacyJhsAPIResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                            `json:"total_count"`
	JSON       wafPackageListResponseLegacyJhsAPIResponseCollectionResultInfoJSON `json:"-"`
}

// wafPackageListResponseLegacyJhsAPIResponseCollectionResultInfoJSON contains the
// JSON metadata for the struct
// [WAFPackageListResponseLegacyJhsAPIResponseCollectionResultInfo]
type wafPackageListResponseLegacyJhsAPIResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFPackageListResponseLegacyJhsAPIResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageListResponseLegacyJhsAPIResponseCollectionResultInfoJSON) RawJSON() string {
	return r.raw
}

type WAFPackageListResponseObject struct {
	Result []WAFPackageListResponseObjectResult `json:"result"`
	JSON   wafPackageListResponseObjectJSON     `json:"-"`
}

// wafPackageListResponseObjectJSON contains the JSON metadata for the struct
// [WAFPackageListResponseObject]
type wafPackageListResponseObjectJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFPackageListResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageListResponseObjectJSON) RawJSON() string {
	return r.raw
}

func (r WAFPackageListResponseObject) implementsFirewallWAFPackageListResponse() {}

// Union satisfied by
// [firewall.WAFPackageListResponseObjectResultLegacyJhsPackageDefinition] or
// [firewall.WAFPackageListResponseObjectResultLegacyJhsAnomalyPackage].
type WAFPackageListResponseObjectResult interface {
	implementsFirewallWAFPackageListResponseObjectResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WAFPackageListResponseObjectResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WAFPackageListResponseObjectResultLegacyJhsPackageDefinition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WAFPackageListResponseObjectResultLegacyJhsAnomalyPackage{}),
		},
	)
}

type WAFPackageListResponseObjectResultLegacyJhsPackageDefinition struct {
	// The unique identifier of a WAF package.
	ID string `json:"id,required"`
	// A summary of the purpose/function of the WAF package.
	Description string `json:"description,required"`
	// The mode that defines how rules within the package are evaluated during the
	// course of a request. When a package uses anomaly detection mode (`anomaly`
	// value), each rule is given a score when triggered. If the total score of all
	// triggered rules exceeds the sensitivity defined in the WAF package, the action
	// configured in the package will be performed. Traditional detection mode
	// (`traditional` value) will decide the action to take when it is triggered by the
	// request. If multiple rules are triggered, the action providing the highest
	// protection will be applied (for example, a 'block' action will win over a
	// 'challenge' action).
	DetectionMode WAFPackageListResponseObjectResultLegacyJhsPackageDefinitionDetectionMode `json:"detection_mode,required"`
	// The name of the WAF package.
	Name string `json:"name,required"`
	// Identifier
	ZoneID string `json:"zone_id,required"`
	// When set to `active`, indicates that the WAF package will be applied to the
	// zone.
	Status WAFPackageListResponseObjectResultLegacyJhsPackageDefinitionStatus `json:"status"`
	JSON   wafPackageListResponseObjectResultLegacyJhsPackageDefinitionJSON   `json:"-"`
}

// wafPackageListResponseObjectResultLegacyJhsPackageDefinitionJSON contains the
// JSON metadata for the struct
// [WAFPackageListResponseObjectResultLegacyJhsPackageDefinition]
type wafPackageListResponseObjectResultLegacyJhsPackageDefinitionJSON struct {
	ID            apijson.Field
	Description   apijson.Field
	DetectionMode apijson.Field
	Name          apijson.Field
	ZoneID        apijson.Field
	Status        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WAFPackageListResponseObjectResultLegacyJhsPackageDefinition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageListResponseObjectResultLegacyJhsPackageDefinitionJSON) RawJSON() string {
	return r.raw
}

func (r WAFPackageListResponseObjectResultLegacyJhsPackageDefinition) implementsFirewallWAFPackageListResponseObjectResult() {
}

// The mode that defines how rules within the package are evaluated during the
// course of a request. When a package uses anomaly detection mode (`anomaly`
// value), each rule is given a score when triggered. If the total score of all
// triggered rules exceeds the sensitivity defined in the WAF package, the action
// configured in the package will be performed. Traditional detection mode
// (`traditional` value) will decide the action to take when it is triggered by the
// request. If multiple rules are triggered, the action providing the highest
// protection will be applied (for example, a 'block' action will win over a
// 'challenge' action).
type WAFPackageListResponseObjectResultLegacyJhsPackageDefinitionDetectionMode string

const (
	WAFPackageListResponseObjectResultLegacyJhsPackageDefinitionDetectionModeAnomaly     WAFPackageListResponseObjectResultLegacyJhsPackageDefinitionDetectionMode = "anomaly"
	WAFPackageListResponseObjectResultLegacyJhsPackageDefinitionDetectionModeTraditional WAFPackageListResponseObjectResultLegacyJhsPackageDefinitionDetectionMode = "traditional"
)

func (r WAFPackageListResponseObjectResultLegacyJhsPackageDefinitionDetectionMode) IsKnown() bool {
	switch r {
	case WAFPackageListResponseObjectResultLegacyJhsPackageDefinitionDetectionModeAnomaly, WAFPackageListResponseObjectResultLegacyJhsPackageDefinitionDetectionModeTraditional:
		return true
	}
	return false
}

// When set to `active`, indicates that the WAF package will be applied to the
// zone.
type WAFPackageListResponseObjectResultLegacyJhsPackageDefinitionStatus string

const (
	WAFPackageListResponseObjectResultLegacyJhsPackageDefinitionStatusActive WAFPackageListResponseObjectResultLegacyJhsPackageDefinitionStatus = "active"
)

func (r WAFPackageListResponseObjectResultLegacyJhsPackageDefinitionStatus) IsKnown() bool {
	switch r {
	case WAFPackageListResponseObjectResultLegacyJhsPackageDefinitionStatusActive:
		return true
	}
	return false
}

type WAFPackageListResponseObjectResultLegacyJhsAnomalyPackage struct {
	// The unique identifier of a WAF package.
	ID string `json:"id,required"`
	// A summary of the purpose/function of the WAF package.
	Description string `json:"description,required"`
	// When a WAF package uses anomaly detection, each rule is given a score when
	// triggered. If the total score of all triggered rules exceeds the sensitivity
	// defined on the WAF package, the action defined on the package will be taken.
	DetectionMode WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageDetectionMode `json:"detection_mode,required"`
	// The name of the WAF package.
	Name string `json:"name,required"`
	// Identifier
	ZoneID string `json:"zone_id,required"`
	// The default action performed by the rules in the WAF package.
	ActionMode WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageActionMode `json:"action_mode"`
	// The sensitivity of the WAF package.
	Sensitivity WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageSensitivity `json:"sensitivity"`
	// When set to `active`, indicates that the WAF package will be applied to the
	// zone.
	Status WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageStatus `json:"status"`
	JSON   wafPackageListResponseObjectResultLegacyJhsAnomalyPackageJSON   `json:"-"`
}

// wafPackageListResponseObjectResultLegacyJhsAnomalyPackageJSON contains the JSON
// metadata for the struct
// [WAFPackageListResponseObjectResultLegacyJhsAnomalyPackage]
type wafPackageListResponseObjectResultLegacyJhsAnomalyPackageJSON struct {
	ID            apijson.Field
	Description   apijson.Field
	DetectionMode apijson.Field
	Name          apijson.Field
	ZoneID        apijson.Field
	ActionMode    apijson.Field
	Sensitivity   apijson.Field
	Status        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WAFPackageListResponseObjectResultLegacyJhsAnomalyPackage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageListResponseObjectResultLegacyJhsAnomalyPackageJSON) RawJSON() string {
	return r.raw
}

func (r WAFPackageListResponseObjectResultLegacyJhsAnomalyPackage) implementsFirewallWAFPackageListResponseObjectResult() {
}

// When a WAF package uses anomaly detection, each rule is given a score when
// triggered. If the total score of all triggered rules exceeds the sensitivity
// defined on the WAF package, the action defined on the package will be taken.
type WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageDetectionMode string

const (
	WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageDetectionModeAnomaly     WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageDetectionMode = "anomaly"
	WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageDetectionModeTraditional WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageDetectionMode = "traditional"
)

func (r WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageDetectionMode) IsKnown() bool {
	switch r {
	case WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageDetectionModeAnomaly, WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageDetectionModeTraditional:
		return true
	}
	return false
}

// The default action performed by the rules in the WAF package.
type WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageActionMode string

const (
	WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageActionModeSimulate  WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageActionMode = "simulate"
	WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageActionModeBlock     WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageActionMode = "block"
	WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageActionModeChallenge WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageActionMode = "challenge"
)

func (r WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageActionMode) IsKnown() bool {
	switch r {
	case WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageActionModeSimulate, WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageActionModeBlock, WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageActionModeChallenge:
		return true
	}
	return false
}

// The sensitivity of the WAF package.
type WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageSensitivity string

const (
	WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageSensitivityHigh   WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageSensitivity = "high"
	WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageSensitivityMedium WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageSensitivity = "medium"
	WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageSensitivityLow    WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageSensitivity = "low"
	WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageSensitivityOff    WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageSensitivity = "off"
)

func (r WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageSensitivity) IsKnown() bool {
	switch r {
	case WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageSensitivityHigh, WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageSensitivityMedium, WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageSensitivityLow, WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageSensitivityOff:
		return true
	}
	return false
}

// When set to `active`, indicates that the WAF package will be applied to the
// zone.
type WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageStatus string

const (
	WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageStatusActive WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageStatus = "active"
)

func (r WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageStatus) IsKnown() bool {
	switch r {
	case WAFPackageListResponseObjectResultLegacyJhsAnomalyPackageStatusActive:
		return true
	}
	return false
}

// Union satisfied by [firewall.WAFPackageGetResponseLegacyJhsAPIResponseSingle] or
// [firewall.WAFPackageGetResponseObject].
type WAFPackageGetResponse interface {
	implementsFirewallWAFPackageGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WAFPackageGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WAFPackageGetResponseLegacyJhsAPIResponseSingle{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WAFPackageGetResponseObject{}),
		},
	)
}

type WAFPackageGetResponseLegacyJhsAPIResponseSingle struct {
	Errors   []shared.ResponseInfo                                 `json:"errors,required"`
	Messages []shared.ResponseInfo                                 `json:"messages,required"`
	Result   WAFPackageGetResponseLegacyJhsAPIResponseSingleResult `json:"result,required,nullable"`
	// Whether the API call was successful
	Success WAFPackageGetResponseLegacyJhsAPIResponseSingleSuccess `json:"success,required"`
	JSON    wafPackageGetResponseLegacyJhsAPIResponseSingleJSON    `json:"-"`
}

// wafPackageGetResponseLegacyJhsAPIResponseSingleJSON contains the JSON metadata
// for the struct [WAFPackageGetResponseLegacyJhsAPIResponseSingle]
type wafPackageGetResponseLegacyJhsAPIResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFPackageGetResponseLegacyJhsAPIResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageGetResponseLegacyJhsAPIResponseSingleJSON) RawJSON() string {
	return r.raw
}

func (r WAFPackageGetResponseLegacyJhsAPIResponseSingle) implementsFirewallWAFPackageGetResponse() {}

// Union satisfied by
// [firewall.WAFPackageGetResponseLegacyJhsAPIResponseSingleResultUnknown] or
// [shared.UnionString].
type WAFPackageGetResponseLegacyJhsAPIResponseSingleResult interface {
	ImplementsFirewallWAFPackageGetResponseLegacyJhsAPIResponseSingleResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WAFPackageGetResponseLegacyJhsAPIResponseSingleResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Whether the API call was successful
type WAFPackageGetResponseLegacyJhsAPIResponseSingleSuccess bool

const (
	WAFPackageGetResponseLegacyJhsAPIResponseSingleSuccessTrue WAFPackageGetResponseLegacyJhsAPIResponseSingleSuccess = true
)

func (r WAFPackageGetResponseLegacyJhsAPIResponseSingleSuccess) IsKnown() bool {
	switch r {
	case WAFPackageGetResponseLegacyJhsAPIResponseSingleSuccessTrue:
		return true
	}
	return false
}

type WAFPackageGetResponseObject struct {
	Result interface{}                     `json:"result"`
	JSON   wafPackageGetResponseObjectJSON `json:"-"`
}

// wafPackageGetResponseObjectJSON contains the JSON metadata for the struct
// [WAFPackageGetResponseObject]
type wafPackageGetResponseObjectJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFPackageGetResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageGetResponseObjectJSON) RawJSON() string {
	return r.raw
}

func (r WAFPackageGetResponseObject) implementsFirewallWAFPackageGetResponse() {}

type WAFPackageListParams struct {
	// The direction used to sort returned packages.
	Direction param.Field[WAFPackageListParamsDirection] `query:"direction"`
	// When set to `all`, all the search requirements must match. When set to `any`,
	// only one of the search requirements has to match.
	Match param.Field[WAFPackageListParamsMatch] `query:"match"`
	// The name of the WAF package.
	Name param.Field[string] `query:"name"`
	// The field used to sort returned packages.
	Order param.Field[WAFPackageListParamsOrder] `query:"order"`
	// The page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The number of packages per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [WAFPackageListParams]'s query parameters as `url.Values`.
func (r WAFPackageListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The direction used to sort returned packages.
type WAFPackageListParamsDirection string

const (
	WAFPackageListParamsDirectionAsc  WAFPackageListParamsDirection = "asc"
	WAFPackageListParamsDirectionDesc WAFPackageListParamsDirection = "desc"
)

func (r WAFPackageListParamsDirection) IsKnown() bool {
	switch r {
	case WAFPackageListParamsDirectionAsc, WAFPackageListParamsDirectionDesc:
		return true
	}
	return false
}

// When set to `all`, all the search requirements must match. When set to `any`,
// only one of the search requirements has to match.
type WAFPackageListParamsMatch string

const (
	WAFPackageListParamsMatchAny WAFPackageListParamsMatch = "any"
	WAFPackageListParamsMatchAll WAFPackageListParamsMatch = "all"
)

func (r WAFPackageListParamsMatch) IsKnown() bool {
	switch r {
	case WAFPackageListParamsMatchAny, WAFPackageListParamsMatchAll:
		return true
	}
	return false
}

// The field used to sort returned packages.
type WAFPackageListParamsOrder string

const (
	WAFPackageListParamsOrderName WAFPackageListParamsOrder = "name"
)

func (r WAFPackageListParamsOrder) IsKnown() bool {
	switch r {
	case WAFPackageListParamsOrderName:
		return true
	}
	return false
}
