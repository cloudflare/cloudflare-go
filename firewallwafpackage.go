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

// FirewallWAFPackageService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewFirewallWAFPackageService] method
// instead.
type FirewallWAFPackageService struct {
	Options []option.RequestOption
	Groups  *FirewallWAFPackageGroupService
	Rules   *FirewallWAFPackageRuleService
}

// NewFirewallWAFPackageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewFirewallWAFPackageService(opts ...option.RequestOption) (r *FirewallWAFPackageService) {
	r = &FirewallWAFPackageService{}
	r.Options = opts
	r.Groups = NewFirewallWAFPackageGroupService(opts...)
	r.Rules = NewFirewallWAFPackageRuleService(opts...)
	return
}

// Fetches WAF packages for a zone.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *FirewallWAFPackageService) List(ctx context.Context, zoneIdentifier string, query FirewallWAFPackageListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[FirewallWAFPackageListResponse], err error) {
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
func (r *FirewallWAFPackageService) ListAutoPaging(ctx context.Context, zoneIdentifier string, query FirewallWAFPackageListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[FirewallWAFPackageListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, zoneIdentifier, query, opts...))
}

// Fetches the details of a WAF package.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *FirewallWAFPackageService) Get(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *FirewallWAFPackageGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Union satisfied by
// [FirewallWAFPackageListResponseLegacyJhsAPIResponseCollection] or
// [FirewallWAFPackageListResponseObject].
type FirewallWAFPackageListResponse interface {
	implementsFirewallWAFPackageListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*FirewallWAFPackageListResponse)(nil)).Elem(), "")
}

type FirewallWAFPackageListResponseLegacyJhsAPIResponseCollection struct {
	Errors   []FirewallWAFPackageListResponseLegacyJhsAPIResponseCollectionError   `json:"errors,required"`
	Messages []FirewallWAFPackageListResponseLegacyJhsAPIResponseCollectionMessage `json:"messages,required"`
	Result   FirewallWAFPackageListResponseLegacyJhsAPIResponseCollectionResult    `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    FirewallWAFPackageListResponseLegacyJhsAPIResponseCollectionSuccess    `json:"success,required"`
	ResultInfo FirewallWAFPackageListResponseLegacyJhsAPIResponseCollectionResultInfo `json:"result_info"`
	JSON       firewallWAFPackageListResponseLegacyJhsAPIResponseCollectionJSON       `json:"-"`
}

// firewallWAFPackageListResponseLegacyJhsAPIResponseCollectionJSON contains the
// JSON metadata for the struct
// [FirewallWAFPackageListResponseLegacyJhsAPIResponseCollection]
type firewallWAFPackageListResponseLegacyJhsAPIResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageListResponseLegacyJhsAPIResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallWAFPackageListResponseLegacyJhsAPIResponseCollectionJSON) RawJSON() string {
	return r.raw
}

func (r FirewallWAFPackageListResponseLegacyJhsAPIResponseCollection) implementsFirewallWAFPackageListResponse() {
}

type FirewallWAFPackageListResponseLegacyJhsAPIResponseCollectionError struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    firewallWAFPackageListResponseLegacyJhsAPIResponseCollectionErrorJSON `json:"-"`
}

// firewallWAFPackageListResponseLegacyJhsAPIResponseCollectionErrorJSON contains
// the JSON metadata for the struct
// [FirewallWAFPackageListResponseLegacyJhsAPIResponseCollectionError]
type firewallWAFPackageListResponseLegacyJhsAPIResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageListResponseLegacyJhsAPIResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallWAFPackageListResponseLegacyJhsAPIResponseCollectionErrorJSON) RawJSON() string {
	return r.raw
}

type FirewallWAFPackageListResponseLegacyJhsAPIResponseCollectionMessage struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    firewallWAFPackageListResponseLegacyJhsAPIResponseCollectionMessageJSON `json:"-"`
}

// firewallWAFPackageListResponseLegacyJhsAPIResponseCollectionMessageJSON contains
// the JSON metadata for the struct
// [FirewallWAFPackageListResponseLegacyJhsAPIResponseCollectionMessage]
type firewallWAFPackageListResponseLegacyJhsAPIResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageListResponseLegacyJhsAPIResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallWAFPackageListResponseLegacyJhsAPIResponseCollectionMessageJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [FirewallWAFPackageListResponseLegacyJhsAPIResponseCollectionResultUnknown],
// [FirewallWAFPackageListResponseLegacyJhsAPIResponseCollectionResultArray] or
// [shared.UnionString].
type FirewallWAFPackageListResponseLegacyJhsAPIResponseCollectionResult interface {
	ImplementsFirewallWAFPackageListResponseLegacyJhsAPIResponseCollectionResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallWAFPackageListResponseLegacyJhsAPIResponseCollectionResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type FirewallWAFPackageListResponseLegacyJhsAPIResponseCollectionResultArray []interface{}

func (r FirewallWAFPackageListResponseLegacyJhsAPIResponseCollectionResultArray) ImplementsFirewallWAFPackageListResponseLegacyJhsAPIResponseCollectionResult() {
}

// Whether the API call was successful
type FirewallWAFPackageListResponseLegacyJhsAPIResponseCollectionSuccess bool

const (
	FirewallWAFPackageListResponseLegacyJhsAPIResponseCollectionSuccessTrue FirewallWAFPackageListResponseLegacyJhsAPIResponseCollectionSuccess = true
)

type FirewallWAFPackageListResponseLegacyJhsAPIResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                    `json:"total_count"`
	JSON       firewallWAFPackageListResponseLegacyJhsAPIResponseCollectionResultInfoJSON `json:"-"`
}

// firewallWAFPackageListResponseLegacyJhsAPIResponseCollectionResultInfoJSON
// contains the JSON metadata for the struct
// [FirewallWAFPackageListResponseLegacyJhsAPIResponseCollectionResultInfo]
type firewallWAFPackageListResponseLegacyJhsAPIResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageListResponseLegacyJhsAPIResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallWAFPackageListResponseLegacyJhsAPIResponseCollectionResultInfoJSON) RawJSON() string {
	return r.raw
}

type FirewallWAFPackageListResponseObject struct {
	Result []FirewallWAFPackageListResponseObjectResult `json:"result"`
	JSON   firewallWAFPackageListResponseObjectJSON     `json:"-"`
}

// firewallWAFPackageListResponseObjectJSON contains the JSON metadata for the
// struct [FirewallWAFPackageListResponseObject]
type firewallWAFPackageListResponseObjectJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageListResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallWAFPackageListResponseObjectJSON) RawJSON() string {
	return r.raw
}

func (r FirewallWAFPackageListResponseObject) implementsFirewallWAFPackageListResponse() {}

// Union satisfied by
// [FirewallWAFPackageListResponseObjectResultLegacyJhsPackageDefinition] or
// [FirewallWAFPackageListResponseObjectResultLegacyJhsAnomalyPackage].
type FirewallWAFPackageListResponseObjectResult interface {
	implementsFirewallWAFPackageListResponseObjectResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*FirewallWAFPackageListResponseObjectResult)(nil)).Elem(), "")
}

type FirewallWAFPackageListResponseObjectResultLegacyJhsPackageDefinition struct {
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
	DetectionMode FirewallWAFPackageListResponseObjectResultLegacyJhsPackageDefinitionDetectionMode `json:"detection_mode,required"`
	// The name of the WAF package.
	Name string `json:"name,required"`
	// Identifier
	ZoneID string `json:"zone_id,required"`
	// When set to `active`, indicates that the WAF package will be applied to the
	// zone.
	Status FirewallWAFPackageListResponseObjectResultLegacyJhsPackageDefinitionStatus `json:"status"`
	JSON   firewallWAFPackageListResponseObjectResultLegacyJhsPackageDefinitionJSON   `json:"-"`
}

// firewallWAFPackageListResponseObjectResultLegacyJhsPackageDefinitionJSON
// contains the JSON metadata for the struct
// [FirewallWAFPackageListResponseObjectResultLegacyJhsPackageDefinition]
type firewallWAFPackageListResponseObjectResultLegacyJhsPackageDefinitionJSON struct {
	ID            apijson.Field
	Description   apijson.Field
	DetectionMode apijson.Field
	Name          apijson.Field
	ZoneID        apijson.Field
	Status        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *FirewallWAFPackageListResponseObjectResultLegacyJhsPackageDefinition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallWAFPackageListResponseObjectResultLegacyJhsPackageDefinitionJSON) RawJSON() string {
	return r.raw
}

func (r FirewallWAFPackageListResponseObjectResultLegacyJhsPackageDefinition) implementsFirewallWAFPackageListResponseObjectResult() {
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
type FirewallWAFPackageListResponseObjectResultLegacyJhsPackageDefinitionDetectionMode string

const (
	FirewallWAFPackageListResponseObjectResultLegacyJhsPackageDefinitionDetectionModeAnomaly     FirewallWAFPackageListResponseObjectResultLegacyJhsPackageDefinitionDetectionMode = "anomaly"
	FirewallWAFPackageListResponseObjectResultLegacyJhsPackageDefinitionDetectionModeTraditional FirewallWAFPackageListResponseObjectResultLegacyJhsPackageDefinitionDetectionMode = "traditional"
)

// When set to `active`, indicates that the WAF package will be applied to the
// zone.
type FirewallWAFPackageListResponseObjectResultLegacyJhsPackageDefinitionStatus string

const (
	FirewallWAFPackageListResponseObjectResultLegacyJhsPackageDefinitionStatusActive FirewallWAFPackageListResponseObjectResultLegacyJhsPackageDefinitionStatus = "active"
)

type FirewallWAFPackageListResponseObjectResultLegacyJhsAnomalyPackage struct {
	// The unique identifier of a WAF package.
	ID string `json:"id,required"`
	// A summary of the purpose/function of the WAF package.
	Description string `json:"description,required"`
	// When a WAF package uses anomaly detection, each rule is given a score when
	// triggered. If the total score of all triggered rules exceeds the sensitivity
	// defined on the WAF package, the action defined on the package will be taken.
	DetectionMode FirewallWAFPackageListResponseObjectResultLegacyJhsAnomalyPackageDetectionMode `json:"detection_mode,required"`
	// The name of the WAF package.
	Name string `json:"name,required"`
	// Identifier
	ZoneID string `json:"zone_id,required"`
	// The default action performed by the rules in the WAF package.
	ActionMode FirewallWAFPackageListResponseObjectResultLegacyJhsAnomalyPackageActionMode `json:"action_mode"`
	// The sensitivity of the WAF package.
	Sensitivity FirewallWAFPackageListResponseObjectResultLegacyJhsAnomalyPackageSensitivity `json:"sensitivity"`
	// When set to `active`, indicates that the WAF package will be applied to the
	// zone.
	Status FirewallWAFPackageListResponseObjectResultLegacyJhsAnomalyPackageStatus `json:"status"`
	JSON   firewallWAFPackageListResponseObjectResultLegacyJhsAnomalyPackageJSON   `json:"-"`
}

// firewallWAFPackageListResponseObjectResultLegacyJhsAnomalyPackageJSON contains
// the JSON metadata for the struct
// [FirewallWAFPackageListResponseObjectResultLegacyJhsAnomalyPackage]
type firewallWAFPackageListResponseObjectResultLegacyJhsAnomalyPackageJSON struct {
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

func (r *FirewallWAFPackageListResponseObjectResultLegacyJhsAnomalyPackage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallWAFPackageListResponseObjectResultLegacyJhsAnomalyPackageJSON) RawJSON() string {
	return r.raw
}

func (r FirewallWAFPackageListResponseObjectResultLegacyJhsAnomalyPackage) implementsFirewallWAFPackageListResponseObjectResult() {
}

// When a WAF package uses anomaly detection, each rule is given a score when
// triggered. If the total score of all triggered rules exceeds the sensitivity
// defined on the WAF package, the action defined on the package will be taken.
type FirewallWAFPackageListResponseObjectResultLegacyJhsAnomalyPackageDetectionMode string

const (
	FirewallWAFPackageListResponseObjectResultLegacyJhsAnomalyPackageDetectionModeAnomaly     FirewallWAFPackageListResponseObjectResultLegacyJhsAnomalyPackageDetectionMode = "anomaly"
	FirewallWAFPackageListResponseObjectResultLegacyJhsAnomalyPackageDetectionModeTraditional FirewallWAFPackageListResponseObjectResultLegacyJhsAnomalyPackageDetectionMode = "traditional"
)

// The default action performed by the rules in the WAF package.
type FirewallWAFPackageListResponseObjectResultLegacyJhsAnomalyPackageActionMode string

const (
	FirewallWAFPackageListResponseObjectResultLegacyJhsAnomalyPackageActionModeSimulate  FirewallWAFPackageListResponseObjectResultLegacyJhsAnomalyPackageActionMode = "simulate"
	FirewallWAFPackageListResponseObjectResultLegacyJhsAnomalyPackageActionModeBlock     FirewallWAFPackageListResponseObjectResultLegacyJhsAnomalyPackageActionMode = "block"
	FirewallWAFPackageListResponseObjectResultLegacyJhsAnomalyPackageActionModeChallenge FirewallWAFPackageListResponseObjectResultLegacyJhsAnomalyPackageActionMode = "challenge"
)

// The sensitivity of the WAF package.
type FirewallWAFPackageListResponseObjectResultLegacyJhsAnomalyPackageSensitivity string

const (
	FirewallWAFPackageListResponseObjectResultLegacyJhsAnomalyPackageSensitivityHigh   FirewallWAFPackageListResponseObjectResultLegacyJhsAnomalyPackageSensitivity = "high"
	FirewallWAFPackageListResponseObjectResultLegacyJhsAnomalyPackageSensitivityMedium FirewallWAFPackageListResponseObjectResultLegacyJhsAnomalyPackageSensitivity = "medium"
	FirewallWAFPackageListResponseObjectResultLegacyJhsAnomalyPackageSensitivityLow    FirewallWAFPackageListResponseObjectResultLegacyJhsAnomalyPackageSensitivity = "low"
	FirewallWAFPackageListResponseObjectResultLegacyJhsAnomalyPackageSensitivityOff    FirewallWAFPackageListResponseObjectResultLegacyJhsAnomalyPackageSensitivity = "off"
)

// When set to `active`, indicates that the WAF package will be applied to the
// zone.
type FirewallWAFPackageListResponseObjectResultLegacyJhsAnomalyPackageStatus string

const (
	FirewallWAFPackageListResponseObjectResultLegacyJhsAnomalyPackageStatusActive FirewallWAFPackageListResponseObjectResultLegacyJhsAnomalyPackageStatus = "active"
)

// Union satisfied by [FirewallWAFPackageGetResponseLegacyJhsAPIResponseSingle] or
// [FirewallWAFPackageGetResponseObject].
type FirewallWAFPackageGetResponse interface {
	implementsFirewallWAFPackageGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*FirewallWAFPackageGetResponse)(nil)).Elem(), "")
}

type FirewallWAFPackageGetResponseLegacyJhsAPIResponseSingle struct {
	Errors   []FirewallWAFPackageGetResponseLegacyJhsAPIResponseSingleError   `json:"errors,required"`
	Messages []FirewallWAFPackageGetResponseLegacyJhsAPIResponseSingleMessage `json:"messages,required"`
	Result   FirewallWAFPackageGetResponseLegacyJhsAPIResponseSingleResult    `json:"result,required,nullable"`
	// Whether the API call was successful
	Success FirewallWAFPackageGetResponseLegacyJhsAPIResponseSingleSuccess `json:"success,required"`
	JSON    firewallWAFPackageGetResponseLegacyJhsAPIResponseSingleJSON    `json:"-"`
}

// firewallWAFPackageGetResponseLegacyJhsAPIResponseSingleJSON contains the JSON
// metadata for the struct
// [FirewallWAFPackageGetResponseLegacyJhsAPIResponseSingle]
type firewallWAFPackageGetResponseLegacyJhsAPIResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageGetResponseLegacyJhsAPIResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallWAFPackageGetResponseLegacyJhsAPIResponseSingleJSON) RawJSON() string {
	return r.raw
}

func (r FirewallWAFPackageGetResponseLegacyJhsAPIResponseSingle) implementsFirewallWAFPackageGetResponse() {
}

type FirewallWAFPackageGetResponseLegacyJhsAPIResponseSingleError struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    firewallWAFPackageGetResponseLegacyJhsAPIResponseSingleErrorJSON `json:"-"`
}

// firewallWAFPackageGetResponseLegacyJhsAPIResponseSingleErrorJSON contains the
// JSON metadata for the struct
// [FirewallWAFPackageGetResponseLegacyJhsAPIResponseSingleError]
type firewallWAFPackageGetResponseLegacyJhsAPIResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageGetResponseLegacyJhsAPIResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallWAFPackageGetResponseLegacyJhsAPIResponseSingleErrorJSON) RawJSON() string {
	return r.raw
}

type FirewallWAFPackageGetResponseLegacyJhsAPIResponseSingleMessage struct {
	Code    int64                                                              `json:"code,required"`
	Message string                                                             `json:"message,required"`
	JSON    firewallWAFPackageGetResponseLegacyJhsAPIResponseSingleMessageJSON `json:"-"`
}

// firewallWAFPackageGetResponseLegacyJhsAPIResponseSingleMessageJSON contains the
// JSON metadata for the struct
// [FirewallWAFPackageGetResponseLegacyJhsAPIResponseSingleMessage]
type firewallWAFPackageGetResponseLegacyJhsAPIResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageGetResponseLegacyJhsAPIResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallWAFPackageGetResponseLegacyJhsAPIResponseSingleMessageJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [FirewallWAFPackageGetResponseLegacyJhsAPIResponseSingleResultUnknown] or
// [shared.UnionString].
type FirewallWAFPackageGetResponseLegacyJhsAPIResponseSingleResult interface {
	ImplementsFirewallWAFPackageGetResponseLegacyJhsAPIResponseSingleResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallWAFPackageGetResponseLegacyJhsAPIResponseSingleResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Whether the API call was successful
type FirewallWAFPackageGetResponseLegacyJhsAPIResponseSingleSuccess bool

const (
	FirewallWAFPackageGetResponseLegacyJhsAPIResponseSingleSuccessTrue FirewallWAFPackageGetResponseLegacyJhsAPIResponseSingleSuccess = true
)

type FirewallWAFPackageGetResponseObject struct {
	Result interface{}                             `json:"result"`
	JSON   firewallWAFPackageGetResponseObjectJSON `json:"-"`
}

// firewallWAFPackageGetResponseObjectJSON contains the JSON metadata for the
// struct [FirewallWAFPackageGetResponseObject]
type firewallWAFPackageGetResponseObjectJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageGetResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallWAFPackageGetResponseObjectJSON) RawJSON() string {
	return r.raw
}

func (r FirewallWAFPackageGetResponseObject) implementsFirewallWAFPackageGetResponse() {}

type FirewallWAFPackageListParams struct {
	// The direction used to sort returned packages.
	Direction param.Field[FirewallWAFPackageListParamsDirection] `query:"direction"`
	// When set to `all`, all the search requirements must match. When set to `any`,
	// only one of the search requirements has to match.
	Match param.Field[FirewallWAFPackageListParamsMatch] `query:"match"`
	// The field used to sort returned packages.
	Order param.Field[FirewallWAFPackageListParamsOrder] `query:"order"`
	// The page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The number of packages per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [FirewallWAFPackageListParams]'s query parameters as
// `url.Values`.
func (r FirewallWAFPackageListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The direction used to sort returned packages.
type FirewallWAFPackageListParamsDirection string

const (
	FirewallWAFPackageListParamsDirectionAsc  FirewallWAFPackageListParamsDirection = "asc"
	FirewallWAFPackageListParamsDirectionDesc FirewallWAFPackageListParamsDirection = "desc"
)

// When set to `all`, all the search requirements must match. When set to `any`,
// only one of the search requirements has to match.
type FirewallWAFPackageListParamsMatch string

const (
	FirewallWAFPackageListParamsMatchAny FirewallWAFPackageListParamsMatch = "any"
	FirewallWAFPackageListParamsMatchAll FirewallWAFPackageListParamsMatch = "all"
)

// The field used to sort returned packages.
type FirewallWAFPackageListParamsOrder string

const (
	FirewallWAFPackageListParamsOrderName FirewallWAFPackageListParamsOrder = "name"
)
