// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// ThreatEventQueryService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventQueryService] method instead.
type ThreatEventQueryService struct {
	Options []option.RequestOption
}

// NewThreatEventQueryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewThreatEventQueryService(opts ...option.RequestOption) (r *ThreatEventQueryService) {
	r = &ThreatEventQueryService{}
	r.Options = opts
	return
}

// Create a new saved event query for the account
func (r *ThreatEventQueryService) New(ctx context.Context, params ThreatEventQueryNewParams, opts ...option.RequestOption) (res *ThreatEventQueryNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/queries/create", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieve all saved event queries for the account
func (r *ThreatEventQueryService) List(ctx context.Context, query ThreatEventQueryListParams, opts ...option.RequestOption) (res *[]ThreatEventQueryListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/queries", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete a saved event query by its ID
func (r *ThreatEventQueryService) Delete(ctx context.Context, queryID int64, body ThreatEventQueryDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/queries/%v", body.AccountID, queryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Update an existing saved event query by its ID
func (r *ThreatEventQueryService) Edit(ctx context.Context, queryID int64, params ThreatEventQueryEditParams, opts ...option.RequestOption) (res *ThreatEventQueryEditResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/queries/%v", params.AccountID, queryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Retrieve a saved event query by its ID
func (r *ThreatEventQueryService) Get(ctx context.Context, queryID int64, query ThreatEventQueryGetParams, opts ...option.RequestOption) (res *ThreatEventQueryGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/queries/%v", query.AccountID, queryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type ThreatEventQueryNewResponse struct {
	// Unique identifier for the saved query
	ID int64 `json:"id" api:"required"`
	// Account ID
	AccountID int64 `json:"account_id" api:"required"`
	// Whether alerts are enabled
	AlertEnabled bool `json:"alert_enabled" api:"required"`
	// Whether alert rollup is enabled
	AlertRollupEnabled bool `json:"alert_rollup_enabled" api:"required"`
	// Creation timestamp
	CreatedAt string `json:"created_at" api:"required"`
	// Name of the saved query
	Name string `json:"name" api:"required"`
	// JSON string containing the query parameters
	QueryJson string `json:"query_json" api:"required"`
	// Whether rule is enabled
	RuleEnabled bool `json:"rule_enabled" api:"required"`
	// Last update timestamp
	UpdatedAt string `json:"updated_at" api:"required"`
	// Email of the user who created the query
	UserEmail string `json:"user_email" api:"required"`
	// Intel Indicator Feed ID (numeric)
	CustomThreatFeedID int64 `json:"custom_threat_feed_id" api:"nullable"`
	// WAF rules list ID for blocking
	RuleListID string `json:"rule_list_id"`
	// Scope for the rule
	RuleScope string                          `json:"rule_scope"`
	JSON      threatEventQueryNewResponseJSON `json:"-"`
}

// threatEventQueryNewResponseJSON contains the JSON metadata for the struct
// [ThreatEventQueryNewResponse]
type threatEventQueryNewResponseJSON struct {
	ID                 apijson.Field
	AccountID          apijson.Field
	AlertEnabled       apijson.Field
	AlertRollupEnabled apijson.Field
	CreatedAt          apijson.Field
	Name               apijson.Field
	QueryJson          apijson.Field
	RuleEnabled        apijson.Field
	UpdatedAt          apijson.Field
	UserEmail          apijson.Field
	CustomThreatFeedID apijson.Field
	RuleListID         apijson.Field
	RuleScope          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ThreatEventQueryNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventQueryNewResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventQueryListResponse struct {
	// Unique identifier for the saved query
	ID int64 `json:"id" api:"required"`
	// Account ID
	AccountID int64 `json:"account_id" api:"required"`
	// Whether alerts are enabled
	AlertEnabled bool `json:"alert_enabled" api:"required"`
	// Whether alert rollup is enabled
	AlertRollupEnabled bool `json:"alert_rollup_enabled" api:"required"`
	// Creation timestamp
	CreatedAt string `json:"created_at" api:"required"`
	// Name of the saved query
	Name string `json:"name" api:"required"`
	// JSON string containing the query parameters
	QueryJson string `json:"query_json" api:"required"`
	// Whether rule is enabled
	RuleEnabled bool `json:"rule_enabled" api:"required"`
	// Last update timestamp
	UpdatedAt string `json:"updated_at" api:"required"`
	// Email of the user who created the query
	UserEmail string `json:"user_email" api:"required"`
	// Intel Indicator Feed ID (numeric)
	CustomThreatFeedID int64 `json:"custom_threat_feed_id" api:"nullable"`
	// WAF rules list ID for blocking
	RuleListID string `json:"rule_list_id"`
	// Scope for the rule
	RuleScope string                           `json:"rule_scope"`
	JSON      threatEventQueryListResponseJSON `json:"-"`
}

// threatEventQueryListResponseJSON contains the JSON metadata for the struct
// [ThreatEventQueryListResponse]
type threatEventQueryListResponseJSON struct {
	ID                 apijson.Field
	AccountID          apijson.Field
	AlertEnabled       apijson.Field
	AlertRollupEnabled apijson.Field
	CreatedAt          apijson.Field
	Name               apijson.Field
	QueryJson          apijson.Field
	RuleEnabled        apijson.Field
	UpdatedAt          apijson.Field
	UserEmail          apijson.Field
	CustomThreatFeedID apijson.Field
	RuleListID         apijson.Field
	RuleScope          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ThreatEventQueryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventQueryListResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventQueryEditResponse struct {
	// Unique identifier for the saved query
	ID int64 `json:"id" api:"required"`
	// Account ID
	AccountID int64 `json:"account_id" api:"required"`
	// Whether alerts are enabled
	AlertEnabled bool `json:"alert_enabled" api:"required"`
	// Whether alert rollup is enabled
	AlertRollupEnabled bool `json:"alert_rollup_enabled" api:"required"`
	// Creation timestamp
	CreatedAt string `json:"created_at" api:"required"`
	// Name of the saved query
	Name string `json:"name" api:"required"`
	// JSON string containing the query parameters
	QueryJson string `json:"query_json" api:"required"`
	// Whether rule is enabled
	RuleEnabled bool `json:"rule_enabled" api:"required"`
	// Last update timestamp
	UpdatedAt string `json:"updated_at" api:"required"`
	// Email of the user who created the query
	UserEmail string `json:"user_email" api:"required"`
	// Intel Indicator Feed ID (numeric)
	CustomThreatFeedID int64 `json:"custom_threat_feed_id" api:"nullable"`
	// WAF rules list ID for blocking
	RuleListID string `json:"rule_list_id"`
	// Scope for the rule
	RuleScope string                           `json:"rule_scope"`
	JSON      threatEventQueryEditResponseJSON `json:"-"`
}

// threatEventQueryEditResponseJSON contains the JSON metadata for the struct
// [ThreatEventQueryEditResponse]
type threatEventQueryEditResponseJSON struct {
	ID                 apijson.Field
	AccountID          apijson.Field
	AlertEnabled       apijson.Field
	AlertRollupEnabled apijson.Field
	CreatedAt          apijson.Field
	Name               apijson.Field
	QueryJson          apijson.Field
	RuleEnabled        apijson.Field
	UpdatedAt          apijson.Field
	UserEmail          apijson.Field
	CustomThreatFeedID apijson.Field
	RuleListID         apijson.Field
	RuleScope          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ThreatEventQueryEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventQueryEditResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventQueryGetResponse struct {
	// Unique identifier for the saved query
	ID int64 `json:"id" api:"required"`
	// Account ID
	AccountID int64 `json:"account_id" api:"required"`
	// Whether alerts are enabled
	AlertEnabled bool `json:"alert_enabled" api:"required"`
	// Whether alert rollup is enabled
	AlertRollupEnabled bool `json:"alert_rollup_enabled" api:"required"`
	// Creation timestamp
	CreatedAt string `json:"created_at" api:"required"`
	// Name of the saved query
	Name string `json:"name" api:"required"`
	// JSON string containing the query parameters
	QueryJson string `json:"query_json" api:"required"`
	// Whether rule is enabled
	RuleEnabled bool `json:"rule_enabled" api:"required"`
	// Last update timestamp
	UpdatedAt string `json:"updated_at" api:"required"`
	// Email of the user who created the query
	UserEmail string `json:"user_email" api:"required"`
	// Intel Indicator Feed ID (numeric)
	CustomThreatFeedID int64 `json:"custom_threat_feed_id" api:"nullable"`
	// WAF rules list ID for blocking
	RuleListID string `json:"rule_list_id"`
	// Scope for the rule
	RuleScope string                          `json:"rule_scope"`
	JSON      threatEventQueryGetResponseJSON `json:"-"`
}

// threatEventQueryGetResponseJSON contains the JSON metadata for the struct
// [ThreatEventQueryGetResponse]
type threatEventQueryGetResponseJSON struct {
	ID                 apijson.Field
	AccountID          apijson.Field
	AlertEnabled       apijson.Field
	AlertRollupEnabled apijson.Field
	CreatedAt          apijson.Field
	Name               apijson.Field
	QueryJson          apijson.Field
	RuleEnabled        apijson.Field
	UpdatedAt          apijson.Field
	UserEmail          apijson.Field
	CustomThreatFeedID apijson.Field
	RuleListID         apijson.Field
	RuleScope          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ThreatEventQueryGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventQueryGetResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventQueryNewParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Enable alerts for this query
	AlertEnabled param.Field[bool] `json:"alert_enabled" api:"required"`
	// Enable alert rollup for this query
	AlertRollupEnabled param.Field[bool] `json:"alert_rollup_enabled" api:"required"`
	// Unique name for the saved query
	Name param.Field[string] `json:"name" api:"required"`
	// JSON string containing the query parameters
	QueryJson param.Field[string] `json:"query_json" api:"required"`
	// Enable rule for this query
	RuleEnabled param.Field[bool] `json:"rule_enabled" api:"required"`
	// Scope for the rule
	RuleScope param.Field[string] `json:"rule_scope"`
}

func (r ThreatEventQueryNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventQueryListParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ThreatEventQueryDeleteParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ThreatEventQueryEditParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Enable alerts for this query
	AlertEnabled param.Field[bool] `json:"alert_enabled"`
	// Enable alert rollup for this query
	AlertRollupEnabled param.Field[bool] `json:"alert_rollup_enabled"`
	// Unique name for the saved query
	Name param.Field[string] `json:"name"`
	// JSON string containing the query parameters
	QueryJson param.Field[string] `json:"query_json"`
	// Enable rule for this query
	RuleEnabled param.Field[bool] `json:"rule_enabled"`
	// Scope for the rule
	RuleScope param.Field[string] `json:"rule_scope"`
}

func (r ThreatEventQueryEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventQueryGetParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}
