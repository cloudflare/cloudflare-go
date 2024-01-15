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
func (r *AccountDlpPayloadLogService) DlpPayloadLogSettingsGetSettings(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountDlpPayloadLogDlpPayloadLogSettingsGetSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/payload_log", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the DLP payload log settings for this account.
func (r *AccountDlpPayloadLogService) DlpPayloadLogSettingsUpdateSettings(ctx context.Context, accountIdentifier string, body AccountDlpPayloadLogDlpPayloadLogSettingsUpdateSettingsParams, opts ...option.RequestOption) (res *AccountDlpPayloadLogDlpPayloadLogSettingsUpdateSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/payload_log", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AccountDlpPayloadLogDlpPayloadLogSettingsGetSettingsResponse struct {
	Errors   []AccountDlpPayloadLogDlpPayloadLogSettingsGetSettingsResponseError   `json:"errors"`
	Messages []AccountDlpPayloadLogDlpPayloadLogSettingsGetSettingsResponseMessage `json:"messages"`
	Result   AccountDlpPayloadLogDlpPayloadLogSettingsGetSettingsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountDlpPayloadLogDlpPayloadLogSettingsGetSettingsResponseSuccess `json:"success"`
	JSON    accountDlpPayloadLogDlpPayloadLogSettingsGetSettingsResponseJSON    `json:"-"`
}

// accountDlpPayloadLogDlpPayloadLogSettingsGetSettingsResponseJSON contains the
// JSON metadata for the struct
// [AccountDlpPayloadLogDlpPayloadLogSettingsGetSettingsResponse]
type accountDlpPayloadLogDlpPayloadLogSettingsGetSettingsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpPayloadLogDlpPayloadLogSettingsGetSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDlpPayloadLogDlpPayloadLogSettingsGetSettingsResponseError struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    accountDlpPayloadLogDlpPayloadLogSettingsGetSettingsResponseErrorJSON `json:"-"`
}

// accountDlpPayloadLogDlpPayloadLogSettingsGetSettingsResponseErrorJSON contains
// the JSON metadata for the struct
// [AccountDlpPayloadLogDlpPayloadLogSettingsGetSettingsResponseError]
type accountDlpPayloadLogDlpPayloadLogSettingsGetSettingsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpPayloadLogDlpPayloadLogSettingsGetSettingsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDlpPayloadLogDlpPayloadLogSettingsGetSettingsResponseMessage struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    accountDlpPayloadLogDlpPayloadLogSettingsGetSettingsResponseMessageJSON `json:"-"`
}

// accountDlpPayloadLogDlpPayloadLogSettingsGetSettingsResponseMessageJSON contains
// the JSON metadata for the struct
// [AccountDlpPayloadLogDlpPayloadLogSettingsGetSettingsResponseMessage]
type accountDlpPayloadLogDlpPayloadLogSettingsGetSettingsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpPayloadLogDlpPayloadLogSettingsGetSettingsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDlpPayloadLogDlpPayloadLogSettingsGetSettingsResponseResult struct {
	PublicKey string                                                                 `json:"public_key,required,nullable"`
	JSON      accountDlpPayloadLogDlpPayloadLogSettingsGetSettingsResponseResultJSON `json:"-"`
}

// accountDlpPayloadLogDlpPayloadLogSettingsGetSettingsResponseResultJSON contains
// the JSON metadata for the struct
// [AccountDlpPayloadLogDlpPayloadLogSettingsGetSettingsResponseResult]
type accountDlpPayloadLogDlpPayloadLogSettingsGetSettingsResponseResultJSON struct {
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpPayloadLogDlpPayloadLogSettingsGetSettingsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountDlpPayloadLogDlpPayloadLogSettingsGetSettingsResponseSuccess bool

const (
	AccountDlpPayloadLogDlpPayloadLogSettingsGetSettingsResponseSuccessTrue AccountDlpPayloadLogDlpPayloadLogSettingsGetSettingsResponseSuccess = true
)

type AccountDlpPayloadLogDlpPayloadLogSettingsUpdateSettingsResponse struct {
	Errors   []AccountDlpPayloadLogDlpPayloadLogSettingsUpdateSettingsResponseError   `json:"errors"`
	Messages []AccountDlpPayloadLogDlpPayloadLogSettingsUpdateSettingsResponseMessage `json:"messages"`
	Result   AccountDlpPayloadLogDlpPayloadLogSettingsUpdateSettingsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountDlpPayloadLogDlpPayloadLogSettingsUpdateSettingsResponseSuccess `json:"success"`
	JSON    accountDlpPayloadLogDlpPayloadLogSettingsUpdateSettingsResponseJSON    `json:"-"`
}

// accountDlpPayloadLogDlpPayloadLogSettingsUpdateSettingsResponseJSON contains the
// JSON metadata for the struct
// [AccountDlpPayloadLogDlpPayloadLogSettingsUpdateSettingsResponse]
type accountDlpPayloadLogDlpPayloadLogSettingsUpdateSettingsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpPayloadLogDlpPayloadLogSettingsUpdateSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDlpPayloadLogDlpPayloadLogSettingsUpdateSettingsResponseError struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    accountDlpPayloadLogDlpPayloadLogSettingsUpdateSettingsResponseErrorJSON `json:"-"`
}

// accountDlpPayloadLogDlpPayloadLogSettingsUpdateSettingsResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountDlpPayloadLogDlpPayloadLogSettingsUpdateSettingsResponseError]
type accountDlpPayloadLogDlpPayloadLogSettingsUpdateSettingsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpPayloadLogDlpPayloadLogSettingsUpdateSettingsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDlpPayloadLogDlpPayloadLogSettingsUpdateSettingsResponseMessage struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    accountDlpPayloadLogDlpPayloadLogSettingsUpdateSettingsResponseMessageJSON `json:"-"`
}

// accountDlpPayloadLogDlpPayloadLogSettingsUpdateSettingsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountDlpPayloadLogDlpPayloadLogSettingsUpdateSettingsResponseMessage]
type accountDlpPayloadLogDlpPayloadLogSettingsUpdateSettingsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpPayloadLogDlpPayloadLogSettingsUpdateSettingsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDlpPayloadLogDlpPayloadLogSettingsUpdateSettingsResponseResult struct {
	PublicKey string                                                                    `json:"public_key,required,nullable"`
	JSON      accountDlpPayloadLogDlpPayloadLogSettingsUpdateSettingsResponseResultJSON `json:"-"`
}

// accountDlpPayloadLogDlpPayloadLogSettingsUpdateSettingsResponseResultJSON
// contains the JSON metadata for the struct
// [AccountDlpPayloadLogDlpPayloadLogSettingsUpdateSettingsResponseResult]
type accountDlpPayloadLogDlpPayloadLogSettingsUpdateSettingsResponseResultJSON struct {
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpPayloadLogDlpPayloadLogSettingsUpdateSettingsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountDlpPayloadLogDlpPayloadLogSettingsUpdateSettingsResponseSuccess bool

const (
	AccountDlpPayloadLogDlpPayloadLogSettingsUpdateSettingsResponseSuccessTrue AccountDlpPayloadLogDlpPayloadLogSettingsUpdateSettingsResponseSuccess = true
)

type AccountDlpPayloadLogDlpPayloadLogSettingsUpdateSettingsParams struct {
	// The public key to use when encrypting extracted payloads, as a base64 string
	PublicKey param.Field[string] `json:"public_key,required"`
}

func (r AccountDlpPayloadLogDlpPayloadLogSettingsUpdateSettingsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
