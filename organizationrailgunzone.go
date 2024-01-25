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

// OrganizationRailgunZoneService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewOrganizationRailgunZoneService] method instead.
type OrganizationRailgunZoneService struct {
	Options []option.RequestOption
}

// NewOrganizationRailgunZoneService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrganizationRailgunZoneService(opts ...option.RequestOption) (r *OrganizationRailgunZoneService) {
	r = &OrganizationRailgunZoneService{}
	r.Options = opts
	return
}

// Lists the zones that are currently using this Railgun.
func (r *OrganizationRailgunZoneService) List(ctx context.Context, organizationIdentifier string, identifier string, opts ...option.RequestOption) (res *OrganizationRailgunZoneListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s/railguns/%s/zones", organizationIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type OrganizationRailgunZoneListResponse struct {
	Errors     []OrganizationRailgunZoneListResponseError    `json:"errors"`
	Messages   []OrganizationRailgunZoneListResponseMessage  `json:"messages"`
	Result     []OrganizationRailgunZoneListResponseResult   `json:"result"`
	ResultInfo OrganizationRailgunZoneListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success OrganizationRailgunZoneListResponseSuccess `json:"success"`
	JSON    organizationRailgunZoneListResponseJSON    `json:"-"`
}

// organizationRailgunZoneListResponseJSON contains the JSON metadata for the
// struct [OrganizationRailgunZoneListResponse]
type organizationRailgunZoneListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationRailgunZoneListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationRailgunZoneListResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    organizationRailgunZoneListResponseErrorJSON `json:"-"`
}

// organizationRailgunZoneListResponseErrorJSON contains the JSON metadata for the
// struct [OrganizationRailgunZoneListResponseError]
type organizationRailgunZoneListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationRailgunZoneListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationRailgunZoneListResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    organizationRailgunZoneListResponseMessageJSON `json:"-"`
}

// organizationRailgunZoneListResponseMessageJSON contains the JSON metadata for
// the struct [OrganizationRailgunZoneListResponseMessage]
type organizationRailgunZoneListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationRailgunZoneListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationRailgunZoneListResponseResult struct {
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
	OriginalRegistrar string                                        `json:"original_registrar,required,nullable"`
	JSON              organizationRailgunZoneListResponseResultJSON `json:"-"`
}

// organizationRailgunZoneListResponseResultJSON contains the JSON metadata for the
// struct [OrganizationRailgunZoneListResponseResult]
type organizationRailgunZoneListResponseResultJSON struct {
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

func (r *OrganizationRailgunZoneListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationRailgunZoneListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                           `json:"total_count"`
	JSON       organizationRailgunZoneListResponseResultInfoJSON `json:"-"`
}

// organizationRailgunZoneListResponseResultInfoJSON contains the JSON metadata for
// the struct [OrganizationRailgunZoneListResponseResultInfo]
type organizationRailgunZoneListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationRailgunZoneListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type OrganizationRailgunZoneListResponseSuccess bool

const (
	OrganizationRailgunZoneListResponseSuccessTrue OrganizationRailgunZoneListResponseSuccess = true
)
