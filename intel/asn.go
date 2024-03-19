// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// ASNService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewASNService] method instead.
type ASNService struct {
	Options []option.RequestOption
	Subnets *ASNSubnetService
}

// NewASNService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewASNService(opts ...option.RequestOption) (r *ASNService) {
	r = &ASNService{}
	r.Options = opts
	r.Subnets = NewASNSubnetService(opts...)
	return
}

// Get ASN Overview
func (r *ASNService) Get(ctx context.Context, asn IntelASNParam, query ASNGetParams, opts ...option.RequestOption) (res *IntelASN, err error) {
	opts = append(r.Options[:], opts...)
	var env ASNGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/asn/%v", query.AccountID, asn)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IntelASN = int64

type IntelASNParam = int64

type ASNGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ASNGetResponseEnvelope struct {
	Errors   []ASNGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ASNGetResponseEnvelopeMessages `json:"messages,required"`
	Result   IntelASN                         `json:"result,required"`
	// Whether the API call was successful
	Success ASNGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    asnGetResponseEnvelopeJSON    `json:"-"`
}

// asnGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ASNGetResponseEnvelope]
type asnGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ASNGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r asnGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ASNGetResponseEnvelopeErrors struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    asnGetResponseEnvelopeErrorsJSON `json:"-"`
}

// asnGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ASNGetResponseEnvelopeErrors]
type asnGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ASNGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r asnGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ASNGetResponseEnvelopeMessages struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    asnGetResponseEnvelopeMessagesJSON `json:"-"`
}

// asnGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [ASNGetResponseEnvelopeMessages]
type asnGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ASNGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r asnGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ASNGetResponseEnvelopeSuccess bool

const (
	ASNGetResponseEnvelopeSuccessTrue ASNGetResponseEnvelopeSuccess = true
)

func (r ASNGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ASNGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
