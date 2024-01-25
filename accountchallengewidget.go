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
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountChallengeWidgetService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountChallengeWidgetService]
// method instead.
type AccountChallengeWidgetService struct {
	Options []option.RequestOption
}

// NewAccountChallengeWidgetService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountChallengeWidgetService(opts ...option.RequestOption) (r *AccountChallengeWidgetService) {
	r = &AccountChallengeWidgetService{}
	r.Options = opts
	return
}

// Lists challenge widgets.
func (r *AccountChallengeWidgetService) New(ctx context.Context, accountIdentifier string, params AccountChallengeWidgetNewParams, opts ...option.RequestOption) (res *AccountChallengeWidgetNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/challenges/widgets", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Show a single challenge widget configuration.
func (r *AccountChallengeWidgetService) Get(ctx context.Context, accountIdentifier string, sitekey string, opts ...option.RequestOption) (res *AccountChallengeWidgetGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/challenges/widgets/%s", accountIdentifier, sitekey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the configuration of a widget.
func (r *AccountChallengeWidgetService) Update(ctx context.Context, accountIdentifier string, sitekey string, body AccountChallengeWidgetUpdateParams, opts ...option.RequestOption) (res *AccountChallengeWidgetUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/challenges/widgets/%s", accountIdentifier, sitekey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists all turnstile widgets of an account.
func (r *AccountChallengeWidgetService) List(ctx context.Context, accountIdentifier string, query AccountChallengeWidgetListParams, opts ...option.RequestOption) (res *shared.Page[AccountChallengeWidgetListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/challenges/widgets", accountIdentifier)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Destroy a Turnstile Widget.
func (r *AccountChallengeWidgetService) Delete(ctx context.Context, accountIdentifier string, sitekey string, opts ...option.RequestOption) (res *AccountChallengeWidgetDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/challenges/widgets/%s", accountIdentifier, sitekey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Generate a new secret key for this widget. If `invalidate_immediately` is set to
// `false`, the previous secret remains valid for 2 hours.
//
// Note that secrets cannot be rotated again during the grace period.
func (r *AccountChallengeWidgetService) RotateSecret(ctx context.Context, accountIdentifier string, sitekey string, body AccountChallengeWidgetRotateSecretParams, opts ...option.RequestOption) (res *AccountChallengeWidgetRotateSecretResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/challenges/widgets/%s/rotate_secret", accountIdentifier, sitekey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountChallengeWidgetNewResponse struct {
	Errors   []AccountChallengeWidgetNewResponseError   `json:"errors"`
	Messages []AccountChallengeWidgetNewResponseMessage `json:"messages"`
	// A Turnstile widget's detailed configuration
	Result     AccountChallengeWidgetNewResponseResult     `json:"result"`
	ResultInfo AccountChallengeWidgetNewResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success bool                                  `json:"success"`
	JSON    accountChallengeWidgetNewResponseJSON `json:"-"`
}

// accountChallengeWidgetNewResponseJSON contains the JSON metadata for the struct
// [AccountChallengeWidgetNewResponse]
type accountChallengeWidgetNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountChallengeWidgetNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountChallengeWidgetNewResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accountChallengeWidgetNewResponseErrorJSON `json:"-"`
}

// accountChallengeWidgetNewResponseErrorJSON contains the JSON metadata for the
// struct [AccountChallengeWidgetNewResponseError]
type accountChallengeWidgetNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountChallengeWidgetNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountChallengeWidgetNewResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountChallengeWidgetNewResponseMessageJSON `json:"-"`
}

// accountChallengeWidgetNewResponseMessageJSON contains the JSON metadata for the
// struct [AccountChallengeWidgetNewResponseMessage]
type accountChallengeWidgetNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountChallengeWidgetNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Turnstile widget's detailed configuration
type AccountChallengeWidgetNewResponseResult struct {
	// If bot_fight_mode is set to `true`, Cloudflare issues computationally expensive
	// challenges in response to malicious bots (ENT only).
	BotFightMode bool `json:"bot_fight_mode,required"`
	// If Turnstile is embedded on a Cloudflare site and the widget should grant
	// challenge clearance, this setting can determine the clearance level to be set
	ClearanceLevel AccountChallengeWidgetNewResponseResultClearanceLevel `json:"clearance_level,required"`
	// When the widget was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	Domains   []string  `json:"domains,required"`
	// Widget Mode
	Mode AccountChallengeWidgetNewResponseResultMode `json:"mode,required"`
	// When the widget was modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Human readable widget name. Not unique. Cloudflare suggests that you set this to
	// a meaningful string to make it easier to identify your widget, and where it is
	// used.
	Name string `json:"name,required"`
	// Do not show any Cloudflare branding on the widget (ENT only).
	Offlabel bool `json:"offlabel,required"`
	// Region where this widget can be used.
	Region AccountChallengeWidgetNewResponseResultRegion `json:"region,required"`
	// Secret key for this widget.
	Secret string `json:"secret,required"`
	// Widget item identifier tag.
	Sitekey string                                      `json:"sitekey,required"`
	JSON    accountChallengeWidgetNewResponseResultJSON `json:"-"`
}

// accountChallengeWidgetNewResponseResultJSON contains the JSON metadata for the
// struct [AccountChallengeWidgetNewResponseResult]
type accountChallengeWidgetNewResponseResultJSON struct {
	BotFightMode   apijson.Field
	ClearanceLevel apijson.Field
	CreatedOn      apijson.Field
	Domains        apijson.Field
	Mode           apijson.Field
	ModifiedOn     apijson.Field
	Name           apijson.Field
	Offlabel       apijson.Field
	Region         apijson.Field
	Secret         apijson.Field
	Sitekey        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountChallengeWidgetNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// If Turnstile is embedded on a Cloudflare site and the widget should grant
// challenge clearance, this setting can determine the clearance level to be set
type AccountChallengeWidgetNewResponseResultClearanceLevel string

const (
	AccountChallengeWidgetNewResponseResultClearanceLevelNoClearance AccountChallengeWidgetNewResponseResultClearanceLevel = "no_clearance"
	AccountChallengeWidgetNewResponseResultClearanceLevelJschallenge AccountChallengeWidgetNewResponseResultClearanceLevel = "jschallenge"
	AccountChallengeWidgetNewResponseResultClearanceLevelManaged     AccountChallengeWidgetNewResponseResultClearanceLevel = "managed"
	AccountChallengeWidgetNewResponseResultClearanceLevelInteractive AccountChallengeWidgetNewResponseResultClearanceLevel = "interactive"
)

// Widget Mode
type AccountChallengeWidgetNewResponseResultMode string

const (
	AccountChallengeWidgetNewResponseResultModeNonInteractive AccountChallengeWidgetNewResponseResultMode = "non-interactive"
	AccountChallengeWidgetNewResponseResultModeInvisible      AccountChallengeWidgetNewResponseResultMode = "invisible"
	AccountChallengeWidgetNewResponseResultModeManaged        AccountChallengeWidgetNewResponseResultMode = "managed"
)

// Region where this widget can be used.
type AccountChallengeWidgetNewResponseResultRegion string

const (
	AccountChallengeWidgetNewResponseResultRegionWorld AccountChallengeWidgetNewResponseResultRegion = "world"
)

type AccountChallengeWidgetNewResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count,required"`
	// Current page within paginated list of results
	Page float64 `json:"page,required"`
	// Number of results per page of results
	PerPage float64 `json:"per_page,required"`
	// Total results available without any search parameters
	TotalCount float64                                         `json:"total_count,required"`
	JSON       accountChallengeWidgetNewResponseResultInfoJSON `json:"-"`
}

// accountChallengeWidgetNewResponseResultInfoJSON contains the JSON metadata for
// the struct [AccountChallengeWidgetNewResponseResultInfo]
type accountChallengeWidgetNewResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountChallengeWidgetNewResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountChallengeWidgetGetResponse struct {
	Errors   []AccountChallengeWidgetGetResponseError   `json:"errors"`
	Messages []AccountChallengeWidgetGetResponseMessage `json:"messages"`
	// A Turnstile widget's detailed configuration
	Result AccountChallengeWidgetGetResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                                  `json:"success"`
	JSON    accountChallengeWidgetGetResponseJSON `json:"-"`
}

// accountChallengeWidgetGetResponseJSON contains the JSON metadata for the struct
// [AccountChallengeWidgetGetResponse]
type accountChallengeWidgetGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountChallengeWidgetGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountChallengeWidgetGetResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accountChallengeWidgetGetResponseErrorJSON `json:"-"`
}

// accountChallengeWidgetGetResponseErrorJSON contains the JSON metadata for the
// struct [AccountChallengeWidgetGetResponseError]
type accountChallengeWidgetGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountChallengeWidgetGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountChallengeWidgetGetResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountChallengeWidgetGetResponseMessageJSON `json:"-"`
}

// accountChallengeWidgetGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountChallengeWidgetGetResponseMessage]
type accountChallengeWidgetGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountChallengeWidgetGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Turnstile widget's detailed configuration
type AccountChallengeWidgetGetResponseResult struct {
	// If bot_fight_mode is set to `true`, Cloudflare issues computationally expensive
	// challenges in response to malicious bots (ENT only).
	BotFightMode bool `json:"bot_fight_mode,required"`
	// If Turnstile is embedded on a Cloudflare site and the widget should grant
	// challenge clearance, this setting can determine the clearance level to be set
	ClearanceLevel AccountChallengeWidgetGetResponseResultClearanceLevel `json:"clearance_level,required"`
	// When the widget was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	Domains   []string  `json:"domains,required"`
	// Widget Mode
	Mode AccountChallengeWidgetGetResponseResultMode `json:"mode,required"`
	// When the widget was modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Human readable widget name. Not unique. Cloudflare suggests that you set this to
	// a meaningful string to make it easier to identify your widget, and where it is
	// used.
	Name string `json:"name,required"`
	// Do not show any Cloudflare branding on the widget (ENT only).
	Offlabel bool `json:"offlabel,required"`
	// Region where this widget can be used.
	Region AccountChallengeWidgetGetResponseResultRegion `json:"region,required"`
	// Secret key for this widget.
	Secret string `json:"secret,required"`
	// Widget item identifier tag.
	Sitekey string                                      `json:"sitekey,required"`
	JSON    accountChallengeWidgetGetResponseResultJSON `json:"-"`
}

// accountChallengeWidgetGetResponseResultJSON contains the JSON metadata for the
// struct [AccountChallengeWidgetGetResponseResult]
type accountChallengeWidgetGetResponseResultJSON struct {
	BotFightMode   apijson.Field
	ClearanceLevel apijson.Field
	CreatedOn      apijson.Field
	Domains        apijson.Field
	Mode           apijson.Field
	ModifiedOn     apijson.Field
	Name           apijson.Field
	Offlabel       apijson.Field
	Region         apijson.Field
	Secret         apijson.Field
	Sitekey        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountChallengeWidgetGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// If Turnstile is embedded on a Cloudflare site and the widget should grant
// challenge clearance, this setting can determine the clearance level to be set
type AccountChallengeWidgetGetResponseResultClearanceLevel string

const (
	AccountChallengeWidgetGetResponseResultClearanceLevelNoClearance AccountChallengeWidgetGetResponseResultClearanceLevel = "no_clearance"
	AccountChallengeWidgetGetResponseResultClearanceLevelJschallenge AccountChallengeWidgetGetResponseResultClearanceLevel = "jschallenge"
	AccountChallengeWidgetGetResponseResultClearanceLevelManaged     AccountChallengeWidgetGetResponseResultClearanceLevel = "managed"
	AccountChallengeWidgetGetResponseResultClearanceLevelInteractive AccountChallengeWidgetGetResponseResultClearanceLevel = "interactive"
)

// Widget Mode
type AccountChallengeWidgetGetResponseResultMode string

const (
	AccountChallengeWidgetGetResponseResultModeNonInteractive AccountChallengeWidgetGetResponseResultMode = "non-interactive"
	AccountChallengeWidgetGetResponseResultModeInvisible      AccountChallengeWidgetGetResponseResultMode = "invisible"
	AccountChallengeWidgetGetResponseResultModeManaged        AccountChallengeWidgetGetResponseResultMode = "managed"
)

// Region where this widget can be used.
type AccountChallengeWidgetGetResponseResultRegion string

const (
	AccountChallengeWidgetGetResponseResultRegionWorld AccountChallengeWidgetGetResponseResultRegion = "world"
)

type AccountChallengeWidgetUpdateResponse struct {
	Errors   []AccountChallengeWidgetUpdateResponseError   `json:"errors"`
	Messages []AccountChallengeWidgetUpdateResponseMessage `json:"messages"`
	// A Turnstile widget's detailed configuration
	Result AccountChallengeWidgetUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                                     `json:"success"`
	JSON    accountChallengeWidgetUpdateResponseJSON `json:"-"`
}

// accountChallengeWidgetUpdateResponseJSON contains the JSON metadata for the
// struct [AccountChallengeWidgetUpdateResponse]
type accountChallengeWidgetUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountChallengeWidgetUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountChallengeWidgetUpdateResponseError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountChallengeWidgetUpdateResponseErrorJSON `json:"-"`
}

// accountChallengeWidgetUpdateResponseErrorJSON contains the JSON metadata for the
// struct [AccountChallengeWidgetUpdateResponseError]
type accountChallengeWidgetUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountChallengeWidgetUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountChallengeWidgetUpdateResponseMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accountChallengeWidgetUpdateResponseMessageJSON `json:"-"`
}

// accountChallengeWidgetUpdateResponseMessageJSON contains the JSON metadata for
// the struct [AccountChallengeWidgetUpdateResponseMessage]
type accountChallengeWidgetUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountChallengeWidgetUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Turnstile widget's detailed configuration
type AccountChallengeWidgetUpdateResponseResult struct {
	// If bot_fight_mode is set to `true`, Cloudflare issues computationally expensive
	// challenges in response to malicious bots (ENT only).
	BotFightMode bool `json:"bot_fight_mode,required"`
	// If Turnstile is embedded on a Cloudflare site and the widget should grant
	// challenge clearance, this setting can determine the clearance level to be set
	ClearanceLevel AccountChallengeWidgetUpdateResponseResultClearanceLevel `json:"clearance_level,required"`
	// When the widget was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	Domains   []string  `json:"domains,required"`
	// Widget Mode
	Mode AccountChallengeWidgetUpdateResponseResultMode `json:"mode,required"`
	// When the widget was modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Human readable widget name. Not unique. Cloudflare suggests that you set this to
	// a meaningful string to make it easier to identify your widget, and where it is
	// used.
	Name string `json:"name,required"`
	// Do not show any Cloudflare branding on the widget (ENT only).
	Offlabel bool `json:"offlabel,required"`
	// Region where this widget can be used.
	Region AccountChallengeWidgetUpdateResponseResultRegion `json:"region,required"`
	// Secret key for this widget.
	Secret string `json:"secret,required"`
	// Widget item identifier tag.
	Sitekey string                                         `json:"sitekey,required"`
	JSON    accountChallengeWidgetUpdateResponseResultJSON `json:"-"`
}

// accountChallengeWidgetUpdateResponseResultJSON contains the JSON metadata for
// the struct [AccountChallengeWidgetUpdateResponseResult]
type accountChallengeWidgetUpdateResponseResultJSON struct {
	BotFightMode   apijson.Field
	ClearanceLevel apijson.Field
	CreatedOn      apijson.Field
	Domains        apijson.Field
	Mode           apijson.Field
	ModifiedOn     apijson.Field
	Name           apijson.Field
	Offlabel       apijson.Field
	Region         apijson.Field
	Secret         apijson.Field
	Sitekey        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountChallengeWidgetUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// If Turnstile is embedded on a Cloudflare site and the widget should grant
// challenge clearance, this setting can determine the clearance level to be set
type AccountChallengeWidgetUpdateResponseResultClearanceLevel string

const (
	AccountChallengeWidgetUpdateResponseResultClearanceLevelNoClearance AccountChallengeWidgetUpdateResponseResultClearanceLevel = "no_clearance"
	AccountChallengeWidgetUpdateResponseResultClearanceLevelJschallenge AccountChallengeWidgetUpdateResponseResultClearanceLevel = "jschallenge"
	AccountChallengeWidgetUpdateResponseResultClearanceLevelManaged     AccountChallengeWidgetUpdateResponseResultClearanceLevel = "managed"
	AccountChallengeWidgetUpdateResponseResultClearanceLevelInteractive AccountChallengeWidgetUpdateResponseResultClearanceLevel = "interactive"
)

// Widget Mode
type AccountChallengeWidgetUpdateResponseResultMode string

const (
	AccountChallengeWidgetUpdateResponseResultModeNonInteractive AccountChallengeWidgetUpdateResponseResultMode = "non-interactive"
	AccountChallengeWidgetUpdateResponseResultModeInvisible      AccountChallengeWidgetUpdateResponseResultMode = "invisible"
	AccountChallengeWidgetUpdateResponseResultModeManaged        AccountChallengeWidgetUpdateResponseResultMode = "managed"
)

// Region where this widget can be used.
type AccountChallengeWidgetUpdateResponseResultRegion string

const (
	AccountChallengeWidgetUpdateResponseResultRegionWorld AccountChallengeWidgetUpdateResponseResultRegion = "world"
)

// A Turnstile Widgets configuration as it appears in listings
type AccountChallengeWidgetListResponse struct {
	// If bot_fight_mode is set to `true`, Cloudflare issues computationally expensive
	// challenges in response to malicious bots (ENT only).
	BotFightMode bool `json:"bot_fight_mode,required"`
	// If Turnstile is embedded on a Cloudflare site and the widget should grant
	// challenge clearance, this setting can determine the clearance level to be set
	ClearanceLevel AccountChallengeWidgetListResponseClearanceLevel `json:"clearance_level,required"`
	// When the widget was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	Domains   []string  `json:"domains,required"`
	// Widget Mode
	Mode AccountChallengeWidgetListResponseMode `json:"mode,required"`
	// When the widget was modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Human readable widget name. Not unique. Cloudflare suggests that you set this to
	// a meaningful string to make it easier to identify your widget, and where it is
	// used.
	Name string `json:"name,required"`
	// Do not show any Cloudflare branding on the widget (ENT only).
	Offlabel bool `json:"offlabel,required"`
	// Region where this widget can be used.
	Region AccountChallengeWidgetListResponseRegion `json:"region,required"`
	// Widget item identifier tag.
	Sitekey string                                 `json:"sitekey,required"`
	JSON    accountChallengeWidgetListResponseJSON `json:"-"`
}

// accountChallengeWidgetListResponseJSON contains the JSON metadata for the struct
// [AccountChallengeWidgetListResponse]
type accountChallengeWidgetListResponseJSON struct {
	BotFightMode   apijson.Field
	ClearanceLevel apijson.Field
	CreatedOn      apijson.Field
	Domains        apijson.Field
	Mode           apijson.Field
	ModifiedOn     apijson.Field
	Name           apijson.Field
	Offlabel       apijson.Field
	Region         apijson.Field
	Sitekey        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountChallengeWidgetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// If Turnstile is embedded on a Cloudflare site and the widget should grant
// challenge clearance, this setting can determine the clearance level to be set
type AccountChallengeWidgetListResponseClearanceLevel string

const (
	AccountChallengeWidgetListResponseClearanceLevelNoClearance AccountChallengeWidgetListResponseClearanceLevel = "no_clearance"
	AccountChallengeWidgetListResponseClearanceLevelJschallenge AccountChallengeWidgetListResponseClearanceLevel = "jschallenge"
	AccountChallengeWidgetListResponseClearanceLevelManaged     AccountChallengeWidgetListResponseClearanceLevel = "managed"
	AccountChallengeWidgetListResponseClearanceLevelInteractive AccountChallengeWidgetListResponseClearanceLevel = "interactive"
)

// Widget Mode
type AccountChallengeWidgetListResponseMode string

const (
	AccountChallengeWidgetListResponseModeNonInteractive AccountChallengeWidgetListResponseMode = "non-interactive"
	AccountChallengeWidgetListResponseModeInvisible      AccountChallengeWidgetListResponseMode = "invisible"
	AccountChallengeWidgetListResponseModeManaged        AccountChallengeWidgetListResponseMode = "managed"
)

// Region where this widget can be used.
type AccountChallengeWidgetListResponseRegion string

const (
	AccountChallengeWidgetListResponseRegionWorld AccountChallengeWidgetListResponseRegion = "world"
)

type AccountChallengeWidgetDeleteResponse struct {
	Errors   []AccountChallengeWidgetDeleteResponseError   `json:"errors"`
	Messages []AccountChallengeWidgetDeleteResponseMessage `json:"messages"`
	// A Turnstile widget's detailed configuration
	Result AccountChallengeWidgetDeleteResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                                     `json:"success"`
	JSON    accountChallengeWidgetDeleteResponseJSON `json:"-"`
}

// accountChallengeWidgetDeleteResponseJSON contains the JSON metadata for the
// struct [AccountChallengeWidgetDeleteResponse]
type accountChallengeWidgetDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountChallengeWidgetDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountChallengeWidgetDeleteResponseError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountChallengeWidgetDeleteResponseErrorJSON `json:"-"`
}

// accountChallengeWidgetDeleteResponseErrorJSON contains the JSON metadata for the
// struct [AccountChallengeWidgetDeleteResponseError]
type accountChallengeWidgetDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountChallengeWidgetDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountChallengeWidgetDeleteResponseMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accountChallengeWidgetDeleteResponseMessageJSON `json:"-"`
}

// accountChallengeWidgetDeleteResponseMessageJSON contains the JSON metadata for
// the struct [AccountChallengeWidgetDeleteResponseMessage]
type accountChallengeWidgetDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountChallengeWidgetDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Turnstile widget's detailed configuration
type AccountChallengeWidgetDeleteResponseResult struct {
	// If bot_fight_mode is set to `true`, Cloudflare issues computationally expensive
	// challenges in response to malicious bots (ENT only).
	BotFightMode bool `json:"bot_fight_mode,required"`
	// If Turnstile is embedded on a Cloudflare site and the widget should grant
	// challenge clearance, this setting can determine the clearance level to be set
	ClearanceLevel AccountChallengeWidgetDeleteResponseResultClearanceLevel `json:"clearance_level,required"`
	// When the widget was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	Domains   []string  `json:"domains,required"`
	// Widget Mode
	Mode AccountChallengeWidgetDeleteResponseResultMode `json:"mode,required"`
	// When the widget was modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Human readable widget name. Not unique. Cloudflare suggests that you set this to
	// a meaningful string to make it easier to identify your widget, and where it is
	// used.
	Name string `json:"name,required"`
	// Do not show any Cloudflare branding on the widget (ENT only).
	Offlabel bool `json:"offlabel,required"`
	// Region where this widget can be used.
	Region AccountChallengeWidgetDeleteResponseResultRegion `json:"region,required"`
	// Secret key for this widget.
	Secret string `json:"secret,required"`
	// Widget item identifier tag.
	Sitekey string                                         `json:"sitekey,required"`
	JSON    accountChallengeWidgetDeleteResponseResultJSON `json:"-"`
}

// accountChallengeWidgetDeleteResponseResultJSON contains the JSON metadata for
// the struct [AccountChallengeWidgetDeleteResponseResult]
type accountChallengeWidgetDeleteResponseResultJSON struct {
	BotFightMode   apijson.Field
	ClearanceLevel apijson.Field
	CreatedOn      apijson.Field
	Domains        apijson.Field
	Mode           apijson.Field
	ModifiedOn     apijson.Field
	Name           apijson.Field
	Offlabel       apijson.Field
	Region         apijson.Field
	Secret         apijson.Field
	Sitekey        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountChallengeWidgetDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// If Turnstile is embedded on a Cloudflare site and the widget should grant
// challenge clearance, this setting can determine the clearance level to be set
type AccountChallengeWidgetDeleteResponseResultClearanceLevel string

const (
	AccountChallengeWidgetDeleteResponseResultClearanceLevelNoClearance AccountChallengeWidgetDeleteResponseResultClearanceLevel = "no_clearance"
	AccountChallengeWidgetDeleteResponseResultClearanceLevelJschallenge AccountChallengeWidgetDeleteResponseResultClearanceLevel = "jschallenge"
	AccountChallengeWidgetDeleteResponseResultClearanceLevelManaged     AccountChallengeWidgetDeleteResponseResultClearanceLevel = "managed"
	AccountChallengeWidgetDeleteResponseResultClearanceLevelInteractive AccountChallengeWidgetDeleteResponseResultClearanceLevel = "interactive"
)

// Widget Mode
type AccountChallengeWidgetDeleteResponseResultMode string

const (
	AccountChallengeWidgetDeleteResponseResultModeNonInteractive AccountChallengeWidgetDeleteResponseResultMode = "non-interactive"
	AccountChallengeWidgetDeleteResponseResultModeInvisible      AccountChallengeWidgetDeleteResponseResultMode = "invisible"
	AccountChallengeWidgetDeleteResponseResultModeManaged        AccountChallengeWidgetDeleteResponseResultMode = "managed"
)

// Region where this widget can be used.
type AccountChallengeWidgetDeleteResponseResultRegion string

const (
	AccountChallengeWidgetDeleteResponseResultRegionWorld AccountChallengeWidgetDeleteResponseResultRegion = "world"
)

type AccountChallengeWidgetRotateSecretResponse struct {
	Errors   []AccountChallengeWidgetRotateSecretResponseError   `json:"errors"`
	Messages []AccountChallengeWidgetRotateSecretResponseMessage `json:"messages"`
	// A Turnstile widget's detailed configuration
	Result AccountChallengeWidgetRotateSecretResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                                           `json:"success"`
	JSON    accountChallengeWidgetRotateSecretResponseJSON `json:"-"`
}

// accountChallengeWidgetRotateSecretResponseJSON contains the JSON metadata for
// the struct [AccountChallengeWidgetRotateSecretResponse]
type accountChallengeWidgetRotateSecretResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountChallengeWidgetRotateSecretResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountChallengeWidgetRotateSecretResponseError struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accountChallengeWidgetRotateSecretResponseErrorJSON `json:"-"`
}

// accountChallengeWidgetRotateSecretResponseErrorJSON contains the JSON metadata
// for the struct [AccountChallengeWidgetRotateSecretResponseError]
type accountChallengeWidgetRotateSecretResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountChallengeWidgetRotateSecretResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountChallengeWidgetRotateSecretResponseMessage struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    accountChallengeWidgetRotateSecretResponseMessageJSON `json:"-"`
}

// accountChallengeWidgetRotateSecretResponseMessageJSON contains the JSON metadata
// for the struct [AccountChallengeWidgetRotateSecretResponseMessage]
type accountChallengeWidgetRotateSecretResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountChallengeWidgetRotateSecretResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Turnstile widget's detailed configuration
type AccountChallengeWidgetRotateSecretResponseResult struct {
	// If bot_fight_mode is set to `true`, Cloudflare issues computationally expensive
	// challenges in response to malicious bots (ENT only).
	BotFightMode bool `json:"bot_fight_mode,required"`
	// If Turnstile is embedded on a Cloudflare site and the widget should grant
	// challenge clearance, this setting can determine the clearance level to be set
	ClearanceLevel AccountChallengeWidgetRotateSecretResponseResultClearanceLevel `json:"clearance_level,required"`
	// When the widget was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	Domains   []string  `json:"domains,required"`
	// Widget Mode
	Mode AccountChallengeWidgetRotateSecretResponseResultMode `json:"mode,required"`
	// When the widget was modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Human readable widget name. Not unique. Cloudflare suggests that you set this to
	// a meaningful string to make it easier to identify your widget, and where it is
	// used.
	Name string `json:"name,required"`
	// Do not show any Cloudflare branding on the widget (ENT only).
	Offlabel bool `json:"offlabel,required"`
	// Region where this widget can be used.
	Region AccountChallengeWidgetRotateSecretResponseResultRegion `json:"region,required"`
	// Secret key for this widget.
	Secret string `json:"secret,required"`
	// Widget item identifier tag.
	Sitekey string                                               `json:"sitekey,required"`
	JSON    accountChallengeWidgetRotateSecretResponseResultJSON `json:"-"`
}

// accountChallengeWidgetRotateSecretResponseResultJSON contains the JSON metadata
// for the struct [AccountChallengeWidgetRotateSecretResponseResult]
type accountChallengeWidgetRotateSecretResponseResultJSON struct {
	BotFightMode   apijson.Field
	ClearanceLevel apijson.Field
	CreatedOn      apijson.Field
	Domains        apijson.Field
	Mode           apijson.Field
	ModifiedOn     apijson.Field
	Name           apijson.Field
	Offlabel       apijson.Field
	Region         apijson.Field
	Secret         apijson.Field
	Sitekey        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountChallengeWidgetRotateSecretResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// If Turnstile is embedded on a Cloudflare site and the widget should grant
// challenge clearance, this setting can determine the clearance level to be set
type AccountChallengeWidgetRotateSecretResponseResultClearanceLevel string

const (
	AccountChallengeWidgetRotateSecretResponseResultClearanceLevelNoClearance AccountChallengeWidgetRotateSecretResponseResultClearanceLevel = "no_clearance"
	AccountChallengeWidgetRotateSecretResponseResultClearanceLevelJschallenge AccountChallengeWidgetRotateSecretResponseResultClearanceLevel = "jschallenge"
	AccountChallengeWidgetRotateSecretResponseResultClearanceLevelManaged     AccountChallengeWidgetRotateSecretResponseResultClearanceLevel = "managed"
	AccountChallengeWidgetRotateSecretResponseResultClearanceLevelInteractive AccountChallengeWidgetRotateSecretResponseResultClearanceLevel = "interactive"
)

// Widget Mode
type AccountChallengeWidgetRotateSecretResponseResultMode string

const (
	AccountChallengeWidgetRotateSecretResponseResultModeNonInteractive AccountChallengeWidgetRotateSecretResponseResultMode = "non-interactive"
	AccountChallengeWidgetRotateSecretResponseResultModeInvisible      AccountChallengeWidgetRotateSecretResponseResultMode = "invisible"
	AccountChallengeWidgetRotateSecretResponseResultModeManaged        AccountChallengeWidgetRotateSecretResponseResultMode = "managed"
)

// Region where this widget can be used.
type AccountChallengeWidgetRotateSecretResponseResultRegion string

const (
	AccountChallengeWidgetRotateSecretResponseResultRegionWorld AccountChallengeWidgetRotateSecretResponseResultRegion = "world"
)

type AccountChallengeWidgetNewParams struct {
	Domains param.Field[[]string] `json:"domains,required"`
	// Widget Mode
	Mode param.Field[AccountChallengeWidgetNewParamsMode] `json:"mode,required"`
	// Human readable widget name. Not unique. Cloudflare suggests that you set this to
	// a meaningful string to make it easier to identify your widget, and where it is
	// used.
	Name param.Field[string] `json:"name,required"`
	// Direction to order widgets.
	Direction param.Field[AccountChallengeWidgetNewParamsDirection] `query:"direction"`
	// Field to order widgets by.
	Order param.Field[AccountChallengeWidgetNewParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of items per page.
	PerPage param.Field[float64] `query:"per_page"`
	// If bot_fight_mode is set to `true`, Cloudflare issues computationally expensive
	// challenges in response to malicious bots (ENT only).
	BotFightMode param.Field[bool] `json:"bot_fight_mode"`
	// If Turnstile is embedded on a Cloudflare site and the widget should grant
	// challenge clearance, this setting can determine the clearance level to be set
	ClearanceLevel param.Field[AccountChallengeWidgetNewParamsClearanceLevel] `json:"clearance_level"`
	// Do not show any Cloudflare branding on the widget (ENT only).
	Offlabel param.Field[bool] `json:"offlabel"`
	// Region where this widget can be used.
	Region param.Field[AccountChallengeWidgetNewParamsRegion] `json:"region"`
}

func (r AccountChallengeWidgetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AccountChallengeWidgetNewParams]'s query parameters as
// `url.Values`.
func (r AccountChallengeWidgetNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Widget Mode
type AccountChallengeWidgetNewParamsMode string

const (
	AccountChallengeWidgetNewParamsModeNonInteractive AccountChallengeWidgetNewParamsMode = "non-interactive"
	AccountChallengeWidgetNewParamsModeInvisible      AccountChallengeWidgetNewParamsMode = "invisible"
	AccountChallengeWidgetNewParamsModeManaged        AccountChallengeWidgetNewParamsMode = "managed"
)

// Direction to order widgets.
type AccountChallengeWidgetNewParamsDirection string

const (
	AccountChallengeWidgetNewParamsDirectionAsc  AccountChallengeWidgetNewParamsDirection = "asc"
	AccountChallengeWidgetNewParamsDirectionDesc AccountChallengeWidgetNewParamsDirection = "desc"
)

// Field to order widgets by.
type AccountChallengeWidgetNewParamsOrder string

const (
	AccountChallengeWidgetNewParamsOrderID         AccountChallengeWidgetNewParamsOrder = "id"
	AccountChallengeWidgetNewParamsOrderSitekey    AccountChallengeWidgetNewParamsOrder = "sitekey"
	AccountChallengeWidgetNewParamsOrderName       AccountChallengeWidgetNewParamsOrder = "name"
	AccountChallengeWidgetNewParamsOrderCreatedOn  AccountChallengeWidgetNewParamsOrder = "created_on"
	AccountChallengeWidgetNewParamsOrderModifiedOn AccountChallengeWidgetNewParamsOrder = "modified_on"
)

// If Turnstile is embedded on a Cloudflare site and the widget should grant
// challenge clearance, this setting can determine the clearance level to be set
type AccountChallengeWidgetNewParamsClearanceLevel string

const (
	AccountChallengeWidgetNewParamsClearanceLevelNoClearance AccountChallengeWidgetNewParamsClearanceLevel = "no_clearance"
	AccountChallengeWidgetNewParamsClearanceLevelJschallenge AccountChallengeWidgetNewParamsClearanceLevel = "jschallenge"
	AccountChallengeWidgetNewParamsClearanceLevelManaged     AccountChallengeWidgetNewParamsClearanceLevel = "managed"
	AccountChallengeWidgetNewParamsClearanceLevelInteractive AccountChallengeWidgetNewParamsClearanceLevel = "interactive"
)

// Region where this widget can be used.
type AccountChallengeWidgetNewParamsRegion string

const (
	AccountChallengeWidgetNewParamsRegionWorld AccountChallengeWidgetNewParamsRegion = "world"
)

type AccountChallengeWidgetUpdateParams struct {
	Domains param.Field[[]string] `json:"domains,required"`
	// Widget Mode
	Mode param.Field[AccountChallengeWidgetUpdateParamsMode] `json:"mode,required"`
	// Human readable widget name. Not unique. Cloudflare suggests that you set this to
	// a meaningful string to make it easier to identify your widget, and where it is
	// used.
	Name param.Field[string] `json:"name,required"`
	// If bot_fight_mode is set to `true`, Cloudflare issues computationally expensive
	// challenges in response to malicious bots (ENT only).
	BotFightMode param.Field[bool] `json:"bot_fight_mode"`
	// If Turnstile is embedded on a Cloudflare site and the widget should grant
	// challenge clearance, this setting can determine the clearance level to be set
	ClearanceLevel param.Field[AccountChallengeWidgetUpdateParamsClearanceLevel] `json:"clearance_level"`
	// Do not show any Cloudflare branding on the widget (ENT only).
	Offlabel param.Field[bool] `json:"offlabel"`
}

func (r AccountChallengeWidgetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Widget Mode
type AccountChallengeWidgetUpdateParamsMode string

const (
	AccountChallengeWidgetUpdateParamsModeNonInteractive AccountChallengeWidgetUpdateParamsMode = "non-interactive"
	AccountChallengeWidgetUpdateParamsModeInvisible      AccountChallengeWidgetUpdateParamsMode = "invisible"
	AccountChallengeWidgetUpdateParamsModeManaged        AccountChallengeWidgetUpdateParamsMode = "managed"
)

// If Turnstile is embedded on a Cloudflare site and the widget should grant
// challenge clearance, this setting can determine the clearance level to be set
type AccountChallengeWidgetUpdateParamsClearanceLevel string

const (
	AccountChallengeWidgetUpdateParamsClearanceLevelNoClearance AccountChallengeWidgetUpdateParamsClearanceLevel = "no_clearance"
	AccountChallengeWidgetUpdateParamsClearanceLevelJschallenge AccountChallengeWidgetUpdateParamsClearanceLevel = "jschallenge"
	AccountChallengeWidgetUpdateParamsClearanceLevelManaged     AccountChallengeWidgetUpdateParamsClearanceLevel = "managed"
	AccountChallengeWidgetUpdateParamsClearanceLevelInteractive AccountChallengeWidgetUpdateParamsClearanceLevel = "interactive"
)

type AccountChallengeWidgetListParams struct {
	// Direction to order widgets.
	Direction param.Field[AccountChallengeWidgetListParamsDirection] `query:"direction"`
	// Field to order widgets by.
	Order param.Field[AccountChallengeWidgetListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of items per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [AccountChallengeWidgetListParams]'s query parameters as
// `url.Values`.
func (r AccountChallengeWidgetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order widgets.
type AccountChallengeWidgetListParamsDirection string

const (
	AccountChallengeWidgetListParamsDirectionAsc  AccountChallengeWidgetListParamsDirection = "asc"
	AccountChallengeWidgetListParamsDirectionDesc AccountChallengeWidgetListParamsDirection = "desc"
)

// Field to order widgets by.
type AccountChallengeWidgetListParamsOrder string

const (
	AccountChallengeWidgetListParamsOrderID         AccountChallengeWidgetListParamsOrder = "id"
	AccountChallengeWidgetListParamsOrderSitekey    AccountChallengeWidgetListParamsOrder = "sitekey"
	AccountChallengeWidgetListParamsOrderName       AccountChallengeWidgetListParamsOrder = "name"
	AccountChallengeWidgetListParamsOrderCreatedOn  AccountChallengeWidgetListParamsOrder = "created_on"
	AccountChallengeWidgetListParamsOrderModifiedOn AccountChallengeWidgetListParamsOrder = "modified_on"
)

type AccountChallengeWidgetRotateSecretParams struct {
	// If `invalidate_immediately` is set to `false`, the previous secret will remain
	// valid for two hours. Otherwise, the secret is immediately invalidated, and
	// requests using it will be rejected.
	InvalidateImmediately param.Field[bool] `json:"invalidate_immediately"`
}

func (r AccountChallengeWidgetRotateSecretParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
