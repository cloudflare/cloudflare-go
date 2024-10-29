// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudflare

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
	"github.com/cloudflare/cloudflare-go/v3/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v3/shared"
)

// ResourceSharingResourceService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewResourceSharingResourceService] method instead.
type ResourceSharingResourceService struct {
	Options []option.RequestOption
}

// NewResourceSharingResourceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewResourceSharingResourceService(opts ...option.RequestOption) (r *ResourceSharingResourceService) {
	r = &ResourceSharingResourceService{}
	r.Options = opts
	return
}

// Create a new share resource
func (r *ResourceSharingResourceService) New(ctx context.Context, shareIdentifier string, params ResourceSharingResourceNewParams, opts ...option.RequestOption) (res *ResourceSharingResourceNewResponse, err error) {
	var env ResourceSharingResourceNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if shareIdentifier == "" {
		err = errors.New("missing required share_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares/%s/resources", params.AccountID, shareIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update is not immediate, an updated share resource object with a new status will
// be returned.
func (r *ResourceSharingResourceService) Update(ctx context.Context, shareIdentifier string, resourceIdentifier string, params ResourceSharingResourceUpdateParams, opts ...option.RequestOption) (res *ResourceSharingResourceUpdateResponse, err error) {
	var env ResourceSharingResourceUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if shareIdentifier == "" {
		err = errors.New("missing required share_identifier parameter")
		return
	}
	if resourceIdentifier == "" {
		err = errors.New("missing required resource_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares/%s/resources/%s", params.AccountID, shareIdentifier, resourceIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List share resources by share ID.
func (r *ResourceSharingResourceService) List(ctx context.Context, shareIdentifier string, params ResourceSharingResourceListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[ResourceSharingResourceListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if shareIdentifier == "" {
		err = errors.New("missing required share_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares/%s/resources", params.AccountID, shareIdentifier)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// List share resources by share ID.
func (r *ResourceSharingResourceService) ListAutoPaging(ctx context.Context, shareIdentifier string, params ResourceSharingResourceListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[ResourceSharingResourceListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, shareIdentifier, params, opts...))
}

// Deletion is not immediate, an updated share resource object with a new status
// will be returned.
func (r *ResourceSharingResourceService) Delete(ctx context.Context, shareIdentifier string, resourceIdentifier string, body ResourceSharingResourceDeleteParams, opts ...option.RequestOption) (res *ResourceSharingResourceDeleteResponse, err error) {
	var env ResourceSharingResourceDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if shareIdentifier == "" {
		err = errors.New("missing required share_identifier parameter")
		return
	}
	if resourceIdentifier == "" {
		err = errors.New("missing required resource_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares/%s/resources/%s", body.AccountID, shareIdentifier, resourceIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get share resource by ID.
func (r *ResourceSharingResourceService) Get(ctx context.Context, shareIdentifier string, resourceIdentifier string, query ResourceSharingResourceGetParams, opts ...option.RequestOption) (res *ResourceSharingResourceGetResponse, err error) {
	var env ResourceSharingResourceGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if shareIdentifier == "" {
		err = errors.New("missing required share_identifier parameter")
		return
	}
	if resourceIdentifier == "" {
		err = errors.New("missing required resource_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares/%s/resources/%s", query.AccountID, shareIdentifier, resourceIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ResourceSharingResourceNewResponse struct {
	// Share Resource identifier.
	ID string `json:"id,required"`
	// When the share was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// Resource Metadata.
	Meta interface{} `json:"meta,required"`
	// When the share was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// Account identifier.
	ResourceAccountID string `json:"resource_account_id,required"`
	// Share Resource identifier.
	ResourceID string `json:"resource_id,required"`
	// Resource Type.
	ResourceType ResourceSharingResourceNewResponseResourceType `json:"resource_type,required"`
	// Resource Version.
	ResourceVersion int64 `json:"resource_version,required"`
	// Resource Status.
	Status ResourceSharingResourceNewResponseStatus `json:"status,required"`
	JSON   resourceSharingResourceNewResponseJSON   `json:"-"`
}

// resourceSharingResourceNewResponseJSON contains the JSON metadata for the struct
// [ResourceSharingResourceNewResponse]
type resourceSharingResourceNewResponseJSON struct {
	ID                apijson.Field
	Created           apijson.Field
	Meta              apijson.Field
	Modified          apijson.Field
	ResourceAccountID apijson.Field
	ResourceID        apijson.Field
	ResourceType      apijson.Field
	ResourceVersion   apijson.Field
	Status            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ResourceSharingResourceNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSharingResourceNewResponseJSON) RawJSON() string {
	return r.raw
}

// Resource Type.
type ResourceSharingResourceNewResponseResourceType string

const (
	ResourceSharingResourceNewResponseResourceTypeCustomRuleset ResourceSharingResourceNewResponseResourceType = "custom-ruleset"
)

func (r ResourceSharingResourceNewResponseResourceType) IsKnown() bool {
	switch r {
	case ResourceSharingResourceNewResponseResourceTypeCustomRuleset:
		return true
	}
	return false
}

// Resource Status.
type ResourceSharingResourceNewResponseStatus string

const (
	ResourceSharingResourceNewResponseStatusActive   ResourceSharingResourceNewResponseStatus = "active"
	ResourceSharingResourceNewResponseStatusDeleting ResourceSharingResourceNewResponseStatus = "deleting"
	ResourceSharingResourceNewResponseStatusDeleted  ResourceSharingResourceNewResponseStatus = "deleted"
)

func (r ResourceSharingResourceNewResponseStatus) IsKnown() bool {
	switch r {
	case ResourceSharingResourceNewResponseStatusActive, ResourceSharingResourceNewResponseStatusDeleting, ResourceSharingResourceNewResponseStatusDeleted:
		return true
	}
	return false
}

type ResourceSharingResourceUpdateResponse struct {
	// Share Resource identifier.
	ID string `json:"id,required"`
	// When the share was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// Resource Metadata.
	Meta interface{} `json:"meta,required"`
	// When the share was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// Account identifier.
	ResourceAccountID string `json:"resource_account_id,required"`
	// Share Resource identifier.
	ResourceID string `json:"resource_id,required"`
	// Resource Type.
	ResourceType ResourceSharingResourceUpdateResponseResourceType `json:"resource_type,required"`
	// Resource Version.
	ResourceVersion int64 `json:"resource_version,required"`
	// Resource Status.
	Status ResourceSharingResourceUpdateResponseStatus `json:"status,required"`
	JSON   resourceSharingResourceUpdateResponseJSON   `json:"-"`
}

// resourceSharingResourceUpdateResponseJSON contains the JSON metadata for the
// struct [ResourceSharingResourceUpdateResponse]
type resourceSharingResourceUpdateResponseJSON struct {
	ID                apijson.Field
	Created           apijson.Field
	Meta              apijson.Field
	Modified          apijson.Field
	ResourceAccountID apijson.Field
	ResourceID        apijson.Field
	ResourceType      apijson.Field
	ResourceVersion   apijson.Field
	Status            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ResourceSharingResourceUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSharingResourceUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Resource Type.
type ResourceSharingResourceUpdateResponseResourceType string

const (
	ResourceSharingResourceUpdateResponseResourceTypeCustomRuleset ResourceSharingResourceUpdateResponseResourceType = "custom-ruleset"
)

func (r ResourceSharingResourceUpdateResponseResourceType) IsKnown() bool {
	switch r {
	case ResourceSharingResourceUpdateResponseResourceTypeCustomRuleset:
		return true
	}
	return false
}

// Resource Status.
type ResourceSharingResourceUpdateResponseStatus string

const (
	ResourceSharingResourceUpdateResponseStatusActive   ResourceSharingResourceUpdateResponseStatus = "active"
	ResourceSharingResourceUpdateResponseStatusDeleting ResourceSharingResourceUpdateResponseStatus = "deleting"
	ResourceSharingResourceUpdateResponseStatusDeleted  ResourceSharingResourceUpdateResponseStatus = "deleted"
)

func (r ResourceSharingResourceUpdateResponseStatus) IsKnown() bool {
	switch r {
	case ResourceSharingResourceUpdateResponseStatusActive, ResourceSharingResourceUpdateResponseStatusDeleting, ResourceSharingResourceUpdateResponseStatusDeleted:
		return true
	}
	return false
}

type ResourceSharingResourceListResponse struct {
	// Share Resource identifier.
	ID string `json:"id,required"`
	// When the share was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// Resource Metadata.
	Meta interface{} `json:"meta,required"`
	// When the share was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// Account identifier.
	ResourceAccountID string `json:"resource_account_id,required"`
	// Share Resource identifier.
	ResourceID string `json:"resource_id,required"`
	// Resource Type.
	ResourceType ResourceSharingResourceListResponseResourceType `json:"resource_type,required"`
	// Resource Version.
	ResourceVersion int64 `json:"resource_version,required"`
	// Resource Status.
	Status ResourceSharingResourceListResponseStatus `json:"status,required"`
	JSON   resourceSharingResourceListResponseJSON   `json:"-"`
}

// resourceSharingResourceListResponseJSON contains the JSON metadata for the
// struct [ResourceSharingResourceListResponse]
type resourceSharingResourceListResponseJSON struct {
	ID                apijson.Field
	Created           apijson.Field
	Meta              apijson.Field
	Modified          apijson.Field
	ResourceAccountID apijson.Field
	ResourceID        apijson.Field
	ResourceType      apijson.Field
	ResourceVersion   apijson.Field
	Status            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ResourceSharingResourceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSharingResourceListResponseJSON) RawJSON() string {
	return r.raw
}

// Resource Type.
type ResourceSharingResourceListResponseResourceType string

const (
	ResourceSharingResourceListResponseResourceTypeCustomRuleset ResourceSharingResourceListResponseResourceType = "custom-ruleset"
)

func (r ResourceSharingResourceListResponseResourceType) IsKnown() bool {
	switch r {
	case ResourceSharingResourceListResponseResourceTypeCustomRuleset:
		return true
	}
	return false
}

// Resource Status.
type ResourceSharingResourceListResponseStatus string

const (
	ResourceSharingResourceListResponseStatusActive   ResourceSharingResourceListResponseStatus = "active"
	ResourceSharingResourceListResponseStatusDeleting ResourceSharingResourceListResponseStatus = "deleting"
	ResourceSharingResourceListResponseStatusDeleted  ResourceSharingResourceListResponseStatus = "deleted"
)

func (r ResourceSharingResourceListResponseStatus) IsKnown() bool {
	switch r {
	case ResourceSharingResourceListResponseStatusActive, ResourceSharingResourceListResponseStatusDeleting, ResourceSharingResourceListResponseStatusDeleted:
		return true
	}
	return false
}

type ResourceSharingResourceDeleteResponse struct {
	// Share Resource identifier.
	ID string `json:"id,required"`
	// When the share was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// Resource Metadata.
	Meta interface{} `json:"meta,required"`
	// When the share was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// Account identifier.
	ResourceAccountID string `json:"resource_account_id,required"`
	// Share Resource identifier.
	ResourceID string `json:"resource_id,required"`
	// Resource Type.
	ResourceType ResourceSharingResourceDeleteResponseResourceType `json:"resource_type,required"`
	// Resource Version.
	ResourceVersion int64 `json:"resource_version,required"`
	// Resource Status.
	Status ResourceSharingResourceDeleteResponseStatus `json:"status,required"`
	JSON   resourceSharingResourceDeleteResponseJSON   `json:"-"`
}

// resourceSharingResourceDeleteResponseJSON contains the JSON metadata for the
// struct [ResourceSharingResourceDeleteResponse]
type resourceSharingResourceDeleteResponseJSON struct {
	ID                apijson.Field
	Created           apijson.Field
	Meta              apijson.Field
	Modified          apijson.Field
	ResourceAccountID apijson.Field
	ResourceID        apijson.Field
	ResourceType      apijson.Field
	ResourceVersion   apijson.Field
	Status            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ResourceSharingResourceDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSharingResourceDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Resource Type.
type ResourceSharingResourceDeleteResponseResourceType string

const (
	ResourceSharingResourceDeleteResponseResourceTypeCustomRuleset ResourceSharingResourceDeleteResponseResourceType = "custom-ruleset"
)

func (r ResourceSharingResourceDeleteResponseResourceType) IsKnown() bool {
	switch r {
	case ResourceSharingResourceDeleteResponseResourceTypeCustomRuleset:
		return true
	}
	return false
}

// Resource Status.
type ResourceSharingResourceDeleteResponseStatus string

const (
	ResourceSharingResourceDeleteResponseStatusActive   ResourceSharingResourceDeleteResponseStatus = "active"
	ResourceSharingResourceDeleteResponseStatusDeleting ResourceSharingResourceDeleteResponseStatus = "deleting"
	ResourceSharingResourceDeleteResponseStatusDeleted  ResourceSharingResourceDeleteResponseStatus = "deleted"
)

func (r ResourceSharingResourceDeleteResponseStatus) IsKnown() bool {
	switch r {
	case ResourceSharingResourceDeleteResponseStatusActive, ResourceSharingResourceDeleteResponseStatusDeleting, ResourceSharingResourceDeleteResponseStatusDeleted:
		return true
	}
	return false
}

type ResourceSharingResourceGetResponse struct {
	// Share Resource identifier.
	ID string `json:"id,required"`
	// When the share was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// Resource Metadata.
	Meta interface{} `json:"meta,required"`
	// When the share was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// Account identifier.
	ResourceAccountID string `json:"resource_account_id,required"`
	// Share Resource identifier.
	ResourceID string `json:"resource_id,required"`
	// Resource Type.
	ResourceType ResourceSharingResourceGetResponseResourceType `json:"resource_type,required"`
	// Resource Version.
	ResourceVersion int64 `json:"resource_version,required"`
	// Resource Status.
	Status ResourceSharingResourceGetResponseStatus `json:"status,required"`
	JSON   resourceSharingResourceGetResponseJSON   `json:"-"`
}

// resourceSharingResourceGetResponseJSON contains the JSON metadata for the struct
// [ResourceSharingResourceGetResponse]
type resourceSharingResourceGetResponseJSON struct {
	ID                apijson.Field
	Created           apijson.Field
	Meta              apijson.Field
	Modified          apijson.Field
	ResourceAccountID apijson.Field
	ResourceID        apijson.Field
	ResourceType      apijson.Field
	ResourceVersion   apijson.Field
	Status            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ResourceSharingResourceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSharingResourceGetResponseJSON) RawJSON() string {
	return r.raw
}

// Resource Type.
type ResourceSharingResourceGetResponseResourceType string

const (
	ResourceSharingResourceGetResponseResourceTypeCustomRuleset ResourceSharingResourceGetResponseResourceType = "custom-ruleset"
)

func (r ResourceSharingResourceGetResponseResourceType) IsKnown() bool {
	switch r {
	case ResourceSharingResourceGetResponseResourceTypeCustomRuleset:
		return true
	}
	return false
}

// Resource Status.
type ResourceSharingResourceGetResponseStatus string

const (
	ResourceSharingResourceGetResponseStatusActive   ResourceSharingResourceGetResponseStatus = "active"
	ResourceSharingResourceGetResponseStatusDeleting ResourceSharingResourceGetResponseStatus = "deleting"
	ResourceSharingResourceGetResponseStatusDeleted  ResourceSharingResourceGetResponseStatus = "deleted"
)

func (r ResourceSharingResourceGetResponseStatus) IsKnown() bool {
	switch r {
	case ResourceSharingResourceGetResponseStatusActive, ResourceSharingResourceGetResponseStatusDeleting, ResourceSharingResourceGetResponseStatusDeleted:
		return true
	}
	return false
}

type ResourceSharingResourceNewParams struct {
	// Account identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Resource Metadata.
	Meta param.Field[interface{}] `json:"meta,required"`
	// Account identifier.
	ResourceAccountID param.Field[string] `json:"resource_account_id,required"`
	// Share Resource identifier.
	ResourceID param.Field[string] `json:"resource_id,required"`
	// Resource Type.
	ResourceType param.Field[ResourceSharingResourceNewParamsResourceType] `json:"resource_type,required"`
}

func (r ResourceSharingResourceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Resource Type.
type ResourceSharingResourceNewParamsResourceType string

const (
	ResourceSharingResourceNewParamsResourceTypeCustomRuleset ResourceSharingResourceNewParamsResourceType = "custom-ruleset"
)

func (r ResourceSharingResourceNewParamsResourceType) IsKnown() bool {
	switch r {
	case ResourceSharingResourceNewParamsResourceTypeCustomRuleset:
		return true
	}
	return false
}

type ResourceSharingResourceNewResponseEnvelope struct {
	Errors []shared.ResponseInfo `json:"errors,required"`
	// Whether the API call was successful.
	Success bool                                           `json:"success,required"`
	Result  ResourceSharingResourceNewResponse             `json:"result"`
	JSON    resourceSharingResourceNewResponseEnvelopeJSON `json:"-"`
}

// resourceSharingResourceNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [ResourceSharingResourceNewResponseEnvelope]
type resourceSharingResourceNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceSharingResourceNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSharingResourceNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ResourceSharingResourceUpdateParams struct {
	// Account identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Resource Metadata.
	Meta param.Field[interface{}] `json:"meta,required"`
}

func (r ResourceSharingResourceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ResourceSharingResourceUpdateResponseEnvelope struct {
	Errors []shared.ResponseInfo `json:"errors,required"`
	// Whether the API call was successful.
	Success bool                                              `json:"success,required"`
	Result  ResourceSharingResourceUpdateResponse             `json:"result"`
	JSON    resourceSharingResourceUpdateResponseEnvelopeJSON `json:"-"`
}

// resourceSharingResourceUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [ResourceSharingResourceUpdateResponseEnvelope]
type resourceSharingResourceUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceSharingResourceUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSharingResourceUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ResourceSharingResourceListParams struct {
	// Account identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Page number.
	Page param.Field[int64] `query:"page"`
	// Number of objects to return per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Filter share resources by resource_type.
	ResourceType param.Field[ResourceSharingResourceListParamsResourceType] `query:"resource_type"`
	// Filter share resources by status.
	Status param.Field[ResourceSharingResourceListParamsStatus] `query:"status"`
}

// URLQuery serializes [ResourceSharingResourceListParams]'s query parameters as
// `url.Values`.
func (r ResourceSharingResourceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filter share resources by resource_type.
type ResourceSharingResourceListParamsResourceType string

const (
	ResourceSharingResourceListParamsResourceTypeCustomRuleset ResourceSharingResourceListParamsResourceType = "custom-ruleset"
)

func (r ResourceSharingResourceListParamsResourceType) IsKnown() bool {
	switch r {
	case ResourceSharingResourceListParamsResourceTypeCustomRuleset:
		return true
	}
	return false
}

// Filter share resources by status.
type ResourceSharingResourceListParamsStatus string

const (
	ResourceSharingResourceListParamsStatusActive   ResourceSharingResourceListParamsStatus = "active"
	ResourceSharingResourceListParamsStatusDeleting ResourceSharingResourceListParamsStatus = "deleting"
	ResourceSharingResourceListParamsStatusDeleted  ResourceSharingResourceListParamsStatus = "deleted"
)

func (r ResourceSharingResourceListParamsStatus) IsKnown() bool {
	switch r {
	case ResourceSharingResourceListParamsStatusActive, ResourceSharingResourceListParamsStatusDeleting, ResourceSharingResourceListParamsStatusDeleted:
		return true
	}
	return false
}

type ResourceSharingResourceDeleteParams struct {
	// Account identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ResourceSharingResourceDeleteResponseEnvelope struct {
	Errors []shared.ResponseInfo `json:"errors,required"`
	// Whether the API call was successful.
	Success bool                                              `json:"success,required"`
	Result  ResourceSharingResourceDeleteResponse             `json:"result"`
	JSON    resourceSharingResourceDeleteResponseEnvelopeJSON `json:"-"`
}

// resourceSharingResourceDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [ResourceSharingResourceDeleteResponseEnvelope]
type resourceSharingResourceDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceSharingResourceDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSharingResourceDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ResourceSharingResourceGetParams struct {
	// Account identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ResourceSharingResourceGetResponseEnvelope struct {
	Errors []shared.ResponseInfo `json:"errors,required"`
	// Whether the API call was successful.
	Success bool                                           `json:"success,required"`
	Result  ResourceSharingResourceGetResponse             `json:"result"`
	JSON    resourceSharingResourceGetResponseEnvelopeJSON `json:"-"`
}

// resourceSharingResourceGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ResourceSharingResourceGetResponseEnvelope]
type resourceSharingResourceGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceSharingResourceGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSharingResourceGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
