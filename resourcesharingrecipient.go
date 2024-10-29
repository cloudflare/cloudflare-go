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

// ResourceSharingRecipientService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewResourceSharingRecipientService] method instead.
type ResourceSharingRecipientService struct {
	Options []option.RequestOption
}

// NewResourceSharingRecipientService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewResourceSharingRecipientService(opts ...option.RequestOption) (r *ResourceSharingRecipientService) {
	r = &ResourceSharingRecipientService{}
	r.Options = opts
	return
}

// Create a new share recipient
func (r *ResourceSharingRecipientService) New(ctx context.Context, shareIdentifier string, params ResourceSharingRecipientNewParams, opts ...option.RequestOption) (res *ResourceSharingRecipientNewResponse, err error) {
	var env ResourceSharingRecipientNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.PathAccountID.Value == "" {
		err = errors.New("missing required path_account_id parameter")
		return
	}
	if shareIdentifier == "" {
		err = errors.New("missing required share_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares/%s/recipients", params.PathAccountID, shareIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List share recipients by share ID.
func (r *ResourceSharingRecipientService) List(ctx context.Context, shareIdentifier string, params ResourceSharingRecipientListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[ResourceSharingRecipientListResponse], err error) {
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
	path := fmt.Sprintf("accounts/%s/shares/%s/recipients", params.AccountID, shareIdentifier)
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

// List share recipients by share ID.
func (r *ResourceSharingRecipientService) ListAutoPaging(ctx context.Context, shareIdentifier string, params ResourceSharingRecipientListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[ResourceSharingRecipientListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, shareIdentifier, params, opts...))
}

// Deletion is not immediate, an updated share recipient object with a new status
// will be returned.
func (r *ResourceSharingRecipientService) Delete(ctx context.Context, shareIdentifier string, recipientIdentifier string, body ResourceSharingRecipientDeleteParams, opts ...option.RequestOption) (res *ResourceSharingRecipientDeleteResponse, err error) {
	var env ResourceSharingRecipientDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if shareIdentifier == "" {
		err = errors.New("missing required share_identifier parameter")
		return
	}
	if recipientIdentifier == "" {
		err = errors.New("missing required recipient_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares/%s/recipients/%s", body.AccountID, shareIdentifier, recipientIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get share recipient by ID.
func (r *ResourceSharingRecipientService) Get(ctx context.Context, shareIdentifier string, recipientIdentifier string, query ResourceSharingRecipientGetParams, opts ...option.RequestOption) (res *ResourceSharingRecipientGetResponse, err error) {
	var env ResourceSharingRecipientGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if shareIdentifier == "" {
		err = errors.New("missing required share_identifier parameter")
		return
	}
	if recipientIdentifier == "" {
		err = errors.New("missing required recipient_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares/%s/recipients/%s", query.AccountID, shareIdentifier, recipientIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ResourceSharingRecipientNewResponse struct {
	// Share Recipient identifier tag.
	ID string `json:"id,required"`
	// Account identifier.
	AccountID string `json:"account_id,required"`
	// Share Recipient association status.
	AssociationStatus ResourceSharingRecipientNewResponseAssociationStatus `json:"association_status,required"`
	// When the share was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// When the share was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// Share Recipient status message.
	StatusMessage string                                  `json:"status_message,required"`
	JSON          resourceSharingRecipientNewResponseJSON `json:"-"`
}

// resourceSharingRecipientNewResponseJSON contains the JSON metadata for the
// struct [ResourceSharingRecipientNewResponse]
type resourceSharingRecipientNewResponseJSON struct {
	ID                apijson.Field
	AccountID         apijson.Field
	AssociationStatus apijson.Field
	Created           apijson.Field
	Modified          apijson.Field
	StatusMessage     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ResourceSharingRecipientNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSharingRecipientNewResponseJSON) RawJSON() string {
	return r.raw
}

// Share Recipient association status.
type ResourceSharingRecipientNewResponseAssociationStatus string

const (
	ResourceSharingRecipientNewResponseAssociationStatusAssociating    ResourceSharingRecipientNewResponseAssociationStatus = "associating"
	ResourceSharingRecipientNewResponseAssociationStatusAssociated     ResourceSharingRecipientNewResponseAssociationStatus = "associated"
	ResourceSharingRecipientNewResponseAssociationStatusDisassociating ResourceSharingRecipientNewResponseAssociationStatus = "disassociating"
	ResourceSharingRecipientNewResponseAssociationStatusDisassociated  ResourceSharingRecipientNewResponseAssociationStatus = "disassociated"
)

func (r ResourceSharingRecipientNewResponseAssociationStatus) IsKnown() bool {
	switch r {
	case ResourceSharingRecipientNewResponseAssociationStatusAssociating, ResourceSharingRecipientNewResponseAssociationStatusAssociated, ResourceSharingRecipientNewResponseAssociationStatusDisassociating, ResourceSharingRecipientNewResponseAssociationStatusDisassociated:
		return true
	}
	return false
}

type ResourceSharingRecipientListResponse struct {
	// Share Recipient identifier tag.
	ID string `json:"id,required"`
	// Account identifier.
	AccountID string `json:"account_id,required"`
	// Share Recipient association status.
	AssociationStatus ResourceSharingRecipientListResponseAssociationStatus `json:"association_status,required"`
	// When the share was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// When the share was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// Share Recipient status message.
	StatusMessage string                                   `json:"status_message,required"`
	JSON          resourceSharingRecipientListResponseJSON `json:"-"`
}

// resourceSharingRecipientListResponseJSON contains the JSON metadata for the
// struct [ResourceSharingRecipientListResponse]
type resourceSharingRecipientListResponseJSON struct {
	ID                apijson.Field
	AccountID         apijson.Field
	AssociationStatus apijson.Field
	Created           apijson.Field
	Modified          apijson.Field
	StatusMessage     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ResourceSharingRecipientListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSharingRecipientListResponseJSON) RawJSON() string {
	return r.raw
}

// Share Recipient association status.
type ResourceSharingRecipientListResponseAssociationStatus string

const (
	ResourceSharingRecipientListResponseAssociationStatusAssociating    ResourceSharingRecipientListResponseAssociationStatus = "associating"
	ResourceSharingRecipientListResponseAssociationStatusAssociated     ResourceSharingRecipientListResponseAssociationStatus = "associated"
	ResourceSharingRecipientListResponseAssociationStatusDisassociating ResourceSharingRecipientListResponseAssociationStatus = "disassociating"
	ResourceSharingRecipientListResponseAssociationStatusDisassociated  ResourceSharingRecipientListResponseAssociationStatus = "disassociated"
)

func (r ResourceSharingRecipientListResponseAssociationStatus) IsKnown() bool {
	switch r {
	case ResourceSharingRecipientListResponseAssociationStatusAssociating, ResourceSharingRecipientListResponseAssociationStatusAssociated, ResourceSharingRecipientListResponseAssociationStatusDisassociating, ResourceSharingRecipientListResponseAssociationStatusDisassociated:
		return true
	}
	return false
}

type ResourceSharingRecipientDeleteResponse struct {
	// Share Recipient identifier tag.
	ID string `json:"id,required"`
	// Account identifier.
	AccountID string `json:"account_id,required"`
	// Share Recipient association status.
	AssociationStatus ResourceSharingRecipientDeleteResponseAssociationStatus `json:"association_status,required"`
	// When the share was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// When the share was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// Share Recipient status message.
	StatusMessage string                                     `json:"status_message,required"`
	JSON          resourceSharingRecipientDeleteResponseJSON `json:"-"`
}

// resourceSharingRecipientDeleteResponseJSON contains the JSON metadata for the
// struct [ResourceSharingRecipientDeleteResponse]
type resourceSharingRecipientDeleteResponseJSON struct {
	ID                apijson.Field
	AccountID         apijson.Field
	AssociationStatus apijson.Field
	Created           apijson.Field
	Modified          apijson.Field
	StatusMessage     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ResourceSharingRecipientDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSharingRecipientDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Share Recipient association status.
type ResourceSharingRecipientDeleteResponseAssociationStatus string

const (
	ResourceSharingRecipientDeleteResponseAssociationStatusAssociating    ResourceSharingRecipientDeleteResponseAssociationStatus = "associating"
	ResourceSharingRecipientDeleteResponseAssociationStatusAssociated     ResourceSharingRecipientDeleteResponseAssociationStatus = "associated"
	ResourceSharingRecipientDeleteResponseAssociationStatusDisassociating ResourceSharingRecipientDeleteResponseAssociationStatus = "disassociating"
	ResourceSharingRecipientDeleteResponseAssociationStatusDisassociated  ResourceSharingRecipientDeleteResponseAssociationStatus = "disassociated"
)

func (r ResourceSharingRecipientDeleteResponseAssociationStatus) IsKnown() bool {
	switch r {
	case ResourceSharingRecipientDeleteResponseAssociationStatusAssociating, ResourceSharingRecipientDeleteResponseAssociationStatusAssociated, ResourceSharingRecipientDeleteResponseAssociationStatusDisassociating, ResourceSharingRecipientDeleteResponseAssociationStatusDisassociated:
		return true
	}
	return false
}

type ResourceSharingRecipientGetResponse struct {
	// Share Recipient identifier tag.
	ID string `json:"id,required"`
	// Account identifier.
	AccountID string `json:"account_id,required"`
	// Share Recipient association status.
	AssociationStatus ResourceSharingRecipientGetResponseAssociationStatus `json:"association_status,required"`
	// When the share was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// When the share was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// Share Recipient status message.
	StatusMessage string                                  `json:"status_message,required"`
	JSON          resourceSharingRecipientGetResponseJSON `json:"-"`
}

// resourceSharingRecipientGetResponseJSON contains the JSON metadata for the
// struct [ResourceSharingRecipientGetResponse]
type resourceSharingRecipientGetResponseJSON struct {
	ID                apijson.Field
	AccountID         apijson.Field
	AssociationStatus apijson.Field
	Created           apijson.Field
	Modified          apijson.Field
	StatusMessage     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ResourceSharingRecipientGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSharingRecipientGetResponseJSON) RawJSON() string {
	return r.raw
}

// Share Recipient association status.
type ResourceSharingRecipientGetResponseAssociationStatus string

const (
	ResourceSharingRecipientGetResponseAssociationStatusAssociating    ResourceSharingRecipientGetResponseAssociationStatus = "associating"
	ResourceSharingRecipientGetResponseAssociationStatusAssociated     ResourceSharingRecipientGetResponseAssociationStatus = "associated"
	ResourceSharingRecipientGetResponseAssociationStatusDisassociating ResourceSharingRecipientGetResponseAssociationStatus = "disassociating"
	ResourceSharingRecipientGetResponseAssociationStatusDisassociated  ResourceSharingRecipientGetResponseAssociationStatus = "disassociated"
)

func (r ResourceSharingRecipientGetResponseAssociationStatus) IsKnown() bool {
	switch r {
	case ResourceSharingRecipientGetResponseAssociationStatusAssociating, ResourceSharingRecipientGetResponseAssociationStatusAssociated, ResourceSharingRecipientGetResponseAssociationStatusDisassociating, ResourceSharingRecipientGetResponseAssociationStatusDisassociated:
		return true
	}
	return false
}

type ResourceSharingRecipientNewParams struct {
	// Account identifier.
	PathAccountID param.Field[string] `path:"account_id,required"`
	// Account identifier.
	BodyAccountID param.Field[string] `json:"account_id"`
	// Organization identifier.
	OrganizationID param.Field[string] `json:"organization_id"`
}

func (r ResourceSharingRecipientNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ResourceSharingRecipientNewResponseEnvelope struct {
	Errors []shared.ResponseInfo `json:"errors,required"`
	// Whether the API call was successful.
	Success bool                                            `json:"success,required"`
	Result  ResourceSharingRecipientNewResponse             `json:"result"`
	JSON    resourceSharingRecipientNewResponseEnvelopeJSON `json:"-"`
}

// resourceSharingRecipientNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [ResourceSharingRecipientNewResponseEnvelope]
type resourceSharingRecipientNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceSharingRecipientNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSharingRecipientNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ResourceSharingRecipientListParams struct {
	// Account identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Page number.
	Page param.Field[int64] `query:"page"`
	// Number of objects to return per page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [ResourceSharingRecipientListParams]'s query parameters as
// `url.Values`.
func (r ResourceSharingRecipientListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ResourceSharingRecipientDeleteParams struct {
	// Account identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ResourceSharingRecipientDeleteResponseEnvelope struct {
	Errors []shared.ResponseInfo `json:"errors,required"`
	// Whether the API call was successful.
	Success bool                                               `json:"success,required"`
	Result  ResourceSharingRecipientDeleteResponse             `json:"result"`
	JSON    resourceSharingRecipientDeleteResponseEnvelopeJSON `json:"-"`
}

// resourceSharingRecipientDeleteResponseEnvelopeJSON contains the JSON metadata
// for the struct [ResourceSharingRecipientDeleteResponseEnvelope]
type resourceSharingRecipientDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceSharingRecipientDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSharingRecipientDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ResourceSharingRecipientGetParams struct {
	// Account identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ResourceSharingRecipientGetResponseEnvelope struct {
	Errors []shared.ResponseInfo `json:"errors,required"`
	// Whether the API call was successful.
	Success bool                                            `json:"success,required"`
	Result  ResourceSharingRecipientGetResponse             `json:"result"`
	JSON    resourceSharingRecipientGetResponseEnvelopeJSON `json:"-"`
}

// resourceSharingRecipientGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ResourceSharingRecipientGetResponseEnvelope]
type resourceSharingRecipientGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceSharingRecipientGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceSharingRecipientGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
