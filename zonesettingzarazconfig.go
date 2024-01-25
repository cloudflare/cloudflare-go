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

// ZoneSettingZarazConfigService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingZarazConfigService]
// method instead.
type ZoneSettingZarazConfigService struct {
	Options []option.RequestOption
}

// NewZoneSettingZarazConfigService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingZarazConfigService(opts ...option.RequestOption) (r *ZoneSettingZarazConfigService) {
	r = &ZoneSettingZarazConfigService{}
	r.Options = opts
	return
}

// Gets latest Zaraz configuration for a zone. It can be preview or published
// configuration, whichever was the last updated. Secret variables values will not
// be included.
func (r *ZoneSettingZarazConfigService) Get(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZarazConfigResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/zaraz/v2/config", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates Zaraz configuration for a zone.
func (r *ZoneSettingZarazConfigService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingZarazConfigUpdateParams, opts ...option.RequestOption) (res *ZarazConfigResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/zaraz/v2/config", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ZarazConfigResponse struct {
	Errors   []ZarazConfigResponseError   `json:"errors"`
	Messages []ZarazConfigResponseMessage `json:"messages"`
	Result   ZarazConfigResponseResult    `json:"result"`
	// Whether the API call was successful
	Success bool                    `json:"success"`
	JSON    zarazConfigResponseJSON `json:"-"`
}

// zarazConfigResponseJSON contains the JSON metadata for the struct
// [ZarazConfigResponse]
type zarazConfigResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigResponseError struct {
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    zarazConfigResponseErrorJSON `json:"-"`
}

// zarazConfigResponseErrorJSON contains the JSON metadata for the struct
// [ZarazConfigResponseError]
type zarazConfigResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigResponseMessage struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    zarazConfigResponseMessageJSON `json:"-"`
}

// zarazConfigResponseMessageJSON contains the JSON metadata for the struct
// [ZarazConfigResponseMessage]
type zarazConfigResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazConfigResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigResponseResult struct {
	// Consent management configuration.
	Consent ZarazConfigResponseResultConsent `json:"consent"`
	// Data layer compatibility mode enabled.
	DataLayer bool `json:"dataLayer"`
	// The key for Zaraz debug mode.
	DebugKey string `json:"debugKey"`
	// Single Page Application support enabled.
	HistoryChange bool `json:"historyChange"`
	// General Zaraz settings.
	Settings ZarazConfigResponseResultSettings `json:"settings"`
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
	ZarazVersion int64                         `json:"zarazVersion"`
	JSON         zarazConfigResponseResultJSON `json:"-"`
}

// zarazConfigResponseResultJSON contains the JSON metadata for the struct
// [ZarazConfigResponseResult]
type zarazConfigResponseResultJSON struct {
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

func (r *ZarazConfigResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Consent management configuration.
type ZarazConfigResponseResultConsent struct {
	Enabled                bool                                                   `json:"enabled,required"`
	ButtonTextTranslations ZarazConfigResponseResultConsentButtonTextTranslations `json:"buttonTextTranslations"`
	CompanyEmail           string                                                 `json:"companyEmail"`
	CompanyName            string                                                 `json:"companyName"`
	CompanyStreetAddress   string                                                 `json:"companyStreetAddress"`
	ConsentModalIntroHTML  string                                                 `json:"consentModalIntroHTML"`
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
	PurposesWithTranslations interface{}                          `json:"purposesWithTranslations"`
	JSON                     zarazConfigResponseResultConsentJSON `json:"-"`
}

// zarazConfigResponseResultConsentJSON contains the JSON metadata for the struct
// [ZarazConfigResponseResultConsent]
type zarazConfigResponseResultConsentJSON struct {
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

func (r *ZarazConfigResponseResultConsent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazConfigResponseResultConsentButtonTextTranslations struct {
	// Object where keys are language codes
	AcceptAll interface{} `json:"accept_all,required"`
	// Object where keys are language codes
	ConfirmMyChoices interface{} `json:"confirm_my_choices,required"`
	// Object where keys are language codes
	RejectAll interface{}                                                `json:"reject_all,required"`
	JSON      zarazConfigResponseResultConsentButtonTextTranslationsJSON `json:"-"`
}

// zarazConfigResponseResultConsentButtonTextTranslationsJSON contains the JSON
// metadata for the struct [ZarazConfigResponseResultConsentButtonTextTranslations]
type zarazConfigResponseResultConsentButtonTextTranslationsJSON struct {
	AcceptAll        apijson.Field
	ConfirmMyChoices apijson.Field
	RejectAll        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZarazConfigResponseResultConsentButtonTextTranslations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// General Zaraz settings.
type ZarazConfigResponseResultSettings struct {
	// Automatic injection of Zaraz scripts enabled.
	AutoInjectScript bool `json:"autoInjectScript,required"`
	// Details of the worker that receives and edits Zaraz Context object.
	ContextEnricher ZarazConfigResponseResultSettingsContextEnricher `json:"contextEnricher"`
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
	JSON      zarazConfigResponseResultSettingsJSON `json:"-"`
}

// zarazConfigResponseResultSettingsJSON contains the JSON metadata for the struct
// [ZarazConfigResponseResultSettings]
type zarazConfigResponseResultSettingsJSON struct {
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

func (r *ZarazConfigResponseResultSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Details of the worker that receives and edits Zaraz Context object.
type ZarazConfigResponseResultSettingsContextEnricher struct {
	EscapedWorkerName string                                               `json:"escapedWorkerName,required"`
	WorkerTag         string                                               `json:"workerTag,required"`
	JSON              zarazConfigResponseResultSettingsContextEnricherJSON `json:"-"`
}

// zarazConfigResponseResultSettingsContextEnricherJSON contains the JSON metadata
// for the struct [ZarazConfigResponseResultSettingsContextEnricher]
type zarazConfigResponseResultSettingsContextEnricherJSON struct {
	EscapedWorkerName apijson.Field
	WorkerTag         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZarazConfigResponseResultSettingsContextEnricher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingZarazConfigUpdateParams struct {
	// Consent management configuration.
	Consent param.Field[ZoneSettingZarazConfigUpdateParamsConsent] `json:"consent"`
	// Data layer compatibility mode enabled.
	DataLayer param.Field[bool] `json:"dataLayer"`
	// The key for Zaraz debug mode.
	DebugKey param.Field[string] `json:"debugKey"`
	// Single Page Application support enabled.
	HistoryChange param.Field[bool] `json:"historyChange"`
	// General Zaraz settings.
	Settings param.Field[ZoneSettingZarazConfigUpdateParamsSettings] `json:"settings"`
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

func (r ZoneSettingZarazConfigUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Consent management configuration.
type ZoneSettingZarazConfigUpdateParamsConsent struct {
	Enabled                param.Field[bool]                                                            `json:"enabled,required"`
	ButtonTextTranslations param.Field[ZoneSettingZarazConfigUpdateParamsConsentButtonTextTranslations] `json:"buttonTextTranslations"`
	CompanyEmail           param.Field[string]                                                          `json:"companyEmail"`
	CompanyName            param.Field[string]                                                          `json:"companyName"`
	CompanyStreetAddress   param.Field[string]                                                          `json:"companyStreetAddress"`
	ConsentModalIntroHTML  param.Field[string]                                                          `json:"consentModalIntroHTML"`
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

func (r ZoneSettingZarazConfigUpdateParamsConsent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingZarazConfigUpdateParamsConsentButtonTextTranslations struct {
	// Object where keys are language codes
	AcceptAll param.Field[interface{}] `json:"accept_all,required"`
	// Object where keys are language codes
	ConfirmMyChoices param.Field[interface{}] `json:"confirm_my_choices,required"`
	// Object where keys are language codes
	RejectAll param.Field[interface{}] `json:"reject_all,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsConsentButtonTextTranslations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// General Zaraz settings.
type ZoneSettingZarazConfigUpdateParamsSettings struct {
	// Automatic injection of Zaraz scripts enabled.
	AutoInjectScript param.Field[bool] `json:"autoInjectScript,required"`
	// Details of the worker that receives and edits Zaraz Context object.
	ContextEnricher param.Field[ZoneSettingZarazConfigUpdateParamsSettingsContextEnricher] `json:"contextEnricher"`
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

func (r ZoneSettingZarazConfigUpdateParamsSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Details of the worker that receives and edits Zaraz Context object.
type ZoneSettingZarazConfigUpdateParamsSettingsContextEnricher struct {
	EscapedWorkerName param.Field[string] `json:"escapedWorkerName,required"`
	WorkerTag         param.Field[string] `json:"workerTag,required"`
}

func (r ZoneSettingZarazConfigUpdateParamsSettingsContextEnricher) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
