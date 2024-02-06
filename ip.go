// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// IPService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewIPService] method instead.
type IPService struct {
	Options []option.RequestOption
}

// NewIPService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewIPService(opts ...option.RequestOption) (r *IPService) {
	r = &IPService{}
	r.Options = opts
	return
}

// Get IPs used on the Cloudflare/JD Cloud network, see
// https://www.cloudflare.com/ips for Cloudflare IPs or
// https://developers.cloudflare.com/china-network/reference/infrastructure/ for JD
// Cloud IPs.
func (r *IPService) List(ctx context.Context, query IPListParams, opts ...option.RequestOption) (res *IPListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "ips"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type IPListResponse struct {
	Errors   []IPListResponseError   `json:"errors"`
	Messages []IPListResponseMessage `json:"messages"`
	Result   IPListResponseResult    `json:"result"`
	// Whether the API call was successful
	Success IPListResponseSuccess `json:"success"`
	JSON    ipListResponseJSON    `json:"-"`
}

// ipListResponseJSON contains the JSON metadata for the struct [IPListResponse]
type ipListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IPListResponseError struct {
	Code    int64                   `json:"code,required"`
	Message string                  `json:"message,required"`
	JSON    ipListResponseErrorJSON `json:"-"`
}

// ipListResponseErrorJSON contains the JSON metadata for the struct
// [IPListResponseError]
type ipListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IPListResponseMessage struct {
	Code    int64                     `json:"code,required"`
	Message string                    `json:"message,required"`
	JSON    ipListResponseMessageJSON `json:"-"`
}

// ipListResponseMessageJSON contains the JSON metadata for the struct
// [IPListResponseMessage]
type ipListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [IPListResponseResult12h4PjeVIPs] or
// [IPListResponseResult12h4PjeVIPsJdcloud].
type IPListResponseResult interface {
	implementsIPListResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*IPListResponseResult)(nil)).Elem(), "")
}

type IPListResponseResult12h4PjeVIPs struct {
	// A digest of the IP data. Useful for determining if the data has changed.
	Etag string `json:"etag"`
	// List of Cloudflare IPv4 CIDR addresses.
	Ipv4Cidrs []string `json:"ipv4_cidrs"`
	// List of Cloudflare IPv6 CIDR addresses.
	IPV6Cidrs []string                            `json:"ipv6_cidrs"`
	JSON      ipListResponseResult12h4PjeViPsJSON `json:"-"`
}

// ipListResponseResult12h4PjeViPsJSON contains the JSON metadata for the struct
// [IPListResponseResult12h4PjeVIPs]
type ipListResponseResult12h4PjeViPsJSON struct {
	Etag        apijson.Field
	Ipv4Cidrs   apijson.Field
	IPV6Cidrs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPListResponseResult12h4PjeVIPs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r IPListResponseResult12h4PjeVIPs) implementsIPListResponseResult() {}

type IPListResponseResult12h4PjeVIPsJdcloud struct {
	// A digest of the IP data. Useful for determining if the data has changed.
	Etag string `json:"etag"`
	// List of Cloudflare IPv4 CIDR addresses.
	Ipv4Cidrs []string `json:"ipv4_cidrs"`
	// List of Cloudflare IPv6 CIDR addresses.
	IPV6Cidrs []string `json:"ipv6_cidrs"`
	// List IPv4 and IPv6 CIDRs, only populated if `?networks=jdcloud` is used.
	JdcloudCidrs []string                                   `json:"jdcloud_cidrs"`
	JSON         ipListResponseResult12h4PjeViPsJdcloudJSON `json:"-"`
}

// ipListResponseResult12h4PjeViPsJdcloudJSON contains the JSON metadata for the
// struct [IPListResponseResult12h4PjeVIPsJdcloud]
type ipListResponseResult12h4PjeViPsJdcloudJSON struct {
	Etag         apijson.Field
	Ipv4Cidrs    apijson.Field
	IPV6Cidrs    apijson.Field
	JdcloudCidrs apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IPListResponseResult12h4PjeVIPsJdcloud) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r IPListResponseResult12h4PjeVIPsJdcloud) implementsIPListResponseResult() {}

// Whether the API call was successful
type IPListResponseSuccess bool

const (
	IPListResponseSuccessTrue IPListResponseSuccess = true
)

type IPListParams struct {
	// Specified as `jdcloud` to list IPs used by JD Cloud data centers.
	Networks param.Field[string] `query:"networks"`
}

// URLQuery serializes [IPListParams]'s query parameters as `url.Values`.
func (r IPListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
