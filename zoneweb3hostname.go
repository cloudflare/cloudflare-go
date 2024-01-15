// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneWeb3HostnameService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneWeb3HostnameService] method
// instead.
type ZoneWeb3HostnameService struct {
	Options            []option.RequestOption
	IpfsUniversalPaths *ZoneWeb3HostnameIpfsUniversalPathService
}

// NewZoneWeb3HostnameService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneWeb3HostnameService(opts ...option.RequestOption) (r *ZoneWeb3HostnameService) {
	r = &ZoneWeb3HostnameService{}
	r.Options = opts
	r.IpfsUniversalPaths = NewZoneWeb3HostnameIpfsUniversalPathService(opts...)
	return
}

// Web3 Hostname Details
func (r *ZoneWeb3HostnameService) Get(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *ZoneWeb3HostnameGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Edit Web3 Hostname
func (r *ZoneWeb3HostnameService) Update(ctx context.Context, zoneIdentifier string, identifier string, body ZoneWeb3HostnameUpdateParams, opts ...option.RequestOption) (res *ZoneWeb3HostnameUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Delete Web3 Hostname
func (r *ZoneWeb3HostnameService) Delete(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *ZoneWeb3HostnameDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create Web3 Hostname
func (r *ZoneWeb3HostnameService) Web3HostnameNewWeb3Hostname(ctx context.Context, zoneIdentifier string, body ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameParams, opts ...option.RequestOption) (res *ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/web3/hostnames", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List Web3 Hostnames
func (r *ZoneWeb3HostnameService) Web3HostnameListWeb3Hostnames(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/web3/hostnames", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneWeb3HostnameGetResponse struct {
	Errors   []ZoneWeb3HostnameGetResponseError   `json:"errors"`
	Messages []ZoneWeb3HostnameGetResponseMessage `json:"messages"`
	Result   ZoneWeb3HostnameGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneWeb3HostnameGetResponseSuccess `json:"success"`
	JSON    zoneWeb3HostnameGetResponseJSON    `json:"-"`
}

// zoneWeb3HostnameGetResponseJSON contains the JSON metadata for the struct
// [ZoneWeb3HostnameGetResponse]
type zoneWeb3HostnameGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWeb3HostnameGetResponseError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    zoneWeb3HostnameGetResponseErrorJSON `json:"-"`
}

// zoneWeb3HostnameGetResponseErrorJSON contains the JSON metadata for the struct
// [ZoneWeb3HostnameGetResponseError]
type zoneWeb3HostnameGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWeb3HostnameGetResponseMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    zoneWeb3HostnameGetResponseMessageJSON `json:"-"`
}

// zoneWeb3HostnameGetResponseMessageJSON contains the JSON metadata for the struct
// [ZoneWeb3HostnameGetResponseMessage]
type zoneWeb3HostnameGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWeb3HostnameGetResponseResult struct {
	// Identifier
	ID        string    `json:"id"`
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the hostname.
	Description string `json:"description"`
	// DNSLink value used if the target is ipfs.
	Dnslink    string    `json:"dnslink"`
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The hostname that will point to the target gateway via CNAME.
	Name string `json:"name"`
	// Status of the hostname's activation.
	Status ZoneWeb3HostnameGetResponseResultStatus `json:"status"`
	// Target gateway of the hostname.
	Target ZoneWeb3HostnameGetResponseResultTarget `json:"target"`
	JSON   zoneWeb3HostnameGetResponseResultJSON   `json:"-"`
}

// zoneWeb3HostnameGetResponseResultJSON contains the JSON metadata for the struct
// [ZoneWeb3HostnameGetResponseResult]
type zoneWeb3HostnameGetResponseResultJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	Dnslink     apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	Status      apijson.Field
	Target      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the hostname's activation.
type ZoneWeb3HostnameGetResponseResultStatus string

const (
	ZoneWeb3HostnameGetResponseResultStatusActive   ZoneWeb3HostnameGetResponseResultStatus = "active"
	ZoneWeb3HostnameGetResponseResultStatusPending  ZoneWeb3HostnameGetResponseResultStatus = "pending"
	ZoneWeb3HostnameGetResponseResultStatusDeleting ZoneWeb3HostnameGetResponseResultStatus = "deleting"
	ZoneWeb3HostnameGetResponseResultStatusError    ZoneWeb3HostnameGetResponseResultStatus = "error"
)

// Target gateway of the hostname.
type ZoneWeb3HostnameGetResponseResultTarget string

const (
	ZoneWeb3HostnameGetResponseResultTargetEthereum          ZoneWeb3HostnameGetResponseResultTarget = "ethereum"
	ZoneWeb3HostnameGetResponseResultTargetIpfs              ZoneWeb3HostnameGetResponseResultTarget = "ipfs"
	ZoneWeb3HostnameGetResponseResultTargetIpfsUniversalPath ZoneWeb3HostnameGetResponseResultTarget = "ipfs_universal_path"
	ZoneWeb3HostnameGetResponseResultTargetPolygon           ZoneWeb3HostnameGetResponseResultTarget = "polygon"
)

// Whether the API call was successful
type ZoneWeb3HostnameGetResponseSuccess bool

const (
	ZoneWeb3HostnameGetResponseSuccessTrue ZoneWeb3HostnameGetResponseSuccess = true
)

type ZoneWeb3HostnameUpdateResponse struct {
	Errors   []ZoneWeb3HostnameUpdateResponseError   `json:"errors"`
	Messages []ZoneWeb3HostnameUpdateResponseMessage `json:"messages"`
	Result   ZoneWeb3HostnameUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneWeb3HostnameUpdateResponseSuccess `json:"success"`
	JSON    zoneWeb3HostnameUpdateResponseJSON    `json:"-"`
}

// zoneWeb3HostnameUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneWeb3HostnameUpdateResponse]
type zoneWeb3HostnameUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWeb3HostnameUpdateResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    zoneWeb3HostnameUpdateResponseErrorJSON `json:"-"`
}

// zoneWeb3HostnameUpdateResponseErrorJSON contains the JSON metadata for the
// struct [ZoneWeb3HostnameUpdateResponseError]
type zoneWeb3HostnameUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWeb3HostnameUpdateResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    zoneWeb3HostnameUpdateResponseMessageJSON `json:"-"`
}

// zoneWeb3HostnameUpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneWeb3HostnameUpdateResponseMessage]
type zoneWeb3HostnameUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWeb3HostnameUpdateResponseResult struct {
	// Identifier
	ID        string    `json:"id"`
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the hostname.
	Description string `json:"description"`
	// DNSLink value used if the target is ipfs.
	Dnslink    string    `json:"dnslink"`
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The hostname that will point to the target gateway via CNAME.
	Name string `json:"name"`
	// Status of the hostname's activation.
	Status ZoneWeb3HostnameUpdateResponseResultStatus `json:"status"`
	// Target gateway of the hostname.
	Target ZoneWeb3HostnameUpdateResponseResultTarget `json:"target"`
	JSON   zoneWeb3HostnameUpdateResponseResultJSON   `json:"-"`
}

// zoneWeb3HostnameUpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneWeb3HostnameUpdateResponseResult]
type zoneWeb3HostnameUpdateResponseResultJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	Dnslink     apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	Status      apijson.Field
	Target      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the hostname's activation.
type ZoneWeb3HostnameUpdateResponseResultStatus string

const (
	ZoneWeb3HostnameUpdateResponseResultStatusActive   ZoneWeb3HostnameUpdateResponseResultStatus = "active"
	ZoneWeb3HostnameUpdateResponseResultStatusPending  ZoneWeb3HostnameUpdateResponseResultStatus = "pending"
	ZoneWeb3HostnameUpdateResponseResultStatusDeleting ZoneWeb3HostnameUpdateResponseResultStatus = "deleting"
	ZoneWeb3HostnameUpdateResponseResultStatusError    ZoneWeb3HostnameUpdateResponseResultStatus = "error"
)

// Target gateway of the hostname.
type ZoneWeb3HostnameUpdateResponseResultTarget string

const (
	ZoneWeb3HostnameUpdateResponseResultTargetEthereum          ZoneWeb3HostnameUpdateResponseResultTarget = "ethereum"
	ZoneWeb3HostnameUpdateResponseResultTargetIpfs              ZoneWeb3HostnameUpdateResponseResultTarget = "ipfs"
	ZoneWeb3HostnameUpdateResponseResultTargetIpfsUniversalPath ZoneWeb3HostnameUpdateResponseResultTarget = "ipfs_universal_path"
	ZoneWeb3HostnameUpdateResponseResultTargetPolygon           ZoneWeb3HostnameUpdateResponseResultTarget = "polygon"
)

// Whether the API call was successful
type ZoneWeb3HostnameUpdateResponseSuccess bool

const (
	ZoneWeb3HostnameUpdateResponseSuccessTrue ZoneWeb3HostnameUpdateResponseSuccess = true
)

type ZoneWeb3HostnameDeleteResponse struct {
	Errors   []ZoneWeb3HostnameDeleteResponseError   `json:"errors"`
	Messages []ZoneWeb3HostnameDeleteResponseMessage `json:"messages"`
	Result   ZoneWeb3HostnameDeleteResponseResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success ZoneWeb3HostnameDeleteResponseSuccess `json:"success"`
	JSON    zoneWeb3HostnameDeleteResponseJSON    `json:"-"`
}

// zoneWeb3HostnameDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneWeb3HostnameDeleteResponse]
type zoneWeb3HostnameDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWeb3HostnameDeleteResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    zoneWeb3HostnameDeleteResponseErrorJSON `json:"-"`
}

// zoneWeb3HostnameDeleteResponseErrorJSON contains the JSON metadata for the
// struct [ZoneWeb3HostnameDeleteResponseError]
type zoneWeb3HostnameDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWeb3HostnameDeleteResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    zoneWeb3HostnameDeleteResponseMessageJSON `json:"-"`
}

// zoneWeb3HostnameDeleteResponseMessageJSON contains the JSON metadata for the
// struct [ZoneWeb3HostnameDeleteResponseMessage]
type zoneWeb3HostnameDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWeb3HostnameDeleteResponseResult struct {
	// Identifier
	ID   string                                   `json:"id,required"`
	JSON zoneWeb3HostnameDeleteResponseResultJSON `json:"-"`
}

// zoneWeb3HostnameDeleteResponseResultJSON contains the JSON metadata for the
// struct [ZoneWeb3HostnameDeleteResponseResult]
type zoneWeb3HostnameDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneWeb3HostnameDeleteResponseSuccess bool

const (
	ZoneWeb3HostnameDeleteResponseSuccessTrue ZoneWeb3HostnameDeleteResponseSuccess = true
)

type ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponse struct {
	Errors   []ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseError   `json:"errors"`
	Messages []ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseMessage `json:"messages"`
	Result   ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseSuccess `json:"success"`
	JSON    zoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseJSON    `json:"-"`
}

// zoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseJSON contains the JSON
// metadata for the struct [ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponse]
type zoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseError struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseErrorJSON `json:"-"`
}

// zoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseErrorJSON contains the JSON
// metadata for the struct
// [ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseError]
type zoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseMessage struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    zoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseMessageJSON `json:"-"`
}

// zoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseMessageJSON contains the JSON
// metadata for the struct
// [ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseMessage]
type zoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseResult struct {
	// Identifier
	ID        string    `json:"id"`
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the hostname.
	Description string `json:"description"`
	// DNSLink value used if the target is ipfs.
	Dnslink    string    `json:"dnslink"`
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The hostname that will point to the target gateway via CNAME.
	Name string `json:"name"`
	// Status of the hostname's activation.
	Status ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseResultStatus `json:"status"`
	// Target gateway of the hostname.
	Target ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseResultTarget `json:"target"`
	JSON   zoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseResultJSON   `json:"-"`
}

// zoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseResultJSON contains the JSON
// metadata for the struct
// [ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseResult]
type zoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseResultJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	Dnslink     apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	Status      apijson.Field
	Target      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the hostname's activation.
type ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseResultStatus string

const (
	ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseResultStatusActive   ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseResultStatus = "active"
	ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseResultStatusPending  ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseResultStatus = "pending"
	ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseResultStatusDeleting ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseResultStatus = "deleting"
	ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseResultStatusError    ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseResultStatus = "error"
)

// Target gateway of the hostname.
type ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseResultTarget string

const (
	ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseResultTargetEthereum          ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseResultTarget = "ethereum"
	ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseResultTargetIpfs              ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseResultTarget = "ipfs"
	ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseResultTargetIpfsUniversalPath ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseResultTarget = "ipfs_universal_path"
	ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseResultTargetPolygon           ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseResultTarget = "polygon"
)

// Whether the API call was successful
type ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseSuccess bool

const (
	ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseSuccessTrue ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameResponseSuccess = true
)

type ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponse struct {
	Errors     []ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseError    `json:"errors"`
	Messages   []ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseMessage  `json:"messages"`
	Result     []ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseResult   `json:"result"`
	ResultInfo ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseSuccess `json:"success"`
	JSON    zoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseJSON    `json:"-"`
}

// zoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseJSON contains the JSON
// metadata for the struct [ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponse]
type zoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseError struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    zoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseErrorJSON `json:"-"`
}

// zoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseErrorJSON contains the JSON
// metadata for the struct
// [ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseError]
type zoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseMessage struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    zoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseMessageJSON `json:"-"`
}

// zoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseMessageJSON contains the
// JSON metadata for the struct
// [ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseMessage]
type zoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseResult struct {
	// Identifier
	ID        string    `json:"id"`
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the hostname.
	Description string `json:"description"`
	// DNSLink value used if the target is ipfs.
	Dnslink    string    `json:"dnslink"`
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The hostname that will point to the target gateway via CNAME.
	Name string `json:"name"`
	// Status of the hostname's activation.
	Status ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseResultStatus `json:"status"`
	// Target gateway of the hostname.
	Target ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseResultTarget `json:"target"`
	JSON   zoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseResultJSON   `json:"-"`
}

// zoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseResultJSON contains the
// JSON metadata for the struct
// [ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseResult]
type zoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseResultJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	Dnslink     apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	Status      apijson.Field
	Target      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the hostname's activation.
type ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseResultStatus string

const (
	ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseResultStatusActive   ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseResultStatus = "active"
	ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseResultStatusPending  ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseResultStatus = "pending"
	ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseResultStatusDeleting ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseResultStatus = "deleting"
	ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseResultStatusError    ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseResultStatus = "error"
)

// Target gateway of the hostname.
type ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseResultTarget string

const (
	ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseResultTargetEthereum          ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseResultTarget = "ethereum"
	ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseResultTargetIpfs              ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseResultTarget = "ipfs"
	ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseResultTargetIpfsUniversalPath ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseResultTarget = "ipfs_universal_path"
	ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseResultTargetPolygon           ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseResultTarget = "polygon"
)

type ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                             `json:"total_count"`
	JSON       zoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseResultInfoJSON `json:"-"`
}

// zoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseResultInfoJSON contains the
// JSON metadata for the struct
// [ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseResultInfo]
type zoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseSuccess bool

const (
	ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseSuccessTrue ZoneWeb3HostnameWeb3HostnameListWeb3HostnamesResponseSuccess = true
)

type ZoneWeb3HostnameUpdateParams struct {
	// An optional description of the hostname.
	Description param.Field[string] `json:"description"`
	// DNSLink value used if the target is ipfs.
	Dnslink param.Field[string] `json:"dnslink"`
}

func (r ZoneWeb3HostnameUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameParams struct {
	// Target gateway of the hostname.
	Target param.Field[ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameParamsTarget] `json:"target,required"`
	// An optional description of the hostname.
	Description param.Field[string] `json:"description"`
	// DNSLink value used if the target is ipfs.
	Dnslink param.Field[string] `json:"dnslink"`
}

func (r ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Target gateway of the hostname.
type ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameParamsTarget string

const (
	ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameParamsTargetEthereum          ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameParamsTarget = "ethereum"
	ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameParamsTargetIpfs              ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameParamsTarget = "ipfs"
	ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameParamsTargetIpfsUniversalPath ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameParamsTarget = "ipfs_universal_path"
	ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameParamsTargetPolygon           ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameParamsTarget = "polygon"
)
