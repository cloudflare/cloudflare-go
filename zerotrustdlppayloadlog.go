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

// ZeroTrustDLPPayloadLogService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZeroTrustDLPPayloadLogService]
// method instead.
type ZeroTrustDLPPayloadLogService struct {
	Options []option.RequestOption
}

// NewZeroTrustDLPPayloadLogService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZeroTrustDLPPayloadLogService(opts ...option.RequestOption) (r *ZeroTrustDLPPayloadLogService) {
	r = &ZeroTrustDLPPayloadLogService{}
	r.Options = opts
	return
}

// Updates the DLP payload log settings for this account.
func (r *ZeroTrustDLPPayloadLogService) Update(ctx context.Context, params ZeroTrustDLPPayloadLogUpdateParams, opts ...option.RequestOption) (res *ZeroTrustDLPPayloadLogUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDLPPayloadLogUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/payload_log", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets the current DLP payload log settings for this account.
func (r *ZeroTrustDLPPayloadLogService) Get(ctx context.Context, query ZeroTrustDLPPayloadLogGetParams, opts ...option.RequestOption) (res *ZeroTrustDLPPayloadLogGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDLPPayloadLogGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/payload_log", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustDLPPayloadLogUpdateResponse struct {
	PublicKey string                                   `json:"public_key,required,nullable"`
	JSON      zeroTrustDLPPayloadLogUpdateResponseJSON `json:"-"`
}

// zeroTrustDLPPayloadLogUpdateResponseJSON contains the JSON metadata for the
// struct [ZeroTrustDLPPayloadLogUpdateResponse]
type zeroTrustDLPPayloadLogUpdateResponseJSON struct {
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPPayloadLogUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDLPPayloadLogGetResponse struct {
	PublicKey string                                `json:"public_key,required,nullable"`
	JSON      zeroTrustDLPPayloadLogGetResponseJSON `json:"-"`
}

// zeroTrustDLPPayloadLogGetResponseJSON contains the JSON metadata for the struct
// [ZeroTrustDLPPayloadLogGetResponse]
type zeroTrustDLPPayloadLogGetResponseJSON struct {
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPPayloadLogGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDLPPayloadLogUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The public key to use when encrypting extracted payloads, as a base64 string
	PublicKey param.Field[string] `json:"public_key,required"`
}

func (r ZeroTrustDLPPayloadLogUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustDLPPayloadLogUpdateResponseEnvelope struct {
	Errors   []ZeroTrustDLPPayloadLogUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDLPPayloadLogUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustDLPPayloadLogUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustDLPPayloadLogUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustDLPPayloadLogUpdateResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustDLPPayloadLogUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustDLPPayloadLogUpdateResponseEnvelope]
type zeroTrustDLPPayloadLogUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPPayloadLogUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDLPPayloadLogUpdateResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zeroTrustDLPPayloadLogUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDLPPayloadLogUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustDLPPayloadLogUpdateResponseEnvelopeErrors]
type zeroTrustDLPPayloadLogUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPPayloadLogUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDLPPayloadLogUpdateResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zeroTrustDLPPayloadLogUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDLPPayloadLogUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustDLPPayloadLogUpdateResponseEnvelopeMessages]
type zeroTrustDLPPayloadLogUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPPayloadLogUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustDLPPayloadLogUpdateResponseEnvelopeSuccess bool

const (
	ZeroTrustDLPPayloadLogUpdateResponseEnvelopeSuccessTrue ZeroTrustDLPPayloadLogUpdateResponseEnvelopeSuccess = true
)

type ZeroTrustDLPPayloadLogGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ZeroTrustDLPPayloadLogGetResponseEnvelope struct {
	Errors   []ZeroTrustDLPPayloadLogGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDLPPayloadLogGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustDLPPayloadLogGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustDLPPayloadLogGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustDLPPayloadLogGetResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustDLPPayloadLogGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustDLPPayloadLogGetResponseEnvelope]
type zeroTrustDLPPayloadLogGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPPayloadLogGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDLPPayloadLogGetResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zeroTrustDLPPayloadLogGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDLPPayloadLogGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustDLPPayloadLogGetResponseEnvelopeErrors]
type zeroTrustDLPPayloadLogGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPPayloadLogGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDLPPayloadLogGetResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    zeroTrustDLPPayloadLogGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDLPPayloadLogGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustDLPPayloadLogGetResponseEnvelopeMessages]
type zeroTrustDLPPayloadLogGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPPayloadLogGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustDLPPayloadLogGetResponseEnvelopeSuccess bool

const (
	ZeroTrustDLPPayloadLogGetResponseEnvelopeSuccessTrue ZeroTrustDLPPayloadLogGetResponseEnvelopeSuccess = true
)
