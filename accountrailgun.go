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
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountRailgunService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountRailgunService] method
// instead.
type AccountRailgunService struct {
	Options     []option.RequestOption
	Connections *AccountRailgunConnectionService
}

// NewAccountRailgunService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountRailgunService(opts ...option.RequestOption) (r *AccountRailgunService) {
	r = &AccountRailgunService{}
	r.Options = opts
	r.Connections = NewAccountRailgunConnectionService(opts...)
	return
}

// Railgun details
func (r *AccountRailgunService) Get(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *RailgunResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/railguns/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a Railgun.
func (r *AccountRailgunService) Update(ctx context.Context, accountIdentifier string, identifier string, body AccountRailgunUpdateParams, opts ...option.RequestOption) (res *RailgunResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/railguns/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Disable and delete a Railgun. This will immediately disable the Railgun for any
// connected zones.
func (r *AccountRailgunService) Delete(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *RailgunResponseSingleID, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/railguns/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create Railgun
func (r *AccountRailgunService) AccountRailgunsNewRailgun(ctx context.Context, accountIdentifier string, body AccountRailgunAccountRailgunsNewRailgunParams, opts ...option.RequestOption) (res *RailgunResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/railguns", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List, search, sort and filter your Railguns.
func (r *AccountRailgunService) AccountRailgunsListRailguns(ctx context.Context, accountIdentifier string, query AccountRailgunAccountRailgunsListRailgunsParams, opts ...option.RequestOption) (res *RailgunResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/railguns", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RailgunResponseCollection struct {
	Errors     []RailgunResponseCollectionError    `json:"errors"`
	Messages   []RailgunResponseCollectionMessage  `json:"messages"`
	Result     []RailgunResponseCollectionResult   `json:"result"`
	ResultInfo RailgunResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success RailgunResponseCollectionSuccess `json:"success"`
	JSON    railgunResponseCollectionJSON    `json:"-"`
}

// railgunResponseCollectionJSON contains the JSON metadata for the struct
// [RailgunResponseCollection]
type railgunResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RailgunResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RailgunResponseCollectionError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    railgunResponseCollectionErrorJSON `json:"-"`
}

// railgunResponseCollectionErrorJSON contains the JSON metadata for the struct
// [RailgunResponseCollectionError]
type railgunResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RailgunResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RailgunResponseCollectionMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    railgunResponseCollectionMessageJSON `json:"-"`
}

// railgunResponseCollectionMessageJSON contains the JSON metadata for the struct
// [RailgunResponseCollectionMessage]
type railgunResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RailgunResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RailgunResponseCollectionResult struct {
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
	Status RailgunResponseCollectionResultStatus `json:"status,required"`
	// The version of the Railgun receiver.
	Version string `json:"version,required"`
	// The number of zones using this Railgun.
	ZonesConnected float64 `json:"zones_connected,required"`
	// Defined when the Railgun version is out of date from the latest release from
	// Cloudflare.
	UpgradeInfo RailgunResponseCollectionResultUpgradeInfo `json:"upgrade_info"`
	JSON        railgunResponseCollectionResultJSON        `json:"-"`
}

// railgunResponseCollectionResultJSON contains the JSON metadata for the struct
// [RailgunResponseCollectionResult]
type railgunResponseCollectionResultJSON struct {
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

func (r *RailgunResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the Railgun.
type RailgunResponseCollectionResultStatus string

const (
	RailgunResponseCollectionResultStatusInitializing RailgunResponseCollectionResultStatus = "initializing"
	RailgunResponseCollectionResultStatusActive       RailgunResponseCollectionResultStatus = "active"
)

// Defined when the Railgun version is out of date from the latest release from
// Cloudflare.
type RailgunResponseCollectionResultUpgradeInfo struct {
	// An HTTP link to download the latest Railgun binary.
	DownloadLink string `json:"download_link"`
	// Latest version of the Railgun receiver available to install.
	LatestVersion string                                         `json:"latest_version"`
	JSON          railgunResponseCollectionResultUpgradeInfoJSON `json:"-"`
}

// railgunResponseCollectionResultUpgradeInfoJSON contains the JSON metadata for
// the struct [RailgunResponseCollectionResultUpgradeInfo]
type railgunResponseCollectionResultUpgradeInfoJSON struct {
	DownloadLink  apijson.Field
	LatestVersion apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RailgunResponseCollectionResultUpgradeInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RailgunResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                 `json:"total_count"`
	JSON       railgunResponseCollectionResultInfoJSON `json:"-"`
}

// railgunResponseCollectionResultInfoJSON contains the JSON metadata for the
// struct [RailgunResponseCollectionResultInfo]
type railgunResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RailgunResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RailgunResponseCollectionSuccess bool

const (
	RailgunResponseCollectionSuccessTrue RailgunResponseCollectionSuccess = true
)

type RailgunResponseSingle struct {
	Errors   []RailgunResponseSingleError   `json:"errors"`
	Messages []RailgunResponseSingleMessage `json:"messages"`
	Result   interface{}                    `json:"result"`
	// Whether the API call was successful
	Success RailgunResponseSingleSuccess `json:"success"`
	JSON    railgunResponseSingleJSON    `json:"-"`
}

// railgunResponseSingleJSON contains the JSON metadata for the struct
// [RailgunResponseSingle]
type railgunResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RailgunResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RailgunResponseSingleError struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    railgunResponseSingleErrorJSON `json:"-"`
}

// railgunResponseSingleErrorJSON contains the JSON metadata for the struct
// [RailgunResponseSingleError]
type railgunResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RailgunResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RailgunResponseSingleMessage struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    railgunResponseSingleMessageJSON `json:"-"`
}

// railgunResponseSingleMessageJSON contains the JSON metadata for the struct
// [RailgunResponseSingleMessage]
type railgunResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RailgunResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RailgunResponseSingleSuccess bool

const (
	RailgunResponseSingleSuccessTrue RailgunResponseSingleSuccess = true
)

type RailgunResponseSingleID struct {
	Errors   []RailgunResponseSingleIDError   `json:"errors"`
	Messages []RailgunResponseSingleIDMessage `json:"messages"`
	Result   RailgunResponseSingleIDResult    `json:"result"`
	// Whether the API call was successful
	Success RailgunResponseSingleIDSuccess `json:"success"`
	JSON    railgunResponseSingleIDJSON    `json:"-"`
}

// railgunResponseSingleIDJSON contains the JSON metadata for the struct
// [RailgunResponseSingleID]
type railgunResponseSingleIDJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RailgunResponseSingleID) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RailgunResponseSingleIDError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    railgunResponseSingleIDErrorJSON `json:"-"`
}

// railgunResponseSingleIDErrorJSON contains the JSON metadata for the struct
// [RailgunResponseSingleIDError]
type railgunResponseSingleIDErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RailgunResponseSingleIDError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RailgunResponseSingleIDMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    railgunResponseSingleIDMessageJSON `json:"-"`
}

// railgunResponseSingleIDMessageJSON contains the JSON metadata for the struct
// [RailgunResponseSingleIDMessage]
type railgunResponseSingleIDMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RailgunResponseSingleIDMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RailgunResponseSingleIDResult struct {
	// Railgun identifier tag.
	ID   string                            `json:"id"`
	JSON railgunResponseSingleIDResultJSON `json:"-"`
}

// railgunResponseSingleIDResultJSON contains the JSON metadata for the struct
// [RailgunResponseSingleIDResult]
type railgunResponseSingleIDResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RailgunResponseSingleIDResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RailgunResponseSingleIDSuccess bool

const (
	RailgunResponseSingleIDSuccessTrue RailgunResponseSingleIDSuccess = true
)

type AccountRailgunUpdateParams struct {
	// Flag to determine if the Railgun is accepting connections.
	Enabled param.Field[bool] `json:"enabled,required"`
	// Readable identifier of the Railgun.
	Name param.Field[string] `json:"name,required"`
	// Defined when the Railgun version is out of date from the latest release from
	// Cloudflare.
	UpgradeInfo param.Field[AccountRailgunUpdateParamsUpgradeInfo] `json:"upgrade_info"`
}

func (r AccountRailgunUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Defined when the Railgun version is out of date from the latest release from
// Cloudflare.
type AccountRailgunUpdateParamsUpgradeInfo struct {
	// An HTTP link to download the latest Railgun binary.
	DownloadLink param.Field[string] `json:"download_link"`
	// Latest version of the Railgun receiver available to install.
	LatestVersion param.Field[string] `json:"latest_version"`
}

func (r AccountRailgunUpdateParamsUpgradeInfo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountRailgunAccountRailgunsNewRailgunParams struct {
	// Readable identifier of the Railgun.
	Name param.Field[string] `json:"name,required"`
}

func (r AccountRailgunAccountRailgunsNewRailgunParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountRailgunAccountRailgunsListRailgunsParams struct {
	// Sort Railguns in ascending or descending order.
	Direction param.Field[AccountRailgunAccountRailgunsListRailgunsParamsDirection] `query:"direction"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of items per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [AccountRailgunAccountRailgunsListRailgunsParams]'s query
// parameters as `url.Values`.
func (r AccountRailgunAccountRailgunsListRailgunsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort Railguns in ascending or descending order.
type AccountRailgunAccountRailgunsListRailgunsParamsDirection string

const (
	AccountRailgunAccountRailgunsListRailgunsParamsDirectionAsc  AccountRailgunAccountRailgunsListRailgunsParamsDirection = "asc"
	AccountRailgunAccountRailgunsListRailgunsParamsDirectionDesc AccountRailgunAccountRailgunsListRailgunsParamsDirection = "desc"
)
