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

// ZoneManagedHeaderService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneManagedHeaderService] method
// instead.
type ZoneManagedHeaderService struct {
	Options []option.RequestOption
}

// NewZoneManagedHeaderService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneManagedHeaderService(opts ...option.RequestOption) (r *ZoneManagedHeaderService) {
	r = &ZoneManagedHeaderService{}
	r.Options = opts
	return
}

// Fetches a list of all Managed Transforms.
func (r *ZoneManagedHeaderService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneManagedHeaderListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/managed_headers", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the status of one or more Managed Transforms.
func (r *ZoneManagedHeaderService) ManagedTransformsUpdateStatusOfManagedTransforms(ctx context.Context, zoneID string, body ZoneManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsParams, opts ...option.RequestOption) (res *ZoneManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/managed_headers", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type ZoneManagedHeaderListResponse struct {
	ManagedRequestHeaders  []ZoneManagedHeaderListResponseManagedRequestHeader  `json:"managed_request_headers"`
	ManagedResponseHeaders []ZoneManagedHeaderListResponseManagedResponseHeader `json:"managed_response_headers"`
	JSON                   zoneManagedHeaderListResponseJSON                    `json:"-"`
}

// zoneManagedHeaderListResponseJSON contains the JSON metadata for the struct
// [ZoneManagedHeaderListResponse]
type zoneManagedHeaderListResponseJSON struct {
	ManagedRequestHeaders  apijson.Field
	ManagedResponseHeaders apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneManagedHeaderListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneManagedHeaderListResponseManagedRequestHeader struct {
	// Human-readable identifier of the Managed Transform.
	ID string `json:"id"`
	// When true, the Managed Transform is enabled.
	Enabled bool                                                  `json:"enabled"`
	JSON    zoneManagedHeaderListResponseManagedRequestHeaderJSON `json:"-"`
}

// zoneManagedHeaderListResponseManagedRequestHeaderJSON contains the JSON metadata
// for the struct [ZoneManagedHeaderListResponseManagedRequestHeader]
type zoneManagedHeaderListResponseManagedRequestHeaderJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneManagedHeaderListResponseManagedRequestHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneManagedHeaderListResponseManagedResponseHeader struct {
	// Human-readable identifier of the Managed Transform.
	ID string `json:"id"`
	// When true, the Managed Transform is enabled.
	Enabled bool                                                   `json:"enabled"`
	JSON    zoneManagedHeaderListResponseManagedResponseHeaderJSON `json:"-"`
}

// zoneManagedHeaderListResponseManagedResponseHeaderJSON contains the JSON
// metadata for the struct [ZoneManagedHeaderListResponseManagedResponseHeader]
type zoneManagedHeaderListResponseManagedResponseHeaderJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneManagedHeaderListResponseManagedResponseHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponse struct {
	ManagedRequestHeaders  []ZoneManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponseManagedRequestHeader  `json:"managed_request_headers"`
	ManagedResponseHeaders []ZoneManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponseManagedResponseHeader `json:"managed_response_headers"`
	JSON                   zoneManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponseJSON                    `json:"-"`
}

// zoneManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponseJSON
// contains the JSON metadata for the struct
// [ZoneManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponse]
type zoneManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponseJSON struct {
	ManagedRequestHeaders  apijson.Field
	ManagedResponseHeaders apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponseManagedRequestHeader struct {
	// Human-readable identifier of the Managed Transform.
	ID string `json:"id"`
	// When true, the Managed Transform is available in the current Cloudflare plan.
	Available bool `json:"available"`
	// When true, the Managed Transform is enabled.
	Enabled bool                                                                                              `json:"enabled"`
	JSON    zoneManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponseManagedRequestHeaderJSON `json:"-"`
}

// zoneManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponseManagedRequestHeaderJSON
// contains the JSON metadata for the struct
// [ZoneManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponseManagedRequestHeader]
type zoneManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponseManagedRequestHeaderJSON struct {
	ID          apijson.Field
	Available   apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponseManagedRequestHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponseManagedResponseHeader struct {
	// Human-readable identifier of the Managed Transform.
	ID string `json:"id"`
	// When true, the Managed Transform is available in the current Cloudflare plan.
	Available bool `json:"available"`
	// When true, the Managed Transform is enabled.
	Enabled bool                                                                                               `json:"enabled"`
	JSON    zoneManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponseManagedResponseHeaderJSON `json:"-"`
}

// zoneManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponseManagedResponseHeaderJSON
// contains the JSON metadata for the struct
// [ZoneManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponseManagedResponseHeader]
type zoneManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponseManagedResponseHeaderJSON struct {
	ID          apijson.Field
	Available   apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsResponseManagedResponseHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsParams struct {
	ManagedRequestHeaders  param.Field[[]ZoneManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsParamsManagedRequestHeader]  `json:"managed_request_headers,required"`
	ManagedResponseHeaders param.Field[[]ZoneManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsParamsManagedResponseHeader] `json:"managed_response_headers,required"`
}

func (r ZoneManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsParamsManagedRequestHeader struct {
	// Human-readable identifier of the Managed Transform.
	ID param.Field[string] `json:"id"`
	// When true, the Managed Transform is enabled.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZoneManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsParamsManagedRequestHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsParamsManagedResponseHeader struct {
	// Human-readable identifier of the Managed Transform.
	ID param.Field[string] `json:"id"`
	// When true, the Managed Transform is enabled.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZoneManagedHeaderManagedTransformsUpdateStatusOfManagedTransformsParamsManagedResponseHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
