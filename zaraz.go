package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

type ZarazConfig struct {
	DebugKey      string              `json:"debugKey"`
	Tools         map[string]Tool     `json:"tools"`
	Triggers      map[string]Trigger  `json:"triggers"`
	ZarazVersion  int64               `json:"zarazVersion"`
	Consent       Consent             `json:"consent,omitempty"`
	DataLayer     bool                `json:"dataLayer,omitempty"`
	Dlp           []any               `json:"dlp,omitempty"`
	HistoryChange bool                `json:"historyChange,omitempty"`
	Settings      ConfigSettings      `json:"settings,omitempty"`
	Variables     map[string]Variable `json:"variables,omitempty"`
}

type Worker struct {
	EscapedWorkerName string `json:"escapedWorkerName"`
	WorkerTag         string `json:"workerTag"`
	MutableId         string `json:"mutableId,omitempty"`
}

type ConfigSettings struct {
	AutoInjectScript    bool   `json:"autoInjectScript"`
	InjectIframes       bool   `json:"injectIframes,omitempty"`
	Ecommerce           bool   `json:"ecommerce,omitempty"`
	HideQueryParams     bool   `json:"hideQueryParams,omitempty"`
	HideIpAddress       bool   `json:"hideIPAddress,omitempty"`
	HideUserAgent       bool   `json:"hideUserAgent,omitempty"`
	HideExternalReferer bool   `json:"hideExternalReferer,omitempty"`
	CookieDomain        string `json:"cookieDomain,omitempty"`
	InitPath            string `json:"initPath,omitempty"`
	ScriptPath          string `json:"scriptPath,omitempty"`
	TrackPath           string `json:"trackPath,omitempty"`
	EventsApiPath       string `json:"eventsApiPath,omitempty"`
	McRootPath          string `json:"mcRootPath,omitempty"`
	ContextEnricher     Worker `json:"contextEnricher,omitempty"`
}

type NeoEvent struct {
	BlockingTriggers []string       `json:"blockingTriggers"`
	FiringTriggers   []string       `json:"firingTriggers"`
	Data             map[string]any `json:"data"`
	ActionType       string         `json:"actionType,omitempty"`
}

type ToolType string

const (
	ToolLibrary   ToolType = "library"
	ToolComponent ToolType = "component"
	ToolCustomMc  ToolType = "custom-mc"
)

type Tool struct {
	BlockingTriggers []string       `json:"blockingTriggers"`
	Enabled          bool           `json:"enabled"`
	DefaultFields    map[string]any `json:"defaultFields"`
	Name             string         `json:"name"`
	NeoEvents        []NeoEvent     `json:"neoEvents"`
	Type             ToolType       `json:"type"`
	DefaultPurpose   string         `json:"defaultPurpose,omitempty"`
	Library          string         `json:"library,omitempty"`
	Component        string         `json:"component,omitempty"`
	Permissions      []string       `json:"permissions,omitempty"`
	Settings         map[string]any `json:"settings,omitempty"`
	Worker           Worker         `json:"worker,omitempty"`
}

type TriggerSystem string

const Pageload TriggerSystem = "pageload"

type LoadRuleOp string

type RuleType string

const (
	ClickListener     RuleType = "clickListener"
	Timer             RuleType = "timer"
	FormSubmission    RuleType = "formSubmission"
	VariableMatch     RuleType = "variableMatch"
	ScrollDepth       RuleType = "scrollDepth"
	ElementVisibility RuleType = "elementVisibility"
	ClientEval        RuleType = "clientEval"
)

type SelectorType string

const (
	Xpath SelectorType = "xpath"
	Css   SelectorType = "css"
)

type RuleSettings struct {
	Type        SelectorType `json:"type,omitempty"`
	Selector    string       `json:"selector,omitempty"`
	WaitForTags int          `json:"waitForTags,omitempty"`
	Interval    int          `json:"interval,omitempty"`
	Limit       int          `json:"limit,omitempty"`
	Validate    bool         `json:"validate,omitempty"`
	Variable    string       `json:"variable,omitempty"`
	Match       string       `json:"match,omitempty"`
	Positions   string       `json:"positions,omitempty"`
	Op          LoadRuleOp   `json:"op,omitempty"`
	Value       string       `json:"value,omitempty"`
}

type TriggerRule struct {
	Id       string       `json:"id"`
	Match    string       `json:"match,omitempty"`
	Op       LoadRuleOp   `json:"op,omitempty"`
	Value    string       `json:"value,omitempty"`
	Action   RuleType     `json:"action"`
	Settings RuleSettings `json:"settings"`
}

type Trigger struct {
	Name         string        `json:"name"`
	Description  string        `json:"description,omitempty"`
	LoadRules    []TriggerRule `json:"loadRules"`
	ExcludeRules []TriggerRule `json:"excludeRules"`
	ClientRules  []any         `json:"clientRules,omitempty"` // what is this?
	System       TriggerSystem `json:"system,omitempty"`
}

type VariableType string

const (
	VarString VariableType = "string"
	VarSecret VariableType = "secret"
	VarWorker VariableType = "worker"
)

type Variable struct {
	Name  string       `json:"name"`
	Type  VariableType `json:"type"`
	Value interface{}  `json:"value"`
}

type ButtonTextTranslations struct {
	AcceptAll        map[string]string `json:"accept_all"`
	RejectAll        map[string]string `json:"reject_all"`
	ConfirmMyChoices map[string]string `json:"confirm_my_choices"`
}

type Purpose struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type PurposeWithTranslations struct {
	Name        map[string]string `json:"name"`
	Description map[string]string `json:"description"`
	Order       int               `json:"order"`
}

type Consent struct {
	Enabled                               bool                               `json:"enabled"`
	ButtonTextTranslations                ButtonTextTranslations             `json:"buttonTextTranslations,omitempty"`
	CompanyEmail                          string                             `json:"companyEmail,omitempty"`
	CompanyName                           string                             `json:"companyName,omitempty"`
	CompanyStreetAddress                  string                             `json:"companyStreetAddress,omitempty"`
	ConsentModalIntroHTML                 string                             `json:"consentModalIntroHTML,omitempty"`
	ConsentModalIntroHTMLWithTranslations map[string]string                  `json:"consentModalIntroHTMLWithTranslations,omitempty"`
	CookieName                            string                             `json:"cookieName,omitempty"`
	CustomCSS                             string                             `json:"customCSS,omitempty"`
	CustomIntroDisclaimerDismissed        bool                               `json:"customIntroDisclaimerDismissed,omitempty"`
	DefaultLanguage                       string                             `json:"defaultLanguage,omitempty"`
	HideModal                             bool                               `json:"hideModal,omitempty"`
	Purposes                              map[string]Purpose                 `json:"purposes,omitempty"`
	PurposesWithTranslations              map[string]PurposeWithTranslations `json:"purposesWithTranslations,omitempty"`
}

type ZarazConfigResponse struct {
	Result ZarazConfig `json:"result"`
	Response
}

type ZarazWorkflowResponse struct {
	Result string `json:"result"`
	Response
}

type ZarazPublishResponse struct {
	Result string `json:"result"`
	Response
}

type UpdateZarazConfigParams = ZarazConfig

type UpdateZarazWorkflowParams = string

type PublishZarazConfigParams = string

type ZarazHistoryRow struct {
	ID          int64      `json:"id,omitempty"`
	UserId      string     `json:"usedId,omitempty"`
	Description string     `json:"description,omitempty"`
	CreatedAt   *time.Time `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
}

type GetZarazConfigHistoryListResponse struct {
	Result []ZarazHistoryRow `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

type GetZarazConfigHistoryListParams struct {
	ResultInfo
}

type GetZarazConfigsByIdResponse = map[string]interface{}

// listZarazConfigHistoryDefaultPageSize represents the default per_page size of the API.
var listZarazConfigHistoryDefaultPageSize int = 100

func (api *API) GetZarazConfig(ctx context.Context, rc *ResourceContainer) (ZarazConfigResponse, error) {
	if rc.Identifier == "" {
		return ZarazConfigResponse{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/settings/zaraz/v2/config", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ZarazConfigResponse{}, err
	}

	var recordResp ZarazConfigResponse
	err = json.Unmarshal(res, &recordResp)
	if err != nil {
		return ZarazConfigResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return recordResp, nil
}

func (api *API) UpdateZarazConfig(ctx context.Context, rc *ResourceContainer, params UpdateZarazConfigParams) (ZarazConfigResponse, error) {
	if rc.Identifier == "" {
		return ZarazConfigResponse{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/settings/zaraz/v2/config", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return ZarazConfigResponse{}, err
	}

	var updateResp ZarazConfigResponse
	err = json.Unmarshal(res, &updateResp)
	if err != nil {
		return ZarazConfigResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return updateResp, nil
}

func (api *API) GetZarazWorkflow(ctx context.Context, rc *ResourceContainer) (ZarazWorkflowResponse, error) {
	if rc.Identifier == "" {
		return ZarazWorkflowResponse{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/settings/zaraz/v2/workflow", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ZarazWorkflowResponse{}, err
	}

	var response ZarazWorkflowResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return ZarazWorkflowResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response, nil
}

func (api *API) UpdateZarazWorkflow(ctx context.Context, rc *ResourceContainer, params UpdateZarazWorkflowParams) (ZarazWorkflowResponse, error) {
	if rc.Identifier == "" {
		return ZarazWorkflowResponse{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/settings/zaraz/v2/workflow", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return ZarazWorkflowResponse{}, err
	}

	var response ZarazWorkflowResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return ZarazWorkflowResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response, nil
}

func (api *API) PublishZarazConfig(ctx context.Context, rc *ResourceContainer, params PublishZarazConfigParams) (ZarazPublishResponse, error) {
	if rc.Identifier == "" {
		return ZarazPublishResponse{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/settings/zaraz/v2/publish", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return ZarazPublishResponse{}, err
	}

	var response ZarazPublishResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return ZarazPublishResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response, nil
}

func (api *API) GetZarazConfigHistoryList(ctx context.Context, rc *ResourceContainer, params GetZarazConfigHistoryListParams) ([]ZarazHistoryRow, *ResultInfo, error) {
	if rc.Identifier == "" {
		return nil, nil, ErrMissingZoneID
	}

	autoPaginate := true
	if params.PerPage >= 1 || params.Page >= 1 {
		autoPaginate = false
	}

	if params.PerPage < 1 {
		params.PerPage = listZarazConfigHistoryDefaultPageSize
	}

	if params.Page < 1 {
		params.Page = 1
	}

	var records []ZarazHistoryRow
	var lastResultInfo ResultInfo

	for {
		uri := buildURI(fmt.Sprintf("/zones/%s/settings/zaraz/v2/history", rc.Identifier), params)
		res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
		if err != nil {
			return []ZarazHistoryRow{}, &ResultInfo{}, err
		}
		var listResponse GetZarazConfigHistoryListResponse
		err = json.Unmarshal(res, &listResponse)
		if err != nil {
			return []ZarazHistoryRow{}, &ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
		}
		records = append(records, listResponse.Result...)
		lastResultInfo = listResponse.ResultInfo
		params.ResultInfo = listResponse.ResultInfo.Next()
		if params.ResultInfo.Done() || !autoPaginate {
			break
		}
	}
	return records, &lastResultInfo, nil
}

func (api *API) GetDefaultZarazConfig(ctx context.Context, rc *ResourceContainer) (ZarazConfigResponse, error) {
	if rc.Identifier == "" {
		return ZarazConfigResponse{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/settings/zaraz/v2/default", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ZarazConfigResponse{}, err
	}

	var recordResp ZarazConfigResponse
	err = json.Unmarshal(res, &recordResp)
	if err != nil {
		return ZarazConfigResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return recordResp, nil
}

func (api *API) ExportZarazConfig(ctx context.Context, rc *ResourceContainer) error {
	if rc.Identifier == "" {
		return ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/settings/zaraz/v2/export", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return err
	}

	var recordResp ZarazConfig
	err = json.Unmarshal(res, &recordResp)
	if err != nil {
		return fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return nil
}
