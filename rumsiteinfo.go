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

// RumSiteInfoService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRumSiteInfoService] method
// instead.
type RumSiteInfoService struct {
	Options []option.RequestOption
}

// NewRumSiteInfoService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRumSiteInfoService(opts ...option.RequestOption) (r *RumSiteInfoService) {
	r = &RumSiteInfoService{}
	r.Options = opts
	return
}

// Creates a new Web Analytics site.
func (r *RumSiteInfoService) New(ctx context.Context, accountID string, body RumSiteInfoNewParams, opts ...option.RequestOption) (res *RumSiteInfoNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RumSiteInfoNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rum/site_info", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing Web Analytics site.
func (r *RumSiteInfoService) Update(ctx context.Context, accountID string, siteID string, body RumSiteInfoUpdateParams, opts ...option.RequestOption) (res *RumSiteInfoUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RumSiteInfoUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rum/site_info/%s", accountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all Web Analytics sites of an account.
func (r *RumSiteInfoService) List(ctx context.Context, accountID string, query RumSiteInfoListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[RumSiteInfoListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/rum/site_info/list", accountID)
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

// Lists all Web Analytics sites of an account.
func (r *RumSiteInfoService) ListAutoPaging(ctx context.Context, accountID string, query RumSiteInfoListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[RumSiteInfoListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, accountID, query, opts...))
}

// Deletes an existing Web Analytics site.
func (r *RumSiteInfoService) Delete(ctx context.Context, accountID string, siteID string, opts ...option.RequestOption) (res *RumSiteInfoDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RumSiteInfoDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rum/site_info/%s", accountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves a Web Analytics site.
func (r *RumSiteInfoService) Get(ctx context.Context, accountID string, siteID string, opts ...option.RequestOption) (res *RumSiteInfoGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RumSiteInfoGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rum/site_info/%s", accountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RumSiteInfoNewResponse struct {
	// If enabled, the JavaScript snippet is automatically injected for orange-clouded
	// sites.
	AutoInstall bool      `json:"auto_install"`
	Created     time.Time `json:"created" format:"date-time"`
	// A list of rules.
	Rules   []RumSiteInfoNewResponseRule  `json:"rules"`
	Ruleset RumSiteInfoNewResponseRuleset `json:"ruleset"`
	// The Web Analytics site identifier.
	SiteTag string `json:"site_tag"`
	// The Web Analytics site token.
	SiteToken string `json:"site_token"`
	// Encoded JavaScript snippet.
	Snippet string                     `json:"snippet"`
	JSON    rumSiteInfoNewResponseJSON `json:"-"`
}

// rumSiteInfoNewResponseJSON contains the JSON metadata for the struct
// [RumSiteInfoNewResponse]
type rumSiteInfoNewResponseJSON struct {
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

func (r *RumSiteInfoNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoNewResponseRule struct {
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
	JSON     rumSiteInfoNewResponseRuleJSON `json:"-"`
}

// rumSiteInfoNewResponseRuleJSON contains the JSON metadata for the struct
// [RumSiteInfoNewResponseRule]
type rumSiteInfoNewResponseRuleJSON struct {
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

func (r *RumSiteInfoNewResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoNewResponseRuleset struct {
	// The Web Analytics ruleset identifier.
	ID string `json:"id"`
	// Whether the ruleset is enabled.
	Enabled  bool   `json:"enabled"`
	ZoneName string `json:"zone_name"`
	// The zone identifier.
	ZoneTag string                            `json:"zone_tag"`
	JSON    rumSiteInfoNewResponseRulesetJSON `json:"-"`
}

// rumSiteInfoNewResponseRulesetJSON contains the JSON metadata for the struct
// [RumSiteInfoNewResponseRuleset]
type rumSiteInfoNewResponseRulesetJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	ZoneName    apijson.Field
	ZoneTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoNewResponseRuleset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoUpdateResponse struct {
	// If enabled, the JavaScript snippet is automatically injected for orange-clouded
	// sites.
	AutoInstall bool      `json:"auto_install"`
	Created     time.Time `json:"created" format:"date-time"`
	// A list of rules.
	Rules   []RumSiteInfoUpdateResponseRule  `json:"rules"`
	Ruleset RumSiteInfoUpdateResponseRuleset `json:"ruleset"`
	// The Web Analytics site identifier.
	SiteTag string `json:"site_tag"`
	// The Web Analytics site token.
	SiteToken string `json:"site_token"`
	// Encoded JavaScript snippet.
	Snippet string                        `json:"snippet"`
	JSON    rumSiteInfoUpdateResponseJSON `json:"-"`
}

// rumSiteInfoUpdateResponseJSON contains the JSON metadata for the struct
// [RumSiteInfoUpdateResponse]
type rumSiteInfoUpdateResponseJSON struct {
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

func (r *RumSiteInfoUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoUpdateResponseRule struct {
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
	Paths    []string                          `json:"paths"`
	Priority float64                           `json:"priority"`
	JSON     rumSiteInfoUpdateResponseRuleJSON `json:"-"`
}

// rumSiteInfoUpdateResponseRuleJSON contains the JSON metadata for the struct
// [RumSiteInfoUpdateResponseRule]
type rumSiteInfoUpdateResponseRuleJSON struct {
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

func (r *RumSiteInfoUpdateResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoUpdateResponseRuleset struct {
	// The Web Analytics ruleset identifier.
	ID string `json:"id"`
	// Whether the ruleset is enabled.
	Enabled  bool   `json:"enabled"`
	ZoneName string `json:"zone_name"`
	// The zone identifier.
	ZoneTag string                               `json:"zone_tag"`
	JSON    rumSiteInfoUpdateResponseRulesetJSON `json:"-"`
}

// rumSiteInfoUpdateResponseRulesetJSON contains the JSON metadata for the struct
// [RumSiteInfoUpdateResponseRuleset]
type rumSiteInfoUpdateResponseRulesetJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	ZoneName    apijson.Field
	ZoneTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoUpdateResponseRuleset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoListResponse struct {
	// If enabled, the JavaScript snippet is automatically injected for orange-clouded
	// sites.
	AutoInstall bool      `json:"auto_install"`
	Created     time.Time `json:"created" format:"date-time"`
	// A list of rules.
	Rules   []RumSiteInfoListResponseRule  `json:"rules"`
	Ruleset RumSiteInfoListResponseRuleset `json:"ruleset"`
	// The Web Analytics site identifier.
	SiteTag string `json:"site_tag"`
	// The Web Analytics site token.
	SiteToken string `json:"site_token"`
	// Encoded JavaScript snippet.
	Snippet string                      `json:"snippet"`
	JSON    rumSiteInfoListResponseJSON `json:"-"`
}

// rumSiteInfoListResponseJSON contains the JSON metadata for the struct
// [RumSiteInfoListResponse]
type rumSiteInfoListResponseJSON struct {
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

func (r *RumSiteInfoListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoListResponseRule struct {
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
	Paths    []string                        `json:"paths"`
	Priority float64                         `json:"priority"`
	JSON     rumSiteInfoListResponseRuleJSON `json:"-"`
}

// rumSiteInfoListResponseRuleJSON contains the JSON metadata for the struct
// [RumSiteInfoListResponseRule]
type rumSiteInfoListResponseRuleJSON struct {
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

func (r *RumSiteInfoListResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoListResponseRuleset struct {
	// The Web Analytics ruleset identifier.
	ID string `json:"id"`
	// Whether the ruleset is enabled.
	Enabled  bool   `json:"enabled"`
	ZoneName string `json:"zone_name"`
	// The zone identifier.
	ZoneTag string                             `json:"zone_tag"`
	JSON    rumSiteInfoListResponseRulesetJSON `json:"-"`
}

// rumSiteInfoListResponseRulesetJSON contains the JSON metadata for the struct
// [RumSiteInfoListResponseRuleset]
type rumSiteInfoListResponseRulesetJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	ZoneName    apijson.Field
	ZoneTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoListResponseRuleset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoDeleteResponse struct {
	// The Web Analytics site identifier.
	SiteTag string                        `json:"site_tag"`
	JSON    rumSiteInfoDeleteResponseJSON `json:"-"`
}

// rumSiteInfoDeleteResponseJSON contains the JSON metadata for the struct
// [RumSiteInfoDeleteResponse]
type rumSiteInfoDeleteResponseJSON struct {
	SiteTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoGetResponse struct {
	// If enabled, the JavaScript snippet is automatically injected for orange-clouded
	// sites.
	AutoInstall bool      `json:"auto_install"`
	Created     time.Time `json:"created" format:"date-time"`
	// A list of rules.
	Rules   []RumSiteInfoGetResponseRule  `json:"rules"`
	Ruleset RumSiteInfoGetResponseRuleset `json:"ruleset"`
	// The Web Analytics site identifier.
	SiteTag string `json:"site_tag"`
	// The Web Analytics site token.
	SiteToken string `json:"site_token"`
	// Encoded JavaScript snippet.
	Snippet string                     `json:"snippet"`
	JSON    rumSiteInfoGetResponseJSON `json:"-"`
}

// rumSiteInfoGetResponseJSON contains the JSON metadata for the struct
// [RumSiteInfoGetResponse]
type rumSiteInfoGetResponseJSON struct {
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

func (r *RumSiteInfoGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoGetResponseRule struct {
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
	JSON     rumSiteInfoGetResponseRuleJSON `json:"-"`
}

// rumSiteInfoGetResponseRuleJSON contains the JSON metadata for the struct
// [RumSiteInfoGetResponseRule]
type rumSiteInfoGetResponseRuleJSON struct {
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

func (r *RumSiteInfoGetResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoGetResponseRuleset struct {
	// The Web Analytics ruleset identifier.
	ID string `json:"id"`
	// Whether the ruleset is enabled.
	Enabled  bool   `json:"enabled"`
	ZoneName string `json:"zone_name"`
	// The zone identifier.
	ZoneTag string                            `json:"zone_tag"`
	JSON    rumSiteInfoGetResponseRulesetJSON `json:"-"`
}

// rumSiteInfoGetResponseRulesetJSON contains the JSON metadata for the struct
// [RumSiteInfoGetResponseRuleset]
type rumSiteInfoGetResponseRulesetJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	ZoneName    apijson.Field
	ZoneTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoGetResponseRuleset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoNewParams struct {
	// If enabled, the JavaScript snippet is automatically injected for orange-clouded
	// sites.
	AutoInstall param.Field[bool] `json:"auto_install"`
	// The hostname to use for gray-clouded sites.
	Host param.Field[string] `json:"host"`
	// The zone identifier.
	ZoneTag param.Field[string] `json:"zone_tag"`
}

func (r RumSiteInfoNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RumSiteInfoNewResponseEnvelope struct {
	Result RumSiteInfoNewResponse             `json:"result"`
	JSON   rumSiteInfoNewResponseEnvelopeJSON `json:"-"`
}

// rumSiteInfoNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RumSiteInfoNewResponseEnvelope]
type rumSiteInfoNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoUpdateParams struct {
	// If enabled, the JavaScript snippet is automatically injected for orange-clouded
	// sites.
	AutoInstall param.Field[bool] `json:"auto_install"`
	// The hostname to use for gray-clouded sites.
	Host param.Field[string] `json:"host"`
	// The zone identifier.
	ZoneTag param.Field[string] `json:"zone_tag"`
}

func (r RumSiteInfoUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RumSiteInfoUpdateResponseEnvelope struct {
	Result RumSiteInfoUpdateResponse             `json:"result"`
	JSON   rumSiteInfoUpdateResponseEnvelopeJSON `json:"-"`
}

// rumSiteInfoUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [RumSiteInfoUpdateResponseEnvelope]
type rumSiteInfoUpdateResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoListParams struct {
	// The property used to sort the list of results.
	OrderBy param.Field[RumSiteInfoListParamsOrderBy] `query:"order_by"`
	// Current page within the paginated list of results.
	Page param.Field[float64] `query:"page"`
	// Number of items to return per page of results.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [RumSiteInfoListParams]'s query parameters as `url.Values`.
func (r RumSiteInfoListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The property used to sort the list of results.
type RumSiteInfoListParamsOrderBy string

const (
	RumSiteInfoListParamsOrderByHost    RumSiteInfoListParamsOrderBy = "host"
	RumSiteInfoListParamsOrderByCreated RumSiteInfoListParamsOrderBy = "created"
)

type RumSiteInfoDeleteResponseEnvelope struct {
	Result RumSiteInfoDeleteResponse             `json:"result"`
	JSON   rumSiteInfoDeleteResponseEnvelopeJSON `json:"-"`
}

// rumSiteInfoDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [RumSiteInfoDeleteResponseEnvelope]
type rumSiteInfoDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoGetResponseEnvelope struct {
	Result RumSiteInfoGetResponse             `json:"result"`
	JSON   rumSiteInfoGetResponseEnvelopeJSON `json:"-"`
}

// rumSiteInfoGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RumSiteInfoGetResponseEnvelope]
type rumSiteInfoGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
