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
	var env RumSiteInfoNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rum/site_info", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves a Web Analytics site.
func (r *RumSiteInfoService) Get(ctx context.Context, accountIdentifier string, siteIdentifier string, opts ...option.RequestOption) (res *RumSiteInfoGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RumSiteInfoGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rum/site_info/%s", accountIdentifier, siteIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing Web Analytics site.
func (r *RumSiteInfoService) Update(ctx context.Context, accountIdentifier string, siteIdentifier string, body RumSiteInfoUpdateParams, opts ...option.RequestOption) (res *RumSiteInfoUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RumSiteInfoUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rum/site_info/%s", accountIdentifier, siteIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all Web Analytics sites of an account.
func (r *RumSiteInfoService) List(ctx context.Context, accountIdentifier string, query RumSiteInfoListParams, opts ...option.RequestOption) (res *[]RumSiteInfoListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RumSiteInfoListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rum/site_info/list", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an existing Web Analytics site.
func (r *RumSiteInfoService) Delete(ctx context.Context, accountIdentifier string, siteIdentifier string, opts ...option.RequestOption) (res *RumSiteInfoDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RumSiteInfoDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rum/site_info/%s", accountIdentifier, siteIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
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
	Errors   []RumSiteInfoNewResponseEnvelopeErrors   `json:"errors"`
	Messages []RumSiteInfoNewResponseEnvelopeMessages `json:"messages"`
	Result   RumSiteInfoNewResponse                   `json:"result"`
	// Whether the API call was successful.
	Success bool                               `json:"success"`
	JSON    rumSiteInfoNewResponseEnvelopeJSON `json:"-"`
}

// rumSiteInfoNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RumSiteInfoNewResponseEnvelope]
type rumSiteInfoNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoNewResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    rumSiteInfoNewResponseEnvelopeErrorsJSON `json:"-"`
}

// rumSiteInfoNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RumSiteInfoNewResponseEnvelopeErrors]
type rumSiteInfoNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoNewResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    rumSiteInfoNewResponseEnvelopeMessagesJSON `json:"-"`
}

// rumSiteInfoNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RumSiteInfoNewResponseEnvelopeMessages]
type rumSiteInfoNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoGetResponseEnvelope struct {
	Errors   []RumSiteInfoGetResponseEnvelopeErrors   `json:"errors"`
	Messages []RumSiteInfoGetResponseEnvelopeMessages `json:"messages"`
	Result   RumSiteInfoGetResponse                   `json:"result"`
	// Whether the API call was successful.
	Success bool                               `json:"success"`
	JSON    rumSiteInfoGetResponseEnvelopeJSON `json:"-"`
}

// rumSiteInfoGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RumSiteInfoGetResponseEnvelope]
type rumSiteInfoGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoGetResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    rumSiteInfoGetResponseEnvelopeErrorsJSON `json:"-"`
}

// rumSiteInfoGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RumSiteInfoGetResponseEnvelopeErrors]
type rumSiteInfoGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoGetResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    rumSiteInfoGetResponseEnvelopeMessagesJSON `json:"-"`
}

// rumSiteInfoGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RumSiteInfoGetResponseEnvelopeMessages]
type rumSiteInfoGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
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
	Errors   []RumSiteInfoUpdateResponseEnvelopeErrors   `json:"errors"`
	Messages []RumSiteInfoUpdateResponseEnvelopeMessages `json:"messages"`
	Result   RumSiteInfoUpdateResponse                   `json:"result"`
	// Whether the API call was successful.
	Success bool                                  `json:"success"`
	JSON    rumSiteInfoUpdateResponseEnvelopeJSON `json:"-"`
}

// rumSiteInfoUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [RumSiteInfoUpdateResponseEnvelope]
type rumSiteInfoUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoUpdateResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    rumSiteInfoUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// rumSiteInfoUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RumSiteInfoUpdateResponseEnvelopeErrors]
type rumSiteInfoUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoUpdateResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    rumSiteInfoUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// rumSiteInfoUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RumSiteInfoUpdateResponseEnvelopeMessages]
type rumSiteInfoUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
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

type RumSiteInfoListResponseEnvelope struct {
	Errors     []RumSiteInfoListResponseEnvelopeErrors   `json:"errors"`
	Messages   []RumSiteInfoListResponseEnvelopeMessages `json:"messages"`
	Result     []RumSiteInfoListResponse                 `json:"result"`
	ResultInfo RumSiteInfoListResponseEnvelopeResultInfo `json:"result_info"`
	// Whether the API call was successful.
	Success bool                                `json:"success"`
	JSON    rumSiteInfoListResponseEnvelopeJSON `json:"-"`
}

// rumSiteInfoListResponseEnvelopeJSON contains the JSON metadata for the struct
// [RumSiteInfoListResponseEnvelope]
type rumSiteInfoListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoListResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    rumSiteInfoListResponseEnvelopeErrorsJSON `json:"-"`
}

// rumSiteInfoListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RumSiteInfoListResponseEnvelopeErrors]
type rumSiteInfoListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoListResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    rumSiteInfoListResponseEnvelopeMessagesJSON `json:"-"`
}

// rumSiteInfoListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RumSiteInfoListResponseEnvelopeMessages]
type rumSiteInfoListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoListResponseEnvelopeResultInfo struct {
	// The total number of items on the current page.
	Count int64 `json:"count"`
	// Current page within the paginated list of results.
	Page int64 `json:"page"`
	// The maximum number of items to return per page of results.
	PerPage int64 `json:"per_page"`
	// The total number of items.
	TotalCount int64 `json:"total_count"`
	// The total number of pages.
	TotalPages int64                                         `json:"total_pages,nullable"`
	JSON       rumSiteInfoListResponseEnvelopeResultInfoJSON `json:"-"`
}

// rumSiteInfoListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [RumSiteInfoListResponseEnvelopeResultInfo]
type rumSiteInfoListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoDeleteResponseEnvelope struct {
	Errors   []RumSiteInfoDeleteResponseEnvelopeErrors   `json:"errors"`
	Messages []RumSiteInfoDeleteResponseEnvelopeMessages `json:"messages"`
	Result   RumSiteInfoDeleteResponse                   `json:"result"`
	// Whether the API call was successful.
	Success bool                                  `json:"success"`
	JSON    rumSiteInfoDeleteResponseEnvelopeJSON `json:"-"`
}

// rumSiteInfoDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [RumSiteInfoDeleteResponseEnvelope]
type rumSiteInfoDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoDeleteResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    rumSiteInfoDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// rumSiteInfoDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RumSiteInfoDeleteResponseEnvelopeErrors]
type rumSiteInfoDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumSiteInfoDeleteResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    rumSiteInfoDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// rumSiteInfoDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RumSiteInfoDeleteResponseEnvelopeMessages]
type rumSiteInfoDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumSiteInfoDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
