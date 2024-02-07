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

// OriginPostQuantumEncryptionService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewOriginPostQuantumEncryptionService] method instead.
type OriginPostQuantumEncryptionService struct {
	Options []option.RequestOption
}

// NewOriginPostQuantumEncryptionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewOriginPostQuantumEncryptionService(opts ...option.RequestOption) (r *OriginPostQuantumEncryptionService) {
	r = &OriginPostQuantumEncryptionService{}
	r.Options = opts
	return
}

// Instructs Cloudflare to use Post-Quantum (PQ) key agreement algorithms when
// connecting to your origin. Preferred instructs Cloudflare to opportunistically
// send a Post-Quantum keyshare in the first message to the origin (for fastest
// connections when the origin supports and prefers PQ), supported means that PQ
// algorithms are advertised but only used when requested by the origin, and off
// means that PQ algorithms are not advertised
func (r *OriginPostQuantumEncryptionService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *OriginPostQuantumEncryptionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/cache/origin_post_quantum_encryption", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Instructs Cloudflare to use Post-Quantum (PQ) key agreement algorithms when
// connecting to your origin. Preferred instructs Cloudflare to opportunistically
// send a Post-Quantum keyshare in the first message to the origin (for fastest
// connections when the origin supports and prefers PQ), supported means that PQ
// algorithms are advertised but only used when requested by the origin, and off
// means that PQ algorithms are not advertised
func (r *OriginPostQuantumEncryptionService) Update(ctx context.Context, zoneID string, body OriginPostQuantumEncryptionUpdateParams, opts ...option.RequestOption) (res *OriginPostQuantumEncryptionUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/cache/origin_post_quantum_encryption", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type OriginPostQuantumEncryptionGetResponse struct {
	Errors   []OriginPostQuantumEncryptionGetResponseError   `json:"errors"`
	Messages []OriginPostQuantumEncryptionGetResponseMessage `json:"messages"`
	Result   interface{}                                     `json:"result"`
	// Whether the API call was successful
	Success OriginPostQuantumEncryptionGetResponseSuccess `json:"success"`
	JSON    originPostQuantumEncryptionGetResponseJSON    `json:"-"`
}

// originPostQuantumEncryptionGetResponseJSON contains the JSON metadata for the
// struct [OriginPostQuantumEncryptionGetResponse]
type originPostQuantumEncryptionGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginPostQuantumEncryptionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OriginPostQuantumEncryptionGetResponseError struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    originPostQuantumEncryptionGetResponseErrorJSON `json:"-"`
}

// originPostQuantumEncryptionGetResponseErrorJSON contains the JSON metadata for
// the struct [OriginPostQuantumEncryptionGetResponseError]
type originPostQuantumEncryptionGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginPostQuantumEncryptionGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OriginPostQuantumEncryptionGetResponseMessage struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    originPostQuantumEncryptionGetResponseMessageJSON `json:"-"`
}

// originPostQuantumEncryptionGetResponseMessageJSON contains the JSON metadata for
// the struct [OriginPostQuantumEncryptionGetResponseMessage]
type originPostQuantumEncryptionGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginPostQuantumEncryptionGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type OriginPostQuantumEncryptionGetResponseSuccess bool

const (
	OriginPostQuantumEncryptionGetResponseSuccessTrue OriginPostQuantumEncryptionGetResponseSuccess = true
)

type OriginPostQuantumEncryptionUpdateResponse struct {
	Errors   []OriginPostQuantumEncryptionUpdateResponseError   `json:"errors"`
	Messages []OriginPostQuantumEncryptionUpdateResponseMessage `json:"messages"`
	Result   interface{}                                        `json:"result"`
	// Whether the API call was successful
	Success OriginPostQuantumEncryptionUpdateResponseSuccess `json:"success"`
	JSON    originPostQuantumEncryptionUpdateResponseJSON    `json:"-"`
}

// originPostQuantumEncryptionUpdateResponseJSON contains the JSON metadata for the
// struct [OriginPostQuantumEncryptionUpdateResponse]
type originPostQuantumEncryptionUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginPostQuantumEncryptionUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OriginPostQuantumEncryptionUpdateResponseError struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    originPostQuantumEncryptionUpdateResponseErrorJSON `json:"-"`
}

// originPostQuantumEncryptionUpdateResponseErrorJSON contains the JSON metadata
// for the struct [OriginPostQuantumEncryptionUpdateResponseError]
type originPostQuantumEncryptionUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginPostQuantumEncryptionUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OriginPostQuantumEncryptionUpdateResponseMessage struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    originPostQuantumEncryptionUpdateResponseMessageJSON `json:"-"`
}

// originPostQuantumEncryptionUpdateResponseMessageJSON contains the JSON metadata
// for the struct [OriginPostQuantumEncryptionUpdateResponseMessage]
type originPostQuantumEncryptionUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginPostQuantumEncryptionUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type OriginPostQuantumEncryptionUpdateResponseSuccess bool

const (
	OriginPostQuantumEncryptionUpdateResponseSuccessTrue OriginPostQuantumEncryptionUpdateResponseSuccess = true
)

type OriginPostQuantumEncryptionUpdateParams struct {
	// Value of the Origin Post Quantum Encryption Setting.
	Value param.Field[OriginPostQuantumEncryptionUpdateParamsValue] `json:"value,required"`
}

func (r OriginPostQuantumEncryptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the Origin Post Quantum Encryption Setting.
type OriginPostQuantumEncryptionUpdateParamsValue string

const (
	OriginPostQuantumEncryptionUpdateParamsValuePreferred OriginPostQuantumEncryptionUpdateParamsValue = "preferred"
	OriginPostQuantumEncryptionUpdateParamsValueSupported OriginPostQuantumEncryptionUpdateParamsValue = "supported"
	OriginPostQuantumEncryptionUpdateParamsValueOff       OriginPostQuantumEncryptionUpdateParamsValue = "off"
)
