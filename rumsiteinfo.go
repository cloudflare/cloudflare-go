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

// RUMSiteInfoService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRUMSiteInfoService] method
// instead.
type RUMSiteInfoService struct {
	Options []option.RequestOption
}

// NewRUMSiteInfoService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRUMSiteInfoService(opts ...option.RequestOption) (r *RUMSiteInfoService) {
	r = &RUMSiteInfoService{}
	r.Options = opts
	return
}

// Creates a new Web Analytics site.
func (r *RUMSiteInfoService) New(ctx context.Context, accountID string, body RUMSiteInfoNewParams, opts ...option.RequestOption) (res *RUMSiteInfoNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RUMSiteInfoNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rum/site_info", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing Web Analytics site.
func (r *RUMSiteInfoService) Update(ctx context.Context, accountID string, siteID string, body RUMSiteInfoUpdateParams, opts ...option.RequestOption) (res *RUMSiteInfoUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RUMSiteInfoUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rum/site_info/%s", accountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all Web Analytics sites of an account.
func (r *RUMSiteInfoService) List(ctx context.Context, accountID string, query RUMSiteInfoListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[RUMSiteInfoListResponse], err error) {
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
func (r *RUMSiteInfoService) ListAutoPaging(ctx context.Context, accountID string, query RUMSiteInfoListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[RUMSiteInfoListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, accountID, query, opts...))
}

// Deletes an existing Web Analytics site.
func (r *RUMSiteInfoService) Delete(ctx context.Context, accountID string, siteID string, opts ...option.RequestOption) (res *RUMSiteInfoDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RUMSiteInfoDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rum/site_info/%s", accountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves a Web Analytics site.
func (r *RUMSiteInfoService) Get(ctx context.Context, accountID string, siteID string, opts ...option.RequestOption) (res *RUMSiteInfoGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RUMSiteInfoGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rum/site_info/%s", accountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RUMSiteInfoNewResponse struct {
	// If enabled, the JavaScript snippet is automatically injected for orange-clouded
	// sites.
	AutoInstall bool      `json:"auto_install"`
	Created     time.Time `json:"created" format:"date-time"`
	// A list of rules.
	Rules   []RUMSiteInfoNewResponseRule  `json:"rules"`
	Ruleset RUMSiteInfoNewResponseRuleset `json:"ruleset"`
	// The Web Analytics site identifier.
	SiteTag string `json:"site_tag"`
	// The Web Analytics site token.
	SiteToken string `json:"site_token"`
	// Encoded JavaScript snippet.
	Snippet string                     `json:"snippet"`
	JSON    rumSiteInfoNewResponseJSON `json:"-"`
}

// rumSiteInfoNewResponseJSON contains the JSON metadata for the struct
// [RUMSiteInfoNewResponse]
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

func (r *RUMSiteInfoNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RUMSiteInfoNewResponseRule struct {
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
// [RUMSiteInfoNewResponseRule]
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

func (r *RUMSiteInfoNewResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RUMSiteInfoNewResponseRuleset struct {
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
// [RUMSiteInfoNewResponseRuleset]
type rumSiteInfoNewResponseRulesetJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	ZoneName    apijson.Field
	ZoneTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RUMSiteInfoNewResponseRuleset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RUMSiteInfoUpdateResponse struct {
	// If enabled, the JavaScript snippet is automatically injected for orange-clouded
	// sites.
	AutoInstall bool      `json:"auto_install"`
	Created     time.Time `json:"created" format:"date-time"`
	// A list of rules.
	Rules   []RUMSiteInfoUpdateResponseRule  `json:"rules"`
	Ruleset RUMSiteInfoUpdateResponseRuleset `json:"ruleset"`
	// The Web Analytics site identifier.
	SiteTag string `json:"site_tag"`
	// The Web Analytics site token.
	SiteToken string `json:"site_token"`
	// Encoded JavaScript snippet.
	Snippet string                        `json:"snippet"`
	JSON    rumSiteInfoUpdateResponseJSON `json:"-"`
}

// rumSiteInfoUpdateResponseJSON contains the JSON metadata for the struct
// [RUMSiteInfoUpdateResponse]
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

func (r *RUMSiteInfoUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RUMSiteInfoUpdateResponseRule struct {
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
// [RUMSiteInfoUpdateResponseRule]
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

func (r *RUMSiteInfoUpdateResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RUMSiteInfoUpdateResponseRuleset struct {
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
// [RUMSiteInfoUpdateResponseRuleset]
type rumSiteInfoUpdateResponseRulesetJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	ZoneName    apijson.Field
	ZoneTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RUMSiteInfoUpdateResponseRuleset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RUMSiteInfoListResponse struct {
	// If enabled, the JavaScript snippet is automatically injected for orange-clouded
	// sites.
	AutoInstall bool      `json:"auto_install"`
	Created     time.Time `json:"created" format:"date-time"`
	// A list of rules.
	Rules   []RUMSiteInfoListResponseRule  `json:"rules"`
	Ruleset RUMSiteInfoListResponseRuleset `json:"ruleset"`
	// The Web Analytics site identifier.
	SiteTag string `json:"site_tag"`
	// The Web Analytics site token.
	SiteToken string `json:"site_token"`
	// Encoded JavaScript snippet.
	Snippet string                      `json:"snippet"`
	JSON    rumSiteInfoListResponseJSON `json:"-"`
}

// rumSiteInfoListResponseJSON contains the JSON metadata for the struct
// [RUMSiteInfoListResponse]
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

func (r *RUMSiteInfoListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RUMSiteInfoListResponseRule struct {
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
// [RUMSiteInfoListResponseRule]
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

func (r *RUMSiteInfoListResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RUMSiteInfoListResponseRuleset struct {
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
// [RUMSiteInfoListResponseRuleset]
type rumSiteInfoListResponseRulesetJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	ZoneName    apijson.Field
	ZoneTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RUMSiteInfoListResponseRuleset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RUMSiteInfoDeleteResponse struct {
	// The Web Analytics site identifier.
	SiteTag string                        `json:"site_tag"`
	JSON    rumSiteInfoDeleteResponseJSON `json:"-"`
}

// rumSiteInfoDeleteResponseJSON contains the JSON metadata for the struct
// [RUMSiteInfoDeleteResponse]
type rumSiteInfoDeleteResponseJSON struct {
	SiteTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RUMSiteInfoDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RUMSiteInfoGetResponse struct {
	// If enabled, the JavaScript snippet is automatically injected for orange-clouded
	// sites.
	AutoInstall bool      `json:"auto_install"`
	Created     time.Time `json:"created" format:"date-time"`
	// A list of rules.
	Rules   []RUMSiteInfoGetResponseRule  `json:"rules"`
	Ruleset RUMSiteInfoGetResponseRuleset `json:"ruleset"`
	// The Web Analytics site identifier.
	SiteTag string `json:"site_tag"`
	// The Web Analytics site token.
	SiteToken string `json:"site_token"`
	// Encoded JavaScript snippet.
	Snippet string                     `json:"snippet"`
	JSON    rumSiteInfoGetResponseJSON `json:"-"`
}

// rumSiteInfoGetResponseJSON contains the JSON metadata for the struct
// [RUMSiteInfoGetResponse]
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

func (r *RUMSiteInfoGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RUMSiteInfoGetResponseRule struct {
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
// [RUMSiteInfoGetResponseRule]
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

func (r *RUMSiteInfoGetResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RUMSiteInfoGetResponseRuleset struct {
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
// [RUMSiteInfoGetResponseRuleset]
type rumSiteInfoGetResponseRulesetJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	ZoneName    apijson.Field
	ZoneTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RUMSiteInfoGetResponseRuleset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RUMSiteInfoNewParams struct {
	// If enabled, the JavaScript snippet is automatically injected for orange-clouded
	// sites.
	AutoInstall param.Field[bool] `json:"auto_install"`
	// The hostname to use for gray-clouded sites.
	Host param.Field[string] `json:"host"`
	// The zone identifier.
	ZoneTag param.Field[string] `json:"zone_tag"`
}

func (r RUMSiteInfoNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RUMSiteInfoNewResponseEnvelope struct {
	Result RUMSiteInfoNewResponse             `json:"result"`
	JSON   rumSiteInfoNewResponseEnvelopeJSON `json:"-"`
}

// rumSiteInfoNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RUMSiteInfoNewResponseEnvelope]
type rumSiteInfoNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RUMSiteInfoNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RUMSiteInfoUpdateParams struct {
	// If enabled, the JavaScript snippet is automatically injected for orange-clouded
	// sites.
	AutoInstall param.Field[bool] `json:"auto_install"`
	// The hostname to use for gray-clouded sites.
	Host param.Field[string] `json:"host"`
	// The zone identifier.
	ZoneTag param.Field[string] `json:"zone_tag"`
}

func (r RUMSiteInfoUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RUMSiteInfoUpdateResponseEnvelope struct {
	Result RUMSiteInfoUpdateResponse             `json:"result"`
	JSON   rumSiteInfoUpdateResponseEnvelopeJSON `json:"-"`
}

// rumSiteInfoUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [RUMSiteInfoUpdateResponseEnvelope]
type rumSiteInfoUpdateResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RUMSiteInfoUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RUMSiteInfoListParams struct {
	// The property used to sort the list of results.
	OrderBy param.Field[RUMSiteInfoListParamsOrderBy] `query:"order_by"`
	// Current page within the paginated list of results.
	Page param.Field[float64] `query:"page"`
	// Number of items to return per page of results.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [RUMSiteInfoListParams]'s query parameters as `url.Values`.
func (r RUMSiteInfoListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The property used to sort the list of results.
type RUMSiteInfoListParamsOrderBy string

const (
	RUMSiteInfoListParamsOrderByHost    RUMSiteInfoListParamsOrderBy = "host"
	RUMSiteInfoListParamsOrderByCreated RUMSiteInfoListParamsOrderBy = "created"
)

type RUMSiteInfoDeleteResponseEnvelope struct {
	Result RUMSiteInfoDeleteResponse             `json:"result"`
	JSON   rumSiteInfoDeleteResponseEnvelopeJSON `json:"-"`
}

// rumSiteInfoDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [RUMSiteInfoDeleteResponseEnvelope]
type rumSiteInfoDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RUMSiteInfoDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RUMSiteInfoGetResponseEnvelope struct {
	Result RUMSiteInfoGetResponse             `json:"result"`
	JSON   rumSiteInfoGetResponseEnvelopeJSON `json:"-"`
}

// rumSiteInfoGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RUMSiteInfoGetResponseEnvelope]
type rumSiteInfoGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RUMSiteInfoGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
