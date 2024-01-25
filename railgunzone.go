// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RailgunZoneService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRailgunZoneService] method
// instead.
type RailgunZoneService struct {
	Options []option.RequestOption
}

// NewRailgunZoneService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRailgunZoneService(opts ...option.RequestOption) (r *RailgunZoneService) {
	r = &RailgunZoneService{}
	r.Options = opts
	return
}

// List the zones that are currently using this Railgun.
func (r *RailgunZoneService) List(ctx context.Context, identifier string, opts ...option.RequestOption) (res *RailgunZoneListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("railguns/%s/zones", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type RailgunZoneListResponse struct {
	Errors     []RailgunZoneListResponseError    `json:"errors"`
	Messages   []RailgunZoneListResponseMessage  `json:"messages"`
	Result     []RailgunZoneListResponseResult   `json:"result"`
	ResultInfo RailgunZoneListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success RailgunZoneListResponseSuccess `json:"success"`
	JSON    railgunZoneListResponseJSON    `json:"-"`
}

// railgunZoneListResponseJSON contains the JSON metadata for the struct
// [RailgunZoneListResponse]
type railgunZoneListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RailgunZoneListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RailgunZoneListResponseError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    railgunZoneListResponseErrorJSON `json:"-"`
}

// railgunZoneListResponseErrorJSON contains the JSON metadata for the struct
// [RailgunZoneListResponseError]
type railgunZoneListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RailgunZoneListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RailgunZoneListResponseMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    railgunZoneListResponseMessageJSON `json:"-"`
}

// railgunZoneListResponseMessageJSON contains the JSON metadata for the struct
// [RailgunZoneListResponseMessage]
type railgunZoneListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RailgunZoneListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RailgunZoneListResponseResult struct {
	// Identifier
	ID string `json:"id,required"`
	// The last time proof of ownership was detected and the zone was made active
	ActivatedOn time.Time `json:"activated_on,required,nullable" format:"date-time"`
	// When the zone was created
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The interval (in seconds) from when development mode expires (positive integer)
	// or last expired (negative integer) for the domain. If development mode has never
	// been enabled, this value is 0.
	DevelopmentMode float64 `json:"development_mode,required"`
	// When the zone was last modified
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The domain name
	Name string `json:"name,required"`
	// DNS host at the time of switching to Cloudflare
	OriginalDnshost string `json:"original_dnshost,required,nullable"`
	// Original name servers before moving to Cloudflare Notes: Is this only available
	// for full zones?
	OriginalNameServers []string `json:"original_name_servers,required,nullable" format:"hostname"`
	// Registrar for the domain at the time of switching to Cloudflare
	OriginalRegistrar string                            `json:"original_registrar,required,nullable"`
	JSON              railgunZoneListResponseResultJSON `json:"-"`
}

// railgunZoneListResponseResultJSON contains the JSON metadata for the struct
// [RailgunZoneListResponseResult]
type railgunZoneListResponseResultJSON struct {
	ID                  apijson.Field
	ActivatedOn         apijson.Field
	CreatedOn           apijson.Field
	DevelopmentMode     apijson.Field
	ModifiedOn          apijson.Field
	Name                apijson.Field
	OriginalDnshost     apijson.Field
	OriginalNameServers apijson.Field
	OriginalRegistrar   apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RailgunZoneListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RailgunZoneListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                               `json:"total_count"`
	JSON       railgunZoneListResponseResultInfoJSON `json:"-"`
}

// railgunZoneListResponseResultInfoJSON contains the JSON metadata for the struct
// [RailgunZoneListResponseResultInfo]
type railgunZoneListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RailgunZoneListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RailgunZoneListResponseSuccess bool

const (
	RailgunZoneListResponseSuccessTrue RailgunZoneListResponseSuccess = true
)
