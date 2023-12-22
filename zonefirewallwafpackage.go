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
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneFirewallWafPackageService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneFirewallWafPackageService]
// method instead.
type ZoneFirewallWafPackageService struct {
	Options []option.RequestOption
	Groups  *ZoneFirewallWafPackageGroupService
	Rules   *ZoneFirewallWafPackageRuleService
}

// NewZoneFirewallWafPackageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneFirewallWafPackageService(opts ...option.RequestOption) (r *ZoneFirewallWafPackageService) {
	r = &ZoneFirewallWafPackageService{}
	r.Options = opts
	r.Groups = NewZoneFirewallWafPackageGroupService(opts...)
	r.Rules = NewZoneFirewallWafPackageRuleService(opts...)
	return
}

// Fetches the details of a WAF package.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *ZoneFirewallWafPackageService) Get(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *PackageResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a WAF package. You can update the sensitivity and the action of an
// anomaly detection WAF package.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *ZoneFirewallWafPackageService) Update(ctx context.Context, zoneIdentifier string, identifier string, body ZoneFirewallWafPackageUpdateParams, opts ...option.RequestOption) (res *ZoneFirewallWafPackageUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Fetches WAF packages for a zone.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *ZoneFirewallWafPackageService) WafPackagesListWafPackages(ctx context.Context, zoneIdentifier string, query ZoneFirewallWafPackageWafPackagesListWafPackagesParams, opts ...option.RequestOption) (res *PackageResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/waf/packages", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Union satisfied by [APIResponseCollection] or [PackageResponseCollectionObject].
type PackageResponseCollection interface {
	implementsPackageResponseCollection()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PackageResponseCollection)(nil)).Elem(), "")
}

type PackageResponseCollectionObject struct {
	Result []PackageResponseCollectionObjectResult `json:"result"`
	JSON   packageResponseCollectionObjectJSON     `json:"-"`
}

// packageResponseCollectionObjectJSON contains the JSON metadata for the struct
// [PackageResponseCollectionObject]
type packageResponseCollectionObjectJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PackageResponseCollectionObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PackageResponseCollectionObject) implementsPackageResponseCollection() {}

// Union satisfied by [PackageResponseCollectionObjectResultPackageDefinition] or
// [PackageResponseCollectionObjectResultAnomalyPackage].
type PackageResponseCollectionObjectResult interface {
	implementsPackageResponseCollectionObjectResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PackageResponseCollectionObjectResult)(nil)).Elem(), "")
}

type PackageResponseCollectionObjectResultPackageDefinition struct {
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
	DetectionMode PackageResponseCollectionObjectResultPackageDefinitionDetectionMode `json:"detection_mode,required"`
	// The name of the WAF package.
	Name string `json:"name,required"`
	// Identifier
	ZoneID string `json:"zone_id,required"`
	// When set to `active`, indicates that the WAF package will be applied to the
	// zone.
	Status PackageResponseCollectionObjectResultPackageDefinitionStatus `json:"status"`
	JSON   packageResponseCollectionObjectResultPackageDefinitionJSON   `json:"-"`
}

// packageResponseCollectionObjectResultPackageDefinitionJSON contains the JSON
// metadata for the struct [PackageResponseCollectionObjectResultPackageDefinition]
type packageResponseCollectionObjectResultPackageDefinitionJSON struct {
	ID            apijson.Field
	Description   apijson.Field
	DetectionMode apijson.Field
	Name          apijson.Field
	ZoneID        apijson.Field
	Status        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PackageResponseCollectionObjectResultPackageDefinition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PackageResponseCollectionObjectResultPackageDefinition) implementsPackageResponseCollectionObjectResult() {
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
type PackageResponseCollectionObjectResultPackageDefinitionDetectionMode string

const (
	PackageResponseCollectionObjectResultPackageDefinitionDetectionModeAnomaly     PackageResponseCollectionObjectResultPackageDefinitionDetectionMode = "anomaly"
	PackageResponseCollectionObjectResultPackageDefinitionDetectionModeTraditional PackageResponseCollectionObjectResultPackageDefinitionDetectionMode = "traditional"
)

// When set to `active`, indicates that the WAF package will be applied to the
// zone.
type PackageResponseCollectionObjectResultPackageDefinitionStatus string

const (
	PackageResponseCollectionObjectResultPackageDefinitionStatusActive PackageResponseCollectionObjectResultPackageDefinitionStatus = "active"
)

type PackageResponseCollectionObjectResultAnomalyPackage struct {
	// The unique identifier of a WAF package.
	ID string `json:"id,required"`
	// The default action performed by the rules in the WAF package.
	ActionMode PackageResponseCollectionObjectResultAnomalyPackageActionMode `json:"action_mode,required"`
	// A summary of the purpose/function of the WAF package.
	Description string `json:"description,required"`
	// When a WAF package uses anomaly detection, each rule is given a score when
	// triggered. If the total score of all triggered rules exceeds the sensitivity
	// defined on the WAF package, the action defined on the package will be taken.
	DetectionMode string `json:"detection_mode,required"`
	// The name of the WAF package.
	Name string `json:"name,required"`
	// The sensitivity of the WAF package.
	Sensitivity PackageResponseCollectionObjectResultAnomalyPackageSensitivity `json:"sensitivity,required"`
	// Identifier
	ZoneID string `json:"zone_id,required"`
	// When set to `active`, indicates that the WAF package will be applied to the
	// zone.
	Status PackageResponseCollectionObjectResultAnomalyPackageStatus `json:"status"`
	JSON   packageResponseCollectionObjectResultAnomalyPackageJSON   `json:"-"`
}

// packageResponseCollectionObjectResultAnomalyPackageJSON contains the JSON
// metadata for the struct [PackageResponseCollectionObjectResultAnomalyPackage]
type packageResponseCollectionObjectResultAnomalyPackageJSON struct {
	ID            apijson.Field
	ActionMode    apijson.Field
	Description   apijson.Field
	DetectionMode apijson.Field
	Name          apijson.Field
	Sensitivity   apijson.Field
	ZoneID        apijson.Field
	Status        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PackageResponseCollectionObjectResultAnomalyPackage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PackageResponseCollectionObjectResultAnomalyPackage) implementsPackageResponseCollectionObjectResult() {
}

// The default action performed by the rules in the WAF package.
type PackageResponseCollectionObjectResultAnomalyPackageActionMode string

const (
	PackageResponseCollectionObjectResultAnomalyPackageActionModeSimulate  PackageResponseCollectionObjectResultAnomalyPackageActionMode = "simulate"
	PackageResponseCollectionObjectResultAnomalyPackageActionModeBlock     PackageResponseCollectionObjectResultAnomalyPackageActionMode = "block"
	PackageResponseCollectionObjectResultAnomalyPackageActionModeChallenge PackageResponseCollectionObjectResultAnomalyPackageActionMode = "challenge"
)

// The sensitivity of the WAF package.
type PackageResponseCollectionObjectResultAnomalyPackageSensitivity string

const (
	PackageResponseCollectionObjectResultAnomalyPackageSensitivityHigh   PackageResponseCollectionObjectResultAnomalyPackageSensitivity = "high"
	PackageResponseCollectionObjectResultAnomalyPackageSensitivityMedium PackageResponseCollectionObjectResultAnomalyPackageSensitivity = "medium"
	PackageResponseCollectionObjectResultAnomalyPackageSensitivityLow    PackageResponseCollectionObjectResultAnomalyPackageSensitivity = "low"
	PackageResponseCollectionObjectResultAnomalyPackageSensitivityOff    PackageResponseCollectionObjectResultAnomalyPackageSensitivity = "off"
)

// When set to `active`, indicates that the WAF package will be applied to the
// zone.
type PackageResponseCollectionObjectResultAnomalyPackageStatus string

const (
	PackageResponseCollectionObjectResultAnomalyPackageStatusActive PackageResponseCollectionObjectResultAnomalyPackageStatus = "active"
)

// Union satisfied by [APIResponseSingle] or [PackageResponseSingleObject].
type PackageResponseSingle interface {
	implementsPackageResponseSingle()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*PackageResponseSingle)(nil)).Elem(), "")
}

type PackageResponseSingleObject struct {
	Result interface{}                     `json:"result"`
	JSON   packageResponseSingleObjectJSON `json:"-"`
}

// packageResponseSingleObjectJSON contains the JSON metadata for the struct
// [PackageResponseSingleObject]
type packageResponseSingleObjectJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PackageResponseSingleObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PackageResponseSingleObject) implementsPackageResponseSingle() {}

type ZoneFirewallWafPackageUpdateResponse struct {
	Result ZoneFirewallWafPackageUpdateResponseResult `json:"result"`
	JSON   zoneFirewallWafPackageUpdateResponseJSON   `json:"-"`
}

// zoneFirewallWafPackageUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneFirewallWafPackageUpdateResponse]
type zoneFirewallWafPackageUpdateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallWafPackageUpdateResponseResult struct {
	// The unique identifier of a WAF package.
	ID string `json:"id,required"`
	// The default action performed by the rules in the WAF package.
	ActionMode ZoneFirewallWafPackageUpdateResponseResultActionMode `json:"action_mode,required"`
	// A summary of the purpose/function of the WAF package.
	Description string `json:"description,required"`
	// When a WAF package uses anomaly detection, each rule is given a score when
	// triggered. If the total score of all triggered rules exceeds the sensitivity
	// defined on the WAF package, the action defined on the package will be taken.
	DetectionMode string `json:"detection_mode,required"`
	// The name of the WAF package.
	Name string `json:"name,required"`
	// The sensitivity of the WAF package.
	Sensitivity ZoneFirewallWafPackageUpdateResponseResultSensitivity `json:"sensitivity,required"`
	// Identifier
	ZoneID string `json:"zone_id,required"`
	// When set to `active`, indicates that the WAF package will be applied to the
	// zone.
	Status ZoneFirewallWafPackageUpdateResponseResultStatus `json:"status"`
	JSON   zoneFirewallWafPackageUpdateResponseResultJSON   `json:"-"`
}

// zoneFirewallWafPackageUpdateResponseResultJSON contains the JSON metadata for
// the struct [ZoneFirewallWafPackageUpdateResponseResult]
type zoneFirewallWafPackageUpdateResponseResultJSON struct {
	ID            apijson.Field
	ActionMode    apijson.Field
	Description   apijson.Field
	DetectionMode apijson.Field
	Name          apijson.Field
	Sensitivity   apijson.Field
	ZoneID        apijson.Field
	Status        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The default action performed by the rules in the WAF package.
type ZoneFirewallWafPackageUpdateResponseResultActionMode string

const (
	ZoneFirewallWafPackageUpdateResponseResultActionModeSimulate  ZoneFirewallWafPackageUpdateResponseResultActionMode = "simulate"
	ZoneFirewallWafPackageUpdateResponseResultActionModeBlock     ZoneFirewallWafPackageUpdateResponseResultActionMode = "block"
	ZoneFirewallWafPackageUpdateResponseResultActionModeChallenge ZoneFirewallWafPackageUpdateResponseResultActionMode = "challenge"
)

// The sensitivity of the WAF package.
type ZoneFirewallWafPackageUpdateResponseResultSensitivity string

const (
	ZoneFirewallWafPackageUpdateResponseResultSensitivityHigh   ZoneFirewallWafPackageUpdateResponseResultSensitivity = "high"
	ZoneFirewallWafPackageUpdateResponseResultSensitivityMedium ZoneFirewallWafPackageUpdateResponseResultSensitivity = "medium"
	ZoneFirewallWafPackageUpdateResponseResultSensitivityLow    ZoneFirewallWafPackageUpdateResponseResultSensitivity = "low"
	ZoneFirewallWafPackageUpdateResponseResultSensitivityOff    ZoneFirewallWafPackageUpdateResponseResultSensitivity = "off"
)

// When set to `active`, indicates that the WAF package will be applied to the
// zone.
type ZoneFirewallWafPackageUpdateResponseResultStatus string

const (
	ZoneFirewallWafPackageUpdateResponseResultStatusActive ZoneFirewallWafPackageUpdateResponseResultStatus = "active"
)

type ZoneFirewallWafPackageUpdateParams struct {
	// The default action performed by the rules in the WAF package.
	ActionMode param.Field[ZoneFirewallWafPackageUpdateParamsActionMode] `json:"action_mode"`
	// The sensitivity of the WAF package.
	Sensitivity param.Field[ZoneFirewallWafPackageUpdateParamsSensitivity] `json:"sensitivity"`
}

func (r ZoneFirewallWafPackageUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The default action performed by the rules in the WAF package.
type ZoneFirewallWafPackageUpdateParamsActionMode string

const (
	ZoneFirewallWafPackageUpdateParamsActionModeSimulate  ZoneFirewallWafPackageUpdateParamsActionMode = "simulate"
	ZoneFirewallWafPackageUpdateParamsActionModeBlock     ZoneFirewallWafPackageUpdateParamsActionMode = "block"
	ZoneFirewallWafPackageUpdateParamsActionModeChallenge ZoneFirewallWafPackageUpdateParamsActionMode = "challenge"
)

// The sensitivity of the WAF package.
type ZoneFirewallWafPackageUpdateParamsSensitivity string

const (
	ZoneFirewallWafPackageUpdateParamsSensitivityHigh   ZoneFirewallWafPackageUpdateParamsSensitivity = "high"
	ZoneFirewallWafPackageUpdateParamsSensitivityMedium ZoneFirewallWafPackageUpdateParamsSensitivity = "medium"
	ZoneFirewallWafPackageUpdateParamsSensitivityLow    ZoneFirewallWafPackageUpdateParamsSensitivity = "low"
	ZoneFirewallWafPackageUpdateParamsSensitivityOff    ZoneFirewallWafPackageUpdateParamsSensitivity = "off"
)

type ZoneFirewallWafPackageWafPackagesListWafPackagesParams struct {
	// The direction used to sort returned packages.
	Direction param.Field[ZoneFirewallWafPackageWafPackagesListWafPackagesParamsDirection] `query:"direction"`
	// When set to `all`, all the search requirements must match. When set to `any`,
	// only one of the search requirements has to match.
	Match param.Field[ZoneFirewallWafPackageWafPackagesListWafPackagesParamsMatch] `query:"match"`
	// The field used to sort returned packages.
	Order param.Field[ZoneFirewallWafPackageWafPackagesListWafPackagesParamsOrder] `query:"order"`
	// The page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The number of packages per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [ZoneFirewallWafPackageWafPackagesListWafPackagesParams]'s
// query parameters as `url.Values`.
func (r ZoneFirewallWafPackageWafPackagesListWafPackagesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The direction used to sort returned packages.
type ZoneFirewallWafPackageWafPackagesListWafPackagesParamsDirection string

const (
	ZoneFirewallWafPackageWafPackagesListWafPackagesParamsDirectionAsc  ZoneFirewallWafPackageWafPackagesListWafPackagesParamsDirection = "asc"
	ZoneFirewallWafPackageWafPackagesListWafPackagesParamsDirectionDesc ZoneFirewallWafPackageWafPackagesListWafPackagesParamsDirection = "desc"
)

// When set to `all`, all the search requirements must match. When set to `any`,
// only one of the search requirements has to match.
type ZoneFirewallWafPackageWafPackagesListWafPackagesParamsMatch string

const (
	ZoneFirewallWafPackageWafPackagesListWafPackagesParamsMatchAny ZoneFirewallWafPackageWafPackagesListWafPackagesParamsMatch = "any"
	ZoneFirewallWafPackageWafPackagesListWafPackagesParamsMatchAll ZoneFirewallWafPackageWafPackagesListWafPackagesParamsMatch = "all"
)

// The field used to sort returned packages.
type ZoneFirewallWafPackageWafPackagesListWafPackagesParamsOrder string

const (
	ZoneFirewallWafPackageWafPackagesListWafPackagesParamsOrderName ZoneFirewallWafPackageWafPackagesListWafPackagesParamsOrder = "name"
)
