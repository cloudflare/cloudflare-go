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
func (r *ZoneWeb3HostnameIpfsUniversalPathContentListService) Web3HostnameIpfsUniversalPathGatewayContentListDetails(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *ContentListDetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update IPFS Universal Path Gateway Content List
func (r *ZoneWeb3HostnameIpfsUniversalPathContentListService) Web3HostnameUpdateIpfsUniversalPathGatewayContentList(ctx context.Context, zoneIdentifier string, identifier string, body ZoneWeb3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListParams, opts ...option.RequestOption) (res *ContentListDetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ContentListDetailsResponse struct {
	Errors   []ContentListDetailsResponseError   `json:"errors"`
	Messages []ContentListDetailsResponseMessage `json:"messages"`
	Result   ContentListDetailsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ContentListDetailsResponseSuccess `json:"success"`
	JSON    contentListDetailsResponseJSON    `json:"-"`
}

// contentListDetailsResponseJSON contains the JSON metadata for the struct
// [ContentListDetailsResponse]
type contentListDetailsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentListDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContentListDetailsResponseError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    contentListDetailsResponseErrorJSON `json:"-"`
}

// contentListDetailsResponseErrorJSON contains the JSON metadata for the struct
// [ContentListDetailsResponseError]
type contentListDetailsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentListDetailsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContentListDetailsResponseMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    contentListDetailsResponseMessageJSON `json:"-"`
}

// contentListDetailsResponseMessageJSON contains the JSON metadata for the struct
// [ContentListDetailsResponseMessage]
type contentListDetailsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentListDetailsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContentListDetailsResponseResult struct {
	// Behavior of the content list.
	Action ContentListDetailsResponseResultAction `json:"action"`
	JSON   contentListDetailsResponseResultJSON   `json:"-"`
}

// contentListDetailsResponseResultJSON contains the JSON metadata for the struct
// [ContentListDetailsResponseResult]
type contentListDetailsResponseResultJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentListDetailsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Behavior of the content list.
type ContentListDetailsResponseResultAction string

const (
	ContentListDetailsResponseResultActionBlock ContentListDetailsResponseResultAction = "block"
)

// Whether the API call was successful
type ContentListDetailsResponseSuccess bool

const (
	ContentListDetailsResponseSuccessTrue ContentListDetailsResponseSuccess = true
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
