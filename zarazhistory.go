// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// ZarazHistoryService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZarazHistoryService] method
// instead.
type ZarazHistoryService struct {
	Options []option.RequestOption
	Configs *ZarazHistoryConfigService
}

// NewZarazHistoryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZarazHistoryService(opts ...option.RequestOption) (r *ZarazHistoryService) {
	r = &ZarazHistoryService{}
	r.Options = opts
	r.Configs = NewZarazHistoryConfigService(opts...)
	return
}

// Restores a historical published Zaraz configuration by ID for a zone.
func (r *ZarazHistoryService) Update(ctx context.Context, zoneID string, body ZarazHistoryUpdateParams, opts ...option.RequestOption) (res *ZarazHistoryUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZarazHistoryUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/zaraz/history", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists a history of published Zaraz configuration records for a zone.
func (r *ZarazHistoryService) List(ctx context.Context, zoneID string, query ZarazHistoryListParams, opts ...option.RequestOption) (res *[]ZarazHistoryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZarazHistoryListResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/zaraz/history", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Zaraz configuration
type ZarazHistoryUpdateResponse struct {
	// Data layer compatibility mode enabled.
	DataLayer bool `json:"dataLayer,required"`
	// The key for Zaraz debug mode.
	DebugKey string `json:"debugKey,required"`
	// General Zaraz settings.
	Settings ZarazHistoryUpdateResponseSettings `json:"settings,required"`
	// Tools set up under Zaraz configuration, where key is the alpha-numeric tool ID
	// and value is the tool configuration object.
	Tools map[string]ZarazHistoryUpdateResponseTool `json:"tools,required"`
	// Triggers set up under Zaraz configuration, where key is the trigger
	// alpha-numeric ID and value is the trigger configuration.
	Triggers map[string]ZarazHistoryUpdateResponseTrigger `json:"triggers,required"`
	// Variables set up under Zaraz configuration, where key is the variable
	// alpha-numeric ID and value is the variable configuration. Values of variables of
	// type secret are not included.
	Variables map[string]ZarazHistoryUpdateResponseVariable `json:"variables,required"`
	// Zaraz internal version of the config.
	ZarazVersion int64 `json:"zarazVersion,required"`
	// Consent management configuration.
	Consent ZarazHistoryUpdateResponseConsent `json:"consent"`
	// Single Page Application support enabled.
	HistoryChange bool                           `json:"historyChange"`
	JSON          zarazHistoryUpdateResponseJSON `json:"-"`
}

// zarazHistoryUpdateResponseJSON contains the JSON metadata for the struct
// [ZarazHistoryUpdateResponse]
type zarazHistoryUpdateResponseJSON struct {
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

func (r *ZarazHistoryUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// General Zaraz settings.
type ZarazHistoryUpdateResponseSettings struct {
	// Automatic injection of Zaraz scripts enabled.
	AutoInjectScript bool `json:"autoInjectScript,required"`
	// Details of the worker that receives and edits Zaraz Context object.
	ContextEnricher ZarazHistoryUpdateResponseSettingsContextEnricher `json:"contextEnricher"`
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
	JSON      zarazHistoryUpdateResponseSettingsJSON `json:"-"`
}

// zarazHistoryUpdateResponseSettingsJSON contains the JSON metadata for the struct
// [ZarazHistoryUpdateResponseSettings]
type zarazHistoryUpdateResponseSettingsJSON struct {
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

func (r *ZarazHistoryUpdateResponseSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details of the worker that receives and edits Zaraz Context object.
type ZarazHistoryUpdateResponseSettingsContextEnricher struct {
	EscapedWorkerName string                                                `json:"escapedWorkerName,required"`
	WorkerTag         string                                                `json:"workerTag,required"`
	JSON              zarazHistoryUpdateResponseSettingsContextEnricherJSON `json:"-"`
}

// zarazHistoryUpdateResponseSettingsContextEnricherJSON contains the JSON metadata
// for the struct [ZarazHistoryUpdateResponseSettingsContextEnricher]
type zarazHistoryUpdateResponseSettingsContextEnricherJSON struct {
	EscapedWorkerName apijson.Field
	WorkerTag         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseSettingsContextEnricher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [ZarazHistoryUpdateResponseToolsZarazManagedComponent] or
// [ZarazHistoryUpdateResponseToolsZarazCustomManagedComponent].
type ZarazHistoryUpdateResponseTool interface {
	implementsZarazHistoryUpdateResponseTool()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZarazHistoryUpdateResponseTool)(nil)).Elem(), "")
}

type ZarazHistoryUpdateResponseToolsZarazManagedComponent struct {
	// List of blocking trigger IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Tool's internal name
	Component string `json:"component,required"`
	// Default fields for tool's actions
	DefaultFields map[string]ZarazHistoryUpdateResponseToolsZarazManagedComponentDefaultField `json:"defaultFields,required"`
	// Whether tool is enabled
	Enabled bool `json:"enabled,required"`
	// Tool's name defined by the user
	Name string `json:"name,required"`
	// List of permissions granted to the component
	Permissions []string `json:"permissions,required"`
	// Tool's settings
	Settings map[string]ZarazHistoryUpdateResponseToolsZarazManagedComponentSetting `json:"settings,required"`
	Type     ZarazHistoryUpdateResponseToolsZarazManagedComponentType               `json:"type,required"`
	// Actions configured on a tool. Either this or neoEvents field is required.
	Actions map[string]ZarazHistoryUpdateResponseToolsZarazManagedComponentAction `json:"actions"`
	// Default consent purpose ID
	DefaultPurpose string `json:"defaultPurpose"`
	// DEPRECATED - List of actions configured on a tool. Either this or actions field
	// is required. If both are present, actions field will take precedence.
	NeoEvents []ZarazHistoryUpdateResponseToolsZarazManagedComponentNeoEvent `json:"neoEvents"`
	JSON      zarazHistoryUpdateResponseToolsZarazManagedComponentJSON       `json:"-"`
}

// zarazHistoryUpdateResponseToolsZarazManagedComponentJSON contains the JSON
// metadata for the struct [ZarazHistoryUpdateResponseToolsZarazManagedComponent]
type zarazHistoryUpdateResponseToolsZarazManagedComponentJSON struct {
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

func (r *ZarazHistoryUpdateResponseToolsZarazManagedComponent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazHistoryUpdateResponseToolsZarazManagedComponent) implementsZarazHistoryUpdateResponseTool() {
}

// Union satisfied by [shared.UnionString] or [shared.UnionBool].
type ZarazHistoryUpdateResponseToolsZarazManagedComponentDefaultField interface {
	ImplementsZarazHistoryUpdateResponseToolsZarazManagedComponentDefaultField()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazHistoryUpdateResponseToolsZarazManagedComponentDefaultField)(nil)).Elem(),
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
type ZarazHistoryUpdateResponseToolsZarazManagedComponentSetting interface {
	ImplementsZarazHistoryUpdateResponseToolsZarazManagedComponentSetting()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazHistoryUpdateResponseToolsZarazManagedComponentSetting)(nil)).Elem(),
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

type ZarazHistoryUpdateResponseToolsZarazManagedComponentType string

const (
	ZarazHistoryUpdateResponseToolsZarazManagedComponentTypeComponent ZarazHistoryUpdateResponseToolsZarazManagedComponentType = "component"
)

type ZarazHistoryUpdateResponseToolsZarazManagedComponentAction struct {
	// Tool event type
	ActionType string `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Event payload
	Data interface{} `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers []string                                                       `json:"firingTriggers,required"`
	JSON           zarazHistoryUpdateResponseToolsZarazManagedComponentActionJSON `json:"-"`
}

// zarazHistoryUpdateResponseToolsZarazManagedComponentActionJSON contains the JSON
// metadata for the struct
// [ZarazHistoryUpdateResponseToolsZarazManagedComponentAction]
type zarazHistoryUpdateResponseToolsZarazManagedComponentActionJSON struct {
	ActionType       apijson.Field
	BlockingTriggers apijson.Field
	Data             apijson.Field
	FiringTriggers   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseToolsZarazManagedComponentAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryUpdateResponseToolsZarazManagedComponentNeoEvent struct {
	// Tool event type
	ActionType string `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Event payload
	Data interface{} `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers []string                                                         `json:"firingTriggers,required"`
	JSON           zarazHistoryUpdateResponseToolsZarazManagedComponentNeoEventJSON `json:"-"`
}

// zarazHistoryUpdateResponseToolsZarazManagedComponentNeoEventJSON contains the
// JSON metadata for the struct
// [ZarazHistoryUpdateResponseToolsZarazManagedComponentNeoEvent]
type zarazHistoryUpdateResponseToolsZarazManagedComponentNeoEventJSON struct {
	ActionType       apijson.Field
	BlockingTriggers apijson.Field
	Data             apijson.Field
	FiringTriggers   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseToolsZarazManagedComponentNeoEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryUpdateResponseToolsZarazCustomManagedComponent struct {
	// List of blocking trigger IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Tool's internal name
	Component string `json:"component,required"`
	// Default fields for tool's actions
	DefaultFields map[string]ZarazHistoryUpdateResponseToolsZarazCustomManagedComponentDefaultField `json:"defaultFields,required"`
	// Whether tool is enabled
	Enabled bool `json:"enabled,required"`
	// Tool's name defined by the user
	Name string `json:"name,required"`
	// List of permissions granted to the component
	Permissions []string `json:"permissions,required"`
	// Tool's settings
	Settings map[string]ZarazHistoryUpdateResponseToolsZarazCustomManagedComponentSetting `json:"settings,required"`
	Type     ZarazHistoryUpdateResponseToolsZarazCustomManagedComponentType               `json:"type,required"`
	// Cloudflare worker that acts as a managed component
	Worker ZarazHistoryUpdateResponseToolsZarazCustomManagedComponentWorker `json:"worker,required"`
	// Actions configured on a tool. Either this or neoEvents field is required.
	Actions map[string]ZarazHistoryUpdateResponseToolsZarazCustomManagedComponentAction `json:"actions"`
	// Default consent purpose ID
	DefaultPurpose string `json:"defaultPurpose"`
	// DEPRECATED - List of actions configured on a tool. Either this or actions field
	// is required. If both are present, actions field will take precedence.
	NeoEvents []ZarazHistoryUpdateResponseToolsZarazCustomManagedComponentNeoEvent `json:"neoEvents"`
	JSON      zarazHistoryUpdateResponseToolsZarazCustomManagedComponentJSON       `json:"-"`
}

// zarazHistoryUpdateResponseToolsZarazCustomManagedComponentJSON contains the JSON
// metadata for the struct
// [ZarazHistoryUpdateResponseToolsZarazCustomManagedComponent]
type zarazHistoryUpdateResponseToolsZarazCustomManagedComponentJSON struct {
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

func (r *ZarazHistoryUpdateResponseToolsZarazCustomManagedComponent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazHistoryUpdateResponseToolsZarazCustomManagedComponent) implementsZarazHistoryUpdateResponseTool() {
}

// Union satisfied by [shared.UnionString] or [shared.UnionBool].
type ZarazHistoryUpdateResponseToolsZarazCustomManagedComponentDefaultField interface {
	ImplementsZarazHistoryUpdateResponseToolsZarazCustomManagedComponentDefaultField()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazHistoryUpdateResponseToolsZarazCustomManagedComponentDefaultField)(nil)).Elem(),
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
type ZarazHistoryUpdateResponseToolsZarazCustomManagedComponentSetting interface {
	ImplementsZarazHistoryUpdateResponseToolsZarazCustomManagedComponentSetting()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazHistoryUpdateResponseToolsZarazCustomManagedComponentSetting)(nil)).Elem(),
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

type ZarazHistoryUpdateResponseToolsZarazCustomManagedComponentType string

const (
	ZarazHistoryUpdateResponseToolsZarazCustomManagedComponentTypeCustomMc ZarazHistoryUpdateResponseToolsZarazCustomManagedComponentType = "custom-mc"
)

// Cloudflare worker that acts as a managed component
type ZarazHistoryUpdateResponseToolsZarazCustomManagedComponentWorker struct {
	EscapedWorkerName string                                                               `json:"escapedWorkerName,required"`
	WorkerTag         string                                                               `json:"workerTag,required"`
	JSON              zarazHistoryUpdateResponseToolsZarazCustomManagedComponentWorkerJSON `json:"-"`
}

// zarazHistoryUpdateResponseToolsZarazCustomManagedComponentWorkerJSON contains
// the JSON metadata for the struct
// [ZarazHistoryUpdateResponseToolsZarazCustomManagedComponentWorker]
type zarazHistoryUpdateResponseToolsZarazCustomManagedComponentWorkerJSON struct {
	EscapedWorkerName apijson.Field
	WorkerTag         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseToolsZarazCustomManagedComponentWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryUpdateResponseToolsZarazCustomManagedComponentAction struct {
	// Tool event type
	ActionType string `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Event payload
	Data interface{} `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers []string                                                             `json:"firingTriggers,required"`
	JSON           zarazHistoryUpdateResponseToolsZarazCustomManagedComponentActionJSON `json:"-"`
}

// zarazHistoryUpdateResponseToolsZarazCustomManagedComponentActionJSON contains
// the JSON metadata for the struct
// [ZarazHistoryUpdateResponseToolsZarazCustomManagedComponentAction]
type zarazHistoryUpdateResponseToolsZarazCustomManagedComponentActionJSON struct {
	ActionType       apijson.Field
	BlockingTriggers apijson.Field
	Data             apijson.Field
	FiringTriggers   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseToolsZarazCustomManagedComponentAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryUpdateResponseToolsZarazCustomManagedComponentNeoEvent struct {
	// Tool event type
	ActionType string `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Event payload
	Data interface{} `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers []string                                                               `json:"firingTriggers,required"`
	JSON           zarazHistoryUpdateResponseToolsZarazCustomManagedComponentNeoEventJSON `json:"-"`
}

// zarazHistoryUpdateResponseToolsZarazCustomManagedComponentNeoEventJSON contains
// the JSON metadata for the struct
// [ZarazHistoryUpdateResponseToolsZarazCustomManagedComponentNeoEvent]
type zarazHistoryUpdateResponseToolsZarazCustomManagedComponentNeoEventJSON struct {
	ActionType       apijson.Field
	BlockingTriggers apijson.Field
	Data             apijson.Field
	FiringTriggers   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseToolsZarazCustomManagedComponentNeoEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryUpdateResponseTrigger struct {
	// Rules defining when the trigger is not fired.
	ExcludeRules []ZarazHistoryUpdateResponseTriggersExcludeRule `json:"excludeRules,required"`
	// Rules defining when the trigger is fired.
	LoadRules []ZarazHistoryUpdateResponseTriggersLoadRule `json:"loadRules,required"`
	// Trigger name.
	Name string `json:"name,required"`
	// Trigger description.
	Description string                                   `json:"description"`
	System      ZarazHistoryUpdateResponseTriggersSystem `json:"system"`
	JSON        zarazHistoryUpdateResponseTriggerJSON    `json:"-"`
}

// zarazHistoryUpdateResponseTriggerJSON contains the JSON metadata for the struct
// [ZarazHistoryUpdateResponseTrigger]
type zarazHistoryUpdateResponseTriggerJSON struct {
	ExcludeRules apijson.Field
	LoadRules    apijson.Field
	Name         apijson.Field
	Description  apijson.Field
	System       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [ZarazHistoryUpdateResponseTriggersExcludeRulesZarazLoadRule],
// [ZarazHistoryUpdateResponseTriggersExcludeRulesZarazClickListenerRule],
// [ZarazHistoryUpdateResponseTriggersExcludeRulesZarazTimerRule],
// [ZarazHistoryUpdateResponseTriggersExcludeRulesZarazFormSubmissionRule],
// [ZarazHistoryUpdateResponseTriggersExcludeRulesZarazVariableMatchRule],
// [ZarazHistoryUpdateResponseTriggersExcludeRulesZarazScrollDepthRule] or
// [ZarazHistoryUpdateResponseTriggersExcludeRulesZarazElementVisibilityRule].
type ZarazHistoryUpdateResponseTriggersExcludeRule interface {
	implementsZarazHistoryUpdateResponseTriggersExcludeRule()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZarazHistoryUpdateResponseTriggersExcludeRule)(nil)).Elem(), "")
}

type ZarazHistoryUpdateResponseTriggersExcludeRulesZarazLoadRule struct {
	ID    string                                                          `json:"id,required"`
	Match string                                                          `json:"match,required"`
	Op    ZarazHistoryUpdateResponseTriggersExcludeRulesZarazLoadRuleOp   `json:"op,required"`
	Value string                                                          `json:"value,required"`
	JSON  zarazHistoryUpdateResponseTriggersExcludeRulesZarazLoadRuleJSON `json:"-"`
}

// zarazHistoryUpdateResponseTriggersExcludeRulesZarazLoadRuleJSON contains the
// JSON metadata for the struct
// [ZarazHistoryUpdateResponseTriggersExcludeRulesZarazLoadRule]
type zarazHistoryUpdateResponseTriggersExcludeRulesZarazLoadRuleJSON struct {
	ID          apijson.Field
	Match       apijson.Field
	Op          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseTriggersExcludeRulesZarazLoadRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazHistoryUpdateResponseTriggersExcludeRulesZarazLoadRule) implementsZarazHistoryUpdateResponseTriggersExcludeRule() {
}

type ZarazHistoryUpdateResponseTriggersExcludeRulesZarazLoadRuleOp string

const (
	ZarazHistoryUpdateResponseTriggersExcludeRulesZarazLoadRuleOpContains           ZarazHistoryUpdateResponseTriggersExcludeRulesZarazLoadRuleOp = "CONTAINS"
	ZarazHistoryUpdateResponseTriggersExcludeRulesZarazLoadRuleOpEquals             ZarazHistoryUpdateResponseTriggersExcludeRulesZarazLoadRuleOp = "EQUALS"
	ZarazHistoryUpdateResponseTriggersExcludeRulesZarazLoadRuleOpStartsWith         ZarazHistoryUpdateResponseTriggersExcludeRulesZarazLoadRuleOp = "STARTS_WITH"
	ZarazHistoryUpdateResponseTriggersExcludeRulesZarazLoadRuleOpEndsWith           ZarazHistoryUpdateResponseTriggersExcludeRulesZarazLoadRuleOp = "ENDS_WITH"
	ZarazHistoryUpdateResponseTriggersExcludeRulesZarazLoadRuleOpMatchRegex         ZarazHistoryUpdateResponseTriggersExcludeRulesZarazLoadRuleOp = "MATCH_REGEX"
	ZarazHistoryUpdateResponseTriggersExcludeRulesZarazLoadRuleOpNotMatchRegex      ZarazHistoryUpdateResponseTriggersExcludeRulesZarazLoadRuleOp = "NOT_MATCH_REGEX"
	ZarazHistoryUpdateResponseTriggersExcludeRulesZarazLoadRuleOpGreaterThan        ZarazHistoryUpdateResponseTriggersExcludeRulesZarazLoadRuleOp = "GREATER_THAN"
	ZarazHistoryUpdateResponseTriggersExcludeRulesZarazLoadRuleOpGreaterThanOrEqual ZarazHistoryUpdateResponseTriggersExcludeRulesZarazLoadRuleOp = "GREATER_THAN_OR_EQUAL"
	ZarazHistoryUpdateResponseTriggersExcludeRulesZarazLoadRuleOpLessThan           ZarazHistoryUpdateResponseTriggersExcludeRulesZarazLoadRuleOp = "LESS_THAN"
	ZarazHistoryUpdateResponseTriggersExcludeRulesZarazLoadRuleOpLessThanOrEqual    ZarazHistoryUpdateResponseTriggersExcludeRulesZarazLoadRuleOp = "LESS_THAN_OR_EQUAL"
)

type ZarazHistoryUpdateResponseTriggersExcludeRulesZarazClickListenerRule struct {
	ID       string                                                                       `json:"id,required"`
	Action   ZarazHistoryUpdateResponseTriggersExcludeRulesZarazClickListenerRuleAction   `json:"action,required"`
	Settings ZarazHistoryUpdateResponseTriggersExcludeRulesZarazClickListenerRuleSettings `json:"settings,required"`
	JSON     zarazHistoryUpdateResponseTriggersExcludeRulesZarazClickListenerRuleJSON     `json:"-"`
}

// zarazHistoryUpdateResponseTriggersExcludeRulesZarazClickListenerRuleJSON
// contains the JSON metadata for the struct
// [ZarazHistoryUpdateResponseTriggersExcludeRulesZarazClickListenerRule]
type zarazHistoryUpdateResponseTriggersExcludeRulesZarazClickListenerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseTriggersExcludeRulesZarazClickListenerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazHistoryUpdateResponseTriggersExcludeRulesZarazClickListenerRule) implementsZarazHistoryUpdateResponseTriggersExcludeRule() {
}

type ZarazHistoryUpdateResponseTriggersExcludeRulesZarazClickListenerRuleAction string

const (
	ZarazHistoryUpdateResponseTriggersExcludeRulesZarazClickListenerRuleActionClickListener ZarazHistoryUpdateResponseTriggersExcludeRulesZarazClickListenerRuleAction = "clickListener"
)

type ZarazHistoryUpdateResponseTriggersExcludeRulesZarazClickListenerRuleSettings struct {
	Selector    string                                                                           `json:"selector,required"`
	Type        ZarazHistoryUpdateResponseTriggersExcludeRulesZarazClickListenerRuleSettingsType `json:"type,required"`
	WaitForTags int64                                                                            `json:"waitForTags,required"`
	JSON        zarazHistoryUpdateResponseTriggersExcludeRulesZarazClickListenerRuleSettingsJSON `json:"-"`
}

// zarazHistoryUpdateResponseTriggersExcludeRulesZarazClickListenerRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazHistoryUpdateResponseTriggersExcludeRulesZarazClickListenerRuleSettings]
type zarazHistoryUpdateResponseTriggersExcludeRulesZarazClickListenerRuleSettingsJSON struct {
	Selector    apijson.Field
	Type        apijson.Field
	WaitForTags apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseTriggersExcludeRulesZarazClickListenerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryUpdateResponseTriggersExcludeRulesZarazClickListenerRuleSettingsType string

const (
	ZarazHistoryUpdateResponseTriggersExcludeRulesZarazClickListenerRuleSettingsTypeXpath ZarazHistoryUpdateResponseTriggersExcludeRulesZarazClickListenerRuleSettingsType = "xpath"
	ZarazHistoryUpdateResponseTriggersExcludeRulesZarazClickListenerRuleSettingsTypeCss   ZarazHistoryUpdateResponseTriggersExcludeRulesZarazClickListenerRuleSettingsType = "css"
)

type ZarazHistoryUpdateResponseTriggersExcludeRulesZarazTimerRule struct {
	ID       string                                                               `json:"id,required"`
	Action   ZarazHistoryUpdateResponseTriggersExcludeRulesZarazTimerRuleAction   `json:"action,required"`
	Settings ZarazHistoryUpdateResponseTriggersExcludeRulesZarazTimerRuleSettings `json:"settings,required"`
	JSON     zarazHistoryUpdateResponseTriggersExcludeRulesZarazTimerRuleJSON     `json:"-"`
}

// zarazHistoryUpdateResponseTriggersExcludeRulesZarazTimerRuleJSON contains the
// JSON metadata for the struct
// [ZarazHistoryUpdateResponseTriggersExcludeRulesZarazTimerRule]
type zarazHistoryUpdateResponseTriggersExcludeRulesZarazTimerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseTriggersExcludeRulesZarazTimerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazHistoryUpdateResponseTriggersExcludeRulesZarazTimerRule) implementsZarazHistoryUpdateResponseTriggersExcludeRule() {
}

type ZarazHistoryUpdateResponseTriggersExcludeRulesZarazTimerRuleAction string

const (
	ZarazHistoryUpdateResponseTriggersExcludeRulesZarazTimerRuleActionTimer ZarazHistoryUpdateResponseTriggersExcludeRulesZarazTimerRuleAction = "timer"
)

type ZarazHistoryUpdateResponseTriggersExcludeRulesZarazTimerRuleSettings struct {
	Interval int64                                                                    `json:"interval,required"`
	Limit    int64                                                                    `json:"limit,required"`
	JSON     zarazHistoryUpdateResponseTriggersExcludeRulesZarazTimerRuleSettingsJSON `json:"-"`
}

// zarazHistoryUpdateResponseTriggersExcludeRulesZarazTimerRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazHistoryUpdateResponseTriggersExcludeRulesZarazTimerRuleSettings]
type zarazHistoryUpdateResponseTriggersExcludeRulesZarazTimerRuleSettingsJSON struct {
	Interval    apijson.Field
	Limit       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseTriggersExcludeRulesZarazTimerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryUpdateResponseTriggersExcludeRulesZarazFormSubmissionRule struct {
	ID       string                                                                        `json:"id,required"`
	Action   ZarazHistoryUpdateResponseTriggersExcludeRulesZarazFormSubmissionRuleAction   `json:"action,required"`
	Settings ZarazHistoryUpdateResponseTriggersExcludeRulesZarazFormSubmissionRuleSettings `json:"settings,required"`
	JSON     zarazHistoryUpdateResponseTriggersExcludeRulesZarazFormSubmissionRuleJSON     `json:"-"`
}

// zarazHistoryUpdateResponseTriggersExcludeRulesZarazFormSubmissionRuleJSON
// contains the JSON metadata for the struct
// [ZarazHistoryUpdateResponseTriggersExcludeRulesZarazFormSubmissionRule]
type zarazHistoryUpdateResponseTriggersExcludeRulesZarazFormSubmissionRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseTriggersExcludeRulesZarazFormSubmissionRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazHistoryUpdateResponseTriggersExcludeRulesZarazFormSubmissionRule) implementsZarazHistoryUpdateResponseTriggersExcludeRule() {
}

type ZarazHistoryUpdateResponseTriggersExcludeRulesZarazFormSubmissionRuleAction string

const (
	ZarazHistoryUpdateResponseTriggersExcludeRulesZarazFormSubmissionRuleActionFormSubmission ZarazHistoryUpdateResponseTriggersExcludeRulesZarazFormSubmissionRuleAction = "formSubmission"
)

type ZarazHistoryUpdateResponseTriggersExcludeRulesZarazFormSubmissionRuleSettings struct {
	Selector string                                                                            `json:"selector,required"`
	Validate bool                                                                              `json:"validate,required"`
	JSON     zarazHistoryUpdateResponseTriggersExcludeRulesZarazFormSubmissionRuleSettingsJSON `json:"-"`
}

// zarazHistoryUpdateResponseTriggersExcludeRulesZarazFormSubmissionRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazHistoryUpdateResponseTriggersExcludeRulesZarazFormSubmissionRuleSettings]
type zarazHistoryUpdateResponseTriggersExcludeRulesZarazFormSubmissionRuleSettingsJSON struct {
	Selector    apijson.Field
	Validate    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseTriggersExcludeRulesZarazFormSubmissionRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryUpdateResponseTriggersExcludeRulesZarazVariableMatchRule struct {
	ID       string                                                                       `json:"id,required"`
	Action   ZarazHistoryUpdateResponseTriggersExcludeRulesZarazVariableMatchRuleAction   `json:"action,required"`
	Settings ZarazHistoryUpdateResponseTriggersExcludeRulesZarazVariableMatchRuleSettings `json:"settings,required"`
	JSON     zarazHistoryUpdateResponseTriggersExcludeRulesZarazVariableMatchRuleJSON     `json:"-"`
}

// zarazHistoryUpdateResponseTriggersExcludeRulesZarazVariableMatchRuleJSON
// contains the JSON metadata for the struct
// [ZarazHistoryUpdateResponseTriggersExcludeRulesZarazVariableMatchRule]
type zarazHistoryUpdateResponseTriggersExcludeRulesZarazVariableMatchRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseTriggersExcludeRulesZarazVariableMatchRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazHistoryUpdateResponseTriggersExcludeRulesZarazVariableMatchRule) implementsZarazHistoryUpdateResponseTriggersExcludeRule() {
}

type ZarazHistoryUpdateResponseTriggersExcludeRulesZarazVariableMatchRuleAction string

const (
	ZarazHistoryUpdateResponseTriggersExcludeRulesZarazVariableMatchRuleActionVariableMatch ZarazHistoryUpdateResponseTriggersExcludeRulesZarazVariableMatchRuleAction = "variableMatch"
)

type ZarazHistoryUpdateResponseTriggersExcludeRulesZarazVariableMatchRuleSettings struct {
	Match    string                                                                           `json:"match,required"`
	Variable string                                                                           `json:"variable,required"`
	JSON     zarazHistoryUpdateResponseTriggersExcludeRulesZarazVariableMatchRuleSettingsJSON `json:"-"`
}

// zarazHistoryUpdateResponseTriggersExcludeRulesZarazVariableMatchRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazHistoryUpdateResponseTriggersExcludeRulesZarazVariableMatchRuleSettings]
type zarazHistoryUpdateResponseTriggersExcludeRulesZarazVariableMatchRuleSettingsJSON struct {
	Match       apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseTriggersExcludeRulesZarazVariableMatchRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryUpdateResponseTriggersExcludeRulesZarazScrollDepthRule struct {
	ID       string                                                                     `json:"id,required"`
	Action   ZarazHistoryUpdateResponseTriggersExcludeRulesZarazScrollDepthRuleAction   `json:"action,required"`
	Settings ZarazHistoryUpdateResponseTriggersExcludeRulesZarazScrollDepthRuleSettings `json:"settings,required"`
	JSON     zarazHistoryUpdateResponseTriggersExcludeRulesZarazScrollDepthRuleJSON     `json:"-"`
}

// zarazHistoryUpdateResponseTriggersExcludeRulesZarazScrollDepthRuleJSON contains
// the JSON metadata for the struct
// [ZarazHistoryUpdateResponseTriggersExcludeRulesZarazScrollDepthRule]
type zarazHistoryUpdateResponseTriggersExcludeRulesZarazScrollDepthRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseTriggersExcludeRulesZarazScrollDepthRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazHistoryUpdateResponseTriggersExcludeRulesZarazScrollDepthRule) implementsZarazHistoryUpdateResponseTriggersExcludeRule() {
}

type ZarazHistoryUpdateResponseTriggersExcludeRulesZarazScrollDepthRuleAction string

const (
	ZarazHistoryUpdateResponseTriggersExcludeRulesZarazScrollDepthRuleActionScrollDepth ZarazHistoryUpdateResponseTriggersExcludeRulesZarazScrollDepthRuleAction = "scrollDepth"
)

type ZarazHistoryUpdateResponseTriggersExcludeRulesZarazScrollDepthRuleSettings struct {
	Positions string                                                                         `json:"positions,required"`
	JSON      zarazHistoryUpdateResponseTriggersExcludeRulesZarazScrollDepthRuleSettingsJSON `json:"-"`
}

// zarazHistoryUpdateResponseTriggersExcludeRulesZarazScrollDepthRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazHistoryUpdateResponseTriggersExcludeRulesZarazScrollDepthRuleSettings]
type zarazHistoryUpdateResponseTriggersExcludeRulesZarazScrollDepthRuleSettingsJSON struct {
	Positions   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseTriggersExcludeRulesZarazScrollDepthRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryUpdateResponseTriggersExcludeRulesZarazElementVisibilityRule struct {
	ID       string                                                                           `json:"id,required"`
	Action   ZarazHistoryUpdateResponseTriggersExcludeRulesZarazElementVisibilityRuleAction   `json:"action,required"`
	Settings ZarazHistoryUpdateResponseTriggersExcludeRulesZarazElementVisibilityRuleSettings `json:"settings,required"`
	JSON     zarazHistoryUpdateResponseTriggersExcludeRulesZarazElementVisibilityRuleJSON     `json:"-"`
}

// zarazHistoryUpdateResponseTriggersExcludeRulesZarazElementVisibilityRuleJSON
// contains the JSON metadata for the struct
// [ZarazHistoryUpdateResponseTriggersExcludeRulesZarazElementVisibilityRule]
type zarazHistoryUpdateResponseTriggersExcludeRulesZarazElementVisibilityRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseTriggersExcludeRulesZarazElementVisibilityRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazHistoryUpdateResponseTriggersExcludeRulesZarazElementVisibilityRule) implementsZarazHistoryUpdateResponseTriggersExcludeRule() {
}

type ZarazHistoryUpdateResponseTriggersExcludeRulesZarazElementVisibilityRuleAction string

const (
	ZarazHistoryUpdateResponseTriggersExcludeRulesZarazElementVisibilityRuleActionElementVisibility ZarazHistoryUpdateResponseTriggersExcludeRulesZarazElementVisibilityRuleAction = "elementVisibility"
)

type ZarazHistoryUpdateResponseTriggersExcludeRulesZarazElementVisibilityRuleSettings struct {
	Selector string                                                                               `json:"selector,required"`
	JSON     zarazHistoryUpdateResponseTriggersExcludeRulesZarazElementVisibilityRuleSettingsJSON `json:"-"`
}

// zarazHistoryUpdateResponseTriggersExcludeRulesZarazElementVisibilityRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazHistoryUpdateResponseTriggersExcludeRulesZarazElementVisibilityRuleSettings]
type zarazHistoryUpdateResponseTriggersExcludeRulesZarazElementVisibilityRuleSettingsJSON struct {
	Selector    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseTriggersExcludeRulesZarazElementVisibilityRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [ZarazHistoryUpdateResponseTriggersLoadRulesZarazLoadRule],
// [ZarazHistoryUpdateResponseTriggersLoadRulesZarazClickListenerRule],
// [ZarazHistoryUpdateResponseTriggersLoadRulesZarazTimerRule],
// [ZarazHistoryUpdateResponseTriggersLoadRulesZarazFormSubmissionRule],
// [ZarazHistoryUpdateResponseTriggersLoadRulesZarazVariableMatchRule],
// [ZarazHistoryUpdateResponseTriggersLoadRulesZarazScrollDepthRule] or
// [ZarazHistoryUpdateResponseTriggersLoadRulesZarazElementVisibilityRule].
type ZarazHistoryUpdateResponseTriggersLoadRule interface {
	implementsZarazHistoryUpdateResponseTriggersLoadRule()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZarazHistoryUpdateResponseTriggersLoadRule)(nil)).Elem(), "")
}

type ZarazHistoryUpdateResponseTriggersLoadRulesZarazLoadRule struct {
	ID    string                                                       `json:"id,required"`
	Match string                                                       `json:"match,required"`
	Op    ZarazHistoryUpdateResponseTriggersLoadRulesZarazLoadRuleOp   `json:"op,required"`
	Value string                                                       `json:"value,required"`
	JSON  zarazHistoryUpdateResponseTriggersLoadRulesZarazLoadRuleJSON `json:"-"`
}

// zarazHistoryUpdateResponseTriggersLoadRulesZarazLoadRuleJSON contains the JSON
// metadata for the struct
// [ZarazHistoryUpdateResponseTriggersLoadRulesZarazLoadRule]
type zarazHistoryUpdateResponseTriggersLoadRulesZarazLoadRuleJSON struct {
	ID          apijson.Field
	Match       apijson.Field
	Op          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseTriggersLoadRulesZarazLoadRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazHistoryUpdateResponseTriggersLoadRulesZarazLoadRule) implementsZarazHistoryUpdateResponseTriggersLoadRule() {
}

type ZarazHistoryUpdateResponseTriggersLoadRulesZarazLoadRuleOp string

const (
	ZarazHistoryUpdateResponseTriggersLoadRulesZarazLoadRuleOpContains           ZarazHistoryUpdateResponseTriggersLoadRulesZarazLoadRuleOp = "CONTAINS"
	ZarazHistoryUpdateResponseTriggersLoadRulesZarazLoadRuleOpEquals             ZarazHistoryUpdateResponseTriggersLoadRulesZarazLoadRuleOp = "EQUALS"
	ZarazHistoryUpdateResponseTriggersLoadRulesZarazLoadRuleOpStartsWith         ZarazHistoryUpdateResponseTriggersLoadRulesZarazLoadRuleOp = "STARTS_WITH"
	ZarazHistoryUpdateResponseTriggersLoadRulesZarazLoadRuleOpEndsWith           ZarazHistoryUpdateResponseTriggersLoadRulesZarazLoadRuleOp = "ENDS_WITH"
	ZarazHistoryUpdateResponseTriggersLoadRulesZarazLoadRuleOpMatchRegex         ZarazHistoryUpdateResponseTriggersLoadRulesZarazLoadRuleOp = "MATCH_REGEX"
	ZarazHistoryUpdateResponseTriggersLoadRulesZarazLoadRuleOpNotMatchRegex      ZarazHistoryUpdateResponseTriggersLoadRulesZarazLoadRuleOp = "NOT_MATCH_REGEX"
	ZarazHistoryUpdateResponseTriggersLoadRulesZarazLoadRuleOpGreaterThan        ZarazHistoryUpdateResponseTriggersLoadRulesZarazLoadRuleOp = "GREATER_THAN"
	ZarazHistoryUpdateResponseTriggersLoadRulesZarazLoadRuleOpGreaterThanOrEqual ZarazHistoryUpdateResponseTriggersLoadRulesZarazLoadRuleOp = "GREATER_THAN_OR_EQUAL"
	ZarazHistoryUpdateResponseTriggersLoadRulesZarazLoadRuleOpLessThan           ZarazHistoryUpdateResponseTriggersLoadRulesZarazLoadRuleOp = "LESS_THAN"
	ZarazHistoryUpdateResponseTriggersLoadRulesZarazLoadRuleOpLessThanOrEqual    ZarazHistoryUpdateResponseTriggersLoadRulesZarazLoadRuleOp = "LESS_THAN_OR_EQUAL"
)

type ZarazHistoryUpdateResponseTriggersLoadRulesZarazClickListenerRule struct {
	ID       string                                                                    `json:"id,required"`
	Action   ZarazHistoryUpdateResponseTriggersLoadRulesZarazClickListenerRuleAction   `json:"action,required"`
	Settings ZarazHistoryUpdateResponseTriggersLoadRulesZarazClickListenerRuleSettings `json:"settings,required"`
	JSON     zarazHistoryUpdateResponseTriggersLoadRulesZarazClickListenerRuleJSON     `json:"-"`
}

// zarazHistoryUpdateResponseTriggersLoadRulesZarazClickListenerRuleJSON contains
// the JSON metadata for the struct
// [ZarazHistoryUpdateResponseTriggersLoadRulesZarazClickListenerRule]
type zarazHistoryUpdateResponseTriggersLoadRulesZarazClickListenerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseTriggersLoadRulesZarazClickListenerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazHistoryUpdateResponseTriggersLoadRulesZarazClickListenerRule) implementsZarazHistoryUpdateResponseTriggersLoadRule() {
}

type ZarazHistoryUpdateResponseTriggersLoadRulesZarazClickListenerRuleAction string

const (
	ZarazHistoryUpdateResponseTriggersLoadRulesZarazClickListenerRuleActionClickListener ZarazHistoryUpdateResponseTriggersLoadRulesZarazClickListenerRuleAction = "clickListener"
)

type ZarazHistoryUpdateResponseTriggersLoadRulesZarazClickListenerRuleSettings struct {
	Selector    string                                                                        `json:"selector,required"`
	Type        ZarazHistoryUpdateResponseTriggersLoadRulesZarazClickListenerRuleSettingsType `json:"type,required"`
	WaitForTags int64                                                                         `json:"waitForTags,required"`
	JSON        zarazHistoryUpdateResponseTriggersLoadRulesZarazClickListenerRuleSettingsJSON `json:"-"`
}

// zarazHistoryUpdateResponseTriggersLoadRulesZarazClickListenerRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazHistoryUpdateResponseTriggersLoadRulesZarazClickListenerRuleSettings]
type zarazHistoryUpdateResponseTriggersLoadRulesZarazClickListenerRuleSettingsJSON struct {
	Selector    apijson.Field
	Type        apijson.Field
	WaitForTags apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseTriggersLoadRulesZarazClickListenerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryUpdateResponseTriggersLoadRulesZarazClickListenerRuleSettingsType string

const (
	ZarazHistoryUpdateResponseTriggersLoadRulesZarazClickListenerRuleSettingsTypeXpath ZarazHistoryUpdateResponseTriggersLoadRulesZarazClickListenerRuleSettingsType = "xpath"
	ZarazHistoryUpdateResponseTriggersLoadRulesZarazClickListenerRuleSettingsTypeCss   ZarazHistoryUpdateResponseTriggersLoadRulesZarazClickListenerRuleSettingsType = "css"
)

type ZarazHistoryUpdateResponseTriggersLoadRulesZarazTimerRule struct {
	ID       string                                                            `json:"id,required"`
	Action   ZarazHistoryUpdateResponseTriggersLoadRulesZarazTimerRuleAction   `json:"action,required"`
	Settings ZarazHistoryUpdateResponseTriggersLoadRulesZarazTimerRuleSettings `json:"settings,required"`
	JSON     zarazHistoryUpdateResponseTriggersLoadRulesZarazTimerRuleJSON     `json:"-"`
}

// zarazHistoryUpdateResponseTriggersLoadRulesZarazTimerRuleJSON contains the JSON
// metadata for the struct
// [ZarazHistoryUpdateResponseTriggersLoadRulesZarazTimerRule]
type zarazHistoryUpdateResponseTriggersLoadRulesZarazTimerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseTriggersLoadRulesZarazTimerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazHistoryUpdateResponseTriggersLoadRulesZarazTimerRule) implementsZarazHistoryUpdateResponseTriggersLoadRule() {
}

type ZarazHistoryUpdateResponseTriggersLoadRulesZarazTimerRuleAction string

const (
	ZarazHistoryUpdateResponseTriggersLoadRulesZarazTimerRuleActionTimer ZarazHistoryUpdateResponseTriggersLoadRulesZarazTimerRuleAction = "timer"
)

type ZarazHistoryUpdateResponseTriggersLoadRulesZarazTimerRuleSettings struct {
	Interval int64                                                                 `json:"interval,required"`
	Limit    int64                                                                 `json:"limit,required"`
	JSON     zarazHistoryUpdateResponseTriggersLoadRulesZarazTimerRuleSettingsJSON `json:"-"`
}

// zarazHistoryUpdateResponseTriggersLoadRulesZarazTimerRuleSettingsJSON contains
// the JSON metadata for the struct
// [ZarazHistoryUpdateResponseTriggersLoadRulesZarazTimerRuleSettings]
type zarazHistoryUpdateResponseTriggersLoadRulesZarazTimerRuleSettingsJSON struct {
	Interval    apijson.Field
	Limit       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseTriggersLoadRulesZarazTimerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryUpdateResponseTriggersLoadRulesZarazFormSubmissionRule struct {
	ID       string                                                                     `json:"id,required"`
	Action   ZarazHistoryUpdateResponseTriggersLoadRulesZarazFormSubmissionRuleAction   `json:"action,required"`
	Settings ZarazHistoryUpdateResponseTriggersLoadRulesZarazFormSubmissionRuleSettings `json:"settings,required"`
	JSON     zarazHistoryUpdateResponseTriggersLoadRulesZarazFormSubmissionRuleJSON     `json:"-"`
}

// zarazHistoryUpdateResponseTriggersLoadRulesZarazFormSubmissionRuleJSON contains
// the JSON metadata for the struct
// [ZarazHistoryUpdateResponseTriggersLoadRulesZarazFormSubmissionRule]
type zarazHistoryUpdateResponseTriggersLoadRulesZarazFormSubmissionRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseTriggersLoadRulesZarazFormSubmissionRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazHistoryUpdateResponseTriggersLoadRulesZarazFormSubmissionRule) implementsZarazHistoryUpdateResponseTriggersLoadRule() {
}

type ZarazHistoryUpdateResponseTriggersLoadRulesZarazFormSubmissionRuleAction string

const (
	ZarazHistoryUpdateResponseTriggersLoadRulesZarazFormSubmissionRuleActionFormSubmission ZarazHistoryUpdateResponseTriggersLoadRulesZarazFormSubmissionRuleAction = "formSubmission"
)

type ZarazHistoryUpdateResponseTriggersLoadRulesZarazFormSubmissionRuleSettings struct {
	Selector string                                                                         `json:"selector,required"`
	Validate bool                                                                           `json:"validate,required"`
	JSON     zarazHistoryUpdateResponseTriggersLoadRulesZarazFormSubmissionRuleSettingsJSON `json:"-"`
}

// zarazHistoryUpdateResponseTriggersLoadRulesZarazFormSubmissionRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazHistoryUpdateResponseTriggersLoadRulesZarazFormSubmissionRuleSettings]
type zarazHistoryUpdateResponseTriggersLoadRulesZarazFormSubmissionRuleSettingsJSON struct {
	Selector    apijson.Field
	Validate    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseTriggersLoadRulesZarazFormSubmissionRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryUpdateResponseTriggersLoadRulesZarazVariableMatchRule struct {
	ID       string                                                                    `json:"id,required"`
	Action   ZarazHistoryUpdateResponseTriggersLoadRulesZarazVariableMatchRuleAction   `json:"action,required"`
	Settings ZarazHistoryUpdateResponseTriggersLoadRulesZarazVariableMatchRuleSettings `json:"settings,required"`
	JSON     zarazHistoryUpdateResponseTriggersLoadRulesZarazVariableMatchRuleJSON     `json:"-"`
}

// zarazHistoryUpdateResponseTriggersLoadRulesZarazVariableMatchRuleJSON contains
// the JSON metadata for the struct
// [ZarazHistoryUpdateResponseTriggersLoadRulesZarazVariableMatchRule]
type zarazHistoryUpdateResponseTriggersLoadRulesZarazVariableMatchRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseTriggersLoadRulesZarazVariableMatchRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazHistoryUpdateResponseTriggersLoadRulesZarazVariableMatchRule) implementsZarazHistoryUpdateResponseTriggersLoadRule() {
}

type ZarazHistoryUpdateResponseTriggersLoadRulesZarazVariableMatchRuleAction string

const (
	ZarazHistoryUpdateResponseTriggersLoadRulesZarazVariableMatchRuleActionVariableMatch ZarazHistoryUpdateResponseTriggersLoadRulesZarazVariableMatchRuleAction = "variableMatch"
)

type ZarazHistoryUpdateResponseTriggersLoadRulesZarazVariableMatchRuleSettings struct {
	Match    string                                                                        `json:"match,required"`
	Variable string                                                                        `json:"variable,required"`
	JSON     zarazHistoryUpdateResponseTriggersLoadRulesZarazVariableMatchRuleSettingsJSON `json:"-"`
}

// zarazHistoryUpdateResponseTriggersLoadRulesZarazVariableMatchRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazHistoryUpdateResponseTriggersLoadRulesZarazVariableMatchRuleSettings]
type zarazHistoryUpdateResponseTriggersLoadRulesZarazVariableMatchRuleSettingsJSON struct {
	Match       apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseTriggersLoadRulesZarazVariableMatchRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryUpdateResponseTriggersLoadRulesZarazScrollDepthRule struct {
	ID       string                                                                  `json:"id,required"`
	Action   ZarazHistoryUpdateResponseTriggersLoadRulesZarazScrollDepthRuleAction   `json:"action,required"`
	Settings ZarazHistoryUpdateResponseTriggersLoadRulesZarazScrollDepthRuleSettings `json:"settings,required"`
	JSON     zarazHistoryUpdateResponseTriggersLoadRulesZarazScrollDepthRuleJSON     `json:"-"`
}

// zarazHistoryUpdateResponseTriggersLoadRulesZarazScrollDepthRuleJSON contains the
// JSON metadata for the struct
// [ZarazHistoryUpdateResponseTriggersLoadRulesZarazScrollDepthRule]
type zarazHistoryUpdateResponseTriggersLoadRulesZarazScrollDepthRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseTriggersLoadRulesZarazScrollDepthRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazHistoryUpdateResponseTriggersLoadRulesZarazScrollDepthRule) implementsZarazHistoryUpdateResponseTriggersLoadRule() {
}

type ZarazHistoryUpdateResponseTriggersLoadRulesZarazScrollDepthRuleAction string

const (
	ZarazHistoryUpdateResponseTriggersLoadRulesZarazScrollDepthRuleActionScrollDepth ZarazHistoryUpdateResponseTriggersLoadRulesZarazScrollDepthRuleAction = "scrollDepth"
)

type ZarazHistoryUpdateResponseTriggersLoadRulesZarazScrollDepthRuleSettings struct {
	Positions string                                                                      `json:"positions,required"`
	JSON      zarazHistoryUpdateResponseTriggersLoadRulesZarazScrollDepthRuleSettingsJSON `json:"-"`
}

// zarazHistoryUpdateResponseTriggersLoadRulesZarazScrollDepthRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazHistoryUpdateResponseTriggersLoadRulesZarazScrollDepthRuleSettings]
type zarazHistoryUpdateResponseTriggersLoadRulesZarazScrollDepthRuleSettingsJSON struct {
	Positions   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseTriggersLoadRulesZarazScrollDepthRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryUpdateResponseTriggersLoadRulesZarazElementVisibilityRule struct {
	ID       string                                                                        `json:"id,required"`
	Action   ZarazHistoryUpdateResponseTriggersLoadRulesZarazElementVisibilityRuleAction   `json:"action,required"`
	Settings ZarazHistoryUpdateResponseTriggersLoadRulesZarazElementVisibilityRuleSettings `json:"settings,required"`
	JSON     zarazHistoryUpdateResponseTriggersLoadRulesZarazElementVisibilityRuleJSON     `json:"-"`
}

// zarazHistoryUpdateResponseTriggersLoadRulesZarazElementVisibilityRuleJSON
// contains the JSON metadata for the struct
// [ZarazHistoryUpdateResponseTriggersLoadRulesZarazElementVisibilityRule]
type zarazHistoryUpdateResponseTriggersLoadRulesZarazElementVisibilityRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseTriggersLoadRulesZarazElementVisibilityRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazHistoryUpdateResponseTriggersLoadRulesZarazElementVisibilityRule) implementsZarazHistoryUpdateResponseTriggersLoadRule() {
}

type ZarazHistoryUpdateResponseTriggersLoadRulesZarazElementVisibilityRuleAction string

const (
	ZarazHistoryUpdateResponseTriggersLoadRulesZarazElementVisibilityRuleActionElementVisibility ZarazHistoryUpdateResponseTriggersLoadRulesZarazElementVisibilityRuleAction = "elementVisibility"
)

type ZarazHistoryUpdateResponseTriggersLoadRulesZarazElementVisibilityRuleSettings struct {
	Selector string                                                                            `json:"selector,required"`
	JSON     zarazHistoryUpdateResponseTriggersLoadRulesZarazElementVisibilityRuleSettingsJSON `json:"-"`
}

// zarazHistoryUpdateResponseTriggersLoadRulesZarazElementVisibilityRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazHistoryUpdateResponseTriggersLoadRulesZarazElementVisibilityRuleSettings]
type zarazHistoryUpdateResponseTriggersLoadRulesZarazElementVisibilityRuleSettingsJSON struct {
	Selector    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseTriggersLoadRulesZarazElementVisibilityRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryUpdateResponseTriggersSystem string

const (
	ZarazHistoryUpdateResponseTriggersSystemPageload ZarazHistoryUpdateResponseTriggersSystem = "pageload"
)

// Union satisfied by [ZarazHistoryUpdateResponseVariablesObject] or
// [ZarazHistoryUpdateResponseVariablesObject].
type ZarazHistoryUpdateResponseVariable interface {
	implementsZarazHistoryUpdateResponseVariable()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZarazHistoryUpdateResponseVariable)(nil)).Elem(), "")
}

type ZarazHistoryUpdateResponseVariablesObject struct {
	Name  string                                        `json:"name,required"`
	Type  ZarazHistoryUpdateResponseVariablesObjectType `json:"type,required"`
	Value string                                        `json:"value,required"`
	JSON  zarazHistoryUpdateResponseVariablesObjectJSON `json:"-"`
}

// zarazHistoryUpdateResponseVariablesObjectJSON contains the JSON metadata for the
// struct [ZarazHistoryUpdateResponseVariablesObject]
type zarazHistoryUpdateResponseVariablesObjectJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseVariablesObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazHistoryUpdateResponseVariablesObject) implementsZarazHistoryUpdateResponseVariable() {}

type ZarazHistoryUpdateResponseVariablesObjectType string

const (
	ZarazHistoryUpdateResponseVariablesObjectTypeString ZarazHistoryUpdateResponseVariablesObjectType = "string"
	ZarazHistoryUpdateResponseVariablesObjectTypeSecret ZarazHistoryUpdateResponseVariablesObjectType = "secret"
)

// Consent management configuration.
type ZarazHistoryUpdateResponseConsent struct {
	Enabled                bool                                                    `json:"enabled,required"`
	ButtonTextTranslations ZarazHistoryUpdateResponseConsentButtonTextTranslations `json:"buttonTextTranslations"`
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
	Purposes map[string]ZarazHistoryUpdateResponseConsentPurpose `json:"purposes"`
	// Object where keys are purpose alpha-numeric IDs
	PurposesWithTranslations map[string]ZarazHistoryUpdateResponseConsentPurposesWithTranslation `json:"purposesWithTranslations"`
	JSON                     zarazHistoryUpdateResponseConsentJSON                               `json:"-"`
}

// zarazHistoryUpdateResponseConsentJSON contains the JSON metadata for the struct
// [ZarazHistoryUpdateResponseConsent]
type zarazHistoryUpdateResponseConsentJSON struct {
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

func (r *ZarazHistoryUpdateResponseConsent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryUpdateResponseConsentButtonTextTranslations struct {
	// Object where keys are language codes
	AcceptAll map[string]string `json:"accept_all,required"`
	// Object where keys are language codes
	ConfirmMyChoices map[string]string `json:"confirm_my_choices,required"`
	// Object where keys are language codes
	RejectAll map[string]string                                           `json:"reject_all,required"`
	JSON      zarazHistoryUpdateResponseConsentButtonTextTranslationsJSON `json:"-"`
}

// zarazHistoryUpdateResponseConsentButtonTextTranslationsJSON contains the JSON
// metadata for the struct
// [ZarazHistoryUpdateResponseConsentButtonTextTranslations]
type zarazHistoryUpdateResponseConsentButtonTextTranslationsJSON struct {
	AcceptAll        apijson.Field
	ConfirmMyChoices apijson.Field
	RejectAll        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseConsentButtonTextTranslations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryUpdateResponseConsentPurpose struct {
	Description string                                       `json:"description,required"`
	Name        string                                       `json:"name,required"`
	JSON        zarazHistoryUpdateResponseConsentPurposeJSON `json:"-"`
}

// zarazHistoryUpdateResponseConsentPurposeJSON contains the JSON metadata for the
// struct [ZarazHistoryUpdateResponseConsentPurpose]
type zarazHistoryUpdateResponseConsentPurposeJSON struct {
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseConsentPurpose) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryUpdateResponseConsentPurposesWithTranslation struct {
	// Object where keys are language codes
	Description map[string]string `json:"description,required"`
	// Object where keys are language codes
	Name  map[string]string                                            `json:"name,required"`
	Order int64                                                        `json:"order,required"`
	JSON  zarazHistoryUpdateResponseConsentPurposesWithTranslationJSON `json:"-"`
}

// zarazHistoryUpdateResponseConsentPurposesWithTranslationJSON contains the JSON
// metadata for the struct
// [ZarazHistoryUpdateResponseConsentPurposesWithTranslation]
type zarazHistoryUpdateResponseConsentPurposesWithTranslationJSON struct {
	Description apijson.Field
	Name        apijson.Field
	Order       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseConsentPurposesWithTranslation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryListResponse struct {
	// ID of the configuration
	ID int64 `json:"id,required"`
	// Date and time the configuration was created
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Configuration description provided by the user who published this configuration
	Description string `json:"description,required"`
	// Date and time the configuration was last updated
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
	// Alpha-numeric ID of the account user who published the configuration
	UserID string                       `json:"userId,required"`
	JSON   zarazHistoryListResponseJSON `json:"-"`
}

// zarazHistoryListResponseJSON contains the JSON metadata for the struct
// [ZarazHistoryListResponse]
type zarazHistoryListResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	UpdatedAt   apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryUpdateParams struct {
	// ID of the Zaraz configuration to restore.
	Body param.Field[int64] `json:"body,required"`
}

func (r ZarazHistoryUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZarazHistoryUpdateResponseEnvelope struct {
	Errors   []ZarazHistoryUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZarazHistoryUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Zaraz configuration
	Result ZarazHistoryUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success bool                                   `json:"success,required"`
	JSON    zarazHistoryUpdateResponseEnvelopeJSON `json:"-"`
}

// zarazHistoryUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZarazHistoryUpdateResponseEnvelope]
type zarazHistoryUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryUpdateResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    zarazHistoryUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// zarazHistoryUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZarazHistoryUpdateResponseEnvelopeErrors]
type zarazHistoryUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryUpdateResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zarazHistoryUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// zarazHistoryUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZarazHistoryUpdateResponseEnvelopeMessages]
type zarazHistoryUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryListParams struct {
	// Maximum amount of results to list. Default value is 10.
	Limit param.Field[int64] `query:"limit"`
	// Ordinal number to start listing the results with. Default value is 0.
	Offset param.Field[int64] `query:"offset"`
	// The field to sort by. Default is updated_at.
	SortField param.Field[ZarazHistoryListParamsSortField] `query:"sortField"`
	// Sorting order. Default is DESC.
	SortOrder param.Field[ZarazHistoryListParamsSortOrder] `query:"sortOrder"`
}

// URLQuery serializes [ZarazHistoryListParams]'s query parameters as `url.Values`.
func (r ZarazHistoryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The field to sort by. Default is updated_at.
type ZarazHistoryListParamsSortField string

const (
	ZarazHistoryListParamsSortFieldID          ZarazHistoryListParamsSortField = "id"
	ZarazHistoryListParamsSortFieldUserID      ZarazHistoryListParamsSortField = "user_id"
	ZarazHistoryListParamsSortFieldDescription ZarazHistoryListParamsSortField = "description"
	ZarazHistoryListParamsSortFieldCreatedAt   ZarazHistoryListParamsSortField = "created_at"
	ZarazHistoryListParamsSortFieldUpdatedAt   ZarazHistoryListParamsSortField = "updated_at"
)

// Sorting order. Default is DESC.
type ZarazHistoryListParamsSortOrder string

const (
	ZarazHistoryListParamsSortOrderDesc ZarazHistoryListParamsSortOrder = "DESC"
	ZarazHistoryListParamsSortOrderAsc  ZarazHistoryListParamsSortOrder = "ASC"
)

type ZarazHistoryListResponseEnvelope struct {
	Errors   []ZarazHistoryListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZarazHistoryListResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZarazHistoryListResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success bool                                 `json:"success,required"`
	JSON    zarazHistoryListResponseEnvelopeJSON `json:"-"`
}

// zarazHistoryListResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZarazHistoryListResponseEnvelope]
type zarazHistoryListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryListResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    zarazHistoryListResponseEnvelopeErrorsJSON `json:"-"`
}

// zarazHistoryListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZarazHistoryListResponseEnvelopeErrors]
type zarazHistoryListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryListResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    zarazHistoryListResponseEnvelopeMessagesJSON `json:"-"`
}

// zarazHistoryListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ZarazHistoryListResponseEnvelopeMessages]
type zarazHistoryListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
