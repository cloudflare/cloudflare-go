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

// RumRuleService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewRumRuleService] method instead.
type RumRuleService struct {
	Options []option.RequestOption
}

// NewRumRuleService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRumRuleService(opts ...option.RequestOption) (r *RumRuleService) {
	r = &RumRuleService{}
	r.Options = opts
	return
}

// Creates a new rule in a Web Analytics ruleset.
func (r *RumRuleService) New(ctx context.Context, accountIdentifier string, rulesetIdentifier string, body RumRuleNewParams, opts ...option.RequestOption) (res *RumRuleNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rum/v2/%s/rule", accountIdentifier, rulesetIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates a rule in a Web Analytics ruleset.
func (r *RumRuleService) Update(ctx context.Context, accountIdentifier string, rulesetIdentifier string, ruleIdentifier string, body RumRuleUpdateParams, opts ...option.RequestOption) (res *RumRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rum/v2/%s/rule/%s", accountIdentifier, rulesetIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists all the rules in a Web Analytics ruleset.
func (r *RumRuleService) List(ctx context.Context, accountIdentifier string, rulesetIdentifier string, opts ...option.RequestOption) (res *RumRuleListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rum/v2/%s/rules", accountIdentifier, rulesetIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes an existing rule from a Web Analytics ruleset.
func (r *RumRuleService) Delete(ctx context.Context, accountIdentifier string, rulesetIdentifier string, ruleIdentifier string, opts ...option.RequestOption) (res *RumRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rum/v2/%s/rule/%s", accountIdentifier, rulesetIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type RumRuleNewResponse struct {
	Errors   []RumRuleNewResponseError   `json:"errors"`
	Messages []RumRuleNewResponseMessage `json:"messages"`
	Result   RumRuleNewResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success bool                   `json:"success"`
	JSON    rumRuleNewResponseJSON `json:"-"`
}

// rumRuleNewResponseJSON contains the JSON metadata for the struct
// [RumRuleNewResponse]
type rumRuleNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumRuleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumRuleNewResponseError struct {
	Code    int64                       `json:"code,required"`
	Message string                      `json:"message,required"`
	JSON    rumRuleNewResponseErrorJSON `json:"-"`
}

// rumRuleNewResponseErrorJSON contains the JSON metadata for the struct
// [RumRuleNewResponseError]
type rumRuleNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumRuleNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumRuleNewResponseMessage struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    rumRuleNewResponseMessageJSON `json:"-"`
}

// rumRuleNewResponseMessageJSON contains the JSON metadata for the struct
// [RumRuleNewResponseMessage]
type rumRuleNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumRuleNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumRuleNewResponseResult struct {
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
	JSON     rumRuleNewResponseResultJSON `json:"-"`
}

// rumRuleNewResponseResultJSON contains the JSON metadata for the struct
// [RumRuleNewResponseResult]
type rumRuleNewResponseResultJSON struct {
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

func (r *RumRuleNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumRuleUpdateResponse struct {
	Errors   []RumRuleUpdateResponseError   `json:"errors"`
	Messages []RumRuleUpdateResponseMessage `json:"messages"`
	Result   RumRuleUpdateResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success bool                      `json:"success"`
	JSON    rumRuleUpdateResponseJSON `json:"-"`
}

// rumRuleUpdateResponseJSON contains the JSON metadata for the struct
// [RumRuleUpdateResponse]
type rumRuleUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumRuleUpdateResponseError struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    rumRuleUpdateResponseErrorJSON `json:"-"`
}

// rumRuleUpdateResponseErrorJSON contains the JSON metadata for the struct
// [RumRuleUpdateResponseError]
type rumRuleUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumRuleUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumRuleUpdateResponseMessage struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    rumRuleUpdateResponseMessageJSON `json:"-"`
}

// rumRuleUpdateResponseMessageJSON contains the JSON metadata for the struct
// [RumRuleUpdateResponseMessage]
type rumRuleUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumRuleUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumRuleUpdateResponseResult struct {
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
	JSON     rumRuleUpdateResponseResultJSON `json:"-"`
}

// rumRuleUpdateResponseResultJSON contains the JSON metadata for the struct
// [RumRuleUpdateResponseResult]
type rumRuleUpdateResponseResultJSON struct {
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

func (r *RumRuleUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumRuleListResponse struct {
	Errors   []RumRuleListResponseError   `json:"errors"`
	Messages []RumRuleListResponseMessage `json:"messages"`
	Result   RumRuleListResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success bool                    `json:"success"`
	JSON    rumRuleListResponseJSON `json:"-"`
}

// rumRuleListResponseJSON contains the JSON metadata for the struct
// [RumRuleListResponse]
type rumRuleListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumRuleListResponseError struct {
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    rumRuleListResponseErrorJSON `json:"-"`
}

// rumRuleListResponseErrorJSON contains the JSON metadata for the struct
// [RumRuleListResponseError]
type rumRuleListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumRuleListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumRuleListResponseMessage struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    rumRuleListResponseMessageJSON `json:"-"`
}

// rumRuleListResponseMessageJSON contains the JSON metadata for the struct
// [RumRuleListResponseMessage]
type rumRuleListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumRuleListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumRuleListResponseResult struct {
	// A list of rules.
	Rules   []RumRuleListResponseResultRule  `json:"rules"`
	Ruleset RumRuleListResponseResultRuleset `json:"ruleset"`
	JSON    rumRuleListResponseResultJSON    `json:"-"`
}

// rumRuleListResponseResultJSON contains the JSON metadata for the struct
// [RumRuleListResponseResult]
type rumRuleListResponseResultJSON struct {
	Rules       apijson.Field
	Ruleset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumRuleListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumRuleListResponseResultRule struct {
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
	JSON     rumRuleListResponseResultRuleJSON `json:"-"`
}

// rumRuleListResponseResultRuleJSON contains the JSON metadata for the struct
// [RumRuleListResponseResultRule]
type rumRuleListResponseResultRuleJSON struct {
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

func (r *RumRuleListResponseResultRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumRuleListResponseResultRuleset struct {
	// The Web Analytics ruleset identifier.
	ID string `json:"id"`
	// Whether the ruleset is enabled.
	Enabled  bool   `json:"enabled"`
	ZoneName string `json:"zone_name"`
	// The zone identifier.
	ZoneTag string                               `json:"zone_tag"`
	JSON    rumRuleListResponseResultRulesetJSON `json:"-"`
}

// rumRuleListResponseResultRulesetJSON contains the JSON metadata for the struct
// [RumRuleListResponseResultRuleset]
type rumRuleListResponseResultRulesetJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	ZoneName    apijson.Field
	ZoneTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumRuleListResponseResultRuleset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumRuleDeleteResponse struct {
	Errors   []RumRuleDeleteResponseError   `json:"errors"`
	Messages []RumRuleDeleteResponseMessage `json:"messages"`
	Result   RumRuleDeleteResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success bool                      `json:"success"`
	JSON    rumRuleDeleteResponseJSON `json:"-"`
}

// rumRuleDeleteResponseJSON contains the JSON metadata for the struct
// [RumRuleDeleteResponse]
type rumRuleDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumRuleDeleteResponseError struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    rumRuleDeleteResponseErrorJSON `json:"-"`
}

// rumRuleDeleteResponseErrorJSON contains the JSON metadata for the struct
// [RumRuleDeleteResponseError]
type rumRuleDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumRuleDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumRuleDeleteResponseMessage struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    rumRuleDeleteResponseMessageJSON `json:"-"`
}

// rumRuleDeleteResponseMessageJSON contains the JSON metadata for the struct
// [RumRuleDeleteResponseMessage]
type rumRuleDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumRuleDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumRuleDeleteResponseResult struct {
	// The Web Analytics rule identifier.
	ID   string                          `json:"id"`
	JSON rumRuleDeleteResponseResultJSON `json:"-"`
}

// rumRuleDeleteResponseResultJSON contains the JSON metadata for the struct
// [RumRuleDeleteResponseResult]
type rumRuleDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumRuleDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RumRuleNewParams struct {
	Host param.Field[string] `json:"host"`
	// Whether the rule includes or excludes traffic from being measured.
	Inclusive param.Field[bool] `json:"inclusive"`
	// Whether the rule is paused or not.
	IsPaused param.Field[bool]     `json:"is_paused"`
	Paths    param.Field[[]string] `json:"paths"`
}

func (r RumRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RumRuleUpdateParams struct {
	Host param.Field[string] `json:"host"`
	// Whether the rule includes or excludes traffic from being measured.
	Inclusive param.Field[bool] `json:"inclusive"`
	// Whether the rule is paused or not.
	IsPaused param.Field[bool]     `json:"is_paused"`
	Paths    param.Field[[]string] `json:"paths"`
}

func (r RumRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
