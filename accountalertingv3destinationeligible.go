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

// AccountAlertingV3DestinationEligibleService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountAlertingV3DestinationEligibleService] method instead.
type AccountAlertingV3DestinationEligibleService struct {
	Options []option.RequestOption
}

// NewAccountAlertingV3DestinationEligibleService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountAlertingV3DestinationEligibleService(opts ...option.RequestOption) (r *AccountAlertingV3DestinationEligibleService) {
	r = &AccountAlertingV3DestinationEligibleService{}
	r.Options = opts
	return
}

// Get a list of all delivery mechanism types for which an account is eligible.
func (r *AccountAlertingV3DestinationEligibleService) NotificationMechanismEligibilityGetDeliveryMechanismEligibility(ctx context.Context, identifier string, opts ...option.RequestOption) (res *EligibilityResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/eligible", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type EligibilityResponseCollection struct {
	Errors     []EligibilityResponseCollectionError    `json:"errors"`
	Messages   []EligibilityResponseCollectionMessage  `json:"messages"`
	Result     interface{}                             `json:"result"`
	ResultInfo EligibilityResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success EligibilityResponseCollectionSuccess `json:"success"`
	JSON    eligibilityResponseCollectionJSON    `json:"-"`
}

// eligibilityResponseCollectionJSON contains the JSON metadata for the struct
// [EligibilityResponseCollection]
type eligibilityResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EligibilityResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EligibilityResponseCollectionError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    eligibilityResponseCollectionErrorJSON `json:"-"`
}

// eligibilityResponseCollectionErrorJSON contains the JSON metadata for the struct
// [EligibilityResponseCollectionError]
type eligibilityResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EligibilityResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EligibilityResponseCollectionMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    eligibilityResponseCollectionMessageJSON `json:"-"`
}

// eligibilityResponseCollectionMessageJSON contains the JSON metadata for the
// struct [EligibilityResponseCollectionMessage]
type eligibilityResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EligibilityResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EligibilityResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                     `json:"total_count"`
	JSON       eligibilityResponseCollectionResultInfoJSON `json:"-"`
}

// eligibilityResponseCollectionResultInfoJSON contains the JSON metadata for the
// struct [EligibilityResponseCollectionResultInfo]
type eligibilityResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EligibilityResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EligibilityResponseCollectionSuccess bool

const (
	EligibilityResponseCollectionSuccessTrue EligibilityResponseCollectionSuccess = true
)
