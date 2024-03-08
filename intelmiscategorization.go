// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
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
func (r *IntelMiscategorizationService) New(ctx context.Context, params IntelMiscategorizationNewParams, opts ...option.RequestOption) (res *IntelMiscategorizationNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IntelMiscategorizationNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/miscategorization", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [IntelMiscategorizationNewResponseUnknown] or
// [shared.UnionString].
type IntelMiscategorizationNewResponse interface {
	ImplementsIntelMiscategorizationNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IntelMiscategorizationNewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type IntelMiscategorizationNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Content category IDs to add.
	ContentAdds param.Field[interface{}] `json:"content_adds"`
	// Content category IDs to remove.
	ContentRemoves param.Field[interface{}]                                  `json:"content_removes"`
	IndicatorType  param.Field[IntelMiscategorizationNewParamsIndicatorType] `json:"indicator_type"`
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

func (r IntelMiscategorizationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IntelMiscategorizationNewParamsIndicatorType string

const (
	IntelMiscategorizationNewParamsIndicatorTypeDomain IntelMiscategorizationNewParamsIndicatorType = "domain"
	IntelMiscategorizationNewParamsIndicatorTypeIPV4   IntelMiscategorizationNewParamsIndicatorType = "ipv4"
	IntelMiscategorizationNewParamsIndicatorTypeIPV6   IntelMiscategorizationNewParamsIndicatorType = "ipv6"
	IntelMiscategorizationNewParamsIndicatorTypeURL    IntelMiscategorizationNewParamsIndicatorType = "url"
)

type IntelMiscategorizationNewResponseEnvelope struct {
	Errors   []IntelMiscategorizationNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IntelMiscategorizationNewResponseEnvelopeMessages `json:"messages,required"`
	Result   IntelMiscategorizationNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success IntelMiscategorizationNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    intelMiscategorizationNewResponseEnvelopeJSON    `json:"-"`
}

// intelMiscategorizationNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [IntelMiscategorizationNewResponseEnvelope]
type intelMiscategorizationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelMiscategorizationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelMiscategorizationNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type IntelMiscategorizationNewResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    intelMiscategorizationNewResponseEnvelopeErrorsJSON `json:"-"`
}

// intelMiscategorizationNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [IntelMiscategorizationNewResponseEnvelopeErrors]
type intelMiscategorizationNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelMiscategorizationNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelMiscategorizationNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IntelMiscategorizationNewResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    intelMiscategorizationNewResponseEnvelopeMessagesJSON `json:"-"`
}

// intelMiscategorizationNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [IntelMiscategorizationNewResponseEnvelopeMessages]
type intelMiscategorizationNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelMiscategorizationNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelMiscategorizationNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IntelMiscategorizationNewResponseEnvelopeSuccess bool

const (
	IntelMiscategorizationNewResponseEnvelopeSuccessTrue IntelMiscategorizationNewResponseEnvelopeSuccess = true
)
