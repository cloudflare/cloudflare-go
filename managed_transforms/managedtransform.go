// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package managed_transforms

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// ManagedTransformService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewManagedTransformService] method instead.
type ManagedTransformService struct {
	Options []option.RequestOption
}

// NewManagedTransformService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewManagedTransformService(opts ...option.RequestOption) (r *ManagedTransformService) {
	r = &ManagedTransformService{}
	r.Options = opts
	return
}

// Fetches a list of all Managed Transforms.
func (r *ManagedTransformService) List(ctx context.Context, query ManagedTransformListParams, opts ...option.RequestOption) (res *ManagedTransformListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/managed_headers", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the status of one or more Managed Transforms.
func (r *ManagedTransformService) Edit(ctx context.Context, params ManagedTransformEditParams, opts ...option.RequestOption) (res *ManagedTransformEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/managed_headers", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

type RequestModel struct {
	// Human-readable identifier of the Managed Transform.
	ID string `json:"id"`
	// When true, the Managed Transform is enabled.
	Enabled bool             `json:"enabled"`
	JSON    requestModelJSON `json:"-"`
}

// requestModelJSON contains the JSON metadata for the struct [RequestModel]
type requestModelJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestModelJSON) RawJSON() string {
	return r.raw
}

type RequestModelParam struct {
	// Human-readable identifier of the Managed Transform.
	ID param.Field[string] `json:"id"`
	// When true, the Managed Transform is enabled.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r RequestModelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ManagedTransformListResponse struct {
	ManagedRequestHeaders  []RequestModel                   `json:"managed_request_headers"`
	ManagedResponseHeaders []RequestModel                   `json:"managed_response_headers"`
	JSON                   managedTransformListResponseJSON `json:"-"`
}

// managedTransformListResponseJSON contains the JSON metadata for the struct
// [ManagedTransformListResponse]
type managedTransformListResponseJSON struct {
	ManagedRequestHeaders  apijson.Field
	ManagedResponseHeaders apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ManagedTransformListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedTransformListResponseJSON) RawJSON() string {
	return r.raw
}

type ManagedTransformEditResponse struct {
	ManagedRequestHeaders  []ManagedTransformEditResponseManagedRequestHeader  `json:"managed_request_headers"`
	ManagedResponseHeaders []ManagedTransformEditResponseManagedResponseHeader `json:"managed_response_headers"`
	JSON                   managedTransformEditResponseJSON                    `json:"-"`
}

// managedTransformEditResponseJSON contains the JSON metadata for the struct
// [ManagedTransformEditResponse]
type managedTransformEditResponseJSON struct {
	ManagedRequestHeaders  apijson.Field
	ManagedResponseHeaders apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ManagedTransformEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedTransformEditResponseJSON) RawJSON() string {
	return r.raw
}

type ManagedTransformEditResponseManagedRequestHeader struct {
	// Human-readable identifier of the Managed Transform.
	ID string `json:"id"`
	// When true, the Managed Transform is available in the current Cloudflare plan.
	Available bool `json:"available"`
	// When true, the Managed Transform is enabled.
	Enabled bool                                                 `json:"enabled"`
	JSON    managedTransformEditResponseManagedRequestHeaderJSON `json:"-"`
}

// managedTransformEditResponseManagedRequestHeaderJSON contains the JSON metadata
// for the struct [ManagedTransformEditResponseManagedRequestHeader]
type managedTransformEditResponseManagedRequestHeaderJSON struct {
	ID          apijson.Field
	Available   apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ManagedTransformEditResponseManagedRequestHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedTransformEditResponseManagedRequestHeaderJSON) RawJSON() string {
	return r.raw
}

type ManagedTransformEditResponseManagedResponseHeader struct {
	// Human-readable identifier of the Managed Transform.
	ID string `json:"id"`
	// When true, the Managed Transform is available in the current Cloudflare plan.
	Available bool `json:"available"`
	// When true, the Managed Transform is enabled.
	Enabled bool                                                  `json:"enabled"`
	JSON    managedTransformEditResponseManagedResponseHeaderJSON `json:"-"`
}

// managedTransformEditResponseManagedResponseHeaderJSON contains the JSON metadata
// for the struct [ManagedTransformEditResponseManagedResponseHeader]
type managedTransformEditResponseManagedResponseHeaderJSON struct {
	ID          apijson.Field
	Available   apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ManagedTransformEditResponseManagedResponseHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedTransformEditResponseManagedResponseHeaderJSON) RawJSON() string {
	return r.raw
}

type ManagedTransformListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ManagedTransformEditParams struct {
	// Identifier
	ZoneID                 param.Field[string]              `path:"zone_id,required"`
	ManagedRequestHeaders  param.Field[[]RequestModelParam] `json:"managed_request_headers,required"`
	ManagedResponseHeaders param.Field[[]RequestModelParam] `json:"managed_response_headers,required"`
}

func (r ManagedTransformEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
