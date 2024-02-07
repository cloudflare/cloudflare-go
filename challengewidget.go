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

// ChallengeWidgetService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewChallengeWidgetService] method
// instead.
type ChallengeWidgetService struct {
	Options []option.RequestOption
}

// NewChallengeWidgetService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewChallengeWidgetService(opts ...option.RequestOption) (r *ChallengeWidgetService) {
	r = &ChallengeWidgetService{}
	r.Options = opts
	return
}

// Lists challenge widgets.
func (r *ChallengeWidgetService) New(ctx context.Context, accountIdentifier string, params ChallengeWidgetNewParams, opts ...option.RequestOption) (res *ChallengeWidgetNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/challenges/widgets", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Show a single challenge widget configuration.
func (r *ChallengeWidgetService) Get(ctx context.Context, accountIdentifier string, sitekey string, opts ...option.RequestOption) (res *ChallengeWidgetGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/challenges/widgets/%s", accountIdentifier, sitekey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the configuration of a widget.
func (r *ChallengeWidgetService) Update(ctx context.Context, accountIdentifier string, sitekey string, body ChallengeWidgetUpdateParams, opts ...option.RequestOption) (res *ChallengeWidgetUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/challenges/widgets/%s", accountIdentifier, sitekey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists all turnstile widgets of an account.
func (r *ChallengeWidgetService) List(ctx context.Context, accountIdentifier string, query ChallengeWidgetListParams, opts ...option.RequestOption) (res *ChallengeWidgetListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/challenges/widgets", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Destroy a Turnstile Widget.
func (r *ChallengeWidgetService) Delete(ctx context.Context, accountIdentifier string, sitekey string, opts ...option.RequestOption) (res *ChallengeWidgetDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/challenges/widgets/%s", accountIdentifier, sitekey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Generate a new secret key for this widget. If `invalidate_immediately` is set to
// `false`, the previous secret remains valid for 2 hours.
//
// Note that secrets cannot be rotated again during the grace period.
func (r *ChallengeWidgetService) RotateSecret(ctx context.Context, accountIdentifier string, sitekey string, body ChallengeWidgetRotateSecretParams, opts ...option.RequestOption) (res *ChallengeWidgetRotateSecretResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/challenges/widgets/%s/rotate_secret", accountIdentifier, sitekey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ChallengeWidgetNewResponse struct {
	Errors   []ChallengeWidgetNewResponseError   `json:"errors"`
	Messages []ChallengeWidgetNewResponseMessage `json:"messages"`
	// A Turnstile widget's detailed configuration
	Result     ChallengeWidgetNewResponseResult     `json:"result"`
	ResultInfo ChallengeWidgetNewResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success bool                           `json:"success"`
	JSON    challengeWidgetNewResponseJSON `json:"-"`
}

// challengeWidgetNewResponseJSON contains the JSON metadata for the struct
// [ChallengeWidgetNewResponse]
type challengeWidgetNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChallengeWidgetNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ChallengeWidgetNewResponseError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    challengeWidgetNewResponseErrorJSON `json:"-"`
}

// challengeWidgetNewResponseErrorJSON contains the JSON metadata for the struct
// [ChallengeWidgetNewResponseError]
type challengeWidgetNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChallengeWidgetNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ChallengeWidgetNewResponseMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    challengeWidgetNewResponseMessageJSON `json:"-"`
}

// challengeWidgetNewResponseMessageJSON contains the JSON metadata for the struct
// [ChallengeWidgetNewResponseMessage]
type challengeWidgetNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChallengeWidgetNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Turnstile widget's detailed configuration
type ChallengeWidgetNewResponseResult struct {
	// If bot_fight_mode is set to `true`, Cloudflare issues computationally expensive
	// challenges in response to malicious bots (ENT only).
	BotFightMode bool `json:"bot_fight_mode,required"`
	// If Turnstile is embedded on a Cloudflare site and the widget should grant
	// challenge clearance, this setting can determine the clearance level to be set
	ClearanceLevel ChallengeWidgetNewResponseResultClearanceLevel `json:"clearance_level,required"`
	// When the widget was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	Domains   []string  `json:"domains,required"`
	// Widget Mode
	Mode ChallengeWidgetNewResponseResultMode `json:"mode,required"`
	// When the widget was modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Human readable widget name. Not unique. Cloudflare suggests that you set this to
	// a meaningful string to make it easier to identify your widget, and where it is
	// used.
	Name string `json:"name,required"`
	// Do not show any Cloudflare branding on the widget (ENT only).
	Offlabel bool `json:"offlabel,required"`
	// Region where this widget can be used.
	Region ChallengeWidgetNewResponseResultRegion `json:"region,required"`
	// Secret key for this widget.
	Secret string `json:"secret,required"`
	// Widget item identifier tag.
	Sitekey string                               `json:"sitekey,required"`
	JSON    challengeWidgetNewResponseResultJSON `json:"-"`
}

// challengeWidgetNewResponseResultJSON contains the JSON metadata for the struct
// [ChallengeWidgetNewResponseResult]
type challengeWidgetNewResponseResultJSON struct {
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

func (r *ChallengeWidgetNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// If Turnstile is embedded on a Cloudflare site and the widget should grant
// challenge clearance, this setting can determine the clearance level to be set
type ChallengeWidgetNewResponseResultClearanceLevel string

const (
	ChallengeWidgetNewResponseResultClearanceLevelNoClearance ChallengeWidgetNewResponseResultClearanceLevel = "no_clearance"
	ChallengeWidgetNewResponseResultClearanceLevelJschallenge ChallengeWidgetNewResponseResultClearanceLevel = "jschallenge"
	ChallengeWidgetNewResponseResultClearanceLevelManaged     ChallengeWidgetNewResponseResultClearanceLevel = "managed"
	ChallengeWidgetNewResponseResultClearanceLevelInteractive ChallengeWidgetNewResponseResultClearanceLevel = "interactive"
)

// Widget Mode
type ChallengeWidgetNewResponseResultMode string

const (
	ChallengeWidgetNewResponseResultModeNonInteractive ChallengeWidgetNewResponseResultMode = "non-interactive"
	ChallengeWidgetNewResponseResultModeInvisible      ChallengeWidgetNewResponseResultMode = "invisible"
	ChallengeWidgetNewResponseResultModeManaged        ChallengeWidgetNewResponseResultMode = "managed"
)

// Region where this widget can be used.
type ChallengeWidgetNewResponseResultRegion string

const (
	ChallengeWidgetNewResponseResultRegionWorld ChallengeWidgetNewResponseResultRegion = "world"
)

type ChallengeWidgetNewResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count,required"`
	// Current page within paginated list of results
	Page float64 `json:"page,required"`
	// Number of results per page of results
	PerPage float64 `json:"per_page,required"`
	// Total results available without any search parameters
	TotalCount float64                                  `json:"total_count,required"`
	JSON       challengeWidgetNewResponseResultInfoJSON `json:"-"`
}

// challengeWidgetNewResponseResultInfoJSON contains the JSON metadata for the
// struct [ChallengeWidgetNewResponseResultInfo]
type challengeWidgetNewResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChallengeWidgetNewResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ChallengeWidgetGetResponse struct {
	Errors   []ChallengeWidgetGetResponseError   `json:"errors"`
	Messages []ChallengeWidgetGetResponseMessage `json:"messages"`
	// A Turnstile widget's detailed configuration
	Result ChallengeWidgetGetResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                           `json:"success"`
	JSON    challengeWidgetGetResponseJSON `json:"-"`
}

// challengeWidgetGetResponseJSON contains the JSON metadata for the struct
// [ChallengeWidgetGetResponse]
type challengeWidgetGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChallengeWidgetGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ChallengeWidgetGetResponseError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    challengeWidgetGetResponseErrorJSON `json:"-"`
}

// challengeWidgetGetResponseErrorJSON contains the JSON metadata for the struct
// [ChallengeWidgetGetResponseError]
type challengeWidgetGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChallengeWidgetGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ChallengeWidgetGetResponseMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    challengeWidgetGetResponseMessageJSON `json:"-"`
}

// challengeWidgetGetResponseMessageJSON contains the JSON metadata for the struct
// [ChallengeWidgetGetResponseMessage]
type challengeWidgetGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChallengeWidgetGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Turnstile widget's detailed configuration
type ChallengeWidgetGetResponseResult struct {
	// If bot_fight_mode is set to `true`, Cloudflare issues computationally expensive
	// challenges in response to malicious bots (ENT only).
	BotFightMode bool `json:"bot_fight_mode,required"`
	// If Turnstile is embedded on a Cloudflare site and the widget should grant
	// challenge clearance, this setting can determine the clearance level to be set
	ClearanceLevel ChallengeWidgetGetResponseResultClearanceLevel `json:"clearance_level,required"`
	// When the widget was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	Domains   []string  `json:"domains,required"`
	// Widget Mode
	Mode ChallengeWidgetGetResponseResultMode `json:"mode,required"`
	// When the widget was modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Human readable widget name. Not unique. Cloudflare suggests that you set this to
	// a meaningful string to make it easier to identify your widget, and where it is
	// used.
	Name string `json:"name,required"`
	// Do not show any Cloudflare branding on the widget (ENT only).
	Offlabel bool `json:"offlabel,required"`
	// Region where this widget can be used.
	Region ChallengeWidgetGetResponseResultRegion `json:"region,required"`
	// Secret key for this widget.
	Secret string `json:"secret,required"`
	// Widget item identifier tag.
	Sitekey string                               `json:"sitekey,required"`
	JSON    challengeWidgetGetResponseResultJSON `json:"-"`
}

// challengeWidgetGetResponseResultJSON contains the JSON metadata for the struct
// [ChallengeWidgetGetResponseResult]
type challengeWidgetGetResponseResultJSON struct {
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

func (r *ChallengeWidgetGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// If Turnstile is embedded on a Cloudflare site and the widget should grant
// challenge clearance, this setting can determine the clearance level to be set
type ChallengeWidgetGetResponseResultClearanceLevel string

const (
	ChallengeWidgetGetResponseResultClearanceLevelNoClearance ChallengeWidgetGetResponseResultClearanceLevel = "no_clearance"
	ChallengeWidgetGetResponseResultClearanceLevelJschallenge ChallengeWidgetGetResponseResultClearanceLevel = "jschallenge"
	ChallengeWidgetGetResponseResultClearanceLevelManaged     ChallengeWidgetGetResponseResultClearanceLevel = "managed"
	ChallengeWidgetGetResponseResultClearanceLevelInteractive ChallengeWidgetGetResponseResultClearanceLevel = "interactive"
)

// Widget Mode
type ChallengeWidgetGetResponseResultMode string

const (
	ChallengeWidgetGetResponseResultModeNonInteractive ChallengeWidgetGetResponseResultMode = "non-interactive"
	ChallengeWidgetGetResponseResultModeInvisible      ChallengeWidgetGetResponseResultMode = "invisible"
	ChallengeWidgetGetResponseResultModeManaged        ChallengeWidgetGetResponseResultMode = "managed"
)

// Region where this widget can be used.
type ChallengeWidgetGetResponseResultRegion string

const (
	ChallengeWidgetGetResponseResultRegionWorld ChallengeWidgetGetResponseResultRegion = "world"
)

type ChallengeWidgetUpdateResponse struct {
	Errors   []ChallengeWidgetUpdateResponseError   `json:"errors"`
	Messages []ChallengeWidgetUpdateResponseMessage `json:"messages"`
	// A Turnstile widget's detailed configuration
	Result ChallengeWidgetUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                              `json:"success"`
	JSON    challengeWidgetUpdateResponseJSON `json:"-"`
}

// challengeWidgetUpdateResponseJSON contains the JSON metadata for the struct
// [ChallengeWidgetUpdateResponse]
type challengeWidgetUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChallengeWidgetUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ChallengeWidgetUpdateResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    challengeWidgetUpdateResponseErrorJSON `json:"-"`
}

// challengeWidgetUpdateResponseErrorJSON contains the JSON metadata for the struct
// [ChallengeWidgetUpdateResponseError]
type challengeWidgetUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChallengeWidgetUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ChallengeWidgetUpdateResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    challengeWidgetUpdateResponseMessageJSON `json:"-"`
}

// challengeWidgetUpdateResponseMessageJSON contains the JSON metadata for the
// struct [ChallengeWidgetUpdateResponseMessage]
type challengeWidgetUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChallengeWidgetUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Turnstile widget's detailed configuration
type ChallengeWidgetUpdateResponseResult struct {
	// If bot_fight_mode is set to `true`, Cloudflare issues computationally expensive
	// challenges in response to malicious bots (ENT only).
	BotFightMode bool `json:"bot_fight_mode,required"`
	// If Turnstile is embedded on a Cloudflare site and the widget should grant
	// challenge clearance, this setting can determine the clearance level to be set
	ClearanceLevel ChallengeWidgetUpdateResponseResultClearanceLevel `json:"clearance_level,required"`
	// When the widget was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	Domains   []string  `json:"domains,required"`
	// Widget Mode
	Mode ChallengeWidgetUpdateResponseResultMode `json:"mode,required"`
	// When the widget was modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Human readable widget name. Not unique. Cloudflare suggests that you set this to
	// a meaningful string to make it easier to identify your widget, and where it is
	// used.
	Name string `json:"name,required"`
	// Do not show any Cloudflare branding on the widget (ENT only).
	Offlabel bool `json:"offlabel,required"`
	// Region where this widget can be used.
	Region ChallengeWidgetUpdateResponseResultRegion `json:"region,required"`
	// Secret key for this widget.
	Secret string `json:"secret,required"`
	// Widget item identifier tag.
	Sitekey string                                  `json:"sitekey,required"`
	JSON    challengeWidgetUpdateResponseResultJSON `json:"-"`
}

// challengeWidgetUpdateResponseResultJSON contains the JSON metadata for the
// struct [ChallengeWidgetUpdateResponseResult]
type challengeWidgetUpdateResponseResultJSON struct {
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

func (r *ChallengeWidgetUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// If Turnstile is embedded on a Cloudflare site and the widget should grant
// challenge clearance, this setting can determine the clearance level to be set
type ChallengeWidgetUpdateResponseResultClearanceLevel string

const (
	ChallengeWidgetUpdateResponseResultClearanceLevelNoClearance ChallengeWidgetUpdateResponseResultClearanceLevel = "no_clearance"
	ChallengeWidgetUpdateResponseResultClearanceLevelJschallenge ChallengeWidgetUpdateResponseResultClearanceLevel = "jschallenge"
	ChallengeWidgetUpdateResponseResultClearanceLevelManaged     ChallengeWidgetUpdateResponseResultClearanceLevel = "managed"
	ChallengeWidgetUpdateResponseResultClearanceLevelInteractive ChallengeWidgetUpdateResponseResultClearanceLevel = "interactive"
)

// Widget Mode
type ChallengeWidgetUpdateResponseResultMode string

const (
	ChallengeWidgetUpdateResponseResultModeNonInteractive ChallengeWidgetUpdateResponseResultMode = "non-interactive"
	ChallengeWidgetUpdateResponseResultModeInvisible      ChallengeWidgetUpdateResponseResultMode = "invisible"
	ChallengeWidgetUpdateResponseResultModeManaged        ChallengeWidgetUpdateResponseResultMode = "managed"
)

// Region where this widget can be used.
type ChallengeWidgetUpdateResponseResultRegion string

const (
	ChallengeWidgetUpdateResponseResultRegionWorld ChallengeWidgetUpdateResponseResultRegion = "world"
)

type ChallengeWidgetListResponse struct {
	Errors     []ChallengeWidgetListResponseError    `json:"errors"`
	Messages   []ChallengeWidgetListResponseMessage  `json:"messages"`
	Result     []ChallengeWidgetListResponseResult   `json:"result"`
	ResultInfo ChallengeWidgetListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success bool                            `json:"success"`
	JSON    challengeWidgetListResponseJSON `json:"-"`
}

// challengeWidgetListResponseJSON contains the JSON metadata for the struct
// [ChallengeWidgetListResponse]
type challengeWidgetListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChallengeWidgetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ChallengeWidgetListResponseError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    challengeWidgetListResponseErrorJSON `json:"-"`
}

// challengeWidgetListResponseErrorJSON contains the JSON metadata for the struct
// [ChallengeWidgetListResponseError]
type challengeWidgetListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChallengeWidgetListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ChallengeWidgetListResponseMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    challengeWidgetListResponseMessageJSON `json:"-"`
}

// challengeWidgetListResponseMessageJSON contains the JSON metadata for the struct
// [ChallengeWidgetListResponseMessage]
type challengeWidgetListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChallengeWidgetListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Turnstile Widgets configuration as it appears in listings
type ChallengeWidgetListResponseResult struct {
	// If bot_fight_mode is set to `true`, Cloudflare issues computationally expensive
	// challenges in response to malicious bots (ENT only).
	BotFightMode bool `json:"bot_fight_mode,required"`
	// If Turnstile is embedded on a Cloudflare site and the widget should grant
	// challenge clearance, this setting can determine the clearance level to be set
	ClearanceLevel ChallengeWidgetListResponseResultClearanceLevel `json:"clearance_level,required"`
	// When the widget was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	Domains   []string  `json:"domains,required"`
	// Widget Mode
	Mode ChallengeWidgetListResponseResultMode `json:"mode,required"`
	// When the widget was modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Human readable widget name. Not unique. Cloudflare suggests that you set this to
	// a meaningful string to make it easier to identify your widget, and where it is
	// used.
	Name string `json:"name,required"`
	// Do not show any Cloudflare branding on the widget (ENT only).
	Offlabel bool `json:"offlabel,required"`
	// Region where this widget can be used.
	Region ChallengeWidgetListResponseResultRegion `json:"region,required"`
	// Widget item identifier tag.
	Sitekey string                                `json:"sitekey,required"`
	JSON    challengeWidgetListResponseResultJSON `json:"-"`
}

// challengeWidgetListResponseResultJSON contains the JSON metadata for the struct
// [ChallengeWidgetListResponseResult]
type challengeWidgetListResponseResultJSON struct {
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

func (r *ChallengeWidgetListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// If Turnstile is embedded on a Cloudflare site and the widget should grant
// challenge clearance, this setting can determine the clearance level to be set
type ChallengeWidgetListResponseResultClearanceLevel string

const (
	ChallengeWidgetListResponseResultClearanceLevelNoClearance ChallengeWidgetListResponseResultClearanceLevel = "no_clearance"
	ChallengeWidgetListResponseResultClearanceLevelJschallenge ChallengeWidgetListResponseResultClearanceLevel = "jschallenge"
	ChallengeWidgetListResponseResultClearanceLevelManaged     ChallengeWidgetListResponseResultClearanceLevel = "managed"
	ChallengeWidgetListResponseResultClearanceLevelInteractive ChallengeWidgetListResponseResultClearanceLevel = "interactive"
)

// Widget Mode
type ChallengeWidgetListResponseResultMode string

const (
	ChallengeWidgetListResponseResultModeNonInteractive ChallengeWidgetListResponseResultMode = "non-interactive"
	ChallengeWidgetListResponseResultModeInvisible      ChallengeWidgetListResponseResultMode = "invisible"
	ChallengeWidgetListResponseResultModeManaged        ChallengeWidgetListResponseResultMode = "managed"
)

// Region where this widget can be used.
type ChallengeWidgetListResponseResultRegion string

const (
	ChallengeWidgetListResponseResultRegionWorld ChallengeWidgetListResponseResultRegion = "world"
)

type ChallengeWidgetListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count,required"`
	// Current page within paginated list of results
	Page float64 `json:"page,required"`
	// Number of results per page of results
	PerPage float64 `json:"per_page,required"`
	// Total results available without any search parameters
	TotalCount float64                                   `json:"total_count,required"`
	JSON       challengeWidgetListResponseResultInfoJSON `json:"-"`
}

// challengeWidgetListResponseResultInfoJSON contains the JSON metadata for the
// struct [ChallengeWidgetListResponseResultInfo]
type challengeWidgetListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChallengeWidgetListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ChallengeWidgetDeleteResponse struct {
	Errors   []ChallengeWidgetDeleteResponseError   `json:"errors"`
	Messages []ChallengeWidgetDeleteResponseMessage `json:"messages"`
	// A Turnstile widget's detailed configuration
	Result ChallengeWidgetDeleteResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                              `json:"success"`
	JSON    challengeWidgetDeleteResponseJSON `json:"-"`
}

// challengeWidgetDeleteResponseJSON contains the JSON metadata for the struct
// [ChallengeWidgetDeleteResponse]
type challengeWidgetDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChallengeWidgetDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ChallengeWidgetDeleteResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    challengeWidgetDeleteResponseErrorJSON `json:"-"`
}

// challengeWidgetDeleteResponseErrorJSON contains the JSON metadata for the struct
// [ChallengeWidgetDeleteResponseError]
type challengeWidgetDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChallengeWidgetDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ChallengeWidgetDeleteResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    challengeWidgetDeleteResponseMessageJSON `json:"-"`
}

// challengeWidgetDeleteResponseMessageJSON contains the JSON metadata for the
// struct [ChallengeWidgetDeleteResponseMessage]
type challengeWidgetDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChallengeWidgetDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Turnstile widget's detailed configuration
type ChallengeWidgetDeleteResponseResult struct {
	// If bot_fight_mode is set to `true`, Cloudflare issues computationally expensive
	// challenges in response to malicious bots (ENT only).
	BotFightMode bool `json:"bot_fight_mode,required"`
	// If Turnstile is embedded on a Cloudflare site and the widget should grant
	// challenge clearance, this setting can determine the clearance level to be set
	ClearanceLevel ChallengeWidgetDeleteResponseResultClearanceLevel `json:"clearance_level,required"`
	// When the widget was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	Domains   []string  `json:"domains,required"`
	// Widget Mode
	Mode ChallengeWidgetDeleteResponseResultMode `json:"mode,required"`
	// When the widget was modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Human readable widget name. Not unique. Cloudflare suggests that you set this to
	// a meaningful string to make it easier to identify your widget, and where it is
	// used.
	Name string `json:"name,required"`
	// Do not show any Cloudflare branding on the widget (ENT only).
	Offlabel bool `json:"offlabel,required"`
	// Region where this widget can be used.
	Region ChallengeWidgetDeleteResponseResultRegion `json:"region,required"`
	// Secret key for this widget.
	Secret string `json:"secret,required"`
	// Widget item identifier tag.
	Sitekey string                                  `json:"sitekey,required"`
	JSON    challengeWidgetDeleteResponseResultJSON `json:"-"`
}

// challengeWidgetDeleteResponseResultJSON contains the JSON metadata for the
// struct [ChallengeWidgetDeleteResponseResult]
type challengeWidgetDeleteResponseResultJSON struct {
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

func (r *ChallengeWidgetDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// If Turnstile is embedded on a Cloudflare site and the widget should grant
// challenge clearance, this setting can determine the clearance level to be set
type ChallengeWidgetDeleteResponseResultClearanceLevel string

const (
	ChallengeWidgetDeleteResponseResultClearanceLevelNoClearance ChallengeWidgetDeleteResponseResultClearanceLevel = "no_clearance"
	ChallengeWidgetDeleteResponseResultClearanceLevelJschallenge ChallengeWidgetDeleteResponseResultClearanceLevel = "jschallenge"
	ChallengeWidgetDeleteResponseResultClearanceLevelManaged     ChallengeWidgetDeleteResponseResultClearanceLevel = "managed"
	ChallengeWidgetDeleteResponseResultClearanceLevelInteractive ChallengeWidgetDeleteResponseResultClearanceLevel = "interactive"
)

// Widget Mode
type ChallengeWidgetDeleteResponseResultMode string

const (
	ChallengeWidgetDeleteResponseResultModeNonInteractive ChallengeWidgetDeleteResponseResultMode = "non-interactive"
	ChallengeWidgetDeleteResponseResultModeInvisible      ChallengeWidgetDeleteResponseResultMode = "invisible"
	ChallengeWidgetDeleteResponseResultModeManaged        ChallengeWidgetDeleteResponseResultMode = "managed"
)

// Region where this widget can be used.
type ChallengeWidgetDeleteResponseResultRegion string

const (
	ChallengeWidgetDeleteResponseResultRegionWorld ChallengeWidgetDeleteResponseResultRegion = "world"
)

type ChallengeWidgetRotateSecretResponse struct {
	Errors   []ChallengeWidgetRotateSecretResponseError   `json:"errors"`
	Messages []ChallengeWidgetRotateSecretResponseMessage `json:"messages"`
	// A Turnstile widget's detailed configuration
	Result ChallengeWidgetRotateSecretResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                                    `json:"success"`
	JSON    challengeWidgetRotateSecretResponseJSON `json:"-"`
}

// challengeWidgetRotateSecretResponseJSON contains the JSON metadata for the
// struct [ChallengeWidgetRotateSecretResponse]
type challengeWidgetRotateSecretResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChallengeWidgetRotateSecretResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ChallengeWidgetRotateSecretResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    challengeWidgetRotateSecretResponseErrorJSON `json:"-"`
}

// challengeWidgetRotateSecretResponseErrorJSON contains the JSON metadata for the
// struct [ChallengeWidgetRotateSecretResponseError]
type challengeWidgetRotateSecretResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChallengeWidgetRotateSecretResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ChallengeWidgetRotateSecretResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    challengeWidgetRotateSecretResponseMessageJSON `json:"-"`
}

// challengeWidgetRotateSecretResponseMessageJSON contains the JSON metadata for
// the struct [ChallengeWidgetRotateSecretResponseMessage]
type challengeWidgetRotateSecretResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChallengeWidgetRotateSecretResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Turnstile widget's detailed configuration
type ChallengeWidgetRotateSecretResponseResult struct {
	// If bot_fight_mode is set to `true`, Cloudflare issues computationally expensive
	// challenges in response to malicious bots (ENT only).
	BotFightMode bool `json:"bot_fight_mode,required"`
	// If Turnstile is embedded on a Cloudflare site and the widget should grant
	// challenge clearance, this setting can determine the clearance level to be set
	ClearanceLevel ChallengeWidgetRotateSecretResponseResultClearanceLevel `json:"clearance_level,required"`
	// When the widget was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	Domains   []string  `json:"domains,required"`
	// Widget Mode
	Mode ChallengeWidgetRotateSecretResponseResultMode `json:"mode,required"`
	// When the widget was modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Human readable widget name. Not unique. Cloudflare suggests that you set this to
	// a meaningful string to make it easier to identify your widget, and where it is
	// used.
	Name string `json:"name,required"`
	// Do not show any Cloudflare branding on the widget (ENT only).
	Offlabel bool `json:"offlabel,required"`
	// Region where this widget can be used.
	Region ChallengeWidgetRotateSecretResponseResultRegion `json:"region,required"`
	// Secret key for this widget.
	Secret string `json:"secret,required"`
	// Widget item identifier tag.
	Sitekey string                                        `json:"sitekey,required"`
	JSON    challengeWidgetRotateSecretResponseResultJSON `json:"-"`
}

// challengeWidgetRotateSecretResponseResultJSON contains the JSON metadata for the
// struct [ChallengeWidgetRotateSecretResponseResult]
type challengeWidgetRotateSecretResponseResultJSON struct {
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

func (r *ChallengeWidgetRotateSecretResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// If Turnstile is embedded on a Cloudflare site and the widget should grant
// challenge clearance, this setting can determine the clearance level to be set
type ChallengeWidgetRotateSecretResponseResultClearanceLevel string

const (
	ChallengeWidgetRotateSecretResponseResultClearanceLevelNoClearance ChallengeWidgetRotateSecretResponseResultClearanceLevel = "no_clearance"
	ChallengeWidgetRotateSecretResponseResultClearanceLevelJschallenge ChallengeWidgetRotateSecretResponseResultClearanceLevel = "jschallenge"
	ChallengeWidgetRotateSecretResponseResultClearanceLevelManaged     ChallengeWidgetRotateSecretResponseResultClearanceLevel = "managed"
	ChallengeWidgetRotateSecretResponseResultClearanceLevelInteractive ChallengeWidgetRotateSecretResponseResultClearanceLevel = "interactive"
)

// Widget Mode
type ChallengeWidgetRotateSecretResponseResultMode string

const (
	ChallengeWidgetRotateSecretResponseResultModeNonInteractive ChallengeWidgetRotateSecretResponseResultMode = "non-interactive"
	ChallengeWidgetRotateSecretResponseResultModeInvisible      ChallengeWidgetRotateSecretResponseResultMode = "invisible"
	ChallengeWidgetRotateSecretResponseResultModeManaged        ChallengeWidgetRotateSecretResponseResultMode = "managed"
)

// Region where this widget can be used.
type ChallengeWidgetRotateSecretResponseResultRegion string

const (
	ChallengeWidgetRotateSecretResponseResultRegionWorld ChallengeWidgetRotateSecretResponseResultRegion = "world"
)

type ChallengeWidgetNewParams struct {
	Domains param.Field[[]string] `json:"domains,required"`
	// Widget Mode
	Mode param.Field[ChallengeWidgetNewParamsMode] `json:"mode,required"`
	// Human readable widget name. Not unique. Cloudflare suggests that you set this to
	// a meaningful string to make it easier to identify your widget, and where it is
	// used.
	Name param.Field[string] `json:"name,required"`
	// Direction to order widgets.
	Direction param.Field[ChallengeWidgetNewParamsDirection] `query:"direction"`
	// Field to order widgets by.
	Order param.Field[ChallengeWidgetNewParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of items per page.
	PerPage param.Field[float64] `query:"per_page"`
	// If bot_fight_mode is set to `true`, Cloudflare issues computationally expensive
	// challenges in response to malicious bots (ENT only).
	BotFightMode param.Field[bool] `json:"bot_fight_mode"`
	// If Turnstile is embedded on a Cloudflare site and the widget should grant
	// challenge clearance, this setting can determine the clearance level to be set
	ClearanceLevel param.Field[ChallengeWidgetNewParamsClearanceLevel] `json:"clearance_level"`
	// Do not show any Cloudflare branding on the widget (ENT only).
	Offlabel param.Field[bool] `json:"offlabel"`
	// Region where this widget can be used.
	Region param.Field[ChallengeWidgetNewParamsRegion] `json:"region"`
}

func (r ChallengeWidgetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [ChallengeWidgetNewParams]'s query parameters as
// `url.Values`.
func (r ChallengeWidgetNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Widget Mode
type ChallengeWidgetNewParamsMode string

const (
	ChallengeWidgetNewParamsModeNonInteractive ChallengeWidgetNewParamsMode = "non-interactive"
	ChallengeWidgetNewParamsModeInvisible      ChallengeWidgetNewParamsMode = "invisible"
	ChallengeWidgetNewParamsModeManaged        ChallengeWidgetNewParamsMode = "managed"
)

// Direction to order widgets.
type ChallengeWidgetNewParamsDirection string

const (
	ChallengeWidgetNewParamsDirectionAsc  ChallengeWidgetNewParamsDirection = "asc"
	ChallengeWidgetNewParamsDirectionDesc ChallengeWidgetNewParamsDirection = "desc"
)

// Field to order widgets by.
type ChallengeWidgetNewParamsOrder string

const (
	ChallengeWidgetNewParamsOrderID         ChallengeWidgetNewParamsOrder = "id"
	ChallengeWidgetNewParamsOrderSitekey    ChallengeWidgetNewParamsOrder = "sitekey"
	ChallengeWidgetNewParamsOrderName       ChallengeWidgetNewParamsOrder = "name"
	ChallengeWidgetNewParamsOrderCreatedOn  ChallengeWidgetNewParamsOrder = "created_on"
	ChallengeWidgetNewParamsOrderModifiedOn ChallengeWidgetNewParamsOrder = "modified_on"
)

// If Turnstile is embedded on a Cloudflare site and the widget should grant
// challenge clearance, this setting can determine the clearance level to be set
type ChallengeWidgetNewParamsClearanceLevel string

const (
	ChallengeWidgetNewParamsClearanceLevelNoClearance ChallengeWidgetNewParamsClearanceLevel = "no_clearance"
	ChallengeWidgetNewParamsClearanceLevelJschallenge ChallengeWidgetNewParamsClearanceLevel = "jschallenge"
	ChallengeWidgetNewParamsClearanceLevelManaged     ChallengeWidgetNewParamsClearanceLevel = "managed"
	ChallengeWidgetNewParamsClearanceLevelInteractive ChallengeWidgetNewParamsClearanceLevel = "interactive"
)

// Region where this widget can be used.
type ChallengeWidgetNewParamsRegion string

const (
	ChallengeWidgetNewParamsRegionWorld ChallengeWidgetNewParamsRegion = "world"
)

type ChallengeWidgetUpdateParams struct {
	Domains param.Field[[]string] `json:"domains,required"`
	// Widget Mode
	Mode param.Field[ChallengeWidgetUpdateParamsMode] `json:"mode,required"`
	// Human readable widget name. Not unique. Cloudflare suggests that you set this to
	// a meaningful string to make it easier to identify your widget, and where it is
	// used.
	Name param.Field[string] `json:"name,required"`
	// If bot_fight_mode is set to `true`, Cloudflare issues computationally expensive
	// challenges in response to malicious bots (ENT only).
	BotFightMode param.Field[bool] `json:"bot_fight_mode"`
	// If Turnstile is embedded on a Cloudflare site and the widget should grant
	// challenge clearance, this setting can determine the clearance level to be set
	ClearanceLevel param.Field[ChallengeWidgetUpdateParamsClearanceLevel] `json:"clearance_level"`
	// Do not show any Cloudflare branding on the widget (ENT only).
	Offlabel param.Field[bool] `json:"offlabel"`
}

func (r ChallengeWidgetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Widget Mode
type ChallengeWidgetUpdateParamsMode string

const (
	ChallengeWidgetUpdateParamsModeNonInteractive ChallengeWidgetUpdateParamsMode = "non-interactive"
	ChallengeWidgetUpdateParamsModeInvisible      ChallengeWidgetUpdateParamsMode = "invisible"
	ChallengeWidgetUpdateParamsModeManaged        ChallengeWidgetUpdateParamsMode = "managed"
)

// If Turnstile is embedded on a Cloudflare site and the widget should grant
// challenge clearance, this setting can determine the clearance level to be set
type ChallengeWidgetUpdateParamsClearanceLevel string

const (
	ChallengeWidgetUpdateParamsClearanceLevelNoClearance ChallengeWidgetUpdateParamsClearanceLevel = "no_clearance"
	ChallengeWidgetUpdateParamsClearanceLevelJschallenge ChallengeWidgetUpdateParamsClearanceLevel = "jschallenge"
	ChallengeWidgetUpdateParamsClearanceLevelManaged     ChallengeWidgetUpdateParamsClearanceLevel = "managed"
	ChallengeWidgetUpdateParamsClearanceLevelInteractive ChallengeWidgetUpdateParamsClearanceLevel = "interactive"
)

type ChallengeWidgetListParams struct {
	// Direction to order widgets.
	Direction param.Field[ChallengeWidgetListParamsDirection] `query:"direction"`
	// Field to order widgets by.
	Order param.Field[ChallengeWidgetListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of items per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [ChallengeWidgetListParams]'s query parameters as
// `url.Values`.
func (r ChallengeWidgetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order widgets.
type ChallengeWidgetListParamsDirection string

const (
	ChallengeWidgetListParamsDirectionAsc  ChallengeWidgetListParamsDirection = "asc"
	ChallengeWidgetListParamsDirectionDesc ChallengeWidgetListParamsDirection = "desc"
)

// Field to order widgets by.
type ChallengeWidgetListParamsOrder string

const (
	ChallengeWidgetListParamsOrderID         ChallengeWidgetListParamsOrder = "id"
	ChallengeWidgetListParamsOrderSitekey    ChallengeWidgetListParamsOrder = "sitekey"
	ChallengeWidgetListParamsOrderName       ChallengeWidgetListParamsOrder = "name"
	ChallengeWidgetListParamsOrderCreatedOn  ChallengeWidgetListParamsOrder = "created_on"
	ChallengeWidgetListParamsOrderModifiedOn ChallengeWidgetListParamsOrder = "modified_on"
)

type ChallengeWidgetRotateSecretParams struct {
	// If `invalidate_immediately` is set to `false`, the previous secret will remain
	// valid for two hours. Otherwise, the secret is immediately invalidated, and
	// requests using it will be rejected.
	InvalidateImmediately param.Field[bool] `json:"invalidate_immediately"`
}

func (r ChallengeWidgetRotateSecretParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
