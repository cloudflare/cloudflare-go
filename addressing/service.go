// File generated from our OpenAPI spec by Stainless.

package addressing

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// ServiceService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewServiceService] method instead.
type ServiceService struct {
	Options []option.RequestOption
}

// NewServiceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewServiceService(opts ...option.RequestOption) (r *ServiceService) {
	r = &ServiceService{}
	r.Options = opts
	return
}

// Bring-Your-Own IP (BYOIP) prefixes onboarded to Cloudflare must be bound to a
// service running on the Cloudflare network to enable a Cloudflare product on the
// IP addresses. This endpoint can be used as a reference of available services on
// the Cloudflare network, and their service IDs.
func (r *ServiceService) List(ctx context.Context, query ServiceListParams, opts ...option.RequestOption) (res *[]ServiceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ServiceListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/services", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ServiceListResponse struct {
	// Identifier
	ID string `json:"id"`
	// Name of a service running on the Cloudflare network
	Name string                  `json:"name"`
	JSON serviceListResponseJSON `json:"-"`
}

// serviceListResponseJSON contains the JSON metadata for the struct
// [ServiceListResponse]
type serviceListResponseJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceListResponseJSON) RawJSON() string {
	return r.raw
}

type ServiceListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ServiceListResponseEnvelope struct {
	Errors   []ServiceListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ServiceListResponseEnvelopeMessages `json:"messages,required"`
	Result   []ServiceListResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success ServiceListResponseEnvelopeSuccess `json:"success,required"`
	JSON    serviceListResponseEnvelopeJSON    `json:"-"`
}

// serviceListResponseEnvelopeJSON contains the JSON metadata for the struct
// [ServiceListResponseEnvelope]
type serviceListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ServiceListResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    serviceListResponseEnvelopeErrorsJSON `json:"-"`
}

// serviceListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ServiceListResponseEnvelopeErrors]
type serviceListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ServiceListResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    serviceListResponseEnvelopeMessagesJSON `json:"-"`
}

// serviceListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ServiceListResponseEnvelopeMessages]
type serviceListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ServiceListResponseEnvelopeSuccess bool

const (
	ServiceListResponseEnvelopeSuccessTrue ServiceListResponseEnvelopeSuccess = true
)
