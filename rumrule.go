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

// RUMRuleService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewRUMRuleService] method instead.
type RUMRuleService struct {
	Options []option.RequestOption
}

// NewRUMRuleService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRUMRuleService(opts ...option.RequestOption) (r *RUMRuleService) {
	r = &RUMRuleService{}
	r.Options = opts
	return
}

// Creates a new rule in a Web Analytics ruleset.
func (r *RUMRuleService) New(ctx context.Context, accountID string, rulesetID string, body RUMRuleNewParams, opts ...option.RequestOption) (res *RUMRuleNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RUMRuleNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rum/v2/%s/rule", accountID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a rule in a Web Analytics ruleset.
func (r *RUMRuleService) Update(ctx context.Context, accountID string, rulesetID string, ruleID string, body RUMRuleUpdateParams, opts ...option.RequestOption) (res *RUMRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RUMRuleUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rum/v2/%s/rule/%s", accountID, rulesetID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all the rules in a Web Analytics ruleset.
func (r *RUMRuleService) List(ctx context.Context, accountID string, rulesetID string, opts ...option.RequestOption) (res *RUMRuleListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RUMRuleListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rum/v2/%s/rules", accountID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an existing rule from a Web Analytics ruleset.
func (r *RUMRuleService) Delete(ctx context.Context, accountID string, rulesetID string, ruleID string, opts ...option.RequestOption) (res *RUMRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RUMRuleDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rum/v2/%s/rule/%s", accountID, rulesetID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RUMRuleNewResponse struct {
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
	Paths    []string               `json:"paths"`
	Priority float64                `json:"priority"`
	JSON     rumRuleNewResponseJSON `json:"-"`
}

// rumRuleNewResponseJSON contains the JSON metadata for the struct
// [RUMRuleNewResponse]
type rumRuleNewResponseJSON struct {
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

func (r *RUMRuleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RUMRuleUpdateResponse struct {
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
	Paths    []string                  `json:"paths"`
	Priority float64                   `json:"priority"`
	JSON     rumRuleUpdateResponseJSON `json:"-"`
}

// rumRuleUpdateResponseJSON contains the JSON metadata for the struct
// [RUMRuleUpdateResponse]
type rumRuleUpdateResponseJSON struct {
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

func (r *RUMRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RUMRuleListResponse struct {
	// A list of rules.
	Rules   []RUMRuleListResponseRule  `json:"rules"`
	Ruleset RUMRuleListResponseRuleset `json:"ruleset"`
	JSON    rumRuleListResponseJSON    `json:"-"`
}

// rumRuleListResponseJSON contains the JSON metadata for the struct
// [RUMRuleListResponse]
type rumRuleListResponseJSON struct {
	Rules       apijson.Field
	Ruleset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RUMRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RUMRuleListResponseRule struct {
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
	JSON     rumRuleListResponseRuleJSON `json:"-"`
}

// rumRuleListResponseRuleJSON contains the JSON metadata for the struct
// [RUMRuleListResponseRule]
type rumRuleListResponseRuleJSON struct {
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

func (r *RUMRuleListResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RUMRuleListResponseRuleset struct {
	// The Web Analytics ruleset identifier.
	ID string `json:"id"`
	// Whether the ruleset is enabled.
	Enabled  bool   `json:"enabled"`
	ZoneName string `json:"zone_name"`
	// The zone identifier.
	ZoneTag string                         `json:"zone_tag"`
	JSON    rumRuleListResponseRulesetJSON `json:"-"`
}

// rumRuleListResponseRulesetJSON contains the JSON metadata for the struct
// [RUMRuleListResponseRuleset]
type rumRuleListResponseRulesetJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	ZoneName    apijson.Field
	ZoneTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RUMRuleListResponseRuleset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RUMRuleDeleteResponse struct {
	// The Web Analytics rule identifier.
	ID   string                    `json:"id"`
	JSON rumRuleDeleteResponseJSON `json:"-"`
}

// rumRuleDeleteResponseJSON contains the JSON metadata for the struct
// [RUMRuleDeleteResponse]
type rumRuleDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RUMRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RUMRuleNewParams struct {
	Host param.Field[string] `json:"host"`
	// Whether the rule includes or excludes traffic from being measured.
	Inclusive param.Field[bool] `json:"inclusive"`
	// Whether the rule is paused or not.
	IsPaused param.Field[bool]     `json:"is_paused"`
	Paths    param.Field[[]string] `json:"paths"`
}

func (r RUMRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RUMRuleNewResponseEnvelope struct {
	Result RUMRuleNewResponse             `json:"result"`
	JSON   rumRuleNewResponseEnvelopeJSON `json:"-"`
}

// rumRuleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RUMRuleNewResponseEnvelope]
type rumRuleNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RUMRuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RUMRuleUpdateParams struct {
	Host param.Field[string] `json:"host"`
	// Whether the rule includes or excludes traffic from being measured.
	Inclusive param.Field[bool] `json:"inclusive"`
	// Whether the rule is paused or not.
	IsPaused param.Field[bool]     `json:"is_paused"`
	Paths    param.Field[[]string] `json:"paths"`
}

func (r RUMRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RUMRuleUpdateResponseEnvelope struct {
	Result RUMRuleUpdateResponse             `json:"result"`
	JSON   rumRuleUpdateResponseEnvelopeJSON `json:"-"`
}

// rumRuleUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [RUMRuleUpdateResponseEnvelope]
type rumRuleUpdateResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RUMRuleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RUMRuleListResponseEnvelope struct {
	Result RUMRuleListResponse             `json:"result"`
	JSON   rumRuleListResponseEnvelopeJSON `json:"-"`
}

// rumRuleListResponseEnvelopeJSON contains the JSON metadata for the struct
// [RUMRuleListResponseEnvelope]
type rumRuleListResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RUMRuleListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RUMRuleDeleteResponseEnvelope struct {
	Result RUMRuleDeleteResponse             `json:"result"`
	JSON   rumRuleDeleteResponseEnvelopeJSON `json:"-"`
}

// rumRuleDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [RUMRuleDeleteResponseEnvelope]
type rumRuleDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RUMRuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
