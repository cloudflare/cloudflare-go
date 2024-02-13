// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ManagedHeaderService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewManagedHeaderService] method
// instead.
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
func (r *ManagedHeaderService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ManagedHeaderListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/managed_headers", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the status of one or more Managed Transforms.
func (r *ManagedHeaderService) ManagedTransformsUpdateStatusOfManagedTransforms(ctx context.Context, zoneID string, body ManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsParams, opts ...option.RequestOption) (res *ManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/managed_headers", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type ManagedHeaderListResponse struct {
	ManagedRequestHeaders  []ManagedHeaderListResponseManagedRequestHeader  `json:"managed_request_headers"`
	ManagedResponseHeaders []ManagedHeaderListResponseManagedResponseHeader `json:"managed_response_headers"`
	JSON                   managedHeaderListResponseJSON                    `json:"-"`
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

type ManagedHeaderListResponseManagedRequestHeader struct {
	// Human-readable identifier of the Managed Transform.
	ID string `json:"id"`
	// When true, the Managed Transform is enabled.
	Enabled bool                                              `json:"enabled"`
	JSON    managedHeaderListResponseManagedRequestHeaderJSON `json:"-"`
}

// managedHeaderListResponseManagedRequestHeaderJSON contains the JSON metadata for
// the struct [ManagedHeaderListResponseManagedRequestHeader]
type managedHeaderListResponseManagedRequestHeaderJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ManagedHeaderListResponseManagedRequestHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ManagedHeaderListResponseManagedResponseHeader struct {
	// Human-readable identifier of the Managed Transform.
	ID string `json:"id"`
	// When true, the Managed Transform is enabled.
	Enabled bool                                               `json:"enabled"`
	JSON    managedHeaderListResponseManagedResponseHeaderJSON `json:"-"`
}

// managedHeaderListResponseManagedResponseHeaderJSON contains the JSON metadata
// for the struct [ManagedHeaderListResponseManagedResponseHeader]
type managedHeaderListResponseManagedResponseHeaderJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ManagedHeaderListResponseManagedResponseHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponse struct {
	ManagedRequestHeaders  []ManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponseManagedRequestHeader  `json:"managed_request_headers"`
	ManagedResponseHeaders []ManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponseManagedResponseHeader `json:"managed_response_headers"`
	JSON                   managedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponseJSON                    `json:"-"`
}

// managedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponseJSON
// contains the JSON metadata for the struct
// [ManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponse]
type managedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponseJSON struct {
	ManagedRequestHeaders  apijson.Field
	ManagedResponseHeaders apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponseManagedRequestHeader struct {
	// Human-readable identifier of the Managed Transform.
	ID string `json:"id"`
	// When true, the Managed Transform is available in the current Cloudflare plan.
	Available bool `json:"available"`
	// When true, the Managed Transform is enabled.
	Enabled bool                                                                                          `json:"enabled"`
	JSON    managedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponseManagedRequestHeaderJSON `json:"-"`
}

// managedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponseManagedRequestHeaderJSON
// contains the JSON metadata for the struct
// [ManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponseManagedRequestHeader]
type managedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponseManagedRequestHeaderJSON struct {
	ID          apijson.Field
	Available   apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponseManagedRequestHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponseManagedResponseHeader struct {
	// Human-readable identifier of the Managed Transform.
	ID string `json:"id"`
	// When true, the Managed Transform is available in the current Cloudflare plan.
	Available bool `json:"available"`
	// When true, the Managed Transform is enabled.
	Enabled bool                                                                                           `json:"enabled"`
	JSON    managedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponseManagedResponseHeaderJSON `json:"-"`
}

// managedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponseManagedResponseHeaderJSON
// contains the JSON metadata for the struct
// [ManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponseManagedResponseHeader]
type managedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponseManagedResponseHeaderJSON struct {
	ID          apijson.Field
	Available   apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponseManagedResponseHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsParams struct {
	ManagedRequestHeaders  param.Field[[]ManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsParamsManagedRequestHeader]  `json:"managed_request_headers,required"`
	ManagedResponseHeaders param.Field[[]ManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsParamsManagedResponseHeader] `json:"managed_response_headers,required"`
}

func (r ManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsParamsManagedRequestHeader struct {
	// Human-readable identifier of the Managed Transform.
	ID param.Field[string] `json:"id"`
	// When true, the Managed Transform is enabled.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsParamsManagedRequestHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsParamsManagedResponseHeader struct {
	// Human-readable identifier of the Managed Transform.
	ID param.Field[string] `json:"id"`
	// When true, the Managed Transform is enabled.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsParamsManagedResponseHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
