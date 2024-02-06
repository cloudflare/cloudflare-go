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
	var env AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/service_tokens/%s/refresh", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponse struct {
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
	Name      string                                                                       `json:"name"`
	UpdatedAt time.Time                                                                    `json:"updated_at" format:"date-time"`
	JSON      accessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseJSON `json:"-"`
}

// accessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseJSON
// contains the JSON metadata for the struct
// [AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponse]
type accessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseJSON struct {
	ID          apijson.Field
	ClientID    apijson.Field
	CreatedAt   apijson.Field
	Duration    apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseEnvelope struct {
	Errors   []AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseEnvelopeErrors   `json:"errors"`
	Messages []AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseEnvelopeMessages `json:"messages"`
	Result   AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponse                   `json:"result"`
	// Whether the API call was successful
	Success AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseEnvelopeSuccess `json:"success"`
	JSON    accessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseEnvelopeJSON    `json:"-"`
}

// accessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseEnvelope]
type accessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseEnvelopeErrors struct {
	Code    int64                                                                                      `json:"code,required"`
	Message string                                                                                     `json:"message,required"`
	JSON    accessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseEnvelopeErrorsJSON `json:"-"`
}

// accessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseEnvelopeErrors]
type accessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseEnvelopeMessages struct {
	Code    int64                                                                                        `json:"code,required"`
	Message string                                                                                       `json:"message,required"`
	JSON    accessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseEnvelopeMessagesJSON `json:"-"`
}

// accessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseEnvelopeMessages]
type accessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseEnvelopeSuccess bool

const (
	AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseEnvelopeSuccessTrue AccessServiceTokenRefreshAccessServiceTokensRefreshAServiceTokenResponseEnvelopeSuccess = true
)
