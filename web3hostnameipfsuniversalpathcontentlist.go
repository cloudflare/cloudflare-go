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
func (r *Web3HostnameIpfsUniversalPathContentListService) Web3HostnameIpfsUniversalPathGatewayContentListDetails(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *Web3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env Web3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseEnvelope
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update IPFS Universal Path Gateway Content List
func (r *Web3HostnameIpfsUniversalPathContentListService) Web3HostnameUpdateIpfsUniversalPathGatewayContentList(ctx context.Context, zoneIdentifier string, identifier string, body Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListParams, opts ...option.RequestOption) (res *Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseEnvelope
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Web3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponse struct {
	// Behavior of the content list.
	Action Web3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseAction `json:"action"`
	JSON   web3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseJSON   `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseJSON
// contains the JSON metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponse]
type web3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Behavior of the content list.
type Web3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseAction string

const (
	Web3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseActionBlock Web3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseAction = "block"
)

type Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponse struct {
	// Behavior of the content list.
	Action Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseAction `json:"action"`
	JSON   web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseJSON   `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseJSON
// contains the JSON metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponse]
type web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Behavior of the content list.
type Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseAction string

const (
	Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseActionBlock Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseAction = "block"
)

type Web3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseEnvelope struct {
	Errors   []Web3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []Web3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseEnvelopeMessages `json:"messages,required"`
	Result   Web3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success Web3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseEnvelopeSuccess `json:"success,required"`
	JSON    web3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseEnvelopeJSON    `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseEnvelope]
type web3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseEnvelopeErrors struct {
	Code    int64                                                                                                                    `json:"code,required"`
	Message string                                                                                                                   `json:"message,required"`
	JSON    web3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseEnvelopeErrorsJSON `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseEnvelopeErrors]
type web3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseEnvelopeMessages struct {
	Code    int64                                                                                                                      `json:"code,required"`
	Message string                                                                                                                     `json:"message,required"`
	JSON    web3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseEnvelopeMessagesJSON `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseEnvelopeMessages]
type web3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type Web3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseEnvelopeSuccess bool

const (
	Web3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseEnvelopeSuccessTrue Web3HostnameIpfsUniversalPathContentListWeb3HostnameIpfsUniversalPathGatewayContentListDetailsResponseEnvelopeSuccess = true
)

type Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListParams struct {
	// Behavior of the content list.
	Action param.Field[Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListParamsAction] `json:"action,required"`
	// Content list entries.
	Entries param.Field[[]Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListParamsEntry] `json:"entries,required"`
}

func (r Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Behavior of the content list.
type Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListParamsAction string

const (
	Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListParamsActionBlock Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListParamsAction = "block"
)

// Content list entry to be blocked.
type Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListParamsEntry struct {
	// CID or content path of content to block.
	Content param.Field[string] `json:"content"`
	// An optional description of the content list entry.
	Description param.Field[string] `json:"description"`
	// Type of content list entry to block.
	Type param.Field[Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListParamsEntriesType] `json:"type"`
}

func (r Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListParamsEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of content list entry to block.
type Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListParamsEntriesType string

const (
	Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListParamsEntriesTypeCid         Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListParamsEntriesType = "cid"
	Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListParamsEntriesTypeContentPath Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListParamsEntriesType = "content_path"
)

type Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseEnvelope struct {
	Errors   []Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseEnvelopeMessages `json:"messages,required"`
	Result   Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseEnvelopeSuccess `json:"success,required"`
	JSON    web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseEnvelopeJSON    `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseEnvelope]
type web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseEnvelopeErrors struct {
	Code    int64                                                                                                                   `json:"code,required"`
	Message string                                                                                                                  `json:"message,required"`
	JSON    web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseEnvelopeErrorsJSON `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseEnvelopeErrors]
type web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseEnvelopeMessages struct {
	Code    int64                                                                                                                     `json:"code,required"`
	Message string                                                                                                                    `json:"message,required"`
	JSON    web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseEnvelopeMessagesJSON `json:"-"`
}

// web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseEnvelopeMessages]
type web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseEnvelopeSuccess bool

const (
	Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseEnvelopeSuccessTrue Web3HostnameIpfsUniversalPathContentListWeb3HostnameUpdateIpfsUniversalPathGatewayContentListResponseEnvelopeSuccess = true
)
