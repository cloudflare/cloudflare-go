// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneWeb3HostnameIpfsUniversalPathContentListService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneWeb3HostnameIpfsUniversalPathContentListService] method instead.
type ZoneWeb3HostnameIpfsUniversalPathContentListService struct {
	Options []option.RequestOption
	Entries *ZoneWeb3HostnameIpfsUniversalPathContentListEntryService
}

// NewZoneWeb3HostnameIpfsUniversalPathContentListService generates a new service
// that applies the given options to each request. These options are applied after
// the parent client's options (if there is one), and before any request-specific
// options.
func NewZoneWeb3HostnameIpfsUniversalPathContentListService(opts ...option.RequestOption) (r *ZoneWeb3HostnameIpfsUniversalPathContentListService) {
	r = &ZoneWeb3HostnameIpfsUniversalPathContentListService{}
	r.Options = opts
	r.Entries = NewZoneWeb3HostnameIpfsUniversalPathContentListEntryService(opts...)
	return
}

// IPFS Universal Path Gateway Content List Details
func (r *ZoneWeb3HostnameIpfsUniversalPathContentListService) Web3HostnameIpfsUniversalPathGatewayContentListDetails(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update IPFS Universal Path Gateway Content List
func (r *ZoneWeb3HostnameIpfsUniversalPathContentListService) Web3HostnameUpdateIpfsUniversalPathGatewayContentList(ctx context.Context, zoneIdentifier string, identifier string, body ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListParams, opts ...option.RequestOption) (res *ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponse struct {
	Errors   []ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseError   `json:"errors"`
	Messages []ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseMessage `json:"messages"`
	Result   ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseSuccess `json:"success"`
	JSON    zoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseJSON    `json:"-"`
}

// zoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseJSON
// contains the JSON metadata for the struct
// [ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponse]
type zoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseError struct {
	Code    int64                                                                                                               `json:"code,required"`
	Message string                                                                                                              `json:"message,required"`
	JSON    zoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseErrorJSON `json:"-"`
}

// zoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseError]
type zoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseMessage struct {
	Code    int64                                                                                                                 `json:"code,required"`
	Message string                                                                                                                `json:"message,required"`
	JSON    zoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseMessageJSON `json:"-"`
}

// zoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseMessage]
type zoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseResult struct {
	// Behavior of the content list.
	Action ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseResultAction `json:"action"`
	JSON   zoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseResultJSON   `json:"-"`
}

// zoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseResult]
type zoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseResultJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Behavior of the content list.
type ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseResultAction string

const (
	ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseResultActionBlock ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseResultAction = "block"
)

// Whether the API call was successful
type ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseSuccess bool

const (
	ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseSuccessTrue ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseSuccess = true
)

type ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponse struct {
	Errors   []ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseError   `json:"errors"`
	Messages []ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseMessage `json:"messages"`
	Result   ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseSuccess `json:"success"`
	JSON    zoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseJSON    `json:"-"`
}

// zoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseJSON
// contains the JSON metadata for the struct
// [ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponse]
type zoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseError struct {
	Code    int64                                                                                                              `json:"code,required"`
	Message string                                                                                                             `json:"message,required"`
	JSON    zoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseErrorJSON `json:"-"`
}

// zoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseError]
type zoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseMessage struct {
	Code    int64                                                                                                                `json:"code,required"`
	Message string                                                                                                               `json:"message,required"`
	JSON    zoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseMessageJSON `json:"-"`
}

// zoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseMessage]
type zoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseResult struct {
	// Behavior of the content list.
	Action ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseResultAction `json:"action"`
	JSON   zoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseResultJSON   `json:"-"`
}

// zoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseResult]
type zoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseResultJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Behavior of the content list.
type ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseResultAction string

const (
	ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseResultActionBlock ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseResultAction = "block"
)

// Whether the API call was successful
type ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseSuccess bool

const (
	ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseSuccessTrue ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseSuccess = true
)

type ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListParams struct {
	// Behavior of the content list.
	Action param.Field[ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListParamsAction] `json:"action,required"`
	// Content list entries.
	Entries param.Field[[]ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListParamsEntry] `json:"entries,required"`
}

func (r ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Behavior of the content list.
type ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListParamsAction string

const (
	ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListParamsActionBlock ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListParamsAction = "block"
)

// Content list entry to be blocked.
type ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListParamsEntry struct {
	// CID or content path of content to block.
	Content param.Field[string] `json:"content"`
	// An optional description of the content list entry.
	Description param.Field[string] `json:"description"`
	// Type of content list entry to block.
	Type param.Field[ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListParamsEntriesType] `json:"type"`
}

func (r ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListParamsEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of content list entry to block.
type ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListParamsEntriesType string

const (
	ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListParamsEntriesTypeCid         ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListParamsEntriesType = "cid"
	ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListParamsEntriesTypeContentPath ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListParamsEntriesType = "content_path"
)
