// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package managed_headers

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// ManagedHeaderService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewManagedHeaderService] method instead.
type ManagedHeaderService struct {
	Options []option.RequestOption
}

// NewManagedHeaderService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewManagedHeaderService(opts ...option.RequestOption) (r *ManagedHeaderService) {
	r = &ManagedHeaderService{}
	r.Options = opts
	return
}

// Fetches a list of all Managed Transforms.
func (r *ManagedHeaderService) List(ctx context.Context, query ManagedHeaderListParams, opts ...option.RequestOption) (res *ManagedHeaderListResponse, err error) {
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
func (r *ManagedHeaderService) Edit(ctx context.Context, params ManagedHeaderEditParams, opts ...option.RequestOption) (res *ManagedHeaderEditResponse, err error) {
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

type ManagedHeaderListResponse struct {
	ManagedRequestHeaders  []RequestModel                `json:"managed_request_headers"`
	ManagedResponseHeaders []RequestModel                `json:"managed_response_headers"`
	JSON                   managedHeaderListResponseJSON `json:"-"`
}

// managedHeaderListResponseJSON contains the JSON metadata for the struct
// [ManagedHeaderListResponse]
type managedHeaderListResponseJSON struct {
	ManagedRequestHeaders  apijson.Field
	ManagedResponseHeaders apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ManagedHeaderListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedHeaderListResponseJSON) RawJSON() string {
	return r.raw
}

type ManagedHeaderEditResponse struct {
	ManagedRequestHeaders  []ManagedHeaderEditResponseManagedRequestHeader  `json:"managed_request_headers"`
	ManagedResponseHeaders []ManagedHeaderEditResponseManagedResponseHeader `json:"managed_response_headers"`
	JSON                   managedHeaderEditResponseJSON                    `json:"-"`
}

// managedHeaderEditResponseJSON contains the JSON metadata for the struct
// [ManagedHeaderEditResponse]
type managedHeaderEditResponseJSON struct {
	ManagedRequestHeaders  apijson.Field
	ManagedResponseHeaders apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ManagedHeaderEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedHeaderEditResponseJSON) RawJSON() string {
	return r.raw
}

type ManagedHeaderEditResponseManagedRequestHeader struct {
	// Human-readable identifier of the Managed Transform.
	ID string `json:"id"`
	// When true, the Managed Transform is available in the current Cloudflare plan.
	Available bool `json:"available"`
	// When true, the Managed Transform is enabled.
	Enabled bool                                              `json:"enabled"`
	JSON    managedHeaderEditResponseManagedRequestHeaderJSON `json:"-"`
}

// managedHeaderEditResponseManagedRequestHeaderJSON contains the JSON metadata for
// the struct [ManagedHeaderEditResponseManagedRequestHeader]
type managedHeaderEditResponseManagedRequestHeaderJSON struct {
	ID          apijson.Field
	Available   apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ManagedHeaderEditResponseManagedRequestHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedHeaderEditResponseManagedRequestHeaderJSON) RawJSON() string {
	return r.raw
}

type ManagedHeaderEditResponseManagedResponseHeader struct {
	// Human-readable identifier of the Managed Transform.
	ID string `json:"id"`
	// When true, the Managed Transform is available in the current Cloudflare plan.
	Available bool `json:"available"`
	// When true, the Managed Transform is enabled.
	Enabled bool                                               `json:"enabled"`
	JSON    managedHeaderEditResponseManagedResponseHeaderJSON `json:"-"`
}

// managedHeaderEditResponseManagedResponseHeaderJSON contains the JSON metadata
// for the struct [ManagedHeaderEditResponseManagedResponseHeader]
type managedHeaderEditResponseManagedResponseHeaderJSON struct {
	ID          apijson.Field
	Available   apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ManagedHeaderEditResponseManagedResponseHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedHeaderEditResponseManagedResponseHeaderJSON) RawJSON() string {
	return r.raw
}

type ManagedHeaderListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ManagedHeaderEditParams struct {
	// Identifier
	ZoneID                 param.Field[string]              `path:"zone_id,required"`
	ManagedRequestHeaders  param.Field[[]RequestModelParam] `json:"managed_request_headers,required"`
	ManagedResponseHeaders param.Field[[]RequestModelParam] `json:"managed_response_headers,required"`
}

func (r ManagedHeaderEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
