// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountCustomNAvailabilityService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountCustomNAvailabilityService] method instead.
type AccountCustomNAvailabilityService struct {
	Options []option.RequestOption
}

// NewAccountCustomNAvailabilityService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountCustomNAvailabilityService(opts ...option.RequestOption) (r *AccountCustomNAvailabilityService) {
	r = &AccountCustomNAvailabilityService{}
	r.Options = opts
	return
}

// Get Eligible Zones for Account Custom Nameservers
func (r *AccountCustomNAvailabilityService) AccountLevelCustomNameserversGetEligibleZonesForAccountCustomNameservers(ctx context.Context, identifier string, opts ...option.RequestOption) (res *AvailabilityResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/custom_ns/availability", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AvailabilityResponse struct {
	Errors     []AvailabilityResponseError    `json:"errors"`
	Messages   []AvailabilityResponseMessage  `json:"messages"`
	Result     []string                       `json:"result" format:"hostname"`
	ResultInfo AvailabilityResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AvailabilityResponseSuccess `json:"success"`
	JSON    availabilityResponseJSON    `json:"-"`
}

// availabilityResponseJSON contains the JSON metadata for the struct
// [AvailabilityResponse]
type availabilityResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailabilityResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AvailabilityResponseError struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    availabilityResponseErrorJSON `json:"-"`
}

// availabilityResponseErrorJSON contains the JSON metadata for the struct
// [AvailabilityResponseError]
type availabilityResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailabilityResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AvailabilityResponseMessage struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    availabilityResponseMessageJSON `json:"-"`
}

// availabilityResponseMessageJSON contains the JSON metadata for the struct
// [AvailabilityResponseMessage]
type availabilityResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailabilityResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AvailabilityResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                            `json:"total_count"`
	JSON       availabilityResponseResultInfoJSON `json:"-"`
}

// availabilityResponseResultInfoJSON contains the JSON metadata for the struct
// [AvailabilityResponseResultInfo]
type availabilityResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailabilityResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AvailabilityResponseSuccess bool

const (
	AvailabilityResponseSuccessTrue AvailabilityResponseSuccess = true
)
