// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// RuleCatchAllService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRuleCatchAllService] method
// instead.
type RuleCatchAllService struct {
	Options []option.RequestOption
}

// NewRuleCatchAllService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRuleCatchAllService(opts ...option.RequestOption) (r *RuleCatchAllService) {
	r = &RuleCatchAllService{}
	r.Options = opts
	return
}

// Enable or disable catch-all routing rule, or change action to forward to
// specific destination address.
func (r *RuleCatchAllService) Update(ctx context.Context, zoneIdentifier string, body RuleCatchAllUpdateParams, opts ...option.RequestOption) (res *RuleCatchAllUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleCatchAllUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/rules/catch_all", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get information on the default catch-all routing rule.
func (r *RuleCatchAllService) Get(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *RuleCatchAllGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleCatchAllGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/rules/catch_all", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Action for the catch-all routing rule.
type CatchAllAction struct {
	// Type of action for catch-all rule.
	Type  CatchAllActionType `json:"type,required"`
	Value []string           `json:"value"`
	JSON  catchAllActionJSON `json:"-"`
}

// catchAllActionJSON contains the JSON metadata for the struct [CatchAllAction]
type catchAllActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CatchAllAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catchAllActionJSON) RawJSON() string {
	return r.raw
}

// Type of action for catch-all rule.
type CatchAllActionType string

const (
	CatchAllActionTypeDrop    CatchAllActionType = "drop"
	CatchAllActionTypeForward CatchAllActionType = "forward"
	CatchAllActionTypeWorker  CatchAllActionType = "worker"
)

func (r CatchAllActionType) IsKnown() bool {
	switch r {
	case CatchAllActionTypeDrop, CatchAllActionTypeForward, CatchAllActionTypeWorker:
		return true
	}
	return false
}

// Action for the catch-all routing rule.
type CatchAllActionParam struct {
	// Type of action for catch-all rule.
	Type  param.Field[CatchAllActionType] `json:"type,required"`
	Value param.Field[[]string]           `json:"value"`
}

func (r CatchAllActionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matcher for catch-all routing rule.
type CatchAllMatcher struct {
	// Type of matcher. Default is 'all'.
	Type CatchAllMatcherType `json:"type,required"`
	JSON catchAllMatcherJSON `json:"-"`
}

// catchAllMatcherJSON contains the JSON metadata for the struct [CatchAllMatcher]
type catchAllMatcherJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CatchAllMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r catchAllMatcherJSON) RawJSON() string {
	return r.raw
}

// Type of matcher. Default is 'all'.
type CatchAllMatcherType string

const (
	CatchAllMatcherTypeAll CatchAllMatcherType = "all"
)

func (r CatchAllMatcherType) IsKnown() bool {
	switch r {
	case CatchAllMatcherTypeAll:
		return true
	}
	return false
}

// Matcher for catch-all routing rule.
type CatchAllMatcherParam struct {
	// Type of matcher. Default is 'all'.
	Type param.Field[CatchAllMatcherType] `json:"type,required"`
}

func (r CatchAllMatcherParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleCatchAllUpdateResponse struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions for the catch-all routing rule.
	Actions []CatchAllAction `json:"actions"`
	// Routing rule status.
	Enabled RuleCatchAllUpdateResponseEnabled `json:"enabled"`
	// List of matchers for the catch-all routing rule.
	Matchers []CatchAllMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string                         `json:"tag"`
	JSON ruleCatchAllUpdateResponseJSON `json:"-"`
}

// ruleCatchAllUpdateResponseJSON contains the JSON metadata for the struct
// [RuleCatchAllUpdateResponse]
type ruleCatchAllUpdateResponseJSON struct {
	ID          apijson.Field
	Actions     apijson.Field
	Enabled     apijson.Field
	Matchers    apijson.Field
	Name        apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleCatchAllUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleCatchAllUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Routing rule status.
type RuleCatchAllUpdateResponseEnabled bool

const (
	RuleCatchAllUpdateResponseEnabledTrue  RuleCatchAllUpdateResponseEnabled = true
	RuleCatchAllUpdateResponseEnabledFalse RuleCatchAllUpdateResponseEnabled = false
)

func (r RuleCatchAllUpdateResponseEnabled) IsKnown() bool {
	switch r {
	case RuleCatchAllUpdateResponseEnabledTrue, RuleCatchAllUpdateResponseEnabledFalse:
		return true
	}
	return false
}

type RuleCatchAllGetResponse struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions for the catch-all routing rule.
	Actions []CatchAllAction `json:"actions"`
	// Routing rule status.
	Enabled RuleCatchAllGetResponseEnabled `json:"enabled"`
	// List of matchers for the catch-all routing rule.
	Matchers []CatchAllMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string                      `json:"tag"`
	JSON ruleCatchAllGetResponseJSON `json:"-"`
}

// ruleCatchAllGetResponseJSON contains the JSON metadata for the struct
// [RuleCatchAllGetResponse]
type ruleCatchAllGetResponseJSON struct {
	ID          apijson.Field
	Actions     apijson.Field
	Enabled     apijson.Field
	Matchers    apijson.Field
	Name        apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleCatchAllGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleCatchAllGetResponseJSON) RawJSON() string {
	return r.raw
}

// Routing rule status.
type RuleCatchAllGetResponseEnabled bool

const (
	RuleCatchAllGetResponseEnabledTrue  RuleCatchAllGetResponseEnabled = true
	RuleCatchAllGetResponseEnabledFalse RuleCatchAllGetResponseEnabled = false
)

func (r RuleCatchAllGetResponseEnabled) IsKnown() bool {
	switch r {
	case RuleCatchAllGetResponseEnabledTrue, RuleCatchAllGetResponseEnabledFalse:
		return true
	}
	return false
}

type RuleCatchAllUpdateParams struct {
	// List actions for the catch-all routing rule.
	Actions param.Field[[]CatchAllActionParam] `json:"actions,required"`
	// List of matchers for the catch-all routing rule.
	Matchers param.Field[[]CatchAllMatcherParam] `json:"matchers,required"`
	// Routing rule status.
	Enabled param.Field[RuleCatchAllUpdateParamsEnabled] `json:"enabled"`
	// Routing rule name.
	Name param.Field[string] `json:"name"`
}

func (r RuleCatchAllUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Routing rule status.
type RuleCatchAllUpdateParamsEnabled bool

const (
	RuleCatchAllUpdateParamsEnabledTrue  RuleCatchAllUpdateParamsEnabled = true
	RuleCatchAllUpdateParamsEnabledFalse RuleCatchAllUpdateParamsEnabled = false
)

func (r RuleCatchAllUpdateParamsEnabled) IsKnown() bool {
	switch r {
	case RuleCatchAllUpdateParamsEnabledTrue, RuleCatchAllUpdateParamsEnabledFalse:
		return true
	}
	return false
}

type RuleCatchAllUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo      `json:"errors,required"`
	Messages []shared.ResponseInfo      `json:"messages,required"`
	Result   RuleCatchAllUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success RuleCatchAllUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleCatchAllUpdateResponseEnvelopeJSON    `json:"-"`
}

// ruleCatchAllUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleCatchAllUpdateResponseEnvelope]
type ruleCatchAllUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleCatchAllUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleCatchAllUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleCatchAllUpdateResponseEnvelopeSuccess bool

const (
	RuleCatchAllUpdateResponseEnvelopeSuccessTrue RuleCatchAllUpdateResponseEnvelopeSuccess = true
)

func (r RuleCatchAllUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleCatchAllUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RuleCatchAllGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo   `json:"errors,required"`
	Messages []shared.ResponseInfo   `json:"messages,required"`
	Result   RuleCatchAllGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success RuleCatchAllGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleCatchAllGetResponseEnvelopeJSON    `json:"-"`
}

// ruleCatchAllGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleCatchAllGetResponseEnvelope]
type ruleCatchAllGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleCatchAllGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleCatchAllGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleCatchAllGetResponseEnvelopeSuccess bool

const (
	RuleCatchAllGetResponseEnvelopeSuccessTrue RuleCatchAllGetResponseEnvelopeSuccess = true
)

func (r RuleCatchAllGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleCatchAllGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
