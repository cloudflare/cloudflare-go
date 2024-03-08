// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// IntelASNService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewIntelASNService] method instead.
type IntelASNService struct {
	Options []option.RequestOption
	Subnets *IntelASNSubnetService
}

// NewIntelASNService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIntelASNService(opts ...option.RequestOption) (r *IntelASNService) {
	r = &IntelASNService{}
	r.Options = opts
	r.Subnets = NewIntelASNSubnetService(opts...)
	return
}

// Get ASN Overview
func (r *IntelASNService) Get(ctx context.Context, asn int64, query IntelASNGetParams, opts ...option.RequestOption) (res *int64, err error) {
	opts = append(r.Options[:], opts...)
	var env IntelASNGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/asn/%v", query.AccountID, asn)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IntelASNGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type IntelASNGetResponseEnvelope struct {
	Errors   []IntelASNGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IntelASNGetResponseEnvelopeMessages `json:"messages,required"`
	Result   int64                                 `json:"result,required"`
	// Whether the API call was successful
	Success IntelASNGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    intelASNGetResponseEnvelopeJSON    `json:"-"`
}

// intelASNGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [IntelASNGetResponseEnvelope]
type intelASNGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelASNGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelASNGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type IntelASNGetResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    intelASNGetResponseEnvelopeErrorsJSON `json:"-"`
}

// intelASNGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [IntelASNGetResponseEnvelopeErrors]
type intelASNGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelASNGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelASNGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IntelASNGetResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    intelASNGetResponseEnvelopeMessagesJSON `json:"-"`
}

// intelASNGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [IntelASNGetResponseEnvelopeMessages]
type intelASNGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelASNGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelASNGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IntelASNGetResponseEnvelopeSuccess bool

const (
	IntelASNGetResponseEnvelopeSuccessTrue IntelASNGetResponseEnvelopeSuccess = true
)
