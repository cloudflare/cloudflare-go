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

// RailgunService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewRailgunService] method instead.
type RailgunService struct {
	Options []option.RequestOption
	Zones   *RailgunZoneService
}

// NewRailgunService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRailgunService(opts ...option.RequestOption) (r *RailgunService) {
	r = &RailgunService{}
	r.Options = opts
	r.Zones = NewRailgunZoneService(opts...)
	return
}

// Railgun details
func (r *RailgunService) Get(ctx context.Context, identifier string, opts ...option.RequestOption) (res *RailgunGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("railguns/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Enable or disable a Railgun for all zones connected to it.
func (r *RailgunService) Update(ctx context.Context, identifier string, body RailgunUpdateParams, opts ...option.RequestOption) (res *RailgunUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("railguns/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Disable and delete a Railgun. This will immediately disable that Railgun for any
// connected zones.
func (r *RailgunService) Delete(ctx context.Context, identifier string, opts ...option.RequestOption) (res *RailgunDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("railguns/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create Railgun
func (r *RailgunService) RailgunNewRailgun(ctx context.Context, body RailgunRailgunNewRailgunParams, opts ...option.RequestOption) (res *RailgunRailgunNewRailgunResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "railguns"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List, search, sort and filter your Railguns.
func (r *RailgunService) RailgunListRailguns(ctx context.Context, query RailgunRailgunListRailgunsParams, opts ...option.RequestOption) (res *shared.Page[RailgunRailgunListRailgunsResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "railguns"
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

type RailgunGetResponse struct {
	Errors   []RailgunGetResponseError   `json:"errors"`
	Messages []RailgunGetResponseMessage `json:"messages"`
	Result   interface{}                 `json:"result"`
	// Whether the API call was successful
	Success RailgunGetResponseSuccess `json:"success"`
	JSON    railgunGetResponseJSON    `json:"-"`
}

// railgunGetResponseJSON contains the JSON metadata for the struct
// [RailgunGetResponse]
type railgunGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RailgunGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RailgunGetResponseError struct {
	Code    int64                       `json:"code,required"`
	Message string                      `json:"message,required"`
	JSON    railgunGetResponseErrorJSON `json:"-"`
}

// railgunGetResponseErrorJSON contains the JSON metadata for the struct
// [RailgunGetResponseError]
type railgunGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RailgunGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RailgunGetResponseMessage struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    railgunGetResponseMessageJSON `json:"-"`
}

// railgunGetResponseMessageJSON contains the JSON metadata for the struct
// [RailgunGetResponseMessage]
type railgunGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RailgunGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RailgunGetResponseSuccess bool

const (
	RailgunGetResponseSuccessTrue RailgunGetResponseSuccess = true
)

type RailgunUpdateResponse struct {
	Errors   []RailgunUpdateResponseError   `json:"errors"`
	Messages []RailgunUpdateResponseMessage `json:"messages"`
	Result   interface{}                    `json:"result"`
	// Whether the API call was successful
	Success RailgunUpdateResponseSuccess `json:"success"`
	JSON    railgunUpdateResponseJSON    `json:"-"`
}

// railgunUpdateResponseJSON contains the JSON metadata for the struct
// [RailgunUpdateResponse]
type railgunUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RailgunUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RailgunUpdateResponseError struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    railgunUpdateResponseErrorJSON `json:"-"`
}

// railgunUpdateResponseErrorJSON contains the JSON metadata for the struct
// [RailgunUpdateResponseError]
type railgunUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RailgunUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RailgunUpdateResponseMessage struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    railgunUpdateResponseMessageJSON `json:"-"`
}

// railgunUpdateResponseMessageJSON contains the JSON metadata for the struct
// [RailgunUpdateResponseMessage]
type railgunUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RailgunUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RailgunUpdateResponseSuccess bool

const (
	RailgunUpdateResponseSuccessTrue RailgunUpdateResponseSuccess = true
)

type RailgunDeleteResponse struct {
	Errors   []RailgunDeleteResponseError   `json:"errors"`
	Messages []RailgunDeleteResponseMessage `json:"messages"`
	Result   RailgunDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success RailgunDeleteResponseSuccess `json:"success"`
	JSON    railgunDeleteResponseJSON    `json:"-"`
}

// railgunDeleteResponseJSON contains the JSON metadata for the struct
// [RailgunDeleteResponse]
type railgunDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RailgunDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RailgunDeleteResponseError struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    railgunDeleteResponseErrorJSON `json:"-"`
}

// railgunDeleteResponseErrorJSON contains the JSON metadata for the struct
// [RailgunDeleteResponseError]
type railgunDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RailgunDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RailgunDeleteResponseMessage struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    railgunDeleteResponseMessageJSON `json:"-"`
}

// railgunDeleteResponseMessageJSON contains the JSON metadata for the struct
// [RailgunDeleteResponseMessage]
type railgunDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RailgunDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RailgunDeleteResponseResult struct {
	// Railgun identifier tag.
	ID   string                          `json:"id"`
	JSON railgunDeleteResponseResultJSON `json:"-"`
}

// railgunDeleteResponseResultJSON contains the JSON metadata for the struct
// [RailgunDeleteResponseResult]
type railgunDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RailgunDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RailgunDeleteResponseSuccess bool

const (
	RailgunDeleteResponseSuccessTrue RailgunDeleteResponseSuccess = true
)

type RailgunRailgunNewRailgunResponse struct {
	Errors   []RailgunRailgunNewRailgunResponseError   `json:"errors"`
	Messages []RailgunRailgunNewRailgunResponseMessage `json:"messages"`
	Result   interface{}                               `json:"result"`
	// Whether the API call was successful
	Success RailgunRailgunNewRailgunResponseSuccess `json:"success"`
	JSON    railgunRailgunNewRailgunResponseJSON    `json:"-"`
}

// railgunRailgunNewRailgunResponseJSON contains the JSON metadata for the struct
// [RailgunRailgunNewRailgunResponse]
type railgunRailgunNewRailgunResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RailgunRailgunNewRailgunResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RailgunRailgunNewRailgunResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    railgunRailgunNewRailgunResponseErrorJSON `json:"-"`
}

// railgunRailgunNewRailgunResponseErrorJSON contains the JSON metadata for the
// struct [RailgunRailgunNewRailgunResponseError]
type railgunRailgunNewRailgunResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RailgunRailgunNewRailgunResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RailgunRailgunNewRailgunResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    railgunRailgunNewRailgunResponseMessageJSON `json:"-"`
}

// railgunRailgunNewRailgunResponseMessageJSON contains the JSON metadata for the
// struct [RailgunRailgunNewRailgunResponseMessage]
type railgunRailgunNewRailgunResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RailgunRailgunNewRailgunResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RailgunRailgunNewRailgunResponseSuccess bool

const (
	RailgunRailgunNewRailgunResponseSuccessTrue RailgunRailgunNewRailgunResponseSuccess = true
)

type RailgunRailgunListRailgunsResponse struct {
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
	Status RailgunRailgunListRailgunsResponseStatus `json:"status,required"`
	// The version of the Railgun receiver.
	Version string `json:"version,required"`
	// The number of zones using this Railgun.
	ZonesConnected float64 `json:"zones_connected,required"`
	// Defined when the Railgun version is out of date from the latest release from
	// Cloudflare.
	UpgradeInfo RailgunRailgunListRailgunsResponseUpgradeInfo `json:"upgrade_info"`
	JSON        railgunRailgunListRailgunsResponseJSON        `json:"-"`
}

// railgunRailgunListRailgunsResponseJSON contains the JSON metadata for the struct
// [RailgunRailgunListRailgunsResponse]
type railgunRailgunListRailgunsResponseJSON struct {
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

func (r *RailgunRailgunListRailgunsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the Railgun.
type RailgunRailgunListRailgunsResponseStatus string

const (
	RailgunRailgunListRailgunsResponseStatusInitializing RailgunRailgunListRailgunsResponseStatus = "initializing"
	RailgunRailgunListRailgunsResponseStatusActive       RailgunRailgunListRailgunsResponseStatus = "active"
)

// Defined when the Railgun version is out of date from the latest release from
// Cloudflare.
type RailgunRailgunListRailgunsResponseUpgradeInfo struct {
	// An HTTP link to download the latest Railgun binary.
	DownloadLink string `json:"download_link"`
	// Latest version of the Railgun receiver available to install.
	LatestVersion string                                            `json:"latest_version"`
	JSON          railgunRailgunListRailgunsResponseUpgradeInfoJSON `json:"-"`
}

// railgunRailgunListRailgunsResponseUpgradeInfoJSON contains the JSON metadata for
// the struct [RailgunRailgunListRailgunsResponseUpgradeInfo]
type railgunRailgunListRailgunsResponseUpgradeInfoJSON struct {
	DownloadLink  apijson.Field
	LatestVersion apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RailgunRailgunListRailgunsResponseUpgradeInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RailgunUpdateParams struct {
	// Flag to determine if the Railgun is accepting connections.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RailgunUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RailgunRailgunNewRailgunParams struct {
	// Readable identifier of the Railgun.
	Name param.Field[string] `json:"name,required"`
}

func (r RailgunRailgunNewRailgunParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RailgunRailgunListRailgunsParams struct {
	// Sort Railguns in ascending or descending order.
	Direction param.Field[RailgunRailgunListRailgunsParamsDirection] `query:"direction"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of items per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [RailgunRailgunListRailgunsParams]'s query parameters as
// `url.Values`.
func (r RailgunRailgunListRailgunsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort Railguns in ascending or descending order.
type RailgunRailgunListRailgunsParamsDirection string

const (
	RailgunRailgunListRailgunsParamsDirectionAsc  RailgunRailgunListRailgunsParamsDirection = "asc"
	RailgunRailgunListRailgunsParamsDirectionDesc RailgunRailgunListRailgunsParamsDirection = "desc"
)
