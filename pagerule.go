// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// PageruleService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewPageruleService] method instead.
type PageruleService struct {
	Options []option.RequestOption
}

// NewPageruleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPageruleService(opts ...option.RequestOption) (r *PageruleService) {
	r = &PageruleService{}
	r.Options = opts
	return
}

// Updates one or more fields of an existing Page Rule.
func (r *PageruleService) Update(ctx context.Context, zoneIdentifier string, identifier string, body PageruleUpdateParams, opts ...option.RequestOption) (res *PageruleResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/pagerules/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type PageruleUpdateParams struct {
	// The set of actions to perform if the targets of this rule match the request.
	// Actions can redirect to another URL or override settings, but not both.
	Actions param.Field[[]PageruleUpdateParamsAction] `json:"actions"`
	// The priority of the rule, used to define which Page Rule is processed over
	// another. A higher number indicates a higher priority. For example, if you have a
	// catch-all Page Rule (rule A: `/images/*`) but want a more specific Page Rule to
	// take precedence (rule B: `/images/special/*`), specify a higher priority for
	// rule B so it overrides rule A.
	Priority param.Field[int64] `json:"priority"`
	// The status of the Page Rule.
	Status param.Field[PageruleUpdateParamsStatus] `json:"status"`
	// The rule targets to evaluate on each request.
	Targets param.Field[[]PageruleUpdateParamsTarget] `json:"targets"`
}

func (r PageruleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PageruleUpdateParamsAction struct {
	// The type of route.
	Name  param.Field[PageruleUpdateParamsActionsName]  `json:"name"`
	Value param.Field[PageruleUpdateParamsActionsValue] `json:"value"`
}

func (r PageruleUpdateParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of route.
type PageruleUpdateParamsActionsName string

const (
	PageruleUpdateParamsActionsNameForwardURL PageruleUpdateParamsActionsName = "forward_url"
)

type PageruleUpdateParamsActionsValue struct {
	// The response type for the URL redirect.
	Type param.Field[PageruleUpdateParamsActionsValueType] `json:"type"`
	// The URL to redirect the request to. Notes: ${num} refers to the position of '\*'
	// in the constraint value.
	URL param.Field[string] `json:"url"`
}

func (r PageruleUpdateParamsActionsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The response type for the URL redirect.
type PageruleUpdateParamsActionsValueType string

const (
	PageruleUpdateParamsActionsValueTypeTemporary PageruleUpdateParamsActionsValueType = "temporary"
	PageruleUpdateParamsActionsValueTypePermanent PageruleUpdateParamsActionsValueType = "permanent"
)

// The status of the Page Rule.
type PageruleUpdateParamsStatus string

const (
	PageruleUpdateParamsStatusActive   PageruleUpdateParamsStatus = "active"
	PageruleUpdateParamsStatusDisabled PageruleUpdateParamsStatus = "disabled"
)

// A request condition target.
type PageruleUpdateParamsTarget struct {
	// The constraint of a target.
	Constraint param.Field[PageruleUpdateParamsTargetsConstraint] `json:"constraint,required"`
	// A target based on the URL of the request.
	Target param.Field[PageruleUpdateParamsTargetsTarget] `json:"target,required"`
}

func (r PageruleUpdateParamsTarget) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The constraint of a target.
type PageruleUpdateParamsTargetsConstraint struct {
	// The matches operator can use asterisks and pipes as wildcard and 'or' operators.
	Operator param.Field[PageruleUpdateParamsTargetsConstraintOperator] `json:"operator"`
	// The URL pattern to match against the current request. The pattern may contain up
	// to four asterisks ('\*') as placeholders.
	Value param.Field[string] `json:"value"`
}

func (r PageruleUpdateParamsTargetsConstraint) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The matches operator can use asterisks and pipes as wildcard and 'or' operators.
type PageruleUpdateParamsTargetsConstraintOperator string

const (
	PageruleUpdateParamsTargetsConstraintOperatorMatches    PageruleUpdateParamsTargetsConstraintOperator = "matches"
	PageruleUpdateParamsTargetsConstraintOperatorContains   PageruleUpdateParamsTargetsConstraintOperator = "contains"
	PageruleUpdateParamsTargetsConstraintOperatorEquals     PageruleUpdateParamsTargetsConstraintOperator = "equals"
	PageruleUpdateParamsTargetsConstraintOperatorNotEqual   PageruleUpdateParamsTargetsConstraintOperator = "not_equal"
	PageruleUpdateParamsTargetsConstraintOperatorNotContain PageruleUpdateParamsTargetsConstraintOperator = "not_contain"
)

// A target based on the URL of the request.
type PageruleUpdateParamsTargetsTarget string

const (
	PageruleUpdateParamsTargetsTargetURL PageruleUpdateParamsTargetsTarget = "url"
)
