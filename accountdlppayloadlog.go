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

// AccountDlpPayloadLogService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountDlpPayloadLogService]
// method instead.
type AccountDlpPayloadLogService struct {
	Options []option.RequestOption
}

// NewAccountDlpPayloadLogService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountDlpPayloadLogService(opts ...option.RequestOption) (r *AccountDlpPayloadLogService) {
	r = &AccountDlpPayloadLogService{}
	r.Options = opts
	return
}

// Gets the current DLP payload log settings for this account.
func (r *AccountDlpPayloadLogService) DlpPayloadLogSettingsGetSettings(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *GetSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/payload_log", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the DLP payload log settings for this account.
func (r *AccountDlpPayloadLogService) DlpPayloadLogSettingsUpdateSettings(ctx context.Context, accountIdentifier string, body AccountDlpPayloadLogDlpPayloadLogSettingsUpdateSettingsParams, opts ...option.RequestOption) (res *UpdateSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/payload_log", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type GetSettingsResponse struct {
	Errors   []GetSettingsResponseError   `json:"errors"`
	Messages []GetSettingsResponseMessage `json:"messages"`
	Result   GetSettingsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success GetSettingsResponseSuccess `json:"success"`
	JSON    getSettingsResponseJSON    `json:"-"`
}

// getSettingsResponseJSON contains the JSON metadata for the struct
// [GetSettingsResponse]
type getSettingsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GetSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GetSettingsResponseError struct {
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    getSettingsResponseErrorJSON `json:"-"`
}

// getSettingsResponseErrorJSON contains the JSON metadata for the struct
// [GetSettingsResponseError]
type getSettingsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GetSettingsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GetSettingsResponseMessage struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    getSettingsResponseMessageJSON `json:"-"`
}

// getSettingsResponseMessageJSON contains the JSON metadata for the struct
// [GetSettingsResponseMessage]
type getSettingsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GetSettingsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GetSettingsResponseResult struct {
	PublicKey string                        `json:"public_key,required,nullable"`
	JSON      getSettingsResponseResultJSON `json:"-"`
}

// getSettingsResponseResultJSON contains the JSON metadata for the struct
// [GetSettingsResponseResult]
type getSettingsResponseResultJSON struct {
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GetSettingsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GetSettingsResponseSuccess bool

const (
	GetSettingsResponseSuccessTrue GetSettingsResponseSuccess = true
)

type UpdateSettingsResponse struct {
	Errors   []UpdateSettingsResponseError   `json:"errors"`
	Messages []UpdateSettingsResponseMessage `json:"messages"`
	Result   UpdateSettingsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success UpdateSettingsResponseSuccess `json:"success"`
	JSON    updateSettingsResponseJSON    `json:"-"`
}

// updateSettingsResponseJSON contains the JSON metadata for the struct
// [UpdateSettingsResponse]
type updateSettingsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UpdateSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UpdateSettingsResponseError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    updateSettingsResponseErrorJSON `json:"-"`
}

// updateSettingsResponseErrorJSON contains the JSON metadata for the struct
// [UpdateSettingsResponseError]
type updateSettingsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UpdateSettingsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UpdateSettingsResponseMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    updateSettingsResponseMessageJSON `json:"-"`
}

// updateSettingsResponseMessageJSON contains the JSON metadata for the struct
// [UpdateSettingsResponseMessage]
type updateSettingsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UpdateSettingsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UpdateSettingsResponseResult struct {
	PublicKey string                           `json:"public_key,required,nullable"`
	JSON      updateSettingsResponseResultJSON `json:"-"`
}

// updateSettingsResponseResultJSON contains the JSON metadata for the struct
// [UpdateSettingsResponseResult]
type updateSettingsResponseResultJSON struct {
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UpdateSettingsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UpdateSettingsResponseSuccess bool

const (
	UpdateSettingsResponseSuccessTrue UpdateSettingsResponseSuccess = true
)

type AccountDlpPayloadLogDlpPayloadLogSettingsUpdateSettingsParams struct {
	// The public key to use when encrypting extracted payloads, as a base64 string
	PublicKey param.Field[string] `json:"public_key,required"`
}

func (r AccountDlpPayloadLogDlpPayloadLogSettingsUpdateSettingsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
