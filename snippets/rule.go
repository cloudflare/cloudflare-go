// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package snippets

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// RuleService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRuleService] method instead.
type RuleService struct {
	Options []option.RequestOption
}

// NewRuleService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRuleService(opts ...option.RequestOption) (r *RuleService) {
	r = &RuleService{}
	r.Options = opts
	return
}

// Updates all snippet rules belonging to the zone.
func (r *RuleService) Update(ctx context.Context, params RuleUpdateParams, opts ...option.RequestOption) (res *pagination.SinglePage[RuleUpdateResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/snippets/snippet_rules", params.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPut, path, params, &res, opts...)
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

// Updates all snippet rules belonging to the zone.
func (r *RuleService) UpdateAutoPaging(ctx context.Context, params RuleUpdateParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[RuleUpdateResponse] {
	return pagination.NewSinglePageAutoPager(r.Update(ctx, params, opts...))
}

// Fetches all snippet rules belonging to the zone.
func (r *RuleService) List(ctx context.Context, query RuleListParams, opts ...option.RequestOption) (res *pagination.SinglePage[RuleListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/snippets/snippet_rules", query.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// Fetches all snippet rules belonging to the zone.
func (r *RuleService) ListAutoPaging(ctx context.Context, query RuleListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[RuleListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes all snippet rules belonging to the zone.
func (r *RuleService) Delete(ctx context.Context, body RuleDeleteParams, opts ...option.RequestOption) (res *pagination.SinglePage[RuleDeleteResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/snippets/snippet_rules", body.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodDelete, path, nil, &res, opts...)
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

// Deletes all snippet rules belonging to the zone.
func (r *RuleService) DeleteAutoPaging(ctx context.Context, body RuleDeleteParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[RuleDeleteResponse] {
	return pagination.NewSinglePageAutoPager(r.Delete(ctx, body, opts...))
}

// Fetches all snippet rules belonging to the zone.
func (r *RuleService) Get(ctx context.Context, query RuleGetParams, opts ...option.RequestOption) (res *pagination.SinglePage[RuleGetResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/snippets/snippet_rules", query.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// Fetches all snippet rules belonging to the zone.
func (r *RuleService) GetAutoPaging(ctx context.Context, query RuleGetParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[RuleGetResponse] {
	return pagination.NewSinglePageAutoPager(r.Get(ctx, query, opts...))
}

// Define a snippet rule.
type RuleUpdateResponse struct {
	// Specify the unique ID of the rule.
	ID string `json:"id" api:"required"`
	// Define the expression that determines which traffic matches the rule.
	Expression string `json:"expression" api:"required"`
	// Specify the timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// Identify the snippet.
	SnippetName string `json:"snippet_name" api:"required"`
	// Provide an informative description of the rule.
	Description string `json:"description"`
	// Indicate whether to execute the rule.
	Enabled bool                   `json:"enabled"`
	JSON    ruleUpdateResponseJSON `json:"-"`
}

// ruleUpdateResponseJSON contains the JSON metadata for the struct
// [RuleUpdateResponse]
type ruleUpdateResponseJSON struct {
	ID          apijson.Field
	Expression  apijson.Field
	LastUpdated apijson.Field
	SnippetName apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Define a snippet rule.
type RuleListResponse struct {
	// Specify the unique ID of the rule.
	ID string `json:"id" api:"required"`
	// Define the expression that determines which traffic matches the rule.
	Expression string `json:"expression" api:"required"`
	// Specify the timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// Identify the snippet.
	SnippetName string `json:"snippet_name" api:"required"`
	// Provide an informative description of the rule.
	Description string `json:"description"`
	// Indicate whether to execute the rule.
	Enabled bool                 `json:"enabled"`
	JSON    ruleListResponseJSON `json:"-"`
}

// ruleListResponseJSON contains the JSON metadata for the struct
// [RuleListResponse]
type ruleListResponseJSON struct {
	ID          apijson.Field
	Expression  apijson.Field
	LastUpdated apijson.Field
	SnippetName apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleListResponseJSON) RawJSON() string {
	return r.raw
}

// Define a snippet rule.
type RuleDeleteResponse struct {
	// Specify the unique ID of the rule.
	ID string `json:"id" api:"required"`
	// Define the expression that determines which traffic matches the rule.
	Expression string `json:"expression" api:"required"`
	// Specify the timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// Identify the snippet.
	SnippetName string `json:"snippet_name" api:"required"`
	// Provide an informative description of the rule.
	Description string `json:"description"`
	// Indicate whether to execute the rule.
	Enabled bool                   `json:"enabled"`
	JSON    ruleDeleteResponseJSON `json:"-"`
}

// ruleDeleteResponseJSON contains the JSON metadata for the struct
// [RuleDeleteResponse]
type ruleDeleteResponseJSON struct {
	ID          apijson.Field
	Expression  apijson.Field
	LastUpdated apijson.Field
	SnippetName apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Define a snippet rule.
type RuleGetResponse struct {
	// Specify the unique ID of the rule.
	ID string `json:"id" api:"required"`
	// Define the expression that determines which traffic matches the rule.
	Expression string `json:"expression" api:"required"`
	// Specify the timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// Identify the snippet.
	SnippetName string `json:"snippet_name" api:"required"`
	// Provide an informative description of the rule.
	Description string `json:"description"`
	// Indicate whether to execute the rule.
	Enabled bool                `json:"enabled"`
	JSON    ruleGetResponseJSON `json:"-"`
}

// ruleGetResponseJSON contains the JSON metadata for the struct [RuleGetResponse]
type ruleGetResponseJSON struct {
	ID          apijson.Field
	Expression  apijson.Field
	LastUpdated apijson.Field
	SnippetName apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleGetResponseJSON) RawJSON() string {
	return r.raw
}

type RuleUpdateParams struct {
	// Use this field to specify the unique ID of the zone.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// Lists snippet rules.
	Rules param.Field[[]RuleUpdateParamsRule] `json:"rules" api:"required"`
}

func (r RuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define a snippet rule.
type RuleUpdateParamsRule struct {
	// Define the expression that determines which traffic matches the rule.
	Expression param.Field[string] `json:"expression" api:"required"`
	// Identify the snippet.
	SnippetName param.Field[string] `json:"snippet_name" api:"required"`
	// Provide an informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Indicate whether to execute the rule.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r RuleUpdateParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleListParams struct {
	// Use this field to specify the unique ID of the zone.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type RuleDeleteParams struct {
	// Use this field to specify the unique ID of the zone.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type RuleGetParams struct {
	// Use this field to specify the unique ID of the zone.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}
