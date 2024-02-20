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

// Restores a historical published Zaraz configuration by ID for a zone.
func (r *ZarazHistoryService) Replace(ctx context.Context, zoneID string, body ZarazHistoryReplaceParams, opts ...option.RequestOption) (res *ZarazHistoryReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZarazHistoryReplaceResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/zaraz/history", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
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

// Zaraz configuration
type ZarazHistoryReplaceResponse struct {
	// Data layer compatibility mode enabled.
	DataLayer bool `json:"dataLayer,required"`
	// The key for Zaraz debug mode.
	DebugKey string `json:"debugKey,required"`
	// General Zaraz settings.
	Settings ZarazHistoryReplaceResponseSettings `json:"settings,required"`
	// Tools set up under Zaraz configuration, where key is the alpha-numeric tool ID
	// and value is the tool configuration object.
	Tools map[string]ZarazHistoryReplaceResponseTool `json:"tools,required"`
	// Triggers set up under Zaraz configuration, where key is the trigger
	// alpha-numeric ID and value is the trigger configuration.
	Triggers map[string]ZarazHistoryReplaceResponseTrigger `json:"triggers,required"`
	// Variables set up under Zaraz configuration, where key is the variable
	// alpha-numeric ID and value is the variable configuration. Values of variables of
	// type secret are not included.
	Variables map[string]ZarazHistoryReplaceResponseVariable `json:"variables,required"`
	// Zaraz internal version of the config.
	ZarazVersion int64 `json:"zarazVersion,required"`
	// Consent management configuration.
	Consent ZarazHistoryReplaceResponseConsent `json:"consent"`
	// Single Page Application support enabled.
	HistoryChange bool                            `json:"historyChange"`
	JSON          zarazHistoryReplaceResponseJSON `json:"-"`
}

// zarazHistoryReplaceResponseJSON contains the JSON metadata for the struct
// [ZarazHistoryReplaceResponse]
type zarazHistoryReplaceResponseJSON struct {
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

func (r *ZarazHistoryReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// General Zaraz settings.
type ZarazHistoryReplaceResponseSettings struct {
	// Automatic injection of Zaraz scripts enabled.
	AutoInjectScript bool `json:"autoInjectScript,required"`
	// Details of the worker that receives and edits Zaraz Context object.
	ContextEnricher ZarazHistoryReplaceResponseSettingsContextEnricher `json:"contextEnricher"`
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
	TrackPath string                                  `json:"trackPath"`
	JSON      zarazHistoryReplaceResponseSettingsJSON `json:"-"`
}

// zarazHistoryReplaceResponseSettingsJSON contains the JSON metadata for the
// struct [ZarazHistoryReplaceResponseSettings]
type zarazHistoryReplaceResponseSettingsJSON struct {
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

func (r *ZarazHistoryReplaceResponseSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details of the worker that receives and edits Zaraz Context object.
type ZarazHistoryReplaceResponseSettingsContextEnricher struct {
	EscapedWorkerName string                                                 `json:"escapedWorkerName,required"`
	WorkerTag         string                                                 `json:"workerTag,required"`
	JSON              zarazHistoryReplaceResponseSettingsContextEnricherJSON `json:"-"`
}

// zarazHistoryReplaceResponseSettingsContextEnricherJSON contains the JSON
// metadata for the struct [ZarazHistoryReplaceResponseSettingsContextEnricher]
type zarazHistoryReplaceResponseSettingsContextEnricherJSON struct {
	EscapedWorkerName apijson.Field
	WorkerTag         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseSettingsContextEnricher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [ZarazHistoryReplaceResponseToolsZarazManagedComponent] or
// [ZarazHistoryReplaceResponseToolsZarazCustomManagedComponent].
type ZarazHistoryReplaceResponseTool interface {
	implementsZarazHistoryReplaceResponseTool()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZarazHistoryReplaceResponseTool)(nil)).Elem(), "")
}

type ZarazHistoryReplaceResponseToolsZarazManagedComponent struct {
	// List of blocking trigger IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Tool's internal name
	Component string `json:"component,required"`
	// Default fields for tool's actions
	DefaultFields map[string]ZarazHistoryReplaceResponseToolsZarazManagedComponentDefaultField `json:"defaultFields,required"`
	// Whether tool is enabled
	Enabled bool `json:"enabled,required"`
	// Tool's name defined by the user
	Name string `json:"name,required"`
	// List of permissions granted to the component
	Permissions []string `json:"permissions,required"`
	// Tool's settings
	Settings map[string]ZarazHistoryReplaceResponseToolsZarazManagedComponentSetting `json:"settings,required"`
	Type     ZarazHistoryReplaceResponseToolsZarazManagedComponentType               `json:"type,required"`
	// Actions configured on a tool. Either this or neoEvents field is required.
	Actions map[string]ZarazHistoryReplaceResponseToolsZarazManagedComponentAction `json:"actions"`
	// Default consent purpose ID
	DefaultPurpose string `json:"defaultPurpose"`
	// DEPRECATED - List of actions configured on a tool. Either this or actions field
	// is required. If both are present, actions field will take precedence.
	NeoEvents []ZarazHistoryReplaceResponseToolsZarazManagedComponentNeoEvent `json:"neoEvents"`
	JSON      zarazHistoryReplaceResponseToolsZarazManagedComponentJSON       `json:"-"`
}

// zarazHistoryReplaceResponseToolsZarazManagedComponentJSON contains the JSON
// metadata for the struct [ZarazHistoryReplaceResponseToolsZarazManagedComponent]
type zarazHistoryReplaceResponseToolsZarazManagedComponentJSON struct {
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

func (r *ZarazHistoryReplaceResponseToolsZarazManagedComponent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazHistoryReplaceResponseToolsZarazManagedComponent) implementsZarazHistoryReplaceResponseTool() {
}

// Union satisfied by [shared.UnionString] or [shared.UnionBool].
type ZarazHistoryReplaceResponseToolsZarazManagedComponentDefaultField interface {
	ImplementsZarazHistoryReplaceResponseToolsZarazManagedComponentDefaultField()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazHistoryReplaceResponseToolsZarazManagedComponentDefaultField)(nil)).Elem(),
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
type ZarazHistoryReplaceResponseToolsZarazManagedComponentSetting interface {
	ImplementsZarazHistoryReplaceResponseToolsZarazManagedComponentSetting()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazHistoryReplaceResponseToolsZarazManagedComponentSetting)(nil)).Elem(),
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

type ZarazHistoryReplaceResponseToolsZarazManagedComponentType string

const (
	ZarazHistoryReplaceResponseToolsZarazManagedComponentTypeComponent ZarazHistoryReplaceResponseToolsZarazManagedComponentType = "component"
)

type ZarazHistoryReplaceResponseToolsZarazManagedComponentAction struct {
	// Tool event type
	ActionType string `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Event payload
	Data interface{} `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers []string                                                        `json:"firingTriggers,required"`
	JSON           zarazHistoryReplaceResponseToolsZarazManagedComponentActionJSON `json:"-"`
}

// zarazHistoryReplaceResponseToolsZarazManagedComponentActionJSON contains the
// JSON metadata for the struct
// [ZarazHistoryReplaceResponseToolsZarazManagedComponentAction]
type zarazHistoryReplaceResponseToolsZarazManagedComponentActionJSON struct {
	ActionType       apijson.Field
	BlockingTriggers apijson.Field
	Data             apijson.Field
	FiringTriggers   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseToolsZarazManagedComponentAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryReplaceResponseToolsZarazManagedComponentNeoEvent struct {
	// Tool event type
	ActionType string `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Event payload
	Data interface{} `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers []string                                                          `json:"firingTriggers,required"`
	JSON           zarazHistoryReplaceResponseToolsZarazManagedComponentNeoEventJSON `json:"-"`
}

// zarazHistoryReplaceResponseToolsZarazManagedComponentNeoEventJSON contains the
// JSON metadata for the struct
// [ZarazHistoryReplaceResponseToolsZarazManagedComponentNeoEvent]
type zarazHistoryReplaceResponseToolsZarazManagedComponentNeoEventJSON struct {
	ActionType       apijson.Field
	BlockingTriggers apijson.Field
	Data             apijson.Field
	FiringTriggers   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseToolsZarazManagedComponentNeoEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryReplaceResponseToolsZarazCustomManagedComponent struct {
	// List of blocking trigger IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Tool's internal name
	Component string `json:"component,required"`
	// Default fields for tool's actions
	DefaultFields map[string]ZarazHistoryReplaceResponseToolsZarazCustomManagedComponentDefaultField `json:"defaultFields,required"`
	// Whether tool is enabled
	Enabled bool `json:"enabled,required"`
	// Tool's name defined by the user
	Name string `json:"name,required"`
	// List of permissions granted to the component
	Permissions []string `json:"permissions,required"`
	// Tool's settings
	Settings map[string]ZarazHistoryReplaceResponseToolsZarazCustomManagedComponentSetting `json:"settings,required"`
	Type     ZarazHistoryReplaceResponseToolsZarazCustomManagedComponentType               `json:"type,required"`
	// Cloudflare worker that acts as a managed component
	Worker ZarazHistoryReplaceResponseToolsZarazCustomManagedComponentWorker `json:"worker,required"`
	// Actions configured on a tool. Either this or neoEvents field is required.
	Actions map[string]ZarazHistoryReplaceResponseToolsZarazCustomManagedComponentAction `json:"actions"`
	// Default consent purpose ID
	DefaultPurpose string `json:"defaultPurpose"`
	// DEPRECATED - List of actions configured on a tool. Either this or actions field
	// is required. If both are present, actions field will take precedence.
	NeoEvents []ZarazHistoryReplaceResponseToolsZarazCustomManagedComponentNeoEvent `json:"neoEvents"`
	JSON      zarazHistoryReplaceResponseToolsZarazCustomManagedComponentJSON       `json:"-"`
}

// zarazHistoryReplaceResponseToolsZarazCustomManagedComponentJSON contains the
// JSON metadata for the struct
// [ZarazHistoryReplaceResponseToolsZarazCustomManagedComponent]
type zarazHistoryReplaceResponseToolsZarazCustomManagedComponentJSON struct {
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

func (r *ZarazHistoryReplaceResponseToolsZarazCustomManagedComponent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazHistoryReplaceResponseToolsZarazCustomManagedComponent) implementsZarazHistoryReplaceResponseTool() {
}

// Union satisfied by [shared.UnionString] or [shared.UnionBool].
type ZarazHistoryReplaceResponseToolsZarazCustomManagedComponentDefaultField interface {
	ImplementsZarazHistoryReplaceResponseToolsZarazCustomManagedComponentDefaultField()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazHistoryReplaceResponseToolsZarazCustomManagedComponentDefaultField)(nil)).Elem(),
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
type ZarazHistoryReplaceResponseToolsZarazCustomManagedComponentSetting interface {
	ImplementsZarazHistoryReplaceResponseToolsZarazCustomManagedComponentSetting()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZarazHistoryReplaceResponseToolsZarazCustomManagedComponentSetting)(nil)).Elem(),
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

type ZarazHistoryReplaceResponseToolsZarazCustomManagedComponentType string

const (
	ZarazHistoryReplaceResponseToolsZarazCustomManagedComponentTypeCustomMc ZarazHistoryReplaceResponseToolsZarazCustomManagedComponentType = "custom-mc"
)

// Cloudflare worker that acts as a managed component
type ZarazHistoryReplaceResponseToolsZarazCustomManagedComponentWorker struct {
	EscapedWorkerName string                                                                `json:"escapedWorkerName,required"`
	WorkerTag         string                                                                `json:"workerTag,required"`
	JSON              zarazHistoryReplaceResponseToolsZarazCustomManagedComponentWorkerJSON `json:"-"`
}

// zarazHistoryReplaceResponseToolsZarazCustomManagedComponentWorkerJSON contains
// the JSON metadata for the struct
// [ZarazHistoryReplaceResponseToolsZarazCustomManagedComponentWorker]
type zarazHistoryReplaceResponseToolsZarazCustomManagedComponentWorkerJSON struct {
	EscapedWorkerName apijson.Field
	WorkerTag         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseToolsZarazCustomManagedComponentWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryReplaceResponseToolsZarazCustomManagedComponentAction struct {
	// Tool event type
	ActionType string `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Event payload
	Data interface{} `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers []string                                                              `json:"firingTriggers,required"`
	JSON           zarazHistoryReplaceResponseToolsZarazCustomManagedComponentActionJSON `json:"-"`
}

// zarazHistoryReplaceResponseToolsZarazCustomManagedComponentActionJSON contains
// the JSON metadata for the struct
// [ZarazHistoryReplaceResponseToolsZarazCustomManagedComponentAction]
type zarazHistoryReplaceResponseToolsZarazCustomManagedComponentActionJSON struct {
	ActionType       apijson.Field
	BlockingTriggers apijson.Field
	Data             apijson.Field
	FiringTriggers   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseToolsZarazCustomManagedComponentAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryReplaceResponseToolsZarazCustomManagedComponentNeoEvent struct {
	// Tool event type
	ActionType string `json:"actionType,required"`
	// List of blocking triggers IDs
	BlockingTriggers []string `json:"blockingTriggers,required"`
	// Event payload
	Data interface{} `json:"data,required"`
	// List of firing triggers IDs
	FiringTriggers []string                                                                `json:"firingTriggers,required"`
	JSON           zarazHistoryReplaceResponseToolsZarazCustomManagedComponentNeoEventJSON `json:"-"`
}

// zarazHistoryReplaceResponseToolsZarazCustomManagedComponentNeoEventJSON contains
// the JSON metadata for the struct
// [ZarazHistoryReplaceResponseToolsZarazCustomManagedComponentNeoEvent]
type zarazHistoryReplaceResponseToolsZarazCustomManagedComponentNeoEventJSON struct {
	ActionType       apijson.Field
	BlockingTriggers apijson.Field
	Data             apijson.Field
	FiringTriggers   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseToolsZarazCustomManagedComponentNeoEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryReplaceResponseTrigger struct {
	// Rules defining when the trigger is not fired.
	ExcludeRules []ZarazHistoryReplaceResponseTriggersExcludeRule `json:"excludeRules,required"`
	// Rules defining when the trigger is fired.
	LoadRules []ZarazHistoryReplaceResponseTriggersLoadRule `json:"loadRules,required"`
	// Trigger name.
	Name string `json:"name,required"`
	// Trigger description.
	Description string                                    `json:"description"`
	System      ZarazHistoryReplaceResponseTriggersSystem `json:"system"`
	JSON        zarazHistoryReplaceResponseTriggerJSON    `json:"-"`
}

// zarazHistoryReplaceResponseTriggerJSON contains the JSON metadata for the struct
// [ZarazHistoryReplaceResponseTrigger]
type zarazHistoryReplaceResponseTriggerJSON struct {
	ExcludeRules apijson.Field
	LoadRules    apijson.Field
	Name         apijson.Field
	Description  apijson.Field
	System       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [ZarazHistoryReplaceResponseTriggersExcludeRulesZarazLoadRule],
// [ZarazHistoryReplaceResponseTriggersExcludeRulesZarazClickListenerRule],
// [ZarazHistoryReplaceResponseTriggersExcludeRulesZarazTimerRule],
// [ZarazHistoryReplaceResponseTriggersExcludeRulesZarazFormSubmissionRule],
// [ZarazHistoryReplaceResponseTriggersExcludeRulesZarazVariableMatchRule],
// [ZarazHistoryReplaceResponseTriggersExcludeRulesZarazScrollDepthRule] or
// [ZarazHistoryReplaceResponseTriggersExcludeRulesZarazElementVisibilityRule].
type ZarazHistoryReplaceResponseTriggersExcludeRule interface {
	implementsZarazHistoryReplaceResponseTriggersExcludeRule()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZarazHistoryReplaceResponseTriggersExcludeRule)(nil)).Elem(), "")
}

type ZarazHistoryReplaceResponseTriggersExcludeRulesZarazLoadRule struct {
	ID    string                                                           `json:"id,required"`
	Match string                                                           `json:"match,required"`
	Op    ZarazHistoryReplaceResponseTriggersExcludeRulesZarazLoadRuleOp   `json:"op,required"`
	Value string                                                           `json:"value,required"`
	JSON  zarazHistoryReplaceResponseTriggersExcludeRulesZarazLoadRuleJSON `json:"-"`
}

// zarazHistoryReplaceResponseTriggersExcludeRulesZarazLoadRuleJSON contains the
// JSON metadata for the struct
// [ZarazHistoryReplaceResponseTriggersExcludeRulesZarazLoadRule]
type zarazHistoryReplaceResponseTriggersExcludeRulesZarazLoadRuleJSON struct {
	ID          apijson.Field
	Match       apijson.Field
	Op          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseTriggersExcludeRulesZarazLoadRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazHistoryReplaceResponseTriggersExcludeRulesZarazLoadRule) implementsZarazHistoryReplaceResponseTriggersExcludeRule() {
}

type ZarazHistoryReplaceResponseTriggersExcludeRulesZarazLoadRuleOp string

const (
	ZarazHistoryReplaceResponseTriggersExcludeRulesZarazLoadRuleOpContains           ZarazHistoryReplaceResponseTriggersExcludeRulesZarazLoadRuleOp = "CONTAINS"
	ZarazHistoryReplaceResponseTriggersExcludeRulesZarazLoadRuleOpEquals             ZarazHistoryReplaceResponseTriggersExcludeRulesZarazLoadRuleOp = "EQUALS"
	ZarazHistoryReplaceResponseTriggersExcludeRulesZarazLoadRuleOpStartsWith         ZarazHistoryReplaceResponseTriggersExcludeRulesZarazLoadRuleOp = "STARTS_WITH"
	ZarazHistoryReplaceResponseTriggersExcludeRulesZarazLoadRuleOpEndsWith           ZarazHistoryReplaceResponseTriggersExcludeRulesZarazLoadRuleOp = "ENDS_WITH"
	ZarazHistoryReplaceResponseTriggersExcludeRulesZarazLoadRuleOpMatchRegex         ZarazHistoryReplaceResponseTriggersExcludeRulesZarazLoadRuleOp = "MATCH_REGEX"
	ZarazHistoryReplaceResponseTriggersExcludeRulesZarazLoadRuleOpNotMatchRegex      ZarazHistoryReplaceResponseTriggersExcludeRulesZarazLoadRuleOp = "NOT_MATCH_REGEX"
	ZarazHistoryReplaceResponseTriggersExcludeRulesZarazLoadRuleOpGreaterThan        ZarazHistoryReplaceResponseTriggersExcludeRulesZarazLoadRuleOp = "GREATER_THAN"
	ZarazHistoryReplaceResponseTriggersExcludeRulesZarazLoadRuleOpGreaterThanOrEqual ZarazHistoryReplaceResponseTriggersExcludeRulesZarazLoadRuleOp = "GREATER_THAN_OR_EQUAL"
	ZarazHistoryReplaceResponseTriggersExcludeRulesZarazLoadRuleOpLessThan           ZarazHistoryReplaceResponseTriggersExcludeRulesZarazLoadRuleOp = "LESS_THAN"
	ZarazHistoryReplaceResponseTriggersExcludeRulesZarazLoadRuleOpLessThanOrEqual    ZarazHistoryReplaceResponseTriggersExcludeRulesZarazLoadRuleOp = "LESS_THAN_OR_EQUAL"
)

type ZarazHistoryReplaceResponseTriggersExcludeRulesZarazClickListenerRule struct {
	ID       string                                                                        `json:"id,required"`
	Action   ZarazHistoryReplaceResponseTriggersExcludeRulesZarazClickListenerRuleAction   `json:"action,required"`
	Settings ZarazHistoryReplaceResponseTriggersExcludeRulesZarazClickListenerRuleSettings `json:"settings,required"`
	JSON     zarazHistoryReplaceResponseTriggersExcludeRulesZarazClickListenerRuleJSON     `json:"-"`
}

// zarazHistoryReplaceResponseTriggersExcludeRulesZarazClickListenerRuleJSON
// contains the JSON metadata for the struct
// [ZarazHistoryReplaceResponseTriggersExcludeRulesZarazClickListenerRule]
type zarazHistoryReplaceResponseTriggersExcludeRulesZarazClickListenerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseTriggersExcludeRulesZarazClickListenerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazHistoryReplaceResponseTriggersExcludeRulesZarazClickListenerRule) implementsZarazHistoryReplaceResponseTriggersExcludeRule() {
}

type ZarazHistoryReplaceResponseTriggersExcludeRulesZarazClickListenerRuleAction string

const (
	ZarazHistoryReplaceResponseTriggersExcludeRulesZarazClickListenerRuleActionClickListener ZarazHistoryReplaceResponseTriggersExcludeRulesZarazClickListenerRuleAction = "clickListener"
)

type ZarazHistoryReplaceResponseTriggersExcludeRulesZarazClickListenerRuleSettings struct {
	Selector    string                                                                            `json:"selector,required"`
	Type        ZarazHistoryReplaceResponseTriggersExcludeRulesZarazClickListenerRuleSettingsType `json:"type,required"`
	WaitForTags int64                                                                             `json:"waitForTags,required"`
	JSON        zarazHistoryReplaceResponseTriggersExcludeRulesZarazClickListenerRuleSettingsJSON `json:"-"`
}

// zarazHistoryReplaceResponseTriggersExcludeRulesZarazClickListenerRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazHistoryReplaceResponseTriggersExcludeRulesZarazClickListenerRuleSettings]
type zarazHistoryReplaceResponseTriggersExcludeRulesZarazClickListenerRuleSettingsJSON struct {
	Selector    apijson.Field
	Type        apijson.Field
	WaitForTags apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseTriggersExcludeRulesZarazClickListenerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryReplaceResponseTriggersExcludeRulesZarazClickListenerRuleSettingsType string

const (
	ZarazHistoryReplaceResponseTriggersExcludeRulesZarazClickListenerRuleSettingsTypeXpath ZarazHistoryReplaceResponseTriggersExcludeRulesZarazClickListenerRuleSettingsType = "xpath"
	ZarazHistoryReplaceResponseTriggersExcludeRulesZarazClickListenerRuleSettingsTypeCss   ZarazHistoryReplaceResponseTriggersExcludeRulesZarazClickListenerRuleSettingsType = "css"
)

type ZarazHistoryReplaceResponseTriggersExcludeRulesZarazTimerRule struct {
	ID       string                                                                `json:"id,required"`
	Action   ZarazHistoryReplaceResponseTriggersExcludeRulesZarazTimerRuleAction   `json:"action,required"`
	Settings ZarazHistoryReplaceResponseTriggersExcludeRulesZarazTimerRuleSettings `json:"settings,required"`
	JSON     zarazHistoryReplaceResponseTriggersExcludeRulesZarazTimerRuleJSON     `json:"-"`
}

// zarazHistoryReplaceResponseTriggersExcludeRulesZarazTimerRuleJSON contains the
// JSON metadata for the struct
// [ZarazHistoryReplaceResponseTriggersExcludeRulesZarazTimerRule]
type zarazHistoryReplaceResponseTriggersExcludeRulesZarazTimerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseTriggersExcludeRulesZarazTimerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazHistoryReplaceResponseTriggersExcludeRulesZarazTimerRule) implementsZarazHistoryReplaceResponseTriggersExcludeRule() {
}

type ZarazHistoryReplaceResponseTriggersExcludeRulesZarazTimerRuleAction string

const (
	ZarazHistoryReplaceResponseTriggersExcludeRulesZarazTimerRuleActionTimer ZarazHistoryReplaceResponseTriggersExcludeRulesZarazTimerRuleAction = "timer"
)

type ZarazHistoryReplaceResponseTriggersExcludeRulesZarazTimerRuleSettings struct {
	Interval int64                                                                     `json:"interval,required"`
	Limit    int64                                                                     `json:"limit,required"`
	JSON     zarazHistoryReplaceResponseTriggersExcludeRulesZarazTimerRuleSettingsJSON `json:"-"`
}

// zarazHistoryReplaceResponseTriggersExcludeRulesZarazTimerRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazHistoryReplaceResponseTriggersExcludeRulesZarazTimerRuleSettings]
type zarazHistoryReplaceResponseTriggersExcludeRulesZarazTimerRuleSettingsJSON struct {
	Interval    apijson.Field
	Limit       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseTriggersExcludeRulesZarazTimerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryReplaceResponseTriggersExcludeRulesZarazFormSubmissionRule struct {
	ID       string                                                                         `json:"id,required"`
	Action   ZarazHistoryReplaceResponseTriggersExcludeRulesZarazFormSubmissionRuleAction   `json:"action,required"`
	Settings ZarazHistoryReplaceResponseTriggersExcludeRulesZarazFormSubmissionRuleSettings `json:"settings,required"`
	JSON     zarazHistoryReplaceResponseTriggersExcludeRulesZarazFormSubmissionRuleJSON     `json:"-"`
}

// zarazHistoryReplaceResponseTriggersExcludeRulesZarazFormSubmissionRuleJSON
// contains the JSON metadata for the struct
// [ZarazHistoryReplaceResponseTriggersExcludeRulesZarazFormSubmissionRule]
type zarazHistoryReplaceResponseTriggersExcludeRulesZarazFormSubmissionRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseTriggersExcludeRulesZarazFormSubmissionRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazHistoryReplaceResponseTriggersExcludeRulesZarazFormSubmissionRule) implementsZarazHistoryReplaceResponseTriggersExcludeRule() {
}

type ZarazHistoryReplaceResponseTriggersExcludeRulesZarazFormSubmissionRuleAction string

const (
	ZarazHistoryReplaceResponseTriggersExcludeRulesZarazFormSubmissionRuleActionFormSubmission ZarazHistoryReplaceResponseTriggersExcludeRulesZarazFormSubmissionRuleAction = "formSubmission"
)

type ZarazHistoryReplaceResponseTriggersExcludeRulesZarazFormSubmissionRuleSettings struct {
	Selector string                                                                             `json:"selector,required"`
	Validate bool                                                                               `json:"validate,required"`
	JSON     zarazHistoryReplaceResponseTriggersExcludeRulesZarazFormSubmissionRuleSettingsJSON `json:"-"`
}

// zarazHistoryReplaceResponseTriggersExcludeRulesZarazFormSubmissionRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazHistoryReplaceResponseTriggersExcludeRulesZarazFormSubmissionRuleSettings]
type zarazHistoryReplaceResponseTriggersExcludeRulesZarazFormSubmissionRuleSettingsJSON struct {
	Selector    apijson.Field
	Validate    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseTriggersExcludeRulesZarazFormSubmissionRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryReplaceResponseTriggersExcludeRulesZarazVariableMatchRule struct {
	ID       string                                                                        `json:"id,required"`
	Action   ZarazHistoryReplaceResponseTriggersExcludeRulesZarazVariableMatchRuleAction   `json:"action,required"`
	Settings ZarazHistoryReplaceResponseTriggersExcludeRulesZarazVariableMatchRuleSettings `json:"settings,required"`
	JSON     zarazHistoryReplaceResponseTriggersExcludeRulesZarazVariableMatchRuleJSON     `json:"-"`
}

// zarazHistoryReplaceResponseTriggersExcludeRulesZarazVariableMatchRuleJSON
// contains the JSON metadata for the struct
// [ZarazHistoryReplaceResponseTriggersExcludeRulesZarazVariableMatchRule]
type zarazHistoryReplaceResponseTriggersExcludeRulesZarazVariableMatchRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseTriggersExcludeRulesZarazVariableMatchRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazHistoryReplaceResponseTriggersExcludeRulesZarazVariableMatchRule) implementsZarazHistoryReplaceResponseTriggersExcludeRule() {
}

type ZarazHistoryReplaceResponseTriggersExcludeRulesZarazVariableMatchRuleAction string

const (
	ZarazHistoryReplaceResponseTriggersExcludeRulesZarazVariableMatchRuleActionVariableMatch ZarazHistoryReplaceResponseTriggersExcludeRulesZarazVariableMatchRuleAction = "variableMatch"
)

type ZarazHistoryReplaceResponseTriggersExcludeRulesZarazVariableMatchRuleSettings struct {
	Match    string                                                                            `json:"match,required"`
	Variable string                                                                            `json:"variable,required"`
	JSON     zarazHistoryReplaceResponseTriggersExcludeRulesZarazVariableMatchRuleSettingsJSON `json:"-"`
}

// zarazHistoryReplaceResponseTriggersExcludeRulesZarazVariableMatchRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazHistoryReplaceResponseTriggersExcludeRulesZarazVariableMatchRuleSettings]
type zarazHistoryReplaceResponseTriggersExcludeRulesZarazVariableMatchRuleSettingsJSON struct {
	Match       apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseTriggersExcludeRulesZarazVariableMatchRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryReplaceResponseTriggersExcludeRulesZarazScrollDepthRule struct {
	ID       string                                                                      `json:"id,required"`
	Action   ZarazHistoryReplaceResponseTriggersExcludeRulesZarazScrollDepthRuleAction   `json:"action,required"`
	Settings ZarazHistoryReplaceResponseTriggersExcludeRulesZarazScrollDepthRuleSettings `json:"settings,required"`
	JSON     zarazHistoryReplaceResponseTriggersExcludeRulesZarazScrollDepthRuleJSON     `json:"-"`
}

// zarazHistoryReplaceResponseTriggersExcludeRulesZarazScrollDepthRuleJSON contains
// the JSON metadata for the struct
// [ZarazHistoryReplaceResponseTriggersExcludeRulesZarazScrollDepthRule]
type zarazHistoryReplaceResponseTriggersExcludeRulesZarazScrollDepthRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseTriggersExcludeRulesZarazScrollDepthRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazHistoryReplaceResponseTriggersExcludeRulesZarazScrollDepthRule) implementsZarazHistoryReplaceResponseTriggersExcludeRule() {
}

type ZarazHistoryReplaceResponseTriggersExcludeRulesZarazScrollDepthRuleAction string

const (
	ZarazHistoryReplaceResponseTriggersExcludeRulesZarazScrollDepthRuleActionScrollDepth ZarazHistoryReplaceResponseTriggersExcludeRulesZarazScrollDepthRuleAction = "scrollDepth"
)

type ZarazHistoryReplaceResponseTriggersExcludeRulesZarazScrollDepthRuleSettings struct {
	Positions string                                                                          `json:"positions,required"`
	JSON      zarazHistoryReplaceResponseTriggersExcludeRulesZarazScrollDepthRuleSettingsJSON `json:"-"`
}

// zarazHistoryReplaceResponseTriggersExcludeRulesZarazScrollDepthRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazHistoryReplaceResponseTriggersExcludeRulesZarazScrollDepthRuleSettings]
type zarazHistoryReplaceResponseTriggersExcludeRulesZarazScrollDepthRuleSettingsJSON struct {
	Positions   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseTriggersExcludeRulesZarazScrollDepthRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryReplaceResponseTriggersExcludeRulesZarazElementVisibilityRule struct {
	ID       string                                                                            `json:"id,required"`
	Action   ZarazHistoryReplaceResponseTriggersExcludeRulesZarazElementVisibilityRuleAction   `json:"action,required"`
	Settings ZarazHistoryReplaceResponseTriggersExcludeRulesZarazElementVisibilityRuleSettings `json:"settings,required"`
	JSON     zarazHistoryReplaceResponseTriggersExcludeRulesZarazElementVisibilityRuleJSON     `json:"-"`
}

// zarazHistoryReplaceResponseTriggersExcludeRulesZarazElementVisibilityRuleJSON
// contains the JSON metadata for the struct
// [ZarazHistoryReplaceResponseTriggersExcludeRulesZarazElementVisibilityRule]
type zarazHistoryReplaceResponseTriggersExcludeRulesZarazElementVisibilityRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseTriggersExcludeRulesZarazElementVisibilityRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazHistoryReplaceResponseTriggersExcludeRulesZarazElementVisibilityRule) implementsZarazHistoryReplaceResponseTriggersExcludeRule() {
}

type ZarazHistoryReplaceResponseTriggersExcludeRulesZarazElementVisibilityRuleAction string

const (
	ZarazHistoryReplaceResponseTriggersExcludeRulesZarazElementVisibilityRuleActionElementVisibility ZarazHistoryReplaceResponseTriggersExcludeRulesZarazElementVisibilityRuleAction = "elementVisibility"
)

type ZarazHistoryReplaceResponseTriggersExcludeRulesZarazElementVisibilityRuleSettings struct {
	Selector string                                                                                `json:"selector,required"`
	JSON     zarazHistoryReplaceResponseTriggersExcludeRulesZarazElementVisibilityRuleSettingsJSON `json:"-"`
}

// zarazHistoryReplaceResponseTriggersExcludeRulesZarazElementVisibilityRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazHistoryReplaceResponseTriggersExcludeRulesZarazElementVisibilityRuleSettings]
type zarazHistoryReplaceResponseTriggersExcludeRulesZarazElementVisibilityRuleSettingsJSON struct {
	Selector    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseTriggersExcludeRulesZarazElementVisibilityRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [ZarazHistoryReplaceResponseTriggersLoadRulesZarazLoadRule],
// [ZarazHistoryReplaceResponseTriggersLoadRulesZarazClickListenerRule],
// [ZarazHistoryReplaceResponseTriggersLoadRulesZarazTimerRule],
// [ZarazHistoryReplaceResponseTriggersLoadRulesZarazFormSubmissionRule],
// [ZarazHistoryReplaceResponseTriggersLoadRulesZarazVariableMatchRule],
// [ZarazHistoryReplaceResponseTriggersLoadRulesZarazScrollDepthRule] or
// [ZarazHistoryReplaceResponseTriggersLoadRulesZarazElementVisibilityRule].
type ZarazHistoryReplaceResponseTriggersLoadRule interface {
	implementsZarazHistoryReplaceResponseTriggersLoadRule()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZarazHistoryReplaceResponseTriggersLoadRule)(nil)).Elem(), "")
}

type ZarazHistoryReplaceResponseTriggersLoadRulesZarazLoadRule struct {
	ID    string                                                        `json:"id,required"`
	Match string                                                        `json:"match,required"`
	Op    ZarazHistoryReplaceResponseTriggersLoadRulesZarazLoadRuleOp   `json:"op,required"`
	Value string                                                        `json:"value,required"`
	JSON  zarazHistoryReplaceResponseTriggersLoadRulesZarazLoadRuleJSON `json:"-"`
}

// zarazHistoryReplaceResponseTriggersLoadRulesZarazLoadRuleJSON contains the JSON
// metadata for the struct
// [ZarazHistoryReplaceResponseTriggersLoadRulesZarazLoadRule]
type zarazHistoryReplaceResponseTriggersLoadRulesZarazLoadRuleJSON struct {
	ID          apijson.Field
	Match       apijson.Field
	Op          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseTriggersLoadRulesZarazLoadRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazHistoryReplaceResponseTriggersLoadRulesZarazLoadRule) implementsZarazHistoryReplaceResponseTriggersLoadRule() {
}

type ZarazHistoryReplaceResponseTriggersLoadRulesZarazLoadRuleOp string

const (
	ZarazHistoryReplaceResponseTriggersLoadRulesZarazLoadRuleOpContains           ZarazHistoryReplaceResponseTriggersLoadRulesZarazLoadRuleOp = "CONTAINS"
	ZarazHistoryReplaceResponseTriggersLoadRulesZarazLoadRuleOpEquals             ZarazHistoryReplaceResponseTriggersLoadRulesZarazLoadRuleOp = "EQUALS"
	ZarazHistoryReplaceResponseTriggersLoadRulesZarazLoadRuleOpStartsWith         ZarazHistoryReplaceResponseTriggersLoadRulesZarazLoadRuleOp = "STARTS_WITH"
	ZarazHistoryReplaceResponseTriggersLoadRulesZarazLoadRuleOpEndsWith           ZarazHistoryReplaceResponseTriggersLoadRulesZarazLoadRuleOp = "ENDS_WITH"
	ZarazHistoryReplaceResponseTriggersLoadRulesZarazLoadRuleOpMatchRegex         ZarazHistoryReplaceResponseTriggersLoadRulesZarazLoadRuleOp = "MATCH_REGEX"
	ZarazHistoryReplaceResponseTriggersLoadRulesZarazLoadRuleOpNotMatchRegex      ZarazHistoryReplaceResponseTriggersLoadRulesZarazLoadRuleOp = "NOT_MATCH_REGEX"
	ZarazHistoryReplaceResponseTriggersLoadRulesZarazLoadRuleOpGreaterThan        ZarazHistoryReplaceResponseTriggersLoadRulesZarazLoadRuleOp = "GREATER_THAN"
	ZarazHistoryReplaceResponseTriggersLoadRulesZarazLoadRuleOpGreaterThanOrEqual ZarazHistoryReplaceResponseTriggersLoadRulesZarazLoadRuleOp = "GREATER_THAN_OR_EQUAL"
	ZarazHistoryReplaceResponseTriggersLoadRulesZarazLoadRuleOpLessThan           ZarazHistoryReplaceResponseTriggersLoadRulesZarazLoadRuleOp = "LESS_THAN"
	ZarazHistoryReplaceResponseTriggersLoadRulesZarazLoadRuleOpLessThanOrEqual    ZarazHistoryReplaceResponseTriggersLoadRulesZarazLoadRuleOp = "LESS_THAN_OR_EQUAL"
)

type ZarazHistoryReplaceResponseTriggersLoadRulesZarazClickListenerRule struct {
	ID       string                                                                     `json:"id,required"`
	Action   ZarazHistoryReplaceResponseTriggersLoadRulesZarazClickListenerRuleAction   `json:"action,required"`
	Settings ZarazHistoryReplaceResponseTriggersLoadRulesZarazClickListenerRuleSettings `json:"settings,required"`
	JSON     zarazHistoryReplaceResponseTriggersLoadRulesZarazClickListenerRuleJSON     `json:"-"`
}

// zarazHistoryReplaceResponseTriggersLoadRulesZarazClickListenerRuleJSON contains
// the JSON metadata for the struct
// [ZarazHistoryReplaceResponseTriggersLoadRulesZarazClickListenerRule]
type zarazHistoryReplaceResponseTriggersLoadRulesZarazClickListenerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseTriggersLoadRulesZarazClickListenerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazHistoryReplaceResponseTriggersLoadRulesZarazClickListenerRule) implementsZarazHistoryReplaceResponseTriggersLoadRule() {
}

type ZarazHistoryReplaceResponseTriggersLoadRulesZarazClickListenerRuleAction string

const (
	ZarazHistoryReplaceResponseTriggersLoadRulesZarazClickListenerRuleActionClickListener ZarazHistoryReplaceResponseTriggersLoadRulesZarazClickListenerRuleAction = "clickListener"
)

type ZarazHistoryReplaceResponseTriggersLoadRulesZarazClickListenerRuleSettings struct {
	Selector    string                                                                         `json:"selector,required"`
	Type        ZarazHistoryReplaceResponseTriggersLoadRulesZarazClickListenerRuleSettingsType `json:"type,required"`
	WaitForTags int64                                                                          `json:"waitForTags,required"`
	JSON        zarazHistoryReplaceResponseTriggersLoadRulesZarazClickListenerRuleSettingsJSON `json:"-"`
}

// zarazHistoryReplaceResponseTriggersLoadRulesZarazClickListenerRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazHistoryReplaceResponseTriggersLoadRulesZarazClickListenerRuleSettings]
type zarazHistoryReplaceResponseTriggersLoadRulesZarazClickListenerRuleSettingsJSON struct {
	Selector    apijson.Field
	Type        apijson.Field
	WaitForTags apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseTriggersLoadRulesZarazClickListenerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryReplaceResponseTriggersLoadRulesZarazClickListenerRuleSettingsType string

const (
	ZarazHistoryReplaceResponseTriggersLoadRulesZarazClickListenerRuleSettingsTypeXpath ZarazHistoryReplaceResponseTriggersLoadRulesZarazClickListenerRuleSettingsType = "xpath"
	ZarazHistoryReplaceResponseTriggersLoadRulesZarazClickListenerRuleSettingsTypeCss   ZarazHistoryReplaceResponseTriggersLoadRulesZarazClickListenerRuleSettingsType = "css"
)

type ZarazHistoryReplaceResponseTriggersLoadRulesZarazTimerRule struct {
	ID       string                                                             `json:"id,required"`
	Action   ZarazHistoryReplaceResponseTriggersLoadRulesZarazTimerRuleAction   `json:"action,required"`
	Settings ZarazHistoryReplaceResponseTriggersLoadRulesZarazTimerRuleSettings `json:"settings,required"`
	JSON     zarazHistoryReplaceResponseTriggersLoadRulesZarazTimerRuleJSON     `json:"-"`
}

// zarazHistoryReplaceResponseTriggersLoadRulesZarazTimerRuleJSON contains the JSON
// metadata for the struct
// [ZarazHistoryReplaceResponseTriggersLoadRulesZarazTimerRule]
type zarazHistoryReplaceResponseTriggersLoadRulesZarazTimerRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseTriggersLoadRulesZarazTimerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazHistoryReplaceResponseTriggersLoadRulesZarazTimerRule) implementsZarazHistoryReplaceResponseTriggersLoadRule() {
}

type ZarazHistoryReplaceResponseTriggersLoadRulesZarazTimerRuleAction string

const (
	ZarazHistoryReplaceResponseTriggersLoadRulesZarazTimerRuleActionTimer ZarazHistoryReplaceResponseTriggersLoadRulesZarazTimerRuleAction = "timer"
)

type ZarazHistoryReplaceResponseTriggersLoadRulesZarazTimerRuleSettings struct {
	Interval int64                                                                  `json:"interval,required"`
	Limit    int64                                                                  `json:"limit,required"`
	JSON     zarazHistoryReplaceResponseTriggersLoadRulesZarazTimerRuleSettingsJSON `json:"-"`
}

// zarazHistoryReplaceResponseTriggersLoadRulesZarazTimerRuleSettingsJSON contains
// the JSON metadata for the struct
// [ZarazHistoryReplaceResponseTriggersLoadRulesZarazTimerRuleSettings]
type zarazHistoryReplaceResponseTriggersLoadRulesZarazTimerRuleSettingsJSON struct {
	Interval    apijson.Field
	Limit       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseTriggersLoadRulesZarazTimerRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryReplaceResponseTriggersLoadRulesZarazFormSubmissionRule struct {
	ID       string                                                                      `json:"id,required"`
	Action   ZarazHistoryReplaceResponseTriggersLoadRulesZarazFormSubmissionRuleAction   `json:"action,required"`
	Settings ZarazHistoryReplaceResponseTriggersLoadRulesZarazFormSubmissionRuleSettings `json:"settings,required"`
	JSON     zarazHistoryReplaceResponseTriggersLoadRulesZarazFormSubmissionRuleJSON     `json:"-"`
}

// zarazHistoryReplaceResponseTriggersLoadRulesZarazFormSubmissionRuleJSON contains
// the JSON metadata for the struct
// [ZarazHistoryReplaceResponseTriggersLoadRulesZarazFormSubmissionRule]
type zarazHistoryReplaceResponseTriggersLoadRulesZarazFormSubmissionRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseTriggersLoadRulesZarazFormSubmissionRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazHistoryReplaceResponseTriggersLoadRulesZarazFormSubmissionRule) implementsZarazHistoryReplaceResponseTriggersLoadRule() {
}

type ZarazHistoryReplaceResponseTriggersLoadRulesZarazFormSubmissionRuleAction string

const (
	ZarazHistoryReplaceResponseTriggersLoadRulesZarazFormSubmissionRuleActionFormSubmission ZarazHistoryReplaceResponseTriggersLoadRulesZarazFormSubmissionRuleAction = "formSubmission"
)

type ZarazHistoryReplaceResponseTriggersLoadRulesZarazFormSubmissionRuleSettings struct {
	Selector string                                                                          `json:"selector,required"`
	Validate bool                                                                            `json:"validate,required"`
	JSON     zarazHistoryReplaceResponseTriggersLoadRulesZarazFormSubmissionRuleSettingsJSON `json:"-"`
}

// zarazHistoryReplaceResponseTriggersLoadRulesZarazFormSubmissionRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazHistoryReplaceResponseTriggersLoadRulesZarazFormSubmissionRuleSettings]
type zarazHistoryReplaceResponseTriggersLoadRulesZarazFormSubmissionRuleSettingsJSON struct {
	Selector    apijson.Field
	Validate    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseTriggersLoadRulesZarazFormSubmissionRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryReplaceResponseTriggersLoadRulesZarazVariableMatchRule struct {
	ID       string                                                                     `json:"id,required"`
	Action   ZarazHistoryReplaceResponseTriggersLoadRulesZarazVariableMatchRuleAction   `json:"action,required"`
	Settings ZarazHistoryReplaceResponseTriggersLoadRulesZarazVariableMatchRuleSettings `json:"settings,required"`
	JSON     zarazHistoryReplaceResponseTriggersLoadRulesZarazVariableMatchRuleJSON     `json:"-"`
}

// zarazHistoryReplaceResponseTriggersLoadRulesZarazVariableMatchRuleJSON contains
// the JSON metadata for the struct
// [ZarazHistoryReplaceResponseTriggersLoadRulesZarazVariableMatchRule]
type zarazHistoryReplaceResponseTriggersLoadRulesZarazVariableMatchRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseTriggersLoadRulesZarazVariableMatchRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazHistoryReplaceResponseTriggersLoadRulesZarazVariableMatchRule) implementsZarazHistoryReplaceResponseTriggersLoadRule() {
}

type ZarazHistoryReplaceResponseTriggersLoadRulesZarazVariableMatchRuleAction string

const (
	ZarazHistoryReplaceResponseTriggersLoadRulesZarazVariableMatchRuleActionVariableMatch ZarazHistoryReplaceResponseTriggersLoadRulesZarazVariableMatchRuleAction = "variableMatch"
)

type ZarazHistoryReplaceResponseTriggersLoadRulesZarazVariableMatchRuleSettings struct {
	Match    string                                                                         `json:"match,required"`
	Variable string                                                                         `json:"variable,required"`
	JSON     zarazHistoryReplaceResponseTriggersLoadRulesZarazVariableMatchRuleSettingsJSON `json:"-"`
}

// zarazHistoryReplaceResponseTriggersLoadRulesZarazVariableMatchRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazHistoryReplaceResponseTriggersLoadRulesZarazVariableMatchRuleSettings]
type zarazHistoryReplaceResponseTriggersLoadRulesZarazVariableMatchRuleSettingsJSON struct {
	Match       apijson.Field
	Variable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseTriggersLoadRulesZarazVariableMatchRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryReplaceResponseTriggersLoadRulesZarazScrollDepthRule struct {
	ID       string                                                                   `json:"id,required"`
	Action   ZarazHistoryReplaceResponseTriggersLoadRulesZarazScrollDepthRuleAction   `json:"action,required"`
	Settings ZarazHistoryReplaceResponseTriggersLoadRulesZarazScrollDepthRuleSettings `json:"settings,required"`
	JSON     zarazHistoryReplaceResponseTriggersLoadRulesZarazScrollDepthRuleJSON     `json:"-"`
}

// zarazHistoryReplaceResponseTriggersLoadRulesZarazScrollDepthRuleJSON contains
// the JSON metadata for the struct
// [ZarazHistoryReplaceResponseTriggersLoadRulesZarazScrollDepthRule]
type zarazHistoryReplaceResponseTriggersLoadRulesZarazScrollDepthRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseTriggersLoadRulesZarazScrollDepthRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazHistoryReplaceResponseTriggersLoadRulesZarazScrollDepthRule) implementsZarazHistoryReplaceResponseTriggersLoadRule() {
}

type ZarazHistoryReplaceResponseTriggersLoadRulesZarazScrollDepthRuleAction string

const (
	ZarazHistoryReplaceResponseTriggersLoadRulesZarazScrollDepthRuleActionScrollDepth ZarazHistoryReplaceResponseTriggersLoadRulesZarazScrollDepthRuleAction = "scrollDepth"
)

type ZarazHistoryReplaceResponseTriggersLoadRulesZarazScrollDepthRuleSettings struct {
	Positions string                                                                       `json:"positions,required"`
	JSON      zarazHistoryReplaceResponseTriggersLoadRulesZarazScrollDepthRuleSettingsJSON `json:"-"`
}

// zarazHistoryReplaceResponseTriggersLoadRulesZarazScrollDepthRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazHistoryReplaceResponseTriggersLoadRulesZarazScrollDepthRuleSettings]
type zarazHistoryReplaceResponseTriggersLoadRulesZarazScrollDepthRuleSettingsJSON struct {
	Positions   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseTriggersLoadRulesZarazScrollDepthRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryReplaceResponseTriggersLoadRulesZarazElementVisibilityRule struct {
	ID       string                                                                         `json:"id,required"`
	Action   ZarazHistoryReplaceResponseTriggersLoadRulesZarazElementVisibilityRuleAction   `json:"action,required"`
	Settings ZarazHistoryReplaceResponseTriggersLoadRulesZarazElementVisibilityRuleSettings `json:"settings,required"`
	JSON     zarazHistoryReplaceResponseTriggersLoadRulesZarazElementVisibilityRuleJSON     `json:"-"`
}

// zarazHistoryReplaceResponseTriggersLoadRulesZarazElementVisibilityRuleJSON
// contains the JSON metadata for the struct
// [ZarazHistoryReplaceResponseTriggersLoadRulesZarazElementVisibilityRule]
type zarazHistoryReplaceResponseTriggersLoadRulesZarazElementVisibilityRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseTriggersLoadRulesZarazElementVisibilityRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazHistoryReplaceResponseTriggersLoadRulesZarazElementVisibilityRule) implementsZarazHistoryReplaceResponseTriggersLoadRule() {
}

type ZarazHistoryReplaceResponseTriggersLoadRulesZarazElementVisibilityRuleAction string

const (
	ZarazHistoryReplaceResponseTriggersLoadRulesZarazElementVisibilityRuleActionElementVisibility ZarazHistoryReplaceResponseTriggersLoadRulesZarazElementVisibilityRuleAction = "elementVisibility"
)

type ZarazHistoryReplaceResponseTriggersLoadRulesZarazElementVisibilityRuleSettings struct {
	Selector string                                                                             `json:"selector,required"`
	JSON     zarazHistoryReplaceResponseTriggersLoadRulesZarazElementVisibilityRuleSettingsJSON `json:"-"`
}

// zarazHistoryReplaceResponseTriggersLoadRulesZarazElementVisibilityRuleSettingsJSON
// contains the JSON metadata for the struct
// [ZarazHistoryReplaceResponseTriggersLoadRulesZarazElementVisibilityRuleSettings]
type zarazHistoryReplaceResponseTriggersLoadRulesZarazElementVisibilityRuleSettingsJSON struct {
	Selector    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseTriggersLoadRulesZarazElementVisibilityRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryReplaceResponseTriggersSystem string

const (
	ZarazHistoryReplaceResponseTriggersSystemPageload ZarazHistoryReplaceResponseTriggersSystem = "pageload"
)

// Union satisfied by [ZarazHistoryReplaceResponseVariablesObject] or
// [ZarazHistoryReplaceResponseVariablesObject].
type ZarazHistoryReplaceResponseVariable interface {
	implementsZarazHistoryReplaceResponseVariable()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZarazHistoryReplaceResponseVariable)(nil)).Elem(), "")
}

type ZarazHistoryReplaceResponseVariablesObject struct {
	Name  string                                         `json:"name,required"`
	Type  ZarazHistoryReplaceResponseVariablesObjectType `json:"type,required"`
	Value string                                         `json:"value,required"`
	JSON  zarazHistoryReplaceResponseVariablesObjectJSON `json:"-"`
}

// zarazHistoryReplaceResponseVariablesObjectJSON contains the JSON metadata for
// the struct [ZarazHistoryReplaceResponseVariablesObject]
type zarazHistoryReplaceResponseVariablesObjectJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseVariablesObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZarazHistoryReplaceResponseVariablesObject) implementsZarazHistoryReplaceResponseVariable() {}

type ZarazHistoryReplaceResponseVariablesObjectType string

const (
	ZarazHistoryReplaceResponseVariablesObjectTypeString ZarazHistoryReplaceResponseVariablesObjectType = "string"
	ZarazHistoryReplaceResponseVariablesObjectTypeSecret ZarazHistoryReplaceResponseVariablesObjectType = "secret"
)

// Consent management configuration.
type ZarazHistoryReplaceResponseConsent struct {
	Enabled                bool                                                     `json:"enabled,required"`
	ButtonTextTranslations ZarazHistoryReplaceResponseConsentButtonTextTranslations `json:"buttonTextTranslations"`
	CompanyEmail           string                                                   `json:"companyEmail"`
	CompanyName            string                                                   `json:"companyName"`
	CompanyStreetAddress   string                                                   `json:"companyStreetAddress"`
	ConsentModalIntroHTML  string                                                   `json:"consentModalIntroHTML"`
	// Object where keys are language codes
	ConsentModalIntroHTMLWithTranslations map[string]string `json:"consentModalIntroHTMLWithTranslations"`
	CookieName                            string            `json:"cookieName"`
	CustomCss                             string            `json:"customCSS"`
	CustomIntroDisclaimerDismissed        bool              `json:"customIntroDisclaimerDismissed"`
	DefaultLanguage                       string            `json:"defaultLanguage"`
	HideModal                             bool              `json:"hideModal"`
	// Object where keys are purpose alpha-numeric IDs
	Purposes map[string]ZarazHistoryReplaceResponseConsentPurpose `json:"purposes"`
	// Object where keys are purpose alpha-numeric IDs
	PurposesWithTranslations map[string]ZarazHistoryReplaceResponseConsentPurposesWithTranslation `json:"purposesWithTranslations"`
	JSON                     zarazHistoryReplaceResponseConsentJSON                               `json:"-"`
}

// zarazHistoryReplaceResponseConsentJSON contains the JSON metadata for the struct
// [ZarazHistoryReplaceResponseConsent]
type zarazHistoryReplaceResponseConsentJSON struct {
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

func (r *ZarazHistoryReplaceResponseConsent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryReplaceResponseConsentButtonTextTranslations struct {
	// Object where keys are language codes
	AcceptAll map[string]string `json:"accept_all,required"`
	// Object where keys are language codes
	ConfirmMyChoices map[string]string `json:"confirm_my_choices,required"`
	// Object where keys are language codes
	RejectAll map[string]string                                            `json:"reject_all,required"`
	JSON      zarazHistoryReplaceResponseConsentButtonTextTranslationsJSON `json:"-"`
}

// zarazHistoryReplaceResponseConsentButtonTextTranslationsJSON contains the JSON
// metadata for the struct
// [ZarazHistoryReplaceResponseConsentButtonTextTranslations]
type zarazHistoryReplaceResponseConsentButtonTextTranslationsJSON struct {
	AcceptAll        apijson.Field
	ConfirmMyChoices apijson.Field
	RejectAll        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseConsentButtonTextTranslations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryReplaceResponseConsentPurpose struct {
	Description string                                        `json:"description,required"`
	Name        string                                        `json:"name,required"`
	JSON        zarazHistoryReplaceResponseConsentPurposeJSON `json:"-"`
}

// zarazHistoryReplaceResponseConsentPurposeJSON contains the JSON metadata for the
// struct [ZarazHistoryReplaceResponseConsentPurpose]
type zarazHistoryReplaceResponseConsentPurposeJSON struct {
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseConsentPurpose) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryReplaceResponseConsentPurposesWithTranslation struct {
	// Object where keys are language codes
	Description map[string]string `json:"description,required"`
	// Object where keys are language codes
	Name  map[string]string                                             `json:"name,required"`
	Order int64                                                         `json:"order,required"`
	JSON  zarazHistoryReplaceResponseConsentPurposesWithTranslationJSON `json:"-"`
}

// zarazHistoryReplaceResponseConsentPurposesWithTranslationJSON contains the JSON
// metadata for the struct
// [ZarazHistoryReplaceResponseConsentPurposesWithTranslation]
type zarazHistoryReplaceResponseConsentPurposesWithTranslationJSON struct {
	Description apijson.Field
	Name        apijson.Field
	Order       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseConsentPurposesWithTranslation) UnmarshalJSON(data []byte) (err error) {
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

type ZarazHistoryReplaceParams struct {
	// ID of the Zaraz configuration to restore.
	Body param.Field[int64] `json:"body,required"`
}

func (r ZarazHistoryReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZarazHistoryReplaceResponseEnvelope struct {
	Errors   []ZarazHistoryReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZarazHistoryReplaceResponseEnvelopeMessages `json:"messages,required"`
	// Zaraz configuration
	Result ZarazHistoryReplaceResponse `json:"result,required"`
	// Whether the API call was successful
	Success bool                                    `json:"success,required"`
	JSON    zarazHistoryReplaceResponseEnvelopeJSON `json:"-"`
}

// zarazHistoryReplaceResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZarazHistoryReplaceResponseEnvelope]
type zarazHistoryReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryReplaceResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zarazHistoryReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// zarazHistoryReplaceResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZarazHistoryReplaceResponseEnvelopeErrors]
type zarazHistoryReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryReplaceResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zarazHistoryReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// zarazHistoryReplaceResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZarazHistoryReplaceResponseEnvelopeMessages]
type zarazHistoryReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
