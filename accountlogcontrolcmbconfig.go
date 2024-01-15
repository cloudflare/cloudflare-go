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

// AccountLogControlCmbConfigService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountLogControlCmbConfigService] method instead.
type AccountLogControlCmbConfigService struct {
	Options []option.RequestOption
}

// NewAccountLogControlCmbConfigService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountLogControlCmbConfigService(opts ...option.RequestOption) (r *AccountLogControlCmbConfigService) {
	r = &AccountLogControlCmbConfigService{}
	r.Options = opts
	return
}

// Deletes CMB config.
func (r *AccountLogControlCmbConfigService) Delete(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountLogControlCmbConfigDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/logs/control/cmb/config", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Gets CMB config.
func (r *AccountLogControlCmbConfigService) GetAccountsAccountIdentifierLogsControlCmbConfig(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountLogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/logs/control/cmb/config", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates CMB config.
func (r *AccountLogControlCmbConfigService) PutAccountsAccountIdentifierLogsControlCmbConfig(ctx context.Context, accountIdentifier string, body AccountLogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigParams, opts ...option.RequestOption) (res *AccountLogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/logs/control/cmb/config", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountLogControlCmbConfigDeleteResponse struct {
	Errors   []AccountLogControlCmbConfigDeleteResponseError   `json:"errors"`
	Messages []AccountLogControlCmbConfigDeleteResponseMessage `json:"messages"`
	Result   interface{}                                       `json:"result,nullable"`
	// Whether the API call was successful
	Success AccountLogControlCmbConfigDeleteResponseSuccess `json:"success"`
	JSON    accountLogControlCmbConfigDeleteResponseJSON    `json:"-"`
}

// accountLogControlCmbConfigDeleteResponseJSON contains the JSON metadata for the
// struct [AccountLogControlCmbConfigDeleteResponse]
type accountLogControlCmbConfigDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogControlCmbConfigDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLogControlCmbConfigDeleteResponseError struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accountLogControlCmbConfigDeleteResponseErrorJSON `json:"-"`
}

// accountLogControlCmbConfigDeleteResponseErrorJSON contains the JSON metadata for
// the struct [AccountLogControlCmbConfigDeleteResponseError]
type accountLogControlCmbConfigDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogControlCmbConfigDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLogControlCmbConfigDeleteResponseMessage struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accountLogControlCmbConfigDeleteResponseMessageJSON `json:"-"`
}

// accountLogControlCmbConfigDeleteResponseMessageJSON contains the JSON metadata
// for the struct [AccountLogControlCmbConfigDeleteResponseMessage]
type accountLogControlCmbConfigDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogControlCmbConfigDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountLogControlCmbConfigDeleteResponseSuccess bool

const (
	AccountLogControlCmbConfigDeleteResponseSuccessTrue AccountLogControlCmbConfigDeleteResponseSuccess = true
)

type AccountLogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponse struct {
	Errors   []AccountLogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseError   `json:"errors"`
	Messages []AccountLogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseMessage `json:"messages"`
	Result   AccountLogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success AccountLogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseSuccess `json:"success"`
	JSON    accountLogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseJSON    `json:"-"`
}

// accountLogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseJSON
// contains the JSON metadata for the struct
// [AccountLogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponse]
type accountLogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseError struct {
	Code    int64                                                                                       `json:"code,required"`
	Message string                                                                                      `json:"message,required"`
	JSON    accountLogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseErrorJSON `json:"-"`
}

// accountLogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountLogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseError]
type accountLogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseMessage struct {
	Code    int64                                                                                         `json:"code,required"`
	Message string                                                                                        `json:"message,required"`
	JSON    accountLogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseMessageJSON `json:"-"`
}

// accountLogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountLogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseMessage]
type accountLogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseResult struct {
	// Comma-separated list of regions.
	Regions string                                                                                       `json:"regions"`
	JSON    accountLogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseResultJSON `json:"-"`
}

// accountLogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseResultJSON
// contains the JSON metadata for the struct
// [AccountLogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseResult]
type accountLogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseResultJSON struct {
	Regions     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountLogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseSuccess bool

const (
	AccountLogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseSuccessTrue AccountLogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseSuccess = true
)

type AccountLogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponse struct {
	Errors   []AccountLogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseError   `json:"errors"`
	Messages []AccountLogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseMessage `json:"messages"`
	Result   AccountLogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success AccountLogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseSuccess `json:"success"`
	JSON    accountLogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseJSON    `json:"-"`
}

// accountLogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseJSON
// contains the JSON metadata for the struct
// [AccountLogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponse]
type accountLogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseError struct {
	Code    int64                                                                                       `json:"code,required"`
	Message string                                                                                      `json:"message,required"`
	JSON    accountLogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseErrorJSON `json:"-"`
}

// accountLogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountLogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseError]
type accountLogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseMessage struct {
	Code    int64                                                                                         `json:"code,required"`
	Message string                                                                                        `json:"message,required"`
	JSON    accountLogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseMessageJSON `json:"-"`
}

// accountLogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountLogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseMessage]
type accountLogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseResult struct {
	// Comma-separated list of regions.
	Regions string                                                                                       `json:"regions"`
	JSON    accountLogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseResultJSON `json:"-"`
}

// accountLogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseResultJSON
// contains the JSON metadata for the struct
// [AccountLogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseResult]
type accountLogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseResultJSON struct {
	Regions     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountLogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseSuccess bool

const (
	AccountLogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseSuccessTrue AccountLogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseSuccess = true
)

type AccountLogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigParams struct {
	// Comma-separated list of regions.
	Regions param.Field[string] `json:"regions"`
}

func (r AccountLogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
