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

// WebAnalyticSiteInfoService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewWebAnalyticSiteInfoService]
// method instead.
type WebAnalyticSiteInfoService struct {
	Options []option.RequestOption
}

// NewWebAnalyticSiteInfoService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWebAnalyticSiteInfoService(opts ...option.RequestOption) (r *WebAnalyticSiteInfoService) {
	r = &WebAnalyticSiteInfoService{}
	r.Options = opts
	return
}

// Retrieves a Web Analytics site.
func (r *WebAnalyticSiteInfoService) Get(ctx context.Context, accountIdentifier string, siteIdentifier string, opts ...option.RequestOption) (res *WebAnalyticSiteInfoGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rum/site_info/%s", accountIdentifier, siteIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an existing Web Analytics site.
func (r *WebAnalyticSiteInfoService) Update(ctx context.Context, accountIdentifier string, siteIdentifier string, body WebAnalyticSiteInfoUpdateParams, opts ...option.RequestOption) (res *WebAnalyticSiteInfoUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rum/site_info/%s", accountIdentifier, siteIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists all Web Analytics sites of an account.
func (r *WebAnalyticSiteInfoService) List(ctx context.Context, accountIdentifier string, query WebAnalyticSiteInfoListParams, opts ...option.RequestOption) (res *shared.Page[WebAnalyticSiteInfoListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/rum/site_info/list", accountIdentifier)
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

// Deletes an existing Web Analytics site.
func (r *WebAnalyticSiteInfoService) Delete(ctx context.Context, accountIdentifier string, siteIdentifier string, opts ...option.RequestOption) (res *WebAnalyticSiteInfoDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rum/site_info/%s", accountIdentifier, siteIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type WebAnalyticSiteInfoGetResponse struct {
	Errors   []WebAnalyticSiteInfoGetResponseError   `json:"errors"`
	Messages []WebAnalyticSiteInfoGetResponseMessage `json:"messages"`
	Result   WebAnalyticSiteInfoGetResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success bool                               `json:"success"`
	JSON    webAnalyticSiteInfoGetResponseJSON `json:"-"`
}

// webAnalyticSiteInfoGetResponseJSON contains the JSON metadata for the struct
// [WebAnalyticSiteInfoGetResponse]
type webAnalyticSiteInfoGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebAnalyticSiteInfoGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticSiteInfoGetResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    webAnalyticSiteInfoGetResponseErrorJSON `json:"-"`
}

// webAnalyticSiteInfoGetResponseErrorJSON contains the JSON metadata for the
// struct [WebAnalyticSiteInfoGetResponseError]
type webAnalyticSiteInfoGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebAnalyticSiteInfoGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticSiteInfoGetResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    webAnalyticSiteInfoGetResponseMessageJSON `json:"-"`
}

// webAnalyticSiteInfoGetResponseMessageJSON contains the JSON metadata for the
// struct [WebAnalyticSiteInfoGetResponseMessage]
type webAnalyticSiteInfoGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebAnalyticSiteInfoGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticSiteInfoGetResponseResult struct {
	// If enabled, the JavaScript snippet is automatically injected for orange-clouded
	// sites.
	AutoInstall bool      `json:"auto_install"`
	Created     time.Time `json:"created" format:"date-time"`
	// A list of rules.
	Rules   []WebAnalyticSiteInfoGetResponseResultRule  `json:"rules"`
	Ruleset WebAnalyticSiteInfoGetResponseResultRuleset `json:"ruleset"`
	// The Web Analytics site identifier.
	SiteTag string `json:"site_tag"`
	// The Web Analytics site token.
	SiteToken string `json:"site_token"`
	// Encoded JavaScript snippet.
	Snippet string                                   `json:"snippet"`
	JSON    webAnalyticSiteInfoGetResponseResultJSON `json:"-"`
}

// webAnalyticSiteInfoGetResponseResultJSON contains the JSON metadata for the
// struct [WebAnalyticSiteInfoGetResponseResult]
type webAnalyticSiteInfoGetResponseResultJSON struct {
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

func (r *WebAnalyticSiteInfoGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticSiteInfoGetResponseResultRule struct {
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
	Paths    []string                                     `json:"paths"`
	Priority float64                                      `json:"priority"`
	JSON     webAnalyticSiteInfoGetResponseResultRuleJSON `json:"-"`
}

// webAnalyticSiteInfoGetResponseResultRuleJSON contains the JSON metadata for the
// struct [WebAnalyticSiteInfoGetResponseResultRule]
type webAnalyticSiteInfoGetResponseResultRuleJSON struct {
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

func (r *WebAnalyticSiteInfoGetResponseResultRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticSiteInfoGetResponseResultRuleset struct {
	// The Web Analytics ruleset identifier.
	ID string `json:"id"`
	// Whether the ruleset is enabled.
	Enabled  bool   `json:"enabled"`
	ZoneName string `json:"zone_name"`
	// The zone identifier.
	ZoneTag string                                          `json:"zone_tag"`
	JSON    webAnalyticSiteInfoGetResponseResultRulesetJSON `json:"-"`
}

// webAnalyticSiteInfoGetResponseResultRulesetJSON contains the JSON metadata for
// the struct [WebAnalyticSiteInfoGetResponseResultRuleset]
type webAnalyticSiteInfoGetResponseResultRulesetJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	ZoneName    apijson.Field
	ZoneTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebAnalyticSiteInfoGetResponseResultRuleset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticSiteInfoUpdateResponse struct {
	Errors   []WebAnalyticSiteInfoUpdateResponseError   `json:"errors"`
	Messages []WebAnalyticSiteInfoUpdateResponseMessage `json:"messages"`
	Result   WebAnalyticSiteInfoUpdateResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success bool                                  `json:"success"`
	JSON    webAnalyticSiteInfoUpdateResponseJSON `json:"-"`
}

// webAnalyticSiteInfoUpdateResponseJSON contains the JSON metadata for the struct
// [WebAnalyticSiteInfoUpdateResponse]
type webAnalyticSiteInfoUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebAnalyticSiteInfoUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticSiteInfoUpdateResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    webAnalyticSiteInfoUpdateResponseErrorJSON `json:"-"`
}

// webAnalyticSiteInfoUpdateResponseErrorJSON contains the JSON metadata for the
// struct [WebAnalyticSiteInfoUpdateResponseError]
type webAnalyticSiteInfoUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebAnalyticSiteInfoUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticSiteInfoUpdateResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    webAnalyticSiteInfoUpdateResponseMessageJSON `json:"-"`
}

// webAnalyticSiteInfoUpdateResponseMessageJSON contains the JSON metadata for the
// struct [WebAnalyticSiteInfoUpdateResponseMessage]
type webAnalyticSiteInfoUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebAnalyticSiteInfoUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticSiteInfoUpdateResponseResult struct {
	// If enabled, the JavaScript snippet is automatically injected for orange-clouded
	// sites.
	AutoInstall bool      `json:"auto_install"`
	Created     time.Time `json:"created" format:"date-time"`
	// A list of rules.
	Rules   []WebAnalyticSiteInfoUpdateResponseResultRule  `json:"rules"`
	Ruleset WebAnalyticSiteInfoUpdateResponseResultRuleset `json:"ruleset"`
	// The Web Analytics site identifier.
	SiteTag string `json:"site_tag"`
	// The Web Analytics site token.
	SiteToken string `json:"site_token"`
	// Encoded JavaScript snippet.
	Snippet string                                      `json:"snippet"`
	JSON    webAnalyticSiteInfoUpdateResponseResultJSON `json:"-"`
}

// webAnalyticSiteInfoUpdateResponseResultJSON contains the JSON metadata for the
// struct [WebAnalyticSiteInfoUpdateResponseResult]
type webAnalyticSiteInfoUpdateResponseResultJSON struct {
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

func (r *WebAnalyticSiteInfoUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticSiteInfoUpdateResponseResultRule struct {
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
	Paths    []string                                        `json:"paths"`
	Priority float64                                         `json:"priority"`
	JSON     webAnalyticSiteInfoUpdateResponseResultRuleJSON `json:"-"`
}

// webAnalyticSiteInfoUpdateResponseResultRuleJSON contains the JSON metadata for
// the struct [WebAnalyticSiteInfoUpdateResponseResultRule]
type webAnalyticSiteInfoUpdateResponseResultRuleJSON struct {
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

func (r *WebAnalyticSiteInfoUpdateResponseResultRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticSiteInfoUpdateResponseResultRuleset struct {
	// The Web Analytics ruleset identifier.
	ID string `json:"id"`
	// Whether the ruleset is enabled.
	Enabled  bool   `json:"enabled"`
	ZoneName string `json:"zone_name"`
	// The zone identifier.
	ZoneTag string                                             `json:"zone_tag"`
	JSON    webAnalyticSiteInfoUpdateResponseResultRulesetJSON `json:"-"`
}

// webAnalyticSiteInfoUpdateResponseResultRulesetJSON contains the JSON metadata
// for the struct [WebAnalyticSiteInfoUpdateResponseResultRuleset]
type webAnalyticSiteInfoUpdateResponseResultRulesetJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	ZoneName    apijson.Field
	ZoneTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebAnalyticSiteInfoUpdateResponseResultRuleset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticSiteInfoListResponse struct {
	// If enabled, the JavaScript snippet is automatically injected for orange-clouded
	// sites.
	AutoInstall bool      `json:"auto_install"`
	Created     time.Time `json:"created" format:"date-time"`
	// A list of rules.
	Rules   []WebAnalyticSiteInfoListResponseRule  `json:"rules"`
	Ruleset WebAnalyticSiteInfoListResponseRuleset `json:"ruleset"`
	// The Web Analytics site identifier.
	SiteTag string `json:"site_tag"`
	// The Web Analytics site token.
	SiteToken string `json:"site_token"`
	// Encoded JavaScript snippet.
	Snippet string                              `json:"snippet"`
	JSON    webAnalyticSiteInfoListResponseJSON `json:"-"`
}

// webAnalyticSiteInfoListResponseJSON contains the JSON metadata for the struct
// [WebAnalyticSiteInfoListResponse]
type webAnalyticSiteInfoListResponseJSON struct {
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

func (r *WebAnalyticSiteInfoListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticSiteInfoListResponseRule struct {
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
	JSON     webAnalyticSiteInfoListResponseRuleJSON `json:"-"`
}

// webAnalyticSiteInfoListResponseRuleJSON contains the JSON metadata for the
// struct [WebAnalyticSiteInfoListResponseRule]
type webAnalyticSiteInfoListResponseRuleJSON struct {
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

func (r *WebAnalyticSiteInfoListResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticSiteInfoListResponseRuleset struct {
	// The Web Analytics ruleset identifier.
	ID string `json:"id"`
	// Whether the ruleset is enabled.
	Enabled  bool   `json:"enabled"`
	ZoneName string `json:"zone_name"`
	// The zone identifier.
	ZoneTag string                                     `json:"zone_tag"`
	JSON    webAnalyticSiteInfoListResponseRulesetJSON `json:"-"`
}

// webAnalyticSiteInfoListResponseRulesetJSON contains the JSON metadata for the
// struct [WebAnalyticSiteInfoListResponseRuleset]
type webAnalyticSiteInfoListResponseRulesetJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	ZoneName    apijson.Field
	ZoneTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebAnalyticSiteInfoListResponseRuleset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticSiteInfoDeleteResponse struct {
	Errors   []WebAnalyticSiteInfoDeleteResponseError   `json:"errors"`
	Messages []WebAnalyticSiteInfoDeleteResponseMessage `json:"messages"`
	Result   WebAnalyticSiteInfoDeleteResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success bool                                  `json:"success"`
	JSON    webAnalyticSiteInfoDeleteResponseJSON `json:"-"`
}

// webAnalyticSiteInfoDeleteResponseJSON contains the JSON metadata for the struct
// [WebAnalyticSiteInfoDeleteResponse]
type webAnalyticSiteInfoDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebAnalyticSiteInfoDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticSiteInfoDeleteResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    webAnalyticSiteInfoDeleteResponseErrorJSON `json:"-"`
}

// webAnalyticSiteInfoDeleteResponseErrorJSON contains the JSON metadata for the
// struct [WebAnalyticSiteInfoDeleteResponseError]
type webAnalyticSiteInfoDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebAnalyticSiteInfoDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticSiteInfoDeleteResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    webAnalyticSiteInfoDeleteResponseMessageJSON `json:"-"`
}

// webAnalyticSiteInfoDeleteResponseMessageJSON contains the JSON metadata for the
// struct [WebAnalyticSiteInfoDeleteResponseMessage]
type webAnalyticSiteInfoDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebAnalyticSiteInfoDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticSiteInfoDeleteResponseResult struct {
	// The Web Analytics site identifier.
	SiteTag string                                      `json:"site_tag"`
	JSON    webAnalyticSiteInfoDeleteResponseResultJSON `json:"-"`
}

// webAnalyticSiteInfoDeleteResponseResultJSON contains the JSON metadata for the
// struct [WebAnalyticSiteInfoDeleteResponseResult]
type webAnalyticSiteInfoDeleteResponseResultJSON struct {
	SiteTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebAnalyticSiteInfoDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticSiteInfoUpdateParams struct {
	// If enabled, the JavaScript snippet is automatically injected for orange-clouded
	// sites.
	AutoInstall param.Field[bool] `json:"auto_install"`
	// The hostname to use for gray-clouded sites.
	Host param.Field[string] `json:"host"`
	// The zone identifier.
	ZoneTag param.Field[string] `json:"zone_tag"`
}

func (r WebAnalyticSiteInfoUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WebAnalyticSiteInfoListParams struct {
	// The property used to sort the list of results.
	OrderBy param.Field[WebAnalyticSiteInfoListParamsOrderBy] `query:"order_by"`
	// Current page within the paginated list of results.
	Page param.Field[float64] `query:"page"`
	// Number of items to return per page of results.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [WebAnalyticSiteInfoListParams]'s query parameters as
// `url.Values`.
func (r WebAnalyticSiteInfoListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The property used to sort the list of results.
type WebAnalyticSiteInfoListParamsOrderBy string

const (
	WebAnalyticSiteInfoListParamsOrderByHost    WebAnalyticSiteInfoListParamsOrderBy = "host"
	WebAnalyticSiteInfoListParamsOrderByCreated WebAnalyticSiteInfoListParamsOrderBy = "created"
)
