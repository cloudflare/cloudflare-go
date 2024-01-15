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

// ZoneRailgunService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneRailgunService] method
// instead.
type ZoneRailgunService struct {
	Options   []option.RequestOption
	Diagnoses *ZoneRailgunDiagnosisService
}

// NewZoneRailgunService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneRailgunService(opts ...option.RequestOption) (r *ZoneRailgunService) {
	r = &ZoneRailgunService{}
	r.Options = opts
	r.Diagnoses = NewZoneRailgunDiagnosisService(opts...)
	return
}

// Lists details about a specific Railgun.
func (r *ZoneRailgunService) Get(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *ZoneRailgunGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/railguns/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Connect or disconnect a Railgun.
func (r *ZoneRailgunService) Update(ctx context.Context, zoneIdentifier string, identifier string, body ZoneRailgunUpdateParams, opts ...option.RequestOption) (res *ZoneRailgunUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/railguns/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// A list of available Railguns the zone can use.
func (r *ZoneRailgunService) RailgunConnectionsForAZoneListAvailableRailguns(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneRailgunRailgunConnectionsForAZoneListAvailableRailgunsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/railguns", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneRailgunGetResponse struct {
	Errors   []ZoneRailgunGetResponseError   `json:"errors"`
	Messages []ZoneRailgunGetResponseMessage `json:"messages"`
	Result   interface{}                     `json:"result"`
	// Whether the API call was successful
	Success ZoneRailgunGetResponseSuccess `json:"success"`
	JSON    zoneRailgunGetResponseJSON    `json:"-"`
}

// zoneRailgunGetResponseJSON contains the JSON metadata for the struct
// [ZoneRailgunGetResponse]
type zoneRailgunGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRailgunGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRailgunGetResponseError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    zoneRailgunGetResponseErrorJSON `json:"-"`
}

// zoneRailgunGetResponseErrorJSON contains the JSON metadata for the struct
// [ZoneRailgunGetResponseError]
type zoneRailgunGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRailgunGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRailgunGetResponseMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    zoneRailgunGetResponseMessageJSON `json:"-"`
}

// zoneRailgunGetResponseMessageJSON contains the JSON metadata for the struct
// [ZoneRailgunGetResponseMessage]
type zoneRailgunGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRailgunGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneRailgunGetResponseSuccess bool

const (
	ZoneRailgunGetResponseSuccessTrue ZoneRailgunGetResponseSuccess = true
)

type ZoneRailgunUpdateResponse struct {
	Errors   []ZoneRailgunUpdateResponseError   `json:"errors"`
	Messages []ZoneRailgunUpdateResponseMessage `json:"messages"`
	Result   interface{}                        `json:"result"`
	// Whether the API call was successful
	Success ZoneRailgunUpdateResponseSuccess `json:"success"`
	JSON    zoneRailgunUpdateResponseJSON    `json:"-"`
}

// zoneRailgunUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneRailgunUpdateResponse]
type zoneRailgunUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRailgunUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRailgunUpdateResponseError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    zoneRailgunUpdateResponseErrorJSON `json:"-"`
}

// zoneRailgunUpdateResponseErrorJSON contains the JSON metadata for the struct
// [ZoneRailgunUpdateResponseError]
type zoneRailgunUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRailgunUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRailgunUpdateResponseMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    zoneRailgunUpdateResponseMessageJSON `json:"-"`
}

// zoneRailgunUpdateResponseMessageJSON contains the JSON metadata for the struct
// [ZoneRailgunUpdateResponseMessage]
type zoneRailgunUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRailgunUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneRailgunUpdateResponseSuccess bool

const (
	ZoneRailgunUpdateResponseSuccessTrue ZoneRailgunUpdateResponseSuccess = true
)

type ZoneRailgunRailgunConnectionsForAZoneListAvailableRailgunsResponse struct {
	Errors     []ZoneRailgunRailgunConnectionsForAZoneListAvailableRailgunsResponseError    `json:"errors"`
	Messages   []ZoneRailgunRailgunConnectionsForAZoneListAvailableRailgunsResponseMessage  `json:"messages"`
	Result     []interface{}                                                                `json:"result"`
	ResultInfo ZoneRailgunRailgunConnectionsForAZoneListAvailableRailgunsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZoneRailgunRailgunConnectionsForAZoneListAvailableRailgunsResponseSuccess `json:"success"`
	JSON    zoneRailgunRailgunConnectionsForAZoneListAvailableRailgunsResponseJSON    `json:"-"`
}

// zoneRailgunRailgunConnectionsForAZoneListAvailableRailgunsResponseJSON contains
// the JSON metadata for the struct
// [ZoneRailgunRailgunConnectionsForAZoneListAvailableRailgunsResponse]
type zoneRailgunRailgunConnectionsForAZoneListAvailableRailgunsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRailgunRailgunConnectionsForAZoneListAvailableRailgunsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRailgunRailgunConnectionsForAZoneListAvailableRailgunsResponseError struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    zoneRailgunRailgunConnectionsForAZoneListAvailableRailgunsResponseErrorJSON `json:"-"`
}

// zoneRailgunRailgunConnectionsForAZoneListAvailableRailgunsResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneRailgunRailgunConnectionsForAZoneListAvailableRailgunsResponseError]
type zoneRailgunRailgunConnectionsForAZoneListAvailableRailgunsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRailgunRailgunConnectionsForAZoneListAvailableRailgunsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRailgunRailgunConnectionsForAZoneListAvailableRailgunsResponseMessage struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    zoneRailgunRailgunConnectionsForAZoneListAvailableRailgunsResponseMessageJSON `json:"-"`
}

// zoneRailgunRailgunConnectionsForAZoneListAvailableRailgunsResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneRailgunRailgunConnectionsForAZoneListAvailableRailgunsResponseMessage]
type zoneRailgunRailgunConnectionsForAZoneListAvailableRailgunsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRailgunRailgunConnectionsForAZoneListAvailableRailgunsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRailgunRailgunConnectionsForAZoneListAvailableRailgunsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                          `json:"total_count"`
	JSON       zoneRailgunRailgunConnectionsForAZoneListAvailableRailgunsResponseResultInfoJSON `json:"-"`
}

// zoneRailgunRailgunConnectionsForAZoneListAvailableRailgunsResponseResultInfoJSON
// contains the JSON metadata for the struct
// [ZoneRailgunRailgunConnectionsForAZoneListAvailableRailgunsResponseResultInfo]
type zoneRailgunRailgunConnectionsForAZoneListAvailableRailgunsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRailgunRailgunConnectionsForAZoneListAvailableRailgunsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneRailgunRailgunConnectionsForAZoneListAvailableRailgunsResponseSuccess bool

const (
	ZoneRailgunRailgunConnectionsForAZoneListAvailableRailgunsResponseSuccessTrue ZoneRailgunRailgunConnectionsForAZoneListAvailableRailgunsResponseSuccess = true
)

type ZoneRailgunUpdateParams struct {
	// A flag indicating whether the given zone is connected to the Railgun.
	Connected param.Field[bool] `json:"connected,required"`
}

func (r ZoneRailgunUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
