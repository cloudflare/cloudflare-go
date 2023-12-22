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

// ZoneWorkerFilterService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneWorkerFilterService] method
// instead.
type ZoneWorkerFilterService struct {
	Options []option.RequestOption
}

// NewZoneWorkerFilterService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneWorkerFilterService(opts ...option.RequestOption) (r *ZoneWorkerFilterService) {
	r = &ZoneWorkerFilterService{}
	r.Options = opts
	return
}

// Update Filter
func (r *ZoneWorkerFilterService) Update(ctx context.Context, zoneID string, filterID string, body ZoneWorkerFilterUpdateParams, opts ...option.RequestOption) (res *FilterResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/workers/filters/%s", zoneID, filterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete Filter
func (r *ZoneWorkerFilterService) Delete(ctx context.Context, zoneID string, filterID string, opts ...option.RequestOption) (res *APIResponseSingleIDKYar7dC1, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/workers/filters/%s", zoneID, filterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create Filter
func (r *ZoneWorkerFilterService) WorkerFiltersDeprecatedNewFilter(ctx context.Context, zoneID string, body ZoneWorkerFilterWorkerFiltersDeprecatedNewFilterParams, opts ...option.RequestOption) (res *APIResponseSingleIDKYar7dC1, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/workers/filters", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List Filters
func (r *ZoneWorkerFilterService) WorkerFiltersDeprecatedListFilters(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *FilterResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/workers/filters", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type FilterResponseCollection struct {
	Errors   []FilterResponseCollectionError   `json:"errors"`
	Messages []FilterResponseCollectionMessage `json:"messages"`
	Result   []FilterResponseCollectionResult  `json:"result"`
	// Whether the API call was successful
	Success FilterResponseCollectionSuccess `json:"success"`
	JSON    filterResponseCollectionJSON    `json:"-"`
}

// filterResponseCollectionJSON contains the JSON metadata for the struct
// [FilterResponseCollection]
type filterResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterResponseCollectionError struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    filterResponseCollectionErrorJSON `json:"-"`
}

// filterResponseCollectionErrorJSON contains the JSON metadata for the struct
// [FilterResponseCollectionError]
type filterResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterResponseCollectionMessage struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    filterResponseCollectionMessageJSON `json:"-"`
}

// filterResponseCollectionMessageJSON contains the JSON metadata for the struct
// [FilterResponseCollectionMessage]
type filterResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterResponseCollectionResult struct {
	// Identifier
	ID      string                             `json:"id,required"`
	Enabled bool                               `json:"enabled,required"`
	Pattern string                             `json:"pattern,required"`
	JSON    filterResponseCollectionResultJSON `json:"-"`
}

// filterResponseCollectionResultJSON contains the JSON metadata for the struct
// [FilterResponseCollectionResult]
type filterResponseCollectionResultJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Pattern     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FilterResponseCollectionSuccess bool

const (
	FilterResponseCollectionSuccessTrue FilterResponseCollectionSuccess = true
)

type FilterResponseSingle struct {
	Errors   []FilterResponseSingleError   `json:"errors"`
	Messages []FilterResponseSingleMessage `json:"messages"`
	Result   FilterResponseSingleResult    `json:"result"`
	// Whether the API call was successful
	Success FilterResponseSingleSuccess `json:"success"`
	JSON    filterResponseSingleJSON    `json:"-"`
}

// filterResponseSingleJSON contains the JSON metadata for the struct
// [FilterResponseSingle]
type filterResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterResponseSingleError struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    filterResponseSingleErrorJSON `json:"-"`
}

// filterResponseSingleErrorJSON contains the JSON metadata for the struct
// [FilterResponseSingleError]
type filterResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterResponseSingleMessage struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    filterResponseSingleMessageJSON `json:"-"`
}

// filterResponseSingleMessageJSON contains the JSON metadata for the struct
// [FilterResponseSingleMessage]
type filterResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterResponseSingleResult struct {
	// Identifier
	ID      string                         `json:"id,required"`
	Enabled bool                           `json:"enabled,required"`
	Pattern string                         `json:"pattern,required"`
	JSON    filterResponseSingleResultJSON `json:"-"`
}

// filterResponseSingleResultJSON contains the JSON metadata for the struct
// [FilterResponseSingleResult]
type filterResponseSingleResultJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Pattern     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterResponseSingleResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FilterResponseSingleSuccess bool

const (
	FilterResponseSingleSuccessTrue FilterResponseSingleSuccess = true
)

type ZoneWorkerFilterUpdateParams struct {
	Enabled param.Field[bool]   `json:"enabled,required"`
	Pattern param.Field[string] `json:"pattern,required"`
}

func (r ZoneWorkerFilterUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneWorkerFilterWorkerFiltersDeprecatedNewFilterParams struct {
	Enabled param.Field[bool]   `json:"enabled,required"`
	Pattern param.Field[string] `json:"pattern,required"`
}

func (r ZoneWorkerFilterWorkerFiltersDeprecatedNewFilterParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
