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
func (r *ZoneFirewallWafPackageService) Get(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *ZoneFirewallWafPackageGetResponse, err error) {
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
func (r *ZoneFirewallWafPackageService) WafPackagesListWafPackages(ctx context.Context, zoneIdentifier string, query ZoneFirewallWafPackageWafPackagesListWafPackagesParams, opts ...option.RequestOption) (res *ZoneFirewallWafPackageWafPackagesListWafPackagesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/waf/packages", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Union satisfied by [ZoneFirewallWafPackageGetResponseDZw70ubTAPIResponseSingle]
// or [ZoneFirewallWafPackageGetResponseObject].
type ZoneFirewallWafPackageGetResponse interface {
	implementsZoneFirewallWafPackageGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneFirewallWafPackageGetResponse)(nil)).Elem(), "")
}

type ZoneFirewallWafPackageGetResponseDZw70ubTAPIResponseSingle struct {
	Errors   []ZoneFirewallWafPackageGetResponseDZw70ubTAPIResponseSingleError   `json:"errors"`
	Messages []ZoneFirewallWafPackageGetResponseDZw70ubTAPIResponseSingleMessage `json:"messages"`
	Result   ZoneFirewallWafPackageGetResponseDZw70ubTAPIResponseSingleResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success ZoneFirewallWafPackageGetResponseDZw70ubTAPIResponseSingleSuccess `json:"success"`
	JSON    zoneFirewallWafPackageGetResponseDZw70ubTapiResponseSingleJSON    `json:"-"`
}

// zoneFirewallWafPackageGetResponseDZw70ubTapiResponseSingleJSON contains the JSON
// metadata for the struct
// [ZoneFirewallWafPackageGetResponseDZw70ubTAPIResponseSingle]
type zoneFirewallWafPackageGetResponseDZw70ubTapiResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageGetResponseDZw70ubTAPIResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallWafPackageGetResponseDZw70ubTAPIResponseSingle) implementsZoneFirewallWafPackageGetResponse() {
}

type ZoneFirewallWafPackageGetResponseDZw70ubTAPIResponseSingleError struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    zoneFirewallWafPackageGetResponseDZw70ubTapiResponseSingleErrorJSON `json:"-"`
}

// zoneFirewallWafPackageGetResponseDZw70ubTapiResponseSingleErrorJSON contains the
// JSON metadata for the struct
// [ZoneFirewallWafPackageGetResponseDZw70ubTAPIResponseSingleError]
type zoneFirewallWafPackageGetResponseDZw70ubTapiResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageGetResponseDZw70ubTAPIResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallWafPackageGetResponseDZw70ubTAPIResponseSingleMessage struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    zoneFirewallWafPackageGetResponseDZw70ubTapiResponseSingleMessageJSON `json:"-"`
}

// zoneFirewallWafPackageGetResponseDZw70ubTapiResponseSingleMessageJSON contains
// the JSON metadata for the struct
// [ZoneFirewallWafPackageGetResponseDZw70ubTAPIResponseSingleMessage]
type zoneFirewallWafPackageGetResponseDZw70ubTapiResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageGetResponseDZw70ubTAPIResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [ZoneFirewallWafPackageGetResponseDZw70ubTAPIResponseSingleResultUnknown] or
// [shared.UnionString].
type ZoneFirewallWafPackageGetResponseDZw70ubTAPIResponseSingleResult interface {
	ImplementsZoneFirewallWafPackageGetResponseDZw70ubTapiResponseSingleResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneFirewallWafPackageGetResponseDZw70ubTAPIResponseSingleResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Whether the API call was successful
type ZoneFirewallWafPackageGetResponseDZw70ubTAPIResponseSingleSuccess bool

const (
	ZoneFirewallWafPackageGetResponseDZw70ubTAPIResponseSingleSuccessTrue ZoneFirewallWafPackageGetResponseDZw70ubTAPIResponseSingleSuccess = true
)

type ZoneFirewallWafPackageGetResponseObject struct {
	Result interface{}                                 `json:"result"`
	JSON   zoneFirewallWafPackageGetResponseObjectJSON `json:"-"`
}

// zoneFirewallWafPackageGetResponseObjectJSON contains the JSON metadata for the
// struct [ZoneFirewallWafPackageGetResponseObject]
type zoneFirewallWafPackageGetResponseObjectJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageGetResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallWafPackageGetResponseObject) implementsZoneFirewallWafPackageGetResponse() {}

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

// Union satisfied by
// [ZoneFirewallWafPackageWafPackagesListWafPackagesResponseDZw70ubTAPIResponseCollection]
// or [ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObject].
type ZoneFirewallWafPackageWafPackagesListWafPackagesResponse interface {
	implementsZoneFirewallWafPackageWafPackagesListWafPackagesResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneFirewallWafPackageWafPackagesListWafPackagesResponse)(nil)).Elem(), "")
}

type ZoneFirewallWafPackageWafPackagesListWafPackagesResponseDZw70ubTAPIResponseCollection struct {
	Errors     []ZoneFirewallWafPackageWafPackagesListWafPackagesResponseDZw70ubTAPIResponseCollectionError    `json:"errors"`
	Messages   []ZoneFirewallWafPackageWafPackagesListWafPackagesResponseDZw70ubTAPIResponseCollectionMessage  `json:"messages"`
	Result     []interface{}                                                                                   `json:"result,nullable"`
	ResultInfo ZoneFirewallWafPackageWafPackagesListWafPackagesResponseDZw70ubTAPIResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZoneFirewallWafPackageWafPackagesListWafPackagesResponseDZw70ubTAPIResponseCollectionSuccess `json:"success"`
	JSON    zoneFirewallWafPackageWafPackagesListWafPackagesResponseDZw70ubTapiResponseCollectionJSON    `json:"-"`
}

// zoneFirewallWafPackageWafPackagesListWafPackagesResponseDZw70ubTapiResponseCollectionJSON
// contains the JSON metadata for the struct
// [ZoneFirewallWafPackageWafPackagesListWafPackagesResponseDZw70ubTAPIResponseCollection]
type zoneFirewallWafPackageWafPackagesListWafPackagesResponseDZw70ubTapiResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageWafPackagesListWafPackagesResponseDZw70ubTAPIResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallWafPackageWafPackagesListWafPackagesResponseDZw70ubTAPIResponseCollection) implementsZoneFirewallWafPackageWafPackagesListWafPackagesResponse() {
}

type ZoneFirewallWafPackageWafPackagesListWafPackagesResponseDZw70ubTAPIResponseCollectionError struct {
	Code    int64                                                                                          `json:"code,required"`
	Message string                                                                                         `json:"message,required"`
	JSON    zoneFirewallWafPackageWafPackagesListWafPackagesResponseDZw70ubTapiResponseCollectionErrorJSON `json:"-"`
}

// zoneFirewallWafPackageWafPackagesListWafPackagesResponseDZw70ubTapiResponseCollectionErrorJSON
// contains the JSON metadata for the struct
// [ZoneFirewallWafPackageWafPackagesListWafPackagesResponseDZw70ubTAPIResponseCollectionError]
type zoneFirewallWafPackageWafPackagesListWafPackagesResponseDZw70ubTapiResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageWafPackagesListWafPackagesResponseDZw70ubTAPIResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallWafPackageWafPackagesListWafPackagesResponseDZw70ubTAPIResponseCollectionMessage struct {
	Code    int64                                                                                            `json:"code,required"`
	Message string                                                                                           `json:"message,required"`
	JSON    zoneFirewallWafPackageWafPackagesListWafPackagesResponseDZw70ubTapiResponseCollectionMessageJSON `json:"-"`
}

// zoneFirewallWafPackageWafPackagesListWafPackagesResponseDZw70ubTapiResponseCollectionMessageJSON
// contains the JSON metadata for the struct
// [ZoneFirewallWafPackageWafPackagesListWafPackagesResponseDZw70ubTAPIResponseCollectionMessage]
type zoneFirewallWafPackageWafPackagesListWafPackagesResponseDZw70ubTapiResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageWafPackagesListWafPackagesResponseDZw70ubTAPIResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallWafPackageWafPackagesListWafPackagesResponseDZw70ubTAPIResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                             `json:"total_count"`
	JSON       zoneFirewallWafPackageWafPackagesListWafPackagesResponseDZw70ubTapiResponseCollectionResultInfoJSON `json:"-"`
}

// zoneFirewallWafPackageWafPackagesListWafPackagesResponseDZw70ubTapiResponseCollectionResultInfoJSON
// contains the JSON metadata for the struct
// [ZoneFirewallWafPackageWafPackagesListWafPackagesResponseDZw70ubTAPIResponseCollectionResultInfo]
type zoneFirewallWafPackageWafPackagesListWafPackagesResponseDZw70ubTapiResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageWafPackagesListWafPackagesResponseDZw70ubTAPIResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneFirewallWafPackageWafPackagesListWafPackagesResponseDZw70ubTAPIResponseCollectionSuccess bool

const (
	ZoneFirewallWafPackageWafPackagesListWafPackagesResponseDZw70ubTAPIResponseCollectionSuccessTrue ZoneFirewallWafPackageWafPackagesListWafPackagesResponseDZw70ubTAPIResponseCollectionSuccess = true
)

type ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObject struct {
	Result []ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResult `json:"result"`
	JSON   zoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectJSON     `json:"-"`
}

// zoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectJSON contains the
// JSON metadata for the struct
// [ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObject]
type zoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObject) implementsZoneFirewallWafPackageWafPackagesListWafPackagesResponse() {
}

// Union satisfied by
// [ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTPackageDefinition]
// or
// [ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTAnomalyPackage].
type ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResult interface {
	implementsZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResult)(nil)).Elem(), "")
}

type ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTPackageDefinition struct {
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
	DetectionMode ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTPackageDefinitionDetectionMode `json:"detection_mode,required"`
	// The name of the WAF package.
	Name string `json:"name,required"`
	// Identifier
	ZoneID string `json:"zone_id,required"`
	// When set to `active`, indicates that the WAF package will be applied to the
	// zone.
	Status ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTPackageDefinitionStatus `json:"status"`
	JSON   zoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTPackageDefinitionJSON   `json:"-"`
}

// zoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTPackageDefinitionJSON
// contains the JSON metadata for the struct
// [ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTPackageDefinition]
type zoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTPackageDefinitionJSON struct {
	ID            apijson.Field
	Description   apijson.Field
	DetectionMode apijson.Field
	Name          apijson.Field
	ZoneID        apijson.Field
	Status        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTPackageDefinition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTPackageDefinition) implementsZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResult() {
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
type ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTPackageDefinitionDetectionMode string

const (
	ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTPackageDefinitionDetectionModeAnomaly     ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTPackageDefinitionDetectionMode = "anomaly"
	ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTPackageDefinitionDetectionModeTraditional ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTPackageDefinitionDetectionMode = "traditional"
)

// When set to `active`, indicates that the WAF package will be applied to the
// zone.
type ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTPackageDefinitionStatus string

const (
	ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTPackageDefinitionStatusActive ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTPackageDefinitionStatus = "active"
)

type ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTAnomalyPackage struct {
	// The unique identifier of a WAF package.
	ID string `json:"id,required"`
	// The default action performed by the rules in the WAF package.
	ActionMode ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTAnomalyPackageActionMode `json:"action_mode,required"`
	// A summary of the purpose/function of the WAF package.
	Description string `json:"description,required"`
	// When a WAF package uses anomaly detection, each rule is given a score when
	// triggered. If the total score of all triggered rules exceeds the sensitivity
	// defined on the WAF package, the action defined on the package will be taken.
	DetectionMode string `json:"detection_mode,required"`
	// The name of the WAF package.
	Name string `json:"name,required"`
	// The sensitivity of the WAF package.
	Sensitivity ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTAnomalyPackageSensitivity `json:"sensitivity,required"`
	// Identifier
	ZoneID string `json:"zone_id,required"`
	// When set to `active`, indicates that the WAF package will be applied to the
	// zone.
	Status ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTAnomalyPackageStatus `json:"status"`
	JSON   zoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTAnomalyPackageJSON   `json:"-"`
}

// zoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTAnomalyPackageJSON
// contains the JSON metadata for the struct
// [ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTAnomalyPackage]
type zoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTAnomalyPackageJSON struct {
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

func (r *ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTAnomalyPackage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTAnomalyPackage) implementsZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResult() {
}

// The default action performed by the rules in the WAF package.
type ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTAnomalyPackageActionMode string

const (
	ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTAnomalyPackageActionModeSimulate  ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTAnomalyPackageActionMode = "simulate"
	ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTAnomalyPackageActionModeBlock     ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTAnomalyPackageActionMode = "block"
	ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTAnomalyPackageActionModeChallenge ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTAnomalyPackageActionMode = "challenge"
)

// The sensitivity of the WAF package.
type ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTAnomalyPackageSensitivity string

const (
	ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTAnomalyPackageSensitivityHigh   ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTAnomalyPackageSensitivity = "high"
	ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTAnomalyPackageSensitivityMedium ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTAnomalyPackageSensitivity = "medium"
	ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTAnomalyPackageSensitivityLow    ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTAnomalyPackageSensitivity = "low"
	ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTAnomalyPackageSensitivityOff    ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTAnomalyPackageSensitivity = "off"
)

// When set to `active`, indicates that the WAF package will be applied to the
// zone.
type ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTAnomalyPackageStatus string

const (
	ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTAnomalyPackageStatusActive ZoneFirewallWafPackageWafPackagesListWafPackagesResponseObjectResultDZw70ubTAnomalyPackageStatus = "active"
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
