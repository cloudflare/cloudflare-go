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

// SSLRecommendationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSSLRecommendationService] method
// instead.
type SSLRecommendationService struct {
	Options []option.RequestOption
}

// NewSSLRecommendationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSSLRecommendationService(opts ...option.RequestOption) (r *SSLRecommendationService) {
	r = &SSLRecommendationService{}
	r.Options = opts
	return
}

// Retrieve the SSL/TLS Recommender's recommendation for a zone.
func (r *SSLRecommendationService) Get(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *SSLRecommendationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SSLRecommendationGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/ssl/recommendation", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SSLRecommendationGetResponse struct {
	// Identifier of a recommedation result.
	ID         string                            `json:"id"`
	ModifiedOn time.Time                         `json:"modified_on" format:"date-time"`
	Value      SSLRecommendationGetResponseValue `json:"value"`
	JSON       sslRecommendationGetResponseJSON  `json:"-"`
}

// sslRecommendationGetResponseJSON contains the JSON metadata for the struct
// [SSLRecommendationGetResponse]
type sslRecommendationGetResponseJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLRecommendationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLRecommendationGetResponseValue string

const (
	SSLRecommendationGetResponseValueFlexible SSLRecommendationGetResponseValue = "flexible"
	SSLRecommendationGetResponseValueFull     SSLRecommendationGetResponseValue = "full"
	SSLRecommendationGetResponseValueStrict   SSLRecommendationGetResponseValue = "strict"
)

type SSLRecommendationGetResponseEnvelope struct {
	Errors   []SSLRecommendationGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SSLRecommendationGetResponseEnvelopeMessages `json:"messages,required"`
	Result   SSLRecommendationGetResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success SSLRecommendationGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    sslRecommendationGetResponseEnvelopeJSON    `json:"-"`
}

// sslRecommendationGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SSLRecommendationGetResponseEnvelope]
type sslRecommendationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLRecommendationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLRecommendationGetResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    sslRecommendationGetResponseEnvelopeErrorsJSON `json:"-"`
}

// sslRecommendationGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SSLRecommendationGetResponseEnvelopeErrors]
type sslRecommendationGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLRecommendationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SSLRecommendationGetResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    sslRecommendationGetResponseEnvelopeMessagesJSON `json:"-"`
}

// sslRecommendationGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SSLRecommendationGetResponseEnvelopeMessages]
type sslRecommendationGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLRecommendationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SSLRecommendationGetResponseEnvelopeSuccess bool

const (
	SSLRecommendationGetResponseEnvelopeSuccessTrue SSLRecommendationGetResponseEnvelopeSuccess = true
)
