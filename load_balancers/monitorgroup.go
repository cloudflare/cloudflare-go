// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v6/shared"
)

// MonitorGroupService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMonitorGroupService] method instead.
type MonitorGroupService struct {
	Options []option.RequestOption
}

// NewMonitorGroupService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMonitorGroupService(opts ...option.RequestOption) (r *MonitorGroupService) {
	r = &MonitorGroupService{}
	r.Options = opts
	return
}

// Create a new monitor group.
func (r *MonitorGroupService) New(ctx context.Context, params MonitorGroupNewParams, opts ...option.RequestOption) (res *MonitorGroup, err error) {
	var env MonitorGroupNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/load_balancers/monitor_groups", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify a configured monitor group.
func (r *MonitorGroupService) Update(ctx context.Context, monitorGroupID string, params MonitorGroupUpdateParams, opts ...option.RequestOption) (res *MonitorGroup, err error) {
	var env MonitorGroupUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if monitorGroupID == "" {
		err = errors.New("missing required monitor_group_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/load_balancers/monitor_groups/%s", params.AccountID, monitorGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List configured monitor groups.
func (r *MonitorGroupService) List(ctx context.Context, query MonitorGroupListParams, opts ...option.RequestOption) (res *pagination.SinglePage[MonitorGroup], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/load_balancers/monitor_groups", query.AccountID)
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

// List configured monitor groups.
func (r *MonitorGroupService) ListAutoPaging(ctx context.Context, query MonitorGroupListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[MonitorGroup] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a configured monitor group.
func (r *MonitorGroupService) Delete(ctx context.Context, monitorGroupID string, body MonitorGroupDeleteParams, opts ...option.RequestOption) (res *MonitorGroup, err error) {
	var env MonitorGroupDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if monitorGroupID == "" {
		err = errors.New("missing required monitor_group_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/load_balancers/monitor_groups/%s", body.AccountID, monitorGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Apply changes to an existing monitor group, overwriting the supplied properties.
func (r *MonitorGroupService) Edit(ctx context.Context, monitorGroupID string, params MonitorGroupEditParams, opts ...option.RequestOption) (res *MonitorGroup, err error) {
	var env MonitorGroupEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if monitorGroupID == "" {
		err = errors.New("missing required monitor_group_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/load_balancers/monitor_groups/%s", params.AccountID, monitorGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a single configured monitor group.
func (r *MonitorGroupService) Get(ctx context.Context, monitorGroupID string, query MonitorGroupGetParams, opts ...option.RequestOption) (res *MonitorGroup, err error) {
	var env MonitorGroupGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if monitorGroupID == "" {
		err = errors.New("missing required monitor_group_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/load_balancers/monitor_groups/%s", query.AccountID, monitorGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MonitorGroup struct {
	// The ID of the Monitor Group to use for checking the health of origins within
	// this pool.
	ID string `json:"id,required"`
	// A short description of the monitor group
	Description string `json:"description,required"`
	// List of monitors in this group
	Members []MonitorGroupMember `json:"members,required"`
	// The timestamp of when the monitor group was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The timestamp of when the monitor group was last updated
	UpdatedAt time.Time        `json:"updated_at" format:"date-time"`
	JSON      monitorGroupJSON `json:"-"`
}

// monitorGroupJSON contains the JSON metadata for the struct [MonitorGroup]
type monitorGroupJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Members     apijson.Field
	CreatedAt   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorGroupJSON) RawJSON() string {
	return r.raw
}

type MonitorGroupMember struct {
	// Whether this monitor is enabled in the group
	Enabled bool `json:"enabled,required"`
	// The ID of the Monitor to use for checking the health of origins within this
	// pool.
	MonitorID string `json:"monitor_id,required"`
	// Whether this monitor is used for monitoring only (does not affect pool health)
	MonitoringOnly bool `json:"monitoring_only,required"`
	// Whether this monitor must be healthy for the pool to be considered healthy
	MustBeHealthy bool `json:"must_be_healthy,required"`
	// The timestamp of when the monitor was added to the group
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The timestamp of when the monitor group member was last updated
	UpdatedAt time.Time              `json:"updated_at" format:"date-time"`
	JSON      monitorGroupMemberJSON `json:"-"`
}

// monitorGroupMemberJSON contains the JSON metadata for the struct
// [MonitorGroupMember]
type monitorGroupMemberJSON struct {
	Enabled        apijson.Field
	MonitorID      apijson.Field
	MonitoringOnly apijson.Field
	MustBeHealthy  apijson.Field
	CreatedAt      apijson.Field
	UpdatedAt      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *MonitorGroupMember) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorGroupMemberJSON) RawJSON() string {
	return r.raw
}

type MonitorGroupParam struct {
	// The ID of the Monitor Group to use for checking the health of origins within
	// this pool.
	ID param.Field[string] `json:"id,required"`
	// A short description of the monitor group
	Description param.Field[string] `json:"description,required"`
	// List of monitors in this group
	Members param.Field[[]MonitorGroupMemberParam] `json:"members,required"`
}

func (r MonitorGroupParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MonitorGroupMemberParam struct {
	// Whether this monitor is enabled in the group
	Enabled param.Field[bool] `json:"enabled,required"`
	// The ID of the Monitor to use for checking the health of origins within this
	// pool.
	MonitorID param.Field[string] `json:"monitor_id,required"`
	// Whether this monitor is used for monitoring only (does not affect pool health)
	MonitoringOnly param.Field[bool] `json:"monitoring_only,required"`
	// Whether this monitor must be healthy for the pool to be considered healthy
	MustBeHealthy param.Field[bool] `json:"must_be_healthy,required"`
}

func (r MonitorGroupMemberParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MonitorGroupNewParams struct {
	// Identifier.
	AccountID    param.Field[string] `path:"account_id,required"`
	MonitorGroup MonitorGroupParam   `json:"monitor_group,required"`
}

func (r MonitorGroupNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.MonitorGroup)
}

type MonitorGroupNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   MonitorGroup          `json:"result,required"`
	// Whether the API call was successful.
	Success MonitorGroupNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    monitorGroupNewResponseEnvelopeJSON    `json:"-"`
}

// monitorGroupNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [MonitorGroupNewResponseEnvelope]
type monitorGroupNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorGroupNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorGroupNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type MonitorGroupNewResponseEnvelopeSuccess bool

const (
	MonitorGroupNewResponseEnvelopeSuccessTrue MonitorGroupNewResponseEnvelopeSuccess = true
)

func (r MonitorGroupNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MonitorGroupNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type MonitorGroupUpdateParams struct {
	// Identifier.
	AccountID    param.Field[string] `path:"account_id,required"`
	MonitorGroup MonitorGroupParam   `json:"monitor_group,required"`
}

func (r MonitorGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.MonitorGroup)
}

type MonitorGroupUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   MonitorGroup          `json:"result,required"`
	// Whether the API call was successful.
	Success MonitorGroupUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    monitorGroupUpdateResponseEnvelopeJSON    `json:"-"`
}

// monitorGroupUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [MonitorGroupUpdateResponseEnvelope]
type monitorGroupUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorGroupUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorGroupUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type MonitorGroupUpdateResponseEnvelopeSuccess bool

const (
	MonitorGroupUpdateResponseEnvelopeSuccessTrue MonitorGroupUpdateResponseEnvelopeSuccess = true
)

func (r MonitorGroupUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MonitorGroupUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type MonitorGroupListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type MonitorGroupDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type MonitorGroupDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   MonitorGroup          `json:"result,required"`
	// Whether the API call was successful.
	Success MonitorGroupDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    monitorGroupDeleteResponseEnvelopeJSON    `json:"-"`
}

// monitorGroupDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [MonitorGroupDeleteResponseEnvelope]
type monitorGroupDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorGroupDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorGroupDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type MonitorGroupDeleteResponseEnvelopeSuccess bool

const (
	MonitorGroupDeleteResponseEnvelopeSuccessTrue MonitorGroupDeleteResponseEnvelopeSuccess = true
)

func (r MonitorGroupDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MonitorGroupDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type MonitorGroupEditParams struct {
	// Identifier.
	AccountID    param.Field[string] `path:"account_id,required"`
	MonitorGroup MonitorGroupParam   `json:"monitor_group,required"`
}

func (r MonitorGroupEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.MonitorGroup)
}

type MonitorGroupEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   MonitorGroup          `json:"result,required"`
	// Whether the API call was successful.
	Success MonitorGroupEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    monitorGroupEditResponseEnvelopeJSON    `json:"-"`
}

// monitorGroupEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [MonitorGroupEditResponseEnvelope]
type monitorGroupEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorGroupEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorGroupEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type MonitorGroupEditResponseEnvelopeSuccess bool

const (
	MonitorGroupEditResponseEnvelopeSuccessTrue MonitorGroupEditResponseEnvelopeSuccess = true
)

func (r MonitorGroupEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MonitorGroupEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type MonitorGroupGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type MonitorGroupGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   MonitorGroup          `json:"result,required"`
	// Whether the API call was successful.
	Success MonitorGroupGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    monitorGroupGetResponseEnvelopeJSON    `json:"-"`
}

// monitorGroupGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [MonitorGroupGetResponseEnvelope]
type monitorGroupGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorGroupGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorGroupGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type MonitorGroupGetResponseEnvelopeSuccess bool

const (
	MonitorGroupGetResponseEnvelopeSuccessTrue MonitorGroupGetResponseEnvelopeSuccess = true
)

func (r MonitorGroupGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MonitorGroupGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
