// File generated from our OpenAPI spec by Stainless.

package intel

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

// MiscategorizationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewMiscategorizationService] method
// instead.
type MiscategorizationService struct {
	Options []option.RequestOption
}

// NewMiscategorizationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMiscategorizationService(opts ...option.RequestOption) (r *MiscategorizationService) {
	r = &MiscategorizationService{}
	r.Options = opts
	return
}

// Create Miscategorization
func (r *MiscategorizationService) New(ctx context.Context, params MiscategorizationNewParams, opts ...option.RequestOption) (res *MiscategorizationNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MiscategorizationNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/miscategorization", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [intel.MiscategorizationNewResponseUnknown] or
// [shared.UnionString].
type MiscategorizationNewResponse interface {
	ImplementsIntelMiscategorizationNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MiscategorizationNewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type MiscategorizationNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Content category IDs to add.
	ContentAdds param.Field[interface{}] `json:"content_adds"`
	// Content category IDs to remove.
	ContentRemoves param.Field[interface{}]                             `json:"content_removes"`
	IndicatorType  param.Field[MiscategorizationNewParamsIndicatorType] `json:"indicator_type"`
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

func (r MiscategorizationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MiscategorizationNewParamsIndicatorType string

const (
	MiscategorizationNewParamsIndicatorTypeDomain MiscategorizationNewParamsIndicatorType = "domain"
	MiscategorizationNewParamsIndicatorTypeIPV4   MiscategorizationNewParamsIndicatorType = "ipv4"
	MiscategorizationNewParamsIndicatorTypeIPV6   MiscategorizationNewParamsIndicatorType = "ipv6"
	MiscategorizationNewParamsIndicatorTypeURL    MiscategorizationNewParamsIndicatorType = "url"
)

type MiscategorizationNewResponseEnvelope struct {
	Errors   []MiscategorizationNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MiscategorizationNewResponseEnvelopeMessages `json:"messages,required"`
	Result   MiscategorizationNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MiscategorizationNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    miscategorizationNewResponseEnvelopeJSON    `json:"-"`
}

// miscategorizationNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [MiscategorizationNewResponseEnvelope]
type miscategorizationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MiscategorizationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r miscategorizationNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MiscategorizationNewResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    miscategorizationNewResponseEnvelopeErrorsJSON `json:"-"`
}

// miscategorizationNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [MiscategorizationNewResponseEnvelopeErrors]
type miscategorizationNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MiscategorizationNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r miscategorizationNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MiscategorizationNewResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    miscategorizationNewResponseEnvelopeMessagesJSON `json:"-"`
}

// miscategorizationNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [MiscategorizationNewResponseEnvelopeMessages]
type miscategorizationNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MiscategorizationNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r miscategorizationNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MiscategorizationNewResponseEnvelopeSuccess bool

const (
	MiscategorizationNewResponseEnvelopeSuccessTrue MiscategorizationNewResponseEnvelopeSuccess = true
)
