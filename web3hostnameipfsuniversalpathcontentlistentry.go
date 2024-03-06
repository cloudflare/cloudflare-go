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

// Web3HostnameIPFSUniversalPathContentListEntryService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWeb3HostnameIPFSUniversalPathContentListEntryService] method instead.
type Web3HostnameIPFSUniversalPathContentListEntryService struct {
	Options []option.RequestOption
}

// NewWeb3HostnameIPFSUniversalPathContentListEntryService generates a new service
// that applies the given options to each request. These options are applied after
// the parent client's options (if there is one), and before any request-specific
// options.
func NewWeb3HostnameIPFSUniversalPathContentListEntryService(opts ...option.RequestOption) (r *Web3HostnameIPFSUniversalPathContentListEntryService) {
	r = &Web3HostnameIPFSUniversalPathContentListEntryService{}
	r.Options = opts
	return
}

// Create IPFS Universal Path Gateway Content List Entry
func (r *Web3HostnameIPFSUniversalPathContentListEntryService) New(ctx context.Context, zoneIdentifier string, identifier string, body Web3HostnameIPFSUniversalPathContentListEntryNewParams, opts ...option.RequestOption) (res *Web3HostnameIPFSUniversalPathContentListEntryNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env Web3HostnameIPFSUniversalPathContentListEntryNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list/entries", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Edit IPFS Universal Path Gateway Content List Entry
func (r *Web3HostnameIPFSUniversalPathContentListEntryService) Update(ctx context.Context, zoneIdentifier string, identifier string, contentListEntryIdentifier string, body Web3HostnameIPFSUniversalPathContentListEntryUpdateParams, opts ...option.RequestOption) (res *Web3HostnameIPFSUniversalPathContentListEntryUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env Web3HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list/entries/%s", zoneIdentifier, identifier, contentListEntryIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List IPFS Universal Path Gateway Content List Entries
func (r *Web3HostnameIPFSUniversalPathContentListEntryService) List(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *Web3HostnameIPFSUniversalPathContentListEntryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env Web3HostnameIPFSUniversalPathContentListEntryListResponseEnvelope
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list/entries", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete IPFS Universal Path Gateway Content List Entry
func (r *Web3HostnameIPFSUniversalPathContentListEntryService) Delete(ctx context.Context, zoneIdentifier string, identifier string, contentListEntryIdentifier string, opts ...option.RequestOption) (res *Web3HostnameIPFSUniversalPathContentListEntryDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env Web3HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list/entries/%s", zoneIdentifier, identifier, contentListEntryIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// IPFS Universal Path Gateway Content List Entry Details
func (r *Web3HostnameIPFSUniversalPathContentListEntryService) Get(ctx context.Context, zoneIdentifier string, identifier string, contentListEntryIdentifier string, opts ...option.RequestOption) (res *Web3HostnameIPFSUniversalPathContentListEntryGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env Web3HostnameIPFSUniversalPathContentListEntryGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list/entries/%s", zoneIdentifier, identifier, contentListEntryIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Content list entry to be blocked.
type Web3HostnameIPFSUniversalPathContentListEntryNewResponse struct {
	// Identifier
	ID string `json:"id"`
	// CID or content path of content to block.
	Content   string    `json:"content"`
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the content list entry.
	Description string    `json:"description"`
	ModifiedOn  time.Time `json:"modified_on" format:"date-time"`
	// Type of content list entry to block.
	Type Web3HostnameIPFSUniversalPathContentListEntryNewResponseType `json:"type"`
	JSON web3HostnameIPFSUniversalPathContentListEntryNewResponseJSON `json:"-"`
}

// web3HostnameIPFSUniversalPathContentListEntryNewResponseJSON contains the JSON
// metadata for the struct
// [Web3HostnameIPFSUniversalPathContentListEntryNewResponse]
type web3HostnameIPFSUniversalPathContentListEntryNewResponseJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	ModifiedOn  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIPFSUniversalPathContentListEntryNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3HostnameIPFSUniversalPathContentListEntryNewResponseJSON) RawJSON() string {
	return r.raw
}

// Type of content list entry to block.
type Web3HostnameIPFSUniversalPathContentListEntryNewResponseType string

const (
	Web3HostnameIPFSUniversalPathContentListEntryNewResponseTypeCid         Web3HostnameIPFSUniversalPathContentListEntryNewResponseType = "cid"
	Web3HostnameIPFSUniversalPathContentListEntryNewResponseTypeContentPath Web3HostnameIPFSUniversalPathContentListEntryNewResponseType = "content_path"
)

// Content list entry to be blocked.
type Web3HostnameIPFSUniversalPathContentListEntryUpdateResponse struct {
	// Identifier
	ID string `json:"id"`
	// CID or content path of content to block.
	Content   string    `json:"content"`
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the content list entry.
	Description string    `json:"description"`
	ModifiedOn  time.Time `json:"modified_on" format:"date-time"`
	// Type of content list entry to block.
	Type Web3HostnameIPFSUniversalPathContentListEntryUpdateResponseType `json:"type"`
	JSON web3HostnameIPFSUniversalPathContentListEntryUpdateResponseJSON `json:"-"`
}

// web3HostnameIPFSUniversalPathContentListEntryUpdateResponseJSON contains the
// JSON metadata for the struct
// [Web3HostnameIPFSUniversalPathContentListEntryUpdateResponse]
type web3HostnameIPFSUniversalPathContentListEntryUpdateResponseJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	ModifiedOn  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIPFSUniversalPathContentListEntryUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3HostnameIPFSUniversalPathContentListEntryUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Type of content list entry to block.
type Web3HostnameIPFSUniversalPathContentListEntryUpdateResponseType string

const (
	Web3HostnameIPFSUniversalPathContentListEntryUpdateResponseTypeCid         Web3HostnameIPFSUniversalPathContentListEntryUpdateResponseType = "cid"
	Web3HostnameIPFSUniversalPathContentListEntryUpdateResponseTypeContentPath Web3HostnameIPFSUniversalPathContentListEntryUpdateResponseType = "content_path"
)

type Web3HostnameIPFSUniversalPathContentListEntryListResponse struct {
	// Content list entries.
	Entries []Web3HostnameIPFSUniversalPathContentListEntryListResponseEntry `json:"entries"`
	JSON    web3HostnameIPFSUniversalPathContentListEntryListResponseJSON    `json:"-"`
}

// web3HostnameIPFSUniversalPathContentListEntryListResponseJSON contains the JSON
// metadata for the struct
// [Web3HostnameIPFSUniversalPathContentListEntryListResponse]
type web3HostnameIPFSUniversalPathContentListEntryListResponseJSON struct {
	Entries     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIPFSUniversalPathContentListEntryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3HostnameIPFSUniversalPathContentListEntryListResponseJSON) RawJSON() string {
	return r.raw
}

// Content list entry to be blocked.
type Web3HostnameIPFSUniversalPathContentListEntryListResponseEntry struct {
	// Identifier
	ID string `json:"id"`
	// CID or content path of content to block.
	Content   string    `json:"content"`
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the content list entry.
	Description string    `json:"description"`
	ModifiedOn  time.Time `json:"modified_on" format:"date-time"`
	// Type of content list entry to block.
	Type Web3HostnameIPFSUniversalPathContentListEntryListResponseEntriesType `json:"type"`
	JSON web3HostnameIPFSUniversalPathContentListEntryListResponseEntryJSON   `json:"-"`
}

// web3HostnameIPFSUniversalPathContentListEntryListResponseEntryJSON contains the
// JSON metadata for the struct
// [Web3HostnameIPFSUniversalPathContentListEntryListResponseEntry]
type web3HostnameIPFSUniversalPathContentListEntryListResponseEntryJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	ModifiedOn  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIPFSUniversalPathContentListEntryListResponseEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3HostnameIPFSUniversalPathContentListEntryListResponseEntryJSON) RawJSON() string {
	return r.raw
}

// Type of content list entry to block.
type Web3HostnameIPFSUniversalPathContentListEntryListResponseEntriesType string

const (
	Web3HostnameIPFSUniversalPathContentListEntryListResponseEntriesTypeCid         Web3HostnameIPFSUniversalPathContentListEntryListResponseEntriesType = "cid"
	Web3HostnameIPFSUniversalPathContentListEntryListResponseEntriesTypeContentPath Web3HostnameIPFSUniversalPathContentListEntryListResponseEntriesType = "content_path"
)

type Web3HostnameIPFSUniversalPathContentListEntryDeleteResponse struct {
	// Identifier
	ID   string                                                          `json:"id,required"`
	JSON web3HostnameIPFSUniversalPathContentListEntryDeleteResponseJSON `json:"-"`
}

// web3HostnameIPFSUniversalPathContentListEntryDeleteResponseJSON contains the
// JSON metadata for the struct
// [Web3HostnameIPFSUniversalPathContentListEntryDeleteResponse]
type web3HostnameIPFSUniversalPathContentListEntryDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIPFSUniversalPathContentListEntryDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3HostnameIPFSUniversalPathContentListEntryDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Content list entry to be blocked.
type Web3HostnameIPFSUniversalPathContentListEntryGetResponse struct {
	// Identifier
	ID string `json:"id"`
	// CID or content path of content to block.
	Content   string    `json:"content"`
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the content list entry.
	Description string    `json:"description"`
	ModifiedOn  time.Time `json:"modified_on" format:"date-time"`
	// Type of content list entry to block.
	Type Web3HostnameIPFSUniversalPathContentListEntryGetResponseType `json:"type"`
	JSON web3HostnameIPFSUniversalPathContentListEntryGetResponseJSON `json:"-"`
}

// web3HostnameIPFSUniversalPathContentListEntryGetResponseJSON contains the JSON
// metadata for the struct
// [Web3HostnameIPFSUniversalPathContentListEntryGetResponse]
type web3HostnameIPFSUniversalPathContentListEntryGetResponseJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	ModifiedOn  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIPFSUniversalPathContentListEntryGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3HostnameIPFSUniversalPathContentListEntryGetResponseJSON) RawJSON() string {
	return r.raw
}

// Type of content list entry to block.
type Web3HostnameIPFSUniversalPathContentListEntryGetResponseType string

const (
	Web3HostnameIPFSUniversalPathContentListEntryGetResponseTypeCid         Web3HostnameIPFSUniversalPathContentListEntryGetResponseType = "cid"
	Web3HostnameIPFSUniversalPathContentListEntryGetResponseTypeContentPath Web3HostnameIPFSUniversalPathContentListEntryGetResponseType = "content_path"
)

type Web3HostnameIPFSUniversalPathContentListEntryNewParams struct {
	// CID or content path of content to block.
	Content param.Field[string] `json:"content,required"`
	// Type of content list entry to block.
	Type param.Field[Web3HostnameIPFSUniversalPathContentListEntryNewParamsType] `json:"type,required"`
	// An optional description of the content list entry.
	Description param.Field[string] `json:"description"`
}

func (r Web3HostnameIPFSUniversalPathContentListEntryNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of content list entry to block.
type Web3HostnameIPFSUniversalPathContentListEntryNewParamsType string

const (
	Web3HostnameIPFSUniversalPathContentListEntryNewParamsTypeCid         Web3HostnameIPFSUniversalPathContentListEntryNewParamsType = "cid"
	Web3HostnameIPFSUniversalPathContentListEntryNewParamsTypeContentPath Web3HostnameIPFSUniversalPathContentListEntryNewParamsType = "content_path"
)

type Web3HostnameIPFSUniversalPathContentListEntryNewResponseEnvelope struct {
	Errors   []Web3HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []Web3HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeMessages `json:"messages,required"`
	// Content list entry to be blocked.
	Result Web3HostnameIPFSUniversalPathContentListEntryNewResponse `json:"result,required"`
	// Whether the API call was successful
	Success Web3HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    web3HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeJSON    `json:"-"`
}

// web3HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [Web3HostnameIPFSUniversalPathContentListEntryNewResponseEnvelope]
type web3HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIPFSUniversalPathContentListEntryNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type Web3HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeErrors struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    web3HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeErrorsJSON `json:"-"`
}

// web3HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [Web3HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeErrors]
type web3HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type Web3HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeMessages struct {
	Code    int64                                                                        `json:"code,required"`
	Message string                                                                       `json:"message,required"`
	JSON    web3HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeMessagesJSON `json:"-"`
}

// web3HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [Web3HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeMessages]
type web3HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type Web3HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeSuccess bool

const (
	Web3HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeSuccessTrue Web3HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeSuccess = true
)

type Web3HostnameIPFSUniversalPathContentListEntryUpdateParams struct {
	// CID or content path of content to block.
	Content param.Field[string] `json:"content,required"`
	// Type of content list entry to block.
	Type param.Field[Web3HostnameIPFSUniversalPathContentListEntryUpdateParamsType] `json:"type,required"`
	// An optional description of the content list entry.
	Description param.Field[string] `json:"description"`
}

func (r Web3HostnameIPFSUniversalPathContentListEntryUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of content list entry to block.
type Web3HostnameIPFSUniversalPathContentListEntryUpdateParamsType string

const (
	Web3HostnameIPFSUniversalPathContentListEntryUpdateParamsTypeCid         Web3HostnameIPFSUniversalPathContentListEntryUpdateParamsType = "cid"
	Web3HostnameIPFSUniversalPathContentListEntryUpdateParamsTypeContentPath Web3HostnameIPFSUniversalPathContentListEntryUpdateParamsType = "content_path"
)

type Web3HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelope struct {
	Errors   []Web3HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []Web3HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Content list entry to be blocked.
	Result Web3HostnameIPFSUniversalPathContentListEntryUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success Web3HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    web3HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeJSON    `json:"-"`
}

// web3HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [Web3HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelope]
type web3HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type Web3HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeErrors struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    web3HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// web3HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [Web3HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeErrors]
type web3HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type Web3HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeMessages struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    web3HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// web3HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [Web3HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeMessages]
type web3HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type Web3HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeSuccess bool

const (
	Web3HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeSuccessTrue Web3HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeSuccess = true
)

type Web3HostnameIPFSUniversalPathContentListEntryListResponseEnvelope struct {
	Errors   []Web3HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []Web3HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeMessages `json:"messages,required"`
	Result   Web3HostnameIPFSUniversalPathContentListEntryListResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    Web3HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo Web3HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       web3HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeJSON       `json:"-"`
}

// web3HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [Web3HostnameIPFSUniversalPathContentListEntryListResponseEnvelope]
type web3HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIPFSUniversalPathContentListEntryListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type Web3HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeErrors struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    web3HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeErrorsJSON `json:"-"`
}

// web3HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [Web3HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeErrors]
type web3HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type Web3HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeMessages struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    web3HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeMessagesJSON `json:"-"`
}

// web3HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [Web3HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeMessages]
type web3HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type Web3HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeSuccess bool

const (
	Web3HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeSuccessTrue Web3HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeSuccess = true
)

type Web3HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                         `json:"total_count"`
	JSON       web3HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeResultInfoJSON `json:"-"`
}

// web3HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [Web3HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeResultInfo]
type web3HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type Web3HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelope struct {
	Errors   []Web3HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []Web3HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   Web3HostnameIPFSUniversalPathContentListEntryDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success Web3HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    web3HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeJSON    `json:"-"`
}

// web3HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [Web3HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelope]
type web3HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type Web3HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeErrors struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    web3HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// web3HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [Web3HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeErrors]
type web3HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type Web3HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeMessages struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    web3HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// web3HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [Web3HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeMessages]
type web3HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type Web3HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeSuccess bool

const (
	Web3HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeSuccessTrue Web3HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeSuccess = true
)

type Web3HostnameIPFSUniversalPathContentListEntryGetResponseEnvelope struct {
	Errors   []Web3HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []Web3HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeMessages `json:"messages,required"`
	// Content list entry to be blocked.
	Result Web3HostnameIPFSUniversalPathContentListEntryGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success Web3HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    web3HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeJSON    `json:"-"`
}

// web3HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [Web3HostnameIPFSUniversalPathContentListEntryGetResponseEnvelope]
type web3HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIPFSUniversalPathContentListEntryGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type Web3HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeErrors struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    web3HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeErrorsJSON `json:"-"`
}

// web3HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [Web3HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeErrors]
type web3HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type Web3HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeMessages struct {
	Code    int64                                                                        `json:"code,required"`
	Message string                                                                       `json:"message,required"`
	JSON    web3HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeMessagesJSON `json:"-"`
}

// web3HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [Web3HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeMessages]
type web3HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type Web3HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeSuccess bool

const (
	Web3HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeSuccessTrue Web3HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeSuccess = true
)
