// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package challenges

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
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
func (r *WidgetService) New(ctx context.Context, params WidgetNewParams, opts ...option.RequestOption) (res *ChallengesWidget, err error) {
	opts = append(r.Options[:], opts...)
	var env WidgetNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/challenges/widgets", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update the configuration of a widget.
func (r *WidgetService) Update(ctx context.Context, sitekey string, params WidgetUpdateParams, opts ...option.RequestOption) (res *ChallengesWidget, err error) {
	opts = append(r.Options[:], opts...)
	var env WidgetUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/challenges/widgets/%s", params.AccountID, sitekey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all turnstile widgets of an account.
func (r *WidgetService) List(ctx context.Context, params WidgetListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[ChallengesWidgetList], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/challenges/widgets", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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
func (r *WidgetService) ListAutoPaging(ctx context.Context, params WidgetListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[ChallengesWidgetList] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Destroy a Turnstile Widget.
func (r *WidgetService) Delete(ctx context.Context, sitekey string, body WidgetDeleteParams, opts ...option.RequestOption) (res *ChallengesWidget, err error) {
	opts = append(r.Options[:], opts...)
	var env WidgetDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/challenges/widgets/%s", body.AccountID, sitekey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Show a single challenge widget configuration.
func (r *WidgetService) Get(ctx context.Context, sitekey string, query WidgetGetParams, opts ...option.RequestOption) (res *ChallengesWidget, err error) {
	opts = append(r.Options[:], opts...)
	var env WidgetGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/challenges/widgets/%s", query.AccountID, sitekey)
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
func (r *WidgetService) RotateSecret(ctx context.Context, sitekey string, params WidgetRotateSecretParams, opts ...option.RequestOption) (res *ChallengesWidget, err error) {
	opts = append(r.Options[:], opts...)
	var env WidgetRotateSecretResponseEnvelope
	path := fmt.Sprintf("accounts/%s/challenges/widgets/%s/rotate_secret", params.AccountID, sitekey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A Turnstile widget's detailed configuration
type ChallengesWidget struct {
	// If bot_fight_mode is set to `true`, Cloudflare issues computationally expensive
	// challenges in response to malicious bots (ENT only).
	BotFightMode bool `json:"bot_fight_mode,required"`
	// If Turnstile is embedded on a Cloudflare site and the widget should grant
	// challenge clearance, this setting can determine the clearance level to be set
	ClearanceLevel ChallengesWidgetClearanceLevel `json:"clearance_level,required"`
	// When the widget was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	Domains   []string  `json:"domains,required"`
	// Widget Mode
	Mode ChallengesWidgetMode `json:"mode,required"`
	// When the widget was modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Human readable widget name. Not unique. Cloudflare suggests that you set this to
	// a meaningful string to make it easier to identify your widget, and where it is
	// used.
	Name string `json:"name,required"`
	// Do not show any Cloudflare branding on the widget (ENT only).
	Offlabel bool `json:"offlabel,required"`
	// Region where this widget can be used.
	Region ChallengesWidgetRegion `json:"region,required"`
	// Secret key for this widget.
	Secret string `json:"secret,required"`
	// Widget item identifier tag.
	Sitekey string               `json:"sitekey,required"`
	JSON    challengesWidgetJSON `json:"-"`
}

// challengesWidgetJSON contains the JSON metadata for the struct
// [ChallengesWidget]
type challengesWidgetJSON struct {
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

func (r *ChallengesWidget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r challengesWidgetJSON) RawJSON() string {
	return r.raw
}

// If Turnstile is embedded on a Cloudflare site and the widget should grant
// challenge clearance, this setting can determine the clearance level to be set
type ChallengesWidgetClearanceLevel string

const (
	ChallengesWidgetClearanceLevelNoClearance ChallengesWidgetClearanceLevel = "no_clearance"
	ChallengesWidgetClearanceLevelJschallenge ChallengesWidgetClearanceLevel = "jschallenge"
	ChallengesWidgetClearanceLevelManaged     ChallengesWidgetClearanceLevel = "managed"
	ChallengesWidgetClearanceLevelInteractive ChallengesWidgetClearanceLevel = "interactive"
)

func (r ChallengesWidgetClearanceLevel) IsKnown() bool {
	switch r {
	case ChallengesWidgetClearanceLevelNoClearance, ChallengesWidgetClearanceLevelJschallenge, ChallengesWidgetClearanceLevelManaged, ChallengesWidgetClearanceLevelInteractive:
		return true
	}
	return false
}

// Widget Mode
type ChallengesWidgetMode string

const (
	ChallengesWidgetModeNonInteractive ChallengesWidgetMode = "non-interactive"
	ChallengesWidgetModeInvisible      ChallengesWidgetMode = "invisible"
	ChallengesWidgetModeManaged        ChallengesWidgetMode = "managed"
)

func (r ChallengesWidgetMode) IsKnown() bool {
	switch r {
	case ChallengesWidgetModeNonInteractive, ChallengesWidgetModeInvisible, ChallengesWidgetModeManaged:
		return true
	}
	return false
}

// Region where this widget can be used.
type ChallengesWidgetRegion string

const (
	ChallengesWidgetRegionWorld ChallengesWidgetRegion = "world"
)

func (r ChallengesWidgetRegion) IsKnown() bool {
	switch r {
	case ChallengesWidgetRegionWorld:
		return true
	}
	return false
}

// A Turnstile Widgets configuration as it appears in listings
type ChallengesWidgetList struct {
	// If bot_fight_mode is set to `true`, Cloudflare issues computationally expensive
	// challenges in response to malicious bots (ENT only).
	BotFightMode bool `json:"bot_fight_mode,required"`
	// If Turnstile is embedded on a Cloudflare site and the widget should grant
	// challenge clearance, this setting can determine the clearance level to be set
	ClearanceLevel ChallengesWidgetListClearanceLevel `json:"clearance_level,required"`
	// When the widget was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	Domains   []string  `json:"domains,required"`
	// Widget Mode
	Mode ChallengesWidgetListMode `json:"mode,required"`
	// When the widget was modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Human readable widget name. Not unique. Cloudflare suggests that you set this to
	// a meaningful string to make it easier to identify your widget, and where it is
	// used.
	Name string `json:"name,required"`
	// Do not show any Cloudflare branding on the widget (ENT only).
	Offlabel bool `json:"offlabel,required"`
	// Region where this widget can be used.
	Region ChallengesWidgetListRegion `json:"region,required"`
	// Widget item identifier tag.
	Sitekey string                   `json:"sitekey,required"`
	JSON    challengesWidgetListJSON `json:"-"`
}

// challengesWidgetListJSON contains the JSON metadata for the struct
// [ChallengesWidgetList]
type challengesWidgetListJSON struct {
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

func (r *ChallengesWidgetList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r challengesWidgetListJSON) RawJSON() string {
	return r.raw
}

// If Turnstile is embedded on a Cloudflare site and the widget should grant
// challenge clearance, this setting can determine the clearance level to be set
type ChallengesWidgetListClearanceLevel string

const (
	ChallengesWidgetListClearanceLevelNoClearance ChallengesWidgetListClearanceLevel = "no_clearance"
	ChallengesWidgetListClearanceLevelJschallenge ChallengesWidgetListClearanceLevel = "jschallenge"
	ChallengesWidgetListClearanceLevelManaged     ChallengesWidgetListClearanceLevel = "managed"
	ChallengesWidgetListClearanceLevelInteractive ChallengesWidgetListClearanceLevel = "interactive"
)

func (r ChallengesWidgetListClearanceLevel) IsKnown() bool {
	switch r {
	case ChallengesWidgetListClearanceLevelNoClearance, ChallengesWidgetListClearanceLevelJschallenge, ChallengesWidgetListClearanceLevelManaged, ChallengesWidgetListClearanceLevelInteractive:
		return true
	}
	return false
}

// Widget Mode
type ChallengesWidgetListMode string

const (
	ChallengesWidgetListModeNonInteractive ChallengesWidgetListMode = "non-interactive"
	ChallengesWidgetListModeInvisible      ChallengesWidgetListMode = "invisible"
	ChallengesWidgetListModeManaged        ChallengesWidgetListMode = "managed"
)

func (r ChallengesWidgetListMode) IsKnown() bool {
	switch r {
	case ChallengesWidgetListModeNonInteractive, ChallengesWidgetListModeInvisible, ChallengesWidgetListModeManaged:
		return true
	}
	return false
}

// Region where this widget can be used.
type ChallengesWidgetListRegion string

const (
	ChallengesWidgetListRegionWorld ChallengesWidgetListRegion = "world"
)

func (r ChallengesWidgetListRegion) IsKnown() bool {
	switch r {
	case ChallengesWidgetListRegionWorld:
		return true
	}
	return false
}

type ChallengesWidgetListItem = string

type ChallengesWidgetListItemParam = string

type WidgetNewParams struct {
	// Identifier
	AccountID param.Field[string]                          `path:"account_id,required"`
	Domains   param.Field[[]ChallengesWidgetListItemParam] `json:"domains,required"`
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
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
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

func (r WidgetNewParamsMode) IsKnown() bool {
	switch r {
	case WidgetNewParamsModeNonInteractive, WidgetNewParamsModeInvisible, WidgetNewParamsModeManaged:
		return true
	}
	return false
}

// Direction to order widgets.
type WidgetNewParamsDirection string

const (
	WidgetNewParamsDirectionAsc  WidgetNewParamsDirection = "asc"
	WidgetNewParamsDirectionDesc WidgetNewParamsDirection = "desc"
)

func (r WidgetNewParamsDirection) IsKnown() bool {
	switch r {
	case WidgetNewParamsDirectionAsc, WidgetNewParamsDirectionDesc:
		return true
	}
	return false
}

// Field to order widgets by.
type WidgetNewParamsOrder string

const (
	WidgetNewParamsOrderID         WidgetNewParamsOrder = "id"
	WidgetNewParamsOrderSitekey    WidgetNewParamsOrder = "sitekey"
	WidgetNewParamsOrderName       WidgetNewParamsOrder = "name"
	WidgetNewParamsOrderCreatedOn  WidgetNewParamsOrder = "created_on"
	WidgetNewParamsOrderModifiedOn WidgetNewParamsOrder = "modified_on"
)

func (r WidgetNewParamsOrder) IsKnown() bool {
	switch r {
	case WidgetNewParamsOrderID, WidgetNewParamsOrderSitekey, WidgetNewParamsOrderName, WidgetNewParamsOrderCreatedOn, WidgetNewParamsOrderModifiedOn:
		return true
	}
	return false
}

// If Turnstile is embedded on a Cloudflare site and the widget should grant
// challenge clearance, this setting can determine the clearance level to be set
type WidgetNewParamsClearanceLevel string

const (
	WidgetNewParamsClearanceLevelNoClearance WidgetNewParamsClearanceLevel = "no_clearance"
	WidgetNewParamsClearanceLevelJschallenge WidgetNewParamsClearanceLevel = "jschallenge"
	WidgetNewParamsClearanceLevelManaged     WidgetNewParamsClearanceLevel = "managed"
	WidgetNewParamsClearanceLevelInteractive WidgetNewParamsClearanceLevel = "interactive"
)

func (r WidgetNewParamsClearanceLevel) IsKnown() bool {
	switch r {
	case WidgetNewParamsClearanceLevelNoClearance, WidgetNewParamsClearanceLevelJschallenge, WidgetNewParamsClearanceLevelManaged, WidgetNewParamsClearanceLevelInteractive:
		return true
	}
	return false
}

// Region where this widget can be used.
type WidgetNewParamsRegion string

const (
	WidgetNewParamsRegionWorld WidgetNewParamsRegion = "world"
)

func (r WidgetNewParamsRegion) IsKnown() bool {
	switch r {
	case WidgetNewParamsRegionWorld:
		return true
	}
	return false
}

type WidgetNewResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// A Turnstile widget's detailed configuration
	Result     ChallengesWidget                    `json:"result"`
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
	// Identifier
	AccountID param.Field[string]                          `path:"account_id,required"`
	Domains   param.Field[[]ChallengesWidgetListItemParam] `json:"domains,required"`
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

func (r WidgetUpdateParamsMode) IsKnown() bool {
	switch r {
	case WidgetUpdateParamsModeNonInteractive, WidgetUpdateParamsModeInvisible, WidgetUpdateParamsModeManaged:
		return true
	}
	return false
}

// If Turnstile is embedded on a Cloudflare site and the widget should grant
// challenge clearance, this setting can determine the clearance level to be set
type WidgetUpdateParamsClearanceLevel string

const (
	WidgetUpdateParamsClearanceLevelNoClearance WidgetUpdateParamsClearanceLevel = "no_clearance"
	WidgetUpdateParamsClearanceLevelJschallenge WidgetUpdateParamsClearanceLevel = "jschallenge"
	WidgetUpdateParamsClearanceLevelManaged     WidgetUpdateParamsClearanceLevel = "managed"
	WidgetUpdateParamsClearanceLevelInteractive WidgetUpdateParamsClearanceLevel = "interactive"
)

func (r WidgetUpdateParamsClearanceLevel) IsKnown() bool {
	switch r {
	case WidgetUpdateParamsClearanceLevelNoClearance, WidgetUpdateParamsClearanceLevelJschallenge, WidgetUpdateParamsClearanceLevelManaged, WidgetUpdateParamsClearanceLevelInteractive:
		return true
	}
	return false
}

type WidgetUpdateResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// A Turnstile widget's detailed configuration
	Result ChallengesWidget                 `json:"result"`
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

type WidgetListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
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
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order widgets.
type WidgetListParamsDirection string

const (
	WidgetListParamsDirectionAsc  WidgetListParamsDirection = "asc"
	WidgetListParamsDirectionDesc WidgetListParamsDirection = "desc"
)

func (r WidgetListParamsDirection) IsKnown() bool {
	switch r {
	case WidgetListParamsDirectionAsc, WidgetListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to order widgets by.
type WidgetListParamsOrder string

const (
	WidgetListParamsOrderID         WidgetListParamsOrder = "id"
	WidgetListParamsOrderSitekey    WidgetListParamsOrder = "sitekey"
	WidgetListParamsOrderName       WidgetListParamsOrder = "name"
	WidgetListParamsOrderCreatedOn  WidgetListParamsOrder = "created_on"
	WidgetListParamsOrderModifiedOn WidgetListParamsOrder = "modified_on"
)

func (r WidgetListParamsOrder) IsKnown() bool {
	switch r {
	case WidgetListParamsOrderID, WidgetListParamsOrderSitekey, WidgetListParamsOrderName, WidgetListParamsOrderCreatedOn, WidgetListParamsOrderModifiedOn:
		return true
	}
	return false
}

type WidgetDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type WidgetDeleteResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// A Turnstile widget's detailed configuration
	Result ChallengesWidget                 `json:"result"`
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

type WidgetGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type WidgetGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// A Turnstile widget's detailed configuration
	Result ChallengesWidget              `json:"result"`
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

type WidgetRotateSecretParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// If `invalidate_immediately` is set to `false`, the previous secret will remain
	// valid for two hours. Otherwise, the secret is immediately invalidated, and
	// requests using it will be rejected.
	InvalidateImmediately param.Field[bool] `json:"invalidate_immediately"`
}

func (r WidgetRotateSecretParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WidgetRotateSecretResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// A Turnstile widget's detailed configuration
	Result ChallengesWidget                       `json:"result"`
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
