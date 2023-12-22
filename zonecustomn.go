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

// Get Account Custom Nameserver Related Zone Metadata
func (r *ZoneCustomNService) AccountLevelCustomNameserversUsageForAZoneGetAccountCustomNameserverRelatedZoneMetadata(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *GetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/custom_ns", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Set Account Custom Nameserver Related Zone Metadata
func (r *ZoneCustomNService) AccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadata(ctx context.Context, zoneIdentifier string, body ZoneCustomNAccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadataParams, opts ...option.RequestOption) (res *EmptyResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/custom_ns", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type GetResponse struct {
	// Whether zone uses account-level custom nameservers.
	Enabled    bool                  `json:"enabled"`
	Errors     []GetResponseError    `json:"errors"`
	Messages   []GetResponseMessage  `json:"messages"`
	Result     []interface{}         `json:"result,nullable"`
	ResultInfo GetResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success GetResponseSuccess `json:"success"`
	JSON    getResponseJSON    `json:"-"`
}

// getResponseJSON contains the JSON metadata for the struct [GetResponse]
type getResponseJSON struct {
	Enabled     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GetResponseError struct {
	Code    int64                `json:"code,required"`
	Message string               `json:"message,required"`
	JSON    getResponseErrorJSON `json:"-"`
}

// getResponseErrorJSON contains the JSON metadata for the struct
// [GetResponseError]
type getResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GetResponseMessage struct {
	Code    int64                  `json:"code,required"`
	Message string                 `json:"message,required"`
	JSON    getResponseMessageJSON `json:"-"`
}

// getResponseMessageJSON contains the JSON metadata for the struct
// [GetResponseMessage]
type getResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GetResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                   `json:"total_count"`
	JSON       getResponseResultInfoJSON `json:"-"`
}

// getResponseResultInfoJSON contains the JSON metadata for the struct
// [GetResponseResultInfo]
type getResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GetResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GetResponseSuccess bool

const (
	GetResponseSuccessTrue GetResponseSuccess = true
)

type ZoneCustomNAccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadataParams struct {
	// Whether zone uses account-level custom nameservers.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZoneCustomNAccountLevelCustomNameserversUsageForAZoneSetAccountCustomNameserverRelatedZoneMetadataParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
