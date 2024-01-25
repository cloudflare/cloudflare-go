// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// OrganizationRailgunService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewOrganizationRailgunService]
// method instead.
type OrganizationRailgunService struct {
	Options []option.RequestOption
	Zones   *OrganizationRailgunZoneService
}

// NewOrganizationRailgunService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrganizationRailgunService(opts ...option.RequestOption) (r *OrganizationRailgunService) {
	r = &OrganizationRailgunService{}
	r.Options = opts
	r.Zones = NewOrganizationRailgunZoneService(opts...)
	return
}

// Railgun details
func (r *OrganizationRailgunService) Get(ctx context.Context, organizationIdentifier string, identifier string, opts ...option.RequestOption) (res *OrganizationRailgunGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s/railguns/%s", organizationIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Enable or disable a Railgun for all zones connected to it.
func (r *OrganizationRailgunService) Update(ctx context.Context, organizationIdentifier string, identifier string, body OrganizationRailgunUpdateParams, opts ...option.RequestOption) (res *OrganizationRailgunUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s/railguns/%s", organizationIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Disable and delete a Railgun. This will immediately disable the Railgun for any
// connected zones.
func (r *OrganizationRailgunService) Delete(ctx context.Context, organizationIdentifier string, identifier string, opts ...option.RequestOption) (res *OrganizationRailgunDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s/railguns/%s", organizationIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create Railgun
func (r *OrganizationRailgunService) OrganizationRailgunNewRailgun(ctx context.Context, organizationIdentifier string, body OrganizationRailgunOrganizationRailgunNewRailgunParams, opts ...option.RequestOption) (res *OrganizationRailgunOrganizationRailgunNewRailgunResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s/railguns", organizationIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List, search, sort and filter your Railguns.
func (r *OrganizationRailgunService) OrganizationRailgunListRailguns(ctx context.Context, organizationIdentifier string, query OrganizationRailgunOrganizationRailgunListRailgunsParams, opts ...option.RequestOption) (res *shared.Page[OrganizationRailgunOrganizationRailgunListRailgunsResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("organizations/%s/railguns", organizationIdentifier)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

type OrganizationRailgunGetResponse struct {
	Errors   []OrganizationRailgunGetResponseError   `json:"errors"`
	Messages []OrganizationRailgunGetResponseMessage `json:"messages"`
	Result   interface{}                             `json:"result"`
	// Whether the API call was successful
	Success OrganizationRailgunGetResponseSuccess `json:"success"`
	JSON    organizationRailgunGetResponseJSON    `json:"-"`
}

// organizationRailgunGetResponseJSON contains the JSON metadata for the struct
// [OrganizationRailgunGetResponse]
type organizationRailgunGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationRailgunGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationRailgunGetResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    organizationRailgunGetResponseErrorJSON `json:"-"`
}

// organizationRailgunGetResponseErrorJSON contains the JSON metadata for the
// struct [OrganizationRailgunGetResponseError]
type organizationRailgunGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationRailgunGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationRailgunGetResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    organizationRailgunGetResponseMessageJSON `json:"-"`
}

// organizationRailgunGetResponseMessageJSON contains the JSON metadata for the
// struct [OrganizationRailgunGetResponseMessage]
type organizationRailgunGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationRailgunGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type OrganizationRailgunGetResponseSuccess bool

const (
	OrganizationRailgunGetResponseSuccessTrue OrganizationRailgunGetResponseSuccess = true
)

type OrganizationRailgunUpdateResponse struct {
	Errors   []OrganizationRailgunUpdateResponseError   `json:"errors"`
	Messages []OrganizationRailgunUpdateResponseMessage `json:"messages"`
	Result   interface{}                                `json:"result"`
	// Whether the API call was successful
	Success OrganizationRailgunUpdateResponseSuccess `json:"success"`
	JSON    organizationRailgunUpdateResponseJSON    `json:"-"`
}

// organizationRailgunUpdateResponseJSON contains the JSON metadata for the struct
// [OrganizationRailgunUpdateResponse]
type organizationRailgunUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationRailgunUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationRailgunUpdateResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    organizationRailgunUpdateResponseErrorJSON `json:"-"`
}

// organizationRailgunUpdateResponseErrorJSON contains the JSON metadata for the
// struct [OrganizationRailgunUpdateResponseError]
type organizationRailgunUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationRailgunUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationRailgunUpdateResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    organizationRailgunUpdateResponseMessageJSON `json:"-"`
}

// organizationRailgunUpdateResponseMessageJSON contains the JSON metadata for the
// struct [OrganizationRailgunUpdateResponseMessage]
type organizationRailgunUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationRailgunUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type OrganizationRailgunUpdateResponseSuccess bool

const (
	OrganizationRailgunUpdateResponseSuccessTrue OrganizationRailgunUpdateResponseSuccess = true
)

type OrganizationRailgunDeleteResponse struct {
	Errors   []OrganizationRailgunDeleteResponseError   `json:"errors"`
	Messages []OrganizationRailgunDeleteResponseMessage `json:"messages"`
	Result   OrganizationRailgunDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success OrganizationRailgunDeleteResponseSuccess `json:"success"`
	JSON    organizationRailgunDeleteResponseJSON    `json:"-"`
}

// organizationRailgunDeleteResponseJSON contains the JSON metadata for the struct
// [OrganizationRailgunDeleteResponse]
type organizationRailgunDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationRailgunDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationRailgunDeleteResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    organizationRailgunDeleteResponseErrorJSON `json:"-"`
}

// organizationRailgunDeleteResponseErrorJSON contains the JSON metadata for the
// struct [OrganizationRailgunDeleteResponseError]
type organizationRailgunDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationRailgunDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationRailgunDeleteResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    organizationRailgunDeleteResponseMessageJSON `json:"-"`
}

// organizationRailgunDeleteResponseMessageJSON contains the JSON metadata for the
// struct [OrganizationRailgunDeleteResponseMessage]
type organizationRailgunDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationRailgunDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationRailgunDeleteResponseResult struct {
	// Railgun identifier tag.
	ID   string                                      `json:"id"`
	JSON organizationRailgunDeleteResponseResultJSON `json:"-"`
}

// organizationRailgunDeleteResponseResultJSON contains the JSON metadata for the
// struct [OrganizationRailgunDeleteResponseResult]
type organizationRailgunDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationRailgunDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type OrganizationRailgunDeleteResponseSuccess bool

const (
	OrganizationRailgunDeleteResponseSuccessTrue OrganizationRailgunDeleteResponseSuccess = true
)

type OrganizationRailgunOrganizationRailgunNewRailgunResponse struct {
	Errors   []OrganizationRailgunOrganizationRailgunNewRailgunResponseError   `json:"errors"`
	Messages []OrganizationRailgunOrganizationRailgunNewRailgunResponseMessage `json:"messages"`
	Result   interface{}                                                       `json:"result"`
	// Whether the API call was successful
	Success OrganizationRailgunOrganizationRailgunNewRailgunResponseSuccess `json:"success"`
	JSON    organizationRailgunOrganizationRailgunNewRailgunResponseJSON    `json:"-"`
}

// organizationRailgunOrganizationRailgunNewRailgunResponseJSON contains the JSON
// metadata for the struct
// [OrganizationRailgunOrganizationRailgunNewRailgunResponse]
type organizationRailgunOrganizationRailgunNewRailgunResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationRailgunOrganizationRailgunNewRailgunResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationRailgunOrganizationRailgunNewRailgunResponseError struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    organizationRailgunOrganizationRailgunNewRailgunResponseErrorJSON `json:"-"`
}

// organizationRailgunOrganizationRailgunNewRailgunResponseErrorJSON contains the
// JSON metadata for the struct
// [OrganizationRailgunOrganizationRailgunNewRailgunResponseError]
type organizationRailgunOrganizationRailgunNewRailgunResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationRailgunOrganizationRailgunNewRailgunResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationRailgunOrganizationRailgunNewRailgunResponseMessage struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    organizationRailgunOrganizationRailgunNewRailgunResponseMessageJSON `json:"-"`
}

// organizationRailgunOrganizationRailgunNewRailgunResponseMessageJSON contains the
// JSON metadata for the struct
// [OrganizationRailgunOrganizationRailgunNewRailgunResponseMessage]
type organizationRailgunOrganizationRailgunNewRailgunResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationRailgunOrganizationRailgunNewRailgunResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type OrganizationRailgunOrganizationRailgunNewRailgunResponseSuccess bool

const (
	OrganizationRailgunOrganizationRailgunNewRailgunResponseSuccessTrue OrganizationRailgunOrganizationRailgunNewRailgunResponseSuccess = true
)

type OrganizationRailgunOrganizationRailgunListRailgunsResponse struct {
	// Railgun identifier tag.
	ID string `json:"id,required"`
	// When the Railgun was activated.
	ActivatedOn   time.Time `json:"activated_on,required" format:"date-time"`
	ActivationKey string    `json:"activation_key,required"`
	// The build identifier for the Railgun receiver.
	Build string `json:"build,required"`
	// When the Railgun was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Flag to determine if the Railgun is accepting connections.
	Enabled bool `json:"enabled,required"`
	// When the Railgun was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Readable identifier of the Railgun.
	Name string `json:"name,required"`
	// The revision of the Railgun receiver.
	Revision string `json:"revision,required"`
	// Status of the Railgun.
	Status OrganizationRailgunOrganizationRailgunListRailgunsResponseStatus `json:"status,required"`
	// The version of the Railgun receiver.
	Version string `json:"version,required"`
	// The number of zones using this Railgun.
	ZonesConnected float64 `json:"zones_connected,required"`
	// Defined when the Railgun version is out of date from the latest release from
	// Cloudflare.
	UpgradeInfo OrganizationRailgunOrganizationRailgunListRailgunsResponseUpgradeInfo `json:"upgrade_info"`
	JSON        organizationRailgunOrganizationRailgunListRailgunsResponseJSON        `json:"-"`
}

// organizationRailgunOrganizationRailgunListRailgunsResponseJSON contains the JSON
// metadata for the struct
// [OrganizationRailgunOrganizationRailgunListRailgunsResponse]
type organizationRailgunOrganizationRailgunListRailgunsResponseJSON struct {
	ID             apijson.Field
	ActivatedOn    apijson.Field
	ActivationKey  apijson.Field
	Build          apijson.Field
	CreatedOn      apijson.Field
	Enabled        apijson.Field
	ModifiedOn     apijson.Field
	Name           apijson.Field
	Revision       apijson.Field
	Status         apijson.Field
	Version        apijson.Field
	ZonesConnected apijson.Field
	UpgradeInfo    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *OrganizationRailgunOrganizationRailgunListRailgunsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the Railgun.
type OrganizationRailgunOrganizationRailgunListRailgunsResponseStatus string

const (
	OrganizationRailgunOrganizationRailgunListRailgunsResponseStatusInitializing OrganizationRailgunOrganizationRailgunListRailgunsResponseStatus = "initializing"
	OrganizationRailgunOrganizationRailgunListRailgunsResponseStatusActive       OrganizationRailgunOrganizationRailgunListRailgunsResponseStatus = "active"
)

// Defined when the Railgun version is out of date from the latest release from
// Cloudflare.
type OrganizationRailgunOrganizationRailgunListRailgunsResponseUpgradeInfo struct {
	// An HTTP link to download the latest Railgun binary.
	DownloadLink string `json:"download_link"`
	// Latest version of the Railgun receiver available to install.
	LatestVersion string                                                                    `json:"latest_version"`
	JSON          organizationRailgunOrganizationRailgunListRailgunsResponseUpgradeInfoJSON `json:"-"`
}

// organizationRailgunOrganizationRailgunListRailgunsResponseUpgradeInfoJSON
// contains the JSON metadata for the struct
// [OrganizationRailgunOrganizationRailgunListRailgunsResponseUpgradeInfo]
type organizationRailgunOrganizationRailgunListRailgunsResponseUpgradeInfoJSON struct {
	DownloadLink  apijson.Field
	LatestVersion apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OrganizationRailgunOrganizationRailgunListRailgunsResponseUpgradeInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationRailgunUpdateParams struct {
	// Flag to determine if the Railgun is accepting connections.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r OrganizationRailgunUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationRailgunOrganizationRailgunNewRailgunParams struct {
	// Readable identifier of the Railgun.
	Name param.Field[string] `json:"name,required"`
}

func (r OrganizationRailgunOrganizationRailgunNewRailgunParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationRailgunOrganizationRailgunListRailgunsParams struct {
	// Sort Railguns in ascending or descending order.
	Direction param.Field[OrganizationRailgunOrganizationRailgunListRailgunsParamsDirection] `query:"direction"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of items per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [OrganizationRailgunOrganizationRailgunListRailgunsParams]'s
// query parameters as `url.Values`.
func (r OrganizationRailgunOrganizationRailgunListRailgunsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort Railguns in ascending or descending order.
type OrganizationRailgunOrganizationRailgunListRailgunsParamsDirection string

const (
	OrganizationRailgunOrganizationRailgunListRailgunsParamsDirectionAsc  OrganizationRailgunOrganizationRailgunListRailgunsParamsDirection = "asc"
	OrganizationRailgunOrganizationRailgunListRailgunsParamsDirectionDesc OrganizationRailgunOrganizationRailgunListRailgunsParamsDirection = "desc"
)
