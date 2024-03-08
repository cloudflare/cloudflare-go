// File generated from our OpenAPI spec by Stainless.

package billing

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// ProfileService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewProfileService] method instead.
type ProfileService struct {
	Options []option.RequestOption
}

// NewProfileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProfileService(opts ...option.RequestOption) (r *ProfileService) {
	r = &ProfileService{}
	r.Options = opts
	return
}

// Gets the current billing profile for the account.
func (r *ProfileService) Get(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *ProfileGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ProfileGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/billing/profile", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [billing.ProfileGetResponseUnknown] or [shared.UnionString].
type ProfileGetResponse interface {
	ImplementsBillingProfileGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProfileGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ProfileGetResponseEnvelope struct {
	Errors   []ProfileGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ProfileGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ProfileGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ProfileGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    profileGetResponseEnvelopeJSON    `json:"-"`
}

// profileGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ProfileGetResponseEnvelope]
type profileGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ProfileGetResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    profileGetResponseEnvelopeErrorsJSON `json:"-"`
}

// profileGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ProfileGetResponseEnvelopeErrors]
type profileGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ProfileGetResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    profileGetResponseEnvelopeMessagesJSON `json:"-"`
}

// profileGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [ProfileGetResponseEnvelopeMessages]
type profileGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProfileGetResponseEnvelopeSuccess bool

const (
	ProfileGetResponseEnvelopeSuccessTrue ProfileGetResponseEnvelopeSuccess = true
)
