// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
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

type ZarazDefaultGetResponse struct {
	// Consent management configuration.
	Consent ZarazDefaultGetResponseConsent `json:"consent"`
	// Data layer compatibility mode enabled.
	DataLayer bool `json:"dataLayer"`
	// The key for Zaraz debug mode.
	DebugKey string `json:"debugKey"`
	// Single Page Application support enabled.
	HistoryChange bool `json:"historyChange"`
	// General Zaraz settings.
	Settings ZarazDefaultGetResponseSettings `json:"settings"`
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
	ZarazVersion int64                       `json:"zarazVersion"`
	JSON         zarazDefaultGetResponseJSON `json:"-"`
}

// zarazDefaultGetResponseJSON contains the JSON metadata for the struct
// [ZarazDefaultGetResponse]
type zarazDefaultGetResponseJSON struct {
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

func (r *ZarazDefaultGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Consent management configuration.
type ZarazDefaultGetResponseConsent struct {
	Enabled                bool                                                 `json:"enabled,required"`
	ButtonTextTranslations ZarazDefaultGetResponseConsentButtonTextTranslations `json:"buttonTextTranslations"`
	CompanyEmail           string                                               `json:"companyEmail"`
	CompanyName            string                                               `json:"companyName"`
	CompanyStreetAddress   string                                               `json:"companyStreetAddress"`
	ConsentModalIntroHTML  string                                               `json:"consentModalIntroHTML"`
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
	PurposesWithTranslations interface{}                        `json:"purposesWithTranslations"`
	JSON                     zarazDefaultGetResponseConsentJSON `json:"-"`
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
	AcceptAll interface{} `json:"accept_all,required"`
	// Object where keys are language codes
	ConfirmMyChoices interface{} `json:"confirm_my_choices,required"`
	// Object where keys are language codes
	RejectAll interface{}                                              `json:"reject_all,required"`
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

type ZarazDefaultGetResponseEnvelope struct {
	Errors   []ZarazDefaultGetResponseEnvelopeErrors   `json:"errors"`
	Messages []ZarazDefaultGetResponseEnvelopeMessages `json:"messages"`
	Result   ZarazDefaultGetResponse                   `json:"result"`
	// Whether the API call was successful
	Success bool                                `json:"success"`
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
