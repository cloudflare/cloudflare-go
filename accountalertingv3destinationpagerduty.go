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

// AccountAlertingV3DestinationPagerdutyService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountAlertingV3DestinationPagerdutyService] method instead.
type AccountAlertingV3DestinationPagerdutyService struct {
	Options []option.RequestOption
}

// NewAccountAlertingV3DestinationPagerdutyService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountAlertingV3DestinationPagerdutyService(opts ...option.RequestOption) (r *AccountAlertingV3DestinationPagerdutyService) {
	r = &AccountAlertingV3DestinationPagerdutyService{}
	r.Options = opts
	return
}

// Get a list of all configured PagerDuty services.
func (r *AccountAlertingV3DestinationPagerdutyService) NotificationDestinationsWithPagerDutyListPagerDutyServices(ctx context.Context, identifier string, opts ...option.RequestOption) (res *PagerdutyResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/pagerduty", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PagerdutyResponseCollection struct {
	Errors     []PagerdutyResponseCollectionError    `json:"errors"`
	Messages   []PagerdutyResponseCollectionMessage  `json:"messages"`
	Result     []PagerdutyResponseCollectionResult   `json:"result"`
	ResultInfo PagerdutyResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success PagerdutyResponseCollectionSuccess `json:"success"`
	JSON    pagerdutyResponseCollectionJSON    `json:"-"`
}

// pagerdutyResponseCollectionJSON contains the JSON metadata for the struct
// [PagerdutyResponseCollection]
type pagerdutyResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagerdutyResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PagerdutyResponseCollectionError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    pagerdutyResponseCollectionErrorJSON `json:"-"`
}

// pagerdutyResponseCollectionErrorJSON contains the JSON metadata for the struct
// [PagerdutyResponseCollectionError]
type pagerdutyResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagerdutyResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PagerdutyResponseCollectionMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    pagerdutyResponseCollectionMessageJSON `json:"-"`
}

// pagerdutyResponseCollectionMessageJSON contains the JSON metadata for the struct
// [PagerdutyResponseCollectionMessage]
type pagerdutyResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagerdutyResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PagerdutyResponseCollectionResult struct {
	// UUID
	ID string `json:"id"`
	// The name of the pagerduty service.
	Name string                                `json:"name"`
	JSON pagerdutyResponseCollectionResultJSON `json:"-"`
}

// pagerdutyResponseCollectionResultJSON contains the JSON metadata for the struct
// [PagerdutyResponseCollectionResult]
type pagerdutyResponseCollectionResultJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagerdutyResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PagerdutyResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                   `json:"total_count"`
	JSON       pagerdutyResponseCollectionResultInfoJSON `json:"-"`
}

// pagerdutyResponseCollectionResultInfoJSON contains the JSON metadata for the
// struct [PagerdutyResponseCollectionResultInfo]
type pagerdutyResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PagerdutyResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PagerdutyResponseCollectionSuccess bool

const (
	PagerdutyResponseCollectionSuccessTrue PagerdutyResponseCollectionSuccess = true
)
