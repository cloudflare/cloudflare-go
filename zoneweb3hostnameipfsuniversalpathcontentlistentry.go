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

// ZoneWeb3HostnameIpfsUniversalPathContentListEntryService contains methods and
// other services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneWeb3HostnameIpfsUniversalPathContentListEntryService] method
// instead.
type ZoneWeb3HostnameIpfsUniversalPathContentListEntryService struct {
	Options []option.RequestOption
}

// NewZoneWeb3HostnameIpfsUniversalPathContentListEntryService generates a new
// service that applies the given options to each request. These options are
// applied after the parent client's options (if there is one), and before any
// request-specific options.
func NewZoneWeb3HostnameIpfsUniversalPathContentListEntryService(opts ...option.RequestOption) (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryService) {
	r = &ZoneWeb3HostnameIpfsUniversalPathContentListEntryService{}
	r.Options = opts
	return
}

// IPFS Universal Path Gateway Content List Entry Details
func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryService) Get(ctx context.Context, zoneIdentifier string, identifier string, contentListEntryIdentifier string, opts ...option.RequestOption) (res *ZoneWeb3HostnameIpfsUniversalPathContentListEntryGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list/entries/%s", zoneIdentifier, identifier, contentListEntryIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Edit IPFS Universal Path Gateway Content List Entry
func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryService) Update(ctx context.Context, zoneIdentifier string, identifier string, contentListEntryIdentifier string, body ZoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateParams, opts ...option.RequestOption) (res *ZoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list/entries/%s", zoneIdentifier, identifier, contentListEntryIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete IPFS Universal Path Gateway Content List Entry
func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryService) Delete(ctx context.Context, zoneIdentifier string, identifier string, contentListEntryIdentifier string, opts ...option.RequestOption) (res *ZoneWeb3HostnameIpfsUniversalPathContentListEntryDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list/entries/%s", zoneIdentifier, identifier, contentListEntryIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create IPFS Universal Path Gateway Content List Entry
func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryService) Web3HostnameNewIpfsUniversalPathGatewayContentListEntry(ctx context.Context, zoneIdentifier string, identifier string, body ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryParams, opts ...option.RequestOption) (res *ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list/entries", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List IPFS Universal Path Gateway Content List Entries
func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryService) Web3HostnameListIpfsUniversalPathGatewayContentListEntries(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list/entries", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneWeb3HostnameIpfsUniversalPathContentListEntryGetResponse struct {
	Errors   []ZoneWeb3HostnameIpfsUniversalPathContentListEntryGetResponseError   `json:"errors"`
	Messages []ZoneWeb3HostnameIpfsUniversalPathContentListEntryGetResponseMessage `json:"messages"`
	// Content list entry to be blocked.
	Result ZoneWeb3HostnameIpfsUniversalPathContentListEntryGetResponseResult `json:"result"`
	// Whether the API call was successful
	Success ZoneWeb3HostnameIpfsUniversalPathContentListEntryGetResponseSuccess `json:"success"`
	JSON    zoneWeb3HostnameIpfsUniversalPathContentListEntryGetResponseJSON    `json:"-"`
}

// zoneWeb3HostnameIpfsUniversalPathContentListEntryGetResponseJSON contains the
// JSON metadata for the struct
// [ZoneWeb3HostnameIpfsUniversalPathContentListEntryGetResponse]
type zoneWeb3HostnameIpfsUniversalPathContentListEntryGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWeb3HostnameIpfsUniversalPathContentListEntryGetResponseError struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    zoneWeb3HostnameIpfsUniversalPathContentListEntryGetResponseErrorJSON `json:"-"`
}

// zoneWeb3HostnameIpfsUniversalPathContentListEntryGetResponseErrorJSON contains
// the JSON metadata for the struct
// [ZoneWeb3HostnameIpfsUniversalPathContentListEntryGetResponseError]
type zoneWeb3HostnameIpfsUniversalPathContentListEntryGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWeb3HostnameIpfsUniversalPathContentListEntryGetResponseMessage struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    zoneWeb3HostnameIpfsUniversalPathContentListEntryGetResponseMessageJSON `json:"-"`
}

// zoneWeb3HostnameIpfsUniversalPathContentListEntryGetResponseMessageJSON contains
// the JSON metadata for the struct
// [ZoneWeb3HostnameIpfsUniversalPathContentListEntryGetResponseMessage]
type zoneWeb3HostnameIpfsUniversalPathContentListEntryGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Content list entry to be blocked.
type ZoneWeb3HostnameIpfsUniversalPathContentListEntryGetResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// CID or content path of content to block.
	Content   string    `json:"content"`
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the content list entry.
	Description string    `json:"description"`
	ModifiedOn  time.Time `json:"modified_on" format:"date-time"`
	// Type of content list entry to block.
	Type ZoneWeb3HostnameIpfsUniversalPathContentListEntryGetResponseResultType `json:"type"`
	JSON zoneWeb3HostnameIpfsUniversalPathContentListEntryGetResponseResultJSON `json:"-"`
}

// zoneWeb3HostnameIpfsUniversalPathContentListEntryGetResponseResultJSON contains
// the JSON metadata for the struct
// [ZoneWeb3HostnameIpfsUniversalPathContentListEntryGetResponseResult]
type zoneWeb3HostnameIpfsUniversalPathContentListEntryGetResponseResultJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	ModifiedOn  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of content list entry to block.
type ZoneWeb3HostnameIpfsUniversalPathContentListEntryGetResponseResultType string

const (
	ZoneWeb3HostnameIpfsUniversalPathContentListEntryGetResponseResultTypeCid         ZoneWeb3HostnameIpfsUniversalPathContentListEntryGetResponseResultType = "cid"
	ZoneWeb3HostnameIpfsUniversalPathContentListEntryGetResponseResultTypeContentPath ZoneWeb3HostnameIpfsUniversalPathContentListEntryGetResponseResultType = "content_path"
)

// Whether the API call was successful
type ZoneWeb3HostnameIpfsUniversalPathContentListEntryGetResponseSuccess bool

const (
	ZoneWeb3HostnameIpfsUniversalPathContentListEntryGetResponseSuccessTrue ZoneWeb3HostnameIpfsUniversalPathContentListEntryGetResponseSuccess = true
)

type ZoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateResponse struct {
	Errors   []ZoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateResponseError   `json:"errors"`
	Messages []ZoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateResponseMessage `json:"messages"`
	// Content list entry to be blocked.
	Result ZoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success ZoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateResponseSuccess `json:"success"`
	JSON    zoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateResponseJSON    `json:"-"`
}

// zoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateResponseJSON contains the
// JSON metadata for the struct
// [ZoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateResponse]
type zoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateResponseError struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    zoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateResponseErrorJSON `json:"-"`
}

// zoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateResponseError]
type zoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateResponseMessage struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    zoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateResponseMessageJSON `json:"-"`
}

// zoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateResponseMessage]
type zoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Content list entry to be blocked.
type ZoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// CID or content path of content to block.
	Content   string    `json:"content"`
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the content list entry.
	Description string    `json:"description"`
	ModifiedOn  time.Time `json:"modified_on" format:"date-time"`
	// Type of content list entry to block.
	Type ZoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateResponseResultType `json:"type"`
	JSON zoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateResponseResultJSON `json:"-"`
}

// zoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateResponseResult]
type zoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateResponseResultJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	ModifiedOn  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of content list entry to block.
type ZoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateResponseResultType string

const (
	ZoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateResponseResultTypeCid         ZoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateResponseResultType = "cid"
	ZoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateResponseResultTypeContentPath ZoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateResponseResultType = "content_path"
)

// Whether the API call was successful
type ZoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateResponseSuccess bool

const (
	ZoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateResponseSuccessTrue ZoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateResponseSuccess = true
)

type ZoneWeb3HostnameIpfsUniversalPathContentListEntryDeleteResponse struct {
	Errors   []ZoneWeb3HostnameIpfsUniversalPathContentListEntryDeleteResponseError   `json:"errors"`
	Messages []ZoneWeb3HostnameIpfsUniversalPathContentListEntryDeleteResponseMessage `json:"messages"`
	Result   ZoneWeb3HostnameIpfsUniversalPathContentListEntryDeleteResponseResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success ZoneWeb3HostnameIpfsUniversalPathContentListEntryDeleteResponseSuccess `json:"success"`
	JSON    zoneWeb3HostnameIpfsUniversalPathContentListEntryDeleteResponseJSON    `json:"-"`
}

// zoneWeb3HostnameIpfsUniversalPathContentListEntryDeleteResponseJSON contains the
// JSON metadata for the struct
// [ZoneWeb3HostnameIpfsUniversalPathContentListEntryDeleteResponse]
type zoneWeb3HostnameIpfsUniversalPathContentListEntryDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWeb3HostnameIpfsUniversalPathContentListEntryDeleteResponseError struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    zoneWeb3HostnameIpfsUniversalPathContentListEntryDeleteResponseErrorJSON `json:"-"`
}

// zoneWeb3HostnameIpfsUniversalPathContentListEntryDeleteResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneWeb3HostnameIpfsUniversalPathContentListEntryDeleteResponseError]
type zoneWeb3HostnameIpfsUniversalPathContentListEntryDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWeb3HostnameIpfsUniversalPathContentListEntryDeleteResponseMessage struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    zoneWeb3HostnameIpfsUniversalPathContentListEntryDeleteResponseMessageJSON `json:"-"`
}

// zoneWeb3HostnameIpfsUniversalPathContentListEntryDeleteResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneWeb3HostnameIpfsUniversalPathContentListEntryDeleteResponseMessage]
type zoneWeb3HostnameIpfsUniversalPathContentListEntryDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWeb3HostnameIpfsUniversalPathContentListEntryDeleteResponseResult struct {
	// Identifier
	ID   string                                                                    `json:"id,required"`
	JSON zoneWeb3HostnameIpfsUniversalPathContentListEntryDeleteResponseResultJSON `json:"-"`
}

// zoneWeb3HostnameIpfsUniversalPathContentListEntryDeleteResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneWeb3HostnameIpfsUniversalPathContentListEntryDeleteResponseResult]
type zoneWeb3HostnameIpfsUniversalPathContentListEntryDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneWeb3HostnameIpfsUniversalPathContentListEntryDeleteResponseSuccess bool

const (
	ZoneWeb3HostnameIpfsUniversalPathContentListEntryDeleteResponseSuccessTrue ZoneWeb3HostnameIpfsUniversalPathContentListEntryDeleteResponseSuccess = true
)

type ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryResponse struct {
	Errors   []ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryResponseError   `json:"errors"`
	Messages []ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryResponseMessage `json:"messages"`
	// Content list entry to be blocked.
	Result ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryResponseResult `json:"result"`
	// Whether the API call was successful
	Success ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryResponseSuccess `json:"success"`
	JSON    zoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryResponseJSON    `json:"-"`
}

// zoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryResponseJSON
// contains the JSON metadata for the struct
// [ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryResponse]
type zoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryResponseError struct {
	Code    int64                                                                                                                     `json:"code,required"`
	Message string                                                                                                                    `json:"message,required"`
	JSON    zoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryResponseErrorJSON `json:"-"`
}

// zoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryResponseError]
type zoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryResponseMessage struct {
	Code    int64                                                                                                                       `json:"code,required"`
	Message string                                                                                                                      `json:"message,required"`
	JSON    zoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryResponseMessageJSON `json:"-"`
}

// zoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryResponseMessage]
type zoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Content list entry to be blocked.
type ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// CID or content path of content to block.
	Content   string    `json:"content"`
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the content list entry.
	Description string    `json:"description"`
	ModifiedOn  time.Time `json:"modified_on" format:"date-time"`
	// Type of content list entry to block.
	Type ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryResponseResultType `json:"type"`
	JSON zoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryResponseResultJSON `json:"-"`
}

// zoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryResponseResult]
type zoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryResponseResultJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	ModifiedOn  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of content list entry to block.
type ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryResponseResultType string

const (
	ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryResponseResultTypeCid         ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryResponseResultType = "cid"
	ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryResponseResultTypeContentPath ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryResponseResultType = "content_path"
)

// Whether the API call was successful
type ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryResponseSuccess bool

const (
	ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryResponseSuccessTrue ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryResponseSuccess = true
)

type ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponse struct {
	Errors     []ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseError    `json:"errors"`
	Messages   []ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseMessage  `json:"messages"`
	Result     ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseResult     `json:"result"`
	ResultInfo ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseSuccess `json:"success"`
	JSON    zoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseJSON    `json:"-"`
}

// zoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseJSON
// contains the JSON metadata for the struct
// [ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponse]
type zoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseError struct {
	Code    int64                                                                                                                        `json:"code,required"`
	Message string                                                                                                                       `json:"message,required"`
	JSON    zoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseErrorJSON `json:"-"`
}

// zoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseError]
type zoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseMessage struct {
	Code    int64                                                                                                                          `json:"code,required"`
	Message string                                                                                                                         `json:"message,required"`
	JSON    zoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseMessageJSON `json:"-"`
}

// zoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseMessage]
type zoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseResult struct {
	// Content list entries.
	Entries []ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseResultEntry `json:"entries"`
	JSON    zoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseResultJSON    `json:"-"`
}

// zoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseResult]
type zoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseResultJSON struct {
	Entries     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Content list entry to be blocked.
type ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseResultEntry struct {
	// Identifier
	ID string `json:"id"`
	// CID or content path of content to block.
	Content   string    `json:"content"`
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the content list entry.
	Description string    `json:"description"`
	ModifiedOn  time.Time `json:"modified_on" format:"date-time"`
	// Type of content list entry to block.
	Type ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseResultEntriesType `json:"type"`
	JSON zoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseResultEntryJSON   `json:"-"`
}

// zoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseResultEntryJSON
// contains the JSON metadata for the struct
// [ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseResultEntry]
type zoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseResultEntryJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	ModifiedOn  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseResultEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of content list entry to block.
type ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseResultEntriesType string

const (
	ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseResultEntriesTypeCid         ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseResultEntriesType = "cid"
	ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseResultEntriesTypeContentPath ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseResultEntriesType = "content_path"
)

type ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                                                           `json:"total_count"`
	JSON       zoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseResultInfoJSON `json:"-"`
}

// zoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseResultInfoJSON
// contains the JSON metadata for the struct
// [ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseResultInfo]
type zoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseSuccess bool

const (
	ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseSuccessTrue ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameListIpfsUniversalPathGatewayContentListEntriesResponseSuccess = true
)

type ZoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateParams struct {
	// CID or content path of content to block.
	Content param.Field[string] `json:"content,required"`
	// Type of content list entry to block.
	Type param.Field[ZoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateParamsType] `json:"type,required"`
	// An optional description of the content list entry.
	Description param.Field[string] `json:"description"`
}

func (r ZoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of content list entry to block.
type ZoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateParamsType string

const (
	ZoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateParamsTypeCid         ZoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateParamsType = "cid"
	ZoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateParamsTypeContentPath ZoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateParamsType = "content_path"
)

type ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryParams struct {
	// CID or content path of content to block.
	Content param.Field[string] `json:"content,required"`
	// Type of content list entry to block.
	Type param.Field[ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryParamsType] `json:"type,required"`
	// An optional description of the content list entry.
	Description param.Field[string] `json:"description"`
}

func (r ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of content list entry to block.
type ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryParamsType string

const (
	ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryParamsTypeCid         ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryParamsType = "cid"
	ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryParamsTypeContentPath ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryParamsType = "content_path"
)
