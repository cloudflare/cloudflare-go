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
func (r *ChallengeWidgetService) New(ctx context.Context, accountIdentifier string, params ChallengeWidgetNewParams, opts ...option.RequestOption) (res *NcChallengesAdminWidgetDetail, err error) {
	opts = append(r.Options[:], opts...)
	var env ChallengeWidgetNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/challenges/widgets", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update the configuration of a widget.
func (r *ChallengeWidgetService) Update(ctx context.Context, accountIdentifier string, sitekey string, body ChallengeWidgetUpdateParams, opts ...option.RequestOption) (res *NcChallengesAdminWidgetDetail, err error) {
	opts = append(r.Options[:], opts...)
	var env ChallengeWidgetUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/challenges/widgets/%s", accountIdentifier, sitekey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all turnstile widgets of an account.
func (r *ChallengeWidgetService) List(ctx context.Context, accountIdentifier string, query ChallengeWidgetListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[NcChallengesAdminWidgetList], err error) {
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
func (r *ChallengeWidgetService) ListAutoPaging(ctx context.Context, accountIdentifier string, query ChallengeWidgetListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[NcChallengesAdminWidgetList] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, accountIdentifier, query, opts...))
}

// Destroy a Turnstile Widget.
func (r *ChallengeWidgetService) Delete(ctx context.Context, accountIdentifier string, sitekey string, opts ...option.RequestOption) (res *NcChallengesAdminWidgetDetail, err error) {
	opts = append(r.Options[:], opts...)
	var env ChallengeWidgetDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/challenges/widgets/%s", accountIdentifier, sitekey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Show a single challenge widget configuration.
func (r *ChallengeWidgetService) Get(ctx context.Context, accountIdentifier string, sitekey string, opts ...option.RequestOption) (res *NcChallengesAdminWidgetDetail, err error) {
	opts = append(r.Options[:], opts...)
	var env ChallengeWidgetGetResponseEnvelope
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
func (r *ChallengeWidgetService) RotateSecret(ctx context.Context, accountIdentifier string, sitekey string, body ChallengeWidgetRotateSecretParams, opts ...option.RequestOption) (res *NcChallengesAdminWidgetDetail, err error) {
	opts = append(r.Options[:], opts...)
	var env ChallengeWidgetRotateSecretResponseEnvelope
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

type ChallengeWidgetNewResponseEnvelope struct {
	Errors   []ChallengeWidgetNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ChallengeWidgetNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// A Turnstile widget's detailed configuration
	Result     NcChallengesAdminWidgetDetail                `json:"result"`
	ResultInfo ChallengeWidgetNewResponseEnvelopeResultInfo `json:"result_info"`
	JSON       challengeWidgetNewResponseEnvelopeJSON       `json:"-"`
}

// challengeWidgetNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ChallengeWidgetNewResponseEnvelope]
type challengeWidgetNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChallengeWidgetNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ChallengeWidgetNewResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    challengeWidgetNewResponseEnvelopeErrorsJSON `json:"-"`
}

// challengeWidgetNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ChallengeWidgetNewResponseEnvelopeErrors]
type challengeWidgetNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChallengeWidgetNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ChallengeWidgetNewResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    challengeWidgetNewResponseEnvelopeMessagesJSON `json:"-"`
}

// challengeWidgetNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ChallengeWidgetNewResponseEnvelopeMessages]
type challengeWidgetNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChallengeWidgetNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ChallengeWidgetNewResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count,required"`
	// Current page within paginated list of results
	Page float64 `json:"page,required"`
	// Number of results per page of results
	PerPage float64 `json:"per_page,required"`
	// Total results available without any search parameters
	TotalCount float64                                          `json:"total_count,required"`
	JSON       challengeWidgetNewResponseEnvelopeResultInfoJSON `json:"-"`
}

// challengeWidgetNewResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [ChallengeWidgetNewResponseEnvelopeResultInfo]
type challengeWidgetNewResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChallengeWidgetNewResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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

type ChallengeWidgetUpdateResponseEnvelope struct {
	Errors   []ChallengeWidgetUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ChallengeWidgetUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// A Turnstile widget's detailed configuration
	Result NcChallengesAdminWidgetDetail             `json:"result"`
	JSON   challengeWidgetUpdateResponseEnvelopeJSON `json:"-"`
}

// challengeWidgetUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [ChallengeWidgetUpdateResponseEnvelope]
type challengeWidgetUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChallengeWidgetUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ChallengeWidgetUpdateResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    challengeWidgetUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// challengeWidgetUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ChallengeWidgetUpdateResponseEnvelopeErrors]
type challengeWidgetUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChallengeWidgetUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ChallengeWidgetUpdateResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    challengeWidgetUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// challengeWidgetUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ChallengeWidgetUpdateResponseEnvelopeMessages]
type challengeWidgetUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChallengeWidgetUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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

type ChallengeWidgetDeleteResponseEnvelope struct {
	Errors   []ChallengeWidgetDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ChallengeWidgetDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// A Turnstile widget's detailed configuration
	Result NcChallengesAdminWidgetDetail             `json:"result"`
	JSON   challengeWidgetDeleteResponseEnvelopeJSON `json:"-"`
}

// challengeWidgetDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [ChallengeWidgetDeleteResponseEnvelope]
type challengeWidgetDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChallengeWidgetDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ChallengeWidgetDeleteResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    challengeWidgetDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// challengeWidgetDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ChallengeWidgetDeleteResponseEnvelopeErrors]
type challengeWidgetDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChallengeWidgetDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ChallengeWidgetDeleteResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    challengeWidgetDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// challengeWidgetDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ChallengeWidgetDeleteResponseEnvelopeMessages]
type challengeWidgetDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChallengeWidgetDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ChallengeWidgetGetResponseEnvelope struct {
	Errors   []ChallengeWidgetGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ChallengeWidgetGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// A Turnstile widget's detailed configuration
	Result NcChallengesAdminWidgetDetail          `json:"result"`
	JSON   challengeWidgetGetResponseEnvelopeJSON `json:"-"`
}

// challengeWidgetGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ChallengeWidgetGetResponseEnvelope]
type challengeWidgetGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChallengeWidgetGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ChallengeWidgetGetResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    challengeWidgetGetResponseEnvelopeErrorsJSON `json:"-"`
}

// challengeWidgetGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ChallengeWidgetGetResponseEnvelopeErrors]
type challengeWidgetGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChallengeWidgetGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ChallengeWidgetGetResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    challengeWidgetGetResponseEnvelopeMessagesJSON `json:"-"`
}

// challengeWidgetGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ChallengeWidgetGetResponseEnvelopeMessages]
type challengeWidgetGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChallengeWidgetGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ChallengeWidgetRotateSecretParams struct {
	// If `invalidate_immediately` is set to `false`, the previous secret will remain
	// valid for two hours. Otherwise, the secret is immediately invalidated, and
	// requests using it will be rejected.
	InvalidateImmediately param.Field[bool] `json:"invalidate_immediately"`
}

func (r ChallengeWidgetRotateSecretParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ChallengeWidgetRotateSecretResponseEnvelope struct {
	Errors   []ChallengeWidgetRotateSecretResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ChallengeWidgetRotateSecretResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// A Turnstile widget's detailed configuration
	Result NcChallengesAdminWidgetDetail                   `json:"result"`
	JSON   challengeWidgetRotateSecretResponseEnvelopeJSON `json:"-"`
}

// challengeWidgetRotateSecretResponseEnvelopeJSON contains the JSON metadata for
// the struct [ChallengeWidgetRotateSecretResponseEnvelope]
type challengeWidgetRotateSecretResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChallengeWidgetRotateSecretResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ChallengeWidgetRotateSecretResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    challengeWidgetRotateSecretResponseEnvelopeErrorsJSON `json:"-"`
}

// challengeWidgetRotateSecretResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ChallengeWidgetRotateSecretResponseEnvelopeErrors]
type challengeWidgetRotateSecretResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChallengeWidgetRotateSecretResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ChallengeWidgetRotateSecretResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    challengeWidgetRotateSecretResponseEnvelopeMessagesJSON `json:"-"`
}

// challengeWidgetRotateSecretResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ChallengeWidgetRotateSecretResponseEnvelopeMessages]
type challengeWidgetRotateSecretResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChallengeWidgetRotateSecretResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
