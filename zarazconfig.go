// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
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
	path := fmt.Sprintf("zones/%s/settings/zaraz/config", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates Zaraz configuration for a zone.
func (r *ZarazConfigService) Update(ctx context.Context, zoneID string, body ZarazConfigUpdateParams, opts ...option.RequestOption) (res *ZarazConfigUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/zaraz/config", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ZarazConfigGetResponse struct {
	Errors   []ZarazConfigGetResponseError   `json:"errors"`
	Messages []ZarazConfigGetResponseMessage `json:"messages"`
	Result   ZarazConfigGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success bool                       `json:"success"`
	JSON    zarazConfigGetResponseJSON `json:"-"`
}

// zarazConfigGetResponseJSON contains the JSON metadata for the struct
// [ZarazConfigGetResponse]
type zarazConfigGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigGetResponseError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    zarazConfigGetResponseErrorJSON `json:"-"`
}

// zarazConfigGetResponseErrorJSON contains the JSON metadata for the struct
// [ZarazConfigGetResponseError]
type zarazConfigGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigGetResponseMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    zarazConfigGetResponseMessageJSON `json:"-"`
}

// zarazConfigGetResponseMessageJSON contains the JSON metadata for the struct
// [ZarazConfigGetResponseMessage]
type zarazConfigGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigGetResponseResult struct {
	// Consent management configuration.
	Consent ZarazConfigGetResponseResultConsent `json:"consent"`
	// Data layer compatibility mode enabled.
	DataLayer bool `json:"dataLayer"`
	// The key for Zaraz debug mode.
	DebugKey string `json:"debugKey"`
	// Single Page Application support enabled.
	HistoryChange bool `json:"historyChange"`
	// General Zaraz settings.
	Settings ZarazConfigGetResponseResultSettings `json:"settings"`
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
	ZarazVersion int64                            `json:"zarazVersion"`
	JSON         zarazConfigGetResponseResultJSON `json:"-"`
}

// zarazConfigGetResponseResultJSON contains the JSON metadata for the struct
// [ZarazConfigGetResponseResult]
type zarazConfigGetResponseResultJSON struct {
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

func (r *ZarazConfigGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Consent management configuration.
type ZarazConfigGetResponseResultConsent struct {
	Enabled                bool                                                      `json:"enabled,required"`
	ButtonTextTranslations ZarazConfigGetResponseResultConsentButtonTextTranslations `json:"buttonTextTranslations"`
	CompanyEmail           string                                                    `json:"companyEmail"`
	CompanyName            string                                                    `json:"companyName"`
	CompanyStreetAddress   string                                                    `json:"companyStreetAddress"`
	ConsentModalIntroHTML  string                                                    `json:"consentModalIntroHTML"`
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
	PurposesWithTranslations interface{}                             `json:"purposesWithTranslations"`
	JSON                     zarazConfigGetResponseResultConsentJSON `json:"-"`
}

// zarazConfigGetResponseResultConsentJSON contains the JSON metadata for the
// struct [ZarazConfigGetResponseResultConsent]
type zarazConfigGetResponseResultConsentJSON struct {
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

func (r *ZarazConfigGetResponseResultConsent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigGetResponseResultConsentButtonTextTranslations struct {
	// Object where keys are language codes
	AcceptAll interface{} `json:"accept_all,required"`
	// Object where keys are language codes
	ConfirmMyChoices interface{} `json:"confirm_my_choices,required"`
	// Object where keys are language codes
	RejectAll interface{}                                                   `json:"reject_all,required"`
	JSON      zarazConfigGetResponseResultConsentButtonTextTranslationsJSON `json:"-"`
}

// zarazConfigGetResponseResultConsentButtonTextTranslationsJSON contains the JSON
// metadata for the struct
// [ZarazConfigGetResponseResultConsentButtonTextTranslations]
type zarazConfigGetResponseResultConsentButtonTextTranslationsJSON struct {
	AcceptAll        apijson.Field
	ConfirmMyChoices apijson.Field
	RejectAll        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazConfigGetResponseResultConsentButtonTextTranslations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// General Zaraz settings.
type ZarazConfigGetResponseResultSettings struct {
	// Automatic injection of Zaraz scripts enabled.
	AutoInjectScript bool `json:"autoInjectScript,required"`
	// Details of the worker that receives and edits Zaraz Context object.
	ContextEnricher ZarazConfigGetResponseResultSettingsContextEnricher `json:"contextEnricher"`
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
	TrackPath string                                   `json:"trackPath"`
	JSON      zarazConfigGetResponseResultSettingsJSON `json:"-"`
}

// zarazConfigGetResponseResultSettingsJSON contains the JSON metadata for the
// struct [ZarazConfigGetResponseResultSettings]
type zarazConfigGetResponseResultSettingsJSON struct {
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

func (r *ZarazConfigGetResponseResultSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details of the worker that receives and edits Zaraz Context object.
type ZarazConfigGetResponseResultSettingsContextEnricher struct {
	EscapedWorkerName string                                                  `json:"escapedWorkerName,required"`
	WorkerTag         string                                                  `json:"workerTag,required"`
	JSON              zarazConfigGetResponseResultSettingsContextEnricherJSON `json:"-"`
}

// zarazConfigGetResponseResultSettingsContextEnricherJSON contains the JSON
// metadata for the struct [ZarazConfigGetResponseResultSettingsContextEnricher]
type zarazConfigGetResponseResultSettingsContextEnricherJSON struct {
	EscapedWorkerName apijson.Field
	WorkerTag         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZarazConfigGetResponseResultSettingsContextEnricher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigUpdateResponse struct {
	Errors   []ZarazConfigUpdateResponseError   `json:"errors"`
	Messages []ZarazConfigUpdateResponseMessage `json:"messages"`
	Result   ZarazConfigUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success bool                          `json:"success"`
	JSON    zarazConfigUpdateResponseJSON `json:"-"`
}

// zarazConfigUpdateResponseJSON contains the JSON metadata for the struct
// [ZarazConfigUpdateResponse]
type zarazConfigUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigUpdateResponseError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    zarazConfigUpdateResponseErrorJSON `json:"-"`
}

// zarazConfigUpdateResponseErrorJSON contains the JSON metadata for the struct
// [ZarazConfigUpdateResponseError]
type zarazConfigUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigUpdateResponseMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    zarazConfigUpdateResponseMessageJSON `json:"-"`
}

// zarazConfigUpdateResponseMessageJSON contains the JSON metadata for the struct
// [ZarazConfigUpdateResponseMessage]
type zarazConfigUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigUpdateResponseResult struct {
	// Consent management configuration.
	Consent ZarazConfigUpdateResponseResultConsent `json:"consent"`
	// Data layer compatibility mode enabled.
	DataLayer bool `json:"dataLayer"`
	// The key for Zaraz debug mode.
	DebugKey string `json:"debugKey"`
	// Single Page Application support enabled.
	HistoryChange bool `json:"historyChange"`
	// General Zaraz settings.
	Settings ZarazConfigUpdateResponseResultSettings `json:"settings"`
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
	ZarazVersion int64                               `json:"zarazVersion"`
	JSON         zarazConfigUpdateResponseResultJSON `json:"-"`
}

// zarazConfigUpdateResponseResultJSON contains the JSON metadata for the struct
// [ZarazConfigUpdateResponseResult]
type zarazConfigUpdateResponseResultJSON struct {
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

func (r *ZarazConfigUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Consent management configuration.
type ZarazConfigUpdateResponseResultConsent struct {
	Enabled                bool                                                         `json:"enabled,required"`
	ButtonTextTranslations ZarazConfigUpdateResponseResultConsentButtonTextTranslations `json:"buttonTextTranslations"`
	CompanyEmail           string                                                       `json:"companyEmail"`
	CompanyName            string                                                       `json:"companyName"`
	CompanyStreetAddress   string                                                       `json:"companyStreetAddress"`
	ConsentModalIntroHTML  string                                                       `json:"consentModalIntroHTML"`
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
	PurposesWithTranslations interface{}                                `json:"purposesWithTranslations"`
	JSON                     zarazConfigUpdateResponseResultConsentJSON `json:"-"`
}

// zarazConfigUpdateResponseResultConsentJSON contains the JSON metadata for the
// struct [ZarazConfigUpdateResponseResultConsent]
type zarazConfigUpdateResponseResultConsentJSON struct {
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

func (r *ZarazConfigUpdateResponseResultConsent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigUpdateResponseResultConsentButtonTextTranslations struct {
	// Object where keys are language codes
	AcceptAll interface{} `json:"accept_all,required"`
	// Object where keys are language codes
	ConfirmMyChoices interface{} `json:"confirm_my_choices,required"`
	// Object where keys are language codes
	RejectAll interface{}                                                      `json:"reject_all,required"`
	JSON      zarazConfigUpdateResponseResultConsentButtonTextTranslationsJSON `json:"-"`
}

// zarazConfigUpdateResponseResultConsentButtonTextTranslationsJSON contains the
// JSON metadata for the struct
// [ZarazConfigUpdateResponseResultConsentButtonTextTranslations]
type zarazConfigUpdateResponseResultConsentButtonTextTranslationsJSON struct {
	AcceptAll        apijson.Field
	ConfirmMyChoices apijson.Field
	RejectAll        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseResultConsentButtonTextTranslations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// General Zaraz settings.
type ZarazConfigUpdateResponseResultSettings struct {
	// Automatic injection of Zaraz scripts enabled.
	AutoInjectScript bool `json:"autoInjectScript,required"`
	// Details of the worker that receives and edits Zaraz Context object.
	ContextEnricher ZarazConfigUpdateResponseResultSettingsContextEnricher `json:"contextEnricher"`
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
	TrackPath string                                      `json:"trackPath"`
	JSON      zarazConfigUpdateResponseResultSettingsJSON `json:"-"`
}

// zarazConfigUpdateResponseResultSettingsJSON contains the JSON metadata for the
// struct [ZarazConfigUpdateResponseResultSettings]
type zarazConfigUpdateResponseResultSettingsJSON struct {
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

func (r *ZarazConfigUpdateResponseResultSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details of the worker that receives and edits Zaraz Context object.
type ZarazConfigUpdateResponseResultSettingsContextEnricher struct {
	EscapedWorkerName string                                                     `json:"escapedWorkerName,required"`
	WorkerTag         string                                                     `json:"workerTag,required"`
	JSON              zarazConfigUpdateResponseResultSettingsContextEnricherJSON `json:"-"`
}

// zarazConfigUpdateResponseResultSettingsContextEnricherJSON contains the JSON
// metadata for the struct [ZarazConfigUpdateResponseResultSettingsContextEnricher]
type zarazConfigUpdateResponseResultSettingsContextEnricherJSON struct {
	EscapedWorkerName apijson.Field
	WorkerTag         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZarazConfigUpdateResponseResultSettingsContextEnricher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigUpdateParams struct {
	// Consent management configuration.
	Consent param.Field[ZarazConfigUpdateParamsConsent] `json:"consent"`
	// Data layer compatibility mode enabled.
	DataLayer param.Field[bool] `json:"dataLayer"`
	// The key for Zaraz debug mode.
	DebugKey param.Field[string] `json:"debugKey"`
	// Single Page Application support enabled.
	HistoryChange param.Field[bool] `json:"historyChange"`
	// General Zaraz settings.
	Settings param.Field[ZarazConfigUpdateParamsSettings] `json:"settings"`
	// Tools set up under Zaraz configuration, where key is the alpha-numeric tool ID
	// and value is the tool configuration object.
	Tools param.Field[interface{}] `json:"tools"`
	// Triggers set up under Zaraz configuration, where key is the trigger
	// alpha-numeric ID and value is the trigger configuration.
	Triggers param.Field[interface{}] `json:"triggers"`
	// Variables set up under Zaraz configuration, where key is the variable
	// alpha-numeric ID and value is the variable configuration. Values of variables of
	// type secret are not included.
	Variables param.Field[interface{}] `json:"variables"`
	// Zaraz internal version of the config.
	ZarazVersion param.Field[int64] `json:"zarazVersion"`
}

func (r ZarazConfigUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Consent management configuration.
type ZarazConfigUpdateParamsConsent struct {
	Enabled                param.Field[bool]                                                 `json:"enabled,required"`
	ButtonTextTranslations param.Field[ZarazConfigUpdateParamsConsentButtonTextTranslations] `json:"buttonTextTranslations"`
	CompanyEmail           param.Field[string]                                               `json:"companyEmail"`
	CompanyName            param.Field[string]                                               `json:"companyName"`
	CompanyStreetAddress   param.Field[string]                                               `json:"companyStreetAddress"`
	ConsentModalIntroHTML  param.Field[string]                                               `json:"consentModalIntroHTML"`
	// Object where keys are language codes
	ConsentModalIntroHTMLWithTranslations param.Field[interface{}] `json:"consentModalIntroHTMLWithTranslations"`
	CookieName                            param.Field[string]      `json:"cookieName"`
	CustomCss                             param.Field[string]      `json:"customCSS"`
	CustomIntroDisclaimerDismissed        param.Field[bool]        `json:"customIntroDisclaimerDismissed"`
	DefaultLanguage                       param.Field[string]      `json:"defaultLanguage"`
	HideModal                             param.Field[bool]        `json:"hideModal"`
	// Object where keys are purpose alpha-numeric IDs
	Purposes param.Field[interface{}] `json:"purposes"`
	// Object where keys are purpose alpha-numeric IDs
	PurposesWithTranslations param.Field[interface{}] `json:"purposesWithTranslations"`
}

func (r ZarazConfigUpdateParamsConsent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZarazConfigUpdateParamsConsentButtonTextTranslations struct {
	// Object where keys are language codes
	AcceptAll param.Field[interface{}] `json:"accept_all,required"`
	// Object where keys are language codes
	ConfirmMyChoices param.Field[interface{}] `json:"confirm_my_choices,required"`
	// Object where keys are language codes
	RejectAll param.Field[interface{}] `json:"reject_all,required"`
}

func (r ZarazConfigUpdateParamsConsentButtonTextTranslations) MarshalJSON() (data []byte, err error) {
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
