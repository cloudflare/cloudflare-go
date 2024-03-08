// File generated from our OpenAPI spec by Stainless.

package challenges

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// WidgetService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewWidgetService] method instead.
type WidgetService struct {
	Options []option.RequestOption
}

// NewWidgetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWidgetService(opts ...option.RequestOption) (r *WidgetService) {
	r = &WidgetService{}
	r.Options = opts
	return
}

// Lists challenge widgets.
func (r *WidgetService) New(ctx context.Context, accountIdentifier string, params WidgetNewParams, opts ...option.RequestOption) (res *NcChallengesAdminWidgetDetail, err error) {
	opts = append(r.Options[:], opts...)
	var env WidgetNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/challenges/widgets", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update the configuration of a widget.
func (r *WidgetService) Update(ctx context.Context, accountIdentifier string, sitekey string, body WidgetUpdateParams, opts ...option.RequestOption) (res *NcChallengesAdminWidgetDetail, err error) {
	opts = append(r.Options[:], opts...)
	var env WidgetUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/challenges/widgets/%s", accountIdentifier, sitekey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all turnstile widgets of an account.
func (r *WidgetService) List(ctx context.Context, accountIdentifier string, query WidgetListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[NcChallengesAdminWidgetList], err error) {
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

// Lists all turnstile widgets of an account.
func (r *WidgetService) ListAutoPaging(ctx context.Context, accountIdentifier string, query WidgetListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[NcChallengesAdminWidgetList] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, accountIdentifier, query, opts...))
}

// Destroy a Turnstile Widget.
func (r *WidgetService) Delete(ctx context.Context, accountIdentifier string, sitekey string, opts ...option.RequestOption) (res *NcChallengesAdminWidgetDetail, err error) {
	opts = append(r.Options[:], opts...)
	var env WidgetDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/challenges/widgets/%s", accountIdentifier, sitekey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Show a single challenge widget configuration.
func (r *WidgetService) Get(ctx context.Context, accountIdentifier string, sitekey string, opts ...option.RequestOption) (res *NcChallengesAdminWidgetDetail, err error) {
	opts = append(r.Options[:], opts...)
	var env WidgetGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/challenges/widgets/%s", accountIdentifier, sitekey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Generate a new secret key for this widget. If `invalidate_immediately` is set to
// `false`, the previous secret remains valid for 2 hours.
//
// Note that secrets cannot be rotated again during the grace period.
func (r *WidgetService) RotateSecret(ctx context.Context, accountIdentifier string, sitekey string, body WidgetRotateSecretParams, opts ...option.RequestOption) (res *NcChallengesAdminWidgetDetail, err error) {
	opts = append(r.Options[:], opts...)
	var env WidgetRotateSecretResponseEnvelope
	path := fmt.Sprintf("accounts/%s/challenges/widgets/%s/rotate_secret", accountIdentifier, sitekey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A Turnstile widget's detailed configuration
type NcChallengesAdminWidgetDetail struct {
	// If bot_fight_mode is set to `true`, Cloudflare issues computationally expensive
	// challenges in response to malicious bots (ENT only).
	BotFightMode bool `json:"bot_fight_mode,required"`
	// If Turnstile is embedded on a Cloudflare site and the widget should grant
	// challenge clearance, this setting can determine the clearance level to be set
	ClearanceLevel NcChallengesAdminWidgetDetailClearanceLevel `json:"clearance_level,required"`
	// When the widget was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	Domains   []string  `json:"domains,required"`
	// Widget Mode
	Mode NcChallengesAdminWidgetDetailMode `json:"mode,required"`
	// When the widget was modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Human readable widget name. Not unique. Cloudflare suggests that you set this to
	// a meaningful string to make it easier to identify your widget, and where it is
	// used.
	Name string `json:"name,required"`
	// Do not show any Cloudflare branding on the widget (ENT only).
	Offlabel bool `json:"offlabel,required"`
	// Region where this widget can be used.
	Region NcChallengesAdminWidgetDetailRegion `json:"region,required"`
	// Secret key for this widget.
	Secret string `json:"secret,required"`
	// Widget item identifier tag.
	Sitekey string                            `json:"sitekey,required"`
	JSON    ncChallengesAdminWidgetDetailJSON `json:"-"`
}

// ncChallengesAdminWidgetDetailJSON contains the JSON metadata for the struct
// [NcChallengesAdminWidgetDetail]
type ncChallengesAdminWidgetDetailJSON struct {
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

func (r *NcChallengesAdminWidgetDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ncChallengesAdminWidgetDetailJSON) RawJSON() string {
	return r.raw
}

// If Turnstile is embedded on a Cloudflare site and the widget should grant
// challenge clearance, this setting can determine the clearance level to be set
type NcChallengesAdminWidgetDetailClearanceLevel string

const (
	NcChallengesAdminWidgetDetailClearanceLevelNoClearance NcChallengesAdminWidgetDetailClearanceLevel = "no_clearance"
	NcChallengesAdminWidgetDetailClearanceLevelJschallenge NcChallengesAdminWidgetDetailClearanceLevel = "jschallenge"
	NcChallengesAdminWidgetDetailClearanceLevelManaged     NcChallengesAdminWidgetDetailClearanceLevel = "managed"
	NcChallengesAdminWidgetDetailClearanceLevelInteractive NcChallengesAdminWidgetDetailClearanceLevel = "interactive"
)

// Widget Mode
type NcChallengesAdminWidgetDetailMode string

const (
	NcChallengesAdminWidgetDetailModeNonInteractive NcChallengesAdminWidgetDetailMode = "non-interactive"
	NcChallengesAdminWidgetDetailModeInvisible      NcChallengesAdminWidgetDetailMode = "invisible"
	NcChallengesAdminWidgetDetailModeManaged        NcChallengesAdminWidgetDetailMode = "managed"
)

// Region where this widget can be used.
type NcChallengesAdminWidgetDetailRegion string

const (
	NcChallengesAdminWidgetDetailRegionWorld NcChallengesAdminWidgetDetailRegion = "world"
)

// A Turnstile Widgets configuration as it appears in listings
type NcChallengesAdminWidgetList struct {
	// If bot_fight_mode is set to `true`, Cloudflare issues computationally expensive
	// challenges in response to malicious bots (ENT only).
	BotFightMode bool `json:"bot_fight_mode,required"`
	// If Turnstile is embedded on a Cloudflare site and the widget should grant
	// challenge clearance, this setting can determine the clearance level to be set
	ClearanceLevel NcChallengesAdminWidgetListClearanceLevel `json:"clearance_level,required"`
	// When the widget was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	Domains   []string  `json:"domains,required"`
	// Widget Mode
	Mode NcChallengesAdminWidgetListMode `json:"mode,required"`
	// When the widget was modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Human readable widget name. Not unique. Cloudflare suggests that you set this to
	// a meaningful string to make it easier to identify your widget, and where it is
	// used.
	Name string `json:"name,required"`
	// Do not show any Cloudflare branding on the widget (ENT only).
	Offlabel bool `json:"offlabel,required"`
	// Region where this widget can be used.
	Region NcChallengesAdminWidgetListRegion `json:"region,required"`
	// Widget item identifier tag.
	Sitekey string                          `json:"sitekey,required"`
	JSON    ncChallengesAdminWidgetListJSON `json:"-"`
}

// ncChallengesAdminWidgetListJSON contains the JSON metadata for the struct
// [NcChallengesAdminWidgetList]
type ncChallengesAdminWidgetListJSON struct {
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

func (r *NcChallengesAdminWidgetList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ncChallengesAdminWidgetListJSON) RawJSON() string {
	return r.raw
}

// If Turnstile is embedded on a Cloudflare site and the widget should grant
// challenge clearance, this setting can determine the clearance level to be set
type NcChallengesAdminWidgetListClearanceLevel string

const (
	NcChallengesAdminWidgetListClearanceLevelNoClearance NcChallengesAdminWidgetListClearanceLevel = "no_clearance"
	NcChallengesAdminWidgetListClearanceLevelJschallenge NcChallengesAdminWidgetListClearanceLevel = "jschallenge"
	NcChallengesAdminWidgetListClearanceLevelManaged     NcChallengesAdminWidgetListClearanceLevel = "managed"
	NcChallengesAdminWidgetListClearanceLevelInteractive NcChallengesAdminWidgetListClearanceLevel = "interactive"
)

// Widget Mode
type NcChallengesAdminWidgetListMode string

const (
	NcChallengesAdminWidgetListModeNonInteractive NcChallengesAdminWidgetListMode = "non-interactive"
	NcChallengesAdminWidgetListModeInvisible      NcChallengesAdminWidgetListMode = "invisible"
	NcChallengesAdminWidgetListModeManaged        NcChallengesAdminWidgetListMode = "managed"
)

// Region where this widget can be used.
type NcChallengesAdminWidgetListRegion string

const (
	NcChallengesAdminWidgetListRegionWorld NcChallengesAdminWidgetListRegion = "world"
)

type WidgetNewParams struct {
	Domains param.Field[[]string] `json:"domains,required"`
	// Widget Mode
	Mode param.Field[WidgetNewParamsMode] `json:"mode,required"`
	// Human readable widget name. Not unique. Cloudflare suggests that you set this to
	// a meaningful string to make it easier to identify your widget, and where it is
	// used.
	Name param.Field[string] `json:"name,required"`
	// Direction to order widgets.
	Direction param.Field[WidgetNewParamsDirection] `query:"direction"`
	// Field to order widgets by.
	Order param.Field[WidgetNewParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of items per page.
	PerPage param.Field[float64] `query:"per_page"`
	// If bot_fight_mode is set to `true`, Cloudflare issues computationally expensive
	// challenges in response to malicious bots (ENT only).
	BotFightMode param.Field[bool] `json:"bot_fight_mode"`
	// If Turnstile is embedded on a Cloudflare site and the widget should grant
	// challenge clearance, this setting can determine the clearance level to be set
	ClearanceLevel param.Field[WidgetNewParamsClearanceLevel] `json:"clearance_level"`
	// Do not show any Cloudflare branding on the widget (ENT only).
	Offlabel param.Field[bool] `json:"offlabel"`
	// Region where this widget can be used.
	Region param.Field[WidgetNewParamsRegion] `json:"region"`
}

func (r WidgetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [WidgetNewParams]'s query parameters as `url.Values`.
func (r WidgetNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Widget Mode
type WidgetNewParamsMode string

const (
	WidgetNewParamsModeNonInteractive WidgetNewParamsMode = "non-interactive"
	WidgetNewParamsModeInvisible      WidgetNewParamsMode = "invisible"
	WidgetNewParamsModeManaged        WidgetNewParamsMode = "managed"
)

// Direction to order widgets.
type WidgetNewParamsDirection string

const (
	WidgetNewParamsDirectionAsc  WidgetNewParamsDirection = "asc"
	WidgetNewParamsDirectionDesc WidgetNewParamsDirection = "desc"
)

// Field to order widgets by.
type WidgetNewParamsOrder string

const (
	WidgetNewParamsOrderID         WidgetNewParamsOrder = "id"
	WidgetNewParamsOrderSitekey    WidgetNewParamsOrder = "sitekey"
	WidgetNewParamsOrderName       WidgetNewParamsOrder = "name"
	WidgetNewParamsOrderCreatedOn  WidgetNewParamsOrder = "created_on"
	WidgetNewParamsOrderModifiedOn WidgetNewParamsOrder = "modified_on"
)

// If Turnstile is embedded on a Cloudflare site and the widget should grant
// challenge clearance, this setting can determine the clearance level to be set
type WidgetNewParamsClearanceLevel string

const (
	WidgetNewParamsClearanceLevelNoClearance WidgetNewParamsClearanceLevel = "no_clearance"
	WidgetNewParamsClearanceLevelJschallenge WidgetNewParamsClearanceLevel = "jschallenge"
	WidgetNewParamsClearanceLevelManaged     WidgetNewParamsClearanceLevel = "managed"
	WidgetNewParamsClearanceLevelInteractive WidgetNewParamsClearanceLevel = "interactive"
)

// Region where this widget can be used.
type WidgetNewParamsRegion string

const (
	WidgetNewParamsRegionWorld WidgetNewParamsRegion = "world"
)

type WidgetNewResponseEnvelope struct {
	Errors   []WidgetNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WidgetNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// A Turnstile widget's detailed configuration
	Result     NcChallengesAdminWidgetDetail       `json:"result"`
	ResultInfo WidgetNewResponseEnvelopeResultInfo `json:"result_info"`
	JSON       widgetNewResponseEnvelopeJSON       `json:"-"`
}

// widgetNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [WidgetNewResponseEnvelope]
type widgetNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WidgetNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r widgetNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WidgetNewResponseEnvelopeErrors struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    widgetNewResponseEnvelopeErrorsJSON `json:"-"`
}

// widgetNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [WidgetNewResponseEnvelopeErrors]
type widgetNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WidgetNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r widgetNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type WidgetNewResponseEnvelopeMessages struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    widgetNewResponseEnvelopeMessagesJSON `json:"-"`
}

// widgetNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [WidgetNewResponseEnvelopeMessages]
type widgetNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WidgetNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r widgetNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type WidgetNewResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count,required"`
	// Current page within paginated list of results
	Page float64 `json:"page,required"`
	// Number of results per page of results
	PerPage float64 `json:"per_page,required"`
	// Total results available without any search parameters
	TotalCount float64                                 `json:"total_count,required"`
	JSON       widgetNewResponseEnvelopeResultInfoJSON `json:"-"`
}

// widgetNewResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [WidgetNewResponseEnvelopeResultInfo]
type widgetNewResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WidgetNewResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r widgetNewResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type WidgetUpdateParams struct {
	Domains param.Field[[]string] `json:"domains,required"`
	// Widget Mode
	Mode param.Field[WidgetUpdateParamsMode] `json:"mode,required"`
	// Human readable widget name. Not unique. Cloudflare suggests that you set this to
	// a meaningful string to make it easier to identify your widget, and where it is
	// used.
	Name param.Field[string] `json:"name,required"`
	// If bot_fight_mode is set to `true`, Cloudflare issues computationally expensive
	// challenges in response to malicious bots (ENT only).
	BotFightMode param.Field[bool] `json:"bot_fight_mode"`
	// If Turnstile is embedded on a Cloudflare site and the widget should grant
	// challenge clearance, this setting can determine the clearance level to be set
	ClearanceLevel param.Field[WidgetUpdateParamsClearanceLevel] `json:"clearance_level"`
	// Do not show any Cloudflare branding on the widget (ENT only).
	Offlabel param.Field[bool] `json:"offlabel"`
}

func (r WidgetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Widget Mode
type WidgetUpdateParamsMode string

const (
	WidgetUpdateParamsModeNonInteractive WidgetUpdateParamsMode = "non-interactive"
	WidgetUpdateParamsModeInvisible      WidgetUpdateParamsMode = "invisible"
	WidgetUpdateParamsModeManaged        WidgetUpdateParamsMode = "managed"
)

// If Turnstile is embedded on a Cloudflare site and the widget should grant
// challenge clearance, this setting can determine the clearance level to be set
type WidgetUpdateParamsClearanceLevel string

const (
	WidgetUpdateParamsClearanceLevelNoClearance WidgetUpdateParamsClearanceLevel = "no_clearance"
	WidgetUpdateParamsClearanceLevelJschallenge WidgetUpdateParamsClearanceLevel = "jschallenge"
	WidgetUpdateParamsClearanceLevelManaged     WidgetUpdateParamsClearanceLevel = "managed"
	WidgetUpdateParamsClearanceLevelInteractive WidgetUpdateParamsClearanceLevel = "interactive"
)

type WidgetUpdateResponseEnvelope struct {
	Errors   []WidgetUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WidgetUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// A Turnstile widget's detailed configuration
	Result NcChallengesAdminWidgetDetail    `json:"result"`
	JSON   widgetUpdateResponseEnvelopeJSON `json:"-"`
}

// widgetUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [WidgetUpdateResponseEnvelope]
type widgetUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WidgetUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r widgetUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WidgetUpdateResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    widgetUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// widgetUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [WidgetUpdateResponseEnvelopeErrors]
type widgetUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WidgetUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r widgetUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type WidgetUpdateResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    widgetUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// widgetUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [WidgetUpdateResponseEnvelopeMessages]
type widgetUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WidgetUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r widgetUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type WidgetListParams struct {
	// Direction to order widgets.
	Direction param.Field[WidgetListParamsDirection] `query:"direction"`
	// Field to order widgets by.
	Order param.Field[WidgetListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of items per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [WidgetListParams]'s query parameters as `url.Values`.
func (r WidgetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order widgets.
type WidgetListParamsDirection string

const (
	WidgetListParamsDirectionAsc  WidgetListParamsDirection = "asc"
	WidgetListParamsDirectionDesc WidgetListParamsDirection = "desc"
)

// Field to order widgets by.
type WidgetListParamsOrder string

const (
	WidgetListParamsOrderID         WidgetListParamsOrder = "id"
	WidgetListParamsOrderSitekey    WidgetListParamsOrder = "sitekey"
	WidgetListParamsOrderName       WidgetListParamsOrder = "name"
	WidgetListParamsOrderCreatedOn  WidgetListParamsOrder = "created_on"
	WidgetListParamsOrderModifiedOn WidgetListParamsOrder = "modified_on"
)

type WidgetDeleteResponseEnvelope struct {
	Errors   []WidgetDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WidgetDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// A Turnstile widget's detailed configuration
	Result NcChallengesAdminWidgetDetail    `json:"result"`
	JSON   widgetDeleteResponseEnvelopeJSON `json:"-"`
}

// widgetDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [WidgetDeleteResponseEnvelope]
type widgetDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WidgetDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r widgetDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WidgetDeleteResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    widgetDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// widgetDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [WidgetDeleteResponseEnvelopeErrors]
type widgetDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WidgetDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r widgetDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type WidgetDeleteResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    widgetDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// widgetDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [WidgetDeleteResponseEnvelopeMessages]
type widgetDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WidgetDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r widgetDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type WidgetGetResponseEnvelope struct {
	Errors   []WidgetGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WidgetGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// A Turnstile widget's detailed configuration
	Result NcChallengesAdminWidgetDetail `json:"result"`
	JSON   widgetGetResponseEnvelopeJSON `json:"-"`
}

// widgetGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [WidgetGetResponseEnvelope]
type widgetGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WidgetGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r widgetGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WidgetGetResponseEnvelopeErrors struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    widgetGetResponseEnvelopeErrorsJSON `json:"-"`
}

// widgetGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [WidgetGetResponseEnvelopeErrors]
type widgetGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WidgetGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r widgetGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type WidgetGetResponseEnvelopeMessages struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    widgetGetResponseEnvelopeMessagesJSON `json:"-"`
}

// widgetGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [WidgetGetResponseEnvelopeMessages]
type widgetGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WidgetGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r widgetGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type WidgetRotateSecretParams struct {
	// If `invalidate_immediately` is set to `false`, the previous secret will remain
	// valid for two hours. Otherwise, the secret is immediately invalidated, and
	// requests using it will be rejected.
	InvalidateImmediately param.Field[bool] `json:"invalidate_immediately"`
}

func (r WidgetRotateSecretParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WidgetRotateSecretResponseEnvelope struct {
	Errors   []WidgetRotateSecretResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WidgetRotateSecretResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// A Turnstile widget's detailed configuration
	Result NcChallengesAdminWidgetDetail          `json:"result"`
	JSON   widgetRotateSecretResponseEnvelopeJSON `json:"-"`
}

// widgetRotateSecretResponseEnvelopeJSON contains the JSON metadata for the struct
// [WidgetRotateSecretResponseEnvelope]
type widgetRotateSecretResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WidgetRotateSecretResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r widgetRotateSecretResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WidgetRotateSecretResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    widgetRotateSecretResponseEnvelopeErrorsJSON `json:"-"`
}

// widgetRotateSecretResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WidgetRotateSecretResponseEnvelopeErrors]
type widgetRotateSecretResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WidgetRotateSecretResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r widgetRotateSecretResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type WidgetRotateSecretResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    widgetRotateSecretResponseEnvelopeMessagesJSON `json:"-"`
}

// widgetRotateSecretResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [WidgetRotateSecretResponseEnvelopeMessages]
type widgetRotateSecretResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WidgetRotateSecretResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r widgetRotateSecretResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
