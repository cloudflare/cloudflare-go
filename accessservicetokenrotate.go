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
	var env AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/service_tokens/%s/rotate", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponse struct {
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
	Name      string                                                                     `json:"name"`
	UpdatedAt time.Time                                                                  `json:"updated_at" format:"date-time"`
	JSON      accessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseJSON `json:"-"`
}

// accessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseJSON
// contains the JSON metadata for the struct
// [AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponse]
type accessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseJSON struct {
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

func (r *AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseEnvelope struct {
	Errors   []AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseEnvelopeJSON    `json:"-"`
}

// accessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseEnvelope]
type accessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseEnvelopeErrors struct {
	Code    int64                                                                                    `json:"code,required"`
	Message string                                                                                   `json:"message,required"`
	JSON    accessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseEnvelopeErrorsJSON `json:"-"`
}

// accessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseEnvelopeErrors]
type accessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseEnvelopeMessages struct {
	Code    int64                                                                                      `json:"code,required"`
	Message string                                                                                     `json:"message,required"`
	JSON    accessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseEnvelopeMessagesJSON `json:"-"`
}

// accessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseEnvelopeMessages]
type accessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseEnvelopeSuccess bool

const (
	AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseEnvelopeSuccessTrue AccessServiceTokenRotateAccessServiceTokensRotateAServiceTokenResponseEnvelopeSuccess = true
)
