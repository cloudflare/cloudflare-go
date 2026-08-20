// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// DEXDeviceISPService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDEXDeviceISPService] method instead.
type DEXDeviceISPService struct {
	Options []option.RequestOption
}

// NewDEXDeviceISPService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDEXDeviceISPService(opts ...option.RequestOption) (r *DEXDeviceISPService) {
	r = &DEXDeviceISPService{}
	r.Options = opts
	return
}

// List ISP information observed for a specific device during traceroute tests.
func (r *DEXDeviceISPService) List(ctx context.Context, deviceID string, params DEXDeviceISPListParams, opts ...option.RequestOption) (res *pagination.V4PagePagination[ISPs], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if deviceID == "" {
		err = errors.New("missing required device_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dex/devices/%s/isps", params.AccountID, deviceID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// List ISP information observed for a specific device during traceroute tests.
func (r *DEXDeviceISPService) ListAutoPaging(ctx context.Context, deviceID string, params DEXDeviceISPListParams, opts ...option.RequestOption) *pagination.V4PagePaginationAutoPager[ISPs] {
	return pagination.NewV4PagePaginationAutoPager(r.List(ctx, deviceID, params, opts...))
}

type ISPs struct {
	ISPs []ISPsISP `json:"isps" api:"required"`
	JSON ispsJSON  `json:"-"`
}

// ispsJSON contains the JSON metadata for the struct [ISPs]
type ispsJSON struct {
	ISPs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ISPs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ispsJSON) RawJSON() string {
	return r.raw
}

type ISPsISP struct {
	// The test that generated this result.
	TestID string `json:"test_id" api:"required"`
	// The specific test result.
	TestResultID string `json:"test_result_id" api:"required"`
	// Timestamp of when the ISP was observed.
	TimeStart time.Time `json:"time_start" api:"required" format:"date-time"`
	// IP address information for the ISP hop. Fields marked as PII-gated (`name`,
	// `address`, `netmask`, and all `location` sub-fields) will be returned as the
	// literal string `"REDACTED"` for callers that do not have the PII permission.
	// `asn`, `aso`, and `version` are always returned regardless of PII access.
	IP   ISPsISPsIP  `json:"ip"`
	JSON ispsISPJSON `json:"-"`
}

// ispsISPJSON contains the JSON metadata for the struct [ISPsISP]
type ispsISPJSON struct {
	TestID       apijson.Field
	TestResultID apijson.Field
	TimeStart    apijson.Field
	IP           apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ISPsISP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ispsISPJSON) RawJSON() string {
	return r.raw
}

// ISPsISPsIP iP address information for the ISP hop. Fields marked as PII-gated (`name`,
// `address`, `netmask`, and all `location` sub-fields) will be returned as the
// literal string `"REDACTED"` for callers that do not have the PII permission.
// `asn`, `aso`, and `version` are always returned regardless of PII access.
type ISPsISPsIP struct {
	// IP address. Returned as `"REDACTED"` without PII permission.
	Address string `json:"address"`
	// Autonomous System Number.
	ASN int64 `json:"asn"`
	// Autonomous System Organization name.
	Aso string `json:"aso"`
	// Geographic location information. All fields are returned as the literal string
	// `"REDACTED"` for callers that do not have the PII permission.
	Location ISPsISPsIPLocation `json:"location"`
	// Named IP address (reverse DNS hostname when available). Returned as `"REDACTED"`
	// without PII permission.
	Name string `json:"name"`
	// Network mask. Returned as `"REDACTED"` without PII permission.
	Netmask string `json:"netmask"`
	// IP version (`1` for IPv4, `2` for IPv6, `0` if unknown).
	Version int64          `json:"version"`
	JSON    ispsISPsIPJSON `json:"-"`
}

// ispsISPsIPJSON contains the JSON metadata for the struct [ISPsISPsIP]
type ispsISPsIPJSON struct {
	Address     apijson.Field
	ASN         apijson.Field
	Aso         apijson.Field
	Location    apijson.Field
	Name        apijson.Field
	Netmask     apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ISPsISPsIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ispsISPsIPJSON) RawJSON() string {
	return r.raw
}

// ISPsISPsIPLocation geographic location information. All fields are returned as the literal string
// `"REDACTED"` for callers that do not have the PII permission.
type ISPsISPsIPLocation struct {
	// City name. Returned as `"REDACTED"` without PII permission.
	City string `json:"city"`
	// Country ISO code. Returned as `"REDACTED"` without PII permission.
	CountryISO string `json:"country_iso"`
	// State/province ISO code. Returned as `"REDACTED"` without PII permission.
	StateISO string `json:"state_iso"`
	// ZIP/postal code. Returned as `"REDACTED"` without PII permission.
	Zip  string                 `json:"zip"`
	JSON ispsISPsIPLocationJSON `json:"-"`
}

// ispsISPsIPLocationJSON contains the JSON metadata for the struct
// [ISPsISPsIPLocation]
type ispsISPsIPLocationJSON struct {
	City        apijson.Field
	CountryISO  apijson.Field
	StateISO    apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ISPsISPsIPLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ispsISPsIPLocationJSON) RawJSON() string {
	return r.raw
}

type DEXDeviceISPListParams struct {
	// Unique identifier linked to an account.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Number of items per page
	PerPage param.Field[int64] `query:"per_page" api:"required"`
	// Cursor for cursor-based pagination. Mutually exclusive with page.
	Cursor param.Field[string] `query:"cursor"`
	// Start time for the query in ISO 8601 format.
	From param.Field[time.Time] `query:"from" format:"date-time"`
	// Page number of paginated results. Mutually exclusive with cursor.
	Page param.Field[int64] `query:"page"`
	// The field to sort results by.
	SortBy param.Field[DEXDeviceISPListParamsSortBy] `query:"sort_by"`
	// The order to sort results.
	SortOrder param.Field[DEXDeviceISPListParamsSortOrder] `query:"sort_order"`
	// End time for the query in ISO 8601 format.
	To param.Field[time.Time] `query:"to" format:"date-time"`
}

// URLQuery serializes [DEXDeviceISPListParams]'s query parameters as `url.Values`.
func (r DEXDeviceISPListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// DEXDeviceISPListParamsSortBy is the field to sort results by.
type DEXDeviceISPListParamsSortBy string

const (
	DEXDeviceISPListParamsSortByTimeStart DEXDeviceISPListParamsSortBy = "time_start"
)

func (r DEXDeviceISPListParamsSortBy) IsKnown() bool {
	switch r {
	case DEXDeviceISPListParamsSortByTimeStart:
		return true
	}
	return false
}

// DEXDeviceISPListParamsSortOrder is the order to sort results.
type DEXDeviceISPListParamsSortOrder string

const (
	DEXDeviceISPListParamsSortOrderAsc  DEXDeviceISPListParamsSortOrder = "ASC"
	DEXDeviceISPListParamsSortOrderDesc DEXDeviceISPListParamsSortOrder = "DESC"
)

func (r DEXDeviceISPListParamsSortOrder) IsKnown() bool {
	switch r {
	case DEXDeviceISPListParamsSortOrderAsc, DEXDeviceISPListParamsSortOrderDesc:
		return true
	}
	return false
}
