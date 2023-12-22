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
func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryService) Get(ctx context.Context, zoneIdentifier string, identifier string, contentListEntryIdentifier string, opts ...option.RequestOption) (res *ContentListEntrySingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list/entries/%s", zoneIdentifier, identifier, contentListEntryIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Edit IPFS Universal Path Gateway Content List Entry
func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryService) Update(ctx context.Context, zoneIdentifier string, identifier string, contentListEntryIdentifier string, body ZoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateParams, opts ...option.RequestOption) (res *ContentListEntrySingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list/entries/%s", zoneIdentifier, identifier, contentListEntryIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete IPFS Universal Path Gateway Content List Entry
func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryService) Delete(ctx context.Context, zoneIdentifier string, identifier string, contentListEntryIdentifier string, opts ...option.RequestOption) (res *APIResponseSingleIDKYar7dC1, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list/entries/%s", zoneIdentifier, identifier, contentListEntryIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create IPFS Universal Path Gateway Content List Entry
func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryService) Web3HostnameNewIpfsUniversalPathGatewayContentListEntry(ctx context.Context, zoneIdentifier string, identifier string, body ZoneWeb3HostnameIpfsUniversalPathContentListEntryWeb3HostnameNewIpfsUniversalPathGatewayContentListEntryParams, opts ...option.RequestOption) (res *ContentListEntrySingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list/entries", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List IPFS Universal Path Gateway Content List Entries
func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryService) Web3HostnameListIpfsUniversalPathGatewayContentListEntries(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *ContentListEntryCollectionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list/entries", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ContentListEntryCollectionResponse struct {
	Errors     []ContentListEntryCollectionResponseError    `json:"errors"`
	Messages   []ContentListEntryCollectionResponseMessage  `json:"messages"`
	Result     ContentListEntryCollectionResponseResult     `json:"result"`
	ResultInfo ContentListEntryCollectionResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ContentListEntryCollectionResponseSuccess `json:"success"`
	JSON    contentListEntryCollectionResponseJSON    `json:"-"`
}

// contentListEntryCollectionResponseJSON contains the JSON metadata for the struct
// [ContentListEntryCollectionResponse]
type contentListEntryCollectionResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentListEntryCollectionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContentListEntryCollectionResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    contentListEntryCollectionResponseErrorJSON `json:"-"`
}

// contentListEntryCollectionResponseErrorJSON contains the JSON metadata for the
// struct [ContentListEntryCollectionResponseError]
type contentListEntryCollectionResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentListEntryCollectionResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContentListEntryCollectionResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    contentListEntryCollectionResponseMessageJSON `json:"-"`
}

// contentListEntryCollectionResponseMessageJSON contains the JSON metadata for the
// struct [ContentListEntryCollectionResponseMessage]
type contentListEntryCollectionResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentListEntryCollectionResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContentListEntryCollectionResponseResult struct {
	// Content list entries.
	Entries []ContentListEntryCollectionResponseResultEntry `json:"entries"`
	JSON    contentListEntryCollectionResponseResultJSON    `json:"-"`
}

// contentListEntryCollectionResponseResultJSON contains the JSON metadata for the
// struct [ContentListEntryCollectionResponseResult]
type contentListEntryCollectionResponseResultJSON struct {
	Entries     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentListEntryCollectionResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Content list entry to be blocked.
type ContentListEntryCollectionResponseResultEntry struct {
	// Identifier
	ID string `json:"id"`
	// CID or content path of content to block.
	Content   string    `json:"content"`
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the content list entry.
	Description string    `json:"description"`
	ModifiedOn  time.Time `json:"modified_on" format:"date-time"`
	// Type of content list entry to block.
	Type ContentListEntryCollectionResponseResultEntriesType `json:"type"`
	JSON contentListEntryCollectionResponseResultEntryJSON   `json:"-"`
}

// contentListEntryCollectionResponseResultEntryJSON contains the JSON metadata for
// the struct [ContentListEntryCollectionResponseResultEntry]
type contentListEntryCollectionResponseResultEntryJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	ModifiedOn  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentListEntryCollectionResponseResultEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of content list entry to block.
type ContentListEntryCollectionResponseResultEntriesType string

const (
	ContentListEntryCollectionResponseResultEntriesTypeCid         ContentListEntryCollectionResponseResultEntriesType = "cid"
	ContentListEntryCollectionResponseResultEntriesTypeContentPath ContentListEntryCollectionResponseResultEntriesType = "content_path"
)

type ContentListEntryCollectionResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                          `json:"total_count"`
	JSON       contentListEntryCollectionResponseResultInfoJSON `json:"-"`
}

// contentListEntryCollectionResponseResultInfoJSON contains the JSON metadata for
// the struct [ContentListEntryCollectionResponseResultInfo]
type contentListEntryCollectionResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentListEntryCollectionResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ContentListEntryCollectionResponseSuccess bool

const (
	ContentListEntryCollectionResponseSuccessTrue ContentListEntryCollectionResponseSuccess = true
)

type ContentListEntrySingleResponse struct {
	Errors   []ContentListEntrySingleResponseError   `json:"errors"`
	Messages []ContentListEntrySingleResponseMessage `json:"messages"`
	// Content list entry to be blocked.
	Result ContentListEntrySingleResponseResult `json:"result"`
	// Whether the API call was successful
	Success ContentListEntrySingleResponseSuccess `json:"success"`
	JSON    contentListEntrySingleResponseJSON    `json:"-"`
}

// contentListEntrySingleResponseJSON contains the JSON metadata for the struct
// [ContentListEntrySingleResponse]
type contentListEntrySingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentListEntrySingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContentListEntrySingleResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    contentListEntrySingleResponseErrorJSON `json:"-"`
}

// contentListEntrySingleResponseErrorJSON contains the JSON metadata for the
// struct [ContentListEntrySingleResponseError]
type contentListEntrySingleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentListEntrySingleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ContentListEntrySingleResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    contentListEntrySingleResponseMessageJSON `json:"-"`
}

// contentListEntrySingleResponseMessageJSON contains the JSON metadata for the
// struct [ContentListEntrySingleResponseMessage]
type contentListEntrySingleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentListEntrySingleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Content list entry to be blocked.
type ContentListEntrySingleResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// CID or content path of content to block.
	Content   string    `json:"content"`
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the content list entry.
	Description string    `json:"description"`
	ModifiedOn  time.Time `json:"modified_on" format:"date-time"`
	// Type of content list entry to block.
	Type ContentListEntrySingleResponseResultType `json:"type"`
	JSON contentListEntrySingleResponseResultJSON `json:"-"`
}

// contentListEntrySingleResponseResultJSON contains the JSON metadata for the
// struct [ContentListEntrySingleResponseResult]
type contentListEntrySingleResponseResultJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	ModifiedOn  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentListEntrySingleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of content list entry to block.
type ContentListEntrySingleResponseResultType string

const (
	ContentListEntrySingleResponseResultTypeCid         ContentListEntrySingleResponseResultType = "cid"
	ContentListEntrySingleResponseResultTypeContentPath ContentListEntrySingleResponseResultType = "content_path"
)

// Whether the API call was successful
type ContentListEntrySingleResponseSuccess bool

const (
	ContentListEntrySingleResponseSuccessTrue ContentListEntrySingleResponseSuccess = true
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
