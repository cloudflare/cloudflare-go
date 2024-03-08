// File generated from our OpenAPI spec by Stainless.

package dcv_delegation

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// UUIDService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewUUIDService] method instead.
type UUIDService struct {
	Options []option.RequestOption
}

// NewUUIDService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUUIDService(opts ...option.RequestOption) (r *UUIDService) {
	r = &UUIDService{}
	r.Options = opts
	return
}

// Retrieve the account and zone specific unique identifier used as part of the
// CNAME target for DCV Delegation.
func (r *UUIDService) Get(ctx context.Context, query UUIDGetParams, opts ...option.RequestOption) (res *TLSCertificatesAndHostnamesUUIDObject, err error) {
	opts = append(r.Options[:], opts...)
	var env UUIDGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/dcv_delegation/uuid", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TLSCertificatesAndHostnamesUUIDObject struct {
	// The DCV Delegation unique identifier.
	UUID string                                    `json:"uuid"`
	JSON tlsCertificatesAndHostnamesUUIDObjectJSON `json:"-"`
}

// tlsCertificatesAndHostnamesUUIDObjectJSON contains the JSON metadata for the
// struct [TLSCertificatesAndHostnamesUUIDObject]
type tlsCertificatesAndHostnamesUUIDObjectJSON struct {
	UUID        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TLSCertificatesAndHostnamesUUIDObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tlsCertificatesAndHostnamesUUIDObjectJSON) RawJSON() string {
	return r.raw
}

type UUIDGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type UUIDGetResponseEnvelope struct {
	Errors   []UUIDGetResponseEnvelopeErrors       `json:"errors,required"`
	Messages []UUIDGetResponseEnvelopeMessages     `json:"messages,required"`
	Result   TLSCertificatesAndHostnamesUUIDObject `json:"result,required"`
	// Whether the API call was successful
	Success UUIDGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    uuidGetResponseEnvelopeJSON    `json:"-"`
}

// uuidGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [UUIDGetResponseEnvelope]
type uuidGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UUIDGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uuidGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type UUIDGetResponseEnvelopeErrors struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    uuidGetResponseEnvelopeErrorsJSON `json:"-"`
}

// uuidGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [UUIDGetResponseEnvelopeErrors]
type uuidGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UUIDGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uuidGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type UUIDGetResponseEnvelopeMessages struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    uuidGetResponseEnvelopeMessagesJSON `json:"-"`
}

// uuidGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [UUIDGetResponseEnvelopeMessages]
type uuidGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UUIDGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uuidGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type UUIDGetResponseEnvelopeSuccess bool

const (
	UUIDGetResponseEnvelopeSuccessTrue UUIDGetResponseEnvelopeSuccess = true
)
