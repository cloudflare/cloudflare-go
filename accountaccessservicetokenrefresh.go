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

// AccountAccessServiceTokenRefreshService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountAccessServiceTokenRefreshService] method instead.
type AccountAccessServiceTokenRefreshService struct {
	Options []option.RequestOption
}

// NewAccountAccessServiceTokenRefreshService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAccessServiceTokenRefreshService(opts ...option.RequestOption) (r *AccountAccessServiceTokenRefreshService) {
	r = &AccountAccessServiceTokenRefreshService{}
	r.Options = opts
	return
}

// Refreshes the expiration of a service token.
func (r *AccountAccessServiceTokenRefreshService) AccessServiceTokensRefreshAServiceToken(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *AccountAccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/service_tokens/%s/refresh", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type AccountAccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponse struct {
	Errors   []AccountAccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseError   `json:"errors"`
	Messages []AccountAccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseMessage `json:"messages"`
	Result   AccountAccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseSuccess `json:"success"`
	JSON    accountAccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseJSON    `json:"-"`
}

// accountAccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseJSON
// contains the JSON metadata for the struct
// [AccountAccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponse]
type accountAccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseError struct {
	Code    int64                                                                                    `json:"code,required"`
	Message string                                                                                   `json:"message,required"`
	JSON    accountAccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseErrorJSON `json:"-"`
}

// accountAccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountAccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseError]
type accountAccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseMessage struct {
	Code    int64                                                                                      `json:"code,required"`
	Message string                                                                                     `json:"message,required"`
	JSON    accountAccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseMessageJSON `json:"-"`
}

// accountAccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountAccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseMessage]
type accountAccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseResult struct {
	// The ID of the service token.
	ID interface{} `json:"id"`
	// The Client ID for the service token. Access will check for this value in the
	// `CF-Access-Client-ID` request header.
	ClientID  string    `json:"client_id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration string `json:"duration"`
	// The name of the service token.
	Name      string                                                                                    `json:"name"`
	UpdatedAt time.Time                                                                                 `json:"updated_at" format:"date-time"`
	JSON      accountAccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseResultJSON `json:"-"`
}

// accountAccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseResultJSON
// contains the JSON metadata for the struct
// [AccountAccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseResult]
type accountAccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseResultJSON struct {
	ID          apijson.Field
	ClientID    apijson.Field
	CreatedAt   apijson.Field
	Duration    apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseSuccess bool

const (
	AccountAccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseSuccessTrue AccountAccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseSuccess = true
)
