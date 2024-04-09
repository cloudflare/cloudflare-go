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

type WAFPackageListResponse struct {
	Errors   interface{} `json:"errors,required"`
	Messages interface{} `json:"messages,required"`
	Result   interface{} `json:"result,required"`
	// Whether the API call was successful
	Success    WAFPackageListResponseSuccess `json:"success"`
	ResultInfo interface{}                   `json:"result_info,required"`
	JSON       wafPackageListResponseJSON    `json:"-"`
	union      WAFPackageListResponseUnion
}

// wafPackageListResponseJSON contains the JSON metadata for the struct
// [WAFPackageListResponse]
type wafPackageListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r wafPackageListResponseJSON) RawJSON() string {
	return r.raw
}

func (r *WAFPackageListResponse) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r WAFPackageListResponse) AsUnion() WAFPackageListResponseUnion {
	return r.union
}

// Union satisfied by
// [firewall.WAFPackageListResponseFirewallAPIResponseCollection] or
// [firewall.WAFPackageListResponseObject].
type WAFPackageListResponseUnion interface {
	implementsFirewallWAFPackageListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WAFPackageListResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WAFPackageListResponseFirewallAPIResponseCollection{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WAFPackageListResponseObject{}),
		},
	)
}

type WAFPackageListResponseFirewallAPIResponseCollection struct {
	Errors   []shared.ResponseInfo                                        `json:"errors,required"`
	Messages []shared.ResponseInfo                                        `json:"messages,required"`
	Result   shared.UnnamedSchemaRef67bbb1ccdd42c3e2937b9fd19f791151Union `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    WAFPackageListResponseFirewallAPIResponseCollectionSuccess    `json:"success,required"`
	ResultInfo WAFPackageListResponseFirewallAPIResponseCollectionResultInfo `json:"result_info"`
	JSON       wafPackageListResponseFirewallAPIResponseCollectionJSON       `json:"-"`
}

// wafPackageListResponseFirewallAPIResponseCollectionJSON contains the JSON
// metadata for the struct [WAFPackageListResponseFirewallAPIResponseCollection]
type wafPackageListResponseFirewallAPIResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFPackageListResponseFirewallAPIResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageListResponseFirewallAPIResponseCollectionJSON) RawJSON() string {
	return r.raw
}

func (r WAFPackageListResponseFirewallAPIResponseCollection) implementsFirewallWAFPackageListResponse() {
}

// Whether the API call was successful
type WAFPackageListResponseFirewallAPIResponseCollectionSuccess bool

const (
	WAFPackageListResponseFirewallAPIResponseCollectionSuccessTrue WAFPackageListResponseFirewallAPIResponseCollectionSuccess = true
)

func (r WAFPackageListResponseFirewallAPIResponseCollectionSuccess) IsKnown() bool {
	switch r {
	case WAFPackageListResponseFirewallAPIResponseCollectionSuccessTrue:
		return true
	}
	return false
}

type WAFPackageListResponseFirewallAPIResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                           `json:"total_count"`
	JSON       wafPackageListResponseFirewallAPIResponseCollectionResultInfoJSON `json:"-"`
}

// wafPackageListResponseFirewallAPIResponseCollectionResultInfoJSON contains the
// JSON metadata for the struct
// [WAFPackageListResponseFirewallAPIResponseCollectionResultInfo]
type wafPackageListResponseFirewallAPIResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFPackageListResponseFirewallAPIResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageListResponseFirewallAPIResponseCollectionResultInfoJSON) RawJSON() string {
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

type WAFPackageListResponseObjectResult struct {
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
	DetectionMode WAFPackageListResponseObjectResultDetectionMode `json:"detection_mode,required"`
	// Identifier
	ID string `json:"id,required"`
	// The name of the WAF package.
	Name string `json:"name,required"`
	// When set to `active`, indicates that the WAF package will be applied to the
	// zone.
	Status WAFPackageListResponseObjectResultStatus `json:"status"`
	// Identifier
	ZoneID string `json:"zone_id,required"`
	// The default action performed by the rules in the WAF package.
	ActionMode WAFPackageListResponseObjectResultActionMode `json:"action_mode"`
	// The sensitivity of the WAF package.
	Sensitivity WAFPackageListResponseObjectResultSensitivity `json:"sensitivity"`
	JSON        wafPackageListResponseObjectResultJSON        `json:"-"`
	union       WAFPackageListResponseObjectResultUnion
}

// wafPackageListResponseObjectResultJSON contains the JSON metadata for the struct
// [WAFPackageListResponseObjectResult]
type wafPackageListResponseObjectResultJSON struct {
	Description   apijson.Field
	DetectionMode apijson.Field
	ID            apijson.Field
	Name          apijson.Field
	Status        apijson.Field
	ZoneID        apijson.Field
	ActionMode    apijson.Field
	Sensitivity   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r wafPackageListResponseObjectResultJSON) RawJSON() string {
	return r.raw
}

func (r *WAFPackageListResponseObjectResult) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r WAFPackageListResponseObjectResult) AsUnion() WAFPackageListResponseObjectResultUnion {
	return r.union
}

// Union satisfied by
// [firewall.WAFPackageListResponseObjectResultFirewallPackageDefinition] or
// [firewall.WAFPackageListResponseObjectResultFirewallAnomalyPackage].
type WAFPackageListResponseObjectResultUnion interface {
	implementsFirewallWAFPackageListResponseObjectResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WAFPackageListResponseObjectResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WAFPackageListResponseObjectResultFirewallPackageDefinition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WAFPackageListResponseObjectResultFirewallAnomalyPackage{}),
		},
	)
}

type WAFPackageListResponseObjectResultFirewallPackageDefinition struct {
	// Identifier
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
	DetectionMode WAFPackageListResponseObjectResultFirewallPackageDefinitionDetectionMode `json:"detection_mode,required"`
	// The name of the WAF package.
	Name string `json:"name,required"`
	// Identifier
	ZoneID string `json:"zone_id,required"`
	// When set to `active`, indicates that the WAF package will be applied to the
	// zone.
	Status WAFPackageListResponseObjectResultFirewallPackageDefinitionStatus `json:"status"`
	JSON   wafPackageListResponseObjectResultFirewallPackageDefinitionJSON   `json:"-"`
}

// wafPackageListResponseObjectResultFirewallPackageDefinitionJSON contains the
// JSON metadata for the struct
// [WAFPackageListResponseObjectResultFirewallPackageDefinition]
type wafPackageListResponseObjectResultFirewallPackageDefinitionJSON struct {
	ID            apijson.Field
	Description   apijson.Field
	DetectionMode apijson.Field
	Name          apijson.Field
	ZoneID        apijson.Field
	Status        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WAFPackageListResponseObjectResultFirewallPackageDefinition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageListResponseObjectResultFirewallPackageDefinitionJSON) RawJSON() string {
	return r.raw
}

func (r WAFPackageListResponseObjectResultFirewallPackageDefinition) implementsFirewallWAFPackageListResponseObjectResult() {
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
type WAFPackageListResponseObjectResultFirewallPackageDefinitionDetectionMode string

const (
	WAFPackageListResponseObjectResultFirewallPackageDefinitionDetectionModeAnomaly     WAFPackageListResponseObjectResultFirewallPackageDefinitionDetectionMode = "anomaly"
	WAFPackageListResponseObjectResultFirewallPackageDefinitionDetectionModeTraditional WAFPackageListResponseObjectResultFirewallPackageDefinitionDetectionMode = "traditional"
)

func (r WAFPackageListResponseObjectResultFirewallPackageDefinitionDetectionMode) IsKnown() bool {
	switch r {
	case WAFPackageListResponseObjectResultFirewallPackageDefinitionDetectionModeAnomaly, WAFPackageListResponseObjectResultFirewallPackageDefinitionDetectionModeTraditional:
		return true
	}
	return false
}

// When set to `active`, indicates that the WAF package will be applied to the
// zone.
type WAFPackageListResponseObjectResultFirewallPackageDefinitionStatus string

const (
	WAFPackageListResponseObjectResultFirewallPackageDefinitionStatusActive WAFPackageListResponseObjectResultFirewallPackageDefinitionStatus = "active"
)

func (r WAFPackageListResponseObjectResultFirewallPackageDefinitionStatus) IsKnown() bool {
	switch r {
	case WAFPackageListResponseObjectResultFirewallPackageDefinitionStatusActive:
		return true
	}
	return false
}

type WAFPackageListResponseObjectResultFirewallAnomalyPackage struct {
	// Identifier
	ID string `json:"id,required"`
	// A summary of the purpose/function of the WAF package.
	Description string `json:"description,required"`
	// When a WAF package uses anomaly detection, each rule is given a score when
	// triggered. If the total score of all triggered rules exceeds the sensitivity
	// defined on the WAF package, the action defined on the package will be taken.
	DetectionMode WAFPackageListResponseObjectResultFirewallAnomalyPackageDetectionMode `json:"detection_mode,required"`
	// The name of the WAF package.
	Name string `json:"name,required"`
	// Identifier
	ZoneID string `json:"zone_id,required"`
	// The default action performed by the rules in the WAF package.
	ActionMode WAFPackageListResponseObjectResultFirewallAnomalyPackageActionMode `json:"action_mode"`
	// The sensitivity of the WAF package.
	Sensitivity WAFPackageListResponseObjectResultFirewallAnomalyPackageSensitivity `json:"sensitivity"`
	// When set to `active`, indicates that the WAF package will be applied to the
	// zone.
	Status WAFPackageListResponseObjectResultFirewallAnomalyPackageStatus `json:"status"`
	JSON   wafPackageListResponseObjectResultFirewallAnomalyPackageJSON   `json:"-"`
}

// wafPackageListResponseObjectResultFirewallAnomalyPackageJSON contains the JSON
// metadata for the struct
// [WAFPackageListResponseObjectResultFirewallAnomalyPackage]
type wafPackageListResponseObjectResultFirewallAnomalyPackageJSON struct {
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

func (r *WAFPackageListResponseObjectResultFirewallAnomalyPackage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageListResponseObjectResultFirewallAnomalyPackageJSON) RawJSON() string {
	return r.raw
}

func (r WAFPackageListResponseObjectResultFirewallAnomalyPackage) implementsFirewallWAFPackageListResponseObjectResult() {
}

// When a WAF package uses anomaly detection, each rule is given a score when
// triggered. If the total score of all triggered rules exceeds the sensitivity
// defined on the WAF package, the action defined on the package will be taken.
type WAFPackageListResponseObjectResultFirewallAnomalyPackageDetectionMode string

const (
	WAFPackageListResponseObjectResultFirewallAnomalyPackageDetectionModeAnomaly     WAFPackageListResponseObjectResultFirewallAnomalyPackageDetectionMode = "anomaly"
	WAFPackageListResponseObjectResultFirewallAnomalyPackageDetectionModeTraditional WAFPackageListResponseObjectResultFirewallAnomalyPackageDetectionMode = "traditional"
)

func (r WAFPackageListResponseObjectResultFirewallAnomalyPackageDetectionMode) IsKnown() bool {
	switch r {
	case WAFPackageListResponseObjectResultFirewallAnomalyPackageDetectionModeAnomaly, WAFPackageListResponseObjectResultFirewallAnomalyPackageDetectionModeTraditional:
		return true
	}
	return false
}

// The default action performed by the rules in the WAF package.
type WAFPackageListResponseObjectResultFirewallAnomalyPackageActionMode string

const (
	WAFPackageListResponseObjectResultFirewallAnomalyPackageActionModeSimulate  WAFPackageListResponseObjectResultFirewallAnomalyPackageActionMode = "simulate"
	WAFPackageListResponseObjectResultFirewallAnomalyPackageActionModeBlock     WAFPackageListResponseObjectResultFirewallAnomalyPackageActionMode = "block"
	WAFPackageListResponseObjectResultFirewallAnomalyPackageActionModeChallenge WAFPackageListResponseObjectResultFirewallAnomalyPackageActionMode = "challenge"
)

func (r WAFPackageListResponseObjectResultFirewallAnomalyPackageActionMode) IsKnown() bool {
	switch r {
	case WAFPackageListResponseObjectResultFirewallAnomalyPackageActionModeSimulate, WAFPackageListResponseObjectResultFirewallAnomalyPackageActionModeBlock, WAFPackageListResponseObjectResultFirewallAnomalyPackageActionModeChallenge:
		return true
	}
	return false
}

// The sensitivity of the WAF package.
type WAFPackageListResponseObjectResultFirewallAnomalyPackageSensitivity string

const (
	WAFPackageListResponseObjectResultFirewallAnomalyPackageSensitivityHigh   WAFPackageListResponseObjectResultFirewallAnomalyPackageSensitivity = "high"
	WAFPackageListResponseObjectResultFirewallAnomalyPackageSensitivityMedium WAFPackageListResponseObjectResultFirewallAnomalyPackageSensitivity = "medium"
	WAFPackageListResponseObjectResultFirewallAnomalyPackageSensitivityLow    WAFPackageListResponseObjectResultFirewallAnomalyPackageSensitivity = "low"
	WAFPackageListResponseObjectResultFirewallAnomalyPackageSensitivityOff    WAFPackageListResponseObjectResultFirewallAnomalyPackageSensitivity = "off"
)

func (r WAFPackageListResponseObjectResultFirewallAnomalyPackageSensitivity) IsKnown() bool {
	switch r {
	case WAFPackageListResponseObjectResultFirewallAnomalyPackageSensitivityHigh, WAFPackageListResponseObjectResultFirewallAnomalyPackageSensitivityMedium, WAFPackageListResponseObjectResultFirewallAnomalyPackageSensitivityLow, WAFPackageListResponseObjectResultFirewallAnomalyPackageSensitivityOff:
		return true
	}
	return false
}

// When set to `active`, indicates that the WAF package will be applied to the
// zone.
type WAFPackageListResponseObjectResultFirewallAnomalyPackageStatus string

const (
	WAFPackageListResponseObjectResultFirewallAnomalyPackageStatusActive WAFPackageListResponseObjectResultFirewallAnomalyPackageStatus = "active"
)

func (r WAFPackageListResponseObjectResultFirewallAnomalyPackageStatus) IsKnown() bool {
	switch r {
	case WAFPackageListResponseObjectResultFirewallAnomalyPackageStatusActive:
		return true
	}
	return false
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
type WAFPackageListResponseObjectResultDetectionMode string

const (
	WAFPackageListResponseObjectResultDetectionModeAnomaly     WAFPackageListResponseObjectResultDetectionMode = "anomaly"
	WAFPackageListResponseObjectResultDetectionModeTraditional WAFPackageListResponseObjectResultDetectionMode = "traditional"
)

func (r WAFPackageListResponseObjectResultDetectionMode) IsKnown() bool {
	switch r {
	case WAFPackageListResponseObjectResultDetectionModeAnomaly, WAFPackageListResponseObjectResultDetectionModeTraditional:
		return true
	}
	return false
}

// When set to `active`, indicates that the WAF package will be applied to the
// zone.
type WAFPackageListResponseObjectResultStatus string

const (
	WAFPackageListResponseObjectResultStatusActive WAFPackageListResponseObjectResultStatus = "active"
)

func (r WAFPackageListResponseObjectResultStatus) IsKnown() bool {
	switch r {
	case WAFPackageListResponseObjectResultStatusActive:
		return true
	}
	return false
}

// The default action performed by the rules in the WAF package.
type WAFPackageListResponseObjectResultActionMode string

const (
	WAFPackageListResponseObjectResultActionModeSimulate  WAFPackageListResponseObjectResultActionMode = "simulate"
	WAFPackageListResponseObjectResultActionModeBlock     WAFPackageListResponseObjectResultActionMode = "block"
	WAFPackageListResponseObjectResultActionModeChallenge WAFPackageListResponseObjectResultActionMode = "challenge"
)

func (r WAFPackageListResponseObjectResultActionMode) IsKnown() bool {
	switch r {
	case WAFPackageListResponseObjectResultActionModeSimulate, WAFPackageListResponseObjectResultActionModeBlock, WAFPackageListResponseObjectResultActionModeChallenge:
		return true
	}
	return false
}

// The sensitivity of the WAF package.
type WAFPackageListResponseObjectResultSensitivity string

const (
	WAFPackageListResponseObjectResultSensitivityHigh   WAFPackageListResponseObjectResultSensitivity = "high"
	WAFPackageListResponseObjectResultSensitivityMedium WAFPackageListResponseObjectResultSensitivity = "medium"
	WAFPackageListResponseObjectResultSensitivityLow    WAFPackageListResponseObjectResultSensitivity = "low"
	WAFPackageListResponseObjectResultSensitivityOff    WAFPackageListResponseObjectResultSensitivity = "off"
)

func (r WAFPackageListResponseObjectResultSensitivity) IsKnown() bool {
	switch r {
	case WAFPackageListResponseObjectResultSensitivityHigh, WAFPackageListResponseObjectResultSensitivityMedium, WAFPackageListResponseObjectResultSensitivityLow, WAFPackageListResponseObjectResultSensitivityOff:
		return true
	}
	return false
}

// Whether the API call was successful
type WAFPackageListResponseSuccess bool

const (
	WAFPackageListResponseSuccessTrue WAFPackageListResponseSuccess = true
)

func (r WAFPackageListResponseSuccess) IsKnown() bool {
	switch r {
	case WAFPackageListResponseSuccessTrue:
		return true
	}
	return false
}

type WAFPackageGetResponse struct {
	Errors   interface{} `json:"errors,required"`
	Messages interface{} `json:"messages,required"`
	Result   interface{} `json:"result,required"`
	// Whether the API call was successful
	Success WAFPackageGetResponseSuccess `json:"success"`
	JSON    wafPackageGetResponseJSON    `json:"-"`
	union   WAFPackageGetResponseUnion
}

// wafPackageGetResponseJSON contains the JSON metadata for the struct
// [WAFPackageGetResponse]
type wafPackageGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r wafPackageGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *WAFPackageGetResponse) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r WAFPackageGetResponse) AsUnion() WAFPackageGetResponseUnion {
	return r.union
}

// Union satisfied by [firewall.WAFPackageGetResponseFirewallAPIResponseSingle] or
// [firewall.WAFPackageGetResponseObject].
type WAFPackageGetResponseUnion interface {
	implementsFirewallWAFPackageGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WAFPackageGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WAFPackageGetResponseFirewallAPIResponseSingle{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WAFPackageGetResponseObject{}),
		},
	)
}

type WAFPackageGetResponseFirewallAPIResponseSingle struct {
	Errors   []shared.ResponseInfo                                        `json:"errors,required"`
	Messages []shared.ResponseInfo                                        `json:"messages,required"`
	Result   shared.UnnamedSchemaRef8d6a37a1e4190f86652802244d29525fUnion `json:"result,required"`
	// Whether the API call was successful
	Success WAFPackageGetResponseFirewallAPIResponseSingleSuccess `json:"success,required"`
	JSON    wafPackageGetResponseFirewallAPIResponseSingleJSON    `json:"-"`
}

// wafPackageGetResponseFirewallAPIResponseSingleJSON contains the JSON metadata
// for the struct [WAFPackageGetResponseFirewallAPIResponseSingle]
type wafPackageGetResponseFirewallAPIResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFPackageGetResponseFirewallAPIResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageGetResponseFirewallAPIResponseSingleJSON) RawJSON() string {
	return r.raw
}

func (r WAFPackageGetResponseFirewallAPIResponseSingle) implementsFirewallWAFPackageGetResponse() {}

// Whether the API call was successful
type WAFPackageGetResponseFirewallAPIResponseSingleSuccess bool

const (
	WAFPackageGetResponseFirewallAPIResponseSingleSuccessTrue WAFPackageGetResponseFirewallAPIResponseSingleSuccess = true
)

func (r WAFPackageGetResponseFirewallAPIResponseSingleSuccess) IsKnown() bool {
	switch r {
	case WAFPackageGetResponseFirewallAPIResponseSingleSuccessTrue:
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

// Whether the API call was successful
type WAFPackageGetResponseSuccess bool

const (
	WAFPackageGetResponseSuccessTrue WAFPackageGetResponseSuccess = true
)

func (r WAFPackageGetResponseSuccess) IsKnown() bool {
	switch r {
	case WAFPackageGetResponseSuccessTrue:
		return true
	}
	return false
}

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
