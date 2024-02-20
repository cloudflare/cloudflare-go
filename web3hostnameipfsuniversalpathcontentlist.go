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

// Web3HostnameIpfsUniversalPathContentListService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWeb3HostnameIpfsUniversalPathContentListService] method instead.
type Web3HostnameIpfsUniversalPathContentListService struct {
	Options []option.RequestOption
	Entries *Web3HostnameIpfsUniversalPathContentListEntryService
}

// NewWeb3HostnameIpfsUniversalPathContentListService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewWeb3HostnameIpfsUniversalPathContentListService(opts ...option.RequestOption) (r *Web3HostnameIpfsUniversalPathContentListService) {
	r = &Web3HostnameIpfsUniversalPathContentListService{}
	r.Options = opts
	r.Entries = NewWeb3HostnameIpfsUniversalPathContentListEntryService(opts...)
	return
}

// IPFS Universal Path Gateway Content List Details
func (r *Web3HostnameIpfsUniversalPathContentListService) List(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *Web3HostnameIpfsUniversalPathContentListListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env Web3HostnameIpfsUniversalPathContentListListResponseEnvelope
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update IPFS Universal Path Gateway Content List
func (r *Web3HostnameIpfsUniversalPathContentListService) Replace(ctx context.Context, zoneIdentifier string, identifier string, body Web3HostnameIpfsUniversalPathContentListReplaceParams, opts ...option.RequestOption) (res *Web3HostnameIpfsUniversalPathContentListReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env Web3HostnameIpfsUniversalPathContentListReplaceResponseEnvelope
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Web3HostnameIpfsUniversalPathContentListListResponse struct {
	// Behavior of the content list.
	Action Web3HostnameIpfsUniversalPathContentListListResponseAction `json:"action"`
	JSON   web3HostnameIpfsUniversalPathContentListListResponseJSON   `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListListResponseJSON contains the JSON
// metadata for the struct [Web3HostnameIpfsUniversalPathContentListListResponse]
type web3HostnameIpfsUniversalPathContentListListResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Behavior of the content list.
type Web3HostnameIpfsUniversalPathContentListListResponseAction string

const (
	Web3HostnameIpfsUniversalPathContentListListResponseActionBlock Web3HostnameIpfsUniversalPathContentListListResponseAction = "block"
)

type Web3HostnameIpfsUniversalPathContentListReplaceResponse struct {
	// Behavior of the content list.
	Action Web3HostnameIpfsUniversalPathContentListReplaceResponseAction `json:"action"`
	JSON   web3HostnameIpfsUniversalPathContentListReplaceResponseJSON   `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListReplaceResponseJSON contains the JSON
// metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListReplaceResponse]
type web3HostnameIpfsUniversalPathContentListReplaceResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Behavior of the content list.
type Web3HostnameIpfsUniversalPathContentListReplaceResponseAction string

const (
	Web3HostnameIpfsUniversalPathContentListReplaceResponseActionBlock Web3HostnameIpfsUniversalPathContentListReplaceResponseAction = "block"
)

type Web3HostnameIpfsUniversalPathContentListListResponseEnvelope struct {
	Errors   []Web3HostnameIpfsUniversalPathContentListListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []Web3HostnameIpfsUniversalPathContentListListResponseEnvelopeMessages `json:"messages,required"`
	Result   Web3HostnameIpfsUniversalPathContentListListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success Web3HostnameIpfsUniversalPathContentListListResponseEnvelopeSuccess `json:"success,required"`
	JSON    web3HostnameIpfsUniversalPathContentListListResponseEnvelopeJSON    `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListListResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListListResponseEnvelope]
type web3HostnameIpfsUniversalPathContentListListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameIpfsUniversalPathContentListListResponseEnvelopeErrors struct {
	Code    int64                                                                  `json:"code,required"`
	Message string                                                                 `json:"message,required"`
	JSON    web3HostnameIpfsUniversalPathContentListListResponseEnvelopeErrorsJSON `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListListResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListListResponseEnvelopeErrors]
type web3HostnameIpfsUniversalPathContentListListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameIpfsUniversalPathContentListListResponseEnvelopeMessages struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    web3HostnameIpfsUniversalPathContentListListResponseEnvelopeMessagesJSON `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListListResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListListResponseEnvelopeMessages]
type web3HostnameIpfsUniversalPathContentListListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type Web3HostnameIpfsUniversalPathContentListListResponseEnvelopeSuccess bool

const (
	Web3HostnameIpfsUniversalPathContentListListResponseEnvelopeSuccessTrue Web3HostnameIpfsUniversalPathContentListListResponseEnvelopeSuccess = true
)

type Web3HostnameIpfsUniversalPathContentListReplaceParams struct {
	// Behavior of the content list.
	Action param.Field[Web3HostnameIpfsUniversalPathContentListReplaceParamsAction] `json:"action,required"`
	// Content list entries.
	Entries param.Field[[]Web3HostnameIpfsUniversalPathContentListReplaceParamsEntry] `json:"entries,required"`
}

func (r Web3HostnameIpfsUniversalPathContentListReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Behavior of the content list.
type Web3HostnameIpfsUniversalPathContentListReplaceParamsAction string

const (
	Web3HostnameIpfsUniversalPathContentListReplaceParamsActionBlock Web3HostnameIpfsUniversalPathContentListReplaceParamsAction = "block"
)

// Content list entry to be blocked.
type Web3HostnameIpfsUniversalPathContentListReplaceParamsEntry struct {
	// CID or content path of content to block.
	Content param.Field[string] `json:"content"`
	// An optional description of the content list entry.
	Description param.Field[string] `json:"description"`
	// Type of content list entry to block.
	Type param.Field[Web3HostnameIpfsUniversalPathContentListReplaceParamsEntriesType] `json:"type"`
}

func (r Web3HostnameIpfsUniversalPathContentListReplaceParamsEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of content list entry to block.
type Web3HostnameIpfsUniversalPathContentListReplaceParamsEntriesType string

const (
	Web3HostnameIpfsUniversalPathContentListReplaceParamsEntriesTypeCid         Web3HostnameIpfsUniversalPathContentListReplaceParamsEntriesType = "cid"
	Web3HostnameIpfsUniversalPathContentListReplaceParamsEntriesTypeContentPath Web3HostnameIpfsUniversalPathContentListReplaceParamsEntriesType = "content_path"
)

type Web3HostnameIpfsUniversalPathContentListReplaceResponseEnvelope struct {
	Errors   []Web3HostnameIpfsUniversalPathContentListReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []Web3HostnameIpfsUniversalPathContentListReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   Web3HostnameIpfsUniversalPathContentListReplaceResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success Web3HostnameIpfsUniversalPathContentListReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    web3HostnameIpfsUniversalPathContentListReplaceResponseEnvelopeJSON    `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListReplaceResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListReplaceResponseEnvelope]
type web3HostnameIpfsUniversalPathContentListReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameIpfsUniversalPathContentListReplaceResponseEnvelopeErrors struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    web3HostnameIpfsUniversalPathContentListReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListReplaceResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListReplaceResponseEnvelopeErrors]
type web3HostnameIpfsUniversalPathContentListReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameIpfsUniversalPathContentListReplaceResponseEnvelopeMessages struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    web3HostnameIpfsUniversalPathContentListReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListReplaceResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListReplaceResponseEnvelopeMessages]
type web3HostnameIpfsUniversalPathContentListReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type Web3HostnameIpfsUniversalPathContentListReplaceResponseEnvelopeSuccess bool

const (
	Web3HostnameIpfsUniversalPathContentListReplaceResponseEnvelopeSuccessTrue Web3HostnameIpfsUniversalPathContentListReplaceResponseEnvelopeSuccess = true
)
