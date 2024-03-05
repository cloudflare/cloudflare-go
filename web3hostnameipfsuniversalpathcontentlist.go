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

// Web3HostnameIPFSUniversalPathContentListService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWeb3HostnameIPFSUniversalPathContentListService] method instead.
type Web3HostnameIPFSUniversalPathContentListService struct {
	Options []option.RequestOption
	Entries *Web3HostnameIPFSUniversalPathContentListEntryService
}

// NewWeb3HostnameIPFSUniversalPathContentListService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewWeb3HostnameIPFSUniversalPathContentListService(opts ...option.RequestOption) (r *Web3HostnameIPFSUniversalPathContentListService) {
	r = &Web3HostnameIPFSUniversalPathContentListService{}
	r.Options = opts
	r.Entries = NewWeb3HostnameIPFSUniversalPathContentListEntryService(opts...)
	return
}

// Update IPFS Universal Path Gateway Content List
func (r *Web3HostnameIPFSUniversalPathContentListService) Update(ctx context.Context, zoneIdentifier string, identifier string, body Web3HostnameIPFSUniversalPathContentListUpdateParams, opts ...option.RequestOption) (res *DwebConfigContentListDetails, err error) {
	opts = append(r.Options[:], opts...)
	var env Web3HostnameIPFSUniversalPathContentListUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// IPFS Universal Path Gateway Content List Details
func (r *Web3HostnameIPFSUniversalPathContentListService) Get(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *DwebConfigContentListDetails, err error) {
	opts = append(r.Options[:], opts...)
	var env Web3HostnameIPFSUniversalPathContentListGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DwebConfigContentListDetails struct {
	// Behavior of the content list.
	Action DwebConfigContentListDetailsAction `json:"action"`
	JSON   dwebConfigContentListDetailsJSON   `json:"-"`
}

// dwebConfigContentListDetailsJSON contains the JSON metadata for the struct
// [DwebConfigContentListDetails]
type dwebConfigContentListDetailsJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DwebConfigContentListDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Behavior of the content list.
type DwebConfigContentListDetailsAction string

const (
	DwebConfigContentListDetailsActionBlock DwebConfigContentListDetailsAction = "block"
)

type Web3HostnameIPFSUniversalPathContentListUpdateParams struct {
	// Behavior of the content list.
	Action param.Field[Web3HostnameIPFSUniversalPathContentListUpdateParamsAction] `json:"action,required"`
	// Content list entries.
	Entries param.Field[[]DwebConfigContentListEntryParam] `json:"entries,required"`
}

func (r Web3HostnameIPFSUniversalPathContentListUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Behavior of the content list.
type Web3HostnameIPFSUniversalPathContentListUpdateParamsAction string

const (
	Web3HostnameIPFSUniversalPathContentListUpdateParamsActionBlock Web3HostnameIPFSUniversalPathContentListUpdateParamsAction = "block"
)

type Web3HostnameIPFSUniversalPathContentListUpdateResponseEnvelope struct {
	Errors   []Web3HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []Web3HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   DwebConfigContentListDetails                                             `json:"result,required"`
	// Whether the API call was successful
	Success Web3HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    web3HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeJSON    `json:"-"`
}

// web3HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [Web3HostnameIPFSUniversalPathContentListUpdateResponseEnvelope]
type web3HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIPFSUniversalPathContentListUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeErrors struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    web3HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// web3HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [Web3HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeErrors]
type web3HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeMessages struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    web3HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// web3HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [Web3HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeMessages]
type web3HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type Web3HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeSuccess bool

const (
	Web3HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeSuccessTrue Web3HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeSuccess = true
)

type Web3HostnameIPFSUniversalPathContentListGetResponseEnvelope struct {
	Errors   []Web3HostnameIPFSUniversalPathContentListGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []Web3HostnameIPFSUniversalPathContentListGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DwebConfigContentListDetails                                          `json:"result,required"`
	// Whether the API call was successful
	Success Web3HostnameIPFSUniversalPathContentListGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    web3HostnameIPFSUniversalPathContentListGetResponseEnvelopeJSON    `json:"-"`
}

// web3HostnameIPFSUniversalPathContentListGetResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [Web3HostnameIPFSUniversalPathContentListGetResponseEnvelope]
type web3HostnameIPFSUniversalPathContentListGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIPFSUniversalPathContentListGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameIPFSUniversalPathContentListGetResponseEnvelopeErrors struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    web3HostnameIPFSUniversalPathContentListGetResponseEnvelopeErrorsJSON `json:"-"`
}

// web3HostnameIPFSUniversalPathContentListGetResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [Web3HostnameIPFSUniversalPathContentListGetResponseEnvelopeErrors]
type web3HostnameIPFSUniversalPathContentListGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIPFSUniversalPathContentListGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Web3HostnameIPFSUniversalPathContentListGetResponseEnvelopeMessages struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    web3HostnameIPFSUniversalPathContentListGetResponseEnvelopeMessagesJSON `json:"-"`
}

// web3HostnameIPFSUniversalPathContentListGetResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [Web3HostnameIPFSUniversalPathContentListGetResponseEnvelopeMessages]
type web3HostnameIPFSUniversalPathContentListGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3HostnameIPFSUniversalPathContentListGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type Web3HostnameIPFSUniversalPathContentListGetResponseEnvelopeSuccess bool

const (
	Web3HostnameIPFSUniversalPathContentListGetResponseEnvelopeSuccessTrue Web3HostnameIPFSUniversalPathContentListGetResponseEnvelopeSuccess = true
)
