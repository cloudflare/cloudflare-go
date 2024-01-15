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

// AccountRumV2RuleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountRumV2RuleService] method
// instead.
type AccountRumV2RuleService struct {
	Options []option.RequestOption
}

// NewAccountRumV2RuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountRumV2RuleService(opts ...option.RequestOption) (r *AccountRumV2RuleService) {
	r = &AccountRumV2RuleService{}
	r.Options = opts
	return
}

// Updates a rule in a Web Analytics ruleset.
func (r *AccountRumV2RuleService) Update(ctx context.Context, accountIdentifier string, rulesetIdentifier string, ruleIdentifier string, body AccountRumV2RuleUpdateParams, opts ...option.RequestOption) (res *Rule, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rum/v2/%s/rule/%s", accountIdentifier, rulesetIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type Rule struct {
	Errors   []RuleError   `json:"errors"`
	Messages []RuleMessage `json:"messages"`
	Result   RuleResult    `json:"result"`
	// Whether the API call was successful.
	Success bool     `json:"success"`
	JSON    ruleJSON `json:"-"`
}

// ruleJSON contains the JSON metadata for the struct [Rule]
type ruleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Rule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleError struct {
	Code    int64         `json:"code,required"`
	Message string        `json:"message,required"`
	JSON    ruleErrorJSON `json:"-"`
}

// ruleErrorJSON contains the JSON metadata for the struct [RuleError]
type ruleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleMessage struct {
	Code    int64           `json:"code,required"`
	Message string          `json:"message,required"`
	JSON    ruleMessageJSON `json:"-"`
}

// ruleMessageJSON contains the JSON metadata for the struct [RuleMessage]
type ruleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RuleResult struct {
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
	Paths    []string       `json:"paths"`
	Priority float64        `json:"priority"`
	JSON     ruleResultJSON `json:"-"`
}

// ruleResultJSON contains the JSON metadata for the struct [RuleResult]
type ruleResultJSON struct {
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

func (r *RuleResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRumV2RuleUpdateParams struct {
	Host param.Field[string] `json:"host"`
	// Whether the rule includes or excludes traffic from being measured.
	Inclusive param.Field[bool] `json:"inclusive"`
	// Whether the rule is paused or not.
	IsPaused param.Field[bool]     `json:"is_paused"`
	Paths    param.Field[[]string] `json:"paths"`
}

func (r AccountRumV2RuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
