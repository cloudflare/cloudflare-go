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

// ZarazExportService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZarazExportService] method
// instead.
type ZarazExportService struct {
	Options []option.RequestOption
}

// NewZarazExportService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZarazExportService(opts ...option.RequestOption) (r *ZarazExportService) {
	r = &ZarazExportService{}
	r.Options = opts
	return
}

// Exports full current published Zaraz configuration for a zone, secret variables
// included.
func (r *ZarazExportService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZarazExportGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/zaraz/export", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Zaraz configuration
type ZarazExportGetResponse struct {
	// Data layer compatibility mode enabled.
	DataLayer bool `json:"dataLayer,required"`
	// The key for Zaraz debug mode.
	DebugKey string `json:"debugKey,required"`
	// General Zaraz settings.
	Settings ZarazExportGetResponseSettings `json:"settings,required"`
	// Tools set up under Zaraz configuration, where key is the alpha-numeric tool ID
	// and value is the tool configuration object.
	Tools map[string]ZarazExportGetResponseTool `json:"tools,required"`
	// Triggers set up under Zaraz configuration, where key is the trigger
	// alpha-numeric ID and value is the trigger configuration.
	Triggers map[string]ZarazExportGetResponseTrigger `json:"triggers,required"`
	// Variables set up under Zaraz configuration, where key is the variable
	// alpha-numeric ID and value is the variable configuration. Values of variables of
	// type secret are not included.
	Variables map[string]ZarazExportGetResponseVariable `json:"variables,required"`
	// Zaraz internal version of the config.
	ZarazVersion int64 `json:"zarazVersion,required"`
	// Consent management configuration.
	Consent ZarazExportGetResponseConsent `json:"consent"`
	// Single Page Application support enabled.
	HistoryChange bool                       `json:"historyChange"`
	JSON          zarazExportGetResponseJSON `json:"-"`
}

// zarazExportGetResponseJSON contains the JSON metadata for the struct
// [ZarazExportGetResponse]
type zarazExportGetResponseJSON struct {
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

func (r *ZarazExportGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// General Zaraz settings.
type ZarazExportGetResponseSettings struct {
	// Automatic injection of Zaraz scripts enabled.
	AutoInjectScript bool `json:"autoInjectScript,required"`
	// Details of the worker that receives and edits Zaraz Context object.
	ContextEnricher ZarazExportGetResponseSettingsContextEnricher `json:"contextEnricher"`
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
	JSON      zarazExportGetResponseSettingsJSON `json:"-"`
}

// zarazExportGetResponseSettingsJSON contains the JSON metadata for the struct
// [ZarazExportGetResponseSettings]
type zarazExportGetResponseSettingsJSON struct {
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

func (r *ZarazExportGetResponseSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details of the worker that receives and edits Zaraz Context object.
type ZarazExportGetResponseSettingsContextEnricher struct {
	EscapedWorkerName string                                            `json:"escapedWorkerName,required"`
	WorkerTag         string                                            `json:"workerTag,required"`
	JSON              zarazExportGetResponseSettingsContextEnricherJSON `json:"-"`
}

// zarazExportGetResponseSettingsContextEnricherJSON contains the JSON metadata for
// the struct [ZarazExportGetResponseSettingsContextEnricher]
type zarazExportGetResponseSettingsContextEnricherJSON struct {
	EscapedWorkerName apijson.Field
	WorkerTag         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZarazExportGetResponseSettingsContextEnricher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [ZarazExportGetResponseToolsZarazManagedComponent] or
// [ZarazExportGetResponseToolsZarazCustomManagedComponent].
type ZarazExportGetResponseTool interface {
	implementsZarazExportGetResponseTool()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZarazExportGetResponseTool)(nil)).Elem(), "")
}

type ZarazExportGetResponseToolsZarazManagedComponent struct {
	// List of blocking trigger IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Tool's internal name
	Component string `json:"component,required"`
	// Default fields for tool's actions
	DefaultFields map[string]ZarazExportGetResponseToolsZarazManagedComponentDefaultField `json:"defaultFields,required"`
	// Whether tool is enabled
	Enabled bool `json:"enabled,required"`
	// Tool's name defined by the user
	Name string `json:"name,required"`
	// List of permissions granted to the component
	Permissions []string `json:"permissions,required"`
	// Tool's settings
	Settings map[string]ZarazExportGetResponseToolsZarazManagedComponentSetting `json:"settings,required"`
	Type     ZarazExportGetResponseToolsZarazManagedComponentType               `json:"type,required"`
	// Actions configured on a tool. Either this or neoEvents field is required.
	Actions map[string]ZarazExportGetResponseToolsZarazManagedComponentAction `json:"actions"`
	// Default consent purpose ID
	DefaultPurpose string `json:"defaultPurpose"`
	// DEPRECATED - List of actions configured on a tool. Either this or actions field
	// is required. If both are present, actions field will take precedence.
	NeoEvents []ZarazExportGetResponseToolsZarazManagedComponentNeoEvent `json:"neoEvents"`
	JSON      zarazExportGetResponseToolsZarazManagedComponentJSON       `json:"-"`
}

// zarazExportGetResponseToolsZarazManagedComponentJSON contains the JSON metadata
// for the struct [ZarazExportGetResponseToolsZarazManagedComponent]
type zarazExportGetResponseToolsZarazManagedComponentJSON struct {
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

func (r *ZarazExportGetResponseToolsZarazManagedComponent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazExportGetResponseToolsZarazManagedComponent) implementsZarazExportGetResponseTool() {}

// Union satisfied by [shared.UnionString] or [shared.UnionBool].
type ZarazExportGetResponseToolsZarazManagedComponentDefaultField interface {
	ImplementsZarazExportGetResponseToolsZarazManagedComponentDefaultField()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazExportGetResponseToolsZarazManagedComponentDefaultField)(nil)).Elem(),
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
type ZarazExportGetResponseToolsZarazManagedComponentSetting interface {
	ImplementsZarazExportGetResponseToolsZarazManagedComponentSetting()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazExportGetResponseToolsZarazManagedComponentSetting)(nil)).Elem(),
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

type ZarazExportGetResponseToolsZarazManagedComponentType string

const (
	ZarazExportGetResponseToolsZarazManagedComponentTypeComponent ZarazExportGetResponseToolsZarazManagedComponentType = "component"
)

type ZarazExportGetResponseToolsZarazManagedComponentAction struct {
	// Tool event type
	ActionType string `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Event payload
	Data interface{} `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers []string                                                   `json:"firingTriggers,required"`
	JSON           zarazExportGetResponseToolsZarazManagedComponentActionJSON `json:"-"`
}

// zarazExportGetResponseToolsZarazManagedComponentActionJSON contains the JSON
// metadata for the struct [ZarazExportGetResponseToolsZarazManagedComponentAction]
type zarazExportGetResponseToolsZarazManagedComponentActionJSON struct {
	ActionType       apijson.Field
	BlockingTriggers apijson.Field
	Data             apijson.Field
	FiringTriggers   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazExportGetResponseToolsZarazManagedComponentAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazExportGetResponseToolsZarazManagedComponentNeoEvent struct {
	// Tool event type
	ActionType string `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Event payload
	Data interface{} `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers []string                                                     `json:"firingTriggers,required"`
	JSON           zarazExportGetResponseToolsZarazManagedComponentNeoEventJSON `json:"-"`
}

// zarazExportGetResponseToolsZarazManagedComponentNeoEventJSON contains the JSON
// metadata for the struct
// [ZarazExportGetResponseToolsZarazManagedComponentNeoEvent]
type zarazExportGetResponseToolsZarazManagedComponentNeoEventJSON struct {
	ActionType       apijson.Field
	BlockingTriggers apijson.Field
	Data             apijson.Field
	FiringTriggers   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazExportGetResponseToolsZarazManagedComponentNeoEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazExportGetResponseToolsZarazCustomManagedComponent struct {
	// List of blocking trigger IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Tool's internal name
	Component string `json:"component,required"`
	// Default fields for tool's actions
	DefaultFields map[string]ZarazExportGetResponseToolsZarazCustomManagedComponentDefaultField `json:"defaultFields,required"`
	// Whether tool is enabled
	Enabled bool `json:"enabled,required"`
	// Tool's name defined by the user
	Name string `json:"name,required"`
	// List of permissions granted to the component
	Permissions []string `json:"permissions,required"`
	// Tool's settings
	Settings map[string]ZarazExportGetResponseToolsZarazCustomManagedComponentSetting `json:"settings,required"`
	Type     ZarazExportGetResponseToolsZarazCustomManagedComponentType               `json:"type,required"`
	// Cloudflare worker that acts as a managed component
	Worker ZarazExportGetResponseToolsZarazCustomManagedComponentWorker `json:"worker,required"`
	// Actions configured on a tool. Either this or neoEvents field is required.
	Actions map[string]ZarazExportGetResponseToolsZarazCustomManagedComponentAction `json:"actions"`
	// Default consent purpose ID
	DefaultPurpose string `json:"defaultPurpose"`
	// DEPRECATED - List of actions configured on a tool. Either this or actions field
	// is required. If both are present, actions field will take precedence.
	NeoEvents []ZarazExportGetResponseToolsZarazCustomManagedComponentNeoEvent `json:"neoEvents"`
	JSON      zarazExportGetResponseToolsZarazCustomManagedComponentJSON       `json:"-"`
}

// zarazExportGetResponseToolsZarazCustomManagedComponentJSON contains the JSON
// metadata for the struct [ZarazExportGetResponseToolsZarazCustomManagedComponent]
type zarazExportGetResponseToolsZarazCustomManagedComponentJSON struct {
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

func (r *ZarazExportGetResponseToolsZarazCustomManagedComponent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazExportGetResponseToolsZarazCustomManagedComponent) implementsZarazExportGetResponseTool() {
}

// Union satisfied by [shared.UnionString] or [shared.UnionBool].
type ZarazExportGetResponseToolsZarazCustomManagedComponentDefaultField interface {
	ImplementsZarazExportGetResponseToolsZarazCustomManagedComponentDefaultField()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazExportGetResponseToolsZarazCustomManagedComponentDefaultField)(nil)).Elem(),
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
type ZarazExportGetResponseToolsZarazCustomManagedComponentSetting interface {
	ImplementsZarazExportGetResponseToolsZarazCustomManagedComponentSetting()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazExportGetResponseToolsZarazCustomManagedComponentSetting)(nil)).Elem(),
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

type ZarazExportGetResponseToolsZarazCustomManagedComponentType string

const (
	ZarazExportGetResponseToolsZarazCustomManagedComponentTypeCustomMc ZarazExportGetResponseToolsZarazCustomManagedComponentType = "custom-mc"
)

// Cloudflare worker that acts as a managed component
type ZarazExportGetResponseToolsZarazCustomManagedComponentWorker struct {
	EscapedWorkerName string                                                           `json:"escapedWorkerName,required"`
	WorkerTag         string                                                           `json:"workerTag,required"`
	JSON              zarazExportGetResponseToolsZarazCustomManagedComponentWorkerJSON `json:"-"`
}

// zarazExportGetResponseToolsZarazCustomManagedComponentWorkerJSON contains the
// JSON metadata for the struct
// [ZarazExportGetResponseToolsZarazCustomManagedComponentWorker]
type zarazExportGetResponseToolsZarazCustomManagedComponentWorkerJSON struct {
	EscapedWorkerName apijson.Field
	WorkerTag         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZarazExportGetResponseToolsZarazCustomManagedComponentWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazExportGetResponseToolsZarazCustomManagedComponentAction struct {
	// Tool event type
	ActionType string `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Event payload
	Data interface{} `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers []string                                                         `json:"firingTriggers,required"`
	JSON           zarazExportGetResponseToolsZarazCustomManagedComponentActionJSON `json:"-"`
}

// zarazExportGetResponseToolsZarazCustomManagedComponentActionJSON contains the
// JSON metadata for the struct
// [ZarazExportGetResponseToolsZarazCustomManagedComponentAction]
type zarazExportGetResponseToolsZarazCustomManagedComponentActionJSON struct {
	ActionType       apijson.Field
	BlockingTriggers apijson.Field
	Data             apijson.Field
	FiringTriggers   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazExportGetResponseToolsZarazCustomManagedComponentAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazExportGetResponseToolsZarazCustomManagedComponentNeoEvent struct {
	// Tool event type
	ActionType string `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Event payload
	Data interface{} `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers []string                                                           `json:"firingTriggers,required"`
	JSON           zarazExportGetResponseToolsZarazCustomManagedComponentNeoEventJSON `json:"-"`
}

// zarazExportGetResponseToolsZarazCustomManagedComponentNeoEventJSON contains the
// JSON metadata for the struct
// [ZarazExportGetResponseToolsZarazCustomManagedComponentNeoEvent]
type zarazExportGetResponseToolsZarazCustomManagedComponentNeoEventJSON struct {
	ActionType       apijson.Field
	BlockingTriggers apijson.Field
	Data             apijson.Field
	FiringTriggers   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazExportGetResponseToolsZarazCustomManagedComponentNeoEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazExportGetResponseTrigger struct {
	// Rules defining when the trigger is not fired.
	ExcludeRules []ZarazExportGetResponseTriggersExcludeRule `json:"excludeRules,required"`
	// Rules defining when the trigger is fired.
	LoadRules []ZarazExportGetResponseTriggersLoadRule `json:"loadRules,required"`
	// Trigger name.
	Name string `json:"name,required"`
	// Trigger description.
	Description string                               `json:"description"`
	System      ZarazExportGetResponseTriggersSystem `json:"system"`
	JSON        zarazExportGetResponseTriggerJSON    `json:"-"`
}

// zarazExportGetResponseTriggerJSON contains the JSON metadata for the struct
// [ZarazExportGetResponseTrigger]
type zarazExportGetResponseTriggerJSON struct {
	ExcludeRules apijson.Field
	LoadRules    apijson.Field
	Name         apijson.Field
	Description  apijson.Field
	System       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZarazExportGetResponseTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [ZarazExportGetResponseTriggersExcludeRulesZarazLoadRule],
// [ZarazExportGetResponseTriggersExcludeRulesZarazClickListenerRule],
// [ZarazExportGetResponseTriggersExcludeRulesZarazTimerRule],
// [ZarazExportGetResponseTriggersExcludeRulesZarazFormSubmissionRule],
// [ZarazExportGetResponseTriggersExcludeRulesZarazVariableMatchRule],
// [ZarazExportGetResponseTriggersExcludeRulesZarazScrollDepthRule] or
// [ZarazExportGetResponseTriggersExcludeRulesZarazElementVisibilityRule].
type ZarazExportGetResponseTriggersExcludeRule interface {
	implementsZarazExportGetResponseTriggersExcludeRule()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZarazExportGetResponseTriggersExcludeRule)(nil)).Elem(), "")
}

type ZarazExportGetResponseTriggersExcludeRulesZarazLoadRule struct {
	ID    string                                                      `json:"id,required"`
	Match string                                                      `json:"match,required"`
	Op    ZarazExportGetResponseTriggersExcludeRulesZarazLoadRuleOp   `json:"op,required"`
	Value string                                                      `json:"value,required"`
	JSON  zarazExportGetResponseTriggersExcludeRulesZarazLoadRuleJSON `json:"-"`
}

// zarazExportGetResponseTriggersExcludeRulesZarazLoadRuleJSON contains the JSON
// metadata for the struct
// [ZarazExportGetResponseTriggersExcludeRulesZarazLoadRule]
type zarazExportGetResponseTriggersExcludeRulesZarazLoadRuleJSON struct {
	ID          apijson.Field
	Match       apijson.Field
	Op          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazExportGetResponseTriggersExcludeRulesZarazLoadRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazExportGetResponseTriggersExcludeRulesZarazLoadRule) implementsZarazExportGetResponseTriggersExcludeRule() {
}

type ZarazExportGetResponseTriggersExcludeRulesZarazLoadRuleOp string

const (
	ZarazExportGetResponseTriggersExcludeRulesZarazLoadRuleOpContains           ZarazExportGetResponseTriggersExcludeRulesZarazLoadRuleOp = "CONTAINS"
	ZarazExportGetResponseTriggersExcludeRulesZarazLoadRuleOpEquals             ZarazExportGetResponseTriggersExcludeRulesZarazLoadRuleOp = "EQUALS"
	ZarazExportGetResponseTriggersExcludeRulesZarazLoadRuleOpStartsWith         ZarazExportGetResponseTriggersExcludeRulesZarazLoadRuleOp = "STARTS_WITH"
	ZarazExportGetResponseTriggersExcludeRulesZarazLoadRuleOpEndsWith           ZarazExportGetResponseTriggersExcludeRulesZarazLoadRuleOp = "ENDS_WITH"
	ZarazExportGetResponseTriggersExcludeRulesZarazLoadRuleOpMatchRegex         ZarazExportGetResponseTriggersExcludeRulesZarazLoadRuleOp = "MATCH_REGEX"
	ZarazExportGetResponseTriggersExcludeRulesZarazLoadRuleOpNotMatchRegex      ZarazExportGetResponseTriggersExcludeRulesZarazLoadRuleOp = "NOT_MATCH_REGEX"
	ZarazExportGetResponseTriggersExcludeRulesZarazLoadRuleOpGreaterThan        ZarazExportGetResponseTriggersExcludeRulesZarazLoadRuleOp = "GREATER_THAN"
	ZarazExportGetResponseTriggersExcludeRulesZarazLoadRuleOpGreaterThanOrEqual ZarazExportGetResponseTriggersExcludeRulesZarazLoadRuleOp = "GREATER_THAN_OR_EQUAL"
	ZarazExportGetResponseTriggersExcludeRulesZarazLoadRuleOpLessThan           ZarazExportGetResponseTriggersExcludeRulesZarazLoadRuleOp = "LESS_THAN"
	ZarazExportGetResponseTriggersExcludeRulesZarazLoadRuleOpLessThanOrEqual    ZarazExportGetResponseTriggersExcludeRulesZarazLoadRuleOp = "LESS_THAN_OR_EQUAL"
)

type ZarazExportGetResponseTriggersExcludeRulesZarazClickListenerRule struct {
	ID       string                                                                   `json:"id,required"`
	Action   ZarazExportGetResponseTriggersExcludeRulesZarazClickListenerRuleAction   `json:"action,required"`
	Settings ZarazExportGetResponseTriggersExcludeRulesZarazClickListenerRuleSettings `json:"settings,required"`
	JSON     zarazExportGetResponseTriggersExcludeRulesZarazClickListenerRuleJSON     `json:"-"`
}

// zarazExportGetResponseTriggersExcludeRulesZarazClickListenerRuleJSON contains
// the JSON metadata for the struct
// [ZarazExportGetResponseTriggersExcludeRulesZarazClickListenerRule]
type zarazExportGetResponseTriggersExcludeRulesZarazClickListenerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazExportGetResponseTriggersExcludeRulesZarazClickListenerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazExportGetResponseTriggersExcludeRulesZarazClickListenerRule) implementsZarazExportGetResponseTriggersExcludeRule() {
}

type ZarazExportGetResponseTriggersExcludeRulesZarazClickListenerRuleAction string

const (
	ZarazExportGetResponseTriggersExcludeRulesZarazClickListenerRuleActionClickListener ZarazExportGetResponseTriggersExcludeRulesZarazClickListenerRuleAction = "clickListener"
)

type ZarazExportGetResponseTriggersExcludeRulesZarazClickListenerRuleSettings struct {
	Selector    string                                                                       `json:"selector,required"`
	Type        ZarazExportGetResponseTriggersExcludeRulesZarazClickListenerRuleSettingsType `json:"type,required"`
	WaitForTags int64                                                                        `json:"waitForTags,required"`
	JSON        zarazExportGetResponseTriggersExcludeRulesZarazClickListenerRuleSettingsJSON `json:"-"`
}

// zarazExportGetResponseTriggersExcludeRulesZarazClickListenerRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazExportGetResponseTriggersExcludeRulesZarazClickListenerRuleSettings]
type zarazExportGetResponseTriggersExcludeRulesZarazClickListenerRuleSettingsJSON struct {
	Selector    apijson.Field
	Type        apijson.Field
	WaitForTags apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazExportGetResponseTriggersExcludeRulesZarazClickListenerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazExportGetResponseTriggersExcludeRulesZarazClickListenerRuleSettingsType string

const (
	ZarazExportGetResponseTriggersExcludeRulesZarazClickListenerRuleSettingsTypeXpath ZarazExportGetResponseTriggersExcludeRulesZarazClickListenerRuleSettingsType = "xpath"
	ZarazExportGetResponseTriggersExcludeRulesZarazClickListenerRuleSettingsTypeCss   ZarazExportGetResponseTriggersExcludeRulesZarazClickListenerRuleSettingsType = "css"
)

type ZarazExportGetResponseTriggersExcludeRulesZarazTimerRule struct {
	ID       string                                                           `json:"id,required"`
	Action   ZarazExportGetResponseTriggersExcludeRulesZarazTimerRuleAction   `json:"action,required"`
	Settings ZarazExportGetResponseTriggersExcludeRulesZarazTimerRuleSettings `json:"settings,required"`
	JSON     zarazExportGetResponseTriggersExcludeRulesZarazTimerRuleJSON     `json:"-"`
}

// zarazExportGetResponseTriggersExcludeRulesZarazTimerRuleJSON contains the JSON
// metadata for the struct
// [ZarazExportGetResponseTriggersExcludeRulesZarazTimerRule]
type zarazExportGetResponseTriggersExcludeRulesZarazTimerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazExportGetResponseTriggersExcludeRulesZarazTimerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazExportGetResponseTriggersExcludeRulesZarazTimerRule) implementsZarazExportGetResponseTriggersExcludeRule() {
}

type ZarazExportGetResponseTriggersExcludeRulesZarazTimerRuleAction string

const (
	ZarazExportGetResponseTriggersExcludeRulesZarazTimerRuleActionTimer ZarazExportGetResponseTriggersExcludeRulesZarazTimerRuleAction = "timer"
)

type ZarazExportGetResponseTriggersExcludeRulesZarazTimerRuleSettings struct {
	Interval int64                                                                `json:"interval,required"`
	Limit    int64                                                                `json:"limit,required"`
	JSON     zarazExportGetResponseTriggersExcludeRulesZarazTimerRuleSettingsJSON `json:"-"`
}

// zarazExportGetResponseTriggersExcludeRulesZarazTimerRuleSettingsJSON contains
// the JSON metadata for the struct
// [ZarazExportGetResponseTriggersExcludeRulesZarazTimerRuleSettings]
type zarazExportGetResponseTriggersExcludeRulesZarazTimerRuleSettingsJSON struct {
	Interval    apijson.Field
	Limit       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazExportGetResponseTriggersExcludeRulesZarazTimerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazExportGetResponseTriggersExcludeRulesZarazFormSubmissionRule struct {
	ID       string                                                                    `json:"id,required"`
	Action   ZarazExportGetResponseTriggersExcludeRulesZarazFormSubmissionRuleAction   `json:"action,required"`
	Settings ZarazExportGetResponseTriggersExcludeRulesZarazFormSubmissionRuleSettings `json:"settings,required"`
	JSON     zarazExportGetResponseTriggersExcludeRulesZarazFormSubmissionRuleJSON     `json:"-"`
}

// zarazExportGetResponseTriggersExcludeRulesZarazFormSubmissionRuleJSON contains
// the JSON metadata for the struct
// [ZarazExportGetResponseTriggersExcludeRulesZarazFormSubmissionRule]
type zarazExportGetResponseTriggersExcludeRulesZarazFormSubmissionRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazExportGetResponseTriggersExcludeRulesZarazFormSubmissionRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazExportGetResponseTriggersExcludeRulesZarazFormSubmissionRule) implementsZarazExportGetResponseTriggersExcludeRule() {
}

type ZarazExportGetResponseTriggersExcludeRulesZarazFormSubmissionRuleAction string

const (
	ZarazExportGetResponseTriggersExcludeRulesZarazFormSubmissionRuleActionFormSubmission ZarazExportGetResponseTriggersExcludeRulesZarazFormSubmissionRuleAction = "formSubmission"
)

type ZarazExportGetResponseTriggersExcludeRulesZarazFormSubmissionRuleSettings struct {
	Selector string                                                                        `json:"selector,required"`
	Validate bool                                                                          `json:"validate,required"`
	JSON     zarazExportGetResponseTriggersExcludeRulesZarazFormSubmissionRuleSettingsJSON `json:"-"`
}

// zarazExportGetResponseTriggersExcludeRulesZarazFormSubmissionRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazExportGetResponseTriggersExcludeRulesZarazFormSubmissionRuleSettings]
type zarazExportGetResponseTriggersExcludeRulesZarazFormSubmissionRuleSettingsJSON struct {
	Selector    apijson.Field
	Validate    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazExportGetResponseTriggersExcludeRulesZarazFormSubmissionRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazExportGetResponseTriggersExcludeRulesZarazVariableMatchRule struct {
	ID       string                                                                   `json:"id,required"`
	Action   ZarazExportGetResponseTriggersExcludeRulesZarazVariableMatchRuleAction   `json:"action,required"`
	Settings ZarazExportGetResponseTriggersExcludeRulesZarazVariableMatchRuleSettings `json:"settings,required"`
	JSON     zarazExportGetResponseTriggersExcludeRulesZarazVariableMatchRuleJSON     `json:"-"`
}

// zarazExportGetResponseTriggersExcludeRulesZarazVariableMatchRuleJSON contains
// the JSON metadata for the struct
// [ZarazExportGetResponseTriggersExcludeRulesZarazVariableMatchRule]
type zarazExportGetResponseTriggersExcludeRulesZarazVariableMatchRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazExportGetResponseTriggersExcludeRulesZarazVariableMatchRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazExportGetResponseTriggersExcludeRulesZarazVariableMatchRule) implementsZarazExportGetResponseTriggersExcludeRule() {
}

type ZarazExportGetResponseTriggersExcludeRulesZarazVariableMatchRuleAction string

const (
	ZarazExportGetResponseTriggersExcludeRulesZarazVariableMatchRuleActionVariableMatch ZarazExportGetResponseTriggersExcludeRulesZarazVariableMatchRuleAction = "variableMatch"
)

type ZarazExportGetResponseTriggersExcludeRulesZarazVariableMatchRuleSettings struct {
	Match    string                                                                       `json:"match,required"`
	Variable string                                                                       `json:"variable,required"`
	JSON     zarazExportGetResponseTriggersExcludeRulesZarazVariableMatchRuleSettingsJSON `json:"-"`
}

// zarazExportGetResponseTriggersExcludeRulesZarazVariableMatchRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazExportGetResponseTriggersExcludeRulesZarazVariableMatchRuleSettings]
type zarazExportGetResponseTriggersExcludeRulesZarazVariableMatchRuleSettingsJSON struct {
	Match       apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazExportGetResponseTriggersExcludeRulesZarazVariableMatchRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazExportGetResponseTriggersExcludeRulesZarazScrollDepthRule struct {
	ID       string                                                                 `json:"id,required"`
	Action   ZarazExportGetResponseTriggersExcludeRulesZarazScrollDepthRuleAction   `json:"action,required"`
	Settings ZarazExportGetResponseTriggersExcludeRulesZarazScrollDepthRuleSettings `json:"settings,required"`
	JSON     zarazExportGetResponseTriggersExcludeRulesZarazScrollDepthRuleJSON     `json:"-"`
}

// zarazExportGetResponseTriggersExcludeRulesZarazScrollDepthRuleJSON contains the
// JSON metadata for the struct
// [ZarazExportGetResponseTriggersExcludeRulesZarazScrollDepthRule]
type zarazExportGetResponseTriggersExcludeRulesZarazScrollDepthRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazExportGetResponseTriggersExcludeRulesZarazScrollDepthRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazExportGetResponseTriggersExcludeRulesZarazScrollDepthRule) implementsZarazExportGetResponseTriggersExcludeRule() {
}

type ZarazExportGetResponseTriggersExcludeRulesZarazScrollDepthRuleAction string

const (
	ZarazExportGetResponseTriggersExcludeRulesZarazScrollDepthRuleActionScrollDepth ZarazExportGetResponseTriggersExcludeRulesZarazScrollDepthRuleAction = "scrollDepth"
)

type ZarazExportGetResponseTriggersExcludeRulesZarazScrollDepthRuleSettings struct {
	Positions string                                                                     `json:"positions,required"`
	JSON      zarazExportGetResponseTriggersExcludeRulesZarazScrollDepthRuleSettingsJSON `json:"-"`
}

// zarazExportGetResponseTriggersExcludeRulesZarazScrollDepthRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazExportGetResponseTriggersExcludeRulesZarazScrollDepthRuleSettings]
type zarazExportGetResponseTriggersExcludeRulesZarazScrollDepthRuleSettingsJSON struct {
	Positions   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazExportGetResponseTriggersExcludeRulesZarazScrollDepthRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazExportGetResponseTriggersExcludeRulesZarazElementVisibilityRule struct {
	ID       string                                                                       `json:"id,required"`
	Action   ZarazExportGetResponseTriggersExcludeRulesZarazElementVisibilityRuleAction   `json:"action,required"`
	Settings ZarazExportGetResponseTriggersExcludeRulesZarazElementVisibilityRuleSettings `json:"settings,required"`
	JSON     zarazExportGetResponseTriggersExcludeRulesZarazElementVisibilityRuleJSON     `json:"-"`
}

// zarazExportGetResponseTriggersExcludeRulesZarazElementVisibilityRuleJSON
// contains the JSON metadata for the struct
// [ZarazExportGetResponseTriggersExcludeRulesZarazElementVisibilityRule]
type zarazExportGetResponseTriggersExcludeRulesZarazElementVisibilityRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazExportGetResponseTriggersExcludeRulesZarazElementVisibilityRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazExportGetResponseTriggersExcludeRulesZarazElementVisibilityRule) implementsZarazExportGetResponseTriggersExcludeRule() {
}

type ZarazExportGetResponseTriggersExcludeRulesZarazElementVisibilityRuleAction string

const (
	ZarazExportGetResponseTriggersExcludeRulesZarazElementVisibilityRuleActionElementVisibility ZarazExportGetResponseTriggersExcludeRulesZarazElementVisibilityRuleAction = "elementVisibility"
)

type ZarazExportGetResponseTriggersExcludeRulesZarazElementVisibilityRuleSettings struct {
	Selector string                                                                           `json:"selector,required"`
	JSON     zarazExportGetResponseTriggersExcludeRulesZarazElementVisibilityRuleSettingsJSON `json:"-"`
}

// zarazExportGetResponseTriggersExcludeRulesZarazElementVisibilityRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazExportGetResponseTriggersExcludeRulesZarazElementVisibilityRuleSettings]
type zarazExportGetResponseTriggersExcludeRulesZarazElementVisibilityRuleSettingsJSON struct {
	Selector    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazExportGetResponseTriggersExcludeRulesZarazElementVisibilityRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [ZarazExportGetResponseTriggersLoadRulesZarazLoadRule],
// [ZarazExportGetResponseTriggersLoadRulesZarazClickListenerRule],
// [ZarazExportGetResponseTriggersLoadRulesZarazTimerRule],
// [ZarazExportGetResponseTriggersLoadRulesZarazFormSubmissionRule],
// [ZarazExportGetResponseTriggersLoadRulesZarazVariableMatchRule],
// [ZarazExportGetResponseTriggersLoadRulesZarazScrollDepthRule] or
// [ZarazExportGetResponseTriggersLoadRulesZarazElementVisibilityRule].
type ZarazExportGetResponseTriggersLoadRule interface {
	implementsZarazExportGetResponseTriggersLoadRule()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZarazExportGetResponseTriggersLoadRule)(nil)).Elem(), "")
}

type ZarazExportGetResponseTriggersLoadRulesZarazLoadRule struct {
	ID    string                                                   `json:"id,required"`
	Match string                                                   `json:"match,required"`
	Op    ZarazExportGetResponseTriggersLoadRulesZarazLoadRuleOp   `json:"op,required"`
	Value string                                                   `json:"value,required"`
	JSON  zarazExportGetResponseTriggersLoadRulesZarazLoadRuleJSON `json:"-"`
}

// zarazExportGetResponseTriggersLoadRulesZarazLoadRuleJSON contains the JSON
// metadata for the struct [ZarazExportGetResponseTriggersLoadRulesZarazLoadRule]
type zarazExportGetResponseTriggersLoadRulesZarazLoadRuleJSON struct {
	ID          apijson.Field
	Match       apijson.Field
	Op          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazExportGetResponseTriggersLoadRulesZarazLoadRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazExportGetResponseTriggersLoadRulesZarazLoadRule) implementsZarazExportGetResponseTriggersLoadRule() {
}

type ZarazExportGetResponseTriggersLoadRulesZarazLoadRuleOp string

const (
	ZarazExportGetResponseTriggersLoadRulesZarazLoadRuleOpContains           ZarazExportGetResponseTriggersLoadRulesZarazLoadRuleOp = "CONTAINS"
	ZarazExportGetResponseTriggersLoadRulesZarazLoadRuleOpEquals             ZarazExportGetResponseTriggersLoadRulesZarazLoadRuleOp = "EQUALS"
	ZarazExportGetResponseTriggersLoadRulesZarazLoadRuleOpStartsWith         ZarazExportGetResponseTriggersLoadRulesZarazLoadRuleOp = "STARTS_WITH"
	ZarazExportGetResponseTriggersLoadRulesZarazLoadRuleOpEndsWith           ZarazExportGetResponseTriggersLoadRulesZarazLoadRuleOp = "ENDS_WITH"
	ZarazExportGetResponseTriggersLoadRulesZarazLoadRuleOpMatchRegex         ZarazExportGetResponseTriggersLoadRulesZarazLoadRuleOp = "MATCH_REGEX"
	ZarazExportGetResponseTriggersLoadRulesZarazLoadRuleOpNotMatchRegex      ZarazExportGetResponseTriggersLoadRulesZarazLoadRuleOp = "NOT_MATCH_REGEX"
	ZarazExportGetResponseTriggersLoadRulesZarazLoadRuleOpGreaterThan        ZarazExportGetResponseTriggersLoadRulesZarazLoadRuleOp = "GREATER_THAN"
	ZarazExportGetResponseTriggersLoadRulesZarazLoadRuleOpGreaterThanOrEqual ZarazExportGetResponseTriggersLoadRulesZarazLoadRuleOp = "GREATER_THAN_OR_EQUAL"
	ZarazExportGetResponseTriggersLoadRulesZarazLoadRuleOpLessThan           ZarazExportGetResponseTriggersLoadRulesZarazLoadRuleOp = "LESS_THAN"
	ZarazExportGetResponseTriggersLoadRulesZarazLoadRuleOpLessThanOrEqual    ZarazExportGetResponseTriggersLoadRulesZarazLoadRuleOp = "LESS_THAN_OR_EQUAL"
)

type ZarazExportGetResponseTriggersLoadRulesZarazClickListenerRule struct {
	ID       string                                                                `json:"id,required"`
	Action   ZarazExportGetResponseTriggersLoadRulesZarazClickListenerRuleAction   `json:"action,required"`
	Settings ZarazExportGetResponseTriggersLoadRulesZarazClickListenerRuleSettings `json:"settings,required"`
	JSON     zarazExportGetResponseTriggersLoadRulesZarazClickListenerRuleJSON     `json:"-"`
}

// zarazExportGetResponseTriggersLoadRulesZarazClickListenerRuleJSON contains the
// JSON metadata for the struct
// [ZarazExportGetResponseTriggersLoadRulesZarazClickListenerRule]
type zarazExportGetResponseTriggersLoadRulesZarazClickListenerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazExportGetResponseTriggersLoadRulesZarazClickListenerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazExportGetResponseTriggersLoadRulesZarazClickListenerRule) implementsZarazExportGetResponseTriggersLoadRule() {
}

type ZarazExportGetResponseTriggersLoadRulesZarazClickListenerRuleAction string

const (
	ZarazExportGetResponseTriggersLoadRulesZarazClickListenerRuleActionClickListener ZarazExportGetResponseTriggersLoadRulesZarazClickListenerRuleAction = "clickListener"
)

type ZarazExportGetResponseTriggersLoadRulesZarazClickListenerRuleSettings struct {
	Selector    string                                                                    `json:"selector,required"`
	Type        ZarazExportGetResponseTriggersLoadRulesZarazClickListenerRuleSettingsType `json:"type,required"`
	WaitForTags int64                                                                     `json:"waitForTags,required"`
	JSON        zarazExportGetResponseTriggersLoadRulesZarazClickListenerRuleSettingsJSON `json:"-"`
}

// zarazExportGetResponseTriggersLoadRulesZarazClickListenerRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazExportGetResponseTriggersLoadRulesZarazClickListenerRuleSettings]
type zarazExportGetResponseTriggersLoadRulesZarazClickListenerRuleSettingsJSON struct {
	Selector    apijson.Field
	Type        apijson.Field
	WaitForTags apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazExportGetResponseTriggersLoadRulesZarazClickListenerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazExportGetResponseTriggersLoadRulesZarazClickListenerRuleSettingsType string

const (
	ZarazExportGetResponseTriggersLoadRulesZarazClickListenerRuleSettingsTypeXpath ZarazExportGetResponseTriggersLoadRulesZarazClickListenerRuleSettingsType = "xpath"
	ZarazExportGetResponseTriggersLoadRulesZarazClickListenerRuleSettingsTypeCss   ZarazExportGetResponseTriggersLoadRulesZarazClickListenerRuleSettingsType = "css"
)

type ZarazExportGetResponseTriggersLoadRulesZarazTimerRule struct {
	ID       string                                                        `json:"id,required"`
	Action   ZarazExportGetResponseTriggersLoadRulesZarazTimerRuleAction   `json:"action,required"`
	Settings ZarazExportGetResponseTriggersLoadRulesZarazTimerRuleSettings `json:"settings,required"`
	JSON     zarazExportGetResponseTriggersLoadRulesZarazTimerRuleJSON     `json:"-"`
}

// zarazExportGetResponseTriggersLoadRulesZarazTimerRuleJSON contains the JSON
// metadata for the struct [ZarazExportGetResponseTriggersLoadRulesZarazTimerRule]
type zarazExportGetResponseTriggersLoadRulesZarazTimerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazExportGetResponseTriggersLoadRulesZarazTimerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazExportGetResponseTriggersLoadRulesZarazTimerRule) implementsZarazExportGetResponseTriggersLoadRule() {
}

type ZarazExportGetResponseTriggersLoadRulesZarazTimerRuleAction string

const (
	ZarazExportGetResponseTriggersLoadRulesZarazTimerRuleActionTimer ZarazExportGetResponseTriggersLoadRulesZarazTimerRuleAction = "timer"
)

type ZarazExportGetResponseTriggersLoadRulesZarazTimerRuleSettings struct {
	Interval int64                                                             `json:"interval,required"`
	Limit    int64                                                             `json:"limit,required"`
	JSON     zarazExportGetResponseTriggersLoadRulesZarazTimerRuleSettingsJSON `json:"-"`
}

// zarazExportGetResponseTriggersLoadRulesZarazTimerRuleSettingsJSON contains the
// JSON metadata for the struct
// [ZarazExportGetResponseTriggersLoadRulesZarazTimerRuleSettings]
type zarazExportGetResponseTriggersLoadRulesZarazTimerRuleSettingsJSON struct {
	Interval    apijson.Field
	Limit       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazExportGetResponseTriggersLoadRulesZarazTimerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazExportGetResponseTriggersLoadRulesZarazFormSubmissionRule struct {
	ID       string                                                                 `json:"id,required"`
	Action   ZarazExportGetResponseTriggersLoadRulesZarazFormSubmissionRuleAction   `json:"action,required"`
	Settings ZarazExportGetResponseTriggersLoadRulesZarazFormSubmissionRuleSettings `json:"settings,required"`
	JSON     zarazExportGetResponseTriggersLoadRulesZarazFormSubmissionRuleJSON     `json:"-"`
}

// zarazExportGetResponseTriggersLoadRulesZarazFormSubmissionRuleJSON contains the
// JSON metadata for the struct
// [ZarazExportGetResponseTriggersLoadRulesZarazFormSubmissionRule]
type zarazExportGetResponseTriggersLoadRulesZarazFormSubmissionRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazExportGetResponseTriggersLoadRulesZarazFormSubmissionRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazExportGetResponseTriggersLoadRulesZarazFormSubmissionRule) implementsZarazExportGetResponseTriggersLoadRule() {
}

type ZarazExportGetResponseTriggersLoadRulesZarazFormSubmissionRuleAction string

const (
	ZarazExportGetResponseTriggersLoadRulesZarazFormSubmissionRuleActionFormSubmission ZarazExportGetResponseTriggersLoadRulesZarazFormSubmissionRuleAction = "formSubmission"
)

type ZarazExportGetResponseTriggersLoadRulesZarazFormSubmissionRuleSettings struct {
	Selector string                                                                     `json:"selector,required"`
	Validate bool                                                                       `json:"validate,required"`
	JSON     zarazExportGetResponseTriggersLoadRulesZarazFormSubmissionRuleSettingsJSON `json:"-"`
}

// zarazExportGetResponseTriggersLoadRulesZarazFormSubmissionRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazExportGetResponseTriggersLoadRulesZarazFormSubmissionRuleSettings]
type zarazExportGetResponseTriggersLoadRulesZarazFormSubmissionRuleSettingsJSON struct {
	Selector    apijson.Field
	Validate    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazExportGetResponseTriggersLoadRulesZarazFormSubmissionRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazExportGetResponseTriggersLoadRulesZarazVariableMatchRule struct {
	ID       string                                                                `json:"id,required"`
	Action   ZarazExportGetResponseTriggersLoadRulesZarazVariableMatchRuleAction   `json:"action,required"`
	Settings ZarazExportGetResponseTriggersLoadRulesZarazVariableMatchRuleSettings `json:"settings,required"`
	JSON     zarazExportGetResponseTriggersLoadRulesZarazVariableMatchRuleJSON     `json:"-"`
}

// zarazExportGetResponseTriggersLoadRulesZarazVariableMatchRuleJSON contains the
// JSON metadata for the struct
// [ZarazExportGetResponseTriggersLoadRulesZarazVariableMatchRule]
type zarazExportGetResponseTriggersLoadRulesZarazVariableMatchRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazExportGetResponseTriggersLoadRulesZarazVariableMatchRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazExportGetResponseTriggersLoadRulesZarazVariableMatchRule) implementsZarazExportGetResponseTriggersLoadRule() {
}

type ZarazExportGetResponseTriggersLoadRulesZarazVariableMatchRuleAction string

const (
	ZarazExportGetResponseTriggersLoadRulesZarazVariableMatchRuleActionVariableMatch ZarazExportGetResponseTriggersLoadRulesZarazVariableMatchRuleAction = "variableMatch"
)

type ZarazExportGetResponseTriggersLoadRulesZarazVariableMatchRuleSettings struct {
	Match    string                                                                    `json:"match,required"`
	Variable string                                                                    `json:"variable,required"`
	JSON     zarazExportGetResponseTriggersLoadRulesZarazVariableMatchRuleSettingsJSON `json:"-"`
}

// zarazExportGetResponseTriggersLoadRulesZarazVariableMatchRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazExportGetResponseTriggersLoadRulesZarazVariableMatchRuleSettings]
type zarazExportGetResponseTriggersLoadRulesZarazVariableMatchRuleSettingsJSON struct {
	Match       apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazExportGetResponseTriggersLoadRulesZarazVariableMatchRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazExportGetResponseTriggersLoadRulesZarazScrollDepthRule struct {
	ID       string                                                              `json:"id,required"`
	Action   ZarazExportGetResponseTriggersLoadRulesZarazScrollDepthRuleAction   `json:"action,required"`
	Settings ZarazExportGetResponseTriggersLoadRulesZarazScrollDepthRuleSettings `json:"settings,required"`
	JSON     zarazExportGetResponseTriggersLoadRulesZarazScrollDepthRuleJSON     `json:"-"`
}

// zarazExportGetResponseTriggersLoadRulesZarazScrollDepthRuleJSON contains the
// JSON metadata for the struct
// [ZarazExportGetResponseTriggersLoadRulesZarazScrollDepthRule]
type zarazExportGetResponseTriggersLoadRulesZarazScrollDepthRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazExportGetResponseTriggersLoadRulesZarazScrollDepthRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazExportGetResponseTriggersLoadRulesZarazScrollDepthRule) implementsZarazExportGetResponseTriggersLoadRule() {
}

type ZarazExportGetResponseTriggersLoadRulesZarazScrollDepthRuleAction string

const (
	ZarazExportGetResponseTriggersLoadRulesZarazScrollDepthRuleActionScrollDepth ZarazExportGetResponseTriggersLoadRulesZarazScrollDepthRuleAction = "scrollDepth"
)

type ZarazExportGetResponseTriggersLoadRulesZarazScrollDepthRuleSettings struct {
	Positions string                                                                  `json:"positions,required"`
	JSON      zarazExportGetResponseTriggersLoadRulesZarazScrollDepthRuleSettingsJSON `json:"-"`
}

// zarazExportGetResponseTriggersLoadRulesZarazScrollDepthRuleSettingsJSON contains
// the JSON metadata for the struct
// [ZarazExportGetResponseTriggersLoadRulesZarazScrollDepthRuleSettings]
type zarazExportGetResponseTriggersLoadRulesZarazScrollDepthRuleSettingsJSON struct {
	Positions   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazExportGetResponseTriggersLoadRulesZarazScrollDepthRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazExportGetResponseTriggersLoadRulesZarazElementVisibilityRule struct {
	ID       string                                                                    `json:"id,required"`
	Action   ZarazExportGetResponseTriggersLoadRulesZarazElementVisibilityRuleAction   `json:"action,required"`
	Settings ZarazExportGetResponseTriggersLoadRulesZarazElementVisibilityRuleSettings `json:"settings,required"`
	JSON     zarazExportGetResponseTriggersLoadRulesZarazElementVisibilityRuleJSON     `json:"-"`
}

// zarazExportGetResponseTriggersLoadRulesZarazElementVisibilityRuleJSON contains
// the JSON metadata for the struct
// [ZarazExportGetResponseTriggersLoadRulesZarazElementVisibilityRule]
type zarazExportGetResponseTriggersLoadRulesZarazElementVisibilityRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazExportGetResponseTriggersLoadRulesZarazElementVisibilityRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazExportGetResponseTriggersLoadRulesZarazElementVisibilityRule) implementsZarazExportGetResponseTriggersLoadRule() {
}

type ZarazExportGetResponseTriggersLoadRulesZarazElementVisibilityRuleAction string

const (
	ZarazExportGetResponseTriggersLoadRulesZarazElementVisibilityRuleActionElementVisibility ZarazExportGetResponseTriggersLoadRulesZarazElementVisibilityRuleAction = "elementVisibility"
)

type ZarazExportGetResponseTriggersLoadRulesZarazElementVisibilityRuleSettings struct {
	Selector string                                                                        `json:"selector,required"`
	JSON     zarazExportGetResponseTriggersLoadRulesZarazElementVisibilityRuleSettingsJSON `json:"-"`
}

// zarazExportGetResponseTriggersLoadRulesZarazElementVisibilityRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazExportGetResponseTriggersLoadRulesZarazElementVisibilityRuleSettings]
type zarazExportGetResponseTriggersLoadRulesZarazElementVisibilityRuleSettingsJSON struct {
	Selector    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazExportGetResponseTriggersLoadRulesZarazElementVisibilityRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazExportGetResponseTriggersSystem string

const (
	ZarazExportGetResponseTriggersSystemPageload ZarazExportGetResponseTriggersSystem = "pageload"
)

// Union satisfied by [ZarazExportGetResponseVariablesObject] or
// [ZarazExportGetResponseVariablesObject].
type ZarazExportGetResponseVariable interface {
	implementsZarazExportGetResponseVariable()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZarazExportGetResponseVariable)(nil)).Elem(), "")
}

type ZarazExportGetResponseVariablesObject struct {
	Name  string                                    `json:"name,required"`
	Type  ZarazExportGetResponseVariablesObjectType `json:"type,required"`
	Value string                                    `json:"value,required"`
	JSON  zarazExportGetResponseVariablesObjectJSON `json:"-"`
}

// zarazExportGetResponseVariablesObjectJSON contains the JSON metadata for the
// struct [ZarazExportGetResponseVariablesObject]
type zarazExportGetResponseVariablesObjectJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazExportGetResponseVariablesObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazExportGetResponseVariablesObject) implementsZarazExportGetResponseVariable() {}

type ZarazExportGetResponseVariablesObjectType string

const (
	ZarazExportGetResponseVariablesObjectTypeString ZarazExportGetResponseVariablesObjectType = "string"
	ZarazExportGetResponseVariablesObjectTypeSecret ZarazExportGetResponseVariablesObjectType = "secret"
)

// Consent management configuration.
type ZarazExportGetResponseConsent struct {
	Enabled                bool                                                `json:"enabled,required"`
	ButtonTextTranslations ZarazExportGetResponseConsentButtonTextTranslations `json:"buttonTextTranslations"`
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
	Purposes map[string]ZarazExportGetResponseConsentPurpose `json:"purposes"`
	// Object where keys are purpose alpha-numeric IDs
	PurposesWithTranslations map[string]ZarazExportGetResponseConsentPurposesWithTranslation `json:"purposesWithTranslations"`
	JSON                     zarazExportGetResponseConsentJSON                               `json:"-"`
}

// zarazExportGetResponseConsentJSON contains the JSON metadata for the struct
// [ZarazExportGetResponseConsent]
type zarazExportGetResponseConsentJSON struct {
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

func (r *ZarazExportGetResponseConsent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazExportGetResponseConsentButtonTextTranslations struct {
	// Object where keys are language codes
	AcceptAll map[string]string `json:"accept_all,required"`
	// Object where keys are language codes
	ConfirmMyChoices map[string]string `json:"confirm_my_choices,required"`
	// Object where keys are language codes
	RejectAll map[string]string                                       `json:"reject_all,required"`
	JSON      zarazExportGetResponseConsentButtonTextTranslationsJSON `json:"-"`
}

// zarazExportGetResponseConsentButtonTextTranslationsJSON contains the JSON
// metadata for the struct [ZarazExportGetResponseConsentButtonTextTranslations]
type zarazExportGetResponseConsentButtonTextTranslationsJSON struct {
	AcceptAll        apijson.Field
	ConfirmMyChoices apijson.Field
	RejectAll        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazExportGetResponseConsentButtonTextTranslations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazExportGetResponseConsentPurpose struct {
	Description string                                   `json:"description,required"`
	Name        string                                   `json:"name,required"`
	JSON        zarazExportGetResponseConsentPurposeJSON `json:"-"`
}

// zarazExportGetResponseConsentPurposeJSON contains the JSON metadata for the
// struct [ZarazExportGetResponseConsentPurpose]
type zarazExportGetResponseConsentPurposeJSON struct {
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazExportGetResponseConsentPurpose) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazExportGetResponseConsentPurposesWithTranslation struct {
	// Object where keys are language codes
	Description map[string]string `json:"description,required"`
	// Object where keys are language codes
	Name  map[string]string                                        `json:"name,required"`
	Order int64                                                    `json:"order,required"`
	JSON  zarazExportGetResponseConsentPurposesWithTranslationJSON `json:"-"`
}

// zarazExportGetResponseConsentPurposesWithTranslationJSON contains the JSON
// metadata for the struct [ZarazExportGetResponseConsentPurposesWithTranslation]
type zarazExportGetResponseConsentPurposesWithTranslationJSON struct {
	Description apijson.Field
	Name        apijson.Field
	Order       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazExportGetResponseConsentPurposesWithTranslation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
