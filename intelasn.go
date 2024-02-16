// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// IntelAsnService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewIntelAsnService] method instead.
type IntelAsnService struct {
	Options []option.RequestOption
	Subnets *IntelAsnSubnetService
}

// NewIntelAsnService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIntelAsnService(opts ...option.RequestOption) (r *IntelAsnService) {
	r = &IntelAsnService{}
	r.Options = opts
	r.Subnets = NewIntelAsnSubnetService(opts...)
	return
}

// Get ASN Overview
func (r *IntelAsnService) Get(ctx context.Context, accountID string, asn int64, opts ...option.RequestOption) (res *int64, err error) {
	opts = append(r.Options[:], opts...)
	var env IntelAsnGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/asn/%v", accountID, asn)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IntelAsnGetResponseEnvelope struct {
	Errors   []IntelAsnGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IntelAsnGetResponseEnvelopeMessages `json:"messages,required"`
	Result   int64                                 `json:"result,required"`
	// Whether the API call was successful
	Success IntelAsnGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    intelAsnGetResponseEnvelopeJSON    `json:"-"`
}

// intelAsnGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [IntelAsnGetResponseEnvelope]
type intelAsnGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelAsnGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelAsnGetResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    intelAsnGetResponseEnvelopeErrorsJSON `json:"-"`
}

// intelAsnGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [IntelAsnGetResponseEnvelopeErrors]
type intelAsnGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelAsnGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelAsnGetResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    intelAsnGetResponseEnvelopeMessagesJSON `json:"-"`
}

// intelAsnGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [IntelAsnGetResponseEnvelopeMessages]
type intelAsnGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelAsnGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type IntelAsnGetResponseEnvelopeSuccess bool

const (
	IntelAsnGetResponseEnvelopeSuccessTrue IntelAsnGetResponseEnvelopeSuccess = true
)
