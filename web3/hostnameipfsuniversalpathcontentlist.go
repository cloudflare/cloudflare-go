// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web3

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
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
func (r *HostnameIPFSUniversalPathContentListService) Update(ctx context.Context, zoneIdentifier string, identifier string, body HostnameIPFSUniversalPathContentListUpdateParams, opts ...option.RequestOption) (res *DistributedWebConfigContentList, err error) {
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
func (r *HostnameIPFSUniversalPathContentListService) Get(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *DistributedWebConfigContentList, err error) {
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

type DistributedWebConfigContentList struct {
	// Behavior of the content list.
	Action DistributedWebConfigContentListAction `json:"action"`
	JSON   distributedWebConfigContentListJSON   `json:"-"`
}

// distributedWebConfigContentListJSON contains the JSON metadata for the struct
// [DistributedWebConfigContentList]
type distributedWebConfigContentListJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DistributedWebConfigContentList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r distributedWebConfigContentListJSON) RawJSON() string {
	return r.raw
}

// Behavior of the content list.
type DistributedWebConfigContentListAction string

const (
	DistributedWebConfigContentListActionBlock DistributedWebConfigContentListAction = "block"
)

func (r DistributedWebConfigContentListAction) IsKnown() bool {
	switch r {
	case DistributedWebConfigContentListActionBlock:
		return true
	}
	return false
}

type HostnameIPFSUniversalPathContentListUpdateParams struct {
	// Behavior of the content list.
	Action param.Field[HostnameIPFSUniversalPathContentListUpdateParamsAction] `json:"action,required"`
	// Content list entries.
	Entries param.Field[[]DistributedWebConfigContentListEntryParam] `json:"entries,required"`
}

func (r HostnameIPFSUniversalPathContentListUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Behavior of the content list.
type HostnameIPFSUniversalPathContentListUpdateParamsAction string

const (
	HostnameIPFSUniversalPathContentListUpdateParamsActionBlock HostnameIPFSUniversalPathContentListUpdateParamsAction = "block"
)

func (r HostnameIPFSUniversalPathContentListUpdateParamsAction) IsKnown() bool {
	switch r {
	case HostnameIPFSUniversalPathContentListUpdateParamsActionBlock:
		return true
	}
	return false
}

type HostnameIPFSUniversalPathContentListUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo           `json:"errors,required"`
	Messages []shared.ResponseInfo           `json:"messages,required"`
	Result   DistributedWebConfigContentList `json:"result,required"`
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

// Whether the API call was successful
type HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeSuccess bool

const (
	HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeSuccessTrue HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeSuccess = true
)

func (r HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type HostnameIPFSUniversalPathContentListGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo           `json:"errors,required"`
	Messages []shared.ResponseInfo           `json:"messages,required"`
	Result   DistributedWebConfigContentList `json:"result,required"`
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

// Whether the API call was successful
type HostnameIPFSUniversalPathContentListGetResponseEnvelopeSuccess bool

const (
	HostnameIPFSUniversalPathContentListGetResponseEnvelopeSuccessTrue HostnameIPFSUniversalPathContentListGetResponseEnvelopeSuccess = true
)

func (r HostnameIPFSUniversalPathContentListGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HostnameIPFSUniversalPathContentListGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
