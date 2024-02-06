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

// AccessServiceTokenRotateService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccessServiceTokenRotateService] method instead.
type AccessServiceTokenRotateService struct {
	Options []option.RequestOption
}

// NewAccessServiceTokenRotateService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccessServiceTokenRotateService(opts ...option.RequestOption) (r *AccessServiceTokenRotateService) {
	r = &AccessServiceTokenRotateService{}
	r.Options = opts
	return
}

// Generates a new Client Secret for a service token and revokes the old one.
func (r *AccessServiceTokenRotateService) AccessServiceTokensRotateAServiceToken(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/service_tokens/%s/rotate", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponse struct {
	Errors   []AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseError   `json:"errors"`
	Messages []AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseMessage `json:"messages"`
	Result   AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseSuccess `json:"success"`
	JSON    accessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseJSON    `json:"-"`
}

// accessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseJSON
// contains the JSON metadata for the struct
// [AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponse]
type accessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseError struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    accessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseErrorJSON `json:"-"`
}

// accessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseErrorJSON
// contains the JSON metadata for the struct
// [AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseError]
type accessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseMessage struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    accessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseMessageJSON `json:"-"`
}

// accessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseMessageJSON
// contains the JSON metadata for the struct
// [AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseMessage]
type accessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseResult struct {
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
	Name      string                                                                           `json:"name"`
	UpdatedAt time.Time                                                                        `json:"updated_at" format:"date-time"`
	JSON      accessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseResultJSON `json:"-"`
}

// accessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseResultJSON
// contains the JSON metadata for the struct
// [AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseResult]
type accessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseResultJSON struct {
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

func (r *AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseSuccess bool

const (
	AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseSuccessTrue AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseSuccess = true
)
