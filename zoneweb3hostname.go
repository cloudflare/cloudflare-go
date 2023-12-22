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
func (r *ZoneWeb3HostnameService) Get(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *Web3HostnameSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Edit Web3 Hostname
func (r *ZoneWeb3HostnameService) Update(ctx context.Context, zoneIdentifier string, identifier string, body ZoneWeb3HostnameUpdateParams, opts ...option.RequestOption) (res *Web3HostnameSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Delete Web3 Hostname
func (r *ZoneWeb3HostnameService) Delete(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *APIResponseSingleIDKYar7dC1, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create Web3 Hostname
func (r *ZoneWeb3HostnameService) Web3HostnameNewWeb3Hostname(ctx context.Context, zoneIdentifier string, body ZoneWeb3HostnameWeb3HostnameNewWeb3HostnameParams, opts ...option.RequestOption) (res *Web3HostnameSingleResponse, err error) {
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

type Web3HostnameSingleResponse struct {
	Errors   []Web3HostnameSingleResponseError   `json:"errors"`
	Messages []Web3HostnameSingleResponseMessage `json:"messages"`
	Result   Web3HostnameSingleResponseResult    `json:"result"`
	// Whether the API call was successful
	Success Web3HostnameSingleResponseSuccess `json:"success"`
	JSON    web3HostnameSingleResponseJSON    `json:"-"`
}

// web3HostnameSingleResponseJSON contains the JSON metadata for the struct
// [Web3HostnameSingleResponse]
type web3HostnameSingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameSingleResponseError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    web3HostnameSingleResponseErrorJSON `json:"-"`
}

// web3HostnameSingleResponseErrorJSON contains the JSON metadata for the struct
// [Web3HostnameSingleResponseError]
type web3HostnameSingleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameSingleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameSingleResponseMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    web3HostnameSingleResponseMessageJSON `json:"-"`
}

// web3HostnameSingleResponseMessageJSON contains the JSON metadata for the struct
// [Web3HostnameSingleResponseMessage]
type web3HostnameSingleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameSingleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameSingleResponseResult struct {
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
	Status Web3HostnameSingleResponseResultStatus `json:"status"`
	// Target gateway of the hostname.
	Target Web3HostnameSingleResponseResultTarget `json:"target"`
	JSON   web3HostnameSingleResponseResultJSON   `json:"-"`
}

// web3HostnameSingleResponseResultJSON contains the JSON metadata for the struct
// [Web3HostnameSingleResponseResult]
type web3HostnameSingleResponseResultJSON struct {
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

func (r *Web3HostnameSingleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the hostname's activation.
type Web3HostnameSingleResponseResultStatus string

const (
	Web3HostnameSingleResponseResultStatusActive   Web3HostnameSingleResponseResultStatus = "active"
	Web3HostnameSingleResponseResultStatusPending  Web3HostnameSingleResponseResultStatus = "pending"
	Web3HostnameSingleResponseResultStatusDeleting Web3HostnameSingleResponseResultStatus = "deleting"
	Web3HostnameSingleResponseResultStatusError    Web3HostnameSingleResponseResultStatus = "error"
)

// Target gateway of the hostname.
type Web3HostnameSingleResponseResultTarget string

const (
	Web3HostnameSingleResponseResultTargetEthereum          Web3HostnameSingleResponseResultTarget = "ethereum"
	Web3HostnameSingleResponseResultTargetIpfs              Web3HostnameSingleResponseResultTarget = "ipfs"
	Web3HostnameSingleResponseResultTargetIpfsUniversalPath Web3HostnameSingleResponseResultTarget = "ipfs_universal_path"
	Web3HostnameSingleResponseResultTargetPolygon           Web3HostnameSingleResponseResultTarget = "polygon"
)

// Whether the API call was successful
type Web3HostnameSingleResponseSuccess bool

const (
	Web3HostnameSingleResponseSuccessTrue Web3HostnameSingleResponseSuccess = true
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
