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
func (r *ZoneWorkerFilterService) Update(ctx context.Context, zoneID string, filterID string, body ZoneWorkerFilterUpdateParams, opts ...option.RequestOption) (res *ZoneWorkerFilterUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/workers/filters/%s", zoneID, filterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete Filter
func (r *ZoneWorkerFilterService) Delete(ctx context.Context, zoneID string, filterID string, opts ...option.RequestOption) (res *ZoneWorkerFilterDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/workers/filters/%s", zoneID, filterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create Filter
func (r *ZoneWorkerFilterService) WorkerFiltersDeprecatedNewFilter(ctx context.Context, zoneID string, body ZoneWorkerFilterWorkerFiltersDeprecatedNewFilterParams, opts ...option.RequestOption) (res *ZoneWorkerFilterWorkerFiltersDeprecatedNewFilterResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/workers/filters", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List Filters
func (r *ZoneWorkerFilterService) WorkerFiltersDeprecatedListFilters(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneWorkerFilterWorkerFiltersDeprecatedListFiltersResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/workers/filters", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneWorkerFilterUpdateResponse struct {
	Errors   []ZoneWorkerFilterUpdateResponseError   `json:"errors"`
	Messages []ZoneWorkerFilterUpdateResponseMessage `json:"messages"`
	Result   ZoneWorkerFilterUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneWorkerFilterUpdateResponseSuccess `json:"success"`
	JSON    zoneWorkerFilterUpdateResponseJSON    `json:"-"`
}

// zoneWorkerFilterUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneWorkerFilterUpdateResponse]
type zoneWorkerFilterUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerFilterUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWorkerFilterUpdateResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    zoneWorkerFilterUpdateResponseErrorJSON `json:"-"`
}

// zoneWorkerFilterUpdateResponseErrorJSON contains the JSON metadata for the
// struct [ZoneWorkerFilterUpdateResponseError]
type zoneWorkerFilterUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerFilterUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWorkerFilterUpdateResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    zoneWorkerFilterUpdateResponseMessageJSON `json:"-"`
}

// zoneWorkerFilterUpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneWorkerFilterUpdateResponseMessage]
type zoneWorkerFilterUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerFilterUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWorkerFilterUpdateResponseResult struct {
	// Identifier
	ID      string                                   `json:"id,required"`
	Enabled bool                                     `json:"enabled,required"`
	Pattern string                                   `json:"pattern,required"`
	JSON    zoneWorkerFilterUpdateResponseResultJSON `json:"-"`
}

// zoneWorkerFilterUpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneWorkerFilterUpdateResponseResult]
type zoneWorkerFilterUpdateResponseResultJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Pattern     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerFilterUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneWorkerFilterUpdateResponseSuccess bool

const (
	ZoneWorkerFilterUpdateResponseSuccessTrue ZoneWorkerFilterUpdateResponseSuccess = true
)

type ZoneWorkerFilterDeleteResponse struct {
	Errors   []ZoneWorkerFilterDeleteResponseError   `json:"errors"`
	Messages []ZoneWorkerFilterDeleteResponseMessage `json:"messages"`
	Result   ZoneWorkerFilterDeleteResponseResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success ZoneWorkerFilterDeleteResponseSuccess `json:"success"`
	JSON    zoneWorkerFilterDeleteResponseJSON    `json:"-"`
}

// zoneWorkerFilterDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneWorkerFilterDeleteResponse]
type zoneWorkerFilterDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerFilterDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWorkerFilterDeleteResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    zoneWorkerFilterDeleteResponseErrorJSON `json:"-"`
}

// zoneWorkerFilterDeleteResponseErrorJSON contains the JSON metadata for the
// struct [ZoneWorkerFilterDeleteResponseError]
type zoneWorkerFilterDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerFilterDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWorkerFilterDeleteResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    zoneWorkerFilterDeleteResponseMessageJSON `json:"-"`
}

// zoneWorkerFilterDeleteResponseMessageJSON contains the JSON metadata for the
// struct [ZoneWorkerFilterDeleteResponseMessage]
type zoneWorkerFilterDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerFilterDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWorkerFilterDeleteResponseResult struct {
	// Identifier
	ID   string                                   `json:"id,required"`
	JSON zoneWorkerFilterDeleteResponseResultJSON `json:"-"`
}

// zoneWorkerFilterDeleteResponseResultJSON contains the JSON metadata for the
// struct [ZoneWorkerFilterDeleteResponseResult]
type zoneWorkerFilterDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerFilterDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneWorkerFilterDeleteResponseSuccess bool

const (
	ZoneWorkerFilterDeleteResponseSuccessTrue ZoneWorkerFilterDeleteResponseSuccess = true
)

type ZoneWorkerFilterWorkerFiltersDeprecatedNewFilterResponse struct {
	Errors   []ZoneWorkerFilterWorkerFiltersDeprecatedNewFilterResponseError   `json:"errors"`
	Messages []ZoneWorkerFilterWorkerFiltersDeprecatedNewFilterResponseMessage `json:"messages"`
	Result   ZoneWorkerFilterWorkerFiltersDeprecatedNewFilterResponseResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success ZoneWorkerFilterWorkerFiltersDeprecatedNewFilterResponseSuccess `json:"success"`
	JSON    zoneWorkerFilterWorkerFiltersDeprecatedNewFilterResponseJSON    `json:"-"`
}

// zoneWorkerFilterWorkerFiltersDeprecatedNewFilterResponseJSON contains the JSON
// metadata for the struct
// [ZoneWorkerFilterWorkerFiltersDeprecatedNewFilterResponse]
type zoneWorkerFilterWorkerFiltersDeprecatedNewFilterResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerFilterWorkerFiltersDeprecatedNewFilterResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWorkerFilterWorkerFiltersDeprecatedNewFilterResponseError struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    zoneWorkerFilterWorkerFiltersDeprecatedNewFilterResponseErrorJSON `json:"-"`
}

// zoneWorkerFilterWorkerFiltersDeprecatedNewFilterResponseErrorJSON contains the
// JSON metadata for the struct
// [ZoneWorkerFilterWorkerFiltersDeprecatedNewFilterResponseError]
type zoneWorkerFilterWorkerFiltersDeprecatedNewFilterResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerFilterWorkerFiltersDeprecatedNewFilterResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWorkerFilterWorkerFiltersDeprecatedNewFilterResponseMessage struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    zoneWorkerFilterWorkerFiltersDeprecatedNewFilterResponseMessageJSON `json:"-"`
}

// zoneWorkerFilterWorkerFiltersDeprecatedNewFilterResponseMessageJSON contains the
// JSON metadata for the struct
// [ZoneWorkerFilterWorkerFiltersDeprecatedNewFilterResponseMessage]
type zoneWorkerFilterWorkerFiltersDeprecatedNewFilterResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerFilterWorkerFiltersDeprecatedNewFilterResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWorkerFilterWorkerFiltersDeprecatedNewFilterResponseResult struct {
	// Identifier
	ID   string                                                             `json:"id,required"`
	JSON zoneWorkerFilterWorkerFiltersDeprecatedNewFilterResponseResultJSON `json:"-"`
}

// zoneWorkerFilterWorkerFiltersDeprecatedNewFilterResponseResultJSON contains the
// JSON metadata for the struct
// [ZoneWorkerFilterWorkerFiltersDeprecatedNewFilterResponseResult]
type zoneWorkerFilterWorkerFiltersDeprecatedNewFilterResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerFilterWorkerFiltersDeprecatedNewFilterResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneWorkerFilterWorkerFiltersDeprecatedNewFilterResponseSuccess bool

const (
	ZoneWorkerFilterWorkerFiltersDeprecatedNewFilterResponseSuccessTrue ZoneWorkerFilterWorkerFiltersDeprecatedNewFilterResponseSuccess = true
)

type ZoneWorkerFilterWorkerFiltersDeprecatedListFiltersResponse struct {
	Errors   []ZoneWorkerFilterWorkerFiltersDeprecatedListFiltersResponseError   `json:"errors"`
	Messages []ZoneWorkerFilterWorkerFiltersDeprecatedListFiltersResponseMessage `json:"messages"`
	Result   []ZoneWorkerFilterWorkerFiltersDeprecatedListFiltersResponseResult  `json:"result"`
	// Whether the API call was successful
	Success ZoneWorkerFilterWorkerFiltersDeprecatedListFiltersResponseSuccess `json:"success"`
	JSON    zoneWorkerFilterWorkerFiltersDeprecatedListFiltersResponseJSON    `json:"-"`
}

// zoneWorkerFilterWorkerFiltersDeprecatedListFiltersResponseJSON contains the JSON
// metadata for the struct
// [ZoneWorkerFilterWorkerFiltersDeprecatedListFiltersResponse]
type zoneWorkerFilterWorkerFiltersDeprecatedListFiltersResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerFilterWorkerFiltersDeprecatedListFiltersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWorkerFilterWorkerFiltersDeprecatedListFiltersResponseError struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    zoneWorkerFilterWorkerFiltersDeprecatedListFiltersResponseErrorJSON `json:"-"`
}

// zoneWorkerFilterWorkerFiltersDeprecatedListFiltersResponseErrorJSON contains the
// JSON metadata for the struct
// [ZoneWorkerFilterWorkerFiltersDeprecatedListFiltersResponseError]
type zoneWorkerFilterWorkerFiltersDeprecatedListFiltersResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerFilterWorkerFiltersDeprecatedListFiltersResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWorkerFilterWorkerFiltersDeprecatedListFiltersResponseMessage struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    zoneWorkerFilterWorkerFiltersDeprecatedListFiltersResponseMessageJSON `json:"-"`
}

// zoneWorkerFilterWorkerFiltersDeprecatedListFiltersResponseMessageJSON contains
// the JSON metadata for the struct
// [ZoneWorkerFilterWorkerFiltersDeprecatedListFiltersResponseMessage]
type zoneWorkerFilterWorkerFiltersDeprecatedListFiltersResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerFilterWorkerFiltersDeprecatedListFiltersResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWorkerFilterWorkerFiltersDeprecatedListFiltersResponseResult struct {
	// Identifier
	ID      string                                                               `json:"id,required"`
	Enabled bool                                                                 `json:"enabled,required"`
	Pattern string                                                               `json:"pattern,required"`
	JSON    zoneWorkerFilterWorkerFiltersDeprecatedListFiltersResponseResultJSON `json:"-"`
}

// zoneWorkerFilterWorkerFiltersDeprecatedListFiltersResponseResultJSON contains
// the JSON metadata for the struct
// [ZoneWorkerFilterWorkerFiltersDeprecatedListFiltersResponseResult]
type zoneWorkerFilterWorkerFiltersDeprecatedListFiltersResponseResultJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Pattern     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerFilterWorkerFiltersDeprecatedListFiltersResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneWorkerFilterWorkerFiltersDeprecatedListFiltersResponseSuccess bool

const (
	ZoneWorkerFilterWorkerFiltersDeprecatedListFiltersResponseSuccessTrue ZoneWorkerFilterWorkerFiltersDeprecatedListFiltersResponseSuccess = true
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
