// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// ZarazConfigService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZarazConfigService] method
// instead.
type ZarazConfigService struct {
	Options []option.RequestOption
}

// NewZarazConfigService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZarazConfigService(opts ...option.RequestOption) (r *ZarazConfigService) {
	r = &ZarazConfigService{}
	r.Options = opts
	return
}

// Gets latest Zaraz configuration for a zone. It can be preview or published
// configuration, whichever was the last updated. Secret variables values will not
// be included.
func (r *ZarazConfigService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZarazConfigGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZarazConfigGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/zaraz/config", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates Zaraz configuration for a zone.
func (r *ZarazConfigService) Replace(ctx context.Context, zoneID string, body ZarazConfigReplaceParams, opts ...option.RequestOption) (res *ZarazConfigReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZarazConfigReplaceResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/zaraz/config", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Zaraz configuration
type ZarazConfigGetResponse struct {
	// Data layer compatibility mode enabled.
	DataLayer bool `json:"dataLayer,required"`
	// The key for Zaraz debug mode.
	DebugKey string `json:"debugKey,required"`
	// General Zaraz settings.
	Settings ZarazConfigGetResponseSettings `json:"settings,required"`
	// Tools set up under Zaraz configuration, where key is the alpha-numeric tool ID
	// and value is the tool configuration object.
	Tools map[string]ZarazConfigGetResponseTool `json:"tools,required"`
	// Triggers set up under Zaraz configuration, where key is the trigger
	// alpha-numeric ID and value is the trigger configuration.
	Triggers map[string]ZarazConfigGetResponseTrigger `json:"triggers,required"`
	// Variables set up under Zaraz configuration, where key is the variable
	// alpha-numeric ID and value is the variable configuration. Values of variables of
	// type secret are not included.
	Variables map[string]ZarazConfigGetResponseVariable `json:"variables,required"`
	// Zaraz internal version of the config.
	ZarazVersion int64 `json:"zarazVersion,required"`
	// Consent management configuration.
	Consent ZarazConfigGetResponseConsent `json:"consent"`
	// Single Page Application support enabled.
	HistoryChange bool                       `json:"historyChange"`
	JSON          zarazConfigGetResponseJSON `json:"-"`
}

// zarazConfigGetResponseJSON contains the JSON metadata for the struct
// [ZarazConfigGetResponse]
type zarazConfigGetResponseJSON struct {
	DataLayer     apijson.Field
	DebugKey      apijson.Field
	Settings      apijson.Field
	Tools         apijson.Field
	Triggers      apijson.Field
	Variables     apijson.Field
	ZarazVersion  apijson.Field
	Consent       apijson.Field
	HistoryChange apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZarazConfigGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// General Zaraz settings.
type ZarazConfigGetResponseSettings struct {
	// Automatic injection of Zaraz scripts enabled.
	AutoInjectScript bool `json:"autoInjectScript,required"`
	// Details of the worker that receives and edits Zaraz Context object.
	ContextEnricher ZarazConfigGetResponseSettingsContextEnricher `json:"contextEnricher"`
	// The domain Zaraz will use for writing and reading its cookies.
	CookieDomain string `json:"cookieDomain"`
	// Ecommerce API enabled.
	Ecommerce bool `json:"ecommerce"`
	// Custom endpoint for server-side track events.
	EventsAPIPath string `json:"eventsApiPath"`
	// Hiding external referrer URL enabled.
	HideExternalReferer bool `json:"hideExternalReferer"`
	// Trimming IP address enabled.
	HideIPAddress bool `json:"hideIPAddress"`
	// Removing URL query params enabled.
	HideQueryParams bool `json:"hideQueryParams"`
	// Removing sensitive data from User Aagent string enabled.
	HideUserAgent bool `json:"hideUserAgent"`
	// Custom endpoint for Zaraz init script.
	InitPath string `json:"initPath"`
	// Injection of Zaraz scripts into iframes enabled.
	InjectIframes bool `json:"injectIframes"`
	// Custom path for Managed Components server functionalities.
	McRootPath string `json:"mcRootPath"`
	// Custom endpoint for Zaraz main script.
	ScriptPath string `json:"scriptPath"`
	// Custom endpoint for Zaraz tracking requests.
	TrackPath string                             `json:"trackPath"`
	JSON      zarazConfigGetResponseSettingsJSON `json:"-"`
}

// zarazConfigGetResponseSettingsJSON contains the JSON metadata for the struct
// [ZarazConfigGetResponseSettings]
type zarazConfigGetResponseSettingsJSON struct {
	AutoInjectScript    apijson.Field
	ContextEnricher     apijson.Field
	CookieDomain        apijson.Field
	Ecommerce           apijson.Field
	EventsAPIPath       apijson.Field
	HideExternalReferer apijson.Field
	HideIPAddress       apijson.Field
	HideQueryParams     apijson.Field
	HideUserAgent       apijson.Field
	InitPath            apijson.Field
	InjectIframes       apijson.Field
	McRootPath          apijson.Field
	ScriptPath          apijson.Field
	TrackPath           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ZarazConfigGetResponseSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details of the worker that receives and edits Zaraz Context object.
type ZarazConfigGetResponseSettingsContextEnricher struct {
	EscapedWorkerName string                                            `json:"escapedWorkerName,required"`
	WorkerTag         string                                            `json:"workerTag,required"`
	JSON              zarazConfigGetResponseSettingsContextEnricherJSON `json:"-"`
}

// zarazConfigGetResponseSettingsContextEnricherJSON contains the JSON metadata for
// the struct [ZarazConfigGetResponseSettingsContextEnricher]
type zarazConfigGetResponseSettingsContextEnricherJSON struct {
	EscapedWorkerName apijson.Field
	WorkerTag         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZarazConfigGetResponseSettingsContextEnricher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [ZarazConfigGetResponseToolsZarazManagedComponent] or
// [ZarazConfigGetResponseToolsZarazCustomManagedComponent].
type ZarazConfigGetResponseTool interface {
	implementsZarazConfigGetResponseTool()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZarazConfigGetResponseTool)(nil)).Elem(), "")
}

type ZarazConfigGetResponseToolsZarazManagedComponent struct {
	// List of blocking trigger IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Tool's internal name
	Component string `json:"component,required"`
	// Default fields for tool's actions
	DefaultFields map[string]ZarazConfigGetResponseToolsZarazManagedComponentDefaultField `json:"defaultFields,required"`
	// Whether tool is enabled
	Enabled bool `json:"enabled,required"`
	// Tool's name defined by the user
	Name string `json:"name,required"`
	// List of permissions granted to the component
	Permissions []string `json:"permissions,required"`
	// Tool's settings
	Settings map[string]ZarazConfigGetResponseToolsZarazManagedComponentSetting `json:"settings,required"`
	Type     ZarazConfigGetResponseToolsZarazManagedComponentType               `json:"type,required"`
	// Actions configured on a tool. Either this or neoEvents field is required.
	Actions map[string]ZarazConfigGetResponseToolsZarazManagedComponentAction `json:"actions"`
	// Default consent purpose ID
	DefaultPurpose string `json:"defaultPurpose"`
	// DEPRECATED - List of actions configured on a tool. Either this or actions field
	// is required. If both are present, actions field will take precedence.
	NeoEvents []ZarazConfigGetResponseToolsZarazManagedComponentNeoEvent `json:"neoEvents"`
	JSON      zarazConfigGetResponseToolsZarazManagedComponentJSON       `json:"-"`
}

// zarazConfigGetResponseToolsZarazManagedComponentJSON contains the JSON metadata
// for the struct [ZarazConfigGetResponseToolsZarazManagedComponent]
type zarazConfigGetResponseToolsZarazManagedComponentJSON struct {
	BlockingTriggers apijson.Field
	Component        apijson.Field
	DefaultFields    apijson.Field
	Enabled          apijson.Field
	Name             apijson.Field
	Permissions      apijson.Field
	Settings         apijson.Field
	Type             apijson.Field
	Actions          apijson.Field
	DefaultPurpose   apijson.Field
	NeoEvents        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazConfigGetResponseToolsZarazManagedComponent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigGetResponseToolsZarazManagedComponent) implementsZarazConfigGetResponseTool() {}

// Union satisfied by [shared.UnionString] or [shared.UnionBool].
type ZarazConfigGetResponseToolsZarazManagedComponentDefaultField interface {
	ImplementsZarazConfigGetResponseToolsZarazManagedComponentDefaultField()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazConfigGetResponseToolsZarazManagedComponentDefaultField)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

// Union satisfied by [shared.UnionString] or [shared.UnionBool].
type ZarazConfigGetResponseToolsZarazManagedComponentSetting interface {
	ImplementsZarazConfigGetResponseToolsZarazManagedComponentSetting()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazConfigGetResponseToolsZarazManagedComponentSetting)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

type ZarazConfigGetResponseToolsZarazManagedComponentType string

const (
	ZarazConfigGetResponseToolsZarazManagedComponentTypeComponent ZarazConfigGetResponseToolsZarazManagedComponentType = "component"
)

type ZarazConfigGetResponseToolsZarazManagedComponentAction struct {
	// Tool event type
	ActionType string `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Event payload
	Data interface{} `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers []string                                                   `json:"firingTriggers,required"`
	JSON           zarazConfigGetResponseToolsZarazManagedComponentActionJSON `json:"-"`
}

// zarazConfigGetResponseToolsZarazManagedComponentActionJSON contains the JSON
// metadata for the struct [ZarazConfigGetResponseToolsZarazManagedComponentAction]
type zarazConfigGetResponseToolsZarazManagedComponentActionJSON struct {
	ActionType       apijson.Field
	BlockingTriggers apijson.Field
	Data             apijson.Field
	FiringTriggers   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazConfigGetResponseToolsZarazManagedComponentAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigGetResponseToolsZarazManagedComponentNeoEvent struct {
	// Tool event type
	ActionType string `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Event payload
	Data interface{} `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers []string                                                     `json:"firingTriggers,required"`
	JSON           zarazConfigGetResponseToolsZarazManagedComponentNeoEventJSON `json:"-"`
}

// zarazConfigGetResponseToolsZarazManagedComponentNeoEventJSON contains the JSON
// metadata for the struct
// [ZarazConfigGetResponseToolsZarazManagedComponentNeoEvent]
type zarazConfigGetResponseToolsZarazManagedComponentNeoEventJSON struct {
	ActionType       apijson.Field
	BlockingTriggers apijson.Field
	Data             apijson.Field
	FiringTriggers   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazConfigGetResponseToolsZarazManagedComponentNeoEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigGetResponseToolsZarazCustomManagedComponent struct {
	// List of blocking trigger IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Tool's internal name
	Component string `json:"component,required"`
	// Default fields for tool's actions
	DefaultFields map[string]ZarazConfigGetResponseToolsZarazCustomManagedComponentDefaultField `json:"defaultFields,required"`
	// Whether tool is enabled
	Enabled bool `json:"enabled,required"`
	// Tool's name defined by the user
	Name string `json:"name,required"`
	// List of permissions granted to the component
	Permissions []string `json:"permissions,required"`
	// Tool's settings
	Settings map[string]ZarazConfigGetResponseToolsZarazCustomManagedComponentSetting `json:"settings,required"`
	Type     ZarazConfigGetResponseToolsZarazCustomManagedComponentType               `json:"type,required"`
	// Cloudflare worker that acts as a managed component
	Worker ZarazConfigGetResponseToolsZarazCustomManagedComponentWorker `json:"worker,required"`
	// Actions configured on a tool. Either this or neoEvents field is required.
	Actions map[string]ZarazConfigGetResponseToolsZarazCustomManagedComponentAction `json:"actions"`
	// Default consent purpose ID
	DefaultPurpose string `json:"defaultPurpose"`
	// DEPRECATED - List of actions configured on a tool. Either this or actions field
	// is required. If both are present, actions field will take precedence.
	NeoEvents []ZarazConfigGetResponseToolsZarazCustomManagedComponentNeoEvent `json:"neoEvents"`
	JSON      zarazConfigGetResponseToolsZarazCustomManagedComponentJSON       `json:"-"`
}

// zarazConfigGetResponseToolsZarazCustomManagedComponentJSON contains the JSON
// metadata for the struct [ZarazConfigGetResponseToolsZarazCustomManagedComponent]
type zarazConfigGetResponseToolsZarazCustomManagedComponentJSON struct {
	BlockingTriggers apijson.Field
	Component        apijson.Field
	DefaultFields    apijson.Field
	Enabled          apijson.Field
	Name             apijson.Field
	Permissions      apijson.Field
	Settings         apijson.Field
	Type             apijson.Field
	Worker           apijson.Field
	Actions          apijson.Field
	DefaultPurpose   apijson.Field
	NeoEvents        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazConfigGetResponseToolsZarazCustomManagedComponent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigGetResponseToolsZarazCustomManagedComponent) implementsZarazConfigGetResponseTool() {
}

// Union satisfied by [shared.UnionString] or [shared.UnionBool].
type ZarazConfigGetResponseToolsZarazCustomManagedComponentDefaultField interface {
	ImplementsZarazConfigGetResponseToolsZarazCustomManagedComponentDefaultField()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazConfigGetResponseToolsZarazCustomManagedComponentDefaultField)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

// Union satisfied by [shared.UnionString] or [shared.UnionBool].
type ZarazConfigGetResponseToolsZarazCustomManagedComponentSetting interface {
	ImplementsZarazConfigGetResponseToolsZarazCustomManagedComponentSetting()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazConfigGetResponseToolsZarazCustomManagedComponentSetting)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

type ZarazConfigGetResponseToolsZarazCustomManagedComponentType string

const (
	ZarazConfigGetResponseToolsZarazCustomManagedComponentTypeCustomMc ZarazConfigGetResponseToolsZarazCustomManagedComponentType = "custom-mc"
)

// Cloudflare worker that acts as a managed component
type ZarazConfigGetResponseToolsZarazCustomManagedComponentWorker struct {
	EscapedWorkerName string                                                           `json:"escapedWorkerName,required"`
	WorkerTag         string                                                           `json:"workerTag,required"`
	JSON              zarazConfigGetResponseToolsZarazCustomManagedComponentWorkerJSON `json:"-"`
}

// zarazConfigGetResponseToolsZarazCustomManagedComponentWorkerJSON contains the
// JSON metadata for the struct
// [ZarazConfigGetResponseToolsZarazCustomManagedComponentWorker]
type zarazConfigGetResponseToolsZarazCustomManagedComponentWorkerJSON struct {
	EscapedWorkerName apijson.Field
	WorkerTag         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZarazConfigGetResponseToolsZarazCustomManagedComponentWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigGetResponseToolsZarazCustomManagedComponentAction struct {
	// Tool event type
	ActionType string `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Event payload
	Data interface{} `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers []string                                                         `json:"firingTriggers,required"`
	JSON           zarazConfigGetResponseToolsZarazCustomManagedComponentActionJSON `json:"-"`
}

// zarazConfigGetResponseToolsZarazCustomManagedComponentActionJSON contains the
// JSON metadata for the struct
// [ZarazConfigGetResponseToolsZarazCustomManagedComponentAction]
type zarazConfigGetResponseToolsZarazCustomManagedComponentActionJSON struct {
	ActionType       apijson.Field
	BlockingTriggers apijson.Field
	Data             apijson.Field
	FiringTriggers   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazConfigGetResponseToolsZarazCustomManagedComponentAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigGetResponseToolsZarazCustomManagedComponentNeoEvent struct {
	// Tool event type
	ActionType string `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Event payload
	Data interface{} `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers []string                                                           `json:"firingTriggers,required"`
	JSON           zarazConfigGetResponseToolsZarazCustomManagedComponentNeoEventJSON `json:"-"`
}

// zarazConfigGetResponseToolsZarazCustomManagedComponentNeoEventJSON contains the
// JSON metadata for the struct
// [ZarazConfigGetResponseToolsZarazCustomManagedComponentNeoEvent]
type zarazConfigGetResponseToolsZarazCustomManagedComponentNeoEventJSON struct {
	ActionType       apijson.Field
	BlockingTriggers apijson.Field
	Data             apijson.Field
	FiringTriggers   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazConfigGetResponseToolsZarazCustomManagedComponentNeoEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigGetResponseTrigger struct {
	// Rules defining when the trigger is not fired.
	ExcludeRules []ZarazConfigGetResponseTriggersExcludeRule `json:"excludeRules,required"`
	// Rules defining when the trigger is fired.
	LoadRules []ZarazConfigGetResponseTriggersLoadRule `json:"loadRules,required"`
	// Trigger name.
	Name string `json:"name,required"`
	// Trigger description.
	Description string                               `json:"description"`
	System      ZarazConfigGetResponseTriggersSystem `json:"system"`
	JSON        zarazConfigGetResponseTriggerJSON    `json:"-"`
}

// zarazConfigGetResponseTriggerJSON contains the JSON metadata for the struct
// [ZarazConfigGetResponseTrigger]
type zarazConfigGetResponseTriggerJSON struct {
	ExcludeRules apijson.Field
	LoadRules    apijson.Field
	Name         apijson.Field
	Description  apijson.Field
	System       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZarazConfigGetResponseTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [ZarazConfigGetResponseTriggersExcludeRulesZarazLoadRule],
// [ZarazConfigGetResponseTriggersExcludeRulesZarazClickListenerRule],
// [ZarazConfigGetResponseTriggersExcludeRulesZarazTimerRule],
// [ZarazConfigGetResponseTriggersExcludeRulesZarazFormSubmissionRule],
// [ZarazConfigGetResponseTriggersExcludeRulesZarazVariableMatchRule],
// [ZarazConfigGetResponseTriggersExcludeRulesZarazScrollDepthRule] or
// [ZarazConfigGetResponseTriggersExcludeRulesZarazElementVisibilityRule].
type ZarazConfigGetResponseTriggersExcludeRule interface {
	implementsZarazConfigGetResponseTriggersExcludeRule()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZarazConfigGetResponseTriggersExcludeRule)(nil)).Elem(), "")
}

type ZarazConfigGetResponseTriggersExcludeRulesZarazLoadRule struct {
	ID    string                                                      `json:"id,required"`
	Match string                                                      `json:"match,required"`
	Op    ZarazConfigGetResponseTriggersExcludeRulesZarazLoadRuleOp   `json:"op,required"`
	Value string                                                      `json:"value,required"`
	JSON  zarazConfigGetResponseTriggersExcludeRulesZarazLoadRuleJSON `json:"-"`
}

// zarazConfigGetResponseTriggersExcludeRulesZarazLoadRuleJSON contains the JSON
// metadata for the struct
// [ZarazConfigGetResponseTriggersExcludeRulesZarazLoadRule]
type zarazConfigGetResponseTriggersExcludeRulesZarazLoadRuleJSON struct {
	ID          apijson.Field
	Match       apijson.Field
	Op          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigGetResponseTriggersExcludeRulesZarazLoadRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigGetResponseTriggersExcludeRulesZarazLoadRule) implementsZarazConfigGetResponseTriggersExcludeRule() {
}

type ZarazConfigGetResponseTriggersExcludeRulesZarazLoadRuleOp string

const (
	ZarazConfigGetResponseTriggersExcludeRulesZarazLoadRuleOpContains           ZarazConfigGetResponseTriggersExcludeRulesZarazLoadRuleOp = "CONTAINS"
	ZarazConfigGetResponseTriggersExcludeRulesZarazLoadRuleOpEquals             ZarazConfigGetResponseTriggersExcludeRulesZarazLoadRuleOp = "EQUALS"
	ZarazConfigGetResponseTriggersExcludeRulesZarazLoadRuleOpStartsWith         ZarazConfigGetResponseTriggersExcludeRulesZarazLoadRuleOp = "STARTS_WITH"
	ZarazConfigGetResponseTriggersExcludeRulesZarazLoadRuleOpEndsWith           ZarazConfigGetResponseTriggersExcludeRulesZarazLoadRuleOp = "ENDS_WITH"
	ZarazConfigGetResponseTriggersExcludeRulesZarazLoadRuleOpMatchRegex         ZarazConfigGetResponseTriggersExcludeRulesZarazLoadRuleOp = "MATCH_REGEX"
	ZarazConfigGetResponseTriggersExcludeRulesZarazLoadRuleOpNotMatchRegex      ZarazConfigGetResponseTriggersExcludeRulesZarazLoadRuleOp = "NOT_MATCH_REGEX"
	ZarazConfigGetResponseTriggersExcludeRulesZarazLoadRuleOpGreaterThan        ZarazConfigGetResponseTriggersExcludeRulesZarazLoadRuleOp = "GREATER_THAN"
	ZarazConfigGetResponseTriggersExcludeRulesZarazLoadRuleOpGreaterThanOrEqual ZarazConfigGetResponseTriggersExcludeRulesZarazLoadRuleOp = "GREATER_THAN_OR_EQUAL"
	ZarazConfigGetResponseTriggersExcludeRulesZarazLoadRuleOpLessThan           ZarazConfigGetResponseTriggersExcludeRulesZarazLoadRuleOp = "LESS_THAN"
	ZarazConfigGetResponseTriggersExcludeRulesZarazLoadRuleOpLessThanOrEqual    ZarazConfigGetResponseTriggersExcludeRulesZarazLoadRuleOp = "LESS_THAN_OR_EQUAL"
)

type ZarazConfigGetResponseTriggersExcludeRulesZarazClickListenerRule struct {
	ID       string                                                                   `json:"id,required"`
	Action   ZarazConfigGetResponseTriggersExcludeRulesZarazClickListenerRuleAction   `json:"action,required"`
	Settings ZarazConfigGetResponseTriggersExcludeRulesZarazClickListenerRuleSettings `json:"settings,required"`
	JSON     zarazConfigGetResponseTriggersExcludeRulesZarazClickListenerRuleJSON     `json:"-"`
}

// zarazConfigGetResponseTriggersExcludeRulesZarazClickListenerRuleJSON contains
// the JSON metadata for the struct
// [ZarazConfigGetResponseTriggersExcludeRulesZarazClickListenerRule]
type zarazConfigGetResponseTriggersExcludeRulesZarazClickListenerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigGetResponseTriggersExcludeRulesZarazClickListenerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigGetResponseTriggersExcludeRulesZarazClickListenerRule) implementsZarazConfigGetResponseTriggersExcludeRule() {
}

type ZarazConfigGetResponseTriggersExcludeRulesZarazClickListenerRuleAction string

const (
	ZarazConfigGetResponseTriggersExcludeRulesZarazClickListenerRuleActionClickListener ZarazConfigGetResponseTriggersExcludeRulesZarazClickListenerRuleAction = "clickListener"
)

type ZarazConfigGetResponseTriggersExcludeRulesZarazClickListenerRuleSettings struct {
	Selector    string                                                                       `json:"selector,required"`
	Type        ZarazConfigGetResponseTriggersExcludeRulesZarazClickListenerRuleSettingsType `json:"type,required"`
	WaitForTags int64                                                                        `json:"waitForTags,required"`
	JSON        zarazConfigGetResponseTriggersExcludeRulesZarazClickListenerRuleSettingsJSON `json:"-"`
}

// zarazConfigGetResponseTriggersExcludeRulesZarazClickListenerRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazConfigGetResponseTriggersExcludeRulesZarazClickListenerRuleSettings]
type zarazConfigGetResponseTriggersExcludeRulesZarazClickListenerRuleSettingsJSON struct {
	Selector    apijson.Field
	Type        apijson.Field
	WaitForTags apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigGetResponseTriggersExcludeRulesZarazClickListenerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigGetResponseTriggersExcludeRulesZarazClickListenerRuleSettingsType string

const (
	ZarazConfigGetResponseTriggersExcludeRulesZarazClickListenerRuleSettingsTypeXpath ZarazConfigGetResponseTriggersExcludeRulesZarazClickListenerRuleSettingsType = "xpath"
	ZarazConfigGetResponseTriggersExcludeRulesZarazClickListenerRuleSettingsTypeCss   ZarazConfigGetResponseTriggersExcludeRulesZarazClickListenerRuleSettingsType = "css"
)

type ZarazConfigGetResponseTriggersExcludeRulesZarazTimerRule struct {
	ID       string                                                           `json:"id,required"`
	Action   ZarazConfigGetResponseTriggersExcludeRulesZarazTimerRuleAction   `json:"action,required"`
	Settings ZarazConfigGetResponseTriggersExcludeRulesZarazTimerRuleSettings `json:"settings,required"`
	JSON     zarazConfigGetResponseTriggersExcludeRulesZarazTimerRuleJSON     `json:"-"`
}

// zarazConfigGetResponseTriggersExcludeRulesZarazTimerRuleJSON contains the JSON
// metadata for the struct
// [ZarazConfigGetResponseTriggersExcludeRulesZarazTimerRule]
type zarazConfigGetResponseTriggersExcludeRulesZarazTimerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigGetResponseTriggersExcludeRulesZarazTimerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigGetResponseTriggersExcludeRulesZarazTimerRule) implementsZarazConfigGetResponseTriggersExcludeRule() {
}

type ZarazConfigGetResponseTriggersExcludeRulesZarazTimerRuleAction string

const (
	ZarazConfigGetResponseTriggersExcludeRulesZarazTimerRuleActionTimer ZarazConfigGetResponseTriggersExcludeRulesZarazTimerRuleAction = "timer"
)

type ZarazConfigGetResponseTriggersExcludeRulesZarazTimerRuleSettings struct {
	Interval int64                                                                `json:"interval,required"`
	Limit    int64                                                                `json:"limit,required"`
	JSON     zarazConfigGetResponseTriggersExcludeRulesZarazTimerRuleSettingsJSON `json:"-"`
}

// zarazConfigGetResponseTriggersExcludeRulesZarazTimerRuleSettingsJSON contains
// the JSON metadata for the struct
// [ZarazConfigGetResponseTriggersExcludeRulesZarazTimerRuleSettings]
type zarazConfigGetResponseTriggersExcludeRulesZarazTimerRuleSettingsJSON struct {
	Interval    apijson.Field
	Limit       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigGetResponseTriggersExcludeRulesZarazTimerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigGetResponseTriggersExcludeRulesZarazFormSubmissionRule struct {
	ID       string                                                                    `json:"id,required"`
	Action   ZarazConfigGetResponseTriggersExcludeRulesZarazFormSubmissionRuleAction   `json:"action,required"`
	Settings ZarazConfigGetResponseTriggersExcludeRulesZarazFormSubmissionRuleSettings `json:"settings,required"`
	JSON     zarazConfigGetResponseTriggersExcludeRulesZarazFormSubmissionRuleJSON     `json:"-"`
}

// zarazConfigGetResponseTriggersExcludeRulesZarazFormSubmissionRuleJSON contains
// the JSON metadata for the struct
// [ZarazConfigGetResponseTriggersExcludeRulesZarazFormSubmissionRule]
type zarazConfigGetResponseTriggersExcludeRulesZarazFormSubmissionRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigGetResponseTriggersExcludeRulesZarazFormSubmissionRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigGetResponseTriggersExcludeRulesZarazFormSubmissionRule) implementsZarazConfigGetResponseTriggersExcludeRule() {
}

type ZarazConfigGetResponseTriggersExcludeRulesZarazFormSubmissionRuleAction string

const (
	ZarazConfigGetResponseTriggersExcludeRulesZarazFormSubmissionRuleActionFormSubmission ZarazConfigGetResponseTriggersExcludeRulesZarazFormSubmissionRuleAction = "formSubmission"
)

type ZarazConfigGetResponseTriggersExcludeRulesZarazFormSubmissionRuleSettings struct {
	Selector string                                                                        `json:"selector,required"`
	Validate bool                                                                          `json:"validate,required"`
	JSON     zarazConfigGetResponseTriggersExcludeRulesZarazFormSubmissionRuleSettingsJSON `json:"-"`
}

// zarazConfigGetResponseTriggersExcludeRulesZarazFormSubmissionRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazConfigGetResponseTriggersExcludeRulesZarazFormSubmissionRuleSettings]
type zarazConfigGetResponseTriggersExcludeRulesZarazFormSubmissionRuleSettingsJSON struct {
	Selector    apijson.Field
	Validate    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigGetResponseTriggersExcludeRulesZarazFormSubmissionRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigGetResponseTriggersExcludeRulesZarazVariableMatchRule struct {
	ID       string                                                                   `json:"id,required"`
	Action   ZarazConfigGetResponseTriggersExcludeRulesZarazVariableMatchRuleAction   `json:"action,required"`
	Settings ZarazConfigGetResponseTriggersExcludeRulesZarazVariableMatchRuleSettings `json:"settings,required"`
	JSON     zarazConfigGetResponseTriggersExcludeRulesZarazVariableMatchRuleJSON     `json:"-"`
}

// zarazConfigGetResponseTriggersExcludeRulesZarazVariableMatchRuleJSON contains
// the JSON metadata for the struct
// [ZarazConfigGetResponseTriggersExcludeRulesZarazVariableMatchRule]
type zarazConfigGetResponseTriggersExcludeRulesZarazVariableMatchRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigGetResponseTriggersExcludeRulesZarazVariableMatchRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigGetResponseTriggersExcludeRulesZarazVariableMatchRule) implementsZarazConfigGetResponseTriggersExcludeRule() {
}

type ZarazConfigGetResponseTriggersExcludeRulesZarazVariableMatchRuleAction string

const (
	ZarazConfigGetResponseTriggersExcludeRulesZarazVariableMatchRuleActionVariableMatch ZarazConfigGetResponseTriggersExcludeRulesZarazVariableMatchRuleAction = "variableMatch"
)

type ZarazConfigGetResponseTriggersExcludeRulesZarazVariableMatchRuleSettings struct {
	Match    string                                                                       `json:"match,required"`
	Variable string                                                                       `json:"variable,required"`
	JSON     zarazConfigGetResponseTriggersExcludeRulesZarazVariableMatchRuleSettingsJSON `json:"-"`
}

// zarazConfigGetResponseTriggersExcludeRulesZarazVariableMatchRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazConfigGetResponseTriggersExcludeRulesZarazVariableMatchRuleSettings]
type zarazConfigGetResponseTriggersExcludeRulesZarazVariableMatchRuleSettingsJSON struct {
	Match       apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigGetResponseTriggersExcludeRulesZarazVariableMatchRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigGetResponseTriggersExcludeRulesZarazScrollDepthRule struct {
	ID       string                                                                 `json:"id,required"`
	Action   ZarazConfigGetResponseTriggersExcludeRulesZarazScrollDepthRuleAction   `json:"action,required"`
	Settings ZarazConfigGetResponseTriggersExcludeRulesZarazScrollDepthRuleSettings `json:"settings,required"`
	JSON     zarazConfigGetResponseTriggersExcludeRulesZarazScrollDepthRuleJSON     `json:"-"`
}

// zarazConfigGetResponseTriggersExcludeRulesZarazScrollDepthRuleJSON contains the
// JSON metadata for the struct
// [ZarazConfigGetResponseTriggersExcludeRulesZarazScrollDepthRule]
type zarazConfigGetResponseTriggersExcludeRulesZarazScrollDepthRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigGetResponseTriggersExcludeRulesZarazScrollDepthRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigGetResponseTriggersExcludeRulesZarazScrollDepthRule) implementsZarazConfigGetResponseTriggersExcludeRule() {
}

type ZarazConfigGetResponseTriggersExcludeRulesZarazScrollDepthRuleAction string

const (
	ZarazConfigGetResponseTriggersExcludeRulesZarazScrollDepthRuleActionScrollDepth ZarazConfigGetResponseTriggersExcludeRulesZarazScrollDepthRuleAction = "scrollDepth"
)

type ZarazConfigGetResponseTriggersExcludeRulesZarazScrollDepthRuleSettings struct {
	Positions string                                                                     `json:"positions,required"`
	JSON      zarazConfigGetResponseTriggersExcludeRulesZarazScrollDepthRuleSettingsJSON `json:"-"`
}

// zarazConfigGetResponseTriggersExcludeRulesZarazScrollDepthRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazConfigGetResponseTriggersExcludeRulesZarazScrollDepthRuleSettings]
type zarazConfigGetResponseTriggersExcludeRulesZarazScrollDepthRuleSettingsJSON struct {
	Positions   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigGetResponseTriggersExcludeRulesZarazScrollDepthRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigGetResponseTriggersExcludeRulesZarazElementVisibilityRule struct {
	ID       string                                                                       `json:"id,required"`
	Action   ZarazConfigGetResponseTriggersExcludeRulesZarazElementVisibilityRuleAction   `json:"action,required"`
	Settings ZarazConfigGetResponseTriggersExcludeRulesZarazElementVisibilityRuleSettings `json:"settings,required"`
	JSON     zarazConfigGetResponseTriggersExcludeRulesZarazElementVisibilityRuleJSON     `json:"-"`
}

// zarazConfigGetResponseTriggersExcludeRulesZarazElementVisibilityRuleJSON
// contains the JSON metadata for the struct
// [ZarazConfigGetResponseTriggersExcludeRulesZarazElementVisibilityRule]
type zarazConfigGetResponseTriggersExcludeRulesZarazElementVisibilityRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigGetResponseTriggersExcludeRulesZarazElementVisibilityRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigGetResponseTriggersExcludeRulesZarazElementVisibilityRule) implementsZarazConfigGetResponseTriggersExcludeRule() {
}

type ZarazConfigGetResponseTriggersExcludeRulesZarazElementVisibilityRuleAction string

const (
	ZarazConfigGetResponseTriggersExcludeRulesZarazElementVisibilityRuleActionElementVisibility ZarazConfigGetResponseTriggersExcludeRulesZarazElementVisibilityRuleAction = "elementVisibility"
)

type ZarazConfigGetResponseTriggersExcludeRulesZarazElementVisibilityRuleSettings struct {
	Selector string                                                                           `json:"selector,required"`
	JSON     zarazConfigGetResponseTriggersExcludeRulesZarazElementVisibilityRuleSettingsJSON `json:"-"`
}

// zarazConfigGetResponseTriggersExcludeRulesZarazElementVisibilityRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazConfigGetResponseTriggersExcludeRulesZarazElementVisibilityRuleSettings]
type zarazConfigGetResponseTriggersExcludeRulesZarazElementVisibilityRuleSettingsJSON struct {
	Selector    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigGetResponseTriggersExcludeRulesZarazElementVisibilityRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [ZarazConfigGetResponseTriggersLoadRulesZarazLoadRule],
// [ZarazConfigGetResponseTriggersLoadRulesZarazClickListenerRule],
// [ZarazConfigGetResponseTriggersLoadRulesZarazTimerRule],
// [ZarazConfigGetResponseTriggersLoadRulesZarazFormSubmissionRule],
// [ZarazConfigGetResponseTriggersLoadRulesZarazVariableMatchRule],
// [ZarazConfigGetResponseTriggersLoadRulesZarazScrollDepthRule] or
// [ZarazConfigGetResponseTriggersLoadRulesZarazElementVisibilityRule].
type ZarazConfigGetResponseTriggersLoadRule interface {
	implementsZarazConfigGetResponseTriggersLoadRule()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZarazConfigGetResponseTriggersLoadRule)(nil)).Elem(), "")
}

type ZarazConfigGetResponseTriggersLoadRulesZarazLoadRule struct {
	ID    string                                                   `json:"id,required"`
	Match string                                                   `json:"match,required"`
	Op    ZarazConfigGetResponseTriggersLoadRulesZarazLoadRuleOp   `json:"op,required"`
	Value string                                                   `json:"value,required"`
	JSON  zarazConfigGetResponseTriggersLoadRulesZarazLoadRuleJSON `json:"-"`
}

// zarazConfigGetResponseTriggersLoadRulesZarazLoadRuleJSON contains the JSON
// metadata for the struct [ZarazConfigGetResponseTriggersLoadRulesZarazLoadRule]
type zarazConfigGetResponseTriggersLoadRulesZarazLoadRuleJSON struct {
	ID          apijson.Field
	Match       apijson.Field
	Op          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigGetResponseTriggersLoadRulesZarazLoadRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigGetResponseTriggersLoadRulesZarazLoadRule) implementsZarazConfigGetResponseTriggersLoadRule() {
}

type ZarazConfigGetResponseTriggersLoadRulesZarazLoadRuleOp string

const (
	ZarazConfigGetResponseTriggersLoadRulesZarazLoadRuleOpContains           ZarazConfigGetResponseTriggersLoadRulesZarazLoadRuleOp = "CONTAINS"
	ZarazConfigGetResponseTriggersLoadRulesZarazLoadRuleOpEquals             ZarazConfigGetResponseTriggersLoadRulesZarazLoadRuleOp = "EQUALS"
	ZarazConfigGetResponseTriggersLoadRulesZarazLoadRuleOpStartsWith         ZarazConfigGetResponseTriggersLoadRulesZarazLoadRuleOp = "STARTS_WITH"
	ZarazConfigGetResponseTriggersLoadRulesZarazLoadRuleOpEndsWith           ZarazConfigGetResponseTriggersLoadRulesZarazLoadRuleOp = "ENDS_WITH"
	ZarazConfigGetResponseTriggersLoadRulesZarazLoadRuleOpMatchRegex         ZarazConfigGetResponseTriggersLoadRulesZarazLoadRuleOp = "MATCH_REGEX"
	ZarazConfigGetResponseTriggersLoadRulesZarazLoadRuleOpNotMatchRegex      ZarazConfigGetResponseTriggersLoadRulesZarazLoadRuleOp = "NOT_MATCH_REGEX"
	ZarazConfigGetResponseTriggersLoadRulesZarazLoadRuleOpGreaterThan        ZarazConfigGetResponseTriggersLoadRulesZarazLoadRuleOp = "GREATER_THAN"
	ZarazConfigGetResponseTriggersLoadRulesZarazLoadRuleOpGreaterThanOrEqual ZarazConfigGetResponseTriggersLoadRulesZarazLoadRuleOp = "GREATER_THAN_OR_EQUAL"
	ZarazConfigGetResponseTriggersLoadRulesZarazLoadRuleOpLessThan           ZarazConfigGetResponseTriggersLoadRulesZarazLoadRuleOp = "LESS_THAN"
	ZarazConfigGetResponseTriggersLoadRulesZarazLoadRuleOpLessThanOrEqual    ZarazConfigGetResponseTriggersLoadRulesZarazLoadRuleOp = "LESS_THAN_OR_EQUAL"
)

type ZarazConfigGetResponseTriggersLoadRulesZarazClickListenerRule struct {
	ID       string                                                                `json:"id,required"`
	Action   ZarazConfigGetResponseTriggersLoadRulesZarazClickListenerRuleAction   `json:"action,required"`
	Settings ZarazConfigGetResponseTriggersLoadRulesZarazClickListenerRuleSettings `json:"settings,required"`
	JSON     zarazConfigGetResponseTriggersLoadRulesZarazClickListenerRuleJSON     `json:"-"`
}

// zarazConfigGetResponseTriggersLoadRulesZarazClickListenerRuleJSON contains the
// JSON metadata for the struct
// [ZarazConfigGetResponseTriggersLoadRulesZarazClickListenerRule]
type zarazConfigGetResponseTriggersLoadRulesZarazClickListenerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigGetResponseTriggersLoadRulesZarazClickListenerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigGetResponseTriggersLoadRulesZarazClickListenerRule) implementsZarazConfigGetResponseTriggersLoadRule() {
}

type ZarazConfigGetResponseTriggersLoadRulesZarazClickListenerRuleAction string

const (
	ZarazConfigGetResponseTriggersLoadRulesZarazClickListenerRuleActionClickListener ZarazConfigGetResponseTriggersLoadRulesZarazClickListenerRuleAction = "clickListener"
)

type ZarazConfigGetResponseTriggersLoadRulesZarazClickListenerRuleSettings struct {
	Selector    string                                                                    `json:"selector,required"`
	Type        ZarazConfigGetResponseTriggersLoadRulesZarazClickListenerRuleSettingsType `json:"type,required"`
	WaitForTags int64                                                                     `json:"waitForTags,required"`
	JSON        zarazConfigGetResponseTriggersLoadRulesZarazClickListenerRuleSettingsJSON `json:"-"`
}

// zarazConfigGetResponseTriggersLoadRulesZarazClickListenerRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazConfigGetResponseTriggersLoadRulesZarazClickListenerRuleSettings]
type zarazConfigGetResponseTriggersLoadRulesZarazClickListenerRuleSettingsJSON struct {
	Selector    apijson.Field
	Type        apijson.Field
	WaitForTags apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigGetResponseTriggersLoadRulesZarazClickListenerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigGetResponseTriggersLoadRulesZarazClickListenerRuleSettingsType string

const (
	ZarazConfigGetResponseTriggersLoadRulesZarazClickListenerRuleSettingsTypeXpath ZarazConfigGetResponseTriggersLoadRulesZarazClickListenerRuleSettingsType = "xpath"
	ZarazConfigGetResponseTriggersLoadRulesZarazClickListenerRuleSettingsTypeCss   ZarazConfigGetResponseTriggersLoadRulesZarazClickListenerRuleSettingsType = "css"
)

type ZarazConfigGetResponseTriggersLoadRulesZarazTimerRule struct {
	ID       string                                                        `json:"id,required"`
	Action   ZarazConfigGetResponseTriggersLoadRulesZarazTimerRuleAction   `json:"action,required"`
	Settings ZarazConfigGetResponseTriggersLoadRulesZarazTimerRuleSettings `json:"settings,required"`
	JSON     zarazConfigGetResponseTriggersLoadRulesZarazTimerRuleJSON     `json:"-"`
}

// zarazConfigGetResponseTriggersLoadRulesZarazTimerRuleJSON contains the JSON
// metadata for the struct [ZarazConfigGetResponseTriggersLoadRulesZarazTimerRule]
type zarazConfigGetResponseTriggersLoadRulesZarazTimerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigGetResponseTriggersLoadRulesZarazTimerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigGetResponseTriggersLoadRulesZarazTimerRule) implementsZarazConfigGetResponseTriggersLoadRule() {
}

type ZarazConfigGetResponseTriggersLoadRulesZarazTimerRuleAction string

const (
	ZarazConfigGetResponseTriggersLoadRulesZarazTimerRuleActionTimer ZarazConfigGetResponseTriggersLoadRulesZarazTimerRuleAction = "timer"
)

type ZarazConfigGetResponseTriggersLoadRulesZarazTimerRuleSettings struct {
	Interval int64                                                             `json:"interval,required"`
	Limit    int64                                                             `json:"limit,required"`
	JSON     zarazConfigGetResponseTriggersLoadRulesZarazTimerRuleSettingsJSON `json:"-"`
}

// zarazConfigGetResponseTriggersLoadRulesZarazTimerRuleSettingsJSON contains the
// JSON metadata for the struct
// [ZarazConfigGetResponseTriggersLoadRulesZarazTimerRuleSettings]
type zarazConfigGetResponseTriggersLoadRulesZarazTimerRuleSettingsJSON struct {
	Interval    apijson.Field
	Limit       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigGetResponseTriggersLoadRulesZarazTimerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigGetResponseTriggersLoadRulesZarazFormSubmissionRule struct {
	ID       string                                                                 `json:"id,required"`
	Action   ZarazConfigGetResponseTriggersLoadRulesZarazFormSubmissionRuleAction   `json:"action,required"`
	Settings ZarazConfigGetResponseTriggersLoadRulesZarazFormSubmissionRuleSettings `json:"settings,required"`
	JSON     zarazConfigGetResponseTriggersLoadRulesZarazFormSubmissionRuleJSON     `json:"-"`
}

// zarazConfigGetResponseTriggersLoadRulesZarazFormSubmissionRuleJSON contains the
// JSON metadata for the struct
// [ZarazConfigGetResponseTriggersLoadRulesZarazFormSubmissionRule]
type zarazConfigGetResponseTriggersLoadRulesZarazFormSubmissionRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigGetResponseTriggersLoadRulesZarazFormSubmissionRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigGetResponseTriggersLoadRulesZarazFormSubmissionRule) implementsZarazConfigGetResponseTriggersLoadRule() {
}

type ZarazConfigGetResponseTriggersLoadRulesZarazFormSubmissionRuleAction string

const (
	ZarazConfigGetResponseTriggersLoadRulesZarazFormSubmissionRuleActionFormSubmission ZarazConfigGetResponseTriggersLoadRulesZarazFormSubmissionRuleAction = "formSubmission"
)

type ZarazConfigGetResponseTriggersLoadRulesZarazFormSubmissionRuleSettings struct {
	Selector string                                                                     `json:"selector,required"`
	Validate bool                                                                       `json:"validate,required"`
	JSON     zarazConfigGetResponseTriggersLoadRulesZarazFormSubmissionRuleSettingsJSON `json:"-"`
}

// zarazConfigGetResponseTriggersLoadRulesZarazFormSubmissionRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazConfigGetResponseTriggersLoadRulesZarazFormSubmissionRuleSettings]
type zarazConfigGetResponseTriggersLoadRulesZarazFormSubmissionRuleSettingsJSON struct {
	Selector    apijson.Field
	Validate    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigGetResponseTriggersLoadRulesZarazFormSubmissionRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigGetResponseTriggersLoadRulesZarazVariableMatchRule struct {
	ID       string                                                                `json:"id,required"`
	Action   ZarazConfigGetResponseTriggersLoadRulesZarazVariableMatchRuleAction   `json:"action,required"`
	Settings ZarazConfigGetResponseTriggersLoadRulesZarazVariableMatchRuleSettings `json:"settings,required"`
	JSON     zarazConfigGetResponseTriggersLoadRulesZarazVariableMatchRuleJSON     `json:"-"`
}

// zarazConfigGetResponseTriggersLoadRulesZarazVariableMatchRuleJSON contains the
// JSON metadata for the struct
// [ZarazConfigGetResponseTriggersLoadRulesZarazVariableMatchRule]
type zarazConfigGetResponseTriggersLoadRulesZarazVariableMatchRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigGetResponseTriggersLoadRulesZarazVariableMatchRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigGetResponseTriggersLoadRulesZarazVariableMatchRule) implementsZarazConfigGetResponseTriggersLoadRule() {
}

type ZarazConfigGetResponseTriggersLoadRulesZarazVariableMatchRuleAction string

const (
	ZarazConfigGetResponseTriggersLoadRulesZarazVariableMatchRuleActionVariableMatch ZarazConfigGetResponseTriggersLoadRulesZarazVariableMatchRuleAction = "variableMatch"
)

type ZarazConfigGetResponseTriggersLoadRulesZarazVariableMatchRuleSettings struct {
	Match    string                                                                    `json:"match,required"`
	Variable string                                                                    `json:"variable,required"`
	JSON     zarazConfigGetResponseTriggersLoadRulesZarazVariableMatchRuleSettingsJSON `json:"-"`
}

// zarazConfigGetResponseTriggersLoadRulesZarazVariableMatchRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazConfigGetResponseTriggersLoadRulesZarazVariableMatchRuleSettings]
type zarazConfigGetResponseTriggersLoadRulesZarazVariableMatchRuleSettingsJSON struct {
	Match       apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigGetResponseTriggersLoadRulesZarazVariableMatchRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigGetResponseTriggersLoadRulesZarazScrollDepthRule struct {
	ID       string                                                              `json:"id,required"`
	Action   ZarazConfigGetResponseTriggersLoadRulesZarazScrollDepthRuleAction   `json:"action,required"`
	Settings ZarazConfigGetResponseTriggersLoadRulesZarazScrollDepthRuleSettings `json:"settings,required"`
	JSON     zarazConfigGetResponseTriggersLoadRulesZarazScrollDepthRuleJSON     `json:"-"`
}

// zarazConfigGetResponseTriggersLoadRulesZarazScrollDepthRuleJSON contains the
// JSON metadata for the struct
// [ZarazConfigGetResponseTriggersLoadRulesZarazScrollDepthRule]
type zarazConfigGetResponseTriggersLoadRulesZarazScrollDepthRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigGetResponseTriggersLoadRulesZarazScrollDepthRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigGetResponseTriggersLoadRulesZarazScrollDepthRule) implementsZarazConfigGetResponseTriggersLoadRule() {
}

type ZarazConfigGetResponseTriggersLoadRulesZarazScrollDepthRuleAction string

const (
	ZarazConfigGetResponseTriggersLoadRulesZarazScrollDepthRuleActionScrollDepth ZarazConfigGetResponseTriggersLoadRulesZarazScrollDepthRuleAction = "scrollDepth"
)

type ZarazConfigGetResponseTriggersLoadRulesZarazScrollDepthRuleSettings struct {
	Positions string                                                                  `json:"positions,required"`
	JSON      zarazConfigGetResponseTriggersLoadRulesZarazScrollDepthRuleSettingsJSON `json:"-"`
}

// zarazConfigGetResponseTriggersLoadRulesZarazScrollDepthRuleSettingsJSON contains
// the JSON metadata for the struct
// [ZarazConfigGetResponseTriggersLoadRulesZarazScrollDepthRuleSettings]
type zarazConfigGetResponseTriggersLoadRulesZarazScrollDepthRuleSettingsJSON struct {
	Positions   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigGetResponseTriggersLoadRulesZarazScrollDepthRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigGetResponseTriggersLoadRulesZarazElementVisibilityRule struct {
	ID       string                                                                    `json:"id,required"`
	Action   ZarazConfigGetResponseTriggersLoadRulesZarazElementVisibilityRuleAction   `json:"action,required"`
	Settings ZarazConfigGetResponseTriggersLoadRulesZarazElementVisibilityRuleSettings `json:"settings,required"`
	JSON     zarazConfigGetResponseTriggersLoadRulesZarazElementVisibilityRuleJSON     `json:"-"`
}

// zarazConfigGetResponseTriggersLoadRulesZarazElementVisibilityRuleJSON contains
// the JSON metadata for the struct
// [ZarazConfigGetResponseTriggersLoadRulesZarazElementVisibilityRule]
type zarazConfigGetResponseTriggersLoadRulesZarazElementVisibilityRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigGetResponseTriggersLoadRulesZarazElementVisibilityRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigGetResponseTriggersLoadRulesZarazElementVisibilityRule) implementsZarazConfigGetResponseTriggersLoadRule() {
}

type ZarazConfigGetResponseTriggersLoadRulesZarazElementVisibilityRuleAction string

const (
	ZarazConfigGetResponseTriggersLoadRulesZarazElementVisibilityRuleActionElementVisibility ZarazConfigGetResponseTriggersLoadRulesZarazElementVisibilityRuleAction = "elementVisibility"
)

type ZarazConfigGetResponseTriggersLoadRulesZarazElementVisibilityRuleSettings struct {
	Selector string                                                                        `json:"selector,required"`
	JSON     zarazConfigGetResponseTriggersLoadRulesZarazElementVisibilityRuleSettingsJSON `json:"-"`
}

// zarazConfigGetResponseTriggersLoadRulesZarazElementVisibilityRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazConfigGetResponseTriggersLoadRulesZarazElementVisibilityRuleSettings]
type zarazConfigGetResponseTriggersLoadRulesZarazElementVisibilityRuleSettingsJSON struct {
	Selector    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigGetResponseTriggersLoadRulesZarazElementVisibilityRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigGetResponseTriggersSystem string

const (
	ZarazConfigGetResponseTriggersSystemPageload ZarazConfigGetResponseTriggersSystem = "pageload"
)

// Union satisfied by [ZarazConfigGetResponseVariablesObject] or
// [ZarazConfigGetResponseVariablesObject].
type ZarazConfigGetResponseVariable interface {
	implementsZarazConfigGetResponseVariable()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZarazConfigGetResponseVariable)(nil)).Elem(), "")
}

type ZarazConfigGetResponseVariablesObject struct {
	Name  string                                    `json:"name,required"`
	Type  ZarazConfigGetResponseVariablesObjectType `json:"type,required"`
	Value string                                    `json:"value,required"`
	JSON  zarazConfigGetResponseVariablesObjectJSON `json:"-"`
}

// zarazConfigGetResponseVariablesObjectJSON contains the JSON metadata for the
// struct [ZarazConfigGetResponseVariablesObject]
type zarazConfigGetResponseVariablesObjectJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigGetResponseVariablesObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigGetResponseVariablesObject) implementsZarazConfigGetResponseVariable() {}

type ZarazConfigGetResponseVariablesObjectType string

const (
	ZarazConfigGetResponseVariablesObjectTypeString ZarazConfigGetResponseVariablesObjectType = "string"
	ZarazConfigGetResponseVariablesObjectTypeSecret ZarazConfigGetResponseVariablesObjectType = "secret"
)

// Consent management configuration.
type ZarazConfigGetResponseConsent struct {
	Enabled                bool                                                `json:"enabled,required"`
	ButtonTextTranslations ZarazConfigGetResponseConsentButtonTextTranslations `json:"buttonTextTranslations"`
	CompanyEmail           string                                              `json:"companyEmail"`
	CompanyName            string                                              `json:"companyName"`
	CompanyStreetAddress   string                                              `json:"companyStreetAddress"`
	ConsentModalIntroHTML  string                                              `json:"consentModalIntroHTML"`
	// Object where keys are language codes
	ConsentModalIntroHTMLWithTranslations map[string]string `json:"consentModalIntroHTMLWithTranslations"`
	CookieName                            string            `json:"cookieName"`
	CustomCss                             string            `json:"customCSS"`
	CustomIntroDisclaimerDismissed        bool              `json:"customIntroDisclaimerDismissed"`
	DefaultLanguage                       string            `json:"defaultLanguage"`
	HideModal                             bool              `json:"hideModal"`
	// Object where keys are purpose alpha-numeric IDs
	Purposes map[string]ZarazConfigGetResponseConsentPurpose `json:"purposes"`
	// Object where keys are purpose alpha-numeric IDs
	PurposesWithTranslations map[string]ZarazConfigGetResponseConsentPurposesWithTranslation `json:"purposesWithTranslations"`
	JSON                     zarazConfigGetResponseConsentJSON                               `json:"-"`
}

// zarazConfigGetResponseConsentJSON contains the JSON metadata for the struct
// [ZarazConfigGetResponseConsent]
type zarazConfigGetResponseConsentJSON struct {
	Enabled                               apijson.Field
	ButtonTextTranslations                apijson.Field
	CompanyEmail                          apijson.Field
	CompanyName                           apijson.Field
	CompanyStreetAddress                  apijson.Field
	ConsentModalIntroHTML                 apijson.Field
	ConsentModalIntroHTMLWithTranslations apijson.Field
	CookieName                            apijson.Field
	CustomCss                             apijson.Field
	CustomIntroDisclaimerDismissed        apijson.Field
	DefaultLanguage                       apijson.Field
	HideModal                             apijson.Field
	Purposes                              apijson.Field
	PurposesWithTranslations              apijson.Field
	raw                                   string
	ExtraFields                           map[string]apijson.Field
}

func (r *ZarazConfigGetResponseConsent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigGetResponseConsentButtonTextTranslations struct {
	// Object where keys are language codes
	AcceptAll map[string]string `json:"accept_all,required"`
	// Object where keys are language codes
	ConfirmMyChoices map[string]string `json:"confirm_my_choices,required"`
	// Object where keys are language codes
	RejectAll map[string]string                                       `json:"reject_all,required"`
	JSON      zarazConfigGetResponseConsentButtonTextTranslationsJSON `json:"-"`
}

// zarazConfigGetResponseConsentButtonTextTranslationsJSON contains the JSON
// metadata for the struct [ZarazConfigGetResponseConsentButtonTextTranslations]
type zarazConfigGetResponseConsentButtonTextTranslationsJSON struct {
	AcceptAll        apijson.Field
	ConfirmMyChoices apijson.Field
	RejectAll        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazConfigGetResponseConsentButtonTextTranslations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigGetResponseConsentPurpose struct {
	Description string                                   `json:"description,required"`
	Name        string                                   `json:"name,required"`
	JSON        zarazConfigGetResponseConsentPurposeJSON `json:"-"`
}

// zarazConfigGetResponseConsentPurposeJSON contains the JSON metadata for the
// struct [ZarazConfigGetResponseConsentPurpose]
type zarazConfigGetResponseConsentPurposeJSON struct {
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigGetResponseConsentPurpose) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigGetResponseConsentPurposesWithTranslation struct {
	// Object where keys are language codes
	Description map[string]string `json:"description,required"`
	// Object where keys are language codes
	Name  map[string]string                                        `json:"name,required"`
	Order int64                                                    `json:"order,required"`
	JSON  zarazConfigGetResponseConsentPurposesWithTranslationJSON `json:"-"`
}

// zarazConfigGetResponseConsentPurposesWithTranslationJSON contains the JSON
// metadata for the struct [ZarazConfigGetResponseConsentPurposesWithTranslation]
type zarazConfigGetResponseConsentPurposesWithTranslationJSON struct {
	Description apijson.Field
	Name        apijson.Field
	Order       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigGetResponseConsentPurposesWithTranslation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Zaraz configuration
type ZarazConfigReplaceResponse struct {
	// Data layer compatibility mode enabled.
	DataLayer bool `json:"dataLayer,required"`
	// The key for Zaraz debug mode.
	DebugKey string `json:"debugKey,required"`
	// General Zaraz settings.
	Settings ZarazConfigReplaceResponseSettings `json:"settings,required"`
	// Tools set up under Zaraz configuration, where key is the alpha-numeric tool ID
	// and value is the tool configuration object.
	Tools map[string]ZarazConfigReplaceResponseTool `json:"tools,required"`
	// Triggers set up under Zaraz configuration, where key is the trigger
	// alpha-numeric ID and value is the trigger configuration.
	Triggers map[string]ZarazConfigReplaceResponseTrigger `json:"triggers,required"`
	// Variables set up under Zaraz configuration, where key is the variable
	// alpha-numeric ID and value is the variable configuration. Values of variables of
	// type secret are not included.
	Variables map[string]ZarazConfigReplaceResponseVariable `json:"variables,required"`
	// Zaraz internal version of the config.
	ZarazVersion int64 `json:"zarazVersion,required"`
	// Consent management configuration.
	Consent ZarazConfigReplaceResponseConsent `json:"consent"`
	// Single Page Application support enabled.
	HistoryChange bool                           `json:"historyChange"`
	JSON          zarazConfigReplaceResponseJSON `json:"-"`
}

// zarazConfigReplaceResponseJSON contains the JSON metadata for the struct
// [ZarazConfigReplaceResponse]
type zarazConfigReplaceResponseJSON struct {
	DataLayer     apijson.Field
	DebugKey      apijson.Field
	Settings      apijson.Field
	Tools         apijson.Field
	Triggers      apijson.Field
	Variables     apijson.Field
	ZarazVersion  apijson.Field
	Consent       apijson.Field
	HistoryChange apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// General Zaraz settings.
type ZarazConfigReplaceResponseSettings struct {
	// Automatic injection of Zaraz scripts enabled.
	AutoInjectScript bool `json:"autoInjectScript,required"`
	// Details of the worker that receives and edits Zaraz Context object.
	ContextEnricher ZarazConfigReplaceResponseSettingsContextEnricher `json:"contextEnricher"`
	// The domain Zaraz will use for writing and reading its cookies.
	CookieDomain string `json:"cookieDomain"`
	// Ecommerce API enabled.
	Ecommerce bool `json:"ecommerce"`
	// Custom endpoint for server-side track events.
	EventsAPIPath string `json:"eventsApiPath"`
	// Hiding external referrer URL enabled.
	HideExternalReferer bool `json:"hideExternalReferer"`
	// Trimming IP address enabled.
	HideIPAddress bool `json:"hideIPAddress"`
	// Removing URL query params enabled.
	HideQueryParams bool `json:"hideQueryParams"`
	// Removing sensitive data from User Aagent string enabled.
	HideUserAgent bool `json:"hideUserAgent"`
	// Custom endpoint for Zaraz init script.
	InitPath string `json:"initPath"`
	// Injection of Zaraz scripts into iframes enabled.
	InjectIframes bool `json:"injectIframes"`
	// Custom path for Managed Components server functionalities.
	McRootPath string `json:"mcRootPath"`
	// Custom endpoint for Zaraz main script.
	ScriptPath string `json:"scriptPath"`
	// Custom endpoint for Zaraz tracking requests.
	TrackPath string                                 `json:"trackPath"`
	JSON      zarazConfigReplaceResponseSettingsJSON `json:"-"`
}

// zarazConfigReplaceResponseSettingsJSON contains the JSON metadata for the struct
// [ZarazConfigReplaceResponseSettings]
type zarazConfigReplaceResponseSettingsJSON struct {
	AutoInjectScript    apijson.Field
	ContextEnricher     apijson.Field
	CookieDomain        apijson.Field
	Ecommerce           apijson.Field
	EventsAPIPath       apijson.Field
	HideExternalReferer apijson.Field
	HideIPAddress       apijson.Field
	HideQueryParams     apijson.Field
	HideUserAgent       apijson.Field
	InitPath            apijson.Field
	InjectIframes       apijson.Field
	McRootPath          apijson.Field
	ScriptPath          apijson.Field
	TrackPath           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details of the worker that receives and edits Zaraz Context object.
type ZarazConfigReplaceResponseSettingsContextEnricher struct {
	EscapedWorkerName string                                                `json:"escapedWorkerName,required"`
	WorkerTag         string                                                `json:"workerTag,required"`
	JSON              zarazConfigReplaceResponseSettingsContextEnricherJSON `json:"-"`
}

// zarazConfigReplaceResponseSettingsContextEnricherJSON contains the JSON metadata
// for the struct [ZarazConfigReplaceResponseSettingsContextEnricher]
type zarazConfigReplaceResponseSettingsContextEnricherJSON struct {
	EscapedWorkerName apijson.Field
	WorkerTag         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseSettingsContextEnricher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [ZarazConfigReplaceResponseToolsZarazManagedComponent] or
// [ZarazConfigReplaceResponseToolsZarazCustomManagedComponent].
type ZarazConfigReplaceResponseTool interface {
	implementsZarazConfigReplaceResponseTool()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZarazConfigReplaceResponseTool)(nil)).Elem(), "")
}

type ZarazConfigReplaceResponseToolsZarazManagedComponent struct {
	// List of blocking trigger IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Tool's internal name
	Component string `json:"component,required"`
	// Default fields for tool's actions
	DefaultFields map[string]ZarazConfigReplaceResponseToolsZarazManagedComponentDefaultField `json:"defaultFields,required"`
	// Whether tool is enabled
	Enabled bool `json:"enabled,required"`
	// Tool's name defined by the user
	Name string `json:"name,required"`
	// List of permissions granted to the component
	Permissions []string `json:"permissions,required"`
	// Tool's settings
	Settings map[string]ZarazConfigReplaceResponseToolsZarazManagedComponentSetting `json:"settings,required"`
	Type     ZarazConfigReplaceResponseToolsZarazManagedComponentType               `json:"type,required"`
	// Actions configured on a tool. Either this or neoEvents field is required.
	Actions map[string]ZarazConfigReplaceResponseToolsZarazManagedComponentAction `json:"actions"`
	// Default consent purpose ID
	DefaultPurpose string `json:"defaultPurpose"`
	// DEPRECATED - List of actions configured on a tool. Either this or actions field
	// is required. If both are present, actions field will take precedence.
	NeoEvents []ZarazConfigReplaceResponseToolsZarazManagedComponentNeoEvent `json:"neoEvents"`
	JSON      zarazConfigReplaceResponseToolsZarazManagedComponentJSON       `json:"-"`
}

// zarazConfigReplaceResponseToolsZarazManagedComponentJSON contains the JSON
// metadata for the struct [ZarazConfigReplaceResponseToolsZarazManagedComponent]
type zarazConfigReplaceResponseToolsZarazManagedComponentJSON struct {
	BlockingTriggers apijson.Field
	Component        apijson.Field
	DefaultFields    apijson.Field
	Enabled          apijson.Field
	Name             apijson.Field
	Permissions      apijson.Field
	Settings         apijson.Field
	Type             apijson.Field
	Actions          apijson.Field
	DefaultPurpose   apijson.Field
	NeoEvents        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseToolsZarazManagedComponent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigReplaceResponseToolsZarazManagedComponent) implementsZarazConfigReplaceResponseTool() {
}

// Union satisfied by [shared.UnionString] or [shared.UnionBool].
type ZarazConfigReplaceResponseToolsZarazManagedComponentDefaultField interface {
	ImplementsZarazConfigReplaceResponseToolsZarazManagedComponentDefaultField()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazConfigReplaceResponseToolsZarazManagedComponentDefaultField)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

// Union satisfied by [shared.UnionString] or [shared.UnionBool].
type ZarazConfigReplaceResponseToolsZarazManagedComponentSetting interface {
	ImplementsZarazConfigReplaceResponseToolsZarazManagedComponentSetting()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazConfigReplaceResponseToolsZarazManagedComponentSetting)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

type ZarazConfigReplaceResponseToolsZarazManagedComponentType string

const (
	ZarazConfigReplaceResponseToolsZarazManagedComponentTypeComponent ZarazConfigReplaceResponseToolsZarazManagedComponentType = "component"
)

type ZarazConfigReplaceResponseToolsZarazManagedComponentAction struct {
	// Tool event type
	ActionType string `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Event payload
	Data interface{} `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers []string                                                       `json:"firingTriggers,required"`
	JSON           zarazConfigReplaceResponseToolsZarazManagedComponentActionJSON `json:"-"`
}

// zarazConfigReplaceResponseToolsZarazManagedComponentActionJSON contains the JSON
// metadata for the struct
// [ZarazConfigReplaceResponseToolsZarazManagedComponentAction]
type zarazConfigReplaceResponseToolsZarazManagedComponentActionJSON struct {
	ActionType       apijson.Field
	BlockingTriggers apijson.Field
	Data             apijson.Field
	FiringTriggers   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseToolsZarazManagedComponentAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigReplaceResponseToolsZarazManagedComponentNeoEvent struct {
	// Tool event type
	ActionType string `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Event payload
	Data interface{} `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers []string                                                         `json:"firingTriggers,required"`
	JSON           zarazConfigReplaceResponseToolsZarazManagedComponentNeoEventJSON `json:"-"`
}

// zarazConfigReplaceResponseToolsZarazManagedComponentNeoEventJSON contains the
// JSON metadata for the struct
// [ZarazConfigReplaceResponseToolsZarazManagedComponentNeoEvent]
type zarazConfigReplaceResponseToolsZarazManagedComponentNeoEventJSON struct {
	ActionType       apijson.Field
	BlockingTriggers apijson.Field
	Data             apijson.Field
	FiringTriggers   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseToolsZarazManagedComponentNeoEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigReplaceResponseToolsZarazCustomManagedComponent struct {
	// List of blocking trigger IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Tool's internal name
	Component string `json:"component,required"`
	// Default fields for tool's actions
	DefaultFields map[string]ZarazConfigReplaceResponseToolsZarazCustomManagedComponentDefaultField `json:"defaultFields,required"`
	// Whether tool is enabled
	Enabled bool `json:"enabled,required"`
	// Tool's name defined by the user
	Name string `json:"name,required"`
	// List of permissions granted to the component
	Permissions []string `json:"permissions,required"`
	// Tool's settings
	Settings map[string]ZarazConfigReplaceResponseToolsZarazCustomManagedComponentSetting `json:"settings,required"`
	Type     ZarazConfigReplaceResponseToolsZarazCustomManagedComponentType               `json:"type,required"`
	// Cloudflare worker that acts as a managed component
	Worker ZarazConfigReplaceResponseToolsZarazCustomManagedComponentWorker `json:"worker,required"`
	// Actions configured on a tool. Either this or neoEvents field is required.
	Actions map[string]ZarazConfigReplaceResponseToolsZarazCustomManagedComponentAction `json:"actions"`
	// Default consent purpose ID
	DefaultPurpose string `json:"defaultPurpose"`
	// DEPRECATED - List of actions configured on a tool. Either this or actions field
	// is required. If both are present, actions field will take precedence.
	NeoEvents []ZarazConfigReplaceResponseToolsZarazCustomManagedComponentNeoEvent `json:"neoEvents"`
	JSON      zarazConfigReplaceResponseToolsZarazCustomManagedComponentJSON       `json:"-"`
}

// zarazConfigReplaceResponseToolsZarazCustomManagedComponentJSON contains the JSON
// metadata for the struct
// [ZarazConfigReplaceResponseToolsZarazCustomManagedComponent]
type zarazConfigReplaceResponseToolsZarazCustomManagedComponentJSON struct {
	BlockingTriggers apijson.Field
	Component        apijson.Field
	DefaultFields    apijson.Field
	Enabled          apijson.Field
	Name             apijson.Field
	Permissions      apijson.Field
	Settings         apijson.Field
	Type             apijson.Field
	Worker           apijson.Field
	Actions          apijson.Field
	DefaultPurpose   apijson.Field
	NeoEvents        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseToolsZarazCustomManagedComponent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigReplaceResponseToolsZarazCustomManagedComponent) implementsZarazConfigReplaceResponseTool() {
}

// Union satisfied by [shared.UnionString] or [shared.UnionBool].
type ZarazConfigReplaceResponseToolsZarazCustomManagedComponentDefaultField interface {
	ImplementsZarazConfigReplaceResponseToolsZarazCustomManagedComponentDefaultField()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazConfigReplaceResponseToolsZarazCustomManagedComponentDefaultField)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

// Union satisfied by [shared.UnionString] or [shared.UnionBool].
type ZarazConfigReplaceResponseToolsZarazCustomManagedComponentSetting interface {
	ImplementsZarazConfigReplaceResponseToolsZarazCustomManagedComponentSetting()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazConfigReplaceResponseToolsZarazCustomManagedComponentSetting)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

type ZarazConfigReplaceResponseToolsZarazCustomManagedComponentType string

const (
	ZarazConfigReplaceResponseToolsZarazCustomManagedComponentTypeCustomMc ZarazConfigReplaceResponseToolsZarazCustomManagedComponentType = "custom-mc"
)

// Cloudflare worker that acts as a managed component
type ZarazConfigReplaceResponseToolsZarazCustomManagedComponentWorker struct {
	EscapedWorkerName string                                                               `json:"escapedWorkerName,required"`
	WorkerTag         string                                                               `json:"workerTag,required"`
	JSON              zarazConfigReplaceResponseToolsZarazCustomManagedComponentWorkerJSON `json:"-"`
}

// zarazConfigReplaceResponseToolsZarazCustomManagedComponentWorkerJSON contains
// the JSON metadata for the struct
// [ZarazConfigReplaceResponseToolsZarazCustomManagedComponentWorker]
type zarazConfigReplaceResponseToolsZarazCustomManagedComponentWorkerJSON struct {
	EscapedWorkerName apijson.Field
	WorkerTag         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseToolsZarazCustomManagedComponentWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigReplaceResponseToolsZarazCustomManagedComponentAction struct {
	// Tool event type
	ActionType string `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Event payload
	Data interface{} `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers []string                                                             `json:"firingTriggers,required"`
	JSON           zarazConfigReplaceResponseToolsZarazCustomManagedComponentActionJSON `json:"-"`
}

// zarazConfigReplaceResponseToolsZarazCustomManagedComponentActionJSON contains
// the JSON metadata for the struct
// [ZarazConfigReplaceResponseToolsZarazCustomManagedComponentAction]
type zarazConfigReplaceResponseToolsZarazCustomManagedComponentActionJSON struct {
	ActionType       apijson.Field
	BlockingTriggers apijson.Field
	Data             apijson.Field
	FiringTriggers   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseToolsZarazCustomManagedComponentAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigReplaceResponseToolsZarazCustomManagedComponentNeoEvent struct {
	// Tool event type
	ActionType string `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Event payload
	Data interface{} `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers []string                                                               `json:"firingTriggers,required"`
	JSON           zarazConfigReplaceResponseToolsZarazCustomManagedComponentNeoEventJSON `json:"-"`
}

// zarazConfigReplaceResponseToolsZarazCustomManagedComponentNeoEventJSON contains
// the JSON metadata for the struct
// [ZarazConfigReplaceResponseToolsZarazCustomManagedComponentNeoEvent]
type zarazConfigReplaceResponseToolsZarazCustomManagedComponentNeoEventJSON struct {
	ActionType       apijson.Field
	BlockingTriggers apijson.Field
	Data             apijson.Field
	FiringTriggers   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseToolsZarazCustomManagedComponentNeoEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigReplaceResponseTrigger struct {
	// Rules defining when the trigger is not fired.
	ExcludeRules []ZarazConfigReplaceResponseTriggersExcludeRule `json:"excludeRules,required"`
	// Rules defining when the trigger is fired.
	LoadRules []ZarazConfigReplaceResponseTriggersLoadRule `json:"loadRules,required"`
	// Trigger name.
	Name string `json:"name,required"`
	// Trigger description.
	Description string                                   `json:"description"`
	System      ZarazConfigReplaceResponseTriggersSystem `json:"system"`
	JSON        zarazConfigReplaceResponseTriggerJSON    `json:"-"`
}

// zarazConfigReplaceResponseTriggerJSON contains the JSON metadata for the struct
// [ZarazConfigReplaceResponseTrigger]
type zarazConfigReplaceResponseTriggerJSON struct {
	ExcludeRules apijson.Field
	LoadRules    apijson.Field
	Name         apijson.Field
	Description  apijson.Field
	System       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [ZarazConfigReplaceResponseTriggersExcludeRulesZarazLoadRule],
// [ZarazConfigReplaceResponseTriggersExcludeRulesZarazClickListenerRule],
// [ZarazConfigReplaceResponseTriggersExcludeRulesZarazTimerRule],
// [ZarazConfigReplaceResponseTriggersExcludeRulesZarazFormSubmissionRule],
// [ZarazConfigReplaceResponseTriggersExcludeRulesZarazVariableMatchRule],
// [ZarazConfigReplaceResponseTriggersExcludeRulesZarazScrollDepthRule] or
// [ZarazConfigReplaceResponseTriggersExcludeRulesZarazElementVisibilityRule].
type ZarazConfigReplaceResponseTriggersExcludeRule interface {
	implementsZarazConfigReplaceResponseTriggersExcludeRule()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZarazConfigReplaceResponseTriggersExcludeRule)(nil)).Elem(), "")
}

type ZarazConfigReplaceResponseTriggersExcludeRulesZarazLoadRule struct {
	ID    string                                                          `json:"id,required"`
	Match string                                                          `json:"match,required"`
	Op    ZarazConfigReplaceResponseTriggersExcludeRulesZarazLoadRuleOp   `json:"op,required"`
	Value string                                                          `json:"value,required"`
	JSON  zarazConfigReplaceResponseTriggersExcludeRulesZarazLoadRuleJSON `json:"-"`
}

// zarazConfigReplaceResponseTriggersExcludeRulesZarazLoadRuleJSON contains the
// JSON metadata for the struct
// [ZarazConfigReplaceResponseTriggersExcludeRulesZarazLoadRule]
type zarazConfigReplaceResponseTriggersExcludeRulesZarazLoadRuleJSON struct {
	ID          apijson.Field
	Match       apijson.Field
	Op          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseTriggersExcludeRulesZarazLoadRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigReplaceResponseTriggersExcludeRulesZarazLoadRule) implementsZarazConfigReplaceResponseTriggersExcludeRule() {
}

type ZarazConfigReplaceResponseTriggersExcludeRulesZarazLoadRuleOp string

const (
	ZarazConfigReplaceResponseTriggersExcludeRulesZarazLoadRuleOpContains           ZarazConfigReplaceResponseTriggersExcludeRulesZarazLoadRuleOp = "CONTAINS"
	ZarazConfigReplaceResponseTriggersExcludeRulesZarazLoadRuleOpEquals             ZarazConfigReplaceResponseTriggersExcludeRulesZarazLoadRuleOp = "EQUALS"
	ZarazConfigReplaceResponseTriggersExcludeRulesZarazLoadRuleOpStartsWith         ZarazConfigReplaceResponseTriggersExcludeRulesZarazLoadRuleOp = "STARTS_WITH"
	ZarazConfigReplaceResponseTriggersExcludeRulesZarazLoadRuleOpEndsWith           ZarazConfigReplaceResponseTriggersExcludeRulesZarazLoadRuleOp = "ENDS_WITH"
	ZarazConfigReplaceResponseTriggersExcludeRulesZarazLoadRuleOpMatchRegex         ZarazConfigReplaceResponseTriggersExcludeRulesZarazLoadRuleOp = "MATCH_REGEX"
	ZarazConfigReplaceResponseTriggersExcludeRulesZarazLoadRuleOpNotMatchRegex      ZarazConfigReplaceResponseTriggersExcludeRulesZarazLoadRuleOp = "NOT_MATCH_REGEX"
	ZarazConfigReplaceResponseTriggersExcludeRulesZarazLoadRuleOpGreaterThan        ZarazConfigReplaceResponseTriggersExcludeRulesZarazLoadRuleOp = "GREATER_THAN"
	ZarazConfigReplaceResponseTriggersExcludeRulesZarazLoadRuleOpGreaterThanOrEqual ZarazConfigReplaceResponseTriggersExcludeRulesZarazLoadRuleOp = "GREATER_THAN_OR_EQUAL"
	ZarazConfigReplaceResponseTriggersExcludeRulesZarazLoadRuleOpLessThan           ZarazConfigReplaceResponseTriggersExcludeRulesZarazLoadRuleOp = "LESS_THAN"
	ZarazConfigReplaceResponseTriggersExcludeRulesZarazLoadRuleOpLessThanOrEqual    ZarazConfigReplaceResponseTriggersExcludeRulesZarazLoadRuleOp = "LESS_THAN_OR_EQUAL"
)

type ZarazConfigReplaceResponseTriggersExcludeRulesZarazClickListenerRule struct {
	ID       string                                                                       `json:"id,required"`
	Action   ZarazConfigReplaceResponseTriggersExcludeRulesZarazClickListenerRuleAction   `json:"action,required"`
	Settings ZarazConfigReplaceResponseTriggersExcludeRulesZarazClickListenerRuleSettings `json:"settings,required"`
	JSON     zarazConfigReplaceResponseTriggersExcludeRulesZarazClickListenerRuleJSON     `json:"-"`
}

// zarazConfigReplaceResponseTriggersExcludeRulesZarazClickListenerRuleJSON
// contains the JSON metadata for the struct
// [ZarazConfigReplaceResponseTriggersExcludeRulesZarazClickListenerRule]
type zarazConfigReplaceResponseTriggersExcludeRulesZarazClickListenerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseTriggersExcludeRulesZarazClickListenerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigReplaceResponseTriggersExcludeRulesZarazClickListenerRule) implementsZarazConfigReplaceResponseTriggersExcludeRule() {
}

type ZarazConfigReplaceResponseTriggersExcludeRulesZarazClickListenerRuleAction string

const (
	ZarazConfigReplaceResponseTriggersExcludeRulesZarazClickListenerRuleActionClickListener ZarazConfigReplaceResponseTriggersExcludeRulesZarazClickListenerRuleAction = "clickListener"
)

type ZarazConfigReplaceResponseTriggersExcludeRulesZarazClickListenerRuleSettings struct {
	Selector    string                                                                           `json:"selector,required"`
	Type        ZarazConfigReplaceResponseTriggersExcludeRulesZarazClickListenerRuleSettingsType `json:"type,required"`
	WaitForTags int64                                                                            `json:"waitForTags,required"`
	JSON        zarazConfigReplaceResponseTriggersExcludeRulesZarazClickListenerRuleSettingsJSON `json:"-"`
}

// zarazConfigReplaceResponseTriggersExcludeRulesZarazClickListenerRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazConfigReplaceResponseTriggersExcludeRulesZarazClickListenerRuleSettings]
type zarazConfigReplaceResponseTriggersExcludeRulesZarazClickListenerRuleSettingsJSON struct {
	Selector    apijson.Field
	Type        apijson.Field
	WaitForTags apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseTriggersExcludeRulesZarazClickListenerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigReplaceResponseTriggersExcludeRulesZarazClickListenerRuleSettingsType string

const (
	ZarazConfigReplaceResponseTriggersExcludeRulesZarazClickListenerRuleSettingsTypeXpath ZarazConfigReplaceResponseTriggersExcludeRulesZarazClickListenerRuleSettingsType = "xpath"
	ZarazConfigReplaceResponseTriggersExcludeRulesZarazClickListenerRuleSettingsTypeCss   ZarazConfigReplaceResponseTriggersExcludeRulesZarazClickListenerRuleSettingsType = "css"
)

type ZarazConfigReplaceResponseTriggersExcludeRulesZarazTimerRule struct {
	ID       string                                                               `json:"id,required"`
	Action   ZarazConfigReplaceResponseTriggersExcludeRulesZarazTimerRuleAction   `json:"action,required"`
	Settings ZarazConfigReplaceResponseTriggersExcludeRulesZarazTimerRuleSettings `json:"settings,required"`
	JSON     zarazConfigReplaceResponseTriggersExcludeRulesZarazTimerRuleJSON     `json:"-"`
}

// zarazConfigReplaceResponseTriggersExcludeRulesZarazTimerRuleJSON contains the
// JSON metadata for the struct
// [ZarazConfigReplaceResponseTriggersExcludeRulesZarazTimerRule]
type zarazConfigReplaceResponseTriggersExcludeRulesZarazTimerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseTriggersExcludeRulesZarazTimerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigReplaceResponseTriggersExcludeRulesZarazTimerRule) implementsZarazConfigReplaceResponseTriggersExcludeRule() {
}

type ZarazConfigReplaceResponseTriggersExcludeRulesZarazTimerRuleAction string

const (
	ZarazConfigReplaceResponseTriggersExcludeRulesZarazTimerRuleActionTimer ZarazConfigReplaceResponseTriggersExcludeRulesZarazTimerRuleAction = "timer"
)

type ZarazConfigReplaceResponseTriggersExcludeRulesZarazTimerRuleSettings struct {
	Interval int64                                                                    `json:"interval,required"`
	Limit    int64                                                                    `json:"limit,required"`
	JSON     zarazConfigReplaceResponseTriggersExcludeRulesZarazTimerRuleSettingsJSON `json:"-"`
}

// zarazConfigReplaceResponseTriggersExcludeRulesZarazTimerRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazConfigReplaceResponseTriggersExcludeRulesZarazTimerRuleSettings]
type zarazConfigReplaceResponseTriggersExcludeRulesZarazTimerRuleSettingsJSON struct {
	Interval    apijson.Field
	Limit       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseTriggersExcludeRulesZarazTimerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigReplaceResponseTriggersExcludeRulesZarazFormSubmissionRule struct {
	ID       string                                                                        `json:"id,required"`
	Action   ZarazConfigReplaceResponseTriggersExcludeRulesZarazFormSubmissionRuleAction   `json:"action,required"`
	Settings ZarazConfigReplaceResponseTriggersExcludeRulesZarazFormSubmissionRuleSettings `json:"settings,required"`
	JSON     zarazConfigReplaceResponseTriggersExcludeRulesZarazFormSubmissionRuleJSON     `json:"-"`
}

// zarazConfigReplaceResponseTriggersExcludeRulesZarazFormSubmissionRuleJSON
// contains the JSON metadata for the struct
// [ZarazConfigReplaceResponseTriggersExcludeRulesZarazFormSubmissionRule]
type zarazConfigReplaceResponseTriggersExcludeRulesZarazFormSubmissionRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseTriggersExcludeRulesZarazFormSubmissionRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigReplaceResponseTriggersExcludeRulesZarazFormSubmissionRule) implementsZarazConfigReplaceResponseTriggersExcludeRule() {
}

type ZarazConfigReplaceResponseTriggersExcludeRulesZarazFormSubmissionRuleAction string

const (
	ZarazConfigReplaceResponseTriggersExcludeRulesZarazFormSubmissionRuleActionFormSubmission ZarazConfigReplaceResponseTriggersExcludeRulesZarazFormSubmissionRuleAction = "formSubmission"
)

type ZarazConfigReplaceResponseTriggersExcludeRulesZarazFormSubmissionRuleSettings struct {
	Selector string                                                                            `json:"selector,required"`
	Validate bool                                                                              `json:"validate,required"`
	JSON     zarazConfigReplaceResponseTriggersExcludeRulesZarazFormSubmissionRuleSettingsJSON `json:"-"`
}

// zarazConfigReplaceResponseTriggersExcludeRulesZarazFormSubmissionRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazConfigReplaceResponseTriggersExcludeRulesZarazFormSubmissionRuleSettings]
type zarazConfigReplaceResponseTriggersExcludeRulesZarazFormSubmissionRuleSettingsJSON struct {
	Selector    apijson.Field
	Validate    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseTriggersExcludeRulesZarazFormSubmissionRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigReplaceResponseTriggersExcludeRulesZarazVariableMatchRule struct {
	ID       string                                                                       `json:"id,required"`
	Action   ZarazConfigReplaceResponseTriggersExcludeRulesZarazVariableMatchRuleAction   `json:"action,required"`
	Settings ZarazConfigReplaceResponseTriggersExcludeRulesZarazVariableMatchRuleSettings `json:"settings,required"`
	JSON     zarazConfigReplaceResponseTriggersExcludeRulesZarazVariableMatchRuleJSON     `json:"-"`
}

// zarazConfigReplaceResponseTriggersExcludeRulesZarazVariableMatchRuleJSON
// contains the JSON metadata for the struct
// [ZarazConfigReplaceResponseTriggersExcludeRulesZarazVariableMatchRule]
type zarazConfigReplaceResponseTriggersExcludeRulesZarazVariableMatchRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseTriggersExcludeRulesZarazVariableMatchRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigReplaceResponseTriggersExcludeRulesZarazVariableMatchRule) implementsZarazConfigReplaceResponseTriggersExcludeRule() {
}

type ZarazConfigReplaceResponseTriggersExcludeRulesZarazVariableMatchRuleAction string

const (
	ZarazConfigReplaceResponseTriggersExcludeRulesZarazVariableMatchRuleActionVariableMatch ZarazConfigReplaceResponseTriggersExcludeRulesZarazVariableMatchRuleAction = "variableMatch"
)

type ZarazConfigReplaceResponseTriggersExcludeRulesZarazVariableMatchRuleSettings struct {
	Match    string                                                                           `json:"match,required"`
	Variable string                                                                           `json:"variable,required"`
	JSON     zarazConfigReplaceResponseTriggersExcludeRulesZarazVariableMatchRuleSettingsJSON `json:"-"`
}

// zarazConfigReplaceResponseTriggersExcludeRulesZarazVariableMatchRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazConfigReplaceResponseTriggersExcludeRulesZarazVariableMatchRuleSettings]
type zarazConfigReplaceResponseTriggersExcludeRulesZarazVariableMatchRuleSettingsJSON struct {
	Match       apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseTriggersExcludeRulesZarazVariableMatchRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigReplaceResponseTriggersExcludeRulesZarazScrollDepthRule struct {
	ID       string                                                                     `json:"id,required"`
	Action   ZarazConfigReplaceResponseTriggersExcludeRulesZarazScrollDepthRuleAction   `json:"action,required"`
	Settings ZarazConfigReplaceResponseTriggersExcludeRulesZarazScrollDepthRuleSettings `json:"settings,required"`
	JSON     zarazConfigReplaceResponseTriggersExcludeRulesZarazScrollDepthRuleJSON     `json:"-"`
}

// zarazConfigReplaceResponseTriggersExcludeRulesZarazScrollDepthRuleJSON contains
// the JSON metadata for the struct
// [ZarazConfigReplaceResponseTriggersExcludeRulesZarazScrollDepthRule]
type zarazConfigReplaceResponseTriggersExcludeRulesZarazScrollDepthRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseTriggersExcludeRulesZarazScrollDepthRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigReplaceResponseTriggersExcludeRulesZarazScrollDepthRule) implementsZarazConfigReplaceResponseTriggersExcludeRule() {
}

type ZarazConfigReplaceResponseTriggersExcludeRulesZarazScrollDepthRuleAction string

const (
	ZarazConfigReplaceResponseTriggersExcludeRulesZarazScrollDepthRuleActionScrollDepth ZarazConfigReplaceResponseTriggersExcludeRulesZarazScrollDepthRuleAction = "scrollDepth"
)

type ZarazConfigReplaceResponseTriggersExcludeRulesZarazScrollDepthRuleSettings struct {
	Positions string                                                                         `json:"positions,required"`
	JSON      zarazConfigReplaceResponseTriggersExcludeRulesZarazScrollDepthRuleSettingsJSON `json:"-"`
}

// zarazConfigReplaceResponseTriggersExcludeRulesZarazScrollDepthRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazConfigReplaceResponseTriggersExcludeRulesZarazScrollDepthRuleSettings]
type zarazConfigReplaceResponseTriggersExcludeRulesZarazScrollDepthRuleSettingsJSON struct {
	Positions   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseTriggersExcludeRulesZarazScrollDepthRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigReplaceResponseTriggersExcludeRulesZarazElementVisibilityRule struct {
	ID       string                                                                           `json:"id,required"`
	Action   ZarazConfigReplaceResponseTriggersExcludeRulesZarazElementVisibilityRuleAction   `json:"action,required"`
	Settings ZarazConfigReplaceResponseTriggersExcludeRulesZarazElementVisibilityRuleSettings `json:"settings,required"`
	JSON     zarazConfigReplaceResponseTriggersExcludeRulesZarazElementVisibilityRuleJSON     `json:"-"`
}

// zarazConfigReplaceResponseTriggersExcludeRulesZarazElementVisibilityRuleJSON
// contains the JSON metadata for the struct
// [ZarazConfigReplaceResponseTriggersExcludeRulesZarazElementVisibilityRule]
type zarazConfigReplaceResponseTriggersExcludeRulesZarazElementVisibilityRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseTriggersExcludeRulesZarazElementVisibilityRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigReplaceResponseTriggersExcludeRulesZarazElementVisibilityRule) implementsZarazConfigReplaceResponseTriggersExcludeRule() {
}

type ZarazConfigReplaceResponseTriggersExcludeRulesZarazElementVisibilityRuleAction string

const (
	ZarazConfigReplaceResponseTriggersExcludeRulesZarazElementVisibilityRuleActionElementVisibility ZarazConfigReplaceResponseTriggersExcludeRulesZarazElementVisibilityRuleAction = "elementVisibility"
)

type ZarazConfigReplaceResponseTriggersExcludeRulesZarazElementVisibilityRuleSettings struct {
	Selector string                                                                               `json:"selector,required"`
	JSON     zarazConfigReplaceResponseTriggersExcludeRulesZarazElementVisibilityRuleSettingsJSON `json:"-"`
}

// zarazConfigReplaceResponseTriggersExcludeRulesZarazElementVisibilityRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazConfigReplaceResponseTriggersExcludeRulesZarazElementVisibilityRuleSettings]
type zarazConfigReplaceResponseTriggersExcludeRulesZarazElementVisibilityRuleSettingsJSON struct {
	Selector    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseTriggersExcludeRulesZarazElementVisibilityRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [ZarazConfigReplaceResponseTriggersLoadRulesZarazLoadRule],
// [ZarazConfigReplaceResponseTriggersLoadRulesZarazClickListenerRule],
// [ZarazConfigReplaceResponseTriggersLoadRulesZarazTimerRule],
// [ZarazConfigReplaceResponseTriggersLoadRulesZarazFormSubmissionRule],
// [ZarazConfigReplaceResponseTriggersLoadRulesZarazVariableMatchRule],
// [ZarazConfigReplaceResponseTriggersLoadRulesZarazScrollDepthRule] or
// [ZarazConfigReplaceResponseTriggersLoadRulesZarazElementVisibilityRule].
type ZarazConfigReplaceResponseTriggersLoadRule interface {
	implementsZarazConfigReplaceResponseTriggersLoadRule()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZarazConfigReplaceResponseTriggersLoadRule)(nil)).Elem(), "")
}

type ZarazConfigReplaceResponseTriggersLoadRulesZarazLoadRule struct {
	ID    string                                                       `json:"id,required"`
	Match string                                                       `json:"match,required"`
	Op    ZarazConfigReplaceResponseTriggersLoadRulesZarazLoadRuleOp   `json:"op,required"`
	Value string                                                       `json:"value,required"`
	JSON  zarazConfigReplaceResponseTriggersLoadRulesZarazLoadRuleJSON `json:"-"`
}

// zarazConfigReplaceResponseTriggersLoadRulesZarazLoadRuleJSON contains the JSON
// metadata for the struct
// [ZarazConfigReplaceResponseTriggersLoadRulesZarazLoadRule]
type zarazConfigReplaceResponseTriggersLoadRulesZarazLoadRuleJSON struct {
	ID          apijson.Field
	Match       apijson.Field
	Op          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseTriggersLoadRulesZarazLoadRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigReplaceResponseTriggersLoadRulesZarazLoadRule) implementsZarazConfigReplaceResponseTriggersLoadRule() {
}

type ZarazConfigReplaceResponseTriggersLoadRulesZarazLoadRuleOp string

const (
	ZarazConfigReplaceResponseTriggersLoadRulesZarazLoadRuleOpContains           ZarazConfigReplaceResponseTriggersLoadRulesZarazLoadRuleOp = "CONTAINS"
	ZarazConfigReplaceResponseTriggersLoadRulesZarazLoadRuleOpEquals             ZarazConfigReplaceResponseTriggersLoadRulesZarazLoadRuleOp = "EQUALS"
	ZarazConfigReplaceResponseTriggersLoadRulesZarazLoadRuleOpStartsWith         ZarazConfigReplaceResponseTriggersLoadRulesZarazLoadRuleOp = "STARTS_WITH"
	ZarazConfigReplaceResponseTriggersLoadRulesZarazLoadRuleOpEndsWith           ZarazConfigReplaceResponseTriggersLoadRulesZarazLoadRuleOp = "ENDS_WITH"
	ZarazConfigReplaceResponseTriggersLoadRulesZarazLoadRuleOpMatchRegex         ZarazConfigReplaceResponseTriggersLoadRulesZarazLoadRuleOp = "MATCH_REGEX"
	ZarazConfigReplaceResponseTriggersLoadRulesZarazLoadRuleOpNotMatchRegex      ZarazConfigReplaceResponseTriggersLoadRulesZarazLoadRuleOp = "NOT_MATCH_REGEX"
	ZarazConfigReplaceResponseTriggersLoadRulesZarazLoadRuleOpGreaterThan        ZarazConfigReplaceResponseTriggersLoadRulesZarazLoadRuleOp = "GREATER_THAN"
	ZarazConfigReplaceResponseTriggersLoadRulesZarazLoadRuleOpGreaterThanOrEqual ZarazConfigReplaceResponseTriggersLoadRulesZarazLoadRuleOp = "GREATER_THAN_OR_EQUAL"
	ZarazConfigReplaceResponseTriggersLoadRulesZarazLoadRuleOpLessThan           ZarazConfigReplaceResponseTriggersLoadRulesZarazLoadRuleOp = "LESS_THAN"
	ZarazConfigReplaceResponseTriggersLoadRulesZarazLoadRuleOpLessThanOrEqual    ZarazConfigReplaceResponseTriggersLoadRulesZarazLoadRuleOp = "LESS_THAN_OR_EQUAL"
)

type ZarazConfigReplaceResponseTriggersLoadRulesZarazClickListenerRule struct {
	ID       string                                                                    `json:"id,required"`
	Action   ZarazConfigReplaceResponseTriggersLoadRulesZarazClickListenerRuleAction   `json:"action,required"`
	Settings ZarazConfigReplaceResponseTriggersLoadRulesZarazClickListenerRuleSettings `json:"settings,required"`
	JSON     zarazConfigReplaceResponseTriggersLoadRulesZarazClickListenerRuleJSON     `json:"-"`
}

// zarazConfigReplaceResponseTriggersLoadRulesZarazClickListenerRuleJSON contains
// the JSON metadata for the struct
// [ZarazConfigReplaceResponseTriggersLoadRulesZarazClickListenerRule]
type zarazConfigReplaceResponseTriggersLoadRulesZarazClickListenerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseTriggersLoadRulesZarazClickListenerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigReplaceResponseTriggersLoadRulesZarazClickListenerRule) implementsZarazConfigReplaceResponseTriggersLoadRule() {
}

type ZarazConfigReplaceResponseTriggersLoadRulesZarazClickListenerRuleAction string

const (
	ZarazConfigReplaceResponseTriggersLoadRulesZarazClickListenerRuleActionClickListener ZarazConfigReplaceResponseTriggersLoadRulesZarazClickListenerRuleAction = "clickListener"
)

type ZarazConfigReplaceResponseTriggersLoadRulesZarazClickListenerRuleSettings struct {
	Selector    string                                                                        `json:"selector,required"`
	Type        ZarazConfigReplaceResponseTriggersLoadRulesZarazClickListenerRuleSettingsType `json:"type,required"`
	WaitForTags int64                                                                         `json:"waitForTags,required"`
	JSON        zarazConfigReplaceResponseTriggersLoadRulesZarazClickListenerRuleSettingsJSON `json:"-"`
}

// zarazConfigReplaceResponseTriggersLoadRulesZarazClickListenerRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazConfigReplaceResponseTriggersLoadRulesZarazClickListenerRuleSettings]
type zarazConfigReplaceResponseTriggersLoadRulesZarazClickListenerRuleSettingsJSON struct {
	Selector    apijson.Field
	Type        apijson.Field
	WaitForTags apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseTriggersLoadRulesZarazClickListenerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigReplaceResponseTriggersLoadRulesZarazClickListenerRuleSettingsType string

const (
	ZarazConfigReplaceResponseTriggersLoadRulesZarazClickListenerRuleSettingsTypeXpath ZarazConfigReplaceResponseTriggersLoadRulesZarazClickListenerRuleSettingsType = "xpath"
	ZarazConfigReplaceResponseTriggersLoadRulesZarazClickListenerRuleSettingsTypeCss   ZarazConfigReplaceResponseTriggersLoadRulesZarazClickListenerRuleSettingsType = "css"
)

type ZarazConfigReplaceResponseTriggersLoadRulesZarazTimerRule struct {
	ID       string                                                            `json:"id,required"`
	Action   ZarazConfigReplaceResponseTriggersLoadRulesZarazTimerRuleAction   `json:"action,required"`
	Settings ZarazConfigReplaceResponseTriggersLoadRulesZarazTimerRuleSettings `json:"settings,required"`
	JSON     zarazConfigReplaceResponseTriggersLoadRulesZarazTimerRuleJSON     `json:"-"`
}

// zarazConfigReplaceResponseTriggersLoadRulesZarazTimerRuleJSON contains the JSON
// metadata for the struct
// [ZarazConfigReplaceResponseTriggersLoadRulesZarazTimerRule]
type zarazConfigReplaceResponseTriggersLoadRulesZarazTimerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseTriggersLoadRulesZarazTimerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigReplaceResponseTriggersLoadRulesZarazTimerRule) implementsZarazConfigReplaceResponseTriggersLoadRule() {
}

type ZarazConfigReplaceResponseTriggersLoadRulesZarazTimerRuleAction string

const (
	ZarazConfigReplaceResponseTriggersLoadRulesZarazTimerRuleActionTimer ZarazConfigReplaceResponseTriggersLoadRulesZarazTimerRuleAction = "timer"
)

type ZarazConfigReplaceResponseTriggersLoadRulesZarazTimerRuleSettings struct {
	Interval int64                                                                 `json:"interval,required"`
	Limit    int64                                                                 `json:"limit,required"`
	JSON     zarazConfigReplaceResponseTriggersLoadRulesZarazTimerRuleSettingsJSON `json:"-"`
}

// zarazConfigReplaceResponseTriggersLoadRulesZarazTimerRuleSettingsJSON contains
// the JSON metadata for the struct
// [ZarazConfigReplaceResponseTriggersLoadRulesZarazTimerRuleSettings]
type zarazConfigReplaceResponseTriggersLoadRulesZarazTimerRuleSettingsJSON struct {
	Interval    apijson.Field
	Limit       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseTriggersLoadRulesZarazTimerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigReplaceResponseTriggersLoadRulesZarazFormSubmissionRule struct {
	ID       string                                                                     `json:"id,required"`
	Action   ZarazConfigReplaceResponseTriggersLoadRulesZarazFormSubmissionRuleAction   `json:"action,required"`
	Settings ZarazConfigReplaceResponseTriggersLoadRulesZarazFormSubmissionRuleSettings `json:"settings,required"`
	JSON     zarazConfigReplaceResponseTriggersLoadRulesZarazFormSubmissionRuleJSON     `json:"-"`
}

// zarazConfigReplaceResponseTriggersLoadRulesZarazFormSubmissionRuleJSON contains
// the JSON metadata for the struct
// [ZarazConfigReplaceResponseTriggersLoadRulesZarazFormSubmissionRule]
type zarazConfigReplaceResponseTriggersLoadRulesZarazFormSubmissionRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseTriggersLoadRulesZarazFormSubmissionRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigReplaceResponseTriggersLoadRulesZarazFormSubmissionRule) implementsZarazConfigReplaceResponseTriggersLoadRule() {
}

type ZarazConfigReplaceResponseTriggersLoadRulesZarazFormSubmissionRuleAction string

const (
	ZarazConfigReplaceResponseTriggersLoadRulesZarazFormSubmissionRuleActionFormSubmission ZarazConfigReplaceResponseTriggersLoadRulesZarazFormSubmissionRuleAction = "formSubmission"
)

type ZarazConfigReplaceResponseTriggersLoadRulesZarazFormSubmissionRuleSettings struct {
	Selector string                                                                         `json:"selector,required"`
	Validate bool                                                                           `json:"validate,required"`
	JSON     zarazConfigReplaceResponseTriggersLoadRulesZarazFormSubmissionRuleSettingsJSON `json:"-"`
}

// zarazConfigReplaceResponseTriggersLoadRulesZarazFormSubmissionRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazConfigReplaceResponseTriggersLoadRulesZarazFormSubmissionRuleSettings]
type zarazConfigReplaceResponseTriggersLoadRulesZarazFormSubmissionRuleSettingsJSON struct {
	Selector    apijson.Field
	Validate    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseTriggersLoadRulesZarazFormSubmissionRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigReplaceResponseTriggersLoadRulesZarazVariableMatchRule struct {
	ID       string                                                                    `json:"id,required"`
	Action   ZarazConfigReplaceResponseTriggersLoadRulesZarazVariableMatchRuleAction   `json:"action,required"`
	Settings ZarazConfigReplaceResponseTriggersLoadRulesZarazVariableMatchRuleSettings `json:"settings,required"`
	JSON     zarazConfigReplaceResponseTriggersLoadRulesZarazVariableMatchRuleJSON     `json:"-"`
}

// zarazConfigReplaceResponseTriggersLoadRulesZarazVariableMatchRuleJSON contains
// the JSON metadata for the struct
// [ZarazConfigReplaceResponseTriggersLoadRulesZarazVariableMatchRule]
type zarazConfigReplaceResponseTriggersLoadRulesZarazVariableMatchRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseTriggersLoadRulesZarazVariableMatchRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigReplaceResponseTriggersLoadRulesZarazVariableMatchRule) implementsZarazConfigReplaceResponseTriggersLoadRule() {
}

type ZarazConfigReplaceResponseTriggersLoadRulesZarazVariableMatchRuleAction string

const (
	ZarazConfigReplaceResponseTriggersLoadRulesZarazVariableMatchRuleActionVariableMatch ZarazConfigReplaceResponseTriggersLoadRulesZarazVariableMatchRuleAction = "variableMatch"
)

type ZarazConfigReplaceResponseTriggersLoadRulesZarazVariableMatchRuleSettings struct {
	Match    string                                                                        `json:"match,required"`
	Variable string                                                                        `json:"variable,required"`
	JSON     zarazConfigReplaceResponseTriggersLoadRulesZarazVariableMatchRuleSettingsJSON `json:"-"`
}

// zarazConfigReplaceResponseTriggersLoadRulesZarazVariableMatchRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazConfigReplaceResponseTriggersLoadRulesZarazVariableMatchRuleSettings]
type zarazConfigReplaceResponseTriggersLoadRulesZarazVariableMatchRuleSettingsJSON struct {
	Match       apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseTriggersLoadRulesZarazVariableMatchRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigReplaceResponseTriggersLoadRulesZarazScrollDepthRule struct {
	ID       string                                                                  `json:"id,required"`
	Action   ZarazConfigReplaceResponseTriggersLoadRulesZarazScrollDepthRuleAction   `json:"action,required"`
	Settings ZarazConfigReplaceResponseTriggersLoadRulesZarazScrollDepthRuleSettings `json:"settings,required"`
	JSON     zarazConfigReplaceResponseTriggersLoadRulesZarazScrollDepthRuleJSON     `json:"-"`
}

// zarazConfigReplaceResponseTriggersLoadRulesZarazScrollDepthRuleJSON contains the
// JSON metadata for the struct
// [ZarazConfigReplaceResponseTriggersLoadRulesZarazScrollDepthRule]
type zarazConfigReplaceResponseTriggersLoadRulesZarazScrollDepthRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseTriggersLoadRulesZarazScrollDepthRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigReplaceResponseTriggersLoadRulesZarazScrollDepthRule) implementsZarazConfigReplaceResponseTriggersLoadRule() {
}

type ZarazConfigReplaceResponseTriggersLoadRulesZarazScrollDepthRuleAction string

const (
	ZarazConfigReplaceResponseTriggersLoadRulesZarazScrollDepthRuleActionScrollDepth ZarazConfigReplaceResponseTriggersLoadRulesZarazScrollDepthRuleAction = "scrollDepth"
)

type ZarazConfigReplaceResponseTriggersLoadRulesZarazScrollDepthRuleSettings struct {
	Positions string                                                                      `json:"positions,required"`
	JSON      zarazConfigReplaceResponseTriggersLoadRulesZarazScrollDepthRuleSettingsJSON `json:"-"`
}

// zarazConfigReplaceResponseTriggersLoadRulesZarazScrollDepthRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazConfigReplaceResponseTriggersLoadRulesZarazScrollDepthRuleSettings]
type zarazConfigReplaceResponseTriggersLoadRulesZarazScrollDepthRuleSettingsJSON struct {
	Positions   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseTriggersLoadRulesZarazScrollDepthRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigReplaceResponseTriggersLoadRulesZarazElementVisibilityRule struct {
	ID       string                                                                        `json:"id,required"`
	Action   ZarazConfigReplaceResponseTriggersLoadRulesZarazElementVisibilityRuleAction   `json:"action,required"`
	Settings ZarazConfigReplaceResponseTriggersLoadRulesZarazElementVisibilityRuleSettings `json:"settings,required"`
	JSON     zarazConfigReplaceResponseTriggersLoadRulesZarazElementVisibilityRuleJSON     `json:"-"`
}

// zarazConfigReplaceResponseTriggersLoadRulesZarazElementVisibilityRuleJSON
// contains the JSON metadata for the struct
// [ZarazConfigReplaceResponseTriggersLoadRulesZarazElementVisibilityRule]
type zarazConfigReplaceResponseTriggersLoadRulesZarazElementVisibilityRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseTriggersLoadRulesZarazElementVisibilityRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigReplaceResponseTriggersLoadRulesZarazElementVisibilityRule) implementsZarazConfigReplaceResponseTriggersLoadRule() {
}

type ZarazConfigReplaceResponseTriggersLoadRulesZarazElementVisibilityRuleAction string

const (
	ZarazConfigReplaceResponseTriggersLoadRulesZarazElementVisibilityRuleActionElementVisibility ZarazConfigReplaceResponseTriggersLoadRulesZarazElementVisibilityRuleAction = "elementVisibility"
)

type ZarazConfigReplaceResponseTriggersLoadRulesZarazElementVisibilityRuleSettings struct {
	Selector string                                                                            `json:"selector,required"`
	JSON     zarazConfigReplaceResponseTriggersLoadRulesZarazElementVisibilityRuleSettingsJSON `json:"-"`
}

// zarazConfigReplaceResponseTriggersLoadRulesZarazElementVisibilityRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazConfigReplaceResponseTriggersLoadRulesZarazElementVisibilityRuleSettings]
type zarazConfigReplaceResponseTriggersLoadRulesZarazElementVisibilityRuleSettingsJSON struct {
	Selector    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseTriggersLoadRulesZarazElementVisibilityRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigReplaceResponseTriggersSystem string

const (
	ZarazConfigReplaceResponseTriggersSystemPageload ZarazConfigReplaceResponseTriggersSystem = "pageload"
)

// Union satisfied by [ZarazConfigReplaceResponseVariablesObject] or
// [ZarazConfigReplaceResponseVariablesObject].
type ZarazConfigReplaceResponseVariable interface {
	implementsZarazConfigReplaceResponseVariable()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZarazConfigReplaceResponseVariable)(nil)).Elem(), "")
}

type ZarazConfigReplaceResponseVariablesObject struct {
	Name  string                                        `json:"name,required"`
	Type  ZarazConfigReplaceResponseVariablesObjectType `json:"type,required"`
	Value string                                        `json:"value,required"`
	JSON  zarazConfigReplaceResponseVariablesObjectJSON `json:"-"`
}

// zarazConfigReplaceResponseVariablesObjectJSON contains the JSON metadata for the
// struct [ZarazConfigReplaceResponseVariablesObject]
type zarazConfigReplaceResponseVariablesObjectJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseVariablesObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigReplaceResponseVariablesObject) implementsZarazConfigReplaceResponseVariable() {}

type ZarazConfigReplaceResponseVariablesObjectType string

const (
	ZarazConfigReplaceResponseVariablesObjectTypeString ZarazConfigReplaceResponseVariablesObjectType = "string"
	ZarazConfigReplaceResponseVariablesObjectTypeSecret ZarazConfigReplaceResponseVariablesObjectType = "secret"
)

// Consent management configuration.
type ZarazConfigReplaceResponseConsent struct {
	Enabled                bool                                                    `json:"enabled,required"`
	ButtonTextTranslations ZarazConfigReplaceResponseConsentButtonTextTranslations `json:"buttonTextTranslations"`
	CompanyEmail           string                                                  `json:"companyEmail"`
	CompanyName            string                                                  `json:"companyName"`
	CompanyStreetAddress   string                                                  `json:"companyStreetAddress"`
	ConsentModalIntroHTML  string                                                  `json:"consentModalIntroHTML"`
	// Object where keys are language codes
	ConsentModalIntroHTMLWithTranslations map[string]string `json:"consentModalIntroHTMLWithTranslations"`
	CookieName                            string            `json:"cookieName"`
	CustomCss                             string            `json:"customCSS"`
	CustomIntroDisclaimerDismissed        bool              `json:"customIntroDisclaimerDismissed"`
	DefaultLanguage                       string            `json:"defaultLanguage"`
	HideModal                             bool              `json:"hideModal"`
	// Object where keys are purpose alpha-numeric IDs
	Purposes map[string]ZarazConfigReplaceResponseConsentPurpose `json:"purposes"`
	// Object where keys are purpose alpha-numeric IDs
	PurposesWithTranslations map[string]ZarazConfigReplaceResponseConsentPurposesWithTranslation `json:"purposesWithTranslations"`
	JSON                     zarazConfigReplaceResponseConsentJSON                               `json:"-"`
}

// zarazConfigReplaceResponseConsentJSON contains the JSON metadata for the struct
// [ZarazConfigReplaceResponseConsent]
type zarazConfigReplaceResponseConsentJSON struct {
	Enabled                               apijson.Field
	ButtonTextTranslations                apijson.Field
	CompanyEmail                          apijson.Field
	CompanyName                           apijson.Field
	CompanyStreetAddress                  apijson.Field
	ConsentModalIntroHTML                 apijson.Field
	ConsentModalIntroHTMLWithTranslations apijson.Field
	CookieName                            apijson.Field
	CustomCss                             apijson.Field
	CustomIntroDisclaimerDismissed        apijson.Field
	DefaultLanguage                       apijson.Field
	HideModal                             apijson.Field
	Purposes                              apijson.Field
	PurposesWithTranslations              apijson.Field
	raw                                   string
	ExtraFields                           map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseConsent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigReplaceResponseConsentButtonTextTranslations struct {
	// Object where keys are language codes
	AcceptAll map[string]string `json:"accept_all,required"`
	// Object where keys are language codes
	ConfirmMyChoices map[string]string `json:"confirm_my_choices,required"`
	// Object where keys are language codes
	RejectAll map[string]string                                           `json:"reject_all,required"`
	JSON      zarazConfigReplaceResponseConsentButtonTextTranslationsJSON `json:"-"`
}

// zarazConfigReplaceResponseConsentButtonTextTranslationsJSON contains the JSON
// metadata for the struct
// [ZarazConfigReplaceResponseConsentButtonTextTranslations]
type zarazConfigReplaceResponseConsentButtonTextTranslationsJSON struct {
	AcceptAll        apijson.Field
	ConfirmMyChoices apijson.Field
	RejectAll        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseConsentButtonTextTranslations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigReplaceResponseConsentPurpose struct {
	Description string                                       `json:"description,required"`
	Name        string                                       `json:"name,required"`
	JSON        zarazConfigReplaceResponseConsentPurposeJSON `json:"-"`
}

// zarazConfigReplaceResponseConsentPurposeJSON contains the JSON metadata for the
// struct [ZarazConfigReplaceResponseConsentPurpose]
type zarazConfigReplaceResponseConsentPurposeJSON struct {
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseConsentPurpose) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigReplaceResponseConsentPurposesWithTranslation struct {
	// Object where keys are language codes
	Description map[string]string `json:"description,required"`
	// Object where keys are language codes
	Name  map[string]string                                            `json:"name,required"`
	Order int64                                                        `json:"order,required"`
	JSON  zarazConfigReplaceResponseConsentPurposesWithTranslationJSON `json:"-"`
}

// zarazConfigReplaceResponseConsentPurposesWithTranslationJSON contains the JSON
// metadata for the struct
// [ZarazConfigReplaceResponseConsentPurposesWithTranslation]
type zarazConfigReplaceResponseConsentPurposesWithTranslationJSON struct {
	Description apijson.Field
	Name        apijson.Field
	Order       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseConsentPurposesWithTranslation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigGetResponseEnvelope struct {
	Errors   []ZarazConfigGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZarazConfigGetResponseEnvelopeMessages `json:"messages,required"`
	// Zaraz configuration
	Result ZarazConfigGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success bool                               `json:"success,required"`
	JSON    zarazConfigGetResponseEnvelopeJSON `json:"-"`
}

// zarazConfigGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZarazConfigGetResponseEnvelope]
type zarazConfigGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigGetResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    zarazConfigGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zarazConfigGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZarazConfigGetResponseEnvelopeErrors]
type zarazConfigGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigGetResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    zarazConfigGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zarazConfigGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ZarazConfigGetResponseEnvelopeMessages]
type zarazConfigGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigReplaceParams struct {
	// Data layer compatibility mode enabled.
	DataLayer param.Field[bool] `json:"dataLayer,required"`
	// The key for Zaraz debug mode.
	DebugKey param.Field[string] `json:"debugKey,required"`
	// General Zaraz settings.
	Settings param.Field[ZarazConfigReplaceParamsSettings] `json:"settings,required"`
	// Tools set up under Zaraz configuration, where key is the alpha-numeric tool ID
	// and value is the tool configuration object.
	Tools param.Field[map[string]ZarazConfigReplaceParamsTools] `json:"tools,required"`
	// Triggers set up under Zaraz configuration, where key is the trigger
	// alpha-numeric ID and value is the trigger configuration.
	Triggers param.Field[map[string]ZarazConfigReplaceParamsTriggers] `json:"triggers,required"`
	// Variables set up under Zaraz configuration, where key is the variable
	// alpha-numeric ID and value is the variable configuration. Values of variables of
	// type secret are not included.
	Variables param.Field[map[string]ZarazConfigReplaceParamsVariables] `json:"variables,required"`
	// Zaraz internal version of the config.
	ZarazVersion param.Field[int64] `json:"zarazVersion,required"`
	// Consent management configuration.
	Consent param.Field[ZarazConfigReplaceParamsConsent] `json:"consent"`
	// Single Page Application support enabled.
	HistoryChange param.Field[bool] `json:"historyChange"`
}

func (r ZarazConfigReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// General Zaraz settings.
type ZarazConfigReplaceParamsSettings struct {
	// Automatic injection of Zaraz scripts enabled.
	AutoInjectScript param.Field[bool] `json:"autoInjectScript,required"`
	// Details of the worker that receives and edits Zaraz Context object.
	ContextEnricher param.Field[ZarazConfigReplaceParamsSettingsContextEnricher] `json:"contextEnricher"`
	// The domain Zaraz will use for writing and reading its cookies.
	CookieDomain param.Field[string] `json:"cookieDomain"`
	// Ecommerce API enabled.
	Ecommerce param.Field[bool] `json:"ecommerce"`
	// Custom endpoint for server-side track events.
	EventsAPIPath param.Field[string] `json:"eventsApiPath"`
	// Hiding external referrer URL enabled.
	HideExternalReferer param.Field[bool] `json:"hideExternalReferer"`
	// Trimming IP address enabled.
	HideIPAddress param.Field[bool] `json:"hideIPAddress"`
	// Removing URL query params enabled.
	HideQueryParams param.Field[bool] `json:"hideQueryParams"`
	// Removing sensitive data from User Aagent string enabled.
	HideUserAgent param.Field[bool] `json:"hideUserAgent"`
	// Custom endpoint for Zaraz init script.
	InitPath param.Field[string] `json:"initPath"`
	// Injection of Zaraz scripts into iframes enabled.
	InjectIframes param.Field[bool] `json:"injectIframes"`
	// Custom path for Managed Components server functionalities.
	McRootPath param.Field[string] `json:"mcRootPath"`
	// Custom endpoint for Zaraz main script.
	ScriptPath param.Field[string] `json:"scriptPath"`
	// Custom endpoint for Zaraz tracking requests.
	TrackPath param.Field[string] `json:"trackPath"`
}

func (r ZarazConfigReplaceParamsSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Details of the worker that receives and edits Zaraz Context object.
type ZarazConfigReplaceParamsSettingsContextEnricher struct {
	EscapedWorkerName param.Field[string] `json:"escapedWorkerName,required"`
	WorkerTag         param.Field[string] `json:"workerTag,required"`
}

func (r ZarazConfigReplaceParamsSettingsContextEnricher) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [ZarazConfigReplaceParamsToolsZarazLegacyTool],
// [ZarazConfigReplaceParamsToolsZarazManagedComponent],
// [ZarazConfigReplaceParamsToolsZarazCustomManagedComponent].
type ZarazConfigReplaceParamsTools interface {
	implementsZarazConfigReplaceParamsTools()
}

type ZarazConfigReplaceParamsToolsZarazLegacyTool struct {
	// List of blocking trigger IDs
	BlockingTriggers param.Field[[]string] `json:"blockingTriggers,required"`
	// Default fields for tool's actions
	DefaultFields param.Field[map[string]ZarazConfigReplaceParamsToolsZarazLegacyToolDefaultFields] `json:"defaultFields,required"`
	// Whether tool is enabled
	Enabled param.Field[bool] `json:"enabled,required"`
	// Tool's internal name
	Library param.Field[string] `json:"library,required"`
	// Tool's name defined by the user
	Name param.Field[string] `json:"name,required"`
	// List of actions configured on a tool
	NeoEvents param.Field[[]ZarazConfigReplaceParamsToolsZarazLegacyToolNeoEvent] `json:"neoEvents,required"`
	Type      param.Field[ZarazConfigReplaceParamsToolsZarazLegacyToolType]       `json:"type,required"`
	// Default consent purpose ID
	DefaultPurpose param.Field[string] `json:"defaultPurpose"`
}

func (r ZarazConfigReplaceParamsToolsZarazLegacyTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZarazConfigReplaceParamsToolsZarazLegacyTool) implementsZarazConfigReplaceParamsTools() {}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type ZarazConfigReplaceParamsToolsZarazLegacyToolDefaultFields interface {
	ImplementsZarazConfigReplaceParamsToolsZarazLegacyToolDefaultFields()
}

type ZarazConfigReplaceParamsToolsZarazLegacyToolNeoEvent struct {
	// List of blocking triggers IDs
	BlockingTriggers param.Field[[]string] `json:"blockingTriggers,required"`
	// Event payload
	Data param.Field[interface{}] `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers param.Field[[]string] `json:"firingTriggers,required"`
}

func (r ZarazConfigReplaceParamsToolsZarazLegacyToolNeoEvent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigReplaceParamsToolsZarazLegacyToolType string

const (
	ZarazConfigReplaceParamsToolsZarazLegacyToolTypeLibrary ZarazConfigReplaceParamsToolsZarazLegacyToolType = "library"
)

type ZarazConfigReplaceParamsToolsZarazManagedComponent struct {
	// List of blocking trigger IDs
	BlockingTriggers param.Field[[]string] `json:"blockingTriggers,required"`
	// Tool's internal name
	Component param.Field[string] `json:"component,required"`
	// Default fields for tool's actions
	DefaultFields param.Field[map[string]ZarazConfigReplaceParamsToolsZarazManagedComponentDefaultFields] `json:"defaultFields,required"`
	// Whether tool is enabled
	Enabled param.Field[bool] `json:"enabled,required"`
	// Tool's name defined by the user
	Name param.Field[string] `json:"name,required"`
	// List of permissions granted to the component
	Permissions param.Field[[]string] `json:"permissions,required"`
	// Tool's settings
	Settings param.Field[map[string]ZarazConfigReplaceParamsToolsZarazManagedComponentSettings] `json:"settings,required"`
	Type     param.Field[ZarazConfigReplaceParamsToolsZarazManagedComponentType]                `json:"type,required"`
	// Actions configured on a tool. Either this or neoEvents field is required.
	Actions param.Field[map[string]ZarazConfigReplaceParamsToolsZarazManagedComponentActions] `json:"actions"`
	// Default consent purpose ID
	DefaultPurpose param.Field[string] `json:"defaultPurpose"`
	// DEPRECATED - List of actions configured on a tool. Either this or actions field
	// is required. If both are present, actions field will take precedence.
	NeoEvents param.Field[[]ZarazConfigReplaceParamsToolsZarazManagedComponentNeoEvent] `json:"neoEvents"`
}

func (r ZarazConfigReplaceParamsToolsZarazManagedComponent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZarazConfigReplaceParamsToolsZarazManagedComponent) implementsZarazConfigReplaceParamsTools() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type ZarazConfigReplaceParamsToolsZarazManagedComponentDefaultFields interface {
	ImplementsZarazConfigReplaceParamsToolsZarazManagedComponentDefaultFields()
}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type ZarazConfigReplaceParamsToolsZarazManagedComponentSettings interface {
	ImplementsZarazConfigReplaceParamsToolsZarazManagedComponentSettings()
}

type ZarazConfigReplaceParamsToolsZarazManagedComponentType string

const (
	ZarazConfigReplaceParamsToolsZarazManagedComponentTypeComponent ZarazConfigReplaceParamsToolsZarazManagedComponentType = "component"
)

type ZarazConfigReplaceParamsToolsZarazManagedComponentActions struct {
	// Tool event type
	ActionType param.Field[string] `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers param.Field[[]string] `json:"blockingTriggers,required"`
	// Event payload
	Data param.Field[interface{}] `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers param.Field[[]string] `json:"firingTriggers,required"`
}

func (r ZarazConfigReplaceParamsToolsZarazManagedComponentActions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigReplaceParamsToolsZarazManagedComponentNeoEvent struct {
	// Tool event type
	ActionType param.Field[string] `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers param.Field[[]string] `json:"blockingTriggers,required"`
	// Event payload
	Data param.Field[interface{}] `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers param.Field[[]string] `json:"firingTriggers,required"`
}

func (r ZarazConfigReplaceParamsToolsZarazManagedComponentNeoEvent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigReplaceParamsToolsZarazCustomManagedComponent struct {
	// List of blocking trigger IDs
	BlockingTriggers param.Field[[]string] `json:"blockingTriggers,required"`
	// Tool's internal name
	Component param.Field[string] `json:"component,required"`
	// Default fields for tool's actions
	DefaultFields param.Field[map[string]ZarazConfigReplaceParamsToolsZarazCustomManagedComponentDefaultFields] `json:"defaultFields,required"`
	// Whether tool is enabled
	Enabled param.Field[bool] `json:"enabled,required"`
	// Tool's name defined by the user
	Name param.Field[string] `json:"name,required"`
	// List of permissions granted to the component
	Permissions param.Field[[]string] `json:"permissions,required"`
	// Tool's settings
	Settings param.Field[map[string]ZarazConfigReplaceParamsToolsZarazCustomManagedComponentSettings] `json:"settings,required"`
	Type     param.Field[ZarazConfigReplaceParamsToolsZarazCustomManagedComponentType]                `json:"type,required"`
	// Cloudflare worker that acts as a managed component
	Worker param.Field[ZarazConfigReplaceParamsToolsZarazCustomManagedComponentWorker] `json:"worker,required"`
	// Actions configured on a tool. Either this or neoEvents field is required.
	Actions param.Field[map[string]ZarazConfigReplaceParamsToolsZarazCustomManagedComponentActions] `json:"actions"`
	// Default consent purpose ID
	DefaultPurpose param.Field[string] `json:"defaultPurpose"`
	// DEPRECATED - List of actions configured on a tool. Either this or actions field
	// is required. If both are present, actions field will take precedence.
	NeoEvents param.Field[[]ZarazConfigReplaceParamsToolsZarazCustomManagedComponentNeoEvent] `json:"neoEvents"`
}

func (r ZarazConfigReplaceParamsToolsZarazCustomManagedComponent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZarazConfigReplaceParamsToolsZarazCustomManagedComponent) implementsZarazConfigReplaceParamsTools() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type ZarazConfigReplaceParamsToolsZarazCustomManagedComponentDefaultFields interface {
	ImplementsZarazConfigReplaceParamsToolsZarazCustomManagedComponentDefaultFields()
}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type ZarazConfigReplaceParamsToolsZarazCustomManagedComponentSettings interface {
	ImplementsZarazConfigReplaceParamsToolsZarazCustomManagedComponentSettings()
}

type ZarazConfigReplaceParamsToolsZarazCustomManagedComponentType string

const (
	ZarazConfigReplaceParamsToolsZarazCustomManagedComponentTypeCustomMc ZarazConfigReplaceParamsToolsZarazCustomManagedComponentType = "custom-mc"
)

// Cloudflare worker that acts as a managed component
type ZarazConfigReplaceParamsToolsZarazCustomManagedComponentWorker struct {
	EscapedWorkerName param.Field[string] `json:"escapedWorkerName,required"`
	WorkerTag         param.Field[string] `json:"workerTag,required"`
}

func (r ZarazConfigReplaceParamsToolsZarazCustomManagedComponentWorker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigReplaceParamsToolsZarazCustomManagedComponentActions struct {
	// Tool event type
	ActionType param.Field[string] `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers param.Field[[]string] `json:"blockingTriggers,required"`
	// Event payload
	Data param.Field[interface{}] `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers param.Field[[]string] `json:"firingTriggers,required"`
}

func (r ZarazConfigReplaceParamsToolsZarazCustomManagedComponentActions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigReplaceParamsToolsZarazCustomManagedComponentNeoEvent struct {
	// Tool event type
	ActionType param.Field[string] `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers param.Field[[]string] `json:"blockingTriggers,required"`
	// Event payload
	Data param.Field[interface{}] `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers param.Field[[]string] `json:"firingTriggers,required"`
}

func (r ZarazConfigReplaceParamsToolsZarazCustomManagedComponentNeoEvent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigReplaceParamsTriggers struct {
	// Rules defining when the trigger is not fired.
	ExcludeRules param.Field[[]ZarazConfigReplaceParamsTriggersExcludeRule] `json:"excludeRules,required"`
	// Rules defining when the trigger is fired.
	LoadRules param.Field[[]ZarazConfigReplaceParamsTriggersLoadRule] `json:"loadRules,required"`
	// Trigger name.
	Name param.Field[string] `json:"name,required"`
	// Trigger description.
	Description param.Field[string]                                 `json:"description"`
	System      param.Field[ZarazConfigReplaceParamsTriggersSystem] `json:"system"`
}

func (r ZarazConfigReplaceParamsTriggers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [ZarazConfigReplaceParamsTriggersExcludeRulesZarazLoadRule],
// [ZarazConfigReplaceParamsTriggersExcludeRulesZarazClickListenerRule],
// [ZarazConfigReplaceParamsTriggersExcludeRulesZarazTimerRule],
// [ZarazConfigReplaceParamsTriggersExcludeRulesZarazFormSubmissionRule],
// [ZarazConfigReplaceParamsTriggersExcludeRulesZarazVariableMatchRule],
// [ZarazConfigReplaceParamsTriggersExcludeRulesZarazScrollDepthRule],
// [ZarazConfigReplaceParamsTriggersExcludeRulesZarazElementVisibilityRule].
type ZarazConfigReplaceParamsTriggersExcludeRule interface {
	implementsZarazConfigReplaceParamsTriggersExcludeRule()
}

type ZarazConfigReplaceParamsTriggersExcludeRulesZarazLoadRule struct {
	ID    param.Field[string]                                                      `json:"id,required"`
	Match param.Field[string]                                                      `json:"match,required"`
	Op    param.Field[ZarazConfigReplaceParamsTriggersExcludeRulesZarazLoadRuleOp] `json:"op,required"`
	Value param.Field[string]                                                      `json:"value,required"`
}

func (r ZarazConfigReplaceParamsTriggersExcludeRulesZarazLoadRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZarazConfigReplaceParamsTriggersExcludeRulesZarazLoadRule) implementsZarazConfigReplaceParamsTriggersExcludeRule() {
}

type ZarazConfigReplaceParamsTriggersExcludeRulesZarazLoadRuleOp string

const (
	ZarazConfigReplaceParamsTriggersExcludeRulesZarazLoadRuleOpContains           ZarazConfigReplaceParamsTriggersExcludeRulesZarazLoadRuleOp = "CONTAINS"
	ZarazConfigReplaceParamsTriggersExcludeRulesZarazLoadRuleOpEquals             ZarazConfigReplaceParamsTriggersExcludeRulesZarazLoadRuleOp = "EQUALS"
	ZarazConfigReplaceParamsTriggersExcludeRulesZarazLoadRuleOpStartsWith         ZarazConfigReplaceParamsTriggersExcludeRulesZarazLoadRuleOp = "STARTS_WITH"
	ZarazConfigReplaceParamsTriggersExcludeRulesZarazLoadRuleOpEndsWith           ZarazConfigReplaceParamsTriggersExcludeRulesZarazLoadRuleOp = "ENDS_WITH"
	ZarazConfigReplaceParamsTriggersExcludeRulesZarazLoadRuleOpMatchRegex         ZarazConfigReplaceParamsTriggersExcludeRulesZarazLoadRuleOp = "MATCH_REGEX"
	ZarazConfigReplaceParamsTriggersExcludeRulesZarazLoadRuleOpNotMatchRegex      ZarazConfigReplaceParamsTriggersExcludeRulesZarazLoadRuleOp = "NOT_MATCH_REGEX"
	ZarazConfigReplaceParamsTriggersExcludeRulesZarazLoadRuleOpGreaterThan        ZarazConfigReplaceParamsTriggersExcludeRulesZarazLoadRuleOp = "GREATER_THAN"
	ZarazConfigReplaceParamsTriggersExcludeRulesZarazLoadRuleOpGreaterThanOrEqual ZarazConfigReplaceParamsTriggersExcludeRulesZarazLoadRuleOp = "GREATER_THAN_OR_EQUAL"
	ZarazConfigReplaceParamsTriggersExcludeRulesZarazLoadRuleOpLessThan           ZarazConfigReplaceParamsTriggersExcludeRulesZarazLoadRuleOp = "LESS_THAN"
	ZarazConfigReplaceParamsTriggersExcludeRulesZarazLoadRuleOpLessThanOrEqual    ZarazConfigReplaceParamsTriggersExcludeRulesZarazLoadRuleOp = "LESS_THAN_OR_EQUAL"
)

type ZarazConfigReplaceParamsTriggersExcludeRulesZarazClickListenerRule struct {
	ID       param.Field[string]                                                                     `json:"id,required"`
	Action   param.Field[ZarazConfigReplaceParamsTriggersExcludeRulesZarazClickListenerRuleAction]   `json:"action,required"`
	Settings param.Field[ZarazConfigReplaceParamsTriggersExcludeRulesZarazClickListenerRuleSettings] `json:"settings,required"`
}

func (r ZarazConfigReplaceParamsTriggersExcludeRulesZarazClickListenerRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZarazConfigReplaceParamsTriggersExcludeRulesZarazClickListenerRule) implementsZarazConfigReplaceParamsTriggersExcludeRule() {
}

type ZarazConfigReplaceParamsTriggersExcludeRulesZarazClickListenerRuleAction string

const (
	ZarazConfigReplaceParamsTriggersExcludeRulesZarazClickListenerRuleActionClickListener ZarazConfigReplaceParamsTriggersExcludeRulesZarazClickListenerRuleAction = "clickListener"
)

type ZarazConfigReplaceParamsTriggersExcludeRulesZarazClickListenerRuleSettings struct {
	Selector    param.Field[string]                                                                         `json:"selector,required"`
	Type        param.Field[ZarazConfigReplaceParamsTriggersExcludeRulesZarazClickListenerRuleSettingsType] `json:"type,required"`
	WaitForTags param.Field[int64]                                                                          `json:"waitForTags,required"`
}

func (r ZarazConfigReplaceParamsTriggersExcludeRulesZarazClickListenerRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigReplaceParamsTriggersExcludeRulesZarazClickListenerRuleSettingsType string

const (
	ZarazConfigReplaceParamsTriggersExcludeRulesZarazClickListenerRuleSettingsTypeXpath ZarazConfigReplaceParamsTriggersExcludeRulesZarazClickListenerRuleSettingsType = "xpath"
	ZarazConfigReplaceParamsTriggersExcludeRulesZarazClickListenerRuleSettingsTypeCss   ZarazConfigReplaceParamsTriggersExcludeRulesZarazClickListenerRuleSettingsType = "css"
)

type ZarazConfigReplaceParamsTriggersExcludeRulesZarazTimerRule struct {
	ID       param.Field[string]                                                             `json:"id,required"`
	Action   param.Field[ZarazConfigReplaceParamsTriggersExcludeRulesZarazTimerRuleAction]   `json:"action,required"`
	Settings param.Field[ZarazConfigReplaceParamsTriggersExcludeRulesZarazTimerRuleSettings] `json:"settings,required"`
}

func (r ZarazConfigReplaceParamsTriggersExcludeRulesZarazTimerRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZarazConfigReplaceParamsTriggersExcludeRulesZarazTimerRule) implementsZarazConfigReplaceParamsTriggersExcludeRule() {
}

type ZarazConfigReplaceParamsTriggersExcludeRulesZarazTimerRuleAction string

const (
	ZarazConfigReplaceParamsTriggersExcludeRulesZarazTimerRuleActionTimer ZarazConfigReplaceParamsTriggersExcludeRulesZarazTimerRuleAction = "timer"
)

type ZarazConfigReplaceParamsTriggersExcludeRulesZarazTimerRuleSettings struct {
	Interval param.Field[int64] `json:"interval,required"`
	Limit    param.Field[int64] `json:"limit,required"`
}

func (r ZarazConfigReplaceParamsTriggersExcludeRulesZarazTimerRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigReplaceParamsTriggersExcludeRulesZarazFormSubmissionRule struct {
	ID       param.Field[string]                                                                      `json:"id,required"`
	Action   param.Field[ZarazConfigReplaceParamsTriggersExcludeRulesZarazFormSubmissionRuleAction]   `json:"action,required"`
	Settings param.Field[ZarazConfigReplaceParamsTriggersExcludeRulesZarazFormSubmissionRuleSettings] `json:"settings,required"`
}

func (r ZarazConfigReplaceParamsTriggersExcludeRulesZarazFormSubmissionRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZarazConfigReplaceParamsTriggersExcludeRulesZarazFormSubmissionRule) implementsZarazConfigReplaceParamsTriggersExcludeRule() {
}

type ZarazConfigReplaceParamsTriggersExcludeRulesZarazFormSubmissionRuleAction string

const (
	ZarazConfigReplaceParamsTriggersExcludeRulesZarazFormSubmissionRuleActionFormSubmission ZarazConfigReplaceParamsTriggersExcludeRulesZarazFormSubmissionRuleAction = "formSubmission"
)

type ZarazConfigReplaceParamsTriggersExcludeRulesZarazFormSubmissionRuleSettings struct {
	Selector param.Field[string] `json:"selector,required"`
	Validate param.Field[bool]   `json:"validate,required"`
}

func (r ZarazConfigReplaceParamsTriggersExcludeRulesZarazFormSubmissionRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigReplaceParamsTriggersExcludeRulesZarazVariableMatchRule struct {
	ID       param.Field[string]                                                                     `json:"id,required"`
	Action   param.Field[ZarazConfigReplaceParamsTriggersExcludeRulesZarazVariableMatchRuleAction]   `json:"action,required"`
	Settings param.Field[ZarazConfigReplaceParamsTriggersExcludeRulesZarazVariableMatchRuleSettings] `json:"settings,required"`
}

func (r ZarazConfigReplaceParamsTriggersExcludeRulesZarazVariableMatchRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZarazConfigReplaceParamsTriggersExcludeRulesZarazVariableMatchRule) implementsZarazConfigReplaceParamsTriggersExcludeRule() {
}

type ZarazConfigReplaceParamsTriggersExcludeRulesZarazVariableMatchRuleAction string

const (
	ZarazConfigReplaceParamsTriggersExcludeRulesZarazVariableMatchRuleActionVariableMatch ZarazConfigReplaceParamsTriggersExcludeRulesZarazVariableMatchRuleAction = "variableMatch"
)

type ZarazConfigReplaceParamsTriggersExcludeRulesZarazVariableMatchRuleSettings struct {
	Match    param.Field[string] `json:"match,required"`
	Variable param.Field[string] `json:"variable,required"`
}

func (r ZarazConfigReplaceParamsTriggersExcludeRulesZarazVariableMatchRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigReplaceParamsTriggersExcludeRulesZarazScrollDepthRule struct {
	ID       param.Field[string]                                                                   `json:"id,required"`
	Action   param.Field[ZarazConfigReplaceParamsTriggersExcludeRulesZarazScrollDepthRuleAction]   `json:"action,required"`
	Settings param.Field[ZarazConfigReplaceParamsTriggersExcludeRulesZarazScrollDepthRuleSettings] `json:"settings,required"`
}

func (r ZarazConfigReplaceParamsTriggersExcludeRulesZarazScrollDepthRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZarazConfigReplaceParamsTriggersExcludeRulesZarazScrollDepthRule) implementsZarazConfigReplaceParamsTriggersExcludeRule() {
}

type ZarazConfigReplaceParamsTriggersExcludeRulesZarazScrollDepthRuleAction string

const (
	ZarazConfigReplaceParamsTriggersExcludeRulesZarazScrollDepthRuleActionScrollDepth ZarazConfigReplaceParamsTriggersExcludeRulesZarazScrollDepthRuleAction = "scrollDepth"
)

type ZarazConfigReplaceParamsTriggersExcludeRulesZarazScrollDepthRuleSettings struct {
	Positions param.Field[string] `json:"positions,required"`
}

func (r ZarazConfigReplaceParamsTriggersExcludeRulesZarazScrollDepthRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigReplaceParamsTriggersExcludeRulesZarazElementVisibilityRule struct {
	ID       param.Field[string]                                                                         `json:"id,required"`
	Action   param.Field[ZarazConfigReplaceParamsTriggersExcludeRulesZarazElementVisibilityRuleAction]   `json:"action,required"`
	Settings param.Field[ZarazConfigReplaceParamsTriggersExcludeRulesZarazElementVisibilityRuleSettings] `json:"settings,required"`
}

func (r ZarazConfigReplaceParamsTriggersExcludeRulesZarazElementVisibilityRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZarazConfigReplaceParamsTriggersExcludeRulesZarazElementVisibilityRule) implementsZarazConfigReplaceParamsTriggersExcludeRule() {
}

type ZarazConfigReplaceParamsTriggersExcludeRulesZarazElementVisibilityRuleAction string

const (
	ZarazConfigReplaceParamsTriggersExcludeRulesZarazElementVisibilityRuleActionElementVisibility ZarazConfigReplaceParamsTriggersExcludeRulesZarazElementVisibilityRuleAction = "elementVisibility"
)

type ZarazConfigReplaceParamsTriggersExcludeRulesZarazElementVisibilityRuleSettings struct {
	Selector param.Field[string] `json:"selector,required"`
}

func (r ZarazConfigReplaceParamsTriggersExcludeRulesZarazElementVisibilityRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [ZarazConfigReplaceParamsTriggersLoadRulesZarazLoadRule],
// [ZarazConfigReplaceParamsTriggersLoadRulesZarazClickListenerRule],
// [ZarazConfigReplaceParamsTriggersLoadRulesZarazTimerRule],
// [ZarazConfigReplaceParamsTriggersLoadRulesZarazFormSubmissionRule],
// [ZarazConfigReplaceParamsTriggersLoadRulesZarazVariableMatchRule],
// [ZarazConfigReplaceParamsTriggersLoadRulesZarazScrollDepthRule],
// [ZarazConfigReplaceParamsTriggersLoadRulesZarazElementVisibilityRule].
type ZarazConfigReplaceParamsTriggersLoadRule interface {
	implementsZarazConfigReplaceParamsTriggersLoadRule()
}

type ZarazConfigReplaceParamsTriggersLoadRulesZarazLoadRule struct {
	ID    param.Field[string]                                                   `json:"id,required"`
	Match param.Field[string]                                                   `json:"match,required"`
	Op    param.Field[ZarazConfigReplaceParamsTriggersLoadRulesZarazLoadRuleOp] `json:"op,required"`
	Value param.Field[string]                                                   `json:"value,required"`
}

func (r ZarazConfigReplaceParamsTriggersLoadRulesZarazLoadRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZarazConfigReplaceParamsTriggersLoadRulesZarazLoadRule) implementsZarazConfigReplaceParamsTriggersLoadRule() {
}

type ZarazConfigReplaceParamsTriggersLoadRulesZarazLoadRuleOp string

const (
	ZarazConfigReplaceParamsTriggersLoadRulesZarazLoadRuleOpContains           ZarazConfigReplaceParamsTriggersLoadRulesZarazLoadRuleOp = "CONTAINS"
	ZarazConfigReplaceParamsTriggersLoadRulesZarazLoadRuleOpEquals             ZarazConfigReplaceParamsTriggersLoadRulesZarazLoadRuleOp = "EQUALS"
	ZarazConfigReplaceParamsTriggersLoadRulesZarazLoadRuleOpStartsWith         ZarazConfigReplaceParamsTriggersLoadRulesZarazLoadRuleOp = "STARTS_WITH"
	ZarazConfigReplaceParamsTriggersLoadRulesZarazLoadRuleOpEndsWith           ZarazConfigReplaceParamsTriggersLoadRulesZarazLoadRuleOp = "ENDS_WITH"
	ZarazConfigReplaceParamsTriggersLoadRulesZarazLoadRuleOpMatchRegex         ZarazConfigReplaceParamsTriggersLoadRulesZarazLoadRuleOp = "MATCH_REGEX"
	ZarazConfigReplaceParamsTriggersLoadRulesZarazLoadRuleOpNotMatchRegex      ZarazConfigReplaceParamsTriggersLoadRulesZarazLoadRuleOp = "NOT_MATCH_REGEX"
	ZarazConfigReplaceParamsTriggersLoadRulesZarazLoadRuleOpGreaterThan        ZarazConfigReplaceParamsTriggersLoadRulesZarazLoadRuleOp = "GREATER_THAN"
	ZarazConfigReplaceParamsTriggersLoadRulesZarazLoadRuleOpGreaterThanOrEqual ZarazConfigReplaceParamsTriggersLoadRulesZarazLoadRuleOp = "GREATER_THAN_OR_EQUAL"
	ZarazConfigReplaceParamsTriggersLoadRulesZarazLoadRuleOpLessThan           ZarazConfigReplaceParamsTriggersLoadRulesZarazLoadRuleOp = "LESS_THAN"
	ZarazConfigReplaceParamsTriggersLoadRulesZarazLoadRuleOpLessThanOrEqual    ZarazConfigReplaceParamsTriggersLoadRulesZarazLoadRuleOp = "LESS_THAN_OR_EQUAL"
)

type ZarazConfigReplaceParamsTriggersLoadRulesZarazClickListenerRule struct {
	ID       param.Field[string]                                                                  `json:"id,required"`
	Action   param.Field[ZarazConfigReplaceParamsTriggersLoadRulesZarazClickListenerRuleAction]   `json:"action,required"`
	Settings param.Field[ZarazConfigReplaceParamsTriggersLoadRulesZarazClickListenerRuleSettings] `json:"settings,required"`
}

func (r ZarazConfigReplaceParamsTriggersLoadRulesZarazClickListenerRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZarazConfigReplaceParamsTriggersLoadRulesZarazClickListenerRule) implementsZarazConfigReplaceParamsTriggersLoadRule() {
}

type ZarazConfigReplaceParamsTriggersLoadRulesZarazClickListenerRuleAction string

const (
	ZarazConfigReplaceParamsTriggersLoadRulesZarazClickListenerRuleActionClickListener ZarazConfigReplaceParamsTriggersLoadRulesZarazClickListenerRuleAction = "clickListener"
)

type ZarazConfigReplaceParamsTriggersLoadRulesZarazClickListenerRuleSettings struct {
	Selector    param.Field[string]                                                                      `json:"selector,required"`
	Type        param.Field[ZarazConfigReplaceParamsTriggersLoadRulesZarazClickListenerRuleSettingsType] `json:"type,required"`
	WaitForTags param.Field[int64]                                                                       `json:"waitForTags,required"`
}

func (r ZarazConfigReplaceParamsTriggersLoadRulesZarazClickListenerRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigReplaceParamsTriggersLoadRulesZarazClickListenerRuleSettingsType string

const (
	ZarazConfigReplaceParamsTriggersLoadRulesZarazClickListenerRuleSettingsTypeXpath ZarazConfigReplaceParamsTriggersLoadRulesZarazClickListenerRuleSettingsType = "xpath"
	ZarazConfigReplaceParamsTriggersLoadRulesZarazClickListenerRuleSettingsTypeCss   ZarazConfigReplaceParamsTriggersLoadRulesZarazClickListenerRuleSettingsType = "css"
)

type ZarazConfigReplaceParamsTriggersLoadRulesZarazTimerRule struct {
	ID       param.Field[string]                                                          `json:"id,required"`
	Action   param.Field[ZarazConfigReplaceParamsTriggersLoadRulesZarazTimerRuleAction]   `json:"action,required"`
	Settings param.Field[ZarazConfigReplaceParamsTriggersLoadRulesZarazTimerRuleSettings] `json:"settings,required"`
}

func (r ZarazConfigReplaceParamsTriggersLoadRulesZarazTimerRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZarazConfigReplaceParamsTriggersLoadRulesZarazTimerRule) implementsZarazConfigReplaceParamsTriggersLoadRule() {
}

type ZarazConfigReplaceParamsTriggersLoadRulesZarazTimerRuleAction string

const (
	ZarazConfigReplaceParamsTriggersLoadRulesZarazTimerRuleActionTimer ZarazConfigReplaceParamsTriggersLoadRulesZarazTimerRuleAction = "timer"
)

type ZarazConfigReplaceParamsTriggersLoadRulesZarazTimerRuleSettings struct {
	Interval param.Field[int64] `json:"interval,required"`
	Limit    param.Field[int64] `json:"limit,required"`
}

func (r ZarazConfigReplaceParamsTriggersLoadRulesZarazTimerRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigReplaceParamsTriggersLoadRulesZarazFormSubmissionRule struct {
	ID       param.Field[string]                                                                   `json:"id,required"`
	Action   param.Field[ZarazConfigReplaceParamsTriggersLoadRulesZarazFormSubmissionRuleAction]   `json:"action,required"`
	Settings param.Field[ZarazConfigReplaceParamsTriggersLoadRulesZarazFormSubmissionRuleSettings] `json:"settings,required"`
}

func (r ZarazConfigReplaceParamsTriggersLoadRulesZarazFormSubmissionRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZarazConfigReplaceParamsTriggersLoadRulesZarazFormSubmissionRule) implementsZarazConfigReplaceParamsTriggersLoadRule() {
}

type ZarazConfigReplaceParamsTriggersLoadRulesZarazFormSubmissionRuleAction string

const (
	ZarazConfigReplaceParamsTriggersLoadRulesZarazFormSubmissionRuleActionFormSubmission ZarazConfigReplaceParamsTriggersLoadRulesZarazFormSubmissionRuleAction = "formSubmission"
)

type ZarazConfigReplaceParamsTriggersLoadRulesZarazFormSubmissionRuleSettings struct {
	Selector param.Field[string] `json:"selector,required"`
	Validate param.Field[bool]   `json:"validate,required"`
}

func (r ZarazConfigReplaceParamsTriggersLoadRulesZarazFormSubmissionRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigReplaceParamsTriggersLoadRulesZarazVariableMatchRule struct {
	ID       param.Field[string]                                                                  `json:"id,required"`
	Action   param.Field[ZarazConfigReplaceParamsTriggersLoadRulesZarazVariableMatchRuleAction]   `json:"action,required"`
	Settings param.Field[ZarazConfigReplaceParamsTriggersLoadRulesZarazVariableMatchRuleSettings] `json:"settings,required"`
}

func (r ZarazConfigReplaceParamsTriggersLoadRulesZarazVariableMatchRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZarazConfigReplaceParamsTriggersLoadRulesZarazVariableMatchRule) implementsZarazConfigReplaceParamsTriggersLoadRule() {
}

type ZarazConfigReplaceParamsTriggersLoadRulesZarazVariableMatchRuleAction string

const (
	ZarazConfigReplaceParamsTriggersLoadRulesZarazVariableMatchRuleActionVariableMatch ZarazConfigReplaceParamsTriggersLoadRulesZarazVariableMatchRuleAction = "variableMatch"
)

type ZarazConfigReplaceParamsTriggersLoadRulesZarazVariableMatchRuleSettings struct {
	Match    param.Field[string] `json:"match,required"`
	Variable param.Field[string] `json:"variable,required"`
}

func (r ZarazConfigReplaceParamsTriggersLoadRulesZarazVariableMatchRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigReplaceParamsTriggersLoadRulesZarazScrollDepthRule struct {
	ID       param.Field[string]                                                                `json:"id,required"`
	Action   param.Field[ZarazConfigReplaceParamsTriggersLoadRulesZarazScrollDepthRuleAction]   `json:"action,required"`
	Settings param.Field[ZarazConfigReplaceParamsTriggersLoadRulesZarazScrollDepthRuleSettings] `json:"settings,required"`
}

func (r ZarazConfigReplaceParamsTriggersLoadRulesZarazScrollDepthRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZarazConfigReplaceParamsTriggersLoadRulesZarazScrollDepthRule) implementsZarazConfigReplaceParamsTriggersLoadRule() {
}

type ZarazConfigReplaceParamsTriggersLoadRulesZarazScrollDepthRuleAction string

const (
	ZarazConfigReplaceParamsTriggersLoadRulesZarazScrollDepthRuleActionScrollDepth ZarazConfigReplaceParamsTriggersLoadRulesZarazScrollDepthRuleAction = "scrollDepth"
)

type ZarazConfigReplaceParamsTriggersLoadRulesZarazScrollDepthRuleSettings struct {
	Positions param.Field[string] `json:"positions,required"`
}

func (r ZarazConfigReplaceParamsTriggersLoadRulesZarazScrollDepthRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigReplaceParamsTriggersLoadRulesZarazElementVisibilityRule struct {
	ID       param.Field[string]                                                                      `json:"id,required"`
	Action   param.Field[ZarazConfigReplaceParamsTriggersLoadRulesZarazElementVisibilityRuleAction]   `json:"action,required"`
	Settings param.Field[ZarazConfigReplaceParamsTriggersLoadRulesZarazElementVisibilityRuleSettings] `json:"settings,required"`
}

func (r ZarazConfigReplaceParamsTriggersLoadRulesZarazElementVisibilityRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZarazConfigReplaceParamsTriggersLoadRulesZarazElementVisibilityRule) implementsZarazConfigReplaceParamsTriggersLoadRule() {
}

type ZarazConfigReplaceParamsTriggersLoadRulesZarazElementVisibilityRuleAction string

const (
	ZarazConfigReplaceParamsTriggersLoadRulesZarazElementVisibilityRuleActionElementVisibility ZarazConfigReplaceParamsTriggersLoadRulesZarazElementVisibilityRuleAction = "elementVisibility"
)

type ZarazConfigReplaceParamsTriggersLoadRulesZarazElementVisibilityRuleSettings struct {
	Selector param.Field[string] `json:"selector,required"`
}

func (r ZarazConfigReplaceParamsTriggersLoadRulesZarazElementVisibilityRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigReplaceParamsTriggersSystem string

const (
	ZarazConfigReplaceParamsTriggersSystemPageload ZarazConfigReplaceParamsTriggersSystem = "pageload"
)

// Satisfied by [ZarazConfigReplaceParamsVariablesObject],
// [ZarazConfigReplaceParamsVariablesObject].
type ZarazConfigReplaceParamsVariables interface {
	implementsZarazConfigReplaceParamsVariables()
}

type ZarazConfigReplaceParamsVariablesObject struct {
	Name  param.Field[string]                                      `json:"name,required"`
	Type  param.Field[ZarazConfigReplaceParamsVariablesObjectType] `json:"type,required"`
	Value param.Field[string]                                      `json:"value,required"`
}

func (r ZarazConfigReplaceParamsVariablesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZarazConfigReplaceParamsVariablesObject) implementsZarazConfigReplaceParamsVariables() {}

type ZarazConfigReplaceParamsVariablesObjectType string

const (
	ZarazConfigReplaceParamsVariablesObjectTypeString ZarazConfigReplaceParamsVariablesObjectType = "string"
	ZarazConfigReplaceParamsVariablesObjectTypeSecret ZarazConfigReplaceParamsVariablesObjectType = "secret"
)

// Consent management configuration.
type ZarazConfigReplaceParamsConsent struct {
	Enabled                param.Field[bool]                                                  `json:"enabled,required"`
	ButtonTextTranslations param.Field[ZarazConfigReplaceParamsConsentButtonTextTranslations] `json:"buttonTextTranslations"`
	CompanyEmail           param.Field[string]                                                `json:"companyEmail"`
	CompanyName            param.Field[string]                                                `json:"companyName"`
	CompanyStreetAddress   param.Field[string]                                                `json:"companyStreetAddress"`
	ConsentModalIntroHTML  param.Field[string]                                                `json:"consentModalIntroHTML"`
	// Object where keys are language codes
	ConsentModalIntroHTMLWithTranslations param.Field[map[string]string] `json:"consentModalIntroHTMLWithTranslations"`
	CookieName                            param.Field[string]            `json:"cookieName"`
	CustomCss                             param.Field[string]            `json:"customCSS"`
	CustomIntroDisclaimerDismissed        param.Field[bool]              `json:"customIntroDisclaimerDismissed"`
	DefaultLanguage                       param.Field[string]            `json:"defaultLanguage"`
	HideModal                             param.Field[bool]              `json:"hideModal"`
	// Object where keys are purpose alpha-numeric IDs
	Purposes param.Field[map[string]ZarazConfigReplaceParamsConsentPurposes] `json:"purposes"`
	// Object where keys are purpose alpha-numeric IDs
	PurposesWithTranslations param.Field[map[string]ZarazConfigReplaceParamsConsentPurposesWithTranslations] `json:"purposesWithTranslations"`
}

func (r ZarazConfigReplaceParamsConsent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigReplaceParamsConsentButtonTextTranslations struct {
	// Object where keys are language codes
	AcceptAll param.Field[map[string]string] `json:"accept_all,required"`
	// Object where keys are language codes
	ConfirmMyChoices param.Field[map[string]string] `json:"confirm_my_choices,required"`
	// Object where keys are language codes
	RejectAll param.Field[map[string]string] `json:"reject_all,required"`
}

func (r ZarazConfigReplaceParamsConsentButtonTextTranslations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigReplaceParamsConsentPurposes struct {
	Description param.Field[string] `json:"description,required"`
	Name        param.Field[string] `json:"name,required"`
}

func (r ZarazConfigReplaceParamsConsentPurposes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigReplaceParamsConsentPurposesWithTranslations struct {
	// Object where keys are language codes
	Description param.Field[map[string]string] `json:"description,required"`
	// Object where keys are language codes
	Name  param.Field[map[string]string] `json:"name,required"`
	Order param.Field[int64]             `json:"order,required"`
}

func (r ZarazConfigReplaceParamsConsentPurposesWithTranslations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigReplaceResponseEnvelope struct {
	Errors   []ZarazConfigReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZarazConfigReplaceResponseEnvelopeMessages `json:"messages,required"`
	// Zaraz configuration
	Result ZarazConfigReplaceResponse `json:"result,required"`
	// Whether the API call was successful
	Success bool                                   `json:"success,required"`
	JSON    zarazConfigReplaceResponseEnvelopeJSON `json:"-"`
}

// zarazConfigReplaceResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZarazConfigReplaceResponseEnvelope]
type zarazConfigReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigReplaceResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    zarazConfigReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// zarazConfigReplaceResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZarazConfigReplaceResponseEnvelopeErrors]
type zarazConfigReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigReplaceResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zarazConfigReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// zarazConfigReplaceResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZarazConfigReplaceResponseEnvelopeMessages]
type zarazConfigReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
