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

// AccessServiceTokenRefreshService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccessServiceTokenRefreshService] method instead.
type AccessServiceTokenRefreshService struct {
	Options []option.RequestOption
}

// NewAccessServiceTokenRefreshService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccessServiceTokenRefreshService(opts ...option.RequestOption) (r *AccessServiceTokenRefreshService) {
	r = &AccessServiceTokenRefreshService{}
	r.Options = opts
	return
}

// Refreshes the expiration of a service token.
func (r *AccessServiceTokenRefreshService) AccessServiceTokensRefreshAServiceToken(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/service_tokens/%s/refresh", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponse struct {
	Errors   []AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseError   `json:"errors"`
	Messages []AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseMessage `json:"messages"`
	Result   AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseSuccess `json:"success"`
	JSON    accessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseJSON    `json:"-"`
}

// accessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseJSON
// contains the JSON metadata for the struct
// [AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponse]
type accessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseError struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    accessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseErrorJSON `json:"-"`
}

// accessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseErrorJSON
// contains the JSON metadata for the struct
// [AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseError]
type accessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseMessage struct {
	Code    int64                                                                               `json:"code,required"`
	Message string                                                                              `json:"message,required"`
	JSON    accessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseMessageJSON `json:"-"`
}

// accessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseMessageJSON
// contains the JSON metadata for the struct
// [AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseMessage]
type accessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseResult struct {
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
	Name      string                                                                             `json:"name"`
	UpdatedAt time.Time                                                                          `json:"updated_at" format:"date-time"`
	JSON      accessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseResultJSON `json:"-"`
}

// accessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseResultJSON
// contains the JSON metadata for the struct
// [AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseResult]
type accessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseResultJSON struct {
	ID          apijson.Field
	ClientID    apijson.Field
	CreatedAt   apijson.Field
	Duration    apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseSuccess bool

const (
	AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseSuccessTrue AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseSuccess = true
)
