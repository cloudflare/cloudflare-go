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
func (r *AccountRailgunService) Get(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *AccountRailgunGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/railguns/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a Railgun.
func (r *AccountRailgunService) Update(ctx context.Context, accountIdentifier string, identifier string, body AccountRailgunUpdateParams, opts ...option.RequestOption) (res *AccountRailgunUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/railguns/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Disable and delete a Railgun. This will immediately disable the Railgun for any
// connected zones.
func (r *AccountRailgunService) Delete(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *AccountRailgunDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/railguns/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create Railgun
func (r *AccountRailgunService) AccountRailgunsNewRailgun(ctx context.Context, accountIdentifier string, body AccountRailgunAccountRailgunsNewRailgunParams, opts ...option.RequestOption) (res *AccountRailgunAccountRailgunsNewRailgunResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/railguns", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List, search, sort and filter your Railguns.
func (r *AccountRailgunService) AccountRailgunsListRailguns(ctx context.Context, accountIdentifier string, query AccountRailgunAccountRailgunsListRailgunsParams, opts ...option.RequestOption) (res *shared.Page[AccountRailgunAccountRailgunsListRailgunsResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/railguns", accountIdentifier)
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

type AccountRailgunGetResponse struct {
	Errors   []AccountRailgunGetResponseError   `json:"errors"`
	Messages []AccountRailgunGetResponseMessage `json:"messages"`
	Result   interface{}                        `json:"result"`
	// Whether the API call was successful
	Success AccountRailgunGetResponseSuccess `json:"success"`
	JSON    accountRailgunGetResponseJSON    `json:"-"`
}

// accountRailgunGetResponseJSON contains the JSON metadata for the struct
// [AccountRailgunGetResponse]
type accountRailgunGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRailgunGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRailgunGetResponseError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    accountRailgunGetResponseErrorJSON `json:"-"`
}

// accountRailgunGetResponseErrorJSON contains the JSON metadata for the struct
// [AccountRailgunGetResponseError]
type accountRailgunGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRailgunGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRailgunGetResponseMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    accountRailgunGetResponseMessageJSON `json:"-"`
}

// accountRailgunGetResponseMessageJSON contains the JSON metadata for the struct
// [AccountRailgunGetResponseMessage]
type accountRailgunGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRailgunGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountRailgunGetResponseSuccess bool

const (
	AccountRailgunGetResponseSuccessTrue AccountRailgunGetResponseSuccess = true
)

type AccountRailgunUpdateResponse struct {
	Errors   []AccountRailgunUpdateResponseError   `json:"errors"`
	Messages []AccountRailgunUpdateResponseMessage `json:"messages"`
	Result   interface{}                           `json:"result"`
	// Whether the API call was successful
	Success AccountRailgunUpdateResponseSuccess `json:"success"`
	JSON    accountRailgunUpdateResponseJSON    `json:"-"`
}

// accountRailgunUpdateResponseJSON contains the JSON metadata for the struct
// [AccountRailgunUpdateResponse]
type accountRailgunUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRailgunUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRailgunUpdateResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    accountRailgunUpdateResponseErrorJSON `json:"-"`
}

// accountRailgunUpdateResponseErrorJSON contains the JSON metadata for the struct
// [AccountRailgunUpdateResponseError]
type accountRailgunUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRailgunUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRailgunUpdateResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    accountRailgunUpdateResponseMessageJSON `json:"-"`
}

// accountRailgunUpdateResponseMessageJSON contains the JSON metadata for the
// struct [AccountRailgunUpdateResponseMessage]
type accountRailgunUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRailgunUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountRailgunUpdateResponseSuccess bool

const (
	AccountRailgunUpdateResponseSuccessTrue AccountRailgunUpdateResponseSuccess = true
)

type AccountRailgunDeleteResponse struct {
	Errors   []AccountRailgunDeleteResponseError   `json:"errors"`
	Messages []AccountRailgunDeleteResponseMessage `json:"messages"`
	Result   AccountRailgunDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountRailgunDeleteResponseSuccess `json:"success"`
	JSON    accountRailgunDeleteResponseJSON    `json:"-"`
}

// accountRailgunDeleteResponseJSON contains the JSON metadata for the struct
// [AccountRailgunDeleteResponse]
type accountRailgunDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRailgunDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRailgunDeleteResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    accountRailgunDeleteResponseErrorJSON `json:"-"`
}

// accountRailgunDeleteResponseErrorJSON contains the JSON metadata for the struct
// [AccountRailgunDeleteResponseError]
type accountRailgunDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRailgunDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRailgunDeleteResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    accountRailgunDeleteResponseMessageJSON `json:"-"`
}

// accountRailgunDeleteResponseMessageJSON contains the JSON metadata for the
// struct [AccountRailgunDeleteResponseMessage]
type accountRailgunDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRailgunDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRailgunDeleteResponseResult struct {
	// Railgun identifier tag.
	ID   string                                 `json:"id"`
	JSON accountRailgunDeleteResponseResultJSON `json:"-"`
}

// accountRailgunDeleteResponseResultJSON contains the JSON metadata for the struct
// [AccountRailgunDeleteResponseResult]
type accountRailgunDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRailgunDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountRailgunDeleteResponseSuccess bool

const (
	AccountRailgunDeleteResponseSuccessTrue AccountRailgunDeleteResponseSuccess = true
)

type AccountRailgunAccountRailgunsNewRailgunResponse struct {
	Errors   []AccountRailgunAccountRailgunsNewRailgunResponseError   `json:"errors"`
	Messages []AccountRailgunAccountRailgunsNewRailgunResponseMessage `json:"messages"`
	Result   interface{}                                              `json:"result"`
	// Whether the API call was successful
	Success AccountRailgunAccountRailgunsNewRailgunResponseSuccess `json:"success"`
	JSON    accountRailgunAccountRailgunsNewRailgunResponseJSON    `json:"-"`
}

// accountRailgunAccountRailgunsNewRailgunResponseJSON contains the JSON metadata
// for the struct [AccountRailgunAccountRailgunsNewRailgunResponse]
type accountRailgunAccountRailgunsNewRailgunResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRailgunAccountRailgunsNewRailgunResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRailgunAccountRailgunsNewRailgunResponseError struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    accountRailgunAccountRailgunsNewRailgunResponseErrorJSON `json:"-"`
}

// accountRailgunAccountRailgunsNewRailgunResponseErrorJSON contains the JSON
// metadata for the struct [AccountRailgunAccountRailgunsNewRailgunResponseError]
type accountRailgunAccountRailgunsNewRailgunResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRailgunAccountRailgunsNewRailgunResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRailgunAccountRailgunsNewRailgunResponseMessage struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    accountRailgunAccountRailgunsNewRailgunResponseMessageJSON `json:"-"`
}

// accountRailgunAccountRailgunsNewRailgunResponseMessageJSON contains the JSON
// metadata for the struct [AccountRailgunAccountRailgunsNewRailgunResponseMessage]
type accountRailgunAccountRailgunsNewRailgunResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRailgunAccountRailgunsNewRailgunResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountRailgunAccountRailgunsNewRailgunResponseSuccess bool

const (
	AccountRailgunAccountRailgunsNewRailgunResponseSuccessTrue AccountRailgunAccountRailgunsNewRailgunResponseSuccess = true
)

type AccountRailgunAccountRailgunsListRailgunsResponse struct {
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
	Status AccountRailgunAccountRailgunsListRailgunsResponseStatus `json:"status,required"`
	// The version of the Railgun receiver.
	Version string `json:"version,required"`
	// The number of zones using this Railgun.
	ZonesConnected float64 `json:"zones_connected,required"`
	// Defined when the Railgun version is out of date from the latest release from
	// Cloudflare.
	UpgradeInfo AccountRailgunAccountRailgunsListRailgunsResponseUpgradeInfo `json:"upgrade_info"`
	JSON        accountRailgunAccountRailgunsListRailgunsResponseJSON        `json:"-"`
}

// accountRailgunAccountRailgunsListRailgunsResponseJSON contains the JSON metadata
// for the struct [AccountRailgunAccountRailgunsListRailgunsResponse]
type accountRailgunAccountRailgunsListRailgunsResponseJSON struct {
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

func (r *AccountRailgunAccountRailgunsListRailgunsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the Railgun.
type AccountRailgunAccountRailgunsListRailgunsResponseStatus string

const (
	AccountRailgunAccountRailgunsListRailgunsResponseStatusInitializing AccountRailgunAccountRailgunsListRailgunsResponseStatus = "initializing"
	AccountRailgunAccountRailgunsListRailgunsResponseStatusActive       AccountRailgunAccountRailgunsListRailgunsResponseStatus = "active"
)

// Defined when the Railgun version is out of date from the latest release from
// Cloudflare.
type AccountRailgunAccountRailgunsListRailgunsResponseUpgradeInfo struct {
	// An HTTP link to download the latest Railgun binary.
	DownloadLink string `json:"download_link"`
	// Latest version of the Railgun receiver available to install.
	LatestVersion string                                                           `json:"latest_version"`
	JSON          accountRailgunAccountRailgunsListRailgunsResponseUpgradeInfoJSON `json:"-"`
}

// accountRailgunAccountRailgunsListRailgunsResponseUpgradeInfoJSON contains the
// JSON metadata for the struct
// [AccountRailgunAccountRailgunsListRailgunsResponseUpgradeInfo]
type accountRailgunAccountRailgunsListRailgunsResponseUpgradeInfoJSON struct {
	DownloadLink  apijson.Field
	LatestVersion apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountRailgunAccountRailgunsListRailgunsResponseUpgradeInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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
