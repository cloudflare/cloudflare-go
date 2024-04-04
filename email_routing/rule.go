// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// RuleService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewRuleService] method instead.
type RuleService struct {
	Options   []option.RequestOption
	CatchAlls *RuleCatchAllService
}

// NewRuleService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRuleService(opts ...option.RequestOption) (r *RuleService) {
	r = &RuleService{}
	r.Options = opts
	r.CatchAlls = NewRuleCatchAllService(opts...)
	return
}

// Rules consist of a set of criteria for matching emails (such as an email being
// sent to a specific custom email address) plus a set of actions to take on the
// email (like forwarding it to a specific destination address).
func (r *RuleService) New(ctx context.Context, zoneIdentifier string, body RuleNewParams, opts ...option.RequestOption) (res *RuleNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update actions and matches, or enable/disable specific routing rules.
func (r *RuleService) Update(ctx context.Context, zoneIdentifier string, ruleIdentifier string, body RuleUpdateParams, opts ...option.RequestOption) (res *RuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/rules/%s", zoneIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists existing routing rules.
func (r *RuleService) List(ctx context.Context, zoneIdentifier string, query RuleListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[RuleListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/email/routing/rules", zoneIdentifier)
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

// Lists existing routing rules.
func (r *RuleService) ListAutoPaging(ctx context.Context, zoneIdentifier string, query RuleListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[RuleListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, zoneIdentifier, query, opts...))
}

// Delete a specific routing rule.
func (r *RuleService) Delete(ctx context.Context, zoneIdentifier string, ruleIdentifier string, opts ...option.RequestOption) (res *RuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/rules/%s", zoneIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get information for a specific routing rule already created.
func (r *RuleService) Get(ctx context.Context, zoneIdentifier string, ruleIdentifier string, opts ...option.RequestOption) (res *RuleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/rules/%s", zoneIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RuleNewResponse struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions patterns.
	Actions []RuleNewResponseAction `json:"actions"`
	// Routing rule status.
	Enabled RuleNewResponseEnabled `json:"enabled"`
	// Matching patterns to forward to your actions.
	Matchers []RuleNewResponseMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Priority of the routing rule.
	Priority float64 `json:"priority"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string              `json:"tag"`
	JSON ruleNewResponseJSON `json:"-"`
}

// ruleNewResponseJSON contains the JSON metadata for the struct [RuleNewResponse]
type ruleNewResponseJSON struct {
	ID          apijson.Field
	Actions     apijson.Field
	Enabled     apijson.Field
	Matchers    apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseJSON) RawJSON() string {
	return r.raw
}

// Actions pattern.
type RuleNewResponseAction struct {
	// Type of supported action.
	Type  RuleNewResponseActionsType `json:"type,required"`
	Value []string                   `json:"value,required"`
	JSON  ruleNewResponseActionJSON  `json:"-"`
}

// ruleNewResponseActionJSON contains the JSON metadata for the struct
// [RuleNewResponseAction]
type ruleNewResponseActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseActionJSON) RawJSON() string {
	return r.raw
}

// Type of supported action.
type RuleNewResponseActionsType string

const (
	RuleNewResponseActionsTypeDrop    RuleNewResponseActionsType = "drop"
	RuleNewResponseActionsTypeForward RuleNewResponseActionsType = "forward"
	RuleNewResponseActionsTypeWorker  RuleNewResponseActionsType = "worker"
)

func (r RuleNewResponseActionsType) IsKnown() bool {
	switch r {
	case RuleNewResponseActionsTypeDrop, RuleNewResponseActionsTypeForward, RuleNewResponseActionsTypeWorker:
		return true
	}
	return false
}

// Routing rule status.
type RuleNewResponseEnabled bool

const (
	RuleNewResponseEnabledTrue  RuleNewResponseEnabled = true
	RuleNewResponseEnabledFalse RuleNewResponseEnabled = false
)

func (r RuleNewResponseEnabled) IsKnown() bool {
	switch r {
	case RuleNewResponseEnabledTrue, RuleNewResponseEnabledFalse:
		return true
	}
	return false
}

// Matching pattern to forward your actions.
type RuleNewResponseMatcher struct {
	// Field for type matcher.
	Field RuleNewResponseMatchersField `json:"field,required"`
	// Type of matcher.
	Type RuleNewResponseMatchersType `json:"type,required"`
	// Value for matcher.
	Value string                     `json:"value,required"`
	JSON  ruleNewResponseMatcherJSON `json:"-"`
}

// ruleNewResponseMatcherJSON contains the JSON metadata for the struct
// [RuleNewResponseMatcher]
type ruleNewResponseMatcherJSON struct {
	Field       apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseMatcherJSON) RawJSON() string {
	return r.raw
}

// Field for type matcher.
type RuleNewResponseMatchersField string

const (
	RuleNewResponseMatchersFieldTo RuleNewResponseMatchersField = "to"
)

func (r RuleNewResponseMatchersField) IsKnown() bool {
	switch r {
	case RuleNewResponseMatchersFieldTo:
		return true
	}
	return false
}

// Type of matcher.
type RuleNewResponseMatchersType string

const (
	RuleNewResponseMatchersTypeLiteral RuleNewResponseMatchersType = "literal"
)

func (r RuleNewResponseMatchersType) IsKnown() bool {
	switch r {
	case RuleNewResponseMatchersTypeLiteral:
		return true
	}
	return false
}

type RuleUpdateResponse struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions patterns.
	Actions []RuleUpdateResponseAction `json:"actions"`
	// Routing rule status.
	Enabled RuleUpdateResponseEnabled `json:"enabled"`
	// Matching patterns to forward to your actions.
	Matchers []RuleUpdateResponseMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Priority of the routing rule.
	Priority float64 `json:"priority"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string                 `json:"tag"`
	JSON ruleUpdateResponseJSON `json:"-"`
}

// ruleUpdateResponseJSON contains the JSON metadata for the struct
// [RuleUpdateResponse]
type ruleUpdateResponseJSON struct {
	ID          apijson.Field
	Actions     apijson.Field
	Enabled     apijson.Field
	Matchers    apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Actions pattern.
type RuleUpdateResponseAction struct {
	// Type of supported action.
	Type  RuleUpdateResponseActionsType `json:"type,required"`
	Value []string                      `json:"value,required"`
	JSON  ruleUpdateResponseActionJSON  `json:"-"`
}

// ruleUpdateResponseActionJSON contains the JSON metadata for the struct
// [RuleUpdateResponseAction]
type ruleUpdateResponseActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleUpdateResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleUpdateResponseActionJSON) RawJSON() string {
	return r.raw
}

// Type of supported action.
type RuleUpdateResponseActionsType string

const (
	RuleUpdateResponseActionsTypeDrop    RuleUpdateResponseActionsType = "drop"
	RuleUpdateResponseActionsTypeForward RuleUpdateResponseActionsType = "forward"
	RuleUpdateResponseActionsTypeWorker  RuleUpdateResponseActionsType = "worker"
)

func (r RuleUpdateResponseActionsType) IsKnown() bool {
	switch r {
	case RuleUpdateResponseActionsTypeDrop, RuleUpdateResponseActionsTypeForward, RuleUpdateResponseActionsTypeWorker:
		return true
	}
	return false
}

// Routing rule status.
type RuleUpdateResponseEnabled bool

const (
	RuleUpdateResponseEnabledTrue  RuleUpdateResponseEnabled = true
	RuleUpdateResponseEnabledFalse RuleUpdateResponseEnabled = false
)

func (r RuleUpdateResponseEnabled) IsKnown() bool {
	switch r {
	case RuleUpdateResponseEnabledTrue, RuleUpdateResponseEnabledFalse:
		return true
	}
	return false
}

// Matching pattern to forward your actions.
type RuleUpdateResponseMatcher struct {
	// Field for type matcher.
	Field RuleUpdateResponseMatchersField `json:"field,required"`
	// Type of matcher.
	Type RuleUpdateResponseMatchersType `json:"type,required"`
	// Value for matcher.
	Value string                        `json:"value,required"`
	JSON  ruleUpdateResponseMatcherJSON `json:"-"`
}

// ruleUpdateResponseMatcherJSON contains the JSON metadata for the struct
// [RuleUpdateResponseMatcher]
type ruleUpdateResponseMatcherJSON struct {
	Field       apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleUpdateResponseMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleUpdateResponseMatcherJSON) RawJSON() string {
	return r.raw
}

// Field for type matcher.
type RuleUpdateResponseMatchersField string

const (
	RuleUpdateResponseMatchersFieldTo RuleUpdateResponseMatchersField = "to"
)

func (r RuleUpdateResponseMatchersField) IsKnown() bool {
	switch r {
	case RuleUpdateResponseMatchersFieldTo:
		return true
	}
	return false
}

// Type of matcher.
type RuleUpdateResponseMatchersType string

const (
	RuleUpdateResponseMatchersTypeLiteral RuleUpdateResponseMatchersType = "literal"
)

func (r RuleUpdateResponseMatchersType) IsKnown() bool {
	switch r {
	case RuleUpdateResponseMatchersTypeLiteral:
		return true
	}
	return false
}

type RuleListResponse struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions patterns.
	Actions []RuleListResponseAction `json:"actions"`
	// Routing rule status.
	Enabled RuleListResponseEnabled `json:"enabled"`
	// Matching patterns to forward to your actions.
	Matchers []RuleListResponseMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Priority of the routing rule.
	Priority float64 `json:"priority"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string               `json:"tag"`
	JSON ruleListResponseJSON `json:"-"`
}

// ruleListResponseJSON contains the JSON metadata for the struct
// [RuleListResponse]
type ruleListResponseJSON struct {
	ID          apijson.Field
	Actions     apijson.Field
	Enabled     apijson.Field
	Matchers    apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleListResponseJSON) RawJSON() string {
	return r.raw
}

// Actions pattern.
type RuleListResponseAction struct {
	// Type of supported action.
	Type  RuleListResponseActionsType `json:"type,required"`
	Value []string                    `json:"value,required"`
	JSON  ruleListResponseActionJSON  `json:"-"`
}

// ruleListResponseActionJSON contains the JSON metadata for the struct
// [RuleListResponseAction]
type ruleListResponseActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleListResponseActionJSON) RawJSON() string {
	return r.raw
}

// Type of supported action.
type RuleListResponseActionsType string

const (
	RuleListResponseActionsTypeDrop    RuleListResponseActionsType = "drop"
	RuleListResponseActionsTypeForward RuleListResponseActionsType = "forward"
	RuleListResponseActionsTypeWorker  RuleListResponseActionsType = "worker"
)

func (r RuleListResponseActionsType) IsKnown() bool {
	switch r {
	case RuleListResponseActionsTypeDrop, RuleListResponseActionsTypeForward, RuleListResponseActionsTypeWorker:
		return true
	}
	return false
}

// Routing rule status.
type RuleListResponseEnabled bool

const (
	RuleListResponseEnabledTrue  RuleListResponseEnabled = true
	RuleListResponseEnabledFalse RuleListResponseEnabled = false
)

func (r RuleListResponseEnabled) IsKnown() bool {
	switch r {
	case RuleListResponseEnabledTrue, RuleListResponseEnabledFalse:
		return true
	}
	return false
}

// Matching pattern to forward your actions.
type RuleListResponseMatcher struct {
	// Field for type matcher.
	Field RuleListResponseMatchersField `json:"field,required"`
	// Type of matcher.
	Type RuleListResponseMatchersType `json:"type,required"`
	// Value for matcher.
	Value string                      `json:"value,required"`
	JSON  ruleListResponseMatcherJSON `json:"-"`
}

// ruleListResponseMatcherJSON contains the JSON metadata for the struct
// [RuleListResponseMatcher]
type ruleListResponseMatcherJSON struct {
	Field       apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListResponseMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleListResponseMatcherJSON) RawJSON() string {
	return r.raw
}

// Field for type matcher.
type RuleListResponseMatchersField string

const (
	RuleListResponseMatchersFieldTo RuleListResponseMatchersField = "to"
)

func (r RuleListResponseMatchersField) IsKnown() bool {
	switch r {
	case RuleListResponseMatchersFieldTo:
		return true
	}
	return false
}

// Type of matcher.
type RuleListResponseMatchersType string

const (
	RuleListResponseMatchersTypeLiteral RuleListResponseMatchersType = "literal"
)

func (r RuleListResponseMatchersType) IsKnown() bool {
	switch r {
	case RuleListResponseMatchersTypeLiteral:
		return true
	}
	return false
}

type RuleDeleteResponse struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions patterns.
	Actions []RuleDeleteResponseAction `json:"actions"`
	// Routing rule status.
	Enabled RuleDeleteResponseEnabled `json:"enabled"`
	// Matching patterns to forward to your actions.
	Matchers []RuleDeleteResponseMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Priority of the routing rule.
	Priority float64 `json:"priority"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string                 `json:"tag"`
	JSON ruleDeleteResponseJSON `json:"-"`
}

// ruleDeleteResponseJSON contains the JSON metadata for the struct
// [RuleDeleteResponse]
type ruleDeleteResponseJSON struct {
	ID          apijson.Field
	Actions     apijson.Field
	Enabled     apijson.Field
	Matchers    apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Actions pattern.
type RuleDeleteResponseAction struct {
	// Type of supported action.
	Type  RuleDeleteResponseActionsType `json:"type,required"`
	Value []string                      `json:"value,required"`
	JSON  ruleDeleteResponseActionJSON  `json:"-"`
}

// ruleDeleteResponseActionJSON contains the JSON metadata for the struct
// [RuleDeleteResponseAction]
type ruleDeleteResponseActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseActionJSON) RawJSON() string {
	return r.raw
}

// Type of supported action.
type RuleDeleteResponseActionsType string

const (
	RuleDeleteResponseActionsTypeDrop    RuleDeleteResponseActionsType = "drop"
	RuleDeleteResponseActionsTypeForward RuleDeleteResponseActionsType = "forward"
	RuleDeleteResponseActionsTypeWorker  RuleDeleteResponseActionsType = "worker"
)

func (r RuleDeleteResponseActionsType) IsKnown() bool {
	switch r {
	case RuleDeleteResponseActionsTypeDrop, RuleDeleteResponseActionsTypeForward, RuleDeleteResponseActionsTypeWorker:
		return true
	}
	return false
}

// Routing rule status.
type RuleDeleteResponseEnabled bool

const (
	RuleDeleteResponseEnabledTrue  RuleDeleteResponseEnabled = true
	RuleDeleteResponseEnabledFalse RuleDeleteResponseEnabled = false
)

func (r RuleDeleteResponseEnabled) IsKnown() bool {
	switch r {
	case RuleDeleteResponseEnabledTrue, RuleDeleteResponseEnabledFalse:
		return true
	}
	return false
}

// Matching pattern to forward your actions.
type RuleDeleteResponseMatcher struct {
	// Field for type matcher.
	Field RuleDeleteResponseMatchersField `json:"field,required"`
	// Type of matcher.
	Type RuleDeleteResponseMatchersType `json:"type,required"`
	// Value for matcher.
	Value string                        `json:"value,required"`
	JSON  ruleDeleteResponseMatcherJSON `json:"-"`
}

// ruleDeleteResponseMatcherJSON contains the JSON metadata for the struct
// [RuleDeleteResponseMatcher]
type ruleDeleteResponseMatcherJSON struct {
	Field       apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseMatcherJSON) RawJSON() string {
	return r.raw
}

// Field for type matcher.
type RuleDeleteResponseMatchersField string

const (
	RuleDeleteResponseMatchersFieldTo RuleDeleteResponseMatchersField = "to"
)

func (r RuleDeleteResponseMatchersField) IsKnown() bool {
	switch r {
	case RuleDeleteResponseMatchersFieldTo:
		return true
	}
	return false
}

// Type of matcher.
type RuleDeleteResponseMatchersType string

const (
	RuleDeleteResponseMatchersTypeLiteral RuleDeleteResponseMatchersType = "literal"
)

func (r RuleDeleteResponseMatchersType) IsKnown() bool {
	switch r {
	case RuleDeleteResponseMatchersTypeLiteral:
		return true
	}
	return false
}

type RuleGetResponse struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions patterns.
	Actions []RuleGetResponseAction `json:"actions"`
	// Routing rule status.
	Enabled RuleGetResponseEnabled `json:"enabled"`
	// Matching patterns to forward to your actions.
	Matchers []RuleGetResponseMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Priority of the routing rule.
	Priority float64 `json:"priority"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string              `json:"tag"`
	JSON ruleGetResponseJSON `json:"-"`
}

// ruleGetResponseJSON contains the JSON metadata for the struct [RuleGetResponse]
type ruleGetResponseJSON struct {
	ID          apijson.Field
	Actions     apijson.Field
	Enabled     apijson.Field
	Matchers    apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleGetResponseJSON) RawJSON() string {
	return r.raw
}

// Actions pattern.
type RuleGetResponseAction struct {
	// Type of supported action.
	Type  RuleGetResponseActionsType `json:"type,required"`
	Value []string                   `json:"value,required"`
	JSON  ruleGetResponseActionJSON  `json:"-"`
}

// ruleGetResponseActionJSON contains the JSON metadata for the struct
// [RuleGetResponseAction]
type ruleGetResponseActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleGetResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleGetResponseActionJSON) RawJSON() string {
	return r.raw
}

// Type of supported action.
type RuleGetResponseActionsType string

const (
	RuleGetResponseActionsTypeDrop    RuleGetResponseActionsType = "drop"
	RuleGetResponseActionsTypeForward RuleGetResponseActionsType = "forward"
	RuleGetResponseActionsTypeWorker  RuleGetResponseActionsType = "worker"
)

func (r RuleGetResponseActionsType) IsKnown() bool {
	switch r {
	case RuleGetResponseActionsTypeDrop, RuleGetResponseActionsTypeForward, RuleGetResponseActionsTypeWorker:
		return true
	}
	return false
}

// Routing rule status.
type RuleGetResponseEnabled bool

const (
	RuleGetResponseEnabledTrue  RuleGetResponseEnabled = true
	RuleGetResponseEnabledFalse RuleGetResponseEnabled = false
)

func (r RuleGetResponseEnabled) IsKnown() bool {
	switch r {
	case RuleGetResponseEnabledTrue, RuleGetResponseEnabledFalse:
		return true
	}
	return false
}

// Matching pattern to forward your actions.
type RuleGetResponseMatcher struct {
	// Field for type matcher.
	Field RuleGetResponseMatchersField `json:"field,required"`
	// Type of matcher.
	Type RuleGetResponseMatchersType `json:"type,required"`
	// Value for matcher.
	Value string                     `json:"value,required"`
	JSON  ruleGetResponseMatcherJSON `json:"-"`
}

// ruleGetResponseMatcherJSON contains the JSON metadata for the struct
// [RuleGetResponseMatcher]
type ruleGetResponseMatcherJSON struct {
	Field       apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleGetResponseMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleGetResponseMatcherJSON) RawJSON() string {
	return r.raw
}

// Field for type matcher.
type RuleGetResponseMatchersField string

const (
	RuleGetResponseMatchersFieldTo RuleGetResponseMatchersField = "to"
)

func (r RuleGetResponseMatchersField) IsKnown() bool {
	switch r {
	case RuleGetResponseMatchersFieldTo:
		return true
	}
	return false
}

// Type of matcher.
type RuleGetResponseMatchersType string

const (
	RuleGetResponseMatchersTypeLiteral RuleGetResponseMatchersType = "literal"
)

func (r RuleGetResponseMatchersType) IsKnown() bool {
	switch r {
	case RuleGetResponseMatchersTypeLiteral:
		return true
	}
	return false
}

type RuleNewParams struct {
	// List actions patterns.
	Actions param.Field[[]RuleNewParamsAction] `json:"actions,required"`
	// Matching patterns to forward to your actions.
	Matchers param.Field[[]RuleNewParamsMatcher] `json:"matchers,required"`
	// Routing rule status.
	Enabled param.Field[RuleNewParamsEnabled] `json:"enabled"`
	// Routing rule name.
	Name param.Field[string] `json:"name"`
	// Priority of the routing rule.
	Priority param.Field[float64] `json:"priority"`
}

func (r RuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Actions pattern.
type RuleNewParamsAction struct {
	// Type of supported action.
	Type  param.Field[RuleNewParamsActionsType] `json:"type,required"`
	Value param.Field[[]string]                 `json:"value,required"`
}

func (r RuleNewParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of supported action.
type RuleNewParamsActionsType string

const (
	RuleNewParamsActionsTypeDrop    RuleNewParamsActionsType = "drop"
	RuleNewParamsActionsTypeForward RuleNewParamsActionsType = "forward"
	RuleNewParamsActionsTypeWorker  RuleNewParamsActionsType = "worker"
)

func (r RuleNewParamsActionsType) IsKnown() bool {
	switch r {
	case RuleNewParamsActionsTypeDrop, RuleNewParamsActionsTypeForward, RuleNewParamsActionsTypeWorker:
		return true
	}
	return false
}

// Matching pattern to forward your actions.
type RuleNewParamsMatcher struct {
	// Field for type matcher.
	Field param.Field[RuleNewParamsMatchersField] `json:"field,required"`
	// Type of matcher.
	Type param.Field[RuleNewParamsMatchersType] `json:"type,required"`
	// Value for matcher.
	Value param.Field[string] `json:"value,required"`
}

func (r RuleNewParamsMatcher) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Field for type matcher.
type RuleNewParamsMatchersField string

const (
	RuleNewParamsMatchersFieldTo RuleNewParamsMatchersField = "to"
)

func (r RuleNewParamsMatchersField) IsKnown() bool {
	switch r {
	case RuleNewParamsMatchersFieldTo:
		return true
	}
	return false
}

// Type of matcher.
type RuleNewParamsMatchersType string

const (
	RuleNewParamsMatchersTypeLiteral RuleNewParamsMatchersType = "literal"
)

func (r RuleNewParamsMatchersType) IsKnown() bool {
	switch r {
	case RuleNewParamsMatchersTypeLiteral:
		return true
	}
	return false
}

// Routing rule status.
type RuleNewParamsEnabled bool

const (
	RuleNewParamsEnabledTrue  RuleNewParamsEnabled = true
	RuleNewParamsEnabledFalse RuleNewParamsEnabled = false
)

func (r RuleNewParamsEnabled) IsKnown() bool {
	switch r {
	case RuleNewParamsEnabledTrue, RuleNewParamsEnabledFalse:
		return true
	}
	return false
}

type RuleNewResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   RuleNewResponse                                           `json:"result,required"`
	// Whether the API call was successful
	Success RuleNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleNewResponseEnvelopeJSON    `json:"-"`
}

// ruleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleNewResponseEnvelope]
type ruleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleNewResponseEnvelopeSuccess bool

const (
	RuleNewResponseEnvelopeSuccessTrue RuleNewResponseEnvelopeSuccess = true
)

func (r RuleNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RuleUpdateParams struct {
	// List actions patterns.
	Actions param.Field[[]RuleUpdateParamsAction] `json:"actions,required"`
	// Matching patterns to forward to your actions.
	Matchers param.Field[[]RuleUpdateParamsMatcher] `json:"matchers,required"`
	// Routing rule status.
	Enabled param.Field[RuleUpdateParamsEnabled] `json:"enabled"`
	// Routing rule name.
	Name param.Field[string] `json:"name"`
	// Priority of the routing rule.
	Priority param.Field[float64] `json:"priority"`
}

func (r RuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Actions pattern.
type RuleUpdateParamsAction struct {
	// Type of supported action.
	Type  param.Field[RuleUpdateParamsActionsType] `json:"type,required"`
	Value param.Field[[]string]                    `json:"value,required"`
}

func (r RuleUpdateParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of supported action.
type RuleUpdateParamsActionsType string

const (
	RuleUpdateParamsActionsTypeDrop    RuleUpdateParamsActionsType = "drop"
	RuleUpdateParamsActionsTypeForward RuleUpdateParamsActionsType = "forward"
	RuleUpdateParamsActionsTypeWorker  RuleUpdateParamsActionsType = "worker"
)

func (r RuleUpdateParamsActionsType) IsKnown() bool {
	switch r {
	case RuleUpdateParamsActionsTypeDrop, RuleUpdateParamsActionsTypeForward, RuleUpdateParamsActionsTypeWorker:
		return true
	}
	return false
}

// Matching pattern to forward your actions.
type RuleUpdateParamsMatcher struct {
	// Field for type matcher.
	Field param.Field[RuleUpdateParamsMatchersField] `json:"field,required"`
	// Type of matcher.
	Type param.Field[RuleUpdateParamsMatchersType] `json:"type,required"`
	// Value for matcher.
	Value param.Field[string] `json:"value,required"`
}

func (r RuleUpdateParamsMatcher) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Field for type matcher.
type RuleUpdateParamsMatchersField string

const (
	RuleUpdateParamsMatchersFieldTo RuleUpdateParamsMatchersField = "to"
)

func (r RuleUpdateParamsMatchersField) IsKnown() bool {
	switch r {
	case RuleUpdateParamsMatchersFieldTo:
		return true
	}
	return false
}

// Type of matcher.
type RuleUpdateParamsMatchersType string

const (
	RuleUpdateParamsMatchersTypeLiteral RuleUpdateParamsMatchersType = "literal"
)

func (r RuleUpdateParamsMatchersType) IsKnown() bool {
	switch r {
	case RuleUpdateParamsMatchersTypeLiteral:
		return true
	}
	return false
}

// Routing rule status.
type RuleUpdateParamsEnabled bool

const (
	RuleUpdateParamsEnabledTrue  RuleUpdateParamsEnabled = true
	RuleUpdateParamsEnabledFalse RuleUpdateParamsEnabled = false
)

func (r RuleUpdateParamsEnabled) IsKnown() bool {
	switch r {
	case RuleUpdateParamsEnabledTrue, RuleUpdateParamsEnabledFalse:
		return true
	}
	return false
}

type RuleUpdateResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   RuleUpdateResponse                                        `json:"result,required"`
	// Whether the API call was successful
	Success RuleUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleUpdateResponseEnvelopeJSON    `json:"-"`
}

// ruleUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleUpdateResponseEnvelope]
type ruleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleUpdateResponseEnvelopeSuccess bool

const (
	RuleUpdateResponseEnvelopeSuccessTrue RuleUpdateResponseEnvelopeSuccess = true
)

func (r RuleUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RuleListParams struct {
	// Filter by enabled routing rules.
	Enabled param.Field[RuleListParamsEnabled] `query:"enabled"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [RuleListParams]'s query parameters as `url.Values`.
func (r RuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by enabled routing rules.
type RuleListParamsEnabled bool

const (
	RuleListParamsEnabledTrue  RuleListParamsEnabled = true
	RuleListParamsEnabledFalse RuleListParamsEnabled = false
)

func (r RuleListParamsEnabled) IsKnown() bool {
	switch r {
	case RuleListParamsEnabledTrue, RuleListParamsEnabledFalse:
		return true
	}
	return false
}

type RuleDeleteResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   RuleDeleteResponse                                        `json:"result,required"`
	// Whether the API call was successful
	Success RuleDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleDeleteResponseEnvelopeJSON    `json:"-"`
}

// ruleDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleDeleteResponseEnvelope]
type ruleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleDeleteResponseEnvelopeSuccess bool

const (
	RuleDeleteResponseEnvelopeSuccessTrue RuleDeleteResponseEnvelopeSuccess = true
)

func (r RuleDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RuleGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   RuleGetResponse                                           `json:"result,required"`
	// Whether the API call was successful
	Success RuleGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleGetResponseEnvelopeJSON    `json:"-"`
}

// ruleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleGetResponseEnvelope]
type ruleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleGetResponseEnvelopeSuccess bool

const (
	RuleGetResponseEnvelopeSuccessTrue RuleGetResponseEnvelopeSuccess = true
)

func (r RuleGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
