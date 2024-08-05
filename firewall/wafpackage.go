// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
	"github.com/tidwall/gjson"
)

// WAFPackageService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWAFPackageService] method instead.
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
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if zoneIdentifier == "" {
		err = errors.New("missing required zone_identifier parameter")
		return
	}
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
	if zoneIdentifier == "" {
		err = errors.New("missing required zone_identifier parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type WAFPackageListResponse struct {
	// This field can have the runtime type of [[]shared.ResponseInfo].
	Errors interface{} `json:"errors,required"`
	// This field can have the runtime type of [[]shared.ResponseInfo].
	Messages interface{} `json:"messages,required"`
	// This field can have the runtime type of
	// [WAFPackageListResponseFirewallAPIResponseCollectionResultUnion],
	// [[]WAFPackageListResponseResultResult].
	Result interface{} `json:"result,required"`
	// Whether the API call was successful
	Success WAFPackageListResponseSuccess `json:"success"`
	// This field can have the runtime type of
	// [WAFPackageListResponseFirewallAPIResponseCollectionResultInfo].
	ResultInfo interface{}                `json:"result_info,required"`
	JSON       wafPackageListResponseJSON `json:"-"`
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
	*r = WAFPackageListResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [WAFPackageListResponseUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are
// [firewall.WAFPackageListResponseFirewallAPIResponseCollection],
// [firewall.WAFPackageListResponseResult].
func (r WAFPackageListResponse) AsUnion() WAFPackageListResponseUnion {
	return r.union
}

// Union satisfied by
// [firewall.WAFPackageListResponseFirewallAPIResponseCollection] or
// [firewall.WAFPackageListResponseResult].
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
			Type:       reflect.TypeOf(WAFPackageListResponseResult{}),
		},
	)
}

type WAFPackageListResponseFirewallAPIResponseCollection struct {
	Errors   []shared.ResponseInfo                                          `json:"errors,required"`
	Messages []shared.ResponseInfo                                          `json:"messages,required"`
	Result   WAFPackageListResponseFirewallAPIResponseCollectionResultUnion `json:"result,required,nullable"`
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

// Union satisfied by
// [firewall.WAFPackageListResponseFirewallAPIResponseCollectionResultUnknown],
// [firewall.WAFPackageListResponseFirewallAPIResponseCollectionResultArray] or
// [shared.UnionString].
type WAFPackageListResponseFirewallAPIResponseCollectionResultUnion interface {
	ImplementsFirewallWAFPackageListResponseFirewallAPIResponseCollectionResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WAFPackageListResponseFirewallAPIResponseCollectionResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WAFPackageListResponseFirewallAPIResponseCollectionResultArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type WAFPackageListResponseFirewallAPIResponseCollectionResultArray []interface{}

func (r WAFPackageListResponseFirewallAPIResponseCollectionResultArray) ImplementsFirewallWAFPackageListResponseFirewallAPIResponseCollectionResultUnion() {
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

type WAFPackageListResponseResult struct {
	Result []WAFPackageListResponseResultResult `json:"result"`
	JSON   wafPackageListResponseResultJSON     `json:"-"`
}

// wafPackageListResponseResultJSON contains the JSON metadata for the struct
// [WAFPackageListResponseResult]
type wafPackageListResponseResultJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFPackageListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageListResponseResultJSON) RawJSON() string {
	return r.raw
}

func (r WAFPackageListResponseResult) implementsFirewallWAFPackageListResponse() {}

type WAFPackageListResponseResultResult struct {
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
	DetectionMode WAFPackageListResponseResultResultDetectionMode `json:"detection_mode,required"`
	// Identifier
	ID string `json:"id,required"`
	// The name of the WAF package.
	Name string `json:"name,required"`
	// When set to `active`, indicates that the WAF package will be applied to the
	// zone.
	Status WAFPackageListResponseResultResultStatus `json:"status"`
	// Identifier
	ZoneID string `json:"zone_id,required"`
	// The default action performed by the rules in the WAF package.
	ActionMode WAFPackageListResponseResultResultActionMode `json:"action_mode"`
	// The sensitivity of the WAF package.
	Sensitivity WAFPackageListResponseResultResultSensitivity `json:"sensitivity"`
	JSON        wafPackageListResponseResultResultJSON        `json:"-"`
	union       WAFPackageListResponseResultResultUnion
}

// wafPackageListResponseResultResultJSON contains the JSON metadata for the struct
// [WAFPackageListResponseResultResult]
type wafPackageListResponseResultResultJSON struct {
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

func (r wafPackageListResponseResultResultJSON) RawJSON() string {
	return r.raw
}

func (r *WAFPackageListResponseResultResult) UnmarshalJSON(data []byte) (err error) {
	*r = WAFPackageListResponseResultResult{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [WAFPackageListResponseResultResultUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [firewall.WAFPackageListResponseResultResultFirewallPackageDefinition],
// [firewall.WAFPackageListResponseResultResultFirewallAnomalyPackage].
func (r WAFPackageListResponseResultResult) AsUnion() WAFPackageListResponseResultResultUnion {
	return r.union
}

// Union satisfied by
// [firewall.WAFPackageListResponseResultResultFirewallPackageDefinition] or
// [firewall.WAFPackageListResponseResultResultFirewallAnomalyPackage].
type WAFPackageListResponseResultResultUnion interface {
	implementsFirewallWAFPackageListResponseResultResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WAFPackageListResponseResultResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WAFPackageListResponseResultResultFirewallPackageDefinition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WAFPackageListResponseResultResultFirewallAnomalyPackage{}),
		},
	)
}

type WAFPackageListResponseResultResultFirewallPackageDefinition struct {
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
	DetectionMode WAFPackageListResponseResultResultFirewallPackageDefinitionDetectionMode `json:"detection_mode,required"`
	// The name of the WAF package.
	Name string `json:"name,required"`
	// Identifier
	ZoneID string `json:"zone_id,required"`
	// When set to `active`, indicates that the WAF package will be applied to the
	// zone.
	Status WAFPackageListResponseResultResultFirewallPackageDefinitionStatus `json:"status"`
	JSON   wafPackageListResponseResultResultFirewallPackageDefinitionJSON   `json:"-"`
}

// wafPackageListResponseResultResultFirewallPackageDefinitionJSON contains the
// JSON metadata for the struct
// [WAFPackageListResponseResultResultFirewallPackageDefinition]
type wafPackageListResponseResultResultFirewallPackageDefinitionJSON struct {
	ID            apijson.Field
	Description   apijson.Field
	DetectionMode apijson.Field
	Name          apijson.Field
	ZoneID        apijson.Field
	Status        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WAFPackageListResponseResultResultFirewallPackageDefinition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageListResponseResultResultFirewallPackageDefinitionJSON) RawJSON() string {
	return r.raw
}

func (r WAFPackageListResponseResultResultFirewallPackageDefinition) implementsFirewallWAFPackageListResponseResultResult() {
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
type WAFPackageListResponseResultResultFirewallPackageDefinitionDetectionMode string

const (
	WAFPackageListResponseResultResultFirewallPackageDefinitionDetectionModeAnomaly     WAFPackageListResponseResultResultFirewallPackageDefinitionDetectionMode = "anomaly"
	WAFPackageListResponseResultResultFirewallPackageDefinitionDetectionModeTraditional WAFPackageListResponseResultResultFirewallPackageDefinitionDetectionMode = "traditional"
)

func (r WAFPackageListResponseResultResultFirewallPackageDefinitionDetectionMode) IsKnown() bool {
	switch r {
	case WAFPackageListResponseResultResultFirewallPackageDefinitionDetectionModeAnomaly, WAFPackageListResponseResultResultFirewallPackageDefinitionDetectionModeTraditional:
		return true
	}
	return false
}

// When set to `active`, indicates that the WAF package will be applied to the
// zone.
type WAFPackageListResponseResultResultFirewallPackageDefinitionStatus string

const (
	WAFPackageListResponseResultResultFirewallPackageDefinitionStatusActive WAFPackageListResponseResultResultFirewallPackageDefinitionStatus = "active"
)

func (r WAFPackageListResponseResultResultFirewallPackageDefinitionStatus) IsKnown() bool {
	switch r {
	case WAFPackageListResponseResultResultFirewallPackageDefinitionStatusActive:
		return true
	}
	return false
}

type WAFPackageListResponseResultResultFirewallAnomalyPackage struct {
	// Identifier
	ID string `json:"id,required"`
	// A summary of the purpose/function of the WAF package.
	Description string `json:"description,required"`
	// When a WAF package uses anomaly detection, each rule is given a score when
	// triggered. If the total score of all triggered rules exceeds the sensitivity
	// defined on the WAF package, the action defined on the package will be taken.
	DetectionMode WAFPackageListResponseResultResultFirewallAnomalyPackageDetectionMode `json:"detection_mode,required"`
	// The name of the WAF package.
	Name string `json:"name,required"`
	// Identifier
	ZoneID string `json:"zone_id,required"`
	// The default action performed by the rules in the WAF package.
	ActionMode WAFPackageListResponseResultResultFirewallAnomalyPackageActionMode `json:"action_mode"`
	// The sensitivity of the WAF package.
	Sensitivity WAFPackageListResponseResultResultFirewallAnomalyPackageSensitivity `json:"sensitivity"`
	// When set to `active`, indicates that the WAF package will be applied to the
	// zone.
	Status WAFPackageListResponseResultResultFirewallAnomalyPackageStatus `json:"status"`
	JSON   wafPackageListResponseResultResultFirewallAnomalyPackageJSON   `json:"-"`
}

// wafPackageListResponseResultResultFirewallAnomalyPackageJSON contains the JSON
// metadata for the struct
// [WAFPackageListResponseResultResultFirewallAnomalyPackage]
type wafPackageListResponseResultResultFirewallAnomalyPackageJSON struct {
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

func (r *WAFPackageListResponseResultResultFirewallAnomalyPackage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageListResponseResultResultFirewallAnomalyPackageJSON) RawJSON() string {
	return r.raw
}

func (r WAFPackageListResponseResultResultFirewallAnomalyPackage) implementsFirewallWAFPackageListResponseResultResult() {
}

// When a WAF package uses anomaly detection, each rule is given a score when
// triggered. If the total score of all triggered rules exceeds the sensitivity
// defined on the WAF package, the action defined on the package will be taken.
type WAFPackageListResponseResultResultFirewallAnomalyPackageDetectionMode string

const (
	WAFPackageListResponseResultResultFirewallAnomalyPackageDetectionModeAnomaly     WAFPackageListResponseResultResultFirewallAnomalyPackageDetectionMode = "anomaly"
	WAFPackageListResponseResultResultFirewallAnomalyPackageDetectionModeTraditional WAFPackageListResponseResultResultFirewallAnomalyPackageDetectionMode = "traditional"
)

func (r WAFPackageListResponseResultResultFirewallAnomalyPackageDetectionMode) IsKnown() bool {
	switch r {
	case WAFPackageListResponseResultResultFirewallAnomalyPackageDetectionModeAnomaly, WAFPackageListResponseResultResultFirewallAnomalyPackageDetectionModeTraditional:
		return true
	}
	return false
}

// The default action performed by the rules in the WAF package.
type WAFPackageListResponseResultResultFirewallAnomalyPackageActionMode string

const (
	WAFPackageListResponseResultResultFirewallAnomalyPackageActionModeSimulate  WAFPackageListResponseResultResultFirewallAnomalyPackageActionMode = "simulate"
	WAFPackageListResponseResultResultFirewallAnomalyPackageActionModeBlock     WAFPackageListResponseResultResultFirewallAnomalyPackageActionMode = "block"
	WAFPackageListResponseResultResultFirewallAnomalyPackageActionModeChallenge WAFPackageListResponseResultResultFirewallAnomalyPackageActionMode = "challenge"
)

func (r WAFPackageListResponseResultResultFirewallAnomalyPackageActionMode) IsKnown() bool {
	switch r {
	case WAFPackageListResponseResultResultFirewallAnomalyPackageActionModeSimulate, WAFPackageListResponseResultResultFirewallAnomalyPackageActionModeBlock, WAFPackageListResponseResultResultFirewallAnomalyPackageActionModeChallenge:
		return true
	}
	return false
}

// The sensitivity of the WAF package.
type WAFPackageListResponseResultResultFirewallAnomalyPackageSensitivity string

const (
	WAFPackageListResponseResultResultFirewallAnomalyPackageSensitivityHigh   WAFPackageListResponseResultResultFirewallAnomalyPackageSensitivity = "high"
	WAFPackageListResponseResultResultFirewallAnomalyPackageSensitivityMedium WAFPackageListResponseResultResultFirewallAnomalyPackageSensitivity = "medium"
	WAFPackageListResponseResultResultFirewallAnomalyPackageSensitivityLow    WAFPackageListResponseResultResultFirewallAnomalyPackageSensitivity = "low"
	WAFPackageListResponseResultResultFirewallAnomalyPackageSensitivityOff    WAFPackageListResponseResultResultFirewallAnomalyPackageSensitivity = "off"
)

func (r WAFPackageListResponseResultResultFirewallAnomalyPackageSensitivity) IsKnown() bool {
	switch r {
	case WAFPackageListResponseResultResultFirewallAnomalyPackageSensitivityHigh, WAFPackageListResponseResultResultFirewallAnomalyPackageSensitivityMedium, WAFPackageListResponseResultResultFirewallAnomalyPackageSensitivityLow, WAFPackageListResponseResultResultFirewallAnomalyPackageSensitivityOff:
		return true
	}
	return false
}

// When set to `active`, indicates that the WAF package will be applied to the
// zone.
type WAFPackageListResponseResultResultFirewallAnomalyPackageStatus string

const (
	WAFPackageListResponseResultResultFirewallAnomalyPackageStatusActive WAFPackageListResponseResultResultFirewallAnomalyPackageStatus = "active"
)

func (r WAFPackageListResponseResultResultFirewallAnomalyPackageStatus) IsKnown() bool {
	switch r {
	case WAFPackageListResponseResultResultFirewallAnomalyPackageStatusActive:
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
type WAFPackageListResponseResultResultDetectionMode string

const (
	WAFPackageListResponseResultResultDetectionModeAnomaly     WAFPackageListResponseResultResultDetectionMode = "anomaly"
	WAFPackageListResponseResultResultDetectionModeTraditional WAFPackageListResponseResultResultDetectionMode = "traditional"
)

func (r WAFPackageListResponseResultResultDetectionMode) IsKnown() bool {
	switch r {
	case WAFPackageListResponseResultResultDetectionModeAnomaly, WAFPackageListResponseResultResultDetectionModeTraditional:
		return true
	}
	return false
}

// When set to `active`, indicates that the WAF package will be applied to the
// zone.
type WAFPackageListResponseResultResultStatus string

const (
	WAFPackageListResponseResultResultStatusActive WAFPackageListResponseResultResultStatus = "active"
)

func (r WAFPackageListResponseResultResultStatus) IsKnown() bool {
	switch r {
	case WAFPackageListResponseResultResultStatusActive:
		return true
	}
	return false
}

// The default action performed by the rules in the WAF package.
type WAFPackageListResponseResultResultActionMode string

const (
	WAFPackageListResponseResultResultActionModeSimulate  WAFPackageListResponseResultResultActionMode = "simulate"
	WAFPackageListResponseResultResultActionModeBlock     WAFPackageListResponseResultResultActionMode = "block"
	WAFPackageListResponseResultResultActionModeChallenge WAFPackageListResponseResultResultActionMode = "challenge"
)

func (r WAFPackageListResponseResultResultActionMode) IsKnown() bool {
	switch r {
	case WAFPackageListResponseResultResultActionModeSimulate, WAFPackageListResponseResultResultActionModeBlock, WAFPackageListResponseResultResultActionModeChallenge:
		return true
	}
	return false
}

// The sensitivity of the WAF package.
type WAFPackageListResponseResultResultSensitivity string

const (
	WAFPackageListResponseResultResultSensitivityHigh   WAFPackageListResponseResultResultSensitivity = "high"
	WAFPackageListResponseResultResultSensitivityMedium WAFPackageListResponseResultResultSensitivity = "medium"
	WAFPackageListResponseResultResultSensitivityLow    WAFPackageListResponseResultResultSensitivity = "low"
	WAFPackageListResponseResultResultSensitivityOff    WAFPackageListResponseResultResultSensitivity = "off"
)

func (r WAFPackageListResponseResultResultSensitivity) IsKnown() bool {
	switch r {
	case WAFPackageListResponseResultResultSensitivityHigh, WAFPackageListResponseResultResultSensitivityMedium, WAFPackageListResponseResultResultSensitivityLow, WAFPackageListResponseResultResultSensitivityOff:
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
	// This field can have the runtime type of [[]shared.ResponseInfo].
	Errors interface{} `json:"errors,required"`
	// This field can have the runtime type of [[]shared.ResponseInfo].
	Messages interface{} `json:"messages,required"`
	// This field can have the runtime type of
	// [WAFPackageGetResponseFirewallAPIResponseSingleResultUnion], [interface{}].
	Result interface{} `json:"result,required"`
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
	*r = WAFPackageGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [WAFPackageGetResponseUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are
// [firewall.WAFPackageGetResponseFirewallAPIResponseSingle],
// [firewall.WAFPackageGetResponseResult].
func (r WAFPackageGetResponse) AsUnion() WAFPackageGetResponseUnion {
	return r.union
}

// Union satisfied by [firewall.WAFPackageGetResponseFirewallAPIResponseSingle] or
// [firewall.WAFPackageGetResponseResult].
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
			Type:       reflect.TypeOf(WAFPackageGetResponseResult{}),
		},
	)
}

type WAFPackageGetResponseFirewallAPIResponseSingle struct {
	Errors   []shared.ResponseInfo                                     `json:"errors,required"`
	Messages []shared.ResponseInfo                                     `json:"messages,required"`
	Result   WAFPackageGetResponseFirewallAPIResponseSingleResultUnion `json:"result,required"`
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

// Union satisfied by
// [firewall.WAFPackageGetResponseFirewallAPIResponseSingleResultUnknown] or
// [shared.UnionString].
type WAFPackageGetResponseFirewallAPIResponseSingleResultUnion interface {
	ImplementsFirewallWAFPackageGetResponseFirewallAPIResponseSingleResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WAFPackageGetResponseFirewallAPIResponseSingleResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

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

type WAFPackageGetResponseResult struct {
	Result interface{}                     `json:"result"`
	JSON   wafPackageGetResponseResultJSON `json:"-"`
}

// wafPackageGetResponseResultJSON contains the JSON metadata for the struct
// [WAFPackageGetResponseResult]
type wafPackageGetResponseResultJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFPackageGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageGetResponseResultJSON) RawJSON() string {
	return r.raw
}

func (r WAFPackageGetResponseResult) implementsFirewallWAFPackageGetResponse() {}

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
		NestedFormat: apiquery.NestedQueryFormatDots,
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
