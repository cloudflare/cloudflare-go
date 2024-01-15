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

// AccountAccessServiceTokenRotateService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountAccessServiceTokenRotateService] method instead.
type AccountAccessServiceTokenRotateService struct {
	Options []option.RequestOption
}

// NewAccountAccessServiceTokenRotateService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAccessServiceTokenRotateService(opts ...option.RequestOption) (r *AccountAccessServiceTokenRotateService) {
	r = &AccountAccessServiceTokenRotateService{}
	r.Options = opts
	return
}

// Generates a new Client Secret for a service token and revokes the old one.
func (r *AccountAccessServiceTokenRotateService) AccessServiceTokensRotateAServiceToken(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *AccountAccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/service_tokens/%s/rotate", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type AccountAccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponse struct {
	Errors   []AccountAccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseError   `json:"errors"`
	Messages []AccountAccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseMessage `json:"messages"`
	Result   AccountAccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseSuccess `json:"success"`
	JSON    accountAccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseJSON    `json:"-"`
}

// accountAccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseJSON
// contains the JSON metadata for the struct
// [AccountAccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponse]
type accountAccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseError struct {
	Code    int64                                                                                  `json:"code,required"`
	Message string                                                                                 `json:"message,required"`
	JSON    accountAccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseErrorJSON `json:"-"`
}

// accountAccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountAccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseError]
type accountAccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseMessage struct {
	Code    int64                                                                                    `json:"code,required"`
	Message string                                                                                   `json:"message,required"`
	JSON    accountAccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseMessageJSON `json:"-"`
}

// accountAccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountAccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseMessage]
type accountAccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseResult struct {
	// The ID of the service token.
	ID interface{} `json:"id"`
	// The Client ID for the service token. Access will check for this value in the
	// `CF-Access-Client-ID` request header.
	ClientID string `json:"client_id"`
	// The Client Secret for the service token. Access will check for this value in the
	// `CF-Access-Client-Secret` request header.
	ClientSecret string    `json:"client_secret"`
	CreatedAt    time.Time `json:"created_at" format:"date-time"`
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration string `json:"duration"`
	// The name of the service token.
	Name      string                                                                                  `json:"name"`
	UpdatedAt time.Time                                                                               `json:"updated_at" format:"date-time"`
	JSON      accountAccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseResultJSON `json:"-"`
}

// accountAccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseResultJSON
// contains the JSON metadata for the struct
// [AccountAccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseResult]
type accountAccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseResultJSON struct {
	ID           apijson.Field
	ClientID     apijson.Field
	ClientSecret apijson.Field
	CreatedAt    apijson.Field
	Duration     apijson.Field
	Name         apijson.Field
	UpdatedAt    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseSuccess bool

const (
	AccountAccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseSuccessTrue AccountAccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseSuccess = true
)
