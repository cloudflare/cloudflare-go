// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
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

// Get Cloudflare IPs.
func (r *IPService) List(ctx context.Context, opts ...option.RequestOption) (res *IPListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "ips"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type IP struct {
	// A digest of the IP data. Useful for determining if the data has changed.
	Etag string `json:"etag"`
	// List of Cloudflare IPv4 CIDR addresses.
	Ipv4Cidrs []string `json:"ipv4_cidrs"`
	// List of Cloudflare IPv6 CIDR addresses.
	Ipv6Cidrs []string `json:"ipv6_cidrs"`
	JSON      ipJSON   `json:"-"`
}

// ipJSON contains the JSON metadata for the struct [IP]
type ipJSON struct {
	Etag        apijson.Field
	Ipv4Cidrs   apijson.Field
	Ipv6Cidrs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IPListResponse struct {
	Errors   []IPListResponseError   `json:"errors"`
	Messages []IPListResponseMessage `json:"messages"`
	Result   IP                      `json:"result"`
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

// Whether the API call was successful
type IPListResponseSuccess bool

const (
	IPListResponseSuccessTrue IPListResponseSuccess = true
)
