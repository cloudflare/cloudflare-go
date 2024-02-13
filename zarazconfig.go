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

// Updates Zaraz configuration for a zone.
func (r *ZarazConfigService) Update(ctx context.Context, zoneID string, body ZarazConfigUpdateParams, opts ...option.RequestOption) (res *ZarazConfigUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZarazConfigUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/zaraz/config", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
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

// Zaraz configuration
type ZarazConfigUpdateResponse struct {
	// Data layer compatibility mode enabled.
	DataLayer bool `json:"dataLayer,required"`
	// The key for Zaraz debug mode.
	DebugKey string `json:"debugKey,required"`
	// General Zaraz settings.
	Settings ZarazConfigUpdateResponseSettings `json:"settings,required"`
	// Tools set up under Zaraz configuration, where key is the alpha-numeric tool ID
	// and value is the tool configuration object.
	Tools map[string]ZarazConfigUpdateResponseTool `json:"tools,required"`
	// Triggers set up under Zaraz configuration, where key is the trigger
	// alpha-numeric ID and value is the trigger configuration.
	Triggers map[string]ZarazConfigUpdateResponseTrigger `json:"triggers,required"`
	// Variables set up under Zaraz configuration, where key is the variable
	// alpha-numeric ID and value is the variable configuration. Values of variables of
	// type secret are not included.
	Variables map[string]ZarazConfigUpdateResponseVariable `json:"variables,required"`
	// Zaraz internal version of the config.
	ZarazVersion int64 `json:"zarazVersion,required"`
	// Consent management configuration.
	Consent ZarazConfigUpdateResponseConsent `json:"consent"`
	// Single Page Application support enabled.
	HistoryChange bool                          `json:"historyChange"`
	JSON          zarazConfigUpdateResponseJSON `json:"-"`
}

// zarazConfigUpdateResponseJSON contains the JSON metadata for the struct
// [ZarazConfigUpdateResponse]
type zarazConfigUpdateResponseJSON struct {
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

func (r *ZarazConfigUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// General Zaraz settings.
type ZarazConfigUpdateResponseSettings struct {
	// Automatic injection of Zaraz scripts enabled.
	AutoInjectScript bool `json:"autoInjectScript,required"`
	// Details of the worker that receives and edits Zaraz Context object.
	ContextEnricher ZarazConfigUpdateResponseSettingsContextEnricher `json:"contextEnricher"`
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
	TrackPath string                                `json:"trackPath"`
	JSON      zarazConfigUpdateResponseSettingsJSON `json:"-"`
}

// zarazConfigUpdateResponseSettingsJSON contains the JSON metadata for the struct
// [ZarazConfigUpdateResponseSettings]
type zarazConfigUpdateResponseSettingsJSON struct {
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

func (r *ZarazConfigUpdateResponseSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details of the worker that receives and edits Zaraz Context object.
type ZarazConfigUpdateResponseSettingsContextEnricher struct {
	EscapedWorkerName string                                               `json:"escapedWorkerName,required"`
	WorkerTag         string                                               `json:"workerTag,required"`
	JSON              zarazConfigUpdateResponseSettingsContextEnricherJSON `json:"-"`
}

// zarazConfigUpdateResponseSettingsContextEnricherJSON contains the JSON metadata
// for the struct [ZarazConfigUpdateResponseSettingsContextEnricher]
type zarazConfigUpdateResponseSettingsContextEnricherJSON struct {
	EscapedWorkerName apijson.Field
	WorkerTag         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseSettingsContextEnricher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [ZarazConfigUpdateResponseToolsZarazManagedComponent] or
// [ZarazConfigUpdateResponseToolsZarazCustomManagedComponent].
type ZarazConfigUpdateResponseTool interface {
	implementsZarazConfigUpdateResponseTool()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZarazConfigUpdateResponseTool)(nil)).Elem(), "")
}

type ZarazConfigUpdateResponseToolsZarazManagedComponent struct {
	// List of blocking trigger IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Tool's internal name
	Component string `json:"component,required"`
	// Default fields for tool's actions
	DefaultFields map[string]ZarazConfigUpdateResponseToolsZarazManagedComponentDefaultField `json:"defaultFields,required"`
	// Whether tool is enabled
	Enabled bool `json:"enabled,required"`
	// Tool's name defined by the user
	Name string `json:"name,required"`
	// List of permissions granted to the component
	Permissions []string `json:"permissions,required"`
	// Tool's settings
	Settings map[string]ZarazConfigUpdateResponseToolsZarazManagedComponentSetting `json:"settings,required"`
	Type     ZarazConfigUpdateResponseToolsZarazManagedComponentType               `json:"type,required"`
	// Actions configured on a tool. Either this or neoEvents field is required.
	Actions map[string]ZarazConfigUpdateResponseToolsZarazManagedComponentAction `json:"actions"`
	// Default consent purpose ID
	DefaultPurpose string `json:"defaultPurpose"`
	// DEPRECATED - List of actions configured on a tool. Either this or actions field
	// is required. If both are present, actions field will take precedence.
	NeoEvents []ZarazConfigUpdateResponseToolsZarazManagedComponentNeoEvent `json:"neoEvents"`
	JSON      zarazConfigUpdateResponseToolsZarazManagedComponentJSON       `json:"-"`
}

// zarazConfigUpdateResponseToolsZarazManagedComponentJSON contains the JSON
// metadata for the struct [ZarazConfigUpdateResponseToolsZarazManagedComponent]
type zarazConfigUpdateResponseToolsZarazManagedComponentJSON struct {
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

func (r *ZarazConfigUpdateResponseToolsZarazManagedComponent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigUpdateResponseToolsZarazManagedComponent) implementsZarazConfigUpdateResponseTool() {
}

// Union satisfied by [shared.UnionString] or [shared.UnionBool].
type ZarazConfigUpdateResponseToolsZarazManagedComponentDefaultField interface {
	ImplementsZarazConfigUpdateResponseToolsZarazManagedComponentDefaultField()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazConfigUpdateResponseToolsZarazManagedComponentDefaultField)(nil)).Elem(),
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
type ZarazConfigUpdateResponseToolsZarazManagedComponentSetting interface {
	ImplementsZarazConfigUpdateResponseToolsZarazManagedComponentSetting()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazConfigUpdateResponseToolsZarazManagedComponentSetting)(nil)).Elem(),
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

type ZarazConfigUpdateResponseToolsZarazManagedComponentType string

const (
	ZarazConfigUpdateResponseToolsZarazManagedComponentTypeComponent ZarazConfigUpdateResponseToolsZarazManagedComponentType = "component"
)

type ZarazConfigUpdateResponseToolsZarazManagedComponentAction struct {
	// Tool event type
	ActionType string `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Event payload
	Data interface{} `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers []string                                                      `json:"firingTriggers,required"`
	JSON           zarazConfigUpdateResponseToolsZarazManagedComponentActionJSON `json:"-"`
}

// zarazConfigUpdateResponseToolsZarazManagedComponentActionJSON contains the JSON
// metadata for the struct
// [ZarazConfigUpdateResponseToolsZarazManagedComponentAction]
type zarazConfigUpdateResponseToolsZarazManagedComponentActionJSON struct {
	ActionType       apijson.Field
	BlockingTriggers apijson.Field
	Data             apijson.Field
	FiringTriggers   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseToolsZarazManagedComponentAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigUpdateResponseToolsZarazManagedComponentNeoEvent struct {
	// Tool event type
	ActionType string `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Event payload
	Data interface{} `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers []string                                                        `json:"firingTriggers,required"`
	JSON           zarazConfigUpdateResponseToolsZarazManagedComponentNeoEventJSON `json:"-"`
}

// zarazConfigUpdateResponseToolsZarazManagedComponentNeoEventJSON contains the
// JSON metadata for the struct
// [ZarazConfigUpdateResponseToolsZarazManagedComponentNeoEvent]
type zarazConfigUpdateResponseToolsZarazManagedComponentNeoEventJSON struct {
	ActionType       apijson.Field
	BlockingTriggers apijson.Field
	Data             apijson.Field
	FiringTriggers   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseToolsZarazManagedComponentNeoEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigUpdateResponseToolsZarazCustomManagedComponent struct {
	// List of blocking trigger IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Tool's internal name
	Component string `json:"component,required"`
	// Default fields for tool's actions
	DefaultFields map[string]ZarazConfigUpdateResponseToolsZarazCustomManagedComponentDefaultField `json:"defaultFields,required"`
	// Whether tool is enabled
	Enabled bool `json:"enabled,required"`
	// Tool's name defined by the user
	Name string `json:"name,required"`
	// List of permissions granted to the component
	Permissions []string `json:"permissions,required"`
	// Tool's settings
	Settings map[string]ZarazConfigUpdateResponseToolsZarazCustomManagedComponentSetting `json:"settings,required"`
	Type     ZarazConfigUpdateResponseToolsZarazCustomManagedComponentType               `json:"type,required"`
	// Cloudflare worker that acts as a managed component
	Worker ZarazConfigUpdateResponseToolsZarazCustomManagedComponentWorker `json:"worker,required"`
	// Actions configured on a tool. Either this or neoEvents field is required.
	Actions map[string]ZarazConfigUpdateResponseToolsZarazCustomManagedComponentAction `json:"actions"`
	// Default consent purpose ID
	DefaultPurpose string `json:"defaultPurpose"`
	// DEPRECATED - List of actions configured on a tool. Either this or actions field
	// is required. If both are present, actions field will take precedence.
	NeoEvents []ZarazConfigUpdateResponseToolsZarazCustomManagedComponentNeoEvent `json:"neoEvents"`
	JSON      zarazConfigUpdateResponseToolsZarazCustomManagedComponentJSON       `json:"-"`
}

// zarazConfigUpdateResponseToolsZarazCustomManagedComponentJSON contains the JSON
// metadata for the struct
// [ZarazConfigUpdateResponseToolsZarazCustomManagedComponent]
type zarazConfigUpdateResponseToolsZarazCustomManagedComponentJSON struct {
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

func (r *ZarazConfigUpdateResponseToolsZarazCustomManagedComponent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigUpdateResponseToolsZarazCustomManagedComponent) implementsZarazConfigUpdateResponseTool() {
}

// Union satisfied by [shared.UnionString] or [shared.UnionBool].
type ZarazConfigUpdateResponseToolsZarazCustomManagedComponentDefaultField interface {
	ImplementsZarazConfigUpdateResponseToolsZarazCustomManagedComponentDefaultField()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazConfigUpdateResponseToolsZarazCustomManagedComponentDefaultField)(nil)).Elem(),
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
type ZarazConfigUpdateResponseToolsZarazCustomManagedComponentSetting interface {
	ImplementsZarazConfigUpdateResponseToolsZarazCustomManagedComponentSetting()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazConfigUpdateResponseToolsZarazCustomManagedComponentSetting)(nil)).Elem(),
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

type ZarazConfigUpdateResponseToolsZarazCustomManagedComponentType string

const (
	ZarazConfigUpdateResponseToolsZarazCustomManagedComponentTypeCustomMc ZarazConfigUpdateResponseToolsZarazCustomManagedComponentType = "custom-mc"
)

// Cloudflare worker that acts as a managed component
type ZarazConfigUpdateResponseToolsZarazCustomManagedComponentWorker struct {
	EscapedWorkerName string                                                              `json:"escapedWorkerName,required"`
	WorkerTag         string                                                              `json:"workerTag,required"`
	JSON              zarazConfigUpdateResponseToolsZarazCustomManagedComponentWorkerJSON `json:"-"`
}

// zarazConfigUpdateResponseToolsZarazCustomManagedComponentWorkerJSON contains the
// JSON metadata for the struct
// [ZarazConfigUpdateResponseToolsZarazCustomManagedComponentWorker]
type zarazConfigUpdateResponseToolsZarazCustomManagedComponentWorkerJSON struct {
	EscapedWorkerName apijson.Field
	WorkerTag         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseToolsZarazCustomManagedComponentWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigUpdateResponseToolsZarazCustomManagedComponentAction struct {
	// Tool event type
	ActionType string `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Event payload
	Data interface{} `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers []string                                                            `json:"firingTriggers,required"`
	JSON           zarazConfigUpdateResponseToolsZarazCustomManagedComponentActionJSON `json:"-"`
}

// zarazConfigUpdateResponseToolsZarazCustomManagedComponentActionJSON contains the
// JSON metadata for the struct
// [ZarazConfigUpdateResponseToolsZarazCustomManagedComponentAction]
type zarazConfigUpdateResponseToolsZarazCustomManagedComponentActionJSON struct {
	ActionType       apijson.Field
	BlockingTriggers apijson.Field
	Data             apijson.Field
	FiringTriggers   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseToolsZarazCustomManagedComponentAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigUpdateResponseToolsZarazCustomManagedComponentNeoEvent struct {
	// Tool event type
	ActionType string `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Event payload
	Data interface{} `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers []string                                                              `json:"firingTriggers,required"`
	JSON           zarazConfigUpdateResponseToolsZarazCustomManagedComponentNeoEventJSON `json:"-"`
}

// zarazConfigUpdateResponseToolsZarazCustomManagedComponentNeoEventJSON contains
// the JSON metadata for the struct
// [ZarazConfigUpdateResponseToolsZarazCustomManagedComponentNeoEvent]
type zarazConfigUpdateResponseToolsZarazCustomManagedComponentNeoEventJSON struct {
	ActionType       apijson.Field
	BlockingTriggers apijson.Field
	Data             apijson.Field
	FiringTriggers   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseToolsZarazCustomManagedComponentNeoEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigUpdateResponseTrigger struct {
	// Rules defining when the trigger is not fired.
	ExcludeRules []ZarazConfigUpdateResponseTriggersExcludeRule `json:"excludeRules,required"`
	// Rules defining when the trigger is fired.
	LoadRules []ZarazConfigUpdateResponseTriggersLoadRule `json:"loadRules,required"`
	// Trigger name.
	Name string `json:"name,required"`
	// Trigger description.
	Description string                                  `json:"description"`
	System      ZarazConfigUpdateResponseTriggersSystem `json:"system"`
	JSON        zarazConfigUpdateResponseTriggerJSON    `json:"-"`
}

// zarazConfigUpdateResponseTriggerJSON contains the JSON metadata for the struct
// [ZarazConfigUpdateResponseTrigger]
type zarazConfigUpdateResponseTriggerJSON struct {
	ExcludeRules apijson.Field
	LoadRules    apijson.Field
	Name         apijson.Field
	Description  apijson.Field
	System       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [ZarazConfigUpdateResponseTriggersExcludeRulesZarazLoadRule],
// [ZarazConfigUpdateResponseTriggersExcludeRulesZarazClickListenerRule],
// [ZarazConfigUpdateResponseTriggersExcludeRulesZarazTimerRule],
// [ZarazConfigUpdateResponseTriggersExcludeRulesZarazFormSubmissionRule],
// [ZarazConfigUpdateResponseTriggersExcludeRulesZarazVariableMatchRule],
// [ZarazConfigUpdateResponseTriggersExcludeRulesZarazScrollDepthRule] or
// [ZarazConfigUpdateResponseTriggersExcludeRulesZarazElementVisibilityRule].
type ZarazConfigUpdateResponseTriggersExcludeRule interface {
	implementsZarazConfigUpdateResponseTriggersExcludeRule()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZarazConfigUpdateResponseTriggersExcludeRule)(nil)).Elem(), "")
}

type ZarazConfigUpdateResponseTriggersExcludeRulesZarazLoadRule struct {
	ID    string                                                         `json:"id,required"`
	Match string                                                         `json:"match,required"`
	Op    ZarazConfigUpdateResponseTriggersExcludeRulesZarazLoadRuleOp   `json:"op,required"`
	Value string                                                         `json:"value,required"`
	JSON  zarazConfigUpdateResponseTriggersExcludeRulesZarazLoadRuleJSON `json:"-"`
}

// zarazConfigUpdateResponseTriggersExcludeRulesZarazLoadRuleJSON contains the JSON
// metadata for the struct
// [ZarazConfigUpdateResponseTriggersExcludeRulesZarazLoadRule]
type zarazConfigUpdateResponseTriggersExcludeRulesZarazLoadRuleJSON struct {
	ID          apijson.Field
	Match       apijson.Field
	Op          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseTriggersExcludeRulesZarazLoadRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigUpdateResponseTriggersExcludeRulesZarazLoadRule) implementsZarazConfigUpdateResponseTriggersExcludeRule() {
}

type ZarazConfigUpdateResponseTriggersExcludeRulesZarazLoadRuleOp string

const (
	ZarazConfigUpdateResponseTriggersExcludeRulesZarazLoadRuleOpContains           ZarazConfigUpdateResponseTriggersExcludeRulesZarazLoadRuleOp = "CONTAINS"
	ZarazConfigUpdateResponseTriggersExcludeRulesZarazLoadRuleOpEquals             ZarazConfigUpdateResponseTriggersExcludeRulesZarazLoadRuleOp = "EQUALS"
	ZarazConfigUpdateResponseTriggersExcludeRulesZarazLoadRuleOpStartsWith         ZarazConfigUpdateResponseTriggersExcludeRulesZarazLoadRuleOp = "STARTS_WITH"
	ZarazConfigUpdateResponseTriggersExcludeRulesZarazLoadRuleOpEndsWith           ZarazConfigUpdateResponseTriggersExcludeRulesZarazLoadRuleOp = "ENDS_WITH"
	ZarazConfigUpdateResponseTriggersExcludeRulesZarazLoadRuleOpMatchRegex         ZarazConfigUpdateResponseTriggersExcludeRulesZarazLoadRuleOp = "MATCH_REGEX"
	ZarazConfigUpdateResponseTriggersExcludeRulesZarazLoadRuleOpNotMatchRegex      ZarazConfigUpdateResponseTriggersExcludeRulesZarazLoadRuleOp = "NOT_MATCH_REGEX"
	ZarazConfigUpdateResponseTriggersExcludeRulesZarazLoadRuleOpGreaterThan        ZarazConfigUpdateResponseTriggersExcludeRulesZarazLoadRuleOp = "GREATER_THAN"
	ZarazConfigUpdateResponseTriggersExcludeRulesZarazLoadRuleOpGreaterThanOrEqual ZarazConfigUpdateResponseTriggersExcludeRulesZarazLoadRuleOp = "GREATER_THAN_OR_EQUAL"
	ZarazConfigUpdateResponseTriggersExcludeRulesZarazLoadRuleOpLessThan           ZarazConfigUpdateResponseTriggersExcludeRulesZarazLoadRuleOp = "LESS_THAN"
	ZarazConfigUpdateResponseTriggersExcludeRulesZarazLoadRuleOpLessThanOrEqual    ZarazConfigUpdateResponseTriggersExcludeRulesZarazLoadRuleOp = "LESS_THAN_OR_EQUAL"
)

type ZarazConfigUpdateResponseTriggersExcludeRulesZarazClickListenerRule struct {
	ID       string                                                                      `json:"id,required"`
	Action   ZarazConfigUpdateResponseTriggersExcludeRulesZarazClickListenerRuleAction   `json:"action,required"`
	Settings ZarazConfigUpdateResponseTriggersExcludeRulesZarazClickListenerRuleSettings `json:"settings,required"`
	JSON     zarazConfigUpdateResponseTriggersExcludeRulesZarazClickListenerRuleJSON     `json:"-"`
}

// zarazConfigUpdateResponseTriggersExcludeRulesZarazClickListenerRuleJSON contains
// the JSON metadata for the struct
// [ZarazConfigUpdateResponseTriggersExcludeRulesZarazClickListenerRule]
type zarazConfigUpdateResponseTriggersExcludeRulesZarazClickListenerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseTriggersExcludeRulesZarazClickListenerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigUpdateResponseTriggersExcludeRulesZarazClickListenerRule) implementsZarazConfigUpdateResponseTriggersExcludeRule() {
}

type ZarazConfigUpdateResponseTriggersExcludeRulesZarazClickListenerRuleAction string

const (
	ZarazConfigUpdateResponseTriggersExcludeRulesZarazClickListenerRuleActionClickListener ZarazConfigUpdateResponseTriggersExcludeRulesZarazClickListenerRuleAction = "clickListener"
)

type ZarazConfigUpdateResponseTriggersExcludeRulesZarazClickListenerRuleSettings struct {
	Selector    string                                                                          `json:"selector,required"`
	Type        ZarazConfigUpdateResponseTriggersExcludeRulesZarazClickListenerRuleSettingsType `json:"type,required"`
	WaitForTags int64                                                                           `json:"waitForTags,required"`
	JSON        zarazConfigUpdateResponseTriggersExcludeRulesZarazClickListenerRuleSettingsJSON `json:"-"`
}

// zarazConfigUpdateResponseTriggersExcludeRulesZarazClickListenerRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazConfigUpdateResponseTriggersExcludeRulesZarazClickListenerRuleSettings]
type zarazConfigUpdateResponseTriggersExcludeRulesZarazClickListenerRuleSettingsJSON struct {
	Selector    apijson.Field
	Type        apijson.Field
	WaitForTags apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseTriggersExcludeRulesZarazClickListenerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigUpdateResponseTriggersExcludeRulesZarazClickListenerRuleSettingsType string

const (
	ZarazConfigUpdateResponseTriggersExcludeRulesZarazClickListenerRuleSettingsTypeXpath ZarazConfigUpdateResponseTriggersExcludeRulesZarazClickListenerRuleSettingsType = "xpath"
	ZarazConfigUpdateResponseTriggersExcludeRulesZarazClickListenerRuleSettingsTypeCss   ZarazConfigUpdateResponseTriggersExcludeRulesZarazClickListenerRuleSettingsType = "css"
)

type ZarazConfigUpdateResponseTriggersExcludeRulesZarazTimerRule struct {
	ID       string                                                              `json:"id,required"`
	Action   ZarazConfigUpdateResponseTriggersExcludeRulesZarazTimerRuleAction   `json:"action,required"`
	Settings ZarazConfigUpdateResponseTriggersExcludeRulesZarazTimerRuleSettings `json:"settings,required"`
	JSON     zarazConfigUpdateResponseTriggersExcludeRulesZarazTimerRuleJSON     `json:"-"`
}

// zarazConfigUpdateResponseTriggersExcludeRulesZarazTimerRuleJSON contains the
// JSON metadata for the struct
// [ZarazConfigUpdateResponseTriggersExcludeRulesZarazTimerRule]
type zarazConfigUpdateResponseTriggersExcludeRulesZarazTimerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseTriggersExcludeRulesZarazTimerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigUpdateResponseTriggersExcludeRulesZarazTimerRule) implementsZarazConfigUpdateResponseTriggersExcludeRule() {
}

type ZarazConfigUpdateResponseTriggersExcludeRulesZarazTimerRuleAction string

const (
	ZarazConfigUpdateResponseTriggersExcludeRulesZarazTimerRuleActionTimer ZarazConfigUpdateResponseTriggersExcludeRulesZarazTimerRuleAction = "timer"
)

type ZarazConfigUpdateResponseTriggersExcludeRulesZarazTimerRuleSettings struct {
	Interval int64                                                                   `json:"interval,required"`
	Limit    int64                                                                   `json:"limit,required"`
	JSON     zarazConfigUpdateResponseTriggersExcludeRulesZarazTimerRuleSettingsJSON `json:"-"`
}

// zarazConfigUpdateResponseTriggersExcludeRulesZarazTimerRuleSettingsJSON contains
// the JSON metadata for the struct
// [ZarazConfigUpdateResponseTriggersExcludeRulesZarazTimerRuleSettings]
type zarazConfigUpdateResponseTriggersExcludeRulesZarazTimerRuleSettingsJSON struct {
	Interval    apijson.Field
	Limit       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseTriggersExcludeRulesZarazTimerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigUpdateResponseTriggersExcludeRulesZarazFormSubmissionRule struct {
	ID       string                                                                       `json:"id,required"`
	Action   ZarazConfigUpdateResponseTriggersExcludeRulesZarazFormSubmissionRuleAction   `json:"action,required"`
	Settings ZarazConfigUpdateResponseTriggersExcludeRulesZarazFormSubmissionRuleSettings `json:"settings,required"`
	JSON     zarazConfigUpdateResponseTriggersExcludeRulesZarazFormSubmissionRuleJSON     `json:"-"`
}

// zarazConfigUpdateResponseTriggersExcludeRulesZarazFormSubmissionRuleJSON
// contains the JSON metadata for the struct
// [ZarazConfigUpdateResponseTriggersExcludeRulesZarazFormSubmissionRule]
type zarazConfigUpdateResponseTriggersExcludeRulesZarazFormSubmissionRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseTriggersExcludeRulesZarazFormSubmissionRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigUpdateResponseTriggersExcludeRulesZarazFormSubmissionRule) implementsZarazConfigUpdateResponseTriggersExcludeRule() {
}

type ZarazConfigUpdateResponseTriggersExcludeRulesZarazFormSubmissionRuleAction string

const (
	ZarazConfigUpdateResponseTriggersExcludeRulesZarazFormSubmissionRuleActionFormSubmission ZarazConfigUpdateResponseTriggersExcludeRulesZarazFormSubmissionRuleAction = "formSubmission"
)

type ZarazConfigUpdateResponseTriggersExcludeRulesZarazFormSubmissionRuleSettings struct {
	Selector string                                                                           `json:"selector,required"`
	Validate bool                                                                             `json:"validate,required"`
	JSON     zarazConfigUpdateResponseTriggersExcludeRulesZarazFormSubmissionRuleSettingsJSON `json:"-"`
}

// zarazConfigUpdateResponseTriggersExcludeRulesZarazFormSubmissionRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazConfigUpdateResponseTriggersExcludeRulesZarazFormSubmissionRuleSettings]
type zarazConfigUpdateResponseTriggersExcludeRulesZarazFormSubmissionRuleSettingsJSON struct {
	Selector    apijson.Field
	Validate    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseTriggersExcludeRulesZarazFormSubmissionRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigUpdateResponseTriggersExcludeRulesZarazVariableMatchRule struct {
	ID       string                                                                      `json:"id,required"`
	Action   ZarazConfigUpdateResponseTriggersExcludeRulesZarazVariableMatchRuleAction   `json:"action,required"`
	Settings ZarazConfigUpdateResponseTriggersExcludeRulesZarazVariableMatchRuleSettings `json:"settings,required"`
	JSON     zarazConfigUpdateResponseTriggersExcludeRulesZarazVariableMatchRuleJSON     `json:"-"`
}

// zarazConfigUpdateResponseTriggersExcludeRulesZarazVariableMatchRuleJSON contains
// the JSON metadata for the struct
// [ZarazConfigUpdateResponseTriggersExcludeRulesZarazVariableMatchRule]
type zarazConfigUpdateResponseTriggersExcludeRulesZarazVariableMatchRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseTriggersExcludeRulesZarazVariableMatchRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigUpdateResponseTriggersExcludeRulesZarazVariableMatchRule) implementsZarazConfigUpdateResponseTriggersExcludeRule() {
}

type ZarazConfigUpdateResponseTriggersExcludeRulesZarazVariableMatchRuleAction string

const (
	ZarazConfigUpdateResponseTriggersExcludeRulesZarazVariableMatchRuleActionVariableMatch ZarazConfigUpdateResponseTriggersExcludeRulesZarazVariableMatchRuleAction = "variableMatch"
)

type ZarazConfigUpdateResponseTriggersExcludeRulesZarazVariableMatchRuleSettings struct {
	Match    string                                                                          `json:"match,required"`
	Variable string                                                                          `json:"variable,required"`
	JSON     zarazConfigUpdateResponseTriggersExcludeRulesZarazVariableMatchRuleSettingsJSON `json:"-"`
}

// zarazConfigUpdateResponseTriggersExcludeRulesZarazVariableMatchRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazConfigUpdateResponseTriggersExcludeRulesZarazVariableMatchRuleSettings]
type zarazConfigUpdateResponseTriggersExcludeRulesZarazVariableMatchRuleSettingsJSON struct {
	Match       apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseTriggersExcludeRulesZarazVariableMatchRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigUpdateResponseTriggersExcludeRulesZarazScrollDepthRule struct {
	ID       string                                                                    `json:"id,required"`
	Action   ZarazConfigUpdateResponseTriggersExcludeRulesZarazScrollDepthRuleAction   `json:"action,required"`
	Settings ZarazConfigUpdateResponseTriggersExcludeRulesZarazScrollDepthRuleSettings `json:"settings,required"`
	JSON     zarazConfigUpdateResponseTriggersExcludeRulesZarazScrollDepthRuleJSON     `json:"-"`
}

// zarazConfigUpdateResponseTriggersExcludeRulesZarazScrollDepthRuleJSON contains
// the JSON metadata for the struct
// [ZarazConfigUpdateResponseTriggersExcludeRulesZarazScrollDepthRule]
type zarazConfigUpdateResponseTriggersExcludeRulesZarazScrollDepthRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseTriggersExcludeRulesZarazScrollDepthRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigUpdateResponseTriggersExcludeRulesZarazScrollDepthRule) implementsZarazConfigUpdateResponseTriggersExcludeRule() {
}

type ZarazConfigUpdateResponseTriggersExcludeRulesZarazScrollDepthRuleAction string

const (
	ZarazConfigUpdateResponseTriggersExcludeRulesZarazScrollDepthRuleActionScrollDepth ZarazConfigUpdateResponseTriggersExcludeRulesZarazScrollDepthRuleAction = "scrollDepth"
)

type ZarazConfigUpdateResponseTriggersExcludeRulesZarazScrollDepthRuleSettings struct {
	Positions string                                                                        `json:"positions,required"`
	JSON      zarazConfigUpdateResponseTriggersExcludeRulesZarazScrollDepthRuleSettingsJSON `json:"-"`
}

// zarazConfigUpdateResponseTriggersExcludeRulesZarazScrollDepthRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazConfigUpdateResponseTriggersExcludeRulesZarazScrollDepthRuleSettings]
type zarazConfigUpdateResponseTriggersExcludeRulesZarazScrollDepthRuleSettingsJSON struct {
	Positions   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseTriggersExcludeRulesZarazScrollDepthRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigUpdateResponseTriggersExcludeRulesZarazElementVisibilityRule struct {
	ID       string                                                                          `json:"id,required"`
	Action   ZarazConfigUpdateResponseTriggersExcludeRulesZarazElementVisibilityRuleAction   `json:"action,required"`
	Settings ZarazConfigUpdateResponseTriggersExcludeRulesZarazElementVisibilityRuleSettings `json:"settings,required"`
	JSON     zarazConfigUpdateResponseTriggersExcludeRulesZarazElementVisibilityRuleJSON     `json:"-"`
}

// zarazConfigUpdateResponseTriggersExcludeRulesZarazElementVisibilityRuleJSON
// contains the JSON metadata for the struct
// [ZarazConfigUpdateResponseTriggersExcludeRulesZarazElementVisibilityRule]
type zarazConfigUpdateResponseTriggersExcludeRulesZarazElementVisibilityRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseTriggersExcludeRulesZarazElementVisibilityRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigUpdateResponseTriggersExcludeRulesZarazElementVisibilityRule) implementsZarazConfigUpdateResponseTriggersExcludeRule() {
}

type ZarazConfigUpdateResponseTriggersExcludeRulesZarazElementVisibilityRuleAction string

const (
	ZarazConfigUpdateResponseTriggersExcludeRulesZarazElementVisibilityRuleActionElementVisibility ZarazConfigUpdateResponseTriggersExcludeRulesZarazElementVisibilityRuleAction = "elementVisibility"
)

type ZarazConfigUpdateResponseTriggersExcludeRulesZarazElementVisibilityRuleSettings struct {
	Selector string                                                                              `json:"selector,required"`
	JSON     zarazConfigUpdateResponseTriggersExcludeRulesZarazElementVisibilityRuleSettingsJSON `json:"-"`
}

// zarazConfigUpdateResponseTriggersExcludeRulesZarazElementVisibilityRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazConfigUpdateResponseTriggersExcludeRulesZarazElementVisibilityRuleSettings]
type zarazConfigUpdateResponseTriggersExcludeRulesZarazElementVisibilityRuleSettingsJSON struct {
	Selector    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseTriggersExcludeRulesZarazElementVisibilityRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [ZarazConfigUpdateResponseTriggersLoadRulesZarazLoadRule],
// [ZarazConfigUpdateResponseTriggersLoadRulesZarazClickListenerRule],
// [ZarazConfigUpdateResponseTriggersLoadRulesZarazTimerRule],
// [ZarazConfigUpdateResponseTriggersLoadRulesZarazFormSubmissionRule],
// [ZarazConfigUpdateResponseTriggersLoadRulesZarazVariableMatchRule],
// [ZarazConfigUpdateResponseTriggersLoadRulesZarazScrollDepthRule] or
// [ZarazConfigUpdateResponseTriggersLoadRulesZarazElementVisibilityRule].
type ZarazConfigUpdateResponseTriggersLoadRule interface {
	implementsZarazConfigUpdateResponseTriggersLoadRule()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZarazConfigUpdateResponseTriggersLoadRule)(nil)).Elem(), "")
}

type ZarazConfigUpdateResponseTriggersLoadRulesZarazLoadRule struct {
	ID    string                                                      `json:"id,required"`
	Match string                                                      `json:"match,required"`
	Op    ZarazConfigUpdateResponseTriggersLoadRulesZarazLoadRuleOp   `json:"op,required"`
	Value string                                                      `json:"value,required"`
	JSON  zarazConfigUpdateResponseTriggersLoadRulesZarazLoadRuleJSON `json:"-"`
}

// zarazConfigUpdateResponseTriggersLoadRulesZarazLoadRuleJSON contains the JSON
// metadata for the struct
// [ZarazConfigUpdateResponseTriggersLoadRulesZarazLoadRule]
type zarazConfigUpdateResponseTriggersLoadRulesZarazLoadRuleJSON struct {
	ID          apijson.Field
	Match       apijson.Field
	Op          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseTriggersLoadRulesZarazLoadRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigUpdateResponseTriggersLoadRulesZarazLoadRule) implementsZarazConfigUpdateResponseTriggersLoadRule() {
}

type ZarazConfigUpdateResponseTriggersLoadRulesZarazLoadRuleOp string

const (
	ZarazConfigUpdateResponseTriggersLoadRulesZarazLoadRuleOpContains           ZarazConfigUpdateResponseTriggersLoadRulesZarazLoadRuleOp = "CONTAINS"
	ZarazConfigUpdateResponseTriggersLoadRulesZarazLoadRuleOpEquals             ZarazConfigUpdateResponseTriggersLoadRulesZarazLoadRuleOp = "EQUALS"
	ZarazConfigUpdateResponseTriggersLoadRulesZarazLoadRuleOpStartsWith         ZarazConfigUpdateResponseTriggersLoadRulesZarazLoadRuleOp = "STARTS_WITH"
	ZarazConfigUpdateResponseTriggersLoadRulesZarazLoadRuleOpEndsWith           ZarazConfigUpdateResponseTriggersLoadRulesZarazLoadRuleOp = "ENDS_WITH"
	ZarazConfigUpdateResponseTriggersLoadRulesZarazLoadRuleOpMatchRegex         ZarazConfigUpdateResponseTriggersLoadRulesZarazLoadRuleOp = "MATCH_REGEX"
	ZarazConfigUpdateResponseTriggersLoadRulesZarazLoadRuleOpNotMatchRegex      ZarazConfigUpdateResponseTriggersLoadRulesZarazLoadRuleOp = "NOT_MATCH_REGEX"
	ZarazConfigUpdateResponseTriggersLoadRulesZarazLoadRuleOpGreaterThan        ZarazConfigUpdateResponseTriggersLoadRulesZarazLoadRuleOp = "GREATER_THAN"
	ZarazConfigUpdateResponseTriggersLoadRulesZarazLoadRuleOpGreaterThanOrEqual ZarazConfigUpdateResponseTriggersLoadRulesZarazLoadRuleOp = "GREATER_THAN_OR_EQUAL"
	ZarazConfigUpdateResponseTriggersLoadRulesZarazLoadRuleOpLessThan           ZarazConfigUpdateResponseTriggersLoadRulesZarazLoadRuleOp = "LESS_THAN"
	ZarazConfigUpdateResponseTriggersLoadRulesZarazLoadRuleOpLessThanOrEqual    ZarazConfigUpdateResponseTriggersLoadRulesZarazLoadRuleOp = "LESS_THAN_OR_EQUAL"
)

type ZarazConfigUpdateResponseTriggersLoadRulesZarazClickListenerRule struct {
	ID       string                                                                   `json:"id,required"`
	Action   ZarazConfigUpdateResponseTriggersLoadRulesZarazClickListenerRuleAction   `json:"action,required"`
	Settings ZarazConfigUpdateResponseTriggersLoadRulesZarazClickListenerRuleSettings `json:"settings,required"`
	JSON     zarazConfigUpdateResponseTriggersLoadRulesZarazClickListenerRuleJSON     `json:"-"`
}

// zarazConfigUpdateResponseTriggersLoadRulesZarazClickListenerRuleJSON contains
// the JSON metadata for the struct
// [ZarazConfigUpdateResponseTriggersLoadRulesZarazClickListenerRule]
type zarazConfigUpdateResponseTriggersLoadRulesZarazClickListenerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseTriggersLoadRulesZarazClickListenerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigUpdateResponseTriggersLoadRulesZarazClickListenerRule) implementsZarazConfigUpdateResponseTriggersLoadRule() {
}

type ZarazConfigUpdateResponseTriggersLoadRulesZarazClickListenerRuleAction string

const (
	ZarazConfigUpdateResponseTriggersLoadRulesZarazClickListenerRuleActionClickListener ZarazConfigUpdateResponseTriggersLoadRulesZarazClickListenerRuleAction = "clickListener"
)

type ZarazConfigUpdateResponseTriggersLoadRulesZarazClickListenerRuleSettings struct {
	Selector    string                                                                       `json:"selector,required"`
	Type        ZarazConfigUpdateResponseTriggersLoadRulesZarazClickListenerRuleSettingsType `json:"type,required"`
	WaitForTags int64                                                                        `json:"waitForTags,required"`
	JSON        zarazConfigUpdateResponseTriggersLoadRulesZarazClickListenerRuleSettingsJSON `json:"-"`
}

// zarazConfigUpdateResponseTriggersLoadRulesZarazClickListenerRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazConfigUpdateResponseTriggersLoadRulesZarazClickListenerRuleSettings]
type zarazConfigUpdateResponseTriggersLoadRulesZarazClickListenerRuleSettingsJSON struct {
	Selector    apijson.Field
	Type        apijson.Field
	WaitForTags apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseTriggersLoadRulesZarazClickListenerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigUpdateResponseTriggersLoadRulesZarazClickListenerRuleSettingsType string

const (
	ZarazConfigUpdateResponseTriggersLoadRulesZarazClickListenerRuleSettingsTypeXpath ZarazConfigUpdateResponseTriggersLoadRulesZarazClickListenerRuleSettingsType = "xpath"
	ZarazConfigUpdateResponseTriggersLoadRulesZarazClickListenerRuleSettingsTypeCss   ZarazConfigUpdateResponseTriggersLoadRulesZarazClickListenerRuleSettingsType = "css"
)

type ZarazConfigUpdateResponseTriggersLoadRulesZarazTimerRule struct {
	ID       string                                                           `json:"id,required"`
	Action   ZarazConfigUpdateResponseTriggersLoadRulesZarazTimerRuleAction   `json:"action,required"`
	Settings ZarazConfigUpdateResponseTriggersLoadRulesZarazTimerRuleSettings `json:"settings,required"`
	JSON     zarazConfigUpdateResponseTriggersLoadRulesZarazTimerRuleJSON     `json:"-"`
}

// zarazConfigUpdateResponseTriggersLoadRulesZarazTimerRuleJSON contains the JSON
// metadata for the struct
// [ZarazConfigUpdateResponseTriggersLoadRulesZarazTimerRule]
type zarazConfigUpdateResponseTriggersLoadRulesZarazTimerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseTriggersLoadRulesZarazTimerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigUpdateResponseTriggersLoadRulesZarazTimerRule) implementsZarazConfigUpdateResponseTriggersLoadRule() {
}

type ZarazConfigUpdateResponseTriggersLoadRulesZarazTimerRuleAction string

const (
	ZarazConfigUpdateResponseTriggersLoadRulesZarazTimerRuleActionTimer ZarazConfigUpdateResponseTriggersLoadRulesZarazTimerRuleAction = "timer"
)

type ZarazConfigUpdateResponseTriggersLoadRulesZarazTimerRuleSettings struct {
	Interval int64                                                                `json:"interval,required"`
	Limit    int64                                                                `json:"limit,required"`
	JSON     zarazConfigUpdateResponseTriggersLoadRulesZarazTimerRuleSettingsJSON `json:"-"`
}

// zarazConfigUpdateResponseTriggersLoadRulesZarazTimerRuleSettingsJSON contains
// the JSON metadata for the struct
// [ZarazConfigUpdateResponseTriggersLoadRulesZarazTimerRuleSettings]
type zarazConfigUpdateResponseTriggersLoadRulesZarazTimerRuleSettingsJSON struct {
	Interval    apijson.Field
	Limit       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseTriggersLoadRulesZarazTimerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigUpdateResponseTriggersLoadRulesZarazFormSubmissionRule struct {
	ID       string                                                                    `json:"id,required"`
	Action   ZarazConfigUpdateResponseTriggersLoadRulesZarazFormSubmissionRuleAction   `json:"action,required"`
	Settings ZarazConfigUpdateResponseTriggersLoadRulesZarazFormSubmissionRuleSettings `json:"settings,required"`
	JSON     zarazConfigUpdateResponseTriggersLoadRulesZarazFormSubmissionRuleJSON     `json:"-"`
}

// zarazConfigUpdateResponseTriggersLoadRulesZarazFormSubmissionRuleJSON contains
// the JSON metadata for the struct
// [ZarazConfigUpdateResponseTriggersLoadRulesZarazFormSubmissionRule]
type zarazConfigUpdateResponseTriggersLoadRulesZarazFormSubmissionRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseTriggersLoadRulesZarazFormSubmissionRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigUpdateResponseTriggersLoadRulesZarazFormSubmissionRule) implementsZarazConfigUpdateResponseTriggersLoadRule() {
}

type ZarazConfigUpdateResponseTriggersLoadRulesZarazFormSubmissionRuleAction string

const (
	ZarazConfigUpdateResponseTriggersLoadRulesZarazFormSubmissionRuleActionFormSubmission ZarazConfigUpdateResponseTriggersLoadRulesZarazFormSubmissionRuleAction = "formSubmission"
)

type ZarazConfigUpdateResponseTriggersLoadRulesZarazFormSubmissionRuleSettings struct {
	Selector string                                                                        `json:"selector,required"`
	Validate bool                                                                          `json:"validate,required"`
	JSON     zarazConfigUpdateResponseTriggersLoadRulesZarazFormSubmissionRuleSettingsJSON `json:"-"`
}

// zarazConfigUpdateResponseTriggersLoadRulesZarazFormSubmissionRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazConfigUpdateResponseTriggersLoadRulesZarazFormSubmissionRuleSettings]
type zarazConfigUpdateResponseTriggersLoadRulesZarazFormSubmissionRuleSettingsJSON struct {
	Selector    apijson.Field
	Validate    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseTriggersLoadRulesZarazFormSubmissionRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigUpdateResponseTriggersLoadRulesZarazVariableMatchRule struct {
	ID       string                                                                   `json:"id,required"`
	Action   ZarazConfigUpdateResponseTriggersLoadRulesZarazVariableMatchRuleAction   `json:"action,required"`
	Settings ZarazConfigUpdateResponseTriggersLoadRulesZarazVariableMatchRuleSettings `json:"settings,required"`
	JSON     zarazConfigUpdateResponseTriggersLoadRulesZarazVariableMatchRuleJSON     `json:"-"`
}

// zarazConfigUpdateResponseTriggersLoadRulesZarazVariableMatchRuleJSON contains
// the JSON metadata for the struct
// [ZarazConfigUpdateResponseTriggersLoadRulesZarazVariableMatchRule]
type zarazConfigUpdateResponseTriggersLoadRulesZarazVariableMatchRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseTriggersLoadRulesZarazVariableMatchRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigUpdateResponseTriggersLoadRulesZarazVariableMatchRule) implementsZarazConfigUpdateResponseTriggersLoadRule() {
}

type ZarazConfigUpdateResponseTriggersLoadRulesZarazVariableMatchRuleAction string

const (
	ZarazConfigUpdateResponseTriggersLoadRulesZarazVariableMatchRuleActionVariableMatch ZarazConfigUpdateResponseTriggersLoadRulesZarazVariableMatchRuleAction = "variableMatch"
)

type ZarazConfigUpdateResponseTriggersLoadRulesZarazVariableMatchRuleSettings struct {
	Match    string                                                                       `json:"match,required"`
	Variable string                                                                       `json:"variable,required"`
	JSON     zarazConfigUpdateResponseTriggersLoadRulesZarazVariableMatchRuleSettingsJSON `json:"-"`
}

// zarazConfigUpdateResponseTriggersLoadRulesZarazVariableMatchRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazConfigUpdateResponseTriggersLoadRulesZarazVariableMatchRuleSettings]
type zarazConfigUpdateResponseTriggersLoadRulesZarazVariableMatchRuleSettingsJSON struct {
	Match       apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseTriggersLoadRulesZarazVariableMatchRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigUpdateResponseTriggersLoadRulesZarazScrollDepthRule struct {
	ID       string                                                                 `json:"id,required"`
	Action   ZarazConfigUpdateResponseTriggersLoadRulesZarazScrollDepthRuleAction   `json:"action,required"`
	Settings ZarazConfigUpdateResponseTriggersLoadRulesZarazScrollDepthRuleSettings `json:"settings,required"`
	JSON     zarazConfigUpdateResponseTriggersLoadRulesZarazScrollDepthRuleJSON     `json:"-"`
}

// zarazConfigUpdateResponseTriggersLoadRulesZarazScrollDepthRuleJSON contains the
// JSON metadata for the struct
// [ZarazConfigUpdateResponseTriggersLoadRulesZarazScrollDepthRule]
type zarazConfigUpdateResponseTriggersLoadRulesZarazScrollDepthRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseTriggersLoadRulesZarazScrollDepthRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigUpdateResponseTriggersLoadRulesZarazScrollDepthRule) implementsZarazConfigUpdateResponseTriggersLoadRule() {
}

type ZarazConfigUpdateResponseTriggersLoadRulesZarazScrollDepthRuleAction string

const (
	ZarazConfigUpdateResponseTriggersLoadRulesZarazScrollDepthRuleActionScrollDepth ZarazConfigUpdateResponseTriggersLoadRulesZarazScrollDepthRuleAction = "scrollDepth"
)

type ZarazConfigUpdateResponseTriggersLoadRulesZarazScrollDepthRuleSettings struct {
	Positions string                                                                     `json:"positions,required"`
	JSON      zarazConfigUpdateResponseTriggersLoadRulesZarazScrollDepthRuleSettingsJSON `json:"-"`
}

// zarazConfigUpdateResponseTriggersLoadRulesZarazScrollDepthRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazConfigUpdateResponseTriggersLoadRulesZarazScrollDepthRuleSettings]
type zarazConfigUpdateResponseTriggersLoadRulesZarazScrollDepthRuleSettingsJSON struct {
	Positions   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseTriggersLoadRulesZarazScrollDepthRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigUpdateResponseTriggersLoadRulesZarazElementVisibilityRule struct {
	ID       string                                                                       `json:"id,required"`
	Action   ZarazConfigUpdateResponseTriggersLoadRulesZarazElementVisibilityRuleAction   `json:"action,required"`
	Settings ZarazConfigUpdateResponseTriggersLoadRulesZarazElementVisibilityRuleSettings `json:"settings,required"`
	JSON     zarazConfigUpdateResponseTriggersLoadRulesZarazElementVisibilityRuleJSON     `json:"-"`
}

// zarazConfigUpdateResponseTriggersLoadRulesZarazElementVisibilityRuleJSON
// contains the JSON metadata for the struct
// [ZarazConfigUpdateResponseTriggersLoadRulesZarazElementVisibilityRule]
type zarazConfigUpdateResponseTriggersLoadRulesZarazElementVisibilityRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseTriggersLoadRulesZarazElementVisibilityRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigUpdateResponseTriggersLoadRulesZarazElementVisibilityRule) implementsZarazConfigUpdateResponseTriggersLoadRule() {
}

type ZarazConfigUpdateResponseTriggersLoadRulesZarazElementVisibilityRuleAction string

const (
	ZarazConfigUpdateResponseTriggersLoadRulesZarazElementVisibilityRuleActionElementVisibility ZarazConfigUpdateResponseTriggersLoadRulesZarazElementVisibilityRuleAction = "elementVisibility"
)

type ZarazConfigUpdateResponseTriggersLoadRulesZarazElementVisibilityRuleSettings struct {
	Selector string                                                                           `json:"selector,required"`
	JSON     zarazConfigUpdateResponseTriggersLoadRulesZarazElementVisibilityRuleSettingsJSON `json:"-"`
}

// zarazConfigUpdateResponseTriggersLoadRulesZarazElementVisibilityRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazConfigUpdateResponseTriggersLoadRulesZarazElementVisibilityRuleSettings]
type zarazConfigUpdateResponseTriggersLoadRulesZarazElementVisibilityRuleSettingsJSON struct {
	Selector    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseTriggersLoadRulesZarazElementVisibilityRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigUpdateResponseTriggersSystem string

const (
	ZarazConfigUpdateResponseTriggersSystemPageload ZarazConfigUpdateResponseTriggersSystem = "pageload"
)

// Union satisfied by [ZarazConfigUpdateResponseVariablesObject] or
// [ZarazConfigUpdateResponseVariablesObject].
type ZarazConfigUpdateResponseVariable interface {
	implementsZarazConfigUpdateResponseVariable()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZarazConfigUpdateResponseVariable)(nil)).Elem(), "")
}

type ZarazConfigUpdateResponseVariablesObject struct {
	Name  string                                       `json:"name,required"`
	Type  ZarazConfigUpdateResponseVariablesObjectType `json:"type,required"`
	Value string                                       `json:"value,required"`
	JSON  zarazConfigUpdateResponseVariablesObjectJSON `json:"-"`
}

// zarazConfigUpdateResponseVariablesObjectJSON contains the JSON metadata for the
// struct [ZarazConfigUpdateResponseVariablesObject]
type zarazConfigUpdateResponseVariablesObjectJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseVariablesObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazConfigUpdateResponseVariablesObject) implementsZarazConfigUpdateResponseVariable() {}

type ZarazConfigUpdateResponseVariablesObjectType string

const (
	ZarazConfigUpdateResponseVariablesObjectTypeString ZarazConfigUpdateResponseVariablesObjectType = "string"
	ZarazConfigUpdateResponseVariablesObjectTypeSecret ZarazConfigUpdateResponseVariablesObjectType = "secret"
)

// Consent management configuration.
type ZarazConfigUpdateResponseConsent struct {
	Enabled                bool                                                   `json:"enabled,required"`
	ButtonTextTranslations ZarazConfigUpdateResponseConsentButtonTextTranslations `json:"buttonTextTranslations"`
	CompanyEmail           string                                                 `json:"companyEmail"`
	CompanyName            string                                                 `json:"companyName"`
	CompanyStreetAddress   string                                                 `json:"companyStreetAddress"`
	ConsentModalIntroHTML  string                                                 `json:"consentModalIntroHTML"`
	// Object where keys are language codes
	ConsentModalIntroHTMLWithTranslations map[string]string `json:"consentModalIntroHTMLWithTranslations"`
	CookieName                            string            `json:"cookieName"`
	CustomCss                             string            `json:"customCSS"`
	CustomIntroDisclaimerDismissed        bool              `json:"customIntroDisclaimerDismissed"`
	DefaultLanguage                       string            `json:"defaultLanguage"`
	HideModal                             bool              `json:"hideModal"`
	// Object where keys are purpose alpha-numeric IDs
	Purposes map[string]ZarazConfigUpdateResponseConsentPurpose `json:"purposes"`
	// Object where keys are purpose alpha-numeric IDs
	PurposesWithTranslations map[string]ZarazConfigUpdateResponseConsentPurposesWithTranslation `json:"purposesWithTranslations"`
	JSON                     zarazConfigUpdateResponseConsentJSON                               `json:"-"`
}

// zarazConfigUpdateResponseConsentJSON contains the JSON metadata for the struct
// [ZarazConfigUpdateResponseConsent]
type zarazConfigUpdateResponseConsentJSON struct {
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

func (r *ZarazConfigUpdateResponseConsent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigUpdateResponseConsentButtonTextTranslations struct {
	// Object where keys are language codes
	AcceptAll map[string]string `json:"accept_all,required"`
	// Object where keys are language codes
	ConfirmMyChoices map[string]string `json:"confirm_my_choices,required"`
	// Object where keys are language codes
	RejectAll map[string]string                                          `json:"reject_all,required"`
	JSON      zarazConfigUpdateResponseConsentButtonTextTranslationsJSON `json:"-"`
}

// zarazConfigUpdateResponseConsentButtonTextTranslationsJSON contains the JSON
// metadata for the struct [ZarazConfigUpdateResponseConsentButtonTextTranslations]
type zarazConfigUpdateResponseConsentButtonTextTranslationsJSON struct {
	AcceptAll        apijson.Field
	ConfirmMyChoices apijson.Field
	RejectAll        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseConsentButtonTextTranslations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigUpdateResponseConsentPurpose struct {
	Description string                                      `json:"description,required"`
	Name        string                                      `json:"name,required"`
	JSON        zarazConfigUpdateResponseConsentPurposeJSON `json:"-"`
}

// zarazConfigUpdateResponseConsentPurposeJSON contains the JSON metadata for the
// struct [ZarazConfigUpdateResponseConsentPurpose]
type zarazConfigUpdateResponseConsentPurposeJSON struct {
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseConsentPurpose) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigUpdateResponseConsentPurposesWithTranslation struct {
	// Object where keys are language codes
	Description map[string]string `json:"description,required"`
	// Object where keys are language codes
	Name  map[string]string                                           `json:"name,required"`
	Order int64                                                       `json:"order,required"`
	JSON  zarazConfigUpdateResponseConsentPurposesWithTranslationJSON `json:"-"`
}

// zarazConfigUpdateResponseConsentPurposesWithTranslationJSON contains the JSON
// metadata for the struct
// [ZarazConfigUpdateResponseConsentPurposesWithTranslation]
type zarazConfigUpdateResponseConsentPurposesWithTranslationJSON struct {
	Description apijson.Field
	Name        apijson.Field
	Order       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseConsentPurposesWithTranslation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

type ZarazConfigUpdateParams struct {
	// Data layer compatibility mode enabled.
	DataLayer param.Field[bool] `json:"dataLayer,required"`
	// The key for Zaraz debug mode.
	DebugKey param.Field[string] `json:"debugKey,required"`
	// General Zaraz settings.
	Settings param.Field[ZarazConfigUpdateParamsSettings] `json:"settings,required"`
	// Tools set up under Zaraz configuration, where key is the alpha-numeric tool ID
	// and value is the tool configuration object.
	Tools param.Field[map[string]ZarazConfigUpdateParamsTools] `json:"tools,required"`
	// Triggers set up under Zaraz configuration, where key is the trigger
	// alpha-numeric ID and value is the trigger configuration.
	Triggers param.Field[map[string]ZarazConfigUpdateParamsTriggers] `json:"triggers,required"`
	// Variables set up under Zaraz configuration, where key is the variable
	// alpha-numeric ID and value is the variable configuration. Values of variables of
	// type secret are not included.
	Variables param.Field[map[string]ZarazConfigUpdateParamsVariables] `json:"variables,required"`
	// Zaraz internal version of the config.
	ZarazVersion param.Field[int64] `json:"zarazVersion,required"`
	// Consent management configuration.
	Consent param.Field[ZarazConfigUpdateParamsConsent] `json:"consent"`
	// Single Page Application support enabled.
	HistoryChange param.Field[bool] `json:"historyChange"`
}

func (r ZarazConfigUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// General Zaraz settings.
type ZarazConfigUpdateParamsSettings struct {
	// Automatic injection of Zaraz scripts enabled.
	AutoInjectScript param.Field[bool] `json:"autoInjectScript,required"`
	// Details of the worker that receives and edits Zaraz Context object.
	ContextEnricher param.Field[ZarazConfigUpdateParamsSettingsContextEnricher] `json:"contextEnricher"`
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

func (r ZarazConfigUpdateParamsSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Details of the worker that receives and edits Zaraz Context object.
type ZarazConfigUpdateParamsSettingsContextEnricher struct {
	EscapedWorkerName param.Field[string] `json:"escapedWorkerName,required"`
	WorkerTag         param.Field[string] `json:"workerTag,required"`
}

func (r ZarazConfigUpdateParamsSettingsContextEnricher) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [ZarazConfigUpdateParamsToolsZarazLegacyTool],
// [ZarazConfigUpdateParamsToolsZarazManagedComponent],
// [ZarazConfigUpdateParamsToolsZarazCustomManagedComponent].
type ZarazConfigUpdateParamsTools interface {
	implementsZarazConfigUpdateParamsTools()
}

type ZarazConfigUpdateParamsToolsZarazLegacyTool struct {
	// List of blocking trigger IDs
	BlockingTriggers param.Field[[]string] `json:"blockingTriggers,required"`
	// Default fields for tool's actions
	DefaultFields param.Field[map[string]ZarazConfigUpdateParamsToolsZarazLegacyToolDefaultFields] `json:"defaultFields,required"`
	// Whether tool is enabled
	Enabled param.Field[bool] `json:"enabled,required"`
	// Tool's internal name
	Library param.Field[string] `json:"library,required"`
	// Tool's name defined by the user
	Name param.Field[string] `json:"name,required"`
	// List of actions configured on a tool
	NeoEvents param.Field[[]ZarazConfigUpdateParamsToolsZarazLegacyToolNeoEvent] `json:"neoEvents,required"`
	Type      param.Field[ZarazConfigUpdateParamsToolsZarazLegacyToolType]       `json:"type,required"`
	// Default consent purpose ID
	DefaultPurpose param.Field[string] `json:"defaultPurpose"`
}

func (r ZarazConfigUpdateParamsToolsZarazLegacyTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZarazConfigUpdateParamsToolsZarazLegacyTool) implementsZarazConfigUpdateParamsTools() {}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type ZarazConfigUpdateParamsToolsZarazLegacyToolDefaultFields interface {
	ImplementsZarazConfigUpdateParamsToolsZarazLegacyToolDefaultFields()
}

type ZarazConfigUpdateParamsToolsZarazLegacyToolNeoEvent struct {
	// List of blocking triggers IDs
	BlockingTriggers param.Field[[]string] `json:"blockingTriggers,required"`
	// Event payload
	Data param.Field[interface{}] `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers param.Field[[]string] `json:"firingTriggers,required"`
}

func (r ZarazConfigUpdateParamsToolsZarazLegacyToolNeoEvent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigUpdateParamsToolsZarazLegacyToolType string

const (
	ZarazConfigUpdateParamsToolsZarazLegacyToolTypeLibrary ZarazConfigUpdateParamsToolsZarazLegacyToolType = "library"
)

type ZarazConfigUpdateParamsToolsZarazManagedComponent struct {
	// List of blocking trigger IDs
	BlockingTriggers param.Field[[]string] `json:"blockingTriggers,required"`
	// Tool's internal name
	Component param.Field[string] `json:"component,required"`
	// Default fields for tool's actions
	DefaultFields param.Field[map[string]ZarazConfigUpdateParamsToolsZarazManagedComponentDefaultFields] `json:"defaultFields,required"`
	// Whether tool is enabled
	Enabled param.Field[bool] `json:"enabled,required"`
	// Tool's name defined by the user
	Name param.Field[string] `json:"name,required"`
	// List of permissions granted to the component
	Permissions param.Field[[]string] `json:"permissions,required"`
	// Tool's settings
	Settings param.Field[map[string]ZarazConfigUpdateParamsToolsZarazManagedComponentSettings] `json:"settings,required"`
	Type     param.Field[ZarazConfigUpdateParamsToolsZarazManagedComponentType]                `json:"type,required"`
	// Actions configured on a tool. Either this or neoEvents field is required.
	Actions param.Field[map[string]ZarazConfigUpdateParamsToolsZarazManagedComponentActions] `json:"actions"`
	// Default consent purpose ID
	DefaultPurpose param.Field[string] `json:"defaultPurpose"`
	// DEPRECATED - List of actions configured on a tool. Either this or actions field
	// is required. If both are present, actions field will take precedence.
	NeoEvents param.Field[[]ZarazConfigUpdateParamsToolsZarazManagedComponentNeoEvent] `json:"neoEvents"`
}

func (r ZarazConfigUpdateParamsToolsZarazManagedComponent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZarazConfigUpdateParamsToolsZarazManagedComponent) implementsZarazConfigUpdateParamsTools() {}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type ZarazConfigUpdateParamsToolsZarazManagedComponentDefaultFields interface {
	ImplementsZarazConfigUpdateParamsToolsZarazManagedComponentDefaultFields()
}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type ZarazConfigUpdateParamsToolsZarazManagedComponentSettings interface {
	ImplementsZarazConfigUpdateParamsToolsZarazManagedComponentSettings()
}

type ZarazConfigUpdateParamsToolsZarazManagedComponentType string

const (
	ZarazConfigUpdateParamsToolsZarazManagedComponentTypeComponent ZarazConfigUpdateParamsToolsZarazManagedComponentType = "component"
)

type ZarazConfigUpdateParamsToolsZarazManagedComponentActions struct {
	// Tool event type
	ActionType param.Field[string] `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers param.Field[[]string] `json:"blockingTriggers,required"`
	// Event payload
	Data param.Field[interface{}] `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers param.Field[[]string] `json:"firingTriggers,required"`
}

func (r ZarazConfigUpdateParamsToolsZarazManagedComponentActions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigUpdateParamsToolsZarazManagedComponentNeoEvent struct {
	// Tool event type
	ActionType param.Field[string] `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers param.Field[[]string] `json:"blockingTriggers,required"`
	// Event payload
	Data param.Field[interface{}] `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers param.Field[[]string] `json:"firingTriggers,required"`
}

func (r ZarazConfigUpdateParamsToolsZarazManagedComponentNeoEvent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigUpdateParamsToolsZarazCustomManagedComponent struct {
	// List of blocking trigger IDs
	BlockingTriggers param.Field[[]string] `json:"blockingTriggers,required"`
	// Tool's internal name
	Component param.Field[string] `json:"component,required"`
	// Default fields for tool's actions
	DefaultFields param.Field[map[string]ZarazConfigUpdateParamsToolsZarazCustomManagedComponentDefaultFields] `json:"defaultFields,required"`
	// Whether tool is enabled
	Enabled param.Field[bool] `json:"enabled,required"`
	// Tool's name defined by the user
	Name param.Field[string] `json:"name,required"`
	// List of permissions granted to the component
	Permissions param.Field[[]string] `json:"permissions,required"`
	// Tool's settings
	Settings param.Field[map[string]ZarazConfigUpdateParamsToolsZarazCustomManagedComponentSettings] `json:"settings,required"`
	Type     param.Field[ZarazConfigUpdateParamsToolsZarazCustomManagedComponentType]                `json:"type,required"`
	// Cloudflare worker that acts as a managed component
	Worker param.Field[ZarazConfigUpdateParamsToolsZarazCustomManagedComponentWorker] `json:"worker,required"`
	// Actions configured on a tool. Either this or neoEvents field is required.
	Actions param.Field[map[string]ZarazConfigUpdateParamsToolsZarazCustomManagedComponentActions] `json:"actions"`
	// Default consent purpose ID
	DefaultPurpose param.Field[string] `json:"defaultPurpose"`
	// DEPRECATED - List of actions configured on a tool. Either this or actions field
	// is required. If both are present, actions field will take precedence.
	NeoEvents param.Field[[]ZarazConfigUpdateParamsToolsZarazCustomManagedComponentNeoEvent] `json:"neoEvents"`
}

func (r ZarazConfigUpdateParamsToolsZarazCustomManagedComponent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZarazConfigUpdateParamsToolsZarazCustomManagedComponent) implementsZarazConfigUpdateParamsTools() {
}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type ZarazConfigUpdateParamsToolsZarazCustomManagedComponentDefaultFields interface {
	ImplementsZarazConfigUpdateParamsToolsZarazCustomManagedComponentDefaultFields()
}

// Satisfied by [shared.UnionString], [shared.UnionBool].
type ZarazConfigUpdateParamsToolsZarazCustomManagedComponentSettings interface {
	ImplementsZarazConfigUpdateParamsToolsZarazCustomManagedComponentSettings()
}

type ZarazConfigUpdateParamsToolsZarazCustomManagedComponentType string

const (
	ZarazConfigUpdateParamsToolsZarazCustomManagedComponentTypeCustomMc ZarazConfigUpdateParamsToolsZarazCustomManagedComponentType = "custom-mc"
)

// Cloudflare worker that acts as a managed component
type ZarazConfigUpdateParamsToolsZarazCustomManagedComponentWorker struct {
	EscapedWorkerName param.Field[string] `json:"escapedWorkerName,required"`
	WorkerTag         param.Field[string] `json:"workerTag,required"`
}

func (r ZarazConfigUpdateParamsToolsZarazCustomManagedComponentWorker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigUpdateParamsToolsZarazCustomManagedComponentActions struct {
	// Tool event type
	ActionType param.Field[string] `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers param.Field[[]string] `json:"blockingTriggers,required"`
	// Event payload
	Data param.Field[interface{}] `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers param.Field[[]string] `json:"firingTriggers,required"`
}

func (r ZarazConfigUpdateParamsToolsZarazCustomManagedComponentActions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigUpdateParamsToolsZarazCustomManagedComponentNeoEvent struct {
	// Tool event type
	ActionType param.Field[string] `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers param.Field[[]string] `json:"blockingTriggers,required"`
	// Event payload
	Data param.Field[interface{}] `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers param.Field[[]string] `json:"firingTriggers,required"`
}

func (r ZarazConfigUpdateParamsToolsZarazCustomManagedComponentNeoEvent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigUpdateParamsTriggers struct {
	// Rules defining when the trigger is not fired.
	ExcludeRules param.Field[[]ZarazConfigUpdateParamsTriggersExcludeRule] `json:"excludeRules,required"`
	// Rules defining when the trigger is fired.
	LoadRules param.Field[[]ZarazConfigUpdateParamsTriggersLoadRule] `json:"loadRules,required"`
	// Trigger name.
	Name param.Field[string] `json:"name,required"`
	// Trigger description.
	Description param.Field[string]                                `json:"description"`
	System      param.Field[ZarazConfigUpdateParamsTriggersSystem] `json:"system"`
}

func (r ZarazConfigUpdateParamsTriggers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [ZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRule],
// [ZarazConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRule],
// [ZarazConfigUpdateParamsTriggersExcludeRulesZarazTimerRule],
// [ZarazConfigUpdateParamsTriggersExcludeRulesZarazFormSubmissionRule],
// [ZarazConfigUpdateParamsTriggersExcludeRulesZarazVariableMatchRule],
// [ZarazConfigUpdateParamsTriggersExcludeRulesZarazScrollDepthRule],
// [ZarazConfigUpdateParamsTriggersExcludeRulesZarazElementVisibilityRule].
type ZarazConfigUpdateParamsTriggersExcludeRule interface {
	implementsZarazConfigUpdateParamsTriggersExcludeRule()
}

type ZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRule struct {
	ID    param.Field[string]                                                     `json:"id,required"`
	Match param.Field[string]                                                     `json:"match,required"`
	Op    param.Field[ZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOp] `json:"op,required"`
	Value param.Field[string]                                                     `json:"value,required"`
}

func (r ZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRule) implementsZarazConfigUpdateParamsTriggersExcludeRule() {
}

type ZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOp string

const (
	ZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpContains           ZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOp = "CONTAINS"
	ZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpEquals             ZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOp = "EQUALS"
	ZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpStartsWith         ZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOp = "STARTS_WITH"
	ZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpEndsWith           ZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOp = "ENDS_WITH"
	ZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpMatchRegex         ZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOp = "MATCH_REGEX"
	ZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpNotMatchRegex      ZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOp = "NOT_MATCH_REGEX"
	ZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpGreaterThan        ZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOp = "GREATER_THAN"
	ZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpGreaterThanOrEqual ZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOp = "GREATER_THAN_OR_EQUAL"
	ZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpLessThan           ZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOp = "LESS_THAN"
	ZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOpLessThanOrEqual    ZarazConfigUpdateParamsTriggersExcludeRulesZarazLoadRuleOp = "LESS_THAN_OR_EQUAL"
)

type ZarazConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRule struct {
	ID       param.Field[string]                                                                    `json:"id,required"`
	Action   param.Field[ZarazConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleAction]   `json:"action,required"`
	Settings param.Field[ZarazConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleSettings] `json:"settings,required"`
}

func (r ZarazConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZarazConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRule) implementsZarazConfigUpdateParamsTriggersExcludeRule() {
}

type ZarazConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleAction string

const (
	ZarazConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleActionClickListener ZarazConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleAction = "clickListener"
)

type ZarazConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleSettings struct {
	Selector    param.Field[string]                                                                        `json:"selector,required"`
	Type        param.Field[ZarazConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleSettingsType] `json:"type,required"`
	WaitForTags param.Field[int64]                                                                         `json:"waitForTags,required"`
}

func (r ZarazConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleSettingsType string

const (
	ZarazConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleSettingsTypeXpath ZarazConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleSettingsType = "xpath"
	ZarazConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleSettingsTypeCss   ZarazConfigUpdateParamsTriggersExcludeRulesZarazClickListenerRuleSettingsType = "css"
)

type ZarazConfigUpdateParamsTriggersExcludeRulesZarazTimerRule struct {
	ID       param.Field[string]                                                            `json:"id,required"`
	Action   param.Field[ZarazConfigUpdateParamsTriggersExcludeRulesZarazTimerRuleAction]   `json:"action,required"`
	Settings param.Field[ZarazConfigUpdateParamsTriggersExcludeRulesZarazTimerRuleSettings] `json:"settings,required"`
}

func (r ZarazConfigUpdateParamsTriggersExcludeRulesZarazTimerRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZarazConfigUpdateParamsTriggersExcludeRulesZarazTimerRule) implementsZarazConfigUpdateParamsTriggersExcludeRule() {
}

type ZarazConfigUpdateParamsTriggersExcludeRulesZarazTimerRuleAction string

const (
	ZarazConfigUpdateParamsTriggersExcludeRulesZarazTimerRuleActionTimer ZarazConfigUpdateParamsTriggersExcludeRulesZarazTimerRuleAction = "timer"
)

type ZarazConfigUpdateParamsTriggersExcludeRulesZarazTimerRuleSettings struct {
	Interval param.Field[int64] `json:"interval,required"`
	Limit    param.Field[int64] `json:"limit,required"`
}

func (r ZarazConfigUpdateParamsTriggersExcludeRulesZarazTimerRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigUpdateParamsTriggersExcludeRulesZarazFormSubmissionRule struct {
	ID       param.Field[string]                                                                     `json:"id,required"`
	Action   param.Field[ZarazConfigUpdateParamsTriggersExcludeRulesZarazFormSubmissionRuleAction]   `json:"action,required"`
	Settings param.Field[ZarazConfigUpdateParamsTriggersExcludeRulesZarazFormSubmissionRuleSettings] `json:"settings,required"`
}

func (r ZarazConfigUpdateParamsTriggersExcludeRulesZarazFormSubmissionRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZarazConfigUpdateParamsTriggersExcludeRulesZarazFormSubmissionRule) implementsZarazConfigUpdateParamsTriggersExcludeRule() {
}

type ZarazConfigUpdateParamsTriggersExcludeRulesZarazFormSubmissionRuleAction string

const (
	ZarazConfigUpdateParamsTriggersExcludeRulesZarazFormSubmissionRuleActionFormSubmission ZarazConfigUpdateParamsTriggersExcludeRulesZarazFormSubmissionRuleAction = "formSubmission"
)

type ZarazConfigUpdateParamsTriggersExcludeRulesZarazFormSubmissionRuleSettings struct {
	Selector param.Field[string] `json:"selector,required"`
	Validate param.Field[bool]   `json:"validate,required"`
}

func (r ZarazConfigUpdateParamsTriggersExcludeRulesZarazFormSubmissionRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigUpdateParamsTriggersExcludeRulesZarazVariableMatchRule struct {
	ID       param.Field[string]                                                                    `json:"id,required"`
	Action   param.Field[ZarazConfigUpdateParamsTriggersExcludeRulesZarazVariableMatchRuleAction]   `json:"action,required"`
	Settings param.Field[ZarazConfigUpdateParamsTriggersExcludeRulesZarazVariableMatchRuleSettings] `json:"settings,required"`
}

func (r ZarazConfigUpdateParamsTriggersExcludeRulesZarazVariableMatchRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZarazConfigUpdateParamsTriggersExcludeRulesZarazVariableMatchRule) implementsZarazConfigUpdateParamsTriggersExcludeRule() {
}

type ZarazConfigUpdateParamsTriggersExcludeRulesZarazVariableMatchRuleAction string

const (
	ZarazConfigUpdateParamsTriggersExcludeRulesZarazVariableMatchRuleActionVariableMatch ZarazConfigUpdateParamsTriggersExcludeRulesZarazVariableMatchRuleAction = "variableMatch"
)

type ZarazConfigUpdateParamsTriggersExcludeRulesZarazVariableMatchRuleSettings struct {
	Match    param.Field[string] `json:"match,required"`
	Variable param.Field[string] `json:"variable,required"`
}

func (r ZarazConfigUpdateParamsTriggersExcludeRulesZarazVariableMatchRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigUpdateParamsTriggersExcludeRulesZarazScrollDepthRule struct {
	ID       param.Field[string]                                                                  `json:"id,required"`
	Action   param.Field[ZarazConfigUpdateParamsTriggersExcludeRulesZarazScrollDepthRuleAction]   `json:"action,required"`
	Settings param.Field[ZarazConfigUpdateParamsTriggersExcludeRulesZarazScrollDepthRuleSettings] `json:"settings,required"`
}

func (r ZarazConfigUpdateParamsTriggersExcludeRulesZarazScrollDepthRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZarazConfigUpdateParamsTriggersExcludeRulesZarazScrollDepthRule) implementsZarazConfigUpdateParamsTriggersExcludeRule() {
}

type ZarazConfigUpdateParamsTriggersExcludeRulesZarazScrollDepthRuleAction string

const (
	ZarazConfigUpdateParamsTriggersExcludeRulesZarazScrollDepthRuleActionScrollDepth ZarazConfigUpdateParamsTriggersExcludeRulesZarazScrollDepthRuleAction = "scrollDepth"
)

type ZarazConfigUpdateParamsTriggersExcludeRulesZarazScrollDepthRuleSettings struct {
	Positions param.Field[string] `json:"positions,required"`
}

func (r ZarazConfigUpdateParamsTriggersExcludeRulesZarazScrollDepthRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigUpdateParamsTriggersExcludeRulesZarazElementVisibilityRule struct {
	ID       param.Field[string]                                                                        `json:"id,required"`
	Action   param.Field[ZarazConfigUpdateParamsTriggersExcludeRulesZarazElementVisibilityRuleAction]   `json:"action,required"`
	Settings param.Field[ZarazConfigUpdateParamsTriggersExcludeRulesZarazElementVisibilityRuleSettings] `json:"settings,required"`
}

func (r ZarazConfigUpdateParamsTriggersExcludeRulesZarazElementVisibilityRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZarazConfigUpdateParamsTriggersExcludeRulesZarazElementVisibilityRule) implementsZarazConfigUpdateParamsTriggersExcludeRule() {
}

type ZarazConfigUpdateParamsTriggersExcludeRulesZarazElementVisibilityRuleAction string

const (
	ZarazConfigUpdateParamsTriggersExcludeRulesZarazElementVisibilityRuleActionElementVisibility ZarazConfigUpdateParamsTriggersExcludeRulesZarazElementVisibilityRuleAction = "elementVisibility"
)

type ZarazConfigUpdateParamsTriggersExcludeRulesZarazElementVisibilityRuleSettings struct {
	Selector param.Field[string] `json:"selector,required"`
}

func (r ZarazConfigUpdateParamsTriggersExcludeRulesZarazElementVisibilityRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [ZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRule],
// [ZarazConfigUpdateParamsTriggersLoadRulesZarazClickListenerRule],
// [ZarazConfigUpdateParamsTriggersLoadRulesZarazTimerRule],
// [ZarazConfigUpdateParamsTriggersLoadRulesZarazFormSubmissionRule],
// [ZarazConfigUpdateParamsTriggersLoadRulesZarazVariableMatchRule],
// [ZarazConfigUpdateParamsTriggersLoadRulesZarazScrollDepthRule],
// [ZarazConfigUpdateParamsTriggersLoadRulesZarazElementVisibilityRule].
type ZarazConfigUpdateParamsTriggersLoadRule interface {
	implementsZarazConfigUpdateParamsTriggersLoadRule()
}

type ZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRule struct {
	ID    param.Field[string]                                                  `json:"id,required"`
	Match param.Field[string]                                                  `json:"match,required"`
	Op    param.Field[ZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOp] `json:"op,required"`
	Value param.Field[string]                                                  `json:"value,required"`
}

func (r ZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRule) implementsZarazConfigUpdateParamsTriggersLoadRule() {
}

type ZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOp string

const (
	ZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpContains           ZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOp = "CONTAINS"
	ZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpEquals             ZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOp = "EQUALS"
	ZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpStartsWith         ZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOp = "STARTS_WITH"
	ZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpEndsWith           ZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOp = "ENDS_WITH"
	ZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpMatchRegex         ZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOp = "MATCH_REGEX"
	ZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpNotMatchRegex      ZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOp = "NOT_MATCH_REGEX"
	ZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpGreaterThan        ZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOp = "GREATER_THAN"
	ZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpGreaterThanOrEqual ZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOp = "GREATER_THAN_OR_EQUAL"
	ZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpLessThan           ZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOp = "LESS_THAN"
	ZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOpLessThanOrEqual    ZarazConfigUpdateParamsTriggersLoadRulesZarazLoadRuleOp = "LESS_THAN_OR_EQUAL"
)

type ZarazConfigUpdateParamsTriggersLoadRulesZarazClickListenerRule struct {
	ID       param.Field[string]                                                                 `json:"id,required"`
	Action   param.Field[ZarazConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleAction]   `json:"action,required"`
	Settings param.Field[ZarazConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleSettings] `json:"settings,required"`
}

func (r ZarazConfigUpdateParamsTriggersLoadRulesZarazClickListenerRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZarazConfigUpdateParamsTriggersLoadRulesZarazClickListenerRule) implementsZarazConfigUpdateParamsTriggersLoadRule() {
}

type ZarazConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleAction string

const (
	ZarazConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleActionClickListener ZarazConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleAction = "clickListener"
)

type ZarazConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleSettings struct {
	Selector    param.Field[string]                                                                     `json:"selector,required"`
	Type        param.Field[ZarazConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleSettingsType] `json:"type,required"`
	WaitForTags param.Field[int64]                                                                      `json:"waitForTags,required"`
}

func (r ZarazConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleSettingsType string

const (
	ZarazConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleSettingsTypeXpath ZarazConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleSettingsType = "xpath"
	ZarazConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleSettingsTypeCss   ZarazConfigUpdateParamsTriggersLoadRulesZarazClickListenerRuleSettingsType = "css"
)

type ZarazConfigUpdateParamsTriggersLoadRulesZarazTimerRule struct {
	ID       param.Field[string]                                                         `json:"id,required"`
	Action   param.Field[ZarazConfigUpdateParamsTriggersLoadRulesZarazTimerRuleAction]   `json:"action,required"`
	Settings param.Field[ZarazConfigUpdateParamsTriggersLoadRulesZarazTimerRuleSettings] `json:"settings,required"`
}

func (r ZarazConfigUpdateParamsTriggersLoadRulesZarazTimerRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZarazConfigUpdateParamsTriggersLoadRulesZarazTimerRule) implementsZarazConfigUpdateParamsTriggersLoadRule() {
}

type ZarazConfigUpdateParamsTriggersLoadRulesZarazTimerRuleAction string

const (
	ZarazConfigUpdateParamsTriggersLoadRulesZarazTimerRuleActionTimer ZarazConfigUpdateParamsTriggersLoadRulesZarazTimerRuleAction = "timer"
)

type ZarazConfigUpdateParamsTriggersLoadRulesZarazTimerRuleSettings struct {
	Interval param.Field[int64] `json:"interval,required"`
	Limit    param.Field[int64] `json:"limit,required"`
}

func (r ZarazConfigUpdateParamsTriggersLoadRulesZarazTimerRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigUpdateParamsTriggersLoadRulesZarazFormSubmissionRule struct {
	ID       param.Field[string]                                                                  `json:"id,required"`
	Action   param.Field[ZarazConfigUpdateParamsTriggersLoadRulesZarazFormSubmissionRuleAction]   `json:"action,required"`
	Settings param.Field[ZarazConfigUpdateParamsTriggersLoadRulesZarazFormSubmissionRuleSettings] `json:"settings,required"`
}

func (r ZarazConfigUpdateParamsTriggersLoadRulesZarazFormSubmissionRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZarazConfigUpdateParamsTriggersLoadRulesZarazFormSubmissionRule) implementsZarazConfigUpdateParamsTriggersLoadRule() {
}

type ZarazConfigUpdateParamsTriggersLoadRulesZarazFormSubmissionRuleAction string

const (
	ZarazConfigUpdateParamsTriggersLoadRulesZarazFormSubmissionRuleActionFormSubmission ZarazConfigUpdateParamsTriggersLoadRulesZarazFormSubmissionRuleAction = "formSubmission"
)

type ZarazConfigUpdateParamsTriggersLoadRulesZarazFormSubmissionRuleSettings struct {
	Selector param.Field[string] `json:"selector,required"`
	Validate param.Field[bool]   `json:"validate,required"`
}

func (r ZarazConfigUpdateParamsTriggersLoadRulesZarazFormSubmissionRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigUpdateParamsTriggersLoadRulesZarazVariableMatchRule struct {
	ID       param.Field[string]                                                                 `json:"id,required"`
	Action   param.Field[ZarazConfigUpdateParamsTriggersLoadRulesZarazVariableMatchRuleAction]   `json:"action,required"`
	Settings param.Field[ZarazConfigUpdateParamsTriggersLoadRulesZarazVariableMatchRuleSettings] `json:"settings,required"`
}

func (r ZarazConfigUpdateParamsTriggersLoadRulesZarazVariableMatchRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZarazConfigUpdateParamsTriggersLoadRulesZarazVariableMatchRule) implementsZarazConfigUpdateParamsTriggersLoadRule() {
}

type ZarazConfigUpdateParamsTriggersLoadRulesZarazVariableMatchRuleAction string

const (
	ZarazConfigUpdateParamsTriggersLoadRulesZarazVariableMatchRuleActionVariableMatch ZarazConfigUpdateParamsTriggersLoadRulesZarazVariableMatchRuleAction = "variableMatch"
)

type ZarazConfigUpdateParamsTriggersLoadRulesZarazVariableMatchRuleSettings struct {
	Match    param.Field[string] `json:"match,required"`
	Variable param.Field[string] `json:"variable,required"`
}

func (r ZarazConfigUpdateParamsTriggersLoadRulesZarazVariableMatchRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigUpdateParamsTriggersLoadRulesZarazScrollDepthRule struct {
	ID       param.Field[string]                                                               `json:"id,required"`
	Action   param.Field[ZarazConfigUpdateParamsTriggersLoadRulesZarazScrollDepthRuleAction]   `json:"action,required"`
	Settings param.Field[ZarazConfigUpdateParamsTriggersLoadRulesZarazScrollDepthRuleSettings] `json:"settings,required"`
}

func (r ZarazConfigUpdateParamsTriggersLoadRulesZarazScrollDepthRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZarazConfigUpdateParamsTriggersLoadRulesZarazScrollDepthRule) implementsZarazConfigUpdateParamsTriggersLoadRule() {
}

type ZarazConfigUpdateParamsTriggersLoadRulesZarazScrollDepthRuleAction string

const (
	ZarazConfigUpdateParamsTriggersLoadRulesZarazScrollDepthRuleActionScrollDepth ZarazConfigUpdateParamsTriggersLoadRulesZarazScrollDepthRuleAction = "scrollDepth"
)

type ZarazConfigUpdateParamsTriggersLoadRulesZarazScrollDepthRuleSettings struct {
	Positions param.Field[string] `json:"positions,required"`
}

func (r ZarazConfigUpdateParamsTriggersLoadRulesZarazScrollDepthRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigUpdateParamsTriggersLoadRulesZarazElementVisibilityRule struct {
	ID       param.Field[string]                                                                     `json:"id,required"`
	Action   param.Field[ZarazConfigUpdateParamsTriggersLoadRulesZarazElementVisibilityRuleAction]   `json:"action,required"`
	Settings param.Field[ZarazConfigUpdateParamsTriggersLoadRulesZarazElementVisibilityRuleSettings] `json:"settings,required"`
}

func (r ZarazConfigUpdateParamsTriggersLoadRulesZarazElementVisibilityRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZarazConfigUpdateParamsTriggersLoadRulesZarazElementVisibilityRule) implementsZarazConfigUpdateParamsTriggersLoadRule() {
}

type ZarazConfigUpdateParamsTriggersLoadRulesZarazElementVisibilityRuleAction string

const (
	ZarazConfigUpdateParamsTriggersLoadRulesZarazElementVisibilityRuleActionElementVisibility ZarazConfigUpdateParamsTriggersLoadRulesZarazElementVisibilityRuleAction = "elementVisibility"
)

type ZarazConfigUpdateParamsTriggersLoadRulesZarazElementVisibilityRuleSettings struct {
	Selector param.Field[string] `json:"selector,required"`
}

func (r ZarazConfigUpdateParamsTriggersLoadRulesZarazElementVisibilityRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigUpdateParamsTriggersSystem string

const (
	ZarazConfigUpdateParamsTriggersSystemPageload ZarazConfigUpdateParamsTriggersSystem = "pageload"
)

// Satisfied by [ZarazConfigUpdateParamsVariablesObject],
// [ZarazConfigUpdateParamsVariablesObject].
type ZarazConfigUpdateParamsVariables interface {
	implementsZarazConfigUpdateParamsVariables()
}

type ZarazConfigUpdateParamsVariablesObject struct {
	Name  param.Field[string]                                     `json:"name,required"`
	Type  param.Field[ZarazConfigUpdateParamsVariablesObjectType] `json:"type,required"`
	Value param.Field[string]                                     `json:"value,required"`
}

func (r ZarazConfigUpdateParamsVariablesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZarazConfigUpdateParamsVariablesObject) implementsZarazConfigUpdateParamsVariables() {}

type ZarazConfigUpdateParamsVariablesObjectType string

const (
	ZarazConfigUpdateParamsVariablesObjectTypeString ZarazConfigUpdateParamsVariablesObjectType = "string"
	ZarazConfigUpdateParamsVariablesObjectTypeSecret ZarazConfigUpdateParamsVariablesObjectType = "secret"
)

// Consent management configuration.
type ZarazConfigUpdateParamsConsent struct {
	Enabled                param.Field[bool]                                                 `json:"enabled,required"`
	ButtonTextTranslations param.Field[ZarazConfigUpdateParamsConsentButtonTextTranslations] `json:"buttonTextTranslations"`
	CompanyEmail           param.Field[string]                                               `json:"companyEmail"`
	CompanyName            param.Field[string]                                               `json:"companyName"`
	CompanyStreetAddress   param.Field[string]                                               `json:"companyStreetAddress"`
	ConsentModalIntroHTML  param.Field[string]                                               `json:"consentModalIntroHTML"`
	// Object where keys are language codes
	ConsentModalIntroHTMLWithTranslations param.Field[map[string]string] `json:"consentModalIntroHTMLWithTranslations"`
	CookieName                            param.Field[string]            `json:"cookieName"`
	CustomCss                             param.Field[string]            `json:"customCSS"`
	CustomIntroDisclaimerDismissed        param.Field[bool]              `json:"customIntroDisclaimerDismissed"`
	DefaultLanguage                       param.Field[string]            `json:"defaultLanguage"`
	HideModal                             param.Field[bool]              `json:"hideModal"`
	// Object where keys are purpose alpha-numeric IDs
	Purposes param.Field[map[string]ZarazConfigUpdateParamsConsentPurposes] `json:"purposes"`
	// Object where keys are purpose alpha-numeric IDs
	PurposesWithTranslations param.Field[map[string]ZarazConfigUpdateParamsConsentPurposesWithTranslations] `json:"purposesWithTranslations"`
}

func (r ZarazConfigUpdateParamsConsent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigUpdateParamsConsentButtonTextTranslations struct {
	// Object where keys are language codes
	AcceptAll param.Field[map[string]string] `json:"accept_all,required"`
	// Object where keys are language codes
	ConfirmMyChoices param.Field[map[string]string] `json:"confirm_my_choices,required"`
	// Object where keys are language codes
	RejectAll param.Field[map[string]string] `json:"reject_all,required"`
}

func (r ZarazConfigUpdateParamsConsentButtonTextTranslations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigUpdateParamsConsentPurposes struct {
	Description param.Field[string] `json:"description,required"`
	Name        param.Field[string] `json:"name,required"`
}

func (r ZarazConfigUpdateParamsConsentPurposes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigUpdateParamsConsentPurposesWithTranslations struct {
	// Object where keys are language codes
	Description param.Field[map[string]string] `json:"description,required"`
	// Object where keys are language codes
	Name  param.Field[map[string]string] `json:"name,required"`
	Order param.Field[int64]             `json:"order,required"`
}

func (r ZarazConfigUpdateParamsConsentPurposesWithTranslations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigUpdateResponseEnvelope struct {
	Errors   []ZarazConfigUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZarazConfigUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Zaraz configuration
	Result ZarazConfigUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success bool                                  `json:"success,required"`
	JSON    zarazConfigUpdateResponseEnvelopeJSON `json:"-"`
}

// zarazConfigUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZarazConfigUpdateResponseEnvelope]
type zarazConfigUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigUpdateResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    zarazConfigUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// zarazConfigUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZarazConfigUpdateResponseEnvelopeErrors]
type zarazConfigUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigUpdateResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zarazConfigUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// zarazConfigUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ZarazConfigUpdateResponseEnvelopeMessages]
type zarazConfigUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
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
