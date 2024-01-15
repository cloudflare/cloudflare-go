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
func (r *AccountCustomNAvailabilityService) AccountLevelCustomNameserversGetEligibleZonesForAccountCustomNameservers(ctx context.Context, identifier string, opts ...option.RequestOption) (res *AccountCustomNAvailabilityAccountLevelCustomNameserversGetEligibleZonesForAccountCustomNameserversResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/custom_ns/availability", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountCustomNAvailabilityAccountLevelCustomNameserversGetEligibleZonesForAccountCustomNameserversResponse struct {
	Errors     []AccountCustomNAvailabilityAccountLevelCustomNameserversGetEligibleZonesForAccountCustomNameserversResponseError    `json:"errors"`
	Messages   []AccountCustomNAvailabilityAccountLevelCustomNameserversGetEligibleZonesForAccountCustomNameserversResponseMessage  `json:"messages"`
	Result     []string                                                                                                             `json:"result" format:"hostname"`
	ResultInfo AccountCustomNAvailabilityAccountLevelCustomNameserversGetEligibleZonesForAccountCustomNameserversResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountCustomNAvailabilityAccountLevelCustomNameserversGetEligibleZonesForAccountCustomNameserversResponseSuccess `json:"success"`
	JSON    accountCustomNAvailabilityAccountLevelCustomNameserversGetEligibleZonesForAccountCustomNameserversResponseJSON    `json:"-"`
}

// accountCustomNAvailabilityAccountLevelCustomNameserversGetEligibleZonesForAccountCustomNameserversResponseJSON
// contains the JSON metadata for the struct
// [AccountCustomNAvailabilityAccountLevelCustomNameserversGetEligibleZonesForAccountCustomNameserversResponse]
type accountCustomNAvailabilityAccountLevelCustomNameserversGetEligibleZonesForAccountCustomNameserversResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomNAvailabilityAccountLevelCustomNameserversGetEligibleZonesForAccountCustomNameserversResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCustomNAvailabilityAccountLevelCustomNameserversGetEligibleZonesForAccountCustomNameserversResponseError struct {
	Code    int64                                                                                                               `json:"code,required"`
	Message string                                                                                                              `json:"message,required"`
	JSON    accountCustomNAvailabilityAccountLevelCustomNameserversGetEligibleZonesForAccountCustomNameserversResponseErrorJSON `json:"-"`
}

// accountCustomNAvailabilityAccountLevelCustomNameserversGetEligibleZonesForAccountCustomNameserversResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountCustomNAvailabilityAccountLevelCustomNameserversGetEligibleZonesForAccountCustomNameserversResponseError]
type accountCustomNAvailabilityAccountLevelCustomNameserversGetEligibleZonesForAccountCustomNameserversResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomNAvailabilityAccountLevelCustomNameserversGetEligibleZonesForAccountCustomNameserversResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCustomNAvailabilityAccountLevelCustomNameserversGetEligibleZonesForAccountCustomNameserversResponseMessage struct {
	Code    int64                                                                                                                 `json:"code,required"`
	Message string                                                                                                                `json:"message,required"`
	JSON    accountCustomNAvailabilityAccountLevelCustomNameserversGetEligibleZonesForAccountCustomNameserversResponseMessageJSON `json:"-"`
}

// accountCustomNAvailabilityAccountLevelCustomNameserversGetEligibleZonesForAccountCustomNameserversResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountCustomNAvailabilityAccountLevelCustomNameserversGetEligibleZonesForAccountCustomNameserversResponseMessage]
type accountCustomNAvailabilityAccountLevelCustomNameserversGetEligibleZonesForAccountCustomNameserversResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomNAvailabilityAccountLevelCustomNameserversGetEligibleZonesForAccountCustomNameserversResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCustomNAvailabilityAccountLevelCustomNameserversGetEligibleZonesForAccountCustomNameserversResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                                                  `json:"total_count"`
	JSON       accountCustomNAvailabilityAccountLevelCustomNameserversGetEligibleZonesForAccountCustomNameserversResponseResultInfoJSON `json:"-"`
}

// accountCustomNAvailabilityAccountLevelCustomNameserversGetEligibleZonesForAccountCustomNameserversResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountCustomNAvailabilityAccountLevelCustomNameserversGetEligibleZonesForAccountCustomNameserversResponseResultInfo]
type accountCustomNAvailabilityAccountLevelCustomNameserversGetEligibleZonesForAccountCustomNameserversResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomNAvailabilityAccountLevelCustomNameserversGetEligibleZonesForAccountCustomNameserversResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountCustomNAvailabilityAccountLevelCustomNameserversGetEligibleZonesForAccountCustomNameserversResponseSuccess bool

const (
	AccountCustomNAvailabilityAccountLevelCustomNameserversGetEligibleZonesForAccountCustomNameserversResponseSuccessTrue AccountCustomNAvailabilityAccountLevelCustomNameserversGetEligibleZonesForAccountCustomNameserversResponseSuccess = true
)
