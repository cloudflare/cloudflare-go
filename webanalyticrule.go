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

// WebAnalyticRuleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewWebAnalyticRuleService] method
// instead.
type WebAnalyticRuleService struct {
	Options []option.RequestOption
}

// NewWebAnalyticRuleService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWebAnalyticRuleService(opts ...option.RequestOption) (r *WebAnalyticRuleService) {
	r = &WebAnalyticRuleService{}
	r.Options = opts
	return
}

// Creates a new rule in a Web Analytics ruleset.
func (r *WebAnalyticRuleService) New(ctx context.Context, accountIdentifier string, rulesetIdentifier string, body WebAnalyticRuleNewParams, opts ...option.RequestOption) (res *Rule, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rum/v2/%s/rule", accountIdentifier, rulesetIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Modifies one or more rules in a Web Analytics ruleset with a single request.
func (r *WebAnalyticRuleService) Update(ctx context.Context, accountIdentifier string, rulesetIdentifier string, body WebAnalyticRuleUpdateParams, opts ...option.RequestOption) (res *WebAnalyticRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rum/v2/%s/rules", accountIdentifier, rulesetIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists all the rules in a Web Analytics ruleset.
func (r *WebAnalyticRuleService) List(ctx context.Context, accountIdentifier string, rulesetIdentifier string, opts ...option.RequestOption) (res *WebAnalyticRuleListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rum/v2/%s/rules", accountIdentifier, rulesetIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes an existing rule from a Web Analytics ruleset.
func (r *WebAnalyticRuleService) Delete(ctx context.Context, accountIdentifier string, rulesetIdentifier string, ruleIdentifier string, opts ...option.RequestOption) (res *WebAnalyticRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rum/v2/%s/rule/%s", accountIdentifier, rulesetIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type WebAnalyticRuleUpdateResponse struct {
	Errors   []WebAnalyticRuleUpdateResponseError   `json:"errors"`
	Messages []WebAnalyticRuleUpdateResponseMessage `json:"messages"`
	Result   WebAnalyticRuleUpdateResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success bool                              `json:"success"`
	JSON    webAnalyticRuleUpdateResponseJSON `json:"-"`
}

// webAnalyticRuleUpdateResponseJSON contains the JSON metadata for the struct
// [WebAnalyticRuleUpdateResponse]
type webAnalyticRuleUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebAnalyticRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticRuleUpdateResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    webAnalyticRuleUpdateResponseErrorJSON `json:"-"`
}

// webAnalyticRuleUpdateResponseErrorJSON contains the JSON metadata for the struct
// [WebAnalyticRuleUpdateResponseError]
type webAnalyticRuleUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebAnalyticRuleUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticRuleUpdateResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    webAnalyticRuleUpdateResponseMessageJSON `json:"-"`
}

// webAnalyticRuleUpdateResponseMessageJSON contains the JSON metadata for the
// struct [WebAnalyticRuleUpdateResponseMessage]
type webAnalyticRuleUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebAnalyticRuleUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticRuleUpdateResponseResult struct {
	// A list of rules.
	Rules   []WebAnalyticRuleUpdateResponseResultRule  `json:"rules"`
	Ruleset WebAnalyticRuleUpdateResponseResultRuleset `json:"ruleset"`
	JSON    webAnalyticRuleUpdateResponseResultJSON    `json:"-"`
}

// webAnalyticRuleUpdateResponseResultJSON contains the JSON metadata for the
// struct [WebAnalyticRuleUpdateResponseResult]
type webAnalyticRuleUpdateResponseResultJSON struct {
	Rules       apijson.Field
	Ruleset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebAnalyticRuleUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticRuleUpdateResponseResultRule struct {
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
	JSON     webAnalyticRuleUpdateResponseResultRuleJSON `json:"-"`
}

// webAnalyticRuleUpdateResponseResultRuleJSON contains the JSON metadata for the
// struct [WebAnalyticRuleUpdateResponseResultRule]
type webAnalyticRuleUpdateResponseResultRuleJSON struct {
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

func (r *WebAnalyticRuleUpdateResponseResultRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticRuleUpdateResponseResultRuleset struct {
	// The Web Analytics ruleset identifier.
	ID string `json:"id"`
	// Whether the ruleset is enabled.
	Enabled  bool   `json:"enabled"`
	ZoneName string `json:"zone_name"`
	// The zone identifier.
	ZoneTag string                                         `json:"zone_tag"`
	JSON    webAnalyticRuleUpdateResponseResultRulesetJSON `json:"-"`
}

// webAnalyticRuleUpdateResponseResultRulesetJSON contains the JSON metadata for
// the struct [WebAnalyticRuleUpdateResponseResultRuleset]
type webAnalyticRuleUpdateResponseResultRulesetJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	ZoneName    apijson.Field
	ZoneTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebAnalyticRuleUpdateResponseResultRuleset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticRuleListResponse struct {
	Errors   []WebAnalyticRuleListResponseError   `json:"errors"`
	Messages []WebAnalyticRuleListResponseMessage `json:"messages"`
	Result   WebAnalyticRuleListResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success bool                            `json:"success"`
	JSON    webAnalyticRuleListResponseJSON `json:"-"`
}

// webAnalyticRuleListResponseJSON contains the JSON metadata for the struct
// [WebAnalyticRuleListResponse]
type webAnalyticRuleListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebAnalyticRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticRuleListResponseError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    webAnalyticRuleListResponseErrorJSON `json:"-"`
}

// webAnalyticRuleListResponseErrorJSON contains the JSON metadata for the struct
// [WebAnalyticRuleListResponseError]
type webAnalyticRuleListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebAnalyticRuleListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticRuleListResponseMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    webAnalyticRuleListResponseMessageJSON `json:"-"`
}

// webAnalyticRuleListResponseMessageJSON contains the JSON metadata for the struct
// [WebAnalyticRuleListResponseMessage]
type webAnalyticRuleListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebAnalyticRuleListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticRuleListResponseResult struct {
	// A list of rules.
	Rules   []WebAnalyticRuleListResponseResultRule  `json:"rules"`
	Ruleset WebAnalyticRuleListResponseResultRuleset `json:"ruleset"`
	JSON    webAnalyticRuleListResponseResultJSON    `json:"-"`
}

// webAnalyticRuleListResponseResultJSON contains the JSON metadata for the struct
// [WebAnalyticRuleListResponseResult]
type webAnalyticRuleListResponseResultJSON struct {
	Rules       apijson.Field
	Ruleset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebAnalyticRuleListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticRuleListResponseResultRule struct {
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
	Paths    []string                                  `json:"paths"`
	Priority float64                                   `json:"priority"`
	JSON     webAnalyticRuleListResponseResultRuleJSON `json:"-"`
}

// webAnalyticRuleListResponseResultRuleJSON contains the JSON metadata for the
// struct [WebAnalyticRuleListResponseResultRule]
type webAnalyticRuleListResponseResultRuleJSON struct {
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

func (r *WebAnalyticRuleListResponseResultRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticRuleListResponseResultRuleset struct {
	// The Web Analytics ruleset identifier.
	ID string `json:"id"`
	// Whether the ruleset is enabled.
	Enabled  bool   `json:"enabled"`
	ZoneName string `json:"zone_name"`
	// The zone identifier.
	ZoneTag string                                       `json:"zone_tag"`
	JSON    webAnalyticRuleListResponseResultRulesetJSON `json:"-"`
}

// webAnalyticRuleListResponseResultRulesetJSON contains the JSON metadata for the
// struct [WebAnalyticRuleListResponseResultRuleset]
type webAnalyticRuleListResponseResultRulesetJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	ZoneName    apijson.Field
	ZoneTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebAnalyticRuleListResponseResultRuleset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticRuleDeleteResponse struct {
	Errors   []WebAnalyticRuleDeleteResponseError   `json:"errors"`
	Messages []WebAnalyticRuleDeleteResponseMessage `json:"messages"`
	Result   WebAnalyticRuleDeleteResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success bool                              `json:"success"`
	JSON    webAnalyticRuleDeleteResponseJSON `json:"-"`
}

// webAnalyticRuleDeleteResponseJSON contains the JSON metadata for the struct
// [WebAnalyticRuleDeleteResponse]
type webAnalyticRuleDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebAnalyticRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticRuleDeleteResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    webAnalyticRuleDeleteResponseErrorJSON `json:"-"`
}

// webAnalyticRuleDeleteResponseErrorJSON contains the JSON metadata for the struct
// [WebAnalyticRuleDeleteResponseError]
type webAnalyticRuleDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebAnalyticRuleDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticRuleDeleteResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    webAnalyticRuleDeleteResponseMessageJSON `json:"-"`
}

// webAnalyticRuleDeleteResponseMessageJSON contains the JSON metadata for the
// struct [WebAnalyticRuleDeleteResponseMessage]
type webAnalyticRuleDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebAnalyticRuleDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticRuleDeleteResponseResult struct {
	// The Web Analytics rule identifier.
	ID   string                                  `json:"id"`
	JSON webAnalyticRuleDeleteResponseResultJSON `json:"-"`
}

// webAnalyticRuleDeleteResponseResultJSON contains the JSON metadata for the
// struct [WebAnalyticRuleDeleteResponseResult]
type webAnalyticRuleDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebAnalyticRuleDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticRuleNewParams struct {
	Host param.Field[string] `json:"host"`
	// Whether the rule includes or excludes traffic from being measured.
	Inclusive param.Field[bool] `json:"inclusive"`
	// Whether the rule is paused or not.
	IsPaused param.Field[bool]     `json:"is_paused"`
	Paths    param.Field[[]string] `json:"paths"`
}

func (r WebAnalyticRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WebAnalyticRuleUpdateParams struct {
	// A list of rule identifiers to delete.
	DeleteRules param.Field[[]string] `json:"delete_rules"`
	// A list of rules to create or update.
	Rules param.Field[[]WebAnalyticRuleUpdateParamsRule] `json:"rules"`
}

func (r WebAnalyticRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WebAnalyticRuleUpdateParamsRule struct {
	// The Web Analytics rule identifier.
	ID        param.Field[string]   `json:"id"`
	Host      param.Field[string]   `json:"host"`
	Inclusive param.Field[bool]     `json:"inclusive"`
	IsPaused  param.Field[bool]     `json:"is_paused"`
	Paths     param.Field[[]string] `json:"paths"`
}

func (r WebAnalyticRuleUpdateParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
