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

// DLPPayloadLogService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDLPPayloadLogService] method
// instead.
type DLPPayloadLogService struct {
	Options []option.RequestOption
}

// NewDLPPayloadLogService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDLPPayloadLogService(opts ...option.RequestOption) (r *DLPPayloadLogService) {
	r = &DLPPayloadLogService{}
	r.Options = opts
	return
}

// Gets the current DLP payload log settings for this account.
func (r *DLPPayloadLogService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *DLPPayloadLogGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DLPPayloadLogGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/payload_log", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates the DLP payload log settings for this account.
func (r *DLPPayloadLogService) Replace(ctx context.Context, accountID string, body DLPPayloadLogReplaceParams, opts ...option.RequestOption) (res *DLPPayloadLogReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DLPPayloadLogReplaceResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/payload_log", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DLPPayloadLogGetResponse struct {
	PublicKey string                       `json:"public_key,required,nullable"`
	JSON      dlpPayloadLogGetResponseJSON `json:"-"`
}

// dlpPayloadLogGetResponseJSON contains the JSON metadata for the struct
// [DLPPayloadLogGetResponse]
type dlpPayloadLogGetResponseJSON struct {
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPPayloadLogGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPPayloadLogReplaceResponse struct {
	PublicKey string                           `json:"public_key,required,nullable"`
	JSON      dlpPayloadLogReplaceResponseJSON `json:"-"`
}

// dlpPayloadLogReplaceResponseJSON contains the JSON metadata for the struct
// [DLPPayloadLogReplaceResponse]
type dlpPayloadLogReplaceResponseJSON struct {
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPPayloadLogReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPPayloadLogGetResponseEnvelope struct {
	Errors   []DLPPayloadLogGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPPayloadLogGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DLPPayloadLogGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DLPPayloadLogGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    dlpPayloadLogGetResponseEnvelopeJSON    `json:"-"`
}

// dlpPayloadLogGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPPayloadLogGetResponseEnvelope]
type dlpPayloadLogGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPPayloadLogGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPPayloadLogGetResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    dlpPayloadLogGetResponseEnvelopeErrorsJSON `json:"-"`
}

// dlpPayloadLogGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPPayloadLogGetResponseEnvelopeErrors]
type dlpPayloadLogGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPPayloadLogGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPPayloadLogGetResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    dlpPayloadLogGetResponseEnvelopeMessagesJSON `json:"-"`
}

// dlpPayloadLogGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DLPPayloadLogGetResponseEnvelopeMessages]
type dlpPayloadLogGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPPayloadLogGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DLPPayloadLogGetResponseEnvelopeSuccess bool

const (
	DLPPayloadLogGetResponseEnvelopeSuccessTrue DLPPayloadLogGetResponseEnvelopeSuccess = true
)

type DLPPayloadLogReplaceParams struct {
	// The public key to use when encrypting extracted payloads, as a base64 string
	PublicKey param.Field[string] `json:"public_key,required"`
}

func (r DLPPayloadLogReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPPayloadLogReplaceResponseEnvelope struct {
	Errors   []DLPPayloadLogReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPPayloadLogReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   DLPPayloadLogReplaceResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DLPPayloadLogReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    dlpPayloadLogReplaceResponseEnvelopeJSON    `json:"-"`
}

// dlpPayloadLogReplaceResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPPayloadLogReplaceResponseEnvelope]
type dlpPayloadLogReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPPayloadLogReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPPayloadLogReplaceResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    dlpPayloadLogReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// dlpPayloadLogReplaceResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DLPPayloadLogReplaceResponseEnvelopeErrors]
type dlpPayloadLogReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPPayloadLogReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPPayloadLogReplaceResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    dlpPayloadLogReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// dlpPayloadLogReplaceResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DLPPayloadLogReplaceResponseEnvelopeMessages]
type dlpPayloadLogReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPPayloadLogReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DLPPayloadLogReplaceResponseEnvelopeSuccess bool

const (
	DLPPayloadLogReplaceResponseEnvelopeSuccessTrue DLPPayloadLogReplaceResponseEnvelopeSuccess = true
)
