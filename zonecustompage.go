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

// ZoneCustomPageService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneCustomPageService] method
// instead.
type ZoneCustomPageService struct {
	Options []option.RequestOption
}

// NewZoneCustomPageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneCustomPageService(opts ...option.RequestOption) (r *ZoneCustomPageService) {
	r = &ZoneCustomPageService{}
	r.Options = opts
	return
}

// Fetches the details of a custom page.
func (r *ZoneCustomPageService) Get(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *ZoneCustomPageGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/custom_pages/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the configuration of an existing custom page.
func (r *ZoneCustomPageService) Update(ctx context.Context, zoneIdentifier string, identifier string, body ZoneCustomPageUpdateParams, opts ...option.RequestOption) (res *ZoneCustomPageUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/custom_pages/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Fetches all the custom pages at the zone level.
func (r *ZoneCustomPageService) CustomPagesForAZoneListCustomPages(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneCustomPageCustomPagesForAZoneListCustomPagesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/custom_pages", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneCustomPageGetResponse struct {
	Errors   []ZoneCustomPageGetResponseError   `json:"errors"`
	Messages []ZoneCustomPageGetResponseMessage `json:"messages"`
	Result   interface{}                        `json:"result"`
	// Whether the API call was successful
	Success ZoneCustomPageGetResponseSuccess `json:"success"`
	JSON    zoneCustomPageGetResponseJSON    `json:"-"`
}

// zoneCustomPageGetResponseJSON contains the JSON metadata for the struct
// [ZoneCustomPageGetResponse]
type zoneCustomPageGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomPageGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomPageGetResponseError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    zoneCustomPageGetResponseErrorJSON `json:"-"`
}

// zoneCustomPageGetResponseErrorJSON contains the JSON metadata for the struct
// [ZoneCustomPageGetResponseError]
type zoneCustomPageGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomPageGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomPageGetResponseMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    zoneCustomPageGetResponseMessageJSON `json:"-"`
}

// zoneCustomPageGetResponseMessageJSON contains the JSON metadata for the struct
// [ZoneCustomPageGetResponseMessage]
type zoneCustomPageGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomPageGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneCustomPageGetResponseSuccess bool

const (
	ZoneCustomPageGetResponseSuccessTrue ZoneCustomPageGetResponseSuccess = true
)

type ZoneCustomPageUpdateResponse struct {
	Errors   []ZoneCustomPageUpdateResponseError   `json:"errors"`
	Messages []ZoneCustomPageUpdateResponseMessage `json:"messages"`
	Result   interface{}                           `json:"result"`
	// Whether the API call was successful
	Success ZoneCustomPageUpdateResponseSuccess `json:"success"`
	JSON    zoneCustomPageUpdateResponseJSON    `json:"-"`
}

// zoneCustomPageUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneCustomPageUpdateResponse]
type zoneCustomPageUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomPageUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomPageUpdateResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    zoneCustomPageUpdateResponseErrorJSON `json:"-"`
}

// zoneCustomPageUpdateResponseErrorJSON contains the JSON metadata for the struct
// [ZoneCustomPageUpdateResponseError]
type zoneCustomPageUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomPageUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomPageUpdateResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    zoneCustomPageUpdateResponseMessageJSON `json:"-"`
}

// zoneCustomPageUpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneCustomPageUpdateResponseMessage]
type zoneCustomPageUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomPageUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneCustomPageUpdateResponseSuccess bool

const (
	ZoneCustomPageUpdateResponseSuccessTrue ZoneCustomPageUpdateResponseSuccess = true
)

type ZoneCustomPageCustomPagesForAZoneListCustomPagesResponse struct {
	Errors     []ZoneCustomPageCustomPagesForAZoneListCustomPagesResponseError    `json:"errors"`
	Messages   []ZoneCustomPageCustomPagesForAZoneListCustomPagesResponseMessage  `json:"messages"`
	Result     []interface{}                                                      `json:"result"`
	ResultInfo ZoneCustomPageCustomPagesForAZoneListCustomPagesResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZoneCustomPageCustomPagesForAZoneListCustomPagesResponseSuccess `json:"success"`
	JSON    zoneCustomPageCustomPagesForAZoneListCustomPagesResponseJSON    `json:"-"`
}

// zoneCustomPageCustomPagesForAZoneListCustomPagesResponseJSON contains the JSON
// metadata for the struct
// [ZoneCustomPageCustomPagesForAZoneListCustomPagesResponse]
type zoneCustomPageCustomPagesForAZoneListCustomPagesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomPageCustomPagesForAZoneListCustomPagesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomPageCustomPagesForAZoneListCustomPagesResponseError struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    zoneCustomPageCustomPagesForAZoneListCustomPagesResponseErrorJSON `json:"-"`
}

// zoneCustomPageCustomPagesForAZoneListCustomPagesResponseErrorJSON contains the
// JSON metadata for the struct
// [ZoneCustomPageCustomPagesForAZoneListCustomPagesResponseError]
type zoneCustomPageCustomPagesForAZoneListCustomPagesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomPageCustomPagesForAZoneListCustomPagesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomPageCustomPagesForAZoneListCustomPagesResponseMessage struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    zoneCustomPageCustomPagesForAZoneListCustomPagesResponseMessageJSON `json:"-"`
}

// zoneCustomPageCustomPagesForAZoneListCustomPagesResponseMessageJSON contains the
// JSON metadata for the struct
// [ZoneCustomPageCustomPagesForAZoneListCustomPagesResponseMessage]
type zoneCustomPageCustomPagesForAZoneListCustomPagesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomPageCustomPagesForAZoneListCustomPagesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomPageCustomPagesForAZoneListCustomPagesResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                `json:"total_count"`
	JSON       zoneCustomPageCustomPagesForAZoneListCustomPagesResponseResultInfoJSON `json:"-"`
}

// zoneCustomPageCustomPagesForAZoneListCustomPagesResponseResultInfoJSON contains
// the JSON metadata for the struct
// [ZoneCustomPageCustomPagesForAZoneListCustomPagesResponseResultInfo]
type zoneCustomPageCustomPagesForAZoneListCustomPagesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomPageCustomPagesForAZoneListCustomPagesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneCustomPageCustomPagesForAZoneListCustomPagesResponseSuccess bool

const (
	ZoneCustomPageCustomPagesForAZoneListCustomPagesResponseSuccessTrue ZoneCustomPageCustomPagesForAZoneListCustomPagesResponseSuccess = true
)

type ZoneCustomPageUpdateParams struct {
	// The custom page state.
	State param.Field[ZoneCustomPageUpdateParamsState] `json:"state,required"`
	// The URL associated with the custom page.
	URL param.Field[string] `json:"url,required" format:"uri"`
}

func (r ZoneCustomPageUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The custom page state.
type ZoneCustomPageUpdateParamsState string

const (
	ZoneCustomPageUpdateParamsStateDefault    ZoneCustomPageUpdateParamsState = "default"
	ZoneCustomPageUpdateParamsStateCustomized ZoneCustomPageUpdateParamsState = "customized"
)
