// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// ZarazDefaultService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZarazDefaultService] method
// instead.
type ZarazDefaultService struct {
	Options []option.RequestOption
}

// NewZarazDefaultService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZarazDefaultService(opts ...option.RequestOption) (r *ZarazDefaultService) {
	r = &ZarazDefaultService{}
	r.Options = opts
	return
}

// Gets default Zaraz configuration for a zone.
func (r *ZarazDefaultService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZarazDefaultGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZarazDefaultGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/zaraz/default", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Zaraz configuration
type ZarazDefaultGetResponse struct {
	// Data layer compatibility mode enabled.
	DataLayer bool `json:"dataLayer,required"`
	// The key for Zaraz debug mode.
	DebugKey string `json:"debugKey,required"`
	// General Zaraz settings.
	Settings ZarazDefaultGetResponseSettings `json:"settings,required"`
	// Tools set up under Zaraz configuration, where key is the alpha-numeric tool ID
	// and value is the tool configuration object.
	Tools map[string]ZarazDefaultGetResponseTool `json:"tools,required"`
	// Triggers set up under Zaraz configuration, where key is the trigger
	// alpha-numeric ID and value is the trigger configuration.
	Triggers map[string]ZarazDefaultGetResponseTrigger `json:"triggers,required"`
	// Variables set up under Zaraz configuration, where key is the variable
	// alpha-numeric ID and value is the variable configuration. Values of variables of
	// type secret are not included.
	Variables map[string]ZarazDefaultGetResponseVariable `json:"variables,required"`
	// Zaraz internal version of the config.
	ZarazVersion int64 `json:"zarazVersion,required"`
	// Consent management configuration.
	Consent ZarazDefaultGetResponseConsent `json:"consent"`
	// Single Page Application support enabled.
	HistoryChange bool                        `json:"historyChange"`
	JSON          zarazDefaultGetResponseJSON `json:"-"`
}

// zarazDefaultGetResponseJSON contains the JSON metadata for the struct
// [ZarazDefaultGetResponse]
type zarazDefaultGetResponseJSON struct {
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

func (r *ZarazDefaultGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// General Zaraz settings.
type ZarazDefaultGetResponseSettings struct {
	// Automatic injection of Zaraz scripts enabled.
	AutoInjectScript bool `json:"autoInjectScript,required"`
	// Details of the worker that receives and edits Zaraz Context object.
	ContextEnricher ZarazDefaultGetResponseSettingsContextEnricher `json:"contextEnricher"`
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
	TrackPath string                              `json:"trackPath"`
	JSON      zarazDefaultGetResponseSettingsJSON `json:"-"`
}

// zarazDefaultGetResponseSettingsJSON contains the JSON metadata for the struct
// [ZarazDefaultGetResponseSettings]
type zarazDefaultGetResponseSettingsJSON struct {
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

func (r *ZarazDefaultGetResponseSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details of the worker that receives and edits Zaraz Context object.
type ZarazDefaultGetResponseSettingsContextEnricher struct {
	EscapedWorkerName string                                             `json:"escapedWorkerName,required"`
	WorkerTag         string                                             `json:"workerTag,required"`
	JSON              zarazDefaultGetResponseSettingsContextEnricherJSON `json:"-"`
}

// zarazDefaultGetResponseSettingsContextEnricherJSON contains the JSON metadata
// for the struct [ZarazDefaultGetResponseSettingsContextEnricher]
type zarazDefaultGetResponseSettingsContextEnricherJSON struct {
	EscapedWorkerName apijson.Field
	WorkerTag         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseSettingsContextEnricher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [ZarazDefaultGetResponseToolsZarazManagedComponent] or
// [ZarazDefaultGetResponseToolsZarazCustomManagedComponent].
type ZarazDefaultGetResponseTool interface {
	implementsZarazDefaultGetResponseTool()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZarazDefaultGetResponseTool)(nil)).Elem(), "")
}

type ZarazDefaultGetResponseToolsZarazManagedComponent struct {
	// List of blocking trigger IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Tool's internal name
	Component string `json:"component,required"`
	// Default fields for tool's actions
	DefaultFields map[string]ZarazDefaultGetResponseToolsZarazManagedComponentDefaultField `json:"defaultFields,required"`
	// Whether tool is enabled
	Enabled bool `json:"enabled,required"`
	// Tool's name defined by the user
	Name string `json:"name,required"`
	// List of permissions granted to the component
	Permissions []string `json:"permissions,required"`
	// Tool's settings
	Settings map[string]ZarazDefaultGetResponseToolsZarazManagedComponentSetting `json:"settings,required"`
	Type     ZarazDefaultGetResponseToolsZarazManagedComponentType               `json:"type,required"`
	// Actions configured on a tool. Either this or neoEvents field is required.
	Actions map[string]ZarazDefaultGetResponseToolsZarazManagedComponentAction `json:"actions"`
	// Default consent purpose ID
	DefaultPurpose string `json:"defaultPurpose"`
	// DEPRECATED - List of actions configured on a tool. Either this or actions field
	// is required. If both are present, actions field will take precedence.
	NeoEvents []ZarazDefaultGetResponseToolsZarazManagedComponentNeoEvent `json:"neoEvents"`
	JSON      zarazDefaultGetResponseToolsZarazManagedComponentJSON       `json:"-"`
}

// zarazDefaultGetResponseToolsZarazManagedComponentJSON contains the JSON metadata
// for the struct [ZarazDefaultGetResponseToolsZarazManagedComponent]
type zarazDefaultGetResponseToolsZarazManagedComponentJSON struct {
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

func (r *ZarazDefaultGetResponseToolsZarazManagedComponent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazDefaultGetResponseToolsZarazManagedComponent) implementsZarazDefaultGetResponseTool() {}

// Union satisfied by [shared.UnionString] or [shared.UnionBool].
type ZarazDefaultGetResponseToolsZarazManagedComponentDefaultField interface {
	ImplementsZarazDefaultGetResponseToolsZarazManagedComponentDefaultField()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazDefaultGetResponseToolsZarazManagedComponentDefaultField)(nil)).Elem(),
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
type ZarazDefaultGetResponseToolsZarazManagedComponentSetting interface {
	ImplementsZarazDefaultGetResponseToolsZarazManagedComponentSetting()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazDefaultGetResponseToolsZarazManagedComponentSetting)(nil)).Elem(),
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

type ZarazDefaultGetResponseToolsZarazManagedComponentType string

const (
	ZarazDefaultGetResponseToolsZarazManagedComponentTypeComponent ZarazDefaultGetResponseToolsZarazManagedComponentType = "component"
)

type ZarazDefaultGetResponseToolsZarazManagedComponentAction struct {
	// Tool event type
	ActionType string `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Event payload
	Data interface{} `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers []string                                                    `json:"firingTriggers,required"`
	JSON           zarazDefaultGetResponseToolsZarazManagedComponentActionJSON `json:"-"`
}

// zarazDefaultGetResponseToolsZarazManagedComponentActionJSON contains the JSON
// metadata for the struct
// [ZarazDefaultGetResponseToolsZarazManagedComponentAction]
type zarazDefaultGetResponseToolsZarazManagedComponentActionJSON struct {
	ActionType       apijson.Field
	BlockingTriggers apijson.Field
	Data             apijson.Field
	FiringTriggers   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseToolsZarazManagedComponentAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazDefaultGetResponseToolsZarazManagedComponentNeoEvent struct {
	// Tool event type
	ActionType string `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Event payload
	Data interface{} `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers []string                                                      `json:"firingTriggers,required"`
	JSON           zarazDefaultGetResponseToolsZarazManagedComponentNeoEventJSON `json:"-"`
}

// zarazDefaultGetResponseToolsZarazManagedComponentNeoEventJSON contains the JSON
// metadata for the struct
// [ZarazDefaultGetResponseToolsZarazManagedComponentNeoEvent]
type zarazDefaultGetResponseToolsZarazManagedComponentNeoEventJSON struct {
	ActionType       apijson.Field
	BlockingTriggers apijson.Field
	Data             apijson.Field
	FiringTriggers   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseToolsZarazManagedComponentNeoEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazDefaultGetResponseToolsZarazCustomManagedComponent struct {
	// List of blocking trigger IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Tool's internal name
	Component string `json:"component,required"`
	// Default fields for tool's actions
	DefaultFields map[string]ZarazDefaultGetResponseToolsZarazCustomManagedComponentDefaultField `json:"defaultFields,required"`
	// Whether tool is enabled
	Enabled bool `json:"enabled,required"`
	// Tool's name defined by the user
	Name string `json:"name,required"`
	// List of permissions granted to the component
	Permissions []string `json:"permissions,required"`
	// Tool's settings
	Settings map[string]ZarazDefaultGetResponseToolsZarazCustomManagedComponentSetting `json:"settings,required"`
	Type     ZarazDefaultGetResponseToolsZarazCustomManagedComponentType               `json:"type,required"`
	// Cloudflare worker that acts as a managed component
	Worker ZarazDefaultGetResponseToolsZarazCustomManagedComponentWorker `json:"worker,required"`
	// Actions configured on a tool. Either this or neoEvents field is required.
	Actions map[string]ZarazDefaultGetResponseToolsZarazCustomManagedComponentAction `json:"actions"`
	// Default consent purpose ID
	DefaultPurpose string `json:"defaultPurpose"`
	// DEPRECATED - List of actions configured on a tool. Either this or actions field
	// is required. If both are present, actions field will take precedence.
	NeoEvents []ZarazDefaultGetResponseToolsZarazCustomManagedComponentNeoEvent `json:"neoEvents"`
	JSON      zarazDefaultGetResponseToolsZarazCustomManagedComponentJSON       `json:"-"`
}

// zarazDefaultGetResponseToolsZarazCustomManagedComponentJSON contains the JSON
// metadata for the struct
// [ZarazDefaultGetResponseToolsZarazCustomManagedComponent]
type zarazDefaultGetResponseToolsZarazCustomManagedComponentJSON struct {
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

func (r *ZarazDefaultGetResponseToolsZarazCustomManagedComponent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazDefaultGetResponseToolsZarazCustomManagedComponent) implementsZarazDefaultGetResponseTool() {
}

// Union satisfied by [shared.UnionString] or [shared.UnionBool].
type ZarazDefaultGetResponseToolsZarazCustomManagedComponentDefaultField interface {
	ImplementsZarazDefaultGetResponseToolsZarazCustomManagedComponentDefaultField()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazDefaultGetResponseToolsZarazCustomManagedComponentDefaultField)(nil)).Elem(),
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
type ZarazDefaultGetResponseToolsZarazCustomManagedComponentSetting interface {
	ImplementsZarazDefaultGetResponseToolsZarazCustomManagedComponentSetting()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazDefaultGetResponseToolsZarazCustomManagedComponentSetting)(nil)).Elem(),
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

type ZarazDefaultGetResponseToolsZarazCustomManagedComponentType string

const (
	ZarazDefaultGetResponseToolsZarazCustomManagedComponentTypeCustomMc ZarazDefaultGetResponseToolsZarazCustomManagedComponentType = "custom-mc"
)

// Cloudflare worker that acts as a managed component
type ZarazDefaultGetResponseToolsZarazCustomManagedComponentWorker struct {
	EscapedWorkerName string                                                            `json:"escapedWorkerName,required"`
	WorkerTag         string                                                            `json:"workerTag,required"`
	JSON              zarazDefaultGetResponseToolsZarazCustomManagedComponentWorkerJSON `json:"-"`
}

// zarazDefaultGetResponseToolsZarazCustomManagedComponentWorkerJSON contains the
// JSON metadata for the struct
// [ZarazDefaultGetResponseToolsZarazCustomManagedComponentWorker]
type zarazDefaultGetResponseToolsZarazCustomManagedComponentWorkerJSON struct {
	EscapedWorkerName apijson.Field
	WorkerTag         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseToolsZarazCustomManagedComponentWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazDefaultGetResponseToolsZarazCustomManagedComponentAction struct {
	// Tool event type
	ActionType string `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Event payload
	Data interface{} `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers []string                                                          `json:"firingTriggers,required"`
	JSON           zarazDefaultGetResponseToolsZarazCustomManagedComponentActionJSON `json:"-"`
}

// zarazDefaultGetResponseToolsZarazCustomManagedComponentActionJSON contains the
// JSON metadata for the struct
// [ZarazDefaultGetResponseToolsZarazCustomManagedComponentAction]
type zarazDefaultGetResponseToolsZarazCustomManagedComponentActionJSON struct {
	ActionType       apijson.Field
	BlockingTriggers apijson.Field
	Data             apijson.Field
	FiringTriggers   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseToolsZarazCustomManagedComponentAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazDefaultGetResponseToolsZarazCustomManagedComponentNeoEvent struct {
	// Tool event type
	ActionType string `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Event payload
	Data interface{} `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers []string                                                            `json:"firingTriggers,required"`
	JSON           zarazDefaultGetResponseToolsZarazCustomManagedComponentNeoEventJSON `json:"-"`
}

// zarazDefaultGetResponseToolsZarazCustomManagedComponentNeoEventJSON contains the
// JSON metadata for the struct
// [ZarazDefaultGetResponseToolsZarazCustomManagedComponentNeoEvent]
type zarazDefaultGetResponseToolsZarazCustomManagedComponentNeoEventJSON struct {
	ActionType       apijson.Field
	BlockingTriggers apijson.Field
	Data             apijson.Field
	FiringTriggers   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseToolsZarazCustomManagedComponentNeoEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazDefaultGetResponseTrigger struct {
	// Rules defining when the trigger is not fired.
	ExcludeRules []ZarazDefaultGetResponseTriggersExcludeRule `json:"excludeRules,required"`
	// Rules defining when the trigger is fired.
	LoadRules []ZarazDefaultGetResponseTriggersLoadRule `json:"loadRules,required"`
	// Trigger name.
	Name string `json:"name,required"`
	// Trigger description.
	Description string                                `json:"description"`
	System      ZarazDefaultGetResponseTriggersSystem `json:"system"`
	JSON        zarazDefaultGetResponseTriggerJSON    `json:"-"`
}

// zarazDefaultGetResponseTriggerJSON contains the JSON metadata for the struct
// [ZarazDefaultGetResponseTrigger]
type zarazDefaultGetResponseTriggerJSON struct {
	ExcludeRules apijson.Field
	LoadRules    apijson.Field
	Name         apijson.Field
	Description  apijson.Field
	System       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [ZarazDefaultGetResponseTriggersExcludeRulesZarazLoadRule],
// [ZarazDefaultGetResponseTriggersExcludeRulesZarazClickListenerRule],
// [ZarazDefaultGetResponseTriggersExcludeRulesZarazTimerRule],
// [ZarazDefaultGetResponseTriggersExcludeRulesZarazFormSubmissionRule],
// [ZarazDefaultGetResponseTriggersExcludeRulesZarazVariableMatchRule],
// [ZarazDefaultGetResponseTriggersExcludeRulesZarazScrollDepthRule] or
// [ZarazDefaultGetResponseTriggersExcludeRulesZarazElementVisibilityRule].
type ZarazDefaultGetResponseTriggersExcludeRule interface {
	implementsZarazDefaultGetResponseTriggersExcludeRule()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZarazDefaultGetResponseTriggersExcludeRule)(nil)).Elem(), "")
}

type ZarazDefaultGetResponseTriggersExcludeRulesZarazLoadRule struct {
	ID    string                                                       `json:"id,required"`
	Match string                                                       `json:"match,required"`
	Op    ZarazDefaultGetResponseTriggersExcludeRulesZarazLoadRuleOp   `json:"op,required"`
	Value string                                                       `json:"value,required"`
	JSON  zarazDefaultGetResponseTriggersExcludeRulesZarazLoadRuleJSON `json:"-"`
}

// zarazDefaultGetResponseTriggersExcludeRulesZarazLoadRuleJSON contains the JSON
// metadata for the struct
// [ZarazDefaultGetResponseTriggersExcludeRulesZarazLoadRule]
type zarazDefaultGetResponseTriggersExcludeRulesZarazLoadRuleJSON struct {
	ID          apijson.Field
	Match       apijson.Field
	Op          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseTriggersExcludeRulesZarazLoadRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazDefaultGetResponseTriggersExcludeRulesZarazLoadRule) implementsZarazDefaultGetResponseTriggersExcludeRule() {
}

type ZarazDefaultGetResponseTriggersExcludeRulesZarazLoadRuleOp string

const (
	ZarazDefaultGetResponseTriggersExcludeRulesZarazLoadRuleOpContains           ZarazDefaultGetResponseTriggersExcludeRulesZarazLoadRuleOp = "CONTAINS"
	ZarazDefaultGetResponseTriggersExcludeRulesZarazLoadRuleOpEquals             ZarazDefaultGetResponseTriggersExcludeRulesZarazLoadRuleOp = "EQUALS"
	ZarazDefaultGetResponseTriggersExcludeRulesZarazLoadRuleOpStartsWith         ZarazDefaultGetResponseTriggersExcludeRulesZarazLoadRuleOp = "STARTS_WITH"
	ZarazDefaultGetResponseTriggersExcludeRulesZarazLoadRuleOpEndsWith           ZarazDefaultGetResponseTriggersExcludeRulesZarazLoadRuleOp = "ENDS_WITH"
	ZarazDefaultGetResponseTriggersExcludeRulesZarazLoadRuleOpMatchRegex         ZarazDefaultGetResponseTriggersExcludeRulesZarazLoadRuleOp = "MATCH_REGEX"
	ZarazDefaultGetResponseTriggersExcludeRulesZarazLoadRuleOpNotMatchRegex      ZarazDefaultGetResponseTriggersExcludeRulesZarazLoadRuleOp = "NOT_MATCH_REGEX"
	ZarazDefaultGetResponseTriggersExcludeRulesZarazLoadRuleOpGreaterThan        ZarazDefaultGetResponseTriggersExcludeRulesZarazLoadRuleOp = "GREATER_THAN"
	ZarazDefaultGetResponseTriggersExcludeRulesZarazLoadRuleOpGreaterThanOrEqual ZarazDefaultGetResponseTriggersExcludeRulesZarazLoadRuleOp = "GREATER_THAN_OR_EQUAL"
	ZarazDefaultGetResponseTriggersExcludeRulesZarazLoadRuleOpLessThan           ZarazDefaultGetResponseTriggersExcludeRulesZarazLoadRuleOp = "LESS_THAN"
	ZarazDefaultGetResponseTriggersExcludeRulesZarazLoadRuleOpLessThanOrEqual    ZarazDefaultGetResponseTriggersExcludeRulesZarazLoadRuleOp = "LESS_THAN_OR_EQUAL"
)

type ZarazDefaultGetResponseTriggersExcludeRulesZarazClickListenerRule struct {
	ID       string                                                                    `json:"id,required"`
	Action   ZarazDefaultGetResponseTriggersExcludeRulesZarazClickListenerRuleAction   `json:"action,required"`
	Settings ZarazDefaultGetResponseTriggersExcludeRulesZarazClickListenerRuleSettings `json:"settings,required"`
	JSON     zarazDefaultGetResponseTriggersExcludeRulesZarazClickListenerRuleJSON     `json:"-"`
}

// zarazDefaultGetResponseTriggersExcludeRulesZarazClickListenerRuleJSON contains
// the JSON metadata for the struct
// [ZarazDefaultGetResponseTriggersExcludeRulesZarazClickListenerRule]
type zarazDefaultGetResponseTriggersExcludeRulesZarazClickListenerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseTriggersExcludeRulesZarazClickListenerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazDefaultGetResponseTriggersExcludeRulesZarazClickListenerRule) implementsZarazDefaultGetResponseTriggersExcludeRule() {
}

type ZarazDefaultGetResponseTriggersExcludeRulesZarazClickListenerRuleAction string

const (
	ZarazDefaultGetResponseTriggersExcludeRulesZarazClickListenerRuleActionClickListener ZarazDefaultGetResponseTriggersExcludeRulesZarazClickListenerRuleAction = "clickListener"
)

type ZarazDefaultGetResponseTriggersExcludeRulesZarazClickListenerRuleSettings struct {
	Selector    string                                                                        `json:"selector,required"`
	Type        ZarazDefaultGetResponseTriggersExcludeRulesZarazClickListenerRuleSettingsType `json:"type,required"`
	WaitForTags int64                                                                         `json:"waitForTags,required"`
	JSON        zarazDefaultGetResponseTriggersExcludeRulesZarazClickListenerRuleSettingsJSON `json:"-"`
}

// zarazDefaultGetResponseTriggersExcludeRulesZarazClickListenerRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazDefaultGetResponseTriggersExcludeRulesZarazClickListenerRuleSettings]
type zarazDefaultGetResponseTriggersExcludeRulesZarazClickListenerRuleSettingsJSON struct {
	Selector    apijson.Field
	Type        apijson.Field
	WaitForTags apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseTriggersExcludeRulesZarazClickListenerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazDefaultGetResponseTriggersExcludeRulesZarazClickListenerRuleSettingsType string

const (
	ZarazDefaultGetResponseTriggersExcludeRulesZarazClickListenerRuleSettingsTypeXpath ZarazDefaultGetResponseTriggersExcludeRulesZarazClickListenerRuleSettingsType = "xpath"
	ZarazDefaultGetResponseTriggersExcludeRulesZarazClickListenerRuleSettingsTypeCss   ZarazDefaultGetResponseTriggersExcludeRulesZarazClickListenerRuleSettingsType = "css"
)

type ZarazDefaultGetResponseTriggersExcludeRulesZarazTimerRule struct {
	ID       string                                                            `json:"id,required"`
	Action   ZarazDefaultGetResponseTriggersExcludeRulesZarazTimerRuleAction   `json:"action,required"`
	Settings ZarazDefaultGetResponseTriggersExcludeRulesZarazTimerRuleSettings `json:"settings,required"`
	JSON     zarazDefaultGetResponseTriggersExcludeRulesZarazTimerRuleJSON     `json:"-"`
}

// zarazDefaultGetResponseTriggersExcludeRulesZarazTimerRuleJSON contains the JSON
// metadata for the struct
// [ZarazDefaultGetResponseTriggersExcludeRulesZarazTimerRule]
type zarazDefaultGetResponseTriggersExcludeRulesZarazTimerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseTriggersExcludeRulesZarazTimerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazDefaultGetResponseTriggersExcludeRulesZarazTimerRule) implementsZarazDefaultGetResponseTriggersExcludeRule() {
}

type ZarazDefaultGetResponseTriggersExcludeRulesZarazTimerRuleAction string

const (
	ZarazDefaultGetResponseTriggersExcludeRulesZarazTimerRuleActionTimer ZarazDefaultGetResponseTriggersExcludeRulesZarazTimerRuleAction = "timer"
)

type ZarazDefaultGetResponseTriggersExcludeRulesZarazTimerRuleSettings struct {
	Interval int64                                                                 `json:"interval,required"`
	Limit    int64                                                                 `json:"limit,required"`
	JSON     zarazDefaultGetResponseTriggersExcludeRulesZarazTimerRuleSettingsJSON `json:"-"`
}

// zarazDefaultGetResponseTriggersExcludeRulesZarazTimerRuleSettingsJSON contains
// the JSON metadata for the struct
// [ZarazDefaultGetResponseTriggersExcludeRulesZarazTimerRuleSettings]
type zarazDefaultGetResponseTriggersExcludeRulesZarazTimerRuleSettingsJSON struct {
	Interval    apijson.Field
	Limit       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseTriggersExcludeRulesZarazTimerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazDefaultGetResponseTriggersExcludeRulesZarazFormSubmissionRule struct {
	ID       string                                                                     `json:"id,required"`
	Action   ZarazDefaultGetResponseTriggersExcludeRulesZarazFormSubmissionRuleAction   `json:"action,required"`
	Settings ZarazDefaultGetResponseTriggersExcludeRulesZarazFormSubmissionRuleSettings `json:"settings,required"`
	JSON     zarazDefaultGetResponseTriggersExcludeRulesZarazFormSubmissionRuleJSON     `json:"-"`
}

// zarazDefaultGetResponseTriggersExcludeRulesZarazFormSubmissionRuleJSON contains
// the JSON metadata for the struct
// [ZarazDefaultGetResponseTriggersExcludeRulesZarazFormSubmissionRule]
type zarazDefaultGetResponseTriggersExcludeRulesZarazFormSubmissionRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseTriggersExcludeRulesZarazFormSubmissionRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazDefaultGetResponseTriggersExcludeRulesZarazFormSubmissionRule) implementsZarazDefaultGetResponseTriggersExcludeRule() {
}

type ZarazDefaultGetResponseTriggersExcludeRulesZarazFormSubmissionRuleAction string

const (
	ZarazDefaultGetResponseTriggersExcludeRulesZarazFormSubmissionRuleActionFormSubmission ZarazDefaultGetResponseTriggersExcludeRulesZarazFormSubmissionRuleAction = "formSubmission"
)

type ZarazDefaultGetResponseTriggersExcludeRulesZarazFormSubmissionRuleSettings struct {
	Selector string                                                                         `json:"selector,required"`
	Validate bool                                                                           `json:"validate,required"`
	JSON     zarazDefaultGetResponseTriggersExcludeRulesZarazFormSubmissionRuleSettingsJSON `json:"-"`
}

// zarazDefaultGetResponseTriggersExcludeRulesZarazFormSubmissionRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazDefaultGetResponseTriggersExcludeRulesZarazFormSubmissionRuleSettings]
type zarazDefaultGetResponseTriggersExcludeRulesZarazFormSubmissionRuleSettingsJSON struct {
	Selector    apijson.Field
	Validate    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseTriggersExcludeRulesZarazFormSubmissionRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazDefaultGetResponseTriggersExcludeRulesZarazVariableMatchRule struct {
	ID       string                                                                    `json:"id,required"`
	Action   ZarazDefaultGetResponseTriggersExcludeRulesZarazVariableMatchRuleAction   `json:"action,required"`
	Settings ZarazDefaultGetResponseTriggersExcludeRulesZarazVariableMatchRuleSettings `json:"settings,required"`
	JSON     zarazDefaultGetResponseTriggersExcludeRulesZarazVariableMatchRuleJSON     `json:"-"`
}

// zarazDefaultGetResponseTriggersExcludeRulesZarazVariableMatchRuleJSON contains
// the JSON metadata for the struct
// [ZarazDefaultGetResponseTriggersExcludeRulesZarazVariableMatchRule]
type zarazDefaultGetResponseTriggersExcludeRulesZarazVariableMatchRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseTriggersExcludeRulesZarazVariableMatchRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazDefaultGetResponseTriggersExcludeRulesZarazVariableMatchRule) implementsZarazDefaultGetResponseTriggersExcludeRule() {
}

type ZarazDefaultGetResponseTriggersExcludeRulesZarazVariableMatchRuleAction string

const (
	ZarazDefaultGetResponseTriggersExcludeRulesZarazVariableMatchRuleActionVariableMatch ZarazDefaultGetResponseTriggersExcludeRulesZarazVariableMatchRuleAction = "variableMatch"
)

type ZarazDefaultGetResponseTriggersExcludeRulesZarazVariableMatchRuleSettings struct {
	Match    string                                                                        `json:"match,required"`
	Variable string                                                                        `json:"variable,required"`
	JSON     zarazDefaultGetResponseTriggersExcludeRulesZarazVariableMatchRuleSettingsJSON `json:"-"`
}

// zarazDefaultGetResponseTriggersExcludeRulesZarazVariableMatchRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazDefaultGetResponseTriggersExcludeRulesZarazVariableMatchRuleSettings]
type zarazDefaultGetResponseTriggersExcludeRulesZarazVariableMatchRuleSettingsJSON struct {
	Match       apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseTriggersExcludeRulesZarazVariableMatchRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazDefaultGetResponseTriggersExcludeRulesZarazScrollDepthRule struct {
	ID       string                                                                  `json:"id,required"`
	Action   ZarazDefaultGetResponseTriggersExcludeRulesZarazScrollDepthRuleAction   `json:"action,required"`
	Settings ZarazDefaultGetResponseTriggersExcludeRulesZarazScrollDepthRuleSettings `json:"settings,required"`
	JSON     zarazDefaultGetResponseTriggersExcludeRulesZarazScrollDepthRuleJSON     `json:"-"`
}

// zarazDefaultGetResponseTriggersExcludeRulesZarazScrollDepthRuleJSON contains the
// JSON metadata for the struct
// [ZarazDefaultGetResponseTriggersExcludeRulesZarazScrollDepthRule]
type zarazDefaultGetResponseTriggersExcludeRulesZarazScrollDepthRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseTriggersExcludeRulesZarazScrollDepthRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazDefaultGetResponseTriggersExcludeRulesZarazScrollDepthRule) implementsZarazDefaultGetResponseTriggersExcludeRule() {
}

type ZarazDefaultGetResponseTriggersExcludeRulesZarazScrollDepthRuleAction string

const (
	ZarazDefaultGetResponseTriggersExcludeRulesZarazScrollDepthRuleActionScrollDepth ZarazDefaultGetResponseTriggersExcludeRulesZarazScrollDepthRuleAction = "scrollDepth"
)

type ZarazDefaultGetResponseTriggersExcludeRulesZarazScrollDepthRuleSettings struct {
	Positions string                                                                      `json:"positions,required"`
	JSON      zarazDefaultGetResponseTriggersExcludeRulesZarazScrollDepthRuleSettingsJSON `json:"-"`
}

// zarazDefaultGetResponseTriggersExcludeRulesZarazScrollDepthRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazDefaultGetResponseTriggersExcludeRulesZarazScrollDepthRuleSettings]
type zarazDefaultGetResponseTriggersExcludeRulesZarazScrollDepthRuleSettingsJSON struct {
	Positions   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseTriggersExcludeRulesZarazScrollDepthRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazDefaultGetResponseTriggersExcludeRulesZarazElementVisibilityRule struct {
	ID       string                                                                        `json:"id,required"`
	Action   ZarazDefaultGetResponseTriggersExcludeRulesZarazElementVisibilityRuleAction   `json:"action,required"`
	Settings ZarazDefaultGetResponseTriggersExcludeRulesZarazElementVisibilityRuleSettings `json:"settings,required"`
	JSON     zarazDefaultGetResponseTriggersExcludeRulesZarazElementVisibilityRuleJSON     `json:"-"`
}

// zarazDefaultGetResponseTriggersExcludeRulesZarazElementVisibilityRuleJSON
// contains the JSON metadata for the struct
// [ZarazDefaultGetResponseTriggersExcludeRulesZarazElementVisibilityRule]
type zarazDefaultGetResponseTriggersExcludeRulesZarazElementVisibilityRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseTriggersExcludeRulesZarazElementVisibilityRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazDefaultGetResponseTriggersExcludeRulesZarazElementVisibilityRule) implementsZarazDefaultGetResponseTriggersExcludeRule() {
}

type ZarazDefaultGetResponseTriggersExcludeRulesZarazElementVisibilityRuleAction string

const (
	ZarazDefaultGetResponseTriggersExcludeRulesZarazElementVisibilityRuleActionElementVisibility ZarazDefaultGetResponseTriggersExcludeRulesZarazElementVisibilityRuleAction = "elementVisibility"
)

type ZarazDefaultGetResponseTriggersExcludeRulesZarazElementVisibilityRuleSettings struct {
	Selector string                                                                            `json:"selector,required"`
	JSON     zarazDefaultGetResponseTriggersExcludeRulesZarazElementVisibilityRuleSettingsJSON `json:"-"`
}

// zarazDefaultGetResponseTriggersExcludeRulesZarazElementVisibilityRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazDefaultGetResponseTriggersExcludeRulesZarazElementVisibilityRuleSettings]
type zarazDefaultGetResponseTriggersExcludeRulesZarazElementVisibilityRuleSettingsJSON struct {
	Selector    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseTriggersExcludeRulesZarazElementVisibilityRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [ZarazDefaultGetResponseTriggersLoadRulesZarazLoadRule],
// [ZarazDefaultGetResponseTriggersLoadRulesZarazClickListenerRule],
// [ZarazDefaultGetResponseTriggersLoadRulesZarazTimerRule],
// [ZarazDefaultGetResponseTriggersLoadRulesZarazFormSubmissionRule],
// [ZarazDefaultGetResponseTriggersLoadRulesZarazVariableMatchRule],
// [ZarazDefaultGetResponseTriggersLoadRulesZarazScrollDepthRule] or
// [ZarazDefaultGetResponseTriggersLoadRulesZarazElementVisibilityRule].
type ZarazDefaultGetResponseTriggersLoadRule interface {
	implementsZarazDefaultGetResponseTriggersLoadRule()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZarazDefaultGetResponseTriggersLoadRule)(nil)).Elem(), "")
}

type ZarazDefaultGetResponseTriggersLoadRulesZarazLoadRule struct {
	ID    string                                                    `json:"id,required"`
	Match string                                                    `json:"match,required"`
	Op    ZarazDefaultGetResponseTriggersLoadRulesZarazLoadRuleOp   `json:"op,required"`
	Value string                                                    `json:"value,required"`
	JSON  zarazDefaultGetResponseTriggersLoadRulesZarazLoadRuleJSON `json:"-"`
}

// zarazDefaultGetResponseTriggersLoadRulesZarazLoadRuleJSON contains the JSON
// metadata for the struct [ZarazDefaultGetResponseTriggersLoadRulesZarazLoadRule]
type zarazDefaultGetResponseTriggersLoadRulesZarazLoadRuleJSON struct {
	ID          apijson.Field
	Match       apijson.Field
	Op          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseTriggersLoadRulesZarazLoadRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazDefaultGetResponseTriggersLoadRulesZarazLoadRule) implementsZarazDefaultGetResponseTriggersLoadRule() {
}

type ZarazDefaultGetResponseTriggersLoadRulesZarazLoadRuleOp string

const (
	ZarazDefaultGetResponseTriggersLoadRulesZarazLoadRuleOpContains           ZarazDefaultGetResponseTriggersLoadRulesZarazLoadRuleOp = "CONTAINS"
	ZarazDefaultGetResponseTriggersLoadRulesZarazLoadRuleOpEquals             ZarazDefaultGetResponseTriggersLoadRulesZarazLoadRuleOp = "EQUALS"
	ZarazDefaultGetResponseTriggersLoadRulesZarazLoadRuleOpStartsWith         ZarazDefaultGetResponseTriggersLoadRulesZarazLoadRuleOp = "STARTS_WITH"
	ZarazDefaultGetResponseTriggersLoadRulesZarazLoadRuleOpEndsWith           ZarazDefaultGetResponseTriggersLoadRulesZarazLoadRuleOp = "ENDS_WITH"
	ZarazDefaultGetResponseTriggersLoadRulesZarazLoadRuleOpMatchRegex         ZarazDefaultGetResponseTriggersLoadRulesZarazLoadRuleOp = "MATCH_REGEX"
	ZarazDefaultGetResponseTriggersLoadRulesZarazLoadRuleOpNotMatchRegex      ZarazDefaultGetResponseTriggersLoadRulesZarazLoadRuleOp = "NOT_MATCH_REGEX"
	ZarazDefaultGetResponseTriggersLoadRulesZarazLoadRuleOpGreaterThan        ZarazDefaultGetResponseTriggersLoadRulesZarazLoadRuleOp = "GREATER_THAN"
	ZarazDefaultGetResponseTriggersLoadRulesZarazLoadRuleOpGreaterThanOrEqual ZarazDefaultGetResponseTriggersLoadRulesZarazLoadRuleOp = "GREATER_THAN_OR_EQUAL"
	ZarazDefaultGetResponseTriggersLoadRulesZarazLoadRuleOpLessThan           ZarazDefaultGetResponseTriggersLoadRulesZarazLoadRuleOp = "LESS_THAN"
	ZarazDefaultGetResponseTriggersLoadRulesZarazLoadRuleOpLessThanOrEqual    ZarazDefaultGetResponseTriggersLoadRulesZarazLoadRuleOp = "LESS_THAN_OR_EQUAL"
)

type ZarazDefaultGetResponseTriggersLoadRulesZarazClickListenerRule struct {
	ID       string                                                                 `json:"id,required"`
	Action   ZarazDefaultGetResponseTriggersLoadRulesZarazClickListenerRuleAction   `json:"action,required"`
	Settings ZarazDefaultGetResponseTriggersLoadRulesZarazClickListenerRuleSettings `json:"settings,required"`
	JSON     zarazDefaultGetResponseTriggersLoadRulesZarazClickListenerRuleJSON     `json:"-"`
}

// zarazDefaultGetResponseTriggersLoadRulesZarazClickListenerRuleJSON contains the
// JSON metadata for the struct
// [ZarazDefaultGetResponseTriggersLoadRulesZarazClickListenerRule]
type zarazDefaultGetResponseTriggersLoadRulesZarazClickListenerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseTriggersLoadRulesZarazClickListenerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazDefaultGetResponseTriggersLoadRulesZarazClickListenerRule) implementsZarazDefaultGetResponseTriggersLoadRule() {
}

type ZarazDefaultGetResponseTriggersLoadRulesZarazClickListenerRuleAction string

const (
	ZarazDefaultGetResponseTriggersLoadRulesZarazClickListenerRuleActionClickListener ZarazDefaultGetResponseTriggersLoadRulesZarazClickListenerRuleAction = "clickListener"
)

type ZarazDefaultGetResponseTriggersLoadRulesZarazClickListenerRuleSettings struct {
	Selector    string                                                                     `json:"selector,required"`
	Type        ZarazDefaultGetResponseTriggersLoadRulesZarazClickListenerRuleSettingsType `json:"type,required"`
	WaitForTags int64                                                                      `json:"waitForTags,required"`
	JSON        zarazDefaultGetResponseTriggersLoadRulesZarazClickListenerRuleSettingsJSON `json:"-"`
}

// zarazDefaultGetResponseTriggersLoadRulesZarazClickListenerRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazDefaultGetResponseTriggersLoadRulesZarazClickListenerRuleSettings]
type zarazDefaultGetResponseTriggersLoadRulesZarazClickListenerRuleSettingsJSON struct {
	Selector    apijson.Field
	Type        apijson.Field
	WaitForTags apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseTriggersLoadRulesZarazClickListenerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazDefaultGetResponseTriggersLoadRulesZarazClickListenerRuleSettingsType string

const (
	ZarazDefaultGetResponseTriggersLoadRulesZarazClickListenerRuleSettingsTypeXpath ZarazDefaultGetResponseTriggersLoadRulesZarazClickListenerRuleSettingsType = "xpath"
	ZarazDefaultGetResponseTriggersLoadRulesZarazClickListenerRuleSettingsTypeCss   ZarazDefaultGetResponseTriggersLoadRulesZarazClickListenerRuleSettingsType = "css"
)

type ZarazDefaultGetResponseTriggersLoadRulesZarazTimerRule struct {
	ID       string                                                         `json:"id,required"`
	Action   ZarazDefaultGetResponseTriggersLoadRulesZarazTimerRuleAction   `json:"action,required"`
	Settings ZarazDefaultGetResponseTriggersLoadRulesZarazTimerRuleSettings `json:"settings,required"`
	JSON     zarazDefaultGetResponseTriggersLoadRulesZarazTimerRuleJSON     `json:"-"`
}

// zarazDefaultGetResponseTriggersLoadRulesZarazTimerRuleJSON contains the JSON
// metadata for the struct [ZarazDefaultGetResponseTriggersLoadRulesZarazTimerRule]
type zarazDefaultGetResponseTriggersLoadRulesZarazTimerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseTriggersLoadRulesZarazTimerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazDefaultGetResponseTriggersLoadRulesZarazTimerRule) implementsZarazDefaultGetResponseTriggersLoadRule() {
}

type ZarazDefaultGetResponseTriggersLoadRulesZarazTimerRuleAction string

const (
	ZarazDefaultGetResponseTriggersLoadRulesZarazTimerRuleActionTimer ZarazDefaultGetResponseTriggersLoadRulesZarazTimerRuleAction = "timer"
)

type ZarazDefaultGetResponseTriggersLoadRulesZarazTimerRuleSettings struct {
	Interval int64                                                              `json:"interval,required"`
	Limit    int64                                                              `json:"limit,required"`
	JSON     zarazDefaultGetResponseTriggersLoadRulesZarazTimerRuleSettingsJSON `json:"-"`
}

// zarazDefaultGetResponseTriggersLoadRulesZarazTimerRuleSettingsJSON contains the
// JSON metadata for the struct
// [ZarazDefaultGetResponseTriggersLoadRulesZarazTimerRuleSettings]
type zarazDefaultGetResponseTriggersLoadRulesZarazTimerRuleSettingsJSON struct {
	Interval    apijson.Field
	Limit       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseTriggersLoadRulesZarazTimerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazDefaultGetResponseTriggersLoadRulesZarazFormSubmissionRule struct {
	ID       string                                                                  `json:"id,required"`
	Action   ZarazDefaultGetResponseTriggersLoadRulesZarazFormSubmissionRuleAction   `json:"action,required"`
	Settings ZarazDefaultGetResponseTriggersLoadRulesZarazFormSubmissionRuleSettings `json:"settings,required"`
	JSON     zarazDefaultGetResponseTriggersLoadRulesZarazFormSubmissionRuleJSON     `json:"-"`
}

// zarazDefaultGetResponseTriggersLoadRulesZarazFormSubmissionRuleJSON contains the
// JSON metadata for the struct
// [ZarazDefaultGetResponseTriggersLoadRulesZarazFormSubmissionRule]
type zarazDefaultGetResponseTriggersLoadRulesZarazFormSubmissionRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseTriggersLoadRulesZarazFormSubmissionRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazDefaultGetResponseTriggersLoadRulesZarazFormSubmissionRule) implementsZarazDefaultGetResponseTriggersLoadRule() {
}

type ZarazDefaultGetResponseTriggersLoadRulesZarazFormSubmissionRuleAction string

const (
	ZarazDefaultGetResponseTriggersLoadRulesZarazFormSubmissionRuleActionFormSubmission ZarazDefaultGetResponseTriggersLoadRulesZarazFormSubmissionRuleAction = "formSubmission"
)

type ZarazDefaultGetResponseTriggersLoadRulesZarazFormSubmissionRuleSettings struct {
	Selector string                                                                      `json:"selector,required"`
	Validate bool                                                                        `json:"validate,required"`
	JSON     zarazDefaultGetResponseTriggersLoadRulesZarazFormSubmissionRuleSettingsJSON `json:"-"`
}

// zarazDefaultGetResponseTriggersLoadRulesZarazFormSubmissionRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazDefaultGetResponseTriggersLoadRulesZarazFormSubmissionRuleSettings]
type zarazDefaultGetResponseTriggersLoadRulesZarazFormSubmissionRuleSettingsJSON struct {
	Selector    apijson.Field
	Validate    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseTriggersLoadRulesZarazFormSubmissionRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazDefaultGetResponseTriggersLoadRulesZarazVariableMatchRule struct {
	ID       string                                                                 `json:"id,required"`
	Action   ZarazDefaultGetResponseTriggersLoadRulesZarazVariableMatchRuleAction   `json:"action,required"`
	Settings ZarazDefaultGetResponseTriggersLoadRulesZarazVariableMatchRuleSettings `json:"settings,required"`
	JSON     zarazDefaultGetResponseTriggersLoadRulesZarazVariableMatchRuleJSON     `json:"-"`
}

// zarazDefaultGetResponseTriggersLoadRulesZarazVariableMatchRuleJSON contains the
// JSON metadata for the struct
// [ZarazDefaultGetResponseTriggersLoadRulesZarazVariableMatchRule]
type zarazDefaultGetResponseTriggersLoadRulesZarazVariableMatchRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseTriggersLoadRulesZarazVariableMatchRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazDefaultGetResponseTriggersLoadRulesZarazVariableMatchRule) implementsZarazDefaultGetResponseTriggersLoadRule() {
}

type ZarazDefaultGetResponseTriggersLoadRulesZarazVariableMatchRuleAction string

const (
	ZarazDefaultGetResponseTriggersLoadRulesZarazVariableMatchRuleActionVariableMatch ZarazDefaultGetResponseTriggersLoadRulesZarazVariableMatchRuleAction = "variableMatch"
)

type ZarazDefaultGetResponseTriggersLoadRulesZarazVariableMatchRuleSettings struct {
	Match    string                                                                     `json:"match,required"`
	Variable string                                                                     `json:"variable,required"`
	JSON     zarazDefaultGetResponseTriggersLoadRulesZarazVariableMatchRuleSettingsJSON `json:"-"`
}

// zarazDefaultGetResponseTriggersLoadRulesZarazVariableMatchRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazDefaultGetResponseTriggersLoadRulesZarazVariableMatchRuleSettings]
type zarazDefaultGetResponseTriggersLoadRulesZarazVariableMatchRuleSettingsJSON struct {
	Match       apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseTriggersLoadRulesZarazVariableMatchRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazDefaultGetResponseTriggersLoadRulesZarazScrollDepthRule struct {
	ID       string                                                               `json:"id,required"`
	Action   ZarazDefaultGetResponseTriggersLoadRulesZarazScrollDepthRuleAction   `json:"action,required"`
	Settings ZarazDefaultGetResponseTriggersLoadRulesZarazScrollDepthRuleSettings `json:"settings,required"`
	JSON     zarazDefaultGetResponseTriggersLoadRulesZarazScrollDepthRuleJSON     `json:"-"`
}

// zarazDefaultGetResponseTriggersLoadRulesZarazScrollDepthRuleJSON contains the
// JSON metadata for the struct
// [ZarazDefaultGetResponseTriggersLoadRulesZarazScrollDepthRule]
type zarazDefaultGetResponseTriggersLoadRulesZarazScrollDepthRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseTriggersLoadRulesZarazScrollDepthRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazDefaultGetResponseTriggersLoadRulesZarazScrollDepthRule) implementsZarazDefaultGetResponseTriggersLoadRule() {
}

type ZarazDefaultGetResponseTriggersLoadRulesZarazScrollDepthRuleAction string

const (
	ZarazDefaultGetResponseTriggersLoadRulesZarazScrollDepthRuleActionScrollDepth ZarazDefaultGetResponseTriggersLoadRulesZarazScrollDepthRuleAction = "scrollDepth"
)

type ZarazDefaultGetResponseTriggersLoadRulesZarazScrollDepthRuleSettings struct {
	Positions string                                                                   `json:"positions,required"`
	JSON      zarazDefaultGetResponseTriggersLoadRulesZarazScrollDepthRuleSettingsJSON `json:"-"`
}

// zarazDefaultGetResponseTriggersLoadRulesZarazScrollDepthRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazDefaultGetResponseTriggersLoadRulesZarazScrollDepthRuleSettings]
type zarazDefaultGetResponseTriggersLoadRulesZarazScrollDepthRuleSettingsJSON struct {
	Positions   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseTriggersLoadRulesZarazScrollDepthRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazDefaultGetResponseTriggersLoadRulesZarazElementVisibilityRule struct {
	ID       string                                                                     `json:"id,required"`
	Action   ZarazDefaultGetResponseTriggersLoadRulesZarazElementVisibilityRuleAction   `json:"action,required"`
	Settings ZarazDefaultGetResponseTriggersLoadRulesZarazElementVisibilityRuleSettings `json:"settings,required"`
	JSON     zarazDefaultGetResponseTriggersLoadRulesZarazElementVisibilityRuleJSON     `json:"-"`
}

// zarazDefaultGetResponseTriggersLoadRulesZarazElementVisibilityRuleJSON contains
// the JSON metadata for the struct
// [ZarazDefaultGetResponseTriggersLoadRulesZarazElementVisibilityRule]
type zarazDefaultGetResponseTriggersLoadRulesZarazElementVisibilityRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseTriggersLoadRulesZarazElementVisibilityRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazDefaultGetResponseTriggersLoadRulesZarazElementVisibilityRule) implementsZarazDefaultGetResponseTriggersLoadRule() {
}

type ZarazDefaultGetResponseTriggersLoadRulesZarazElementVisibilityRuleAction string

const (
	ZarazDefaultGetResponseTriggersLoadRulesZarazElementVisibilityRuleActionElementVisibility ZarazDefaultGetResponseTriggersLoadRulesZarazElementVisibilityRuleAction = "elementVisibility"
)

type ZarazDefaultGetResponseTriggersLoadRulesZarazElementVisibilityRuleSettings struct {
	Selector string                                                                         `json:"selector,required"`
	JSON     zarazDefaultGetResponseTriggersLoadRulesZarazElementVisibilityRuleSettingsJSON `json:"-"`
}

// zarazDefaultGetResponseTriggersLoadRulesZarazElementVisibilityRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazDefaultGetResponseTriggersLoadRulesZarazElementVisibilityRuleSettings]
type zarazDefaultGetResponseTriggersLoadRulesZarazElementVisibilityRuleSettingsJSON struct {
	Selector    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseTriggersLoadRulesZarazElementVisibilityRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazDefaultGetResponseTriggersSystem string

const (
	ZarazDefaultGetResponseTriggersSystemPageload ZarazDefaultGetResponseTriggersSystem = "pageload"
)

// Union satisfied by [ZarazDefaultGetResponseVariablesObject] or
// [ZarazDefaultGetResponseVariablesObject].
type ZarazDefaultGetResponseVariable interface {
	implementsZarazDefaultGetResponseVariable()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZarazDefaultGetResponseVariable)(nil)).Elem(), "")
}

type ZarazDefaultGetResponseVariablesObject struct {
	Name  string                                     `json:"name,required"`
	Type  ZarazDefaultGetResponseVariablesObjectType `json:"type,required"`
	Value string                                     `json:"value,required"`
	JSON  zarazDefaultGetResponseVariablesObjectJSON `json:"-"`
}

// zarazDefaultGetResponseVariablesObjectJSON contains the JSON metadata for the
// struct [ZarazDefaultGetResponseVariablesObject]
type zarazDefaultGetResponseVariablesObjectJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseVariablesObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazDefaultGetResponseVariablesObject) implementsZarazDefaultGetResponseVariable() {}

type ZarazDefaultGetResponseVariablesObjectType string

const (
	ZarazDefaultGetResponseVariablesObjectTypeString ZarazDefaultGetResponseVariablesObjectType = "string"
	ZarazDefaultGetResponseVariablesObjectTypeSecret ZarazDefaultGetResponseVariablesObjectType = "secret"
)

// Consent management configuration.
type ZarazDefaultGetResponseConsent struct {
	Enabled                bool                                                 `json:"enabled,required"`
	ButtonTextTranslations ZarazDefaultGetResponseConsentButtonTextTranslations `json:"buttonTextTranslations"`
	CompanyEmail           string                                               `json:"companyEmail"`
	CompanyName            string                                               `json:"companyName"`
	CompanyStreetAddress   string                                               `json:"companyStreetAddress"`
	ConsentModalIntroHTML  string                                               `json:"consentModalIntroHTML"`
	// Object where keys are language codes
	ConsentModalIntroHTMLWithTranslations map[string]string `json:"consentModalIntroHTMLWithTranslations"`
	CookieName                            string            `json:"cookieName"`
	CustomCss                             string            `json:"customCSS"`
	CustomIntroDisclaimerDismissed        bool              `json:"customIntroDisclaimerDismissed"`
	DefaultLanguage                       string            `json:"defaultLanguage"`
	HideModal                             bool              `json:"hideModal"`
	// Object where keys are purpose alpha-numeric IDs
	Purposes map[string]ZarazDefaultGetResponseConsentPurpose `json:"purposes"`
	// Object where keys are purpose alpha-numeric IDs
	PurposesWithTranslations map[string]ZarazDefaultGetResponseConsentPurposesWithTranslation `json:"purposesWithTranslations"`
	JSON                     zarazDefaultGetResponseConsentJSON                               `json:"-"`
}

// zarazDefaultGetResponseConsentJSON contains the JSON metadata for the struct
// [ZarazDefaultGetResponseConsent]
type zarazDefaultGetResponseConsentJSON struct {
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

func (r *ZarazDefaultGetResponseConsent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazDefaultGetResponseConsentButtonTextTranslations struct {
	// Object where keys are language codes
	AcceptAll map[string]string `json:"accept_all,required"`
	// Object where keys are language codes
	ConfirmMyChoices map[string]string `json:"confirm_my_choices,required"`
	// Object where keys are language codes
	RejectAll map[string]string                                        `json:"reject_all,required"`
	JSON      zarazDefaultGetResponseConsentButtonTextTranslationsJSON `json:"-"`
}

// zarazDefaultGetResponseConsentButtonTextTranslationsJSON contains the JSON
// metadata for the struct [ZarazDefaultGetResponseConsentButtonTextTranslations]
type zarazDefaultGetResponseConsentButtonTextTranslationsJSON struct {
	AcceptAll        apijson.Field
	ConfirmMyChoices apijson.Field
	RejectAll        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseConsentButtonTextTranslations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazDefaultGetResponseConsentPurpose struct {
	Description string                                    `json:"description,required"`
	Name        string                                    `json:"name,required"`
	JSON        zarazDefaultGetResponseConsentPurposeJSON `json:"-"`
}

// zarazDefaultGetResponseConsentPurposeJSON contains the JSON metadata for the
// struct [ZarazDefaultGetResponseConsentPurpose]
type zarazDefaultGetResponseConsentPurposeJSON struct {
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseConsentPurpose) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazDefaultGetResponseConsentPurposesWithTranslation struct {
	// Object where keys are language codes
	Description map[string]string `json:"description,required"`
	// Object where keys are language codes
	Name  map[string]string                                         `json:"name,required"`
	Order int64                                                     `json:"order,required"`
	JSON  zarazDefaultGetResponseConsentPurposesWithTranslationJSON `json:"-"`
}

// zarazDefaultGetResponseConsentPurposesWithTranslationJSON contains the JSON
// metadata for the struct [ZarazDefaultGetResponseConsentPurposesWithTranslation]
type zarazDefaultGetResponseConsentPurposesWithTranslationJSON struct {
	Description apijson.Field
	Name        apijson.Field
	Order       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseConsentPurposesWithTranslation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazDefaultGetResponseEnvelope struct {
	Errors   []ZarazDefaultGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZarazDefaultGetResponseEnvelopeMessages `json:"messages,required"`
	// Zaraz configuration
	Result ZarazDefaultGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success bool                                `json:"success,required"`
	JSON    zarazDefaultGetResponseEnvelopeJSON `json:"-"`
}

// zarazDefaultGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZarazDefaultGetResponseEnvelope]
type zarazDefaultGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazDefaultGetResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    zarazDefaultGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zarazDefaultGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZarazDefaultGetResponseEnvelopeErrors]
type zarazDefaultGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazDefaultGetResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    zarazDefaultGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zarazDefaultGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ZarazDefaultGetResponseEnvelopeMessages]
type zarazDefaultGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazDefaultGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
