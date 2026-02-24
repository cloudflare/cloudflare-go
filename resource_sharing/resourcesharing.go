// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package resource_sharing

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
	"github.com/cloudflare/cloudflare-go/v6/shared"
)

// ResourceSharingService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewResourceSharingService] method instead.
type ResourceSharingService struct {
	Options    []option.RequestOption
	Recipients *RecipientService
	Resources  *ResourceService
}

// NewResourceSharingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewResourceSharingService(opts ...option.RequestOption) (r *ResourceSharingService) {
	r = &ResourceSharingService{}
	r.Options = opts
	r.Recipients = NewRecipientService(opts...)
	r.Resources = NewResourceService(opts...)
	return
}

// Creates a new resource share for sharing Cloudflare resources with other
// accounts or organizations.
func (r *ResourceSharingService) New(ctx context.Context, params ResourceSharingNewParams, opts ...option.RequestOption) (res *ResourceSharingNewResponse, err error) {
	var env ResourceSharingNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updating is not immediate, an updated share object with a new status will be
// returned.
func (r *ResourceSharingService) Update(ctx context.Context, shareID string, params ResourceSharingUpdateParams, opts ...option.RequestOption) (res *ResourceSharingUpdateResponse, err error) {
	var env ResourceSharingUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if shareID == "" {
		err = errors.New("missing required share_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares/%s", params.AccountID, shareID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all account shares.
func (r *ResourceSharingService) List(ctx context.Context, params ResourceSharingListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[ResourceSharingListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares", params.AccountID)
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

// Lists all account shares.
func (r *ResourceSharingService) ListAutoPaging(ctx context.Context, params ResourceSharingListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[ResourceSharingListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletion is not immediate, an updated share object with a new status will be
// returned.
func (r *ResourceSharingService) Delete(ctx context.Context, shareID string, body ResourceSharingDeleteParams, opts ...option.RequestOption) (res *ResourceSharingDeleteResponse, err error) {
	var env ResourceSharingDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if shareID == "" {
		err = errors.New("missing required share_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares/%s", body.AccountID, shareID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches share by ID.
func (r *ResourceSharingService) Get(ctx context.Context, shareID string, params ResourceSharingGetParams, opts ...option.RequestOption) (res *ResourceSharingGetResponse, err error) {
	var env ResourceSharingGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if shareID == "" {
		err = errors.New("missing required share_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares/%s", params.AccountID, shareID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ResourceSharingNewResponse struct {
	// Share identifier tag.
	ID string `json:"id,required"`
	// Account identifier.
	AccountID string `json:"account_id,required"`
	// The display name of an account.
	AccountName string `json:"account_name,required"`
	// When the share was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// When the share was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// The name of the share.
	Name string `json:"name,required"`
	// Organization identifier.
	OrganizationID string                               `json:"organization_id,required"`
	Status         ResourceSharingNewResponseStatus     `json:"status,required"`
	TargetType     ResourceSharingNewResponseTargetType `json:"target_type,required"`
	// The number of recipients in the 'associated' state. This field is only included
	// when requested via the 'include_recipient_counts' parameter.
	AssociatedRecipientCount int64 `json:"associated_recipient_count"`
	// The number of recipients in the 'associating' state. This field is only included
	// when requested via the 'include_recipient_counts' parameter.
	AssociatingRecipientCount int64 `json:"associating_recipient_count"`
	// The number of recipients in the 'disassociated' state. This field is only
	// included when requested via the 'include_recipient_counts' parameter.
	DisassociatedRecipientCount int64 `json:"disassociated_recipient_count"`
	// The number of recipients in the 'disassociating' state. This field is only
	// included when requested via the 'include_recipient_counts' parameter.
	DisassociatingRecipientCount int64                          `json:"disassociating_recipient_count"`
	Kind                         ResourceSharingNewResponseKind `json:"kind"`
	// A list of resources that are part of the share. This field is only included when
	// requested via the 'include_resources' parameter.
	Resources []ResourceSharingNewResponseResource `json:"resources"`
	JSON      resourceSharingNewResponseJSON       `json:"-"`
}

// resourceSharingNewResponseJSON contains the JSON metadata for the struct
// [ResourceSharingNewResponse]
type resourceSharingNewResponseJSON struct {
	ID                           apijson.Field
	AccountID                    apijson.Field
	AccountName                  apijson.Field
	Created                      apijson.Field
	Modified                     apijson.Field
	Name                         apijson.Field
	OrganizationID               apijson.Field
	Status                       apijson.Field
	TargetType                   apijson.Field
	AssociatedRecipientCount     apijson.Field
	AssociatingRecipientCount    apijson.Field
	DisassociatedRecipientCount  apijson.Field
	DisassociatingRecipientCount apijson.Field
	Kind                         apijson.Field
	Resources                    apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ResourceSharingNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSharingNewResponseJSON) RawJSON() string {
	return r.raw
}

type ResourceSharingNewResponseStatus string

const (
	ResourceSharingNewResponseStatusActive   ResourceSharingNewResponseStatus = "active"
	ResourceSharingNewResponseStatusDeleting ResourceSharingNewResponseStatus = "deleting"
	ResourceSharingNewResponseStatusDeleted  ResourceSharingNewResponseStatus = "deleted"
)

func (r ResourceSharingNewResponseStatus) IsKnown() bool {
	switch r {
	case ResourceSharingNewResponseStatusActive, ResourceSharingNewResponseStatusDeleting, ResourceSharingNewResponseStatusDeleted:
		return true
	}
	return false
}

type ResourceSharingNewResponseTargetType string

const (
	ResourceSharingNewResponseTargetTypeAccount      ResourceSharingNewResponseTargetType = "account"
	ResourceSharingNewResponseTargetTypeOrganization ResourceSharingNewResponseTargetType = "organization"
)

func (r ResourceSharingNewResponseTargetType) IsKnown() bool {
	switch r {
	case ResourceSharingNewResponseTargetTypeAccount, ResourceSharingNewResponseTargetTypeOrganization:
		return true
	}
	return false
}

type ResourceSharingNewResponseKind string

const (
	ResourceSharingNewResponseKindSent     ResourceSharingNewResponseKind = "sent"
	ResourceSharingNewResponseKindReceived ResourceSharingNewResponseKind = "received"
)

func (r ResourceSharingNewResponseKind) IsKnown() bool {
	switch r {
	case ResourceSharingNewResponseKindSent, ResourceSharingNewResponseKindReceived:
		return true
	}
	return false
}

type ResourceSharingNewResponseResource struct {
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
	ResourceType ResourceSharingNewResponseResourcesResourceType `json:"resource_type,required"`
	// Resource Version.
	ResourceVersion int64 `json:"resource_version,required"`
	// Resource Status.
	Status ResourceSharingNewResponseResourcesStatus `json:"status,required"`
	JSON   resourceSharingNewResponseResourceJSON    `json:"-"`
}

// resourceSharingNewResponseResourceJSON contains the JSON metadata for the struct
// [ResourceSharingNewResponseResource]
type resourceSharingNewResponseResourceJSON struct {
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

func (r *ResourceSharingNewResponseResource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSharingNewResponseResourceJSON) RawJSON() string {
	return r.raw
}

// Resource Type.
type ResourceSharingNewResponseResourcesResourceType string

const (
	ResourceSharingNewResponseResourcesResourceTypeCustomRuleset                ResourceSharingNewResponseResourcesResourceType = "custom-ruleset"
	ResourceSharingNewResponseResourcesResourceTypeWidget                       ResourceSharingNewResponseResourcesResourceType = "widget"
	ResourceSharingNewResponseResourcesResourceTypeGatewayPolicy                ResourceSharingNewResponseResourcesResourceType = "gateway-policy"
	ResourceSharingNewResponseResourcesResourceTypeGatewayDestinationIP         ResourceSharingNewResponseResourcesResourceType = "gateway-destination-ip"
	ResourceSharingNewResponseResourcesResourceTypeGatewayBlockPageSettings     ResourceSharingNewResponseResourcesResourceType = "gateway-block-page-settings"
	ResourceSharingNewResponseResourcesResourceTypeGatewayExtendedEmailMatching ResourceSharingNewResponseResourcesResourceType = "gateway-extended-email-matching"
)

func (r ResourceSharingNewResponseResourcesResourceType) IsKnown() bool {
	switch r {
	case ResourceSharingNewResponseResourcesResourceTypeCustomRuleset, ResourceSharingNewResponseResourcesResourceTypeWidget, ResourceSharingNewResponseResourcesResourceTypeGatewayPolicy, ResourceSharingNewResponseResourcesResourceTypeGatewayDestinationIP, ResourceSharingNewResponseResourcesResourceTypeGatewayBlockPageSettings, ResourceSharingNewResponseResourcesResourceTypeGatewayExtendedEmailMatching:
		return true
	}
	return false
}

// Resource Status.
type ResourceSharingNewResponseResourcesStatus string

const (
	ResourceSharingNewResponseResourcesStatusActive   ResourceSharingNewResponseResourcesStatus = "active"
	ResourceSharingNewResponseResourcesStatusDeleting ResourceSharingNewResponseResourcesStatus = "deleting"
	ResourceSharingNewResponseResourcesStatusDeleted  ResourceSharingNewResponseResourcesStatus = "deleted"
)

func (r ResourceSharingNewResponseResourcesStatus) IsKnown() bool {
	switch r {
	case ResourceSharingNewResponseResourcesStatusActive, ResourceSharingNewResponseResourcesStatusDeleting, ResourceSharingNewResponseResourcesStatusDeleted:
		return true
	}
	return false
}

type ResourceSharingUpdateResponse struct {
	// Share identifier tag.
	ID string `json:"id,required"`
	// Account identifier.
	AccountID string `json:"account_id,required"`
	// The display name of an account.
	AccountName string `json:"account_name,required"`
	// When the share was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// When the share was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// The name of the share.
	Name string `json:"name,required"`
	// Organization identifier.
	OrganizationID string                                  `json:"organization_id,required"`
	Status         ResourceSharingUpdateResponseStatus     `json:"status,required"`
	TargetType     ResourceSharingUpdateResponseTargetType `json:"target_type,required"`
	// The number of recipients in the 'associated' state. This field is only included
	// when requested via the 'include_recipient_counts' parameter.
	AssociatedRecipientCount int64 `json:"associated_recipient_count"`
	// The number of recipients in the 'associating' state. This field is only included
	// when requested via the 'include_recipient_counts' parameter.
	AssociatingRecipientCount int64 `json:"associating_recipient_count"`
	// The number of recipients in the 'disassociated' state. This field is only
	// included when requested via the 'include_recipient_counts' parameter.
	DisassociatedRecipientCount int64 `json:"disassociated_recipient_count"`
	// The number of recipients in the 'disassociating' state. This field is only
	// included when requested via the 'include_recipient_counts' parameter.
	DisassociatingRecipientCount int64                             `json:"disassociating_recipient_count"`
	Kind                         ResourceSharingUpdateResponseKind `json:"kind"`
	// A list of resources that are part of the share. This field is only included when
	// requested via the 'include_resources' parameter.
	Resources []ResourceSharingUpdateResponseResource `json:"resources"`
	JSON      resourceSharingUpdateResponseJSON       `json:"-"`
}

// resourceSharingUpdateResponseJSON contains the JSON metadata for the struct
// [ResourceSharingUpdateResponse]
type resourceSharingUpdateResponseJSON struct {
	ID                           apijson.Field
	AccountID                    apijson.Field
	AccountName                  apijson.Field
	Created                      apijson.Field
	Modified                     apijson.Field
	Name                         apijson.Field
	OrganizationID               apijson.Field
	Status                       apijson.Field
	TargetType                   apijson.Field
	AssociatedRecipientCount     apijson.Field
	AssociatingRecipientCount    apijson.Field
	DisassociatedRecipientCount  apijson.Field
	DisassociatingRecipientCount apijson.Field
	Kind                         apijson.Field
	Resources                    apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ResourceSharingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSharingUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ResourceSharingUpdateResponseStatus string

const (
	ResourceSharingUpdateResponseStatusActive   ResourceSharingUpdateResponseStatus = "active"
	ResourceSharingUpdateResponseStatusDeleting ResourceSharingUpdateResponseStatus = "deleting"
	ResourceSharingUpdateResponseStatusDeleted  ResourceSharingUpdateResponseStatus = "deleted"
)

func (r ResourceSharingUpdateResponseStatus) IsKnown() bool {
	switch r {
	case ResourceSharingUpdateResponseStatusActive, ResourceSharingUpdateResponseStatusDeleting, ResourceSharingUpdateResponseStatusDeleted:
		return true
	}
	return false
}

type ResourceSharingUpdateResponseTargetType string

const (
	ResourceSharingUpdateResponseTargetTypeAccount      ResourceSharingUpdateResponseTargetType = "account"
	ResourceSharingUpdateResponseTargetTypeOrganization ResourceSharingUpdateResponseTargetType = "organization"
)

func (r ResourceSharingUpdateResponseTargetType) IsKnown() bool {
	switch r {
	case ResourceSharingUpdateResponseTargetTypeAccount, ResourceSharingUpdateResponseTargetTypeOrganization:
		return true
	}
	return false
}

type ResourceSharingUpdateResponseKind string

const (
	ResourceSharingUpdateResponseKindSent     ResourceSharingUpdateResponseKind = "sent"
	ResourceSharingUpdateResponseKindReceived ResourceSharingUpdateResponseKind = "received"
)

func (r ResourceSharingUpdateResponseKind) IsKnown() bool {
	switch r {
	case ResourceSharingUpdateResponseKindSent, ResourceSharingUpdateResponseKindReceived:
		return true
	}
	return false
}

type ResourceSharingUpdateResponseResource struct {
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
	ResourceType ResourceSharingUpdateResponseResourcesResourceType `json:"resource_type,required"`
	// Resource Version.
	ResourceVersion int64 `json:"resource_version,required"`
	// Resource Status.
	Status ResourceSharingUpdateResponseResourcesStatus `json:"status,required"`
	JSON   resourceSharingUpdateResponseResourceJSON    `json:"-"`
}

// resourceSharingUpdateResponseResourceJSON contains the JSON metadata for the
// struct [ResourceSharingUpdateResponseResource]
type resourceSharingUpdateResponseResourceJSON struct {
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

func (r *ResourceSharingUpdateResponseResource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSharingUpdateResponseResourceJSON) RawJSON() string {
	return r.raw
}

// Resource Type.
type ResourceSharingUpdateResponseResourcesResourceType string

const (
	ResourceSharingUpdateResponseResourcesResourceTypeCustomRuleset                ResourceSharingUpdateResponseResourcesResourceType = "custom-ruleset"
	ResourceSharingUpdateResponseResourcesResourceTypeWidget                       ResourceSharingUpdateResponseResourcesResourceType = "widget"
	ResourceSharingUpdateResponseResourcesResourceTypeGatewayPolicy                ResourceSharingUpdateResponseResourcesResourceType = "gateway-policy"
	ResourceSharingUpdateResponseResourcesResourceTypeGatewayDestinationIP         ResourceSharingUpdateResponseResourcesResourceType = "gateway-destination-ip"
	ResourceSharingUpdateResponseResourcesResourceTypeGatewayBlockPageSettings     ResourceSharingUpdateResponseResourcesResourceType = "gateway-block-page-settings"
	ResourceSharingUpdateResponseResourcesResourceTypeGatewayExtendedEmailMatching ResourceSharingUpdateResponseResourcesResourceType = "gateway-extended-email-matching"
)

func (r ResourceSharingUpdateResponseResourcesResourceType) IsKnown() bool {
	switch r {
	case ResourceSharingUpdateResponseResourcesResourceTypeCustomRuleset, ResourceSharingUpdateResponseResourcesResourceTypeWidget, ResourceSharingUpdateResponseResourcesResourceTypeGatewayPolicy, ResourceSharingUpdateResponseResourcesResourceTypeGatewayDestinationIP, ResourceSharingUpdateResponseResourcesResourceTypeGatewayBlockPageSettings, ResourceSharingUpdateResponseResourcesResourceTypeGatewayExtendedEmailMatching:
		return true
	}
	return false
}

// Resource Status.
type ResourceSharingUpdateResponseResourcesStatus string

const (
	ResourceSharingUpdateResponseResourcesStatusActive   ResourceSharingUpdateResponseResourcesStatus = "active"
	ResourceSharingUpdateResponseResourcesStatusDeleting ResourceSharingUpdateResponseResourcesStatus = "deleting"
	ResourceSharingUpdateResponseResourcesStatusDeleted  ResourceSharingUpdateResponseResourcesStatus = "deleted"
)

func (r ResourceSharingUpdateResponseResourcesStatus) IsKnown() bool {
	switch r {
	case ResourceSharingUpdateResponseResourcesStatusActive, ResourceSharingUpdateResponseResourcesStatusDeleting, ResourceSharingUpdateResponseResourcesStatusDeleted:
		return true
	}
	return false
}

type ResourceSharingListResponse struct {
	// Share identifier tag.
	ID string `json:"id,required"`
	// Account identifier.
	AccountID string `json:"account_id,required"`
	// The display name of an account.
	AccountName string `json:"account_name,required"`
	// When the share was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// When the share was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// The name of the share.
	Name string `json:"name,required"`
	// Organization identifier.
	OrganizationID string                                `json:"organization_id,required"`
	Status         ResourceSharingListResponseStatus     `json:"status,required"`
	TargetType     ResourceSharingListResponseTargetType `json:"target_type,required"`
	// The number of recipients in the 'associated' state. This field is only included
	// when requested via the 'include_recipient_counts' parameter.
	AssociatedRecipientCount int64 `json:"associated_recipient_count"`
	// The number of recipients in the 'associating' state. This field is only included
	// when requested via the 'include_recipient_counts' parameter.
	AssociatingRecipientCount int64 `json:"associating_recipient_count"`
	// The number of recipients in the 'disassociated' state. This field is only
	// included when requested via the 'include_recipient_counts' parameter.
	DisassociatedRecipientCount int64 `json:"disassociated_recipient_count"`
	// The number of recipients in the 'disassociating' state. This field is only
	// included when requested via the 'include_recipient_counts' parameter.
	DisassociatingRecipientCount int64                           `json:"disassociating_recipient_count"`
	Kind                         ResourceSharingListResponseKind `json:"kind"`
	// A list of resources that are part of the share. This field is only included when
	// requested via the 'include_resources' parameter.
	Resources []ResourceSharingListResponseResource `json:"resources"`
	JSON      resourceSharingListResponseJSON       `json:"-"`
}

// resourceSharingListResponseJSON contains the JSON metadata for the struct
// [ResourceSharingListResponse]
type resourceSharingListResponseJSON struct {
	ID                           apijson.Field
	AccountID                    apijson.Field
	AccountName                  apijson.Field
	Created                      apijson.Field
	Modified                     apijson.Field
	Name                         apijson.Field
	OrganizationID               apijson.Field
	Status                       apijson.Field
	TargetType                   apijson.Field
	AssociatedRecipientCount     apijson.Field
	AssociatingRecipientCount    apijson.Field
	DisassociatedRecipientCount  apijson.Field
	DisassociatingRecipientCount apijson.Field
	Kind                         apijson.Field
	Resources                    apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ResourceSharingListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSharingListResponseJSON) RawJSON() string {
	return r.raw
}

type ResourceSharingListResponseStatus string

const (
	ResourceSharingListResponseStatusActive   ResourceSharingListResponseStatus = "active"
	ResourceSharingListResponseStatusDeleting ResourceSharingListResponseStatus = "deleting"
	ResourceSharingListResponseStatusDeleted  ResourceSharingListResponseStatus = "deleted"
)

func (r ResourceSharingListResponseStatus) IsKnown() bool {
	switch r {
	case ResourceSharingListResponseStatusActive, ResourceSharingListResponseStatusDeleting, ResourceSharingListResponseStatusDeleted:
		return true
	}
	return false
}

type ResourceSharingListResponseTargetType string

const (
	ResourceSharingListResponseTargetTypeAccount      ResourceSharingListResponseTargetType = "account"
	ResourceSharingListResponseTargetTypeOrganization ResourceSharingListResponseTargetType = "organization"
)

func (r ResourceSharingListResponseTargetType) IsKnown() bool {
	switch r {
	case ResourceSharingListResponseTargetTypeAccount, ResourceSharingListResponseTargetTypeOrganization:
		return true
	}
	return false
}

type ResourceSharingListResponseKind string

const (
	ResourceSharingListResponseKindSent     ResourceSharingListResponseKind = "sent"
	ResourceSharingListResponseKindReceived ResourceSharingListResponseKind = "received"
)

func (r ResourceSharingListResponseKind) IsKnown() bool {
	switch r {
	case ResourceSharingListResponseKindSent, ResourceSharingListResponseKindReceived:
		return true
	}
	return false
}

type ResourceSharingListResponseResource struct {
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
	ResourceType ResourceSharingListResponseResourcesResourceType `json:"resource_type,required"`
	// Resource Version.
	ResourceVersion int64 `json:"resource_version,required"`
	// Resource Status.
	Status ResourceSharingListResponseResourcesStatus `json:"status,required"`
	JSON   resourceSharingListResponseResourceJSON    `json:"-"`
}

// resourceSharingListResponseResourceJSON contains the JSON metadata for the
// struct [ResourceSharingListResponseResource]
type resourceSharingListResponseResourceJSON struct {
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

func (r *ResourceSharingListResponseResource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSharingListResponseResourceJSON) RawJSON() string {
	return r.raw
}

// Resource Type.
type ResourceSharingListResponseResourcesResourceType string

const (
	ResourceSharingListResponseResourcesResourceTypeCustomRuleset                ResourceSharingListResponseResourcesResourceType = "custom-ruleset"
	ResourceSharingListResponseResourcesResourceTypeWidget                       ResourceSharingListResponseResourcesResourceType = "widget"
	ResourceSharingListResponseResourcesResourceTypeGatewayPolicy                ResourceSharingListResponseResourcesResourceType = "gateway-policy"
	ResourceSharingListResponseResourcesResourceTypeGatewayDestinationIP         ResourceSharingListResponseResourcesResourceType = "gateway-destination-ip"
	ResourceSharingListResponseResourcesResourceTypeGatewayBlockPageSettings     ResourceSharingListResponseResourcesResourceType = "gateway-block-page-settings"
	ResourceSharingListResponseResourcesResourceTypeGatewayExtendedEmailMatching ResourceSharingListResponseResourcesResourceType = "gateway-extended-email-matching"
)

func (r ResourceSharingListResponseResourcesResourceType) IsKnown() bool {
	switch r {
	case ResourceSharingListResponseResourcesResourceTypeCustomRuleset, ResourceSharingListResponseResourcesResourceTypeWidget, ResourceSharingListResponseResourcesResourceTypeGatewayPolicy, ResourceSharingListResponseResourcesResourceTypeGatewayDestinationIP, ResourceSharingListResponseResourcesResourceTypeGatewayBlockPageSettings, ResourceSharingListResponseResourcesResourceTypeGatewayExtendedEmailMatching:
		return true
	}
	return false
}

// Resource Status.
type ResourceSharingListResponseResourcesStatus string

const (
	ResourceSharingListResponseResourcesStatusActive   ResourceSharingListResponseResourcesStatus = "active"
	ResourceSharingListResponseResourcesStatusDeleting ResourceSharingListResponseResourcesStatus = "deleting"
	ResourceSharingListResponseResourcesStatusDeleted  ResourceSharingListResponseResourcesStatus = "deleted"
)

func (r ResourceSharingListResponseResourcesStatus) IsKnown() bool {
	switch r {
	case ResourceSharingListResponseResourcesStatusActive, ResourceSharingListResponseResourcesStatusDeleting, ResourceSharingListResponseResourcesStatusDeleted:
		return true
	}
	return false
}

type ResourceSharingDeleteResponse struct {
	// Share identifier tag.
	ID string `json:"id,required"`
	// Account identifier.
	AccountID string `json:"account_id,required"`
	// The display name of an account.
	AccountName string `json:"account_name,required"`
	// When the share was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// When the share was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// The name of the share.
	Name string `json:"name,required"`
	// Organization identifier.
	OrganizationID string                                  `json:"organization_id,required"`
	Status         ResourceSharingDeleteResponseStatus     `json:"status,required"`
	TargetType     ResourceSharingDeleteResponseTargetType `json:"target_type,required"`
	// The number of recipients in the 'associated' state. This field is only included
	// when requested via the 'include_recipient_counts' parameter.
	AssociatedRecipientCount int64 `json:"associated_recipient_count"`
	// The number of recipients in the 'associating' state. This field is only included
	// when requested via the 'include_recipient_counts' parameter.
	AssociatingRecipientCount int64 `json:"associating_recipient_count"`
	// The number of recipients in the 'disassociated' state. This field is only
	// included when requested via the 'include_recipient_counts' parameter.
	DisassociatedRecipientCount int64 `json:"disassociated_recipient_count"`
	// The number of recipients in the 'disassociating' state. This field is only
	// included when requested via the 'include_recipient_counts' parameter.
	DisassociatingRecipientCount int64                             `json:"disassociating_recipient_count"`
	Kind                         ResourceSharingDeleteResponseKind `json:"kind"`
	// A list of resources that are part of the share. This field is only included when
	// requested via the 'include_resources' parameter.
	Resources []ResourceSharingDeleteResponseResource `json:"resources"`
	JSON      resourceSharingDeleteResponseJSON       `json:"-"`
}

// resourceSharingDeleteResponseJSON contains the JSON metadata for the struct
// [ResourceSharingDeleteResponse]
type resourceSharingDeleteResponseJSON struct {
	ID                           apijson.Field
	AccountID                    apijson.Field
	AccountName                  apijson.Field
	Created                      apijson.Field
	Modified                     apijson.Field
	Name                         apijson.Field
	OrganizationID               apijson.Field
	Status                       apijson.Field
	TargetType                   apijson.Field
	AssociatedRecipientCount     apijson.Field
	AssociatingRecipientCount    apijson.Field
	DisassociatedRecipientCount  apijson.Field
	DisassociatingRecipientCount apijson.Field
	Kind                         apijson.Field
	Resources                    apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ResourceSharingDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSharingDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ResourceSharingDeleteResponseStatus string

const (
	ResourceSharingDeleteResponseStatusActive   ResourceSharingDeleteResponseStatus = "active"
	ResourceSharingDeleteResponseStatusDeleting ResourceSharingDeleteResponseStatus = "deleting"
	ResourceSharingDeleteResponseStatusDeleted  ResourceSharingDeleteResponseStatus = "deleted"
)

func (r ResourceSharingDeleteResponseStatus) IsKnown() bool {
	switch r {
	case ResourceSharingDeleteResponseStatusActive, ResourceSharingDeleteResponseStatusDeleting, ResourceSharingDeleteResponseStatusDeleted:
		return true
	}
	return false
}

type ResourceSharingDeleteResponseTargetType string

const (
	ResourceSharingDeleteResponseTargetTypeAccount      ResourceSharingDeleteResponseTargetType = "account"
	ResourceSharingDeleteResponseTargetTypeOrganization ResourceSharingDeleteResponseTargetType = "organization"
)

func (r ResourceSharingDeleteResponseTargetType) IsKnown() bool {
	switch r {
	case ResourceSharingDeleteResponseTargetTypeAccount, ResourceSharingDeleteResponseTargetTypeOrganization:
		return true
	}
	return false
}

type ResourceSharingDeleteResponseKind string

const (
	ResourceSharingDeleteResponseKindSent     ResourceSharingDeleteResponseKind = "sent"
	ResourceSharingDeleteResponseKindReceived ResourceSharingDeleteResponseKind = "received"
)

func (r ResourceSharingDeleteResponseKind) IsKnown() bool {
	switch r {
	case ResourceSharingDeleteResponseKindSent, ResourceSharingDeleteResponseKindReceived:
		return true
	}
	return false
}

type ResourceSharingDeleteResponseResource struct {
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
	ResourceType ResourceSharingDeleteResponseResourcesResourceType `json:"resource_type,required"`
	// Resource Version.
	ResourceVersion int64 `json:"resource_version,required"`
	// Resource Status.
	Status ResourceSharingDeleteResponseResourcesStatus `json:"status,required"`
	JSON   resourceSharingDeleteResponseResourceJSON    `json:"-"`
}

// resourceSharingDeleteResponseResourceJSON contains the JSON metadata for the
// struct [ResourceSharingDeleteResponseResource]
type resourceSharingDeleteResponseResourceJSON struct {
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

func (r *ResourceSharingDeleteResponseResource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSharingDeleteResponseResourceJSON) RawJSON() string {
	return r.raw
}

// Resource Type.
type ResourceSharingDeleteResponseResourcesResourceType string

const (
	ResourceSharingDeleteResponseResourcesResourceTypeCustomRuleset                ResourceSharingDeleteResponseResourcesResourceType = "custom-ruleset"
	ResourceSharingDeleteResponseResourcesResourceTypeWidget                       ResourceSharingDeleteResponseResourcesResourceType = "widget"
	ResourceSharingDeleteResponseResourcesResourceTypeGatewayPolicy                ResourceSharingDeleteResponseResourcesResourceType = "gateway-policy"
	ResourceSharingDeleteResponseResourcesResourceTypeGatewayDestinationIP         ResourceSharingDeleteResponseResourcesResourceType = "gateway-destination-ip"
	ResourceSharingDeleteResponseResourcesResourceTypeGatewayBlockPageSettings     ResourceSharingDeleteResponseResourcesResourceType = "gateway-block-page-settings"
	ResourceSharingDeleteResponseResourcesResourceTypeGatewayExtendedEmailMatching ResourceSharingDeleteResponseResourcesResourceType = "gateway-extended-email-matching"
)

func (r ResourceSharingDeleteResponseResourcesResourceType) IsKnown() bool {
	switch r {
	case ResourceSharingDeleteResponseResourcesResourceTypeCustomRuleset, ResourceSharingDeleteResponseResourcesResourceTypeWidget, ResourceSharingDeleteResponseResourcesResourceTypeGatewayPolicy, ResourceSharingDeleteResponseResourcesResourceTypeGatewayDestinationIP, ResourceSharingDeleteResponseResourcesResourceTypeGatewayBlockPageSettings, ResourceSharingDeleteResponseResourcesResourceTypeGatewayExtendedEmailMatching:
		return true
	}
	return false
}

// Resource Status.
type ResourceSharingDeleteResponseResourcesStatus string

const (
	ResourceSharingDeleteResponseResourcesStatusActive   ResourceSharingDeleteResponseResourcesStatus = "active"
	ResourceSharingDeleteResponseResourcesStatusDeleting ResourceSharingDeleteResponseResourcesStatus = "deleting"
	ResourceSharingDeleteResponseResourcesStatusDeleted  ResourceSharingDeleteResponseResourcesStatus = "deleted"
)

func (r ResourceSharingDeleteResponseResourcesStatus) IsKnown() bool {
	switch r {
	case ResourceSharingDeleteResponseResourcesStatusActive, ResourceSharingDeleteResponseResourcesStatusDeleting, ResourceSharingDeleteResponseResourcesStatusDeleted:
		return true
	}
	return false
}

type ResourceSharingGetResponse struct {
	// Share identifier tag.
	ID string `json:"id,required"`
	// Account identifier.
	AccountID string `json:"account_id,required"`
	// The display name of an account.
	AccountName string `json:"account_name,required"`
	// When the share was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// When the share was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// The name of the share.
	Name string `json:"name,required"`
	// Organization identifier.
	OrganizationID string                               `json:"organization_id,required"`
	Status         ResourceSharingGetResponseStatus     `json:"status,required"`
	TargetType     ResourceSharingGetResponseTargetType `json:"target_type,required"`
	// The number of recipients in the 'associated' state. This field is only included
	// when requested via the 'include_recipient_counts' parameter.
	AssociatedRecipientCount int64 `json:"associated_recipient_count"`
	// The number of recipients in the 'associating' state. This field is only included
	// when requested via the 'include_recipient_counts' parameter.
	AssociatingRecipientCount int64 `json:"associating_recipient_count"`
	// The number of recipients in the 'disassociated' state. This field is only
	// included when requested via the 'include_recipient_counts' parameter.
	DisassociatedRecipientCount int64 `json:"disassociated_recipient_count"`
	// The number of recipients in the 'disassociating' state. This field is only
	// included when requested via the 'include_recipient_counts' parameter.
	DisassociatingRecipientCount int64                          `json:"disassociating_recipient_count"`
	Kind                         ResourceSharingGetResponseKind `json:"kind"`
	// A list of resources that are part of the share. This field is only included when
	// requested via the 'include_resources' parameter.
	Resources []ResourceSharingGetResponseResource `json:"resources"`
	JSON      resourceSharingGetResponseJSON       `json:"-"`
}

// resourceSharingGetResponseJSON contains the JSON metadata for the struct
// [ResourceSharingGetResponse]
type resourceSharingGetResponseJSON struct {
	ID                           apijson.Field
	AccountID                    apijson.Field
	AccountName                  apijson.Field
	Created                      apijson.Field
	Modified                     apijson.Field
	Name                         apijson.Field
	OrganizationID               apijson.Field
	Status                       apijson.Field
	TargetType                   apijson.Field
	AssociatedRecipientCount     apijson.Field
	AssociatingRecipientCount    apijson.Field
	DisassociatedRecipientCount  apijson.Field
	DisassociatingRecipientCount apijson.Field
	Kind                         apijson.Field
	Resources                    apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ResourceSharingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSharingGetResponseJSON) RawJSON() string {
	return r.raw
}

type ResourceSharingGetResponseStatus string

const (
	ResourceSharingGetResponseStatusActive   ResourceSharingGetResponseStatus = "active"
	ResourceSharingGetResponseStatusDeleting ResourceSharingGetResponseStatus = "deleting"
	ResourceSharingGetResponseStatusDeleted  ResourceSharingGetResponseStatus = "deleted"
)

func (r ResourceSharingGetResponseStatus) IsKnown() bool {
	switch r {
	case ResourceSharingGetResponseStatusActive, ResourceSharingGetResponseStatusDeleting, ResourceSharingGetResponseStatusDeleted:
		return true
	}
	return false
}

type ResourceSharingGetResponseTargetType string

const (
	ResourceSharingGetResponseTargetTypeAccount      ResourceSharingGetResponseTargetType = "account"
	ResourceSharingGetResponseTargetTypeOrganization ResourceSharingGetResponseTargetType = "organization"
)

func (r ResourceSharingGetResponseTargetType) IsKnown() bool {
	switch r {
	case ResourceSharingGetResponseTargetTypeAccount, ResourceSharingGetResponseTargetTypeOrganization:
		return true
	}
	return false
}

type ResourceSharingGetResponseKind string

const (
	ResourceSharingGetResponseKindSent     ResourceSharingGetResponseKind = "sent"
	ResourceSharingGetResponseKindReceived ResourceSharingGetResponseKind = "received"
)

func (r ResourceSharingGetResponseKind) IsKnown() bool {
	switch r {
	case ResourceSharingGetResponseKindSent, ResourceSharingGetResponseKindReceived:
		return true
	}
	return false
}

type ResourceSharingGetResponseResource struct {
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
	ResourceType ResourceSharingGetResponseResourcesResourceType `json:"resource_type,required"`
	// Resource Version.
	ResourceVersion int64 `json:"resource_version,required"`
	// Resource Status.
	Status ResourceSharingGetResponseResourcesStatus `json:"status,required"`
	JSON   resourceSharingGetResponseResourceJSON    `json:"-"`
}

// resourceSharingGetResponseResourceJSON contains the JSON metadata for the struct
// [ResourceSharingGetResponseResource]
type resourceSharingGetResponseResourceJSON struct {
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

func (r *ResourceSharingGetResponseResource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSharingGetResponseResourceJSON) RawJSON() string {
	return r.raw
}

// Resource Type.
type ResourceSharingGetResponseResourcesResourceType string

const (
	ResourceSharingGetResponseResourcesResourceTypeCustomRuleset                ResourceSharingGetResponseResourcesResourceType = "custom-ruleset"
	ResourceSharingGetResponseResourcesResourceTypeWidget                       ResourceSharingGetResponseResourcesResourceType = "widget"
	ResourceSharingGetResponseResourcesResourceTypeGatewayPolicy                ResourceSharingGetResponseResourcesResourceType = "gateway-policy"
	ResourceSharingGetResponseResourcesResourceTypeGatewayDestinationIP         ResourceSharingGetResponseResourcesResourceType = "gateway-destination-ip"
	ResourceSharingGetResponseResourcesResourceTypeGatewayBlockPageSettings     ResourceSharingGetResponseResourcesResourceType = "gateway-block-page-settings"
	ResourceSharingGetResponseResourcesResourceTypeGatewayExtendedEmailMatching ResourceSharingGetResponseResourcesResourceType = "gateway-extended-email-matching"
)

func (r ResourceSharingGetResponseResourcesResourceType) IsKnown() bool {
	switch r {
	case ResourceSharingGetResponseResourcesResourceTypeCustomRuleset, ResourceSharingGetResponseResourcesResourceTypeWidget, ResourceSharingGetResponseResourcesResourceTypeGatewayPolicy, ResourceSharingGetResponseResourcesResourceTypeGatewayDestinationIP, ResourceSharingGetResponseResourcesResourceTypeGatewayBlockPageSettings, ResourceSharingGetResponseResourcesResourceTypeGatewayExtendedEmailMatching:
		return true
	}
	return false
}

// Resource Status.
type ResourceSharingGetResponseResourcesStatus string

const (
	ResourceSharingGetResponseResourcesStatusActive   ResourceSharingGetResponseResourcesStatus = "active"
	ResourceSharingGetResponseResourcesStatusDeleting ResourceSharingGetResponseResourcesStatus = "deleting"
	ResourceSharingGetResponseResourcesStatusDeleted  ResourceSharingGetResponseResourcesStatus = "deleted"
)

func (r ResourceSharingGetResponseResourcesStatus) IsKnown() bool {
	switch r {
	case ResourceSharingGetResponseResourcesStatusActive, ResourceSharingGetResponseResourcesStatusDeleting, ResourceSharingGetResponseResourcesStatusDeleted:
		return true
	}
	return false
}

type ResourceSharingNewParams struct {
	// Account identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of the share.
	Name       param.Field[string]                              `json:"name,required"`
	Recipients param.Field[[]ResourceSharingNewParamsRecipient] `json:"recipients,required"`
	Resources  param.Field[[]ResourceSharingNewParamsResource]  `json:"resources,required"`
}

func (r ResourceSharingNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Account or organization ID must be provided.
type ResourceSharingNewParamsRecipient struct {
	// Account identifier.
	AccountID param.Field[string] `json:"account_id"`
	// Organization identifier.
	OrganizationID param.Field[string] `json:"organization_id"`
}

func (r ResourceSharingNewParamsRecipient) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ResourceSharingNewParamsResource struct {
	// Resource Metadata.
	Meta param.Field[interface{}] `json:"meta,required"`
	// Account identifier.
	ResourceAccountID param.Field[string] `json:"resource_account_id,required"`
	// Share Resource identifier.
	ResourceID param.Field[string] `json:"resource_id,required"`
	// Resource Type.
	ResourceType param.Field[ResourceSharingNewParamsResourcesResourceType] `json:"resource_type,required"`
}

func (r ResourceSharingNewParamsResource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Resource Type.
type ResourceSharingNewParamsResourcesResourceType string

const (
	ResourceSharingNewParamsResourcesResourceTypeCustomRuleset                ResourceSharingNewParamsResourcesResourceType = "custom-ruleset"
	ResourceSharingNewParamsResourcesResourceTypeWidget                       ResourceSharingNewParamsResourcesResourceType = "widget"
	ResourceSharingNewParamsResourcesResourceTypeGatewayPolicy                ResourceSharingNewParamsResourcesResourceType = "gateway-policy"
	ResourceSharingNewParamsResourcesResourceTypeGatewayDestinationIP         ResourceSharingNewParamsResourcesResourceType = "gateway-destination-ip"
	ResourceSharingNewParamsResourcesResourceTypeGatewayBlockPageSettings     ResourceSharingNewParamsResourcesResourceType = "gateway-block-page-settings"
	ResourceSharingNewParamsResourcesResourceTypeGatewayExtendedEmailMatching ResourceSharingNewParamsResourcesResourceType = "gateway-extended-email-matching"
)

func (r ResourceSharingNewParamsResourcesResourceType) IsKnown() bool {
	switch r {
	case ResourceSharingNewParamsResourcesResourceTypeCustomRuleset, ResourceSharingNewParamsResourcesResourceTypeWidget, ResourceSharingNewParamsResourcesResourceTypeGatewayPolicy, ResourceSharingNewParamsResourcesResourceTypeGatewayDestinationIP, ResourceSharingNewParamsResourcesResourceTypeGatewayBlockPageSettings, ResourceSharingNewParamsResourcesResourceTypeGatewayExtendedEmailMatching:
		return true
	}
	return false
}

type ResourceSharingNewResponseEnvelope struct {
	Errors []shared.ResponseInfo `json:"errors,required"`
	// Whether the API call was successful.
	Success bool                                   `json:"success,required"`
	Result  ResourceSharingNewResponse             `json:"result"`
	JSON    resourceSharingNewResponseEnvelopeJSON `json:"-"`
}

// resourceSharingNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ResourceSharingNewResponseEnvelope]
type resourceSharingNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceSharingNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSharingNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ResourceSharingUpdateParams struct {
	// Account identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of the share.
	Name param.Field[string] `json:"name,required"`
}

func (r ResourceSharingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ResourceSharingUpdateResponseEnvelope struct {
	Errors []shared.ResponseInfo `json:"errors,required"`
	// Whether the API call was successful.
	Success bool                                      `json:"success,required"`
	Result  ResourceSharingUpdateResponse             `json:"result"`
	JSON    resourceSharingUpdateResponseEnvelopeJSON `json:"-"`
}

// resourceSharingUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [ResourceSharingUpdateResponseEnvelope]
type resourceSharingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceSharingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSharingUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ResourceSharingListParams struct {
	// Account identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Direction to sort objects.
	Direction param.Field[ResourceSharingListParamsDirection] `query:"direction"`
	// Include recipient counts in the response.
	IncludeRecipientCounts param.Field[bool] `query:"include_recipient_counts"`
	// Include resources in the response.
	IncludeResources param.Field[bool] `query:"include_resources"`
	// Filter shares by kind.
	Kind param.Field[ResourceSharingListParamsKind] `query:"kind"`
	// Order shares by values in the given field.
	Order param.Field[ResourceSharingListParamsOrder] `query:"order"`
	// Page number.
	Page param.Field[int64] `query:"page"`
	// Number of objects to return per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Filter share resources by resource_types.
	ResourceTypes param.Field[[]ResourceSharingListParamsResourceType] `query:"resource_types"`
	// Filter shares by status.
	Status param.Field[ResourceSharingListParamsStatus] `query:"status"`
	// Filter shares by target_type.
	TargetType param.Field[ResourceSharingListParamsTargetType] `query:"target_type"`
}

// URLQuery serializes [ResourceSharingListParams]'s query parameters as
// `url.Values`.
func (r ResourceSharingListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Direction to sort objects.
type ResourceSharingListParamsDirection string

const (
	ResourceSharingListParamsDirectionAsc  ResourceSharingListParamsDirection = "asc"
	ResourceSharingListParamsDirectionDesc ResourceSharingListParamsDirection = "desc"
)

func (r ResourceSharingListParamsDirection) IsKnown() bool {
	switch r {
	case ResourceSharingListParamsDirectionAsc, ResourceSharingListParamsDirectionDesc:
		return true
	}
	return false
}

// Filter shares by kind.
type ResourceSharingListParamsKind string

const (
	ResourceSharingListParamsKindSent     ResourceSharingListParamsKind = "sent"
	ResourceSharingListParamsKindReceived ResourceSharingListParamsKind = "received"
)

func (r ResourceSharingListParamsKind) IsKnown() bool {
	switch r {
	case ResourceSharingListParamsKindSent, ResourceSharingListParamsKindReceived:
		return true
	}
	return false
}

// Order shares by values in the given field.
type ResourceSharingListParamsOrder string

const (
	ResourceSharingListParamsOrderName    ResourceSharingListParamsOrder = "name"
	ResourceSharingListParamsOrderCreated ResourceSharingListParamsOrder = "created"
)

func (r ResourceSharingListParamsOrder) IsKnown() bool {
	switch r {
	case ResourceSharingListParamsOrderName, ResourceSharingListParamsOrderCreated:
		return true
	}
	return false
}

// Resource Type.
type ResourceSharingListParamsResourceType string

const (
	ResourceSharingListParamsResourceTypeCustomRuleset                ResourceSharingListParamsResourceType = "custom-ruleset"
	ResourceSharingListParamsResourceTypeWidget                       ResourceSharingListParamsResourceType = "widget"
	ResourceSharingListParamsResourceTypeGatewayPolicy                ResourceSharingListParamsResourceType = "gateway-policy"
	ResourceSharingListParamsResourceTypeGatewayDestinationIP         ResourceSharingListParamsResourceType = "gateway-destination-ip"
	ResourceSharingListParamsResourceTypeGatewayBlockPageSettings     ResourceSharingListParamsResourceType = "gateway-block-page-settings"
	ResourceSharingListParamsResourceTypeGatewayExtendedEmailMatching ResourceSharingListParamsResourceType = "gateway-extended-email-matching"
)

func (r ResourceSharingListParamsResourceType) IsKnown() bool {
	switch r {
	case ResourceSharingListParamsResourceTypeCustomRuleset, ResourceSharingListParamsResourceTypeWidget, ResourceSharingListParamsResourceTypeGatewayPolicy, ResourceSharingListParamsResourceTypeGatewayDestinationIP, ResourceSharingListParamsResourceTypeGatewayBlockPageSettings, ResourceSharingListParamsResourceTypeGatewayExtendedEmailMatching:
		return true
	}
	return false
}

// Filter shares by status.
type ResourceSharingListParamsStatus string

const (
	ResourceSharingListParamsStatusActive   ResourceSharingListParamsStatus = "active"
	ResourceSharingListParamsStatusDeleting ResourceSharingListParamsStatus = "deleting"
	ResourceSharingListParamsStatusDeleted  ResourceSharingListParamsStatus = "deleted"
)

func (r ResourceSharingListParamsStatus) IsKnown() bool {
	switch r {
	case ResourceSharingListParamsStatusActive, ResourceSharingListParamsStatusDeleting, ResourceSharingListParamsStatusDeleted:
		return true
	}
	return false
}

// Filter shares by target_type.
type ResourceSharingListParamsTargetType string

const (
	ResourceSharingListParamsTargetTypeAccount      ResourceSharingListParamsTargetType = "account"
	ResourceSharingListParamsTargetTypeOrganization ResourceSharingListParamsTargetType = "organization"
)

func (r ResourceSharingListParamsTargetType) IsKnown() bool {
	switch r {
	case ResourceSharingListParamsTargetTypeAccount, ResourceSharingListParamsTargetTypeOrganization:
		return true
	}
	return false
}

type ResourceSharingDeleteParams struct {
	// Account identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ResourceSharingDeleteResponseEnvelope struct {
	Errors []shared.ResponseInfo `json:"errors,required"`
	// Whether the API call was successful.
	Success bool                                      `json:"success,required"`
	Result  ResourceSharingDeleteResponse             `json:"result"`
	JSON    resourceSharingDeleteResponseEnvelopeJSON `json:"-"`
}

// resourceSharingDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [ResourceSharingDeleteResponseEnvelope]
type resourceSharingDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceSharingDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSharingDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ResourceSharingGetParams struct {
	// Account identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Include recipient counts in the response.
	IncludeRecipientCounts param.Field[bool] `query:"include_recipient_counts"`
	// Include resources in the response.
	IncludeResources param.Field[bool] `query:"include_resources"`
}

// URLQuery serializes [ResourceSharingGetParams]'s query parameters as
// `url.Values`.
func (r ResourceSharingGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ResourceSharingGetResponseEnvelope struct {
	Errors []shared.ResponseInfo `json:"errors,required"`
	// Whether the API call was successful.
	Success bool                                   `json:"success,required"`
	Result  ResourceSharingGetResponse             `json:"result"`
	JSON    resourceSharingGetResponseEnvelopeJSON `json:"-"`
}

// resourceSharingGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ResourceSharingGetResponseEnvelope]
type resourceSharingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceSharingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSharingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
