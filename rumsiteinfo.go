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
func (r *RumSiteInfoService) New(ctx context.Context, accountIdentifier string, body RumSiteInfoNewParams, opts ...option.RequestOption) (res *RumSiteInfoNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rum/site_info", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves a Web Analytics site.
func (r *RumSiteInfoService) Get(ctx context.Context, accountIdentifier string, siteIdentifier string, opts ...option.RequestOption) (res *RumSiteInfoGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rum/site_info/%s", accountIdentifier, siteIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an existing Web Analytics site.
func (r *RumSiteInfoService) Update(ctx context.Context, accountIdentifier string, siteIdentifier string, body RumSiteInfoUpdateParams, opts ...option.RequestOption) (res *RumSiteInfoUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rum/site_info/%s", accountIdentifier, siteIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists all Web Analytics sites of an account.
func (r *RumSiteInfoService) List(ctx context.Context, accountIdentifier string, query RumSiteInfoListParams, opts ...option.RequestOption) (res *RumSiteInfoListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rum/site_info/list", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes an existing Web Analytics site.
func (r *RumSiteInfoService) Delete(ctx context.Context, accountIdentifier string, siteIdentifier string, opts ...option.RequestOption) (res *RumSiteInfoDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rum/site_info/%s", accountIdentifier, siteIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type RumSiteInfoNewResponse struct {
	Errors   []RumSiteInfoNewResponseError   `json:"errors"`
	Messages []RumSiteInfoNewResponseMessage `json:"messages"`
	Result   RumSiteInfoNewResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success bool                       `json:"success"`
	JSON    rumSiteInfoNewResponseJSON `json:"-"`
}

// rumSiteInfoNewResponseJSON contains the JSON metadata for the struct
// [RumSiteInfoNewResponse]
type rumSiteInfoNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoNewResponseError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    rumSiteInfoNewResponseErrorJSON `json:"-"`
}

// rumSiteInfoNewResponseErrorJSON contains the JSON metadata for the struct
// [RumSiteInfoNewResponseError]
type rumSiteInfoNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoNewResponseMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    rumSiteInfoNewResponseMessageJSON `json:"-"`
}

// rumSiteInfoNewResponseMessageJSON contains the JSON metadata for the struct
// [RumSiteInfoNewResponseMessage]
type rumSiteInfoNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoNewResponseResult struct {
	// If enabled, the JavaScript snippet is automatically injected for orange-clouded
	// sites.
	AutoInstall bool      `json:"auto_install"`
	Created     time.Time `json:"created" format:"date-time"`
	// A list of rules.
	Rules   []RumSiteInfoNewResponseResultRule  `json:"rules"`
	Ruleset RumSiteInfoNewResponseResultRuleset `json:"ruleset"`
	// The Web Analytics site identifier.
	SiteTag string `json:"site_tag"`
	// The Web Analytics site token.
	SiteToken string `json:"site_token"`
	// Encoded JavaScript snippet.
	Snippet string                           `json:"snippet"`
	JSON    rumSiteInfoNewResponseResultJSON `json:"-"`
}

// rumSiteInfoNewResponseResultJSON contains the JSON metadata for the struct
// [RumSiteInfoNewResponseResult]
type rumSiteInfoNewResponseResultJSON struct {
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

func (r *RumSiteInfoNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoNewResponseResultRule struct {
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
	Paths    []string                             `json:"paths"`
	Priority float64                              `json:"priority"`
	JSON     rumSiteInfoNewResponseResultRuleJSON `json:"-"`
}

// rumSiteInfoNewResponseResultRuleJSON contains the JSON metadata for the struct
// [RumSiteInfoNewResponseResultRule]
type rumSiteInfoNewResponseResultRuleJSON struct {
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

func (r *RumSiteInfoNewResponseResultRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoNewResponseResultRuleset struct {
	// The Web Analytics ruleset identifier.
	ID string `json:"id"`
	// Whether the ruleset is enabled.
	Enabled  bool   `json:"enabled"`
	ZoneName string `json:"zone_name"`
	// The zone identifier.
	ZoneTag string                                  `json:"zone_tag"`
	JSON    rumSiteInfoNewResponseResultRulesetJSON `json:"-"`
}

// rumSiteInfoNewResponseResultRulesetJSON contains the JSON metadata for the
// struct [RumSiteInfoNewResponseResultRuleset]
type rumSiteInfoNewResponseResultRulesetJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	ZoneName    apijson.Field
	ZoneTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoNewResponseResultRuleset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoGetResponse struct {
	Errors   []RumSiteInfoGetResponseError   `json:"errors"`
	Messages []RumSiteInfoGetResponseMessage `json:"messages"`
	Result   RumSiteInfoGetResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success bool                       `json:"success"`
	JSON    rumSiteInfoGetResponseJSON `json:"-"`
}

// rumSiteInfoGetResponseJSON contains the JSON metadata for the struct
// [RumSiteInfoGetResponse]
type rumSiteInfoGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoGetResponseError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    rumSiteInfoGetResponseErrorJSON `json:"-"`
}

// rumSiteInfoGetResponseErrorJSON contains the JSON metadata for the struct
// [RumSiteInfoGetResponseError]
type rumSiteInfoGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoGetResponseMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    rumSiteInfoGetResponseMessageJSON `json:"-"`
}

// rumSiteInfoGetResponseMessageJSON contains the JSON metadata for the struct
// [RumSiteInfoGetResponseMessage]
type rumSiteInfoGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoGetResponseResult struct {
	// If enabled, the JavaScript snippet is automatically injected for orange-clouded
	// sites.
	AutoInstall bool      `json:"auto_install"`
	Created     time.Time `json:"created" format:"date-time"`
	// A list of rules.
	Rules   []RumSiteInfoGetResponseResultRule  `json:"rules"`
	Ruleset RumSiteInfoGetResponseResultRuleset `json:"ruleset"`
	// The Web Analytics site identifier.
	SiteTag string `json:"site_tag"`
	// The Web Analytics site token.
	SiteToken string `json:"site_token"`
	// Encoded JavaScript snippet.
	Snippet string                           `json:"snippet"`
	JSON    rumSiteInfoGetResponseResultJSON `json:"-"`
}

// rumSiteInfoGetResponseResultJSON contains the JSON metadata for the struct
// [RumSiteInfoGetResponseResult]
type rumSiteInfoGetResponseResultJSON struct {
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

func (r *RumSiteInfoGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoGetResponseResultRule struct {
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
	Paths    []string                             `json:"paths"`
	Priority float64                              `json:"priority"`
	JSON     rumSiteInfoGetResponseResultRuleJSON `json:"-"`
}

// rumSiteInfoGetResponseResultRuleJSON contains the JSON metadata for the struct
// [RumSiteInfoGetResponseResultRule]
type rumSiteInfoGetResponseResultRuleJSON struct {
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

func (r *RumSiteInfoGetResponseResultRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoGetResponseResultRuleset struct {
	// The Web Analytics ruleset identifier.
	ID string `json:"id"`
	// Whether the ruleset is enabled.
	Enabled  bool   `json:"enabled"`
	ZoneName string `json:"zone_name"`
	// The zone identifier.
	ZoneTag string                                  `json:"zone_tag"`
	JSON    rumSiteInfoGetResponseResultRulesetJSON `json:"-"`
}

// rumSiteInfoGetResponseResultRulesetJSON contains the JSON metadata for the
// struct [RumSiteInfoGetResponseResultRuleset]
type rumSiteInfoGetResponseResultRulesetJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	ZoneName    apijson.Field
	ZoneTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoGetResponseResultRuleset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoUpdateResponse struct {
	Errors   []RumSiteInfoUpdateResponseError   `json:"errors"`
	Messages []RumSiteInfoUpdateResponseMessage `json:"messages"`
	Result   RumSiteInfoUpdateResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success bool                          `json:"success"`
	JSON    rumSiteInfoUpdateResponseJSON `json:"-"`
}

// rumSiteInfoUpdateResponseJSON contains the JSON metadata for the struct
// [RumSiteInfoUpdateResponse]
type rumSiteInfoUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoUpdateResponseError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    rumSiteInfoUpdateResponseErrorJSON `json:"-"`
}

// rumSiteInfoUpdateResponseErrorJSON contains the JSON metadata for the struct
// [RumSiteInfoUpdateResponseError]
type rumSiteInfoUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoUpdateResponseMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    rumSiteInfoUpdateResponseMessageJSON `json:"-"`
}

// rumSiteInfoUpdateResponseMessageJSON contains the JSON metadata for the struct
// [RumSiteInfoUpdateResponseMessage]
type rumSiteInfoUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoUpdateResponseResult struct {
	// If enabled, the JavaScript snippet is automatically injected for orange-clouded
	// sites.
	AutoInstall bool      `json:"auto_install"`
	Created     time.Time `json:"created" format:"date-time"`
	// A list of rules.
	Rules   []RumSiteInfoUpdateResponseResultRule  `json:"rules"`
	Ruleset RumSiteInfoUpdateResponseResultRuleset `json:"ruleset"`
	// The Web Analytics site identifier.
	SiteTag string `json:"site_tag"`
	// The Web Analytics site token.
	SiteToken string `json:"site_token"`
	// Encoded JavaScript snippet.
	Snippet string                              `json:"snippet"`
	JSON    rumSiteInfoUpdateResponseResultJSON `json:"-"`
}

// rumSiteInfoUpdateResponseResultJSON contains the JSON metadata for the struct
// [RumSiteInfoUpdateResponseResult]
type rumSiteInfoUpdateResponseResultJSON struct {
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

func (r *RumSiteInfoUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoUpdateResponseResultRule struct {
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
	Paths    []string                                `json:"paths"`
	Priority float64                                 `json:"priority"`
	JSON     rumSiteInfoUpdateResponseResultRuleJSON `json:"-"`
}

// rumSiteInfoUpdateResponseResultRuleJSON contains the JSON metadata for the
// struct [RumSiteInfoUpdateResponseResultRule]
type rumSiteInfoUpdateResponseResultRuleJSON struct {
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

func (r *RumSiteInfoUpdateResponseResultRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoUpdateResponseResultRuleset struct {
	// The Web Analytics ruleset identifier.
	ID string `json:"id"`
	// Whether the ruleset is enabled.
	Enabled  bool   `json:"enabled"`
	ZoneName string `json:"zone_name"`
	// The zone identifier.
	ZoneTag string                                     `json:"zone_tag"`
	JSON    rumSiteInfoUpdateResponseResultRulesetJSON `json:"-"`
}

// rumSiteInfoUpdateResponseResultRulesetJSON contains the JSON metadata for the
// struct [RumSiteInfoUpdateResponseResultRuleset]
type rumSiteInfoUpdateResponseResultRulesetJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	ZoneName    apijson.Field
	ZoneTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoUpdateResponseResultRuleset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoListResponse struct {
	Errors     []RumSiteInfoListResponseError    `json:"errors"`
	Messages   []RumSiteInfoListResponseMessage  `json:"messages"`
	Result     []RumSiteInfoListResponseResult   `json:"result"`
	ResultInfo RumSiteInfoListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful.
	Success bool                        `json:"success"`
	JSON    rumSiteInfoListResponseJSON `json:"-"`
}

// rumSiteInfoListResponseJSON contains the JSON metadata for the struct
// [RumSiteInfoListResponse]
type rumSiteInfoListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoListResponseError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    rumSiteInfoListResponseErrorJSON `json:"-"`
}

// rumSiteInfoListResponseErrorJSON contains the JSON metadata for the struct
// [RumSiteInfoListResponseError]
type rumSiteInfoListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoListResponseMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    rumSiteInfoListResponseMessageJSON `json:"-"`
}

// rumSiteInfoListResponseMessageJSON contains the JSON metadata for the struct
// [RumSiteInfoListResponseMessage]
type rumSiteInfoListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoListResponseResult struct {
	// If enabled, the JavaScript snippet is automatically injected for orange-clouded
	// sites.
	AutoInstall bool      `json:"auto_install"`
	Created     time.Time `json:"created" format:"date-time"`
	// A list of rules.
	Rules   []RumSiteInfoListResponseResultRule  `json:"rules"`
	Ruleset RumSiteInfoListResponseResultRuleset `json:"ruleset"`
	// The Web Analytics site identifier.
	SiteTag string `json:"site_tag"`
	// The Web Analytics site token.
	SiteToken string `json:"site_token"`
	// Encoded JavaScript snippet.
	Snippet string                            `json:"snippet"`
	JSON    rumSiteInfoListResponseResultJSON `json:"-"`
}

// rumSiteInfoListResponseResultJSON contains the JSON metadata for the struct
// [RumSiteInfoListResponseResult]
type rumSiteInfoListResponseResultJSON struct {
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

func (r *RumSiteInfoListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoListResponseResultRule struct {
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
	Paths    []string                              `json:"paths"`
	Priority float64                               `json:"priority"`
	JSON     rumSiteInfoListResponseResultRuleJSON `json:"-"`
}

// rumSiteInfoListResponseResultRuleJSON contains the JSON metadata for the struct
// [RumSiteInfoListResponseResultRule]
type rumSiteInfoListResponseResultRuleJSON struct {
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

func (r *RumSiteInfoListResponseResultRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoListResponseResultRuleset struct {
	// The Web Analytics ruleset identifier.
	ID string `json:"id"`
	// Whether the ruleset is enabled.
	Enabled  bool   `json:"enabled"`
	ZoneName string `json:"zone_name"`
	// The zone identifier.
	ZoneTag string                                   `json:"zone_tag"`
	JSON    rumSiteInfoListResponseResultRulesetJSON `json:"-"`
}

// rumSiteInfoListResponseResultRulesetJSON contains the JSON metadata for the
// struct [RumSiteInfoListResponseResultRuleset]
type rumSiteInfoListResponseResultRulesetJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	ZoneName    apijson.Field
	ZoneTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoListResponseResultRuleset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoListResponseResultInfo struct {
	// The total number of items on the current page.
	Count int64 `json:"count"`
	// Current page within the paginated list of results.
	Page int64 `json:"page"`
	// The maximum number of items to return per page of results.
	PerPage int64 `json:"per_page"`
	// The total number of items.
	TotalCount int64 `json:"total_count"`
	// The total number of pages.
	TotalPages int64                                 `json:"total_pages,nullable"`
	JSON       rumSiteInfoListResponseResultInfoJSON `json:"-"`
}

// rumSiteInfoListResponseResultInfoJSON contains the JSON metadata for the struct
// [RumSiteInfoListResponseResultInfo]
type rumSiteInfoListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoDeleteResponse struct {
	Errors   []RumSiteInfoDeleteResponseError   `json:"errors"`
	Messages []RumSiteInfoDeleteResponseMessage `json:"messages"`
	Result   RumSiteInfoDeleteResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success bool                          `json:"success"`
	JSON    rumSiteInfoDeleteResponseJSON `json:"-"`
}

// rumSiteInfoDeleteResponseJSON contains the JSON metadata for the struct
// [RumSiteInfoDeleteResponse]
type rumSiteInfoDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoDeleteResponseError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    rumSiteInfoDeleteResponseErrorJSON `json:"-"`
}

// rumSiteInfoDeleteResponseErrorJSON contains the JSON metadata for the struct
// [RumSiteInfoDeleteResponseError]
type rumSiteInfoDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoDeleteResponseMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    rumSiteInfoDeleteResponseMessageJSON `json:"-"`
}

// rumSiteInfoDeleteResponseMessageJSON contains the JSON metadata for the struct
// [RumSiteInfoDeleteResponseMessage]
type rumSiteInfoDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoDeleteResponseResult struct {
	// The Web Analytics site identifier.
	SiteTag string                              `json:"site_tag"`
	JSON    rumSiteInfoDeleteResponseResultJSON `json:"-"`
}

// rumSiteInfoDeleteResponseResultJSON contains the JSON metadata for the struct
// [RumSiteInfoDeleteResponseResult]
type rumSiteInfoDeleteResponseResultJSON struct {
	SiteTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
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
