// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountStreamKeyService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountStreamKeyService] method
// instead.
type AccountStreamKeyService struct {
	Options []option.RequestOption
}

// NewAccountStreamKeyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountStreamKeyService(opts ...option.RequestOption) (r *AccountStreamKeyService) {
	r = &AccountStreamKeyService{}
	r.Options = opts
	return
}

// Deletes signing keys and revokes all signed URLs generated with the key.
func (r *AccountStreamKeyService) Delete(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *AccountStreamKeyDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/keys/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates an RSA private key in PEM and JWK formats. Key files are only displayed
// once after creation. Keys are created, used, and deleted independently of
// videos, and every key can sign any video.
func (r *AccountStreamKeyService) StreamSigningKeysNewSigningKeys(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountStreamKeyStreamSigningKeysNewSigningKeysResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/keys", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Lists the video ID and creation date and time when a signing key was created.
func (r *AccountStreamKeyService) StreamSigningKeysListSigningKeys(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountStreamKeyStreamSigningKeysListSigningKeysResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/keys", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountStreamKeyDeleteResponse struct {
	Errors   []AccountStreamKeyDeleteResponseError   `json:"errors"`
	Messages []AccountStreamKeyDeleteResponseMessage `json:"messages"`
	Result   string                                  `json:"result"`
	// Whether the API call was successful
	Success AccountStreamKeyDeleteResponseSuccess `json:"success"`
	JSON    accountStreamKeyDeleteResponseJSON    `json:"-"`
}

// accountStreamKeyDeleteResponseJSON contains the JSON metadata for the struct
// [AccountStreamKeyDeleteResponse]
type accountStreamKeyDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamKeyDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamKeyDeleteResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    accountStreamKeyDeleteResponseErrorJSON `json:"-"`
}

// accountStreamKeyDeleteResponseErrorJSON contains the JSON metadata for the
// struct [AccountStreamKeyDeleteResponseError]
type accountStreamKeyDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamKeyDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamKeyDeleteResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accountStreamKeyDeleteResponseMessageJSON `json:"-"`
}

// accountStreamKeyDeleteResponseMessageJSON contains the JSON metadata for the
// struct [AccountStreamKeyDeleteResponseMessage]
type accountStreamKeyDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamKeyDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountStreamKeyDeleteResponseSuccess bool

const (
	AccountStreamKeyDeleteResponseSuccessTrue AccountStreamKeyDeleteResponseSuccess = true
)

type AccountStreamKeyStreamSigningKeysNewSigningKeysResponse struct {
	Errors   []AccountStreamKeyStreamSigningKeysNewSigningKeysResponseError   `json:"errors"`
	Messages []AccountStreamKeyStreamSigningKeysNewSigningKeysResponseMessage `json:"messages"`
	Result   AccountStreamKeyStreamSigningKeysNewSigningKeysResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountStreamKeyStreamSigningKeysNewSigningKeysResponseSuccess `json:"success"`
	JSON    accountStreamKeyStreamSigningKeysNewSigningKeysResponseJSON    `json:"-"`
}

// accountStreamKeyStreamSigningKeysNewSigningKeysResponseJSON contains the JSON
// metadata for the struct
// [AccountStreamKeyStreamSigningKeysNewSigningKeysResponse]
type accountStreamKeyStreamSigningKeysNewSigningKeysResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamKeyStreamSigningKeysNewSigningKeysResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamKeyStreamSigningKeysNewSigningKeysResponseError struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    accountStreamKeyStreamSigningKeysNewSigningKeysResponseErrorJSON `json:"-"`
}

// accountStreamKeyStreamSigningKeysNewSigningKeysResponseErrorJSON contains the
// JSON metadata for the struct
// [AccountStreamKeyStreamSigningKeysNewSigningKeysResponseError]
type accountStreamKeyStreamSigningKeysNewSigningKeysResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamKeyStreamSigningKeysNewSigningKeysResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamKeyStreamSigningKeysNewSigningKeysResponseMessage struct {
	Code    int64                                                              `json:"code,required"`
	Message string                                                             `json:"message,required"`
	JSON    accountStreamKeyStreamSigningKeysNewSigningKeysResponseMessageJSON `json:"-"`
}

// accountStreamKeyStreamSigningKeysNewSigningKeysResponseMessageJSON contains the
// JSON metadata for the struct
// [AccountStreamKeyStreamSigningKeysNewSigningKeysResponseMessage]
type accountStreamKeyStreamSigningKeysNewSigningKeysResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamKeyStreamSigningKeysNewSigningKeysResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamKeyStreamSigningKeysNewSigningKeysResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// The date and time a signing key was created.
	Created time.Time `json:"created" format:"date-time"`
	// The signing key in JWK format.
	Jwk string `json:"jwk"`
	// The signing key in PEM format.
	Pem  string                                                            `json:"pem"`
	JSON accountStreamKeyStreamSigningKeysNewSigningKeysResponseResultJSON `json:"-"`
}

// accountStreamKeyStreamSigningKeysNewSigningKeysResponseResultJSON contains the
// JSON metadata for the struct
// [AccountStreamKeyStreamSigningKeysNewSigningKeysResponseResult]
type accountStreamKeyStreamSigningKeysNewSigningKeysResponseResultJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Jwk         apijson.Field
	Pem         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamKeyStreamSigningKeysNewSigningKeysResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountStreamKeyStreamSigningKeysNewSigningKeysResponseSuccess bool

const (
	AccountStreamKeyStreamSigningKeysNewSigningKeysResponseSuccessTrue AccountStreamKeyStreamSigningKeysNewSigningKeysResponseSuccess = true
)

type AccountStreamKeyStreamSigningKeysListSigningKeysResponse struct {
	Errors   []AccountStreamKeyStreamSigningKeysListSigningKeysResponseError   `json:"errors"`
	Messages []AccountStreamKeyStreamSigningKeysListSigningKeysResponseMessage `json:"messages"`
	Result   []AccountStreamKeyStreamSigningKeysListSigningKeysResponseResult  `json:"result"`
	// Whether the API call was successful
	Success AccountStreamKeyStreamSigningKeysListSigningKeysResponseSuccess `json:"success"`
	JSON    accountStreamKeyStreamSigningKeysListSigningKeysResponseJSON    `json:"-"`
}

// accountStreamKeyStreamSigningKeysListSigningKeysResponseJSON contains the JSON
// metadata for the struct
// [AccountStreamKeyStreamSigningKeysListSigningKeysResponse]
type accountStreamKeyStreamSigningKeysListSigningKeysResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamKeyStreamSigningKeysListSigningKeysResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamKeyStreamSigningKeysListSigningKeysResponseError struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    accountStreamKeyStreamSigningKeysListSigningKeysResponseErrorJSON `json:"-"`
}

// accountStreamKeyStreamSigningKeysListSigningKeysResponseErrorJSON contains the
// JSON metadata for the struct
// [AccountStreamKeyStreamSigningKeysListSigningKeysResponseError]
type accountStreamKeyStreamSigningKeysListSigningKeysResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamKeyStreamSigningKeysListSigningKeysResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamKeyStreamSigningKeysListSigningKeysResponseMessage struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    accountStreamKeyStreamSigningKeysListSigningKeysResponseMessageJSON `json:"-"`
}

// accountStreamKeyStreamSigningKeysListSigningKeysResponseMessageJSON contains the
// JSON metadata for the struct
// [AccountStreamKeyStreamSigningKeysListSigningKeysResponseMessage]
type accountStreamKeyStreamSigningKeysListSigningKeysResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamKeyStreamSigningKeysListSigningKeysResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamKeyStreamSigningKeysListSigningKeysResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// The date and time a signing key was created.
	Created time.Time                                                          `json:"created" format:"date-time"`
	JSON    accountStreamKeyStreamSigningKeysListSigningKeysResponseResultJSON `json:"-"`
}

// accountStreamKeyStreamSigningKeysListSigningKeysResponseResultJSON contains the
// JSON metadata for the struct
// [AccountStreamKeyStreamSigningKeysListSigningKeysResponseResult]
type accountStreamKeyStreamSigningKeysListSigningKeysResponseResultJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamKeyStreamSigningKeysListSigningKeysResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountStreamKeyStreamSigningKeysListSigningKeysResponseSuccess bool

const (
	AccountStreamKeyStreamSigningKeysListSigningKeysResponseSuccessTrue AccountStreamKeyStreamSigningKeysListSigningKeysResponseSuccess = true
)
