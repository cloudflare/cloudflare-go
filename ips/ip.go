// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ips

import (
	"context"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
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
	var env IPListResponseEnvelope
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

// Union satisfied by [ips.IPs] or [ips.JDCloudIPs].
type IPListResponse interface {
	implementsIPsIPListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IPListResponse)(nil)).Elem(),
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
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IPListResponseEnvelope struct {
	Errors   []IPListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IPListResponseEnvelopeMessages `json:"messages,required"`
	Result   IPListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success IPListResponseEnvelopeSuccess `json:"success,required"`
	JSON    ipListResponseEnvelopeJSON    `json:"-"`
}

// ipListResponseEnvelopeJSON contains the JSON metadata for the struct
// [IPListResponseEnvelope]
type ipListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type IPListResponseEnvelopeErrors struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    ipListResponseEnvelopeErrorsJSON `json:"-"`
}

// ipListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [IPListResponseEnvelopeErrors]
type ipListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IPListResponseEnvelopeMessages struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    ipListResponseEnvelopeMessagesJSON `json:"-"`
}

// ipListResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [IPListResponseEnvelopeMessages]
type ipListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipListResponseEnvelopeMessagesJSON) RawJSON() string {
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
