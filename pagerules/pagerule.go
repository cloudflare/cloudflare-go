// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagerules

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/shared"
)

// PageruleService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPageruleService] method instead.
type PageruleService struct {
	Options  []option.RequestOption
	Settings *SettingService
}

// NewPageruleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPageruleService(opts ...option.RequestOption) (r *PageruleService) {
	r = &PageruleService{}
	r.Options = opts
	r.Settings = NewSettingService(opts...)
	return
}

// Creates a new Page Rule.
//
// Deprecated: The Page Rules API is deprecated in favour of the Ruleset Engine.
// See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#page-rules
// for full details.
func (r *PageruleService) New(ctx context.Context, params PageruleNewParams, opts ...option.RequestOption) (res *PageruleNewResponse, err error) {
	var env PageruleNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/pagerules", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Replaces the configuration of an existing Page Rule. The configuration of the
// updated Page Rule will exactly match the data passed in the API request.
//
// Deprecated: The Page Rules API is deprecated in favour of the Ruleset Engine.
// See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#page-rules
// for full details.
func (r *PageruleService) Update(ctx context.Context, pageruleID string, params PageruleUpdateParams, opts ...option.RequestOption) (res *PageruleUpdateResponse, err error) {
	var env PageruleUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if pageruleID == "" {
		err = errors.New("missing required pagerule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/pagerules/%s", params.ZoneID, pageruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches Page Rules in a zone.
//
// Deprecated: The Page Rules API is deprecated in favour of the Ruleset Engine.
// See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#page-rules
// for full details.
func (r *PageruleService) List(ctx context.Context, params PageruleListParams, opts ...option.RequestOption) (res *[]PageruleListResponse, err error) {
	var env PageruleListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/pagerules", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an existing Page Rule.
//
// Deprecated: The Page Rules API is deprecated in favour of the Ruleset Engine.
// See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#page-rules
// for full details.
func (r *PageruleService) Delete(ctx context.Context, pageruleID string, body PageruleDeleteParams, opts ...option.RequestOption) (res *PageruleDeleteResponse, err error) {
	var env PageruleDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if pageruleID == "" {
		err = errors.New("missing required pagerule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/pagerules/%s", body.ZoneID, pageruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates one or more fields of an existing Page Rule.
//
// Deprecated: The Page Rules API is deprecated in favour of the Ruleset Engine.
// See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#page-rules
// for full details.
func (r *PageruleService) Edit(ctx context.Context, pageruleID string, params PageruleEditParams, opts ...option.RequestOption) (res *PageruleEditResponse, err error) {
	var env PageruleEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if pageruleID == "" {
		err = errors.New("missing required pagerule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/pagerules/%s", params.ZoneID, pageruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the details of a Page Rule.
//
// Deprecated: The Page Rules API is deprecated in favour of the Ruleset Engine.
// See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#page-rules
// for full details.
func (r *PageruleService) Get(ctx context.Context, pageruleID string, query PageruleGetParams, opts ...option.RequestOption) (res *PageruleGetResponse, err error) {
	var env PageruleGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if pageruleID == "" {
		err = errors.New("missing required pagerule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/pagerules/%s", query.ZoneID, pageruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// URL target.
type Target struct {
	// String constraint.
	Constraint TargetConstraint `json:"constraint"`
	// A target based on the URL of the request.
	Target TargetTarget `json:"target"`
	JSON   targetJSON   `json:"-"`
}

// targetJSON contains the JSON metadata for the struct [Target]
type targetJSON struct {
	Constraint  apijson.Field
	Target      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Target) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r targetJSON) RawJSON() string {
	return r.raw
}

// String constraint.
type TargetConstraint struct {
	// The matches operator can use asterisks and pipes as wildcard and 'or' operators.
	Operator TargetConstraintOperator `json:"operator,required"`
	// The URL pattern to match against the current request. The pattern may contain up
	// to four asterisks ('\*') as placeholders.
	Value string               `json:"value,required"`
	JSON  targetConstraintJSON `json:"-"`
}

// targetConstraintJSON contains the JSON metadata for the struct
// [TargetConstraint]
type targetConstraintJSON struct {
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TargetConstraint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r targetConstraintJSON) RawJSON() string {
	return r.raw
}

// The matches operator can use asterisks and pipes as wildcard and 'or' operators.
type TargetConstraintOperator string

const (
	TargetConstraintOperatorMatches    TargetConstraintOperator = "matches"
	TargetConstraintOperatorContains   TargetConstraintOperator = "contains"
	TargetConstraintOperatorEquals     TargetConstraintOperator = "equals"
	TargetConstraintOperatorNotEqual   TargetConstraintOperator = "not_equal"
	TargetConstraintOperatorNotContain TargetConstraintOperator = "not_contain"
)

func (r TargetConstraintOperator) IsKnown() bool {
	switch r {
	case TargetConstraintOperatorMatches, TargetConstraintOperatorContains, TargetConstraintOperatorEquals, TargetConstraintOperatorNotEqual, TargetConstraintOperatorNotContain:
		return true
	}
	return false
}

// A target based on the URL of the request.
type TargetTarget string

const (
	TargetTargetURL TargetTarget = "url"
)

func (r TargetTarget) IsKnown() bool {
	switch r {
	case TargetTargetURL:
		return true
	}
	return false
}

// URL target.
type TargetParam struct {
	// String constraint.
	Constraint param.Field[TargetConstraintParam] `json:"constraint"`
	// A target based on the URL of the request.
	Target param.Field[TargetTarget] `json:"target"`
}

func (r TargetParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// String constraint.
type TargetConstraintParam struct {
	// The matches operator can use asterisks and pipes as wildcard and 'or' operators.
	Operator param.Field[TargetConstraintOperator] `json:"operator,required"`
	// The URL pattern to match against the current request. The pattern may contain up
	// to four asterisks ('\*') as placeholders.
	Value param.Field[string] `json:"value,required"`
}

func (r TargetConstraintParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PageruleNewResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// The set of actions to perform if the targets of this rule match the request.
	// Actions can redirect to another URL or override settings, but not both.
	Actions []PageruleNewResponseAction `json:"actions,required"`
	// The timestamp of when the Page Rule was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The timestamp of when the Page Rule was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The priority of the rule, used to define which Page Rule is processed over
	// another. A higher number indicates a higher priority. For example, if you have a
	// catch-all Page Rule (rule A: `/images/*`) but want a more specific Page Rule to
	// take precedence (rule B: `/images/special/*`), specify a higher priority for
	// rule B so it overrides rule A.
	Priority int64 `json:"priority,required"`
	// The status of the Page Rule.
	Status PageruleNewResponseStatus `json:"status,required"`
	// The rule targets to evaluate on each request.
	Targets []Target                `json:"targets,required"`
	JSON    pageruleNewResponseJSON `json:"-"`
}

// pageruleNewResponseJSON contains the JSON metadata for the struct
// [PageruleNewResponse]
type pageruleNewResponseJSON struct {
	ID          apijson.Field
	Actions     apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	Priority    apijson.Field
	Status      apijson.Field
	Targets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleNewResponseJSON) RawJSON() string {
	return r.raw
}

type PageruleNewResponseAction struct {
	// Redirects one URL to another using an `HTTP 301/302` redirect. Refer to
	// [Wildcard matching and referencing](https://developers.cloudflare.com/rules/page-rules/reference/wildcard-matching/).
	ID    PageruleNewResponseActionsID    `json:"id"`
	Value PageruleNewResponseActionsValue `json:"value"`
	JSON  pageruleNewResponseActionJSON   `json:"-"`
}

// pageruleNewResponseActionJSON contains the JSON metadata for the struct
// [PageruleNewResponseAction]
type pageruleNewResponseActionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleNewResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleNewResponseActionJSON) RawJSON() string {
	return r.raw
}

// Redirects one URL to another using an `HTTP 301/302` redirect. Refer to
// [Wildcard matching and referencing](https://developers.cloudflare.com/rules/page-rules/reference/wildcard-matching/).
type PageruleNewResponseActionsID string

const (
	PageruleNewResponseActionsIDForwardingURL PageruleNewResponseActionsID = "forwarding_url"
)

func (r PageruleNewResponseActionsID) IsKnown() bool {
	switch r {
	case PageruleNewResponseActionsIDForwardingURL:
		return true
	}
	return false
}

type PageruleNewResponseActionsValue struct {
	// The status code to use for the URL redirect. 301 is a permanent redirect. 302 is
	// a temporary redirect.
	StatusCode PageruleNewResponseActionsValueStatusCode `json:"status_code"`
	// The URL to redirect the request to. Notes: ${num} refers to the position of '\*'
	// in the constraint value.
	URL  string                              `json:"url"`
	JSON pageruleNewResponseActionsValueJSON `json:"-"`
}

// pageruleNewResponseActionsValueJSON contains the JSON metadata for the struct
// [PageruleNewResponseActionsValue]
type pageruleNewResponseActionsValueJSON struct {
	StatusCode  apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleNewResponseActionsValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleNewResponseActionsValueJSON) RawJSON() string {
	return r.raw
}

// The status code to use for the URL redirect. 301 is a permanent redirect. 302 is
// a temporary redirect.
type PageruleNewResponseActionsValueStatusCode int64

const (
	PageruleNewResponseActionsValueStatusCode301 PageruleNewResponseActionsValueStatusCode = 301
	PageruleNewResponseActionsValueStatusCode302 PageruleNewResponseActionsValueStatusCode = 302
)

func (r PageruleNewResponseActionsValueStatusCode) IsKnown() bool {
	switch r {
	case PageruleNewResponseActionsValueStatusCode301, PageruleNewResponseActionsValueStatusCode302:
		return true
	}
	return false
}

// The status of the Page Rule.
type PageruleNewResponseStatus string

const (
	PageruleNewResponseStatusActive   PageruleNewResponseStatus = "active"
	PageruleNewResponseStatusDisabled PageruleNewResponseStatus = "disabled"
)

func (r PageruleNewResponseStatus) IsKnown() bool {
	switch r {
	case PageruleNewResponseStatusActive, PageruleNewResponseStatusDisabled:
		return true
	}
	return false
}

type PageruleUpdateResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// The set of actions to perform if the targets of this rule match the request.
	// Actions can redirect to another URL or override settings, but not both.
	Actions []PageruleUpdateResponseAction `json:"actions,required"`
	// The timestamp of when the Page Rule was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The timestamp of when the Page Rule was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The priority of the rule, used to define which Page Rule is processed over
	// another. A higher number indicates a higher priority. For example, if you have a
	// catch-all Page Rule (rule A: `/images/*`) but want a more specific Page Rule to
	// take precedence (rule B: `/images/special/*`), specify a higher priority for
	// rule B so it overrides rule A.
	Priority int64 `json:"priority,required"`
	// The status of the Page Rule.
	Status PageruleUpdateResponseStatus `json:"status,required"`
	// The rule targets to evaluate on each request.
	Targets []Target                   `json:"targets,required"`
	JSON    pageruleUpdateResponseJSON `json:"-"`
}

// pageruleUpdateResponseJSON contains the JSON metadata for the struct
// [PageruleUpdateResponse]
type pageruleUpdateResponseJSON struct {
	ID          apijson.Field
	Actions     apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	Priority    apijson.Field
	Status      apijson.Field
	Targets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type PageruleUpdateResponseAction struct {
	// Redirects one URL to another using an `HTTP 301/302` redirect. Refer to
	// [Wildcard matching and referencing](https://developers.cloudflare.com/rules/page-rules/reference/wildcard-matching/).
	ID    PageruleUpdateResponseActionsID    `json:"id"`
	Value PageruleUpdateResponseActionsValue `json:"value"`
	JSON  pageruleUpdateResponseActionJSON   `json:"-"`
}

// pageruleUpdateResponseActionJSON contains the JSON metadata for the struct
// [PageruleUpdateResponseAction]
type pageruleUpdateResponseActionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleUpdateResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleUpdateResponseActionJSON) RawJSON() string {
	return r.raw
}

// Redirects one URL to another using an `HTTP 301/302` redirect. Refer to
// [Wildcard matching and referencing](https://developers.cloudflare.com/rules/page-rules/reference/wildcard-matching/).
type PageruleUpdateResponseActionsID string

const (
	PageruleUpdateResponseActionsIDForwardingURL PageruleUpdateResponseActionsID = "forwarding_url"
)

func (r PageruleUpdateResponseActionsID) IsKnown() bool {
	switch r {
	case PageruleUpdateResponseActionsIDForwardingURL:
		return true
	}
	return false
}

type PageruleUpdateResponseActionsValue struct {
	// The status code to use for the URL redirect. 301 is a permanent redirect. 302 is
	// a temporary redirect.
	StatusCode PageruleUpdateResponseActionsValueStatusCode `json:"status_code"`
	// The URL to redirect the request to. Notes: ${num} refers to the position of '\*'
	// in the constraint value.
	URL  string                                 `json:"url"`
	JSON pageruleUpdateResponseActionsValueJSON `json:"-"`
}

// pageruleUpdateResponseActionsValueJSON contains the JSON metadata for the struct
// [PageruleUpdateResponseActionsValue]
type pageruleUpdateResponseActionsValueJSON struct {
	StatusCode  apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleUpdateResponseActionsValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleUpdateResponseActionsValueJSON) RawJSON() string {
	return r.raw
}

// The status code to use for the URL redirect. 301 is a permanent redirect. 302 is
// a temporary redirect.
type PageruleUpdateResponseActionsValueStatusCode int64

const (
	PageruleUpdateResponseActionsValueStatusCode301 PageruleUpdateResponseActionsValueStatusCode = 301
	PageruleUpdateResponseActionsValueStatusCode302 PageruleUpdateResponseActionsValueStatusCode = 302
)

func (r PageruleUpdateResponseActionsValueStatusCode) IsKnown() bool {
	switch r {
	case PageruleUpdateResponseActionsValueStatusCode301, PageruleUpdateResponseActionsValueStatusCode302:
		return true
	}
	return false
}

// The status of the Page Rule.
type PageruleUpdateResponseStatus string

const (
	PageruleUpdateResponseStatusActive   PageruleUpdateResponseStatus = "active"
	PageruleUpdateResponseStatusDisabled PageruleUpdateResponseStatus = "disabled"
)

func (r PageruleUpdateResponseStatus) IsKnown() bool {
	switch r {
	case PageruleUpdateResponseStatusActive, PageruleUpdateResponseStatusDisabled:
		return true
	}
	return false
}

type PageruleListResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// The set of actions to perform if the targets of this rule match the request.
	// Actions can redirect to another URL or override settings, but not both.
	Actions []PageruleListResponseAction `json:"actions,required"`
	// The timestamp of when the Page Rule was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The timestamp of when the Page Rule was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The priority of the rule, used to define which Page Rule is processed over
	// another. A higher number indicates a higher priority. For example, if you have a
	// catch-all Page Rule (rule A: `/images/*`) but want a more specific Page Rule to
	// take precedence (rule B: `/images/special/*`), specify a higher priority for
	// rule B so it overrides rule A.
	Priority int64 `json:"priority,required"`
	// The status of the Page Rule.
	Status PageruleListResponseStatus `json:"status,required"`
	// The rule targets to evaluate on each request.
	Targets []Target                 `json:"targets,required"`
	JSON    pageruleListResponseJSON `json:"-"`
}

// pageruleListResponseJSON contains the JSON metadata for the struct
// [PageruleListResponse]
type pageruleListResponseJSON struct {
	ID          apijson.Field
	Actions     apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	Priority    apijson.Field
	Status      apijson.Field
	Targets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleListResponseJSON) RawJSON() string {
	return r.raw
}

type PageruleListResponseAction struct {
	// Redirects one URL to another using an `HTTP 301/302` redirect. Refer to
	// [Wildcard matching and referencing](https://developers.cloudflare.com/rules/page-rules/reference/wildcard-matching/).
	ID    PageruleListResponseActionsID    `json:"id"`
	Value PageruleListResponseActionsValue `json:"value"`
	JSON  pageruleListResponseActionJSON   `json:"-"`
}

// pageruleListResponseActionJSON contains the JSON metadata for the struct
// [PageruleListResponseAction]
type pageruleListResponseActionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleListResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleListResponseActionJSON) RawJSON() string {
	return r.raw
}

// Redirects one URL to another using an `HTTP 301/302` redirect. Refer to
// [Wildcard matching and referencing](https://developers.cloudflare.com/rules/page-rules/reference/wildcard-matching/).
type PageruleListResponseActionsID string

const (
	PageruleListResponseActionsIDForwardingURL PageruleListResponseActionsID = "forwarding_url"
)

func (r PageruleListResponseActionsID) IsKnown() bool {
	switch r {
	case PageruleListResponseActionsIDForwardingURL:
		return true
	}
	return false
}

type PageruleListResponseActionsValue struct {
	// The status code to use for the URL redirect. 301 is a permanent redirect. 302 is
	// a temporary redirect.
	StatusCode PageruleListResponseActionsValueStatusCode `json:"status_code"`
	// The URL to redirect the request to. Notes: ${num} refers to the position of '\*'
	// in the constraint value.
	URL  string                               `json:"url"`
	JSON pageruleListResponseActionsValueJSON `json:"-"`
}

// pageruleListResponseActionsValueJSON contains the JSON metadata for the struct
// [PageruleListResponseActionsValue]
type pageruleListResponseActionsValueJSON struct {
	StatusCode  apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleListResponseActionsValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleListResponseActionsValueJSON) RawJSON() string {
	return r.raw
}

// The status code to use for the URL redirect. 301 is a permanent redirect. 302 is
// a temporary redirect.
type PageruleListResponseActionsValueStatusCode int64

const (
	PageruleListResponseActionsValueStatusCode301 PageruleListResponseActionsValueStatusCode = 301
	PageruleListResponseActionsValueStatusCode302 PageruleListResponseActionsValueStatusCode = 302
)

func (r PageruleListResponseActionsValueStatusCode) IsKnown() bool {
	switch r {
	case PageruleListResponseActionsValueStatusCode301, PageruleListResponseActionsValueStatusCode302:
		return true
	}
	return false
}

// The status of the Page Rule.
type PageruleListResponseStatus string

const (
	PageruleListResponseStatusActive   PageruleListResponseStatus = "active"
	PageruleListResponseStatusDisabled PageruleListResponseStatus = "disabled"
)

func (r PageruleListResponseStatus) IsKnown() bool {
	switch r {
	case PageruleListResponseStatusActive, PageruleListResponseStatusDisabled:
		return true
	}
	return false
}

type PageruleDeleteResponse struct {
	// Identifier
	ID   string                     `json:"id,required"`
	JSON pageruleDeleteResponseJSON `json:"-"`
}

// pageruleDeleteResponseJSON contains the JSON metadata for the struct
// [PageruleDeleteResponse]
type pageruleDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type PageruleEditResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// The set of actions to perform if the targets of this rule match the request.
	// Actions can redirect to another URL or override settings, but not both.
	Actions []PageruleEditResponseAction `json:"actions,required"`
	// The timestamp of when the Page Rule was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The timestamp of when the Page Rule was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The priority of the rule, used to define which Page Rule is processed over
	// another. A higher number indicates a higher priority. For example, if you have a
	// catch-all Page Rule (rule A: `/images/*`) but want a more specific Page Rule to
	// take precedence (rule B: `/images/special/*`), specify a higher priority for
	// rule B so it overrides rule A.
	Priority int64 `json:"priority,required"`
	// The status of the Page Rule.
	Status PageruleEditResponseStatus `json:"status,required"`
	// The rule targets to evaluate on each request.
	Targets []Target                 `json:"targets,required"`
	JSON    pageruleEditResponseJSON `json:"-"`
}

// pageruleEditResponseJSON contains the JSON metadata for the struct
// [PageruleEditResponse]
type pageruleEditResponseJSON struct {
	ID          apijson.Field
	Actions     apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	Priority    apijson.Field
	Status      apijson.Field
	Targets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleEditResponseJSON) RawJSON() string {
	return r.raw
}

type PageruleEditResponseAction struct {
	// Redirects one URL to another using an `HTTP 301/302` redirect. Refer to
	// [Wildcard matching and referencing](https://developers.cloudflare.com/rules/page-rules/reference/wildcard-matching/).
	ID    PageruleEditResponseActionsID    `json:"id"`
	Value PageruleEditResponseActionsValue `json:"value"`
	JSON  pageruleEditResponseActionJSON   `json:"-"`
}

// pageruleEditResponseActionJSON contains the JSON metadata for the struct
// [PageruleEditResponseAction]
type pageruleEditResponseActionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleEditResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleEditResponseActionJSON) RawJSON() string {
	return r.raw
}

// Redirects one URL to another using an `HTTP 301/302` redirect. Refer to
// [Wildcard matching and referencing](https://developers.cloudflare.com/rules/page-rules/reference/wildcard-matching/).
type PageruleEditResponseActionsID string

const (
	PageruleEditResponseActionsIDForwardingURL PageruleEditResponseActionsID = "forwarding_url"
)

func (r PageruleEditResponseActionsID) IsKnown() bool {
	switch r {
	case PageruleEditResponseActionsIDForwardingURL:
		return true
	}
	return false
}

type PageruleEditResponseActionsValue struct {
	// The status code to use for the URL redirect. 301 is a permanent redirect. 302 is
	// a temporary redirect.
	StatusCode PageruleEditResponseActionsValueStatusCode `json:"status_code"`
	// The URL to redirect the request to. Notes: ${num} refers to the position of '\*'
	// in the constraint value.
	URL  string                               `json:"url"`
	JSON pageruleEditResponseActionsValueJSON `json:"-"`
}

// pageruleEditResponseActionsValueJSON contains the JSON metadata for the struct
// [PageruleEditResponseActionsValue]
type pageruleEditResponseActionsValueJSON struct {
	StatusCode  apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleEditResponseActionsValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleEditResponseActionsValueJSON) RawJSON() string {
	return r.raw
}

// The status code to use for the URL redirect. 301 is a permanent redirect. 302 is
// a temporary redirect.
type PageruleEditResponseActionsValueStatusCode int64

const (
	PageruleEditResponseActionsValueStatusCode301 PageruleEditResponseActionsValueStatusCode = 301
	PageruleEditResponseActionsValueStatusCode302 PageruleEditResponseActionsValueStatusCode = 302
)

func (r PageruleEditResponseActionsValueStatusCode) IsKnown() bool {
	switch r {
	case PageruleEditResponseActionsValueStatusCode301, PageruleEditResponseActionsValueStatusCode302:
		return true
	}
	return false
}

// The status of the Page Rule.
type PageruleEditResponseStatus string

const (
	PageruleEditResponseStatusActive   PageruleEditResponseStatus = "active"
	PageruleEditResponseStatusDisabled PageruleEditResponseStatus = "disabled"
)

func (r PageruleEditResponseStatus) IsKnown() bool {
	switch r {
	case PageruleEditResponseStatusActive, PageruleEditResponseStatusDisabled:
		return true
	}
	return false
}

type PageruleGetResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// The set of actions to perform if the targets of this rule match the request.
	// Actions can redirect to another URL or override settings, but not both.
	Actions []PageruleGetResponseAction `json:"actions,required"`
	// The timestamp of when the Page Rule was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The timestamp of when the Page Rule was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The priority of the rule, used to define which Page Rule is processed over
	// another. A higher number indicates a higher priority. For example, if you have a
	// catch-all Page Rule (rule A: `/images/*`) but want a more specific Page Rule to
	// take precedence (rule B: `/images/special/*`), specify a higher priority for
	// rule B so it overrides rule A.
	Priority int64 `json:"priority,required"`
	// The status of the Page Rule.
	Status PageruleGetResponseStatus `json:"status,required"`
	// The rule targets to evaluate on each request.
	Targets []Target                `json:"targets,required"`
	JSON    pageruleGetResponseJSON `json:"-"`
}

// pageruleGetResponseJSON contains the JSON metadata for the struct
// [PageruleGetResponse]
type pageruleGetResponseJSON struct {
	ID          apijson.Field
	Actions     apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	Priority    apijson.Field
	Status      apijson.Field
	Targets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleGetResponseJSON) RawJSON() string {
	return r.raw
}

type PageruleGetResponseAction struct {
	// Redirects one URL to another using an `HTTP 301/302` redirect. Refer to
	// [Wildcard matching and referencing](https://developers.cloudflare.com/rules/page-rules/reference/wildcard-matching/).
	ID    PageruleGetResponseActionsID    `json:"id"`
	Value PageruleGetResponseActionsValue `json:"value"`
	JSON  pageruleGetResponseActionJSON   `json:"-"`
}

// pageruleGetResponseActionJSON contains the JSON metadata for the struct
// [PageruleGetResponseAction]
type pageruleGetResponseActionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleGetResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleGetResponseActionJSON) RawJSON() string {
	return r.raw
}

// Redirects one URL to another using an `HTTP 301/302` redirect. Refer to
// [Wildcard matching and referencing](https://developers.cloudflare.com/rules/page-rules/reference/wildcard-matching/).
type PageruleGetResponseActionsID string

const (
	PageruleGetResponseActionsIDForwardingURL PageruleGetResponseActionsID = "forwarding_url"
)

func (r PageruleGetResponseActionsID) IsKnown() bool {
	switch r {
	case PageruleGetResponseActionsIDForwardingURL:
		return true
	}
	return false
}

type PageruleGetResponseActionsValue struct {
	// The status code to use for the URL redirect. 301 is a permanent redirect. 302 is
	// a temporary redirect.
	StatusCode PageruleGetResponseActionsValueStatusCode `json:"status_code"`
	// The URL to redirect the request to. Notes: ${num} refers to the position of '\*'
	// in the constraint value.
	URL  string                              `json:"url"`
	JSON pageruleGetResponseActionsValueJSON `json:"-"`
}

// pageruleGetResponseActionsValueJSON contains the JSON metadata for the struct
// [PageruleGetResponseActionsValue]
type pageruleGetResponseActionsValueJSON struct {
	StatusCode  apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleGetResponseActionsValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleGetResponseActionsValueJSON) RawJSON() string {
	return r.raw
}

// The status code to use for the URL redirect. 301 is a permanent redirect. 302 is
// a temporary redirect.
type PageruleGetResponseActionsValueStatusCode int64

const (
	PageruleGetResponseActionsValueStatusCode301 PageruleGetResponseActionsValueStatusCode = 301
	PageruleGetResponseActionsValueStatusCode302 PageruleGetResponseActionsValueStatusCode = 302
)

func (r PageruleGetResponseActionsValueStatusCode) IsKnown() bool {
	switch r {
	case PageruleGetResponseActionsValueStatusCode301, PageruleGetResponseActionsValueStatusCode302:
		return true
	}
	return false
}

// The status of the Page Rule.
type PageruleGetResponseStatus string

const (
	PageruleGetResponseStatusActive   PageruleGetResponseStatus = "active"
	PageruleGetResponseStatusDisabled PageruleGetResponseStatus = "disabled"
)

func (r PageruleGetResponseStatus) IsKnown() bool {
	switch r {
	case PageruleGetResponseStatusActive, PageruleGetResponseStatusDisabled:
		return true
	}
	return false
}

type PageruleNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The set of actions to perform if the targets of this rule match the request.
	// Actions can redirect to another URL or override settings, but not both.
	Actions param.Field[[]PageruleNewParamsAction] `json:"actions,required"`
	// The rule targets to evaluate on each request.
	Targets param.Field[[]TargetParam] `json:"targets,required"`
	// The priority of the rule, used to define which Page Rule is processed over
	// another. A higher number indicates a higher priority. For example, if you have a
	// catch-all Page Rule (rule A: `/images/*`) but want a more specific Page Rule to
	// take precedence (rule B: `/images/special/*`), specify a higher priority for
	// rule B so it overrides rule A.
	Priority param.Field[int64] `json:"priority"`
	// The status of the Page Rule.
	Status param.Field[PageruleNewParamsStatus] `json:"status"`
}

func (r PageruleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PageruleNewParamsAction struct {
	// Redirects one URL to another using an `HTTP 301/302` redirect. Refer to
	// [Wildcard matching and referencing](https://developers.cloudflare.com/rules/page-rules/reference/wildcard-matching/).
	ID    param.Field[PageruleNewParamsActionsID]    `json:"id"`
	Value param.Field[PageruleNewParamsActionsValue] `json:"value"`
}

func (r PageruleNewParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Redirects one URL to another using an `HTTP 301/302` redirect. Refer to
// [Wildcard matching and referencing](https://developers.cloudflare.com/rules/page-rules/reference/wildcard-matching/).
type PageruleNewParamsActionsID string

const (
	PageruleNewParamsActionsIDForwardingURL PageruleNewParamsActionsID = "forwarding_url"
)

func (r PageruleNewParamsActionsID) IsKnown() bool {
	switch r {
	case PageruleNewParamsActionsIDForwardingURL:
		return true
	}
	return false
}

type PageruleNewParamsActionsValue struct {
	// The status code to use for the URL redirect. 301 is a permanent redirect. 302 is
	// a temporary redirect.
	StatusCode param.Field[PageruleNewParamsActionsValueStatusCode] `json:"status_code"`
	// The URL to redirect the request to. Notes: ${num} refers to the position of '\*'
	// in the constraint value.
	URL param.Field[string] `json:"url"`
}

func (r PageruleNewParamsActionsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The status code to use for the URL redirect. 301 is a permanent redirect. 302 is
// a temporary redirect.
type PageruleNewParamsActionsValueStatusCode int64

const (
	PageruleNewParamsActionsValueStatusCode301 PageruleNewParamsActionsValueStatusCode = 301
	PageruleNewParamsActionsValueStatusCode302 PageruleNewParamsActionsValueStatusCode = 302
)

func (r PageruleNewParamsActionsValueStatusCode) IsKnown() bool {
	switch r {
	case PageruleNewParamsActionsValueStatusCode301, PageruleNewParamsActionsValueStatusCode302:
		return true
	}
	return false
}

// The status of the Page Rule.
type PageruleNewParamsStatus string

const (
	PageruleNewParamsStatusActive   PageruleNewParamsStatus = "active"
	PageruleNewParamsStatusDisabled PageruleNewParamsStatus = "disabled"
)

func (r PageruleNewParamsStatus) IsKnown() bool {
	switch r {
	case PageruleNewParamsStatusActive, PageruleNewParamsStatusDisabled:
		return true
	}
	return false
}

type PageruleNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PageruleNewResponseEnvelopeSuccess `json:"success,required"`
	Result  PageruleNewResponse                `json:"result"`
	JSON    pageruleNewResponseEnvelopeJSON    `json:"-"`
}

// pageruleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [PageruleNewResponseEnvelope]
type pageruleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageruleNewResponseEnvelopeSuccess bool

const (
	PageruleNewResponseEnvelopeSuccessTrue PageruleNewResponseEnvelopeSuccess = true
)

func (r PageruleNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PageruleNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PageruleUpdateParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The set of actions to perform if the targets of this rule match the request.
	// Actions can redirect to another URL or override settings, but not both.
	Actions param.Field[[]PageruleUpdateParamsAction] `json:"actions,required"`
	// The rule targets to evaluate on each request.
	Targets param.Field[[]TargetParam] `json:"targets,required"`
	// The priority of the rule, used to define which Page Rule is processed over
	// another. A higher number indicates a higher priority. For example, if you have a
	// catch-all Page Rule (rule A: `/images/*`) but want a more specific Page Rule to
	// take precedence (rule B: `/images/special/*`), specify a higher priority for
	// rule B so it overrides rule A.
	Priority param.Field[int64] `json:"priority"`
	// The status of the Page Rule.
	Status param.Field[PageruleUpdateParamsStatus] `json:"status"`
}

func (r PageruleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PageruleUpdateParamsAction struct {
	// Redirects one URL to another using an `HTTP 301/302` redirect. Refer to
	// [Wildcard matching and referencing](https://developers.cloudflare.com/rules/page-rules/reference/wildcard-matching/).
	ID    param.Field[PageruleUpdateParamsActionsID]    `json:"id"`
	Value param.Field[PageruleUpdateParamsActionsValue] `json:"value"`
}

func (r PageruleUpdateParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Redirects one URL to another using an `HTTP 301/302` redirect. Refer to
// [Wildcard matching and referencing](https://developers.cloudflare.com/rules/page-rules/reference/wildcard-matching/).
type PageruleUpdateParamsActionsID string

const (
	PageruleUpdateParamsActionsIDForwardingURL PageruleUpdateParamsActionsID = "forwarding_url"
)

func (r PageruleUpdateParamsActionsID) IsKnown() bool {
	switch r {
	case PageruleUpdateParamsActionsIDForwardingURL:
		return true
	}
	return false
}

type PageruleUpdateParamsActionsValue struct {
	// The status code to use for the URL redirect. 301 is a permanent redirect. 302 is
	// a temporary redirect.
	StatusCode param.Field[PageruleUpdateParamsActionsValueStatusCode] `json:"status_code"`
	// The URL to redirect the request to. Notes: ${num} refers to the position of '\*'
	// in the constraint value.
	URL param.Field[string] `json:"url"`
}

func (r PageruleUpdateParamsActionsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The status code to use for the URL redirect. 301 is a permanent redirect. 302 is
// a temporary redirect.
type PageruleUpdateParamsActionsValueStatusCode int64

const (
	PageruleUpdateParamsActionsValueStatusCode301 PageruleUpdateParamsActionsValueStatusCode = 301
	PageruleUpdateParamsActionsValueStatusCode302 PageruleUpdateParamsActionsValueStatusCode = 302
)

func (r PageruleUpdateParamsActionsValueStatusCode) IsKnown() bool {
	switch r {
	case PageruleUpdateParamsActionsValueStatusCode301, PageruleUpdateParamsActionsValueStatusCode302:
		return true
	}
	return false
}

// The status of the Page Rule.
type PageruleUpdateParamsStatus string

const (
	PageruleUpdateParamsStatusActive   PageruleUpdateParamsStatus = "active"
	PageruleUpdateParamsStatusDisabled PageruleUpdateParamsStatus = "disabled"
)

func (r PageruleUpdateParamsStatus) IsKnown() bool {
	switch r {
	case PageruleUpdateParamsStatusActive, PageruleUpdateParamsStatusDisabled:
		return true
	}
	return false
}

type PageruleUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PageruleUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  PageruleUpdateResponse                `json:"result"`
	JSON    pageruleUpdateResponseEnvelopeJSON    `json:"-"`
}

// pageruleUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [PageruleUpdateResponseEnvelope]
type pageruleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageruleUpdateResponseEnvelopeSuccess bool

const (
	PageruleUpdateResponseEnvelopeSuccessTrue PageruleUpdateResponseEnvelopeSuccess = true
)

func (r PageruleUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PageruleUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PageruleListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The direction used to sort returned Page Rules.
	Direction param.Field[PageruleListParamsDirection] `query:"direction"`
	// When set to `all`, all the search requirements must match. When set to `any`,
	// only one of the search requirements has to match.
	Match param.Field[PageruleListParamsMatch] `query:"match"`
	// The field used to sort returned Page Rules.
	Order param.Field[PageruleListParamsOrder] `query:"order"`
	// The status of the Page Rule.
	Status param.Field[PageruleListParamsStatus] `query:"status"`
}

// URLQuery serializes [PageruleListParams]'s query parameters as `url.Values`.
func (r PageruleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The direction used to sort returned Page Rules.
type PageruleListParamsDirection string

const (
	PageruleListParamsDirectionAsc  PageruleListParamsDirection = "asc"
	PageruleListParamsDirectionDesc PageruleListParamsDirection = "desc"
)

func (r PageruleListParamsDirection) IsKnown() bool {
	switch r {
	case PageruleListParamsDirectionAsc, PageruleListParamsDirectionDesc:
		return true
	}
	return false
}

// When set to `all`, all the search requirements must match. When set to `any`,
// only one of the search requirements has to match.
type PageruleListParamsMatch string

const (
	PageruleListParamsMatchAny PageruleListParamsMatch = "any"
	PageruleListParamsMatchAll PageruleListParamsMatch = "all"
)

func (r PageruleListParamsMatch) IsKnown() bool {
	switch r {
	case PageruleListParamsMatchAny, PageruleListParamsMatchAll:
		return true
	}
	return false
}

// The field used to sort returned Page Rules.
type PageruleListParamsOrder string

const (
	PageruleListParamsOrderStatus   PageruleListParamsOrder = "status"
	PageruleListParamsOrderPriority PageruleListParamsOrder = "priority"
)

func (r PageruleListParamsOrder) IsKnown() bool {
	switch r {
	case PageruleListParamsOrderStatus, PageruleListParamsOrderPriority:
		return true
	}
	return false
}

// The status of the Page Rule.
type PageruleListParamsStatus string

const (
	PageruleListParamsStatusActive   PageruleListParamsStatus = "active"
	PageruleListParamsStatusDisabled PageruleListParamsStatus = "disabled"
)

func (r PageruleListParamsStatus) IsKnown() bool {
	switch r {
	case PageruleListParamsStatusActive, PageruleListParamsStatusDisabled:
		return true
	}
	return false
}

type PageruleListResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PageruleListResponseEnvelopeSuccess `json:"success,required"`
	Result  []PageruleListResponse              `json:"result"`
	JSON    pageruleListResponseEnvelopeJSON    `json:"-"`
}

// pageruleListResponseEnvelopeJSON contains the JSON metadata for the struct
// [PageruleListResponseEnvelope]
type pageruleListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageruleListResponseEnvelopeSuccess bool

const (
	PageruleListResponseEnvelopeSuccessTrue PageruleListResponseEnvelopeSuccess = true
)

func (r PageruleListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PageruleListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PageruleDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type PageruleDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PageruleDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  PageruleDeleteResponse                `json:"result,nullable"`
	JSON    pageruleDeleteResponseEnvelopeJSON    `json:"-"`
}

// pageruleDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [PageruleDeleteResponseEnvelope]
type pageruleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageruleDeleteResponseEnvelopeSuccess bool

const (
	PageruleDeleteResponseEnvelopeSuccessTrue PageruleDeleteResponseEnvelopeSuccess = true
)

func (r PageruleDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PageruleDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PageruleEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The set of actions to perform if the targets of this rule match the request.
	// Actions can redirect to another URL or override settings, but not both.
	Actions param.Field[[]PageruleEditParamsAction] `json:"actions"`
	// The priority of the rule, used to define which Page Rule is processed over
	// another. A higher number indicates a higher priority. For example, if you have a
	// catch-all Page Rule (rule A: `/images/*`) but want a more specific Page Rule to
	// take precedence (rule B: `/images/special/*`), specify a higher priority for
	// rule B so it overrides rule A.
	Priority param.Field[int64] `json:"priority"`
	// The status of the Page Rule.
	Status param.Field[PageruleEditParamsStatus] `json:"status"`
	// The rule targets to evaluate on each request.
	Targets param.Field[[]TargetParam] `json:"targets"`
}

func (r PageruleEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PageruleEditParamsAction struct {
	// Redirects one URL to another using an `HTTP 301/302` redirect. Refer to
	// [Wildcard matching and referencing](https://developers.cloudflare.com/rules/page-rules/reference/wildcard-matching/).
	ID    param.Field[PageruleEditParamsActionsID]    `json:"id"`
	Value param.Field[PageruleEditParamsActionsValue] `json:"value"`
}

func (r PageruleEditParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Redirects one URL to another using an `HTTP 301/302` redirect. Refer to
// [Wildcard matching and referencing](https://developers.cloudflare.com/rules/page-rules/reference/wildcard-matching/).
type PageruleEditParamsActionsID string

const (
	PageruleEditParamsActionsIDForwardingURL PageruleEditParamsActionsID = "forwarding_url"
)

func (r PageruleEditParamsActionsID) IsKnown() bool {
	switch r {
	case PageruleEditParamsActionsIDForwardingURL:
		return true
	}
	return false
}

type PageruleEditParamsActionsValue struct {
	// The status code to use for the URL redirect. 301 is a permanent redirect. 302 is
	// a temporary redirect.
	StatusCode param.Field[PageruleEditParamsActionsValueStatusCode] `json:"status_code"`
	// The URL to redirect the request to. Notes: ${num} refers to the position of '\*'
	// in the constraint value.
	URL param.Field[string] `json:"url"`
}

func (r PageruleEditParamsActionsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The status code to use for the URL redirect. 301 is a permanent redirect. 302 is
// a temporary redirect.
type PageruleEditParamsActionsValueStatusCode int64

const (
	PageruleEditParamsActionsValueStatusCode301 PageruleEditParamsActionsValueStatusCode = 301
	PageruleEditParamsActionsValueStatusCode302 PageruleEditParamsActionsValueStatusCode = 302
)

func (r PageruleEditParamsActionsValueStatusCode) IsKnown() bool {
	switch r {
	case PageruleEditParamsActionsValueStatusCode301, PageruleEditParamsActionsValueStatusCode302:
		return true
	}
	return false
}

// The status of the Page Rule.
type PageruleEditParamsStatus string

const (
	PageruleEditParamsStatusActive   PageruleEditParamsStatus = "active"
	PageruleEditParamsStatusDisabled PageruleEditParamsStatus = "disabled"
)

func (r PageruleEditParamsStatus) IsKnown() bool {
	switch r {
	case PageruleEditParamsStatusActive, PageruleEditParamsStatusDisabled:
		return true
	}
	return false
}

type PageruleEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PageruleEditResponseEnvelopeSuccess `json:"success,required"`
	Result  PageruleEditResponse                `json:"result"`
	JSON    pageruleEditResponseEnvelopeJSON    `json:"-"`
}

// pageruleEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [PageruleEditResponseEnvelope]
type pageruleEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageruleEditResponseEnvelopeSuccess bool

const (
	PageruleEditResponseEnvelopeSuccessTrue PageruleEditResponseEnvelopeSuccess = true
)

func (r PageruleEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PageruleEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PageruleGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type PageruleGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PageruleGetResponseEnvelopeSuccess `json:"success,required"`
	Result  PageruleGetResponse                `json:"result"`
	JSON    pageruleGetResponseEnvelopeJSON    `json:"-"`
}

// pageruleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PageruleGetResponseEnvelope]
type pageruleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageruleGetResponseEnvelopeSuccess bool

const (
	PageruleGetResponseEnvelopeSuccessTrue PageruleGetResponseEnvelopeSuccess = true
)

func (r PageruleGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PageruleGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
