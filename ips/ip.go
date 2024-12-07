// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ips

import (
	"context"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
	"github.com/tidwall/gjson"
)

// IPService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIPService] method instead.
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
	var env IPListResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "ips"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IPs struct {
	// A digest of the IP data. Useful for determining if the data has changed.
	Etag string `json:"etag"`
	// List of Cloudflare IPv4 CIDR addresses.
	IPV4CIDRs []string `json:"ipv4_cidrs"`
	// List of Cloudflare IPv6 CIDR addresses.
	IPV6CIDRs []string `json:"ipv6_cidrs"`
	JSON      ipsJSON  `json:"-"`
}

// ipsJSON contains the JSON metadata for the struct [IPs]
type ipsJSON struct {
	Etag        apijson.Field
	IPV4CIDRs   apijson.Field
	IPV6CIDRs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsJSON) RawJSON() string {
	return r.raw
}

func (r IPs) implementsIPsIPListResponse() {}

type JDCloudIPs struct {
	// A digest of the IP data. Useful for determining if the data has changed.
	Etag string `json:"etag"`
	// List of Cloudflare IPv4 CIDR addresses.
	IPV4CIDRs []string `json:"ipv4_cidrs"`
	// List of Cloudflare IPv6 CIDR addresses.
	IPV6CIDRs []string `json:"ipv6_cidrs"`
	// List IPv4 and IPv6 CIDRs, only populated if `?networks=jdcloud` is used.
	JDCloudCIDRs []string       `json:"jdcloud_cidrs"`
	JSON         JDCloudIPsJSON `json:"-"`
}

// JDCloudIPsJSON contains the JSON metadata for the struct [JDCloudIPs]
type JDCloudIPsJSON struct {
	Etag         apijson.Field
	IPV4CIDRs    apijson.Field
	IPV6CIDRs    apijson.Field
	JDCloudCIDRs apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *JDCloudIPs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r JDCloudIPsJSON) RawJSON() string {
	return r.raw
}

func (r JDCloudIPs) implementsIPsIPListResponse() {}

type IPListResponse struct {
	// A digest of the IP data. Useful for determining if the data has changed.
	Etag string `json:"etag"`
	// This field can have the runtime type of [[]string].
	IPV4CIDRs interface{} `json:"ipv4_cidrs"`
	// This field can have the runtime type of [[]string].
	IPV6CIDRs interface{} `json:"ipv6_cidrs"`
	// This field can have the runtime type of [[]string].
	JDCloudCIDRs interface{}        `json:"jdcloud_cidrs"`
	JSON         ipListResponseJSON `json:"-"`
	union        IPListResponseUnion
}

// ipListResponseJSON contains the JSON metadata for the struct [IPListResponse]
type ipListResponseJSON struct {
	Etag         apijson.Field
	IPV4CIDRs    apijson.Field
	IPV6CIDRs    apijson.Field
	JDCloudCIDRs apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r ipListResponseJSON) RawJSON() string {
	return r.raw
}

func (r *IPListResponse) UnmarshalJSON(data []byte) (err error) {
	*r = IPListResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [IPListResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [ips.IPs], [ips.JDCloudIPs].
func (r IPListResponse) AsUnion() IPListResponseUnion {
	return r.union
}

// Union satisfied by [ips.IPs] or [ips.JDCloudIPs].
type IPListResponseUnion interface {
	implementsIPsIPListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IPListResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IPs{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(JDCloudIPs{}),
		},
	)
}

type IPListParams struct {
	// Specified as `jdcloud` to list IPs used by JD Cloud data centers.
	Networks param.Field[string] `query:"networks"`
}

// URLQuery serializes [IPListParams]'s query parameters as `url.Values`.
func (r IPListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type IPListResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success IPListResponseEnvelopeSuccess `json:"success,required"`
	Result  IPListResponse                `json:"result"`
	JSON    ipListResponseEnvelopeJSON    `json:"-"`
}

// ipListResponseEnvelopeJSON contains the JSON metadata for the struct
// [IPListResponseEnvelope]
type ipListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IPListResponseEnvelopeSuccess bool

const (
	IPListResponseEnvelopeSuccessTrue IPListResponseEnvelopeSuccess = true
)

func (r IPListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IPListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
