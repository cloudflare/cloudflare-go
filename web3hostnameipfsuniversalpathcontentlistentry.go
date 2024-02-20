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

// Web3HostnameIpfsUniversalPathContentListEntryService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWeb3HostnameIpfsUniversalPathContentListEntryService] method instead.
type Web3HostnameIpfsUniversalPathContentListEntryService struct {
	Options []option.RequestOption
}

// NewWeb3HostnameIpfsUniversalPathContentListEntryService generates a new service
// that applies the given options to each request. These options are applied after
// the parent client's options (if there is one), and before any request-specific
// options.
func NewWeb3HostnameIpfsUniversalPathContentListEntryService(opts ...option.RequestOption) (r *Web3HostnameIpfsUniversalPathContentListEntryService) {
	r = &Web3HostnameIpfsUniversalPathContentListEntryService{}
	r.Options = opts
	return
}

// Create IPFS Universal Path Gateway Content List Entry
func (r *Web3HostnameIpfsUniversalPathContentListEntryService) New(ctx context.Context, zoneIdentifier string, identifier string, body Web3HostnameIpfsUniversalPathContentListEntryNewParams, opts ...option.RequestOption) (res *Web3HostnameIpfsUniversalPathContentListEntryNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env Web3HostnameIpfsUniversalPathContentListEntryNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list/entries", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Edit IPFS Universal Path Gateway Content List Entry
func (r *Web3HostnameIpfsUniversalPathContentListEntryService) Update(ctx context.Context, zoneIdentifier string, identifier string, contentListEntryIdentifier string, body Web3HostnameIpfsUniversalPathContentListEntryUpdateParams, opts ...option.RequestOption) (res *Web3HostnameIpfsUniversalPathContentListEntryUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env Web3HostnameIpfsUniversalPathContentListEntryUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list/entries/%s", zoneIdentifier, identifier, contentListEntryIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List IPFS Universal Path Gateway Content List Entries
func (r *Web3HostnameIpfsUniversalPathContentListEntryService) List(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *Web3HostnameIpfsUniversalPathContentListEntryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env Web3HostnameIpfsUniversalPathContentListEntryListResponseEnvelope
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list/entries", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete IPFS Universal Path Gateway Content List Entry
func (r *Web3HostnameIpfsUniversalPathContentListEntryService) Delete(ctx context.Context, zoneIdentifier string, identifier string, contentListEntryIdentifier string, opts ...option.RequestOption) (res *Web3HostnameIpfsUniversalPathContentListEntryDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env Web3HostnameIpfsUniversalPathContentListEntryDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list/entries/%s", zoneIdentifier, identifier, contentListEntryIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// IPFS Universal Path Gateway Content List Entry Details
func (r *Web3HostnameIpfsUniversalPathContentListEntryService) Get(ctx context.Context, zoneIdentifier string, identifier string, contentListEntryIdentifier string, opts ...option.RequestOption) (res *Web3HostnameIpfsUniversalPathContentListEntryGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env Web3HostnameIpfsUniversalPathContentListEntryGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list/entries/%s", zoneIdentifier, identifier, contentListEntryIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Content list entry to be blocked.
type Web3HostnameIpfsUniversalPathContentListEntryNewResponse struct {
	// Identifier
	ID string `json:"id"`
	// CID or content path of content to block.
	Content   string    `json:"content"`
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the content list entry.
	Description string    `json:"description"`
	ModifiedOn  time.Time `json:"modified_on" format:"date-time"`
	// Type of content list entry to block.
	Type Web3HostnameIpfsUniversalPathContentListEntryNewResponseType `json:"type"`
	JSON web3HostnameIpfsUniversalPathContentListEntryNewResponseJSON `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListEntryNewResponseJSON contains the JSON
// metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListEntryNewResponse]
type web3HostnameIpfsUniversalPathContentListEntryNewResponseJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	ModifiedOn  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListEntryNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of content list entry to block.
type Web3HostnameIpfsUniversalPathContentListEntryNewResponseType string

const (
	Web3HostnameIpfsUniversalPathContentListEntryNewResponseTypeCid         Web3HostnameIpfsUniversalPathContentListEntryNewResponseType = "cid"
	Web3HostnameIpfsUniversalPathContentListEntryNewResponseTypeContentPath Web3HostnameIpfsUniversalPathContentListEntryNewResponseType = "content_path"
)

// Content list entry to be blocked.
type Web3HostnameIpfsUniversalPathContentListEntryUpdateResponse struct {
	// Identifier
	ID string `json:"id"`
	// CID or content path of content to block.
	Content   string    `json:"content"`
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the content list entry.
	Description string    `json:"description"`
	ModifiedOn  time.Time `json:"modified_on" format:"date-time"`
	// Type of content list entry to block.
	Type Web3HostnameIpfsUniversalPathContentListEntryUpdateResponseType `json:"type"`
	JSON web3HostnameIpfsUniversalPathContentListEntryUpdateResponseJSON `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListEntryUpdateResponseJSON contains the
// JSON metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListEntryUpdateResponse]
type web3HostnameIpfsUniversalPathContentListEntryUpdateResponseJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	ModifiedOn  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListEntryUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of content list entry to block.
type Web3HostnameIpfsUniversalPathContentListEntryUpdateResponseType string

const (
	Web3HostnameIpfsUniversalPathContentListEntryUpdateResponseTypeCid         Web3HostnameIpfsUniversalPathContentListEntryUpdateResponseType = "cid"
	Web3HostnameIpfsUniversalPathContentListEntryUpdateResponseTypeContentPath Web3HostnameIpfsUniversalPathContentListEntryUpdateResponseType = "content_path"
)

type Web3HostnameIpfsUniversalPathContentListEntryListResponse struct {
	// Content list entries.
	Entries []Web3HostnameIpfsUniversalPathContentListEntryListResponseEntry `json:"entries"`
	JSON    web3HostnameIpfsUniversalPathContentListEntryListResponseJSON    `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListEntryListResponseJSON contains the JSON
// metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListEntryListResponse]
type web3HostnameIpfsUniversalPathContentListEntryListResponseJSON struct {
	Entries     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListEntryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Content list entry to be blocked.
type Web3HostnameIpfsUniversalPathContentListEntryListResponseEntry struct {
	// Identifier
	ID string `json:"id"`
	// CID or content path of content to block.
	Content   string    `json:"content"`
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the content list entry.
	Description string    `json:"description"`
	ModifiedOn  time.Time `json:"modified_on" format:"date-time"`
	// Type of content list entry to block.
	Type Web3HostnameIpfsUniversalPathContentListEntryListResponseEntriesType `json:"type"`
	JSON web3HostnameIpfsUniversalPathContentListEntryListResponseEntryJSON   `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListEntryListResponseEntryJSON contains the
// JSON metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListEntryListResponseEntry]
type web3HostnameIpfsUniversalPathContentListEntryListResponseEntryJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	ModifiedOn  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListEntryListResponseEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of content list entry to block.
type Web3HostnameIpfsUniversalPathContentListEntryListResponseEntriesType string

const (
	Web3HostnameIpfsUniversalPathContentListEntryListResponseEntriesTypeCid         Web3HostnameIpfsUniversalPathContentListEntryListResponseEntriesType = "cid"
	Web3HostnameIpfsUniversalPathContentListEntryListResponseEntriesTypeContentPath Web3HostnameIpfsUniversalPathContentListEntryListResponseEntriesType = "content_path"
)

type Web3HostnameIpfsUniversalPathContentListEntryDeleteResponse struct {
	// Identifier
	ID   string                                                          `json:"id,required"`
	JSON web3HostnameIpfsUniversalPathContentListEntryDeleteResponseJSON `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListEntryDeleteResponseJSON contains the
// JSON metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListEntryDeleteResponse]
type web3HostnameIpfsUniversalPathContentListEntryDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListEntryDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Content list entry to be blocked.
type Web3HostnameIpfsUniversalPathContentListEntryGetResponse struct {
	// Identifier
	ID string `json:"id"`
	// CID or content path of content to block.
	Content   string    `json:"content"`
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the content list entry.
	Description string    `json:"description"`
	ModifiedOn  time.Time `json:"modified_on" format:"date-time"`
	// Type of content list entry to block.
	Type Web3HostnameIpfsUniversalPathContentListEntryGetResponseType `json:"type"`
	JSON web3HostnameIpfsUniversalPathContentListEntryGetResponseJSON `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListEntryGetResponseJSON contains the JSON
// metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListEntryGetResponse]
type web3HostnameIpfsUniversalPathContentListEntryGetResponseJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	ModifiedOn  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListEntryGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of content list entry to block.
type Web3HostnameIpfsUniversalPathContentListEntryGetResponseType string

const (
	Web3HostnameIpfsUniversalPathContentListEntryGetResponseTypeCid         Web3HostnameIpfsUniversalPathContentListEntryGetResponseType = "cid"
	Web3HostnameIpfsUniversalPathContentListEntryGetResponseTypeContentPath Web3HostnameIpfsUniversalPathContentListEntryGetResponseType = "content_path"
)

type Web3HostnameIpfsUniversalPathContentListEntryNewParams struct {
	// CID or content path of content to block.
	Content param.Field[string] `json:"content,required"`
	// Type of content list entry to block.
	Type param.Field[Web3HostnameIpfsUniversalPathContentListEntryNewParamsType] `json:"type,required"`
	// An optional description of the content list entry.
	Description param.Field[string] `json:"description"`
}

func (r Web3HostnameIpfsUniversalPathContentListEntryNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of content list entry to block.
type Web3HostnameIpfsUniversalPathContentListEntryNewParamsType string

const (
	Web3HostnameIpfsUniversalPathContentListEntryNewParamsTypeCid         Web3HostnameIpfsUniversalPathContentListEntryNewParamsType = "cid"
	Web3HostnameIpfsUniversalPathContentListEntryNewParamsTypeContentPath Web3HostnameIpfsUniversalPathContentListEntryNewParamsType = "content_path"
)

type Web3HostnameIpfsUniversalPathContentListEntryNewResponseEnvelope struct {
	Errors   []Web3HostnameIpfsUniversalPathContentListEntryNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []Web3HostnameIpfsUniversalPathContentListEntryNewResponseEnvelopeMessages `json:"messages,required"`
	// Content list entry to be blocked.
	Result Web3HostnameIpfsUniversalPathContentListEntryNewResponse `json:"result,required"`
	// Whether the API call was successful
	Success Web3HostnameIpfsUniversalPathContentListEntryNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    web3HostnameIpfsUniversalPathContentListEntryNewResponseEnvelopeJSON    `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListEntryNewResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListEntryNewResponseEnvelope]
type web3HostnameIpfsUniversalPathContentListEntryNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListEntryNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameIpfsUniversalPathContentListEntryNewResponseEnvelopeErrors struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    web3HostnameIpfsUniversalPathContentListEntryNewResponseEnvelopeErrorsJSON `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListEntryNewResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListEntryNewResponseEnvelopeErrors]
type web3HostnameIpfsUniversalPathContentListEntryNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListEntryNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameIpfsUniversalPathContentListEntryNewResponseEnvelopeMessages struct {
	Code    int64                                                                        `json:"code,required"`
	Message string                                                                       `json:"message,required"`
	JSON    web3HostnameIpfsUniversalPathContentListEntryNewResponseEnvelopeMessagesJSON `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListEntryNewResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListEntryNewResponseEnvelopeMessages]
type web3HostnameIpfsUniversalPathContentListEntryNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListEntryNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type Web3HostnameIpfsUniversalPathContentListEntryNewResponseEnvelopeSuccess bool

const (
	Web3HostnameIpfsUniversalPathContentListEntryNewResponseEnvelopeSuccessTrue Web3HostnameIpfsUniversalPathContentListEntryNewResponseEnvelopeSuccess = true
)

type Web3HostnameIpfsUniversalPathContentListEntryUpdateParams struct {
	// CID or content path of content to block.
	Content param.Field[string] `json:"content,required"`
	// Type of content list entry to block.
	Type param.Field[Web3HostnameIpfsUniversalPathContentListEntryUpdateParamsType] `json:"type,required"`
	// An optional description of the content list entry.
	Description param.Field[string] `json:"description"`
}

func (r Web3HostnameIpfsUniversalPathContentListEntryUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of content list entry to block.
type Web3HostnameIpfsUniversalPathContentListEntryUpdateParamsType string

const (
	Web3HostnameIpfsUniversalPathContentListEntryUpdateParamsTypeCid         Web3HostnameIpfsUniversalPathContentListEntryUpdateParamsType = "cid"
	Web3HostnameIpfsUniversalPathContentListEntryUpdateParamsTypeContentPath Web3HostnameIpfsUniversalPathContentListEntryUpdateParamsType = "content_path"
)

type Web3HostnameIpfsUniversalPathContentListEntryUpdateResponseEnvelope struct {
	Errors   []Web3HostnameIpfsUniversalPathContentListEntryUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []Web3HostnameIpfsUniversalPathContentListEntryUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Content list entry to be blocked.
	Result Web3HostnameIpfsUniversalPathContentListEntryUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success Web3HostnameIpfsUniversalPathContentListEntryUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    web3HostnameIpfsUniversalPathContentListEntryUpdateResponseEnvelopeJSON    `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListEntryUpdateResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListEntryUpdateResponseEnvelope]
type web3HostnameIpfsUniversalPathContentListEntryUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListEntryUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameIpfsUniversalPathContentListEntryUpdateResponseEnvelopeErrors struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    web3HostnameIpfsUniversalPathContentListEntryUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListEntryUpdateResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListEntryUpdateResponseEnvelopeErrors]
type web3HostnameIpfsUniversalPathContentListEntryUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListEntryUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameIpfsUniversalPathContentListEntryUpdateResponseEnvelopeMessages struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    web3HostnameIpfsUniversalPathContentListEntryUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListEntryUpdateResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListEntryUpdateResponseEnvelopeMessages]
type web3HostnameIpfsUniversalPathContentListEntryUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListEntryUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type Web3HostnameIpfsUniversalPathContentListEntryUpdateResponseEnvelopeSuccess bool

const (
	Web3HostnameIpfsUniversalPathContentListEntryUpdateResponseEnvelopeSuccessTrue Web3HostnameIpfsUniversalPathContentListEntryUpdateResponseEnvelopeSuccess = true
)

type Web3HostnameIpfsUniversalPathContentListEntryListResponseEnvelope struct {
	Errors   []Web3HostnameIpfsUniversalPathContentListEntryListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []Web3HostnameIpfsUniversalPathContentListEntryListResponseEnvelopeMessages `json:"messages,required"`
	Result   Web3HostnameIpfsUniversalPathContentListEntryListResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    Web3HostnameIpfsUniversalPathContentListEntryListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo Web3HostnameIpfsUniversalPathContentListEntryListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       web3HostnameIpfsUniversalPathContentListEntryListResponseEnvelopeJSON       `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListEntryListResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListEntryListResponseEnvelope]
type web3HostnameIpfsUniversalPathContentListEntryListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListEntryListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameIpfsUniversalPathContentListEntryListResponseEnvelopeErrors struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    web3HostnameIpfsUniversalPathContentListEntryListResponseEnvelopeErrorsJSON `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListEntryListResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListEntryListResponseEnvelopeErrors]
type web3HostnameIpfsUniversalPathContentListEntryListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListEntryListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameIpfsUniversalPathContentListEntryListResponseEnvelopeMessages struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    web3HostnameIpfsUniversalPathContentListEntryListResponseEnvelopeMessagesJSON `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListEntryListResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListEntryListResponseEnvelopeMessages]
type web3HostnameIpfsUniversalPathContentListEntryListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListEntryListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type Web3HostnameIpfsUniversalPathContentListEntryListResponseEnvelopeSuccess bool

const (
	Web3HostnameIpfsUniversalPathContentListEntryListResponseEnvelopeSuccessTrue Web3HostnameIpfsUniversalPathContentListEntryListResponseEnvelopeSuccess = true
)

type Web3HostnameIpfsUniversalPathContentListEntryListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                         `json:"total_count"`
	JSON       web3HostnameIpfsUniversalPathContentListEntryListResponseEnvelopeResultInfoJSON `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListEntryListResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListEntryListResponseEnvelopeResultInfo]
type web3HostnameIpfsUniversalPathContentListEntryListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListEntryListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameIpfsUniversalPathContentListEntryDeleteResponseEnvelope struct {
	Errors   []Web3HostnameIpfsUniversalPathContentListEntryDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []Web3HostnameIpfsUniversalPathContentListEntryDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   Web3HostnameIpfsUniversalPathContentListEntryDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success Web3HostnameIpfsUniversalPathContentListEntryDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    web3HostnameIpfsUniversalPathContentListEntryDeleteResponseEnvelopeJSON    `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListEntryDeleteResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListEntryDeleteResponseEnvelope]
type web3HostnameIpfsUniversalPathContentListEntryDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListEntryDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameIpfsUniversalPathContentListEntryDeleteResponseEnvelopeErrors struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    web3HostnameIpfsUniversalPathContentListEntryDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListEntryDeleteResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListEntryDeleteResponseEnvelopeErrors]
type web3HostnameIpfsUniversalPathContentListEntryDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListEntryDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameIpfsUniversalPathContentListEntryDeleteResponseEnvelopeMessages struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    web3HostnameIpfsUniversalPathContentListEntryDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListEntryDeleteResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListEntryDeleteResponseEnvelopeMessages]
type web3HostnameIpfsUniversalPathContentListEntryDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListEntryDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type Web3HostnameIpfsUniversalPathContentListEntryDeleteResponseEnvelopeSuccess bool

const (
	Web3HostnameIpfsUniversalPathContentListEntryDeleteResponseEnvelopeSuccessTrue Web3HostnameIpfsUniversalPathContentListEntryDeleteResponseEnvelopeSuccess = true
)

type Web3HostnameIpfsUniversalPathContentListEntryGetResponseEnvelope struct {
	Errors   []Web3HostnameIpfsUniversalPathContentListEntryGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []Web3HostnameIpfsUniversalPathContentListEntryGetResponseEnvelopeMessages `json:"messages,required"`
	// Content list entry to be blocked.
	Result Web3HostnameIpfsUniversalPathContentListEntryGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success Web3HostnameIpfsUniversalPathContentListEntryGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    web3HostnameIpfsUniversalPathContentListEntryGetResponseEnvelopeJSON    `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListEntryGetResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListEntryGetResponseEnvelope]
type web3HostnameIpfsUniversalPathContentListEntryGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListEntryGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameIpfsUniversalPathContentListEntryGetResponseEnvelopeErrors struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    web3HostnameIpfsUniversalPathContentListEntryGetResponseEnvelopeErrorsJSON `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListEntryGetResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListEntryGetResponseEnvelopeErrors]
type web3HostnameIpfsUniversalPathContentListEntryGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListEntryGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameIpfsUniversalPathContentListEntryGetResponseEnvelopeMessages struct {
	Code    int64                                                                        `json:"code,required"`
	Message string                                                                       `json:"message,required"`
	JSON    web3HostnameIpfsUniversalPathContentListEntryGetResponseEnvelopeMessagesJSON `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListEntryGetResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListEntryGetResponseEnvelopeMessages]
type web3HostnameIpfsUniversalPathContentListEntryGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListEntryGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type Web3HostnameIpfsUniversalPathContentListEntryGetResponseEnvelopeSuccess bool

const (
	Web3HostnameIpfsUniversalPathContentListEntryGetResponseEnvelopeSuccessTrue Web3HostnameIpfsUniversalPathContentListEntryGetResponseEnvelopeSuccess = true
)
