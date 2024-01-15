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

// ZoneCustomNService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneCustomNService] method
// instead.
type ZoneCustomNService struct {
	Options []option.RequestOption
}

// NewZoneCustomNService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneCustomNService(opts ...option.RequestOption) (r *ZoneCustomNService) {
	r = &ZoneCustomNService{}
	r.Options = opts
	return
}

// Get metadata for account-level custom nameservers on a zone.
func (r *ZoneCustomNService) AccountLevelCustomNameserversUsageForAZoneGetAccountCustomNameserverRelatedZoneMetadata(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneCustomNAccountLevelCustomNameserversUsageForAZoneGetAccountCustomNameserverRelatedZoneMetadataResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/custom_ns", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Set metadata for account-level custom nameservers on a zone.
//
// If you would like new zones in the account to use account custom nameservers by
// default, use PUT /accounts/:identifier to set the account setting
// use_account_custom_ns_by_default to true.
func (r *ZoneCustomNService) AccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadata(ctx context.Context, zoneIdentifier string, body ZoneCustomNAccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadataParams, opts ...option.RequestOption) (res *ZoneCustomNAccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadataResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/custom_ns", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ZoneCustomNAccountLevelCustomNameserversUsageForAZoneGetAccountCustomNameserverRelatedZoneMetadataResponse struct {
	// Whether zone uses account-level custom nameservers.
	Enabled  bool                                                                                                                `json:"enabled"`
	Errors   []ZoneCustomNAccountLevelCustomNameserversUsageForAZoneGetAccountCustomNameserverRelatedZoneMetadataResponseError   `json:"errors"`
	Messages []ZoneCustomNAccountLevelCustomNameserversUsageForAZoneGetAccountCustomNameserverRelatedZoneMetadataResponseMessage `json:"messages"`
	// The number of the name server set to assign to the zone.
	NsSet      float64                                                                                                              `json:"ns_set"`
	Result     []interface{}                                                                                                        `json:"result,nullable"`
	ResultInfo ZoneCustomNAccountLevelCustomNameserversUsageForAZoneGetAccountCustomNameserverRelatedZoneMetadataResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZoneCustomNAccountLevelCustomNameserversUsageForAZoneGetAccountCustomNameserverRelatedZoneMetadataResponseSuccess `json:"success"`
	JSON    zoneCustomNAccountLevelCustomNameserversUsageForAZoneGetAccountCustomNameserverRelatedZoneMetadataResponseJSON    `json:"-"`
}

// zoneCustomNAccountLevelCustomNameserversUsageForAZoneGetAccountCustomNameserverRelatedZoneMetadataResponseJSON
// contains the JSON metadata for the struct
// [ZoneCustomNAccountLevelCustomNameserversUsageForAZoneGetAccountCustomNameserverRelatedZoneMetadataResponse]
type zoneCustomNAccountLevelCustomNameserversUsageForAZoneGetAccountCustomNameserverRelatedZoneMetadataResponseJSON struct {
	Enabled     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
	NsSet       apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomNAccountLevelCustomNameserversUsageForAZoneGetAccountCustomNameserverRelatedZoneMetadataResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomNAccountLevelCustomNameserversUsageForAZoneGetAccountCustomNameserverRelatedZoneMetadataResponseError struct {
	Code    int64                                                                                                               `json:"code,required"`
	Message string                                                                                                              `json:"message,required"`
	JSON    zoneCustomNAccountLevelCustomNameserversUsageForAZoneGetAccountCustomNameserverRelatedZoneMetadataResponseErrorJSON `json:"-"`
}

// zoneCustomNAccountLevelCustomNameserversUsageForAZoneGetAccountCustomNameserverRelatedZoneMetadataResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneCustomNAccountLevelCustomNameserversUsageForAZoneGetAccountCustomNameserverRelatedZoneMetadataResponseError]
type zoneCustomNAccountLevelCustomNameserversUsageForAZoneGetAccountCustomNameserverRelatedZoneMetadataResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomNAccountLevelCustomNameserversUsageForAZoneGetAccountCustomNameserverRelatedZoneMetadataResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomNAccountLevelCustomNameserversUsageForAZoneGetAccountCustomNameserverRelatedZoneMetadataResponseMessage struct {
	Code    int64                                                                                                                 `json:"code,required"`
	Message string                                                                                                                `json:"message,required"`
	JSON    zoneCustomNAccountLevelCustomNameserversUsageForAZoneGetAccountCustomNameserverRelatedZoneMetadataResponseMessageJSON `json:"-"`
}

// zoneCustomNAccountLevelCustomNameserversUsageForAZoneGetAccountCustomNameserverRelatedZoneMetadataResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneCustomNAccountLevelCustomNameserversUsageForAZoneGetAccountCustomNameserverRelatedZoneMetadataResponseMessage]
type zoneCustomNAccountLevelCustomNameserversUsageForAZoneGetAccountCustomNameserverRelatedZoneMetadataResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomNAccountLevelCustomNameserversUsageForAZoneGetAccountCustomNameserverRelatedZoneMetadataResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomNAccountLevelCustomNameserversUsageForAZoneGetAccountCustomNameserverRelatedZoneMetadataResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                                                  `json:"total_count"`
	JSON       zoneCustomNAccountLevelCustomNameserversUsageForAZoneGetAccountCustomNameserverRelatedZoneMetadataResponseResultInfoJSON `json:"-"`
}

// zoneCustomNAccountLevelCustomNameserversUsageForAZoneGetAccountCustomNameserverRelatedZoneMetadataResponseResultInfoJSON
// contains the JSON metadata for the struct
// [ZoneCustomNAccountLevelCustomNameserversUsageForAZoneGetAccountCustomNameserverRelatedZoneMetadataResponseResultInfo]
type zoneCustomNAccountLevelCustomNameserversUsageForAZoneGetAccountCustomNameserverRelatedZoneMetadataResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomNAccountLevelCustomNameserversUsageForAZoneGetAccountCustomNameserverRelatedZoneMetadataResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneCustomNAccountLevelCustomNameserversUsageForAZoneGetAccountCustomNameserverRelatedZoneMetadataResponseSuccess bool

const (
	ZoneCustomNAccountLevelCustomNameserversUsageForAZoneGetAccountCustomNameserverRelatedZoneMetadataResponseSuccessTrue ZoneCustomNAccountLevelCustomNameserversUsageForAZoneGetAccountCustomNameserverRelatedZoneMetadataResponseSuccess = true
)

type ZoneCustomNAccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadataResponse struct {
	Errors     []ZoneCustomNAccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadataResponseError    `json:"errors"`
	Messages   []ZoneCustomNAccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadataResponseMessage  `json:"messages"`
	Result     []interface{}                                                                                                        `json:"result"`
	ResultInfo ZoneCustomNAccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadataResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZoneCustomNAccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadataResponseSuccess `json:"success"`
	JSON    zoneCustomNAccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadataResponseJSON    `json:"-"`
}

// zoneCustomNAccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadataResponseJSON
// contains the JSON metadata for the struct
// [ZoneCustomNAccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadataResponse]
type zoneCustomNAccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadataResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomNAccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadataResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomNAccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadataResponseError struct {
	Code    int64                                                                                                               `json:"code,required"`
	Message string                                                                                                              `json:"message,required"`
	JSON    zoneCustomNAccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadataResponseErrorJSON `json:"-"`
}

// zoneCustomNAccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadataResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneCustomNAccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadataResponseError]
type zoneCustomNAccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadataResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomNAccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadataResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomNAccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadataResponseMessage struct {
	Code    int64                                                                                                                 `json:"code,required"`
	Message string                                                                                                                `json:"message,required"`
	JSON    zoneCustomNAccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadataResponseMessageJSON `json:"-"`
}

// zoneCustomNAccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadataResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneCustomNAccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadataResponseMessage]
type zoneCustomNAccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadataResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomNAccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadataResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomNAccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadataResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                                                  `json:"total_count"`
	JSON       zoneCustomNAccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadataResponseResultInfoJSON `json:"-"`
}

// zoneCustomNAccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadataResponseResultInfoJSON
// contains the JSON metadata for the struct
// [ZoneCustomNAccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadataResponseResultInfo]
type zoneCustomNAccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadataResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomNAccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadataResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneCustomNAccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadataResponseSuccess bool

const (
	ZoneCustomNAccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadataResponseSuccessTrue ZoneCustomNAccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadataResponseSuccess = true
)

type ZoneCustomNAccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadataParams struct {
	// Whether zone uses account-level custom nameservers.
	Enabled param.Field[bool] `json:"enabled"`
	// The number of the name server set to assign to the zone.
	NsSet param.Field[float64] `json:"ns_set"`
}

func (r ZoneCustomNAccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadataParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
