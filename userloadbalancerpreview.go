// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// UserLoadBalancerPreviewService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewUserLoadBalancerPreviewService] method instead.
type UserLoadBalancerPreviewService struct {
	Options []option.RequestOption
}

// NewUserLoadBalancerPreviewService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserLoadBalancerPreviewService(opts ...option.RequestOption) (r *UserLoadBalancerPreviewService) {
	r = &UserLoadBalancerPreviewService{}
	r.Options = opts
	return
}

// Get the result of a previous preview operation using the provided preview_id.
func (r *UserLoadBalancerPreviewService) Get(ctx context.Context, previewID interface{}, opts ...option.RequestOption) (res *UserLoadBalancerPreviewGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserLoadBalancerPreviewGetResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/preview/%v", previewID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type UserLoadBalancerPreviewGetResponse map[string]UserLoadBalancerPreviewGetResponse

type UserLoadBalancerPreviewGetResponseEnvelope struct {
	Errors   []UserLoadBalancerPreviewGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserLoadBalancerPreviewGetResponseEnvelopeMessages `json:"messages,required"`
	// Resulting health data from a preview operation.
	Result UserLoadBalancerPreviewGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success UserLoadBalancerPreviewGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    userLoadBalancerPreviewGetResponseEnvelopeJSON    `json:"-"`
}

// userLoadBalancerPreviewGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [UserLoadBalancerPreviewGetResponseEnvelope]
type userLoadBalancerPreviewGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPreviewGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userLoadBalancerPreviewGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type UserLoadBalancerPreviewGetResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    userLoadBalancerPreviewGetResponseEnvelopeErrorsJSON `json:"-"`
}

// userLoadBalancerPreviewGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [UserLoadBalancerPreviewGetResponseEnvelopeErrors]
type userLoadBalancerPreviewGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPreviewGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userLoadBalancerPreviewGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type UserLoadBalancerPreviewGetResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    userLoadBalancerPreviewGetResponseEnvelopeMessagesJSON `json:"-"`
}

// userLoadBalancerPreviewGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [UserLoadBalancerPreviewGetResponseEnvelopeMessages]
type userLoadBalancerPreviewGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPreviewGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userLoadBalancerPreviewGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type UserLoadBalancerPreviewGetResponseEnvelopeSuccess bool

const (
	UserLoadBalancerPreviewGetResponseEnvelopeSuccessTrue UserLoadBalancerPreviewGetResponseEnvelopeSuccess = true
)
