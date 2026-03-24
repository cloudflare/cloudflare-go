// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_gateway

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
)

// LabelUserService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLabelUserService] method instead.
type LabelUserService struct {
	Options   []option.RequestOption
	Resources *LabelUserResourceService
}

// NewLabelUserService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLabelUserService(opts ...option.RequestOption) (r *LabelUserService) {
	r = &LabelUserService{}
	r.Options = opts
	r.Resources = NewLabelUserResourceService(opts...)
	return
}

// Update all fields on a label
func (r *LabelUserService) Update(ctx context.Context, name string, params LabelUserUpdateParams, opts ...option.RequestOption) (res *LabelUserUpdateResponse, err error) {
	var env LabelUserUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if name == "" {
		err = errors.New("missing required name parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/labels/user/%s", params.ZoneID, name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete user label
func (r *LabelUserService) Delete(ctx context.Context, name string, body LabelUserDeleteParams, opts ...option.RequestOption) (res *LabelUserDeleteResponse, err error) {
	var env LabelUserDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if name == "" {
		err = errors.New("missing required name parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/labels/user/%s", body.ZoneID, name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Create user labels
func (r *LabelUserService) BulkNew(ctx context.Context, params LabelUserBulkNewParams, opts ...option.RequestOption) (res *pagination.SinglePage[LabelUserBulkNewResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/labels/user", params.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
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

// Create user labels
func (r *LabelUserService) BulkNewAutoPaging(ctx context.Context, params LabelUserBulkNewParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[LabelUserBulkNewResponse] {
	return pagination.NewSinglePageAutoPager(r.BulkNew(ctx, params, opts...))
}

// Delete user labels
func (r *LabelUserService) BulkDelete(ctx context.Context, body LabelUserBulkDeleteParams, opts ...option.RequestOption) (res *pagination.SinglePage[LabelUserBulkDeleteResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/labels/user", body.ZoneID)
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

// Delete user labels
func (r *LabelUserService) BulkDeleteAutoPaging(ctx context.Context, body LabelUserBulkDeleteParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[LabelUserBulkDeleteResponse] {
	return pagination.NewSinglePageAutoPager(r.BulkDelete(ctx, body, opts...))
}

// Update certain fields on a label
func (r *LabelUserService) Edit(ctx context.Context, name string, params LabelUserEditParams, opts ...option.RequestOption) (res *LabelUserEditResponse, err error) {
	var env LabelUserEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if name == "" {
		err = errors.New("missing required name parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/labels/user/%s", params.ZoneID, name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieve user label
func (r *LabelUserService) Get(ctx context.Context, name string, params LabelUserGetParams, opts ...option.RequestOption) (res *LabelUserGetResponse, err error) {
	var env LabelUserGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if name == "" {
		err = errors.New("missing required name parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/labels/user/%s", params.ZoneID, name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LabelUserUpdateResponse struct {
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The description of the label
	Description string    `json:"description,required"`
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// Metadata for the label
	Metadata interface{} `json:"metadata,required"`
	// The name of the label
	Name string `json:"name,required"`
	// - `user` - label is owned by the user
	// - `managed` - label is owned by cloudflare
	Source LabelUserUpdateResponseSource `json:"source,required"`
	JSON   labelUserUpdateResponseJSON   `json:"-"`
}

// labelUserUpdateResponseJSON contains the JSON metadata for the struct
// [LabelUserUpdateResponse]
type labelUserUpdateResponseJSON struct {
	CreatedAt   apijson.Field
	Description apijson.Field
	LastUpdated apijson.Field
	Metadata    apijson.Field
	Name        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LabelUserUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r labelUserUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// - `user` - label is owned by the user
// - `managed` - label is owned by cloudflare
type LabelUserUpdateResponseSource string

const (
	LabelUserUpdateResponseSourceUser    LabelUserUpdateResponseSource = "user"
	LabelUserUpdateResponseSourceManaged LabelUserUpdateResponseSource = "managed"
)

func (r LabelUserUpdateResponseSource) IsKnown() bool {
	switch r {
	case LabelUserUpdateResponseSourceUser, LabelUserUpdateResponseSourceManaged:
		return true
	}
	return false
}

type LabelUserDeleteResponse struct {
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The description of the label
	Description string    `json:"description,required"`
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// Metadata for the label
	Metadata interface{} `json:"metadata,required"`
	// The name of the label
	Name string `json:"name,required"`
	// - `user` - label is owned by the user
	// - `managed` - label is owned by cloudflare
	Source LabelUserDeleteResponseSource `json:"source,required"`
	JSON   labelUserDeleteResponseJSON   `json:"-"`
}

// labelUserDeleteResponseJSON contains the JSON metadata for the struct
// [LabelUserDeleteResponse]
type labelUserDeleteResponseJSON struct {
	CreatedAt   apijson.Field
	Description apijson.Field
	LastUpdated apijson.Field
	Metadata    apijson.Field
	Name        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LabelUserDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r labelUserDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// - `user` - label is owned by the user
// - `managed` - label is owned by cloudflare
type LabelUserDeleteResponseSource string

const (
	LabelUserDeleteResponseSourceUser    LabelUserDeleteResponseSource = "user"
	LabelUserDeleteResponseSourceManaged LabelUserDeleteResponseSource = "managed"
)

func (r LabelUserDeleteResponseSource) IsKnown() bool {
	switch r {
	case LabelUserDeleteResponseSourceUser, LabelUserDeleteResponseSourceManaged:
		return true
	}
	return false
}

type LabelUserBulkNewResponse struct {
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The description of the label
	Description string    `json:"description,required"`
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// Metadata for the label
	Metadata interface{} `json:"metadata,required"`
	// The name of the label
	Name string `json:"name,required"`
	// - `user` - label is owned by the user
	// - `managed` - label is owned by cloudflare
	Source LabelUserBulkNewResponseSource `json:"source,required"`
	JSON   labelUserBulkNewResponseJSON   `json:"-"`
}

// labelUserBulkNewResponseJSON contains the JSON metadata for the struct
// [LabelUserBulkNewResponse]
type labelUserBulkNewResponseJSON struct {
	CreatedAt   apijson.Field
	Description apijson.Field
	LastUpdated apijson.Field
	Metadata    apijson.Field
	Name        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LabelUserBulkNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r labelUserBulkNewResponseJSON) RawJSON() string {
	return r.raw
}

// - `user` - label is owned by the user
// - `managed` - label is owned by cloudflare
type LabelUserBulkNewResponseSource string

const (
	LabelUserBulkNewResponseSourceUser    LabelUserBulkNewResponseSource = "user"
	LabelUserBulkNewResponseSourceManaged LabelUserBulkNewResponseSource = "managed"
)

func (r LabelUserBulkNewResponseSource) IsKnown() bool {
	switch r {
	case LabelUserBulkNewResponseSourceUser, LabelUserBulkNewResponseSourceManaged:
		return true
	}
	return false
}

type LabelUserBulkDeleteResponse struct {
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The description of the label
	Description string    `json:"description,required"`
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// Metadata for the label
	Metadata interface{} `json:"metadata,required"`
	// The name of the label
	Name string `json:"name,required"`
	// - `user` - label is owned by the user
	// - `managed` - label is owned by cloudflare
	Source LabelUserBulkDeleteResponseSource `json:"source,required"`
	JSON   labelUserBulkDeleteResponseJSON   `json:"-"`
}

// labelUserBulkDeleteResponseJSON contains the JSON metadata for the struct
// [LabelUserBulkDeleteResponse]
type labelUserBulkDeleteResponseJSON struct {
	CreatedAt   apijson.Field
	Description apijson.Field
	LastUpdated apijson.Field
	Metadata    apijson.Field
	Name        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LabelUserBulkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r labelUserBulkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// - `user` - label is owned by the user
// - `managed` - label is owned by cloudflare
type LabelUserBulkDeleteResponseSource string

const (
	LabelUserBulkDeleteResponseSourceUser    LabelUserBulkDeleteResponseSource = "user"
	LabelUserBulkDeleteResponseSourceManaged LabelUserBulkDeleteResponseSource = "managed"
)

func (r LabelUserBulkDeleteResponseSource) IsKnown() bool {
	switch r {
	case LabelUserBulkDeleteResponseSourceUser, LabelUserBulkDeleteResponseSourceManaged:
		return true
	}
	return false
}

type LabelUserEditResponse struct {
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The description of the label
	Description string    `json:"description,required"`
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// Metadata for the label
	Metadata interface{} `json:"metadata,required"`
	// The name of the label
	Name string `json:"name,required"`
	// - `user` - label is owned by the user
	// - `managed` - label is owned by cloudflare
	Source LabelUserEditResponseSource `json:"source,required"`
	JSON   labelUserEditResponseJSON   `json:"-"`
}

// labelUserEditResponseJSON contains the JSON metadata for the struct
// [LabelUserEditResponse]
type labelUserEditResponseJSON struct {
	CreatedAt   apijson.Field
	Description apijson.Field
	LastUpdated apijson.Field
	Metadata    apijson.Field
	Name        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LabelUserEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r labelUserEditResponseJSON) RawJSON() string {
	return r.raw
}

// - `user` - label is owned by the user
// - `managed` - label is owned by cloudflare
type LabelUserEditResponseSource string

const (
	LabelUserEditResponseSourceUser    LabelUserEditResponseSource = "user"
	LabelUserEditResponseSourceManaged LabelUserEditResponseSource = "managed"
)

func (r LabelUserEditResponseSource) IsKnown() bool {
	switch r {
	case LabelUserEditResponseSourceUser, LabelUserEditResponseSourceManaged:
		return true
	}
	return false
}

type LabelUserGetResponse struct {
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The description of the label
	Description string    `json:"description,required"`
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// Metadata for the label
	Metadata interface{} `json:"metadata,required"`
	// The name of the label
	Name string `json:"name,required"`
	// - `user` - label is owned by the user
	// - `managed` - label is owned by cloudflare
	Source LabelUserGetResponseSource `json:"source,required"`
	// Provides counts of what resources are linked to this label
	MappedResources interface{}              `json:"mapped_resources"`
	JSON            labelUserGetResponseJSON `json:"-"`
}

// labelUserGetResponseJSON contains the JSON metadata for the struct
// [LabelUserGetResponse]
type labelUserGetResponseJSON struct {
	CreatedAt       apijson.Field
	Description     apijson.Field
	LastUpdated     apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	Source          apijson.Field
	MappedResources apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *LabelUserGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r labelUserGetResponseJSON) RawJSON() string {
	return r.raw
}

// - `user` - label is owned by the user
// - `managed` - label is owned by cloudflare
type LabelUserGetResponseSource string

const (
	LabelUserGetResponseSourceUser    LabelUserGetResponseSource = "user"
	LabelUserGetResponseSourceManaged LabelUserGetResponseSource = "managed"
)

func (r LabelUserGetResponseSource) IsKnown() bool {
	switch r {
	case LabelUserGetResponseSourceUser, LabelUserGetResponseSourceManaged:
		return true
	}
	return false
}

type LabelUserUpdateParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The description of the label
	Description param.Field[string] `json:"description"`
	// Metadata for the label
	Metadata param.Field[interface{}] `json:"metadata"`
}

func (r LabelUserUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LabelUserUpdateResponseEnvelope struct {
	Errors   Message                 `json:"errors,required"`
	Messages Message                 `json:"messages,required"`
	Result   LabelUserUpdateResponse `json:"result,required"`
	// Whether the API call was successful.
	Success LabelUserUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    labelUserUpdateResponseEnvelopeJSON    `json:"-"`
}

// labelUserUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [LabelUserUpdateResponseEnvelope]
type labelUserUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LabelUserUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r labelUserUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type LabelUserUpdateResponseEnvelopeSuccess bool

const (
	LabelUserUpdateResponseEnvelopeSuccessTrue LabelUserUpdateResponseEnvelopeSuccess = true
)

func (r LabelUserUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LabelUserUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LabelUserDeleteParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type LabelUserDeleteResponseEnvelope struct {
	Errors   Message                 `json:"errors,required"`
	Messages Message                 `json:"messages,required"`
	Result   LabelUserDeleteResponse `json:"result,required"`
	// Whether the API call was successful.
	Success LabelUserDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    labelUserDeleteResponseEnvelopeJSON    `json:"-"`
}

// labelUserDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [LabelUserDeleteResponseEnvelope]
type labelUserDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LabelUserDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r labelUserDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type LabelUserDeleteResponseEnvelopeSuccess bool

const (
	LabelUserDeleteResponseEnvelopeSuccessTrue LabelUserDeleteResponseEnvelopeSuccess = true
)

func (r LabelUserDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LabelUserDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LabelUserBulkNewParams struct {
	// Identifier.
	ZoneID param.Field[string]          `path:"zone_id,required"`
	Body   []LabelUserBulkNewParamsBody `json:"body,required"`
}

func (r LabelUserBulkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type LabelUserBulkNewParamsBody struct {
	// The name of the label
	Name param.Field[string] `json:"name,required"`
	// The description of the label
	Description param.Field[string] `json:"description"`
	// Metadata for the label
	Metadata param.Field[interface{}] `json:"metadata"`
}

func (r LabelUserBulkNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LabelUserBulkDeleteParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type LabelUserEditParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The description of the label
	Description param.Field[string] `json:"description"`
	// Metadata for the label
	Metadata param.Field[interface{}] `json:"metadata"`
}

func (r LabelUserEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LabelUserEditResponseEnvelope struct {
	Errors   Message               `json:"errors,required"`
	Messages Message               `json:"messages,required"`
	Result   LabelUserEditResponse `json:"result,required"`
	// Whether the API call was successful.
	Success LabelUserEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    labelUserEditResponseEnvelopeJSON    `json:"-"`
}

// labelUserEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [LabelUserEditResponseEnvelope]
type labelUserEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LabelUserEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r labelUserEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type LabelUserEditResponseEnvelopeSuccess bool

const (
	LabelUserEditResponseEnvelopeSuccessTrue LabelUserEditResponseEnvelopeSuccess = true
)

func (r LabelUserEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LabelUserEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LabelUserGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Include `mapped_resources` for each label
	WithMappedResourceCounts param.Field[bool] `query:"with_mapped_resource_counts"`
}

// URLQuery serializes [LabelUserGetParams]'s query parameters as `url.Values`.
func (r LabelUserGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LabelUserGetResponseEnvelope struct {
	Errors   Message              `json:"errors,required"`
	Messages Message              `json:"messages,required"`
	Result   LabelUserGetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success LabelUserGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    labelUserGetResponseEnvelopeJSON    `json:"-"`
}

// labelUserGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [LabelUserGetResponseEnvelope]
type labelUserGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LabelUserGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r labelUserGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type LabelUserGetResponseEnvelopeSuccess bool

const (
	LabelUserGetResponseEnvelopeSuccessTrue LabelUserGetResponseEnvelopeSuccess = true
)

func (r LabelUserGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LabelUserGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
