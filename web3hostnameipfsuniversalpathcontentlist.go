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

// Update IPFS Universal Path Gateway Content List
func (r *Web3HostnameIpfsUniversalPathContentListService) Update(ctx context.Context, zoneIdentifier string, identifier string, body Web3HostnameIpfsUniversalPathContentListUpdateParams, opts ...option.RequestOption) (res *Web3HostnameIpfsUniversalPathContentListUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env Web3HostnameIpfsUniversalPathContentListUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
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

type Web3HostnameIpfsUniversalPathContentListUpdateResponse struct {
	// Behavior of the content list.
	Action Web3HostnameIpfsUniversalPathContentListUpdateResponseAction `json:"action"`
	JSON   web3HostnameIpfsUniversalPathContentListUpdateResponseJSON   `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListUpdateResponseJSON contains the JSON
// metadata for the struct [Web3HostnameIpfsUniversalPathContentListUpdateResponse]
type web3HostnameIpfsUniversalPathContentListUpdateResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Behavior of the content list.
type Web3HostnameIpfsUniversalPathContentListUpdateResponseAction string

const (
	Web3HostnameIpfsUniversalPathContentListUpdateResponseActionBlock Web3HostnameIpfsUniversalPathContentListUpdateResponseAction = "block"
)

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

type Web3HostnameIpfsUniversalPathContentListUpdateParams struct {
	// Behavior of the content list.
	Action param.Field[Web3HostnameIpfsUniversalPathContentListUpdateParamsAction] `json:"action,required"`
	// Content list entries.
	Entries param.Field[[]Web3HostnameIpfsUniversalPathContentListUpdateParamsEntry] `json:"entries,required"`
}

func (r Web3HostnameIpfsUniversalPathContentListUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Behavior of the content list.
type Web3HostnameIpfsUniversalPathContentListUpdateParamsAction string

const (
	Web3HostnameIpfsUniversalPathContentListUpdateParamsActionBlock Web3HostnameIpfsUniversalPathContentListUpdateParamsAction = "block"
)

// Content list entry to be blocked.
type Web3HostnameIpfsUniversalPathContentListUpdateParamsEntry struct {
	// CID or content path of content to block.
	Content param.Field[string] `json:"content"`
	// An optional description of the content list entry.
	Description param.Field[string] `json:"description"`
	// Type of content list entry to block.
	Type param.Field[Web3HostnameIpfsUniversalPathContentListUpdateParamsEntriesType] `json:"type"`
}

func (r Web3HostnameIpfsUniversalPathContentListUpdateParamsEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of content list entry to block.
type Web3HostnameIpfsUniversalPathContentListUpdateParamsEntriesType string

const (
	Web3HostnameIpfsUniversalPathContentListUpdateParamsEntriesTypeCid         Web3HostnameIpfsUniversalPathContentListUpdateParamsEntriesType = "cid"
	Web3HostnameIpfsUniversalPathContentListUpdateParamsEntriesTypeContentPath Web3HostnameIpfsUniversalPathContentListUpdateParamsEntriesType = "content_path"
)

type Web3HostnameIpfsUniversalPathContentListUpdateResponseEnvelope struct {
	Errors   []Web3HostnameIpfsUniversalPathContentListUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []Web3HostnameIpfsUniversalPathContentListUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   Web3HostnameIpfsUniversalPathContentListUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success Web3HostnameIpfsUniversalPathContentListUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    web3HostnameIpfsUniversalPathContentListUpdateResponseEnvelopeJSON    `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListUpdateResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListUpdateResponseEnvelope]
type web3HostnameIpfsUniversalPathContentListUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameIpfsUniversalPathContentListUpdateResponseEnvelopeErrors struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    web3HostnameIpfsUniversalPathContentListUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListUpdateResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListUpdateResponseEnvelopeErrors]
type web3HostnameIpfsUniversalPathContentListUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameIpfsUniversalPathContentListUpdateResponseEnvelopeMessages struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    web3HostnameIpfsUniversalPathContentListUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListUpdateResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListUpdateResponseEnvelopeMessages]
type web3HostnameIpfsUniversalPathContentListUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type Web3HostnameIpfsUniversalPathContentListUpdateResponseEnvelopeSuccess bool

const (
	Web3HostnameIpfsUniversalPathContentListUpdateResponseEnvelopeSuccessTrue Web3HostnameIpfsUniversalPathContentListUpdateResponseEnvelopeSuccess = true
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
