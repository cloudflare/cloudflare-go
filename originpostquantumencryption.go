// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
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
	var env OriginPostQuantumEncryptionGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/origin_post_quantum_encryption", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Instructs Cloudflare to use Post-Quantum (PQ) key agreement algorithms when
// connecting to your origin. Preferred instructs Cloudflare to opportunistically
// send a Post-Quantum keyshare in the first message to the origin (for fastest
// connections when the origin supports and prefers PQ), supported means that PQ
// algorithms are advertised but only used when requested by the origin, and off
// means that PQ algorithms are not advertised
func (r *OriginPostQuantumEncryptionService) Replace(ctx context.Context, zoneID string, body OriginPostQuantumEncryptionReplaceParams, opts ...option.RequestOption) (res *OriginPostQuantumEncryptionReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env OriginPostQuantumEncryptionReplaceResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/origin_post_quantum_encryption", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [OriginPostQuantumEncryptionGetResponseUnknown] or
// [shared.UnionString].
type OriginPostQuantumEncryptionGetResponse interface {
	ImplementsOriginPostQuantumEncryptionGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*OriginPostQuantumEncryptionGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [OriginPostQuantumEncryptionReplaceResponseUnknown] or
// [shared.UnionString].
type OriginPostQuantumEncryptionReplaceResponse interface {
	ImplementsOriginPostQuantumEncryptionReplaceResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*OriginPostQuantumEncryptionReplaceResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type OriginPostQuantumEncryptionGetResponseEnvelope struct {
	Errors   []OriginPostQuantumEncryptionGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []OriginPostQuantumEncryptionGetResponseEnvelopeMessages `json:"messages,required"`
	Result   OriginPostQuantumEncryptionGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success OriginPostQuantumEncryptionGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    originPostQuantumEncryptionGetResponseEnvelopeJSON    `json:"-"`
}

// originPostQuantumEncryptionGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [OriginPostQuantumEncryptionGetResponseEnvelope]
type originPostQuantumEncryptionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginPostQuantumEncryptionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OriginPostQuantumEncryptionGetResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    originPostQuantumEncryptionGetResponseEnvelopeErrorsJSON `json:"-"`
}

// originPostQuantumEncryptionGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [OriginPostQuantumEncryptionGetResponseEnvelopeErrors]
type originPostQuantumEncryptionGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginPostQuantumEncryptionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OriginPostQuantumEncryptionGetResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    originPostQuantumEncryptionGetResponseEnvelopeMessagesJSON `json:"-"`
}

// originPostQuantumEncryptionGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [OriginPostQuantumEncryptionGetResponseEnvelopeMessages]
type originPostQuantumEncryptionGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginPostQuantumEncryptionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type OriginPostQuantumEncryptionGetResponseEnvelopeSuccess bool

const (
	OriginPostQuantumEncryptionGetResponseEnvelopeSuccessTrue OriginPostQuantumEncryptionGetResponseEnvelopeSuccess = true
)

type OriginPostQuantumEncryptionReplaceParams struct {
	// Value of the Origin Post Quantum Encryption Setting.
	Value param.Field[OriginPostQuantumEncryptionReplaceParamsValue] `json:"value,required"`
}

func (r OriginPostQuantumEncryptionReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the Origin Post Quantum Encryption Setting.
type OriginPostQuantumEncryptionReplaceParamsValue string

const (
	OriginPostQuantumEncryptionReplaceParamsValuePreferred OriginPostQuantumEncryptionReplaceParamsValue = "preferred"
	OriginPostQuantumEncryptionReplaceParamsValueSupported OriginPostQuantumEncryptionReplaceParamsValue = "supported"
	OriginPostQuantumEncryptionReplaceParamsValueOff       OriginPostQuantumEncryptionReplaceParamsValue = "off"
)

type OriginPostQuantumEncryptionReplaceResponseEnvelope struct {
	Errors   []OriginPostQuantumEncryptionReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []OriginPostQuantumEncryptionReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   OriginPostQuantumEncryptionReplaceResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success OriginPostQuantumEncryptionReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    originPostQuantumEncryptionReplaceResponseEnvelopeJSON    `json:"-"`
}

// originPostQuantumEncryptionReplaceResponseEnvelopeJSON contains the JSON
// metadata for the struct [OriginPostQuantumEncryptionReplaceResponseEnvelope]
type originPostQuantumEncryptionReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginPostQuantumEncryptionReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OriginPostQuantumEncryptionReplaceResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    originPostQuantumEncryptionReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// originPostQuantumEncryptionReplaceResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [OriginPostQuantumEncryptionReplaceResponseEnvelopeErrors]
type originPostQuantumEncryptionReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginPostQuantumEncryptionReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OriginPostQuantumEncryptionReplaceResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    originPostQuantumEncryptionReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// originPostQuantumEncryptionReplaceResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [OriginPostQuantumEncryptionReplaceResponseEnvelopeMessages]
type originPostQuantumEncryptionReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginPostQuantumEncryptionReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type OriginPostQuantumEncryptionReplaceResponseEnvelopeSuccess bool

const (
	OriginPostQuantumEncryptionReplaceResponseEnvelopeSuccessTrue OriginPostQuantumEncryptionReplaceResponseEnvelopeSuccess = true
)
