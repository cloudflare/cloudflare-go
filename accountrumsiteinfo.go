// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountRumSiteInfoService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountRumSiteInfoService] method
// instead.
type AccountRumSiteInfoService struct {
	Options []option.RequestOption
}

// NewAccountRumSiteInfoService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountRumSiteInfoService(opts ...option.RequestOption) (r *AccountRumSiteInfoService) {
	r = &AccountRumSiteInfoService{}
	r.Options = opts
	return
}

// Creates a new Web Analytics site.
func (r *AccountRumSiteInfoService) New(ctx context.Context, accountIdentifier string, body AccountRumSiteInfoNewParams, opts ...option.RequestOption) (res *AccountRumSiteInfoNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rum/site_info", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountRumSiteInfoNewResponse struct {
	Errors   []AccountRumSiteInfoNewResponseError   `json:"errors"`
	Messages []AccountRumSiteInfoNewResponseMessage `json:"messages"`
	Result   AccountRumSiteInfoNewResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success bool                              `json:"success"`
	JSON    accountRumSiteInfoNewResponseJSON `json:"-"`
}

// accountRumSiteInfoNewResponseJSON contains the JSON metadata for the struct
// [AccountRumSiteInfoNewResponse]
type accountRumSiteInfoNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRumSiteInfoNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRumSiteInfoNewResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    accountRumSiteInfoNewResponseErrorJSON `json:"-"`
}

// accountRumSiteInfoNewResponseErrorJSON contains the JSON metadata for the struct
// [AccountRumSiteInfoNewResponseError]
type accountRumSiteInfoNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRumSiteInfoNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRumSiteInfoNewResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accountRumSiteInfoNewResponseMessageJSON `json:"-"`
}

// accountRumSiteInfoNewResponseMessageJSON contains the JSON metadata for the
// struct [AccountRumSiteInfoNewResponseMessage]
type accountRumSiteInfoNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRumSiteInfoNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRumSiteInfoNewResponseResult struct {
	// If enabled, the JavaScript snippet is automatically injected for orange-clouded
	// sites.
	AutoInstall bool      `json:"auto_install"`
	Created     time.Time `json:"created" format:"date-time"`
	// A list of rules.
	Rules   []AccountRumSiteInfoNewResponseResultRule  `json:"rules"`
	Ruleset AccountRumSiteInfoNewResponseResultRuleset `json:"ruleset"`
	// The Web Analytics site identifier.
	SiteTag string `json:"site_tag"`
	// The Web Analytics site token.
	SiteToken string `json:"site_token"`
	// Encoded JavaScript snippet.
	Snippet string                                  `json:"snippet"`
	JSON    accountRumSiteInfoNewResponseResultJSON `json:"-"`
}

// accountRumSiteInfoNewResponseResultJSON contains the JSON metadata for the
// struct [AccountRumSiteInfoNewResponseResult]
type accountRumSiteInfoNewResponseResultJSON struct {
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

func (r *AccountRumSiteInfoNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRumSiteInfoNewResponseResultRule struct {
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
	Paths    []string                                    `json:"paths"`
	Priority float64                                     `json:"priority"`
	JSON     accountRumSiteInfoNewResponseResultRuleJSON `json:"-"`
}

// accountRumSiteInfoNewResponseResultRuleJSON contains the JSON metadata for the
// struct [AccountRumSiteInfoNewResponseResultRule]
type accountRumSiteInfoNewResponseResultRuleJSON struct {
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

func (r *AccountRumSiteInfoNewResponseResultRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRumSiteInfoNewResponseResultRuleset struct {
	// The Web Analytics ruleset identifier.
	ID string `json:"id"`
	// Whether the ruleset is enabled.
	Enabled  bool   `json:"enabled"`
	ZoneName string `json:"zone_name"`
	// The zone identifier.
	ZoneTag string                                         `json:"zone_tag"`
	JSON    accountRumSiteInfoNewResponseResultRulesetJSON `json:"-"`
}

// accountRumSiteInfoNewResponseResultRulesetJSON contains the JSON metadata for
// the struct [AccountRumSiteInfoNewResponseResultRuleset]
type accountRumSiteInfoNewResponseResultRulesetJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	ZoneName    apijson.Field
	ZoneTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRumSiteInfoNewResponseResultRuleset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRumSiteInfoNewParams struct {
	// If enabled, the JavaScript snippet is automatically injected for orange-clouded
	// sites.
	AutoInstall param.Field[bool] `json:"auto_install"`
	// The hostname to use for gray-clouded sites.
	Host param.Field[string] `json:"host"`
	// The zone identifier.
	ZoneTag param.Field[string] `json:"zone_tag"`
}

func (r AccountRumSiteInfoNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
