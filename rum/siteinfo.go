// File generated from our OpenAPI spec by Stainless.

package rum

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
)

// SiteInfoService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewSiteInfoService] method instead.
type SiteInfoService struct {
	Options []option.RequestOption
}

// NewSiteInfoService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSiteInfoService(opts ...option.RequestOption) (r *SiteInfoService) {
	r = &SiteInfoService{}
	r.Options = opts
	return
}

// Creates a new Web Analytics site.
func (r *SiteInfoService) New(ctx context.Context, params SiteInfoNewParams, opts ...option.RequestOption) (res *SiteInfoNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteInfoNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rum/site_info", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing Web Analytics site.
func (r *SiteInfoService) Update(ctx context.Context, siteID string, params SiteInfoUpdateParams, opts ...option.RequestOption) (res *SiteInfoUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteInfoUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rum/site_info/%s", params.AccountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all Web Analytics sites of an account.
func (r *SiteInfoService) List(ctx context.Context, params SiteInfoListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[SiteInfoListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/rum/site_info/list", params.AccountID)
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

// Lists all Web Analytics sites of an account.
func (r *SiteInfoService) ListAutoPaging(ctx context.Context, params SiteInfoListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[SiteInfoListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes an existing Web Analytics site.
func (r *SiteInfoService) Delete(ctx context.Context, siteID string, body SiteInfoDeleteParams, opts ...option.RequestOption) (res *SiteInfoDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteInfoDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rum/site_info/%s", body.AccountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves a Web Analytics site.
func (r *SiteInfoService) Get(ctx context.Context, siteID string, query SiteInfoGetParams, opts ...option.RequestOption) (res *SiteInfoGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteInfoGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rum/site_info/%s", query.AccountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SiteInfoNewResponse struct {
	// If enabled, the JavaScript snippet is automatically injected for orange-clouded
	// sites.
	AutoInstall bool      `json:"auto_install"`
	Created     time.Time `json:"created" format:"date-time"`
	// A list of rules.
	Rules   []SiteInfoNewResponseRule  `json:"rules"`
	Ruleset SiteInfoNewResponseRuleset `json:"ruleset"`
	// The Web Analytics site identifier.
	SiteTag string `json:"site_tag"`
	// The Web Analytics site token.
	SiteToken string `json:"site_token"`
	// Encoded JavaScript snippet.
	Snippet string                  `json:"snippet"`
	JSON    siteInfoNewResponseJSON `json:"-"`
}

// siteInfoNewResponseJSON contains the JSON metadata for the struct
// [SiteInfoNewResponse]
type siteInfoNewResponseJSON struct {
	AutoInstall apijson.Field
	Created     apijson.Field
	Rules       apijson.Field
	Ruleset     apijson.Field
	SiteTag     apijson.Field
	SiteToken   apijson.Field
	Snippet     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteInfoNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteInfoNewResponseJSON) RawJSON() string {
	return r.raw
}

type SiteInfoNewResponseRule struct {
	// The Web Analytics rule identifier.
	ID      string    `json:"id"`
	Created time.Time `json:"created" format:"date-time"`
	// The hostname the rule will be applied to.
	Host string `json:"host"`
	// Whether the rule includes or excludes traffic from being measured.
	Inclusive bool `json:"inclusive"`
	// Whether the rule is paused or not.
	IsPaused bool `json:"is_paused"`
	// The paths the rule will be applied to.
	Paths    []string                    `json:"paths"`
	Priority float64                     `json:"priority"`
	JSON     siteInfoNewResponseRuleJSON `json:"-"`
}

// siteInfoNewResponseRuleJSON contains the JSON metadata for the struct
// [SiteInfoNewResponseRule]
type siteInfoNewResponseRuleJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Host        apijson.Field
	Inclusive   apijson.Field
	IsPaused    apijson.Field
	Paths       apijson.Field
	Priority    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteInfoNewResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteInfoNewResponseRuleJSON) RawJSON() string {
	return r.raw
}

type SiteInfoNewResponseRuleset struct {
	// The Web Analytics ruleset identifier.
	ID string `json:"id"`
	// Whether the ruleset is enabled.
	Enabled  bool   `json:"enabled"`
	ZoneName string `json:"zone_name"`
	// The zone identifier.
	ZoneTag string                         `json:"zone_tag"`
	JSON    siteInfoNewResponseRulesetJSON `json:"-"`
}

// siteInfoNewResponseRulesetJSON contains the JSON metadata for the struct
// [SiteInfoNewResponseRuleset]
type siteInfoNewResponseRulesetJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	ZoneName    apijson.Field
	ZoneTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteInfoNewResponseRuleset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteInfoNewResponseRulesetJSON) RawJSON() string {
	return r.raw
}

type SiteInfoUpdateResponse struct {
	// If enabled, the JavaScript snippet is automatically injected for orange-clouded
	// sites.
	AutoInstall bool      `json:"auto_install"`
	Created     time.Time `json:"created" format:"date-time"`
	// A list of rules.
	Rules   []SiteInfoUpdateResponseRule  `json:"rules"`
	Ruleset SiteInfoUpdateResponseRuleset `json:"ruleset"`
	// The Web Analytics site identifier.
	SiteTag string `json:"site_tag"`
	// The Web Analytics site token.
	SiteToken string `json:"site_token"`
	// Encoded JavaScript snippet.
	Snippet string                     `json:"snippet"`
	JSON    siteInfoUpdateResponseJSON `json:"-"`
}

// siteInfoUpdateResponseJSON contains the JSON metadata for the struct
// [SiteInfoUpdateResponse]
type siteInfoUpdateResponseJSON struct {
	AutoInstall apijson.Field
	Created     apijson.Field
	Rules       apijson.Field
	Ruleset     apijson.Field
	SiteTag     apijson.Field
	SiteToken   apijson.Field
	Snippet     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteInfoUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteInfoUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type SiteInfoUpdateResponseRule struct {
	// The Web Analytics rule identifier.
	ID      string    `json:"id"`
	Created time.Time `json:"created" format:"date-time"`
	// The hostname the rule will be applied to.
	Host string `json:"host"`
	// Whether the rule includes or excludes traffic from being measured.
	Inclusive bool `json:"inclusive"`
	// Whether the rule is paused or not.
	IsPaused bool `json:"is_paused"`
	// The paths the rule will be applied to.
	Paths    []string                       `json:"paths"`
	Priority float64                        `json:"priority"`
	JSON     siteInfoUpdateResponseRuleJSON `json:"-"`
}

// siteInfoUpdateResponseRuleJSON contains the JSON metadata for the struct
// [SiteInfoUpdateResponseRule]
type siteInfoUpdateResponseRuleJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Host        apijson.Field
	Inclusive   apijson.Field
	IsPaused    apijson.Field
	Paths       apijson.Field
	Priority    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteInfoUpdateResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteInfoUpdateResponseRuleJSON) RawJSON() string {
	return r.raw
}

type SiteInfoUpdateResponseRuleset struct {
	// The Web Analytics ruleset identifier.
	ID string `json:"id"`
	// Whether the ruleset is enabled.
	Enabled  bool   `json:"enabled"`
	ZoneName string `json:"zone_name"`
	// The zone identifier.
	ZoneTag string                            `json:"zone_tag"`
	JSON    siteInfoUpdateResponseRulesetJSON `json:"-"`
}

// siteInfoUpdateResponseRulesetJSON contains the JSON metadata for the struct
// [SiteInfoUpdateResponseRuleset]
type siteInfoUpdateResponseRulesetJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	ZoneName    apijson.Field
	ZoneTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteInfoUpdateResponseRuleset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteInfoUpdateResponseRulesetJSON) RawJSON() string {
	return r.raw
}

type SiteInfoListResponse struct {
	// If enabled, the JavaScript snippet is automatically injected for orange-clouded
	// sites.
	AutoInstall bool      `json:"auto_install"`
	Created     time.Time `json:"created" format:"date-time"`
	// A list of rules.
	Rules   []SiteInfoListResponseRule  `json:"rules"`
	Ruleset SiteInfoListResponseRuleset `json:"ruleset"`
	// The Web Analytics site identifier.
	SiteTag string `json:"site_tag"`
	// The Web Analytics site token.
	SiteToken string `json:"site_token"`
	// Encoded JavaScript snippet.
	Snippet string                   `json:"snippet"`
	JSON    siteInfoListResponseJSON `json:"-"`
}

// siteInfoListResponseJSON contains the JSON metadata for the struct
// [SiteInfoListResponse]
type siteInfoListResponseJSON struct {
	AutoInstall apijson.Field
	Created     apijson.Field
	Rules       apijson.Field
	Ruleset     apijson.Field
	SiteTag     apijson.Field
	SiteToken   apijson.Field
	Snippet     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteInfoListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteInfoListResponseJSON) RawJSON() string {
	return r.raw
}

type SiteInfoListResponseRule struct {
	// The Web Analytics rule identifier.
	ID      string    `json:"id"`
	Created time.Time `json:"created" format:"date-time"`
	// The hostname the rule will be applied to.
	Host string `json:"host"`
	// Whether the rule includes or excludes traffic from being measured.
	Inclusive bool `json:"inclusive"`
	// Whether the rule is paused or not.
	IsPaused bool `json:"is_paused"`
	// The paths the rule will be applied to.
	Paths    []string                     `json:"paths"`
	Priority float64                      `json:"priority"`
	JSON     siteInfoListResponseRuleJSON `json:"-"`
}

// siteInfoListResponseRuleJSON contains the JSON metadata for the struct
// [SiteInfoListResponseRule]
type siteInfoListResponseRuleJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Host        apijson.Field
	Inclusive   apijson.Field
	IsPaused    apijson.Field
	Paths       apijson.Field
	Priority    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteInfoListResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteInfoListResponseRuleJSON) RawJSON() string {
	return r.raw
}

type SiteInfoListResponseRuleset struct {
	// The Web Analytics ruleset identifier.
	ID string `json:"id"`
	// Whether the ruleset is enabled.
	Enabled  bool   `json:"enabled"`
	ZoneName string `json:"zone_name"`
	// The zone identifier.
	ZoneTag string                          `json:"zone_tag"`
	JSON    siteInfoListResponseRulesetJSON `json:"-"`
}

// siteInfoListResponseRulesetJSON contains the JSON metadata for the struct
// [SiteInfoListResponseRuleset]
type siteInfoListResponseRulesetJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	ZoneName    apijson.Field
	ZoneTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteInfoListResponseRuleset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteInfoListResponseRulesetJSON) RawJSON() string {
	return r.raw
}

type SiteInfoDeleteResponse struct {
	// The Web Analytics site identifier.
	SiteTag string                     `json:"site_tag"`
	JSON    siteInfoDeleteResponseJSON `json:"-"`
}

// siteInfoDeleteResponseJSON contains the JSON metadata for the struct
// [SiteInfoDeleteResponse]
type siteInfoDeleteResponseJSON struct {
	SiteTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteInfoDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteInfoDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SiteInfoGetResponse struct {
	// If enabled, the JavaScript snippet is automatically injected for orange-clouded
	// sites.
	AutoInstall bool      `json:"auto_install"`
	Created     time.Time `json:"created" format:"date-time"`
	// A list of rules.
	Rules   []SiteInfoGetResponseRule  `json:"rules"`
	Ruleset SiteInfoGetResponseRuleset `json:"ruleset"`
	// The Web Analytics site identifier.
	SiteTag string `json:"site_tag"`
	// The Web Analytics site token.
	SiteToken string `json:"site_token"`
	// Encoded JavaScript snippet.
	Snippet string                  `json:"snippet"`
	JSON    siteInfoGetResponseJSON `json:"-"`
}

// siteInfoGetResponseJSON contains the JSON metadata for the struct
// [SiteInfoGetResponse]
type siteInfoGetResponseJSON struct {
	AutoInstall apijson.Field
	Created     apijson.Field
	Rules       apijson.Field
	Ruleset     apijson.Field
	SiteTag     apijson.Field
	SiteToken   apijson.Field
	Snippet     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteInfoGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteInfoGetResponseJSON) RawJSON() string {
	return r.raw
}

type SiteInfoGetResponseRule struct {
	// The Web Analytics rule identifier.
	ID      string    `json:"id"`
	Created time.Time `json:"created" format:"date-time"`
	// The hostname the rule will be applied to.
	Host string `json:"host"`
	// Whether the rule includes or excludes traffic from being measured.
	Inclusive bool `json:"inclusive"`
	// Whether the rule is paused or not.
	IsPaused bool `json:"is_paused"`
	// The paths the rule will be applied to.
	Paths    []string                    `json:"paths"`
	Priority float64                     `json:"priority"`
	JSON     siteInfoGetResponseRuleJSON `json:"-"`
}

// siteInfoGetResponseRuleJSON contains the JSON metadata for the struct
// [SiteInfoGetResponseRule]
type siteInfoGetResponseRuleJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Host        apijson.Field
	Inclusive   apijson.Field
	IsPaused    apijson.Field
	Paths       apijson.Field
	Priority    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteInfoGetResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteInfoGetResponseRuleJSON) RawJSON() string {
	return r.raw
}

type SiteInfoGetResponseRuleset struct {
	// The Web Analytics ruleset identifier.
	ID string `json:"id"`
	// Whether the ruleset is enabled.
	Enabled  bool   `json:"enabled"`
	ZoneName string `json:"zone_name"`
	// The zone identifier.
	ZoneTag string                         `json:"zone_tag"`
	JSON    siteInfoGetResponseRulesetJSON `json:"-"`
}

// siteInfoGetResponseRulesetJSON contains the JSON metadata for the struct
// [SiteInfoGetResponseRuleset]
type siteInfoGetResponseRulesetJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	ZoneName    apijson.Field
	ZoneTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteInfoGetResponseRuleset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteInfoGetResponseRulesetJSON) RawJSON() string {
	return r.raw
}

type SiteInfoNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// If enabled, the JavaScript snippet is automatically injected for orange-clouded
	// sites.
	AutoInstall param.Field[bool] `json:"auto_install"`
	// The hostname to use for gray-clouded sites.
	Host param.Field[string] `json:"host"`
	// The zone identifier.
	ZoneTag param.Field[string] `json:"zone_tag"`
}

func (r SiteInfoNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteInfoNewResponseEnvelope struct {
	Result SiteInfoNewResponse             `json:"result"`
	JSON   siteInfoNewResponseEnvelopeJSON `json:"-"`
}

// siteInfoNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteInfoNewResponseEnvelope]
type siteInfoNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteInfoNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteInfoNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SiteInfoUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// If enabled, the JavaScript snippet is automatically injected for orange-clouded
	// sites.
	AutoInstall param.Field[bool] `json:"auto_install"`
	// The hostname to use for gray-clouded sites.
	Host param.Field[string] `json:"host"`
	// The zone identifier.
	ZoneTag param.Field[string] `json:"zone_tag"`
}

func (r SiteInfoUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteInfoUpdateResponseEnvelope struct {
	Result SiteInfoUpdateResponse             `json:"result"`
	JSON   siteInfoUpdateResponseEnvelopeJSON `json:"-"`
}

// siteInfoUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteInfoUpdateResponseEnvelope]
type siteInfoUpdateResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteInfoUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteInfoUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SiteInfoListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The property used to sort the list of results.
	OrderBy param.Field[SiteInfoListParamsOrderBy] `query:"order_by"`
	// Current page within the paginated list of results.
	Page param.Field[float64] `query:"page"`
	// Number of items to return per page of results.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [SiteInfoListParams]'s query parameters as `url.Values`.
func (r SiteInfoListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The property used to sort the list of results.
type SiteInfoListParamsOrderBy string

const (
	SiteInfoListParamsOrderByHost    SiteInfoListParamsOrderBy = "host"
	SiteInfoListParamsOrderByCreated SiteInfoListParamsOrderBy = "created"
)

type SiteInfoDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SiteInfoDeleteResponseEnvelope struct {
	Result SiteInfoDeleteResponse             `json:"result"`
	JSON   siteInfoDeleteResponseEnvelopeJSON `json:"-"`
}

// siteInfoDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteInfoDeleteResponseEnvelope]
type siteInfoDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteInfoDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteInfoDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SiteInfoGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SiteInfoGetResponseEnvelope struct {
	Result SiteInfoGetResponse             `json:"result"`
	JSON   siteInfoGetResponseEnvelopeJSON `json:"-"`
}

// siteInfoGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteInfoGetResponseEnvelope]
type siteInfoGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteInfoGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteInfoGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
