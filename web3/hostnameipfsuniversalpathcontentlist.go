// File generated from our OpenAPI spec by Stainless.

package web3

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// HostnameIPFSUniversalPathContentListService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewHostnameIPFSUniversalPathContentListService] method instead.
type HostnameIPFSUniversalPathContentListService struct {
	Options []option.RequestOption
	Entries *HostnameIPFSUniversalPathContentListEntryService
}

// NewHostnameIPFSUniversalPathContentListService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewHostnameIPFSUniversalPathContentListService(opts ...option.RequestOption) (r *HostnameIPFSUniversalPathContentListService) {
	r = &HostnameIPFSUniversalPathContentListService{}
	r.Options = opts
	r.Entries = NewHostnameIPFSUniversalPathContentListEntryService(opts...)
	return
}

// Update IPFS Universal Path Gateway Content List
func (r *HostnameIPFSUniversalPathContentListService) Update(ctx context.Context, zoneIdentifier string, identifier string, body HostnameIPFSUniversalPathContentListUpdateParams, opts ...option.RequestOption) (res *DwebConfigContentListDetails, err error) {
	opts = append(r.Options[:], opts...)
	var env HostnameIPFSUniversalPathContentListUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// IPFS Universal Path Gateway Content List Details
func (r *HostnameIPFSUniversalPathContentListService) Get(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *DwebConfigContentListDetails, err error) {
	opts = append(r.Options[:], opts...)
	var env HostnameIPFSUniversalPathContentListGetResponseEnvelope
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

func (r dwebConfigContentListDetailsJSON) RawJSON() string {
	return r.raw
}

// Behavior of the content list.
type DwebConfigContentListDetailsAction string

const (
	DwebConfigContentListDetailsActionBlock DwebConfigContentListDetailsAction = "block"
)

type HostnameIPFSUniversalPathContentListUpdateParams struct {
	// Behavior of the content list.
	Action param.Field[HostnameIPFSUniversalPathContentListUpdateParamsAction] `json:"action,required"`
	// Content list entries.
	Entries param.Field[[]DwebConfigContentListEntryParam] `json:"entries,required"`
}

func (r HostnameIPFSUniversalPathContentListUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Behavior of the content list.
type HostnameIPFSUniversalPathContentListUpdateParamsAction string

const (
	HostnameIPFSUniversalPathContentListUpdateParamsActionBlock HostnameIPFSUniversalPathContentListUpdateParamsAction = "block"
)

type HostnameIPFSUniversalPathContentListUpdateResponseEnvelope struct {
	Errors   []HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   DwebConfigContentListDetails                                         `json:"result,required"`
	// Whether the API call was successful
	Success HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    hostnameIPFSUniversalPathContentListUpdateResponseEnvelopeJSON    `json:"-"`
}

// hostnameIPFSUniversalPathContentListUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [HostnameIPFSUniversalPathContentListUpdateResponseEnvelope]
type hostnameIPFSUniversalPathContentListUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameIPFSUniversalPathContentListUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameIPFSUniversalPathContentListUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeErrors struct {
	Code    int64                                                                `json:"code,required"`
	Message string                                                               `json:"message,required"`
	JSON    hostnameIPFSUniversalPathContentListUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// hostnameIPFSUniversalPathContentListUpdateResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeErrors]
type hostnameIPFSUniversalPathContentListUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameIPFSUniversalPathContentListUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeMessages struct {
	Code    int64                                                                  `json:"code,required"`
	Message string                                                                 `json:"message,required"`
	JSON    hostnameIPFSUniversalPathContentListUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// hostnameIPFSUniversalPathContentListUpdateResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeMessages]
type hostnameIPFSUniversalPathContentListUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameIPFSUniversalPathContentListUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeSuccess bool

const (
	HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeSuccessTrue HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeSuccess = true
)

type HostnameIPFSUniversalPathContentListGetResponseEnvelope struct {
	Errors   []HostnameIPFSUniversalPathContentListGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HostnameIPFSUniversalPathContentListGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DwebConfigContentListDetails                                      `json:"result,required"`
	// Whether the API call was successful
	Success HostnameIPFSUniversalPathContentListGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    hostnameIPFSUniversalPathContentListGetResponseEnvelopeJSON    `json:"-"`
}

// hostnameIPFSUniversalPathContentListGetResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [HostnameIPFSUniversalPathContentListGetResponseEnvelope]
type hostnameIPFSUniversalPathContentListGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameIPFSUniversalPathContentListGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameIPFSUniversalPathContentListGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HostnameIPFSUniversalPathContentListGetResponseEnvelopeErrors struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    hostnameIPFSUniversalPathContentListGetResponseEnvelopeErrorsJSON `json:"-"`
}

// hostnameIPFSUniversalPathContentListGetResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [HostnameIPFSUniversalPathContentListGetResponseEnvelopeErrors]
type hostnameIPFSUniversalPathContentListGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameIPFSUniversalPathContentListGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameIPFSUniversalPathContentListGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type HostnameIPFSUniversalPathContentListGetResponseEnvelopeMessages struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    hostnameIPFSUniversalPathContentListGetResponseEnvelopeMessagesJSON `json:"-"`
}

// hostnameIPFSUniversalPathContentListGetResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [HostnameIPFSUniversalPathContentListGetResponseEnvelopeMessages]
type hostnameIPFSUniversalPathContentListGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameIPFSUniversalPathContentListGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameIPFSUniversalPathContentListGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HostnameIPFSUniversalPathContentListGetResponseEnvelopeSuccess bool

const (
	HostnameIPFSUniversalPathContentListGetResponseEnvelopeSuccessTrue HostnameIPFSUniversalPathContentListGetResponseEnvelopeSuccess = true
)
