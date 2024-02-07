// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
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
	path := fmt.Sprintf("zones/%s/settings/zaraz/history", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists a history of published Zaraz configuration records for a zone.
func (r *ZarazHistoryService) List(ctx context.Context, zoneID string, query ZarazHistoryListParams, opts ...option.RequestOption) (res *ZarazHistoryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/zaraz/history", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ZarazHistoryUpdateResponse struct {
	Errors   []ZarazHistoryUpdateResponseError   `json:"errors"`
	Messages []ZarazHistoryUpdateResponseMessage `json:"messages"`
	Result   ZarazHistoryUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success bool                           `json:"success"`
	JSON    zarazHistoryUpdateResponseJSON `json:"-"`
}

// zarazHistoryUpdateResponseJSON contains the JSON metadata for the struct
// [ZarazHistoryUpdateResponse]
type zarazHistoryUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryUpdateResponseError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    zarazHistoryUpdateResponseErrorJSON `json:"-"`
}

// zarazHistoryUpdateResponseErrorJSON contains the JSON metadata for the struct
// [ZarazHistoryUpdateResponseError]
type zarazHistoryUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryUpdateResponseMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    zarazHistoryUpdateResponseMessageJSON `json:"-"`
}

// zarazHistoryUpdateResponseMessageJSON contains the JSON metadata for the struct
// [ZarazHistoryUpdateResponseMessage]
type zarazHistoryUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryUpdateResponseResult struct {
	// Consent management configuration.
	Consent ZarazHistoryUpdateResponseResultConsent `json:"consent"`
	// Data layer compatibility mode enabled.
	DataLayer bool `json:"dataLayer"`
	// The key for Zaraz debug mode.
	DebugKey string `json:"debugKey"`
	// Single Page Application support enabled.
	HistoryChange bool `json:"historyChange"`
	// General Zaraz settings.
	Settings ZarazHistoryUpdateResponseResultSettings `json:"settings"`
	// Tools set up under Zaraz configuration, where key is the alpha-numeric tool ID
	// and value is the tool configuration object.
	Tools interface{} `json:"tools"`
	// Triggers set up under Zaraz configuration, where key is the trigger
	// alpha-numeric ID and value is the trigger configuration.
	Triggers interface{} `json:"triggers"`
	// Variables set up under Zaraz configuration, where key is the variable
	// alpha-numeric ID and value is the variable configuration. Values of variables of
	// type secret are not included.
	Variables interface{} `json:"variables"`
	// Zaraz internal version of the config.
	ZarazVersion int64                                `json:"zarazVersion"`
	JSON         zarazHistoryUpdateResponseResultJSON `json:"-"`
}

// zarazHistoryUpdateResponseResultJSON contains the JSON metadata for the struct
// [ZarazHistoryUpdateResponseResult]
type zarazHistoryUpdateResponseResultJSON struct {
	Consent       apijson.Field
	DataLayer     apijson.Field
	DebugKey      apijson.Field
	HistoryChange apijson.Field
	Settings      apijson.Field
	Tools         apijson.Field
	Triggers      apijson.Field
	Variables     apijson.Field
	ZarazVersion  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Consent management configuration.
type ZarazHistoryUpdateResponseResultConsent struct {
	Enabled                bool                                                          `json:"enabled,required"`
	ButtonTextTranslations ZarazHistoryUpdateResponseResultConsentButtonTextTranslations `json:"buttonTextTranslations"`
	CompanyEmail           string                                                        `json:"companyEmail"`
	CompanyName            string                                                        `json:"companyName"`
	CompanyStreetAddress   string                                                        `json:"companyStreetAddress"`
	ConsentModalIntroHTML  string                                                        `json:"consentModalIntroHTML"`
	// Object where keys are language codes
	ConsentModalIntroHTMLWithTranslations interface{} `json:"consentModalIntroHTMLWithTranslations"`
	CookieName                            string      `json:"cookieName"`
	CustomCss                             string      `json:"customCSS"`
	CustomIntroDisclaimerDismissed        bool        `json:"customIntroDisclaimerDismissed"`
	DefaultLanguage                       string      `json:"defaultLanguage"`
	HideModal                             bool        `json:"hideModal"`
	// Object where keys are purpose alpha-numeric IDs
	Purposes interface{} `json:"purposes"`
	// Object where keys are purpose alpha-numeric IDs
	PurposesWithTranslations interface{}                                 `json:"purposesWithTranslations"`
	JSON                     zarazHistoryUpdateResponseResultConsentJSON `json:"-"`
}

// zarazHistoryUpdateResponseResultConsentJSON contains the JSON metadata for the
// struct [ZarazHistoryUpdateResponseResultConsent]
type zarazHistoryUpdateResponseResultConsentJSON struct {
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

func (r *ZarazHistoryUpdateResponseResultConsent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryUpdateResponseResultConsentButtonTextTranslations struct {
	// Object where keys are language codes
	AcceptAll interface{} `json:"accept_all,required"`
	// Object where keys are language codes
	ConfirmMyChoices interface{} `json:"confirm_my_choices,required"`
	// Object where keys are language codes
	RejectAll interface{}                                                       `json:"reject_all,required"`
	JSON      zarazHistoryUpdateResponseResultConsentButtonTextTranslationsJSON `json:"-"`
}

// zarazHistoryUpdateResponseResultConsentButtonTextTranslationsJSON contains the
// JSON metadata for the struct
// [ZarazHistoryUpdateResponseResultConsentButtonTextTranslations]
type zarazHistoryUpdateResponseResultConsentButtonTextTranslationsJSON struct {
	AcceptAll        apijson.Field
	ConfirmMyChoices apijson.Field
	RejectAll        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseResultConsentButtonTextTranslations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// General Zaraz settings.
type ZarazHistoryUpdateResponseResultSettings struct {
	// Automatic injection of Zaraz scripts enabled.
	AutoInjectScript bool `json:"autoInjectScript,required"`
	// Details of the worker that receives and edits Zaraz Context object.
	ContextEnricher ZarazHistoryUpdateResponseResultSettingsContextEnricher `json:"contextEnricher"`
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
	TrackPath string                                       `json:"trackPath"`
	JSON      zarazHistoryUpdateResponseResultSettingsJSON `json:"-"`
}

// zarazHistoryUpdateResponseResultSettingsJSON contains the JSON metadata for the
// struct [ZarazHistoryUpdateResponseResultSettings]
type zarazHistoryUpdateResponseResultSettingsJSON struct {
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

func (r *ZarazHistoryUpdateResponseResultSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details of the worker that receives and edits Zaraz Context object.
type ZarazHistoryUpdateResponseResultSettingsContextEnricher struct {
	EscapedWorkerName string                                                      `json:"escapedWorkerName,required"`
	WorkerTag         string                                                      `json:"workerTag,required"`
	JSON              zarazHistoryUpdateResponseResultSettingsContextEnricherJSON `json:"-"`
}

// zarazHistoryUpdateResponseResultSettingsContextEnricherJSON contains the JSON
// metadata for the struct
// [ZarazHistoryUpdateResponseResultSettingsContextEnricher]
type zarazHistoryUpdateResponseResultSettingsContextEnricherJSON struct {
	EscapedWorkerName apijson.Field
	WorkerTag         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZarazHistoryUpdateResponseResultSettingsContextEnricher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryListResponse struct {
	Errors   []ZarazHistoryListResponseError   `json:"errors"`
	Messages []ZarazHistoryListResponseMessage `json:"messages"`
	Result   []ZarazHistoryListResponseResult  `json:"result"`
	// Whether the API call was successful
	Success bool                         `json:"success"`
	JSON    zarazHistoryListResponseJSON `json:"-"`
}

// zarazHistoryListResponseJSON contains the JSON metadata for the struct
// [ZarazHistoryListResponse]
type zarazHistoryListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryListResponseError struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    zarazHistoryListResponseErrorJSON `json:"-"`
}

// zarazHistoryListResponseErrorJSON contains the JSON metadata for the struct
// [ZarazHistoryListResponseError]
type zarazHistoryListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryListResponseMessage struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    zarazHistoryListResponseMessageJSON `json:"-"`
}

// zarazHistoryListResponseMessageJSON contains the JSON metadata for the struct
// [ZarazHistoryListResponseMessage]
type zarazHistoryListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryListResponseResult struct {
	// ID of the configuration
	ID int64 `json:"id"`
	// Date and time the configuration was created
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Configuration description provided by the user who published this configuration
	Description string `json:"description"`
	// Date and time the configuration was last updated
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// Alpha-numeric ID of the account user who published the configuration
	UserID string                             `json:"userId"`
	JSON   zarazHistoryListResponseResultJSON `json:"-"`
}

// zarazHistoryListResponseResultJSON contains the JSON metadata for the struct
// [ZarazHistoryListResponseResult]
type zarazHistoryListResponseResultJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	UpdatedAt   apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazHistoryListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazHistoryUpdateParams struct {
	// ID of the Zaraz configuration to restore.
	Body param.Field[int64] `json:"body,required"`
}

func (r ZarazHistoryUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
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
