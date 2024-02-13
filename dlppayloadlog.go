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
func (r *DLPPayloadLogService) DLPPayloadLogSettingsGetSettings(ctx context.Context, accountID string, opts ...option.RequestOption) (res *DLPPayloadLogDLPPayloadLogSettingsGetSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DLPPayloadLogDLPPayloadLogSettingsGetSettingsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/payload_log", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates the DLP payload log settings for this account.
func (r *DLPPayloadLogService) DLPPayloadLogSettingsUpdateSettings(ctx context.Context, accountID string, body DLPPayloadLogDLPPayloadLogSettingsUpdateSettingsParams, opts ...option.RequestOption) (res *DLPPayloadLogDLPPayloadLogSettingsUpdateSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DLPPayloadLogDLPPayloadLogSettingsUpdateSettingsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/payload_log", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DLPPayloadLogDLPPayloadLogSettingsGetSettingsResponse struct {
	PublicKey string                                                    `json:"public_key,required,nullable"`
	JSON      dlpPayloadLogDLPPayloadLogSettingsGetSettingsResponseJSON `json:"-"`
}

// dlpPayloadLogDLPPayloadLogSettingsGetSettingsResponseJSON contains the JSON
// metadata for the struct [DLPPayloadLogDLPPayloadLogSettingsGetSettingsResponse]
type dlpPayloadLogDLPPayloadLogSettingsGetSettingsResponseJSON struct {
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPPayloadLogDLPPayloadLogSettingsGetSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPPayloadLogDLPPayloadLogSettingsUpdateSettingsResponse struct {
	PublicKey string                                                       `json:"public_key,required,nullable"`
	JSON      dlpPayloadLogDLPPayloadLogSettingsUpdateSettingsResponseJSON `json:"-"`
}

// dlpPayloadLogDLPPayloadLogSettingsUpdateSettingsResponseJSON contains the JSON
// metadata for the struct
// [DLPPayloadLogDLPPayloadLogSettingsUpdateSettingsResponse]
type dlpPayloadLogDLPPayloadLogSettingsUpdateSettingsResponseJSON struct {
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPPayloadLogDLPPayloadLogSettingsUpdateSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPPayloadLogDLPPayloadLogSettingsGetSettingsResponseEnvelope struct {
	Errors   []DLPPayloadLogDLPPayloadLogSettingsGetSettingsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPPayloadLogDLPPayloadLogSettingsGetSettingsResponseEnvelopeMessages `json:"messages,required"`
	Result   DLPPayloadLogDLPPayloadLogSettingsGetSettingsResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DLPPayloadLogDLPPayloadLogSettingsGetSettingsResponseEnvelopeSuccess `json:"success,required"`
	JSON    dlpPayloadLogDLPPayloadLogSettingsGetSettingsResponseEnvelopeJSON    `json:"-"`
}

// dlpPayloadLogDLPPayloadLogSettingsGetSettingsResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [DLPPayloadLogDLPPayloadLogSettingsGetSettingsResponseEnvelope]
type dlpPayloadLogDLPPayloadLogSettingsGetSettingsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPPayloadLogDLPPayloadLogSettingsGetSettingsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPPayloadLogDLPPayloadLogSettingsGetSettingsResponseEnvelopeErrors struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    dlpPayloadLogDLPPayloadLogSettingsGetSettingsResponseEnvelopeErrorsJSON `json:"-"`
}

// dlpPayloadLogDLPPayloadLogSettingsGetSettingsResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [DLPPayloadLogDLPPayloadLogSettingsGetSettingsResponseEnvelopeErrors]
type dlpPayloadLogDLPPayloadLogSettingsGetSettingsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPPayloadLogDLPPayloadLogSettingsGetSettingsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPPayloadLogDLPPayloadLogSettingsGetSettingsResponseEnvelopeMessages struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    dlpPayloadLogDLPPayloadLogSettingsGetSettingsResponseEnvelopeMessagesJSON `json:"-"`
}

// dlpPayloadLogDLPPayloadLogSettingsGetSettingsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [DLPPayloadLogDLPPayloadLogSettingsGetSettingsResponseEnvelopeMessages]
type dlpPayloadLogDLPPayloadLogSettingsGetSettingsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPPayloadLogDLPPayloadLogSettingsGetSettingsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DLPPayloadLogDLPPayloadLogSettingsGetSettingsResponseEnvelopeSuccess bool

const (
	DLPPayloadLogDLPPayloadLogSettingsGetSettingsResponseEnvelopeSuccessTrue DLPPayloadLogDLPPayloadLogSettingsGetSettingsResponseEnvelopeSuccess = true
)

type DLPPayloadLogDLPPayloadLogSettingsUpdateSettingsParams struct {
	// The public key to use when encrypting extracted payloads, as a base64 string
	PublicKey param.Field[string] `json:"public_key,required"`
}

func (r DLPPayloadLogDLPPayloadLogSettingsUpdateSettingsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPPayloadLogDLPPayloadLogSettingsUpdateSettingsResponseEnvelope struct {
	Errors   []DLPPayloadLogDLPPayloadLogSettingsUpdateSettingsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPPayloadLogDLPPayloadLogSettingsUpdateSettingsResponseEnvelopeMessages `json:"messages,required"`
	Result   DLPPayloadLogDLPPayloadLogSettingsUpdateSettingsResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DLPPayloadLogDLPPayloadLogSettingsUpdateSettingsResponseEnvelopeSuccess `json:"success,required"`
	JSON    dlpPayloadLogDLPPayloadLogSettingsUpdateSettingsResponseEnvelopeJSON    `json:"-"`
}

// dlpPayloadLogDLPPayloadLogSettingsUpdateSettingsResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [DLPPayloadLogDLPPayloadLogSettingsUpdateSettingsResponseEnvelope]
type dlpPayloadLogDLPPayloadLogSettingsUpdateSettingsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPPayloadLogDLPPayloadLogSettingsUpdateSettingsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPPayloadLogDLPPayloadLogSettingsUpdateSettingsResponseEnvelopeErrors struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    dlpPayloadLogDLPPayloadLogSettingsUpdateSettingsResponseEnvelopeErrorsJSON `json:"-"`
}

// dlpPayloadLogDLPPayloadLogSettingsUpdateSettingsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [DLPPayloadLogDLPPayloadLogSettingsUpdateSettingsResponseEnvelopeErrors]
type dlpPayloadLogDLPPayloadLogSettingsUpdateSettingsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPPayloadLogDLPPayloadLogSettingsUpdateSettingsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPPayloadLogDLPPayloadLogSettingsUpdateSettingsResponseEnvelopeMessages struct {
	Code    int64                                                                        `json:"code,required"`
	Message string                                                                       `json:"message,required"`
	JSON    dlpPayloadLogDLPPayloadLogSettingsUpdateSettingsResponseEnvelopeMessagesJSON `json:"-"`
}

// dlpPayloadLogDLPPayloadLogSettingsUpdateSettingsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [DLPPayloadLogDLPPayloadLogSettingsUpdateSettingsResponseEnvelopeMessages]
type dlpPayloadLogDLPPayloadLogSettingsUpdateSettingsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPPayloadLogDLPPayloadLogSettingsUpdateSettingsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DLPPayloadLogDLPPayloadLogSettingsUpdateSettingsResponseEnvelopeSuccess bool

const (
	DLPPayloadLogDLPPayloadLogSettingsUpdateSettingsResponseEnvelopeSuccessTrue DLPPayloadLogDLPPayloadLogSettingsUpdateSettingsResponseEnvelopeSuccess = true
)
