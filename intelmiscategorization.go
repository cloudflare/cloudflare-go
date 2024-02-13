// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// IntelMiscategorizationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewIntelMiscategorizationService]
// method instead.
type IntelMiscategorizationService struct {
	Options []option.RequestOption
}

// NewIntelMiscategorizationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewIntelMiscategorizationService(opts ...option.RequestOption) (r *IntelMiscategorizationService) {
	r = &IntelMiscategorizationService{}
	r.Options = opts
	return
}

// Create Miscategorization
func (r *IntelMiscategorizationService) MiscategorizationNewMiscategorization(ctx context.Context, accountID string, body IntelMiscategorizationMiscategorizationNewMiscategorizationParams, opts ...option.RequestOption) (res *IntelMiscategorizationMiscategorizationNewMiscategorizationResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IntelMiscategorizationMiscategorizationNewMiscategorizationResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/miscategorization", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by
// [IntelMiscategorizationMiscategorizationNewMiscategorizationResponseUnknown] or
// [shared.UnionString].
type IntelMiscategorizationMiscategorizationNewMiscategorizationResponse interface {
	ImplementsIntelMiscategorizationMiscategorizationNewMiscategorizationResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IntelMiscategorizationMiscategorizationNewMiscategorizationResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type IntelMiscategorizationMiscategorizationNewMiscategorizationParams struct {
	// Content category IDs to add.
	ContentAdds param.Field[interface{}] `json:"content_adds"`
	// Content category IDs to remove.
	ContentRemoves param.Field[interface{}]                                                                    `json:"content_removes"`
	IndicatorType  param.Field[IntelMiscategorizationMiscategorizationNewMiscategorizationParamsIndicatorType] `json:"indicator_type"`
	// Provide only if indicator_type is `ipv4` or `ipv6`.
	IP param.Field[interface{}] `json:"ip"`
	// Security category IDs to add.
	SecurityAdds param.Field[interface{}] `json:"security_adds"`
	// Security category IDs to remove.
	SecurityRemoves param.Field[interface{}] `json:"security_removes"`
	// Provide only if indicator_type is `domain` or `url`. Example if indicator_type
	// is `domain`: `example.com`. Example if indicator_type is `url`:
	// `https://example.com/news/`.
	URL param.Field[string] `json:"url"`
}

func (r IntelMiscategorizationMiscategorizationNewMiscategorizationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IntelMiscategorizationMiscategorizationNewMiscategorizationParamsIndicatorType string

const (
	IntelMiscategorizationMiscategorizationNewMiscategorizationParamsIndicatorTypeDomain IntelMiscategorizationMiscategorizationNewMiscategorizationParamsIndicatorType = "domain"
	IntelMiscategorizationMiscategorizationNewMiscategorizationParamsIndicatorTypeIPV4   IntelMiscategorizationMiscategorizationNewMiscategorizationParamsIndicatorType = "ipv4"
	IntelMiscategorizationMiscategorizationNewMiscategorizationParamsIndicatorTypeIPV6   IntelMiscategorizationMiscategorizationNewMiscategorizationParamsIndicatorType = "ipv6"
	IntelMiscategorizationMiscategorizationNewMiscategorizationParamsIndicatorTypeURL    IntelMiscategorizationMiscategorizationNewMiscategorizationParamsIndicatorType = "url"
)

type IntelMiscategorizationMiscategorizationNewMiscategorizationResponseEnvelope struct {
	Errors   []IntelMiscategorizationMiscategorizationNewMiscategorizationResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IntelMiscategorizationMiscategorizationNewMiscategorizationResponseEnvelopeMessages `json:"messages,required"`
	Result   IntelMiscategorizationMiscategorizationNewMiscategorizationResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success IntelMiscategorizationMiscategorizationNewMiscategorizationResponseEnvelopeSuccess `json:"success,required"`
	JSON    intelMiscategorizationMiscategorizationNewMiscategorizationResponseEnvelopeJSON    `json:"-"`
}

// intelMiscategorizationMiscategorizationNewMiscategorizationResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [IntelMiscategorizationMiscategorizationNewMiscategorizationResponseEnvelope]
type intelMiscategorizationMiscategorizationNewMiscategorizationResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelMiscategorizationMiscategorizationNewMiscategorizationResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelMiscategorizationMiscategorizationNewMiscategorizationResponseEnvelopeErrors struct {
	Code    int64                                                                                 `json:"code,required"`
	Message string                                                                                `json:"message,required"`
	JSON    intelMiscategorizationMiscategorizationNewMiscategorizationResponseEnvelopeErrorsJSON `json:"-"`
}

// intelMiscategorizationMiscategorizationNewMiscategorizationResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [IntelMiscategorizationMiscategorizationNewMiscategorizationResponseEnvelopeErrors]
type intelMiscategorizationMiscategorizationNewMiscategorizationResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelMiscategorizationMiscategorizationNewMiscategorizationResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelMiscategorizationMiscategorizationNewMiscategorizationResponseEnvelopeMessages struct {
	Code    int64                                                                                   `json:"code,required"`
	Message string                                                                                  `json:"message,required"`
	JSON    intelMiscategorizationMiscategorizationNewMiscategorizationResponseEnvelopeMessagesJSON `json:"-"`
}

// intelMiscategorizationMiscategorizationNewMiscategorizationResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [IntelMiscategorizationMiscategorizationNewMiscategorizationResponseEnvelopeMessages]
type intelMiscategorizationMiscategorizationNewMiscategorizationResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelMiscategorizationMiscategorizationNewMiscategorizationResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type IntelMiscategorizationMiscategorizationNewMiscategorizationResponseEnvelopeSuccess bool

const (
	IntelMiscategorizationMiscategorizationNewMiscategorizationResponseEnvelopeSuccessTrue IntelMiscategorizationMiscategorizationNewMiscategorizationResponseEnvelopeSuccess = true
)
